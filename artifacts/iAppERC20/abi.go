// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iAppERC20

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

// IAppERC20MetaData contains all meta data concerning the IAppERC20 contract.
var IAppERC20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IAppERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use IAppERC20MetaData.ABI instead.
var IAppERC20ABI = IAppERC20MetaData.ABI

// IAppERC20 is an auto generated Go binding around an Ethereum contract.
type IAppERC20 struct {
	IAppERC20Caller     // Read-only binding to the contract
	IAppERC20Transactor // Write-only binding to the contract
	IAppERC20Filterer   // Log filterer for contract events
}

// IAppERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IAppERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAppERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IAppERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAppERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAppERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAppERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IAppERC20Session struct {
	Contract     *IAppERC20        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAppERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IAppERC20CallerSession struct {
	Contract *IAppERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IAppERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IAppERC20TransactorSession struct {
	Contract     *IAppERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IAppERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IAppERC20Raw struct {
	Contract *IAppERC20 // Generic contract binding to access the raw methods on
}

// IAppERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IAppERC20CallerRaw struct {
	Contract *IAppERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IAppERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IAppERC20TransactorRaw struct {
	Contract *IAppERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIAppERC20 creates a new instance of IAppERC20, bound to a specific deployed contract.
func NewIAppERC20(address common.Address, backend bind.ContractBackend) (*IAppERC20, error) {
	contract, err := bindIAppERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAppERC20{IAppERC20Caller: IAppERC20Caller{contract: contract}, IAppERC20Transactor: IAppERC20Transactor{contract: contract}, IAppERC20Filterer: IAppERC20Filterer{contract: contract}}, nil
}

// NewIAppERC20Caller creates a new read-only instance of IAppERC20, bound to a specific deployed contract.
func NewIAppERC20Caller(address common.Address, caller bind.ContractCaller) (*IAppERC20Caller, error) {
	contract, err := bindIAppERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAppERC20Caller{contract: contract}, nil
}

// NewIAppERC20Transactor creates a new write-only instance of IAppERC20, bound to a specific deployed contract.
func NewIAppERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IAppERC20Transactor, error) {
	contract, err := bindIAppERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAppERC20Transactor{contract: contract}, nil
}

// NewIAppERC20Filterer creates a new log filterer instance of IAppERC20, bound to a specific deployed contract.
func NewIAppERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IAppERC20Filterer, error) {
	contract, err := bindIAppERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAppERC20Filterer{contract: contract}, nil
}

// bindIAppERC20 binds a generic wrapper to an already deployed contract.
func bindIAppERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IAppERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAppERC20 *IAppERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAppERC20.Contract.IAppERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAppERC20 *IAppERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAppERC20.Contract.IAppERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAppERC20 *IAppERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAppERC20.Contract.IAppERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAppERC20 *IAppERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAppERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAppERC20 *IAppERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAppERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAppERC20 *IAppERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAppERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IAppERC20 *IAppERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IAppERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IAppERC20 *IAppERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IAppERC20.Contract.Allowance(&_IAppERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IAppERC20 *IAppERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IAppERC20.Contract.Allowance(&_IAppERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IAppERC20 *IAppERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IAppERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IAppERC20 *IAppERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _IAppERC20.Contract.BalanceOf(&_IAppERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IAppERC20 *IAppERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IAppERC20.Contract.BalanceOf(&_IAppERC20.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IAppERC20 *IAppERC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _IAppERC20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IAppERC20 *IAppERC20Session) Decimals() (uint8, error) {
	return _IAppERC20.Contract.Decimals(&_IAppERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IAppERC20 *IAppERC20CallerSession) Decimals() (uint8, error) {
	return _IAppERC20.Contract.Decimals(&_IAppERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IAppERC20 *IAppERC20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IAppERC20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IAppERC20 *IAppERC20Session) Name() (string, error) {
	return _IAppERC20.Contract.Name(&_IAppERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IAppERC20 *IAppERC20CallerSession) Name() (string, error) {
	return _IAppERC20.Contract.Name(&_IAppERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IAppERC20 *IAppERC20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IAppERC20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IAppERC20 *IAppERC20Session) Symbol() (string, error) {
	return _IAppERC20.Contract.Symbol(&_IAppERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IAppERC20 *IAppERC20CallerSession) Symbol() (string, error) {
	return _IAppERC20.Contract.Symbol(&_IAppERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IAppERC20 *IAppERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IAppERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IAppERC20 *IAppERC20Session) TotalSupply() (*big.Int, error) {
	return _IAppERC20.Contract.TotalSupply(&_IAppERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IAppERC20 *IAppERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IAppERC20.Contract.TotalSupply(&_IAppERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IAppERC20 *IAppERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IAppERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IAppERC20 *IAppERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IAppERC20.Contract.Approve(&_IAppERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IAppERC20 *IAppERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IAppERC20.Contract.Approve(&_IAppERC20.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IAppERC20 *IAppERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IAppERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IAppERC20 *IAppERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IAppERC20.Contract.Transfer(&_IAppERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IAppERC20 *IAppERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IAppERC20.Contract.Transfer(&_IAppERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IAppERC20 *IAppERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IAppERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IAppERC20 *IAppERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IAppERC20.Contract.TransferFrom(&_IAppERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IAppERC20 *IAppERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IAppERC20.Contract.TransferFrom(&_IAppERC20.TransactOpts, sender, recipient, amount)
}
