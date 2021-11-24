// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package flashLoanAttacker

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

// DataTypesExchange is an auto generated low-level Go binding around an user-defined struct.
type DataTypesExchange struct {
	Path         []common.Address
	AmountOutMin *big.Int
}

// FlashLoanAttackerMetaData contains all meta data concerning the FlashLoanAttacker contract.
var FlashLoanAttackerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditManager\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"leverage\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.Exchange[]\",\"name\":\"paths\",\"type\":\"tuple[]\"}],\"name\":\"attackClose\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"leverage\",\"type\":\"uint256\"}],\"name\":\"attackRepay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// FlashLoanAttackerABI is the input ABI used to generate the binding from.
// Deprecated: Use FlashLoanAttackerMetaData.ABI instead.
var FlashLoanAttackerABI = FlashLoanAttackerMetaData.ABI

// FlashLoanAttacker is an auto generated Go binding around an Ethereum contract.
type FlashLoanAttacker struct {
	FlashLoanAttackerCaller     // Read-only binding to the contract
	FlashLoanAttackerTransactor // Write-only binding to the contract
	FlashLoanAttackerFilterer   // Log filterer for contract events
}

// FlashLoanAttackerCaller is an auto generated read-only Go binding around an Ethereum contract.
type FlashLoanAttackerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlashLoanAttackerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FlashLoanAttackerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlashLoanAttackerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FlashLoanAttackerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlashLoanAttackerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FlashLoanAttackerSession struct {
	Contract     *FlashLoanAttacker // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// FlashLoanAttackerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FlashLoanAttackerCallerSession struct {
	Contract *FlashLoanAttackerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// FlashLoanAttackerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FlashLoanAttackerTransactorSession struct {
	Contract     *FlashLoanAttackerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// FlashLoanAttackerRaw is an auto generated low-level Go binding around an Ethereum contract.
type FlashLoanAttackerRaw struct {
	Contract *FlashLoanAttacker // Generic contract binding to access the raw methods on
}

// FlashLoanAttackerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FlashLoanAttackerCallerRaw struct {
	Contract *FlashLoanAttackerCaller // Generic read-only contract binding to access the raw methods on
}

// FlashLoanAttackerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FlashLoanAttackerTransactorRaw struct {
	Contract *FlashLoanAttackerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFlashLoanAttacker creates a new instance of FlashLoanAttacker, bound to a specific deployed contract.
func NewFlashLoanAttacker(address common.Address, backend bind.ContractBackend) (*FlashLoanAttacker, error) {
	contract, err := bindFlashLoanAttacker(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FlashLoanAttacker{FlashLoanAttackerCaller: FlashLoanAttackerCaller{contract: contract}, FlashLoanAttackerTransactor: FlashLoanAttackerTransactor{contract: contract}, FlashLoanAttackerFilterer: FlashLoanAttackerFilterer{contract: contract}}, nil
}

// NewFlashLoanAttackerCaller creates a new read-only instance of FlashLoanAttacker, bound to a specific deployed contract.
func NewFlashLoanAttackerCaller(address common.Address, caller bind.ContractCaller) (*FlashLoanAttackerCaller, error) {
	contract, err := bindFlashLoanAttacker(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FlashLoanAttackerCaller{contract: contract}, nil
}

// NewFlashLoanAttackerTransactor creates a new write-only instance of FlashLoanAttacker, bound to a specific deployed contract.
func NewFlashLoanAttackerTransactor(address common.Address, transactor bind.ContractTransactor) (*FlashLoanAttackerTransactor, error) {
	contract, err := bindFlashLoanAttacker(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FlashLoanAttackerTransactor{contract: contract}, nil
}

// NewFlashLoanAttackerFilterer creates a new log filterer instance of FlashLoanAttacker, bound to a specific deployed contract.
func NewFlashLoanAttackerFilterer(address common.Address, filterer bind.ContractFilterer) (*FlashLoanAttackerFilterer, error) {
	contract, err := bindFlashLoanAttacker(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FlashLoanAttackerFilterer{contract: contract}, nil
}

// bindFlashLoanAttacker binds a generic wrapper to an already deployed contract.
func bindFlashLoanAttacker(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FlashLoanAttackerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FlashLoanAttacker *FlashLoanAttackerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FlashLoanAttacker.Contract.FlashLoanAttackerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FlashLoanAttacker *FlashLoanAttackerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FlashLoanAttacker.Contract.FlashLoanAttackerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FlashLoanAttacker *FlashLoanAttackerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FlashLoanAttacker.Contract.FlashLoanAttackerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FlashLoanAttacker *FlashLoanAttackerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FlashLoanAttacker.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FlashLoanAttacker *FlashLoanAttackerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FlashLoanAttacker.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FlashLoanAttacker *FlashLoanAttackerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FlashLoanAttacker.Contract.contract.Transact(opts, method, params...)
}

// AttackClose is a paid mutator transaction binding the contract method 0x1c1fab63.
//
// Solidity: function attackClose(uint256 amount, uint256 leverage, (address[],uint256)[] paths) returns()
func (_FlashLoanAttacker *FlashLoanAttackerTransactor) AttackClose(opts *bind.TransactOpts, amount *big.Int, leverage *big.Int, paths []DataTypesExchange) (*types.Transaction, error) {
	return _FlashLoanAttacker.contract.Transact(opts, "attackClose", amount, leverage, paths)
}

// AttackClose is a paid mutator transaction binding the contract method 0x1c1fab63.
//
// Solidity: function attackClose(uint256 amount, uint256 leverage, (address[],uint256)[] paths) returns()
func (_FlashLoanAttacker *FlashLoanAttackerSession) AttackClose(amount *big.Int, leverage *big.Int, paths []DataTypesExchange) (*types.Transaction, error) {
	return _FlashLoanAttacker.Contract.AttackClose(&_FlashLoanAttacker.TransactOpts, amount, leverage, paths)
}

// AttackClose is a paid mutator transaction binding the contract method 0x1c1fab63.
//
// Solidity: function attackClose(uint256 amount, uint256 leverage, (address[],uint256)[] paths) returns()
func (_FlashLoanAttacker *FlashLoanAttackerTransactorSession) AttackClose(amount *big.Int, leverage *big.Int, paths []DataTypesExchange) (*types.Transaction, error) {
	return _FlashLoanAttacker.Contract.AttackClose(&_FlashLoanAttacker.TransactOpts, amount, leverage, paths)
}

// AttackRepay is a paid mutator transaction binding the contract method 0xe6b8925a.
//
// Solidity: function attackRepay(uint256 amount, uint256 leverage) returns()
func (_FlashLoanAttacker *FlashLoanAttackerTransactor) AttackRepay(opts *bind.TransactOpts, amount *big.Int, leverage *big.Int) (*types.Transaction, error) {
	return _FlashLoanAttacker.contract.Transact(opts, "attackRepay", amount, leverage)
}

// AttackRepay is a paid mutator transaction binding the contract method 0xe6b8925a.
//
// Solidity: function attackRepay(uint256 amount, uint256 leverage) returns()
func (_FlashLoanAttacker *FlashLoanAttackerSession) AttackRepay(amount *big.Int, leverage *big.Int) (*types.Transaction, error) {
	return _FlashLoanAttacker.Contract.AttackRepay(&_FlashLoanAttacker.TransactOpts, amount, leverage)
}

// AttackRepay is a paid mutator transaction binding the contract method 0xe6b8925a.
//
// Solidity: function attackRepay(uint256 amount, uint256 leverage) returns()
func (_FlashLoanAttacker *FlashLoanAttackerTransactorSession) AttackRepay(amount *big.Int, leverage *big.Int) (*types.Transaction, error) {
	return _FlashLoanAttacker.Contract.AttackRepay(&_FlashLoanAttacker.TransactOpts, amount, leverage)
}
