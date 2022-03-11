// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package creditConfigurator

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

// CreditConfiguratorABI is the input ABI used to generate the binding from.
const CreditConfiguratorABI = "[{\"inputs\":[{\"internalType\":\"contractCreditManager\",\"name\":\"_creditManager\",\"type\":\"address\"},{\"internalType\":\"contractCreditFacade\",\"name\":\"_creditFacade\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"minBorrowedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxBorrowedAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidationThreshold\",\"type\":\"uint256\"}],\"internalType\":\"structAllowedToken[]\",\"name\":\"allowedTokens\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"targetContract\",\"type\":\"address\"},{\"internalType\":\"enumAdapterType\",\"name\":\"adapterType\",\"type\":\"uint8\"}],\"internalType\":\"structAdapterConfig[]\",\"name\":\"adapters\",\"type\":\"tuple[]\"}],\"internalType\":\"structCreditManagerOpts\",\"name\":\"opts\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"protocol\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"adapter\",\"type\":\"address\"}],\"name\":\"ContractAllowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"protocol\",\"type\":\"address\"}],\"name\":\"ContractForbidden\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newCreditConfigurator\",\"type\":\"address\"}],\"name\":\"CreditConfiguratorUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newCreditFacade\",\"type\":\"address\"}],\"name\":\"CreditFacadeUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chiThreshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fastCheckDelay\",\"type\":\"uint256\"}],\"name\":\"FastCheckParametersUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeInterest\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeLiquidation\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidationPremium\",\"type\":\"uint256\"}],\"name\":\"FeesUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minBorrowedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxBorrowedAmount\",\"type\":\"uint256\"}],\"name\":\"LimitsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newPriceOracle\",\"type\":\"address\"}],\"name\":\"PriceOracleUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenAllowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenForbidden\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidityThreshold\",\"type\":\"uint256\"}],\"name\":\"TokenLiquidationThresholdUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"addTokenToAllowedList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addressProvider\",\"outputs\":[{\"internalType\":\"contractAddressProvider\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"targetContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"adapter\",\"type\":\"address\"}],\"name\":\"allowContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"allowToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"allowedContracts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allowedContractsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"percentage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"times\",\"type\":\"uint256\"}],\"name\":\"calcMaxPossibleDrop\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creditFacade\",\"outputs\":[{\"internalType\":\"contractCreditFacade\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creditManager\",\"outputs\":[{\"internalType\":\"contractCreditManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"targetContract\",\"type\":\"address\"}],\"name\":\"forbidContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"forbidToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chiThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_hfCheckInterval\",\"type\":\"uint256\"}],\"name\":\"setFastCheckParameters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_feeInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_feeLiquidation\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_liquidationPremium\",\"type\":\"uint256\"}],\"name\":\"setFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minBorrowedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxBorrowedAmount\",\"type\":\"uint256\"}],\"name\":\"setLimits\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidationThreshold\",\"type\":\"uint256\"}],\"name\":\"setLiquidationThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlyingToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_creditConfigurator\",\"type\":\"address\"}],\"name\":\"upgradeConfigurator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_creditFacade\",\"type\":\"address\"}],\"name\":\"upgradeCreditFacade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"upgradePriceOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// CreditConfigurator is an auto generated Go binding around an Ethereum contract.
type CreditConfigurator struct {
	CreditConfiguratorCaller     // Read-only binding to the contract
	CreditConfiguratorTransactor // Write-only binding to the contract
	CreditConfiguratorFilterer   // Log filterer for contract events
}

// CreditConfiguratorCaller is an auto generated read-only Go binding around an Ethereum contract.
type CreditConfiguratorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditConfiguratorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CreditConfiguratorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditConfiguratorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CreditConfiguratorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditConfiguratorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CreditConfiguratorSession struct {
	Contract     *CreditConfigurator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// CreditConfiguratorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CreditConfiguratorCallerSession struct {
	Contract *CreditConfiguratorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// CreditConfiguratorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CreditConfiguratorTransactorSession struct {
	Contract     *CreditConfiguratorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// CreditConfiguratorRaw is an auto generated low-level Go binding around an Ethereum contract.
type CreditConfiguratorRaw struct {
	Contract *CreditConfigurator // Generic contract binding to access the raw methods on
}

// CreditConfiguratorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CreditConfiguratorCallerRaw struct {
	Contract *CreditConfiguratorCaller // Generic read-only contract binding to access the raw methods on
}

// CreditConfiguratorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CreditConfiguratorTransactorRaw struct {
	Contract *CreditConfiguratorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCreditConfigurator creates a new instance of CreditConfigurator, bound to a specific deployed contract.
func NewCreditConfigurator(address common.Address, backend bind.ContractBackend) (*CreditConfigurator, error) {
	contract, err := bindCreditConfigurator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CreditConfigurator{CreditConfiguratorCaller: CreditConfiguratorCaller{contract: contract}, CreditConfiguratorTransactor: CreditConfiguratorTransactor{contract: contract}, CreditConfiguratorFilterer: CreditConfiguratorFilterer{contract: contract}}, nil
}

// NewCreditConfiguratorCaller creates a new read-only instance of CreditConfigurator, bound to a specific deployed contract.
func NewCreditConfiguratorCaller(address common.Address, caller bind.ContractCaller) (*CreditConfiguratorCaller, error) {
	contract, err := bindCreditConfigurator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorCaller{contract: contract}, nil
}

// NewCreditConfiguratorTransactor creates a new write-only instance of CreditConfigurator, bound to a specific deployed contract.
func NewCreditConfiguratorTransactor(address common.Address, transactor bind.ContractTransactor) (*CreditConfiguratorTransactor, error) {
	contract, err := bindCreditConfigurator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorTransactor{contract: contract}, nil
}

// NewCreditConfiguratorFilterer creates a new log filterer instance of CreditConfigurator, bound to a specific deployed contract.
func NewCreditConfiguratorFilterer(address common.Address, filterer bind.ContractFilterer) (*CreditConfiguratorFilterer, error) {
	contract, err := bindCreditConfigurator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorFilterer{contract: contract}, nil
}

// bindCreditConfigurator binds a generic wrapper to an already deployed contract.
func bindCreditConfigurator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CreditConfiguratorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CreditConfigurator *CreditConfiguratorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CreditConfigurator.Contract.CreditConfiguratorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CreditConfigurator *CreditConfiguratorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditConfigurator.Contract.CreditConfiguratorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CreditConfigurator *CreditConfiguratorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CreditConfigurator.Contract.CreditConfiguratorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CreditConfigurator *CreditConfiguratorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CreditConfigurator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CreditConfigurator *CreditConfiguratorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditConfigurator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CreditConfigurator *CreditConfiguratorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CreditConfigurator.Contract.contract.Transact(opts, method, params...)
}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_CreditConfigurator *CreditConfiguratorCaller) AddressProvider(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditConfigurator.contract.Call(opts, &out, "addressProvider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_CreditConfigurator *CreditConfiguratorSession) AddressProvider() (common.Address, error) {
	return _CreditConfigurator.Contract.AddressProvider(&_CreditConfigurator.CallOpts)
}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_CreditConfigurator *CreditConfiguratorCallerSession) AddressProvider() (common.Address, error) {
	return _CreditConfigurator.Contract.AddressProvider(&_CreditConfigurator.CallOpts)
}

// AllowedContracts is a free data retrieval call binding the contract method 0x5094cb4f.
//
// Solidity: function allowedContracts(uint256 i) view returns(address)
func (_CreditConfigurator *CreditConfiguratorCaller) AllowedContracts(opts *bind.CallOpts, i *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CreditConfigurator.contract.Call(opts, &out, "allowedContracts", i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllowedContracts is a free data retrieval call binding the contract method 0x5094cb4f.
//
// Solidity: function allowedContracts(uint256 i) view returns(address)
func (_CreditConfigurator *CreditConfiguratorSession) AllowedContracts(i *big.Int) (common.Address, error) {
	return _CreditConfigurator.Contract.AllowedContracts(&_CreditConfigurator.CallOpts, i)
}

// AllowedContracts is a free data retrieval call binding the contract method 0x5094cb4f.
//
// Solidity: function allowedContracts(uint256 i) view returns(address)
func (_CreditConfigurator *CreditConfiguratorCallerSession) AllowedContracts(i *big.Int) (common.Address, error) {
	return _CreditConfigurator.Contract.AllowedContracts(&_CreditConfigurator.CallOpts, i)
}

// AllowedContractsCount is a free data retrieval call binding the contract method 0x50e036ff.
//
// Solidity: function allowedContractsCount() view returns(uint256)
func (_CreditConfigurator *CreditConfiguratorCaller) AllowedContractsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CreditConfigurator.contract.Call(opts, &out, "allowedContractsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllowedContractsCount is a free data retrieval call binding the contract method 0x50e036ff.
//
// Solidity: function allowedContractsCount() view returns(uint256)
func (_CreditConfigurator *CreditConfiguratorSession) AllowedContractsCount() (*big.Int, error) {
	return _CreditConfigurator.Contract.AllowedContractsCount(&_CreditConfigurator.CallOpts)
}

// AllowedContractsCount is a free data retrieval call binding the contract method 0x50e036ff.
//
// Solidity: function allowedContractsCount() view returns(uint256)
func (_CreditConfigurator *CreditConfiguratorCallerSession) AllowedContractsCount() (*big.Int, error) {
	return _CreditConfigurator.Contract.AllowedContractsCount(&_CreditConfigurator.CallOpts)
}

// CalcMaxPossibleDrop is a free data retrieval call binding the contract method 0xb3c61943.
//
// Solidity: function calcMaxPossibleDrop(uint256 percentage, uint256 times) pure returns(uint256 value)
func (_CreditConfigurator *CreditConfiguratorCaller) CalcMaxPossibleDrop(opts *bind.CallOpts, percentage *big.Int, times *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CreditConfigurator.contract.Call(opts, &out, "calcMaxPossibleDrop", percentage, times)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcMaxPossibleDrop is a free data retrieval call binding the contract method 0xb3c61943.
//
// Solidity: function calcMaxPossibleDrop(uint256 percentage, uint256 times) pure returns(uint256 value)
func (_CreditConfigurator *CreditConfiguratorSession) CalcMaxPossibleDrop(percentage *big.Int, times *big.Int) (*big.Int, error) {
	return _CreditConfigurator.Contract.CalcMaxPossibleDrop(&_CreditConfigurator.CallOpts, percentage, times)
}

// CalcMaxPossibleDrop is a free data retrieval call binding the contract method 0xb3c61943.
//
// Solidity: function calcMaxPossibleDrop(uint256 percentage, uint256 times) pure returns(uint256 value)
func (_CreditConfigurator *CreditConfiguratorCallerSession) CalcMaxPossibleDrop(percentage *big.Int, times *big.Int) (*big.Int, error) {
	return _CreditConfigurator.Contract.CalcMaxPossibleDrop(&_CreditConfigurator.CallOpts, percentage, times)
}

// CreditFacade is a free data retrieval call binding the contract method 0x2f7a1881.
//
// Solidity: function creditFacade() view returns(address)
func (_CreditConfigurator *CreditConfiguratorCaller) CreditFacade(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditConfigurator.contract.Call(opts, &out, "creditFacade")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditFacade is a free data retrieval call binding the contract method 0x2f7a1881.
//
// Solidity: function creditFacade() view returns(address)
func (_CreditConfigurator *CreditConfiguratorSession) CreditFacade() (common.Address, error) {
	return _CreditConfigurator.Contract.CreditFacade(&_CreditConfigurator.CallOpts)
}

// CreditFacade is a free data retrieval call binding the contract method 0x2f7a1881.
//
// Solidity: function creditFacade() view returns(address)
func (_CreditConfigurator *CreditConfiguratorCallerSession) CreditFacade() (common.Address, error) {
	return _CreditConfigurator.Contract.CreditFacade(&_CreditConfigurator.CallOpts)
}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_CreditConfigurator *CreditConfiguratorCaller) CreditManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditConfigurator.contract.Call(opts, &out, "creditManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_CreditConfigurator *CreditConfiguratorSession) CreditManager() (common.Address, error) {
	return _CreditConfigurator.Contract.CreditManager(&_CreditConfigurator.CallOpts)
}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_CreditConfigurator *CreditConfiguratorCallerSession) CreditManager() (common.Address, error) {
	return _CreditConfigurator.Contract.CreditManager(&_CreditConfigurator.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CreditConfigurator *CreditConfiguratorCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CreditConfigurator.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CreditConfigurator *CreditConfiguratorSession) Paused() (bool, error) {
	return _CreditConfigurator.Contract.Paused(&_CreditConfigurator.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CreditConfigurator *CreditConfiguratorCallerSession) Paused() (bool, error) {
	return _CreditConfigurator.Contract.Paused(&_CreditConfigurator.CallOpts)
}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_CreditConfigurator *CreditConfiguratorCaller) UnderlyingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditConfigurator.contract.Call(opts, &out, "underlyingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_CreditConfigurator *CreditConfiguratorSession) UnderlyingToken() (common.Address, error) {
	return _CreditConfigurator.Contract.UnderlyingToken(&_CreditConfigurator.CallOpts)
}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_CreditConfigurator *CreditConfiguratorCallerSession) UnderlyingToken() (common.Address, error) {
	return _CreditConfigurator.Contract.UnderlyingToken(&_CreditConfigurator.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_CreditConfigurator *CreditConfiguratorCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CreditConfigurator.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_CreditConfigurator *CreditConfiguratorSession) Version() (*big.Int, error) {
	return _CreditConfigurator.Contract.Version(&_CreditConfigurator.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_CreditConfigurator *CreditConfiguratorCallerSession) Version() (*big.Int, error) {
	return _CreditConfigurator.Contract.Version(&_CreditConfigurator.CallOpts)
}

// AddTokenToAllowedList is a paid mutator transaction binding the contract method 0xdadfb98b.
//
// Solidity: function addTokenToAllowedList(address token) returns()
func (_CreditConfigurator *CreditConfiguratorTransactor) AddTokenToAllowedList(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _CreditConfigurator.contract.Transact(opts, "addTokenToAllowedList", token)
}

// AddTokenToAllowedList is a paid mutator transaction binding the contract method 0xdadfb98b.
//
// Solidity: function addTokenToAllowedList(address token) returns()
func (_CreditConfigurator *CreditConfiguratorSession) AddTokenToAllowedList(token common.Address) (*types.Transaction, error) {
	return _CreditConfigurator.Contract.AddTokenToAllowedList(&_CreditConfigurator.TransactOpts, token)
}

// AddTokenToAllowedList is a paid mutator transaction binding the contract method 0xdadfb98b.
//
// Solidity: function addTokenToAllowedList(address token) returns()
func (_CreditConfigurator *CreditConfiguratorTransactorSession) AddTokenToAllowedList(token common.Address) (*types.Transaction, error) {
	return _CreditConfigurator.Contract.AddTokenToAllowedList(&_CreditConfigurator.TransactOpts, token)
}

// AllowContract is a paid mutator transaction binding the contract method 0x7bccacee.
//
// Solidity: function allowContract(address targetContract, address adapter) returns()
func (_CreditConfigurator *CreditConfiguratorTransactor) AllowContract(opts *bind.TransactOpts, targetContract common.Address, adapter common.Address) (*types.Transaction, error) {
	return _CreditConfigurator.contract.Transact(opts, "allowContract", targetContract, adapter)
}

// AllowContract is a paid mutator transaction binding the contract method 0x7bccacee.
//
// Solidity: function allowContract(address targetContract, address adapter) returns()
func (_CreditConfigurator *CreditConfiguratorSession) AllowContract(targetContract common.Address, adapter common.Address) (*types.Transaction, error) {
	return _CreditConfigurator.Contract.AllowContract(&_CreditConfigurator.TransactOpts, targetContract, adapter)
}

// AllowContract is a paid mutator transaction binding the contract method 0x7bccacee.
//
// Solidity: function allowContract(address targetContract, address adapter) returns()
func (_CreditConfigurator *CreditConfiguratorTransactorSession) AllowContract(targetContract common.Address, adapter common.Address) (*types.Transaction, error) {
	return _CreditConfigurator.Contract.AllowContract(&_CreditConfigurator.TransactOpts, targetContract, adapter)
}

// AllowToken is a paid mutator transaction binding the contract method 0xb53472ef.
//
// Solidity: function allowToken(address token) returns()
func (_CreditConfigurator *CreditConfiguratorTransactor) AllowToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _CreditConfigurator.contract.Transact(opts, "allowToken", token)
}

// AllowToken is a paid mutator transaction binding the contract method 0xb53472ef.
//
// Solidity: function allowToken(address token) returns()
func (_CreditConfigurator *CreditConfiguratorSession) AllowToken(token common.Address) (*types.Transaction, error) {
	return _CreditConfigurator.Contract.AllowToken(&_CreditConfigurator.TransactOpts, token)
}

// AllowToken is a paid mutator transaction binding the contract method 0xb53472ef.
//
// Solidity: function allowToken(address token) returns()
func (_CreditConfigurator *CreditConfiguratorTransactorSession) AllowToken(token common.Address) (*types.Transaction, error) {
	return _CreditConfigurator.Contract.AllowToken(&_CreditConfigurator.TransactOpts, token)
}

// ForbidContract is a paid mutator transaction binding the contract method 0x52438e54.
//
// Solidity: function forbidContract(address targetContract) returns()
func (_CreditConfigurator *CreditConfiguratorTransactor) ForbidContract(opts *bind.TransactOpts, targetContract common.Address) (*types.Transaction, error) {
	return _CreditConfigurator.contract.Transact(opts, "forbidContract", targetContract)
}

// ForbidContract is a paid mutator transaction binding the contract method 0x52438e54.
//
// Solidity: function forbidContract(address targetContract) returns()
func (_CreditConfigurator *CreditConfiguratorSession) ForbidContract(targetContract common.Address) (*types.Transaction, error) {
	return _CreditConfigurator.Contract.ForbidContract(&_CreditConfigurator.TransactOpts, targetContract)
}

// ForbidContract is a paid mutator transaction binding the contract method 0x52438e54.
//
// Solidity: function forbidContract(address targetContract) returns()
func (_CreditConfigurator *CreditConfiguratorTransactorSession) ForbidContract(targetContract common.Address) (*types.Transaction, error) {
	return _CreditConfigurator.Contract.ForbidContract(&_CreditConfigurator.TransactOpts, targetContract)
}

// ForbidToken is a paid mutator transaction binding the contract method 0x24147708.
//
// Solidity: function forbidToken(address token) returns()
func (_CreditConfigurator *CreditConfiguratorTransactor) ForbidToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _CreditConfigurator.contract.Transact(opts, "forbidToken", token)
}

// ForbidToken is a paid mutator transaction binding the contract method 0x24147708.
//
// Solidity: function forbidToken(address token) returns()
func (_CreditConfigurator *CreditConfiguratorSession) ForbidToken(token common.Address) (*types.Transaction, error) {
	return _CreditConfigurator.Contract.ForbidToken(&_CreditConfigurator.TransactOpts, token)
}

// ForbidToken is a paid mutator transaction binding the contract method 0x24147708.
//
// Solidity: function forbidToken(address token) returns()
func (_CreditConfigurator *CreditConfiguratorTransactorSession) ForbidToken(token common.Address) (*types.Transaction, error) {
	return _CreditConfigurator.Contract.ForbidToken(&_CreditConfigurator.TransactOpts, token)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CreditConfigurator *CreditConfiguratorTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditConfigurator.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CreditConfigurator *CreditConfiguratorSession) Pause() (*types.Transaction, error) {
	return _CreditConfigurator.Contract.Pause(&_CreditConfigurator.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CreditConfigurator *CreditConfiguratorTransactorSession) Pause() (*types.Transaction, error) {
	return _CreditConfigurator.Contract.Pause(&_CreditConfigurator.TransactOpts)
}

// SetFastCheckParameters is a paid mutator transaction binding the contract method 0x62061c6d.
//
// Solidity: function setFastCheckParameters(uint256 _chiThreshold, uint256 _hfCheckInterval) returns()
func (_CreditConfigurator *CreditConfiguratorTransactor) SetFastCheckParameters(opts *bind.TransactOpts, _chiThreshold *big.Int, _hfCheckInterval *big.Int) (*types.Transaction, error) {
	return _CreditConfigurator.contract.Transact(opts, "setFastCheckParameters", _chiThreshold, _hfCheckInterval)
}

// SetFastCheckParameters is a paid mutator transaction binding the contract method 0x62061c6d.
//
// Solidity: function setFastCheckParameters(uint256 _chiThreshold, uint256 _hfCheckInterval) returns()
func (_CreditConfigurator *CreditConfiguratorSession) SetFastCheckParameters(_chiThreshold *big.Int, _hfCheckInterval *big.Int) (*types.Transaction, error) {
	return _CreditConfigurator.Contract.SetFastCheckParameters(&_CreditConfigurator.TransactOpts, _chiThreshold, _hfCheckInterval)
}

// SetFastCheckParameters is a paid mutator transaction binding the contract method 0x62061c6d.
//
// Solidity: function setFastCheckParameters(uint256 _chiThreshold, uint256 _hfCheckInterval) returns()
func (_CreditConfigurator *CreditConfiguratorTransactorSession) SetFastCheckParameters(_chiThreshold *big.Int, _hfCheckInterval *big.Int) (*types.Transaction, error) {
	return _CreditConfigurator.Contract.SetFastCheckParameters(&_CreditConfigurator.TransactOpts, _chiThreshold, _hfCheckInterval)
}

// SetFees is a paid mutator transaction binding the contract method 0xcec10c11.
//
// Solidity: function setFees(uint256 _feeInterest, uint256 _feeLiquidation, uint256 _liquidationPremium) returns()
func (_CreditConfigurator *CreditConfiguratorTransactor) SetFees(opts *bind.TransactOpts, _feeInterest *big.Int, _feeLiquidation *big.Int, _liquidationPremium *big.Int) (*types.Transaction, error) {
	return _CreditConfigurator.contract.Transact(opts, "setFees", _feeInterest, _feeLiquidation, _liquidationPremium)
}

// SetFees is a paid mutator transaction binding the contract method 0xcec10c11.
//
// Solidity: function setFees(uint256 _feeInterest, uint256 _feeLiquidation, uint256 _liquidationPremium) returns()
func (_CreditConfigurator *CreditConfiguratorSession) SetFees(_feeInterest *big.Int, _feeLiquidation *big.Int, _liquidationPremium *big.Int) (*types.Transaction, error) {
	return _CreditConfigurator.Contract.SetFees(&_CreditConfigurator.TransactOpts, _feeInterest, _feeLiquidation, _liquidationPremium)
}

// SetFees is a paid mutator transaction binding the contract method 0xcec10c11.
//
// Solidity: function setFees(uint256 _feeInterest, uint256 _feeLiquidation, uint256 _liquidationPremium) returns()
func (_CreditConfigurator *CreditConfiguratorTransactorSession) SetFees(_feeInterest *big.Int, _feeLiquidation *big.Int, _liquidationPremium *big.Int) (*types.Transaction, error) {
	return _CreditConfigurator.Contract.SetFees(&_CreditConfigurator.TransactOpts, _feeInterest, _feeLiquidation, _liquidationPremium)
}

// SetLimits is a paid mutator transaction binding the contract method 0xc4590d3f.
//
// Solidity: function setLimits(uint256 _minBorrowedAmount, uint256 _maxBorrowedAmount) returns()
func (_CreditConfigurator *CreditConfiguratorTransactor) SetLimits(opts *bind.TransactOpts, _minBorrowedAmount *big.Int, _maxBorrowedAmount *big.Int) (*types.Transaction, error) {
	return _CreditConfigurator.contract.Transact(opts, "setLimits", _minBorrowedAmount, _maxBorrowedAmount)
}

// SetLimits is a paid mutator transaction binding the contract method 0xc4590d3f.
//
// Solidity: function setLimits(uint256 _minBorrowedAmount, uint256 _maxBorrowedAmount) returns()
func (_CreditConfigurator *CreditConfiguratorSession) SetLimits(_minBorrowedAmount *big.Int, _maxBorrowedAmount *big.Int) (*types.Transaction, error) {
	return _CreditConfigurator.Contract.SetLimits(&_CreditConfigurator.TransactOpts, _minBorrowedAmount, _maxBorrowedAmount)
}

// SetLimits is a paid mutator transaction binding the contract method 0xc4590d3f.
//
// Solidity: function setLimits(uint256 _minBorrowedAmount, uint256 _maxBorrowedAmount) returns()
func (_CreditConfigurator *CreditConfiguratorTransactorSession) SetLimits(_minBorrowedAmount *big.Int, _maxBorrowedAmount *big.Int) (*types.Transaction, error) {
	return _CreditConfigurator.Contract.SetLimits(&_CreditConfigurator.TransactOpts, _minBorrowedAmount, _maxBorrowedAmount)
}

// SetLiquidationThreshold is a paid mutator transaction binding the contract method 0x0e30428d.
//
// Solidity: function setLiquidationThreshold(address token, uint256 liquidationThreshold) returns()
func (_CreditConfigurator *CreditConfiguratorTransactor) SetLiquidationThreshold(opts *bind.TransactOpts, token common.Address, liquidationThreshold *big.Int) (*types.Transaction, error) {
	return _CreditConfigurator.contract.Transact(opts, "setLiquidationThreshold", token, liquidationThreshold)
}

// SetLiquidationThreshold is a paid mutator transaction binding the contract method 0x0e30428d.
//
// Solidity: function setLiquidationThreshold(address token, uint256 liquidationThreshold) returns()
func (_CreditConfigurator *CreditConfiguratorSession) SetLiquidationThreshold(token common.Address, liquidationThreshold *big.Int) (*types.Transaction, error) {
	return _CreditConfigurator.Contract.SetLiquidationThreshold(&_CreditConfigurator.TransactOpts, token, liquidationThreshold)
}

// SetLiquidationThreshold is a paid mutator transaction binding the contract method 0x0e30428d.
//
// Solidity: function setLiquidationThreshold(address token, uint256 liquidationThreshold) returns()
func (_CreditConfigurator *CreditConfiguratorTransactorSession) SetLiquidationThreshold(token common.Address, liquidationThreshold *big.Int) (*types.Transaction, error) {
	return _CreditConfigurator.Contract.SetLiquidationThreshold(&_CreditConfigurator.TransactOpts, token, liquidationThreshold)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CreditConfigurator *CreditConfiguratorTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditConfigurator.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CreditConfigurator *CreditConfiguratorSession) Unpause() (*types.Transaction, error) {
	return _CreditConfigurator.Contract.Unpause(&_CreditConfigurator.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CreditConfigurator *CreditConfiguratorTransactorSession) Unpause() (*types.Transaction, error) {
	return _CreditConfigurator.Contract.Unpause(&_CreditConfigurator.TransactOpts)
}

// UpgradeConfigurator is a paid mutator transaction binding the contract method 0xbbbae6ab.
//
// Solidity: function upgradeConfigurator(address _creditConfigurator) returns()
func (_CreditConfigurator *CreditConfiguratorTransactor) UpgradeConfigurator(opts *bind.TransactOpts, _creditConfigurator common.Address) (*types.Transaction, error) {
	return _CreditConfigurator.contract.Transact(opts, "upgradeConfigurator", _creditConfigurator)
}

// UpgradeConfigurator is a paid mutator transaction binding the contract method 0xbbbae6ab.
//
// Solidity: function upgradeConfigurator(address _creditConfigurator) returns()
func (_CreditConfigurator *CreditConfiguratorSession) UpgradeConfigurator(_creditConfigurator common.Address) (*types.Transaction, error) {
	return _CreditConfigurator.Contract.UpgradeConfigurator(&_CreditConfigurator.TransactOpts, _creditConfigurator)
}

// UpgradeConfigurator is a paid mutator transaction binding the contract method 0xbbbae6ab.
//
// Solidity: function upgradeConfigurator(address _creditConfigurator) returns()
func (_CreditConfigurator *CreditConfiguratorTransactorSession) UpgradeConfigurator(_creditConfigurator common.Address) (*types.Transaction, error) {
	return _CreditConfigurator.Contract.UpgradeConfigurator(&_CreditConfigurator.TransactOpts, _creditConfigurator)
}

// UpgradeCreditFacade is a paid mutator transaction binding the contract method 0x693ce7f5.
//
// Solidity: function upgradeCreditFacade(address _creditFacade) returns()
func (_CreditConfigurator *CreditConfiguratorTransactor) UpgradeCreditFacade(opts *bind.TransactOpts, _creditFacade common.Address) (*types.Transaction, error) {
	return _CreditConfigurator.contract.Transact(opts, "upgradeCreditFacade", _creditFacade)
}

// UpgradeCreditFacade is a paid mutator transaction binding the contract method 0x693ce7f5.
//
// Solidity: function upgradeCreditFacade(address _creditFacade) returns()
func (_CreditConfigurator *CreditConfiguratorSession) UpgradeCreditFacade(_creditFacade common.Address) (*types.Transaction, error) {
	return _CreditConfigurator.Contract.UpgradeCreditFacade(&_CreditConfigurator.TransactOpts, _creditFacade)
}

// UpgradeCreditFacade is a paid mutator transaction binding the contract method 0x693ce7f5.
//
// Solidity: function upgradeCreditFacade(address _creditFacade) returns()
func (_CreditConfigurator *CreditConfiguratorTransactorSession) UpgradeCreditFacade(_creditFacade common.Address) (*types.Transaction, error) {
	return _CreditConfigurator.Contract.UpgradeCreditFacade(&_CreditConfigurator.TransactOpts, _creditFacade)
}

// UpgradePriceOracle is a paid mutator transaction binding the contract method 0xf0527ac6.
//
// Solidity: function upgradePriceOracle() returns()
func (_CreditConfigurator *CreditConfiguratorTransactor) UpgradePriceOracle(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditConfigurator.contract.Transact(opts, "upgradePriceOracle")
}

// UpgradePriceOracle is a paid mutator transaction binding the contract method 0xf0527ac6.
//
// Solidity: function upgradePriceOracle() returns()
func (_CreditConfigurator *CreditConfiguratorSession) UpgradePriceOracle() (*types.Transaction, error) {
	return _CreditConfigurator.Contract.UpgradePriceOracle(&_CreditConfigurator.TransactOpts)
}

// UpgradePriceOracle is a paid mutator transaction binding the contract method 0xf0527ac6.
//
// Solidity: function upgradePriceOracle() returns()
func (_CreditConfigurator *CreditConfiguratorTransactorSession) UpgradePriceOracle() (*types.Transaction, error) {
	return _CreditConfigurator.Contract.UpgradePriceOracle(&_CreditConfigurator.TransactOpts)
}

// CreditConfiguratorContractAllowedIterator is returned from FilterContractAllowed and is used to iterate over the raw logs and unpacked data for ContractAllowed events raised by the CreditConfigurator contract.
type CreditConfiguratorContractAllowedIterator struct {
	Event *CreditConfiguratorContractAllowed // Event containing the contract specifics and raw log

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
func (it *CreditConfiguratorContractAllowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditConfiguratorContractAllowed)
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
		it.Event = new(CreditConfiguratorContractAllowed)
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
func (it *CreditConfiguratorContractAllowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditConfiguratorContractAllowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditConfiguratorContractAllowed represents a ContractAllowed event raised by the CreditConfigurator contract.
type CreditConfiguratorContractAllowed struct {
	Protocol common.Address
	Adapter  common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterContractAllowed is a free log retrieval operation binding the contract event 0x4bcbefaef68b99503d502f5a6abe7bca2b183ab8ac55457013c77d084ebd1305.
//
// Solidity: event ContractAllowed(address indexed protocol, address indexed adapter)
func (_CreditConfigurator *CreditConfiguratorFilterer) FilterContractAllowed(opts *bind.FilterOpts, protocol []common.Address, adapter []common.Address) (*CreditConfiguratorContractAllowedIterator, error) {

	var protocolRule []interface{}
	for _, protocolItem := range protocol {
		protocolRule = append(protocolRule, protocolItem)
	}
	var adapterRule []interface{}
	for _, adapterItem := range adapter {
		adapterRule = append(adapterRule, adapterItem)
	}

	logs, sub, err := _CreditConfigurator.contract.FilterLogs(opts, "ContractAllowed", protocolRule, adapterRule)
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorContractAllowedIterator{contract: _CreditConfigurator.contract, event: "ContractAllowed", logs: logs, sub: sub}, nil
}

// WatchContractAllowed is a free log subscription operation binding the contract event 0x4bcbefaef68b99503d502f5a6abe7bca2b183ab8ac55457013c77d084ebd1305.
//
// Solidity: event ContractAllowed(address indexed protocol, address indexed adapter)
func (_CreditConfigurator *CreditConfiguratorFilterer) WatchContractAllowed(opts *bind.WatchOpts, sink chan<- *CreditConfiguratorContractAllowed, protocol []common.Address, adapter []common.Address) (event.Subscription, error) {

	var protocolRule []interface{}
	for _, protocolItem := range protocol {
		protocolRule = append(protocolRule, protocolItem)
	}
	var adapterRule []interface{}
	for _, adapterItem := range adapter {
		adapterRule = append(adapterRule, adapterItem)
	}

	logs, sub, err := _CreditConfigurator.contract.WatchLogs(opts, "ContractAllowed", protocolRule, adapterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditConfiguratorContractAllowed)
				if err := _CreditConfigurator.contract.UnpackLog(event, "ContractAllowed", log); err != nil {
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
func (_CreditConfigurator *CreditConfiguratorFilterer) ParseContractAllowed(log types.Log) (*CreditConfiguratorContractAllowed, error) {
	event := new(CreditConfiguratorContractAllowed)
	if err := _CreditConfigurator.contract.UnpackLog(event, "ContractAllowed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CreditConfiguratorContractForbiddenIterator is returned from FilterContractForbidden and is used to iterate over the raw logs and unpacked data for ContractForbidden events raised by the CreditConfigurator contract.
type CreditConfiguratorContractForbiddenIterator struct {
	Event *CreditConfiguratorContractForbidden // Event containing the contract specifics and raw log

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
func (it *CreditConfiguratorContractForbiddenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditConfiguratorContractForbidden)
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
		it.Event = new(CreditConfiguratorContractForbidden)
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
func (it *CreditConfiguratorContractForbiddenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditConfiguratorContractForbiddenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditConfiguratorContractForbidden represents a ContractForbidden event raised by the CreditConfigurator contract.
type CreditConfiguratorContractForbidden struct {
	Protocol common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterContractForbidden is a free log retrieval operation binding the contract event 0xab9f405bf0c19b97f65a7031634db41569cd2f0e0376a610a1e977f9ab22b58f.
//
// Solidity: event ContractForbidden(address indexed protocol)
func (_CreditConfigurator *CreditConfiguratorFilterer) FilterContractForbidden(opts *bind.FilterOpts, protocol []common.Address) (*CreditConfiguratorContractForbiddenIterator, error) {

	var protocolRule []interface{}
	for _, protocolItem := range protocol {
		protocolRule = append(protocolRule, protocolItem)
	}

	logs, sub, err := _CreditConfigurator.contract.FilterLogs(opts, "ContractForbidden", protocolRule)
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorContractForbiddenIterator{contract: _CreditConfigurator.contract, event: "ContractForbidden", logs: logs, sub: sub}, nil
}

// WatchContractForbidden is a free log subscription operation binding the contract event 0xab9f405bf0c19b97f65a7031634db41569cd2f0e0376a610a1e977f9ab22b58f.
//
// Solidity: event ContractForbidden(address indexed protocol)
func (_CreditConfigurator *CreditConfiguratorFilterer) WatchContractForbidden(opts *bind.WatchOpts, sink chan<- *CreditConfiguratorContractForbidden, protocol []common.Address) (event.Subscription, error) {

	var protocolRule []interface{}
	for _, protocolItem := range protocol {
		protocolRule = append(protocolRule, protocolItem)
	}

	logs, sub, err := _CreditConfigurator.contract.WatchLogs(opts, "ContractForbidden", protocolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditConfiguratorContractForbidden)
				if err := _CreditConfigurator.contract.UnpackLog(event, "ContractForbidden", log); err != nil {
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
func (_CreditConfigurator *CreditConfiguratorFilterer) ParseContractForbidden(log types.Log) (*CreditConfiguratorContractForbidden, error) {
	event := new(CreditConfiguratorContractForbidden)
	if err := _CreditConfigurator.contract.UnpackLog(event, "ContractForbidden", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CreditConfiguratorCreditConfiguratorUpgradedIterator is returned from FilterCreditConfiguratorUpgraded and is used to iterate over the raw logs and unpacked data for CreditConfiguratorUpgraded events raised by the CreditConfigurator contract.
type CreditConfiguratorCreditConfiguratorUpgradedIterator struct {
	Event *CreditConfiguratorCreditConfiguratorUpgraded // Event containing the contract specifics and raw log

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
func (it *CreditConfiguratorCreditConfiguratorUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditConfiguratorCreditConfiguratorUpgraded)
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
		it.Event = new(CreditConfiguratorCreditConfiguratorUpgraded)
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
func (it *CreditConfiguratorCreditConfiguratorUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditConfiguratorCreditConfiguratorUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditConfiguratorCreditConfiguratorUpgraded represents a CreditConfiguratorUpgraded event raised by the CreditConfigurator contract.
type CreditConfiguratorCreditConfiguratorUpgraded struct {
	NewCreditConfigurator common.Address
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterCreditConfiguratorUpgraded is a free log retrieval operation binding the contract event 0x5a0b7d0f9c24b39256e112a0584b4c5ce38d8f1dee2e7c56f15b852604cdc886.
//
// Solidity: event CreditConfiguratorUpgraded(address indexed newCreditConfigurator)
func (_CreditConfigurator *CreditConfiguratorFilterer) FilterCreditConfiguratorUpgraded(opts *bind.FilterOpts, newCreditConfigurator []common.Address) (*CreditConfiguratorCreditConfiguratorUpgradedIterator, error) {

	var newCreditConfiguratorRule []interface{}
	for _, newCreditConfiguratorItem := range newCreditConfigurator {
		newCreditConfiguratorRule = append(newCreditConfiguratorRule, newCreditConfiguratorItem)
	}

	logs, sub, err := _CreditConfigurator.contract.FilterLogs(opts, "CreditConfiguratorUpgraded", newCreditConfiguratorRule)
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorCreditConfiguratorUpgradedIterator{contract: _CreditConfigurator.contract, event: "CreditConfiguratorUpgraded", logs: logs, sub: sub}, nil
}

// WatchCreditConfiguratorUpgraded is a free log subscription operation binding the contract event 0x5a0b7d0f9c24b39256e112a0584b4c5ce38d8f1dee2e7c56f15b852604cdc886.
//
// Solidity: event CreditConfiguratorUpgraded(address indexed newCreditConfigurator)
func (_CreditConfigurator *CreditConfiguratorFilterer) WatchCreditConfiguratorUpgraded(opts *bind.WatchOpts, sink chan<- *CreditConfiguratorCreditConfiguratorUpgraded, newCreditConfigurator []common.Address) (event.Subscription, error) {

	var newCreditConfiguratorRule []interface{}
	for _, newCreditConfiguratorItem := range newCreditConfigurator {
		newCreditConfiguratorRule = append(newCreditConfiguratorRule, newCreditConfiguratorItem)
	}

	logs, sub, err := _CreditConfigurator.contract.WatchLogs(opts, "CreditConfiguratorUpgraded", newCreditConfiguratorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditConfiguratorCreditConfiguratorUpgraded)
				if err := _CreditConfigurator.contract.UnpackLog(event, "CreditConfiguratorUpgraded", log); err != nil {
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

// ParseCreditConfiguratorUpgraded is a log parse operation binding the contract event 0x5a0b7d0f9c24b39256e112a0584b4c5ce38d8f1dee2e7c56f15b852604cdc886.
//
// Solidity: event CreditConfiguratorUpgraded(address indexed newCreditConfigurator)
func (_CreditConfigurator *CreditConfiguratorFilterer) ParseCreditConfiguratorUpgraded(log types.Log) (*CreditConfiguratorCreditConfiguratorUpgraded, error) {
	event := new(CreditConfiguratorCreditConfiguratorUpgraded)
	if err := _CreditConfigurator.contract.UnpackLog(event, "CreditConfiguratorUpgraded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CreditConfiguratorCreditFacadeUpgradedIterator is returned from FilterCreditFacadeUpgraded and is used to iterate over the raw logs and unpacked data for CreditFacadeUpgraded events raised by the CreditConfigurator contract.
type CreditConfiguratorCreditFacadeUpgradedIterator struct {
	Event *CreditConfiguratorCreditFacadeUpgraded // Event containing the contract specifics and raw log

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
func (it *CreditConfiguratorCreditFacadeUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditConfiguratorCreditFacadeUpgraded)
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
		it.Event = new(CreditConfiguratorCreditFacadeUpgraded)
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
func (it *CreditConfiguratorCreditFacadeUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditConfiguratorCreditFacadeUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditConfiguratorCreditFacadeUpgraded represents a CreditFacadeUpgraded event raised by the CreditConfigurator contract.
type CreditConfiguratorCreditFacadeUpgraded struct {
	NewCreditFacade common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCreditFacadeUpgraded is a free log retrieval operation binding the contract event 0xa8b21f72cb83bce808df32dc2330217d744a1c22f3e9e44e4b11bbf049d37d9d.
//
// Solidity: event CreditFacadeUpgraded(address indexed newCreditFacade)
func (_CreditConfigurator *CreditConfiguratorFilterer) FilterCreditFacadeUpgraded(opts *bind.FilterOpts, newCreditFacade []common.Address) (*CreditConfiguratorCreditFacadeUpgradedIterator, error) {

	var newCreditFacadeRule []interface{}
	for _, newCreditFacadeItem := range newCreditFacade {
		newCreditFacadeRule = append(newCreditFacadeRule, newCreditFacadeItem)
	}

	logs, sub, err := _CreditConfigurator.contract.FilterLogs(opts, "CreditFacadeUpgraded", newCreditFacadeRule)
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorCreditFacadeUpgradedIterator{contract: _CreditConfigurator.contract, event: "CreditFacadeUpgraded", logs: logs, sub: sub}, nil
}

// WatchCreditFacadeUpgraded is a free log subscription operation binding the contract event 0xa8b21f72cb83bce808df32dc2330217d744a1c22f3e9e44e4b11bbf049d37d9d.
//
// Solidity: event CreditFacadeUpgraded(address indexed newCreditFacade)
func (_CreditConfigurator *CreditConfiguratorFilterer) WatchCreditFacadeUpgraded(opts *bind.WatchOpts, sink chan<- *CreditConfiguratorCreditFacadeUpgraded, newCreditFacade []common.Address) (event.Subscription, error) {

	var newCreditFacadeRule []interface{}
	for _, newCreditFacadeItem := range newCreditFacade {
		newCreditFacadeRule = append(newCreditFacadeRule, newCreditFacadeItem)
	}

	logs, sub, err := _CreditConfigurator.contract.WatchLogs(opts, "CreditFacadeUpgraded", newCreditFacadeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditConfiguratorCreditFacadeUpgraded)
				if err := _CreditConfigurator.contract.UnpackLog(event, "CreditFacadeUpgraded", log); err != nil {
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

// ParseCreditFacadeUpgraded is a log parse operation binding the contract event 0xa8b21f72cb83bce808df32dc2330217d744a1c22f3e9e44e4b11bbf049d37d9d.
//
// Solidity: event CreditFacadeUpgraded(address indexed newCreditFacade)
func (_CreditConfigurator *CreditConfiguratorFilterer) ParseCreditFacadeUpgraded(log types.Log) (*CreditConfiguratorCreditFacadeUpgraded, error) {
	event := new(CreditConfiguratorCreditFacadeUpgraded)
	if err := _CreditConfigurator.contract.UnpackLog(event, "CreditFacadeUpgraded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CreditConfiguratorFastCheckParametersUpdatedIterator is returned from FilterFastCheckParametersUpdated and is used to iterate over the raw logs and unpacked data for FastCheckParametersUpdated events raised by the CreditConfigurator contract.
type CreditConfiguratorFastCheckParametersUpdatedIterator struct {
	Event *CreditConfiguratorFastCheckParametersUpdated // Event containing the contract specifics and raw log

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
func (it *CreditConfiguratorFastCheckParametersUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditConfiguratorFastCheckParametersUpdated)
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
		it.Event = new(CreditConfiguratorFastCheckParametersUpdated)
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
func (it *CreditConfiguratorFastCheckParametersUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditConfiguratorFastCheckParametersUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditConfiguratorFastCheckParametersUpdated represents a FastCheckParametersUpdated event raised by the CreditConfigurator contract.
type CreditConfiguratorFastCheckParametersUpdated struct {
	ChiThreshold   *big.Int
	FastCheckDelay *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterFastCheckParametersUpdated is a free log retrieval operation binding the contract event 0x86c4e28f8d9aa63f858bc3a8f0c0bbed25c8e045c4ac2e280eedc497246fdf29.
//
// Solidity: event FastCheckParametersUpdated(uint256 chiThreshold, uint256 fastCheckDelay)
func (_CreditConfigurator *CreditConfiguratorFilterer) FilterFastCheckParametersUpdated(opts *bind.FilterOpts) (*CreditConfiguratorFastCheckParametersUpdatedIterator, error) {

	logs, sub, err := _CreditConfigurator.contract.FilterLogs(opts, "FastCheckParametersUpdated")
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorFastCheckParametersUpdatedIterator{contract: _CreditConfigurator.contract, event: "FastCheckParametersUpdated", logs: logs, sub: sub}, nil
}

// WatchFastCheckParametersUpdated is a free log subscription operation binding the contract event 0x86c4e28f8d9aa63f858bc3a8f0c0bbed25c8e045c4ac2e280eedc497246fdf29.
//
// Solidity: event FastCheckParametersUpdated(uint256 chiThreshold, uint256 fastCheckDelay)
func (_CreditConfigurator *CreditConfiguratorFilterer) WatchFastCheckParametersUpdated(opts *bind.WatchOpts, sink chan<- *CreditConfiguratorFastCheckParametersUpdated) (event.Subscription, error) {

	logs, sub, err := _CreditConfigurator.contract.WatchLogs(opts, "FastCheckParametersUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditConfiguratorFastCheckParametersUpdated)
				if err := _CreditConfigurator.contract.UnpackLog(event, "FastCheckParametersUpdated", log); err != nil {
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

// ParseFastCheckParametersUpdated is a log parse operation binding the contract event 0x86c4e28f8d9aa63f858bc3a8f0c0bbed25c8e045c4ac2e280eedc497246fdf29.
//
// Solidity: event FastCheckParametersUpdated(uint256 chiThreshold, uint256 fastCheckDelay)
func (_CreditConfigurator *CreditConfiguratorFilterer) ParseFastCheckParametersUpdated(log types.Log) (*CreditConfiguratorFastCheckParametersUpdated, error) {
	event := new(CreditConfiguratorFastCheckParametersUpdated)
	if err := _CreditConfigurator.contract.UnpackLog(event, "FastCheckParametersUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CreditConfiguratorFeesUpdatedIterator is returned from FilterFeesUpdated and is used to iterate over the raw logs and unpacked data for FeesUpdated events raised by the CreditConfigurator contract.
type CreditConfiguratorFeesUpdatedIterator struct {
	Event *CreditConfiguratorFeesUpdated // Event containing the contract specifics and raw log

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
func (it *CreditConfiguratorFeesUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditConfiguratorFeesUpdated)
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
		it.Event = new(CreditConfiguratorFeesUpdated)
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
func (it *CreditConfiguratorFeesUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditConfiguratorFeesUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditConfiguratorFeesUpdated represents a FeesUpdated event raised by the CreditConfigurator contract.
type CreditConfiguratorFeesUpdated struct {
	FeeInterest        *big.Int
	FeeLiquidation     *big.Int
	LiquidationPremium *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterFeesUpdated is a free log retrieval operation binding the contract event 0xcf8a1e1d5f09cf3c97dbb653cd9a4d7aace9292fbc1bb8211febf2d400febbdd.
//
// Solidity: event FeesUpdated(uint256 feeInterest, uint256 feeLiquidation, uint256 liquidationPremium)
func (_CreditConfigurator *CreditConfiguratorFilterer) FilterFeesUpdated(opts *bind.FilterOpts) (*CreditConfiguratorFeesUpdatedIterator, error) {

	logs, sub, err := _CreditConfigurator.contract.FilterLogs(opts, "FeesUpdated")
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorFeesUpdatedIterator{contract: _CreditConfigurator.contract, event: "FeesUpdated", logs: logs, sub: sub}, nil
}

// WatchFeesUpdated is a free log subscription operation binding the contract event 0xcf8a1e1d5f09cf3c97dbb653cd9a4d7aace9292fbc1bb8211febf2d400febbdd.
//
// Solidity: event FeesUpdated(uint256 feeInterest, uint256 feeLiquidation, uint256 liquidationPremium)
func (_CreditConfigurator *CreditConfiguratorFilterer) WatchFeesUpdated(opts *bind.WatchOpts, sink chan<- *CreditConfiguratorFeesUpdated) (event.Subscription, error) {

	logs, sub, err := _CreditConfigurator.contract.WatchLogs(opts, "FeesUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditConfiguratorFeesUpdated)
				if err := _CreditConfigurator.contract.UnpackLog(event, "FeesUpdated", log); err != nil {
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

// ParseFeesUpdated is a log parse operation binding the contract event 0xcf8a1e1d5f09cf3c97dbb653cd9a4d7aace9292fbc1bb8211febf2d400febbdd.
//
// Solidity: event FeesUpdated(uint256 feeInterest, uint256 feeLiquidation, uint256 liquidationPremium)
func (_CreditConfigurator *CreditConfiguratorFilterer) ParseFeesUpdated(log types.Log) (*CreditConfiguratorFeesUpdated, error) {
	event := new(CreditConfiguratorFeesUpdated)
	if err := _CreditConfigurator.contract.UnpackLog(event, "FeesUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CreditConfiguratorLimitsUpdatedIterator is returned from FilterLimitsUpdated and is used to iterate over the raw logs and unpacked data for LimitsUpdated events raised by the CreditConfigurator contract.
type CreditConfiguratorLimitsUpdatedIterator struct {
	Event *CreditConfiguratorLimitsUpdated // Event containing the contract specifics and raw log

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
func (it *CreditConfiguratorLimitsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditConfiguratorLimitsUpdated)
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
		it.Event = new(CreditConfiguratorLimitsUpdated)
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
func (it *CreditConfiguratorLimitsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditConfiguratorLimitsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditConfiguratorLimitsUpdated represents a LimitsUpdated event raised by the CreditConfigurator contract.
type CreditConfiguratorLimitsUpdated struct {
	MinBorrowedAmount *big.Int
	MaxBorrowedAmount *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterLimitsUpdated is a free log retrieval operation binding the contract event 0x4d4981437d0211f9e6843eb024d9ada1fa3a99514d4343d4aece106dd11524bb.
//
// Solidity: event LimitsUpdated(uint256 minBorrowedAmount, uint256 maxBorrowedAmount)
func (_CreditConfigurator *CreditConfiguratorFilterer) FilterLimitsUpdated(opts *bind.FilterOpts) (*CreditConfiguratorLimitsUpdatedIterator, error) {

	logs, sub, err := _CreditConfigurator.contract.FilterLogs(opts, "LimitsUpdated")
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorLimitsUpdatedIterator{contract: _CreditConfigurator.contract, event: "LimitsUpdated", logs: logs, sub: sub}, nil
}

// WatchLimitsUpdated is a free log subscription operation binding the contract event 0x4d4981437d0211f9e6843eb024d9ada1fa3a99514d4343d4aece106dd11524bb.
//
// Solidity: event LimitsUpdated(uint256 minBorrowedAmount, uint256 maxBorrowedAmount)
func (_CreditConfigurator *CreditConfiguratorFilterer) WatchLimitsUpdated(opts *bind.WatchOpts, sink chan<- *CreditConfiguratorLimitsUpdated) (event.Subscription, error) {

	logs, sub, err := _CreditConfigurator.contract.WatchLogs(opts, "LimitsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditConfiguratorLimitsUpdated)
				if err := _CreditConfigurator.contract.UnpackLog(event, "LimitsUpdated", log); err != nil {
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

// ParseLimitsUpdated is a log parse operation binding the contract event 0x4d4981437d0211f9e6843eb024d9ada1fa3a99514d4343d4aece106dd11524bb.
//
// Solidity: event LimitsUpdated(uint256 minBorrowedAmount, uint256 maxBorrowedAmount)
func (_CreditConfigurator *CreditConfiguratorFilterer) ParseLimitsUpdated(log types.Log) (*CreditConfiguratorLimitsUpdated, error) {
	event := new(CreditConfiguratorLimitsUpdated)
	if err := _CreditConfigurator.contract.UnpackLog(event, "LimitsUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CreditConfiguratorPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the CreditConfigurator contract.
type CreditConfiguratorPausedIterator struct {
	Event *CreditConfiguratorPaused // Event containing the contract specifics and raw log

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
func (it *CreditConfiguratorPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditConfiguratorPaused)
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
		it.Event = new(CreditConfiguratorPaused)
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
func (it *CreditConfiguratorPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditConfiguratorPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditConfiguratorPaused represents a Paused event raised by the CreditConfigurator contract.
type CreditConfiguratorPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_CreditConfigurator *CreditConfiguratorFilterer) FilterPaused(opts *bind.FilterOpts) (*CreditConfiguratorPausedIterator, error) {

	logs, sub, err := _CreditConfigurator.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorPausedIterator{contract: _CreditConfigurator.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_CreditConfigurator *CreditConfiguratorFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *CreditConfiguratorPaused) (event.Subscription, error) {

	logs, sub, err := _CreditConfigurator.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditConfiguratorPaused)
				if err := _CreditConfigurator.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_CreditConfigurator *CreditConfiguratorFilterer) ParsePaused(log types.Log) (*CreditConfiguratorPaused, error) {
	event := new(CreditConfiguratorPaused)
	if err := _CreditConfigurator.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CreditConfiguratorPriceOracleUpgradedIterator is returned from FilterPriceOracleUpgraded and is used to iterate over the raw logs and unpacked data for PriceOracleUpgraded events raised by the CreditConfigurator contract.
type CreditConfiguratorPriceOracleUpgradedIterator struct {
	Event *CreditConfiguratorPriceOracleUpgraded // Event containing the contract specifics and raw log

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
func (it *CreditConfiguratorPriceOracleUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditConfiguratorPriceOracleUpgraded)
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
		it.Event = new(CreditConfiguratorPriceOracleUpgraded)
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
func (it *CreditConfiguratorPriceOracleUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditConfiguratorPriceOracleUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditConfiguratorPriceOracleUpgraded represents a PriceOracleUpgraded event raised by the CreditConfigurator contract.
type CreditConfiguratorPriceOracleUpgraded struct {
	NewPriceOracle common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPriceOracleUpgraded is a free log retrieval operation binding the contract event 0x3f82447be465b0b5a4a9e54c74d5f6ae73f2e9537f2cc1590a340524703d0961.
//
// Solidity: event PriceOracleUpgraded(address indexed newPriceOracle)
func (_CreditConfigurator *CreditConfiguratorFilterer) FilterPriceOracleUpgraded(opts *bind.FilterOpts, newPriceOracle []common.Address) (*CreditConfiguratorPriceOracleUpgradedIterator, error) {

	var newPriceOracleRule []interface{}
	for _, newPriceOracleItem := range newPriceOracle {
		newPriceOracleRule = append(newPriceOracleRule, newPriceOracleItem)
	}

	logs, sub, err := _CreditConfigurator.contract.FilterLogs(opts, "PriceOracleUpgraded", newPriceOracleRule)
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorPriceOracleUpgradedIterator{contract: _CreditConfigurator.contract, event: "PriceOracleUpgraded", logs: logs, sub: sub}, nil
}

// WatchPriceOracleUpgraded is a free log subscription operation binding the contract event 0x3f82447be465b0b5a4a9e54c74d5f6ae73f2e9537f2cc1590a340524703d0961.
//
// Solidity: event PriceOracleUpgraded(address indexed newPriceOracle)
func (_CreditConfigurator *CreditConfiguratorFilterer) WatchPriceOracleUpgraded(opts *bind.WatchOpts, sink chan<- *CreditConfiguratorPriceOracleUpgraded, newPriceOracle []common.Address) (event.Subscription, error) {

	var newPriceOracleRule []interface{}
	for _, newPriceOracleItem := range newPriceOracle {
		newPriceOracleRule = append(newPriceOracleRule, newPriceOracleItem)
	}

	logs, sub, err := _CreditConfigurator.contract.WatchLogs(opts, "PriceOracleUpgraded", newPriceOracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditConfiguratorPriceOracleUpgraded)
				if err := _CreditConfigurator.contract.UnpackLog(event, "PriceOracleUpgraded", log); err != nil {
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

// ParsePriceOracleUpgraded is a log parse operation binding the contract event 0x3f82447be465b0b5a4a9e54c74d5f6ae73f2e9537f2cc1590a340524703d0961.
//
// Solidity: event PriceOracleUpgraded(address indexed newPriceOracle)
func (_CreditConfigurator *CreditConfiguratorFilterer) ParsePriceOracleUpgraded(log types.Log) (*CreditConfiguratorPriceOracleUpgraded, error) {
	event := new(CreditConfiguratorPriceOracleUpgraded)
	if err := _CreditConfigurator.contract.UnpackLog(event, "PriceOracleUpgraded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CreditConfiguratorTokenAllowedIterator is returned from FilterTokenAllowed and is used to iterate over the raw logs and unpacked data for TokenAllowed events raised by the CreditConfigurator contract.
type CreditConfiguratorTokenAllowedIterator struct {
	Event *CreditConfiguratorTokenAllowed // Event containing the contract specifics and raw log

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
func (it *CreditConfiguratorTokenAllowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditConfiguratorTokenAllowed)
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
		it.Event = new(CreditConfiguratorTokenAllowed)
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
func (it *CreditConfiguratorTokenAllowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditConfiguratorTokenAllowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditConfiguratorTokenAllowed represents a TokenAllowed event raised by the CreditConfigurator contract.
type CreditConfiguratorTokenAllowed struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTokenAllowed is a free log retrieval operation binding the contract event 0xbeceb48aeaa805aeae57be163cca6249077a18734e408a85aa74e875c4373809.
//
// Solidity: event TokenAllowed(address indexed token)
func (_CreditConfigurator *CreditConfiguratorFilterer) FilterTokenAllowed(opts *bind.FilterOpts, token []common.Address) (*CreditConfiguratorTokenAllowedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CreditConfigurator.contract.FilterLogs(opts, "TokenAllowed", tokenRule)
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorTokenAllowedIterator{contract: _CreditConfigurator.contract, event: "TokenAllowed", logs: logs, sub: sub}, nil
}

// WatchTokenAllowed is a free log subscription operation binding the contract event 0xbeceb48aeaa805aeae57be163cca6249077a18734e408a85aa74e875c4373809.
//
// Solidity: event TokenAllowed(address indexed token)
func (_CreditConfigurator *CreditConfiguratorFilterer) WatchTokenAllowed(opts *bind.WatchOpts, sink chan<- *CreditConfiguratorTokenAllowed, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CreditConfigurator.contract.WatchLogs(opts, "TokenAllowed", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditConfiguratorTokenAllowed)
				if err := _CreditConfigurator.contract.UnpackLog(event, "TokenAllowed", log); err != nil {
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

// ParseTokenAllowed is a log parse operation binding the contract event 0xbeceb48aeaa805aeae57be163cca6249077a18734e408a85aa74e875c4373809.
//
// Solidity: event TokenAllowed(address indexed token)
func (_CreditConfigurator *CreditConfiguratorFilterer) ParseTokenAllowed(log types.Log) (*CreditConfiguratorTokenAllowed, error) {
	event := new(CreditConfiguratorTokenAllowed)
	if err := _CreditConfigurator.contract.UnpackLog(event, "TokenAllowed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CreditConfiguratorTokenForbiddenIterator is returned from FilterTokenForbidden and is used to iterate over the raw logs and unpacked data for TokenForbidden events raised by the CreditConfigurator contract.
type CreditConfiguratorTokenForbiddenIterator struct {
	Event *CreditConfiguratorTokenForbidden // Event containing the contract specifics and raw log

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
func (it *CreditConfiguratorTokenForbiddenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditConfiguratorTokenForbidden)
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
		it.Event = new(CreditConfiguratorTokenForbidden)
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
func (it *CreditConfiguratorTokenForbiddenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditConfiguratorTokenForbiddenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditConfiguratorTokenForbidden represents a TokenForbidden event raised by the CreditConfigurator contract.
type CreditConfiguratorTokenForbidden struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTokenForbidden is a free log retrieval operation binding the contract event 0xf17b849746e74d7186170c9553d4bbf60b4f8bb1ed81fe50c099b934fb078f05.
//
// Solidity: event TokenForbidden(address indexed token)
func (_CreditConfigurator *CreditConfiguratorFilterer) FilterTokenForbidden(opts *bind.FilterOpts, token []common.Address) (*CreditConfiguratorTokenForbiddenIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CreditConfigurator.contract.FilterLogs(opts, "TokenForbidden", tokenRule)
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorTokenForbiddenIterator{contract: _CreditConfigurator.contract, event: "TokenForbidden", logs: logs, sub: sub}, nil
}

// WatchTokenForbidden is a free log subscription operation binding the contract event 0xf17b849746e74d7186170c9553d4bbf60b4f8bb1ed81fe50c099b934fb078f05.
//
// Solidity: event TokenForbidden(address indexed token)
func (_CreditConfigurator *CreditConfiguratorFilterer) WatchTokenForbidden(opts *bind.WatchOpts, sink chan<- *CreditConfiguratorTokenForbidden, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CreditConfigurator.contract.WatchLogs(opts, "TokenForbidden", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditConfiguratorTokenForbidden)
				if err := _CreditConfigurator.contract.UnpackLog(event, "TokenForbidden", log); err != nil {
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

// ParseTokenForbidden is a log parse operation binding the contract event 0xf17b849746e74d7186170c9553d4bbf60b4f8bb1ed81fe50c099b934fb078f05.
//
// Solidity: event TokenForbidden(address indexed token)
func (_CreditConfigurator *CreditConfiguratorFilterer) ParseTokenForbidden(log types.Log) (*CreditConfiguratorTokenForbidden, error) {
	event := new(CreditConfiguratorTokenForbidden)
	if err := _CreditConfigurator.contract.UnpackLog(event, "TokenForbidden", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CreditConfiguratorTokenLiquidationThresholdUpdatedIterator is returned from FilterTokenLiquidationThresholdUpdated and is used to iterate over the raw logs and unpacked data for TokenLiquidationThresholdUpdated events raised by the CreditConfigurator contract.
type CreditConfiguratorTokenLiquidationThresholdUpdatedIterator struct {
	Event *CreditConfiguratorTokenLiquidationThresholdUpdated // Event containing the contract specifics and raw log

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
func (it *CreditConfiguratorTokenLiquidationThresholdUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditConfiguratorTokenLiquidationThresholdUpdated)
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
		it.Event = new(CreditConfiguratorTokenLiquidationThresholdUpdated)
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
func (it *CreditConfiguratorTokenLiquidationThresholdUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditConfiguratorTokenLiquidationThresholdUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditConfiguratorTokenLiquidationThresholdUpdated represents a TokenLiquidationThresholdUpdated event raised by the CreditConfigurator contract.
type CreditConfiguratorTokenLiquidationThresholdUpdated struct {
	Token              common.Address
	LiquidityThreshold *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterTokenLiquidationThresholdUpdated is a free log retrieval operation binding the contract event 0xd5e67007f7834ddcc2f493b83810f7c6ab74e1bba7b88847f6d4adf9e03fbe82.
//
// Solidity: event TokenLiquidationThresholdUpdated(address indexed token, uint256 liquidityThreshold)
func (_CreditConfigurator *CreditConfiguratorFilterer) FilterTokenLiquidationThresholdUpdated(opts *bind.FilterOpts, token []common.Address) (*CreditConfiguratorTokenLiquidationThresholdUpdatedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CreditConfigurator.contract.FilterLogs(opts, "TokenLiquidationThresholdUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorTokenLiquidationThresholdUpdatedIterator{contract: _CreditConfigurator.contract, event: "TokenLiquidationThresholdUpdated", logs: logs, sub: sub}, nil
}

// WatchTokenLiquidationThresholdUpdated is a free log subscription operation binding the contract event 0xd5e67007f7834ddcc2f493b83810f7c6ab74e1bba7b88847f6d4adf9e03fbe82.
//
// Solidity: event TokenLiquidationThresholdUpdated(address indexed token, uint256 liquidityThreshold)
func (_CreditConfigurator *CreditConfiguratorFilterer) WatchTokenLiquidationThresholdUpdated(opts *bind.WatchOpts, sink chan<- *CreditConfiguratorTokenLiquidationThresholdUpdated, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CreditConfigurator.contract.WatchLogs(opts, "TokenLiquidationThresholdUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditConfiguratorTokenLiquidationThresholdUpdated)
				if err := _CreditConfigurator.contract.UnpackLog(event, "TokenLiquidationThresholdUpdated", log); err != nil {
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

// ParseTokenLiquidationThresholdUpdated is a log parse operation binding the contract event 0xd5e67007f7834ddcc2f493b83810f7c6ab74e1bba7b88847f6d4adf9e03fbe82.
//
// Solidity: event TokenLiquidationThresholdUpdated(address indexed token, uint256 liquidityThreshold)
func (_CreditConfigurator *CreditConfiguratorFilterer) ParseTokenLiquidationThresholdUpdated(log types.Log) (*CreditConfiguratorTokenLiquidationThresholdUpdated, error) {
	event := new(CreditConfiguratorTokenLiquidationThresholdUpdated)
	if err := _CreditConfigurator.contract.UnpackLog(event, "TokenLiquidationThresholdUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CreditConfiguratorUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the CreditConfigurator contract.
type CreditConfiguratorUnpausedIterator struct {
	Event *CreditConfiguratorUnpaused // Event containing the contract specifics and raw log

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
func (it *CreditConfiguratorUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditConfiguratorUnpaused)
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
		it.Event = new(CreditConfiguratorUnpaused)
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
func (it *CreditConfiguratorUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditConfiguratorUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditConfiguratorUnpaused represents a Unpaused event raised by the CreditConfigurator contract.
type CreditConfiguratorUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_CreditConfigurator *CreditConfiguratorFilterer) FilterUnpaused(opts *bind.FilterOpts) (*CreditConfiguratorUnpausedIterator, error) {

	logs, sub, err := _CreditConfigurator.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorUnpausedIterator{contract: _CreditConfigurator.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_CreditConfigurator *CreditConfiguratorFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *CreditConfiguratorUnpaused) (event.Subscription, error) {

	logs, sub, err := _CreditConfigurator.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditConfiguratorUnpaused)
				if err := _CreditConfigurator.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_CreditConfigurator *CreditConfiguratorFilterer) ParseUnpaused(log types.Log) (*CreditConfiguratorUnpaused, error) {
	event := new(CreditConfiguratorUnpaused)
	if err := _CreditConfigurator.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	return event, nil
}
