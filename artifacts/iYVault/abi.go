// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iYVault

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

// IYVaultMetaData contains all meta data concerning the IYVault contract.
var IYVaultMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pricePerShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxShares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxShares\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxShares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxLoss\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IYVaultABI is the input ABI used to generate the binding from.
// Deprecated: Use IYVaultMetaData.ABI instead.
var IYVaultABI = IYVaultMetaData.ABI

// IYVault is an auto generated Go binding around an Ethereum contract.
type IYVault struct {
	IYVaultCaller     // Read-only binding to the contract
	IYVaultTransactor // Write-only binding to the contract
	IYVaultFilterer   // Log filterer for contract events
}

// IYVaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type IYVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IYVaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IYVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IYVaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IYVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IYVaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IYVaultSession struct {
	Contract     *IYVault          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IYVaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IYVaultCallerSession struct {
	Contract *IYVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IYVaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IYVaultTransactorSession struct {
	Contract     *IYVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IYVaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type IYVaultRaw struct {
	Contract *IYVault // Generic contract binding to access the raw methods on
}

// IYVaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IYVaultCallerRaw struct {
	Contract *IYVaultCaller // Generic read-only contract binding to access the raw methods on
}

// IYVaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IYVaultTransactorRaw struct {
	Contract *IYVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIYVault creates a new instance of IYVault, bound to a specific deployed contract.
func NewIYVault(address common.Address, backend bind.ContractBackend) (*IYVault, error) {
	contract, err := bindIYVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IYVault{IYVaultCaller: IYVaultCaller{contract: contract}, IYVaultTransactor: IYVaultTransactor{contract: contract}, IYVaultFilterer: IYVaultFilterer{contract: contract}}, nil
}

// NewIYVaultCaller creates a new read-only instance of IYVault, bound to a specific deployed contract.
func NewIYVaultCaller(address common.Address, caller bind.ContractCaller) (*IYVaultCaller, error) {
	contract, err := bindIYVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IYVaultCaller{contract: contract}, nil
}

// NewIYVaultTransactor creates a new write-only instance of IYVault, bound to a specific deployed contract.
func NewIYVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*IYVaultTransactor, error) {
	contract, err := bindIYVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IYVaultTransactor{contract: contract}, nil
}

// NewIYVaultFilterer creates a new log filterer instance of IYVault, bound to a specific deployed contract.
func NewIYVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*IYVaultFilterer, error) {
	contract, err := bindIYVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IYVaultFilterer{contract: contract}, nil
}

// bindIYVault binds a generic wrapper to an already deployed contract.
func bindIYVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IYVaultABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IYVault *IYVaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IYVault.Contract.IYVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IYVault *IYVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IYVault.Contract.IYVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IYVault *IYVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IYVault.Contract.IYVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IYVault *IYVaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IYVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IYVault *IYVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IYVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IYVault *IYVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IYVault.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IYVault *IYVaultCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IYVault.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IYVault *IYVaultSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IYVault.Contract.Allowance(&_IYVault.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IYVault *IYVaultCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IYVault.Contract.Allowance(&_IYVault.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IYVault *IYVaultCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IYVault.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IYVault *IYVaultSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IYVault.Contract.BalanceOf(&_IYVault.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IYVault *IYVaultCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IYVault.Contract.BalanceOf(&_IYVault.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IYVault *IYVaultCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _IYVault.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IYVault *IYVaultSession) Decimals() (uint8, error) {
	return _IYVault.Contract.Decimals(&_IYVault.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IYVault *IYVaultCallerSession) Decimals() (uint8, error) {
	return _IYVault.Contract.Decimals(&_IYVault.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IYVault *IYVaultCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IYVault.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IYVault *IYVaultSession) Name() (string, error) {
	return _IYVault.Contract.Name(&_IYVault.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IYVault *IYVaultCallerSession) Name() (string, error) {
	return _IYVault.Contract.Name(&_IYVault.CallOpts)
}

// PricePerShare is a free data retrieval call binding the contract method 0x99530b06.
//
// Solidity: function pricePerShare() view returns(uint256)
func (_IYVault *IYVaultCaller) PricePerShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IYVault.contract.Call(opts, &out, "pricePerShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PricePerShare is a free data retrieval call binding the contract method 0x99530b06.
//
// Solidity: function pricePerShare() view returns(uint256)
func (_IYVault *IYVaultSession) PricePerShare() (*big.Int, error) {
	return _IYVault.Contract.PricePerShare(&_IYVault.CallOpts)
}

// PricePerShare is a free data retrieval call binding the contract method 0x99530b06.
//
// Solidity: function pricePerShare() view returns(uint256)
func (_IYVault *IYVaultCallerSession) PricePerShare() (*big.Int, error) {
	return _IYVault.Contract.PricePerShare(&_IYVault.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IYVault *IYVaultCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IYVault.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IYVault *IYVaultSession) Symbol() (string, error) {
	return _IYVault.Contract.Symbol(&_IYVault.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IYVault *IYVaultCallerSession) Symbol() (string, error) {
	return _IYVault.Contract.Symbol(&_IYVault.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_IYVault *IYVaultCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IYVault.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_IYVault *IYVaultSession) Token() (common.Address, error) {
	return _IYVault.Contract.Token(&_IYVault.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_IYVault *IYVaultCallerSession) Token() (common.Address, error) {
	return _IYVault.Contract.Token(&_IYVault.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IYVault *IYVaultCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IYVault.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IYVault *IYVaultSession) TotalSupply() (*big.Int, error) {
	return _IYVault.Contract.TotalSupply(&_IYVault.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IYVault *IYVaultCallerSession) TotalSupply() (*big.Int, error) {
	return _IYVault.Contract.TotalSupply(&_IYVault.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IYVault *IYVaultTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IYVault.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IYVault *IYVaultSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IYVault.Contract.Approve(&_IYVault.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IYVault *IYVaultTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IYVault.Contract.Approve(&_IYVault.TransactOpts, spender, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 _amount, address recipient) returns(uint256)
func (_IYVault *IYVaultTransactor) Deposit(opts *bind.TransactOpts, _amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _IYVault.contract.Transact(opts, "deposit", _amount, recipient)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 _amount, address recipient) returns(uint256)
func (_IYVault *IYVaultSession) Deposit(_amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _IYVault.Contract.Deposit(&_IYVault.TransactOpts, _amount, recipient)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 _amount, address recipient) returns(uint256)
func (_IYVault *IYVaultTransactorSession) Deposit(_amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _IYVault.Contract.Deposit(&_IYVault.TransactOpts, _amount, recipient)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns(uint256)
func (_IYVault *IYVaultTransactor) Deposit0(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _IYVault.contract.Transact(opts, "deposit0", _amount)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns(uint256)
func (_IYVault *IYVaultSession) Deposit0(_amount *big.Int) (*types.Transaction, error) {
	return _IYVault.Contract.Deposit0(&_IYVault.TransactOpts, _amount)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns(uint256)
func (_IYVault *IYVaultTransactorSession) Deposit0(_amount *big.Int) (*types.Transaction, error) {
	return _IYVault.Contract.Deposit0(&_IYVault.TransactOpts, _amount)
}

// Deposit1 is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns(uint256)
func (_IYVault *IYVaultTransactor) Deposit1(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IYVault.contract.Transact(opts, "deposit1")
}

// Deposit1 is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns(uint256)
func (_IYVault *IYVaultSession) Deposit1() (*types.Transaction, error) {
	return _IYVault.Contract.Deposit1(&_IYVault.TransactOpts)
}

// Deposit1 is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns(uint256)
func (_IYVault *IYVaultTransactorSession) Deposit1() (*types.Transaction, error) {
	return _IYVault.Contract.Deposit1(&_IYVault.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IYVault *IYVaultTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IYVault.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IYVault *IYVaultSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IYVault.Contract.Transfer(&_IYVault.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IYVault *IYVaultTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IYVault.Contract.Transfer(&_IYVault.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IYVault *IYVaultTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IYVault.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IYVault *IYVaultSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IYVault.Contract.TransferFrom(&_IYVault.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IYVault *IYVaultTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IYVault.Contract.TransferFrom(&_IYVault.TransactOpts, sender, recipient, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 maxShares, address recipient) returns(uint256)
func (_IYVault *IYVaultTransactor) Withdraw(opts *bind.TransactOpts, maxShares *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _IYVault.contract.Transact(opts, "withdraw", maxShares, recipient)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 maxShares, address recipient) returns(uint256)
func (_IYVault *IYVaultSession) Withdraw(maxShares *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _IYVault.Contract.Withdraw(&_IYVault.TransactOpts, maxShares, recipient)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 maxShares, address recipient) returns(uint256)
func (_IYVault *IYVaultTransactorSession) Withdraw(maxShares *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _IYVault.Contract.Withdraw(&_IYVault.TransactOpts, maxShares, recipient)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 maxShares) returns(uint256)
func (_IYVault *IYVaultTransactor) Withdraw0(opts *bind.TransactOpts, maxShares *big.Int) (*types.Transaction, error) {
	return _IYVault.contract.Transact(opts, "withdraw0", maxShares)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 maxShares) returns(uint256)
func (_IYVault *IYVaultSession) Withdraw0(maxShares *big.Int) (*types.Transaction, error) {
	return _IYVault.Contract.Withdraw0(&_IYVault.TransactOpts, maxShares)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 maxShares) returns(uint256)
func (_IYVault *IYVaultTransactorSession) Withdraw0(maxShares *big.Int) (*types.Transaction, error) {
	return _IYVault.Contract.Withdraw0(&_IYVault.TransactOpts, maxShares)
}

// Withdraw1 is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(uint256)
func (_IYVault *IYVaultTransactor) Withdraw1(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IYVault.contract.Transact(opts, "withdraw1")
}

// Withdraw1 is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(uint256)
func (_IYVault *IYVaultSession) Withdraw1() (*types.Transaction, error) {
	return _IYVault.Contract.Withdraw1(&_IYVault.TransactOpts)
}

// Withdraw1 is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(uint256)
func (_IYVault *IYVaultTransactorSession) Withdraw1() (*types.Transaction, error) {
	return _IYVault.Contract.Withdraw1(&_IYVault.TransactOpts)
}

// Withdraw2 is a paid mutator transaction binding the contract method 0xe63697c8.
//
// Solidity: function withdraw(uint256 maxShares, address recipient, uint256 maxLoss) returns(uint256)
func (_IYVault *IYVaultTransactor) Withdraw2(opts *bind.TransactOpts, maxShares *big.Int, recipient common.Address, maxLoss *big.Int) (*types.Transaction, error) {
	return _IYVault.contract.Transact(opts, "withdraw2", maxShares, recipient, maxLoss)
}

// Withdraw2 is a paid mutator transaction binding the contract method 0xe63697c8.
//
// Solidity: function withdraw(uint256 maxShares, address recipient, uint256 maxLoss) returns(uint256)
func (_IYVault *IYVaultSession) Withdraw2(maxShares *big.Int, recipient common.Address, maxLoss *big.Int) (*types.Transaction, error) {
	return _IYVault.Contract.Withdraw2(&_IYVault.TransactOpts, maxShares, recipient, maxLoss)
}

// Withdraw2 is a paid mutator transaction binding the contract method 0xe63697c8.
//
// Solidity: function withdraw(uint256 maxShares, address recipient, uint256 maxLoss) returns(uint256)
func (_IYVault *IYVaultTransactorSession) Withdraw2(maxShares *big.Int, recipient common.Address, maxLoss *big.Int) (*types.Transaction, error) {
	return _IYVault.Contract.Withdraw2(&_IYVault.TransactOpts, maxShares, recipient, maxLoss)
}

// IYVaultApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IYVault contract.
type IYVaultApprovalIterator struct {
	Event *IYVaultApproval // Event containing the contract specifics and raw log

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
func (it *IYVaultApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IYVaultApproval)
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
		it.Event = new(IYVaultApproval)
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
func (it *IYVaultApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IYVaultApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IYVaultApproval represents a Approval event raised by the IYVault contract.
type IYVaultApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IYVault *IYVaultFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IYVaultApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IYVault.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IYVaultApprovalIterator{contract: _IYVault.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IYVault *IYVaultFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IYVaultApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IYVault.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IYVaultApproval)
				if err := _IYVault.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_IYVault *IYVaultFilterer) ParseApproval(log types.Log) (*IYVaultApproval, error) {
	event := new(IYVaultApproval)
	if err := _IYVault.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IYVaultTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IYVault contract.
type IYVaultTransferIterator struct {
	Event *IYVaultTransfer // Event containing the contract specifics and raw log

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
func (it *IYVaultTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IYVaultTransfer)
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
		it.Event = new(IYVaultTransfer)
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
func (it *IYVaultTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IYVaultTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IYVaultTransfer represents a Transfer event raised by the IYVault contract.
type IYVaultTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IYVault *IYVaultFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IYVaultTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IYVault.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IYVaultTransferIterator{contract: _IYVault.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IYVault *IYVaultFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IYVaultTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IYVault.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IYVaultTransfer)
				if err := _IYVault.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_IYVault *IYVaultFilterer) ParseTransfer(log types.Log) (*IYVaultTransfer, error) {
	event := new(IYVaultTransfer)
	if err := _IYVault.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
