// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package signer

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// RegistrySigner is an auto generated low-level Go binding around an user-defined struct.
type RegistrySigner struct {
	Addr           common.Address
	IpAddr         string
	PubKey         [2]*big.Int
	PaillierPubkey string
	Index          *big.Int
}

// RegistryContractMetaData contains all meta data concerning the RegistryContract contract.
var RegistryContractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"Sign\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"ipAddr\",\"type\":\"string\"},{\"internalType\":\"uint256[2]\",\"name\":\"pubKey\",\"type\":\"uint256[2]\"},{\"internalType\":\"string\",\"name\":\"paillierPubkey\",\"type\":\"string\"}],\"name\":\"SignerRegister\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"SingerIsRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"countSigner\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"findSignerByIndex\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"ipAddr\",\"type\":\"string\"},{\"internalType\":\"uint256[2]\",\"name\":\"pubKey\",\"type\":\"uint256[2]\"},{\"internalType\":\"string\",\"name\":\"paillierPubkey\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"internalType\":\"structRegistry.Signer\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllPKs\",\"outputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"\",\"type\":\"uint256[2]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMessage\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getSignerByAddress\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"ipAddr\",\"type\":\"string\"},{\"internalType\":\"uint256[2]\",\"name\":\"pubKey\",\"type\":\"uint256[2]\"},{\"internalType\":\"string\",\"name\":\"paillierPubkey\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"internalType\":\"structRegistry.Signer\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getSignerPubkeyByAddress\",\"outputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"\",\"type\":\"uint256[2]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"requestSign\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// RegistryContractABI is the input ABI used to generate the binding from.
// Deprecated: Use RegistryContractMetaData.ABI instead.
var RegistryContractABI = RegistryContractMetaData.ABI

// RegistryContract is an auto generated Go binding around an Ethereum contract.
type RegistryContract struct {
	RegistryContractCaller     // Read-only binding to the contract
	RegistryContractTransactor // Write-only binding to the contract
	RegistryContractFilterer   // Log filterer for contract events
}

// RegistryContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type RegistryContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RegistryContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RegistryContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RegistryContractSession struct {
	Contract     *RegistryContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RegistryContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RegistryContractCallerSession struct {
	Contract *RegistryContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// RegistryContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RegistryContractTransactorSession struct {
	Contract     *RegistryContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// RegistryContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type RegistryContractRaw struct {
	Contract *RegistryContract // Generic contract binding to access the raw methods on
}

// RegistryContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RegistryContractCallerRaw struct {
	Contract *RegistryContractCaller // Generic read-only contract binding to access the raw methods on
}

// RegistryContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RegistryContractTransactorRaw struct {
	Contract *RegistryContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRegistryContract creates a new instance of RegistryContract, bound to a specific deployed contract.
func NewRegistryContract(address common.Address, backend bind.ContractBackend) (*RegistryContract, error) {
	contract, err := bindRegistryContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RegistryContract{RegistryContractCaller: RegistryContractCaller{contract: contract}, RegistryContractTransactor: RegistryContractTransactor{contract: contract}, RegistryContractFilterer: RegistryContractFilterer{contract: contract}}, nil
}

// NewRegistryContractCaller creates a new read-only instance of RegistryContract, bound to a specific deployed contract.
func NewRegistryContractCaller(address common.Address, caller bind.ContractCaller) (*RegistryContractCaller, error) {
	contract, err := bindRegistryContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryContractCaller{contract: contract}, nil
}

// NewRegistryContractTransactor creates a new write-only instance of RegistryContract, bound to a specific deployed contract.
func NewRegistryContractTransactor(address common.Address, transactor bind.ContractTransactor) (*RegistryContractTransactor, error) {
	contract, err := bindRegistryContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryContractTransactor{contract: contract}, nil
}

// NewRegistryContractFilterer creates a new log filterer instance of RegistryContract, bound to a specific deployed contract.
func NewRegistryContractFilterer(address common.Address, filterer bind.ContractFilterer) (*RegistryContractFilterer, error) {
	contract, err := bindRegistryContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RegistryContractFilterer{contract: contract}, nil
}

// bindRegistryContract binds a generic wrapper to an already deployed contract.
func bindRegistryContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RegistryContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RegistryContract *RegistryContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RegistryContract.Contract.RegistryContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RegistryContract *RegistryContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RegistryContract.Contract.RegistryContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RegistryContract *RegistryContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RegistryContract.Contract.RegistryContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RegistryContract *RegistryContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RegistryContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RegistryContract *RegistryContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RegistryContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RegistryContract *RegistryContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RegistryContract.Contract.contract.Transact(opts, method, params...)
}

// SingerIsRegistered is a free data retrieval call binding the contract method 0xd3ebb973.
//
// Solidity: function SingerIsRegistered(address _addr) view returns(bool)
func (_RegistryContract *RegistryContractCaller) SingerIsRegistered(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var out []interface{}
	err := _RegistryContract.contract.Call(opts, &out, "SingerIsRegistered", _addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SingerIsRegistered is a free data retrieval call binding the contract method 0xd3ebb973.
//
// Solidity: function SingerIsRegistered(address _addr) view returns(bool)
func (_RegistryContract *RegistryContractSession) SingerIsRegistered(_addr common.Address) (bool, error) {
	return _RegistryContract.Contract.SingerIsRegistered(&_RegistryContract.CallOpts, _addr)
}

// SingerIsRegistered is a free data retrieval call binding the contract method 0xd3ebb973.
//
// Solidity: function SingerIsRegistered(address _addr) view returns(bool)
func (_RegistryContract *RegistryContractCallerSession) SingerIsRegistered(_addr common.Address) (bool, error) {
	return _RegistryContract.Contract.SingerIsRegistered(&_RegistryContract.CallOpts, _addr)
}

// CountSigner is a free data retrieval call binding the contract method 0x74583ecd.
//
// Solidity: function countSigner() view returns(uint256)
func (_RegistryContract *RegistryContractCaller) CountSigner(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RegistryContract.contract.Call(opts, &out, "countSigner")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CountSigner is a free data retrieval call binding the contract method 0x74583ecd.
//
// Solidity: function countSigner() view returns(uint256)
func (_RegistryContract *RegistryContractSession) CountSigner() (*big.Int, error) {
	return _RegistryContract.Contract.CountSigner(&_RegistryContract.CallOpts)
}

// CountSigner is a free data retrieval call binding the contract method 0x74583ecd.
//
// Solidity: function countSigner() view returns(uint256)
func (_RegistryContract *RegistryContractCallerSession) CountSigner() (*big.Int, error) {
	return _RegistryContract.Contract.CountSigner(&_RegistryContract.CallOpts)
}

// FindSignerByIndex is a free data retrieval call binding the contract method 0x3886ddce.
//
// Solidity: function findSignerByIndex(uint256 _index) view returns((address,string,uint256[2],string,uint256))
func (_RegistryContract *RegistryContractCaller) FindSignerByIndex(opts *bind.CallOpts, _index *big.Int) (RegistrySigner, error) {
	var out []interface{}
	err := _RegistryContract.contract.Call(opts, &out, "findSignerByIndex", _index)

	if err != nil {
		return *new(RegistrySigner), err
	}

	out0 := *abi.ConvertType(out[0], new(RegistrySigner)).(*RegistrySigner)

	return out0, err

}

// FindSignerByIndex is a free data retrieval call binding the contract method 0x3886ddce.
//
// Solidity: function findSignerByIndex(uint256 _index) view returns((address,string,uint256[2],string,uint256))
func (_RegistryContract *RegistryContractSession) FindSignerByIndex(_index *big.Int) (RegistrySigner, error) {
	return _RegistryContract.Contract.FindSignerByIndex(&_RegistryContract.CallOpts, _index)
}

// FindSignerByIndex is a free data retrieval call binding the contract method 0x3886ddce.
//
// Solidity: function findSignerByIndex(uint256 _index) view returns((address,string,uint256[2],string,uint256))
func (_RegistryContract *RegistryContractCallerSession) FindSignerByIndex(_index *big.Int) (RegistrySigner, error) {
	return _RegistryContract.Contract.FindSignerByIndex(&_RegistryContract.CallOpts, _index)
}

// GetAllPKs is a free data retrieval call binding the contract method 0x189c5fca.
//
// Solidity: function getAllPKs() view returns(uint256[2])
func (_RegistryContract *RegistryContractCaller) GetAllPKs(opts *bind.CallOpts) ([2]*big.Int, error) {
	var out []interface{}
	err := _RegistryContract.contract.Call(opts, &out, "getAllPKs")

	if err != nil {
		return *new([2]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([2]*big.Int)).(*[2]*big.Int)

	return out0, err

}

// GetAllPKs is a free data retrieval call binding the contract method 0x189c5fca.
//
// Solidity: function getAllPKs() view returns(uint256[2])
func (_RegistryContract *RegistryContractSession) GetAllPKs() ([2]*big.Int, error) {
	return _RegistryContract.Contract.GetAllPKs(&_RegistryContract.CallOpts)
}

// GetAllPKs is a free data retrieval call binding the contract method 0x189c5fca.
//
// Solidity: function getAllPKs() view returns(uint256[2])
func (_RegistryContract *RegistryContractCallerSession) GetAllPKs() ([2]*big.Int, error) {
	return _RegistryContract.Contract.GetAllPKs(&_RegistryContract.CallOpts)
}

// GetMessage is a free data retrieval call binding the contract method 0xce6d41de.
//
// Solidity: function getMessage() view returns(bytes)
func (_RegistryContract *RegistryContractCaller) GetMessage(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _RegistryContract.contract.Call(opts, &out, "getMessage")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetMessage is a free data retrieval call binding the contract method 0xce6d41de.
//
// Solidity: function getMessage() view returns(bytes)
func (_RegistryContract *RegistryContractSession) GetMessage() ([]byte, error) {
	return _RegistryContract.Contract.GetMessage(&_RegistryContract.CallOpts)
}

// GetMessage is a free data retrieval call binding the contract method 0xce6d41de.
//
// Solidity: function getMessage() view returns(bytes)
func (_RegistryContract *RegistryContractCallerSession) GetMessage() ([]byte, error) {
	return _RegistryContract.Contract.GetMessage(&_RegistryContract.CallOpts)
}

// GetSignerByAddress is a free data retrieval call binding the contract method 0xc50bafcc.
//
// Solidity: function getSignerByAddress(address addr) view returns((address,string,uint256[2],string,uint256))
func (_RegistryContract *RegistryContractCaller) GetSignerByAddress(opts *bind.CallOpts, addr common.Address) (RegistrySigner, error) {
	var out []interface{}
	err := _RegistryContract.contract.Call(opts, &out, "getSignerByAddress", addr)

	if err != nil {
		return *new(RegistrySigner), err
	}

	out0 := *abi.ConvertType(out[0], new(RegistrySigner)).(*RegistrySigner)

	return out0, err

}

// GetSignerByAddress is a free data retrieval call binding the contract method 0xc50bafcc.
//
// Solidity: function getSignerByAddress(address addr) view returns((address,string,uint256[2],string,uint256))
func (_RegistryContract *RegistryContractSession) GetSignerByAddress(addr common.Address) (RegistrySigner, error) {
	return _RegistryContract.Contract.GetSignerByAddress(&_RegistryContract.CallOpts, addr)
}

// GetSignerByAddress is a free data retrieval call binding the contract method 0xc50bafcc.
//
// Solidity: function getSignerByAddress(address addr) view returns((address,string,uint256[2],string,uint256))
func (_RegistryContract *RegistryContractCallerSession) GetSignerByAddress(addr common.Address) (RegistrySigner, error) {
	return _RegistryContract.Contract.GetSignerByAddress(&_RegistryContract.CallOpts, addr)
}

// GetSignerPubkeyByAddress is a free data retrieval call binding the contract method 0xd999dc79.
//
// Solidity: function getSignerPubkeyByAddress(address addr) view returns(uint256[2])
func (_RegistryContract *RegistryContractCaller) GetSignerPubkeyByAddress(opts *bind.CallOpts, addr common.Address) ([2]*big.Int, error) {
	var out []interface{}
	err := _RegistryContract.contract.Call(opts, &out, "getSignerPubkeyByAddress", addr)

	if err != nil {
		return *new([2]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([2]*big.Int)).(*[2]*big.Int)

	return out0, err

}

// GetSignerPubkeyByAddress is a free data retrieval call binding the contract method 0xd999dc79.
//
// Solidity: function getSignerPubkeyByAddress(address addr) view returns(uint256[2])
func (_RegistryContract *RegistryContractSession) GetSignerPubkeyByAddress(addr common.Address) ([2]*big.Int, error) {
	return _RegistryContract.Contract.GetSignerPubkeyByAddress(&_RegistryContract.CallOpts, addr)
}

// GetSignerPubkeyByAddress is a free data retrieval call binding the contract method 0xd999dc79.
//
// Solidity: function getSignerPubkeyByAddress(address addr) view returns(uint256[2])
func (_RegistryContract *RegistryContractCallerSession) GetSignerPubkeyByAddress(addr common.Address) ([2]*big.Int, error) {
	return _RegistryContract.Contract.GetSignerPubkeyByAddress(&_RegistryContract.CallOpts, addr)
}

// SignerRegister is a paid mutator transaction binding the contract method 0x5116c6ae.
//
// Solidity: function SignerRegister(string ipAddr, uint256[2] pubKey, string paillierPubkey) payable returns()
func (_RegistryContract *RegistryContractTransactor) SignerRegister(opts *bind.TransactOpts, ipAddr string, pubKey [2]*big.Int, paillierPubkey string) (*types.Transaction, error) {
	return _RegistryContract.contract.Transact(opts, "SignerRegister", ipAddr, pubKey, paillierPubkey)
}

// SignerRegister is a paid mutator transaction binding the contract method 0x5116c6ae.
//
// Solidity: function SignerRegister(string ipAddr, uint256[2] pubKey, string paillierPubkey) payable returns()
func (_RegistryContract *RegistryContractSession) SignerRegister(ipAddr string, pubKey [2]*big.Int, paillierPubkey string) (*types.Transaction, error) {
	return _RegistryContract.Contract.SignerRegister(&_RegistryContract.TransactOpts, ipAddr, pubKey, paillierPubkey)
}

// SignerRegister is a paid mutator transaction binding the contract method 0x5116c6ae.
//
// Solidity: function SignerRegister(string ipAddr, uint256[2] pubKey, string paillierPubkey) payable returns()
func (_RegistryContract *RegistryContractTransactorSession) SignerRegister(ipAddr string, pubKey [2]*big.Int, paillierPubkey string) (*types.Transaction, error) {
	return _RegistryContract.Contract.SignerRegister(&_RegistryContract.TransactOpts, ipAddr, pubKey, paillierPubkey)
}

// RequestSign is a paid mutator transaction binding the contract method 0x84498920.
//
// Solidity: function requestSign(bytes _message) payable returns()
func (_RegistryContract *RegistryContractTransactor) RequestSign(opts *bind.TransactOpts, _message []byte) (*types.Transaction, error) {
	return _RegistryContract.contract.Transact(opts, "requestSign", _message)
}

// RequestSign is a paid mutator transaction binding the contract method 0x84498920.
//
// Solidity: function requestSign(bytes _message) payable returns()
func (_RegistryContract *RegistryContractSession) RequestSign(_message []byte) (*types.Transaction, error) {
	return _RegistryContract.Contract.RequestSign(&_RegistryContract.TransactOpts, _message)
}

// RequestSign is a paid mutator transaction binding the contract method 0x84498920.
//
// Solidity: function requestSign(bytes _message) payable returns()
func (_RegistryContract *RegistryContractTransactorSession) RequestSign(_message []byte) (*types.Transaction, error) {
	return _RegistryContract.Contract.RequestSign(&_RegistryContract.TransactOpts, _message)
}

// RegistryContractSignIterator is returned from FilterSign and is used to iterate over the raw logs and unpacked data for Sign events raised by the RegistryContract contract.
type RegistryContractSignIterator struct {
	Event *RegistryContractSign // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RegistryContractSignIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryContractSign)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RegistryContractSign)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RegistryContractSignIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryContractSignIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryContractSign represents a Sign event raised by the RegistryContract contract.
type RegistryContractSign struct {
	Message []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSign is a free log retrieval operation binding the contract event 0xf1ca1132f9dcf9cab9c8d9dfe5c481772f8de990f398d47775f6a6e5443eaee8.
//
// Solidity: event Sign(bytes message)
func (_RegistryContract *RegistryContractFilterer) FilterSign(opts *bind.FilterOpts) (*RegistryContractSignIterator, error) {

	logs, sub, err := _RegistryContract.contract.FilterLogs(opts, "Sign")
	if err != nil {
		return nil, err
	}
	return &RegistryContractSignIterator{contract: _RegistryContract.contract, event: "Sign", logs: logs, sub: sub}, nil
}

// WatchSign is a free log subscription operation binding the contract event 0xf1ca1132f9dcf9cab9c8d9dfe5c481772f8de990f398d47775f6a6e5443eaee8.
//
// Solidity: event Sign(bytes message)
func (_RegistryContract *RegistryContractFilterer) WatchSign(opts *bind.WatchOpts, sink chan<- *RegistryContractSign) (event.Subscription, error) {

	logs, sub, err := _RegistryContract.contract.WatchLogs(opts, "Sign")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryContractSign)
				if err := _RegistryContract.contract.UnpackLog(event, "Sign", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSign is a log parse operation binding the contract event 0xf1ca1132f9dcf9cab9c8d9dfe5c481772f8de990f398d47775f6a6e5443eaee8.
//
// Solidity: event Sign(bytes message)
func (_RegistryContract *RegistryContractFilterer) ParseSign(log types.Log) (*RegistryContractSign, error) {
	event := new(RegistryContractSign)
	if err := _RegistryContract.contract.UnpackLog(event, "Sign", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
