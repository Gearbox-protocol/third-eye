// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package creditManagerMockForPoolTest

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

// CreditManagerMockForPoolTestMetaData contains all meta data concerning the CreditManagerMockForPoolTest contract.
var CreditManagerMockForPoolTestMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_poolService\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"borrowedAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"}],\"name\":\"lendCreditAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolService\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"borrowedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"profit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"loss\",\"type\":\"uint256\"}],\"name\":\"repayCreditAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlyingToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// CreditManagerMockForPoolTestABI is the input ABI used to generate the binding from.
// Deprecated: Use CreditManagerMockForPoolTestMetaData.ABI instead.
var CreditManagerMockForPoolTestABI = CreditManagerMockForPoolTestMetaData.ABI

// CreditManagerMockForPoolTest is an auto generated Go binding around an Ethereum contract.
type CreditManagerMockForPoolTest struct {
	CreditManagerMockForPoolTestCaller     // Read-only binding to the contract
	CreditManagerMockForPoolTestTransactor // Write-only binding to the contract
	CreditManagerMockForPoolTestFilterer   // Log filterer for contract events
}

// CreditManagerMockForPoolTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type CreditManagerMockForPoolTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditManagerMockForPoolTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CreditManagerMockForPoolTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditManagerMockForPoolTestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CreditManagerMockForPoolTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditManagerMockForPoolTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CreditManagerMockForPoolTestSession struct {
	Contract     *CreditManagerMockForPoolTest // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                 // Call options to use throughout this session
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// CreditManagerMockForPoolTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CreditManagerMockForPoolTestCallerSession struct {
	Contract *CreditManagerMockForPoolTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                       // Call options to use throughout this session
}

// CreditManagerMockForPoolTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CreditManagerMockForPoolTestTransactorSession struct {
	Contract     *CreditManagerMockForPoolTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                       // Transaction auth options to use throughout this session
}

// CreditManagerMockForPoolTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type CreditManagerMockForPoolTestRaw struct {
	Contract *CreditManagerMockForPoolTest // Generic contract binding to access the raw methods on
}

// CreditManagerMockForPoolTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CreditManagerMockForPoolTestCallerRaw struct {
	Contract *CreditManagerMockForPoolTestCaller // Generic read-only contract binding to access the raw methods on
}

// CreditManagerMockForPoolTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CreditManagerMockForPoolTestTransactorRaw struct {
	Contract *CreditManagerMockForPoolTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCreditManagerMockForPoolTest creates a new instance of CreditManagerMockForPoolTest, bound to a specific deployed contract.
func NewCreditManagerMockForPoolTest(address common.Address, backend bind.ContractBackend) (*CreditManagerMockForPoolTest, error) {
	contract, err := bindCreditManagerMockForPoolTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CreditManagerMockForPoolTest{CreditManagerMockForPoolTestCaller: CreditManagerMockForPoolTestCaller{contract: contract}, CreditManagerMockForPoolTestTransactor: CreditManagerMockForPoolTestTransactor{contract: contract}, CreditManagerMockForPoolTestFilterer: CreditManagerMockForPoolTestFilterer{contract: contract}}, nil
}

// NewCreditManagerMockForPoolTestCaller creates a new read-only instance of CreditManagerMockForPoolTest, bound to a specific deployed contract.
func NewCreditManagerMockForPoolTestCaller(address common.Address, caller bind.ContractCaller) (*CreditManagerMockForPoolTestCaller, error) {
	contract, err := bindCreditManagerMockForPoolTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CreditManagerMockForPoolTestCaller{contract: contract}, nil
}

// NewCreditManagerMockForPoolTestTransactor creates a new write-only instance of CreditManagerMockForPoolTest, bound to a specific deployed contract.
func NewCreditManagerMockForPoolTestTransactor(address common.Address, transactor bind.ContractTransactor) (*CreditManagerMockForPoolTestTransactor, error) {
	contract, err := bindCreditManagerMockForPoolTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CreditManagerMockForPoolTestTransactor{contract: contract}, nil
}

// NewCreditManagerMockForPoolTestFilterer creates a new log filterer instance of CreditManagerMockForPoolTest, bound to a specific deployed contract.
func NewCreditManagerMockForPoolTestFilterer(address common.Address, filterer bind.ContractFilterer) (*CreditManagerMockForPoolTestFilterer, error) {
	contract, err := bindCreditManagerMockForPoolTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CreditManagerMockForPoolTestFilterer{contract: contract}, nil
}

// bindCreditManagerMockForPoolTest binds a generic wrapper to an already deployed contract.
func bindCreditManagerMockForPoolTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CreditManagerMockForPoolTestABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CreditManagerMockForPoolTest *CreditManagerMockForPoolTestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CreditManagerMockForPoolTest.Contract.CreditManagerMockForPoolTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CreditManagerMockForPoolTest *CreditManagerMockForPoolTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditManagerMockForPoolTest.Contract.CreditManagerMockForPoolTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CreditManagerMockForPoolTest *CreditManagerMockForPoolTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CreditManagerMockForPoolTest.Contract.CreditManagerMockForPoolTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CreditManagerMockForPoolTest *CreditManagerMockForPoolTestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CreditManagerMockForPoolTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CreditManagerMockForPoolTest *CreditManagerMockForPoolTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditManagerMockForPoolTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CreditManagerMockForPoolTest *CreditManagerMockForPoolTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CreditManagerMockForPoolTest.Contract.contract.Transact(opts, method, params...)
}

// PoolService is a free data retrieval call binding the contract method 0x570a7af2.
//
// Solidity: function poolService() view returns(address)
func (_CreditManagerMockForPoolTest *CreditManagerMockForPoolTestCaller) PoolService(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditManagerMockForPoolTest.contract.Call(opts, &out, "poolService")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolService is a free data retrieval call binding the contract method 0x570a7af2.
//
// Solidity: function poolService() view returns(address)
func (_CreditManagerMockForPoolTest *CreditManagerMockForPoolTestSession) PoolService() (common.Address, error) {
	return _CreditManagerMockForPoolTest.Contract.PoolService(&_CreditManagerMockForPoolTest.CallOpts)
}

// PoolService is a free data retrieval call binding the contract method 0x570a7af2.
//
// Solidity: function poolService() view returns(address)
func (_CreditManagerMockForPoolTest *CreditManagerMockForPoolTestCallerSession) PoolService() (common.Address, error) {
	return _CreditManagerMockForPoolTest.Contract.PoolService(&_CreditManagerMockForPoolTest.CallOpts)
}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_CreditManagerMockForPoolTest *CreditManagerMockForPoolTestCaller) UnderlyingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditManagerMockForPoolTest.contract.Call(opts, &out, "underlyingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_CreditManagerMockForPoolTest *CreditManagerMockForPoolTestSession) UnderlyingToken() (common.Address, error) {
	return _CreditManagerMockForPoolTest.Contract.UnderlyingToken(&_CreditManagerMockForPoolTest.CallOpts)
}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_CreditManagerMockForPoolTest *CreditManagerMockForPoolTestCallerSession) UnderlyingToken() (common.Address, error) {
	return _CreditManagerMockForPoolTest.Contract.UnderlyingToken(&_CreditManagerMockForPoolTest.CallOpts)
}

// LendCreditAccount is a paid mutator transaction binding the contract method 0xbf28068b.
//
// Solidity: function lendCreditAccount(uint256 borrowedAmount, address creditAccount) returns()
func (_CreditManagerMockForPoolTest *CreditManagerMockForPoolTestTransactor) LendCreditAccount(opts *bind.TransactOpts, borrowedAmount *big.Int, creditAccount common.Address) (*types.Transaction, error) {
	return _CreditManagerMockForPoolTest.contract.Transact(opts, "lendCreditAccount", borrowedAmount, creditAccount)
}

// LendCreditAccount is a paid mutator transaction binding the contract method 0xbf28068b.
//
// Solidity: function lendCreditAccount(uint256 borrowedAmount, address creditAccount) returns()
func (_CreditManagerMockForPoolTest *CreditManagerMockForPoolTestSession) LendCreditAccount(borrowedAmount *big.Int, creditAccount common.Address) (*types.Transaction, error) {
	return _CreditManagerMockForPoolTest.Contract.LendCreditAccount(&_CreditManagerMockForPoolTest.TransactOpts, borrowedAmount, creditAccount)
}

// LendCreditAccount is a paid mutator transaction binding the contract method 0xbf28068b.
//
// Solidity: function lendCreditAccount(uint256 borrowedAmount, address creditAccount) returns()
func (_CreditManagerMockForPoolTest *CreditManagerMockForPoolTestTransactorSession) LendCreditAccount(borrowedAmount *big.Int, creditAccount common.Address) (*types.Transaction, error) {
	return _CreditManagerMockForPoolTest.Contract.LendCreditAccount(&_CreditManagerMockForPoolTest.TransactOpts, borrowedAmount, creditAccount)
}

// RepayCreditAccount is a paid mutator transaction binding the contract method 0xca9505e4.
//
// Solidity: function repayCreditAccount(uint256 borrowedAmount, uint256 profit, uint256 loss) returns()
func (_CreditManagerMockForPoolTest *CreditManagerMockForPoolTestTransactor) RepayCreditAccount(opts *bind.TransactOpts, borrowedAmount *big.Int, profit *big.Int, loss *big.Int) (*types.Transaction, error) {
	return _CreditManagerMockForPoolTest.contract.Transact(opts, "repayCreditAccount", borrowedAmount, profit, loss)
}

// RepayCreditAccount is a paid mutator transaction binding the contract method 0xca9505e4.
//
// Solidity: function repayCreditAccount(uint256 borrowedAmount, uint256 profit, uint256 loss) returns()
func (_CreditManagerMockForPoolTest *CreditManagerMockForPoolTestSession) RepayCreditAccount(borrowedAmount *big.Int, profit *big.Int, loss *big.Int) (*types.Transaction, error) {
	return _CreditManagerMockForPoolTest.Contract.RepayCreditAccount(&_CreditManagerMockForPoolTest.TransactOpts, borrowedAmount, profit, loss)
}

// RepayCreditAccount is a paid mutator transaction binding the contract method 0xca9505e4.
//
// Solidity: function repayCreditAccount(uint256 borrowedAmount, uint256 profit, uint256 loss) returns()
func (_CreditManagerMockForPoolTest *CreditManagerMockForPoolTestTransactorSession) RepayCreditAccount(borrowedAmount *big.Int, profit *big.Int, loss *big.Int) (*types.Transaction, error) {
	return _CreditManagerMockForPoolTest.Contract.RepayCreditAccount(&_CreditManagerMockForPoolTest.TransactOpts, borrowedAmount, profit, loss)
}
