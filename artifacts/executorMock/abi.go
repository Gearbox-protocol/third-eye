// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package executorMock

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
)

// ExecutorMockMetaData contains all meta data concerning the ExecutorMock contract.
var ExecutorMockMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"calledBy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"setValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"value\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ExecutorMockABI is the input ABI used to generate the binding from.
// Deprecated: Use ExecutorMockMetaData.ABI instead.
var ExecutorMockABI = ExecutorMockMetaData.ABI

// ExecutorMock is an auto generated Go binding around an Ethereum contract.
type ExecutorMock struct {
	ExecutorMockCaller     // Read-only binding to the contract
	ExecutorMockTransactor // Write-only binding to the contract
	ExecutorMockFilterer   // Log filterer for contract events
}

// ExecutorMockCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExecutorMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExecutorMockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExecutorMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExecutorMockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExecutorMockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExecutorMockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExecutorMockSession struct {
	Contract     *ExecutorMock     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExecutorMockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExecutorMockCallerSession struct {
	Contract *ExecutorMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ExecutorMockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExecutorMockTransactorSession struct {
	Contract     *ExecutorMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ExecutorMockRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExecutorMockRaw struct {
	Contract *ExecutorMock // Generic contract binding to access the raw methods on
}

// ExecutorMockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExecutorMockCallerRaw struct {
	Contract *ExecutorMockCaller // Generic read-only contract binding to access the raw methods on
}

// ExecutorMockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExecutorMockTransactorRaw struct {
	Contract *ExecutorMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExecutorMock creates a new instance of ExecutorMock, bound to a specific deployed contract.
func NewExecutorMock(address common.Address, backend bind.ContractBackend) (*ExecutorMock, error) {
	contract, err := bindExecutorMock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExecutorMock{ExecutorMockCaller: ExecutorMockCaller{contract: contract}, ExecutorMockTransactor: ExecutorMockTransactor{contract: contract}, ExecutorMockFilterer: ExecutorMockFilterer{contract: contract}}, nil
}

// NewExecutorMockCaller creates a new read-only instance of ExecutorMock, bound to a specific deployed contract.
func NewExecutorMockCaller(address common.Address, caller bind.ContractCaller) (*ExecutorMockCaller, error) {
	contract, err := bindExecutorMock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExecutorMockCaller{contract: contract}, nil
}

// NewExecutorMockTransactor creates a new write-only instance of ExecutorMock, bound to a specific deployed contract.
func NewExecutorMockTransactor(address common.Address, transactor bind.ContractTransactor) (*ExecutorMockTransactor, error) {
	contract, err := bindExecutorMock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExecutorMockTransactor{contract: contract}, nil
}

// NewExecutorMockFilterer creates a new log filterer instance of ExecutorMock, bound to a specific deployed contract.
func NewExecutorMockFilterer(address common.Address, filterer bind.ContractFilterer) (*ExecutorMockFilterer, error) {
	contract, err := bindExecutorMock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExecutorMockFilterer{contract: contract}, nil
}

// bindExecutorMock binds a generic wrapper to an already deployed contract.
func bindExecutorMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ExecutorMockABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExecutorMock *ExecutorMockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExecutorMock.Contract.ExecutorMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExecutorMock *ExecutorMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExecutorMock.Contract.ExecutorMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExecutorMock *ExecutorMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExecutorMock.Contract.ExecutorMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExecutorMock *ExecutorMockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExecutorMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExecutorMock *ExecutorMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExecutorMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExecutorMock *ExecutorMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExecutorMock.Contract.contract.Transact(opts, method, params...)
}

// CalledBy is a free data retrieval call binding the contract method 0x51ae2a67.
//
// Solidity: function calledBy() view returns(address)
func (_ExecutorMock *ExecutorMockCaller) CalledBy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExecutorMock.contract.Call(opts, &out, "calledBy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CalledBy is a free data retrieval call binding the contract method 0x51ae2a67.
//
// Solidity: function calledBy() view returns(address)
func (_ExecutorMock *ExecutorMockSession) CalledBy() (common.Address, error) {
	return _ExecutorMock.Contract.CalledBy(&_ExecutorMock.CallOpts)
}

// CalledBy is a free data retrieval call binding the contract method 0x51ae2a67.
//
// Solidity: function calledBy() view returns(address)
func (_ExecutorMock *ExecutorMockCallerSession) CalledBy() (common.Address, error) {
	return _ExecutorMock.Contract.CalledBy(&_ExecutorMock.CallOpts)
}

// Value is a free data retrieval call binding the contract method 0x3fa4f245.
//
// Solidity: function value() view returns(uint256)
func (_ExecutorMock *ExecutorMockCaller) Value(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ExecutorMock.contract.Call(opts, &out, "value")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Value is a free data retrieval call binding the contract method 0x3fa4f245.
//
// Solidity: function value() view returns(uint256)
func (_ExecutorMock *ExecutorMockSession) Value() (*big.Int, error) {
	return _ExecutorMock.Contract.Value(&_ExecutorMock.CallOpts)
}

// Value is a free data retrieval call binding the contract method 0x3fa4f245.
//
// Solidity: function value() view returns(uint256)
func (_ExecutorMock *ExecutorMockCallerSession) Value() (*big.Int, error) {
	return _ExecutorMock.Contract.Value(&_ExecutorMock.CallOpts)
}

// SetValue is a paid mutator transaction binding the contract method 0x55241077.
//
// Solidity: function setValue(uint256 _value) returns(uint256)
func (_ExecutorMock *ExecutorMockTransactor) SetValue(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _ExecutorMock.contract.Transact(opts, "setValue", _value)
}

// SetValue is a paid mutator transaction binding the contract method 0x55241077.
//
// Solidity: function setValue(uint256 _value) returns(uint256)
func (_ExecutorMock *ExecutorMockSession) SetValue(_value *big.Int) (*types.Transaction, error) {
	return _ExecutorMock.Contract.SetValue(&_ExecutorMock.TransactOpts, _value)
}

// SetValue is a paid mutator transaction binding the contract method 0x55241077.
//
// Solidity: function setValue(uint256 _value) returns(uint256)
func (_ExecutorMock *ExecutorMockTransactorSession) SetValue(_value *big.Int) (*types.Transaction, error) {
	return _ExecutorMock.Contract.SetValue(&_ExecutorMock.TransactOpts, _value)
}
