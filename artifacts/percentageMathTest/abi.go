// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package percentageMathTest

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

// PercentageMathTestMetaData contains all meta data concerning the PercentageMathTest contract.
var PercentageMathTestMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"percentage\",\"type\":\"uint256\"}],\"name\":\"percentDiv\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"percentage\",\"type\":\"uint256\"}],\"name\":\"percentMul\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// PercentageMathTestABI is the input ABI used to generate the binding from.
// Deprecated: Use PercentageMathTestMetaData.ABI instead.
var PercentageMathTestABI = PercentageMathTestMetaData.ABI

// PercentageMathTest is an auto generated Go binding around an Ethereum contract.
type PercentageMathTest struct {
	PercentageMathTestCaller     // Read-only binding to the contract
	PercentageMathTestTransactor // Write-only binding to the contract
	PercentageMathTestFilterer   // Log filterer for contract events
}

// PercentageMathTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type PercentageMathTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PercentageMathTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PercentageMathTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PercentageMathTestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PercentageMathTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PercentageMathTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PercentageMathTestSession struct {
	Contract     *PercentageMathTest // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PercentageMathTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PercentageMathTestCallerSession struct {
	Contract *PercentageMathTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// PercentageMathTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PercentageMathTestTransactorSession struct {
	Contract     *PercentageMathTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// PercentageMathTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type PercentageMathTestRaw struct {
	Contract *PercentageMathTest // Generic contract binding to access the raw methods on
}

// PercentageMathTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PercentageMathTestCallerRaw struct {
	Contract *PercentageMathTestCaller // Generic read-only contract binding to access the raw methods on
}

// PercentageMathTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PercentageMathTestTransactorRaw struct {
	Contract *PercentageMathTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPercentageMathTest creates a new instance of PercentageMathTest, bound to a specific deployed contract.
func NewPercentageMathTest(address common.Address, backend bind.ContractBackend) (*PercentageMathTest, error) {
	contract, err := bindPercentageMathTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PercentageMathTest{PercentageMathTestCaller: PercentageMathTestCaller{contract: contract}, PercentageMathTestTransactor: PercentageMathTestTransactor{contract: contract}, PercentageMathTestFilterer: PercentageMathTestFilterer{contract: contract}}, nil
}

// NewPercentageMathTestCaller creates a new read-only instance of PercentageMathTest, bound to a specific deployed contract.
func NewPercentageMathTestCaller(address common.Address, caller bind.ContractCaller) (*PercentageMathTestCaller, error) {
	contract, err := bindPercentageMathTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PercentageMathTestCaller{contract: contract}, nil
}

// NewPercentageMathTestTransactor creates a new write-only instance of PercentageMathTest, bound to a specific deployed contract.
func NewPercentageMathTestTransactor(address common.Address, transactor bind.ContractTransactor) (*PercentageMathTestTransactor, error) {
	contract, err := bindPercentageMathTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PercentageMathTestTransactor{contract: contract}, nil
}

// NewPercentageMathTestFilterer creates a new log filterer instance of PercentageMathTest, bound to a specific deployed contract.
func NewPercentageMathTestFilterer(address common.Address, filterer bind.ContractFilterer) (*PercentageMathTestFilterer, error) {
	contract, err := bindPercentageMathTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PercentageMathTestFilterer{contract: contract}, nil
}

// bindPercentageMathTest binds a generic wrapper to an already deployed contract.
func bindPercentageMathTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PercentageMathTestABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PercentageMathTest *PercentageMathTestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PercentageMathTest.Contract.PercentageMathTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PercentageMathTest *PercentageMathTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PercentageMathTest.Contract.PercentageMathTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PercentageMathTest *PercentageMathTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PercentageMathTest.Contract.PercentageMathTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PercentageMathTest *PercentageMathTestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PercentageMathTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PercentageMathTest *PercentageMathTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PercentageMathTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PercentageMathTest *PercentageMathTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PercentageMathTest.Contract.contract.Transact(opts, method, params...)
}

// PercentDiv is a free data retrieval call binding the contract method 0x46c840bb.
//
// Solidity: function percentDiv(uint256 value, uint256 percentage) pure returns(uint256)
func (_PercentageMathTest *PercentageMathTestCaller) PercentDiv(opts *bind.CallOpts, value *big.Int, percentage *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PercentageMathTest.contract.Call(opts, &out, "percentDiv", value, percentage)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PercentDiv is a free data retrieval call binding the contract method 0x46c840bb.
//
// Solidity: function percentDiv(uint256 value, uint256 percentage) pure returns(uint256)
func (_PercentageMathTest *PercentageMathTestSession) PercentDiv(value *big.Int, percentage *big.Int) (*big.Int, error) {
	return _PercentageMathTest.Contract.PercentDiv(&_PercentageMathTest.CallOpts, value, percentage)
}

// PercentDiv is a free data retrieval call binding the contract method 0x46c840bb.
//
// Solidity: function percentDiv(uint256 value, uint256 percentage) pure returns(uint256)
func (_PercentageMathTest *PercentageMathTestCallerSession) PercentDiv(value *big.Int, percentage *big.Int) (*big.Int, error) {
	return _PercentageMathTest.Contract.PercentDiv(&_PercentageMathTest.CallOpts, value, percentage)
}

// PercentMul is a free data retrieval call binding the contract method 0x4bf6a8f0.
//
// Solidity: function percentMul(uint256 value, uint256 percentage) pure returns(uint256)
func (_PercentageMathTest *PercentageMathTestCaller) PercentMul(opts *bind.CallOpts, value *big.Int, percentage *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PercentageMathTest.contract.Call(opts, &out, "percentMul", value, percentage)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PercentMul is a free data retrieval call binding the contract method 0x4bf6a8f0.
//
// Solidity: function percentMul(uint256 value, uint256 percentage) pure returns(uint256)
func (_PercentageMathTest *PercentageMathTestSession) PercentMul(value *big.Int, percentage *big.Int) (*big.Int, error) {
	return _PercentageMathTest.Contract.PercentMul(&_PercentageMathTest.CallOpts, value, percentage)
}

// PercentMul is a free data retrieval call binding the contract method 0x4bf6a8f0.
//
// Solidity: function percentMul(uint256 value, uint256 percentage) pure returns(uint256)
func (_PercentageMathTest *PercentageMathTestCallerSession) PercentMul(value *big.Int, percentage *big.Int) (*big.Int, error) {
	return _PercentageMathTest.Contract.PercentMul(&_PercentageMathTest.CallOpts, value, percentage)
}
