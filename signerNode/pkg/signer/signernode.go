package signer

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"net"
	"signer/pkg/kyber/pairing/bn256"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	paillier "github.com/lyszhang/go-go-gadget-paillier"
	"github.com/segmentio/kafka-go"
	"go.dedis.ch/kyber/v3"
	"go.dedis.ch/kyber/v3/pairing"
	"google.golang.org/grpc"
)

type OracleNode struct {
	UnsafeSignerServer
	server    *grpc.Server
	serverLis net.Listener
	EthClient *ethclient.Client
	Contract  *ContractWrapper
	suite     pairing.Suite

	ecdsaPrivateKey *ecdsa.PrivateKey
	account         common.Address
	chainId         *big.Int
	signerNode      *Signer // 执行签名方案的节点

}

func NewOracleNode(c Config) (*OracleNode, error) {
	server := grpc.NewServer()
	serverLis, err := net.Listen("tcp", c.BindAddress)
	if err != nil {
		return nil, fmt.Errorf("listen on %s: %v", c.BindAddress, err)
	}
	// 创建一个连接以太坊的客户端，TargetAddress是以太坊的目标地址
	EthClient, err := ethclient.Dial(c.Ethereum.Address)
	if err != nil {
		return nil, fmt.Errorf("dial eth client: %v", err)
	}

	// 区块链的ID
	chainId := big.NewInt(c.Ethereum.ChainID)

	Registry, err := NewRegistryContract(common.HexToAddress(c.Contracts.RegistryContractAddress), EthClient)
	if err != nil {
		return nil, fmt.Errorf("Registry contract: %v", err)
	}

	SingerContract, err := NewSingerContract(common.HexToAddress(c.Contracts.SingerContractAddress), EthClient)

	if err != nil {
		return nil, fmt.Errorf("Singer contract: %v", err)
	}

	Contract := &ContractWrapper{
		RegistryContract: Registry,
		SingerContract:   SingerContract,
	}
	suite := bn256.NewSuite()

	ecdsaPrivateKey, err := crypto.HexToECDSA(c.Ethereum.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("hex to ecdsa: %v", err)
	}

	hexAddress, err := AddressFromPrivateKey(ecdsaPrivateKey)

	if err != nil {
		fmt.Errorf("failed to generate private key: %v", err)
	}

	account := common.HexToAddress(hexAddress)

	privateKey := suite.G1().Scalar().Pick(suite.RandomStream())
	if err != nil {
		fmt.Errorf("failed to generate private key: %v", err)
	}
	RAll := make(map[common.Address]kyber.Point)   // SM2
	MtAll := make(map[common.Address]kyber.Scalar) // ECDSA

	// 初始化kafka Writer 和 Reader
	writer := &kafka.Writer{
		Addr:                   kafka.TCP(c.Kafka.IpAddress),
		Topic:                  c.Kafka.Topic,
		RequiredAcks:           kafka.RequireAll,
		Balancer:               &kafka.LeastBytes{},
		AllowAutoTopicCreation: true,
		Async:                  true,
	}

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{c.Kafka.IpAddress},
		Topic:     c.Kafka.Topic,
		Partition: int(c.Kafka.Partition),
		MaxBytes:  10e6, // 10MB
	})

	Signer := NewSigner(
		suite,
		Contract,
		ecdsaPrivateKey,
		EthClient,
		account,
		privateKey,
		chainId,
		writer,
		reader,
		RAll,
		MtAll,
		false,
	)

	node := &OracleNode{
		server:          server,
		serverLis:       serverLis,
		EthClient:       EthClient,
		Contract:        Contract,
		suite:           suite,
		ecdsaPrivateKey: ecdsaPrivateKey,
		account:         account,
		chainId:         chainId,
		signerNode:      Signer,
	}

	RegisterSignerServer(server, node)

	return node, nil
}

func (n *OracleNode) Run() error {

	go func() {
		if err := n.signerNode.ListenAndProcess(n); err != nil {
			fmt.Errorf("Watch and handle R log: %v", err)
		}
	}()

	go func() {
		if err := n.WatchAndHandleValidationRequestsLog(context.Background()); err != nil {
			log.Fatal("Watch and handle SigatureRequest log: %w", err)
		}
	}()

	go func() {
		if err := n.signerNode.WatchAndHandleSignLog(context.Background()); err != nil {
			log.Fatal("Watch and handle SigatureRequest log: %w", err)
		}
	}()

	go func() {
		if err := n.signerNode.WatchAndHandleMulVerifyLog(context.Background()); err != nil {
			log.Fatal("Watch and handle SigatureRequest log: %w", err)
		}
	}()

	if err := n.register(n.serverLis.Addr().String()); err != nil {
		return fmt.Errorf("register: %w", err)
	}

	return n.server.Serve(n.serverLis)
}

func (n *OracleNode) register(ipAddr string) error {

	auth, err := bind.NewKeyedTransactorWithChainID(n.ecdsaPrivateKey, n.chainId)
	if err != nil {
		log.Println("NewKeyedTransactorWithChainID :", err)
	}
	log.Println(n.signerNode.privateKey)
	publicKey := n.suite.G1().Point().Mul(n.signerNode.privateKey, nil)
	publicKeyBig, err := G1PointToBig(publicKey)
	if err != nil {
		log.Println("translate idPk to Big : ", err)
	}
	// 构造paillire 同态加密的密钥
	paillierKeys, err := paillier.GenerateKey(rand.Reader, 1024)
	n.signerNode.paillierKey = paillierKeys
	paillierPubKeyString := paillierKeys.PublicKey.String()
	if err != nil {
		log.Fatalf("Error in paillierPubKey Json: %v", err)
	}
	_, err = n.Contract.RegistryContract.SignerRegister(auth, ipAddr, publicKeyBig, paillierPubKeyString)
	if err != nil {
		return fmt.Errorf("register node: %w", err)

	}

	return nil
}

func (n *OracleNode) WatchAndHandleValidationRequestsLog(ctx context.Context) error {
	sink := make(chan *RegistryContractSign)
	defer close(sink)

	sub, err := n.Contract.WatchSign(
		&bind.WatchOpts{
			Context: context.Background(),
		},
		sink,
	)
	if err != nil {
		return err
	}
	defer sub.Unsubscribe()

	for {
		select {
		case event := <-sink:
			log.Println("Received SignatureRequest event with Message ", event.Message)

			if err != nil {
				log.Println("Is aggregator: ", err)
				continue
			}

			// 报名函数
			auth, err := bind.NewKeyedTransactorWithChainID(n.ecdsaPrivateKey, n.chainId)
			usePublicKey, err := n.Contract.GetUsePublicKeyCount(nil)
			if err != nil {
				log.Println("get publickey has err: ", err)
				return err
			}
			log.Println(usePublicKey)
			if usePublicKey.Cmp(big.NewInt(0)) == 0 {

				_, err := n.Contract.EnRoll(auth)

				if err != nil {
					log.Println("Node Enroll err: ", err)
					return err
				}

			}

			n.signerNode.message = event.Message
			n.signerNode.RAll = make(map[common.Address]kyber.Point)
			n.signerNode.MtAll = make(map[common.Address]kyber.Scalar)

			n.signerNode.Sign(event.Message)
			continue
		case err = <-sub.Err():
			return err
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (n *OracleNode) Stop() {
	n.server.Stop()

	n.EthClient.Close()
}
