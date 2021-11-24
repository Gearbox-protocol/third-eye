// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iWETHGateway

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

// IWETHGatewayMetaData contains all meta data concerning the IWETHGateway contract.
var IWETHGatewayMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"}],\"name\":\"addCollateralETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"referralCode\",\"type\":\"uint16\"}],\"name\":\"addLiquidityETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditManager\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"leverageFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"referralCode\",\"type\":\"uint256\"}],\"name\":\"openCreditAccountETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"removeLiquidityETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"repayCreditAccountETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"unwrapWETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IWETHGatewayABI is the input ABI used to generate the binding from.
// Deprecated: Use IWETHGatewayMetaData.ABI instead.
var IWETHGatewayABI = IWETHGatewayMetaData.ABI

// IWETHGateway is an auto generated Go binding around an Ethereum contract.
type IWETHGateway struct {
	IWETHGatewayCaller     // Read-only binding to the contract
	IWETHGatewayTransactor // Write-only binding to the contract
	IWETHGatewayFilterer   // Log filterer for contract events
}

// IWETHGatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type IWETHGatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWETHGatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IWETHGatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWETHGatewayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IWETHGatewayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWETHGatewaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IWETHGatewaySession struct {
	Contract     *IWETHGateway     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IWETHGatewayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IWETHGatewayCallerSession struct {
	Contract *IWETHGatewayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IWETHGatewayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IWETHGatewayTransactorSession struct {
	Contract     *IWETHGatewayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IWETHGatewayRaw is an auto generated low-level Go binding around an Ethereum contract.
type IWETHGatewayRaw struct {
	Contract *IWETHGateway // Generic contract binding to access the raw methods on
}

// IWETHGatewayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IWETHGatewayCallerRaw struct {
	Contract *IWETHGatewayCaller // Generic read-only contract binding to access the raw methods on
}

// IWETHGatewayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IWETHGatewayTransactorRaw struct {
	Contract *IWETHGatewayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIWETHGateway creates a new instance of IWETHGateway, bound to a specific deployed contract.
func NewIWETHGateway(address common.Address, backend bind.ContractBackend) (*IWETHGateway, error) {
	contract, err := bindIWETHGateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IWETHGateway{IWETHGatewayCaller: IWETHGatewayCaller{contract: contract}, IWETHGatewayTransactor: IWETHGatewayTransactor{contract: contract}, IWETHGatewayFilterer: IWETHGatewayFilterer{contract: contract}}, nil
}

// NewIWETHGatewayCaller creates a new read-only instance of IWETHGateway, bound to a specific deployed contract.
func NewIWETHGatewayCaller(address common.Address, caller bind.ContractCaller) (*IWETHGatewayCaller, error) {
	contract, err := bindIWETHGateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IWETHGatewayCaller{contract: contract}, nil
}

// NewIWETHGatewayTransactor creates a new write-only instance of IWETHGateway, bound to a specific deployed contract.
func NewIWETHGatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*IWETHGatewayTransactor, error) {
	contract, err := bindIWETHGateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IWETHGatewayTransactor{contract: contract}, nil
}

// NewIWETHGatewayFilterer creates a new log filterer instance of IWETHGateway, bound to a specific deployed contract.
func NewIWETHGatewayFilterer(address common.Address, filterer bind.ContractFilterer) (*IWETHGatewayFilterer, error) {
	contract, err := bindIWETHGateway(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IWETHGatewayFilterer{contract: contract}, nil
}

// bindIWETHGateway binds a generic wrapper to an already deployed contract.
func bindIWETHGateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IWETHGatewayABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IWETHGateway *IWETHGatewayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IWETHGateway.Contract.IWETHGatewayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IWETHGateway *IWETHGatewayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWETHGateway.Contract.IWETHGatewayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IWETHGateway *IWETHGatewayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IWETHGateway.Contract.IWETHGatewayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IWETHGateway *IWETHGatewayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IWETHGateway.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IWETHGateway *IWETHGatewayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWETHGateway.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IWETHGateway *IWETHGatewayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IWETHGateway.Contract.contract.Transact(opts, method, params...)
}

// AddCollateralETH is a paid mutator transaction binding the contract method 0x420ac3b2.
//
// Solidity: function addCollateralETH(address creditManager, address onBehalfOf) payable returns()
func (_IWETHGateway *IWETHGatewayTransactor) AddCollateralETH(opts *bind.TransactOpts, creditManager common.Address, onBehalfOf common.Address) (*types.Transaction, error) {
	return _IWETHGateway.contract.Transact(opts, "addCollateralETH", creditManager, onBehalfOf)
}

// AddCollateralETH is a paid mutator transaction binding the contract method 0x420ac3b2.
//
// Solidity: function addCollateralETH(address creditManager, address onBehalfOf) payable returns()
func (_IWETHGateway *IWETHGatewaySession) AddCollateralETH(creditManager common.Address, onBehalfOf common.Address) (*types.Transaction, error) {
	return _IWETHGateway.Contract.AddCollateralETH(&_IWETHGateway.TransactOpts, creditManager, onBehalfOf)
}

// AddCollateralETH is a paid mutator transaction binding the contract method 0x420ac3b2.
//
// Solidity: function addCollateralETH(address creditManager, address onBehalfOf) payable returns()
func (_IWETHGateway *IWETHGatewayTransactorSession) AddCollateralETH(creditManager common.Address, onBehalfOf common.Address) (*types.Transaction, error) {
	return _IWETHGateway.Contract.AddCollateralETH(&_IWETHGateway.TransactOpts, creditManager, onBehalfOf)
}

// AddLiquidityETH is a paid mutator transaction binding the contract method 0xdeecfbc9.
//
// Solidity: function addLiquidityETH(address pool, address onBehalfOf, uint16 referralCode) payable returns()
func (_IWETHGateway *IWETHGatewayTransactor) AddLiquidityETH(opts *bind.TransactOpts, pool common.Address, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _IWETHGateway.contract.Transact(opts, "addLiquidityETH", pool, onBehalfOf, referralCode)
}

// AddLiquidityETH is a paid mutator transaction binding the contract method 0xdeecfbc9.
//
// Solidity: function addLiquidityETH(address pool, address onBehalfOf, uint16 referralCode) payable returns()
func (_IWETHGateway *IWETHGatewaySession) AddLiquidityETH(pool common.Address, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _IWETHGateway.Contract.AddLiquidityETH(&_IWETHGateway.TransactOpts, pool, onBehalfOf, referralCode)
}

// AddLiquidityETH is a paid mutator transaction binding the contract method 0xdeecfbc9.
//
// Solidity: function addLiquidityETH(address pool, address onBehalfOf, uint16 referralCode) payable returns()
func (_IWETHGateway *IWETHGatewayTransactorSession) AddLiquidityETH(pool common.Address, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _IWETHGateway.Contract.AddLiquidityETH(&_IWETHGateway.TransactOpts, pool, onBehalfOf, referralCode)
}

// OpenCreditAccountETH is a paid mutator transaction binding the contract method 0xd8c99bc3.
//
// Solidity: function openCreditAccountETH(address creditManager, address onBehalfOf, uint256 leverageFactor, uint256 referralCode) payable returns()
func (_IWETHGateway *IWETHGatewayTransactor) OpenCreditAccountETH(opts *bind.TransactOpts, creditManager common.Address, onBehalfOf common.Address, leverageFactor *big.Int, referralCode *big.Int) (*types.Transaction, error) {
	return _IWETHGateway.contract.Transact(opts, "openCreditAccountETH", creditManager, onBehalfOf, leverageFactor, referralCode)
}

// OpenCreditAccountETH is a paid mutator transaction binding the contract method 0xd8c99bc3.
//
// Solidity: function openCreditAccountETH(address creditManager, address onBehalfOf, uint256 leverageFactor, uint256 referralCode) payable returns()
func (_IWETHGateway *IWETHGatewaySession) OpenCreditAccountETH(creditManager common.Address, onBehalfOf common.Address, leverageFactor *big.Int, referralCode *big.Int) (*types.Transaction, error) {
	return _IWETHGateway.Contract.OpenCreditAccountETH(&_IWETHGateway.TransactOpts, creditManager, onBehalfOf, leverageFactor, referralCode)
}

// OpenCreditAccountETH is a paid mutator transaction binding the contract method 0xd8c99bc3.
//
// Solidity: function openCreditAccountETH(address creditManager, address onBehalfOf, uint256 leverageFactor, uint256 referralCode) payable returns()
func (_IWETHGateway *IWETHGatewayTransactorSession) OpenCreditAccountETH(creditManager common.Address, onBehalfOf common.Address, leverageFactor *big.Int, referralCode *big.Int) (*types.Transaction, error) {
	return _IWETHGateway.Contract.OpenCreditAccountETH(&_IWETHGateway.TransactOpts, creditManager, onBehalfOf, leverageFactor, referralCode)
}

// RemoveLiquidityETH is a paid mutator transaction binding the contract method 0xe79a4089.
//
// Solidity: function removeLiquidityETH(address pool, uint256 amount, address to) returns()
func (_IWETHGateway *IWETHGatewayTransactor) RemoveLiquidityETH(opts *bind.TransactOpts, pool common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _IWETHGateway.contract.Transact(opts, "removeLiquidityETH", pool, amount, to)
}

// RemoveLiquidityETH is a paid mutator transaction binding the contract method 0xe79a4089.
//
// Solidity: function removeLiquidityETH(address pool, uint256 amount, address to) returns()
func (_IWETHGateway *IWETHGatewaySession) RemoveLiquidityETH(pool common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _IWETHGateway.Contract.RemoveLiquidityETH(&_IWETHGateway.TransactOpts, pool, amount, to)
}

// RemoveLiquidityETH is a paid mutator transaction binding the contract method 0xe79a4089.
//
// Solidity: function removeLiquidityETH(address pool, uint256 amount, address to) returns()
func (_IWETHGateway *IWETHGatewayTransactorSession) RemoveLiquidityETH(pool common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _IWETHGateway.Contract.RemoveLiquidityETH(&_IWETHGateway.TransactOpts, pool, amount, to)
}

// RepayCreditAccountETH is a paid mutator transaction binding the contract method 0xa6eab5c2.
//
// Solidity: function repayCreditAccountETH(address creditManager, address to) payable returns()
func (_IWETHGateway *IWETHGatewayTransactor) RepayCreditAccountETH(opts *bind.TransactOpts, creditManager common.Address, to common.Address) (*types.Transaction, error) {
	return _IWETHGateway.contract.Transact(opts, "repayCreditAccountETH", creditManager, to)
}

// RepayCreditAccountETH is a paid mutator transaction binding the contract method 0xa6eab5c2.
//
// Solidity: function repayCreditAccountETH(address creditManager, address to) payable returns()
func (_IWETHGateway *IWETHGatewaySession) RepayCreditAccountETH(creditManager common.Address, to common.Address) (*types.Transaction, error) {
	return _IWETHGateway.Contract.RepayCreditAccountETH(&_IWETHGateway.TransactOpts, creditManager, to)
}

// RepayCreditAccountETH is a paid mutator transaction binding the contract method 0xa6eab5c2.
//
// Solidity: function repayCreditAccountETH(address creditManager, address to) payable returns()
func (_IWETHGateway *IWETHGatewayTransactorSession) RepayCreditAccountETH(creditManager common.Address, to common.Address) (*types.Transaction, error) {
	return _IWETHGateway.Contract.RepayCreditAccountETH(&_IWETHGateway.TransactOpts, creditManager, to)
}

// UnwrapWETH is a paid mutator transaction binding the contract method 0x5869dba8.
//
// Solidity: function unwrapWETH(address to, uint256 amount) returns()
func (_IWETHGateway *IWETHGatewayTransactor) UnwrapWETH(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IWETHGateway.contract.Transact(opts, "unwrapWETH", to, amount)
}

// UnwrapWETH is a paid mutator transaction binding the contract method 0x5869dba8.
//
// Solidity: function unwrapWETH(address to, uint256 amount) returns()
func (_IWETHGateway *IWETHGatewaySession) UnwrapWETH(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IWETHGateway.Contract.UnwrapWETH(&_IWETHGateway.TransactOpts, to, amount)
}

// UnwrapWETH is a paid mutator transaction binding the contract method 0x5869dba8.
//
// Solidity: function unwrapWETH(address to, uint256 amount) returns()
func (_IWETHGateway *IWETHGatewayTransactorSession) UnwrapWETH(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IWETHGateway.Contract.UnwrapWETH(&_IWETHGateway.TransactOpts, to, amount)
}
