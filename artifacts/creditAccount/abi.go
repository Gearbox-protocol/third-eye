// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package creditAccount

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

// CreditAccountMetaData contains all meta data concerning the CreditAccount contract.
var CreditAccountMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"swapContract\",\"type\":\"address\"}],\"name\":\"approveToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"borrowedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"targetContract\",\"type\":\"address\"}],\"name\":\"cancelAllowance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_creditManager\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_borrowedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_cumulativeIndexAtOpen\",\"type\":\"uint256\"}],\"name\":\"connectTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creditManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cumulativeIndexAtOpen\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"safeTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"since\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_borrowedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_cumulativeIndexAtOpen\",\"type\":\"uint256\"}],\"name\":\"updateParameters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// CreditAccountABI is the input ABI used to generate the binding from.
// Deprecated: Use CreditAccountMetaData.ABI instead.
var CreditAccountABI = CreditAccountMetaData.ABI

// CreditAccount is an auto generated Go binding around an Ethereum contract.
type CreditAccount struct {
	CreditAccountCaller     // Read-only binding to the contract
	CreditAccountTransactor // Write-only binding to the contract
	CreditAccountFilterer   // Log filterer for contract events
}

// CreditAccountCaller is an auto generated read-only Go binding around an Ethereum contract.
type CreditAccountCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditAccountTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CreditAccountTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditAccountFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CreditAccountFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditAccountSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CreditAccountSession struct {
	Contract     *CreditAccount    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CreditAccountCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CreditAccountCallerSession struct {
	Contract *CreditAccountCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// CreditAccountTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CreditAccountTransactorSession struct {
	Contract     *CreditAccountTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// CreditAccountRaw is an auto generated low-level Go binding around an Ethereum contract.
type CreditAccountRaw struct {
	Contract *CreditAccount // Generic contract binding to access the raw methods on
}

// CreditAccountCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CreditAccountCallerRaw struct {
	Contract *CreditAccountCaller // Generic read-only contract binding to access the raw methods on
}

// CreditAccountTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CreditAccountTransactorRaw struct {
	Contract *CreditAccountTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCreditAccount creates a new instance of CreditAccount, bound to a specific deployed contract.
func NewCreditAccount(address common.Address, backend bind.ContractBackend) (*CreditAccount, error) {
	contract, err := bindCreditAccount(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CreditAccount{CreditAccountCaller: CreditAccountCaller{contract: contract}, CreditAccountTransactor: CreditAccountTransactor{contract: contract}, CreditAccountFilterer: CreditAccountFilterer{contract: contract}}, nil
}

// NewCreditAccountCaller creates a new read-only instance of CreditAccount, bound to a specific deployed contract.
func NewCreditAccountCaller(address common.Address, caller bind.ContractCaller) (*CreditAccountCaller, error) {
	contract, err := bindCreditAccount(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CreditAccountCaller{contract: contract}, nil
}

// NewCreditAccountTransactor creates a new write-only instance of CreditAccount, bound to a specific deployed contract.
func NewCreditAccountTransactor(address common.Address, transactor bind.ContractTransactor) (*CreditAccountTransactor, error) {
	contract, err := bindCreditAccount(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CreditAccountTransactor{contract: contract}, nil
}

// NewCreditAccountFilterer creates a new log filterer instance of CreditAccount, bound to a specific deployed contract.
func NewCreditAccountFilterer(address common.Address, filterer bind.ContractFilterer) (*CreditAccountFilterer, error) {
	contract, err := bindCreditAccount(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CreditAccountFilterer{contract: contract}, nil
}

// bindCreditAccount binds a generic wrapper to an already deployed contract.
func bindCreditAccount(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CreditAccountABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CreditAccount *CreditAccountRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CreditAccount.Contract.CreditAccountCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CreditAccount *CreditAccountRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditAccount.Contract.CreditAccountTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CreditAccount *CreditAccountRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CreditAccount.Contract.CreditAccountTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CreditAccount *CreditAccountCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CreditAccount.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CreditAccount *CreditAccountTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditAccount.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CreditAccount *CreditAccountTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CreditAccount.Contract.contract.Transact(opts, method, params...)
}

// BorrowedAmount is a free data retrieval call binding the contract method 0x1afbb7a4.
//
// Solidity: function borrowedAmount() view returns(uint256)
func (_CreditAccount *CreditAccountCaller) BorrowedAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CreditAccount.contract.Call(opts, &out, "borrowedAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowedAmount is a free data retrieval call binding the contract method 0x1afbb7a4.
//
// Solidity: function borrowedAmount() view returns(uint256)
func (_CreditAccount *CreditAccountSession) BorrowedAmount() (*big.Int, error) {
	return _CreditAccount.Contract.BorrowedAmount(&_CreditAccount.CallOpts)
}

// BorrowedAmount is a free data retrieval call binding the contract method 0x1afbb7a4.
//
// Solidity: function borrowedAmount() view returns(uint256)
func (_CreditAccount *CreditAccountCallerSession) BorrowedAmount() (*big.Int, error) {
	return _CreditAccount.Contract.BorrowedAmount(&_CreditAccount.CallOpts)
}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_CreditAccount *CreditAccountCaller) CreditManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditAccount.contract.Call(opts, &out, "creditManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_CreditAccount *CreditAccountSession) CreditManager() (common.Address, error) {
	return _CreditAccount.Contract.CreditManager(&_CreditAccount.CallOpts)
}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_CreditAccount *CreditAccountCallerSession) CreditManager() (common.Address, error) {
	return _CreditAccount.Contract.CreditManager(&_CreditAccount.CallOpts)
}

// CumulativeIndexAtOpen is a free data retrieval call binding the contract method 0x17d11a15.
//
// Solidity: function cumulativeIndexAtOpen() view returns(uint256)
func (_CreditAccount *CreditAccountCaller) CumulativeIndexAtOpen(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CreditAccount.contract.Call(opts, &out, "cumulativeIndexAtOpen")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CumulativeIndexAtOpen is a free data retrieval call binding the contract method 0x17d11a15.
//
// Solidity: function cumulativeIndexAtOpen() view returns(uint256)
func (_CreditAccount *CreditAccountSession) CumulativeIndexAtOpen() (*big.Int, error) {
	return _CreditAccount.Contract.CumulativeIndexAtOpen(&_CreditAccount.CallOpts)
}

// CumulativeIndexAtOpen is a free data retrieval call binding the contract method 0x17d11a15.
//
// Solidity: function cumulativeIndexAtOpen() view returns(uint256)
func (_CreditAccount *CreditAccountCallerSession) CumulativeIndexAtOpen() (*big.Int, error) {
	return _CreditAccount.Contract.CumulativeIndexAtOpen(&_CreditAccount.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_CreditAccount *CreditAccountCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditAccount.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_CreditAccount *CreditAccountSession) Factory() (common.Address, error) {
	return _CreditAccount.Contract.Factory(&_CreditAccount.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_CreditAccount *CreditAccountCallerSession) Factory() (common.Address, error) {
	return _CreditAccount.Contract.Factory(&_CreditAccount.CallOpts)
}

// Since is a free data retrieval call binding the contract method 0x3dc54b40.
//
// Solidity: function since() view returns(uint256)
func (_CreditAccount *CreditAccountCaller) Since(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CreditAccount.contract.Call(opts, &out, "since")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Since is a free data retrieval call binding the contract method 0x3dc54b40.
//
// Solidity: function since() view returns(uint256)
func (_CreditAccount *CreditAccountSession) Since() (*big.Int, error) {
	return _CreditAccount.Contract.Since(&_CreditAccount.CallOpts)
}

// Since is a free data retrieval call binding the contract method 0x3dc54b40.
//
// Solidity: function since() view returns(uint256)
func (_CreditAccount *CreditAccountCallerSession) Since() (*big.Int, error) {
	return _CreditAccount.Contract.Since(&_CreditAccount.CallOpts)
}

// ApproveToken is a paid mutator transaction binding the contract method 0x03105b04.
//
// Solidity: function approveToken(address token, address swapContract) returns()
func (_CreditAccount *CreditAccountTransactor) ApproveToken(opts *bind.TransactOpts, token common.Address, swapContract common.Address) (*types.Transaction, error) {
	return _CreditAccount.contract.Transact(opts, "approveToken", token, swapContract)
}

// ApproveToken is a paid mutator transaction binding the contract method 0x03105b04.
//
// Solidity: function approveToken(address token, address swapContract) returns()
func (_CreditAccount *CreditAccountSession) ApproveToken(token common.Address, swapContract common.Address) (*types.Transaction, error) {
	return _CreditAccount.Contract.ApproveToken(&_CreditAccount.TransactOpts, token, swapContract)
}

// ApproveToken is a paid mutator transaction binding the contract method 0x03105b04.
//
// Solidity: function approveToken(address token, address swapContract) returns()
func (_CreditAccount *CreditAccountTransactorSession) ApproveToken(token common.Address, swapContract common.Address) (*types.Transaction, error) {
	return _CreditAccount.Contract.ApproveToken(&_CreditAccount.TransactOpts, token, swapContract)
}

// CancelAllowance is a paid mutator transaction binding the contract method 0x19a16039.
//
// Solidity: function cancelAllowance(address token, address targetContract) returns()
func (_CreditAccount *CreditAccountTransactor) CancelAllowance(opts *bind.TransactOpts, token common.Address, targetContract common.Address) (*types.Transaction, error) {
	return _CreditAccount.contract.Transact(opts, "cancelAllowance", token, targetContract)
}

// CancelAllowance is a paid mutator transaction binding the contract method 0x19a16039.
//
// Solidity: function cancelAllowance(address token, address targetContract) returns()
func (_CreditAccount *CreditAccountSession) CancelAllowance(token common.Address, targetContract common.Address) (*types.Transaction, error) {
	return _CreditAccount.Contract.CancelAllowance(&_CreditAccount.TransactOpts, token, targetContract)
}

// CancelAllowance is a paid mutator transaction binding the contract method 0x19a16039.
//
// Solidity: function cancelAllowance(address token, address targetContract) returns()
func (_CreditAccount *CreditAccountTransactorSession) CancelAllowance(token common.Address, targetContract common.Address) (*types.Transaction, error) {
	return _CreditAccount.Contract.CancelAllowance(&_CreditAccount.TransactOpts, token, targetContract)
}

// ConnectTo is a paid mutator transaction binding the contract method 0xc75b5a71.
//
// Solidity: function connectTo(address _creditManager, uint256 _borrowedAmount, uint256 _cumulativeIndexAtOpen) returns()
func (_CreditAccount *CreditAccountTransactor) ConnectTo(opts *bind.TransactOpts, _creditManager common.Address, _borrowedAmount *big.Int, _cumulativeIndexAtOpen *big.Int) (*types.Transaction, error) {
	return _CreditAccount.contract.Transact(opts, "connectTo", _creditManager, _borrowedAmount, _cumulativeIndexAtOpen)
}

// ConnectTo is a paid mutator transaction binding the contract method 0xc75b5a71.
//
// Solidity: function connectTo(address _creditManager, uint256 _borrowedAmount, uint256 _cumulativeIndexAtOpen) returns()
func (_CreditAccount *CreditAccountSession) ConnectTo(_creditManager common.Address, _borrowedAmount *big.Int, _cumulativeIndexAtOpen *big.Int) (*types.Transaction, error) {
	return _CreditAccount.Contract.ConnectTo(&_CreditAccount.TransactOpts, _creditManager, _borrowedAmount, _cumulativeIndexAtOpen)
}

// ConnectTo is a paid mutator transaction binding the contract method 0xc75b5a71.
//
// Solidity: function connectTo(address _creditManager, uint256 _borrowedAmount, uint256 _cumulativeIndexAtOpen) returns()
func (_CreditAccount *CreditAccountTransactorSession) ConnectTo(_creditManager common.Address, _borrowedAmount *big.Int, _cumulativeIndexAtOpen *big.Int) (*types.Transaction, error) {
	return _CreditAccount.Contract.ConnectTo(&_CreditAccount.TransactOpts, _creditManager, _borrowedAmount, _cumulativeIndexAtOpen)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address destination, bytes data) returns(bytes)
func (_CreditAccount *CreditAccountTransactor) Execute(opts *bind.TransactOpts, destination common.Address, data []byte) (*types.Transaction, error) {
	return _CreditAccount.contract.Transact(opts, "execute", destination, data)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address destination, bytes data) returns(bytes)
func (_CreditAccount *CreditAccountSession) Execute(destination common.Address, data []byte) (*types.Transaction, error) {
	return _CreditAccount.Contract.Execute(&_CreditAccount.TransactOpts, destination, data)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address destination, bytes data) returns(bytes)
func (_CreditAccount *CreditAccountTransactorSession) Execute(destination common.Address, data []byte) (*types.Transaction, error) {
	return _CreditAccount.Contract.Execute(&_CreditAccount.TransactOpts, destination, data)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_CreditAccount *CreditAccountTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditAccount.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_CreditAccount *CreditAccountSession) Initialize() (*types.Transaction, error) {
	return _CreditAccount.Contract.Initialize(&_CreditAccount.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_CreditAccount *CreditAccountTransactorSession) Initialize() (*types.Transaction, error) {
	return _CreditAccount.Contract.Initialize(&_CreditAccount.TransactOpts)
}

// SafeTransfer is a paid mutator transaction binding the contract method 0xd1660f99.
//
// Solidity: function safeTransfer(address token, address to, uint256 amount) returns()
func (_CreditAccount *CreditAccountTransactor) SafeTransfer(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CreditAccount.contract.Transact(opts, "safeTransfer", token, to, amount)
}

// SafeTransfer is a paid mutator transaction binding the contract method 0xd1660f99.
//
// Solidity: function safeTransfer(address token, address to, uint256 amount) returns()
func (_CreditAccount *CreditAccountSession) SafeTransfer(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CreditAccount.Contract.SafeTransfer(&_CreditAccount.TransactOpts, token, to, amount)
}

// SafeTransfer is a paid mutator transaction binding the contract method 0xd1660f99.
//
// Solidity: function safeTransfer(address token, address to, uint256 amount) returns()
func (_CreditAccount *CreditAccountTransactorSession) SafeTransfer(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CreditAccount.Contract.SafeTransfer(&_CreditAccount.TransactOpts, token, to, amount)
}

// UpdateParameters is a paid mutator transaction binding the contract method 0x16128211.
//
// Solidity: function updateParameters(uint256 _borrowedAmount, uint256 _cumulativeIndexAtOpen) returns()
func (_CreditAccount *CreditAccountTransactor) UpdateParameters(opts *bind.TransactOpts, _borrowedAmount *big.Int, _cumulativeIndexAtOpen *big.Int) (*types.Transaction, error) {
	return _CreditAccount.contract.Transact(opts, "updateParameters", _borrowedAmount, _cumulativeIndexAtOpen)
}

// UpdateParameters is a paid mutator transaction binding the contract method 0x16128211.
//
// Solidity: function updateParameters(uint256 _borrowedAmount, uint256 _cumulativeIndexAtOpen) returns()
func (_CreditAccount *CreditAccountSession) UpdateParameters(_borrowedAmount *big.Int, _cumulativeIndexAtOpen *big.Int) (*types.Transaction, error) {
	return _CreditAccount.Contract.UpdateParameters(&_CreditAccount.TransactOpts, _borrowedAmount, _cumulativeIndexAtOpen)
}

// UpdateParameters is a paid mutator transaction binding the contract method 0x16128211.
//
// Solidity: function updateParameters(uint256 _borrowedAmount, uint256 _cumulativeIndexAtOpen) returns()
func (_CreditAccount *CreditAccountTransactorSession) UpdateParameters(_borrowedAmount *big.Int, _cumulativeIndexAtOpen *big.Int) (*types.Transaction, error) {
	return _CreditAccount.Contract.UpdateParameters(&_CreditAccount.TransactOpts, _borrowedAmount, _cumulativeIndexAtOpen)
}
