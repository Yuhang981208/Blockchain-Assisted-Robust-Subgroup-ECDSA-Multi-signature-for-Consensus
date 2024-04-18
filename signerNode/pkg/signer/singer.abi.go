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

// SingerContractMetaData contains all meta data concerning the SingerContract contract.
var SingerContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registryContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"SignatureBegin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"verifyMulSignature\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"IsEnroll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"aggKey\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"countEnrollNodes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enRoll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"findEnrollNodeByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEnrollPKs\",\"outputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"\",\"type\":\"uint256[2]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUsePublicKeyCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"r_i\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sigma\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"messageAndGtag\",\"type\":\"bytes\"}],\"name\":\"submitECDSA\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"messageAndGtag\",\"type\":\"bytes\"}],\"name\":\"verifyMul\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// SingerContractABI is the input ABI used to generate the binding from.
// Deprecated: Use SingerContractMetaData.ABI instead.
var SingerContractABI = SingerContractMetaData.ABI

// SingerContract is an auto generated Go binding around an Ethereum contract.
type SingerContract struct {
	SingerContractCaller     // Read-only binding to the contract
	SingerContractTransactor // Write-only binding to the contract
	SingerContractFilterer   // Log filterer for contract events
}

// SingerContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type SingerContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SingerContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SingerContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SingerContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SingerContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SingerContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SingerContractSession struct {
	Contract     *SingerContract   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SingerContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SingerContractCallerSession struct {
	Contract *SingerContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// SingerContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SingerContractTransactorSession struct {
	Contract     *SingerContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// SingerContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type SingerContractRaw struct {
	Contract *SingerContract // Generic contract binding to access the raw methods on
}

// SingerContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SingerContractCallerRaw struct {
	Contract *SingerContractCaller // Generic read-only contract binding to access the raw methods on
}

// SingerContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SingerContractTransactorRaw struct {
	Contract *SingerContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSingerContract creates a new instance of SingerContract, bound to a specific deployed contract.
func NewSingerContract(address common.Address, backend bind.ContractBackend) (*SingerContract, error) {
	contract, err := bindSingerContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SingerContract{SingerContractCaller: SingerContractCaller{contract: contract}, SingerContractTransactor: SingerContractTransactor{contract: contract}, SingerContractFilterer: SingerContractFilterer{contract: contract}}, nil
}

// NewSingerContractCaller creates a new read-only instance of SingerContract, bound to a specific deployed contract.
func NewSingerContractCaller(address common.Address, caller bind.ContractCaller) (*SingerContractCaller, error) {
	contract, err := bindSingerContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SingerContractCaller{contract: contract}, nil
}

// NewSingerContractTransactor creates a new write-only instance of SingerContract, bound to a specific deployed contract.
func NewSingerContractTransactor(address common.Address, transactor bind.ContractTransactor) (*SingerContractTransactor, error) {
	contract, err := bindSingerContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SingerContractTransactor{contract: contract}, nil
}

// NewSingerContractFilterer creates a new log filterer instance of SingerContract, bound to a specific deployed contract.
func NewSingerContractFilterer(address common.Address, filterer bind.ContractFilterer) (*SingerContractFilterer, error) {
	contract, err := bindSingerContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SingerContractFilterer{contract: contract}, nil
}

// bindSingerContract binds a generic wrapper to an already deployed contract.
func bindSingerContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SingerContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SingerContract *SingerContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SingerContract.Contract.SingerContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SingerContract *SingerContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SingerContract.Contract.SingerContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SingerContract *SingerContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SingerContract.Contract.SingerContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SingerContract *SingerContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SingerContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SingerContract *SingerContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SingerContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SingerContract *SingerContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SingerContract.Contract.contract.Transact(opts, method, params...)
}

// IsEnroll is a free data retrieval call binding the contract method 0xe81d4903.
//
// Solidity: function IsEnroll(address _addr) view returns(bool)
func (_SingerContract *SingerContractCaller) IsEnroll(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var out []interface{}
	err := _SingerContract.contract.Call(opts, &out, "IsEnroll", _addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEnroll is a free data retrieval call binding the contract method 0xe81d4903.
//
// Solidity: function IsEnroll(address _addr) view returns(bool)
func (_SingerContract *SingerContractSession) IsEnroll(_addr common.Address) (bool, error) {
	return _SingerContract.Contract.IsEnroll(&_SingerContract.CallOpts, _addr)
}

// IsEnroll is a free data retrieval call binding the contract method 0xe81d4903.
//
// Solidity: function IsEnroll(address _addr) view returns(bool)
func (_SingerContract *SingerContractCallerSession) IsEnroll(_addr common.Address) (bool, error) {
	return _SingerContract.Contract.IsEnroll(&_SingerContract.CallOpts, _addr)
}

// CountEnrollNodes is a free data retrieval call binding the contract method 0xfc769ee9.
//
// Solidity: function countEnrollNodes() view returns(uint256)
func (_SingerContract *SingerContractCaller) CountEnrollNodes(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SingerContract.contract.Call(opts, &out, "countEnrollNodes")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CountEnrollNodes is a free data retrieval call binding the contract method 0xfc769ee9.
//
// Solidity: function countEnrollNodes() view returns(uint256)
func (_SingerContract *SingerContractSession) CountEnrollNodes() (*big.Int, error) {
	return _SingerContract.Contract.CountEnrollNodes(&_SingerContract.CallOpts)
}

// CountEnrollNodes is a free data retrieval call binding the contract method 0xfc769ee9.
//
// Solidity: function countEnrollNodes() view returns(uint256)
func (_SingerContract *SingerContractCallerSession) CountEnrollNodes() (*big.Int, error) {
	return _SingerContract.Contract.CountEnrollNodes(&_SingerContract.CallOpts)
}

// FindEnrollNodeByIndex is a free data retrieval call binding the contract method 0x7e985fd0.
//
// Solidity: function findEnrollNodeByIndex(uint256 _index) view returns(address)
func (_SingerContract *SingerContractCaller) FindEnrollNodeByIndex(opts *bind.CallOpts, _index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SingerContract.contract.Call(opts, &out, "findEnrollNodeByIndex", _index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FindEnrollNodeByIndex is a free data retrieval call binding the contract method 0x7e985fd0.
//
// Solidity: function findEnrollNodeByIndex(uint256 _index) view returns(address)
func (_SingerContract *SingerContractSession) FindEnrollNodeByIndex(_index *big.Int) (common.Address, error) {
	return _SingerContract.Contract.FindEnrollNodeByIndex(&_SingerContract.CallOpts, _index)
}

// FindEnrollNodeByIndex is a free data retrieval call binding the contract method 0x7e985fd0.
//
// Solidity: function findEnrollNodeByIndex(uint256 _index) view returns(address)
func (_SingerContract *SingerContractCallerSession) FindEnrollNodeByIndex(_index *big.Int) (common.Address, error) {
	return _SingerContract.Contract.FindEnrollNodeByIndex(&_SingerContract.CallOpts, _index)
}

// GetEnrollPKs is a free data retrieval call binding the contract method 0x4aa4b725.
//
// Solidity: function getEnrollPKs() view returns(uint256[2])
func (_SingerContract *SingerContractCaller) GetEnrollPKs(opts *bind.CallOpts) ([2]*big.Int, error) {
	var out []interface{}
	err := _SingerContract.contract.Call(opts, &out, "getEnrollPKs")

	if err != nil {
		return *new([2]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([2]*big.Int)).(*[2]*big.Int)

	return out0, err

}

// GetEnrollPKs is a free data retrieval call binding the contract method 0x4aa4b725.
//
// Solidity: function getEnrollPKs() view returns(uint256[2])
func (_SingerContract *SingerContractSession) GetEnrollPKs() ([2]*big.Int, error) {
	return _SingerContract.Contract.GetEnrollPKs(&_SingerContract.CallOpts)
}

// GetEnrollPKs is a free data retrieval call binding the contract method 0x4aa4b725.
//
// Solidity: function getEnrollPKs() view returns(uint256[2])
func (_SingerContract *SingerContractCallerSession) GetEnrollPKs() ([2]*big.Int, error) {
	return _SingerContract.Contract.GetEnrollPKs(&_SingerContract.CallOpts)
}

// GetUsePublicKeyCount is a free data retrieval call binding the contract method 0x4068de4a.
//
// Solidity: function getUsePublicKeyCount() view returns(uint256)
func (_SingerContract *SingerContractCaller) GetUsePublicKeyCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SingerContract.contract.Call(opts, &out, "getUsePublicKeyCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUsePublicKeyCount is a free data retrieval call binding the contract method 0x4068de4a.
//
// Solidity: function getUsePublicKeyCount() view returns(uint256)
func (_SingerContract *SingerContractSession) GetUsePublicKeyCount() (*big.Int, error) {
	return _SingerContract.Contract.GetUsePublicKeyCount(&_SingerContract.CallOpts)
}

// GetUsePublicKeyCount is a free data retrieval call binding the contract method 0x4068de4a.
//
// Solidity: function getUsePublicKeyCount() view returns(uint256)
func (_SingerContract *SingerContractCallerSession) GetUsePublicKeyCount() (*big.Int, error) {
	return _SingerContract.Contract.GetUsePublicKeyCount(&_SingerContract.CallOpts)
}

// AggKey is a paid mutator transaction binding the contract method 0x4a98662d.
//
// Solidity: function aggKey() payable returns()
func (_SingerContract *SingerContractTransactor) AggKey(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SingerContract.contract.Transact(opts, "aggKey")
}

// AggKey is a paid mutator transaction binding the contract method 0x4a98662d.
//
// Solidity: function aggKey() payable returns()
func (_SingerContract *SingerContractSession) AggKey() (*types.Transaction, error) {
	return _SingerContract.Contract.AggKey(&_SingerContract.TransactOpts)
}

// AggKey is a paid mutator transaction binding the contract method 0x4a98662d.
//
// Solidity: function aggKey() payable returns()
func (_SingerContract *SingerContractTransactorSession) AggKey() (*types.Transaction, error) {
	return _SingerContract.Contract.AggKey(&_SingerContract.TransactOpts)
}

// EnRoll is a paid mutator transaction binding the contract method 0x8079c59e.
//
// Solidity: function enRoll() returns()
func (_SingerContract *SingerContractTransactor) EnRoll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SingerContract.contract.Transact(opts, "enRoll")
}

// EnRoll is a paid mutator transaction binding the contract method 0x8079c59e.
//
// Solidity: function enRoll() returns()
func (_SingerContract *SingerContractSession) EnRoll() (*types.Transaction, error) {
	return _SingerContract.Contract.EnRoll(&_SingerContract.TransactOpts)
}

// EnRoll is a paid mutator transaction binding the contract method 0x8079c59e.
//
// Solidity: function enRoll() returns()
func (_SingerContract *SingerContractTransactorSession) EnRoll() (*types.Transaction, error) {
	return _SingerContract.Contract.EnRoll(&_SingerContract.TransactOpts)
}

// SubmitECDSA is a paid mutator transaction binding the contract method 0x37a0a99a.
//
// Solidity: function submitECDSA(uint256 r_i, uint256 r, uint256 s, uint256 sigma, bytes messageAndGtag) payable returns()
func (_SingerContract *SingerContractTransactor) SubmitECDSA(opts *bind.TransactOpts, r_i *big.Int, r *big.Int, s *big.Int, sigma *big.Int, messageAndGtag []byte) (*types.Transaction, error) {
	return _SingerContract.contract.Transact(opts, "submitECDSA", r_i, r, s, sigma, messageAndGtag)
}

// SubmitECDSA is a paid mutator transaction binding the contract method 0x37a0a99a.
//
// Solidity: function submitECDSA(uint256 r_i, uint256 r, uint256 s, uint256 sigma, bytes messageAndGtag) payable returns()
func (_SingerContract *SingerContractSession) SubmitECDSA(r_i *big.Int, r *big.Int, s *big.Int, sigma *big.Int, messageAndGtag []byte) (*types.Transaction, error) {
	return _SingerContract.Contract.SubmitECDSA(&_SingerContract.TransactOpts, r_i, r, s, sigma, messageAndGtag)
}

// SubmitECDSA is a paid mutator transaction binding the contract method 0x37a0a99a.
//
// Solidity: function submitECDSA(uint256 r_i, uint256 r, uint256 s, uint256 sigma, bytes messageAndGtag) payable returns()
func (_SingerContract *SingerContractTransactorSession) SubmitECDSA(r_i *big.Int, r *big.Int, s *big.Int, sigma *big.Int, messageAndGtag []byte) (*types.Transaction, error) {
	return _SingerContract.Contract.SubmitECDSA(&_SingerContract.TransactOpts, r_i, r, s, sigma, messageAndGtag)
}

// VerifyMul is a paid mutator transaction binding the contract method 0x8270a2e0.
//
// Solidity: function verifyMul(bytes messageAndGtag) payable returns()
func (_SingerContract *SingerContractTransactor) VerifyMul(opts *bind.TransactOpts, messageAndGtag []byte) (*types.Transaction, error) {
	return _SingerContract.contract.Transact(opts, "verifyMul", messageAndGtag)
}

// VerifyMul is a paid mutator transaction binding the contract method 0x8270a2e0.
//
// Solidity: function verifyMul(bytes messageAndGtag) payable returns()
func (_SingerContract *SingerContractSession) VerifyMul(messageAndGtag []byte) (*types.Transaction, error) {
	return _SingerContract.Contract.VerifyMul(&_SingerContract.TransactOpts, messageAndGtag)
}

// VerifyMul is a paid mutator transaction binding the contract method 0x8270a2e0.
//
// Solidity: function verifyMul(bytes messageAndGtag) payable returns()
func (_SingerContract *SingerContractTransactorSession) VerifyMul(messageAndGtag []byte) (*types.Transaction, error) {
	return _SingerContract.Contract.VerifyMul(&_SingerContract.TransactOpts, messageAndGtag)
}

// SingerContractSignatureBeginIterator is returned from FilterSignatureBegin and is used to iterate over the raw logs and unpacked data for SignatureBegin events raised by the SingerContract contract.
type SingerContractSignatureBeginIterator struct {
	Event *SingerContractSignatureBegin // Event containing the contract specifics and raw log

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
func (it *SingerContractSignatureBeginIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SingerContractSignatureBegin)
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
		it.Event = new(SingerContractSignatureBegin)
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
func (it *SingerContractSignatureBeginIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SingerContractSignatureBeginIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SingerContractSignatureBegin represents a SignatureBegin event raised by the SingerContract contract.
type SingerContractSignatureBegin struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSignatureBegin is a free log retrieval operation binding the contract event 0xec62c1e142d57b58e0b31fcc896eaa4fb5b20bfcaa83c1ecc0857e778fb33edf.
//
// Solidity: event SignatureBegin()
func (_SingerContract *SingerContractFilterer) FilterSignatureBegin(opts *bind.FilterOpts) (*SingerContractSignatureBeginIterator, error) {

	logs, sub, err := _SingerContract.contract.FilterLogs(opts, "SignatureBegin")
	if err != nil {
		return nil, err
	}
	return &SingerContractSignatureBeginIterator{contract: _SingerContract.contract, event: "SignatureBegin", logs: logs, sub: sub}, nil
}

// WatchSignatureBegin is a free log subscription operation binding the contract event 0xec62c1e142d57b58e0b31fcc896eaa4fb5b20bfcaa83c1ecc0857e778fb33edf.
//
// Solidity: event SignatureBegin()
func (_SingerContract *SingerContractFilterer) WatchSignatureBegin(opts *bind.WatchOpts, sink chan<- *SingerContractSignatureBegin) (event.Subscription, error) {

	logs, sub, err := _SingerContract.contract.WatchLogs(opts, "SignatureBegin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SingerContractSignatureBegin)
				if err := _SingerContract.contract.UnpackLog(event, "SignatureBegin", log); err != nil {
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

// ParseSignatureBegin is a log parse operation binding the contract event 0xec62c1e142d57b58e0b31fcc896eaa4fb5b20bfcaa83c1ecc0857e778fb33edf.
//
// Solidity: event SignatureBegin()
func (_SingerContract *SingerContractFilterer) ParseSignatureBegin(log types.Log) (*SingerContractSignatureBegin, error) {
	event := new(SingerContractSignatureBegin)
	if err := _SingerContract.contract.UnpackLog(event, "SignatureBegin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SingerContractVerifyMulSignatureIterator is returned from FilterVerifyMulSignature and is used to iterate over the raw logs and unpacked data for VerifyMulSignature events raised by the SingerContract contract.
type SingerContractVerifyMulSignatureIterator struct {
	Event *SingerContractVerifyMulSignature // Event containing the contract specifics and raw log

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
func (it *SingerContractVerifyMulSignatureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SingerContractVerifyMulSignature)
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
		it.Event = new(SingerContractVerifyMulSignature)
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
func (it *SingerContractVerifyMulSignatureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SingerContractVerifyMulSignatureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SingerContractVerifyMulSignature represents a VerifyMulSignature event raised by the SingerContract contract.
type SingerContractVerifyMulSignature struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterVerifyMulSignature is a free log retrieval operation binding the contract event 0x62257d4dafbcc200d0ab7688037ac0089af1ad7c87181f1df22a3b0f31d26263.
//
// Solidity: event verifyMulSignature()
func (_SingerContract *SingerContractFilterer) FilterVerifyMulSignature(opts *bind.FilterOpts) (*SingerContractVerifyMulSignatureIterator, error) {

	logs, sub, err := _SingerContract.contract.FilterLogs(opts, "verifyMulSignature")
	if err != nil {
		return nil, err
	}
	return &SingerContractVerifyMulSignatureIterator{contract: _SingerContract.contract, event: "verifyMulSignature", logs: logs, sub: sub}, nil
}

// WatchVerifyMulSignature is a free log subscription operation binding the contract event 0x62257d4dafbcc200d0ab7688037ac0089af1ad7c87181f1df22a3b0f31d26263.
//
// Solidity: event verifyMulSignature()
func (_SingerContract *SingerContractFilterer) WatchVerifyMulSignature(opts *bind.WatchOpts, sink chan<- *SingerContractVerifyMulSignature) (event.Subscription, error) {

	logs, sub, err := _SingerContract.contract.WatchLogs(opts, "verifyMulSignature")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SingerContractVerifyMulSignature)
				if err := _SingerContract.contract.UnpackLog(event, "verifyMulSignature", log); err != nil {
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

// ParseVerifyMulSignature is a log parse operation binding the contract event 0x62257d4dafbcc200d0ab7688037ac0089af1ad7c87181f1df22a3b0f31d26263.
//
// Solidity: event verifyMulSignature()
func (_SingerContract *SingerContractFilterer) ParseVerifyMulSignature(log types.Log) (*SingerContractVerifyMulSignature, error) {
	event := new(SingerContractVerifyMulSignature)
	if err := _SingerContract.contract.UnpackLog(event, "verifyMulSignature", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
