// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iAppPoolService

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

// IAppPoolServiceMetaData contains all meta data concerning the IAppPoolService contract.
var IAppPoolServiceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"referralCode\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"removeLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IAppPoolServiceABI is the input ABI used to generate the binding from.
// Deprecated: Use IAppPoolServiceMetaData.ABI instead.
var IAppPoolServiceABI = IAppPoolServiceMetaData.ABI

// IAppPoolService is an auto generated Go binding around an Ethereum contract.
type IAppPoolService struct {
	IAppPoolServiceCaller     // Read-only binding to the contract
	IAppPoolServiceTransactor // Write-only binding to the contract
	IAppPoolServiceFilterer   // Log filterer for contract events
}

// IAppPoolServiceCaller is an auto generated read-only Go binding around an Ethereum contract.
type IAppPoolServiceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAppPoolServiceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IAppPoolServiceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAppPoolServiceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAppPoolServiceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAppPoolServiceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IAppPoolServiceSession struct {
	Contract     *IAppPoolService  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAppPoolServiceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IAppPoolServiceCallerSession struct {
	Contract *IAppPoolServiceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IAppPoolServiceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IAppPoolServiceTransactorSession struct {
	Contract     *IAppPoolServiceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IAppPoolServiceRaw is an auto generated low-level Go binding around an Ethereum contract.
type IAppPoolServiceRaw struct {
	Contract *IAppPoolService // Generic contract binding to access the raw methods on
}

// IAppPoolServiceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IAppPoolServiceCallerRaw struct {
	Contract *IAppPoolServiceCaller // Generic read-only contract binding to access the raw methods on
}

// IAppPoolServiceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IAppPoolServiceTransactorRaw struct {
	Contract *IAppPoolServiceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIAppPoolService creates a new instance of IAppPoolService, bound to a specific deployed contract.
func NewIAppPoolService(address common.Address, backend bind.ContractBackend) (*IAppPoolService, error) {
	contract, err := bindIAppPoolService(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAppPoolService{IAppPoolServiceCaller: IAppPoolServiceCaller{contract: contract}, IAppPoolServiceTransactor: IAppPoolServiceTransactor{contract: contract}, IAppPoolServiceFilterer: IAppPoolServiceFilterer{contract: contract}}, nil
}

// NewIAppPoolServiceCaller creates a new read-only instance of IAppPoolService, bound to a specific deployed contract.
func NewIAppPoolServiceCaller(address common.Address, caller bind.ContractCaller) (*IAppPoolServiceCaller, error) {
	contract, err := bindIAppPoolService(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAppPoolServiceCaller{contract: contract}, nil
}

// NewIAppPoolServiceTransactor creates a new write-only instance of IAppPoolService, bound to a specific deployed contract.
func NewIAppPoolServiceTransactor(address common.Address, transactor bind.ContractTransactor) (*IAppPoolServiceTransactor, error) {
	contract, err := bindIAppPoolService(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAppPoolServiceTransactor{contract: contract}, nil
}

// NewIAppPoolServiceFilterer creates a new log filterer instance of IAppPoolService, bound to a specific deployed contract.
func NewIAppPoolServiceFilterer(address common.Address, filterer bind.ContractFilterer) (*IAppPoolServiceFilterer, error) {
	contract, err := bindIAppPoolService(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAppPoolServiceFilterer{contract: contract}, nil
}

// bindIAppPoolService binds a generic wrapper to an already deployed contract.
func bindIAppPoolService(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IAppPoolServiceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAppPoolService *IAppPoolServiceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAppPoolService.Contract.IAppPoolServiceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAppPoolService *IAppPoolServiceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAppPoolService.Contract.IAppPoolServiceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAppPoolService *IAppPoolServiceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAppPoolService.Contract.IAppPoolServiceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAppPoolService *IAppPoolServiceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAppPoolService.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAppPoolService *IAppPoolServiceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAppPoolService.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAppPoolService *IAppPoolServiceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAppPoolService.Contract.contract.Transact(opts, method, params...)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x9aa5d462.
//
// Solidity: function addLiquidity(uint256 amount, address onBehalfOf, uint256 referralCode) returns()
func (_IAppPoolService *IAppPoolServiceTransactor) AddLiquidity(opts *bind.TransactOpts, amount *big.Int, onBehalfOf common.Address, referralCode *big.Int) (*types.Transaction, error) {
	return _IAppPoolService.contract.Transact(opts, "addLiquidity", amount, onBehalfOf, referralCode)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x9aa5d462.
//
// Solidity: function addLiquidity(uint256 amount, address onBehalfOf, uint256 referralCode) returns()
func (_IAppPoolService *IAppPoolServiceSession) AddLiquidity(amount *big.Int, onBehalfOf common.Address, referralCode *big.Int) (*types.Transaction, error) {
	return _IAppPoolService.Contract.AddLiquidity(&_IAppPoolService.TransactOpts, amount, onBehalfOf, referralCode)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x9aa5d462.
//
// Solidity: function addLiquidity(uint256 amount, address onBehalfOf, uint256 referralCode) returns()
func (_IAppPoolService *IAppPoolServiceTransactorSession) AddLiquidity(amount *big.Int, onBehalfOf common.Address, referralCode *big.Int) (*types.Transaction, error) {
	return _IAppPoolService.Contract.AddLiquidity(&_IAppPoolService.TransactOpts, amount, onBehalfOf, referralCode)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x05fe138b.
//
// Solidity: function removeLiquidity(uint256 amount, address to) returns(uint256)
func (_IAppPoolService *IAppPoolServiceTransactor) RemoveLiquidity(opts *bind.TransactOpts, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _IAppPoolService.contract.Transact(opts, "removeLiquidity", amount, to)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x05fe138b.
//
// Solidity: function removeLiquidity(uint256 amount, address to) returns(uint256)
func (_IAppPoolService *IAppPoolServiceSession) RemoveLiquidity(amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _IAppPoolService.Contract.RemoveLiquidity(&_IAppPoolService.TransactOpts, amount, to)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x05fe138b.
//
// Solidity: function removeLiquidity(uint256 amount, address to) returns(uint256)
func (_IAppPoolService *IAppPoolServiceTransactorSession) RemoveLiquidity(amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _IAppPoolService.Contract.RemoveLiquidity(&_IAppPoolService.TransactOpts, amount, to)
}
