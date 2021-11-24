// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package curveMock

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

// CurveMockMetaData contains all meta data concerning the CurveMock contract.
var CurveMockMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_coins\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"coins\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"dx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"min_dy\",\"type\":\"uint256\"}],\"name\":\"exchange\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"dx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"min_dy\",\"type\":\"uint256\"}],\"name\":\"exchange_underlying\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"dx\",\"type\":\"uint256\"}],\"name\":\"get_dy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"dx\",\"type\":\"uint256\"}],\"name\":\"get_dy_underlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_virtual_price\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// CurveMockABI is the input ABI used to generate the binding from.
// Deprecated: Use CurveMockMetaData.ABI instead.
var CurveMockABI = CurveMockMetaData.ABI

// CurveMock is an auto generated Go binding around an Ethereum contract.
type CurveMock struct {
	CurveMockCaller     // Read-only binding to the contract
	CurveMockTransactor // Write-only binding to the contract
	CurveMockFilterer   // Log filterer for contract events
}

// CurveMockCaller is an auto generated read-only Go binding around an Ethereum contract.
type CurveMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveMockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CurveMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveMockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CurveMockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveMockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CurveMockSession struct {
	Contract     *CurveMock        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CurveMockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CurveMockCallerSession struct {
	Contract *CurveMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// CurveMockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CurveMockTransactorSession struct {
	Contract     *CurveMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// CurveMockRaw is an auto generated low-level Go binding around an Ethereum contract.
type CurveMockRaw struct {
	Contract *CurveMock // Generic contract binding to access the raw methods on
}

// CurveMockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CurveMockCallerRaw struct {
	Contract *CurveMockCaller // Generic read-only contract binding to access the raw methods on
}

// CurveMockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CurveMockTransactorRaw struct {
	Contract *CurveMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCurveMock creates a new instance of CurveMock, bound to a specific deployed contract.
func NewCurveMock(address common.Address, backend bind.ContractBackend) (*CurveMock, error) {
	contract, err := bindCurveMock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CurveMock{CurveMockCaller: CurveMockCaller{contract: contract}, CurveMockTransactor: CurveMockTransactor{contract: contract}, CurveMockFilterer: CurveMockFilterer{contract: contract}}, nil
}

// NewCurveMockCaller creates a new read-only instance of CurveMock, bound to a specific deployed contract.
func NewCurveMockCaller(address common.Address, caller bind.ContractCaller) (*CurveMockCaller, error) {
	contract, err := bindCurveMock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CurveMockCaller{contract: contract}, nil
}

// NewCurveMockTransactor creates a new write-only instance of CurveMock, bound to a specific deployed contract.
func NewCurveMockTransactor(address common.Address, transactor bind.ContractTransactor) (*CurveMockTransactor, error) {
	contract, err := bindCurveMock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CurveMockTransactor{contract: contract}, nil
}

// NewCurveMockFilterer creates a new log filterer instance of CurveMock, bound to a specific deployed contract.
func NewCurveMockFilterer(address common.Address, filterer bind.ContractFilterer) (*CurveMockFilterer, error) {
	contract, err := bindCurveMock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CurveMockFilterer{contract: contract}, nil
}

// bindCurveMock binds a generic wrapper to an already deployed contract.
func bindCurveMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CurveMockABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CurveMock *CurveMockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CurveMock.Contract.CurveMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CurveMock *CurveMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveMock.Contract.CurveMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CurveMock *CurveMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CurveMock.Contract.CurveMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CurveMock *CurveMockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CurveMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CurveMock *CurveMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CurveMock *CurveMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CurveMock.Contract.contract.Transact(opts, method, params...)
}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 ) view returns(address)
func (_CurveMock *CurveMockCaller) Coins(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CurveMock.contract.Call(opts, &out, "coins", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 ) view returns(address)
func (_CurveMock *CurveMockSession) Coins(arg0 *big.Int) (common.Address, error) {
	return _CurveMock.Contract.Coins(&_CurveMock.CallOpts, arg0)
}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 ) view returns(address)
func (_CurveMock *CurveMockCallerSession) Coins(arg0 *big.Int) (common.Address, error) {
	return _CurveMock.Contract.Coins(&_CurveMock.CallOpts, arg0)
}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurveMock *CurveMockCaller) GetDy(opts *bind.CallOpts, i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveMock.contract.Call(opts, &out, "get_dy", i, j, dx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurveMock *CurveMockSession) GetDy(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _CurveMock.Contract.GetDy(&_CurveMock.CallOpts, i, j, dx)
}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurveMock *CurveMockCallerSession) GetDy(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _CurveMock.Contract.GetDy(&_CurveMock.CallOpts, i, j, dx)
}

// GetDyUnderlying is a free data retrieval call binding the contract method 0x07211ef7.
//
// Solidity: function get_dy_underlying(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurveMock *CurveMockCaller) GetDyUnderlying(opts *bind.CallOpts, i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveMock.contract.Call(opts, &out, "get_dy_underlying", i, j, dx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDyUnderlying is a free data retrieval call binding the contract method 0x07211ef7.
//
// Solidity: function get_dy_underlying(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurveMock *CurveMockSession) GetDyUnderlying(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _CurveMock.Contract.GetDyUnderlying(&_CurveMock.CallOpts, i, j, dx)
}

// GetDyUnderlying is a free data retrieval call binding the contract method 0x07211ef7.
//
// Solidity: function get_dy_underlying(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurveMock *CurveMockCallerSession) GetDyUnderlying(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _CurveMock.Contract.GetDyUnderlying(&_CurveMock.CallOpts, i, j, dx)
}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_CurveMock *CurveMockCaller) GetVirtualPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveMock.contract.Call(opts, &out, "get_virtual_price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_CurveMock *CurveMockSession) GetVirtualPrice() (*big.Int, error) {
	return _CurveMock.Contract.GetVirtualPrice(&_CurveMock.CallOpts)
}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_CurveMock *CurveMockCallerSession) GetVirtualPrice() (*big.Int, error) {
	return _CurveMock.Contract.GetVirtualPrice(&_CurveMock.CallOpts)
}

// Exchange is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 dx, uint256 min_dy) returns()
func (_CurveMock *CurveMockTransactor) Exchange(opts *bind.TransactOpts, i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _CurveMock.contract.Transact(opts, "exchange", i, j, dx, min_dy)
}

// Exchange is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 dx, uint256 min_dy) returns()
func (_CurveMock *CurveMockSession) Exchange(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _CurveMock.Contract.Exchange(&_CurveMock.TransactOpts, i, j, dx, min_dy)
}

// Exchange is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 dx, uint256 min_dy) returns()
func (_CurveMock *CurveMockTransactorSession) Exchange(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _CurveMock.Contract.Exchange(&_CurveMock.TransactOpts, i, j, dx, min_dy)
}

// ExchangeUnderlying is a paid mutator transaction binding the contract method 0xa6417ed6.
//
// Solidity: function exchange_underlying(int128 i, int128 j, uint256 dx, uint256 min_dy) returns()
func (_CurveMock *CurveMockTransactor) ExchangeUnderlying(opts *bind.TransactOpts, i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _CurveMock.contract.Transact(opts, "exchange_underlying", i, j, dx, min_dy)
}

// ExchangeUnderlying is a paid mutator transaction binding the contract method 0xa6417ed6.
//
// Solidity: function exchange_underlying(int128 i, int128 j, uint256 dx, uint256 min_dy) returns()
func (_CurveMock *CurveMockSession) ExchangeUnderlying(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _CurveMock.Contract.ExchangeUnderlying(&_CurveMock.TransactOpts, i, j, dx, min_dy)
}

// ExchangeUnderlying is a paid mutator transaction binding the contract method 0xa6417ed6.
//
// Solidity: function exchange_underlying(int128 i, int128 j, uint256 dx, uint256 min_dy) returns()
func (_CurveMock *CurveMockTransactorSession) ExchangeUnderlying(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _CurveMock.Contract.ExchangeUnderlying(&_CurveMock.TransactOpts, i, j, dx, min_dy)
}
