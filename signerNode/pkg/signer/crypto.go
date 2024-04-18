package signer

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"
	"go.dedis.ch/kyber/v3"
	"go.dedis.ch/kyber/v3/pairing"
)

func AddressFromPrivateKey(privateKey *ecdsa.PrivateKey) (string, error) {
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)

	if !ok {
		return "", fmt.Errorf("could not cast to public key ecdsa")
	}
	return crypto.PubkeyToAddress(*publicKeyECDSA).Hex(), nil
}

func HexToScalar(suite kyber.Group, hexScalar string) (kyber.Scalar, error) {
	b, err := hex.DecodeString(hexScalar)
	if byteErr, ok := err.(hex.InvalidByteError); ok {
		return nil, fmt.Errorf("invalid hex character %q in scalar", byte(byteErr))
	} else if err != nil {
		return nil, errors.New("invalid hex data for scalar")
	}
	s := suite.Scalar()
	if err := s.UnmarshalBinary(b); err != nil {
		return nil, fmt.Errorf("unmarshal scalar binary: %w", err)
	}
	return s, nil
}

func G1PointToBig(point kyber.Point) ([2]*big.Int, error) {
	bytes, err := point.MarshalBinary()
	if err != nil {
		return [2]*big.Int{}, fmt.Errorf("marshal public key: %w", err)
	}

	if len(bytes) != 64 {
		return [2]*big.Int{}, fmt.Errorf("invalid public key length")
	}

	return [2]*big.Int{
		new(big.Int).SetBytes(bytes[:32]),
		new(big.Int).SetBytes(bytes[32:64]),
	}, nil
}

func G2PointToBig(point kyber.Point) ([4]*big.Int, error) {
	b, err := point.MarshalBinary()
	if err != nil {
		return [4]*big.Int{}, fmt.Errorf("marshal public key: %w", err)
	}

	if len(b) != 128 {
		return [4]*big.Int{}, fmt.Errorf("invalid public key length")
	}
	return [4]*big.Int{
		new(big.Int).SetBytes(b[32:64]),
		new(big.Int).SetBytes(b[:32]),
		new(big.Int).SetBytes(b[96:128]),
		new(big.Int).SetBytes(b[64:96]),
	}, nil
}

func ScalarToBig(scalar kyber.Scalar) (*big.Int, error) {
	bytes, err := scalar.MarshalBinary()
	if err != nil {
		return nil, fmt.Errorf("marshal signature: %w", err)
	}
	if len(bytes) != 32 {
		return nil, fmt.Errorf("invalid signature length")
	}
	return new(big.Int).SetBytes(bytes), nil
}

func BigToG1Point(suite pairing.Suite, pointBig [2]*big.Int) (kyber.Point, error) {
	pointBytes := make([]byte, 0)
	for k := 0; k < 32-len(pointBig[0].Bytes()); k++ {
		pointBytes = append(pointBytes, 0)
	}
	pointBytes = append(pointBytes, pointBig[0].Bytes()...)
	for k := 0; k < 32-len(pointBig[1].Bytes()); k++ {
		pointBytes = append(pointBytes, 0)
	}
	pointBytes = append(pointBytes, pointBig[1].Bytes()...)

	point := suite.G1().Point().Null()
	err := point.UnmarshalBinary(pointBytes)
	if err != nil {
		return suite.G1().Point().Null(), err
	}
	return point, nil
}

func BigToG2Point(suite pairing.Suite, pointBig [4]*big.Int) (kyber.Point, error) {
	pointBytes := make([]byte, 0)

	for _, z := range [4]int{1, 0, 3, 2} {
		// log.Println(pointBig, z, pointBig[z])
		bigByte := pointBig[z].Bytes()
		sub := 32 - len(bigByte)
		bigByte = make([]byte, 0)
		for i := 0; i < sub; i++ {
			bigByte = append(bigByte, 0)
		}

		bigByte = append(bigByte, pointBig[z].Bytes()...)
		pointBytes = append(pointBytes, bigByte...)

	}

	point := suite.G2().Point()
	err := point.UnmarshalBinary(pointBytes)

	if err != nil {
		return suite.G2().Point().Null(), err
	}
	return point, nil
}

func groupSet(pk [2]*big.Int) []byte {
	hash := sha256.New()
	hash.Write(pk[0].Bytes())
	hash.Write(pk[1].Bytes())
	return hash.Sum(nil)
}
