package signer

// import (
// 	"crypto/ecdsa"
// 	"encoding/hex"
// 	"errors"
// 	"fmt"

// 	"github.com/decred/dcrd/dcrec/secp256k1/v4"
// 	"github.com/ethereum/go-ethereum/crypto"
// 	"go.dedis.ch/kyber/v3"
// )

// func AddressFromPrivateKey(privateKey *ecdsa.PrivateKey) (string, error) {
// 	publicKey := privateKey.Public()
// 	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)

// 	if !ok {
// 		return "", fmt.Errorf("could not cast to public key ecdsa")
// 	}
// 	return crypto.PubkeyToAddress(*publicKeyECDSA).Hex(), nil
// }

// func HexToScalar(suite kyber.Group, hexScalar string) (kyber.Scalar, error) {
// 	b, err := hex.DecodeString(hexScalar)
// 	if byteErr, ok := err.(hex.InvalidByteError); ok {
// 		return nil, fmt.Errorf("invalid hex character %q in scalar", byte(byteErr))
// 	} else if err != nil {
// 		return nil, errors.New("invalid hex data for scalar")
// 	}
// 	s := suite.Scalar()
// 	if err := s.UnmarshalBinary(b); err != nil {
// 		return nil, fmt.Errorf("unmarshal scalar binary: %w", err)
// 	}
// 	return s, nil
// }

// // secp256k1.JacobianPoint to bytes
// func pointToBytes(point secp256k1.JacobianPoint) []byte {
// 	res := make([]byte, 0)
// 	for _, b := range point.X.Bytes() {
// 		res = append(res, b)
// 	}
// 	for _, b := range point.Y.Bytes() {
// 		res = append(res, b)
// 	}
// 	for _, b := range point.Z.Bytes() {
// 		res = append(res, b)
// 	}
// 	return res
// }

// func BytesToPoint(b []byte) secp256k1.JacobianPoint {
// 	X := new(secp256k1.FieldVal)
// 	X.SetBytes((*[32]byte)(b[0:32]))
// 	Y := new(secp256k1.FieldVal)
// 	Y.SetBytes((*[32]byte)(b[32:64]))
// 	Z := new(secp256k1.FieldVal)
// 	Z.SetBytes((*[32]byte)(b[64:96]))

// 	return secp256k1.MakeJacobianPoint(X, Y, Z)
// }
