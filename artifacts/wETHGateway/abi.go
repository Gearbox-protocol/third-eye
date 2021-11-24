// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package wETHGateway

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

// WETHGatewayMetaData contains all meta data concerning the WETHGateway contract.
var WETHGatewayMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addressProvider\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"}],\"name\":\"addCollateralETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"referralCode\",\"type\":\"uint16\"}],\"name\":\"addLiquidityETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditManager\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"leverageFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"referralCode\",\"type\":\"uint256\"}],\"name\":\"openCreditAccountETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"removeLiquidityETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"repayCreditAccountETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"unwrapWETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wethAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// WETHGatewayABI is the input ABI used to generate the binding from.
// Deprecated: Use WETHGatewayMetaData.ABI instead.
var WETHGatewayABI = WETHGatewayMetaData.ABI

// WETHGateway is an auto generated Go binding around an Ethereum contract.
type WETHGateway struct {
	WETHGatewayCaller     // Read-only binding to the contract
	WETHGatewayTransactor // Write-only binding to the contract
	WETHGatewayFilterer   // Log filterer for contract events
}

// WETHGatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type WETHGatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WETHGatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WETHGatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WETHGatewayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WETHGatewayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WETHGatewaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WETHGatewaySession struct {
	Contract     *WETHGateway      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WETHGatewayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WETHGatewayCallerSession struct {
	Contract *WETHGatewayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// WETHGatewayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WETHGatewayTransactorSession struct {
	Contract     *WETHGatewayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// WETHGatewayRaw is an auto generated low-level Go binding around an Ethereum contract.
type WETHGatewayRaw struct {
	Contract *WETHGateway // Generic contract binding to access the raw methods on
}

// WETHGatewayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WETHGatewayCallerRaw struct {
	Contract *WETHGatewayCaller // Generic read-only contract binding to access the raw methods on
}

// WETHGatewayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WETHGatewayTransactorRaw struct {
	Contract *WETHGatewayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWETHGateway creates a new instance of WETHGateway, bound to a specific deployed contract.
func NewWETHGateway(address common.Address, backend bind.ContractBackend) (*WETHGateway, error) {
	contract, err := bindWETHGateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WETHGateway{WETHGatewayCaller: WETHGatewayCaller{contract: contract}, WETHGatewayTransactor: WETHGatewayTransactor{contract: contract}, WETHGatewayFilterer: WETHGatewayFilterer{contract: contract}}, nil
}

// NewWETHGatewayCaller creates a new read-only instance of WETHGateway, bound to a specific deployed contract.
func NewWETHGatewayCaller(address common.Address, caller bind.ContractCaller) (*WETHGatewayCaller, error) {
	contract, err := bindWETHGateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WETHGatewayCaller{contract: contract}, nil
}

// NewWETHGatewayTransactor creates a new write-only instance of WETHGateway, bound to a specific deployed contract.
func NewWETHGatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*WETHGatewayTransactor, error) {
	contract, err := bindWETHGateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WETHGatewayTransactor{contract: contract}, nil
}

// NewWETHGatewayFilterer creates a new log filterer instance of WETHGateway, bound to a specific deployed contract.
func NewWETHGatewayFilterer(address common.Address, filterer bind.ContractFilterer) (*WETHGatewayFilterer, error) {
	contract, err := bindWETHGateway(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WETHGatewayFilterer{contract: contract}, nil
}

// bindWETHGateway binds a generic wrapper to an already deployed contract.
func bindWETHGateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WETHGatewayABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WETHGateway *WETHGatewayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WETHGateway.Contract.WETHGatewayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WETHGateway *WETHGatewayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WETHGateway.Contract.WETHGatewayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WETHGateway *WETHGatewayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WETHGateway.Contract.WETHGatewayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WETHGateway *WETHGatewayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WETHGateway.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WETHGateway *WETHGatewayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WETHGateway.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WETHGateway *WETHGatewayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WETHGateway.Contract.contract.Transact(opts, method, params...)
}

// WethAddress is a free data retrieval call binding the contract method 0x4f0e0ef3.
//
// Solidity: function wethAddress() view returns(address)
func (_WETHGateway *WETHGatewayCaller) WethAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WETHGateway.contract.Call(opts, &out, "wethAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WethAddress is a free data retrieval call binding the contract method 0x4f0e0ef3.
//
// Solidity: function wethAddress() view returns(address)
func (_WETHGateway *WETHGatewaySession) WethAddress() (common.Address, error) {
	return _WETHGateway.Contract.WethAddress(&_WETHGateway.CallOpts)
}

// WethAddress is a free data retrieval call binding the contract method 0x4f0e0ef3.
//
// Solidity: function wethAddress() view returns(address)
func (_WETHGateway *WETHGatewayCallerSession) WethAddress() (common.Address, error) {
	return _WETHGateway.Contract.WethAddress(&_WETHGateway.CallOpts)
}

// AddCollateralETH is a paid mutator transaction binding the contract method 0x420ac3b2.
//
// Solidity: function addCollateralETH(address creditManager, address onBehalfOf) payable returns()
func (_WETHGateway *WETHGatewayTransactor) AddCollateralETH(opts *bind.TransactOpts, creditManager common.Address, onBehalfOf common.Address) (*types.Transaction, error) {
	return _WETHGateway.contract.Transact(opts, "addCollateralETH", creditManager, onBehalfOf)
}

// AddCollateralETH is a paid mutator transaction binding the contract method 0x420ac3b2.
//
// Solidity: function addCollateralETH(address creditManager, address onBehalfOf) payable returns()
func (_WETHGateway *WETHGatewaySession) AddCollateralETH(creditManager common.Address, onBehalfOf common.Address) (*types.Transaction, error) {
	return _WETHGateway.Contract.AddCollateralETH(&_WETHGateway.TransactOpts, creditManager, onBehalfOf)
}

// AddCollateralETH is a paid mutator transaction binding the contract method 0x420ac3b2.
//
// Solidity: function addCollateralETH(address creditManager, address onBehalfOf) payable returns()
func (_WETHGateway *WETHGatewayTransactorSession) AddCollateralETH(creditManager common.Address, onBehalfOf common.Address) (*types.Transaction, error) {
	return _WETHGateway.Contract.AddCollateralETH(&_WETHGateway.TransactOpts, creditManager, onBehalfOf)
}

// AddLiquidityETH is a paid mutator transaction binding the contract method 0xdeecfbc9.
//
// Solidity: function addLiquidityETH(address pool, address onBehalfOf, uint16 referralCode) payable returns()
func (_WETHGateway *WETHGatewayTransactor) AddLiquidityETH(opts *bind.TransactOpts, pool common.Address, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _WETHGateway.contract.Transact(opts, "addLiquidityETH", pool, onBehalfOf, referralCode)
}

// AddLiquidityETH is a paid mutator transaction binding the contract method 0xdeecfbc9.
//
// Solidity: function addLiquidityETH(address pool, address onBehalfOf, uint16 referralCode) payable returns()
func (_WETHGateway *WETHGatewaySession) AddLiquidityETH(pool common.Address, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _WETHGateway.Contract.AddLiquidityETH(&_WETHGateway.TransactOpts, pool, onBehalfOf, referralCode)
}

// AddLiquidityETH is a paid mutator transaction binding the contract method 0xdeecfbc9.
//
// Solidity: function addLiquidityETH(address pool, address onBehalfOf, uint16 referralCode) payable returns()
func (_WETHGateway *WETHGatewayTransactorSession) AddLiquidityETH(pool common.Address, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _WETHGateway.Contract.AddLiquidityETH(&_WETHGateway.TransactOpts, pool, onBehalfOf, referralCode)
}

// OpenCreditAccountETH is a paid mutator transaction binding the contract method 0xd8c99bc3.
//
// Solidity: function openCreditAccountETH(address creditManager, address onBehalfOf, uint256 leverageFactor, uint256 referralCode) payable returns()
func (_WETHGateway *WETHGatewayTransactor) OpenCreditAccountETH(opts *bind.TransactOpts, creditManager common.Address, onBehalfOf common.Address, leverageFactor *big.Int, referralCode *big.Int) (*types.Transaction, error) {
	return _WETHGateway.contract.Transact(opts, "openCreditAccountETH", creditManager, onBehalfOf, leverageFactor, referralCode)
}

// OpenCreditAccountETH is a paid mutator transaction binding the contract method 0xd8c99bc3.
//
// Solidity: function openCreditAccountETH(address creditManager, address onBehalfOf, uint256 leverageFactor, uint256 referralCode) payable returns()
func (_WETHGateway *WETHGatewaySession) OpenCreditAccountETH(creditManager common.Address, onBehalfOf common.Address, leverageFactor *big.Int, referralCode *big.Int) (*types.Transaction, error) {
	return _WETHGateway.Contract.OpenCreditAccountETH(&_WETHGateway.TransactOpts, creditManager, onBehalfOf, leverageFactor, referralCode)
}

// OpenCreditAccountETH is a paid mutator transaction binding the contract method 0xd8c99bc3.
//
// Solidity: function openCreditAccountETH(address creditManager, address onBehalfOf, uint256 leverageFactor, uint256 referralCode) payable returns()
func (_WETHGateway *WETHGatewayTransactorSession) OpenCreditAccountETH(creditManager common.Address, onBehalfOf common.Address, leverageFactor *big.Int, referralCode *big.Int) (*types.Transaction, error) {
	return _WETHGateway.Contract.OpenCreditAccountETH(&_WETHGateway.TransactOpts, creditManager, onBehalfOf, leverageFactor, referralCode)
}

// RemoveLiquidityETH is a paid mutator transaction binding the contract method 0xe79a4089.
//
// Solidity: function removeLiquidityETH(address pool, uint256 amount, address to) returns()
func (_WETHGateway *WETHGatewayTransactor) RemoveLiquidityETH(opts *bind.TransactOpts, pool common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _WETHGateway.contract.Transact(opts, "removeLiquidityETH", pool, amount, to)
}

// RemoveLiquidityETH is a paid mutator transaction binding the contract method 0xe79a4089.
//
// Solidity: function removeLiquidityETH(address pool, uint256 amount, address to) returns()
func (_WETHGateway *WETHGatewaySession) RemoveLiquidityETH(pool common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _WETHGateway.Contract.RemoveLiquidityETH(&_WETHGateway.TransactOpts, pool, amount, to)
}

// RemoveLiquidityETH is a paid mutator transaction binding the contract method 0xe79a4089.
//
// Solidity: function removeLiquidityETH(address pool, uint256 amount, address to) returns()
func (_WETHGateway *WETHGatewayTransactorSession) RemoveLiquidityETH(pool common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _WETHGateway.Contract.RemoveLiquidityETH(&_WETHGateway.TransactOpts, pool, amount, to)
}

// RepayCreditAccountETH is a paid mutator transaction binding the contract method 0xa6eab5c2.
//
// Solidity: function repayCreditAccountETH(address creditManager, address to) payable returns()
func (_WETHGateway *WETHGatewayTransactor) RepayCreditAccountETH(opts *bind.TransactOpts, creditManager common.Address, to common.Address) (*types.Transaction, error) {
	return _WETHGateway.contract.Transact(opts, "repayCreditAccountETH", creditManager, to)
}

// RepayCreditAccountETH is a paid mutator transaction binding the contract method 0xa6eab5c2.
//
// Solidity: function repayCreditAccountETH(address creditManager, address to) payable returns()
func (_WETHGateway *WETHGatewaySession) RepayCreditAccountETH(creditManager common.Address, to common.Address) (*types.Transaction, error) {
	return _WETHGateway.Contract.RepayCreditAccountETH(&_WETHGateway.TransactOpts, creditManager, to)
}

// RepayCreditAccountETH is a paid mutator transaction binding the contract method 0xa6eab5c2.
//
// Solidity: function repayCreditAccountETH(address creditManager, address to) payable returns()
func (_WETHGateway *WETHGatewayTransactorSession) RepayCreditAccountETH(creditManager common.Address, to common.Address) (*types.Transaction, error) {
	return _WETHGateway.Contract.RepayCreditAccountETH(&_WETHGateway.TransactOpts, creditManager, to)
}

// UnwrapWETH is a paid mutator transaction binding the contract method 0x5869dba8.
//
// Solidity: function unwrapWETH(address to, uint256 amount) returns()
func (_WETHGateway *WETHGatewayTransactor) UnwrapWETH(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WETHGateway.contract.Transact(opts, "unwrapWETH", to, amount)
}

// UnwrapWETH is a paid mutator transaction binding the contract method 0x5869dba8.
//
// Solidity: function unwrapWETH(address to, uint256 amount) returns()
func (_WETHGateway *WETHGatewaySession) UnwrapWETH(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WETHGateway.Contract.UnwrapWETH(&_WETHGateway.TransactOpts, to, amount)
}

// UnwrapWETH is a paid mutator transaction binding the contract method 0x5869dba8.
//
// Solidity: function unwrapWETH(address to, uint256 amount) returns()
func (_WETHGateway *WETHGatewayTransactorSession) UnwrapWETH(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WETHGateway.Contract.UnwrapWETH(&_WETHGateway.TransactOpts, to, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WETHGateway *WETHGatewayTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WETHGateway.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WETHGateway *WETHGatewaySession) Receive() (*types.Transaction, error) {
	return _WETHGateway.Contract.Receive(&_WETHGateway.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WETHGateway *WETHGatewayTransactorSession) Receive() (*types.Transaction, error) {
	return _WETHGateway.Contract.Receive(&_WETHGateway.TransactOpts)
}
