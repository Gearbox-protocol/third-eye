// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iAppCreditManager

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

// IAppCreditManagerMetaData contains all meta data concerning the IAppCreditManager contract.
var IAppCreditManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"addCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isLiquidated\",\"type\":\"bool\"}],\"name\":\"calcRepayAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.Exchange[]\",\"name\":\"paths\",\"type\":\"tuple[]\"}],\"name\":\"closeCreditAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultSwapContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"}],\"name\":\"getCreditAccountOrRevert\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"}],\"name\":\"hasOpenedCreditAccount\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"increaseBorrowedAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"leverageFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"referralCode\",\"type\":\"uint256\"}],\"name\":\"openCreditAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"repayCreditAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IAppCreditManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use IAppCreditManagerMetaData.ABI instead.
var IAppCreditManagerABI = IAppCreditManagerMetaData.ABI

// IAppCreditManager is an auto generated Go binding around an Ethereum contract.
type IAppCreditManager struct {
	IAppCreditManagerCaller     // Read-only binding to the contract
	IAppCreditManagerTransactor // Write-only binding to the contract
	IAppCreditManagerFilterer   // Log filterer for contract events
}

// IAppCreditManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IAppCreditManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAppCreditManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IAppCreditManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAppCreditManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAppCreditManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAppCreditManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IAppCreditManagerSession struct {
	Contract     *IAppCreditManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IAppCreditManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IAppCreditManagerCallerSession struct {
	Contract *IAppCreditManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// IAppCreditManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IAppCreditManagerTransactorSession struct {
	Contract     *IAppCreditManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IAppCreditManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IAppCreditManagerRaw struct {
	Contract *IAppCreditManager // Generic contract binding to access the raw methods on
}

// IAppCreditManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IAppCreditManagerCallerRaw struct {
	Contract *IAppCreditManagerCaller // Generic read-only contract binding to access the raw methods on
}

// IAppCreditManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IAppCreditManagerTransactorRaw struct {
	Contract *IAppCreditManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIAppCreditManager creates a new instance of IAppCreditManager, bound to a specific deployed contract.
func NewIAppCreditManager(address common.Address, backend bind.ContractBackend) (*IAppCreditManager, error) {
	contract, err := bindIAppCreditManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAppCreditManager{IAppCreditManagerCaller: IAppCreditManagerCaller{contract: contract}, IAppCreditManagerTransactor: IAppCreditManagerTransactor{contract: contract}, IAppCreditManagerFilterer: IAppCreditManagerFilterer{contract: contract}}, nil
}

// NewIAppCreditManagerCaller creates a new read-only instance of IAppCreditManager, bound to a specific deployed contract.
func NewIAppCreditManagerCaller(address common.Address, caller bind.ContractCaller) (*IAppCreditManagerCaller, error) {
	contract, err := bindIAppCreditManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAppCreditManagerCaller{contract: contract}, nil
}

// NewIAppCreditManagerTransactor creates a new write-only instance of IAppCreditManager, bound to a specific deployed contract.
func NewIAppCreditManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*IAppCreditManagerTransactor, error) {
	contract, err := bindIAppCreditManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAppCreditManagerTransactor{contract: contract}, nil
}

// NewIAppCreditManagerFilterer creates a new log filterer instance of IAppCreditManager, bound to a specific deployed contract.
func NewIAppCreditManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*IAppCreditManagerFilterer, error) {
	contract, err := bindIAppCreditManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAppCreditManagerFilterer{contract: contract}, nil
}

// bindIAppCreditManager binds a generic wrapper to an already deployed contract.
func bindIAppCreditManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IAppCreditManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAppCreditManager *IAppCreditManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAppCreditManager.Contract.IAppCreditManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAppCreditManager *IAppCreditManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAppCreditManager.Contract.IAppCreditManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAppCreditManager *IAppCreditManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAppCreditManager.Contract.IAppCreditManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAppCreditManager *IAppCreditManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAppCreditManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAppCreditManager *IAppCreditManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAppCreditManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAppCreditManager *IAppCreditManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAppCreditManager.Contract.contract.Transact(opts, method, params...)
}

// CalcRepayAmount is a free data retrieval call binding the contract method 0x3ce07355.
//
// Solidity: function calcRepayAmount(address borrower, bool isLiquidated) view returns(uint256)
func (_IAppCreditManager *IAppCreditManagerCaller) CalcRepayAmount(opts *bind.CallOpts, borrower common.Address, isLiquidated bool) (*big.Int, error) {
	var out []interface{}
	err := _IAppCreditManager.contract.Call(opts, &out, "calcRepayAmount", borrower, isLiquidated)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcRepayAmount is a free data retrieval call binding the contract method 0x3ce07355.
//
// Solidity: function calcRepayAmount(address borrower, bool isLiquidated) view returns(uint256)
func (_IAppCreditManager *IAppCreditManagerSession) CalcRepayAmount(borrower common.Address, isLiquidated bool) (*big.Int, error) {
	return _IAppCreditManager.Contract.CalcRepayAmount(&_IAppCreditManager.CallOpts, borrower, isLiquidated)
}

// CalcRepayAmount is a free data retrieval call binding the contract method 0x3ce07355.
//
// Solidity: function calcRepayAmount(address borrower, bool isLiquidated) view returns(uint256)
func (_IAppCreditManager *IAppCreditManagerCallerSession) CalcRepayAmount(borrower common.Address, isLiquidated bool) (*big.Int, error) {
	return _IAppCreditManager.Contract.CalcRepayAmount(&_IAppCreditManager.CallOpts, borrower, isLiquidated)
}

// DefaultSwapContract is a free data retrieval call binding the contract method 0xe0c011b7.
//
// Solidity: function defaultSwapContract() view returns(address)
func (_IAppCreditManager *IAppCreditManagerCaller) DefaultSwapContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IAppCreditManager.contract.Call(opts, &out, "defaultSwapContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DefaultSwapContract is a free data retrieval call binding the contract method 0xe0c011b7.
//
// Solidity: function defaultSwapContract() view returns(address)
func (_IAppCreditManager *IAppCreditManagerSession) DefaultSwapContract() (common.Address, error) {
	return _IAppCreditManager.Contract.DefaultSwapContract(&_IAppCreditManager.CallOpts)
}

// DefaultSwapContract is a free data retrieval call binding the contract method 0xe0c011b7.
//
// Solidity: function defaultSwapContract() view returns(address)
func (_IAppCreditManager *IAppCreditManagerCallerSession) DefaultSwapContract() (common.Address, error) {
	return _IAppCreditManager.Contract.DefaultSwapContract(&_IAppCreditManager.CallOpts)
}

// GetCreditAccountOrRevert is a free data retrieval call binding the contract method 0xe958b704.
//
// Solidity: function getCreditAccountOrRevert(address borrower) view returns(address)
func (_IAppCreditManager *IAppCreditManagerCaller) GetCreditAccountOrRevert(opts *bind.CallOpts, borrower common.Address) (common.Address, error) {
	var out []interface{}
	err := _IAppCreditManager.contract.Call(opts, &out, "getCreditAccountOrRevert", borrower)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCreditAccountOrRevert is a free data retrieval call binding the contract method 0xe958b704.
//
// Solidity: function getCreditAccountOrRevert(address borrower) view returns(address)
func (_IAppCreditManager *IAppCreditManagerSession) GetCreditAccountOrRevert(borrower common.Address) (common.Address, error) {
	return _IAppCreditManager.Contract.GetCreditAccountOrRevert(&_IAppCreditManager.CallOpts, borrower)
}

// GetCreditAccountOrRevert is a free data retrieval call binding the contract method 0xe958b704.
//
// Solidity: function getCreditAccountOrRevert(address borrower) view returns(address)
func (_IAppCreditManager *IAppCreditManagerCallerSession) GetCreditAccountOrRevert(borrower common.Address) (common.Address, error) {
	return _IAppCreditManager.Contract.GetCreditAccountOrRevert(&_IAppCreditManager.CallOpts, borrower)
}

// HasOpenedCreditAccount is a free data retrieval call binding the contract method 0x256ac915.
//
// Solidity: function hasOpenedCreditAccount(address borrower) view returns(bool)
func (_IAppCreditManager *IAppCreditManagerCaller) HasOpenedCreditAccount(opts *bind.CallOpts, borrower common.Address) (bool, error) {
	var out []interface{}
	err := _IAppCreditManager.contract.Call(opts, &out, "hasOpenedCreditAccount", borrower)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasOpenedCreditAccount is a free data retrieval call binding the contract method 0x256ac915.
//
// Solidity: function hasOpenedCreditAccount(address borrower) view returns(bool)
func (_IAppCreditManager *IAppCreditManagerSession) HasOpenedCreditAccount(borrower common.Address) (bool, error) {
	return _IAppCreditManager.Contract.HasOpenedCreditAccount(&_IAppCreditManager.CallOpts, borrower)
}

// HasOpenedCreditAccount is a free data retrieval call binding the contract method 0x256ac915.
//
// Solidity: function hasOpenedCreditAccount(address borrower) view returns(bool)
func (_IAppCreditManager *IAppCreditManagerCallerSession) HasOpenedCreditAccount(borrower common.Address) (bool, error) {
	return _IAppCreditManager.Contract.HasOpenedCreditAccount(&_IAppCreditManager.CallOpts, borrower)
}

// AddCollateral is a paid mutator transaction binding the contract method 0x59781034.
//
// Solidity: function addCollateral(address onBehalfOf, address token, uint256 amount) returns()
func (_IAppCreditManager *IAppCreditManagerTransactor) AddCollateral(opts *bind.TransactOpts, onBehalfOf common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IAppCreditManager.contract.Transact(opts, "addCollateral", onBehalfOf, token, amount)
}

// AddCollateral is a paid mutator transaction binding the contract method 0x59781034.
//
// Solidity: function addCollateral(address onBehalfOf, address token, uint256 amount) returns()
func (_IAppCreditManager *IAppCreditManagerSession) AddCollateral(onBehalfOf common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IAppCreditManager.Contract.AddCollateral(&_IAppCreditManager.TransactOpts, onBehalfOf, token, amount)
}

// AddCollateral is a paid mutator transaction binding the contract method 0x59781034.
//
// Solidity: function addCollateral(address onBehalfOf, address token, uint256 amount) returns()
func (_IAppCreditManager *IAppCreditManagerTransactorSession) AddCollateral(onBehalfOf common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IAppCreditManager.Contract.AddCollateral(&_IAppCreditManager.TransactOpts, onBehalfOf, token, amount)
}

// CloseCreditAccount is a paid mutator transaction binding the contract method 0xab114805.
//
// Solidity: function closeCreditAccount(address to, (address[],uint256)[] paths) returns()
func (_IAppCreditManager *IAppCreditManagerTransactor) CloseCreditAccount(opts *bind.TransactOpts, to common.Address, paths []DataTypesExchange) (*types.Transaction, error) {
	return _IAppCreditManager.contract.Transact(opts, "closeCreditAccount", to, paths)
}

// CloseCreditAccount is a paid mutator transaction binding the contract method 0xab114805.
//
// Solidity: function closeCreditAccount(address to, (address[],uint256)[] paths) returns()
func (_IAppCreditManager *IAppCreditManagerSession) CloseCreditAccount(to common.Address, paths []DataTypesExchange) (*types.Transaction, error) {
	return _IAppCreditManager.Contract.CloseCreditAccount(&_IAppCreditManager.TransactOpts, to, paths)
}

// CloseCreditAccount is a paid mutator transaction binding the contract method 0xab114805.
//
// Solidity: function closeCreditAccount(address to, (address[],uint256)[] paths) returns()
func (_IAppCreditManager *IAppCreditManagerTransactorSession) CloseCreditAccount(to common.Address, paths []DataTypesExchange) (*types.Transaction, error) {
	return _IAppCreditManager.Contract.CloseCreditAccount(&_IAppCreditManager.TransactOpts, to, paths)
}

// IncreaseBorrowedAmount is a paid mutator transaction binding the contract method 0x9efc60d0.
//
// Solidity: function increaseBorrowedAmount(uint256 amount) returns()
func (_IAppCreditManager *IAppCreditManagerTransactor) IncreaseBorrowedAmount(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _IAppCreditManager.contract.Transact(opts, "increaseBorrowedAmount", amount)
}

// IncreaseBorrowedAmount is a paid mutator transaction binding the contract method 0x9efc60d0.
//
// Solidity: function increaseBorrowedAmount(uint256 amount) returns()
func (_IAppCreditManager *IAppCreditManagerSession) IncreaseBorrowedAmount(amount *big.Int) (*types.Transaction, error) {
	return _IAppCreditManager.Contract.IncreaseBorrowedAmount(&_IAppCreditManager.TransactOpts, amount)
}

// IncreaseBorrowedAmount is a paid mutator transaction binding the contract method 0x9efc60d0.
//
// Solidity: function increaseBorrowedAmount(uint256 amount) returns()
func (_IAppCreditManager *IAppCreditManagerTransactorSession) IncreaseBorrowedAmount(amount *big.Int) (*types.Transaction, error) {
	return _IAppCreditManager.Contract.IncreaseBorrowedAmount(&_IAppCreditManager.TransactOpts, amount)
}

// OpenCreditAccount is a paid mutator transaction binding the contract method 0x5288ba4b.
//
// Solidity: function openCreditAccount(uint256 amount, address onBehalfOf, uint256 leverageFactor, uint256 referralCode) returns()
func (_IAppCreditManager *IAppCreditManagerTransactor) OpenCreditAccount(opts *bind.TransactOpts, amount *big.Int, onBehalfOf common.Address, leverageFactor *big.Int, referralCode *big.Int) (*types.Transaction, error) {
	return _IAppCreditManager.contract.Transact(opts, "openCreditAccount", amount, onBehalfOf, leverageFactor, referralCode)
}

// OpenCreditAccount is a paid mutator transaction binding the contract method 0x5288ba4b.
//
// Solidity: function openCreditAccount(uint256 amount, address onBehalfOf, uint256 leverageFactor, uint256 referralCode) returns()
func (_IAppCreditManager *IAppCreditManagerSession) OpenCreditAccount(amount *big.Int, onBehalfOf common.Address, leverageFactor *big.Int, referralCode *big.Int) (*types.Transaction, error) {
	return _IAppCreditManager.Contract.OpenCreditAccount(&_IAppCreditManager.TransactOpts, amount, onBehalfOf, leverageFactor, referralCode)
}

// OpenCreditAccount is a paid mutator transaction binding the contract method 0x5288ba4b.
//
// Solidity: function openCreditAccount(uint256 amount, address onBehalfOf, uint256 leverageFactor, uint256 referralCode) returns()
func (_IAppCreditManager *IAppCreditManagerTransactorSession) OpenCreditAccount(amount *big.Int, onBehalfOf common.Address, leverageFactor *big.Int, referralCode *big.Int) (*types.Transaction, error) {
	return _IAppCreditManager.Contract.OpenCreditAccount(&_IAppCreditManager.TransactOpts, amount, onBehalfOf, leverageFactor, referralCode)
}

// RepayCreditAccount is a paid mutator transaction binding the contract method 0xd692ba33.
//
// Solidity: function repayCreditAccount(address to) returns()
func (_IAppCreditManager *IAppCreditManagerTransactor) RepayCreditAccount(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _IAppCreditManager.contract.Transact(opts, "repayCreditAccount", to)
}

// RepayCreditAccount is a paid mutator transaction binding the contract method 0xd692ba33.
//
// Solidity: function repayCreditAccount(address to) returns()
func (_IAppCreditManager *IAppCreditManagerSession) RepayCreditAccount(to common.Address) (*types.Transaction, error) {
	return _IAppCreditManager.Contract.RepayCreditAccount(&_IAppCreditManager.TransactOpts, to)
}

// RepayCreditAccount is a paid mutator transaction binding the contract method 0xd692ba33.
//
// Solidity: function repayCreditAccount(address to) returns()
func (_IAppCreditManager *IAppCreditManagerTransactorSession) RepayCreditAccount(to common.Address) (*types.Transaction, error) {
	return _IAppCreditManager.Contract.RepayCreditAccount(&_IAppCreditManager.TransactOpts, to)
}
