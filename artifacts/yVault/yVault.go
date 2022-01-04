// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package yVault

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// YVaultABI is the input ABI used to generate the binding from.
const YVaultABI = "[{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pricePerShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// YVault is an auto generated Go binding around an Ethereum contract.
type YVault struct {
	YVaultCaller     // Read-only binding to the contract
	YVaultTransactor // Write-only binding to the contract
	YVaultFilterer   // Log filterer for contract events
}

// YVaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type YVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YVaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type YVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YVaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type YVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YVaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type YVaultSession struct {
	Contract     *YVault           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// YVaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type YVaultCallerSession struct {
	Contract *YVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// YVaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type YVaultTransactorSession struct {
	Contract     *YVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// YVaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type YVaultRaw struct {
	Contract *YVault // Generic contract binding to access the raw methods on
}

// YVaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type YVaultCallerRaw struct {
	Contract *YVaultCaller // Generic read-only contract binding to access the raw methods on
}

// YVaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type YVaultTransactorRaw struct {
	Contract *YVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewYVault creates a new instance of YVault, bound to a specific deployed contract.
func NewYVault(address common.Address, backend bind.ContractBackend) (*YVault, error) {
	contract, err := bindYVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &YVault{YVaultCaller: YVaultCaller{contract: contract}, YVaultTransactor: YVaultTransactor{contract: contract}, YVaultFilterer: YVaultFilterer{contract: contract}}, nil
}

// NewYVaultCaller creates a new read-only instance of YVault, bound to a specific deployed contract.
func NewYVaultCaller(address common.Address, caller bind.ContractCaller) (*YVaultCaller, error) {
	contract, err := bindYVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &YVaultCaller{contract: contract}, nil
}

// NewYVaultTransactor creates a new write-only instance of YVault, bound to a specific deployed contract.
func NewYVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*YVaultTransactor, error) {
	contract, err := bindYVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &YVaultTransactor{contract: contract}, nil
}

// NewYVaultFilterer creates a new log filterer instance of YVault, bound to a specific deployed contract.
func NewYVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*YVaultFilterer, error) {
	contract, err := bindYVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &YVaultFilterer{contract: contract}, nil
}

// bindYVault binds a generic wrapper to an already deployed contract.
func bindYVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(YVaultABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YVault *YVaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YVault.Contract.YVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YVault *YVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YVault.Contract.YVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YVault *YVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YVault.Contract.YVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YVault *YVaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YVault *YVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YVault *YVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YVault.Contract.contract.Transact(opts, method, params...)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_YVault *YVaultCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _YVault.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_YVault *YVaultSession) Decimals() (uint8, error) {
	return _YVault.Contract.Decimals(&_YVault.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_YVault *YVaultCallerSession) Decimals() (uint8, error) {
	return _YVault.Contract.Decimals(&_YVault.CallOpts)
}

// PricePerShare is a free data retrieval call binding the contract method 0x99530b06.
//
// Solidity: function pricePerShare() view returns(uint256)
func (_YVault *YVaultCaller) PricePerShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YVault.contract.Call(opts, &out, "pricePerShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PricePerShare is a free data retrieval call binding the contract method 0x99530b06.
//
// Solidity: function pricePerShare() view returns(uint256)
func (_YVault *YVaultSession) PricePerShare() (*big.Int, error) {
	return _YVault.Contract.PricePerShare(&_YVault.CallOpts)
}

// PricePerShare is a free data retrieval call binding the contract method 0x99530b06.
//
// Solidity: function pricePerShare() view returns(uint256)
func (_YVault *YVaultCallerSession) PricePerShare() (*big.Int, error) {
	return _YVault.Contract.PricePerShare(&_YVault.CallOpts)
}
