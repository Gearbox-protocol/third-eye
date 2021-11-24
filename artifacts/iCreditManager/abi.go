// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iCreditManager

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

// ICreditManagerMetaData contains all meta data concerning the ICreditManager contract.
var ICreditManagerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"AddCollateral\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"remainingFunds\",\"type\":\"uint256\"}],\"name\":\"CloseCreditAccount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"ExecuteOrder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"IncreaseBorrowedAmount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"remainingFunds\",\"type\":\"uint256\"}],\"name\":\"LiquidateCreditAccount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxLeverage\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeInterest\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeLiquidation\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidationDiscount\",\"type\":\"uint256\"}],\"name\":\"NewParameters\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"referralCode\",\"type\":\"uint256\"}],\"name\":\"OpenCreditAccount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"RepayCreditAccount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"TransferAccount\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"addCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"targetContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isLiquidated\",\"type\":\"bool\"}],\"name\":\"calcRepayAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.Exchange[]\",\"name\":\"paths\",\"type\":\"tuple[]\"}],\"name\":\"closeCreditAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"}],\"name\":\"creditAccounts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creditFilter\",\"outputs\":[{\"internalType\":\"contractICreditFilter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultSwapContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeOrder\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeInterest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeLiquidation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"}],\"name\":\"getCreditAccountOrRevert\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"}],\"name\":\"hasOpenedCreditAccount\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"increaseBorrowedAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"force\",\"type\":\"bool\"}],\"name\":\"liquidateCreditAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidationDiscount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxLeverageFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minHealthFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"leverageFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"referralCode\",\"type\":\"uint256\"}],\"name\":\"openCreditAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolService\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"provideCreditAccountAllowance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"repayCreditAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"repayCreditAccountETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferAccountOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlyingToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ICreditManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ICreditManagerMetaData.ABI instead.
var ICreditManagerABI = ICreditManagerMetaData.ABI

// ICreditManager is an auto generated Go binding around an Ethereum contract.
type ICreditManager struct {
	ICreditManagerCaller     // Read-only binding to the contract
	ICreditManagerTransactor // Write-only binding to the contract
	ICreditManagerFilterer   // Log filterer for contract events
}

// ICreditManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ICreditManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICreditManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ICreditManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICreditManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ICreditManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICreditManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ICreditManagerSession struct {
	Contract     *ICreditManager   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ICreditManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ICreditManagerCallerSession struct {
	Contract *ICreditManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ICreditManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ICreditManagerTransactorSession struct {
	Contract     *ICreditManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ICreditManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ICreditManagerRaw struct {
	Contract *ICreditManager // Generic contract binding to access the raw methods on
}

// ICreditManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ICreditManagerCallerRaw struct {
	Contract *ICreditManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ICreditManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ICreditManagerTransactorRaw struct {
	Contract *ICreditManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewICreditManager creates a new instance of ICreditManager, bound to a specific deployed contract.
func NewICreditManager(address common.Address, backend bind.ContractBackend) (*ICreditManager, error) {
	contract, err := bindICreditManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ICreditManager{ICreditManagerCaller: ICreditManagerCaller{contract: contract}, ICreditManagerTransactor: ICreditManagerTransactor{contract: contract}, ICreditManagerFilterer: ICreditManagerFilterer{contract: contract}}, nil
}

// NewICreditManagerCaller creates a new read-only instance of ICreditManager, bound to a specific deployed contract.
func NewICreditManagerCaller(address common.Address, caller bind.ContractCaller) (*ICreditManagerCaller, error) {
	contract, err := bindICreditManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ICreditManagerCaller{contract: contract}, nil
}

// NewICreditManagerTransactor creates a new write-only instance of ICreditManager, bound to a specific deployed contract.
func NewICreditManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ICreditManagerTransactor, error) {
	contract, err := bindICreditManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ICreditManagerTransactor{contract: contract}, nil
}

// NewICreditManagerFilterer creates a new log filterer instance of ICreditManager, bound to a specific deployed contract.
func NewICreditManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ICreditManagerFilterer, error) {
	contract, err := bindICreditManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ICreditManagerFilterer{contract: contract}, nil
}

// bindICreditManager binds a generic wrapper to an already deployed contract.
func bindICreditManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ICreditManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICreditManager *ICreditManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICreditManager.Contract.ICreditManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICreditManager *ICreditManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICreditManager.Contract.ICreditManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICreditManager *ICreditManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICreditManager.Contract.ICreditManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICreditManager *ICreditManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICreditManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICreditManager *ICreditManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICreditManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICreditManager *ICreditManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICreditManager.Contract.contract.Transact(opts, method, params...)
}

// CalcRepayAmount is a free data retrieval call binding the contract method 0x3ce07355.
//
// Solidity: function calcRepayAmount(address borrower, bool isLiquidated) view returns(uint256)
func (_ICreditManager *ICreditManagerCaller) CalcRepayAmount(opts *bind.CallOpts, borrower common.Address, isLiquidated bool) (*big.Int, error) {
	var out []interface{}
	err := _ICreditManager.contract.Call(opts, &out, "calcRepayAmount", borrower, isLiquidated)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcRepayAmount is a free data retrieval call binding the contract method 0x3ce07355.
//
// Solidity: function calcRepayAmount(address borrower, bool isLiquidated) view returns(uint256)
func (_ICreditManager *ICreditManagerSession) CalcRepayAmount(borrower common.Address, isLiquidated bool) (*big.Int, error) {
	return _ICreditManager.Contract.CalcRepayAmount(&_ICreditManager.CallOpts, borrower, isLiquidated)
}

// CalcRepayAmount is a free data retrieval call binding the contract method 0x3ce07355.
//
// Solidity: function calcRepayAmount(address borrower, bool isLiquidated) view returns(uint256)
func (_ICreditManager *ICreditManagerCallerSession) CalcRepayAmount(borrower common.Address, isLiquidated bool) (*big.Int, error) {
	return _ICreditManager.Contract.CalcRepayAmount(&_ICreditManager.CallOpts, borrower, isLiquidated)
}

// CreditAccounts is a free data retrieval call binding the contract method 0x055ee9b5.
//
// Solidity: function creditAccounts(address borrower) view returns(address)
func (_ICreditManager *ICreditManagerCaller) CreditAccounts(opts *bind.CallOpts, borrower common.Address) (common.Address, error) {
	var out []interface{}
	err := _ICreditManager.contract.Call(opts, &out, "creditAccounts", borrower)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditAccounts is a free data retrieval call binding the contract method 0x055ee9b5.
//
// Solidity: function creditAccounts(address borrower) view returns(address)
func (_ICreditManager *ICreditManagerSession) CreditAccounts(borrower common.Address) (common.Address, error) {
	return _ICreditManager.Contract.CreditAccounts(&_ICreditManager.CallOpts, borrower)
}

// CreditAccounts is a free data retrieval call binding the contract method 0x055ee9b5.
//
// Solidity: function creditAccounts(address borrower) view returns(address)
func (_ICreditManager *ICreditManagerCallerSession) CreditAccounts(borrower common.Address) (common.Address, error) {
	return _ICreditManager.Contract.CreditAccounts(&_ICreditManager.CallOpts, borrower)
}

// CreditFilter is a free data retrieval call binding the contract method 0xf93f515b.
//
// Solidity: function creditFilter() view returns(address)
func (_ICreditManager *ICreditManagerCaller) CreditFilter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ICreditManager.contract.Call(opts, &out, "creditFilter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditFilter is a free data retrieval call binding the contract method 0xf93f515b.
//
// Solidity: function creditFilter() view returns(address)
func (_ICreditManager *ICreditManagerSession) CreditFilter() (common.Address, error) {
	return _ICreditManager.Contract.CreditFilter(&_ICreditManager.CallOpts)
}

// CreditFilter is a free data retrieval call binding the contract method 0xf93f515b.
//
// Solidity: function creditFilter() view returns(address)
func (_ICreditManager *ICreditManagerCallerSession) CreditFilter() (common.Address, error) {
	return _ICreditManager.Contract.CreditFilter(&_ICreditManager.CallOpts)
}

// DefaultSwapContract is a free data retrieval call binding the contract method 0xe0c011b7.
//
// Solidity: function defaultSwapContract() view returns(address)
func (_ICreditManager *ICreditManagerCaller) DefaultSwapContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ICreditManager.contract.Call(opts, &out, "defaultSwapContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DefaultSwapContract is a free data retrieval call binding the contract method 0xe0c011b7.
//
// Solidity: function defaultSwapContract() view returns(address)
func (_ICreditManager *ICreditManagerSession) DefaultSwapContract() (common.Address, error) {
	return _ICreditManager.Contract.DefaultSwapContract(&_ICreditManager.CallOpts)
}

// DefaultSwapContract is a free data retrieval call binding the contract method 0xe0c011b7.
//
// Solidity: function defaultSwapContract() view returns(address)
func (_ICreditManager *ICreditManagerCallerSession) DefaultSwapContract() (common.Address, error) {
	return _ICreditManager.Contract.DefaultSwapContract(&_ICreditManager.CallOpts)
}

// FeeInterest is a free data retrieval call binding the contract method 0x5e0b63d3.
//
// Solidity: function feeInterest() view returns(uint256)
func (_ICreditManager *ICreditManagerCaller) FeeInterest(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ICreditManager.contract.Call(opts, &out, "feeInterest")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeInterest is a free data retrieval call binding the contract method 0x5e0b63d3.
//
// Solidity: function feeInterest() view returns(uint256)
func (_ICreditManager *ICreditManagerSession) FeeInterest() (*big.Int, error) {
	return _ICreditManager.Contract.FeeInterest(&_ICreditManager.CallOpts)
}

// FeeInterest is a free data retrieval call binding the contract method 0x5e0b63d3.
//
// Solidity: function feeInterest() view returns(uint256)
func (_ICreditManager *ICreditManagerCallerSession) FeeInterest() (*big.Int, error) {
	return _ICreditManager.Contract.FeeInterest(&_ICreditManager.CallOpts)
}

// FeeLiquidation is a free data retrieval call binding the contract method 0x3915ffaa.
//
// Solidity: function feeLiquidation() view returns(uint256)
func (_ICreditManager *ICreditManagerCaller) FeeLiquidation(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ICreditManager.contract.Call(opts, &out, "feeLiquidation")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeLiquidation is a free data retrieval call binding the contract method 0x3915ffaa.
//
// Solidity: function feeLiquidation() view returns(uint256)
func (_ICreditManager *ICreditManagerSession) FeeLiquidation() (*big.Int, error) {
	return _ICreditManager.Contract.FeeLiquidation(&_ICreditManager.CallOpts)
}

// FeeLiquidation is a free data retrieval call binding the contract method 0x3915ffaa.
//
// Solidity: function feeLiquidation() view returns(uint256)
func (_ICreditManager *ICreditManagerCallerSession) FeeLiquidation() (*big.Int, error) {
	return _ICreditManager.Contract.FeeLiquidation(&_ICreditManager.CallOpts)
}

// GetCreditAccountOrRevert is a free data retrieval call binding the contract method 0xe958b704.
//
// Solidity: function getCreditAccountOrRevert(address borrower) view returns(address)
func (_ICreditManager *ICreditManagerCaller) GetCreditAccountOrRevert(opts *bind.CallOpts, borrower common.Address) (common.Address, error) {
	var out []interface{}
	err := _ICreditManager.contract.Call(opts, &out, "getCreditAccountOrRevert", borrower)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCreditAccountOrRevert is a free data retrieval call binding the contract method 0xe958b704.
//
// Solidity: function getCreditAccountOrRevert(address borrower) view returns(address)
func (_ICreditManager *ICreditManagerSession) GetCreditAccountOrRevert(borrower common.Address) (common.Address, error) {
	return _ICreditManager.Contract.GetCreditAccountOrRevert(&_ICreditManager.CallOpts, borrower)
}

// GetCreditAccountOrRevert is a free data retrieval call binding the contract method 0xe958b704.
//
// Solidity: function getCreditAccountOrRevert(address borrower) view returns(address)
func (_ICreditManager *ICreditManagerCallerSession) GetCreditAccountOrRevert(borrower common.Address) (common.Address, error) {
	return _ICreditManager.Contract.GetCreditAccountOrRevert(&_ICreditManager.CallOpts, borrower)
}

// HasOpenedCreditAccount is a free data retrieval call binding the contract method 0x256ac915.
//
// Solidity: function hasOpenedCreditAccount(address borrower) view returns(bool)
func (_ICreditManager *ICreditManagerCaller) HasOpenedCreditAccount(opts *bind.CallOpts, borrower common.Address) (bool, error) {
	var out []interface{}
	err := _ICreditManager.contract.Call(opts, &out, "hasOpenedCreditAccount", borrower)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasOpenedCreditAccount is a free data retrieval call binding the contract method 0x256ac915.
//
// Solidity: function hasOpenedCreditAccount(address borrower) view returns(bool)
func (_ICreditManager *ICreditManagerSession) HasOpenedCreditAccount(borrower common.Address) (bool, error) {
	return _ICreditManager.Contract.HasOpenedCreditAccount(&_ICreditManager.CallOpts, borrower)
}

// HasOpenedCreditAccount is a free data retrieval call binding the contract method 0x256ac915.
//
// Solidity: function hasOpenedCreditAccount(address borrower) view returns(bool)
func (_ICreditManager *ICreditManagerCallerSession) HasOpenedCreditAccount(borrower common.Address) (bool, error) {
	return _ICreditManager.Contract.HasOpenedCreditAccount(&_ICreditManager.CallOpts, borrower)
}

// LiquidationDiscount is a free data retrieval call binding the contract method 0x8053fcbe.
//
// Solidity: function liquidationDiscount() view returns(uint256)
func (_ICreditManager *ICreditManagerCaller) LiquidationDiscount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ICreditManager.contract.Call(opts, &out, "liquidationDiscount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LiquidationDiscount is a free data retrieval call binding the contract method 0x8053fcbe.
//
// Solidity: function liquidationDiscount() view returns(uint256)
func (_ICreditManager *ICreditManagerSession) LiquidationDiscount() (*big.Int, error) {
	return _ICreditManager.Contract.LiquidationDiscount(&_ICreditManager.CallOpts)
}

// LiquidationDiscount is a free data retrieval call binding the contract method 0x8053fcbe.
//
// Solidity: function liquidationDiscount() view returns(uint256)
func (_ICreditManager *ICreditManagerCallerSession) LiquidationDiscount() (*big.Int, error) {
	return _ICreditManager.Contract.LiquidationDiscount(&_ICreditManager.CallOpts)
}

// MaxAmount is a free data retrieval call binding the contract method 0x5f48f393.
//
// Solidity: function maxAmount() view returns(uint256)
func (_ICreditManager *ICreditManagerCaller) MaxAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ICreditManager.contract.Call(opts, &out, "maxAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxAmount is a free data retrieval call binding the contract method 0x5f48f393.
//
// Solidity: function maxAmount() view returns(uint256)
func (_ICreditManager *ICreditManagerSession) MaxAmount() (*big.Int, error) {
	return _ICreditManager.Contract.MaxAmount(&_ICreditManager.CallOpts)
}

// MaxAmount is a free data retrieval call binding the contract method 0x5f48f393.
//
// Solidity: function maxAmount() view returns(uint256)
func (_ICreditManager *ICreditManagerCallerSession) MaxAmount() (*big.Int, error) {
	return _ICreditManager.Contract.MaxAmount(&_ICreditManager.CallOpts)
}

// MaxLeverageFactor is a free data retrieval call binding the contract method 0xb2c53a6c.
//
// Solidity: function maxLeverageFactor() view returns(uint256)
func (_ICreditManager *ICreditManagerCaller) MaxLeverageFactor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ICreditManager.contract.Call(opts, &out, "maxLeverageFactor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxLeverageFactor is a free data retrieval call binding the contract method 0xb2c53a6c.
//
// Solidity: function maxLeverageFactor() view returns(uint256)
func (_ICreditManager *ICreditManagerSession) MaxLeverageFactor() (*big.Int, error) {
	return _ICreditManager.Contract.MaxLeverageFactor(&_ICreditManager.CallOpts)
}

// MaxLeverageFactor is a free data retrieval call binding the contract method 0xb2c53a6c.
//
// Solidity: function maxLeverageFactor() view returns(uint256)
func (_ICreditManager *ICreditManagerCallerSession) MaxLeverageFactor() (*big.Int, error) {
	return _ICreditManager.Contract.MaxLeverageFactor(&_ICreditManager.CallOpts)
}

// MinAmount is a free data retrieval call binding the contract method 0x9b2cb5d8.
//
// Solidity: function minAmount() view returns(uint256)
func (_ICreditManager *ICreditManagerCaller) MinAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ICreditManager.contract.Call(opts, &out, "minAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinAmount is a free data retrieval call binding the contract method 0x9b2cb5d8.
//
// Solidity: function minAmount() view returns(uint256)
func (_ICreditManager *ICreditManagerSession) MinAmount() (*big.Int, error) {
	return _ICreditManager.Contract.MinAmount(&_ICreditManager.CallOpts)
}

// MinAmount is a free data retrieval call binding the contract method 0x9b2cb5d8.
//
// Solidity: function minAmount() view returns(uint256)
func (_ICreditManager *ICreditManagerCallerSession) MinAmount() (*big.Int, error) {
	return _ICreditManager.Contract.MinAmount(&_ICreditManager.CallOpts)
}

// MinHealthFactor is a free data retrieval call binding the contract method 0xe1b4264c.
//
// Solidity: function minHealthFactor() view returns(uint256)
func (_ICreditManager *ICreditManagerCaller) MinHealthFactor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ICreditManager.contract.Call(opts, &out, "minHealthFactor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinHealthFactor is a free data retrieval call binding the contract method 0xe1b4264c.
//
// Solidity: function minHealthFactor() view returns(uint256)
func (_ICreditManager *ICreditManagerSession) MinHealthFactor() (*big.Int, error) {
	return _ICreditManager.Contract.MinHealthFactor(&_ICreditManager.CallOpts)
}

// MinHealthFactor is a free data retrieval call binding the contract method 0xe1b4264c.
//
// Solidity: function minHealthFactor() view returns(uint256)
func (_ICreditManager *ICreditManagerCallerSession) MinHealthFactor() (*big.Int, error) {
	return _ICreditManager.Contract.MinHealthFactor(&_ICreditManager.CallOpts)
}

// PoolService is a free data retrieval call binding the contract method 0x570a7af2.
//
// Solidity: function poolService() view returns(address)
func (_ICreditManager *ICreditManagerCaller) PoolService(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ICreditManager.contract.Call(opts, &out, "poolService")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolService is a free data retrieval call binding the contract method 0x570a7af2.
//
// Solidity: function poolService() view returns(address)
func (_ICreditManager *ICreditManagerSession) PoolService() (common.Address, error) {
	return _ICreditManager.Contract.PoolService(&_ICreditManager.CallOpts)
}

// PoolService is a free data retrieval call binding the contract method 0x570a7af2.
//
// Solidity: function poolService() view returns(address)
func (_ICreditManager *ICreditManagerCallerSession) PoolService() (common.Address, error) {
	return _ICreditManager.Contract.PoolService(&_ICreditManager.CallOpts)
}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_ICreditManager *ICreditManagerCaller) UnderlyingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ICreditManager.contract.Call(opts, &out, "underlyingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_ICreditManager *ICreditManagerSession) UnderlyingToken() (common.Address, error) {
	return _ICreditManager.Contract.UnderlyingToken(&_ICreditManager.CallOpts)
}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_ICreditManager *ICreditManagerCallerSession) UnderlyingToken() (common.Address, error) {
	return _ICreditManager.Contract.UnderlyingToken(&_ICreditManager.CallOpts)
}

// AddCollateral is a paid mutator transaction binding the contract method 0x59781034.
//
// Solidity: function addCollateral(address onBehalfOf, address token, uint256 amount) returns()
func (_ICreditManager *ICreditManagerTransactor) AddCollateral(opts *bind.TransactOpts, onBehalfOf common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ICreditManager.contract.Transact(opts, "addCollateral", onBehalfOf, token, amount)
}

// AddCollateral is a paid mutator transaction binding the contract method 0x59781034.
//
// Solidity: function addCollateral(address onBehalfOf, address token, uint256 amount) returns()
func (_ICreditManager *ICreditManagerSession) AddCollateral(onBehalfOf common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ICreditManager.Contract.AddCollateral(&_ICreditManager.TransactOpts, onBehalfOf, token, amount)
}

// AddCollateral is a paid mutator transaction binding the contract method 0x59781034.
//
// Solidity: function addCollateral(address onBehalfOf, address token, uint256 amount) returns()
func (_ICreditManager *ICreditManagerTransactorSession) AddCollateral(onBehalfOf common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ICreditManager.Contract.AddCollateral(&_ICreditManager.TransactOpts, onBehalfOf, token, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x7e5465ba.
//
// Solidity: function approve(address targetContract, address token) returns()
func (_ICreditManager *ICreditManagerTransactor) Approve(opts *bind.TransactOpts, targetContract common.Address, token common.Address) (*types.Transaction, error) {
	return _ICreditManager.contract.Transact(opts, "approve", targetContract, token)
}

// Approve is a paid mutator transaction binding the contract method 0x7e5465ba.
//
// Solidity: function approve(address targetContract, address token) returns()
func (_ICreditManager *ICreditManagerSession) Approve(targetContract common.Address, token common.Address) (*types.Transaction, error) {
	return _ICreditManager.Contract.Approve(&_ICreditManager.TransactOpts, targetContract, token)
}

// Approve is a paid mutator transaction binding the contract method 0x7e5465ba.
//
// Solidity: function approve(address targetContract, address token) returns()
func (_ICreditManager *ICreditManagerTransactorSession) Approve(targetContract common.Address, token common.Address) (*types.Transaction, error) {
	return _ICreditManager.Contract.Approve(&_ICreditManager.TransactOpts, targetContract, token)
}

// CloseCreditAccount is a paid mutator transaction binding the contract method 0xab114805.
//
// Solidity: function closeCreditAccount(address to, (address[],uint256)[] paths) returns()
func (_ICreditManager *ICreditManagerTransactor) CloseCreditAccount(opts *bind.TransactOpts, to common.Address, paths []DataTypesExchange) (*types.Transaction, error) {
	return _ICreditManager.contract.Transact(opts, "closeCreditAccount", to, paths)
}

// CloseCreditAccount is a paid mutator transaction binding the contract method 0xab114805.
//
// Solidity: function closeCreditAccount(address to, (address[],uint256)[] paths) returns()
func (_ICreditManager *ICreditManagerSession) CloseCreditAccount(to common.Address, paths []DataTypesExchange) (*types.Transaction, error) {
	return _ICreditManager.Contract.CloseCreditAccount(&_ICreditManager.TransactOpts, to, paths)
}

// CloseCreditAccount is a paid mutator transaction binding the contract method 0xab114805.
//
// Solidity: function closeCreditAccount(address to, (address[],uint256)[] paths) returns()
func (_ICreditManager *ICreditManagerTransactorSession) CloseCreditAccount(to common.Address, paths []DataTypesExchange) (*types.Transaction, error) {
	return _ICreditManager.Contract.CloseCreditAccount(&_ICreditManager.TransactOpts, to, paths)
}

// ExecuteOrder is a paid mutator transaction binding the contract method 0x6ce4074a.
//
// Solidity: function executeOrder(address borrower, address target, bytes data) returns(bytes)
func (_ICreditManager *ICreditManagerTransactor) ExecuteOrder(opts *bind.TransactOpts, borrower common.Address, target common.Address, data []byte) (*types.Transaction, error) {
	return _ICreditManager.contract.Transact(opts, "executeOrder", borrower, target, data)
}

// ExecuteOrder is a paid mutator transaction binding the contract method 0x6ce4074a.
//
// Solidity: function executeOrder(address borrower, address target, bytes data) returns(bytes)
func (_ICreditManager *ICreditManagerSession) ExecuteOrder(borrower common.Address, target common.Address, data []byte) (*types.Transaction, error) {
	return _ICreditManager.Contract.ExecuteOrder(&_ICreditManager.TransactOpts, borrower, target, data)
}

// ExecuteOrder is a paid mutator transaction binding the contract method 0x6ce4074a.
//
// Solidity: function executeOrder(address borrower, address target, bytes data) returns(bytes)
func (_ICreditManager *ICreditManagerTransactorSession) ExecuteOrder(borrower common.Address, target common.Address, data []byte) (*types.Transaction, error) {
	return _ICreditManager.Contract.ExecuteOrder(&_ICreditManager.TransactOpts, borrower, target, data)
}

// IncreaseBorrowedAmount is a paid mutator transaction binding the contract method 0x9efc60d0.
//
// Solidity: function increaseBorrowedAmount(uint256 amount) returns()
func (_ICreditManager *ICreditManagerTransactor) IncreaseBorrowedAmount(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _ICreditManager.contract.Transact(opts, "increaseBorrowedAmount", amount)
}

// IncreaseBorrowedAmount is a paid mutator transaction binding the contract method 0x9efc60d0.
//
// Solidity: function increaseBorrowedAmount(uint256 amount) returns()
func (_ICreditManager *ICreditManagerSession) IncreaseBorrowedAmount(amount *big.Int) (*types.Transaction, error) {
	return _ICreditManager.Contract.IncreaseBorrowedAmount(&_ICreditManager.TransactOpts, amount)
}

// IncreaseBorrowedAmount is a paid mutator transaction binding the contract method 0x9efc60d0.
//
// Solidity: function increaseBorrowedAmount(uint256 amount) returns()
func (_ICreditManager *ICreditManagerTransactorSession) IncreaseBorrowedAmount(amount *big.Int) (*types.Transaction, error) {
	return _ICreditManager.Contract.IncreaseBorrowedAmount(&_ICreditManager.TransactOpts, amount)
}

// LiquidateCreditAccount is a paid mutator transaction binding the contract method 0xa69a7dd6.
//
// Solidity: function liquidateCreditAccount(address borrower, address to, bool force) returns()
func (_ICreditManager *ICreditManagerTransactor) LiquidateCreditAccount(opts *bind.TransactOpts, borrower common.Address, to common.Address, force bool) (*types.Transaction, error) {
	return _ICreditManager.contract.Transact(opts, "liquidateCreditAccount", borrower, to, force)
}

// LiquidateCreditAccount is a paid mutator transaction binding the contract method 0xa69a7dd6.
//
// Solidity: function liquidateCreditAccount(address borrower, address to, bool force) returns()
func (_ICreditManager *ICreditManagerSession) LiquidateCreditAccount(borrower common.Address, to common.Address, force bool) (*types.Transaction, error) {
	return _ICreditManager.Contract.LiquidateCreditAccount(&_ICreditManager.TransactOpts, borrower, to, force)
}

// LiquidateCreditAccount is a paid mutator transaction binding the contract method 0xa69a7dd6.
//
// Solidity: function liquidateCreditAccount(address borrower, address to, bool force) returns()
func (_ICreditManager *ICreditManagerTransactorSession) LiquidateCreditAccount(borrower common.Address, to common.Address, force bool) (*types.Transaction, error) {
	return _ICreditManager.Contract.LiquidateCreditAccount(&_ICreditManager.TransactOpts, borrower, to, force)
}

// OpenCreditAccount is a paid mutator transaction binding the contract method 0x5288ba4b.
//
// Solidity: function openCreditAccount(uint256 amount, address onBehalfOf, uint256 leverageFactor, uint256 referralCode) returns()
func (_ICreditManager *ICreditManagerTransactor) OpenCreditAccount(opts *bind.TransactOpts, amount *big.Int, onBehalfOf common.Address, leverageFactor *big.Int, referralCode *big.Int) (*types.Transaction, error) {
	return _ICreditManager.contract.Transact(opts, "openCreditAccount", amount, onBehalfOf, leverageFactor, referralCode)
}

// OpenCreditAccount is a paid mutator transaction binding the contract method 0x5288ba4b.
//
// Solidity: function openCreditAccount(uint256 amount, address onBehalfOf, uint256 leverageFactor, uint256 referralCode) returns()
func (_ICreditManager *ICreditManagerSession) OpenCreditAccount(amount *big.Int, onBehalfOf common.Address, leverageFactor *big.Int, referralCode *big.Int) (*types.Transaction, error) {
	return _ICreditManager.Contract.OpenCreditAccount(&_ICreditManager.TransactOpts, amount, onBehalfOf, leverageFactor, referralCode)
}

// OpenCreditAccount is a paid mutator transaction binding the contract method 0x5288ba4b.
//
// Solidity: function openCreditAccount(uint256 amount, address onBehalfOf, uint256 leverageFactor, uint256 referralCode) returns()
func (_ICreditManager *ICreditManagerTransactorSession) OpenCreditAccount(amount *big.Int, onBehalfOf common.Address, leverageFactor *big.Int, referralCode *big.Int) (*types.Transaction, error) {
	return _ICreditManager.Contract.OpenCreditAccount(&_ICreditManager.TransactOpts, amount, onBehalfOf, leverageFactor, referralCode)
}

// ProvideCreditAccountAllowance is a paid mutator transaction binding the contract method 0x579122ab.
//
// Solidity: function provideCreditAccountAllowance(address creditAccount, address toContract, address token) returns()
func (_ICreditManager *ICreditManagerTransactor) ProvideCreditAccountAllowance(opts *bind.TransactOpts, creditAccount common.Address, toContract common.Address, token common.Address) (*types.Transaction, error) {
	return _ICreditManager.contract.Transact(opts, "provideCreditAccountAllowance", creditAccount, toContract, token)
}

// ProvideCreditAccountAllowance is a paid mutator transaction binding the contract method 0x579122ab.
//
// Solidity: function provideCreditAccountAllowance(address creditAccount, address toContract, address token) returns()
func (_ICreditManager *ICreditManagerSession) ProvideCreditAccountAllowance(creditAccount common.Address, toContract common.Address, token common.Address) (*types.Transaction, error) {
	return _ICreditManager.Contract.ProvideCreditAccountAllowance(&_ICreditManager.TransactOpts, creditAccount, toContract, token)
}

// ProvideCreditAccountAllowance is a paid mutator transaction binding the contract method 0x579122ab.
//
// Solidity: function provideCreditAccountAllowance(address creditAccount, address toContract, address token) returns()
func (_ICreditManager *ICreditManagerTransactorSession) ProvideCreditAccountAllowance(creditAccount common.Address, toContract common.Address, token common.Address) (*types.Transaction, error) {
	return _ICreditManager.Contract.ProvideCreditAccountAllowance(&_ICreditManager.TransactOpts, creditAccount, toContract, token)
}

// RepayCreditAccount is a paid mutator transaction binding the contract method 0xd692ba33.
//
// Solidity: function repayCreditAccount(address to) returns()
func (_ICreditManager *ICreditManagerTransactor) RepayCreditAccount(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _ICreditManager.contract.Transact(opts, "repayCreditAccount", to)
}

// RepayCreditAccount is a paid mutator transaction binding the contract method 0xd692ba33.
//
// Solidity: function repayCreditAccount(address to) returns()
func (_ICreditManager *ICreditManagerSession) RepayCreditAccount(to common.Address) (*types.Transaction, error) {
	return _ICreditManager.Contract.RepayCreditAccount(&_ICreditManager.TransactOpts, to)
}

// RepayCreditAccount is a paid mutator transaction binding the contract method 0xd692ba33.
//
// Solidity: function repayCreditAccount(address to) returns()
func (_ICreditManager *ICreditManagerTransactorSession) RepayCreditAccount(to common.Address) (*types.Transaction, error) {
	return _ICreditManager.Contract.RepayCreditAccount(&_ICreditManager.TransactOpts, to)
}

// RepayCreditAccountETH is a paid mutator transaction binding the contract method 0xa6eab5c2.
//
// Solidity: function repayCreditAccountETH(address borrower, address to) returns(uint256)
func (_ICreditManager *ICreditManagerTransactor) RepayCreditAccountETH(opts *bind.TransactOpts, borrower common.Address, to common.Address) (*types.Transaction, error) {
	return _ICreditManager.contract.Transact(opts, "repayCreditAccountETH", borrower, to)
}

// RepayCreditAccountETH is a paid mutator transaction binding the contract method 0xa6eab5c2.
//
// Solidity: function repayCreditAccountETH(address borrower, address to) returns(uint256)
func (_ICreditManager *ICreditManagerSession) RepayCreditAccountETH(borrower common.Address, to common.Address) (*types.Transaction, error) {
	return _ICreditManager.Contract.RepayCreditAccountETH(&_ICreditManager.TransactOpts, borrower, to)
}

// RepayCreditAccountETH is a paid mutator transaction binding the contract method 0xa6eab5c2.
//
// Solidity: function repayCreditAccountETH(address borrower, address to) returns(uint256)
func (_ICreditManager *ICreditManagerTransactorSession) RepayCreditAccountETH(borrower common.Address, to common.Address) (*types.Transaction, error) {
	return _ICreditManager.Contract.RepayCreditAccountETH(&_ICreditManager.TransactOpts, borrower, to)
}

// TransferAccountOwnership is a paid mutator transaction binding the contract method 0x5019e20a.
//
// Solidity: function transferAccountOwnership(address newOwner) returns()
func (_ICreditManager *ICreditManagerTransactor) TransferAccountOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ICreditManager.contract.Transact(opts, "transferAccountOwnership", newOwner)
}

// TransferAccountOwnership is a paid mutator transaction binding the contract method 0x5019e20a.
//
// Solidity: function transferAccountOwnership(address newOwner) returns()
func (_ICreditManager *ICreditManagerSession) TransferAccountOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ICreditManager.Contract.TransferAccountOwnership(&_ICreditManager.TransactOpts, newOwner)
}

// TransferAccountOwnership is a paid mutator transaction binding the contract method 0x5019e20a.
//
// Solidity: function transferAccountOwnership(address newOwner) returns()
func (_ICreditManager *ICreditManagerTransactorSession) TransferAccountOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ICreditManager.Contract.TransferAccountOwnership(&_ICreditManager.TransactOpts, newOwner)
}

// ICreditManagerAddCollateralIterator is returned from FilterAddCollateral and is used to iterate over the raw logs and unpacked data for AddCollateral events raised by the ICreditManager contract.
type ICreditManagerAddCollateralIterator struct {
	Event *ICreditManagerAddCollateral // Event containing the contract specifics and raw log

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
func (it *ICreditManagerAddCollateralIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICreditManagerAddCollateral)
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
		it.Event = new(ICreditManagerAddCollateral)
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
func (it *ICreditManagerAddCollateralIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICreditManagerAddCollateralIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICreditManagerAddCollateral represents a AddCollateral event raised by the ICreditManager contract.
type ICreditManagerAddCollateral struct {
	OnBehalfOf common.Address
	Token      common.Address
	Value      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAddCollateral is a free log retrieval operation binding the contract event 0xa32435755c235de2976ed44a75a2f85cb01faf0c894f639fe0c32bb9455fea8f.
//
// Solidity: event AddCollateral(address indexed onBehalfOf, address indexed token, uint256 value)
func (_ICreditManager *ICreditManagerFilterer) FilterAddCollateral(opts *bind.FilterOpts, onBehalfOf []common.Address, token []common.Address) (*ICreditManagerAddCollateralIterator, error) {

	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ICreditManager.contract.FilterLogs(opts, "AddCollateral", onBehalfOfRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &ICreditManagerAddCollateralIterator{contract: _ICreditManager.contract, event: "AddCollateral", logs: logs, sub: sub}, nil
}

// WatchAddCollateral is a free log subscription operation binding the contract event 0xa32435755c235de2976ed44a75a2f85cb01faf0c894f639fe0c32bb9455fea8f.
//
// Solidity: event AddCollateral(address indexed onBehalfOf, address indexed token, uint256 value)
func (_ICreditManager *ICreditManagerFilterer) WatchAddCollateral(opts *bind.WatchOpts, sink chan<- *ICreditManagerAddCollateral, onBehalfOf []common.Address, token []common.Address) (event.Subscription, error) {

	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ICreditManager.contract.WatchLogs(opts, "AddCollateral", onBehalfOfRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICreditManagerAddCollateral)
				if err := _ICreditManager.contract.UnpackLog(event, "AddCollateral", log); err != nil {
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

// ParseAddCollateral is a log parse operation binding the contract event 0xa32435755c235de2976ed44a75a2f85cb01faf0c894f639fe0c32bb9455fea8f.
//
// Solidity: event AddCollateral(address indexed onBehalfOf, address indexed token, uint256 value)
func (_ICreditManager *ICreditManagerFilterer) ParseAddCollateral(log types.Log) (*ICreditManagerAddCollateral, error) {
	event := new(ICreditManagerAddCollateral)
	if err := _ICreditManager.contract.UnpackLog(event, "AddCollateral", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ICreditManagerCloseCreditAccountIterator is returned from FilterCloseCreditAccount and is used to iterate over the raw logs and unpacked data for CloseCreditAccount events raised by the ICreditManager contract.
type ICreditManagerCloseCreditAccountIterator struct {
	Event *ICreditManagerCloseCreditAccount // Event containing the contract specifics and raw log

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
func (it *ICreditManagerCloseCreditAccountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICreditManagerCloseCreditAccount)
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
		it.Event = new(ICreditManagerCloseCreditAccount)
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
func (it *ICreditManagerCloseCreditAccountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICreditManagerCloseCreditAccountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICreditManagerCloseCreditAccount represents a CloseCreditAccount event raised by the ICreditManager contract.
type ICreditManagerCloseCreditAccount struct {
	Owner          common.Address
	To             common.Address
	RemainingFunds *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterCloseCreditAccount is a free log retrieval operation binding the contract event 0xca05b632388199c23de1352b2e96fd72a0ec71611683330b38060c004bbf0a76.
//
// Solidity: event CloseCreditAccount(address indexed owner, address indexed to, uint256 remainingFunds)
func (_ICreditManager *ICreditManagerFilterer) FilterCloseCreditAccount(opts *bind.FilterOpts, owner []common.Address, to []common.Address) (*ICreditManagerCloseCreditAccountIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ICreditManager.contract.FilterLogs(opts, "CloseCreditAccount", ownerRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ICreditManagerCloseCreditAccountIterator{contract: _ICreditManager.contract, event: "CloseCreditAccount", logs: logs, sub: sub}, nil
}

// WatchCloseCreditAccount is a free log subscription operation binding the contract event 0xca05b632388199c23de1352b2e96fd72a0ec71611683330b38060c004bbf0a76.
//
// Solidity: event CloseCreditAccount(address indexed owner, address indexed to, uint256 remainingFunds)
func (_ICreditManager *ICreditManagerFilterer) WatchCloseCreditAccount(opts *bind.WatchOpts, sink chan<- *ICreditManagerCloseCreditAccount, owner []common.Address, to []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ICreditManager.contract.WatchLogs(opts, "CloseCreditAccount", ownerRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICreditManagerCloseCreditAccount)
				if err := _ICreditManager.contract.UnpackLog(event, "CloseCreditAccount", log); err != nil {
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

// ParseCloseCreditAccount is a log parse operation binding the contract event 0xca05b632388199c23de1352b2e96fd72a0ec71611683330b38060c004bbf0a76.
//
// Solidity: event CloseCreditAccount(address indexed owner, address indexed to, uint256 remainingFunds)
func (_ICreditManager *ICreditManagerFilterer) ParseCloseCreditAccount(log types.Log) (*ICreditManagerCloseCreditAccount, error) {
	event := new(ICreditManagerCloseCreditAccount)
	if err := _ICreditManager.contract.UnpackLog(event, "CloseCreditAccount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ICreditManagerExecuteOrderIterator is returned from FilterExecuteOrder and is used to iterate over the raw logs and unpacked data for ExecuteOrder events raised by the ICreditManager contract.
type ICreditManagerExecuteOrderIterator struct {
	Event *ICreditManagerExecuteOrder // Event containing the contract specifics and raw log

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
func (it *ICreditManagerExecuteOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICreditManagerExecuteOrder)
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
		it.Event = new(ICreditManagerExecuteOrder)
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
func (it *ICreditManagerExecuteOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICreditManagerExecuteOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICreditManagerExecuteOrder represents a ExecuteOrder event raised by the ICreditManager contract.
type ICreditManagerExecuteOrder struct {
	Borrower common.Address
	Target   common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterExecuteOrder is a free log retrieval operation binding the contract event 0xaed1eb34af6acd8c1e3911fb2ebb875a66324b03957886bd002227b17f52ab03.
//
// Solidity: event ExecuteOrder(address indexed borrower, address indexed target)
func (_ICreditManager *ICreditManagerFilterer) FilterExecuteOrder(opts *bind.FilterOpts, borrower []common.Address, target []common.Address) (*ICreditManagerExecuteOrderIterator, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _ICreditManager.contract.FilterLogs(opts, "ExecuteOrder", borrowerRule, targetRule)
	if err != nil {
		return nil, err
	}
	return &ICreditManagerExecuteOrderIterator{contract: _ICreditManager.contract, event: "ExecuteOrder", logs: logs, sub: sub}, nil
}

// WatchExecuteOrder is a free log subscription operation binding the contract event 0xaed1eb34af6acd8c1e3911fb2ebb875a66324b03957886bd002227b17f52ab03.
//
// Solidity: event ExecuteOrder(address indexed borrower, address indexed target)
func (_ICreditManager *ICreditManagerFilterer) WatchExecuteOrder(opts *bind.WatchOpts, sink chan<- *ICreditManagerExecuteOrder, borrower []common.Address, target []common.Address) (event.Subscription, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _ICreditManager.contract.WatchLogs(opts, "ExecuteOrder", borrowerRule, targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICreditManagerExecuteOrder)
				if err := _ICreditManager.contract.UnpackLog(event, "ExecuteOrder", log); err != nil {
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

// ParseExecuteOrder is a log parse operation binding the contract event 0xaed1eb34af6acd8c1e3911fb2ebb875a66324b03957886bd002227b17f52ab03.
//
// Solidity: event ExecuteOrder(address indexed borrower, address indexed target)
func (_ICreditManager *ICreditManagerFilterer) ParseExecuteOrder(log types.Log) (*ICreditManagerExecuteOrder, error) {
	event := new(ICreditManagerExecuteOrder)
	if err := _ICreditManager.contract.UnpackLog(event, "ExecuteOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ICreditManagerIncreaseBorrowedAmountIterator is returned from FilterIncreaseBorrowedAmount and is used to iterate over the raw logs and unpacked data for IncreaseBorrowedAmount events raised by the ICreditManager contract.
type ICreditManagerIncreaseBorrowedAmountIterator struct {
	Event *ICreditManagerIncreaseBorrowedAmount // Event containing the contract specifics and raw log

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
func (it *ICreditManagerIncreaseBorrowedAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICreditManagerIncreaseBorrowedAmount)
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
		it.Event = new(ICreditManagerIncreaseBorrowedAmount)
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
func (it *ICreditManagerIncreaseBorrowedAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICreditManagerIncreaseBorrowedAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICreditManagerIncreaseBorrowedAmount represents a IncreaseBorrowedAmount event raised by the ICreditManager contract.
type ICreditManagerIncreaseBorrowedAmount struct {
	Borrower common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterIncreaseBorrowedAmount is a free log retrieval operation binding the contract event 0x9cac51154cc0d835e2f9c9d1f59a9344588cee107f4203bf58a8c797e3a58c45.
//
// Solidity: event IncreaseBorrowedAmount(address indexed borrower, uint256 amount)
func (_ICreditManager *ICreditManagerFilterer) FilterIncreaseBorrowedAmount(opts *bind.FilterOpts, borrower []common.Address) (*ICreditManagerIncreaseBorrowedAmountIterator, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _ICreditManager.contract.FilterLogs(opts, "IncreaseBorrowedAmount", borrowerRule)
	if err != nil {
		return nil, err
	}
	return &ICreditManagerIncreaseBorrowedAmountIterator{contract: _ICreditManager.contract, event: "IncreaseBorrowedAmount", logs: logs, sub: sub}, nil
}

// WatchIncreaseBorrowedAmount is a free log subscription operation binding the contract event 0x9cac51154cc0d835e2f9c9d1f59a9344588cee107f4203bf58a8c797e3a58c45.
//
// Solidity: event IncreaseBorrowedAmount(address indexed borrower, uint256 amount)
func (_ICreditManager *ICreditManagerFilterer) WatchIncreaseBorrowedAmount(opts *bind.WatchOpts, sink chan<- *ICreditManagerIncreaseBorrowedAmount, borrower []common.Address) (event.Subscription, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _ICreditManager.contract.WatchLogs(opts, "IncreaseBorrowedAmount", borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICreditManagerIncreaseBorrowedAmount)
				if err := _ICreditManager.contract.UnpackLog(event, "IncreaseBorrowedAmount", log); err != nil {
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

// ParseIncreaseBorrowedAmount is a log parse operation binding the contract event 0x9cac51154cc0d835e2f9c9d1f59a9344588cee107f4203bf58a8c797e3a58c45.
//
// Solidity: event IncreaseBorrowedAmount(address indexed borrower, uint256 amount)
func (_ICreditManager *ICreditManagerFilterer) ParseIncreaseBorrowedAmount(log types.Log) (*ICreditManagerIncreaseBorrowedAmount, error) {
	event := new(ICreditManagerIncreaseBorrowedAmount)
	if err := _ICreditManager.contract.UnpackLog(event, "IncreaseBorrowedAmount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ICreditManagerLiquidateCreditAccountIterator is returned from FilterLiquidateCreditAccount and is used to iterate over the raw logs and unpacked data for LiquidateCreditAccount events raised by the ICreditManager contract.
type ICreditManagerLiquidateCreditAccountIterator struct {
	Event *ICreditManagerLiquidateCreditAccount // Event containing the contract specifics and raw log

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
func (it *ICreditManagerLiquidateCreditAccountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICreditManagerLiquidateCreditAccount)
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
		it.Event = new(ICreditManagerLiquidateCreditAccount)
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
func (it *ICreditManagerLiquidateCreditAccountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICreditManagerLiquidateCreditAccountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICreditManagerLiquidateCreditAccount represents a LiquidateCreditAccount event raised by the ICreditManager contract.
type ICreditManagerLiquidateCreditAccount struct {
	Owner          common.Address
	Liquidator     common.Address
	RemainingFunds *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLiquidateCreditAccount is a free log retrieval operation binding the contract event 0x5e5da6c348e62989f9cfe029252433fc99009b7d28fa3c20d675520a10ff5896.
//
// Solidity: event LiquidateCreditAccount(address indexed owner, address indexed liquidator, uint256 remainingFunds)
func (_ICreditManager *ICreditManagerFilterer) FilterLiquidateCreditAccount(opts *bind.FilterOpts, owner []common.Address, liquidator []common.Address) (*ICreditManagerLiquidateCreditAccountIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var liquidatorRule []interface{}
	for _, liquidatorItem := range liquidator {
		liquidatorRule = append(liquidatorRule, liquidatorItem)
	}

	logs, sub, err := _ICreditManager.contract.FilterLogs(opts, "LiquidateCreditAccount", ownerRule, liquidatorRule)
	if err != nil {
		return nil, err
	}
	return &ICreditManagerLiquidateCreditAccountIterator{contract: _ICreditManager.contract, event: "LiquidateCreditAccount", logs: logs, sub: sub}, nil
}

// WatchLiquidateCreditAccount is a free log subscription operation binding the contract event 0x5e5da6c348e62989f9cfe029252433fc99009b7d28fa3c20d675520a10ff5896.
//
// Solidity: event LiquidateCreditAccount(address indexed owner, address indexed liquidator, uint256 remainingFunds)
func (_ICreditManager *ICreditManagerFilterer) WatchLiquidateCreditAccount(opts *bind.WatchOpts, sink chan<- *ICreditManagerLiquidateCreditAccount, owner []common.Address, liquidator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var liquidatorRule []interface{}
	for _, liquidatorItem := range liquidator {
		liquidatorRule = append(liquidatorRule, liquidatorItem)
	}

	logs, sub, err := _ICreditManager.contract.WatchLogs(opts, "LiquidateCreditAccount", ownerRule, liquidatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICreditManagerLiquidateCreditAccount)
				if err := _ICreditManager.contract.UnpackLog(event, "LiquidateCreditAccount", log); err != nil {
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

// ParseLiquidateCreditAccount is a log parse operation binding the contract event 0x5e5da6c348e62989f9cfe029252433fc99009b7d28fa3c20d675520a10ff5896.
//
// Solidity: event LiquidateCreditAccount(address indexed owner, address indexed liquidator, uint256 remainingFunds)
func (_ICreditManager *ICreditManagerFilterer) ParseLiquidateCreditAccount(log types.Log) (*ICreditManagerLiquidateCreditAccount, error) {
	event := new(ICreditManagerLiquidateCreditAccount)
	if err := _ICreditManager.contract.UnpackLog(event, "LiquidateCreditAccount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ICreditManagerNewParametersIterator is returned from FilterNewParameters and is used to iterate over the raw logs and unpacked data for NewParameters events raised by the ICreditManager contract.
type ICreditManagerNewParametersIterator struct {
	Event *ICreditManagerNewParameters // Event containing the contract specifics and raw log

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
func (it *ICreditManagerNewParametersIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICreditManagerNewParameters)
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
		it.Event = new(ICreditManagerNewParameters)
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
func (it *ICreditManagerNewParametersIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICreditManagerNewParametersIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICreditManagerNewParameters represents a NewParameters event raised by the ICreditManager contract.
type ICreditManagerNewParameters struct {
	MinAmount           *big.Int
	MaxAmount           *big.Int
	MaxLeverage         *big.Int
	FeeInterest         *big.Int
	FeeLiquidation      *big.Int
	LiquidationDiscount *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterNewParameters is a free log retrieval operation binding the contract event 0xa32137411fc7c20db359079cd84af0e2cad58cd7a182a8a5e23e08e554e88bf0.
//
// Solidity: event NewParameters(uint256 minAmount, uint256 maxAmount, uint256 maxLeverage, uint256 feeInterest, uint256 feeLiquidation, uint256 liquidationDiscount)
func (_ICreditManager *ICreditManagerFilterer) FilterNewParameters(opts *bind.FilterOpts) (*ICreditManagerNewParametersIterator, error) {

	logs, sub, err := _ICreditManager.contract.FilterLogs(opts, "NewParameters")
	if err != nil {
		return nil, err
	}
	return &ICreditManagerNewParametersIterator{contract: _ICreditManager.contract, event: "NewParameters", logs: logs, sub: sub}, nil
}

// WatchNewParameters is a free log subscription operation binding the contract event 0xa32137411fc7c20db359079cd84af0e2cad58cd7a182a8a5e23e08e554e88bf0.
//
// Solidity: event NewParameters(uint256 minAmount, uint256 maxAmount, uint256 maxLeverage, uint256 feeInterest, uint256 feeLiquidation, uint256 liquidationDiscount)
func (_ICreditManager *ICreditManagerFilterer) WatchNewParameters(opts *bind.WatchOpts, sink chan<- *ICreditManagerNewParameters) (event.Subscription, error) {

	logs, sub, err := _ICreditManager.contract.WatchLogs(opts, "NewParameters")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICreditManagerNewParameters)
				if err := _ICreditManager.contract.UnpackLog(event, "NewParameters", log); err != nil {
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

// ParseNewParameters is a log parse operation binding the contract event 0xa32137411fc7c20db359079cd84af0e2cad58cd7a182a8a5e23e08e554e88bf0.
//
// Solidity: event NewParameters(uint256 minAmount, uint256 maxAmount, uint256 maxLeverage, uint256 feeInterest, uint256 feeLiquidation, uint256 liquidationDiscount)
func (_ICreditManager *ICreditManagerFilterer) ParseNewParameters(log types.Log) (*ICreditManagerNewParameters, error) {
	event := new(ICreditManagerNewParameters)
	if err := _ICreditManager.contract.UnpackLog(event, "NewParameters", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ICreditManagerOpenCreditAccountIterator is returned from FilterOpenCreditAccount and is used to iterate over the raw logs and unpacked data for OpenCreditAccount events raised by the ICreditManager contract.
type ICreditManagerOpenCreditAccountIterator struct {
	Event *ICreditManagerOpenCreditAccount // Event containing the contract specifics and raw log

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
func (it *ICreditManagerOpenCreditAccountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICreditManagerOpenCreditAccount)
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
		it.Event = new(ICreditManagerOpenCreditAccount)
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
func (it *ICreditManagerOpenCreditAccountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICreditManagerOpenCreditAccountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICreditManagerOpenCreditAccount represents a OpenCreditAccount event raised by the ICreditManager contract.
type ICreditManagerOpenCreditAccount struct {
	Sender        common.Address
	OnBehalfOf    common.Address
	CreditAccount common.Address
	Amount        *big.Int
	BorrowAmount  *big.Int
	ReferralCode  *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOpenCreditAccount is a free log retrieval operation binding the contract event 0x7b20ae77867a263a1074203a2da261ef0d096c99395c59c9d4a0104b9f334a27.
//
// Solidity: event OpenCreditAccount(address indexed sender, address indexed onBehalfOf, address indexed creditAccount, uint256 amount, uint256 borrowAmount, uint256 referralCode)
func (_ICreditManager *ICreditManagerFilterer) FilterOpenCreditAccount(opts *bind.FilterOpts, sender []common.Address, onBehalfOf []common.Address, creditAccount []common.Address) (*ICreditManagerOpenCreditAccountIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}
	var creditAccountRule []interface{}
	for _, creditAccountItem := range creditAccount {
		creditAccountRule = append(creditAccountRule, creditAccountItem)
	}

	logs, sub, err := _ICreditManager.contract.FilterLogs(opts, "OpenCreditAccount", senderRule, onBehalfOfRule, creditAccountRule)
	if err != nil {
		return nil, err
	}
	return &ICreditManagerOpenCreditAccountIterator{contract: _ICreditManager.contract, event: "OpenCreditAccount", logs: logs, sub: sub}, nil
}

// WatchOpenCreditAccount is a free log subscription operation binding the contract event 0x7b20ae77867a263a1074203a2da261ef0d096c99395c59c9d4a0104b9f334a27.
//
// Solidity: event OpenCreditAccount(address indexed sender, address indexed onBehalfOf, address indexed creditAccount, uint256 amount, uint256 borrowAmount, uint256 referralCode)
func (_ICreditManager *ICreditManagerFilterer) WatchOpenCreditAccount(opts *bind.WatchOpts, sink chan<- *ICreditManagerOpenCreditAccount, sender []common.Address, onBehalfOf []common.Address, creditAccount []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}
	var creditAccountRule []interface{}
	for _, creditAccountItem := range creditAccount {
		creditAccountRule = append(creditAccountRule, creditAccountItem)
	}

	logs, sub, err := _ICreditManager.contract.WatchLogs(opts, "OpenCreditAccount", senderRule, onBehalfOfRule, creditAccountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICreditManagerOpenCreditAccount)
				if err := _ICreditManager.contract.UnpackLog(event, "OpenCreditAccount", log); err != nil {
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

// ParseOpenCreditAccount is a log parse operation binding the contract event 0x7b20ae77867a263a1074203a2da261ef0d096c99395c59c9d4a0104b9f334a27.
//
// Solidity: event OpenCreditAccount(address indexed sender, address indexed onBehalfOf, address indexed creditAccount, uint256 amount, uint256 borrowAmount, uint256 referralCode)
func (_ICreditManager *ICreditManagerFilterer) ParseOpenCreditAccount(log types.Log) (*ICreditManagerOpenCreditAccount, error) {
	event := new(ICreditManagerOpenCreditAccount)
	if err := _ICreditManager.contract.UnpackLog(event, "OpenCreditAccount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ICreditManagerRepayCreditAccountIterator is returned from FilterRepayCreditAccount and is used to iterate over the raw logs and unpacked data for RepayCreditAccount events raised by the ICreditManager contract.
type ICreditManagerRepayCreditAccountIterator struct {
	Event *ICreditManagerRepayCreditAccount // Event containing the contract specifics and raw log

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
func (it *ICreditManagerRepayCreditAccountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICreditManagerRepayCreditAccount)
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
		it.Event = new(ICreditManagerRepayCreditAccount)
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
func (it *ICreditManagerRepayCreditAccountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICreditManagerRepayCreditAccountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICreditManagerRepayCreditAccount represents a RepayCreditAccount event raised by the ICreditManager contract.
type ICreditManagerRepayCreditAccount struct {
	Owner common.Address
	To    common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRepayCreditAccount is a free log retrieval operation binding the contract event 0xe7c7987373a0cc4913d307f23ab8ef02e0333a2af445065e2ef7636cffc6daa7.
//
// Solidity: event RepayCreditAccount(address indexed owner, address indexed to)
func (_ICreditManager *ICreditManagerFilterer) FilterRepayCreditAccount(opts *bind.FilterOpts, owner []common.Address, to []common.Address) (*ICreditManagerRepayCreditAccountIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ICreditManager.contract.FilterLogs(opts, "RepayCreditAccount", ownerRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ICreditManagerRepayCreditAccountIterator{contract: _ICreditManager.contract, event: "RepayCreditAccount", logs: logs, sub: sub}, nil
}

// WatchRepayCreditAccount is a free log subscription operation binding the contract event 0xe7c7987373a0cc4913d307f23ab8ef02e0333a2af445065e2ef7636cffc6daa7.
//
// Solidity: event RepayCreditAccount(address indexed owner, address indexed to)
func (_ICreditManager *ICreditManagerFilterer) WatchRepayCreditAccount(opts *bind.WatchOpts, sink chan<- *ICreditManagerRepayCreditAccount, owner []common.Address, to []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ICreditManager.contract.WatchLogs(opts, "RepayCreditAccount", ownerRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICreditManagerRepayCreditAccount)
				if err := _ICreditManager.contract.UnpackLog(event, "RepayCreditAccount", log); err != nil {
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

// ParseRepayCreditAccount is a log parse operation binding the contract event 0xe7c7987373a0cc4913d307f23ab8ef02e0333a2af445065e2ef7636cffc6daa7.
//
// Solidity: event RepayCreditAccount(address indexed owner, address indexed to)
func (_ICreditManager *ICreditManagerFilterer) ParseRepayCreditAccount(log types.Log) (*ICreditManagerRepayCreditAccount, error) {
	event := new(ICreditManagerRepayCreditAccount)
	if err := _ICreditManager.contract.UnpackLog(event, "RepayCreditAccount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ICreditManagerTransferAccountIterator is returned from FilterTransferAccount and is used to iterate over the raw logs and unpacked data for TransferAccount events raised by the ICreditManager contract.
type ICreditManagerTransferAccountIterator struct {
	Event *ICreditManagerTransferAccount // Event containing the contract specifics and raw log

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
func (it *ICreditManagerTransferAccountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICreditManagerTransferAccount)
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
		it.Event = new(ICreditManagerTransferAccount)
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
func (it *ICreditManagerTransferAccountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICreditManagerTransferAccountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICreditManagerTransferAccount represents a TransferAccount event raised by the ICreditManager contract.
type ICreditManagerTransferAccount struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferAccount is a free log retrieval operation binding the contract event 0x93c70cc9715bef0d83edf2095f3595402279d274f402a73ffc17f1bcb19d863d.
//
// Solidity: event TransferAccount(address indexed oldOwner, address indexed newOwner)
func (_ICreditManager *ICreditManagerFilterer) FilterTransferAccount(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*ICreditManagerTransferAccountIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ICreditManager.contract.FilterLogs(opts, "TransferAccount", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ICreditManagerTransferAccountIterator{contract: _ICreditManager.contract, event: "TransferAccount", logs: logs, sub: sub}, nil
}

// WatchTransferAccount is a free log subscription operation binding the contract event 0x93c70cc9715bef0d83edf2095f3595402279d274f402a73ffc17f1bcb19d863d.
//
// Solidity: event TransferAccount(address indexed oldOwner, address indexed newOwner)
func (_ICreditManager *ICreditManagerFilterer) WatchTransferAccount(opts *bind.WatchOpts, sink chan<- *ICreditManagerTransferAccount, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ICreditManager.contract.WatchLogs(opts, "TransferAccount", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICreditManagerTransferAccount)
				if err := _ICreditManager.contract.UnpackLog(event, "TransferAccount", log); err != nil {
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

// ParseTransferAccount is a log parse operation binding the contract event 0x93c70cc9715bef0d83edf2095f3595402279d274f402a73ffc17f1bcb19d863d.
//
// Solidity: event TransferAccount(address indexed oldOwner, address indexed newOwner)
func (_ICreditManager *ICreditManagerFilterer) ParseTransferAccount(log types.Log) (*ICreditManagerTransferAccount, error) {
	event := new(ICreditManagerTransferAccount)
	if err := _ICreditManager.contract.UnpackLog(event, "TransferAccount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
