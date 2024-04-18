// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "./crypto/BN256G1.sol";
import "./crypto/BN256G2.sol";
import "./Registry.sol";

contract ECDSA {
    Registry private registry;
    uint256[2] private enrollNodePK;
    constructor(address _registryContract) {
        registry = Registry(_registryContract);
    }
    event SignatureBegin();
    event verifyMulSignature();
    address[] private enrollNodes;

    uint256 private MulSignature = 0;
    uint256 private count = 0;
    uint256 private rForMulSignature = 0;
    uint256[2] private aggPublicKey;
    uint256 private usePublickey = 0;
    bool isKeyAgg = false;
    // --------------------------------------------------------------------------------------------------------------------------------------------------------
    // --------------------------------------------------------------------------------------------------------------------------------------------------------
    // --------------------------------------------------------------------------------------------------------------------------------------------------------
    function getUsePublicKeyCount() public view returns (uint256) {
        return usePublickey;
    }

    function enRoll() external {
        require(
            enrollNodes.length <= registry.countSigner() / 2,
            "Enrollment closed!"
        );
        require(
            registry.SingerIsRegistered(msg.sender),
            "The Oracle doesn't registered"
        );
        require(!IsEnroll(msg.sender), "already enrolled");

        enrollNodes.push(msg.sender);
        uint256[2] memory pubkey = registry.getSignerPubkeyByAddress(
            msg.sender
        );
        (enrollNodePK[0], enrollNodePK[1]) = BN256G1.addPoint(
            [enrollNodePK[0], enrollNodePK[1], pubkey[0], pubkey[1]]
        );
        if (enrollNodes.length == registry.countSigner() / 2 + 1) {
            isKeyAgg = false;
            emit SignatureBegin();
        }
    }

    function IsEnroll(address _addr) public view returns (bool) {
        if (enrollNodes.length == 0) return false;
        for (uint32 i = 0; i < enrollNodes.length; i++) {
            if (enrollNodes[i] == _addr) {
                return true;
            }
        }
        return false;
    }

    // 根据索引查找报名节点
    function findEnrollNodeByIndex(
        uint256 _index
    ) public view returns (address) {
        require(_index >= 0 && _index < enrollNodes.length, "not found");
        return enrollNodes[_index];
    }

    function getEnrollPKs() public view returns (uint256[2] memory) {
        return enrollNodePK;
    }

    // 返回报名总数
    function countEnrollNodes() external view returns (uint256) {
        return enrollNodes.length;
    }

    // 验证单个签名的合法性
    function submitECDSA(
        uint256 r_i,
        uint256 r,
        uint256 s,
        uint256 sigma,
        bytes calldata messageAndGtag
    ) external payable {
        require(
            verifySingleECDSA(messageAndGtag, r_i, r, s),
            "single ECDSA signature verify fail!"
        );

        if (rForMulSignature == 0) {
            rForMulSignature = r;
        } else {
            require(rForMulSignature == r, "r is not equal");
        }
        MulSignature = addmod(MulSignature, sigma, BN256G1.NN);
        MulSignature = addmod(MulSignature, s, BN256G1.NN);
        count++;

        if (enrollNodes.length == count) {
            // 此时要开始执行聚合签名算法以及聚合签名验证算法
            emit verifyMulSignature();
        }
    }

    function aggKey() public payable {
        if (!isKeyAgg) {
            isKeyAgg = true;
            uint256[2] memory MulY; // 是后面整个部分
            uint256[2] memory registryNodePK = registry.getAllPKs();

            for (uint i = 0; i < enrollNodes.length; i++) {
                uint256[2] memory publicKey = registry.getSignerPubkeyByAddress(
                    enrollNodes[i]
                );
                bytes32 a = sha256(
                    abi.encodePacked(
                        BN256G1.toBytes(publicKey[0]),
                        BN256G1.toBytes(enrollNodePK[0]),
                        BN256G1.toBytes(registryNodePK[0])
                    )
                );

                uint256[2] memory tmp0;
                (tmp0[0], tmp0[1]) = BN256G1.mulPoint(
                    [publicKey[0], publicKey[1], uint256(a)]
                );

                (MulY[0], MulY[1]) = BN256G1.addPoint(
                    [MulY[0], MulY[1], tmp0[0], tmp0[1]]
                );
            }

            aggPublicKey = MulY;
        }
    }

    function verifyMul(bytes calldata messageAndGtag) external payable {
        uint256 sInverse = BN256G1.modInverse(MulSignature);
        uint256 h_m = uint256(sha256(abi.encodePacked(messageAndGtag)));

        uint256 tmp1 = mulmod(h_m, sInverse, BN256G1.NN);
        uint256[2] memory tmp1Point;
        (tmp1Point[0], tmp1Point[1]) = BN256G1.mulPoint(
            [BN256G1.GX, BN256G1.GY, tmp1]
        );

        uint256 tmp2 = mulmod(rForMulSignature, sInverse, BN256G1.NN);

        uint256[2] memory tmp2Point;
        (tmp2Point[0], tmp2Point[1]) = BN256G1.mulPoint(
            [aggPublicKey[0], aggPublicKey[1], tmp2]
        );
        uint256[2] memory R;
        (R[0], R[1]) = BN256G1.addPoint(
            [tmp1Point[0], tmp1Point[1], tmp2Point[0], tmp2Point[1]]
        );
        uint256 h_R = uint256(sha256(abi.encodePacked(BN256G1.toBytes(R[0]))));
        uint256 r = h_R % BN256G1.NN;
        require(r == rForMulSignature, "ECDSA  check failed!");
    }

    function verifySingleECDSA(
        bytes calldata messageAndGtag,
        uint256 r_i, // k^-1G
        uint256 r,
        uint256 s
    ) private returns (bool) {
        // uint256[2] memory Y; // 是后面整个部分
        uint256[2] memory registryNodePK = registry.getAllPKs();

        uint256[2] memory publicKey = registry.getSignerPubkeyByAddress(
            msg.sender
        );
        bytes32 a = sha256(
            abi.encodePacked(
                BN256G1.toBytes(publicKey[0]),
                BN256G1.toBytes(enrollNodePK[0]),
                BN256G1.toBytes(registryNodePK[0])
            )
        );

        uint256 sInverse = BN256G1.modInverse(s);
        uint256 h_m = uint256(sha256(abi.encodePacked(messageAndGtag)));

        uint256 tmp1 = mulmod(h_m, sInverse, BN256G1.NN);
        uint256[2] memory tmp1Point;
        (tmp1Point[0], tmp1Point[1]) = BN256G1.mulPoint(
            [BN256G1.GX, BN256G1.GY, tmp1]
        );

        uint256 tmp2 = mulmod(r, sInverse, BN256G1.NN);
        tmp2 = mulmod(tmp2, uint256(a), BN256G1.NN);
        uint256[2] memory tmp2Point;
        (tmp2Point[0], tmp2Point[1]) = BN256G1.mulPoint(
            [publicKey[0], publicKey[1], tmp2]
        );
        uint256[2] memory R;
        (R[0], R[1]) = BN256G1.addPoint(
            [tmp1Point[0], tmp1Point[1], tmp2Point[0], tmp2Point[1]]
        );

        return R[0] % BN256G1.NN == r_i;
    }
}
