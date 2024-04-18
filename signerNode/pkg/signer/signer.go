package signer

import (
	"context"

	"crypto/ecdsa"
	"crypto/sha256"
	"errors"
	"fmt"
	"log"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	paillier "github.com/lyszhang/go-go-gadget-paillier"
	"github.com/segmentio/kafka-go"
	"go.dedis.ch/kyber/v3"
	"go.dedis.ch/kyber/v3/pairing"
	"google.golang.org/grpc"
)

type Signer struct {
	sync.RWMutex
	suite           pairing.Suite
	Contract        *ContractWrapper
	ecdsaPrivateKey *ecdsa.PrivateKey
	ethClient       *ethclient.Client
	account         common.Address
	chainId         *big.Int
	privateKey      kyber.Scalar
	paillierKey     *paillier.PrivateKey

	kafkaWriter *kafka.Writer
	kafkaReader *kafka.Reader

	sigPrivateKey kyber.Scalar
	MtAResultAij  []kyber.Scalar // Pi->Pj ki的时候，Pj拿到的aij
	MtAResultBji  []kyber.Scalar // Pi->Pj ki的时候，Pi拿到的bji

	RAll                         map[common.Address]kyber.Point
	message                      []byte
	MtAll                        map[common.Address]kyber.Scalar
	enRollSuccess                bool
	singleSignatureVerifySuccess bool
}

func NewSigner(
	suite pairing.Suite,
	Contract *ContractWrapper,
	ecdsaPrivateKey *ecdsa.PrivateKey,
	ethClient *ethclient.Client,
	account common.Address,
	privateKey kyber.Scalar,
	chainId *big.Int,
	kafkaWriter *kafka.Writer,
	kafkaReader *kafka.Reader,
	RAll map[common.Address]kyber.Point,
	MtAll map[common.Address]kyber.Scalar,
	enRollSuccess bool,

) *Signer {
	return &Signer{
		suite:                        suite,
		Contract:                     Contract,
		ecdsaPrivateKey:              ecdsaPrivateKey,
		ethClient:                    ethClient,
		account:                      account,
		chainId:                      chainId,
		kafkaWriter:                  kafkaWriter,
		kafkaReader:                  kafkaReader,
		RAll:                         RAll,
		MtAll:                        MtAll,
		privateKey:                   privateKey,
		enRollSuccess:                enRollSuccess,
		singleSignatureVerifySuccess: false,
	}
}

func (s *Signer) WatchAndHandleSignLog(ctx context.Context) error {
	sink := make(chan *SingerContractSignatureBegin)
	defer close(sink)

	sub, err := s.Contract.SingerContract.WatchSignatureBegin(
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
			auth, _ := bind.NewKeyedTransactorWithChainID(s.ecdsaPrivateKey, s.chainId)

			s.Contract.AggKey(auth)
			s.enRollSuccess = true
			_ = event
		case err = <-sub.Err():
			return err
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (s *Signer) WatchAndHandleMulVerifyLog(ctx context.Context) error {
	sink := make(chan *SingerContractVerifyMulSignature)
	defer close(sink)

	sub, err := s.Contract.SingerContract.WatchVerifyMulSignature(
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
			auth, _ := bind.NewKeyedTransactorWithChainID(s.ecdsaPrivateKey, s.chainId)
			log.Println(s.message)
			_, err := s.Contract.VerifyMul(auth, s.message)

			if err != nil {
				log.Println("verifyMul has err", err)
			}
			_ = event
		case err = <-sub.Err():
			return err
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

// ECDSA
func (s *Signer) Sign(message []byte) {

	for !s.enRollSuccess {

	}
	// 先产生自己的R，然后在等待一段时间，随后广播, 构造R序列
	count, _ := s.Contract.CountEnrollNodes(nil)
	hash := sha256.New()
	ki := s.suite.G1().Scalar().Pick(s.suite.RandomStream())

	kiInv := s.suite.G1().Scalar().Inv(ki)
	kiInvG := s.suite.G1().Point().Mul(kiInv, nil)
	kiInvGBig, _ := G1PointToBig(kiInvG)

	// ECDSA中的ri就用xi代替
	start := time.Now()
	s.sigPrivateKey = s.suite.G1().Scalar().Pick(s.suite.RandomStream())

	tmpPoint := s.suite.G1().Point().Mul(s.sigPrivateKey, nil)
	log.Println(time.Since(start).Microseconds())

	tmpPointBytes, _ := tmpPoint.MarshalBinary()

	s.sendR(tmpPointBytes)

	timeout := time.After(Timeout)
	nodes, _ := s.Contract.FindEnrollSingers()

loop:
	for {
		select {
		case <-timeout:
			fmt.Errorf("Timeout")
			break loop
		default:

			if len(nodes) == len(s.RAll) {
				break loop
			}
			// time.Sleep(50 * time.Millisecond)
		}
	}

	R := s.suite.G1().Point().Null()
	start = time.Now()
	for key := range s.RAll {
		R = s.suite.G1().Point().Add(R, s.RAll[key])
	}
	log.Println(time.Since(start).Microseconds())
	kByte, err := ki.MarshalBinary()
	s.MtAStart(kByte)

	for len(s.MtAResultAij) < int(count.Int64())-1 || len(s.MtAResultBji) < int(count.Int64())-1 { // 因为有一个节点作为聚合器，不参与整个过程

	}
	log.Println("MtA end")
	start = time.Now()
	tmp := s.suite.G1().Scalar().Mul(ki, s.sigPrivateKey)
	for _, a := range s.MtAResultAij {
		tmp = s.suite.G1().Scalar().Add(tmp, a)
	}

	for _, b := range s.MtAResultBji {
		tmp = s.suite.G1().Scalar().Add(tmp, b)
	}
	log.Println(time.Since(start).Microseconds())
	tmpBytes, _ := tmp.MarshalBinary()

	s.sendR(tmpBytes)

loop1:
	for {
		select {
		case <-timeout:
			fmt.Errorf("Timeout")
			break loop1
		default:
			if len(nodes) == len(s.MtAll) {
				break loop1
			}
			// time.Sleep(50 * time.Millisecond)
		}
	}
	// 此时需要获取到其他人的部分份额
	log.Println("kafka broadcast tmp complete")
	start = time.Now()
	delt := s.suite.G1().Scalar().Zero()
	for key := range s.MtAll {
		delt = s.suite.G1().Scalar().Add(delt, s.MtAll[key])
	}

	delt = s.suite.G1().Scalar().Inv(delt)
	R = s.suite.G1().Point().Mul(delt, R)
	log.Println(time.Since(start).Microseconds())

	s.MtAll = make(map[common.Address]kyber.Scalar)
	Rbig, err := G1PointToBig(R)
	if err != nil {
		log.Println("R translate big has err, ", err)
	}
	start = time.Now()
	hash.Reset()
	hash.Write(Rbig[0].Bytes())
	log.Println(Rbig)
	r := s.suite.G1().Scalar().SetBytes(hash.Sum(nil))
	log.Println(time.Since(start).Microseconds())

	hash.Reset()
	currentPk := s.suite.G1().Point().Mul(s.privateKey, nil)
	currentPkBig, err := G1PointToBig(currentPk)
	if err != nil {
		log.Println("get current PK has err", err)
	}

	J, err := s.Contract.GetEnrollPKs(nil)
	if err != nil {
		log.Println("get J has err", err)
	}

	PK, err := s.Contract.GetAllPKs(nil)
	if err != nil {
		log.Println("get PkByte has err", err)
	}
	start = time.Now()
	hash.Write(currentPkBig[0].Bytes())
	hash.Write(J[0].Bytes())
	hash.Write(PK[0].Bytes())

	a := s.suite.G1().Scalar().SetBytes(hash.Sum(nil))
	s.sigPrivateKey = s.suite.G1().Scalar().Mul(a, s.privateKey)
	log.Println(time.Since(start).Microseconds())

	// 开始第二轮MtA
	s.MtAResultAij = make([]kyber.Scalar, 0)
	s.MtAResultBji = make([]kyber.Scalar, 0)
	s.MtAStart(kByte)

	gtag := groupSet(PK)
	message = append(message, gtag...)
	s.message = message
	hash.Reset()

	start = time.Now()
	hash.Write(message)

	H_m := s.suite.G1().Scalar().SetBytes(hash.Sum(nil))
	singleSignature := s.suite.G1().Scalar().Add(s.suite.G1().Scalar().Mul(ki, H_m), s.suite.G1().Scalar().Mul(s.suite.G1().Scalar().Mul(ki, s.sigPrivateKey), r))
	log.Println(time.Since(start).Microseconds())

	singleSignatureBig, err := ScalarToBig(singleSignature)
	rBig, _ := ScalarToBig(r)
	if err != nil {
		log.Println("get signature big has err", err)

	}

	for len(s.MtAResultAij) < int(count.Int64())-1 || len(s.MtAResultBji) < int(count.Int64())-1 { // 因为有一个节点作为聚合器，不参与整个过程
	}
	log.Println("MtA end")

	start = time.Now()
	sigma := s.suite.G1().Scalar().Zero()
	for _, a := range s.MtAResultAij {
		sigma = s.suite.G1().Scalar().Add(sigma, a)
	}

	for _, b := range s.MtAResultBji {
		sigma = s.suite.G1().Scalar().Add(sigma, b)
	}
	sigma = s.suite.G1().Scalar().Mul(sigma, r)

	// 	si := s.suite.G1().Scalar().Add(sigma, singleSignature)
	// 	siByte, _ := si.MarshalBinary()

	// 	s.sendR(siByte)

	// loop2:
	// 	for {
	// 		select {
	// 		case <-timeout:
	// 			fmt.Errorf("Timeout")
	// 			break loop2
	// 		default:
	// 			if len(nodes) == len(s.MtAll) {
	// 				break loop2
	// 			}
	// 			// time.Sleep(50 * time.Millisecond)
	// 		}
	// 	}
	// 	// 此时需要获取到其他人的部分份额
	// 	log.Println("kafka broadcast si complete")
	// 	start = time.Now()
	// 	sig := s.suite.G1().Scalar().Zero()
	// 	for key := range s.MtAll {
	// 		sig = s.suite.G1().Scalar().Add(sig, s.MtAll[key])
	// 	}
	// 	log.Println(sig)
	// 	s.verifyECDSA(sig, r)
	// 	log.Println(time.Since(start).Microseconds())

	sigmaBig, _ := ScalarToBig(sigma)
	auth, err := bind.NewKeyedTransactorWithChainID(s.ecdsaPrivateKey, s.chainId)
	log.Println(message)
	_, err = s.Contract.SubmitECDSA(auth, kiInvGBig[0], rBig, singleSignatureBig, sigmaBig, message)
	if err != nil {
		log.Println("submit singleSignature has err", err)
	} else {
		log.Println("submit singSignature success")
	}

}

func (s *Signer) MtAStart(k []byte) {

	encrypt_k, err := paillier.Encrypt(&s.paillierKey.PublicKey, k)

	signer, err := s.Contract.GetSignerByAddress(nil, s.account)

	if err != nil {
		log.Println("get current signerNode has err, ", err)
	}

	nodes, err := s.Contract.FindEnrollSingers()
	if err != nil {
		log.Println("get All nodes has err", err)
	}

	var wg sync.WaitGroup
	var mutex sync.Mutex

	for _, n := range nodes {
		if s.account.Cmp(n) == 0 {
			continue
		}
		node, _ := s.Contract.GetSignerByAddress(nil, n)
		conn, err := grpc.Dial(node.IpAddr, grpc.WithInsecure())
		if err != nil {
			log.Println("Find connection by address: ", err)
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			defer mutex.Unlock()
			client := NewSignerClient(conn)
			ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

			result, err := client.MtA(ctx, &MtARequest{K: encrypt_k, Index: signer.Index.Int64()})

			cancel()
			if err != nil {
				log.Println("get MtA result has err :", err)
				return
			}

			mutex.Lock()
			bByte, err := paillier.Decrypt(s.paillierKey, result.B)
			if err != nil {
				log.Println("decrypt b has err", err)
			}
			// if err != nil {
			// 	log.Println("decrypt b has err", err, len(bByte), len(result.B))
			// } else {
			// 	log.Println(len(bByte), len(result.B))
			// }
			b := s.suite.G1().Scalar().SetBytes(bByte)

			s.MtAResultBji = append(s.MtAResultBji, b)

		}()
	}

	wg.Wait()

}

func (s *Signer) ListenAndProcess(o *OracleNode) error {

	for {
		m, err := s.kafkaReader.ReadMessage(context.Background())
		if err != nil {
			break
		}
		// 处理kafka消息
		// 长度是32， 为大整数
		if len(m.Value) == 32 {
			tmp := s.suite.G1().Scalar().SetBytes(m.Value)
			s.MtAll[common.Address(m.Key)] = tmp
		} else if len(m.Value) == 64 {
			RPoint := s.suite.G1().Point()
			err = RPoint.UnmarshalBinary(m.Value)
			if err != nil {
				log.Println("R transform to Point: ", err)

			} else {
				s.RAll[common.Address(m.Key)] = RPoint
			}
		} else {
			log.Println(len(m.Value), m.Value)
			log.Println("message has error")
		}

	}
	return nil
}

func (s *Signer) sendR(R []byte) {
	messages := []kafka.Message{
		{
			Key:   []byte(s.account.String()),
			Value: R,
		},
	}
	var err error
	const retries = 3
	// 重试3次

	for i := 0; i < retries; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		err = s.kafkaWriter.WriteMessages(ctx, messages...)
		if errors.Is(err, kafka.LeaderNotAvailable) || errors.Is(err, context.DeadlineExceeded) {
			time.Sleep(time.Millisecond * 100)
			continue
		}
		if err != nil {
			log.Fatalf("unexpected error %v", err)
		}
		break
	}
}

func (s *Signer) MtA(encrypt_k []byte, index int64) []byte {

	// 接收到加密后的数据，然后进行同态处理
	// 获取交互的节点
	node, err := s.Contract.FindSignerByIndex(nil, big.NewInt(index))
	if err != nil {
		log.Println("get node has err", err)
	}
	paillierPubKey, err := paillier.NewPubkeyFromString(node.PaillierPubkey)
	if err != nil {
		log.Println("get paillierPubKey has err", err)
	}
	xByte, _ := s.sigPrivateKey.MarshalBinary()

	kMulx := paillier.Mul(paillierPubKey, encrypt_k, xByte)

	// 随机选取一个a
	a := s.suite.G1().Scalar().Pick(s.suite.RandomStream())

	s.MtAResultAij = append(s.MtAResultAij, a)
	aByte, err := a.MarshalBinary()
	encrypt_a, err := paillier.Encrypt(paillierPubKey, aByte)

	if err != nil {
		log.Println("encrypt a has err", err)
	}
	encrypt_b := paillier.SubCipher(paillierPubKey, kMulx, encrypt_a)
	return encrypt_b
}

// func (s *Signer) verifyECDSA(si kyber.Scalar, r kyber.Scalar) {
// 	nodes, _ := s.Contract.FindEnrollSingers()
// 	MulY := s.suite.G1().Point().Null()
// 	hash := sha256.New()
// 	J, _ := s.Contract.GetEnrollPKs(nil)

// 	PK, _ := s.Contract.GetAllPKs(nil)

// 	for _, n := range nodes {
// 		currentPK, _ := s.Contract.GetSignerPubkeyByAddress(nil, n)
// 		hash.Write(currentPK[0].Bytes())
// 		hash.Write(J[0].Bytes())
// 		hash.Write(PK[0].Bytes())

// 		a := s.suite.G1().Scalar().SetBytes(hash.Sum(nil))
// 		nodeY, _ := BigToG1Point(s.suite, currentPK)
// 		MulY = s.suite.G1().Point().Add(MulY, s.suite.G1().Point().Mul(a, nodeY))
// 	}
// 	log.Println(G1PointToBig(MulY))

// 	hash.Reset()
// 	hash.Write(s.message)

// 	H_m := s.suite.G1().Scalar().SetBytes(hash.Sum(nil))

// 	R := s.suite.G1().Point().Add(s.suite.G1().Point().Mul(s.suite.G1().Scalar().Mul(H_m, s.suite.G1().Scalar().Inv(si)), nil), s.suite.G1().Point().Mul(s.suite.G1().Scalar().Mul(r, s.suite.G1().Scalar().Inv(si)), MulY))

// 	Rbig, err := G1PointToBig(R)
// 	if err != nil {
// 		log.Println("R translate big has err, ", err)
// 	}

// 	hash.Reset()
// 	hash.Write(Rbig[0].Bytes())

// 	rNew := s.suite.G1().Scalar().SetBytes(hash.Sum(nil))
// 	log.Println(r.Equal(rNew))
// }
