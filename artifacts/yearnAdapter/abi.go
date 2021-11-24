// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package yearnAdapter

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

// YearnAdapterMetaData contains all meta data concerning the YearnAdapter contract.
var YearnAdapterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_creditManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_yVault\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creditFilter\",\"outputs\":[{\"internalType\":\"contractICreditFilter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creditManager\",\"outputs\":[{\"internalType\":\"contractICreditManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pricePerShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxShares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxShares\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxShares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxLoss\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"yVault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// YearnAdapterABI is the input ABI used to generate the binding from.
// Deprecated: Use YearnAdapterMetaData.ABI instead.
var YearnAdapterABI = YearnAdapterMetaData.ABI

// YearnAdapter is an auto generated Go binding around an Ethereum contract.
type YearnAdapter struct {
	YearnAdapterCaller     // Read-only binding to the contract
	YearnAdapterTransactor // Write-only binding to the contract
	YearnAdapterFilterer   // Log filterer for contract events
}

// YearnAdapterCaller is an auto generated read-only Go binding around an Ethereum contract.
type YearnAdapterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YearnAdapterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type YearnAdapterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YearnAdapterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type YearnAdapterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YearnAdapterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type YearnAdapterSession struct {
	Contract     *YearnAdapter     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// YearnAdapterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type YearnAdapterCallerSession struct {
	Contract *YearnAdapterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// YearnAdapterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type YearnAdapterTransactorSession struct {
	Contract     *YearnAdapterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// YearnAdapterRaw is an auto generated low-level Go binding around an Ethereum contract.
type YearnAdapterRaw struct {
	Contract *YearnAdapter // Generic contract binding to access the raw methods on
}

// YearnAdapterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type YearnAdapterCallerRaw struct {
	Contract *YearnAdapterCaller // Generic read-only contract binding to access the raw methods on
}

// YearnAdapterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type YearnAdapterTransactorRaw struct {
	Contract *YearnAdapterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewYearnAdapter creates a new instance of YearnAdapter, bound to a specific deployed contract.
func NewYearnAdapter(address common.Address, backend bind.ContractBackend) (*YearnAdapter, error) {
	contract, err := bindYearnAdapter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &YearnAdapter{YearnAdapterCaller: YearnAdapterCaller{contract: contract}, YearnAdapterTransactor: YearnAdapterTransactor{contract: contract}, YearnAdapterFilterer: YearnAdapterFilterer{contract: contract}}, nil
}

// NewYearnAdapterCaller creates a new read-only instance of YearnAdapter, bound to a specific deployed contract.
func NewYearnAdapterCaller(address common.Address, caller bind.ContractCaller) (*YearnAdapterCaller, error) {
	contract, err := bindYearnAdapter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &YearnAdapterCaller{contract: contract}, nil
}

// NewYearnAdapterTransactor creates a new write-only instance of YearnAdapter, bound to a specific deployed contract.
func NewYearnAdapterTransactor(address common.Address, transactor bind.ContractTransactor) (*YearnAdapterTransactor, error) {
	contract, err := bindYearnAdapter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &YearnAdapterTransactor{contract: contract}, nil
}

// NewYearnAdapterFilterer creates a new log filterer instance of YearnAdapter, bound to a specific deployed contract.
func NewYearnAdapterFilterer(address common.Address, filterer bind.ContractFilterer) (*YearnAdapterFilterer, error) {
	contract, err := bindYearnAdapter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &YearnAdapterFilterer{contract: contract}, nil
}

// bindYearnAdapter binds a generic wrapper to an already deployed contract.
func bindYearnAdapter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(YearnAdapterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YearnAdapter *YearnAdapterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YearnAdapter.Contract.YearnAdapterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YearnAdapter *YearnAdapterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YearnAdapter.Contract.YearnAdapterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YearnAdapter *YearnAdapterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YearnAdapter.Contract.YearnAdapterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YearnAdapter *YearnAdapterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YearnAdapter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YearnAdapter *YearnAdapterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YearnAdapter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YearnAdapter *YearnAdapterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YearnAdapter.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_YearnAdapter *YearnAdapterCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YearnAdapter.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_YearnAdapter *YearnAdapterSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _YearnAdapter.Contract.Allowance(&_YearnAdapter.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_YearnAdapter *YearnAdapterCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _YearnAdapter.Contract.Allowance(&_YearnAdapter.CallOpts, owner, spender)
}

// Approve is a free data retrieval call binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address , uint256 ) pure returns(bool)
func (_YearnAdapter *YearnAdapterCaller) Approve(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (bool, error) {
	var out []interface{}
	err := _YearnAdapter.contract.Call(opts, &out, "approve", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Approve is a free data retrieval call binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address , uint256 ) pure returns(bool)
func (_YearnAdapter *YearnAdapterSession) Approve(arg0 common.Address, arg1 *big.Int) (bool, error) {
	return _YearnAdapter.Contract.Approve(&_YearnAdapter.CallOpts, arg0, arg1)
}

// Approve is a free data retrieval call binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address , uint256 ) pure returns(bool)
func (_YearnAdapter *YearnAdapterCallerSession) Approve(arg0 common.Address, arg1 *big.Int) (bool, error) {
	return _YearnAdapter.Contract.Approve(&_YearnAdapter.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_YearnAdapter *YearnAdapterCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YearnAdapter.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_YearnAdapter *YearnAdapterSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _YearnAdapter.Contract.BalanceOf(&_YearnAdapter.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_YearnAdapter *YearnAdapterCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _YearnAdapter.Contract.BalanceOf(&_YearnAdapter.CallOpts, account)
}

// CreditFilter is a free data retrieval call binding the contract method 0xf93f515b.
//
// Solidity: function creditFilter() view returns(address)
func (_YearnAdapter *YearnAdapterCaller) CreditFilter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YearnAdapter.contract.Call(opts, &out, "creditFilter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditFilter is a free data retrieval call binding the contract method 0xf93f515b.
//
// Solidity: function creditFilter() view returns(address)
func (_YearnAdapter *YearnAdapterSession) CreditFilter() (common.Address, error) {
	return _YearnAdapter.Contract.CreditFilter(&_YearnAdapter.CallOpts)
}

// CreditFilter is a free data retrieval call binding the contract method 0xf93f515b.
//
// Solidity: function creditFilter() view returns(address)
func (_YearnAdapter *YearnAdapterCallerSession) CreditFilter() (common.Address, error) {
	return _YearnAdapter.Contract.CreditFilter(&_YearnAdapter.CallOpts)
}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_YearnAdapter *YearnAdapterCaller) CreditManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YearnAdapter.contract.Call(opts, &out, "creditManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_YearnAdapter *YearnAdapterSession) CreditManager() (common.Address, error) {
	return _YearnAdapter.Contract.CreditManager(&_YearnAdapter.CallOpts)
}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_YearnAdapter *YearnAdapterCallerSession) CreditManager() (common.Address, error) {
	return _YearnAdapter.Contract.CreditManager(&_YearnAdapter.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_YearnAdapter *YearnAdapterCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _YearnAdapter.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_YearnAdapter *YearnAdapterSession) Decimals() (uint8, error) {
	return _YearnAdapter.Contract.Decimals(&_YearnAdapter.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_YearnAdapter *YearnAdapterCallerSession) Decimals() (uint8, error) {
	return _YearnAdapter.Contract.Decimals(&_YearnAdapter.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_YearnAdapter *YearnAdapterCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _YearnAdapter.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_YearnAdapter *YearnAdapterSession) Name() (string, error) {
	return _YearnAdapter.Contract.Name(&_YearnAdapter.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_YearnAdapter *YearnAdapterCallerSession) Name() (string, error) {
	return _YearnAdapter.Contract.Name(&_YearnAdapter.CallOpts)
}

// PricePerShare is a free data retrieval call binding the contract method 0x99530b06.
//
// Solidity: function pricePerShare() view returns(uint256)
func (_YearnAdapter *YearnAdapterCaller) PricePerShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YearnAdapter.contract.Call(opts, &out, "pricePerShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PricePerShare is a free data retrieval call binding the contract method 0x99530b06.
//
// Solidity: function pricePerShare() view returns(uint256)
func (_YearnAdapter *YearnAdapterSession) PricePerShare() (*big.Int, error) {
	return _YearnAdapter.Contract.PricePerShare(&_YearnAdapter.CallOpts)
}

// PricePerShare is a free data retrieval call binding the contract method 0x99530b06.
//
// Solidity: function pricePerShare() view returns(uint256)
func (_YearnAdapter *YearnAdapterCallerSession) PricePerShare() (*big.Int, error) {
	return _YearnAdapter.Contract.PricePerShare(&_YearnAdapter.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_YearnAdapter *YearnAdapterCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _YearnAdapter.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_YearnAdapter *YearnAdapterSession) Symbol() (string, error) {
	return _YearnAdapter.Contract.Symbol(&_YearnAdapter.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_YearnAdapter *YearnAdapterCallerSession) Symbol() (string, error) {
	return _YearnAdapter.Contract.Symbol(&_YearnAdapter.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_YearnAdapter *YearnAdapterCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YearnAdapter.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_YearnAdapter *YearnAdapterSession) Token() (common.Address, error) {
	return _YearnAdapter.Contract.Token(&_YearnAdapter.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_YearnAdapter *YearnAdapterCallerSession) Token() (common.Address, error) {
	return _YearnAdapter.Contract.Token(&_YearnAdapter.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_YearnAdapter *YearnAdapterCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YearnAdapter.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_YearnAdapter *YearnAdapterSession) TotalSupply() (*big.Int, error) {
	return _YearnAdapter.Contract.TotalSupply(&_YearnAdapter.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_YearnAdapter *YearnAdapterCallerSession) TotalSupply() (*big.Int, error) {
	return _YearnAdapter.Contract.TotalSupply(&_YearnAdapter.CallOpts)
}

// Transfer is a free data retrieval call binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address , uint256 ) pure returns(bool)
func (_YearnAdapter *YearnAdapterCaller) Transfer(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (bool, error) {
	var out []interface{}
	err := _YearnAdapter.contract.Call(opts, &out, "transfer", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Transfer is a free data retrieval call binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address , uint256 ) pure returns(bool)
func (_YearnAdapter *YearnAdapterSession) Transfer(arg0 common.Address, arg1 *big.Int) (bool, error) {
	return _YearnAdapter.Contract.Transfer(&_YearnAdapter.CallOpts, arg0, arg1)
}

// Transfer is a free data retrieval call binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address , uint256 ) pure returns(bool)
func (_YearnAdapter *YearnAdapterCallerSession) Transfer(arg0 common.Address, arg1 *big.Int) (bool, error) {
	return _YearnAdapter.Contract.Transfer(&_YearnAdapter.CallOpts, arg0, arg1)
}

// TransferFrom is a free data retrieval call binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address , address , uint256 ) pure returns(bool)
func (_YearnAdapter *YearnAdapterCaller) TransferFrom(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int) (bool, error) {
	var out []interface{}
	err := _YearnAdapter.contract.Call(opts, &out, "transferFrom", arg0, arg1, arg2)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TransferFrom is a free data retrieval call binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address , address , uint256 ) pure returns(bool)
func (_YearnAdapter *YearnAdapterSession) TransferFrom(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (bool, error) {
	return _YearnAdapter.Contract.TransferFrom(&_YearnAdapter.CallOpts, arg0, arg1, arg2)
}

// TransferFrom is a free data retrieval call binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address , address , uint256 ) pure returns(bool)
func (_YearnAdapter *YearnAdapterCallerSession) TransferFrom(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (bool, error) {
	return _YearnAdapter.Contract.TransferFrom(&_YearnAdapter.CallOpts, arg0, arg1, arg2)
}

// YVault is a free data retrieval call binding the contract method 0x33303f8e.
//
// Solidity: function yVault() view returns(address)
func (_YearnAdapter *YearnAdapterCaller) YVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YearnAdapter.contract.Call(opts, &out, "yVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// YVault is a free data retrieval call binding the contract method 0x33303f8e.
//
// Solidity: function yVault() view returns(address)
func (_YearnAdapter *YearnAdapterSession) YVault() (common.Address, error) {
	return _YearnAdapter.Contract.YVault(&_YearnAdapter.CallOpts)
}

// YVault is a free data retrieval call binding the contract method 0x33303f8e.
//
// Solidity: function yVault() view returns(address)
func (_YearnAdapter *YearnAdapterCallerSession) YVault() (common.Address, error) {
	return _YearnAdapter.Contract.YVault(&_YearnAdapter.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 amount, address ) returns(uint256)
func (_YearnAdapter *YearnAdapterTransactor) Deposit(opts *bind.TransactOpts, amount *big.Int, arg1 common.Address) (*types.Transaction, error) {
	return _YearnAdapter.contract.Transact(opts, "deposit", amount, arg1)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 amount, address ) returns(uint256)
func (_YearnAdapter *YearnAdapterSession) Deposit(amount *big.Int, arg1 common.Address) (*types.Transaction, error) {
	return _YearnAdapter.Contract.Deposit(&_YearnAdapter.TransactOpts, amount, arg1)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 amount, address ) returns(uint256)
func (_YearnAdapter *YearnAdapterTransactorSession) Deposit(amount *big.Int, arg1 common.Address) (*types.Transaction, error) {
	return _YearnAdapter.Contract.Deposit(&_YearnAdapter.TransactOpts, amount, arg1)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns(uint256)
func (_YearnAdapter *YearnAdapterTransactor) Deposit0(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _YearnAdapter.contract.Transact(opts, "deposit0", amount)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns(uint256)
func (_YearnAdapter *YearnAdapterSession) Deposit0(amount *big.Int) (*types.Transaction, error) {
	return _YearnAdapter.Contract.Deposit0(&_YearnAdapter.TransactOpts, amount)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns(uint256)
func (_YearnAdapter *YearnAdapterTransactorSession) Deposit0(amount *big.Int) (*types.Transaction, error) {
	return _YearnAdapter.Contract.Deposit0(&_YearnAdapter.TransactOpts, amount)
}

// Deposit1 is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns(uint256)
func (_YearnAdapter *YearnAdapterTransactor) Deposit1(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YearnAdapter.contract.Transact(opts, "deposit1")
}

// Deposit1 is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns(uint256)
func (_YearnAdapter *YearnAdapterSession) Deposit1() (*types.Transaction, error) {
	return _YearnAdapter.Contract.Deposit1(&_YearnAdapter.TransactOpts)
}

// Deposit1 is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns(uint256)
func (_YearnAdapter *YearnAdapterTransactorSession) Deposit1() (*types.Transaction, error) {
	return _YearnAdapter.Contract.Deposit1(&_YearnAdapter.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 maxShares, address ) returns(uint256)
func (_YearnAdapter *YearnAdapterTransactor) Withdraw(opts *bind.TransactOpts, maxShares *big.Int, arg1 common.Address) (*types.Transaction, error) {
	return _YearnAdapter.contract.Transact(opts, "withdraw", maxShares, arg1)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 maxShares, address ) returns(uint256)
func (_YearnAdapter *YearnAdapterSession) Withdraw(maxShares *big.Int, arg1 common.Address) (*types.Transaction, error) {
	return _YearnAdapter.Contract.Withdraw(&_YearnAdapter.TransactOpts, maxShares, arg1)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 maxShares, address ) returns(uint256)
func (_YearnAdapter *YearnAdapterTransactorSession) Withdraw(maxShares *big.Int, arg1 common.Address) (*types.Transaction, error) {
	return _YearnAdapter.Contract.Withdraw(&_YearnAdapter.TransactOpts, maxShares, arg1)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 maxShares) returns(uint256)
func (_YearnAdapter *YearnAdapterTransactor) Withdraw0(opts *bind.TransactOpts, maxShares *big.Int) (*types.Transaction, error) {
	return _YearnAdapter.contract.Transact(opts, "withdraw0", maxShares)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 maxShares) returns(uint256)
func (_YearnAdapter *YearnAdapterSession) Withdraw0(maxShares *big.Int) (*types.Transaction, error) {
	return _YearnAdapter.Contract.Withdraw0(&_YearnAdapter.TransactOpts, maxShares)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 maxShares) returns(uint256)
func (_YearnAdapter *YearnAdapterTransactorSession) Withdraw0(maxShares *big.Int) (*types.Transaction, error) {
	return _YearnAdapter.Contract.Withdraw0(&_YearnAdapter.TransactOpts, maxShares)
}

// Withdraw1 is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(uint256)
func (_YearnAdapter *YearnAdapterTransactor) Withdraw1(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YearnAdapter.contract.Transact(opts, "withdraw1")
}

// Withdraw1 is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(uint256)
func (_YearnAdapter *YearnAdapterSession) Withdraw1() (*types.Transaction, error) {
	return _YearnAdapter.Contract.Withdraw1(&_YearnAdapter.TransactOpts)
}

// Withdraw1 is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(uint256)
func (_YearnAdapter *YearnAdapterTransactorSession) Withdraw1() (*types.Transaction, error) {
	return _YearnAdapter.Contract.Withdraw1(&_YearnAdapter.TransactOpts)
}

// Withdraw2 is a paid mutator transaction binding the contract method 0xe63697c8.
//
// Solidity: function withdraw(uint256 maxShares, address , uint256 maxLoss) returns(uint256 shares)
func (_YearnAdapter *YearnAdapterTransactor) Withdraw2(opts *bind.TransactOpts, maxShares *big.Int, arg1 common.Address, maxLoss *big.Int) (*types.Transaction, error) {
	return _YearnAdapter.contract.Transact(opts, "withdraw2", maxShares, arg1, maxLoss)
}

// Withdraw2 is a paid mutator transaction binding the contract method 0xe63697c8.
//
// Solidity: function withdraw(uint256 maxShares, address , uint256 maxLoss) returns(uint256 shares)
func (_YearnAdapter *YearnAdapterSession) Withdraw2(maxShares *big.Int, arg1 common.Address, maxLoss *big.Int) (*types.Transaction, error) {
	return _YearnAdapter.Contract.Withdraw2(&_YearnAdapter.TransactOpts, maxShares, arg1, maxLoss)
}

// Withdraw2 is a paid mutator transaction binding the contract method 0xe63697c8.
//
// Solidity: function withdraw(uint256 maxShares, address , uint256 maxLoss) returns(uint256 shares)
func (_YearnAdapter *YearnAdapterTransactorSession) Withdraw2(maxShares *big.Int, arg1 common.Address, maxLoss *big.Int) (*types.Transaction, error) {
	return _YearnAdapter.Contract.Withdraw2(&_YearnAdapter.TransactOpts, maxShares, arg1, maxLoss)
}

// YearnAdapterApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the YearnAdapter contract.
type YearnAdapterApprovalIterator struct {
	Event *YearnAdapterApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *YearnAdapterApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnAdapterApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(YearnAdapterApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *YearnAdapterApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnAdapterApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnAdapterApproval represents a Approval event raised by the YearnAdapter contract.
type YearnAdapterApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_YearnAdapter *YearnAdapterFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*YearnAdapterApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _YearnAdapter.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &YearnAdapterApprovalIterator{contract: _YearnAdapter.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_YearnAdapter *YearnAdapterFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *YearnAdapterApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _YearnAdapter.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnAdapterApproval)
				if err := _YearnAdapter.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_YearnAdapter *YearnAdapterFilterer) ParseApproval(log types.Log) (*YearnAdapterApproval, error) {
	event := new(YearnAdapterApproval)
	if err := _YearnAdapter.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnAdapterTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the YearnAdapter contract.
type YearnAdapterTransferIterator struct {
	Event *YearnAdapterTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *YearnAdapterTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnAdapterTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(YearnAdapterTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *YearnAdapterTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnAdapterTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnAdapterTransfer represents a Transfer event raised by the YearnAdapter contract.
type YearnAdapterTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_YearnAdapter *YearnAdapterFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*YearnAdapterTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _YearnAdapter.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &YearnAdapterTransferIterator{contract: _YearnAdapter.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_YearnAdapter *YearnAdapterFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *YearnAdapterTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _YearnAdapter.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnAdapterTransfer)
				if err := _YearnAdapter.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_YearnAdapter *YearnAdapterFilterer) ParseTransfer(log types.Log) (*YearnAdapterTransfer, error) {
	event := new(YearnAdapterTransfer)
	if err := _YearnAdapter.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
