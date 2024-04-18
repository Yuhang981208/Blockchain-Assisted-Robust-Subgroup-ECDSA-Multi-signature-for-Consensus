package signer

// func (a *Aggregator) verifyECDSA(si kyber.Scalar, r kyber.Scalar, rByte []byte) {
// 	hash := sha256.New()

// 	_, PK, err := a.Contract.FindSM2SignersPublicKey()
// 	if err != nil {
// 		log.Println("get PkByte has err", err)
// 	}
// 	nodes, err := a.Contract.FindSM2Signers()

// 	MulY := a.suite.G1().Point().Null()

// 	for _, node := range nodes {
// 		hash.Reset()
// 		hash.Write(PK)
// 		hash.Write(a.message)
// 		hash.Write(rByte)
// 		pk, _ := a.Contract.GetSM2SignerPubkeyByAddress(nil, node)
// 		hash.Write(pk[0].Bytes())

// 		nodeY, _ := BigToG1Point(a.suite, pk)
// 		ai := a.suite.G1().Scalar().SetBytes(hash.Sum(nil))

// 		MulY = a.suite.G1().Point().Add(MulY, a.suite.G1().Point().Mul(ai, nodeY))

// 	}
// 	log.Println(MulY)
// 	hash.Reset()
// 	hash.Write(a.message)
// 	if err != nil {
// 		log.Println("R translate big has err, ", err)
// 	}
// 	H_m := a.suite.G1().Scalar().SetBytes(hash.Sum(nil))

// 	R := a.suite.G1().Point().Add(a.suite.G1().Point().Mul(a.suite.G1().Scalar().Mul(H_m, a.suite.G1().Scalar().Inv(si)), nil), a.suite.G1().Point().Mul(a.suite.G1().Scalar().Mul(r, a.suite.G1().Scalar().Inv(si)), MulY))

// 	Rbig, err := G1PointToBig(R)
// 	if err != nil {
// 		log.Println("R translate big has err, ", err)
// 	}

// 	hash.Reset()
// 	hash.Write(Rbig[0].Bytes())

// 	rNew := a.suite.G1().Scalar().SetBytes(hash.Sum(nil))
// 	log.Println(r.Equal(rNew))

// 	auth, _ := bind.NewKeyedTransactorWithChainID(a.ecdsaPrivateKey, a.chainId)

// 	sBig, _ := ScalarToBig(si)
// 	rBig, _ := ScalarToBig(r)
// 	MulYBig, _ := G1PointToBig(MulY)
// 	_, err = a.Contract.SubmitECDSA(auth, a.message, PK, sBig, rBig, nodes, MulYBig)

// 	if err != nil {
// 		log.Println("submit signature has err :", err)
// 	} else {
// 		log.Println("submit success")
// 	}

// }
