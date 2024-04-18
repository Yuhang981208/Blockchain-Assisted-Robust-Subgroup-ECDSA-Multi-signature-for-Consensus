#!/bin/sh

cd "$(dirname "$0")" || exit 1


baseDir=".."

# solc --optimize --abi $baseDir/contracts/RegistryContract.sol --overwrite -o $baseDir/build/contracts/abi --base-path $baseDir/contracts --include-path $baseDir/node_modules
solc --optimize --abi $baseDir/contracts/Registry.sol --overwrite -o $baseDir/build/contracts/abi --base-path $baseDir/contracts
solc --optimize --abi $baseDir/contracts/ECDSA.sol --overwrite -o $baseDir/build/contracts/abi --base-path $baseDir/contracts


# ECDSA
abigen --abi $baseDir/build/contracts/abi/Registry.abi --pkg signer --type RegistryContract  --out $baseDir/../signerNode/pkg/signer/registry.abi.go
abigen --abi $baseDir/build/contracts/abi/ECDSA.abi --pkg signer --type SingerContract  --out $baseDir/../signerNode/pkg/signer/singer.abi.go
