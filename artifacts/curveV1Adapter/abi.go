// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package curveV1Adapter

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

// CurveV1AdapterMetaData contains all meta data concerning the CurveV1Adapter contract.
var CurveV1AdapterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_creditManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_curvePool\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"coins\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creditFilter\",\"outputs\":[{\"internalType\":\"contractICreditFilter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creditManager\",\"outputs\":[{\"internalType\":\"contractICreditManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"curvePool\",\"outputs\":[{\"internalType\":\"contractICurvePool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"dx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"min_dy\",\"type\":\"uint256\"}],\"name\":\"exchange\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"dx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"min_dy\",\"type\":\"uint256\"}],\"name\":\"exchange_underlying\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"dx\",\"type\":\"uint256\"}],\"name\":\"get_dy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"dx\",\"type\":\"uint256\"}],\"name\":\"get_dy_underlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_virtual_price\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// CurveV1AdapterABI is the input ABI used to generate the binding from.
// Deprecated: Use CurveV1AdapterMetaData.ABI instead.
var CurveV1AdapterABI = CurveV1AdapterMetaData.ABI

// CurveV1Adapter is an auto generated Go binding around an Ethereum contract.
type CurveV1Adapter struct {
	CurveV1AdapterCaller     // Read-only binding to the contract
	CurveV1AdapterTransactor // Write-only binding to the contract
	CurveV1AdapterFilterer   // Log filterer for contract events
}

// CurveV1AdapterCaller is an auto generated read-only Go binding around an Ethereum contract.
type CurveV1AdapterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveV1AdapterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CurveV1AdapterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveV1AdapterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CurveV1AdapterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveV1AdapterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CurveV1AdapterSession struct {
	Contract     *CurveV1Adapter   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CurveV1AdapterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CurveV1AdapterCallerSession struct {
	Contract *CurveV1AdapterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// CurveV1AdapterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CurveV1AdapterTransactorSession struct {
	Contract     *CurveV1AdapterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// CurveV1AdapterRaw is an auto generated low-level Go binding around an Ethereum contract.
type CurveV1AdapterRaw struct {
	Contract *CurveV1Adapter // Generic contract binding to access the raw methods on
}

// CurveV1AdapterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CurveV1AdapterCallerRaw struct {
	Contract *CurveV1AdapterCaller // Generic read-only contract binding to access the raw methods on
}

// CurveV1AdapterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CurveV1AdapterTransactorRaw struct {
	Contract *CurveV1AdapterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCurveV1Adapter creates a new instance of CurveV1Adapter, bound to a specific deployed contract.
func NewCurveV1Adapter(address common.Address, backend bind.ContractBackend) (*CurveV1Adapter, error) {
	contract, err := bindCurveV1Adapter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CurveV1Adapter{CurveV1AdapterCaller: CurveV1AdapterCaller{contract: contract}, CurveV1AdapterTransactor: CurveV1AdapterTransactor{contract: contract}, CurveV1AdapterFilterer: CurveV1AdapterFilterer{contract: contract}}, nil
}

// NewCurveV1AdapterCaller creates a new read-only instance of CurveV1Adapter, bound to a specific deployed contract.
func NewCurveV1AdapterCaller(address common.Address, caller bind.ContractCaller) (*CurveV1AdapterCaller, error) {
	contract, err := bindCurveV1Adapter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CurveV1AdapterCaller{contract: contract}, nil
}

// NewCurveV1AdapterTransactor creates a new write-only instance of CurveV1Adapter, bound to a specific deployed contract.
func NewCurveV1AdapterTransactor(address common.Address, transactor bind.ContractTransactor) (*CurveV1AdapterTransactor, error) {
	contract, err := bindCurveV1Adapter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CurveV1AdapterTransactor{contract: contract}, nil
}

// NewCurveV1AdapterFilterer creates a new log filterer instance of CurveV1Adapter, bound to a specific deployed contract.
func NewCurveV1AdapterFilterer(address common.Address, filterer bind.ContractFilterer) (*CurveV1AdapterFilterer, error) {
	contract, err := bindCurveV1Adapter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CurveV1AdapterFilterer{contract: contract}, nil
}

// bindCurveV1Adapter binds a generic wrapper to an already deployed contract.
func bindCurveV1Adapter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CurveV1AdapterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CurveV1Adapter *CurveV1AdapterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CurveV1Adapter.Contract.CurveV1AdapterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CurveV1Adapter *CurveV1AdapterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveV1Adapter.Contract.CurveV1AdapterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CurveV1Adapter *CurveV1AdapterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CurveV1Adapter.Contract.CurveV1AdapterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CurveV1Adapter *CurveV1AdapterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CurveV1Adapter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CurveV1Adapter *CurveV1AdapterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveV1Adapter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CurveV1Adapter *CurveV1AdapterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CurveV1Adapter.Contract.contract.Transact(opts, method, params...)
}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 i) view returns(address)
func (_CurveV1Adapter *CurveV1AdapterCaller) Coins(opts *bind.CallOpts, i *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CurveV1Adapter.contract.Call(opts, &out, "coins", i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 i) view returns(address)
func (_CurveV1Adapter *CurveV1AdapterSession) Coins(i *big.Int) (common.Address, error) {
	return _CurveV1Adapter.Contract.Coins(&_CurveV1Adapter.CallOpts, i)
}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 i) view returns(address)
func (_CurveV1Adapter *CurveV1AdapterCallerSession) Coins(i *big.Int) (common.Address, error) {
	return _CurveV1Adapter.Contract.Coins(&_CurveV1Adapter.CallOpts, i)
}

// CreditFilter is a free data retrieval call binding the contract method 0xf93f515b.
//
// Solidity: function creditFilter() view returns(address)
func (_CurveV1Adapter *CurveV1AdapterCaller) CreditFilter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveV1Adapter.contract.Call(opts, &out, "creditFilter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditFilter is a free data retrieval call binding the contract method 0xf93f515b.
//
// Solidity: function creditFilter() view returns(address)
func (_CurveV1Adapter *CurveV1AdapterSession) CreditFilter() (common.Address, error) {
	return _CurveV1Adapter.Contract.CreditFilter(&_CurveV1Adapter.CallOpts)
}

// CreditFilter is a free data retrieval call binding the contract method 0xf93f515b.
//
// Solidity: function creditFilter() view returns(address)
func (_CurveV1Adapter *CurveV1AdapterCallerSession) CreditFilter() (common.Address, error) {
	return _CurveV1Adapter.Contract.CreditFilter(&_CurveV1Adapter.CallOpts)
}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_CurveV1Adapter *CurveV1AdapterCaller) CreditManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveV1Adapter.contract.Call(opts, &out, "creditManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_CurveV1Adapter *CurveV1AdapterSession) CreditManager() (common.Address, error) {
	return _CurveV1Adapter.Contract.CreditManager(&_CurveV1Adapter.CallOpts)
}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_CurveV1Adapter *CurveV1AdapterCallerSession) CreditManager() (common.Address, error) {
	return _CurveV1Adapter.Contract.CreditManager(&_CurveV1Adapter.CallOpts)
}

// CurvePool is a free data retrieval call binding the contract method 0x218751b2.
//
// Solidity: function curvePool() view returns(address)
func (_CurveV1Adapter *CurveV1AdapterCaller) CurvePool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveV1Adapter.contract.Call(opts, &out, "curvePool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CurvePool is a free data retrieval call binding the contract method 0x218751b2.
//
// Solidity: function curvePool() view returns(address)
func (_CurveV1Adapter *CurveV1AdapterSession) CurvePool() (common.Address, error) {
	return _CurveV1Adapter.Contract.CurvePool(&_CurveV1Adapter.CallOpts)
}

// CurvePool is a free data retrieval call binding the contract method 0x218751b2.
//
// Solidity: function curvePool() view returns(address)
func (_CurveV1Adapter *CurveV1AdapterCallerSession) CurvePool() (common.Address, error) {
	return _CurveV1Adapter.Contract.CurvePool(&_CurveV1Adapter.CallOpts)
}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurveV1Adapter *CurveV1AdapterCaller) GetDy(opts *bind.CallOpts, i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveV1Adapter.contract.Call(opts, &out, "get_dy", i, j, dx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurveV1Adapter *CurveV1AdapterSession) GetDy(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _CurveV1Adapter.Contract.GetDy(&_CurveV1Adapter.CallOpts, i, j, dx)
}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurveV1Adapter *CurveV1AdapterCallerSession) GetDy(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _CurveV1Adapter.Contract.GetDy(&_CurveV1Adapter.CallOpts, i, j, dx)
}

// GetDyUnderlying is a free data retrieval call binding the contract method 0x07211ef7.
//
// Solidity: function get_dy_underlying(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurveV1Adapter *CurveV1AdapterCaller) GetDyUnderlying(opts *bind.CallOpts, i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveV1Adapter.contract.Call(opts, &out, "get_dy_underlying", i, j, dx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDyUnderlying is a free data retrieval call binding the contract method 0x07211ef7.
//
// Solidity: function get_dy_underlying(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurveV1Adapter *CurveV1AdapterSession) GetDyUnderlying(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _CurveV1Adapter.Contract.GetDyUnderlying(&_CurveV1Adapter.CallOpts, i, j, dx)
}

// GetDyUnderlying is a free data retrieval call binding the contract method 0x07211ef7.
//
// Solidity: function get_dy_underlying(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurveV1Adapter *CurveV1AdapterCallerSession) GetDyUnderlying(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _CurveV1Adapter.Contract.GetDyUnderlying(&_CurveV1Adapter.CallOpts, i, j, dx)
}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_CurveV1Adapter *CurveV1AdapterCaller) GetVirtualPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveV1Adapter.contract.Call(opts, &out, "get_virtual_price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_CurveV1Adapter *CurveV1AdapterSession) GetVirtualPrice() (*big.Int, error) {
	return _CurveV1Adapter.Contract.GetVirtualPrice(&_CurveV1Adapter.CallOpts)
}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_CurveV1Adapter *CurveV1AdapterCallerSession) GetVirtualPrice() (*big.Int, error) {
	return _CurveV1Adapter.Contract.GetVirtualPrice(&_CurveV1Adapter.CallOpts)
}

// Exchange is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 dx, uint256 min_dy) returns()
func (_CurveV1Adapter *CurveV1AdapterTransactor) Exchange(opts *bind.TransactOpts, i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _CurveV1Adapter.contract.Transact(opts, "exchange", i, j, dx, min_dy)
}

// Exchange is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 dx, uint256 min_dy) returns()
func (_CurveV1Adapter *CurveV1AdapterSession) Exchange(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _CurveV1Adapter.Contract.Exchange(&_CurveV1Adapter.TransactOpts, i, j, dx, min_dy)
}

// Exchange is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 dx, uint256 min_dy) returns()
func (_CurveV1Adapter *CurveV1AdapterTransactorSession) Exchange(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _CurveV1Adapter.Contract.Exchange(&_CurveV1Adapter.TransactOpts, i, j, dx, min_dy)
}

// ExchangeUnderlying is a paid mutator transaction binding the contract method 0xa6417ed6.
//
// Solidity: function exchange_underlying(int128 i, int128 j, uint256 dx, uint256 min_dy) returns()
func (_CurveV1Adapter *CurveV1AdapterTransactor) ExchangeUnderlying(opts *bind.TransactOpts, i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _CurveV1Adapter.contract.Transact(opts, "exchange_underlying", i, j, dx, min_dy)
}

// ExchangeUnderlying is a paid mutator transaction binding the contract method 0xa6417ed6.
//
// Solidity: function exchange_underlying(int128 i, int128 j, uint256 dx, uint256 min_dy) returns()
func (_CurveV1Adapter *CurveV1AdapterSession) ExchangeUnderlying(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _CurveV1Adapter.Contract.ExchangeUnderlying(&_CurveV1Adapter.TransactOpts, i, j, dx, min_dy)
}

// ExchangeUnderlying is a paid mutator transaction binding the contract method 0xa6417ed6.
//
// Solidity: function exchange_underlying(int128 i, int128 j, uint256 dx, uint256 min_dy) returns()
func (_CurveV1Adapter *CurveV1AdapterTransactorSession) ExchangeUnderlying(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _CurveV1Adapter.Contract.ExchangeUnderlying(&_CurveV1Adapter.TransactOpts, i, j, dx, min_dy)
}
