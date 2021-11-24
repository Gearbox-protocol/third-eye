// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iUniswapV2Migrator

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

// IUniswapV2MigratorMetaData contains all meta data concerning the IUniswapV2Migrator contract.
var IUniswapV2MigratorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"migrate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IUniswapV2MigratorABI is the input ABI used to generate the binding from.
// Deprecated: Use IUniswapV2MigratorMetaData.ABI instead.
var IUniswapV2MigratorABI = IUniswapV2MigratorMetaData.ABI

// IUniswapV2Migrator is an auto generated Go binding around an Ethereum contract.
type IUniswapV2Migrator struct {
	IUniswapV2MigratorCaller     // Read-only binding to the contract
	IUniswapV2MigratorTransactor // Write-only binding to the contract
	IUniswapV2MigratorFilterer   // Log filterer for contract events
}

// IUniswapV2MigratorCaller is an auto generated read-only Go binding around an Ethereum contract.
type IUniswapV2MigratorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUniswapV2MigratorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IUniswapV2MigratorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUniswapV2MigratorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IUniswapV2MigratorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUniswapV2MigratorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IUniswapV2MigratorSession struct {
	Contract     *IUniswapV2Migrator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IUniswapV2MigratorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IUniswapV2MigratorCallerSession struct {
	Contract *IUniswapV2MigratorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// IUniswapV2MigratorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IUniswapV2MigratorTransactorSession struct {
	Contract     *IUniswapV2MigratorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// IUniswapV2MigratorRaw is an auto generated low-level Go binding around an Ethereum contract.
type IUniswapV2MigratorRaw struct {
	Contract *IUniswapV2Migrator // Generic contract binding to access the raw methods on
}

// IUniswapV2MigratorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IUniswapV2MigratorCallerRaw struct {
	Contract *IUniswapV2MigratorCaller // Generic read-only contract binding to access the raw methods on
}

// IUniswapV2MigratorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IUniswapV2MigratorTransactorRaw struct {
	Contract *IUniswapV2MigratorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIUniswapV2Migrator creates a new instance of IUniswapV2Migrator, bound to a specific deployed contract.
func NewIUniswapV2Migrator(address common.Address, backend bind.ContractBackend) (*IUniswapV2Migrator, error) {
	contract, err := bindIUniswapV2Migrator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IUniswapV2Migrator{IUniswapV2MigratorCaller: IUniswapV2MigratorCaller{contract: contract}, IUniswapV2MigratorTransactor: IUniswapV2MigratorTransactor{contract: contract}, IUniswapV2MigratorFilterer: IUniswapV2MigratorFilterer{contract: contract}}, nil
}

// NewIUniswapV2MigratorCaller creates a new read-only instance of IUniswapV2Migrator, bound to a specific deployed contract.
func NewIUniswapV2MigratorCaller(address common.Address, caller bind.ContractCaller) (*IUniswapV2MigratorCaller, error) {
	contract, err := bindIUniswapV2Migrator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IUniswapV2MigratorCaller{contract: contract}, nil
}

// NewIUniswapV2MigratorTransactor creates a new write-only instance of IUniswapV2Migrator, bound to a specific deployed contract.
func NewIUniswapV2MigratorTransactor(address common.Address, transactor bind.ContractTransactor) (*IUniswapV2MigratorTransactor, error) {
	contract, err := bindIUniswapV2Migrator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IUniswapV2MigratorTransactor{contract: contract}, nil
}

// NewIUniswapV2MigratorFilterer creates a new log filterer instance of IUniswapV2Migrator, bound to a specific deployed contract.
func NewIUniswapV2MigratorFilterer(address common.Address, filterer bind.ContractFilterer) (*IUniswapV2MigratorFilterer, error) {
	contract, err := bindIUniswapV2Migrator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IUniswapV2MigratorFilterer{contract: contract}, nil
}

// bindIUniswapV2Migrator binds a generic wrapper to an already deployed contract.
func bindIUniswapV2Migrator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IUniswapV2MigratorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IUniswapV2Migrator *IUniswapV2MigratorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IUniswapV2Migrator.Contract.IUniswapV2MigratorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IUniswapV2Migrator *IUniswapV2MigratorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUniswapV2Migrator.Contract.IUniswapV2MigratorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IUniswapV2Migrator *IUniswapV2MigratorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IUniswapV2Migrator.Contract.IUniswapV2MigratorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IUniswapV2Migrator *IUniswapV2MigratorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IUniswapV2Migrator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IUniswapV2Migrator *IUniswapV2MigratorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUniswapV2Migrator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IUniswapV2Migrator *IUniswapV2MigratorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IUniswapV2Migrator.Contract.contract.Transact(opts, method, params...)
}

// Migrate is a paid mutator transaction binding the contract method 0xb7df1d25.
//
// Solidity: function migrate(address token, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns()
func (_IUniswapV2Migrator *IUniswapV2MigratorTransactor) Migrate(opts *bind.TransactOpts, token common.Address, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapV2Migrator.contract.Transact(opts, "migrate", token, amountTokenMin, amountETHMin, to, deadline)
}

// Migrate is a paid mutator transaction binding the contract method 0xb7df1d25.
//
// Solidity: function migrate(address token, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns()
func (_IUniswapV2Migrator *IUniswapV2MigratorSession) Migrate(token common.Address, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapV2Migrator.Contract.Migrate(&_IUniswapV2Migrator.TransactOpts, token, amountTokenMin, amountETHMin, to, deadline)
}

// Migrate is a paid mutator transaction binding the contract method 0xb7df1d25.
//
// Solidity: function migrate(address token, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns()
func (_IUniswapV2Migrator *IUniswapV2MigratorTransactorSession) Migrate(token common.Address, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapV2Migrator.Contract.Migrate(&_IUniswapV2Migrator.TransactOpts, token, amountTokenMin, amountETHMin, to, deadline)
}
