// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iCurvePool

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

// ICurvePoolMetaData contains all meta data concerning the ICurvePool contract.
var ICurvePoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"coins\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"dx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"min_dy\",\"type\":\"uint256\"}],\"name\":\"exchange\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"dx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"min_dy\",\"type\":\"uint256\"}],\"name\":\"exchange_underlying\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"dx\",\"type\":\"uint256\"}],\"name\":\"get_dy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"dx\",\"type\":\"uint256\"}],\"name\":\"get_dy_underlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_virtual_price\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ICurvePoolABI is the input ABI used to generate the binding from.
// Deprecated: Use ICurvePoolMetaData.ABI instead.
var ICurvePoolABI = ICurvePoolMetaData.ABI

// ICurvePool is an auto generated Go binding around an Ethereum contract.
type ICurvePool struct {
	ICurvePoolCaller     // Read-only binding to the contract
	ICurvePoolTransactor // Write-only binding to the contract
	ICurvePoolFilterer   // Log filterer for contract events
}

// ICurvePoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type ICurvePoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICurvePoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ICurvePoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICurvePoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ICurvePoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICurvePoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ICurvePoolSession struct {
	Contract     *ICurvePool       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ICurvePoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ICurvePoolCallerSession struct {
	Contract *ICurvePoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ICurvePoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ICurvePoolTransactorSession struct {
	Contract     *ICurvePoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ICurvePoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type ICurvePoolRaw struct {
	Contract *ICurvePool // Generic contract binding to access the raw methods on
}

// ICurvePoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ICurvePoolCallerRaw struct {
	Contract *ICurvePoolCaller // Generic read-only contract binding to access the raw methods on
}

// ICurvePoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ICurvePoolTransactorRaw struct {
	Contract *ICurvePoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewICurvePool creates a new instance of ICurvePool, bound to a specific deployed contract.
func NewICurvePool(address common.Address, backend bind.ContractBackend) (*ICurvePool, error) {
	contract, err := bindICurvePool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ICurvePool{ICurvePoolCaller: ICurvePoolCaller{contract: contract}, ICurvePoolTransactor: ICurvePoolTransactor{contract: contract}, ICurvePoolFilterer: ICurvePoolFilterer{contract: contract}}, nil
}

// NewICurvePoolCaller creates a new read-only instance of ICurvePool, bound to a specific deployed contract.
func NewICurvePoolCaller(address common.Address, caller bind.ContractCaller) (*ICurvePoolCaller, error) {
	contract, err := bindICurvePool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ICurvePoolCaller{contract: contract}, nil
}

// NewICurvePoolTransactor creates a new write-only instance of ICurvePool, bound to a specific deployed contract.
func NewICurvePoolTransactor(address common.Address, transactor bind.ContractTransactor) (*ICurvePoolTransactor, error) {
	contract, err := bindICurvePool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ICurvePoolTransactor{contract: contract}, nil
}

// NewICurvePoolFilterer creates a new log filterer instance of ICurvePool, bound to a specific deployed contract.
func NewICurvePoolFilterer(address common.Address, filterer bind.ContractFilterer) (*ICurvePoolFilterer, error) {
	contract, err := bindICurvePool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ICurvePoolFilterer{contract: contract}, nil
}

// bindICurvePool binds a generic wrapper to an already deployed contract.
func bindICurvePool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ICurvePoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICurvePool *ICurvePoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICurvePool.Contract.ICurvePoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICurvePool *ICurvePoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICurvePool.Contract.ICurvePoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICurvePool *ICurvePoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICurvePool.Contract.ICurvePoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICurvePool *ICurvePoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICurvePool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICurvePool *ICurvePoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICurvePool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICurvePool *ICurvePoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICurvePool.Contract.contract.Transact(opts, method, params...)
}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 ) view returns(address)
func (_ICurvePool *ICurvePoolCaller) Coins(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ICurvePool.contract.Call(opts, &out, "coins", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 ) view returns(address)
func (_ICurvePool *ICurvePoolSession) Coins(arg0 *big.Int) (common.Address, error) {
	return _ICurvePool.Contract.Coins(&_ICurvePool.CallOpts, arg0)
}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 ) view returns(address)
func (_ICurvePool *ICurvePoolCallerSession) Coins(arg0 *big.Int) (common.Address, error) {
	return _ICurvePool.Contract.Coins(&_ICurvePool.CallOpts, arg0)
}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_ICurvePool *ICurvePoolCaller) GetDy(opts *bind.CallOpts, i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ICurvePool.contract.Call(opts, &out, "get_dy", i, j, dx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_ICurvePool *ICurvePoolSession) GetDy(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _ICurvePool.Contract.GetDy(&_ICurvePool.CallOpts, i, j, dx)
}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_ICurvePool *ICurvePoolCallerSession) GetDy(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _ICurvePool.Contract.GetDy(&_ICurvePool.CallOpts, i, j, dx)
}

// GetDyUnderlying is a free data retrieval call binding the contract method 0x07211ef7.
//
// Solidity: function get_dy_underlying(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_ICurvePool *ICurvePoolCaller) GetDyUnderlying(opts *bind.CallOpts, i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ICurvePool.contract.Call(opts, &out, "get_dy_underlying", i, j, dx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDyUnderlying is a free data retrieval call binding the contract method 0x07211ef7.
//
// Solidity: function get_dy_underlying(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_ICurvePool *ICurvePoolSession) GetDyUnderlying(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _ICurvePool.Contract.GetDyUnderlying(&_ICurvePool.CallOpts, i, j, dx)
}

// GetDyUnderlying is a free data retrieval call binding the contract method 0x07211ef7.
//
// Solidity: function get_dy_underlying(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_ICurvePool *ICurvePoolCallerSession) GetDyUnderlying(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _ICurvePool.Contract.GetDyUnderlying(&_ICurvePool.CallOpts, i, j, dx)
}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_ICurvePool *ICurvePoolCaller) GetVirtualPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ICurvePool.contract.Call(opts, &out, "get_virtual_price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_ICurvePool *ICurvePoolSession) GetVirtualPrice() (*big.Int, error) {
	return _ICurvePool.Contract.GetVirtualPrice(&_ICurvePool.CallOpts)
}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_ICurvePool *ICurvePoolCallerSession) GetVirtualPrice() (*big.Int, error) {
	return _ICurvePool.Contract.GetVirtualPrice(&_ICurvePool.CallOpts)
}

// Exchange is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 dx, uint256 min_dy) returns()
func (_ICurvePool *ICurvePoolTransactor) Exchange(opts *bind.TransactOpts, i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _ICurvePool.contract.Transact(opts, "exchange", i, j, dx, min_dy)
}

// Exchange is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 dx, uint256 min_dy) returns()
func (_ICurvePool *ICurvePoolSession) Exchange(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _ICurvePool.Contract.Exchange(&_ICurvePool.TransactOpts, i, j, dx, min_dy)
}

// Exchange is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 dx, uint256 min_dy) returns()
func (_ICurvePool *ICurvePoolTransactorSession) Exchange(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _ICurvePool.Contract.Exchange(&_ICurvePool.TransactOpts, i, j, dx, min_dy)
}

// ExchangeUnderlying is a paid mutator transaction binding the contract method 0xa6417ed6.
//
// Solidity: function exchange_underlying(int128 i, int128 j, uint256 dx, uint256 min_dy) returns()
func (_ICurvePool *ICurvePoolTransactor) ExchangeUnderlying(opts *bind.TransactOpts, i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _ICurvePool.contract.Transact(opts, "exchange_underlying", i, j, dx, min_dy)
}

// ExchangeUnderlying is a paid mutator transaction binding the contract method 0xa6417ed6.
//
// Solidity: function exchange_underlying(int128 i, int128 j, uint256 dx, uint256 min_dy) returns()
func (_ICurvePool *ICurvePoolSession) ExchangeUnderlying(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _ICurvePool.Contract.ExchangeUnderlying(&_ICurvePool.TransactOpts, i, j, dx, min_dy)
}

// ExchangeUnderlying is a paid mutator transaction binding the contract method 0xa6417ed6.
//
// Solidity: function exchange_underlying(int128 i, int128 j, uint256 dx, uint256 min_dy) returns()
func (_ICurvePool *ICurvePoolTransactorSession) ExchangeUnderlying(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _ICurvePool.Contract.ExchangeUnderlying(&_ICurvePool.TransactOpts, i, j, dx, min_dy)
}
