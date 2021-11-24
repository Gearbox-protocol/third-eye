// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iCreditAccount

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

// ICreditAccountMetaData contains all meta data concerning the ICreditAccount contract.
var ICreditAccountMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"swapContract\",\"type\":\"address\"}],\"name\":\"approveToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"borrowedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"targetContract\",\"type\":\"address\"}],\"name\":\"cancelAllowance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_creditManager\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_borrowedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_cumulativeIndexAtOpen\",\"type\":\"uint256\"}],\"name\":\"connectTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creditManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cumulativeIndexAtOpen\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"safeTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"since\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_borrowedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_cumulativeIndexAtOpen\",\"type\":\"uint256\"}],\"name\":\"updateParameters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ICreditAccountABI is the input ABI used to generate the binding from.
// Deprecated: Use ICreditAccountMetaData.ABI instead.
var ICreditAccountABI = ICreditAccountMetaData.ABI

// ICreditAccount is an auto generated Go binding around an Ethereum contract.
type ICreditAccount struct {
	ICreditAccountCaller     // Read-only binding to the contract
	ICreditAccountTransactor // Write-only binding to the contract
	ICreditAccountFilterer   // Log filterer for contract events
}

// ICreditAccountCaller is an auto generated read-only Go binding around an Ethereum contract.
type ICreditAccountCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICreditAccountTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ICreditAccountTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICreditAccountFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ICreditAccountFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICreditAccountSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ICreditAccountSession struct {
	Contract     *ICreditAccount   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ICreditAccountCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ICreditAccountCallerSession struct {
	Contract *ICreditAccountCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ICreditAccountTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ICreditAccountTransactorSession struct {
	Contract     *ICreditAccountTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ICreditAccountRaw is an auto generated low-level Go binding around an Ethereum contract.
type ICreditAccountRaw struct {
	Contract *ICreditAccount // Generic contract binding to access the raw methods on
}

// ICreditAccountCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ICreditAccountCallerRaw struct {
	Contract *ICreditAccountCaller // Generic read-only contract binding to access the raw methods on
}

// ICreditAccountTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ICreditAccountTransactorRaw struct {
	Contract *ICreditAccountTransactor // Generic write-only contract binding to access the raw methods on
}

// NewICreditAccount creates a new instance of ICreditAccount, bound to a specific deployed contract.
func NewICreditAccount(address common.Address, backend bind.ContractBackend) (*ICreditAccount, error) {
	contract, err := bindICreditAccount(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ICreditAccount{ICreditAccountCaller: ICreditAccountCaller{contract: contract}, ICreditAccountTransactor: ICreditAccountTransactor{contract: contract}, ICreditAccountFilterer: ICreditAccountFilterer{contract: contract}}, nil
}

// NewICreditAccountCaller creates a new read-only instance of ICreditAccount, bound to a specific deployed contract.
func NewICreditAccountCaller(address common.Address, caller bind.ContractCaller) (*ICreditAccountCaller, error) {
	contract, err := bindICreditAccount(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ICreditAccountCaller{contract: contract}, nil
}

// NewICreditAccountTransactor creates a new write-only instance of ICreditAccount, bound to a specific deployed contract.
func NewICreditAccountTransactor(address common.Address, transactor bind.ContractTransactor) (*ICreditAccountTransactor, error) {
	contract, err := bindICreditAccount(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ICreditAccountTransactor{contract: contract}, nil
}

// NewICreditAccountFilterer creates a new log filterer instance of ICreditAccount, bound to a specific deployed contract.
func NewICreditAccountFilterer(address common.Address, filterer bind.ContractFilterer) (*ICreditAccountFilterer, error) {
	contract, err := bindICreditAccount(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ICreditAccountFilterer{contract: contract}, nil
}

// bindICreditAccount binds a generic wrapper to an already deployed contract.
func bindICreditAccount(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ICreditAccountABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICreditAccount *ICreditAccountRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICreditAccount.Contract.ICreditAccountCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICreditAccount *ICreditAccountRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICreditAccount.Contract.ICreditAccountTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICreditAccount *ICreditAccountRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICreditAccount.Contract.ICreditAccountTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICreditAccount *ICreditAccountCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICreditAccount.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICreditAccount *ICreditAccountTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICreditAccount.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICreditAccount *ICreditAccountTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICreditAccount.Contract.contract.Transact(opts, method, params...)
}

// BorrowedAmount is a free data retrieval call binding the contract method 0x1afbb7a4.
//
// Solidity: function borrowedAmount() view returns(uint256)
func (_ICreditAccount *ICreditAccountCaller) BorrowedAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ICreditAccount.contract.Call(opts, &out, "borrowedAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowedAmount is a free data retrieval call binding the contract method 0x1afbb7a4.
//
// Solidity: function borrowedAmount() view returns(uint256)
func (_ICreditAccount *ICreditAccountSession) BorrowedAmount() (*big.Int, error) {
	return _ICreditAccount.Contract.BorrowedAmount(&_ICreditAccount.CallOpts)
}

// BorrowedAmount is a free data retrieval call binding the contract method 0x1afbb7a4.
//
// Solidity: function borrowedAmount() view returns(uint256)
func (_ICreditAccount *ICreditAccountCallerSession) BorrowedAmount() (*big.Int, error) {
	return _ICreditAccount.Contract.BorrowedAmount(&_ICreditAccount.CallOpts)
}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_ICreditAccount *ICreditAccountCaller) CreditManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ICreditAccount.contract.Call(opts, &out, "creditManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_ICreditAccount *ICreditAccountSession) CreditManager() (common.Address, error) {
	return _ICreditAccount.Contract.CreditManager(&_ICreditAccount.CallOpts)
}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_ICreditAccount *ICreditAccountCallerSession) CreditManager() (common.Address, error) {
	return _ICreditAccount.Contract.CreditManager(&_ICreditAccount.CallOpts)
}

// CumulativeIndexAtOpen is a free data retrieval call binding the contract method 0x17d11a15.
//
// Solidity: function cumulativeIndexAtOpen() view returns(uint256)
func (_ICreditAccount *ICreditAccountCaller) CumulativeIndexAtOpen(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ICreditAccount.contract.Call(opts, &out, "cumulativeIndexAtOpen")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CumulativeIndexAtOpen is a free data retrieval call binding the contract method 0x17d11a15.
//
// Solidity: function cumulativeIndexAtOpen() view returns(uint256)
func (_ICreditAccount *ICreditAccountSession) CumulativeIndexAtOpen() (*big.Int, error) {
	return _ICreditAccount.Contract.CumulativeIndexAtOpen(&_ICreditAccount.CallOpts)
}

// CumulativeIndexAtOpen is a free data retrieval call binding the contract method 0x17d11a15.
//
// Solidity: function cumulativeIndexAtOpen() view returns(uint256)
func (_ICreditAccount *ICreditAccountCallerSession) CumulativeIndexAtOpen() (*big.Int, error) {
	return _ICreditAccount.Contract.CumulativeIndexAtOpen(&_ICreditAccount.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_ICreditAccount *ICreditAccountCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ICreditAccount.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_ICreditAccount *ICreditAccountSession) Factory() (common.Address, error) {
	return _ICreditAccount.Contract.Factory(&_ICreditAccount.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_ICreditAccount *ICreditAccountCallerSession) Factory() (common.Address, error) {
	return _ICreditAccount.Contract.Factory(&_ICreditAccount.CallOpts)
}

// Since is a free data retrieval call binding the contract method 0x3dc54b40.
//
// Solidity: function since() view returns(uint256)
func (_ICreditAccount *ICreditAccountCaller) Since(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ICreditAccount.contract.Call(opts, &out, "since")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Since is a free data retrieval call binding the contract method 0x3dc54b40.
//
// Solidity: function since() view returns(uint256)
func (_ICreditAccount *ICreditAccountSession) Since() (*big.Int, error) {
	return _ICreditAccount.Contract.Since(&_ICreditAccount.CallOpts)
}

// Since is a free data retrieval call binding the contract method 0x3dc54b40.
//
// Solidity: function since() view returns(uint256)
func (_ICreditAccount *ICreditAccountCallerSession) Since() (*big.Int, error) {
	return _ICreditAccount.Contract.Since(&_ICreditAccount.CallOpts)
}

// ApproveToken is a paid mutator transaction binding the contract method 0x03105b04.
//
// Solidity: function approveToken(address token, address swapContract) returns()
func (_ICreditAccount *ICreditAccountTransactor) ApproveToken(opts *bind.TransactOpts, token common.Address, swapContract common.Address) (*types.Transaction, error) {
	return _ICreditAccount.contract.Transact(opts, "approveToken", token, swapContract)
}

// ApproveToken is a paid mutator transaction binding the contract method 0x03105b04.
//
// Solidity: function approveToken(address token, address swapContract) returns()
func (_ICreditAccount *ICreditAccountSession) ApproveToken(token common.Address, swapContract common.Address) (*types.Transaction, error) {
	return _ICreditAccount.Contract.ApproveToken(&_ICreditAccount.TransactOpts, token, swapContract)
}

// ApproveToken is a paid mutator transaction binding the contract method 0x03105b04.
//
// Solidity: function approveToken(address token, address swapContract) returns()
func (_ICreditAccount *ICreditAccountTransactorSession) ApproveToken(token common.Address, swapContract common.Address) (*types.Transaction, error) {
	return _ICreditAccount.Contract.ApproveToken(&_ICreditAccount.TransactOpts, token, swapContract)
}

// CancelAllowance is a paid mutator transaction binding the contract method 0x19a16039.
//
// Solidity: function cancelAllowance(address token, address targetContract) returns()
func (_ICreditAccount *ICreditAccountTransactor) CancelAllowance(opts *bind.TransactOpts, token common.Address, targetContract common.Address) (*types.Transaction, error) {
	return _ICreditAccount.contract.Transact(opts, "cancelAllowance", token, targetContract)
}

// CancelAllowance is a paid mutator transaction binding the contract method 0x19a16039.
//
// Solidity: function cancelAllowance(address token, address targetContract) returns()
func (_ICreditAccount *ICreditAccountSession) CancelAllowance(token common.Address, targetContract common.Address) (*types.Transaction, error) {
	return _ICreditAccount.Contract.CancelAllowance(&_ICreditAccount.TransactOpts, token, targetContract)
}

// CancelAllowance is a paid mutator transaction binding the contract method 0x19a16039.
//
// Solidity: function cancelAllowance(address token, address targetContract) returns()
func (_ICreditAccount *ICreditAccountTransactorSession) CancelAllowance(token common.Address, targetContract common.Address) (*types.Transaction, error) {
	return _ICreditAccount.Contract.CancelAllowance(&_ICreditAccount.TransactOpts, token, targetContract)
}

// ConnectTo is a paid mutator transaction binding the contract method 0xc75b5a71.
//
// Solidity: function connectTo(address _creditManager, uint256 _borrowedAmount, uint256 _cumulativeIndexAtOpen) returns()
func (_ICreditAccount *ICreditAccountTransactor) ConnectTo(opts *bind.TransactOpts, _creditManager common.Address, _borrowedAmount *big.Int, _cumulativeIndexAtOpen *big.Int) (*types.Transaction, error) {
	return _ICreditAccount.contract.Transact(opts, "connectTo", _creditManager, _borrowedAmount, _cumulativeIndexAtOpen)
}

// ConnectTo is a paid mutator transaction binding the contract method 0xc75b5a71.
//
// Solidity: function connectTo(address _creditManager, uint256 _borrowedAmount, uint256 _cumulativeIndexAtOpen) returns()
func (_ICreditAccount *ICreditAccountSession) ConnectTo(_creditManager common.Address, _borrowedAmount *big.Int, _cumulativeIndexAtOpen *big.Int) (*types.Transaction, error) {
	return _ICreditAccount.Contract.ConnectTo(&_ICreditAccount.TransactOpts, _creditManager, _borrowedAmount, _cumulativeIndexAtOpen)
}

// ConnectTo is a paid mutator transaction binding the contract method 0xc75b5a71.
//
// Solidity: function connectTo(address _creditManager, uint256 _borrowedAmount, uint256 _cumulativeIndexAtOpen) returns()
func (_ICreditAccount *ICreditAccountTransactorSession) ConnectTo(_creditManager common.Address, _borrowedAmount *big.Int, _cumulativeIndexAtOpen *big.Int) (*types.Transaction, error) {
	return _ICreditAccount.Contract.ConnectTo(&_ICreditAccount.TransactOpts, _creditManager, _borrowedAmount, _cumulativeIndexAtOpen)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address destination, bytes data) returns(bytes)
func (_ICreditAccount *ICreditAccountTransactor) Execute(opts *bind.TransactOpts, destination common.Address, data []byte) (*types.Transaction, error) {
	return _ICreditAccount.contract.Transact(opts, "execute", destination, data)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address destination, bytes data) returns(bytes)
func (_ICreditAccount *ICreditAccountSession) Execute(destination common.Address, data []byte) (*types.Transaction, error) {
	return _ICreditAccount.Contract.Execute(&_ICreditAccount.TransactOpts, destination, data)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address destination, bytes data) returns(bytes)
func (_ICreditAccount *ICreditAccountTransactorSession) Execute(destination common.Address, data []byte) (*types.Transaction, error) {
	return _ICreditAccount.Contract.Execute(&_ICreditAccount.TransactOpts, destination, data)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_ICreditAccount *ICreditAccountTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICreditAccount.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_ICreditAccount *ICreditAccountSession) Initialize() (*types.Transaction, error) {
	return _ICreditAccount.Contract.Initialize(&_ICreditAccount.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_ICreditAccount *ICreditAccountTransactorSession) Initialize() (*types.Transaction, error) {
	return _ICreditAccount.Contract.Initialize(&_ICreditAccount.TransactOpts)
}

// SafeTransfer is a paid mutator transaction binding the contract method 0xd1660f99.
//
// Solidity: function safeTransfer(address token, address to, uint256 amount) returns()
func (_ICreditAccount *ICreditAccountTransactor) SafeTransfer(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ICreditAccount.contract.Transact(opts, "safeTransfer", token, to, amount)
}

// SafeTransfer is a paid mutator transaction binding the contract method 0xd1660f99.
//
// Solidity: function safeTransfer(address token, address to, uint256 amount) returns()
func (_ICreditAccount *ICreditAccountSession) SafeTransfer(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ICreditAccount.Contract.SafeTransfer(&_ICreditAccount.TransactOpts, token, to, amount)
}

// SafeTransfer is a paid mutator transaction binding the contract method 0xd1660f99.
//
// Solidity: function safeTransfer(address token, address to, uint256 amount) returns()
func (_ICreditAccount *ICreditAccountTransactorSession) SafeTransfer(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ICreditAccount.Contract.SafeTransfer(&_ICreditAccount.TransactOpts, token, to, amount)
}

// UpdateParameters is a paid mutator transaction binding the contract method 0x16128211.
//
// Solidity: function updateParameters(uint256 _borrowedAmount, uint256 _cumulativeIndexAtOpen) returns()
func (_ICreditAccount *ICreditAccountTransactor) UpdateParameters(opts *bind.TransactOpts, _borrowedAmount *big.Int, _cumulativeIndexAtOpen *big.Int) (*types.Transaction, error) {
	return _ICreditAccount.contract.Transact(opts, "updateParameters", _borrowedAmount, _cumulativeIndexAtOpen)
}

// UpdateParameters is a paid mutator transaction binding the contract method 0x16128211.
//
// Solidity: function updateParameters(uint256 _borrowedAmount, uint256 _cumulativeIndexAtOpen) returns()
func (_ICreditAccount *ICreditAccountSession) UpdateParameters(_borrowedAmount *big.Int, _cumulativeIndexAtOpen *big.Int) (*types.Transaction, error) {
	return _ICreditAccount.Contract.UpdateParameters(&_ICreditAccount.TransactOpts, _borrowedAmount, _cumulativeIndexAtOpen)
}

// UpdateParameters is a paid mutator transaction binding the contract method 0x16128211.
//
// Solidity: function updateParameters(uint256 _borrowedAmount, uint256 _cumulativeIndexAtOpen) returns()
func (_ICreditAccount *ICreditAccountTransactorSession) UpdateParameters(_borrowedAmount *big.Int, _cumulativeIndexAtOpen *big.Int) (*types.Transaction, error) {
	return _ICreditAccount.Contract.UpdateParameters(&_ICreditAccount.TransactOpts, _borrowedAmount, _cumulativeIndexAtOpen)
}
