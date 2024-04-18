package signer

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type ContractWrapper struct {
	*RegistryContract
	*SingerContract
}

// 获取所有用户的公钥字节流
func (r *ContractWrapper) FindSignersPublicKey() ([]byte, error) {
	count, err := r.RegistryContract.CountSigner(nil)
	if err != nil {
		return nil, fmt.Errorf("count oracle nodes: %w", err)
	}
	nodePKsByte := make([]byte, 0)
	for i := int64(0); i < count.Int64(); i++ {
		node, err := r.RegistryContract.FindSignerByIndex(nil, big.NewInt(i))
		if err != nil {
			return nil, fmt.Errorf("find oracle node by index %d: %w", i, err)
		}

		pk := make([]byte, 0)
		for _, num := range node.PubKey {
			pk = append(pk, num.Bytes()...)
		}
		nodePKsByte = append(nodePKsByte, pk...)
	}
	return nodePKsByte, nil
}

// 获取所有用户
func (r *ContractWrapper) FindSigners() ([]common.Address, error) {
	count, err := r.RegistryContract.CountSigner(nil)
	if err != nil {
		return nil, fmt.Errorf("count oracle nodes: %w", err)
	}
	nodes := make([]common.Address, count.Int64())
	for i := int64(0); i < count.Int64(); i++ { // 此时不存在聚合器
		node, err := r.RegistryContract.FindSignerByIndex(nil, big.NewInt(i))
		if err != nil {
			return nil, fmt.Errorf("find oracle node by index %d: %w", i, err)
		}

		nodes[i] = node.Addr
	}
	return nodes, nil
}

// 获取子分组用户
func (n *ContractWrapper) FindEnrollSingers() ([]common.Address, error) {
	count, err := n.SingerContract.CountEnrollNodes(nil)
	if err != nil {
		return nil, fmt.Errorf("count oracle nodes: %w", err)
	}
	nodeEnrolls := make([]common.Address, count.Int64())
	for i := int64(0); i < count.Int64(); i++ {
		enrollNode, err := n.FindEnrollNodeByIndex(nil, big.NewInt(i))
		if err != nil {
			return nil, fmt.Errorf("find enrolloracle node by index %d: %w", i, err)
		}
		nodeEnrolls[i] = enrollNode
		if err != nil {
			return nil, fmt.Errorf("get enroll nodes has err %w", err)
		}
	}
	return nodeEnrolls, nil
}

// 获取子分组用户所有的公钥字节流集合
func (r *ContractWrapper) FindEnrollSingersPublicKey() ([]byte, error) {
	count, err := r.SingerContract.CountEnrollNodes(nil)
	if err != nil {
		return nil, fmt.Errorf("count oracle nodes: %w", err)
	}

	enrollNodePublicKeys := make([]byte, 0)
	for i := int64(0); i < count.Int64(); i++ {
		enrollNode, err := r.FindEnrollNodeByIndex(nil, big.NewInt(i))
		if err != nil {
			return nil, fmt.Errorf("find enrolloracle node by index %d: %w", i, err)
		}
		pk, err := r.RegistryContract.GetSignerPubkeyByAddress(nil, enrollNode)
		if err != nil {
			return nil, fmt.Errorf("find oracle node by index %d: %w", i, err)
		}

		pkByte := make([]byte, 0)
		for _, num := range pk {
			pkByte = append(pkByte, num.Bytes()...)
		}
		enrollNodePublicKeys = append(enrollNodePublicKeys, pkByte...)
	}
	return enrollNodePublicKeys, nil
}
