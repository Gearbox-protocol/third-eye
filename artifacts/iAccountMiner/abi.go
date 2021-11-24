// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iAccountMiner

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

// IAccountMinerMetaData contains all meta data concerning the IAccountMiner contract.
var IAccountMinerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"kind\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"mineAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IAccountMinerABI is the input ABI used to generate the binding from.
// Deprecated: Use IAccountMinerMetaData.ABI instead.
var IAccountMinerABI = IAccountMinerMetaData.ABI

// IAccountMiner is an auto generated Go binding around an Ethereum contract.
type IAccountMiner struct {
	IAccountMinerCaller     // Read-only binding to the contract
	IAccountMinerTransactor // Write-only binding to the contract
	IAccountMinerFilterer   // Log filterer for contract events
}

// IAccountMinerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IAccountMinerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAccountMinerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IAccountMinerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAccountMinerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAccountMinerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAccountMinerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IAccountMinerSession struct {
	Contract     *IAccountMiner    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAccountMinerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IAccountMinerCallerSession struct {
	Contract *IAccountMinerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IAccountMinerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IAccountMinerTransactorSession struct {
	Contract     *IAccountMinerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IAccountMinerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IAccountMinerRaw struct {
	Contract *IAccountMiner // Generic contract binding to access the raw methods on
}

// IAccountMinerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IAccountMinerCallerRaw struct {
	Contract *IAccountMinerCaller // Generic read-only contract binding to access the raw methods on
}

// IAccountMinerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IAccountMinerTransactorRaw struct {
	Contract *IAccountMinerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIAccountMiner creates a new instance of IAccountMiner, bound to a specific deployed contract.
func NewIAccountMiner(address common.Address, backend bind.ContractBackend) (*IAccountMiner, error) {
	contract, err := bindIAccountMiner(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAccountMiner{IAccountMinerCaller: IAccountMinerCaller{contract: contract}, IAccountMinerTransactor: IAccountMinerTransactor{contract: contract}, IAccountMinerFilterer: IAccountMinerFilterer{contract: contract}}, nil
}

// NewIAccountMinerCaller creates a new read-only instance of IAccountMiner, bound to a specific deployed contract.
func NewIAccountMinerCaller(address common.Address, caller bind.ContractCaller) (*IAccountMinerCaller, error) {
	contract, err := bindIAccountMiner(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAccountMinerCaller{contract: contract}, nil
}

// NewIAccountMinerTransactor creates a new write-only instance of IAccountMiner, bound to a specific deployed contract.
func NewIAccountMinerTransactor(address common.Address, transactor bind.ContractTransactor) (*IAccountMinerTransactor, error) {
	contract, err := bindIAccountMiner(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAccountMinerTransactor{contract: contract}, nil
}

// NewIAccountMinerFilterer creates a new log filterer instance of IAccountMiner, bound to a specific deployed contract.
func NewIAccountMinerFilterer(address common.Address, filterer bind.ContractFilterer) (*IAccountMinerFilterer, error) {
	contract, err := bindIAccountMiner(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAccountMinerFilterer{contract: contract}, nil
}

// bindIAccountMiner binds a generic wrapper to an already deployed contract.
func bindIAccountMiner(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IAccountMinerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAccountMiner *IAccountMinerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAccountMiner.Contract.IAccountMinerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAccountMiner *IAccountMinerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAccountMiner.Contract.IAccountMinerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAccountMiner *IAccountMinerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAccountMiner.Contract.IAccountMinerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAccountMiner *IAccountMinerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAccountMiner.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAccountMiner *IAccountMinerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAccountMiner.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAccountMiner *IAccountMinerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAccountMiner.Contract.contract.Transact(opts, method, params...)
}

// Kind is a free data retrieval call binding the contract method 0x04baa00b.
//
// Solidity: function kind() pure returns(bytes32)
func (_IAccountMiner *IAccountMinerCaller) Kind(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IAccountMiner.contract.Call(opts, &out, "kind")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Kind is a free data retrieval call binding the contract method 0x04baa00b.
//
// Solidity: function kind() pure returns(bytes32)
func (_IAccountMiner *IAccountMinerSession) Kind() ([32]byte, error) {
	return _IAccountMiner.Contract.Kind(&_IAccountMiner.CallOpts)
}

// Kind is a free data retrieval call binding the contract method 0x04baa00b.
//
// Solidity: function kind() pure returns(bytes32)
func (_IAccountMiner *IAccountMinerCallerSession) Kind() ([32]byte, error) {
	return _IAccountMiner.Contract.Kind(&_IAccountMiner.CallOpts)
}

// MineAccount is a paid mutator transaction binding the contract method 0x9849e42f.
//
// Solidity: function mineAccount(address user) returns()
func (_IAccountMiner *IAccountMinerTransactor) MineAccount(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _IAccountMiner.contract.Transact(opts, "mineAccount", user)
}

// MineAccount is a paid mutator transaction binding the contract method 0x9849e42f.
//
// Solidity: function mineAccount(address user) returns()
func (_IAccountMiner *IAccountMinerSession) MineAccount(user common.Address) (*types.Transaction, error) {
	return _IAccountMiner.Contract.MineAccount(&_IAccountMiner.TransactOpts, user)
}

// MineAccount is a paid mutator transaction binding the contract method 0x9849e42f.
//
// Solidity: function mineAccount(address user) returns()
func (_IAccountMiner *IAccountMinerTransactorSession) MineAccount(user common.Address) (*types.Transaction, error) {
	return _IAccountMiner.Contract.MineAccount(&_IAccountMiner.TransactOpts, user)
}
