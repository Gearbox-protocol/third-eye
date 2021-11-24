// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iCreditFilter

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

// ICreditFilterMetaData contains all meta data concerning the ICreditFilter contract.
var ICreditFilterMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"protocol\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"adapter\",\"type\":\"address\"}],\"name\":\"ContractAllowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"protocol\",\"type\":\"address\"}],\"name\":\"ContractForbidden\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chiThreshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fastCheckDelay\",\"type\":\"uint256\"}],\"name\":\"NewFastCheckParameters\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidityThreshold\",\"type\":\"uint256\"}],\"name\":\"TokenAllowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"state\",\"type\":\"bool\"}],\"name\":\"TransferAccountAllowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pugin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"state\",\"type\":\"bool\"}],\"name\":\"TransferPluginAllowed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"targetContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"adapter\",\"type\":\"address\"}],\"name\":\"allowContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidationThreshold\",\"type\":\"uint256\"}],\"name\":\"allowToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"allowanceForAccountTransfers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"allowedContracts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allowedContractsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"allowedTokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allowedTokensCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"state\",\"type\":\"bool\"}],\"name\":\"approveAccountTransfers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"}],\"name\":\"calcCreditAccountAccruedInterest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"}],\"name\":\"calcCreditAccountHealthFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"}],\"name\":\"calcThresholdWeightedValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"}],\"name\":\"calcTotalValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"checkAndEnableToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"checkCollateralChange\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"amountIn\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amountOut\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"tokenIn\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"tokenOut\",\"type\":\"address[]\"}],\"name\":\"checkMultiTokenCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"poolService\",\"type\":\"address\"}],\"name\":\"connectCreditManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"allowedContract\",\"type\":\"address\"}],\"name\":\"contractToAdapter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"}],\"name\":\"enabledTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"targetContract\",\"type\":\"address\"}],\"name\":\"forbidContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getCreditAccountTokenById\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"twv\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"}],\"name\":\"initEnabledTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"isTokenAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"liquidationThresholds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"onwer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"}],\"name\":\"revertIfAccountTransferIsNotAllowed\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minHealthFactor\",\"type\":\"uint256\"}],\"name\":\"revertIfCantIncreaseBorrowing\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"revertIfTokenNotAllowed\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlyingToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateUnderlyingTokenLiquidationThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ICreditFilterABI is the input ABI used to generate the binding from.
// Deprecated: Use ICreditFilterMetaData.ABI instead.
var ICreditFilterABI = ICreditFilterMetaData.ABI

// ICreditFilter is an auto generated Go binding around an Ethereum contract.
type ICreditFilter struct {
	ICreditFilterCaller     // Read-only binding to the contract
	ICreditFilterTransactor // Write-only binding to the contract
	ICreditFilterFilterer   // Log filterer for contract events
}

// ICreditFilterCaller is an auto generated read-only Go binding around an Ethereum contract.
type ICreditFilterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICreditFilterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ICreditFilterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICreditFilterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ICreditFilterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICreditFilterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ICreditFilterSession struct {
	Contract     *ICreditFilter    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ICreditFilterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ICreditFilterCallerSession struct {
	Contract *ICreditFilterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ICreditFilterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ICreditFilterTransactorSession struct {
	Contract     *ICreditFilterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ICreditFilterRaw is an auto generated low-level Go binding around an Ethereum contract.
type ICreditFilterRaw struct {
	Contract *ICreditFilter // Generic contract binding to access the raw methods on
}

// ICreditFilterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ICreditFilterCallerRaw struct {
	Contract *ICreditFilterCaller // Generic read-only contract binding to access the raw methods on
}

// ICreditFilterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ICreditFilterTransactorRaw struct {
	Contract *ICreditFilterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewICreditFilter creates a new instance of ICreditFilter, bound to a specific deployed contract.
func NewICreditFilter(address common.Address, backend bind.ContractBackend) (*ICreditFilter, error) {
	contract, err := bindICreditFilter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ICreditFilter{ICreditFilterCaller: ICreditFilterCaller{contract: contract}, ICreditFilterTransactor: ICreditFilterTransactor{contract: contract}, ICreditFilterFilterer: ICreditFilterFilterer{contract: contract}}, nil
}

// NewICreditFilterCaller creates a new read-only instance of ICreditFilter, bound to a specific deployed contract.
func NewICreditFilterCaller(address common.Address, caller bind.ContractCaller) (*ICreditFilterCaller, error) {
	contract, err := bindICreditFilter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ICreditFilterCaller{contract: contract}, nil
}

// NewICreditFilterTransactor creates a new write-only instance of ICreditFilter, bound to a specific deployed contract.
func NewICreditFilterTransactor(address common.Address, transactor bind.ContractTransactor) (*ICreditFilterTransactor, error) {
	contract, err := bindICreditFilter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ICreditFilterTransactor{contract: contract}, nil
}

// NewICreditFilterFilterer creates a new log filterer instance of ICreditFilter, bound to a specific deployed contract.
func NewICreditFilterFilterer(address common.Address, filterer bind.ContractFilterer) (*ICreditFilterFilterer, error) {
	contract, err := bindICreditFilter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ICreditFilterFilterer{contract: contract}, nil
}

// bindICreditFilter binds a generic wrapper to an already deployed contract.
func bindICreditFilter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ICreditFilterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICreditFilter *ICreditFilterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICreditFilter.Contract.ICreditFilterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICreditFilter *ICreditFilterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICreditFilter.Contract.ICreditFilterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICreditFilter *ICreditFilterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICreditFilter.Contract.ICreditFilterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICreditFilter *ICreditFilterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICreditFilter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICreditFilter *ICreditFilterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICreditFilter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICreditFilter *ICreditFilterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICreditFilter.Contract.contract.Transact(opts, method, params...)
}

// AllowanceForAccountTransfers is a free data retrieval call binding the contract method 0x5a29be45.
//
// Solidity: function allowanceForAccountTransfers(address from, address to) view returns(bool)
func (_ICreditFilter *ICreditFilterCaller) AllowanceForAccountTransfers(opts *bind.CallOpts, from common.Address, to common.Address) (bool, error) {
	var out []interface{}
	err := _ICreditFilter.contract.Call(opts, &out, "allowanceForAccountTransfers", from, to)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowanceForAccountTransfers is a free data retrieval call binding the contract method 0x5a29be45.
//
// Solidity: function allowanceForAccountTransfers(address from, address to) view returns(bool)
func (_ICreditFilter *ICreditFilterSession) AllowanceForAccountTransfers(from common.Address, to common.Address) (bool, error) {
	return _ICreditFilter.Contract.AllowanceForAccountTransfers(&_ICreditFilter.CallOpts, from, to)
}

// AllowanceForAccountTransfers is a free data retrieval call binding the contract method 0x5a29be45.
//
// Solidity: function allowanceForAccountTransfers(address from, address to) view returns(bool)
func (_ICreditFilter *ICreditFilterCallerSession) AllowanceForAccountTransfers(from common.Address, to common.Address) (bool, error) {
	return _ICreditFilter.Contract.AllowanceForAccountTransfers(&_ICreditFilter.CallOpts, from, to)
}

// AllowedContracts is a free data retrieval call binding the contract method 0x5094cb4f.
//
// Solidity: function allowedContracts(uint256 id) view returns(address)
func (_ICreditFilter *ICreditFilterCaller) AllowedContracts(opts *bind.CallOpts, id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ICreditFilter.contract.Call(opts, &out, "allowedContracts", id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllowedContracts is a free data retrieval call binding the contract method 0x5094cb4f.
//
// Solidity: function allowedContracts(uint256 id) view returns(address)
func (_ICreditFilter *ICreditFilterSession) AllowedContracts(id *big.Int) (common.Address, error) {
	return _ICreditFilter.Contract.AllowedContracts(&_ICreditFilter.CallOpts, id)
}

// AllowedContracts is a free data retrieval call binding the contract method 0x5094cb4f.
//
// Solidity: function allowedContracts(uint256 id) view returns(address)
func (_ICreditFilter *ICreditFilterCallerSession) AllowedContracts(id *big.Int) (common.Address, error) {
	return _ICreditFilter.Contract.AllowedContracts(&_ICreditFilter.CallOpts, id)
}

// AllowedContractsCount is a free data retrieval call binding the contract method 0x50e036ff.
//
// Solidity: function allowedContractsCount() view returns(uint256)
func (_ICreditFilter *ICreditFilterCaller) AllowedContractsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ICreditFilter.contract.Call(opts, &out, "allowedContractsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllowedContractsCount is a free data retrieval call binding the contract method 0x50e036ff.
//
// Solidity: function allowedContractsCount() view returns(uint256)
func (_ICreditFilter *ICreditFilterSession) AllowedContractsCount() (*big.Int, error) {
	return _ICreditFilter.Contract.AllowedContractsCount(&_ICreditFilter.CallOpts)
}

// AllowedContractsCount is a free data retrieval call binding the contract method 0x50e036ff.
//
// Solidity: function allowedContractsCount() view returns(uint256)
func (_ICreditFilter *ICreditFilterCallerSession) AllowedContractsCount() (*big.Int, error) {
	return _ICreditFilter.Contract.AllowedContractsCount(&_ICreditFilter.CallOpts)
}

// AllowedTokens is a free data retrieval call binding the contract method 0x5e5f2e26.
//
// Solidity: function allowedTokens(uint256 id) view returns(address)
func (_ICreditFilter *ICreditFilterCaller) AllowedTokens(opts *bind.CallOpts, id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ICreditFilter.contract.Call(opts, &out, "allowedTokens", id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllowedTokens is a free data retrieval call binding the contract method 0x5e5f2e26.
//
// Solidity: function allowedTokens(uint256 id) view returns(address)
func (_ICreditFilter *ICreditFilterSession) AllowedTokens(id *big.Int) (common.Address, error) {
	return _ICreditFilter.Contract.AllowedTokens(&_ICreditFilter.CallOpts, id)
}

// AllowedTokens is a free data retrieval call binding the contract method 0x5e5f2e26.
//
// Solidity: function allowedTokens(uint256 id) view returns(address)
func (_ICreditFilter *ICreditFilterCallerSession) AllowedTokens(id *big.Int) (common.Address, error) {
	return _ICreditFilter.Contract.AllowedTokens(&_ICreditFilter.CallOpts, id)
}

// AllowedTokensCount is a free data retrieval call binding the contract method 0x20a05ff7.
//
// Solidity: function allowedTokensCount() view returns(uint256)
func (_ICreditFilter *ICreditFilterCaller) AllowedTokensCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ICreditFilter.contract.Call(opts, &out, "allowedTokensCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllowedTokensCount is a free data retrieval call binding the contract method 0x20a05ff7.
//
// Solidity: function allowedTokensCount() view returns(uint256)
func (_ICreditFilter *ICreditFilterSession) AllowedTokensCount() (*big.Int, error) {
	return _ICreditFilter.Contract.AllowedTokensCount(&_ICreditFilter.CallOpts)
}

// AllowedTokensCount is a free data retrieval call binding the contract method 0x20a05ff7.
//
// Solidity: function allowedTokensCount() view returns(uint256)
func (_ICreditFilter *ICreditFilterCallerSession) AllowedTokensCount() (*big.Int, error) {
	return _ICreditFilter.Contract.AllowedTokensCount(&_ICreditFilter.CallOpts)
}

// CalcCreditAccountAccruedInterest is a free data retrieval call binding the contract method 0x3192195c.
//
// Solidity: function calcCreditAccountAccruedInterest(address creditAccount) view returns(uint256)
func (_ICreditFilter *ICreditFilterCaller) CalcCreditAccountAccruedInterest(opts *bind.CallOpts, creditAccount common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ICreditFilter.contract.Call(opts, &out, "calcCreditAccountAccruedInterest", creditAccount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcCreditAccountAccruedInterest is a free data retrieval call binding the contract method 0x3192195c.
//
// Solidity: function calcCreditAccountAccruedInterest(address creditAccount) view returns(uint256)
func (_ICreditFilter *ICreditFilterSession) CalcCreditAccountAccruedInterest(creditAccount common.Address) (*big.Int, error) {
	return _ICreditFilter.Contract.CalcCreditAccountAccruedInterest(&_ICreditFilter.CallOpts, creditAccount)
}

// CalcCreditAccountAccruedInterest is a free data retrieval call binding the contract method 0x3192195c.
//
// Solidity: function calcCreditAccountAccruedInterest(address creditAccount) view returns(uint256)
func (_ICreditFilter *ICreditFilterCallerSession) CalcCreditAccountAccruedInterest(creditAccount common.Address) (*big.Int, error) {
	return _ICreditFilter.Contract.CalcCreditAccountAccruedInterest(&_ICreditFilter.CallOpts, creditAccount)
}

// CalcCreditAccountHealthFactor is a free data retrieval call binding the contract method 0xdfd59465.
//
// Solidity: function calcCreditAccountHealthFactor(address creditAccount) view returns(uint256)
func (_ICreditFilter *ICreditFilterCaller) CalcCreditAccountHealthFactor(opts *bind.CallOpts, creditAccount common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ICreditFilter.contract.Call(opts, &out, "calcCreditAccountHealthFactor", creditAccount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcCreditAccountHealthFactor is a free data retrieval call binding the contract method 0xdfd59465.
//
// Solidity: function calcCreditAccountHealthFactor(address creditAccount) view returns(uint256)
func (_ICreditFilter *ICreditFilterSession) CalcCreditAccountHealthFactor(creditAccount common.Address) (*big.Int, error) {
	return _ICreditFilter.Contract.CalcCreditAccountHealthFactor(&_ICreditFilter.CallOpts, creditAccount)
}

// CalcCreditAccountHealthFactor is a free data retrieval call binding the contract method 0xdfd59465.
//
// Solidity: function calcCreditAccountHealthFactor(address creditAccount) view returns(uint256)
func (_ICreditFilter *ICreditFilterCallerSession) CalcCreditAccountHealthFactor(creditAccount common.Address) (*big.Int, error) {
	return _ICreditFilter.Contract.CalcCreditAccountHealthFactor(&_ICreditFilter.CallOpts, creditAccount)
}

// CalcThresholdWeightedValue is a free data retrieval call binding the contract method 0x90b1300a.
//
// Solidity: function calcThresholdWeightedValue(address creditAccount) view returns(uint256 total)
func (_ICreditFilter *ICreditFilterCaller) CalcThresholdWeightedValue(opts *bind.CallOpts, creditAccount common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ICreditFilter.contract.Call(opts, &out, "calcThresholdWeightedValue", creditAccount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcThresholdWeightedValue is a free data retrieval call binding the contract method 0x90b1300a.
//
// Solidity: function calcThresholdWeightedValue(address creditAccount) view returns(uint256 total)
func (_ICreditFilter *ICreditFilterSession) CalcThresholdWeightedValue(creditAccount common.Address) (*big.Int, error) {
	return _ICreditFilter.Contract.CalcThresholdWeightedValue(&_ICreditFilter.CallOpts, creditAccount)
}

// CalcThresholdWeightedValue is a free data retrieval call binding the contract method 0x90b1300a.
//
// Solidity: function calcThresholdWeightedValue(address creditAccount) view returns(uint256 total)
func (_ICreditFilter *ICreditFilterCallerSession) CalcThresholdWeightedValue(creditAccount common.Address) (*big.Int, error) {
	return _ICreditFilter.Contract.CalcThresholdWeightedValue(&_ICreditFilter.CallOpts, creditAccount)
}

// CalcTotalValue is a free data retrieval call binding the contract method 0xc7de38a6.
//
// Solidity: function calcTotalValue(address creditAccount) view returns(uint256 total)
func (_ICreditFilter *ICreditFilterCaller) CalcTotalValue(opts *bind.CallOpts, creditAccount common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ICreditFilter.contract.Call(opts, &out, "calcTotalValue", creditAccount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcTotalValue is a free data retrieval call binding the contract method 0xc7de38a6.
//
// Solidity: function calcTotalValue(address creditAccount) view returns(uint256 total)
func (_ICreditFilter *ICreditFilterSession) CalcTotalValue(creditAccount common.Address) (*big.Int, error) {
	return _ICreditFilter.Contract.CalcTotalValue(&_ICreditFilter.CallOpts, creditAccount)
}

// CalcTotalValue is a free data retrieval call binding the contract method 0xc7de38a6.
//
// Solidity: function calcTotalValue(address creditAccount) view returns(uint256 total)
func (_ICreditFilter *ICreditFilterCallerSession) CalcTotalValue(creditAccount common.Address) (*big.Int, error) {
	return _ICreditFilter.Contract.CalcTotalValue(&_ICreditFilter.CallOpts, creditAccount)
}

// ContractToAdapter is a free data retrieval call binding the contract method 0xfdd57645.
//
// Solidity: function contractToAdapter(address allowedContract) view returns(address)
func (_ICreditFilter *ICreditFilterCaller) ContractToAdapter(opts *bind.CallOpts, allowedContract common.Address) (common.Address, error) {
	var out []interface{}
	err := _ICreditFilter.contract.Call(opts, &out, "contractToAdapter", allowedContract)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ContractToAdapter is a free data retrieval call binding the contract method 0xfdd57645.
//
// Solidity: function contractToAdapter(address allowedContract) view returns(address)
func (_ICreditFilter *ICreditFilterSession) ContractToAdapter(allowedContract common.Address) (common.Address, error) {
	return _ICreditFilter.Contract.ContractToAdapter(&_ICreditFilter.CallOpts, allowedContract)
}

// ContractToAdapter is a free data retrieval call binding the contract method 0xfdd57645.
//
// Solidity: function contractToAdapter(address allowedContract) view returns(address)
func (_ICreditFilter *ICreditFilterCallerSession) ContractToAdapter(allowedContract common.Address) (common.Address, error) {
	return _ICreditFilter.Contract.ContractToAdapter(&_ICreditFilter.CallOpts, allowedContract)
}

// EnabledTokens is a free data retrieval call binding the contract method 0xb451cecc.
//
// Solidity: function enabledTokens(address creditAccount) view returns(uint256)
func (_ICreditFilter *ICreditFilterCaller) EnabledTokens(opts *bind.CallOpts, creditAccount common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ICreditFilter.contract.Call(opts, &out, "enabledTokens", creditAccount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EnabledTokens is a free data retrieval call binding the contract method 0xb451cecc.
//
// Solidity: function enabledTokens(address creditAccount) view returns(uint256)
func (_ICreditFilter *ICreditFilterSession) EnabledTokens(creditAccount common.Address) (*big.Int, error) {
	return _ICreditFilter.Contract.EnabledTokens(&_ICreditFilter.CallOpts, creditAccount)
}

// EnabledTokens is a free data retrieval call binding the contract method 0xb451cecc.
//
// Solidity: function enabledTokens(address creditAccount) view returns(uint256)
func (_ICreditFilter *ICreditFilterCallerSession) EnabledTokens(creditAccount common.Address) (*big.Int, error) {
	return _ICreditFilter.Contract.EnabledTokens(&_ICreditFilter.CallOpts, creditAccount)
}

// GetCreditAccountTokenById is a free data retrieval call binding the contract method 0xaf0a6502.
//
// Solidity: function getCreditAccountTokenById(address creditAccount, uint256 id) view returns(address token, uint256 balance, uint256 tv, uint256 twv)
func (_ICreditFilter *ICreditFilterCaller) GetCreditAccountTokenById(opts *bind.CallOpts, creditAccount common.Address, id *big.Int) (struct {
	Token   common.Address
	Balance *big.Int
	Tv      *big.Int
	Twv     *big.Int
}, error) {
	var out []interface{}
	err := _ICreditFilter.contract.Call(opts, &out, "getCreditAccountTokenById", creditAccount, id)

	outstruct := new(struct {
		Token   common.Address
		Balance *big.Int
		Tv      *big.Int
		Twv     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Token = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Balance = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Tv = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Twv = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetCreditAccountTokenById is a free data retrieval call binding the contract method 0xaf0a6502.
//
// Solidity: function getCreditAccountTokenById(address creditAccount, uint256 id) view returns(address token, uint256 balance, uint256 tv, uint256 twv)
func (_ICreditFilter *ICreditFilterSession) GetCreditAccountTokenById(creditAccount common.Address, id *big.Int) (struct {
	Token   common.Address
	Balance *big.Int
	Tv      *big.Int
	Twv     *big.Int
}, error) {
	return _ICreditFilter.Contract.GetCreditAccountTokenById(&_ICreditFilter.CallOpts, creditAccount, id)
}

// GetCreditAccountTokenById is a free data retrieval call binding the contract method 0xaf0a6502.
//
// Solidity: function getCreditAccountTokenById(address creditAccount, uint256 id) view returns(address token, uint256 balance, uint256 tv, uint256 twv)
func (_ICreditFilter *ICreditFilterCallerSession) GetCreditAccountTokenById(creditAccount common.Address, id *big.Int) (struct {
	Token   common.Address
	Balance *big.Int
	Tv      *big.Int
	Twv     *big.Int
}, error) {
	return _ICreditFilter.Contract.GetCreditAccountTokenById(&_ICreditFilter.CallOpts, creditAccount, id)
}

// IsTokenAllowed is a free data retrieval call binding the contract method 0xf9eaee0d.
//
// Solidity: function isTokenAllowed(address token) view returns(bool)
func (_ICreditFilter *ICreditFilterCaller) IsTokenAllowed(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _ICreditFilter.contract.Call(opts, &out, "isTokenAllowed", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTokenAllowed is a free data retrieval call binding the contract method 0xf9eaee0d.
//
// Solidity: function isTokenAllowed(address token) view returns(bool)
func (_ICreditFilter *ICreditFilterSession) IsTokenAllowed(token common.Address) (bool, error) {
	return _ICreditFilter.Contract.IsTokenAllowed(&_ICreditFilter.CallOpts, token)
}

// IsTokenAllowed is a free data retrieval call binding the contract method 0xf9eaee0d.
//
// Solidity: function isTokenAllowed(address token) view returns(bool)
func (_ICreditFilter *ICreditFilterCallerSession) IsTokenAllowed(token common.Address) (bool, error) {
	return _ICreditFilter.Contract.IsTokenAllowed(&_ICreditFilter.CallOpts, token)
}

// LiquidationThresholds is a free data retrieval call binding the contract method 0x78327438.
//
// Solidity: function liquidationThresholds(address token) view returns(uint256)
func (_ICreditFilter *ICreditFilterCaller) LiquidationThresholds(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ICreditFilter.contract.Call(opts, &out, "liquidationThresholds", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LiquidationThresholds is a free data retrieval call binding the contract method 0x78327438.
//
// Solidity: function liquidationThresholds(address token) view returns(uint256)
func (_ICreditFilter *ICreditFilterSession) LiquidationThresholds(token common.Address) (*big.Int, error) {
	return _ICreditFilter.Contract.LiquidationThresholds(&_ICreditFilter.CallOpts, token)
}

// LiquidationThresholds is a free data retrieval call binding the contract method 0x78327438.
//
// Solidity: function liquidationThresholds(address token) view returns(uint256)
func (_ICreditFilter *ICreditFilterCallerSession) LiquidationThresholds(token common.Address) (*big.Int, error) {
	return _ICreditFilter.Contract.LiquidationThresholds(&_ICreditFilter.CallOpts, token)
}

// PriceOracle is a free data retrieval call binding the contract method 0x2630c12f.
//
// Solidity: function priceOracle() view returns(address)
func (_ICreditFilter *ICreditFilterCaller) PriceOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ICreditFilter.contract.Call(opts, &out, "priceOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PriceOracle is a free data retrieval call binding the contract method 0x2630c12f.
//
// Solidity: function priceOracle() view returns(address)
func (_ICreditFilter *ICreditFilterSession) PriceOracle() (common.Address, error) {
	return _ICreditFilter.Contract.PriceOracle(&_ICreditFilter.CallOpts)
}

// PriceOracle is a free data retrieval call binding the contract method 0x2630c12f.
//
// Solidity: function priceOracle() view returns(address)
func (_ICreditFilter *ICreditFilterCallerSession) PriceOracle() (common.Address, error) {
	return _ICreditFilter.Contract.PriceOracle(&_ICreditFilter.CallOpts)
}

// RevertIfAccountTransferIsNotAllowed is a free data retrieval call binding the contract method 0x3b00ae70.
//
// Solidity: function revertIfAccountTransferIsNotAllowed(address onwer, address creditAccount) view returns()
func (_ICreditFilter *ICreditFilterCaller) RevertIfAccountTransferIsNotAllowed(opts *bind.CallOpts, onwer common.Address, creditAccount common.Address) error {
	var out []interface{}
	err := _ICreditFilter.contract.Call(opts, &out, "revertIfAccountTransferIsNotAllowed", onwer, creditAccount)

	if err != nil {
		return err
	}

	return err

}

// RevertIfAccountTransferIsNotAllowed is a free data retrieval call binding the contract method 0x3b00ae70.
//
// Solidity: function revertIfAccountTransferIsNotAllowed(address onwer, address creditAccount) view returns()
func (_ICreditFilter *ICreditFilterSession) RevertIfAccountTransferIsNotAllowed(onwer common.Address, creditAccount common.Address) error {
	return _ICreditFilter.Contract.RevertIfAccountTransferIsNotAllowed(&_ICreditFilter.CallOpts, onwer, creditAccount)
}

// RevertIfAccountTransferIsNotAllowed is a free data retrieval call binding the contract method 0x3b00ae70.
//
// Solidity: function revertIfAccountTransferIsNotAllowed(address onwer, address creditAccount) view returns()
func (_ICreditFilter *ICreditFilterCallerSession) RevertIfAccountTransferIsNotAllowed(onwer common.Address, creditAccount common.Address) error {
	return _ICreditFilter.Contract.RevertIfAccountTransferIsNotAllowed(&_ICreditFilter.CallOpts, onwer, creditAccount)
}

// RevertIfCantIncreaseBorrowing is a free data retrieval call binding the contract method 0xa5757517.
//
// Solidity: function revertIfCantIncreaseBorrowing(address creditAccount, uint256 minHealthFactor) view returns()
func (_ICreditFilter *ICreditFilterCaller) RevertIfCantIncreaseBorrowing(opts *bind.CallOpts, creditAccount common.Address, minHealthFactor *big.Int) error {
	var out []interface{}
	err := _ICreditFilter.contract.Call(opts, &out, "revertIfCantIncreaseBorrowing", creditAccount, minHealthFactor)

	if err != nil {
		return err
	}

	return err

}

// RevertIfCantIncreaseBorrowing is a free data retrieval call binding the contract method 0xa5757517.
//
// Solidity: function revertIfCantIncreaseBorrowing(address creditAccount, uint256 minHealthFactor) view returns()
func (_ICreditFilter *ICreditFilterSession) RevertIfCantIncreaseBorrowing(creditAccount common.Address, minHealthFactor *big.Int) error {
	return _ICreditFilter.Contract.RevertIfCantIncreaseBorrowing(&_ICreditFilter.CallOpts, creditAccount, minHealthFactor)
}

// RevertIfCantIncreaseBorrowing is a free data retrieval call binding the contract method 0xa5757517.
//
// Solidity: function revertIfCantIncreaseBorrowing(address creditAccount, uint256 minHealthFactor) view returns()
func (_ICreditFilter *ICreditFilterCallerSession) RevertIfCantIncreaseBorrowing(creditAccount common.Address, minHealthFactor *big.Int) error {
	return _ICreditFilter.Contract.RevertIfCantIncreaseBorrowing(&_ICreditFilter.CallOpts, creditAccount, minHealthFactor)
}

// RevertIfTokenNotAllowed is a free data retrieval call binding the contract method 0x7dd0ba82.
//
// Solidity: function revertIfTokenNotAllowed(address token) view returns()
func (_ICreditFilter *ICreditFilterCaller) RevertIfTokenNotAllowed(opts *bind.CallOpts, token common.Address) error {
	var out []interface{}
	err := _ICreditFilter.contract.Call(opts, &out, "revertIfTokenNotAllowed", token)

	if err != nil {
		return err
	}

	return err

}

// RevertIfTokenNotAllowed is a free data retrieval call binding the contract method 0x7dd0ba82.
//
// Solidity: function revertIfTokenNotAllowed(address token) view returns()
func (_ICreditFilter *ICreditFilterSession) RevertIfTokenNotAllowed(token common.Address) error {
	return _ICreditFilter.Contract.RevertIfTokenNotAllowed(&_ICreditFilter.CallOpts, token)
}

// RevertIfTokenNotAllowed is a free data retrieval call binding the contract method 0x7dd0ba82.
//
// Solidity: function revertIfTokenNotAllowed(address token) view returns()
func (_ICreditFilter *ICreditFilterCallerSession) RevertIfTokenNotAllowed(token common.Address) error {
	return _ICreditFilter.Contract.RevertIfTokenNotAllowed(&_ICreditFilter.CallOpts, token)
}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_ICreditFilter *ICreditFilterCaller) UnderlyingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ICreditFilter.contract.Call(opts, &out, "underlyingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_ICreditFilter *ICreditFilterSession) UnderlyingToken() (common.Address, error) {
	return _ICreditFilter.Contract.UnderlyingToken(&_ICreditFilter.CallOpts)
}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_ICreditFilter *ICreditFilterCallerSession) UnderlyingToken() (common.Address, error) {
	return _ICreditFilter.Contract.UnderlyingToken(&_ICreditFilter.CallOpts)
}

// AllowContract is a paid mutator transaction binding the contract method 0x7bccacee.
//
// Solidity: function allowContract(address targetContract, address adapter) returns()
func (_ICreditFilter *ICreditFilterTransactor) AllowContract(opts *bind.TransactOpts, targetContract common.Address, adapter common.Address) (*types.Transaction, error) {
	return _ICreditFilter.contract.Transact(opts, "allowContract", targetContract, adapter)
}

// AllowContract is a paid mutator transaction binding the contract method 0x7bccacee.
//
// Solidity: function allowContract(address targetContract, address adapter) returns()
func (_ICreditFilter *ICreditFilterSession) AllowContract(targetContract common.Address, adapter common.Address) (*types.Transaction, error) {
	return _ICreditFilter.Contract.AllowContract(&_ICreditFilter.TransactOpts, targetContract, adapter)
}

// AllowContract is a paid mutator transaction binding the contract method 0x7bccacee.
//
// Solidity: function allowContract(address targetContract, address adapter) returns()
func (_ICreditFilter *ICreditFilterTransactorSession) AllowContract(targetContract common.Address, adapter common.Address) (*types.Transaction, error) {
	return _ICreditFilter.Contract.AllowContract(&_ICreditFilter.TransactOpts, targetContract, adapter)
}

// AllowToken is a paid mutator transaction binding the contract method 0xa147c6c6.
//
// Solidity: function allowToken(address token, uint256 liquidationThreshold) returns()
func (_ICreditFilter *ICreditFilterTransactor) AllowToken(opts *bind.TransactOpts, token common.Address, liquidationThreshold *big.Int) (*types.Transaction, error) {
	return _ICreditFilter.contract.Transact(opts, "allowToken", token, liquidationThreshold)
}

// AllowToken is a paid mutator transaction binding the contract method 0xa147c6c6.
//
// Solidity: function allowToken(address token, uint256 liquidationThreshold) returns()
func (_ICreditFilter *ICreditFilterSession) AllowToken(token common.Address, liquidationThreshold *big.Int) (*types.Transaction, error) {
	return _ICreditFilter.Contract.AllowToken(&_ICreditFilter.TransactOpts, token, liquidationThreshold)
}

// AllowToken is a paid mutator transaction binding the contract method 0xa147c6c6.
//
// Solidity: function allowToken(address token, uint256 liquidationThreshold) returns()
func (_ICreditFilter *ICreditFilterTransactorSession) AllowToken(token common.Address, liquidationThreshold *big.Int) (*types.Transaction, error) {
	return _ICreditFilter.Contract.AllowToken(&_ICreditFilter.TransactOpts, token, liquidationThreshold)
}

// ApproveAccountTransfers is a paid mutator transaction binding the contract method 0x5f27212a.
//
// Solidity: function approveAccountTransfers(address from, bool state) returns()
func (_ICreditFilter *ICreditFilterTransactor) ApproveAccountTransfers(opts *bind.TransactOpts, from common.Address, state bool) (*types.Transaction, error) {
	return _ICreditFilter.contract.Transact(opts, "approveAccountTransfers", from, state)
}

// ApproveAccountTransfers is a paid mutator transaction binding the contract method 0x5f27212a.
//
// Solidity: function approveAccountTransfers(address from, bool state) returns()
func (_ICreditFilter *ICreditFilterSession) ApproveAccountTransfers(from common.Address, state bool) (*types.Transaction, error) {
	return _ICreditFilter.Contract.ApproveAccountTransfers(&_ICreditFilter.TransactOpts, from, state)
}

// ApproveAccountTransfers is a paid mutator transaction binding the contract method 0x5f27212a.
//
// Solidity: function approveAccountTransfers(address from, bool state) returns()
func (_ICreditFilter *ICreditFilterTransactorSession) ApproveAccountTransfers(from common.Address, state bool) (*types.Transaction, error) {
	return _ICreditFilter.Contract.ApproveAccountTransfers(&_ICreditFilter.TransactOpts, from, state)
}

// CheckAndEnableToken is a paid mutator transaction binding the contract method 0x51e3f160.
//
// Solidity: function checkAndEnableToken(address creditAccount, address token) returns()
func (_ICreditFilter *ICreditFilterTransactor) CheckAndEnableToken(opts *bind.TransactOpts, creditAccount common.Address, token common.Address) (*types.Transaction, error) {
	return _ICreditFilter.contract.Transact(opts, "checkAndEnableToken", creditAccount, token)
}

// CheckAndEnableToken is a paid mutator transaction binding the contract method 0x51e3f160.
//
// Solidity: function checkAndEnableToken(address creditAccount, address token) returns()
func (_ICreditFilter *ICreditFilterSession) CheckAndEnableToken(creditAccount common.Address, token common.Address) (*types.Transaction, error) {
	return _ICreditFilter.Contract.CheckAndEnableToken(&_ICreditFilter.TransactOpts, creditAccount, token)
}

// CheckAndEnableToken is a paid mutator transaction binding the contract method 0x51e3f160.
//
// Solidity: function checkAndEnableToken(address creditAccount, address token) returns()
func (_ICreditFilter *ICreditFilterTransactorSession) CheckAndEnableToken(creditAccount common.Address, token common.Address) (*types.Transaction, error) {
	return _ICreditFilter.Contract.CheckAndEnableToken(&_ICreditFilter.TransactOpts, creditAccount, token)
}

// CheckCollateralChange is a paid mutator transaction binding the contract method 0xe1c8ef0d.
//
// Solidity: function checkCollateralChange(address creditAccount, address tokenIn, address tokenOut, uint256 amountIn, uint256 amountOut) returns()
func (_ICreditFilter *ICreditFilterTransactor) CheckCollateralChange(opts *bind.TransactOpts, creditAccount common.Address, tokenIn common.Address, tokenOut common.Address, amountIn *big.Int, amountOut *big.Int) (*types.Transaction, error) {
	return _ICreditFilter.contract.Transact(opts, "checkCollateralChange", creditAccount, tokenIn, tokenOut, amountIn, amountOut)
}

// CheckCollateralChange is a paid mutator transaction binding the contract method 0xe1c8ef0d.
//
// Solidity: function checkCollateralChange(address creditAccount, address tokenIn, address tokenOut, uint256 amountIn, uint256 amountOut) returns()
func (_ICreditFilter *ICreditFilterSession) CheckCollateralChange(creditAccount common.Address, tokenIn common.Address, tokenOut common.Address, amountIn *big.Int, amountOut *big.Int) (*types.Transaction, error) {
	return _ICreditFilter.Contract.CheckCollateralChange(&_ICreditFilter.TransactOpts, creditAccount, tokenIn, tokenOut, amountIn, amountOut)
}

// CheckCollateralChange is a paid mutator transaction binding the contract method 0xe1c8ef0d.
//
// Solidity: function checkCollateralChange(address creditAccount, address tokenIn, address tokenOut, uint256 amountIn, uint256 amountOut) returns()
func (_ICreditFilter *ICreditFilterTransactorSession) CheckCollateralChange(creditAccount common.Address, tokenIn common.Address, tokenOut common.Address, amountIn *big.Int, amountOut *big.Int) (*types.Transaction, error) {
	return _ICreditFilter.Contract.CheckCollateralChange(&_ICreditFilter.TransactOpts, creditAccount, tokenIn, tokenOut, amountIn, amountOut)
}

// CheckMultiTokenCollateral is a paid mutator transaction binding the contract method 0x7e4a6863.
//
// Solidity: function checkMultiTokenCollateral(address creditAccount, uint256[] amountIn, uint256[] amountOut, address[] tokenIn, address[] tokenOut) returns()
func (_ICreditFilter *ICreditFilterTransactor) CheckMultiTokenCollateral(opts *bind.TransactOpts, creditAccount common.Address, amountIn []*big.Int, amountOut []*big.Int, tokenIn []common.Address, tokenOut []common.Address) (*types.Transaction, error) {
	return _ICreditFilter.contract.Transact(opts, "checkMultiTokenCollateral", creditAccount, amountIn, amountOut, tokenIn, tokenOut)
}

// CheckMultiTokenCollateral is a paid mutator transaction binding the contract method 0x7e4a6863.
//
// Solidity: function checkMultiTokenCollateral(address creditAccount, uint256[] amountIn, uint256[] amountOut, address[] tokenIn, address[] tokenOut) returns()
func (_ICreditFilter *ICreditFilterSession) CheckMultiTokenCollateral(creditAccount common.Address, amountIn []*big.Int, amountOut []*big.Int, tokenIn []common.Address, tokenOut []common.Address) (*types.Transaction, error) {
	return _ICreditFilter.Contract.CheckMultiTokenCollateral(&_ICreditFilter.TransactOpts, creditAccount, amountIn, amountOut, tokenIn, tokenOut)
}

// CheckMultiTokenCollateral is a paid mutator transaction binding the contract method 0x7e4a6863.
//
// Solidity: function checkMultiTokenCollateral(address creditAccount, uint256[] amountIn, uint256[] amountOut, address[] tokenIn, address[] tokenOut) returns()
func (_ICreditFilter *ICreditFilterTransactorSession) CheckMultiTokenCollateral(creditAccount common.Address, amountIn []*big.Int, amountOut []*big.Int, tokenIn []common.Address, tokenOut []common.Address) (*types.Transaction, error) {
	return _ICreditFilter.Contract.CheckMultiTokenCollateral(&_ICreditFilter.TransactOpts, creditAccount, amountIn, amountOut, tokenIn, tokenOut)
}

// ConnectCreditManager is a paid mutator transaction binding the contract method 0xcf33d955.
//
// Solidity: function connectCreditManager(address poolService) returns()
func (_ICreditFilter *ICreditFilterTransactor) ConnectCreditManager(opts *bind.TransactOpts, poolService common.Address) (*types.Transaction, error) {
	return _ICreditFilter.contract.Transact(opts, "connectCreditManager", poolService)
}

// ConnectCreditManager is a paid mutator transaction binding the contract method 0xcf33d955.
//
// Solidity: function connectCreditManager(address poolService) returns()
func (_ICreditFilter *ICreditFilterSession) ConnectCreditManager(poolService common.Address) (*types.Transaction, error) {
	return _ICreditFilter.Contract.ConnectCreditManager(&_ICreditFilter.TransactOpts, poolService)
}

// ConnectCreditManager is a paid mutator transaction binding the contract method 0xcf33d955.
//
// Solidity: function connectCreditManager(address poolService) returns()
func (_ICreditFilter *ICreditFilterTransactorSession) ConnectCreditManager(poolService common.Address) (*types.Transaction, error) {
	return _ICreditFilter.Contract.ConnectCreditManager(&_ICreditFilter.TransactOpts, poolService)
}

// ForbidContract is a paid mutator transaction binding the contract method 0x52438e54.
//
// Solidity: function forbidContract(address targetContract) returns()
func (_ICreditFilter *ICreditFilterTransactor) ForbidContract(opts *bind.TransactOpts, targetContract common.Address) (*types.Transaction, error) {
	return _ICreditFilter.contract.Transact(opts, "forbidContract", targetContract)
}

// ForbidContract is a paid mutator transaction binding the contract method 0x52438e54.
//
// Solidity: function forbidContract(address targetContract) returns()
func (_ICreditFilter *ICreditFilterSession) ForbidContract(targetContract common.Address) (*types.Transaction, error) {
	return _ICreditFilter.Contract.ForbidContract(&_ICreditFilter.TransactOpts, targetContract)
}

// ForbidContract is a paid mutator transaction binding the contract method 0x52438e54.
//
// Solidity: function forbidContract(address targetContract) returns()
func (_ICreditFilter *ICreditFilterTransactorSession) ForbidContract(targetContract common.Address) (*types.Transaction, error) {
	return _ICreditFilter.Contract.ForbidContract(&_ICreditFilter.TransactOpts, targetContract)
}

// InitEnabledTokens is a paid mutator transaction binding the contract method 0xe54fe9c8.
//
// Solidity: function initEnabledTokens(address creditAccount) returns()
func (_ICreditFilter *ICreditFilterTransactor) InitEnabledTokens(opts *bind.TransactOpts, creditAccount common.Address) (*types.Transaction, error) {
	return _ICreditFilter.contract.Transact(opts, "initEnabledTokens", creditAccount)
}

// InitEnabledTokens is a paid mutator transaction binding the contract method 0xe54fe9c8.
//
// Solidity: function initEnabledTokens(address creditAccount) returns()
func (_ICreditFilter *ICreditFilterSession) InitEnabledTokens(creditAccount common.Address) (*types.Transaction, error) {
	return _ICreditFilter.Contract.InitEnabledTokens(&_ICreditFilter.TransactOpts, creditAccount)
}

// InitEnabledTokens is a paid mutator transaction binding the contract method 0xe54fe9c8.
//
// Solidity: function initEnabledTokens(address creditAccount) returns()
func (_ICreditFilter *ICreditFilterTransactorSession) InitEnabledTokens(creditAccount common.Address) (*types.Transaction, error) {
	return _ICreditFilter.Contract.InitEnabledTokens(&_ICreditFilter.TransactOpts, creditAccount)
}

// UpdateUnderlyingTokenLiquidationThreshold is a paid mutator transaction binding the contract method 0x40631828.
//
// Solidity: function updateUnderlyingTokenLiquidationThreshold() returns()
func (_ICreditFilter *ICreditFilterTransactor) UpdateUnderlyingTokenLiquidationThreshold(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICreditFilter.contract.Transact(opts, "updateUnderlyingTokenLiquidationThreshold")
}

// UpdateUnderlyingTokenLiquidationThreshold is a paid mutator transaction binding the contract method 0x40631828.
//
// Solidity: function updateUnderlyingTokenLiquidationThreshold() returns()
func (_ICreditFilter *ICreditFilterSession) UpdateUnderlyingTokenLiquidationThreshold() (*types.Transaction, error) {
	return _ICreditFilter.Contract.UpdateUnderlyingTokenLiquidationThreshold(&_ICreditFilter.TransactOpts)
}

// UpdateUnderlyingTokenLiquidationThreshold is a paid mutator transaction binding the contract method 0x40631828.
//
// Solidity: function updateUnderlyingTokenLiquidationThreshold() returns()
func (_ICreditFilter *ICreditFilterTransactorSession) UpdateUnderlyingTokenLiquidationThreshold() (*types.Transaction, error) {
	return _ICreditFilter.Contract.UpdateUnderlyingTokenLiquidationThreshold(&_ICreditFilter.TransactOpts)
}

// ICreditFilterContractAllowedIterator is returned from FilterContractAllowed and is used to iterate over the raw logs and unpacked data for ContractAllowed events raised by the ICreditFilter contract.
type ICreditFilterContractAllowedIterator struct {
	Event *ICreditFilterContractAllowed // Event containing the contract specifics and raw log

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
func (it *ICreditFilterContractAllowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICreditFilterContractAllowed)
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
		it.Event = new(ICreditFilterContractAllowed)
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
func (it *ICreditFilterContractAllowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICreditFilterContractAllowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICreditFilterContractAllowed represents a ContractAllowed event raised by the ICreditFilter contract.
type ICreditFilterContractAllowed struct {
	Protocol common.Address
	Adapter  common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterContractAllowed is a free log retrieval operation binding the contract event 0x4bcbefaef68b99503d502f5a6abe7bca2b183ab8ac55457013c77d084ebd1305.
//
// Solidity: event ContractAllowed(address indexed protocol, address indexed adapter)
func (_ICreditFilter *ICreditFilterFilterer) FilterContractAllowed(opts *bind.FilterOpts, protocol []common.Address, adapter []common.Address) (*ICreditFilterContractAllowedIterator, error) {

	var protocolRule []interface{}
	for _, protocolItem := range protocol {
		protocolRule = append(protocolRule, protocolItem)
	}
	var adapterRule []interface{}
	for _, adapterItem := range adapter {
		adapterRule = append(adapterRule, adapterItem)
	}

	logs, sub, err := _ICreditFilter.contract.FilterLogs(opts, "ContractAllowed", protocolRule, adapterRule)
	if err != nil {
		return nil, err
	}
	return &ICreditFilterContractAllowedIterator{contract: _ICreditFilter.contract, event: "ContractAllowed", logs: logs, sub: sub}, nil
}

// WatchContractAllowed is a free log subscription operation binding the contract event 0x4bcbefaef68b99503d502f5a6abe7bca2b183ab8ac55457013c77d084ebd1305.
//
// Solidity: event ContractAllowed(address indexed protocol, address indexed adapter)
func (_ICreditFilter *ICreditFilterFilterer) WatchContractAllowed(opts *bind.WatchOpts, sink chan<- *ICreditFilterContractAllowed, protocol []common.Address, adapter []common.Address) (event.Subscription, error) {

	var protocolRule []interface{}
	for _, protocolItem := range protocol {
		protocolRule = append(protocolRule, protocolItem)
	}
	var adapterRule []interface{}
	for _, adapterItem := range adapter {
		adapterRule = append(adapterRule, adapterItem)
	}

	logs, sub, err := _ICreditFilter.contract.WatchLogs(opts, "ContractAllowed", protocolRule, adapterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICreditFilterContractAllowed)
				if err := _ICreditFilter.contract.UnpackLog(event, "ContractAllowed", log); err != nil {
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

// ParseContractAllowed is a log parse operation binding the contract event 0x4bcbefaef68b99503d502f5a6abe7bca2b183ab8ac55457013c77d084ebd1305.
//
// Solidity: event ContractAllowed(address indexed protocol, address indexed adapter)
func (_ICreditFilter *ICreditFilterFilterer) ParseContractAllowed(log types.Log) (*ICreditFilterContractAllowed, error) {
	event := new(ICreditFilterContractAllowed)
	if err := _ICreditFilter.contract.UnpackLog(event, "ContractAllowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ICreditFilterContractForbiddenIterator is returned from FilterContractForbidden and is used to iterate over the raw logs and unpacked data for ContractForbidden events raised by the ICreditFilter contract.
type ICreditFilterContractForbiddenIterator struct {
	Event *ICreditFilterContractForbidden // Event containing the contract specifics and raw log

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
func (it *ICreditFilterContractForbiddenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICreditFilterContractForbidden)
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
		it.Event = new(ICreditFilterContractForbidden)
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
func (it *ICreditFilterContractForbiddenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICreditFilterContractForbiddenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICreditFilterContractForbidden represents a ContractForbidden event raised by the ICreditFilter contract.
type ICreditFilterContractForbidden struct {
	Protocol common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterContractForbidden is a free log retrieval operation binding the contract event 0xab9f405bf0c19b97f65a7031634db41569cd2f0e0376a610a1e977f9ab22b58f.
//
// Solidity: event ContractForbidden(address indexed protocol)
func (_ICreditFilter *ICreditFilterFilterer) FilterContractForbidden(opts *bind.FilterOpts, protocol []common.Address) (*ICreditFilterContractForbiddenIterator, error) {

	var protocolRule []interface{}
	for _, protocolItem := range protocol {
		protocolRule = append(protocolRule, protocolItem)
	}

	logs, sub, err := _ICreditFilter.contract.FilterLogs(opts, "ContractForbidden", protocolRule)
	if err != nil {
		return nil, err
	}
	return &ICreditFilterContractForbiddenIterator{contract: _ICreditFilter.contract, event: "ContractForbidden", logs: logs, sub: sub}, nil
}

// WatchContractForbidden is a free log subscription operation binding the contract event 0xab9f405bf0c19b97f65a7031634db41569cd2f0e0376a610a1e977f9ab22b58f.
//
// Solidity: event ContractForbidden(address indexed protocol)
func (_ICreditFilter *ICreditFilterFilterer) WatchContractForbidden(opts *bind.WatchOpts, sink chan<- *ICreditFilterContractForbidden, protocol []common.Address) (event.Subscription, error) {

	var protocolRule []interface{}
	for _, protocolItem := range protocol {
		protocolRule = append(protocolRule, protocolItem)
	}

	logs, sub, err := _ICreditFilter.contract.WatchLogs(opts, "ContractForbidden", protocolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICreditFilterContractForbidden)
				if err := _ICreditFilter.contract.UnpackLog(event, "ContractForbidden", log); err != nil {
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

// ParseContractForbidden is a log parse operation binding the contract event 0xab9f405bf0c19b97f65a7031634db41569cd2f0e0376a610a1e977f9ab22b58f.
//
// Solidity: event ContractForbidden(address indexed protocol)
func (_ICreditFilter *ICreditFilterFilterer) ParseContractForbidden(log types.Log) (*ICreditFilterContractForbidden, error) {
	event := new(ICreditFilterContractForbidden)
	if err := _ICreditFilter.contract.UnpackLog(event, "ContractForbidden", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ICreditFilterNewFastCheckParametersIterator is returned from FilterNewFastCheckParameters and is used to iterate over the raw logs and unpacked data for NewFastCheckParameters events raised by the ICreditFilter contract.
type ICreditFilterNewFastCheckParametersIterator struct {
	Event *ICreditFilterNewFastCheckParameters // Event containing the contract specifics and raw log

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
func (it *ICreditFilterNewFastCheckParametersIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICreditFilterNewFastCheckParameters)
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
		it.Event = new(ICreditFilterNewFastCheckParameters)
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
func (it *ICreditFilterNewFastCheckParametersIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICreditFilterNewFastCheckParametersIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICreditFilterNewFastCheckParameters represents a NewFastCheckParameters event raised by the ICreditFilter contract.
type ICreditFilterNewFastCheckParameters struct {
	ChiThreshold   *big.Int
	FastCheckDelay *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewFastCheckParameters is a free log retrieval operation binding the contract event 0x727652fff0946c19c233fd3eab5fc03db9e9fdd907e902d9136c2a9cac47101c.
//
// Solidity: event NewFastCheckParameters(uint256 chiThreshold, uint256 fastCheckDelay)
func (_ICreditFilter *ICreditFilterFilterer) FilterNewFastCheckParameters(opts *bind.FilterOpts) (*ICreditFilterNewFastCheckParametersIterator, error) {

	logs, sub, err := _ICreditFilter.contract.FilterLogs(opts, "NewFastCheckParameters")
	if err != nil {
		return nil, err
	}
	return &ICreditFilterNewFastCheckParametersIterator{contract: _ICreditFilter.contract, event: "NewFastCheckParameters", logs: logs, sub: sub}, nil
}

// WatchNewFastCheckParameters is a free log subscription operation binding the contract event 0x727652fff0946c19c233fd3eab5fc03db9e9fdd907e902d9136c2a9cac47101c.
//
// Solidity: event NewFastCheckParameters(uint256 chiThreshold, uint256 fastCheckDelay)
func (_ICreditFilter *ICreditFilterFilterer) WatchNewFastCheckParameters(opts *bind.WatchOpts, sink chan<- *ICreditFilterNewFastCheckParameters) (event.Subscription, error) {

	logs, sub, err := _ICreditFilter.contract.WatchLogs(opts, "NewFastCheckParameters")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICreditFilterNewFastCheckParameters)
				if err := _ICreditFilter.contract.UnpackLog(event, "NewFastCheckParameters", log); err != nil {
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

// ParseNewFastCheckParameters is a log parse operation binding the contract event 0x727652fff0946c19c233fd3eab5fc03db9e9fdd907e902d9136c2a9cac47101c.
//
// Solidity: event NewFastCheckParameters(uint256 chiThreshold, uint256 fastCheckDelay)
func (_ICreditFilter *ICreditFilterFilterer) ParseNewFastCheckParameters(log types.Log) (*ICreditFilterNewFastCheckParameters, error) {
	event := new(ICreditFilterNewFastCheckParameters)
	if err := _ICreditFilter.contract.UnpackLog(event, "NewFastCheckParameters", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ICreditFilterTokenAllowedIterator is returned from FilterTokenAllowed and is used to iterate over the raw logs and unpacked data for TokenAllowed events raised by the ICreditFilter contract.
type ICreditFilterTokenAllowedIterator struct {
	Event *ICreditFilterTokenAllowed // Event containing the contract specifics and raw log

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
func (it *ICreditFilterTokenAllowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICreditFilterTokenAllowed)
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
		it.Event = new(ICreditFilterTokenAllowed)
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
func (it *ICreditFilterTokenAllowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICreditFilterTokenAllowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICreditFilterTokenAllowed represents a TokenAllowed event raised by the ICreditFilter contract.
type ICreditFilterTokenAllowed struct {
	Token              common.Address
	LiquidityThreshold *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterTokenAllowed is a free log retrieval operation binding the contract event 0xa52fb6bfa514a4ddcb31de40a5f6c20d767db1f921a8b7747973d93dc5da7a02.
//
// Solidity: event TokenAllowed(address indexed token, uint256 liquidityThreshold)
func (_ICreditFilter *ICreditFilterFilterer) FilterTokenAllowed(opts *bind.FilterOpts, token []common.Address) (*ICreditFilterTokenAllowedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ICreditFilter.contract.FilterLogs(opts, "TokenAllowed", tokenRule)
	if err != nil {
		return nil, err
	}
	return &ICreditFilterTokenAllowedIterator{contract: _ICreditFilter.contract, event: "TokenAllowed", logs: logs, sub: sub}, nil
}

// WatchTokenAllowed is a free log subscription operation binding the contract event 0xa52fb6bfa514a4ddcb31de40a5f6c20d767db1f921a8b7747973d93dc5da7a02.
//
// Solidity: event TokenAllowed(address indexed token, uint256 liquidityThreshold)
func (_ICreditFilter *ICreditFilterFilterer) WatchTokenAllowed(opts *bind.WatchOpts, sink chan<- *ICreditFilterTokenAllowed, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ICreditFilter.contract.WatchLogs(opts, "TokenAllowed", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICreditFilterTokenAllowed)
				if err := _ICreditFilter.contract.UnpackLog(event, "TokenAllowed", log); err != nil {
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

// ParseTokenAllowed is a log parse operation binding the contract event 0xa52fb6bfa514a4ddcb31de40a5f6c20d767db1f921a8b7747973d93dc5da7a02.
//
// Solidity: event TokenAllowed(address indexed token, uint256 liquidityThreshold)
func (_ICreditFilter *ICreditFilterFilterer) ParseTokenAllowed(log types.Log) (*ICreditFilterTokenAllowed, error) {
	event := new(ICreditFilterTokenAllowed)
	if err := _ICreditFilter.contract.UnpackLog(event, "TokenAllowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ICreditFilterTransferAccountAllowedIterator is returned from FilterTransferAccountAllowed and is used to iterate over the raw logs and unpacked data for TransferAccountAllowed events raised by the ICreditFilter contract.
type ICreditFilterTransferAccountAllowedIterator struct {
	Event *ICreditFilterTransferAccountAllowed // Event containing the contract specifics and raw log

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
func (it *ICreditFilterTransferAccountAllowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICreditFilterTransferAccountAllowed)
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
		it.Event = new(ICreditFilterTransferAccountAllowed)
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
func (it *ICreditFilterTransferAccountAllowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICreditFilterTransferAccountAllowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICreditFilterTransferAccountAllowed represents a TransferAccountAllowed event raised by the ICreditFilter contract.
type ICreditFilterTransferAccountAllowed struct {
	From  common.Address
	To    common.Address
	State bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransferAccountAllowed is a free log retrieval operation binding the contract event 0x9b3258bc4904fd6426b99843e206c6c7cdb1fd0f040121c25b71dafbb3851ee0.
//
// Solidity: event TransferAccountAllowed(address indexed from, address indexed to, bool state)
func (_ICreditFilter *ICreditFilterFilterer) FilterTransferAccountAllowed(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ICreditFilterTransferAccountAllowedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ICreditFilter.contract.FilterLogs(opts, "TransferAccountAllowed", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ICreditFilterTransferAccountAllowedIterator{contract: _ICreditFilter.contract, event: "TransferAccountAllowed", logs: logs, sub: sub}, nil
}

// WatchTransferAccountAllowed is a free log subscription operation binding the contract event 0x9b3258bc4904fd6426b99843e206c6c7cdb1fd0f040121c25b71dafbb3851ee0.
//
// Solidity: event TransferAccountAllowed(address indexed from, address indexed to, bool state)
func (_ICreditFilter *ICreditFilterFilterer) WatchTransferAccountAllowed(opts *bind.WatchOpts, sink chan<- *ICreditFilterTransferAccountAllowed, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ICreditFilter.contract.WatchLogs(opts, "TransferAccountAllowed", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICreditFilterTransferAccountAllowed)
				if err := _ICreditFilter.contract.UnpackLog(event, "TransferAccountAllowed", log); err != nil {
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

// ParseTransferAccountAllowed is a log parse operation binding the contract event 0x9b3258bc4904fd6426b99843e206c6c7cdb1fd0f040121c25b71dafbb3851ee0.
//
// Solidity: event TransferAccountAllowed(address indexed from, address indexed to, bool state)
func (_ICreditFilter *ICreditFilterFilterer) ParseTransferAccountAllowed(log types.Log) (*ICreditFilterTransferAccountAllowed, error) {
	event := new(ICreditFilterTransferAccountAllowed)
	if err := _ICreditFilter.contract.UnpackLog(event, "TransferAccountAllowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ICreditFilterTransferPluginAllowedIterator is returned from FilterTransferPluginAllowed and is used to iterate over the raw logs and unpacked data for TransferPluginAllowed events raised by the ICreditFilter contract.
type ICreditFilterTransferPluginAllowedIterator struct {
	Event *ICreditFilterTransferPluginAllowed // Event containing the contract specifics and raw log

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
func (it *ICreditFilterTransferPluginAllowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICreditFilterTransferPluginAllowed)
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
		it.Event = new(ICreditFilterTransferPluginAllowed)
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
func (it *ICreditFilterTransferPluginAllowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICreditFilterTransferPluginAllowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICreditFilterTransferPluginAllowed represents a TransferPluginAllowed event raised by the ICreditFilter contract.
type ICreditFilterTransferPluginAllowed struct {
	Pugin common.Address
	State bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransferPluginAllowed is a free log retrieval operation binding the contract event 0xc7d2592986c53f858769b011e8ce6298936f8609789988e9f5ad4f0a20798897.
//
// Solidity: event TransferPluginAllowed(address indexed pugin, bool state)
func (_ICreditFilter *ICreditFilterFilterer) FilterTransferPluginAllowed(opts *bind.FilterOpts, pugin []common.Address) (*ICreditFilterTransferPluginAllowedIterator, error) {

	var puginRule []interface{}
	for _, puginItem := range pugin {
		puginRule = append(puginRule, puginItem)
	}

	logs, sub, err := _ICreditFilter.contract.FilterLogs(opts, "TransferPluginAllowed", puginRule)
	if err != nil {
		return nil, err
	}
	return &ICreditFilterTransferPluginAllowedIterator{contract: _ICreditFilter.contract, event: "TransferPluginAllowed", logs: logs, sub: sub}, nil
}

// WatchTransferPluginAllowed is a free log subscription operation binding the contract event 0xc7d2592986c53f858769b011e8ce6298936f8609789988e9f5ad4f0a20798897.
//
// Solidity: event TransferPluginAllowed(address indexed pugin, bool state)
func (_ICreditFilter *ICreditFilterFilterer) WatchTransferPluginAllowed(opts *bind.WatchOpts, sink chan<- *ICreditFilterTransferPluginAllowed, pugin []common.Address) (event.Subscription, error) {

	var puginRule []interface{}
	for _, puginItem := range pugin {
		puginRule = append(puginRule, puginItem)
	}

	logs, sub, err := _ICreditFilter.contract.WatchLogs(opts, "TransferPluginAllowed", puginRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICreditFilterTransferPluginAllowed)
				if err := _ICreditFilter.contract.UnpackLog(event, "TransferPluginAllowed", log); err != nil {
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

// ParseTransferPluginAllowed is a log parse operation binding the contract event 0xc7d2592986c53f858769b011e8ce6298936f8609789988e9f5ad4f0a20798897.
//
// Solidity: event TransferPluginAllowed(address indexed pugin, bool state)
func (_ICreditFilter *ICreditFilterFilterer) ParseTransferPluginAllowed(log types.Log) (*ICreditFilterTransferPluginAllowed, error) {
	event := new(ICreditFilterTransferPluginAllowed)
	if err := _ICreditFilter.contract.UnpackLog(event, "TransferPluginAllowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
