// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;
import "./crypto/BN256G1.sol";
contract Registry {
    bytes private message;
    uint256[2] private registryNodePK;
    event Sign(bytes message);

    struct Signer {
        address addr; // eth address
        string ipAddr; // IP
        uint256[2] pubKey; // 用户自己选的公私钥对
        string paillierPubkey; // paillier 的公钥
        uint256 index; // 当前节点的编号
    }

    mapping(address => Signer) private SignerMap;

    address[] private SignerArr;

    // 记录报名节点的地址；

    function SignerRegister(
        string calldata ipAddr,
        uint256[2] calldata pubKey,
        string calldata paillierPubkey
    ) external payable {
        require(SignerMap[msg.sender].addr != msg.sender, "already");
        Signer storage signer = SignerMap[msg.sender];
        signer.addr = msg.sender;
        signer.ipAddr = ipAddr;
        signer.pubKey = pubKey;
        signer.paillierPubkey = paillierPubkey;

        signer.index = SignerArr.length;
        SignerArr.push(msg.sender);
        (registryNodePK[0], registryNodePK[1]) = BN256G1.addPoint(
            [registryNodePK[0], registryNodePK[1], pubKey[0], pubKey[1]]
        );
    }

    function getSignerByAddress(
        address addr
    ) public view returns (Signer memory) {
        return SignerMap[addr];
    }

    function getSignerPubkeyByAddress(
        address addr
    ) public view returns (uint256[2] memory) {
        return SignerMap[addr].pubKey;
    }

    function requestSign(bytes memory _message) external payable {
        message = _message;
        emit Sign(message);
    }

    function countSigner() public view returns (uint256) {
        return SignerArr.length;
    }

    function findSignerByIndex(
        uint256 _index
    ) public view returns (Signer memory) {
        require(_index >= 0 && _index < SignerArr.length, "not found");
        return SignerMap[SignerArr[_index]];
    }

    function SingerIsRegistered(address _addr) public view returns (bool) {
        if (SignerArr.length == 0) return false;
        return SignerArr[SignerMap[_addr].index] == _addr;
    }

    function getMessage() public view returns (bytes memory) {
        return message;
    }
    function getAllPKs() public view returns (uint256[2] memory) {
        return registryNodePK;
    }

    // function isAggregator(address addr) public view returns (bool) {
    //     return aggregator == addr;
    // }

    // function getAggregator() public view returns (address) {
    //     return aggregator;
    // }
    // // test
    // struct testSigner {
    //     address addr; // eth address
    //     string ipAddr; // IP
    //     // uint256[4] pubKey; // 用户自己选的公私钥对
    //     bytes identity; // identity；
    //     uint256[2] identityPubKey; // H(identity), H: string -> G1 point
    //     uint256 index; // 当前节点的编号
    // }
    // mapping(address => testSigner) private testSignerMap;

    // address[] private testSignerArr;

    // function testSignerRegister(
    //     string calldata ipAddr,
    //     uint256[2] calldata pubKey,
    //     bytes calldata identity
    // ) external payable {
    //     require(testSignerMap[msg.sender].addr != msg.sender, "already");
    //     testSigner storage signer = testSignerMap[msg.sender];
    //     signer.addr = msg.sender;
    //     signer.ipAddr = ipAddr;
    //     signer.identityPubKey = pubKey;
    //     signer.identity = identity;

    //     if (testSignerArr.length == 0) {
    //         aggregator = msg.sender;
    //     }

    //     signer.index = testSignerArr.length;
    //     testSignerArr.push(msg.sender);
    // }
}
