// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package creditFilterMock

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

// CreditFilterMockMetaData contains all meta data concerning the CreditFilterMock contract.
var CreditFilterMockMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addressProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_underlyingToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"protocol\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"adapter\",\"type\":\"address\"}],\"name\":\"ContractAllowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"protocol\",\"type\":\"address\"}],\"name\":\"ContractForbidden\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chiThreshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fastCheckDelay\",\"type\":\"uint256\"}],\"name\":\"NewFastCheckParameters\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidityThreshold\",\"type\":\"uint256\"}],\"name\":\"TokenAllowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"state\",\"type\":\"bool\"}],\"name\":\"TransferAccountAllowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pugin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"state\",\"type\":\"bool\"}],\"name\":\"TransferPluginAllowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_allowedTokensMap\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addressProvider\",\"outputs\":[{\"internalType\":\"contractAddressProvider\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"targetContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"adapter\",\"type\":\"address\"}],\"name\":\"allowContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"plugin\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"state\",\"type\":\"bool\"}],\"name\":\"allowPlugin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidationThreshold\",\"type\":\"uint256\"}],\"name\":\"allowToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"allowanceForAccountTransfers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowedAdapters\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"allowedContracts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allowedContractsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowedPlugins\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allowedTokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allowedTokensCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"state\",\"type\":\"bool\"}],\"name\":\"approveAccountTransfers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"}],\"name\":\"calcCreditAccountAccruedInterest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"}],\"name\":\"calcCreditAccountHealthFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"percentage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"times\",\"type\":\"uint256\"}],\"name\":\"calcMaxPossibleDrop\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"}],\"name\":\"calcThresholdWeightedValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"}],\"name\":\"calcTotalValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"checkAndEnableToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"checkCollateralChange\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"amountIn\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amountOut\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"tokenIn\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"tokenOut\",\"type\":\"address[]\"}],\"name\":\"checkMultiTokenCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chiThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_creditManager\",\"type\":\"address\"}],\"name\":\"connectCreditManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"contractToAdapter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creditManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"enabledTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"fastCheckCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"targetContract\",\"type\":\"address\"}],\"name\":\"forbidContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"forbidToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getCreditAccountTokenById\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tvw\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hfCheckInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"}],\"name\":\"initEnabledTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"isTokenAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"liquidationThresholds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolService\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"revertIfAccountTransferIsNotAllowed\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minHealthFactor\",\"type\":\"uint256\"}],\"name\":\"revertIfCantIncreaseBorrowing\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"revertIfTokenNotAllowed\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenMask\",\"type\":\"uint256\"}],\"name\":\"setEnabledTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"blockNum\",\"type\":\"uint256\"}],\"name\":\"setFastCheckBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chiThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_hfCheckInterval\",\"type\":\"uint256\"}],\"name\":\"setFastCheckParameters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenMasksMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlyingToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateUnderlyingTokenLiquidationThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wethAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// CreditFilterMockABI is the input ABI used to generate the binding from.
// Deprecated: Use CreditFilterMockMetaData.ABI instead.
var CreditFilterMockABI = CreditFilterMockMetaData.ABI

// CreditFilterMock is an auto generated Go binding around an Ethereum contract.
type CreditFilterMock struct {
	CreditFilterMockCaller     // Read-only binding to the contract
	CreditFilterMockTransactor // Write-only binding to the contract
	CreditFilterMockFilterer   // Log filterer for contract events
}

// CreditFilterMockCaller is an auto generated read-only Go binding around an Ethereum contract.
type CreditFilterMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditFilterMockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CreditFilterMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditFilterMockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CreditFilterMockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditFilterMockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CreditFilterMockSession struct {
	Contract     *CreditFilterMock // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CreditFilterMockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CreditFilterMockCallerSession struct {
	Contract *CreditFilterMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// CreditFilterMockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CreditFilterMockTransactorSession struct {
	Contract     *CreditFilterMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// CreditFilterMockRaw is an auto generated low-level Go binding around an Ethereum contract.
type CreditFilterMockRaw struct {
	Contract *CreditFilterMock // Generic contract binding to access the raw methods on
}

// CreditFilterMockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CreditFilterMockCallerRaw struct {
	Contract *CreditFilterMockCaller // Generic read-only contract binding to access the raw methods on
}

// CreditFilterMockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CreditFilterMockTransactorRaw struct {
	Contract *CreditFilterMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCreditFilterMock creates a new instance of CreditFilterMock, bound to a specific deployed contract.
func NewCreditFilterMock(address common.Address, backend bind.ContractBackend) (*CreditFilterMock, error) {
	contract, err := bindCreditFilterMock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CreditFilterMock{CreditFilterMockCaller: CreditFilterMockCaller{contract: contract}, CreditFilterMockTransactor: CreditFilterMockTransactor{contract: contract}, CreditFilterMockFilterer: CreditFilterMockFilterer{contract: contract}}, nil
}

// NewCreditFilterMockCaller creates a new read-only instance of CreditFilterMock, bound to a specific deployed contract.
func NewCreditFilterMockCaller(address common.Address, caller bind.ContractCaller) (*CreditFilterMockCaller, error) {
	contract, err := bindCreditFilterMock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CreditFilterMockCaller{contract: contract}, nil
}

// NewCreditFilterMockTransactor creates a new write-only instance of CreditFilterMock, bound to a specific deployed contract.
func NewCreditFilterMockTransactor(address common.Address, transactor bind.ContractTransactor) (*CreditFilterMockTransactor, error) {
	contract, err := bindCreditFilterMock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CreditFilterMockTransactor{contract: contract}, nil
}

// NewCreditFilterMockFilterer creates a new log filterer instance of CreditFilterMock, bound to a specific deployed contract.
func NewCreditFilterMockFilterer(address common.Address, filterer bind.ContractFilterer) (*CreditFilterMockFilterer, error) {
	contract, err := bindCreditFilterMock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CreditFilterMockFilterer{contract: contract}, nil
}

// bindCreditFilterMock binds a generic wrapper to an already deployed contract.
func bindCreditFilterMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CreditFilterMockABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CreditFilterMock *CreditFilterMockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CreditFilterMock.Contract.CreditFilterMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CreditFilterMock *CreditFilterMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditFilterMock.Contract.CreditFilterMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CreditFilterMock *CreditFilterMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CreditFilterMock.Contract.CreditFilterMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CreditFilterMock *CreditFilterMockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CreditFilterMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CreditFilterMock *CreditFilterMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditFilterMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CreditFilterMock *CreditFilterMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CreditFilterMock.Contract.contract.Transact(opts, method, params...)
}

// AllowedTokensMap is a free data retrieval call binding the contract method 0xb2d0d86b.
//
// Solidity: function _allowedTokensMap(address ) view returns(bool)
func (_CreditFilterMock *CreditFilterMockCaller) AllowedTokensMap(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _CreditFilterMock.contract.Call(opts, &out, "_allowedTokensMap", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedTokensMap is a free data retrieval call binding the contract method 0xb2d0d86b.
//
// Solidity: function _allowedTokensMap(address ) view returns(bool)
func (_CreditFilterMock *CreditFilterMockSession) AllowedTokensMap(arg0 common.Address) (bool, error) {
	return _CreditFilterMock.Contract.AllowedTokensMap(&_CreditFilterMock.CallOpts, arg0)
}

// AllowedTokensMap is a free data retrieval call binding the contract method 0xb2d0d86b.
//
// Solidity: function _allowedTokensMap(address ) view returns(bool)
func (_CreditFilterMock *CreditFilterMockCallerSession) AllowedTokensMap(arg0 common.Address) (bool, error) {
	return _CreditFilterMock.Contract.AllowedTokensMap(&_CreditFilterMock.CallOpts, arg0)
}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_CreditFilterMock *CreditFilterMockCaller) AddressProvider(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditFilterMock.contract.Call(opts, &out, "addressProvider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_CreditFilterMock *CreditFilterMockSession) AddressProvider() (common.Address, error) {
	return _CreditFilterMock.Contract.AddressProvider(&_CreditFilterMock.CallOpts)
}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_CreditFilterMock *CreditFilterMockCallerSession) AddressProvider() (common.Address, error) {
	return _CreditFilterMock.Contract.AddressProvider(&_CreditFilterMock.CallOpts)
}

// AllowanceForAccountTransfers is a free data retrieval call binding the contract method 0x5a29be45.
//
// Solidity: function allowanceForAccountTransfers(address from, address to) view returns(bool)
func (_CreditFilterMock *CreditFilterMockCaller) AllowanceForAccountTransfers(opts *bind.CallOpts, from common.Address, to common.Address) (bool, error) {
	var out []interface{}
	err := _CreditFilterMock.contract.Call(opts, &out, "allowanceForAccountTransfers", from, to)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowanceForAccountTransfers is a free data retrieval call binding the contract method 0x5a29be45.
//
// Solidity: function allowanceForAccountTransfers(address from, address to) view returns(bool)
func (_CreditFilterMock *CreditFilterMockSession) AllowanceForAccountTransfers(from common.Address, to common.Address) (bool, error) {
	return _CreditFilterMock.Contract.AllowanceForAccountTransfers(&_CreditFilterMock.CallOpts, from, to)
}

// AllowanceForAccountTransfers is a free data retrieval call binding the contract method 0x5a29be45.
//
// Solidity: function allowanceForAccountTransfers(address from, address to) view returns(bool)
func (_CreditFilterMock *CreditFilterMockCallerSession) AllowanceForAccountTransfers(from common.Address, to common.Address) (bool, error) {
	return _CreditFilterMock.Contract.AllowanceForAccountTransfers(&_CreditFilterMock.CallOpts, from, to)
}

// AllowedAdapters is a free data retrieval call binding the contract method 0x3bdfe4f5.
//
// Solidity: function allowedAdapters(address ) view returns(bool)
func (_CreditFilterMock *CreditFilterMockCaller) AllowedAdapters(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _CreditFilterMock.contract.Call(opts, &out, "allowedAdapters", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedAdapters is a free data retrieval call binding the contract method 0x3bdfe4f5.
//
// Solidity: function allowedAdapters(address ) view returns(bool)
func (_CreditFilterMock *CreditFilterMockSession) AllowedAdapters(arg0 common.Address) (bool, error) {
	return _CreditFilterMock.Contract.AllowedAdapters(&_CreditFilterMock.CallOpts, arg0)
}

// AllowedAdapters is a free data retrieval call binding the contract method 0x3bdfe4f5.
//
// Solidity: function allowedAdapters(address ) view returns(bool)
func (_CreditFilterMock *CreditFilterMockCallerSession) AllowedAdapters(arg0 common.Address) (bool, error) {
	return _CreditFilterMock.Contract.AllowedAdapters(&_CreditFilterMock.CallOpts, arg0)
}

// AllowedContracts is a free data retrieval call binding the contract method 0x5094cb4f.
//
// Solidity: function allowedContracts(uint256 i) view returns(address)
func (_CreditFilterMock *CreditFilterMockCaller) AllowedContracts(opts *bind.CallOpts, i *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CreditFilterMock.contract.Call(opts, &out, "allowedContracts", i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllowedContracts is a free data retrieval call binding the contract method 0x5094cb4f.
//
// Solidity: function allowedContracts(uint256 i) view returns(address)
func (_CreditFilterMock *CreditFilterMockSession) AllowedContracts(i *big.Int) (common.Address, error) {
	return _CreditFilterMock.Contract.AllowedContracts(&_CreditFilterMock.CallOpts, i)
}

// AllowedContracts is a free data retrieval call binding the contract method 0x5094cb4f.
//
// Solidity: function allowedContracts(uint256 i) view returns(address)
func (_CreditFilterMock *CreditFilterMockCallerSession) AllowedContracts(i *big.Int) (common.Address, error) {
	return _CreditFilterMock.Contract.AllowedContracts(&_CreditFilterMock.CallOpts, i)
}

// AllowedContractsCount is a free data retrieval call binding the contract method 0x50e036ff.
//
// Solidity: function allowedContractsCount() view returns(uint256)
func (_CreditFilterMock *CreditFilterMockCaller) AllowedContractsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CreditFilterMock.contract.Call(opts, &out, "allowedContractsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllowedContractsCount is a free data retrieval call binding the contract method 0x50e036ff.
//
// Solidity: function allowedContractsCount() view returns(uint256)
func (_CreditFilterMock *CreditFilterMockSession) AllowedContractsCount() (*big.Int, error) {
	return _CreditFilterMock.Contract.AllowedContractsCount(&_CreditFilterMock.CallOpts)
}

// AllowedContractsCount is a free data retrieval call binding the contract method 0x50e036ff.
//
// Solidity: function allowedContractsCount() view returns(uint256)
func (_CreditFilterMock *CreditFilterMockCallerSession) AllowedContractsCount() (*big.Int, error) {
	return _CreditFilterMock.Contract.AllowedContractsCount(&_CreditFilterMock.CallOpts)
}

// AllowedPlugins is a free data retrieval call binding the contract method 0x5f598edd.
//
// Solidity: function allowedPlugins(address ) view returns(bool)
func (_CreditFilterMock *CreditFilterMockCaller) AllowedPlugins(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _CreditFilterMock.contract.Call(opts, &out, "allowedPlugins", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedPlugins is a free data retrieval call binding the contract method 0x5f598edd.
//
// Solidity: function allowedPlugins(address ) view returns(bool)
func (_CreditFilterMock *CreditFilterMockSession) AllowedPlugins(arg0 common.Address) (bool, error) {
	return _CreditFilterMock.Contract.AllowedPlugins(&_CreditFilterMock.CallOpts, arg0)
}

// AllowedPlugins is a free data retrieval call binding the contract method 0x5f598edd.
//
// Solidity: function allowedPlugins(address ) view returns(bool)
func (_CreditFilterMock *CreditFilterMockCallerSession) AllowedPlugins(arg0 common.Address) (bool, error) {
	return _CreditFilterMock.Contract.AllowedPlugins(&_CreditFilterMock.CallOpts, arg0)
}

// AllowedTokens is a free data retrieval call binding the contract method 0x5e5f2e26.
//
// Solidity: function allowedTokens(uint256 ) view returns(address)
func (_CreditFilterMock *CreditFilterMockCaller) AllowedTokens(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CreditFilterMock.contract.Call(opts, &out, "allowedTokens", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllowedTokens is a free data retrieval call binding the contract method 0x5e5f2e26.
//
// Solidity: function allowedTokens(uint256 ) view returns(address)
func (_CreditFilterMock *CreditFilterMockSession) AllowedTokens(arg0 *big.Int) (common.Address, error) {
	return _CreditFilterMock.Contract.AllowedTokens(&_CreditFilterMock.CallOpts, arg0)
}

// AllowedTokens is a free data retrieval call binding the contract method 0x5e5f2e26.
//
// Solidity: function allowedTokens(uint256 ) view returns(address)
func (_CreditFilterMock *CreditFilterMockCallerSession) AllowedTokens(arg0 *big.Int) (common.Address, error) {
	return _CreditFilterMock.Contract.AllowedTokens(&_CreditFilterMock.CallOpts, arg0)
}

// AllowedTokensCount is a free data retrieval call binding the contract method 0x20a05ff7.
//
// Solidity: function allowedTokensCount() view returns(uint256)
func (_CreditFilterMock *CreditFilterMockCaller) AllowedTokensCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CreditFilterMock.contract.Call(opts, &out, "allowedTokensCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllowedTokensCount is a free data retrieval call binding the contract method 0x20a05ff7.
//
// Solidity: function allowedTokensCount() view returns(uint256)
func (_CreditFilterMock *CreditFilterMockSession) AllowedTokensCount() (*big.Int, error) {
	return _CreditFilterMock.Contract.AllowedTokensCount(&_CreditFilterMock.CallOpts)
}

// AllowedTokensCount is a free data retrieval call binding the contract method 0x20a05ff7.
//
// Solidity: function allowedTokensCount() view returns(uint256)
func (_CreditFilterMock *CreditFilterMockCallerSession) AllowedTokensCount() (*big.Int, error) {
	return _CreditFilterMock.Contract.AllowedTokensCount(&_CreditFilterMock.CallOpts)
}

// CalcCreditAccountAccruedInterest is a free data retrieval call binding the contract method 0x3192195c.
//
// Solidity: function calcCreditAccountAccruedInterest(address creditAccount) view returns(uint256)
func (_CreditFilterMock *CreditFilterMockCaller) CalcCreditAccountAccruedInterest(opts *bind.CallOpts, creditAccount common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CreditFilterMock.contract.Call(opts, &out, "calcCreditAccountAccruedInterest", creditAccount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcCreditAccountAccruedInterest is a free data retrieval call binding the contract method 0x3192195c.
//
// Solidity: function calcCreditAccountAccruedInterest(address creditAccount) view returns(uint256)
func (_CreditFilterMock *CreditFilterMockSession) CalcCreditAccountAccruedInterest(creditAccount common.Address) (*big.Int, error) {
	return _CreditFilterMock.Contract.CalcCreditAccountAccruedInterest(&_CreditFilterMock.CallOpts, creditAccount)
}

// CalcCreditAccountAccruedInterest is a free data retrieval call binding the contract method 0x3192195c.
//
// Solidity: function calcCreditAccountAccruedInterest(address creditAccount) view returns(uint256)
func (_CreditFilterMock *CreditFilterMockCallerSession) CalcCreditAccountAccruedInterest(creditAccount common.Address) (*big.Int, error) {
	return _CreditFilterMock.Contract.CalcCreditAccountAccruedInterest(&_CreditFilterMock.CallOpts, creditAccount)
}

// CalcCreditAccountHealthFactor is a free data retrieval call binding the contract method 0xdfd59465.
//
// Solidity: function calcCreditAccountHealthFactor(address creditAccount) view returns(uint256)
func (_CreditFilterMock *CreditFilterMockCaller) CalcCreditAccountHealthFactor(opts *bind.CallOpts, creditAccount common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CreditFilterMock.contract.Call(opts, &out, "calcCreditAccountHealthFactor", creditAccount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcCreditAccountHealthFactor is a free data retrieval call binding the contract method 0xdfd59465.
//
// Solidity: function calcCreditAccountHealthFactor(address creditAccount) view returns(uint256)
func (_CreditFilterMock *CreditFilterMockSession) CalcCreditAccountHealthFactor(creditAccount common.Address) (*big.Int, error) {
	return _CreditFilterMock.Contract.CalcCreditAccountHealthFactor(&_CreditFilterMock.CallOpts, creditAccount)
}

// CalcCreditAccountHealthFactor is a free data retrieval call binding the contract method 0xdfd59465.
//
// Solidity: function calcCreditAccountHealthFactor(address creditAccount) view returns(uint256)
func (_CreditFilterMock *CreditFilterMockCallerSession) CalcCreditAccountHealthFactor(creditAccount common.Address) (*big.Int, error) {
	return _CreditFilterMock.Contract.CalcCreditAccountHealthFactor(&_CreditFilterMock.CallOpts, creditAccount)
}

// CalcMaxPossibleDrop is a free data retrieval call binding the contract method 0xb3c61943.
//
// Solidity: function calcMaxPossibleDrop(uint256 percentage, uint256 times) pure returns(uint256 value)
func (_CreditFilterMock *CreditFilterMockCaller) CalcMaxPossibleDrop(opts *bind.CallOpts, percentage *big.Int, times *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CreditFilterMock.contract.Call(opts, &out, "calcMaxPossibleDrop", percentage, times)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcMaxPossibleDrop is a free data retrieval call binding the contract method 0xb3c61943.
//
// Solidity: function calcMaxPossibleDrop(uint256 percentage, uint256 times) pure returns(uint256 value)
func (_CreditFilterMock *CreditFilterMockSession) CalcMaxPossibleDrop(percentage *big.Int, times *big.Int) (*big.Int, error) {
	return _CreditFilterMock.Contract.CalcMaxPossibleDrop(&_CreditFilterMock.CallOpts, percentage, times)
}

// CalcMaxPossibleDrop is a free data retrieval call binding the contract method 0xb3c61943.
//
// Solidity: function calcMaxPossibleDrop(uint256 percentage, uint256 times) pure returns(uint256 value)
func (_CreditFilterMock *CreditFilterMockCallerSession) CalcMaxPossibleDrop(percentage *big.Int, times *big.Int) (*big.Int, error) {
	return _CreditFilterMock.Contract.CalcMaxPossibleDrop(&_CreditFilterMock.CallOpts, percentage, times)
}

// CalcThresholdWeightedValue is a free data retrieval call binding the contract method 0x90b1300a.
//
// Solidity: function calcThresholdWeightedValue(address creditAccount) view returns(uint256 total)
func (_CreditFilterMock *CreditFilterMockCaller) CalcThresholdWeightedValue(opts *bind.CallOpts, creditAccount common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CreditFilterMock.contract.Call(opts, &out, "calcThresholdWeightedValue", creditAccount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcThresholdWeightedValue is a free data retrieval call binding the contract method 0x90b1300a.
//
// Solidity: function calcThresholdWeightedValue(address creditAccount) view returns(uint256 total)
func (_CreditFilterMock *CreditFilterMockSession) CalcThresholdWeightedValue(creditAccount common.Address) (*big.Int, error) {
	return _CreditFilterMock.Contract.CalcThresholdWeightedValue(&_CreditFilterMock.CallOpts, creditAccount)
}

// CalcThresholdWeightedValue is a free data retrieval call binding the contract method 0x90b1300a.
//
// Solidity: function calcThresholdWeightedValue(address creditAccount) view returns(uint256 total)
func (_CreditFilterMock *CreditFilterMockCallerSession) CalcThresholdWeightedValue(creditAccount common.Address) (*big.Int, error) {
	return _CreditFilterMock.Contract.CalcThresholdWeightedValue(&_CreditFilterMock.CallOpts, creditAccount)
}

// CalcTotalValue is a free data retrieval call binding the contract method 0xc7de38a6.
//
// Solidity: function calcTotalValue(address creditAccount) view returns(uint256 total)
func (_CreditFilterMock *CreditFilterMockCaller) CalcTotalValue(opts *bind.CallOpts, creditAccount common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CreditFilterMock.contract.Call(opts, &out, "calcTotalValue", creditAccount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcTotalValue is a free data retrieval call binding the contract method 0xc7de38a6.
//
// Solidity: function calcTotalValue(address creditAccount) view returns(uint256 total)
func (_CreditFilterMock *CreditFilterMockSession) CalcTotalValue(creditAccount common.Address) (*big.Int, error) {
	return _CreditFilterMock.Contract.CalcTotalValue(&_CreditFilterMock.CallOpts, creditAccount)
}

// CalcTotalValue is a free data retrieval call binding the contract method 0xc7de38a6.
//
// Solidity: function calcTotalValue(address creditAccount) view returns(uint256 total)
func (_CreditFilterMock *CreditFilterMockCallerSession) CalcTotalValue(creditAccount common.Address) (*big.Int, error) {
	return _CreditFilterMock.Contract.CalcTotalValue(&_CreditFilterMock.CallOpts, creditAccount)
}

// ChiThreshold is a free data retrieval call binding the contract method 0x47dedfc9.
//
// Solidity: function chiThreshold() view returns(uint256)
func (_CreditFilterMock *CreditFilterMockCaller) ChiThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CreditFilterMock.contract.Call(opts, &out, "chiThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChiThreshold is a free data retrieval call binding the contract method 0x47dedfc9.
//
// Solidity: function chiThreshold() view returns(uint256)
func (_CreditFilterMock *CreditFilterMockSession) ChiThreshold() (*big.Int, error) {
	return _CreditFilterMock.Contract.ChiThreshold(&_CreditFilterMock.CallOpts)
}

// ChiThreshold is a free data retrieval call binding the contract method 0x47dedfc9.
//
// Solidity: function chiThreshold() view returns(uint256)
func (_CreditFilterMock *CreditFilterMockCallerSession) ChiThreshold() (*big.Int, error) {
	return _CreditFilterMock.Contract.ChiThreshold(&_CreditFilterMock.CallOpts)
}

// ContractToAdapter is a free data retrieval call binding the contract method 0xfdd57645.
//
// Solidity: function contractToAdapter(address ) view returns(address)
func (_CreditFilterMock *CreditFilterMockCaller) ContractToAdapter(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _CreditFilterMock.contract.Call(opts, &out, "contractToAdapter", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ContractToAdapter is a free data retrieval call binding the contract method 0xfdd57645.
//
// Solidity: function contractToAdapter(address ) view returns(address)
func (_CreditFilterMock *CreditFilterMockSession) ContractToAdapter(arg0 common.Address) (common.Address, error) {
	return _CreditFilterMock.Contract.ContractToAdapter(&_CreditFilterMock.CallOpts, arg0)
}

// ContractToAdapter is a free data retrieval call binding the contract method 0xfdd57645.
//
// Solidity: function contractToAdapter(address ) view returns(address)
func (_CreditFilterMock *CreditFilterMockCallerSession) ContractToAdapter(arg0 common.Address) (common.Address, error) {
	return _CreditFilterMock.Contract.ContractToAdapter(&_CreditFilterMock.CallOpts, arg0)
}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_CreditFilterMock *CreditFilterMockCaller) CreditManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditFilterMock.contract.Call(opts, &out, "creditManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_CreditFilterMock *CreditFilterMockSession) CreditManager() (common.Address, error) {
	return _CreditFilterMock.Contract.CreditManager(&_CreditFilterMock.CallOpts)
}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_CreditFilterMock *CreditFilterMockCallerSession) CreditManager() (common.Address, error) {
	return _CreditFilterMock.Contract.CreditManager(&_CreditFilterMock.CallOpts)
}

// EnabledTokens is a free data retrieval call binding the contract method 0xb451cecc.
//
// Solidity: function enabledTokens(address ) view returns(uint256)
func (_CreditFilterMock *CreditFilterMockCaller) EnabledTokens(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CreditFilterMock.contract.Call(opts, &out, "enabledTokens", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EnabledTokens is a free data retrieval call binding the contract method 0xb451cecc.
//
// Solidity: function enabledTokens(address ) view returns(uint256)
func (_CreditFilterMock *CreditFilterMockSession) EnabledTokens(arg0 common.Address) (*big.Int, error) {
	return _CreditFilterMock.Contract.EnabledTokens(&_CreditFilterMock.CallOpts, arg0)
}

// EnabledTokens is a free data retrieval call binding the contract method 0xb451cecc.
//
// Solidity: function enabledTokens(address ) view returns(uint256)
func (_CreditFilterMock *CreditFilterMockCallerSession) EnabledTokens(arg0 common.Address) (*big.Int, error) {
	return _CreditFilterMock.Contract.EnabledTokens(&_CreditFilterMock.CallOpts, arg0)
}

// FastCheckCounter is a free data retrieval call binding the contract method 0x4cba294a.
//
// Solidity: function fastCheckCounter(address ) view returns(uint256)
func (_CreditFilterMock *CreditFilterMockCaller) FastCheckCounter(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CreditFilterMock.contract.Call(opts, &out, "fastCheckCounter", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FastCheckCounter is a free data retrieval call binding the contract method 0x4cba294a.
//
// Solidity: function fastCheckCounter(address ) view returns(uint256)
func (_CreditFilterMock *CreditFilterMockSession) FastCheckCounter(arg0 common.Address) (*big.Int, error) {
	return _CreditFilterMock.Contract.FastCheckCounter(&_CreditFilterMock.CallOpts, arg0)
}

// FastCheckCounter is a free data retrieval call binding the contract method 0x4cba294a.
//
// Solidity: function fastCheckCounter(address ) view returns(uint256)
func (_CreditFilterMock *CreditFilterMockCallerSession) FastCheckCounter(arg0 common.Address) (*big.Int, error) {
	return _CreditFilterMock.Contract.FastCheckCounter(&_CreditFilterMock.CallOpts, arg0)
}

// GetCreditAccountTokenById is a free data retrieval call binding the contract method 0xaf0a6502.
//
// Solidity: function getCreditAccountTokenById(address creditAccount, uint256 id) view returns(address token, uint256 balance, uint256 tv, uint256 tvw)
func (_CreditFilterMock *CreditFilterMockCaller) GetCreditAccountTokenById(opts *bind.CallOpts, creditAccount common.Address, id *big.Int) (struct {
	Token   common.Address
	Balance *big.Int
	Tv      *big.Int
	Tvw     *big.Int
}, error) {
	var out []interface{}
	err := _CreditFilterMock.contract.Call(opts, &out, "getCreditAccountTokenById", creditAccount, id)

	outstruct := new(struct {
		Token   common.Address
		Balance *big.Int
		Tv      *big.Int
		Tvw     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Token = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Balance = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Tv = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Tvw = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetCreditAccountTokenById is a free data retrieval call binding the contract method 0xaf0a6502.
//
// Solidity: function getCreditAccountTokenById(address creditAccount, uint256 id) view returns(address token, uint256 balance, uint256 tv, uint256 tvw)
func (_CreditFilterMock *CreditFilterMockSession) GetCreditAccountTokenById(creditAccount common.Address, id *big.Int) (struct {
	Token   common.Address
	Balance *big.Int
	Tv      *big.Int
	Tvw     *big.Int
}, error) {
	return _CreditFilterMock.Contract.GetCreditAccountTokenById(&_CreditFilterMock.CallOpts, creditAccount, id)
}

// GetCreditAccountTokenById is a free data retrieval call binding the contract method 0xaf0a6502.
//
// Solidity: function getCreditAccountTokenById(address creditAccount, uint256 id) view returns(address token, uint256 balance, uint256 tv, uint256 tvw)
func (_CreditFilterMock *CreditFilterMockCallerSession) GetCreditAccountTokenById(creditAccount common.Address, id *big.Int) (struct {
	Token   common.Address
	Balance *big.Int
	Tv      *big.Int
	Tvw     *big.Int
}, error) {
	return _CreditFilterMock.Contract.GetCreditAccountTokenById(&_CreditFilterMock.CallOpts, creditAccount, id)
}

// HfCheckInterval is a free data retrieval call binding the contract method 0xe6dee2cc.
//
// Solidity: function hfCheckInterval() view returns(uint256)
func (_CreditFilterMock *CreditFilterMockCaller) HfCheckInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CreditFilterMock.contract.Call(opts, &out, "hfCheckInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HfCheckInterval is a free data retrieval call binding the contract method 0xe6dee2cc.
//
// Solidity: function hfCheckInterval() view returns(uint256)
func (_CreditFilterMock *CreditFilterMockSession) HfCheckInterval() (*big.Int, error) {
	return _CreditFilterMock.Contract.HfCheckInterval(&_CreditFilterMock.CallOpts)
}

// HfCheckInterval is a free data retrieval call binding the contract method 0xe6dee2cc.
//
// Solidity: function hfCheckInterval() view returns(uint256)
func (_CreditFilterMock *CreditFilterMockCallerSession) HfCheckInterval() (*big.Int, error) {
	return _CreditFilterMock.Contract.HfCheckInterval(&_CreditFilterMock.CallOpts)
}

// IsTokenAllowed is a free data retrieval call binding the contract method 0xf9eaee0d.
//
// Solidity: function isTokenAllowed(address token) view returns(bool)
func (_CreditFilterMock *CreditFilterMockCaller) IsTokenAllowed(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _CreditFilterMock.contract.Call(opts, &out, "isTokenAllowed", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTokenAllowed is a free data retrieval call binding the contract method 0xf9eaee0d.
//
// Solidity: function isTokenAllowed(address token) view returns(bool)
func (_CreditFilterMock *CreditFilterMockSession) IsTokenAllowed(token common.Address) (bool, error) {
	return _CreditFilterMock.Contract.IsTokenAllowed(&_CreditFilterMock.CallOpts, token)
}

// IsTokenAllowed is a free data retrieval call binding the contract method 0xf9eaee0d.
//
// Solidity: function isTokenAllowed(address token) view returns(bool)
func (_CreditFilterMock *CreditFilterMockCallerSession) IsTokenAllowed(token common.Address) (bool, error) {
	return _CreditFilterMock.Contract.IsTokenAllowed(&_CreditFilterMock.CallOpts, token)
}

// LiquidationThresholds is a free data retrieval call binding the contract method 0x78327438.
//
// Solidity: function liquidationThresholds(address ) view returns(uint256)
func (_CreditFilterMock *CreditFilterMockCaller) LiquidationThresholds(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CreditFilterMock.contract.Call(opts, &out, "liquidationThresholds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LiquidationThresholds is a free data retrieval call binding the contract method 0x78327438.
//
// Solidity: function liquidationThresholds(address ) view returns(uint256)
func (_CreditFilterMock *CreditFilterMockSession) LiquidationThresholds(arg0 common.Address) (*big.Int, error) {
	return _CreditFilterMock.Contract.LiquidationThresholds(&_CreditFilterMock.CallOpts, arg0)
}

// LiquidationThresholds is a free data retrieval call binding the contract method 0x78327438.
//
// Solidity: function liquidationThresholds(address ) view returns(uint256)
func (_CreditFilterMock *CreditFilterMockCallerSession) LiquidationThresholds(arg0 common.Address) (*big.Int, error) {
	return _CreditFilterMock.Contract.LiquidationThresholds(&_CreditFilterMock.CallOpts, arg0)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CreditFilterMock *CreditFilterMockCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CreditFilterMock.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CreditFilterMock *CreditFilterMockSession) Paused() (bool, error) {
	return _CreditFilterMock.Contract.Paused(&_CreditFilterMock.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CreditFilterMock *CreditFilterMockCallerSession) Paused() (bool, error) {
	return _CreditFilterMock.Contract.Paused(&_CreditFilterMock.CallOpts)
}

// PoolService is a free data retrieval call binding the contract method 0x570a7af2.
//
// Solidity: function poolService() view returns(address)
func (_CreditFilterMock *CreditFilterMockCaller) PoolService(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditFilterMock.contract.Call(opts, &out, "poolService")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolService is a free data retrieval call binding the contract method 0x570a7af2.
//
// Solidity: function poolService() view returns(address)
func (_CreditFilterMock *CreditFilterMockSession) PoolService() (common.Address, error) {
	return _CreditFilterMock.Contract.PoolService(&_CreditFilterMock.CallOpts)
}

// PoolService is a free data retrieval call binding the contract method 0x570a7af2.
//
// Solidity: function poolService() view returns(address)
func (_CreditFilterMock *CreditFilterMockCallerSession) PoolService() (common.Address, error) {
	return _CreditFilterMock.Contract.PoolService(&_CreditFilterMock.CallOpts)
}

// PriceOracle is a free data retrieval call binding the contract method 0x2630c12f.
//
// Solidity: function priceOracle() view returns(address)
func (_CreditFilterMock *CreditFilterMockCaller) PriceOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditFilterMock.contract.Call(opts, &out, "priceOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PriceOracle is a free data retrieval call binding the contract method 0x2630c12f.
//
// Solidity: function priceOracle() view returns(address)
func (_CreditFilterMock *CreditFilterMockSession) PriceOracle() (common.Address, error) {
	return _CreditFilterMock.Contract.PriceOracle(&_CreditFilterMock.CallOpts)
}

// PriceOracle is a free data retrieval call binding the contract method 0x2630c12f.
//
// Solidity: function priceOracle() view returns(address)
func (_CreditFilterMock *CreditFilterMockCallerSession) PriceOracle() (common.Address, error) {
	return _CreditFilterMock.Contract.PriceOracle(&_CreditFilterMock.CallOpts)
}

// RevertIfAccountTransferIsNotAllowed is a free data retrieval call binding the contract method 0x3b00ae70.
//
// Solidity: function revertIfAccountTransferIsNotAllowed(address owner, address newOwner) view returns()
func (_CreditFilterMock *CreditFilterMockCaller) RevertIfAccountTransferIsNotAllowed(opts *bind.CallOpts, owner common.Address, newOwner common.Address) error {
	var out []interface{}
	err := _CreditFilterMock.contract.Call(opts, &out, "revertIfAccountTransferIsNotAllowed", owner, newOwner)

	if err != nil {
		return err
	}

	return err

}

// RevertIfAccountTransferIsNotAllowed is a free data retrieval call binding the contract method 0x3b00ae70.
//
// Solidity: function revertIfAccountTransferIsNotAllowed(address owner, address newOwner) view returns()
func (_CreditFilterMock *CreditFilterMockSession) RevertIfAccountTransferIsNotAllowed(owner common.Address, newOwner common.Address) error {
	return _CreditFilterMock.Contract.RevertIfAccountTransferIsNotAllowed(&_CreditFilterMock.CallOpts, owner, newOwner)
}

// RevertIfAccountTransferIsNotAllowed is a free data retrieval call binding the contract method 0x3b00ae70.
//
// Solidity: function revertIfAccountTransferIsNotAllowed(address owner, address newOwner) view returns()
func (_CreditFilterMock *CreditFilterMockCallerSession) RevertIfAccountTransferIsNotAllowed(owner common.Address, newOwner common.Address) error {
	return _CreditFilterMock.Contract.RevertIfAccountTransferIsNotAllowed(&_CreditFilterMock.CallOpts, owner, newOwner)
}

// RevertIfCantIncreaseBorrowing is a free data retrieval call binding the contract method 0xa5757517.
//
// Solidity: function revertIfCantIncreaseBorrowing(address creditAccount, uint256 minHealthFactor) view returns()
func (_CreditFilterMock *CreditFilterMockCaller) RevertIfCantIncreaseBorrowing(opts *bind.CallOpts, creditAccount common.Address, minHealthFactor *big.Int) error {
	var out []interface{}
	err := _CreditFilterMock.contract.Call(opts, &out, "revertIfCantIncreaseBorrowing", creditAccount, minHealthFactor)

	if err != nil {
		return err
	}

	return err

}

// RevertIfCantIncreaseBorrowing is a free data retrieval call binding the contract method 0xa5757517.
//
// Solidity: function revertIfCantIncreaseBorrowing(address creditAccount, uint256 minHealthFactor) view returns()
func (_CreditFilterMock *CreditFilterMockSession) RevertIfCantIncreaseBorrowing(creditAccount common.Address, minHealthFactor *big.Int) error {
	return _CreditFilterMock.Contract.RevertIfCantIncreaseBorrowing(&_CreditFilterMock.CallOpts, creditAccount, minHealthFactor)
}

// RevertIfCantIncreaseBorrowing is a free data retrieval call binding the contract method 0xa5757517.
//
// Solidity: function revertIfCantIncreaseBorrowing(address creditAccount, uint256 minHealthFactor) view returns()
func (_CreditFilterMock *CreditFilterMockCallerSession) RevertIfCantIncreaseBorrowing(creditAccount common.Address, minHealthFactor *big.Int) error {
	return _CreditFilterMock.Contract.RevertIfCantIncreaseBorrowing(&_CreditFilterMock.CallOpts, creditAccount, minHealthFactor)
}

// RevertIfTokenNotAllowed is a free data retrieval call binding the contract method 0x7dd0ba82.
//
// Solidity: function revertIfTokenNotAllowed(address token) view returns()
func (_CreditFilterMock *CreditFilterMockCaller) RevertIfTokenNotAllowed(opts *bind.CallOpts, token common.Address) error {
	var out []interface{}
	err := _CreditFilterMock.contract.Call(opts, &out, "revertIfTokenNotAllowed", token)

	if err != nil {
		return err
	}

	return err

}

// RevertIfTokenNotAllowed is a free data retrieval call binding the contract method 0x7dd0ba82.
//
// Solidity: function revertIfTokenNotAllowed(address token) view returns()
func (_CreditFilterMock *CreditFilterMockSession) RevertIfTokenNotAllowed(token common.Address) error {
	return _CreditFilterMock.Contract.RevertIfTokenNotAllowed(&_CreditFilterMock.CallOpts, token)
}

// RevertIfTokenNotAllowed is a free data retrieval call binding the contract method 0x7dd0ba82.
//
// Solidity: function revertIfTokenNotAllowed(address token) view returns()
func (_CreditFilterMock *CreditFilterMockCallerSession) RevertIfTokenNotAllowed(token common.Address) error {
	return _CreditFilterMock.Contract.RevertIfTokenNotAllowed(&_CreditFilterMock.CallOpts, token)
}

// TokenMasksMap is a free data retrieval call binding the contract method 0xf67c5bd0.
//
// Solidity: function tokenMasksMap(address ) view returns(uint256)
func (_CreditFilterMock *CreditFilterMockCaller) TokenMasksMap(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CreditFilterMock.contract.Call(opts, &out, "tokenMasksMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenMasksMap is a free data retrieval call binding the contract method 0xf67c5bd0.
//
// Solidity: function tokenMasksMap(address ) view returns(uint256)
func (_CreditFilterMock *CreditFilterMockSession) TokenMasksMap(arg0 common.Address) (*big.Int, error) {
	return _CreditFilterMock.Contract.TokenMasksMap(&_CreditFilterMock.CallOpts, arg0)
}

// TokenMasksMap is a free data retrieval call binding the contract method 0xf67c5bd0.
//
// Solidity: function tokenMasksMap(address ) view returns(uint256)
func (_CreditFilterMock *CreditFilterMockCallerSession) TokenMasksMap(arg0 common.Address) (*big.Int, error) {
	return _CreditFilterMock.Contract.TokenMasksMap(&_CreditFilterMock.CallOpts, arg0)
}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_CreditFilterMock *CreditFilterMockCaller) UnderlyingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditFilterMock.contract.Call(opts, &out, "underlyingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_CreditFilterMock *CreditFilterMockSession) UnderlyingToken() (common.Address, error) {
	return _CreditFilterMock.Contract.UnderlyingToken(&_CreditFilterMock.CallOpts)
}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_CreditFilterMock *CreditFilterMockCallerSession) UnderlyingToken() (common.Address, error) {
	return _CreditFilterMock.Contract.UnderlyingToken(&_CreditFilterMock.CallOpts)
}

// WethAddress is a free data retrieval call binding the contract method 0x4f0e0ef3.
//
// Solidity: function wethAddress() view returns(address)
func (_CreditFilterMock *CreditFilterMockCaller) WethAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditFilterMock.contract.Call(opts, &out, "wethAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WethAddress is a free data retrieval call binding the contract method 0x4f0e0ef3.
//
// Solidity: function wethAddress() view returns(address)
func (_CreditFilterMock *CreditFilterMockSession) WethAddress() (common.Address, error) {
	return _CreditFilterMock.Contract.WethAddress(&_CreditFilterMock.CallOpts)
}

// WethAddress is a free data retrieval call binding the contract method 0x4f0e0ef3.
//
// Solidity: function wethAddress() view returns(address)
func (_CreditFilterMock *CreditFilterMockCallerSession) WethAddress() (common.Address, error) {
	return _CreditFilterMock.Contract.WethAddress(&_CreditFilterMock.CallOpts)
}

// AllowContract is a paid mutator transaction binding the contract method 0x7bccacee.
//
// Solidity: function allowContract(address targetContract, address adapter) returns()
func (_CreditFilterMock *CreditFilterMockTransactor) AllowContract(opts *bind.TransactOpts, targetContract common.Address, adapter common.Address) (*types.Transaction, error) {
	return _CreditFilterMock.contract.Transact(opts, "allowContract", targetContract, adapter)
}

// AllowContract is a paid mutator transaction binding the contract method 0x7bccacee.
//
// Solidity: function allowContract(address targetContract, address adapter) returns()
func (_CreditFilterMock *CreditFilterMockSession) AllowContract(targetContract common.Address, adapter common.Address) (*types.Transaction, error) {
	return _CreditFilterMock.Contract.AllowContract(&_CreditFilterMock.TransactOpts, targetContract, adapter)
}

// AllowContract is a paid mutator transaction binding the contract method 0x7bccacee.
//
// Solidity: function allowContract(address targetContract, address adapter) returns()
func (_CreditFilterMock *CreditFilterMockTransactorSession) AllowContract(targetContract common.Address, adapter common.Address) (*types.Transaction, error) {
	return _CreditFilterMock.Contract.AllowContract(&_CreditFilterMock.TransactOpts, targetContract, adapter)
}

// AllowPlugin is a paid mutator transaction binding the contract method 0x2e2986dd.
//
// Solidity: function allowPlugin(address plugin, bool state) returns()
func (_CreditFilterMock *CreditFilterMockTransactor) AllowPlugin(opts *bind.TransactOpts, plugin common.Address, state bool) (*types.Transaction, error) {
	return _CreditFilterMock.contract.Transact(opts, "allowPlugin", plugin, state)
}

// AllowPlugin is a paid mutator transaction binding the contract method 0x2e2986dd.
//
// Solidity: function allowPlugin(address plugin, bool state) returns()
func (_CreditFilterMock *CreditFilterMockSession) AllowPlugin(plugin common.Address, state bool) (*types.Transaction, error) {
	return _CreditFilterMock.Contract.AllowPlugin(&_CreditFilterMock.TransactOpts, plugin, state)
}

// AllowPlugin is a paid mutator transaction binding the contract method 0x2e2986dd.
//
// Solidity: function allowPlugin(address plugin, bool state) returns()
func (_CreditFilterMock *CreditFilterMockTransactorSession) AllowPlugin(plugin common.Address, state bool) (*types.Transaction, error) {
	return _CreditFilterMock.Contract.AllowPlugin(&_CreditFilterMock.TransactOpts, plugin, state)
}

// AllowToken is a paid mutator transaction binding the contract method 0xa147c6c6.
//
// Solidity: function allowToken(address token, uint256 liquidationThreshold) returns()
func (_CreditFilterMock *CreditFilterMockTransactor) AllowToken(opts *bind.TransactOpts, token common.Address, liquidationThreshold *big.Int) (*types.Transaction, error) {
	return _CreditFilterMock.contract.Transact(opts, "allowToken", token, liquidationThreshold)
}

// AllowToken is a paid mutator transaction binding the contract method 0xa147c6c6.
//
// Solidity: function allowToken(address token, uint256 liquidationThreshold) returns()
func (_CreditFilterMock *CreditFilterMockSession) AllowToken(token common.Address, liquidationThreshold *big.Int) (*types.Transaction, error) {
	return _CreditFilterMock.Contract.AllowToken(&_CreditFilterMock.TransactOpts, token, liquidationThreshold)
}

// AllowToken is a paid mutator transaction binding the contract method 0xa147c6c6.
//
// Solidity: function allowToken(address token, uint256 liquidationThreshold) returns()
func (_CreditFilterMock *CreditFilterMockTransactorSession) AllowToken(token common.Address, liquidationThreshold *big.Int) (*types.Transaction, error) {
	return _CreditFilterMock.Contract.AllowToken(&_CreditFilterMock.TransactOpts, token, liquidationThreshold)
}

// ApproveAccountTransfers is a paid mutator transaction binding the contract method 0x5f27212a.
//
// Solidity: function approveAccountTransfers(address from, bool state) returns()
func (_CreditFilterMock *CreditFilterMockTransactor) ApproveAccountTransfers(opts *bind.TransactOpts, from common.Address, state bool) (*types.Transaction, error) {
	return _CreditFilterMock.contract.Transact(opts, "approveAccountTransfers", from, state)
}

// ApproveAccountTransfers is a paid mutator transaction binding the contract method 0x5f27212a.
//
// Solidity: function approveAccountTransfers(address from, bool state) returns()
func (_CreditFilterMock *CreditFilterMockSession) ApproveAccountTransfers(from common.Address, state bool) (*types.Transaction, error) {
	return _CreditFilterMock.Contract.ApproveAccountTransfers(&_CreditFilterMock.TransactOpts, from, state)
}

// ApproveAccountTransfers is a paid mutator transaction binding the contract method 0x5f27212a.
//
// Solidity: function approveAccountTransfers(address from, bool state) returns()
func (_CreditFilterMock *CreditFilterMockTransactorSession) ApproveAccountTransfers(from common.Address, state bool) (*types.Transaction, error) {
	return _CreditFilterMock.Contract.ApproveAccountTransfers(&_CreditFilterMock.TransactOpts, from, state)
}

// CheckAndEnableToken is a paid mutator transaction binding the contract method 0x51e3f160.
//
// Solidity: function checkAndEnableToken(address creditAccount, address token) returns()
func (_CreditFilterMock *CreditFilterMockTransactor) CheckAndEnableToken(opts *bind.TransactOpts, creditAccount common.Address, token common.Address) (*types.Transaction, error) {
	return _CreditFilterMock.contract.Transact(opts, "checkAndEnableToken", creditAccount, token)
}

// CheckAndEnableToken is a paid mutator transaction binding the contract method 0x51e3f160.
//
// Solidity: function checkAndEnableToken(address creditAccount, address token) returns()
func (_CreditFilterMock *CreditFilterMockSession) CheckAndEnableToken(creditAccount common.Address, token common.Address) (*types.Transaction, error) {
	return _CreditFilterMock.Contract.CheckAndEnableToken(&_CreditFilterMock.TransactOpts, creditAccount, token)
}

// CheckAndEnableToken is a paid mutator transaction binding the contract method 0x51e3f160.
//
// Solidity: function checkAndEnableToken(address creditAccount, address token) returns()
func (_CreditFilterMock *CreditFilterMockTransactorSession) CheckAndEnableToken(creditAccount common.Address, token common.Address) (*types.Transaction, error) {
	return _CreditFilterMock.Contract.CheckAndEnableToken(&_CreditFilterMock.TransactOpts, creditAccount, token)
}

// CheckCollateralChange is a paid mutator transaction binding the contract method 0xe1c8ef0d.
//
// Solidity: function checkCollateralChange(address creditAccount, address tokenIn, address tokenOut, uint256 amountIn, uint256 amountOut) returns()
func (_CreditFilterMock *CreditFilterMockTransactor) CheckCollateralChange(opts *bind.TransactOpts, creditAccount common.Address, tokenIn common.Address, tokenOut common.Address, amountIn *big.Int, amountOut *big.Int) (*types.Transaction, error) {
	return _CreditFilterMock.contract.Transact(opts, "checkCollateralChange", creditAccount, tokenIn, tokenOut, amountIn, amountOut)
}

// CheckCollateralChange is a paid mutator transaction binding the contract method 0xe1c8ef0d.
//
// Solidity: function checkCollateralChange(address creditAccount, address tokenIn, address tokenOut, uint256 amountIn, uint256 amountOut) returns()
func (_CreditFilterMock *CreditFilterMockSession) CheckCollateralChange(creditAccount common.Address, tokenIn common.Address, tokenOut common.Address, amountIn *big.Int, amountOut *big.Int) (*types.Transaction, error) {
	return _CreditFilterMock.Contract.CheckCollateralChange(&_CreditFilterMock.TransactOpts, creditAccount, tokenIn, tokenOut, amountIn, amountOut)
}

// CheckCollateralChange is a paid mutator transaction binding the contract method 0xe1c8ef0d.
//
// Solidity: function checkCollateralChange(address creditAccount, address tokenIn, address tokenOut, uint256 amountIn, uint256 amountOut) returns()
func (_CreditFilterMock *CreditFilterMockTransactorSession) CheckCollateralChange(creditAccount common.Address, tokenIn common.Address, tokenOut common.Address, amountIn *big.Int, amountOut *big.Int) (*types.Transaction, error) {
	return _CreditFilterMock.Contract.CheckCollateralChange(&_CreditFilterMock.TransactOpts, creditAccount, tokenIn, tokenOut, amountIn, amountOut)
}

// CheckMultiTokenCollateral is a paid mutator transaction binding the contract method 0x7e4a6863.
//
// Solidity: function checkMultiTokenCollateral(address creditAccount, uint256[] amountIn, uint256[] amountOut, address[] tokenIn, address[] tokenOut) returns()
func (_CreditFilterMock *CreditFilterMockTransactor) CheckMultiTokenCollateral(opts *bind.TransactOpts, creditAccount common.Address, amountIn []*big.Int, amountOut []*big.Int, tokenIn []common.Address, tokenOut []common.Address) (*types.Transaction, error) {
	return _CreditFilterMock.contract.Transact(opts, "checkMultiTokenCollateral", creditAccount, amountIn, amountOut, tokenIn, tokenOut)
}

// CheckMultiTokenCollateral is a paid mutator transaction binding the contract method 0x7e4a6863.
//
// Solidity: function checkMultiTokenCollateral(address creditAccount, uint256[] amountIn, uint256[] amountOut, address[] tokenIn, address[] tokenOut) returns()
func (_CreditFilterMock *CreditFilterMockSession) CheckMultiTokenCollateral(creditAccount common.Address, amountIn []*big.Int, amountOut []*big.Int, tokenIn []common.Address, tokenOut []common.Address) (*types.Transaction, error) {
	return _CreditFilterMock.Contract.CheckMultiTokenCollateral(&_CreditFilterMock.TransactOpts, creditAccount, amountIn, amountOut, tokenIn, tokenOut)
}

// CheckMultiTokenCollateral is a paid mutator transaction binding the contract method 0x7e4a6863.
//
// Solidity: function checkMultiTokenCollateral(address creditAccount, uint256[] amountIn, uint256[] amountOut, address[] tokenIn, address[] tokenOut) returns()
func (_CreditFilterMock *CreditFilterMockTransactorSession) CheckMultiTokenCollateral(creditAccount common.Address, amountIn []*big.Int, amountOut []*big.Int, tokenIn []common.Address, tokenOut []common.Address) (*types.Transaction, error) {
	return _CreditFilterMock.Contract.CheckMultiTokenCollateral(&_CreditFilterMock.TransactOpts, creditAccount, amountIn, amountOut, tokenIn, tokenOut)
}

// ConnectCreditManager is a paid mutator transaction binding the contract method 0xcf33d955.
//
// Solidity: function connectCreditManager(address _creditManager) returns()
func (_CreditFilterMock *CreditFilterMockTransactor) ConnectCreditManager(opts *bind.TransactOpts, _creditManager common.Address) (*types.Transaction, error) {
	return _CreditFilterMock.contract.Transact(opts, "connectCreditManager", _creditManager)
}

// ConnectCreditManager is a paid mutator transaction binding the contract method 0xcf33d955.
//
// Solidity: function connectCreditManager(address _creditManager) returns()
func (_CreditFilterMock *CreditFilterMockSession) ConnectCreditManager(_creditManager common.Address) (*types.Transaction, error) {
	return _CreditFilterMock.Contract.ConnectCreditManager(&_CreditFilterMock.TransactOpts, _creditManager)
}

// ConnectCreditManager is a paid mutator transaction binding the contract method 0xcf33d955.
//
// Solidity: function connectCreditManager(address _creditManager) returns()
func (_CreditFilterMock *CreditFilterMockTransactorSession) ConnectCreditManager(_creditManager common.Address) (*types.Transaction, error) {
	return _CreditFilterMock.Contract.ConnectCreditManager(&_CreditFilterMock.TransactOpts, _creditManager)
}

// ForbidContract is a paid mutator transaction binding the contract method 0x52438e54.
//
// Solidity: function forbidContract(address targetContract) returns()
func (_CreditFilterMock *CreditFilterMockTransactor) ForbidContract(opts *bind.TransactOpts, targetContract common.Address) (*types.Transaction, error) {
	return _CreditFilterMock.contract.Transact(opts, "forbidContract", targetContract)
}

// ForbidContract is a paid mutator transaction binding the contract method 0x52438e54.
//
// Solidity: function forbidContract(address targetContract) returns()
func (_CreditFilterMock *CreditFilterMockSession) ForbidContract(targetContract common.Address) (*types.Transaction, error) {
	return _CreditFilterMock.Contract.ForbidContract(&_CreditFilterMock.TransactOpts, targetContract)
}

// ForbidContract is a paid mutator transaction binding the contract method 0x52438e54.
//
// Solidity: function forbidContract(address targetContract) returns()
func (_CreditFilterMock *CreditFilterMockTransactorSession) ForbidContract(targetContract common.Address) (*types.Transaction, error) {
	return _CreditFilterMock.Contract.ForbidContract(&_CreditFilterMock.TransactOpts, targetContract)
}

// ForbidToken is a paid mutator transaction binding the contract method 0x24147708.
//
// Solidity: function forbidToken(address token) returns()
func (_CreditFilterMock *CreditFilterMockTransactor) ForbidToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _CreditFilterMock.contract.Transact(opts, "forbidToken", token)
}

// ForbidToken is a paid mutator transaction binding the contract method 0x24147708.
//
// Solidity: function forbidToken(address token) returns()
func (_CreditFilterMock *CreditFilterMockSession) ForbidToken(token common.Address) (*types.Transaction, error) {
	return _CreditFilterMock.Contract.ForbidToken(&_CreditFilterMock.TransactOpts, token)
}

// ForbidToken is a paid mutator transaction binding the contract method 0x24147708.
//
// Solidity: function forbidToken(address token) returns()
func (_CreditFilterMock *CreditFilterMockTransactorSession) ForbidToken(token common.Address) (*types.Transaction, error) {
	return _CreditFilterMock.Contract.ForbidToken(&_CreditFilterMock.TransactOpts, token)
}

// InitEnabledTokens is a paid mutator transaction binding the contract method 0xe54fe9c8.
//
// Solidity: function initEnabledTokens(address creditAccount) returns()
func (_CreditFilterMock *CreditFilterMockTransactor) InitEnabledTokens(opts *bind.TransactOpts, creditAccount common.Address) (*types.Transaction, error) {
	return _CreditFilterMock.contract.Transact(opts, "initEnabledTokens", creditAccount)
}

// InitEnabledTokens is a paid mutator transaction binding the contract method 0xe54fe9c8.
//
// Solidity: function initEnabledTokens(address creditAccount) returns()
func (_CreditFilterMock *CreditFilterMockSession) InitEnabledTokens(creditAccount common.Address) (*types.Transaction, error) {
	return _CreditFilterMock.Contract.InitEnabledTokens(&_CreditFilterMock.TransactOpts, creditAccount)
}

// InitEnabledTokens is a paid mutator transaction binding the contract method 0xe54fe9c8.
//
// Solidity: function initEnabledTokens(address creditAccount) returns()
func (_CreditFilterMock *CreditFilterMockTransactorSession) InitEnabledTokens(creditAccount common.Address) (*types.Transaction, error) {
	return _CreditFilterMock.Contract.InitEnabledTokens(&_CreditFilterMock.TransactOpts, creditAccount)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CreditFilterMock *CreditFilterMockTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditFilterMock.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CreditFilterMock *CreditFilterMockSession) Pause() (*types.Transaction, error) {
	return _CreditFilterMock.Contract.Pause(&_CreditFilterMock.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CreditFilterMock *CreditFilterMockTransactorSession) Pause() (*types.Transaction, error) {
	return _CreditFilterMock.Contract.Pause(&_CreditFilterMock.TransactOpts)
}

// SetEnabledTokens is a paid mutator transaction binding the contract method 0x044078b3.
//
// Solidity: function setEnabledTokens(address creditAccount, uint256 tokenMask) returns()
func (_CreditFilterMock *CreditFilterMockTransactor) SetEnabledTokens(opts *bind.TransactOpts, creditAccount common.Address, tokenMask *big.Int) (*types.Transaction, error) {
	return _CreditFilterMock.contract.Transact(opts, "setEnabledTokens", creditAccount, tokenMask)
}

// SetEnabledTokens is a paid mutator transaction binding the contract method 0x044078b3.
//
// Solidity: function setEnabledTokens(address creditAccount, uint256 tokenMask) returns()
func (_CreditFilterMock *CreditFilterMockSession) SetEnabledTokens(creditAccount common.Address, tokenMask *big.Int) (*types.Transaction, error) {
	return _CreditFilterMock.Contract.SetEnabledTokens(&_CreditFilterMock.TransactOpts, creditAccount, tokenMask)
}

// SetEnabledTokens is a paid mutator transaction binding the contract method 0x044078b3.
//
// Solidity: function setEnabledTokens(address creditAccount, uint256 tokenMask) returns()
func (_CreditFilterMock *CreditFilterMockTransactorSession) SetEnabledTokens(creditAccount common.Address, tokenMask *big.Int) (*types.Transaction, error) {
	return _CreditFilterMock.Contract.SetEnabledTokens(&_CreditFilterMock.TransactOpts, creditAccount, tokenMask)
}

// SetFastCheckBlock is a paid mutator transaction binding the contract method 0xa77e0ba8.
//
// Solidity: function setFastCheckBlock(address creditAccount, uint256 blockNum) returns()
func (_CreditFilterMock *CreditFilterMockTransactor) SetFastCheckBlock(opts *bind.TransactOpts, creditAccount common.Address, blockNum *big.Int) (*types.Transaction, error) {
	return _CreditFilterMock.contract.Transact(opts, "setFastCheckBlock", creditAccount, blockNum)
}

// SetFastCheckBlock is a paid mutator transaction binding the contract method 0xa77e0ba8.
//
// Solidity: function setFastCheckBlock(address creditAccount, uint256 blockNum) returns()
func (_CreditFilterMock *CreditFilterMockSession) SetFastCheckBlock(creditAccount common.Address, blockNum *big.Int) (*types.Transaction, error) {
	return _CreditFilterMock.Contract.SetFastCheckBlock(&_CreditFilterMock.TransactOpts, creditAccount, blockNum)
}

// SetFastCheckBlock is a paid mutator transaction binding the contract method 0xa77e0ba8.
//
// Solidity: function setFastCheckBlock(address creditAccount, uint256 blockNum) returns()
func (_CreditFilterMock *CreditFilterMockTransactorSession) SetFastCheckBlock(creditAccount common.Address, blockNum *big.Int) (*types.Transaction, error) {
	return _CreditFilterMock.Contract.SetFastCheckBlock(&_CreditFilterMock.TransactOpts, creditAccount, blockNum)
}

// SetFastCheckParameters is a paid mutator transaction binding the contract method 0x62061c6d.
//
// Solidity: function setFastCheckParameters(uint256 _chiThreshold, uint256 _hfCheckInterval) returns()
func (_CreditFilterMock *CreditFilterMockTransactor) SetFastCheckParameters(opts *bind.TransactOpts, _chiThreshold *big.Int, _hfCheckInterval *big.Int) (*types.Transaction, error) {
	return _CreditFilterMock.contract.Transact(opts, "setFastCheckParameters", _chiThreshold, _hfCheckInterval)
}

// SetFastCheckParameters is a paid mutator transaction binding the contract method 0x62061c6d.
//
// Solidity: function setFastCheckParameters(uint256 _chiThreshold, uint256 _hfCheckInterval) returns()
func (_CreditFilterMock *CreditFilterMockSession) SetFastCheckParameters(_chiThreshold *big.Int, _hfCheckInterval *big.Int) (*types.Transaction, error) {
	return _CreditFilterMock.Contract.SetFastCheckParameters(&_CreditFilterMock.TransactOpts, _chiThreshold, _hfCheckInterval)
}

// SetFastCheckParameters is a paid mutator transaction binding the contract method 0x62061c6d.
//
// Solidity: function setFastCheckParameters(uint256 _chiThreshold, uint256 _hfCheckInterval) returns()
func (_CreditFilterMock *CreditFilterMockTransactorSession) SetFastCheckParameters(_chiThreshold *big.Int, _hfCheckInterval *big.Int) (*types.Transaction, error) {
	return _CreditFilterMock.Contract.SetFastCheckParameters(&_CreditFilterMock.TransactOpts, _chiThreshold, _hfCheckInterval)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CreditFilterMock *CreditFilterMockTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditFilterMock.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CreditFilterMock *CreditFilterMockSession) Unpause() (*types.Transaction, error) {
	return _CreditFilterMock.Contract.Unpause(&_CreditFilterMock.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CreditFilterMock *CreditFilterMockTransactorSession) Unpause() (*types.Transaction, error) {
	return _CreditFilterMock.Contract.Unpause(&_CreditFilterMock.TransactOpts)
}

// UpdateUnderlyingTokenLiquidationThreshold is a paid mutator transaction binding the contract method 0x40631828.
//
// Solidity: function updateUnderlyingTokenLiquidationThreshold() returns()
func (_CreditFilterMock *CreditFilterMockTransactor) UpdateUnderlyingTokenLiquidationThreshold(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditFilterMock.contract.Transact(opts, "updateUnderlyingTokenLiquidationThreshold")
}

// UpdateUnderlyingTokenLiquidationThreshold is a paid mutator transaction binding the contract method 0x40631828.
//
// Solidity: function updateUnderlyingTokenLiquidationThreshold() returns()
func (_CreditFilterMock *CreditFilterMockSession) UpdateUnderlyingTokenLiquidationThreshold() (*types.Transaction, error) {
	return _CreditFilterMock.Contract.UpdateUnderlyingTokenLiquidationThreshold(&_CreditFilterMock.TransactOpts)
}

// UpdateUnderlyingTokenLiquidationThreshold is a paid mutator transaction binding the contract method 0x40631828.
//
// Solidity: function updateUnderlyingTokenLiquidationThreshold() returns()
func (_CreditFilterMock *CreditFilterMockTransactorSession) UpdateUnderlyingTokenLiquidationThreshold() (*types.Transaction, error) {
	return _CreditFilterMock.Contract.UpdateUnderlyingTokenLiquidationThreshold(&_CreditFilterMock.TransactOpts)
}

// CreditFilterMockContractAllowedIterator is returned from FilterContractAllowed and is used to iterate over the raw logs and unpacked data for ContractAllowed events raised by the CreditFilterMock contract.
type CreditFilterMockContractAllowedIterator struct {
	Event *CreditFilterMockContractAllowed // Event containing the contract specifics and raw log

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
func (it *CreditFilterMockContractAllowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditFilterMockContractAllowed)
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
		it.Event = new(CreditFilterMockContractAllowed)
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
func (it *CreditFilterMockContractAllowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditFilterMockContractAllowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditFilterMockContractAllowed represents a ContractAllowed event raised by the CreditFilterMock contract.
type CreditFilterMockContractAllowed struct {
	Protocol common.Address
	Adapter  common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterContractAllowed is a free log retrieval operation binding the contract event 0x4bcbefaef68b99503d502f5a6abe7bca2b183ab8ac55457013c77d084ebd1305.
//
// Solidity: event ContractAllowed(address indexed protocol, address indexed adapter)
func (_CreditFilterMock *CreditFilterMockFilterer) FilterContractAllowed(opts *bind.FilterOpts, protocol []common.Address, adapter []common.Address) (*CreditFilterMockContractAllowedIterator, error) {

	var protocolRule []interface{}
	for _, protocolItem := range protocol {
		protocolRule = append(protocolRule, protocolItem)
	}
	var adapterRule []interface{}
	for _, adapterItem := range adapter {
		adapterRule = append(adapterRule, adapterItem)
	}

	logs, sub, err := _CreditFilterMock.contract.FilterLogs(opts, "ContractAllowed", protocolRule, adapterRule)
	if err != nil {
		return nil, err
	}
	return &CreditFilterMockContractAllowedIterator{contract: _CreditFilterMock.contract, event: "ContractAllowed", logs: logs, sub: sub}, nil
}

// WatchContractAllowed is a free log subscription operation binding the contract event 0x4bcbefaef68b99503d502f5a6abe7bca2b183ab8ac55457013c77d084ebd1305.
//
// Solidity: event ContractAllowed(address indexed protocol, address indexed adapter)
func (_CreditFilterMock *CreditFilterMockFilterer) WatchContractAllowed(opts *bind.WatchOpts, sink chan<- *CreditFilterMockContractAllowed, protocol []common.Address, adapter []common.Address) (event.Subscription, error) {

	var protocolRule []interface{}
	for _, protocolItem := range protocol {
		protocolRule = append(protocolRule, protocolItem)
	}
	var adapterRule []interface{}
	for _, adapterItem := range adapter {
		adapterRule = append(adapterRule, adapterItem)
	}

	logs, sub, err := _CreditFilterMock.contract.WatchLogs(opts, "ContractAllowed", protocolRule, adapterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditFilterMockContractAllowed)
				if err := _CreditFilterMock.contract.UnpackLog(event, "ContractAllowed", log); err != nil {
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
func (_CreditFilterMock *CreditFilterMockFilterer) ParseContractAllowed(log types.Log) (*CreditFilterMockContractAllowed, error) {
	event := new(CreditFilterMockContractAllowed)
	if err := _CreditFilterMock.contract.UnpackLog(event, "ContractAllowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditFilterMockContractForbiddenIterator is returned from FilterContractForbidden and is used to iterate over the raw logs and unpacked data for ContractForbidden events raised by the CreditFilterMock contract.
type CreditFilterMockContractForbiddenIterator struct {
	Event *CreditFilterMockContractForbidden // Event containing the contract specifics and raw log

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
func (it *CreditFilterMockContractForbiddenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditFilterMockContractForbidden)
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
		it.Event = new(CreditFilterMockContractForbidden)
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
func (it *CreditFilterMockContractForbiddenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditFilterMockContractForbiddenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditFilterMockContractForbidden represents a ContractForbidden event raised by the CreditFilterMock contract.
type CreditFilterMockContractForbidden struct {
	Protocol common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterContractForbidden is a free log retrieval operation binding the contract event 0xab9f405bf0c19b97f65a7031634db41569cd2f0e0376a610a1e977f9ab22b58f.
//
// Solidity: event ContractForbidden(address indexed protocol)
func (_CreditFilterMock *CreditFilterMockFilterer) FilterContractForbidden(opts *bind.FilterOpts, protocol []common.Address) (*CreditFilterMockContractForbiddenIterator, error) {

	var protocolRule []interface{}
	for _, protocolItem := range protocol {
		protocolRule = append(protocolRule, protocolItem)
	}

	logs, sub, err := _CreditFilterMock.contract.FilterLogs(opts, "ContractForbidden", protocolRule)
	if err != nil {
		return nil, err
	}
	return &CreditFilterMockContractForbiddenIterator{contract: _CreditFilterMock.contract, event: "ContractForbidden", logs: logs, sub: sub}, nil
}

// WatchContractForbidden is a free log subscription operation binding the contract event 0xab9f405bf0c19b97f65a7031634db41569cd2f0e0376a610a1e977f9ab22b58f.
//
// Solidity: event ContractForbidden(address indexed protocol)
func (_CreditFilterMock *CreditFilterMockFilterer) WatchContractForbidden(opts *bind.WatchOpts, sink chan<- *CreditFilterMockContractForbidden, protocol []common.Address) (event.Subscription, error) {

	var protocolRule []interface{}
	for _, protocolItem := range protocol {
		protocolRule = append(protocolRule, protocolItem)
	}

	logs, sub, err := _CreditFilterMock.contract.WatchLogs(opts, "ContractForbidden", protocolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditFilterMockContractForbidden)
				if err := _CreditFilterMock.contract.UnpackLog(event, "ContractForbidden", log); err != nil {
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
func (_CreditFilterMock *CreditFilterMockFilterer) ParseContractForbidden(log types.Log) (*CreditFilterMockContractForbidden, error) {
	event := new(CreditFilterMockContractForbidden)
	if err := _CreditFilterMock.contract.UnpackLog(event, "ContractForbidden", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditFilterMockNewFastCheckParametersIterator is returned from FilterNewFastCheckParameters and is used to iterate over the raw logs and unpacked data for NewFastCheckParameters events raised by the CreditFilterMock contract.
type CreditFilterMockNewFastCheckParametersIterator struct {
	Event *CreditFilterMockNewFastCheckParameters // Event containing the contract specifics and raw log

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
func (it *CreditFilterMockNewFastCheckParametersIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditFilterMockNewFastCheckParameters)
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
		it.Event = new(CreditFilterMockNewFastCheckParameters)
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
func (it *CreditFilterMockNewFastCheckParametersIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditFilterMockNewFastCheckParametersIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditFilterMockNewFastCheckParameters represents a NewFastCheckParameters event raised by the CreditFilterMock contract.
type CreditFilterMockNewFastCheckParameters struct {
	ChiThreshold   *big.Int
	FastCheckDelay *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewFastCheckParameters is a free log retrieval operation binding the contract event 0x727652fff0946c19c233fd3eab5fc03db9e9fdd907e902d9136c2a9cac47101c.
//
// Solidity: event NewFastCheckParameters(uint256 chiThreshold, uint256 fastCheckDelay)
func (_CreditFilterMock *CreditFilterMockFilterer) FilterNewFastCheckParameters(opts *bind.FilterOpts) (*CreditFilterMockNewFastCheckParametersIterator, error) {

	logs, sub, err := _CreditFilterMock.contract.FilterLogs(opts, "NewFastCheckParameters")
	if err != nil {
		return nil, err
	}
	return &CreditFilterMockNewFastCheckParametersIterator{contract: _CreditFilterMock.contract, event: "NewFastCheckParameters", logs: logs, sub: sub}, nil
}

// WatchNewFastCheckParameters is a free log subscription operation binding the contract event 0x727652fff0946c19c233fd3eab5fc03db9e9fdd907e902d9136c2a9cac47101c.
//
// Solidity: event NewFastCheckParameters(uint256 chiThreshold, uint256 fastCheckDelay)
func (_CreditFilterMock *CreditFilterMockFilterer) WatchNewFastCheckParameters(opts *bind.WatchOpts, sink chan<- *CreditFilterMockNewFastCheckParameters) (event.Subscription, error) {

	logs, sub, err := _CreditFilterMock.contract.WatchLogs(opts, "NewFastCheckParameters")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditFilterMockNewFastCheckParameters)
				if err := _CreditFilterMock.contract.UnpackLog(event, "NewFastCheckParameters", log); err != nil {
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
func (_CreditFilterMock *CreditFilterMockFilterer) ParseNewFastCheckParameters(log types.Log) (*CreditFilterMockNewFastCheckParameters, error) {
	event := new(CreditFilterMockNewFastCheckParameters)
	if err := _CreditFilterMock.contract.UnpackLog(event, "NewFastCheckParameters", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditFilterMockPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the CreditFilterMock contract.
type CreditFilterMockPausedIterator struct {
	Event *CreditFilterMockPaused // Event containing the contract specifics and raw log

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
func (it *CreditFilterMockPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditFilterMockPaused)
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
		it.Event = new(CreditFilterMockPaused)
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
func (it *CreditFilterMockPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditFilterMockPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditFilterMockPaused represents a Paused event raised by the CreditFilterMock contract.
type CreditFilterMockPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_CreditFilterMock *CreditFilterMockFilterer) FilterPaused(opts *bind.FilterOpts) (*CreditFilterMockPausedIterator, error) {

	logs, sub, err := _CreditFilterMock.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &CreditFilterMockPausedIterator{contract: _CreditFilterMock.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_CreditFilterMock *CreditFilterMockFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *CreditFilterMockPaused) (event.Subscription, error) {

	logs, sub, err := _CreditFilterMock.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditFilterMockPaused)
				if err := _CreditFilterMock.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_CreditFilterMock *CreditFilterMockFilterer) ParsePaused(log types.Log) (*CreditFilterMockPaused, error) {
	event := new(CreditFilterMockPaused)
	if err := _CreditFilterMock.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditFilterMockTokenAllowedIterator is returned from FilterTokenAllowed and is used to iterate over the raw logs and unpacked data for TokenAllowed events raised by the CreditFilterMock contract.
type CreditFilterMockTokenAllowedIterator struct {
	Event *CreditFilterMockTokenAllowed // Event containing the contract specifics and raw log

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
func (it *CreditFilterMockTokenAllowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditFilterMockTokenAllowed)
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
		it.Event = new(CreditFilterMockTokenAllowed)
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
func (it *CreditFilterMockTokenAllowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditFilterMockTokenAllowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditFilterMockTokenAllowed represents a TokenAllowed event raised by the CreditFilterMock contract.
type CreditFilterMockTokenAllowed struct {
	Token              common.Address
	LiquidityThreshold *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterTokenAllowed is a free log retrieval operation binding the contract event 0xa52fb6bfa514a4ddcb31de40a5f6c20d767db1f921a8b7747973d93dc5da7a02.
//
// Solidity: event TokenAllowed(address indexed token, uint256 liquidityThreshold)
func (_CreditFilterMock *CreditFilterMockFilterer) FilterTokenAllowed(opts *bind.FilterOpts, token []common.Address) (*CreditFilterMockTokenAllowedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CreditFilterMock.contract.FilterLogs(opts, "TokenAllowed", tokenRule)
	if err != nil {
		return nil, err
	}
	return &CreditFilterMockTokenAllowedIterator{contract: _CreditFilterMock.contract, event: "TokenAllowed", logs: logs, sub: sub}, nil
}

// WatchTokenAllowed is a free log subscription operation binding the contract event 0xa52fb6bfa514a4ddcb31de40a5f6c20d767db1f921a8b7747973d93dc5da7a02.
//
// Solidity: event TokenAllowed(address indexed token, uint256 liquidityThreshold)
func (_CreditFilterMock *CreditFilterMockFilterer) WatchTokenAllowed(opts *bind.WatchOpts, sink chan<- *CreditFilterMockTokenAllowed, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CreditFilterMock.contract.WatchLogs(opts, "TokenAllowed", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditFilterMockTokenAllowed)
				if err := _CreditFilterMock.contract.UnpackLog(event, "TokenAllowed", log); err != nil {
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
func (_CreditFilterMock *CreditFilterMockFilterer) ParseTokenAllowed(log types.Log) (*CreditFilterMockTokenAllowed, error) {
	event := new(CreditFilterMockTokenAllowed)
	if err := _CreditFilterMock.contract.UnpackLog(event, "TokenAllowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditFilterMockTransferAccountAllowedIterator is returned from FilterTransferAccountAllowed and is used to iterate over the raw logs and unpacked data for TransferAccountAllowed events raised by the CreditFilterMock contract.
type CreditFilterMockTransferAccountAllowedIterator struct {
	Event *CreditFilterMockTransferAccountAllowed // Event containing the contract specifics and raw log

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
func (it *CreditFilterMockTransferAccountAllowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditFilterMockTransferAccountAllowed)
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
		it.Event = new(CreditFilterMockTransferAccountAllowed)
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
func (it *CreditFilterMockTransferAccountAllowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditFilterMockTransferAccountAllowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditFilterMockTransferAccountAllowed represents a TransferAccountAllowed event raised by the CreditFilterMock contract.
type CreditFilterMockTransferAccountAllowed struct {
	From  common.Address
	To    common.Address
	State bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransferAccountAllowed is a free log retrieval operation binding the contract event 0x9b3258bc4904fd6426b99843e206c6c7cdb1fd0f040121c25b71dafbb3851ee0.
//
// Solidity: event TransferAccountAllowed(address indexed from, address indexed to, bool state)
func (_CreditFilterMock *CreditFilterMockFilterer) FilterTransferAccountAllowed(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CreditFilterMockTransferAccountAllowedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CreditFilterMock.contract.FilterLogs(opts, "TransferAccountAllowed", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CreditFilterMockTransferAccountAllowedIterator{contract: _CreditFilterMock.contract, event: "TransferAccountAllowed", logs: logs, sub: sub}, nil
}

// WatchTransferAccountAllowed is a free log subscription operation binding the contract event 0x9b3258bc4904fd6426b99843e206c6c7cdb1fd0f040121c25b71dafbb3851ee0.
//
// Solidity: event TransferAccountAllowed(address indexed from, address indexed to, bool state)
func (_CreditFilterMock *CreditFilterMockFilterer) WatchTransferAccountAllowed(opts *bind.WatchOpts, sink chan<- *CreditFilterMockTransferAccountAllowed, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CreditFilterMock.contract.WatchLogs(opts, "TransferAccountAllowed", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditFilterMockTransferAccountAllowed)
				if err := _CreditFilterMock.contract.UnpackLog(event, "TransferAccountAllowed", log); err != nil {
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
func (_CreditFilterMock *CreditFilterMockFilterer) ParseTransferAccountAllowed(log types.Log) (*CreditFilterMockTransferAccountAllowed, error) {
	event := new(CreditFilterMockTransferAccountAllowed)
	if err := _CreditFilterMock.contract.UnpackLog(event, "TransferAccountAllowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditFilterMockTransferPluginAllowedIterator is returned from FilterTransferPluginAllowed and is used to iterate over the raw logs and unpacked data for TransferPluginAllowed events raised by the CreditFilterMock contract.
type CreditFilterMockTransferPluginAllowedIterator struct {
	Event *CreditFilterMockTransferPluginAllowed // Event containing the contract specifics and raw log

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
func (it *CreditFilterMockTransferPluginAllowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditFilterMockTransferPluginAllowed)
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
		it.Event = new(CreditFilterMockTransferPluginAllowed)
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
func (it *CreditFilterMockTransferPluginAllowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditFilterMockTransferPluginAllowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditFilterMockTransferPluginAllowed represents a TransferPluginAllowed event raised by the CreditFilterMock contract.
type CreditFilterMockTransferPluginAllowed struct {
	Pugin common.Address
	State bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransferPluginAllowed is a free log retrieval operation binding the contract event 0xc7d2592986c53f858769b011e8ce6298936f8609789988e9f5ad4f0a20798897.
//
// Solidity: event TransferPluginAllowed(address indexed pugin, bool state)
func (_CreditFilterMock *CreditFilterMockFilterer) FilterTransferPluginAllowed(opts *bind.FilterOpts, pugin []common.Address) (*CreditFilterMockTransferPluginAllowedIterator, error) {

	var puginRule []interface{}
	for _, puginItem := range pugin {
		puginRule = append(puginRule, puginItem)
	}

	logs, sub, err := _CreditFilterMock.contract.FilterLogs(opts, "TransferPluginAllowed", puginRule)
	if err != nil {
		return nil, err
	}
	return &CreditFilterMockTransferPluginAllowedIterator{contract: _CreditFilterMock.contract, event: "TransferPluginAllowed", logs: logs, sub: sub}, nil
}

// WatchTransferPluginAllowed is a free log subscription operation binding the contract event 0xc7d2592986c53f858769b011e8ce6298936f8609789988e9f5ad4f0a20798897.
//
// Solidity: event TransferPluginAllowed(address indexed pugin, bool state)
func (_CreditFilterMock *CreditFilterMockFilterer) WatchTransferPluginAllowed(opts *bind.WatchOpts, sink chan<- *CreditFilterMockTransferPluginAllowed, pugin []common.Address) (event.Subscription, error) {

	var puginRule []interface{}
	for _, puginItem := range pugin {
		puginRule = append(puginRule, puginItem)
	}

	logs, sub, err := _CreditFilterMock.contract.WatchLogs(opts, "TransferPluginAllowed", puginRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditFilterMockTransferPluginAllowed)
				if err := _CreditFilterMock.contract.UnpackLog(event, "TransferPluginAllowed", log); err != nil {
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
func (_CreditFilterMock *CreditFilterMockFilterer) ParseTransferPluginAllowed(log types.Log) (*CreditFilterMockTransferPluginAllowed, error) {
	event := new(CreditFilterMockTransferPluginAllowed)
	if err := _CreditFilterMock.contract.UnpackLog(event, "TransferPluginAllowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditFilterMockUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the CreditFilterMock contract.
type CreditFilterMockUnpausedIterator struct {
	Event *CreditFilterMockUnpaused // Event containing the contract specifics and raw log

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
func (it *CreditFilterMockUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditFilterMockUnpaused)
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
		it.Event = new(CreditFilterMockUnpaused)
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
func (it *CreditFilterMockUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditFilterMockUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditFilterMockUnpaused represents a Unpaused event raised by the CreditFilterMock contract.
type CreditFilterMockUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_CreditFilterMock *CreditFilterMockFilterer) FilterUnpaused(opts *bind.FilterOpts) (*CreditFilterMockUnpausedIterator, error) {

	logs, sub, err := _CreditFilterMock.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &CreditFilterMockUnpausedIterator{contract: _CreditFilterMock.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_CreditFilterMock *CreditFilterMockFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *CreditFilterMockUnpaused) (event.Subscription, error) {

	logs, sub, err := _CreditFilterMock.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditFilterMockUnpaused)
				if err := _CreditFilterMock.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_CreditFilterMock *CreditFilterMockFilterer) ParseUnpaused(log types.Log) (*CreditFilterMockUnpaused, error) {
	event := new(CreditFilterMockUnpaused)
	if err := _CreditFilterMock.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
