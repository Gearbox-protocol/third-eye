// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package creditManagerMockForFilter

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

// CreditManagerMockForFilterMetaData contains all meta data concerning the CreditManagerMockForFilter contract.
var CreditManagerMockForFilterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"calcLinearCumulative_RAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"checkAndEnableToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"checkCollateralChange\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_creditFilterAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_underlyingToken\",\"type\":\"address\"}],\"name\":\"connectFilter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeInterest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeLiquidation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeSuccess\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"healthFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"}],\"name\":\"initEnabledTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidationDiscount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxLeverageFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minHealthFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolService\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"setFeeLiquidation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"setLinearCumulative\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"setLiquidationDiscount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"setMaxLeverageFactor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlyingToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateUnderlyingTokenLiquidationThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// CreditManagerMockForFilterABI is the input ABI used to generate the binding from.
// Deprecated: Use CreditManagerMockForFilterMetaData.ABI instead.
var CreditManagerMockForFilterABI = CreditManagerMockForFilterMetaData.ABI

// CreditManagerMockForFilter is an auto generated Go binding around an Ethereum contract.
type CreditManagerMockForFilter struct {
	CreditManagerMockForFilterCaller     // Read-only binding to the contract
	CreditManagerMockForFilterTransactor // Write-only binding to the contract
	CreditManagerMockForFilterFilterer   // Log filterer for contract events
}

// CreditManagerMockForFilterCaller is an auto generated read-only Go binding around an Ethereum contract.
type CreditManagerMockForFilterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditManagerMockForFilterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CreditManagerMockForFilterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditManagerMockForFilterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CreditManagerMockForFilterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditManagerMockForFilterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CreditManagerMockForFilterSession struct {
	Contract     *CreditManagerMockForFilter // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// CreditManagerMockForFilterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CreditManagerMockForFilterCallerSession struct {
	Contract *CreditManagerMockForFilterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// CreditManagerMockForFilterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CreditManagerMockForFilterTransactorSession struct {
	Contract     *CreditManagerMockForFilterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// CreditManagerMockForFilterRaw is an auto generated low-level Go binding around an Ethereum contract.
type CreditManagerMockForFilterRaw struct {
	Contract *CreditManagerMockForFilter // Generic contract binding to access the raw methods on
}

// CreditManagerMockForFilterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CreditManagerMockForFilterCallerRaw struct {
	Contract *CreditManagerMockForFilterCaller // Generic read-only contract binding to access the raw methods on
}

// CreditManagerMockForFilterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CreditManagerMockForFilterTransactorRaw struct {
	Contract *CreditManagerMockForFilterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCreditManagerMockForFilter creates a new instance of CreditManagerMockForFilter, bound to a specific deployed contract.
func NewCreditManagerMockForFilter(address common.Address, backend bind.ContractBackend) (*CreditManagerMockForFilter, error) {
	contract, err := bindCreditManagerMockForFilter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CreditManagerMockForFilter{CreditManagerMockForFilterCaller: CreditManagerMockForFilterCaller{contract: contract}, CreditManagerMockForFilterTransactor: CreditManagerMockForFilterTransactor{contract: contract}, CreditManagerMockForFilterFilterer: CreditManagerMockForFilterFilterer{contract: contract}}, nil
}

// NewCreditManagerMockForFilterCaller creates a new read-only instance of CreditManagerMockForFilter, bound to a specific deployed contract.
func NewCreditManagerMockForFilterCaller(address common.Address, caller bind.ContractCaller) (*CreditManagerMockForFilterCaller, error) {
	contract, err := bindCreditManagerMockForFilter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CreditManagerMockForFilterCaller{contract: contract}, nil
}

// NewCreditManagerMockForFilterTransactor creates a new write-only instance of CreditManagerMockForFilter, bound to a specific deployed contract.
func NewCreditManagerMockForFilterTransactor(address common.Address, transactor bind.ContractTransactor) (*CreditManagerMockForFilterTransactor, error) {
	contract, err := bindCreditManagerMockForFilter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CreditManagerMockForFilterTransactor{contract: contract}, nil
}

// NewCreditManagerMockForFilterFilterer creates a new log filterer instance of CreditManagerMockForFilter, bound to a specific deployed contract.
func NewCreditManagerMockForFilterFilterer(address common.Address, filterer bind.ContractFilterer) (*CreditManagerMockForFilterFilterer, error) {
	contract, err := bindCreditManagerMockForFilter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CreditManagerMockForFilterFilterer{contract: contract}, nil
}

// bindCreditManagerMockForFilter binds a generic wrapper to an already deployed contract.
func bindCreditManagerMockForFilter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CreditManagerMockForFilterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CreditManagerMockForFilter *CreditManagerMockForFilterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CreditManagerMockForFilter.Contract.CreditManagerMockForFilterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CreditManagerMockForFilter *CreditManagerMockForFilterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditManagerMockForFilter.Contract.CreditManagerMockForFilterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CreditManagerMockForFilter *CreditManagerMockForFilterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CreditManagerMockForFilter.Contract.CreditManagerMockForFilterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CreditManagerMockForFilter *CreditManagerMockForFilterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CreditManagerMockForFilter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CreditManagerMockForFilter *CreditManagerMockForFilterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditManagerMockForFilter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CreditManagerMockForFilter *CreditManagerMockForFilterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CreditManagerMockForFilter.Contract.contract.Transact(opts, method, params...)
}

// CalcLinearCumulativeRAY is a free data retrieval call binding the contract method 0x0fce70fb.
//
// Solidity: function calcLinearCumulative_RAY() view returns(uint256)
func (_CreditManagerMockForFilter *CreditManagerMockForFilterCaller) CalcLinearCumulativeRAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CreditManagerMockForFilter.contract.Call(opts, &out, "calcLinearCumulative_RAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcLinearCumulativeRAY is a free data retrieval call binding the contract method 0x0fce70fb.
//
// Solidity: function calcLinearCumulative_RAY() view returns(uint256)
func (_CreditManagerMockForFilter *CreditManagerMockForFilterSession) CalcLinearCumulativeRAY() (*big.Int, error) {
	return _CreditManagerMockForFilter.Contract.CalcLinearCumulativeRAY(&_CreditManagerMockForFilter.CallOpts)
}

// CalcLinearCumulativeRAY is a free data retrieval call binding the contract method 0x0fce70fb.
//
// Solidity: function calcLinearCumulative_RAY() view returns(uint256)
func (_CreditManagerMockForFilter *CreditManagerMockForFilterCallerSession) CalcLinearCumulativeRAY() (*big.Int, error) {
	return _CreditManagerMockForFilter.Contract.CalcLinearCumulativeRAY(&_CreditManagerMockForFilter.CallOpts)
}

// FeeInterest is a free data retrieval call binding the contract method 0x5e0b63d3.
//
// Solidity: function feeInterest() view returns(uint256)
func (_CreditManagerMockForFilter *CreditManagerMockForFilterCaller) FeeInterest(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CreditManagerMockForFilter.contract.Call(opts, &out, "feeInterest")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeInterest is a free data retrieval call binding the contract method 0x5e0b63d3.
//
// Solidity: function feeInterest() view returns(uint256)
func (_CreditManagerMockForFilter *CreditManagerMockForFilterSession) FeeInterest() (*big.Int, error) {
	return _CreditManagerMockForFilter.Contract.FeeInterest(&_CreditManagerMockForFilter.CallOpts)
}

// FeeInterest is a free data retrieval call binding the contract method 0x5e0b63d3.
//
// Solidity: function feeInterest() view returns(uint256)
func (_CreditManagerMockForFilter *CreditManagerMockForFilterCallerSession) FeeInterest() (*big.Int, error) {
	return _CreditManagerMockForFilter.Contract.FeeInterest(&_CreditManagerMockForFilter.CallOpts)
}

// FeeLiquidation is a free data retrieval call binding the contract method 0x3915ffaa.
//
// Solidity: function feeLiquidation() view returns(uint256)
func (_CreditManagerMockForFilter *CreditManagerMockForFilterCaller) FeeLiquidation(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CreditManagerMockForFilter.contract.Call(opts, &out, "feeLiquidation")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeLiquidation is a free data retrieval call binding the contract method 0x3915ffaa.
//
// Solidity: function feeLiquidation() view returns(uint256)
func (_CreditManagerMockForFilter *CreditManagerMockForFilterSession) FeeLiquidation() (*big.Int, error) {
	return _CreditManagerMockForFilter.Contract.FeeLiquidation(&_CreditManagerMockForFilter.CallOpts)
}

// FeeLiquidation is a free data retrieval call binding the contract method 0x3915ffaa.
//
// Solidity: function feeLiquidation() view returns(uint256)
func (_CreditManagerMockForFilter *CreditManagerMockForFilterCallerSession) FeeLiquidation() (*big.Int, error) {
	return _CreditManagerMockForFilter.Contract.FeeLiquidation(&_CreditManagerMockForFilter.CallOpts)
}

// FeeSuccess is a free data retrieval call binding the contract method 0x9e52f12e.
//
// Solidity: function feeSuccess() view returns(uint256)
func (_CreditManagerMockForFilter *CreditManagerMockForFilterCaller) FeeSuccess(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CreditManagerMockForFilter.contract.Call(opts, &out, "feeSuccess")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeSuccess is a free data retrieval call binding the contract method 0x9e52f12e.
//
// Solidity: function feeSuccess() view returns(uint256)
func (_CreditManagerMockForFilter *CreditManagerMockForFilterSession) FeeSuccess() (*big.Int, error) {
	return _CreditManagerMockForFilter.Contract.FeeSuccess(&_CreditManagerMockForFilter.CallOpts)
}

// FeeSuccess is a free data retrieval call binding the contract method 0x9e52f12e.
//
// Solidity: function feeSuccess() view returns(uint256)
func (_CreditManagerMockForFilter *CreditManagerMockForFilterCallerSession) FeeSuccess() (*big.Int, error) {
	return _CreditManagerMockForFilter.Contract.FeeSuccess(&_CreditManagerMockForFilter.CallOpts)
}

// HealthFactor is a free data retrieval call binding the contract method 0x22841f01.
//
// Solidity: function healthFactor() view returns(uint256)
func (_CreditManagerMockForFilter *CreditManagerMockForFilterCaller) HealthFactor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CreditManagerMockForFilter.contract.Call(opts, &out, "healthFactor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HealthFactor is a free data retrieval call binding the contract method 0x22841f01.
//
// Solidity: function healthFactor() view returns(uint256)
func (_CreditManagerMockForFilter *CreditManagerMockForFilterSession) HealthFactor() (*big.Int, error) {
	return _CreditManagerMockForFilter.Contract.HealthFactor(&_CreditManagerMockForFilter.CallOpts)
}

// HealthFactor is a free data retrieval call binding the contract method 0x22841f01.
//
// Solidity: function healthFactor() view returns(uint256)
func (_CreditManagerMockForFilter *CreditManagerMockForFilterCallerSession) HealthFactor() (*big.Int, error) {
	return _CreditManagerMockForFilter.Contract.HealthFactor(&_CreditManagerMockForFilter.CallOpts)
}

// LiquidationDiscount is a free data retrieval call binding the contract method 0x8053fcbe.
//
// Solidity: function liquidationDiscount() view returns(uint256)
func (_CreditManagerMockForFilter *CreditManagerMockForFilterCaller) LiquidationDiscount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CreditManagerMockForFilter.contract.Call(opts, &out, "liquidationDiscount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LiquidationDiscount is a free data retrieval call binding the contract method 0x8053fcbe.
//
// Solidity: function liquidationDiscount() view returns(uint256)
func (_CreditManagerMockForFilter *CreditManagerMockForFilterSession) LiquidationDiscount() (*big.Int, error) {
	return _CreditManagerMockForFilter.Contract.LiquidationDiscount(&_CreditManagerMockForFilter.CallOpts)
}

// LiquidationDiscount is a free data retrieval call binding the contract method 0x8053fcbe.
//
// Solidity: function liquidationDiscount() view returns(uint256)
func (_CreditManagerMockForFilter *CreditManagerMockForFilterCallerSession) LiquidationDiscount() (*big.Int, error) {
	return _CreditManagerMockForFilter.Contract.LiquidationDiscount(&_CreditManagerMockForFilter.CallOpts)
}

// MaxLeverageFactor is a free data retrieval call binding the contract method 0xb2c53a6c.
//
// Solidity: function maxLeverageFactor() view returns(uint256)
func (_CreditManagerMockForFilter *CreditManagerMockForFilterCaller) MaxLeverageFactor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CreditManagerMockForFilter.contract.Call(opts, &out, "maxLeverageFactor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxLeverageFactor is a free data retrieval call binding the contract method 0xb2c53a6c.
//
// Solidity: function maxLeverageFactor() view returns(uint256)
func (_CreditManagerMockForFilter *CreditManagerMockForFilterSession) MaxLeverageFactor() (*big.Int, error) {
	return _CreditManagerMockForFilter.Contract.MaxLeverageFactor(&_CreditManagerMockForFilter.CallOpts)
}

// MaxLeverageFactor is a free data retrieval call binding the contract method 0xb2c53a6c.
//
// Solidity: function maxLeverageFactor() view returns(uint256)
func (_CreditManagerMockForFilter *CreditManagerMockForFilterCallerSession) MaxLeverageFactor() (*big.Int, error) {
	return _CreditManagerMockForFilter.Contract.MaxLeverageFactor(&_CreditManagerMockForFilter.CallOpts)
}

// MinHealthFactor is a free data retrieval call binding the contract method 0xe1b4264c.
//
// Solidity: function minHealthFactor() view returns(uint256)
func (_CreditManagerMockForFilter *CreditManagerMockForFilterCaller) MinHealthFactor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CreditManagerMockForFilter.contract.Call(opts, &out, "minHealthFactor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinHealthFactor is a free data retrieval call binding the contract method 0xe1b4264c.
//
// Solidity: function minHealthFactor() view returns(uint256)
func (_CreditManagerMockForFilter *CreditManagerMockForFilterSession) MinHealthFactor() (*big.Int, error) {
	return _CreditManagerMockForFilter.Contract.MinHealthFactor(&_CreditManagerMockForFilter.CallOpts)
}

// MinHealthFactor is a free data retrieval call binding the contract method 0xe1b4264c.
//
// Solidity: function minHealthFactor() view returns(uint256)
func (_CreditManagerMockForFilter *CreditManagerMockForFilterCallerSession) MinHealthFactor() (*big.Int, error) {
	return _CreditManagerMockForFilter.Contract.MinHealthFactor(&_CreditManagerMockForFilter.CallOpts)
}

// PoolService is a free data retrieval call binding the contract method 0x570a7af2.
//
// Solidity: function poolService() view returns(address)
func (_CreditManagerMockForFilter *CreditManagerMockForFilterCaller) PoolService(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditManagerMockForFilter.contract.Call(opts, &out, "poolService")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolService is a free data retrieval call binding the contract method 0x570a7af2.
//
// Solidity: function poolService() view returns(address)
func (_CreditManagerMockForFilter *CreditManagerMockForFilterSession) PoolService() (common.Address, error) {
	return _CreditManagerMockForFilter.Contract.PoolService(&_CreditManagerMockForFilter.CallOpts)
}

// PoolService is a free data retrieval call binding the contract method 0x570a7af2.
//
// Solidity: function poolService() view returns(address)
func (_CreditManagerMockForFilter *CreditManagerMockForFilterCallerSession) PoolService() (common.Address, error) {
	return _CreditManagerMockForFilter.Contract.PoolService(&_CreditManagerMockForFilter.CallOpts)
}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_CreditManagerMockForFilter *CreditManagerMockForFilterCaller) UnderlyingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditManagerMockForFilter.contract.Call(opts, &out, "underlyingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_CreditManagerMockForFilter *CreditManagerMockForFilterSession) UnderlyingToken() (common.Address, error) {
	return _CreditManagerMockForFilter.Contract.UnderlyingToken(&_CreditManagerMockForFilter.CallOpts)
}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_CreditManagerMockForFilter *CreditManagerMockForFilterCallerSession) UnderlyingToken() (common.Address, error) {
	return _CreditManagerMockForFilter.Contract.UnderlyingToken(&_CreditManagerMockForFilter.CallOpts)
}

// CheckAndEnableToken is a paid mutator transaction binding the contract method 0x51e3f160.
//
// Solidity: function checkAndEnableToken(address creditAccount, address token) returns()
func (_CreditManagerMockForFilter *CreditManagerMockForFilterTransactor) CheckAndEnableToken(opts *bind.TransactOpts, creditAccount common.Address, token common.Address) (*types.Transaction, error) {
	return _CreditManagerMockForFilter.contract.Transact(opts, "checkAndEnableToken", creditAccount, token)
}

// CheckAndEnableToken is a paid mutator transaction binding the contract method 0x51e3f160.
//
// Solidity: function checkAndEnableToken(address creditAccount, address token) returns()
func (_CreditManagerMockForFilter *CreditManagerMockForFilterSession) CheckAndEnableToken(creditAccount common.Address, token common.Address) (*types.Transaction, error) {
	return _CreditManagerMockForFilter.Contract.CheckAndEnableToken(&_CreditManagerMockForFilter.TransactOpts, creditAccount, token)
}

// CheckAndEnableToken is a paid mutator transaction binding the contract method 0x51e3f160.
//
// Solidity: function checkAndEnableToken(address creditAccount, address token) returns()
func (_CreditManagerMockForFilter *CreditManagerMockForFilterTransactorSession) CheckAndEnableToken(creditAccount common.Address, token common.Address) (*types.Transaction, error) {
	return _CreditManagerMockForFilter.Contract.CheckAndEnableToken(&_CreditManagerMockForFilter.TransactOpts, creditAccount, token)
}

// CheckCollateralChange is a paid mutator transaction binding the contract method 0xe1c8ef0d.
//
// Solidity: function checkCollateralChange(address creditAccount, address tokenIn, address tokenOut, uint256 amountIn, uint256 amountOut) returns()
func (_CreditManagerMockForFilter *CreditManagerMockForFilterTransactor) CheckCollateralChange(opts *bind.TransactOpts, creditAccount common.Address, tokenIn common.Address, tokenOut common.Address, amountIn *big.Int, amountOut *big.Int) (*types.Transaction, error) {
	return _CreditManagerMockForFilter.contract.Transact(opts, "checkCollateralChange", creditAccount, tokenIn, tokenOut, amountIn, amountOut)
}

// CheckCollateralChange is a paid mutator transaction binding the contract method 0xe1c8ef0d.
//
// Solidity: function checkCollateralChange(address creditAccount, address tokenIn, address tokenOut, uint256 amountIn, uint256 amountOut) returns()
func (_CreditManagerMockForFilter *CreditManagerMockForFilterSession) CheckCollateralChange(creditAccount common.Address, tokenIn common.Address, tokenOut common.Address, amountIn *big.Int, amountOut *big.Int) (*types.Transaction, error) {
	return _CreditManagerMockForFilter.Contract.CheckCollateralChange(&_CreditManagerMockForFilter.TransactOpts, creditAccount, tokenIn, tokenOut, amountIn, amountOut)
}

// CheckCollateralChange is a paid mutator transaction binding the contract method 0xe1c8ef0d.
//
// Solidity: function checkCollateralChange(address creditAccount, address tokenIn, address tokenOut, uint256 amountIn, uint256 amountOut) returns()
func (_CreditManagerMockForFilter *CreditManagerMockForFilterTransactorSession) CheckCollateralChange(creditAccount common.Address, tokenIn common.Address, tokenOut common.Address, amountIn *big.Int, amountOut *big.Int) (*types.Transaction, error) {
	return _CreditManagerMockForFilter.Contract.CheckCollateralChange(&_CreditManagerMockForFilter.TransactOpts, creditAccount, tokenIn, tokenOut, amountIn, amountOut)
}

// ConnectFilter is a paid mutator transaction binding the contract method 0xaeb1e31c.
//
// Solidity: function connectFilter(address _creditFilterAddress, address _underlyingToken) returns()
func (_CreditManagerMockForFilter *CreditManagerMockForFilterTransactor) ConnectFilter(opts *bind.TransactOpts, _creditFilterAddress common.Address, _underlyingToken common.Address) (*types.Transaction, error) {
	return _CreditManagerMockForFilter.contract.Transact(opts, "connectFilter", _creditFilterAddress, _underlyingToken)
}

// ConnectFilter is a paid mutator transaction binding the contract method 0xaeb1e31c.
//
// Solidity: function connectFilter(address _creditFilterAddress, address _underlyingToken) returns()
func (_CreditManagerMockForFilter *CreditManagerMockForFilterSession) ConnectFilter(_creditFilterAddress common.Address, _underlyingToken common.Address) (*types.Transaction, error) {
	return _CreditManagerMockForFilter.Contract.ConnectFilter(&_CreditManagerMockForFilter.TransactOpts, _creditFilterAddress, _underlyingToken)
}

// ConnectFilter is a paid mutator transaction binding the contract method 0xaeb1e31c.
//
// Solidity: function connectFilter(address _creditFilterAddress, address _underlyingToken) returns()
func (_CreditManagerMockForFilter *CreditManagerMockForFilterTransactorSession) ConnectFilter(_creditFilterAddress common.Address, _underlyingToken common.Address) (*types.Transaction, error) {
	return _CreditManagerMockForFilter.Contract.ConnectFilter(&_CreditManagerMockForFilter.TransactOpts, _creditFilterAddress, _underlyingToken)
}

// InitEnabledTokens is a paid mutator transaction binding the contract method 0xe54fe9c8.
//
// Solidity: function initEnabledTokens(address creditAccount) returns()
func (_CreditManagerMockForFilter *CreditManagerMockForFilterTransactor) InitEnabledTokens(opts *bind.TransactOpts, creditAccount common.Address) (*types.Transaction, error) {
	return _CreditManagerMockForFilter.contract.Transact(opts, "initEnabledTokens", creditAccount)
}

// InitEnabledTokens is a paid mutator transaction binding the contract method 0xe54fe9c8.
//
// Solidity: function initEnabledTokens(address creditAccount) returns()
func (_CreditManagerMockForFilter *CreditManagerMockForFilterSession) InitEnabledTokens(creditAccount common.Address) (*types.Transaction, error) {
	return _CreditManagerMockForFilter.Contract.InitEnabledTokens(&_CreditManagerMockForFilter.TransactOpts, creditAccount)
}

// InitEnabledTokens is a paid mutator transaction binding the contract method 0xe54fe9c8.
//
// Solidity: function initEnabledTokens(address creditAccount) returns()
func (_CreditManagerMockForFilter *CreditManagerMockForFilterTransactorSession) InitEnabledTokens(creditAccount common.Address) (*types.Transaction, error) {
	return _CreditManagerMockForFilter.Contract.InitEnabledTokens(&_CreditManagerMockForFilter.TransactOpts, creditAccount)
}

// SetFeeLiquidation is a paid mutator transaction binding the contract method 0x115c2ec7.
//
// Solidity: function setFeeLiquidation(uint256 _value) returns()
func (_CreditManagerMockForFilter *CreditManagerMockForFilterTransactor) SetFeeLiquidation(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _CreditManagerMockForFilter.contract.Transact(opts, "setFeeLiquidation", _value)
}

// SetFeeLiquidation is a paid mutator transaction binding the contract method 0x115c2ec7.
//
// Solidity: function setFeeLiquidation(uint256 _value) returns()
func (_CreditManagerMockForFilter *CreditManagerMockForFilterSession) SetFeeLiquidation(_value *big.Int) (*types.Transaction, error) {
	return _CreditManagerMockForFilter.Contract.SetFeeLiquidation(&_CreditManagerMockForFilter.TransactOpts, _value)
}

// SetFeeLiquidation is a paid mutator transaction binding the contract method 0x115c2ec7.
//
// Solidity: function setFeeLiquidation(uint256 _value) returns()
func (_CreditManagerMockForFilter *CreditManagerMockForFilterTransactorSession) SetFeeLiquidation(_value *big.Int) (*types.Transaction, error) {
	return _CreditManagerMockForFilter.Contract.SetFeeLiquidation(&_CreditManagerMockForFilter.TransactOpts, _value)
}

// SetLinearCumulative is a paid mutator transaction binding the contract method 0x134b4ac5.
//
// Solidity: function setLinearCumulative(uint256 newValue) returns()
func (_CreditManagerMockForFilter *CreditManagerMockForFilterTransactor) SetLinearCumulative(opts *bind.TransactOpts, newValue *big.Int) (*types.Transaction, error) {
	return _CreditManagerMockForFilter.contract.Transact(opts, "setLinearCumulative", newValue)
}

// SetLinearCumulative is a paid mutator transaction binding the contract method 0x134b4ac5.
//
// Solidity: function setLinearCumulative(uint256 newValue) returns()
func (_CreditManagerMockForFilter *CreditManagerMockForFilterSession) SetLinearCumulative(newValue *big.Int) (*types.Transaction, error) {
	return _CreditManagerMockForFilter.Contract.SetLinearCumulative(&_CreditManagerMockForFilter.TransactOpts, newValue)
}

// SetLinearCumulative is a paid mutator transaction binding the contract method 0x134b4ac5.
//
// Solidity: function setLinearCumulative(uint256 newValue) returns()
func (_CreditManagerMockForFilter *CreditManagerMockForFilterTransactorSession) SetLinearCumulative(newValue *big.Int) (*types.Transaction, error) {
	return _CreditManagerMockForFilter.Contract.SetLinearCumulative(&_CreditManagerMockForFilter.TransactOpts, newValue)
}

// SetLiquidationDiscount is a paid mutator transaction binding the contract method 0xc5e86154.
//
// Solidity: function setLiquidationDiscount(uint256 _value) returns()
func (_CreditManagerMockForFilter *CreditManagerMockForFilterTransactor) SetLiquidationDiscount(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _CreditManagerMockForFilter.contract.Transact(opts, "setLiquidationDiscount", _value)
}

// SetLiquidationDiscount is a paid mutator transaction binding the contract method 0xc5e86154.
//
// Solidity: function setLiquidationDiscount(uint256 _value) returns()
func (_CreditManagerMockForFilter *CreditManagerMockForFilterSession) SetLiquidationDiscount(_value *big.Int) (*types.Transaction, error) {
	return _CreditManagerMockForFilter.Contract.SetLiquidationDiscount(&_CreditManagerMockForFilter.TransactOpts, _value)
}

// SetLiquidationDiscount is a paid mutator transaction binding the contract method 0xc5e86154.
//
// Solidity: function setLiquidationDiscount(uint256 _value) returns()
func (_CreditManagerMockForFilter *CreditManagerMockForFilterTransactorSession) SetLiquidationDiscount(_value *big.Int) (*types.Transaction, error) {
	return _CreditManagerMockForFilter.Contract.SetLiquidationDiscount(&_CreditManagerMockForFilter.TransactOpts, _value)
}

// SetMaxLeverageFactor is a paid mutator transaction binding the contract method 0x7629da6f.
//
// Solidity: function setMaxLeverageFactor(uint256 _value) returns()
func (_CreditManagerMockForFilter *CreditManagerMockForFilterTransactor) SetMaxLeverageFactor(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _CreditManagerMockForFilter.contract.Transact(opts, "setMaxLeverageFactor", _value)
}

// SetMaxLeverageFactor is a paid mutator transaction binding the contract method 0x7629da6f.
//
// Solidity: function setMaxLeverageFactor(uint256 _value) returns()
func (_CreditManagerMockForFilter *CreditManagerMockForFilterSession) SetMaxLeverageFactor(_value *big.Int) (*types.Transaction, error) {
	return _CreditManagerMockForFilter.Contract.SetMaxLeverageFactor(&_CreditManagerMockForFilter.TransactOpts, _value)
}

// SetMaxLeverageFactor is a paid mutator transaction binding the contract method 0x7629da6f.
//
// Solidity: function setMaxLeverageFactor(uint256 _value) returns()
func (_CreditManagerMockForFilter *CreditManagerMockForFilterTransactorSession) SetMaxLeverageFactor(_value *big.Int) (*types.Transaction, error) {
	return _CreditManagerMockForFilter.Contract.SetMaxLeverageFactor(&_CreditManagerMockForFilter.TransactOpts, _value)
}

// UpdateUnderlyingTokenLiquidationThreshold is a paid mutator transaction binding the contract method 0x40631828.
//
// Solidity: function updateUnderlyingTokenLiquidationThreshold() returns()
func (_CreditManagerMockForFilter *CreditManagerMockForFilterTransactor) UpdateUnderlyingTokenLiquidationThreshold(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditManagerMockForFilter.contract.Transact(opts, "updateUnderlyingTokenLiquidationThreshold")
}

// UpdateUnderlyingTokenLiquidationThreshold is a paid mutator transaction binding the contract method 0x40631828.
//
// Solidity: function updateUnderlyingTokenLiquidationThreshold() returns()
func (_CreditManagerMockForFilter *CreditManagerMockForFilterSession) UpdateUnderlyingTokenLiquidationThreshold() (*types.Transaction, error) {
	return _CreditManagerMockForFilter.Contract.UpdateUnderlyingTokenLiquidationThreshold(&_CreditManagerMockForFilter.TransactOpts)
}

// UpdateUnderlyingTokenLiquidationThreshold is a paid mutator transaction binding the contract method 0x40631828.
//
// Solidity: function updateUnderlyingTokenLiquidationThreshold() returns()
func (_CreditManagerMockForFilter *CreditManagerMockForFilterTransactorSession) UpdateUnderlyingTokenLiquidationThreshold() (*types.Transaction, error) {
	return _CreditManagerMockForFilter.Contract.UpdateUnderlyingTokenLiquidationThreshold(&_CreditManagerMockForFilter.TransactOpts)
}
