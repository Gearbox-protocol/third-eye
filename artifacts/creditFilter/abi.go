// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package creditFilter

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

// CreditFilterABI is the input ABI used to generate the binding from.
const CreditFilterABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addressProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_underlyingToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"protocol\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"adapter\",\"type\":\"address\"}],\"name\":\"ContractAllowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"protocol\",\"type\":\"address\"}],\"name\":\"ContractForbidden\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chiThreshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fastCheckDelay\",\"type\":\"uint256\"}],\"name\":\"NewFastCheckParameters\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newPriceOracle\",\"type\":\"address\"}],\"name\":\"PriceOracleUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidityThreshold\",\"type\":\"uint256\"}],\"name\":\"TokenAllowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenForbidden\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"state\",\"type\":\"bool\"}],\"name\":\"TransferAccountAllowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pugin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"state\",\"type\":\"bool\"}],\"name\":\"TransferPluginAllowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"addressProvider\",\"outputs\":[{\"internalType\":\"contractAddressProvider\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"targetContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"adapter\",\"type\":\"address\"}],\"name\":\"allowContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"plugin\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"state\",\"type\":\"bool\"}],\"name\":\"allowPlugin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidationThreshold\",\"type\":\"uint256\"}],\"name\":\"allowToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"allowanceForAccountTransfers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowedAdapters\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"allowedContracts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allowedContractsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowedPlugins\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allowedTokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allowedTokensCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"state\",\"type\":\"bool\"}],\"name\":\"approveAccountTransfers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"}],\"name\":\"calcCreditAccountAccruedInterest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"}],\"name\":\"calcCreditAccountHealthFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"percentage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"times\",\"type\":\"uint256\"}],\"name\":\"calcMaxPossibleDrop\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"}],\"name\":\"calcThresholdWeightedValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"}],\"name\":\"calcTotalValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"checkAndEnableToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"checkCollateralChange\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"amountIn\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amountOut\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"tokenIn\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"tokenOut\",\"type\":\"address[]\"}],\"name\":\"checkMultiTokenCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chiThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_creditManager\",\"type\":\"address\"}],\"name\":\"connectCreditManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"contractToAdapter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creditManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"enabledTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"fastCheckCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"targetContract\",\"type\":\"address\"}],\"name\":\"forbidContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"forbidToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getCreditAccountTokenById\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tvw\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hfCheckInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"}],\"name\":\"initEnabledTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isTokenAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"liquidationThresholds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolService\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"revertIfAccountTransferIsNotAllowed\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minHealthFactor\",\"type\":\"uint256\"}],\"name\":\"revertIfCantIncreaseBorrowing\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"revertIfTokenNotAllowed\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chiThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_hfCheckInterval\",\"type\":\"uint256\"}],\"name\":\"setFastCheckParameters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenMasksMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlyingToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateUnderlyingTokenLiquidationThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"upgradePriceOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wethAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// CreditFilter is an auto generated Go binding around an Ethereum contract.
type CreditFilter struct {
	CreditFilterCaller     // Read-only binding to the contract
	CreditFilterTransactor // Write-only binding to the contract
	CreditFilterFilterer   // Log filterer for contract events
}

// CreditFilterCaller is an auto generated read-only Go binding around an Ethereum contract.
type CreditFilterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditFilterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CreditFilterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditFilterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CreditFilterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditFilterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CreditFilterSession struct {
	Contract     *CreditFilter     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CreditFilterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CreditFilterCallerSession struct {
	Contract *CreditFilterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// CreditFilterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CreditFilterTransactorSession struct {
	Contract     *CreditFilterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// CreditFilterRaw is an auto generated low-level Go binding around an Ethereum contract.
type CreditFilterRaw struct {
	Contract *CreditFilter // Generic contract binding to access the raw methods on
}

// CreditFilterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CreditFilterCallerRaw struct {
	Contract *CreditFilterCaller // Generic read-only contract binding to access the raw methods on
}

// CreditFilterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CreditFilterTransactorRaw struct {
	Contract *CreditFilterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCreditFilter creates a new instance of CreditFilter, bound to a specific deployed contract.
func NewCreditFilter(address common.Address, backend bind.ContractBackend) (*CreditFilter, error) {
	contract, err := bindCreditFilter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CreditFilter{CreditFilterCaller: CreditFilterCaller{contract: contract}, CreditFilterTransactor: CreditFilterTransactor{contract: contract}, CreditFilterFilterer: CreditFilterFilterer{contract: contract}}, nil
}

// NewCreditFilterCaller creates a new read-only instance of CreditFilter, bound to a specific deployed contract.
func NewCreditFilterCaller(address common.Address, caller bind.ContractCaller) (*CreditFilterCaller, error) {
	contract, err := bindCreditFilter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CreditFilterCaller{contract: contract}, nil
}

// NewCreditFilterTransactor creates a new write-only instance of CreditFilter, bound to a specific deployed contract.
func NewCreditFilterTransactor(address common.Address, transactor bind.ContractTransactor) (*CreditFilterTransactor, error) {
	contract, err := bindCreditFilter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CreditFilterTransactor{contract: contract}, nil
}

// NewCreditFilterFilterer creates a new log filterer instance of CreditFilter, bound to a specific deployed contract.
func NewCreditFilterFilterer(address common.Address, filterer bind.ContractFilterer) (*CreditFilterFilterer, error) {
	contract, err := bindCreditFilter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CreditFilterFilterer{contract: contract}, nil
}

// bindCreditFilter binds a generic wrapper to an already deployed contract.
func bindCreditFilter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CreditFilterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CreditFilter *CreditFilterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CreditFilter.Contract.CreditFilterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CreditFilter *CreditFilterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditFilter.Contract.CreditFilterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CreditFilter *CreditFilterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CreditFilter.Contract.CreditFilterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CreditFilter *CreditFilterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CreditFilter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CreditFilter *CreditFilterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditFilter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CreditFilter *CreditFilterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CreditFilter.Contract.contract.Transact(opts, method, params...)
}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_CreditFilter *CreditFilterCaller) AddressProvider(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditFilter.contract.Call(opts, &out, "addressProvider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_CreditFilter *CreditFilterSession) AddressProvider() (common.Address, error) {
	return _CreditFilter.Contract.AddressProvider(&_CreditFilter.CallOpts)
}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_CreditFilter *CreditFilterCallerSession) AddressProvider() (common.Address, error) {
	return _CreditFilter.Contract.AddressProvider(&_CreditFilter.CallOpts)
}

// AllowanceForAccountTransfers is a free data retrieval call binding the contract method 0x5a29be45.
//
// Solidity: function allowanceForAccountTransfers(address from, address to) view returns(bool)
func (_CreditFilter *CreditFilterCaller) AllowanceForAccountTransfers(opts *bind.CallOpts, from common.Address, to common.Address) (bool, error) {
	var out []interface{}
	err := _CreditFilter.contract.Call(opts, &out, "allowanceForAccountTransfers", from, to)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowanceForAccountTransfers is a free data retrieval call binding the contract method 0x5a29be45.
//
// Solidity: function allowanceForAccountTransfers(address from, address to) view returns(bool)
func (_CreditFilter *CreditFilterSession) AllowanceForAccountTransfers(from common.Address, to common.Address) (bool, error) {
	return _CreditFilter.Contract.AllowanceForAccountTransfers(&_CreditFilter.CallOpts, from, to)
}

// AllowanceForAccountTransfers is a free data retrieval call binding the contract method 0x5a29be45.
//
// Solidity: function allowanceForAccountTransfers(address from, address to) view returns(bool)
func (_CreditFilter *CreditFilterCallerSession) AllowanceForAccountTransfers(from common.Address, to common.Address) (bool, error) {
	return _CreditFilter.Contract.AllowanceForAccountTransfers(&_CreditFilter.CallOpts, from, to)
}

// AllowedAdapters is a free data retrieval call binding the contract method 0x3bdfe4f5.
//
// Solidity: function allowedAdapters(address ) view returns(bool)
func (_CreditFilter *CreditFilterCaller) AllowedAdapters(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _CreditFilter.contract.Call(opts, &out, "allowedAdapters", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedAdapters is a free data retrieval call binding the contract method 0x3bdfe4f5.
//
// Solidity: function allowedAdapters(address ) view returns(bool)
func (_CreditFilter *CreditFilterSession) AllowedAdapters(arg0 common.Address) (bool, error) {
	return _CreditFilter.Contract.AllowedAdapters(&_CreditFilter.CallOpts, arg0)
}

// AllowedAdapters is a free data retrieval call binding the contract method 0x3bdfe4f5.
//
// Solidity: function allowedAdapters(address ) view returns(bool)
func (_CreditFilter *CreditFilterCallerSession) AllowedAdapters(arg0 common.Address) (bool, error) {
	return _CreditFilter.Contract.AllowedAdapters(&_CreditFilter.CallOpts, arg0)
}

// AllowedContracts is a free data retrieval call binding the contract method 0x5094cb4f.
//
// Solidity: function allowedContracts(uint256 i) view returns(address)
func (_CreditFilter *CreditFilterCaller) AllowedContracts(opts *bind.CallOpts, i *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CreditFilter.contract.Call(opts, &out, "allowedContracts", i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllowedContracts is a free data retrieval call binding the contract method 0x5094cb4f.
//
// Solidity: function allowedContracts(uint256 i) view returns(address)
func (_CreditFilter *CreditFilterSession) AllowedContracts(i *big.Int) (common.Address, error) {
	return _CreditFilter.Contract.AllowedContracts(&_CreditFilter.CallOpts, i)
}

// AllowedContracts is a free data retrieval call binding the contract method 0x5094cb4f.
//
// Solidity: function allowedContracts(uint256 i) view returns(address)
func (_CreditFilter *CreditFilterCallerSession) AllowedContracts(i *big.Int) (common.Address, error) {
	return _CreditFilter.Contract.AllowedContracts(&_CreditFilter.CallOpts, i)
}

// AllowedContractsCount is a free data retrieval call binding the contract method 0x50e036ff.
//
// Solidity: function allowedContractsCount() view returns(uint256)
func (_CreditFilter *CreditFilterCaller) AllowedContractsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CreditFilter.contract.Call(opts, &out, "allowedContractsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllowedContractsCount is a free data retrieval call binding the contract method 0x50e036ff.
//
// Solidity: function allowedContractsCount() view returns(uint256)
func (_CreditFilter *CreditFilterSession) AllowedContractsCount() (*big.Int, error) {
	return _CreditFilter.Contract.AllowedContractsCount(&_CreditFilter.CallOpts)
}

// AllowedContractsCount is a free data retrieval call binding the contract method 0x50e036ff.
//
// Solidity: function allowedContractsCount() view returns(uint256)
func (_CreditFilter *CreditFilterCallerSession) AllowedContractsCount() (*big.Int, error) {
	return _CreditFilter.Contract.AllowedContractsCount(&_CreditFilter.CallOpts)
}

// AllowedPlugins is a free data retrieval call binding the contract method 0x5f598edd.
//
// Solidity: function allowedPlugins(address ) view returns(bool)
func (_CreditFilter *CreditFilterCaller) AllowedPlugins(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _CreditFilter.contract.Call(opts, &out, "allowedPlugins", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedPlugins is a free data retrieval call binding the contract method 0x5f598edd.
//
// Solidity: function allowedPlugins(address ) view returns(bool)
func (_CreditFilter *CreditFilterSession) AllowedPlugins(arg0 common.Address) (bool, error) {
	return _CreditFilter.Contract.AllowedPlugins(&_CreditFilter.CallOpts, arg0)
}

// AllowedPlugins is a free data retrieval call binding the contract method 0x5f598edd.
//
// Solidity: function allowedPlugins(address ) view returns(bool)
func (_CreditFilter *CreditFilterCallerSession) AllowedPlugins(arg0 common.Address) (bool, error) {
	return _CreditFilter.Contract.AllowedPlugins(&_CreditFilter.CallOpts, arg0)
}

// AllowedTokens is a free data retrieval call binding the contract method 0x5e5f2e26.
//
// Solidity: function allowedTokens(uint256 ) view returns(address)
func (_CreditFilter *CreditFilterCaller) AllowedTokens(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CreditFilter.contract.Call(opts, &out, "allowedTokens", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllowedTokens is a free data retrieval call binding the contract method 0x5e5f2e26.
//
// Solidity: function allowedTokens(uint256 ) view returns(address)
func (_CreditFilter *CreditFilterSession) AllowedTokens(arg0 *big.Int) (common.Address, error) {
	return _CreditFilter.Contract.AllowedTokens(&_CreditFilter.CallOpts, arg0)
}

// AllowedTokens is a free data retrieval call binding the contract method 0x5e5f2e26.
//
// Solidity: function allowedTokens(uint256 ) view returns(address)
func (_CreditFilter *CreditFilterCallerSession) AllowedTokens(arg0 *big.Int) (common.Address, error) {
	return _CreditFilter.Contract.AllowedTokens(&_CreditFilter.CallOpts, arg0)
}

// AllowedTokensCount is a free data retrieval call binding the contract method 0x20a05ff7.
//
// Solidity: function allowedTokensCount() view returns(uint256)
func (_CreditFilter *CreditFilterCaller) AllowedTokensCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CreditFilter.contract.Call(opts, &out, "allowedTokensCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllowedTokensCount is a free data retrieval call binding the contract method 0x20a05ff7.
//
// Solidity: function allowedTokensCount() view returns(uint256)
func (_CreditFilter *CreditFilterSession) AllowedTokensCount() (*big.Int, error) {
	return _CreditFilter.Contract.AllowedTokensCount(&_CreditFilter.CallOpts)
}

// AllowedTokensCount is a free data retrieval call binding the contract method 0x20a05ff7.
//
// Solidity: function allowedTokensCount() view returns(uint256)
func (_CreditFilter *CreditFilterCallerSession) AllowedTokensCount() (*big.Int, error) {
	return _CreditFilter.Contract.AllowedTokensCount(&_CreditFilter.CallOpts)
}

// CalcCreditAccountAccruedInterest is a free data retrieval call binding the contract method 0x3192195c.
//
// Solidity: function calcCreditAccountAccruedInterest(address creditAccount) view returns(uint256)
func (_CreditFilter *CreditFilterCaller) CalcCreditAccountAccruedInterest(opts *bind.CallOpts, creditAccount common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CreditFilter.contract.Call(opts, &out, "calcCreditAccountAccruedInterest", creditAccount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcCreditAccountAccruedInterest is a free data retrieval call binding the contract method 0x3192195c.
//
// Solidity: function calcCreditAccountAccruedInterest(address creditAccount) view returns(uint256)
func (_CreditFilter *CreditFilterSession) CalcCreditAccountAccruedInterest(creditAccount common.Address) (*big.Int, error) {
	return _CreditFilter.Contract.CalcCreditAccountAccruedInterest(&_CreditFilter.CallOpts, creditAccount)
}

// CalcCreditAccountAccruedInterest is a free data retrieval call binding the contract method 0x3192195c.
//
// Solidity: function calcCreditAccountAccruedInterest(address creditAccount) view returns(uint256)
func (_CreditFilter *CreditFilterCallerSession) CalcCreditAccountAccruedInterest(creditAccount common.Address) (*big.Int, error) {
	return _CreditFilter.Contract.CalcCreditAccountAccruedInterest(&_CreditFilter.CallOpts, creditAccount)
}

// CalcCreditAccountHealthFactor is a free data retrieval call binding the contract method 0xdfd59465.
//
// Solidity: function calcCreditAccountHealthFactor(address creditAccount) view returns(uint256)
func (_CreditFilter *CreditFilterCaller) CalcCreditAccountHealthFactor(opts *bind.CallOpts, creditAccount common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CreditFilter.contract.Call(opts, &out, "calcCreditAccountHealthFactor", creditAccount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcCreditAccountHealthFactor is a free data retrieval call binding the contract method 0xdfd59465.
//
// Solidity: function calcCreditAccountHealthFactor(address creditAccount) view returns(uint256)
func (_CreditFilter *CreditFilterSession) CalcCreditAccountHealthFactor(creditAccount common.Address) (*big.Int, error) {
	return _CreditFilter.Contract.CalcCreditAccountHealthFactor(&_CreditFilter.CallOpts, creditAccount)
}

// CalcCreditAccountHealthFactor is a free data retrieval call binding the contract method 0xdfd59465.
//
// Solidity: function calcCreditAccountHealthFactor(address creditAccount) view returns(uint256)
func (_CreditFilter *CreditFilterCallerSession) CalcCreditAccountHealthFactor(creditAccount common.Address) (*big.Int, error) {
	return _CreditFilter.Contract.CalcCreditAccountHealthFactor(&_CreditFilter.CallOpts, creditAccount)
}

// CalcMaxPossibleDrop is a free data retrieval call binding the contract method 0xb3c61943.
//
// Solidity: function calcMaxPossibleDrop(uint256 percentage, uint256 times) pure returns(uint256 value)
func (_CreditFilter *CreditFilterCaller) CalcMaxPossibleDrop(opts *bind.CallOpts, percentage *big.Int, times *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CreditFilter.contract.Call(opts, &out, "calcMaxPossibleDrop", percentage, times)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcMaxPossibleDrop is a free data retrieval call binding the contract method 0xb3c61943.
//
// Solidity: function calcMaxPossibleDrop(uint256 percentage, uint256 times) pure returns(uint256 value)
func (_CreditFilter *CreditFilterSession) CalcMaxPossibleDrop(percentage *big.Int, times *big.Int) (*big.Int, error) {
	return _CreditFilter.Contract.CalcMaxPossibleDrop(&_CreditFilter.CallOpts, percentage, times)
}

// CalcMaxPossibleDrop is a free data retrieval call binding the contract method 0xb3c61943.
//
// Solidity: function calcMaxPossibleDrop(uint256 percentage, uint256 times) pure returns(uint256 value)
func (_CreditFilter *CreditFilterCallerSession) CalcMaxPossibleDrop(percentage *big.Int, times *big.Int) (*big.Int, error) {
	return _CreditFilter.Contract.CalcMaxPossibleDrop(&_CreditFilter.CallOpts, percentage, times)
}

// CalcThresholdWeightedValue is a free data retrieval call binding the contract method 0x90b1300a.
//
// Solidity: function calcThresholdWeightedValue(address creditAccount) view returns(uint256 total)
func (_CreditFilter *CreditFilterCaller) CalcThresholdWeightedValue(opts *bind.CallOpts, creditAccount common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CreditFilter.contract.Call(opts, &out, "calcThresholdWeightedValue", creditAccount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcThresholdWeightedValue is a free data retrieval call binding the contract method 0x90b1300a.
//
// Solidity: function calcThresholdWeightedValue(address creditAccount) view returns(uint256 total)
func (_CreditFilter *CreditFilterSession) CalcThresholdWeightedValue(creditAccount common.Address) (*big.Int, error) {
	return _CreditFilter.Contract.CalcThresholdWeightedValue(&_CreditFilter.CallOpts, creditAccount)
}

// CalcThresholdWeightedValue is a free data retrieval call binding the contract method 0x90b1300a.
//
// Solidity: function calcThresholdWeightedValue(address creditAccount) view returns(uint256 total)
func (_CreditFilter *CreditFilterCallerSession) CalcThresholdWeightedValue(creditAccount common.Address) (*big.Int, error) {
	return _CreditFilter.Contract.CalcThresholdWeightedValue(&_CreditFilter.CallOpts, creditAccount)
}

// CalcTotalValue is a free data retrieval call binding the contract method 0xc7de38a6.
//
// Solidity: function calcTotalValue(address creditAccount) view returns(uint256 total)
func (_CreditFilter *CreditFilterCaller) CalcTotalValue(opts *bind.CallOpts, creditAccount common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CreditFilter.contract.Call(opts, &out, "calcTotalValue", creditAccount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcTotalValue is a free data retrieval call binding the contract method 0xc7de38a6.
//
// Solidity: function calcTotalValue(address creditAccount) view returns(uint256 total)
func (_CreditFilter *CreditFilterSession) CalcTotalValue(creditAccount common.Address) (*big.Int, error) {
	return _CreditFilter.Contract.CalcTotalValue(&_CreditFilter.CallOpts, creditAccount)
}

// CalcTotalValue is a free data retrieval call binding the contract method 0xc7de38a6.
//
// Solidity: function calcTotalValue(address creditAccount) view returns(uint256 total)
func (_CreditFilter *CreditFilterCallerSession) CalcTotalValue(creditAccount common.Address) (*big.Int, error) {
	return _CreditFilter.Contract.CalcTotalValue(&_CreditFilter.CallOpts, creditAccount)
}

// ChiThreshold is a free data retrieval call binding the contract method 0x47dedfc9.
//
// Solidity: function chiThreshold() view returns(uint256)
func (_CreditFilter *CreditFilterCaller) ChiThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CreditFilter.contract.Call(opts, &out, "chiThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChiThreshold is a free data retrieval call binding the contract method 0x47dedfc9.
//
// Solidity: function chiThreshold() view returns(uint256)
func (_CreditFilter *CreditFilterSession) ChiThreshold() (*big.Int, error) {
	return _CreditFilter.Contract.ChiThreshold(&_CreditFilter.CallOpts)
}

// ChiThreshold is a free data retrieval call binding the contract method 0x47dedfc9.
//
// Solidity: function chiThreshold() view returns(uint256)
func (_CreditFilter *CreditFilterCallerSession) ChiThreshold() (*big.Int, error) {
	return _CreditFilter.Contract.ChiThreshold(&_CreditFilter.CallOpts)
}

// ContractToAdapter is a free data retrieval call binding the contract method 0xfdd57645.
//
// Solidity: function contractToAdapter(address ) view returns(address)
func (_CreditFilter *CreditFilterCaller) ContractToAdapter(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _CreditFilter.contract.Call(opts, &out, "contractToAdapter", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ContractToAdapter is a free data retrieval call binding the contract method 0xfdd57645.
//
// Solidity: function contractToAdapter(address ) view returns(address)
func (_CreditFilter *CreditFilterSession) ContractToAdapter(arg0 common.Address) (common.Address, error) {
	return _CreditFilter.Contract.ContractToAdapter(&_CreditFilter.CallOpts, arg0)
}

// ContractToAdapter is a free data retrieval call binding the contract method 0xfdd57645.
//
// Solidity: function contractToAdapter(address ) view returns(address)
func (_CreditFilter *CreditFilterCallerSession) ContractToAdapter(arg0 common.Address) (common.Address, error) {
	return _CreditFilter.Contract.ContractToAdapter(&_CreditFilter.CallOpts, arg0)
}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_CreditFilter *CreditFilterCaller) CreditManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditFilter.contract.Call(opts, &out, "creditManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_CreditFilter *CreditFilterSession) CreditManager() (common.Address, error) {
	return _CreditFilter.Contract.CreditManager(&_CreditFilter.CallOpts)
}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_CreditFilter *CreditFilterCallerSession) CreditManager() (common.Address, error) {
	return _CreditFilter.Contract.CreditManager(&_CreditFilter.CallOpts)
}

// EnabledTokens is a free data retrieval call binding the contract method 0xb451cecc.
//
// Solidity: function enabledTokens(address ) view returns(uint256)
func (_CreditFilter *CreditFilterCaller) EnabledTokens(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CreditFilter.contract.Call(opts, &out, "enabledTokens", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EnabledTokens is a free data retrieval call binding the contract method 0xb451cecc.
//
// Solidity: function enabledTokens(address ) view returns(uint256)
func (_CreditFilter *CreditFilterSession) EnabledTokens(arg0 common.Address) (*big.Int, error) {
	return _CreditFilter.Contract.EnabledTokens(&_CreditFilter.CallOpts, arg0)
}

// EnabledTokens is a free data retrieval call binding the contract method 0xb451cecc.
//
// Solidity: function enabledTokens(address ) view returns(uint256)
func (_CreditFilter *CreditFilterCallerSession) EnabledTokens(arg0 common.Address) (*big.Int, error) {
	return _CreditFilter.Contract.EnabledTokens(&_CreditFilter.CallOpts, arg0)
}

// FastCheckCounter is a free data retrieval call binding the contract method 0x4cba294a.
//
// Solidity: function fastCheckCounter(address ) view returns(uint256)
func (_CreditFilter *CreditFilterCaller) FastCheckCounter(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CreditFilter.contract.Call(opts, &out, "fastCheckCounter", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FastCheckCounter is a free data retrieval call binding the contract method 0x4cba294a.
//
// Solidity: function fastCheckCounter(address ) view returns(uint256)
func (_CreditFilter *CreditFilterSession) FastCheckCounter(arg0 common.Address) (*big.Int, error) {
	return _CreditFilter.Contract.FastCheckCounter(&_CreditFilter.CallOpts, arg0)
}

// FastCheckCounter is a free data retrieval call binding the contract method 0x4cba294a.
//
// Solidity: function fastCheckCounter(address ) view returns(uint256)
func (_CreditFilter *CreditFilterCallerSession) FastCheckCounter(arg0 common.Address) (*big.Int, error) {
	return _CreditFilter.Contract.FastCheckCounter(&_CreditFilter.CallOpts, arg0)
}

// GetCreditAccountTokenById is a free data retrieval call binding the contract method 0xaf0a6502.
//
// Solidity: function getCreditAccountTokenById(address creditAccount, uint256 id) view returns(address token, uint256 balance, uint256 tv, uint256 tvw)
func (_CreditFilter *CreditFilterCaller) GetCreditAccountTokenById(opts *bind.CallOpts, creditAccount common.Address, id *big.Int) (struct {
	Token   common.Address
	Balance *big.Int
	Tv      *big.Int
	Tvw     *big.Int
}, error) {
	var out []interface{}
	err := _CreditFilter.contract.Call(opts, &out, "getCreditAccountTokenById", creditAccount, id)

	outstruct := new(struct {
		Token   common.Address
		Balance *big.Int
		Tv      *big.Int
		Tvw     *big.Int
	})

	outstruct.Token = out[0].(common.Address)
	outstruct.Balance = out[1].(*big.Int)
	outstruct.Tv = out[2].(*big.Int)
	outstruct.Tvw = out[3].(*big.Int)

	return *outstruct, err

}

// GetCreditAccountTokenById is a free data retrieval call binding the contract method 0xaf0a6502.
//
// Solidity: function getCreditAccountTokenById(address creditAccount, uint256 id) view returns(address token, uint256 balance, uint256 tv, uint256 tvw)
func (_CreditFilter *CreditFilterSession) GetCreditAccountTokenById(creditAccount common.Address, id *big.Int) (struct {
	Token   common.Address
	Balance *big.Int
	Tv      *big.Int
	Tvw     *big.Int
}, error) {
	return _CreditFilter.Contract.GetCreditAccountTokenById(&_CreditFilter.CallOpts, creditAccount, id)
}

// GetCreditAccountTokenById is a free data retrieval call binding the contract method 0xaf0a6502.
//
// Solidity: function getCreditAccountTokenById(address creditAccount, uint256 id) view returns(address token, uint256 balance, uint256 tv, uint256 tvw)
func (_CreditFilter *CreditFilterCallerSession) GetCreditAccountTokenById(creditAccount common.Address, id *big.Int) (struct {
	Token   common.Address
	Balance *big.Int
	Tv      *big.Int
	Tvw     *big.Int
}, error) {
	return _CreditFilter.Contract.GetCreditAccountTokenById(&_CreditFilter.CallOpts, creditAccount, id)
}

// HfCheckInterval is a free data retrieval call binding the contract method 0xe6dee2cc.
//
// Solidity: function hfCheckInterval() view returns(uint256)
func (_CreditFilter *CreditFilterCaller) HfCheckInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CreditFilter.contract.Call(opts, &out, "hfCheckInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HfCheckInterval is a free data retrieval call binding the contract method 0xe6dee2cc.
//
// Solidity: function hfCheckInterval() view returns(uint256)
func (_CreditFilter *CreditFilterSession) HfCheckInterval() (*big.Int, error) {
	return _CreditFilter.Contract.HfCheckInterval(&_CreditFilter.CallOpts)
}

// HfCheckInterval is a free data retrieval call binding the contract method 0xe6dee2cc.
//
// Solidity: function hfCheckInterval() view returns(uint256)
func (_CreditFilter *CreditFilterCallerSession) HfCheckInterval() (*big.Int, error) {
	return _CreditFilter.Contract.HfCheckInterval(&_CreditFilter.CallOpts)
}

// IsTokenAllowed is a free data retrieval call binding the contract method 0xf9eaee0d.
//
// Solidity: function isTokenAllowed(address ) view returns(bool)
func (_CreditFilter *CreditFilterCaller) IsTokenAllowed(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _CreditFilter.contract.Call(opts, &out, "isTokenAllowed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTokenAllowed is a free data retrieval call binding the contract method 0xf9eaee0d.
//
// Solidity: function isTokenAllowed(address ) view returns(bool)
func (_CreditFilter *CreditFilterSession) IsTokenAllowed(arg0 common.Address) (bool, error) {
	return _CreditFilter.Contract.IsTokenAllowed(&_CreditFilter.CallOpts, arg0)
}

// IsTokenAllowed is a free data retrieval call binding the contract method 0xf9eaee0d.
//
// Solidity: function isTokenAllowed(address ) view returns(bool)
func (_CreditFilter *CreditFilterCallerSession) IsTokenAllowed(arg0 common.Address) (bool, error) {
	return _CreditFilter.Contract.IsTokenAllowed(&_CreditFilter.CallOpts, arg0)
}

// LiquidationThresholds is a free data retrieval call binding the contract method 0x78327438.
//
// Solidity: function liquidationThresholds(address ) view returns(uint256)
func (_CreditFilter *CreditFilterCaller) LiquidationThresholds(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CreditFilter.contract.Call(opts, &out, "liquidationThresholds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LiquidationThresholds is a free data retrieval call binding the contract method 0x78327438.
//
// Solidity: function liquidationThresholds(address ) view returns(uint256)
func (_CreditFilter *CreditFilterSession) LiquidationThresholds(arg0 common.Address) (*big.Int, error) {
	return _CreditFilter.Contract.LiquidationThresholds(&_CreditFilter.CallOpts, arg0)
}

// LiquidationThresholds is a free data retrieval call binding the contract method 0x78327438.
//
// Solidity: function liquidationThresholds(address ) view returns(uint256)
func (_CreditFilter *CreditFilterCallerSession) LiquidationThresholds(arg0 common.Address) (*big.Int, error) {
	return _CreditFilter.Contract.LiquidationThresholds(&_CreditFilter.CallOpts, arg0)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CreditFilter *CreditFilterCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CreditFilter.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CreditFilter *CreditFilterSession) Paused() (bool, error) {
	return _CreditFilter.Contract.Paused(&_CreditFilter.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CreditFilter *CreditFilterCallerSession) Paused() (bool, error) {
	return _CreditFilter.Contract.Paused(&_CreditFilter.CallOpts)
}

// PoolService is a free data retrieval call binding the contract method 0x570a7af2.
//
// Solidity: function poolService() view returns(address)
func (_CreditFilter *CreditFilterCaller) PoolService(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditFilter.contract.Call(opts, &out, "poolService")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolService is a free data retrieval call binding the contract method 0x570a7af2.
//
// Solidity: function poolService() view returns(address)
func (_CreditFilter *CreditFilterSession) PoolService() (common.Address, error) {
	return _CreditFilter.Contract.PoolService(&_CreditFilter.CallOpts)
}

// PoolService is a free data retrieval call binding the contract method 0x570a7af2.
//
// Solidity: function poolService() view returns(address)
func (_CreditFilter *CreditFilterCallerSession) PoolService() (common.Address, error) {
	return _CreditFilter.Contract.PoolService(&_CreditFilter.CallOpts)
}

// PriceOracle is a free data retrieval call binding the contract method 0x2630c12f.
//
// Solidity: function priceOracle() view returns(address)
func (_CreditFilter *CreditFilterCaller) PriceOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditFilter.contract.Call(opts, &out, "priceOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PriceOracle is a free data retrieval call binding the contract method 0x2630c12f.
//
// Solidity: function priceOracle() view returns(address)
func (_CreditFilter *CreditFilterSession) PriceOracle() (common.Address, error) {
	return _CreditFilter.Contract.PriceOracle(&_CreditFilter.CallOpts)
}

// PriceOracle is a free data retrieval call binding the contract method 0x2630c12f.
//
// Solidity: function priceOracle() view returns(address)
func (_CreditFilter *CreditFilterCallerSession) PriceOracle() (common.Address, error) {
	return _CreditFilter.Contract.PriceOracle(&_CreditFilter.CallOpts)
}

// RevertIfAccountTransferIsNotAllowed is a free data retrieval call binding the contract method 0x3b00ae70.
//
// Solidity: function revertIfAccountTransferIsNotAllowed(address owner, address newOwner) view returns()
func (_CreditFilter *CreditFilterCaller) RevertIfAccountTransferIsNotAllowed(opts *bind.CallOpts, owner common.Address, newOwner common.Address) error {
	var out []interface{}
	err := _CreditFilter.contract.Call(opts, &out, "revertIfAccountTransferIsNotAllowed", owner, newOwner)

	if err != nil {
		return err
	}

	return err

}

// RevertIfAccountTransferIsNotAllowed is a free data retrieval call binding the contract method 0x3b00ae70.
//
// Solidity: function revertIfAccountTransferIsNotAllowed(address owner, address newOwner) view returns()
func (_CreditFilter *CreditFilterSession) RevertIfAccountTransferIsNotAllowed(owner common.Address, newOwner common.Address) error {
	return _CreditFilter.Contract.RevertIfAccountTransferIsNotAllowed(&_CreditFilter.CallOpts, owner, newOwner)
}

// RevertIfAccountTransferIsNotAllowed is a free data retrieval call binding the contract method 0x3b00ae70.
//
// Solidity: function revertIfAccountTransferIsNotAllowed(address owner, address newOwner) view returns()
func (_CreditFilter *CreditFilterCallerSession) RevertIfAccountTransferIsNotAllowed(owner common.Address, newOwner common.Address) error {
	return _CreditFilter.Contract.RevertIfAccountTransferIsNotAllowed(&_CreditFilter.CallOpts, owner, newOwner)
}

// RevertIfCantIncreaseBorrowing is a free data retrieval call binding the contract method 0xa5757517.
//
// Solidity: function revertIfCantIncreaseBorrowing(address creditAccount, uint256 minHealthFactor) view returns()
func (_CreditFilter *CreditFilterCaller) RevertIfCantIncreaseBorrowing(opts *bind.CallOpts, creditAccount common.Address, minHealthFactor *big.Int) error {
	var out []interface{}
	err := _CreditFilter.contract.Call(opts, &out, "revertIfCantIncreaseBorrowing", creditAccount, minHealthFactor)

	if err != nil {
		return err
	}

	return err

}

// RevertIfCantIncreaseBorrowing is a free data retrieval call binding the contract method 0xa5757517.
//
// Solidity: function revertIfCantIncreaseBorrowing(address creditAccount, uint256 minHealthFactor) view returns()
func (_CreditFilter *CreditFilterSession) RevertIfCantIncreaseBorrowing(creditAccount common.Address, minHealthFactor *big.Int) error {
	return _CreditFilter.Contract.RevertIfCantIncreaseBorrowing(&_CreditFilter.CallOpts, creditAccount, minHealthFactor)
}

// RevertIfCantIncreaseBorrowing is a free data retrieval call binding the contract method 0xa5757517.
//
// Solidity: function revertIfCantIncreaseBorrowing(address creditAccount, uint256 minHealthFactor) view returns()
func (_CreditFilter *CreditFilterCallerSession) RevertIfCantIncreaseBorrowing(creditAccount common.Address, minHealthFactor *big.Int) error {
	return _CreditFilter.Contract.RevertIfCantIncreaseBorrowing(&_CreditFilter.CallOpts, creditAccount, minHealthFactor)
}

// RevertIfTokenNotAllowed is a free data retrieval call binding the contract method 0x7dd0ba82.
//
// Solidity: function revertIfTokenNotAllowed(address token) view returns()
func (_CreditFilter *CreditFilterCaller) RevertIfTokenNotAllowed(opts *bind.CallOpts, token common.Address) error {
	var out []interface{}
	err := _CreditFilter.contract.Call(opts, &out, "revertIfTokenNotAllowed", token)

	if err != nil {
		return err
	}

	return err

}

// RevertIfTokenNotAllowed is a free data retrieval call binding the contract method 0x7dd0ba82.
//
// Solidity: function revertIfTokenNotAllowed(address token) view returns()
func (_CreditFilter *CreditFilterSession) RevertIfTokenNotAllowed(token common.Address) error {
	return _CreditFilter.Contract.RevertIfTokenNotAllowed(&_CreditFilter.CallOpts, token)
}

// RevertIfTokenNotAllowed is a free data retrieval call binding the contract method 0x7dd0ba82.
//
// Solidity: function revertIfTokenNotAllowed(address token) view returns()
func (_CreditFilter *CreditFilterCallerSession) RevertIfTokenNotAllowed(token common.Address) error {
	return _CreditFilter.Contract.RevertIfTokenNotAllowed(&_CreditFilter.CallOpts, token)
}

// TokenMasksMap is a free data retrieval call binding the contract method 0xf67c5bd0.
//
// Solidity: function tokenMasksMap(address ) view returns(uint256)
func (_CreditFilter *CreditFilterCaller) TokenMasksMap(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CreditFilter.contract.Call(opts, &out, "tokenMasksMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenMasksMap is a free data retrieval call binding the contract method 0xf67c5bd0.
//
// Solidity: function tokenMasksMap(address ) view returns(uint256)
func (_CreditFilter *CreditFilterSession) TokenMasksMap(arg0 common.Address) (*big.Int, error) {
	return _CreditFilter.Contract.TokenMasksMap(&_CreditFilter.CallOpts, arg0)
}

// TokenMasksMap is a free data retrieval call binding the contract method 0xf67c5bd0.
//
// Solidity: function tokenMasksMap(address ) view returns(uint256)
func (_CreditFilter *CreditFilterCallerSession) TokenMasksMap(arg0 common.Address) (*big.Int, error) {
	return _CreditFilter.Contract.TokenMasksMap(&_CreditFilter.CallOpts, arg0)
}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_CreditFilter *CreditFilterCaller) UnderlyingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditFilter.contract.Call(opts, &out, "underlyingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_CreditFilter *CreditFilterSession) UnderlyingToken() (common.Address, error) {
	return _CreditFilter.Contract.UnderlyingToken(&_CreditFilter.CallOpts)
}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_CreditFilter *CreditFilterCallerSession) UnderlyingToken() (common.Address, error) {
	return _CreditFilter.Contract.UnderlyingToken(&_CreditFilter.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_CreditFilter *CreditFilterCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CreditFilter.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_CreditFilter *CreditFilterSession) Version() (*big.Int, error) {
	return _CreditFilter.Contract.Version(&_CreditFilter.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_CreditFilter *CreditFilterCallerSession) Version() (*big.Int, error) {
	return _CreditFilter.Contract.Version(&_CreditFilter.CallOpts)
}

// WethAddress is a free data retrieval call binding the contract method 0x4f0e0ef3.
//
// Solidity: function wethAddress() view returns(address)
func (_CreditFilter *CreditFilterCaller) WethAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditFilter.contract.Call(opts, &out, "wethAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WethAddress is a free data retrieval call binding the contract method 0x4f0e0ef3.
//
// Solidity: function wethAddress() view returns(address)
func (_CreditFilter *CreditFilterSession) WethAddress() (common.Address, error) {
	return _CreditFilter.Contract.WethAddress(&_CreditFilter.CallOpts)
}

// WethAddress is a free data retrieval call binding the contract method 0x4f0e0ef3.
//
// Solidity: function wethAddress() view returns(address)
func (_CreditFilter *CreditFilterCallerSession) WethAddress() (common.Address, error) {
	return _CreditFilter.Contract.WethAddress(&_CreditFilter.CallOpts)
}

// AllowContract is a paid mutator transaction binding the contract method 0x7bccacee.
//
// Solidity: function allowContract(address targetContract, address adapter) returns()
func (_CreditFilter *CreditFilterTransactor) AllowContract(opts *bind.TransactOpts, targetContract common.Address, adapter common.Address) (*types.Transaction, error) {
	return _CreditFilter.contract.Transact(opts, "allowContract", targetContract, adapter)
}

// AllowContract is a paid mutator transaction binding the contract method 0x7bccacee.
//
// Solidity: function allowContract(address targetContract, address adapter) returns()
func (_CreditFilter *CreditFilterSession) AllowContract(targetContract common.Address, adapter common.Address) (*types.Transaction, error) {
	return _CreditFilter.Contract.AllowContract(&_CreditFilter.TransactOpts, targetContract, adapter)
}

// AllowContract is a paid mutator transaction binding the contract method 0x7bccacee.
//
// Solidity: function allowContract(address targetContract, address adapter) returns()
func (_CreditFilter *CreditFilterTransactorSession) AllowContract(targetContract common.Address, adapter common.Address) (*types.Transaction, error) {
	return _CreditFilter.Contract.AllowContract(&_CreditFilter.TransactOpts, targetContract, adapter)
}

// AllowPlugin is a paid mutator transaction binding the contract method 0x2e2986dd.
//
// Solidity: function allowPlugin(address plugin, bool state) returns()
func (_CreditFilter *CreditFilterTransactor) AllowPlugin(opts *bind.TransactOpts, plugin common.Address, state bool) (*types.Transaction, error) {
	return _CreditFilter.contract.Transact(opts, "allowPlugin", plugin, state)
}

// AllowPlugin is a paid mutator transaction binding the contract method 0x2e2986dd.
//
// Solidity: function allowPlugin(address plugin, bool state) returns()
func (_CreditFilter *CreditFilterSession) AllowPlugin(plugin common.Address, state bool) (*types.Transaction, error) {
	return _CreditFilter.Contract.AllowPlugin(&_CreditFilter.TransactOpts, plugin, state)
}

// AllowPlugin is a paid mutator transaction binding the contract method 0x2e2986dd.
//
// Solidity: function allowPlugin(address plugin, bool state) returns()
func (_CreditFilter *CreditFilterTransactorSession) AllowPlugin(plugin common.Address, state bool) (*types.Transaction, error) {
	return _CreditFilter.Contract.AllowPlugin(&_CreditFilter.TransactOpts, plugin, state)
}

// AllowToken is a paid mutator transaction binding the contract method 0xa147c6c6.
//
// Solidity: function allowToken(address token, uint256 liquidationThreshold) returns()
func (_CreditFilter *CreditFilterTransactor) AllowToken(opts *bind.TransactOpts, token common.Address, liquidationThreshold *big.Int) (*types.Transaction, error) {
	return _CreditFilter.contract.Transact(opts, "allowToken", token, liquidationThreshold)
}

// AllowToken is a paid mutator transaction binding the contract method 0xa147c6c6.
//
// Solidity: function allowToken(address token, uint256 liquidationThreshold) returns()
func (_CreditFilter *CreditFilterSession) AllowToken(token common.Address, liquidationThreshold *big.Int) (*types.Transaction, error) {
	return _CreditFilter.Contract.AllowToken(&_CreditFilter.TransactOpts, token, liquidationThreshold)
}

// AllowToken is a paid mutator transaction binding the contract method 0xa147c6c6.
//
// Solidity: function allowToken(address token, uint256 liquidationThreshold) returns()
func (_CreditFilter *CreditFilterTransactorSession) AllowToken(token common.Address, liquidationThreshold *big.Int) (*types.Transaction, error) {
	return _CreditFilter.Contract.AllowToken(&_CreditFilter.TransactOpts, token, liquidationThreshold)
}

// ApproveAccountTransfers is a paid mutator transaction binding the contract method 0x5f27212a.
//
// Solidity: function approveAccountTransfers(address from, bool state) returns()
func (_CreditFilter *CreditFilterTransactor) ApproveAccountTransfers(opts *bind.TransactOpts, from common.Address, state bool) (*types.Transaction, error) {
	return _CreditFilter.contract.Transact(opts, "approveAccountTransfers", from, state)
}

// ApproveAccountTransfers is a paid mutator transaction binding the contract method 0x5f27212a.
//
// Solidity: function approveAccountTransfers(address from, bool state) returns()
func (_CreditFilter *CreditFilterSession) ApproveAccountTransfers(from common.Address, state bool) (*types.Transaction, error) {
	return _CreditFilter.Contract.ApproveAccountTransfers(&_CreditFilter.TransactOpts, from, state)
}

// ApproveAccountTransfers is a paid mutator transaction binding the contract method 0x5f27212a.
//
// Solidity: function approveAccountTransfers(address from, bool state) returns()
func (_CreditFilter *CreditFilterTransactorSession) ApproveAccountTransfers(from common.Address, state bool) (*types.Transaction, error) {
	return _CreditFilter.Contract.ApproveAccountTransfers(&_CreditFilter.TransactOpts, from, state)
}

// CheckAndEnableToken is a paid mutator transaction binding the contract method 0x51e3f160.
//
// Solidity: function checkAndEnableToken(address creditAccount, address token) returns()
func (_CreditFilter *CreditFilterTransactor) CheckAndEnableToken(opts *bind.TransactOpts, creditAccount common.Address, token common.Address) (*types.Transaction, error) {
	return _CreditFilter.contract.Transact(opts, "checkAndEnableToken", creditAccount, token)
}

// CheckAndEnableToken is a paid mutator transaction binding the contract method 0x51e3f160.
//
// Solidity: function checkAndEnableToken(address creditAccount, address token) returns()
func (_CreditFilter *CreditFilterSession) CheckAndEnableToken(creditAccount common.Address, token common.Address) (*types.Transaction, error) {
	return _CreditFilter.Contract.CheckAndEnableToken(&_CreditFilter.TransactOpts, creditAccount, token)
}

// CheckAndEnableToken is a paid mutator transaction binding the contract method 0x51e3f160.
//
// Solidity: function checkAndEnableToken(address creditAccount, address token) returns()
func (_CreditFilter *CreditFilterTransactorSession) CheckAndEnableToken(creditAccount common.Address, token common.Address) (*types.Transaction, error) {
	return _CreditFilter.Contract.CheckAndEnableToken(&_CreditFilter.TransactOpts, creditAccount, token)
}

// CheckCollateralChange is a paid mutator transaction binding the contract method 0xe1c8ef0d.
//
// Solidity: function checkCollateralChange(address creditAccount, address tokenIn, address tokenOut, uint256 amountIn, uint256 amountOut) returns()
func (_CreditFilter *CreditFilterTransactor) CheckCollateralChange(opts *bind.TransactOpts, creditAccount common.Address, tokenIn common.Address, tokenOut common.Address, amountIn *big.Int, amountOut *big.Int) (*types.Transaction, error) {
	return _CreditFilter.contract.Transact(opts, "checkCollateralChange", creditAccount, tokenIn, tokenOut, amountIn, amountOut)
}

// CheckCollateralChange is a paid mutator transaction binding the contract method 0xe1c8ef0d.
//
// Solidity: function checkCollateralChange(address creditAccount, address tokenIn, address tokenOut, uint256 amountIn, uint256 amountOut) returns()
func (_CreditFilter *CreditFilterSession) CheckCollateralChange(creditAccount common.Address, tokenIn common.Address, tokenOut common.Address, amountIn *big.Int, amountOut *big.Int) (*types.Transaction, error) {
	return _CreditFilter.Contract.CheckCollateralChange(&_CreditFilter.TransactOpts, creditAccount, tokenIn, tokenOut, amountIn, amountOut)
}

// CheckCollateralChange is a paid mutator transaction binding the contract method 0xe1c8ef0d.
//
// Solidity: function checkCollateralChange(address creditAccount, address tokenIn, address tokenOut, uint256 amountIn, uint256 amountOut) returns()
func (_CreditFilter *CreditFilterTransactorSession) CheckCollateralChange(creditAccount common.Address, tokenIn common.Address, tokenOut common.Address, amountIn *big.Int, amountOut *big.Int) (*types.Transaction, error) {
	return _CreditFilter.Contract.CheckCollateralChange(&_CreditFilter.TransactOpts, creditAccount, tokenIn, tokenOut, amountIn, amountOut)
}

// CheckMultiTokenCollateral is a paid mutator transaction binding the contract method 0x7e4a6863.
//
// Solidity: function checkMultiTokenCollateral(address creditAccount, uint256[] amountIn, uint256[] amountOut, address[] tokenIn, address[] tokenOut) returns()
func (_CreditFilter *CreditFilterTransactor) CheckMultiTokenCollateral(opts *bind.TransactOpts, creditAccount common.Address, amountIn []*big.Int, amountOut []*big.Int, tokenIn []common.Address, tokenOut []common.Address) (*types.Transaction, error) {
	return _CreditFilter.contract.Transact(opts, "checkMultiTokenCollateral", creditAccount, amountIn, amountOut, tokenIn, tokenOut)
}

// CheckMultiTokenCollateral is a paid mutator transaction binding the contract method 0x7e4a6863.
//
// Solidity: function checkMultiTokenCollateral(address creditAccount, uint256[] amountIn, uint256[] amountOut, address[] tokenIn, address[] tokenOut) returns()
func (_CreditFilter *CreditFilterSession) CheckMultiTokenCollateral(creditAccount common.Address, amountIn []*big.Int, amountOut []*big.Int, tokenIn []common.Address, tokenOut []common.Address) (*types.Transaction, error) {
	return _CreditFilter.Contract.CheckMultiTokenCollateral(&_CreditFilter.TransactOpts, creditAccount, amountIn, amountOut, tokenIn, tokenOut)
}

// CheckMultiTokenCollateral is a paid mutator transaction binding the contract method 0x7e4a6863.
//
// Solidity: function checkMultiTokenCollateral(address creditAccount, uint256[] amountIn, uint256[] amountOut, address[] tokenIn, address[] tokenOut) returns()
func (_CreditFilter *CreditFilterTransactorSession) CheckMultiTokenCollateral(creditAccount common.Address, amountIn []*big.Int, amountOut []*big.Int, tokenIn []common.Address, tokenOut []common.Address) (*types.Transaction, error) {
	return _CreditFilter.Contract.CheckMultiTokenCollateral(&_CreditFilter.TransactOpts, creditAccount, amountIn, amountOut, tokenIn, tokenOut)
}

// ConnectCreditManager is a paid mutator transaction binding the contract method 0xcf33d955.
//
// Solidity: function connectCreditManager(address _creditManager) returns()
func (_CreditFilter *CreditFilterTransactor) ConnectCreditManager(opts *bind.TransactOpts, _creditManager common.Address) (*types.Transaction, error) {
	return _CreditFilter.contract.Transact(opts, "connectCreditManager", _creditManager)
}

// ConnectCreditManager is a paid mutator transaction binding the contract method 0xcf33d955.
//
// Solidity: function connectCreditManager(address _creditManager) returns()
func (_CreditFilter *CreditFilterSession) ConnectCreditManager(_creditManager common.Address) (*types.Transaction, error) {
	return _CreditFilter.Contract.ConnectCreditManager(&_CreditFilter.TransactOpts, _creditManager)
}

// ConnectCreditManager is a paid mutator transaction binding the contract method 0xcf33d955.
//
// Solidity: function connectCreditManager(address _creditManager) returns()
func (_CreditFilter *CreditFilterTransactorSession) ConnectCreditManager(_creditManager common.Address) (*types.Transaction, error) {
	return _CreditFilter.Contract.ConnectCreditManager(&_CreditFilter.TransactOpts, _creditManager)
}

// ForbidContract is a paid mutator transaction binding the contract method 0x52438e54.
//
// Solidity: function forbidContract(address targetContract) returns()
func (_CreditFilter *CreditFilterTransactor) ForbidContract(opts *bind.TransactOpts, targetContract common.Address) (*types.Transaction, error) {
	return _CreditFilter.contract.Transact(opts, "forbidContract", targetContract)
}

// ForbidContract is a paid mutator transaction binding the contract method 0x52438e54.
//
// Solidity: function forbidContract(address targetContract) returns()
func (_CreditFilter *CreditFilterSession) ForbidContract(targetContract common.Address) (*types.Transaction, error) {
	return _CreditFilter.Contract.ForbidContract(&_CreditFilter.TransactOpts, targetContract)
}

// ForbidContract is a paid mutator transaction binding the contract method 0x52438e54.
//
// Solidity: function forbidContract(address targetContract) returns()
func (_CreditFilter *CreditFilterTransactorSession) ForbidContract(targetContract common.Address) (*types.Transaction, error) {
	return _CreditFilter.Contract.ForbidContract(&_CreditFilter.TransactOpts, targetContract)
}

// ForbidToken is a paid mutator transaction binding the contract method 0x24147708.
//
// Solidity: function forbidToken(address token) returns()
func (_CreditFilter *CreditFilterTransactor) ForbidToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _CreditFilter.contract.Transact(opts, "forbidToken", token)
}

// ForbidToken is a paid mutator transaction binding the contract method 0x24147708.
//
// Solidity: function forbidToken(address token) returns()
func (_CreditFilter *CreditFilterSession) ForbidToken(token common.Address) (*types.Transaction, error) {
	return _CreditFilter.Contract.ForbidToken(&_CreditFilter.TransactOpts, token)
}

// ForbidToken is a paid mutator transaction binding the contract method 0x24147708.
//
// Solidity: function forbidToken(address token) returns()
func (_CreditFilter *CreditFilterTransactorSession) ForbidToken(token common.Address) (*types.Transaction, error) {
	return _CreditFilter.Contract.ForbidToken(&_CreditFilter.TransactOpts, token)
}

// InitEnabledTokens is a paid mutator transaction binding the contract method 0xe54fe9c8.
//
// Solidity: function initEnabledTokens(address creditAccount) returns()
func (_CreditFilter *CreditFilterTransactor) InitEnabledTokens(opts *bind.TransactOpts, creditAccount common.Address) (*types.Transaction, error) {
	return _CreditFilter.contract.Transact(opts, "initEnabledTokens", creditAccount)
}

// InitEnabledTokens is a paid mutator transaction binding the contract method 0xe54fe9c8.
//
// Solidity: function initEnabledTokens(address creditAccount) returns()
func (_CreditFilter *CreditFilterSession) InitEnabledTokens(creditAccount common.Address) (*types.Transaction, error) {
	return _CreditFilter.Contract.InitEnabledTokens(&_CreditFilter.TransactOpts, creditAccount)
}

// InitEnabledTokens is a paid mutator transaction binding the contract method 0xe54fe9c8.
//
// Solidity: function initEnabledTokens(address creditAccount) returns()
func (_CreditFilter *CreditFilterTransactorSession) InitEnabledTokens(creditAccount common.Address) (*types.Transaction, error) {
	return _CreditFilter.Contract.InitEnabledTokens(&_CreditFilter.TransactOpts, creditAccount)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CreditFilter *CreditFilterTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditFilter.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CreditFilter *CreditFilterSession) Pause() (*types.Transaction, error) {
	return _CreditFilter.Contract.Pause(&_CreditFilter.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CreditFilter *CreditFilterTransactorSession) Pause() (*types.Transaction, error) {
	return _CreditFilter.Contract.Pause(&_CreditFilter.TransactOpts)
}

// SetFastCheckParameters is a paid mutator transaction binding the contract method 0x62061c6d.
//
// Solidity: function setFastCheckParameters(uint256 _chiThreshold, uint256 _hfCheckInterval) returns()
func (_CreditFilter *CreditFilterTransactor) SetFastCheckParameters(opts *bind.TransactOpts, _chiThreshold *big.Int, _hfCheckInterval *big.Int) (*types.Transaction, error) {
	return _CreditFilter.contract.Transact(opts, "setFastCheckParameters", _chiThreshold, _hfCheckInterval)
}

// SetFastCheckParameters is a paid mutator transaction binding the contract method 0x62061c6d.
//
// Solidity: function setFastCheckParameters(uint256 _chiThreshold, uint256 _hfCheckInterval) returns()
func (_CreditFilter *CreditFilterSession) SetFastCheckParameters(_chiThreshold *big.Int, _hfCheckInterval *big.Int) (*types.Transaction, error) {
	return _CreditFilter.Contract.SetFastCheckParameters(&_CreditFilter.TransactOpts, _chiThreshold, _hfCheckInterval)
}

// SetFastCheckParameters is a paid mutator transaction binding the contract method 0x62061c6d.
//
// Solidity: function setFastCheckParameters(uint256 _chiThreshold, uint256 _hfCheckInterval) returns()
func (_CreditFilter *CreditFilterTransactorSession) SetFastCheckParameters(_chiThreshold *big.Int, _hfCheckInterval *big.Int) (*types.Transaction, error) {
	return _CreditFilter.Contract.SetFastCheckParameters(&_CreditFilter.TransactOpts, _chiThreshold, _hfCheckInterval)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CreditFilter *CreditFilterTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditFilter.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CreditFilter *CreditFilterSession) Unpause() (*types.Transaction, error) {
	return _CreditFilter.Contract.Unpause(&_CreditFilter.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CreditFilter *CreditFilterTransactorSession) Unpause() (*types.Transaction, error) {
	return _CreditFilter.Contract.Unpause(&_CreditFilter.TransactOpts)
}

// UpdateUnderlyingTokenLiquidationThreshold is a paid mutator transaction binding the contract method 0x40631828.
//
// Solidity: function updateUnderlyingTokenLiquidationThreshold() returns()
func (_CreditFilter *CreditFilterTransactor) UpdateUnderlyingTokenLiquidationThreshold(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditFilter.contract.Transact(opts, "updateUnderlyingTokenLiquidationThreshold")
}

// UpdateUnderlyingTokenLiquidationThreshold is a paid mutator transaction binding the contract method 0x40631828.
//
// Solidity: function updateUnderlyingTokenLiquidationThreshold() returns()
func (_CreditFilter *CreditFilterSession) UpdateUnderlyingTokenLiquidationThreshold() (*types.Transaction, error) {
	return _CreditFilter.Contract.UpdateUnderlyingTokenLiquidationThreshold(&_CreditFilter.TransactOpts)
}

// UpdateUnderlyingTokenLiquidationThreshold is a paid mutator transaction binding the contract method 0x40631828.
//
// Solidity: function updateUnderlyingTokenLiquidationThreshold() returns()
func (_CreditFilter *CreditFilterTransactorSession) UpdateUnderlyingTokenLiquidationThreshold() (*types.Transaction, error) {
	return _CreditFilter.Contract.UpdateUnderlyingTokenLiquidationThreshold(&_CreditFilter.TransactOpts)
}

// UpgradePriceOracle is a paid mutator transaction binding the contract method 0xf0527ac6.
//
// Solidity: function upgradePriceOracle() returns()
func (_CreditFilter *CreditFilterTransactor) UpgradePriceOracle(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditFilter.contract.Transact(opts, "upgradePriceOracle")
}

// UpgradePriceOracle is a paid mutator transaction binding the contract method 0xf0527ac6.
//
// Solidity: function upgradePriceOracle() returns()
func (_CreditFilter *CreditFilterSession) UpgradePriceOracle() (*types.Transaction, error) {
	return _CreditFilter.Contract.UpgradePriceOracle(&_CreditFilter.TransactOpts)
}

// UpgradePriceOracle is a paid mutator transaction binding the contract method 0xf0527ac6.
//
// Solidity: function upgradePriceOracle() returns()
func (_CreditFilter *CreditFilterTransactorSession) UpgradePriceOracle() (*types.Transaction, error) {
	return _CreditFilter.Contract.UpgradePriceOracle(&_CreditFilter.TransactOpts)
}

// CreditFilterContractAllowedIterator is returned from FilterContractAllowed and is used to iterate over the raw logs and unpacked data for ContractAllowed events raised by the CreditFilter contract.
type CreditFilterContractAllowedIterator struct {
	Event *CreditFilterContractAllowed // Event containing the contract specifics and raw log

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
func (it *CreditFilterContractAllowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditFilterContractAllowed)
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
		it.Event = new(CreditFilterContractAllowed)
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
func (it *CreditFilterContractAllowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditFilterContractAllowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditFilterContractAllowed represents a ContractAllowed event raised by the CreditFilter contract.
type CreditFilterContractAllowed struct {
	Protocol common.Address
	Adapter  common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterContractAllowed is a free log retrieval operation binding the contract event 0x4bcbefaef68b99503d502f5a6abe7bca2b183ab8ac55457013c77d084ebd1305.
//
// Solidity: event ContractAllowed(address indexed protocol, address indexed adapter)
func (_CreditFilter *CreditFilterFilterer) FilterContractAllowed(opts *bind.FilterOpts, protocol []common.Address, adapter []common.Address) (*CreditFilterContractAllowedIterator, error) {

	var protocolRule []interface{}
	for _, protocolItem := range protocol {
		protocolRule = append(protocolRule, protocolItem)
	}
	var adapterRule []interface{}
	for _, adapterItem := range adapter {
		adapterRule = append(adapterRule, adapterItem)
	}

	logs, sub, err := _CreditFilter.contract.FilterLogs(opts, "ContractAllowed", protocolRule, adapterRule)
	if err != nil {
		return nil, err
	}
	return &CreditFilterContractAllowedIterator{contract: _CreditFilter.contract, event: "ContractAllowed", logs: logs, sub: sub}, nil
}

// WatchContractAllowed is a free log subscription operation binding the contract event 0x4bcbefaef68b99503d502f5a6abe7bca2b183ab8ac55457013c77d084ebd1305.
//
// Solidity: event ContractAllowed(address indexed protocol, address indexed adapter)
func (_CreditFilter *CreditFilterFilterer) WatchContractAllowed(opts *bind.WatchOpts, sink chan<- *CreditFilterContractAllowed, protocol []common.Address, adapter []common.Address) (event.Subscription, error) {

	var protocolRule []interface{}
	for _, protocolItem := range protocol {
		protocolRule = append(protocolRule, protocolItem)
	}
	var adapterRule []interface{}
	for _, adapterItem := range adapter {
		adapterRule = append(adapterRule, adapterItem)
	}

	logs, sub, err := _CreditFilter.contract.WatchLogs(opts, "ContractAllowed", protocolRule, adapterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditFilterContractAllowed)
				if err := _CreditFilter.contract.UnpackLog(event, "ContractAllowed", log); err != nil {
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
func (_CreditFilter *CreditFilterFilterer) ParseContractAllowed(log types.Log) (*CreditFilterContractAllowed, error) {
	event := new(CreditFilterContractAllowed)
	if err := _CreditFilter.contract.UnpackLog(event, "ContractAllowed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CreditFilterContractForbiddenIterator is returned from FilterContractForbidden and is used to iterate over the raw logs and unpacked data for ContractForbidden events raised by the CreditFilter contract.
type CreditFilterContractForbiddenIterator struct {
	Event *CreditFilterContractForbidden // Event containing the contract specifics and raw log

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
func (it *CreditFilterContractForbiddenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditFilterContractForbidden)
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
		it.Event = new(CreditFilterContractForbidden)
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
func (it *CreditFilterContractForbiddenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditFilterContractForbiddenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditFilterContractForbidden represents a ContractForbidden event raised by the CreditFilter contract.
type CreditFilterContractForbidden struct {
	Protocol common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterContractForbidden is a free log retrieval operation binding the contract event 0xab9f405bf0c19b97f65a7031634db41569cd2f0e0376a610a1e977f9ab22b58f.
//
// Solidity: event ContractForbidden(address indexed protocol)
func (_CreditFilter *CreditFilterFilterer) FilterContractForbidden(opts *bind.FilterOpts, protocol []common.Address) (*CreditFilterContractForbiddenIterator, error) {

	var protocolRule []interface{}
	for _, protocolItem := range protocol {
		protocolRule = append(protocolRule, protocolItem)
	}

	logs, sub, err := _CreditFilter.contract.FilterLogs(opts, "ContractForbidden", protocolRule)
	if err != nil {
		return nil, err
	}
	return &CreditFilterContractForbiddenIterator{contract: _CreditFilter.contract, event: "ContractForbidden", logs: logs, sub: sub}, nil
}

// WatchContractForbidden is a free log subscription operation binding the contract event 0xab9f405bf0c19b97f65a7031634db41569cd2f0e0376a610a1e977f9ab22b58f.
//
// Solidity: event ContractForbidden(address indexed protocol)
func (_CreditFilter *CreditFilterFilterer) WatchContractForbidden(opts *bind.WatchOpts, sink chan<- *CreditFilterContractForbidden, protocol []common.Address) (event.Subscription, error) {

	var protocolRule []interface{}
	for _, protocolItem := range protocol {
		protocolRule = append(protocolRule, protocolItem)
	}

	logs, sub, err := _CreditFilter.contract.WatchLogs(opts, "ContractForbidden", protocolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditFilterContractForbidden)
				if err := _CreditFilter.contract.UnpackLog(event, "ContractForbidden", log); err != nil {
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
func (_CreditFilter *CreditFilterFilterer) ParseContractForbidden(log types.Log) (*CreditFilterContractForbidden, error) {
	event := new(CreditFilterContractForbidden)
	if err := _CreditFilter.contract.UnpackLog(event, "ContractForbidden", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CreditFilterNewFastCheckParametersIterator is returned from FilterNewFastCheckParameters and is used to iterate over the raw logs and unpacked data for NewFastCheckParameters events raised by the CreditFilter contract.
type CreditFilterNewFastCheckParametersIterator struct {
	Event *CreditFilterNewFastCheckParameters // Event containing the contract specifics and raw log

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
func (it *CreditFilterNewFastCheckParametersIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditFilterNewFastCheckParameters)
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
		it.Event = new(CreditFilterNewFastCheckParameters)
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
func (it *CreditFilterNewFastCheckParametersIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditFilterNewFastCheckParametersIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditFilterNewFastCheckParameters represents a NewFastCheckParameters event raised by the CreditFilter contract.
type CreditFilterNewFastCheckParameters struct {
	ChiThreshold   *big.Int
	FastCheckDelay *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewFastCheckParameters is a free log retrieval operation binding the contract event 0x727652fff0946c19c233fd3eab5fc03db9e9fdd907e902d9136c2a9cac47101c.
//
// Solidity: event NewFastCheckParameters(uint256 chiThreshold, uint256 fastCheckDelay)
func (_CreditFilter *CreditFilterFilterer) FilterNewFastCheckParameters(opts *bind.FilterOpts) (*CreditFilterNewFastCheckParametersIterator, error) {

	logs, sub, err := _CreditFilter.contract.FilterLogs(opts, "NewFastCheckParameters")
	if err != nil {
		return nil, err
	}
	return &CreditFilterNewFastCheckParametersIterator{contract: _CreditFilter.contract, event: "NewFastCheckParameters", logs: logs, sub: sub}, nil
}

// WatchNewFastCheckParameters is a free log subscription operation binding the contract event 0x727652fff0946c19c233fd3eab5fc03db9e9fdd907e902d9136c2a9cac47101c.
//
// Solidity: event NewFastCheckParameters(uint256 chiThreshold, uint256 fastCheckDelay)
func (_CreditFilter *CreditFilterFilterer) WatchNewFastCheckParameters(opts *bind.WatchOpts, sink chan<- *CreditFilterNewFastCheckParameters) (event.Subscription, error) {

	logs, sub, err := _CreditFilter.contract.WatchLogs(opts, "NewFastCheckParameters")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditFilterNewFastCheckParameters)
				if err := _CreditFilter.contract.UnpackLog(event, "NewFastCheckParameters", log); err != nil {
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
func (_CreditFilter *CreditFilterFilterer) ParseNewFastCheckParameters(log types.Log) (*CreditFilterNewFastCheckParameters, error) {
	event := new(CreditFilterNewFastCheckParameters)
	if err := _CreditFilter.contract.UnpackLog(event, "NewFastCheckParameters", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CreditFilterPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the CreditFilter contract.
type CreditFilterPausedIterator struct {
	Event *CreditFilterPaused // Event containing the contract specifics and raw log

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
func (it *CreditFilterPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditFilterPaused)
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
		it.Event = new(CreditFilterPaused)
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
func (it *CreditFilterPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditFilterPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditFilterPaused represents a Paused event raised by the CreditFilter contract.
type CreditFilterPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_CreditFilter *CreditFilterFilterer) FilterPaused(opts *bind.FilterOpts) (*CreditFilterPausedIterator, error) {

	logs, sub, err := _CreditFilter.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &CreditFilterPausedIterator{contract: _CreditFilter.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_CreditFilter *CreditFilterFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *CreditFilterPaused) (event.Subscription, error) {

	logs, sub, err := _CreditFilter.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditFilterPaused)
				if err := _CreditFilter.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_CreditFilter *CreditFilterFilterer) ParsePaused(log types.Log) (*CreditFilterPaused, error) {
	event := new(CreditFilterPaused)
	if err := _CreditFilter.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CreditFilterPriceOracleUpdatedIterator is returned from FilterPriceOracleUpdated and is used to iterate over the raw logs and unpacked data for PriceOracleUpdated events raised by the CreditFilter contract.
type CreditFilterPriceOracleUpdatedIterator struct {
	Event *CreditFilterPriceOracleUpdated // Event containing the contract specifics and raw log

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
func (it *CreditFilterPriceOracleUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditFilterPriceOracleUpdated)
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
		it.Event = new(CreditFilterPriceOracleUpdated)
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
func (it *CreditFilterPriceOracleUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditFilterPriceOracleUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditFilterPriceOracleUpdated represents a PriceOracleUpdated event raised by the CreditFilter contract.
type CreditFilterPriceOracleUpdated struct {
	NewPriceOracle common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPriceOracleUpdated is a free log retrieval operation binding the contract event 0xefe8ab924ca486283a79dc604baa67add51afb82af1db8ac386ebbba643cdffd.
//
// Solidity: event PriceOracleUpdated(address indexed newPriceOracle)
func (_CreditFilter *CreditFilterFilterer) FilterPriceOracleUpdated(opts *bind.FilterOpts, newPriceOracle []common.Address) (*CreditFilterPriceOracleUpdatedIterator, error) {

	var newPriceOracleRule []interface{}
	for _, newPriceOracleItem := range newPriceOracle {
		newPriceOracleRule = append(newPriceOracleRule, newPriceOracleItem)
	}

	logs, sub, err := _CreditFilter.contract.FilterLogs(opts, "PriceOracleUpdated", newPriceOracleRule)
	if err != nil {
		return nil, err
	}
	return &CreditFilterPriceOracleUpdatedIterator{contract: _CreditFilter.contract, event: "PriceOracleUpdated", logs: logs, sub: sub}, nil
}

// WatchPriceOracleUpdated is a free log subscription operation binding the contract event 0xefe8ab924ca486283a79dc604baa67add51afb82af1db8ac386ebbba643cdffd.
//
// Solidity: event PriceOracleUpdated(address indexed newPriceOracle)
func (_CreditFilter *CreditFilterFilterer) WatchPriceOracleUpdated(opts *bind.WatchOpts, sink chan<- *CreditFilterPriceOracleUpdated, newPriceOracle []common.Address) (event.Subscription, error) {

	var newPriceOracleRule []interface{}
	for _, newPriceOracleItem := range newPriceOracle {
		newPriceOracleRule = append(newPriceOracleRule, newPriceOracleItem)
	}

	logs, sub, err := _CreditFilter.contract.WatchLogs(opts, "PriceOracleUpdated", newPriceOracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditFilterPriceOracleUpdated)
				if err := _CreditFilter.contract.UnpackLog(event, "PriceOracleUpdated", log); err != nil {
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

// ParsePriceOracleUpdated is a log parse operation binding the contract event 0xefe8ab924ca486283a79dc604baa67add51afb82af1db8ac386ebbba643cdffd.
//
// Solidity: event PriceOracleUpdated(address indexed newPriceOracle)
func (_CreditFilter *CreditFilterFilterer) ParsePriceOracleUpdated(log types.Log) (*CreditFilterPriceOracleUpdated, error) {
	event := new(CreditFilterPriceOracleUpdated)
	if err := _CreditFilter.contract.UnpackLog(event, "PriceOracleUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CreditFilterTokenAllowedIterator is returned from FilterTokenAllowed and is used to iterate over the raw logs and unpacked data for TokenAllowed events raised by the CreditFilter contract.
type CreditFilterTokenAllowedIterator struct {
	Event *CreditFilterTokenAllowed // Event containing the contract specifics and raw log

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
func (it *CreditFilterTokenAllowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditFilterTokenAllowed)
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
		it.Event = new(CreditFilterTokenAllowed)
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
func (it *CreditFilterTokenAllowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditFilterTokenAllowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditFilterTokenAllowed represents a TokenAllowed event raised by the CreditFilter contract.
type CreditFilterTokenAllowed struct {
	Token              common.Address
	LiquidityThreshold *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterTokenAllowed is a free log retrieval operation binding the contract event 0xa52fb6bfa514a4ddcb31de40a5f6c20d767db1f921a8b7747973d93dc5da7a02.
//
// Solidity: event TokenAllowed(address indexed token, uint256 liquidityThreshold)
func (_CreditFilter *CreditFilterFilterer) FilterTokenAllowed(opts *bind.FilterOpts, token []common.Address) (*CreditFilterTokenAllowedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CreditFilter.contract.FilterLogs(opts, "TokenAllowed", tokenRule)
	if err != nil {
		return nil, err
	}
	return &CreditFilterTokenAllowedIterator{contract: _CreditFilter.contract, event: "TokenAllowed", logs: logs, sub: sub}, nil
}

// WatchTokenAllowed is a free log subscription operation binding the contract event 0xa52fb6bfa514a4ddcb31de40a5f6c20d767db1f921a8b7747973d93dc5da7a02.
//
// Solidity: event TokenAllowed(address indexed token, uint256 liquidityThreshold)
func (_CreditFilter *CreditFilterFilterer) WatchTokenAllowed(opts *bind.WatchOpts, sink chan<- *CreditFilterTokenAllowed, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CreditFilter.contract.WatchLogs(opts, "TokenAllowed", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditFilterTokenAllowed)
				if err := _CreditFilter.contract.UnpackLog(event, "TokenAllowed", log); err != nil {
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
func (_CreditFilter *CreditFilterFilterer) ParseTokenAllowed(log types.Log) (*CreditFilterTokenAllowed, error) {
	event := new(CreditFilterTokenAllowed)
	if err := _CreditFilter.contract.UnpackLog(event, "TokenAllowed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CreditFilterTokenForbiddenIterator is returned from FilterTokenForbidden and is used to iterate over the raw logs and unpacked data for TokenForbidden events raised by the CreditFilter contract.
type CreditFilterTokenForbiddenIterator struct {
	Event *CreditFilterTokenForbidden // Event containing the contract specifics and raw log

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
func (it *CreditFilterTokenForbiddenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditFilterTokenForbidden)
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
		it.Event = new(CreditFilterTokenForbidden)
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
func (it *CreditFilterTokenForbiddenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditFilterTokenForbiddenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditFilterTokenForbidden represents a TokenForbidden event raised by the CreditFilter contract.
type CreditFilterTokenForbidden struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTokenForbidden is a free log retrieval operation binding the contract event 0xf17b849746e74d7186170c9553d4bbf60b4f8bb1ed81fe50c099b934fb078f05.
//
// Solidity: event TokenForbidden(address indexed token)
func (_CreditFilter *CreditFilterFilterer) FilterTokenForbidden(opts *bind.FilterOpts, token []common.Address) (*CreditFilterTokenForbiddenIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CreditFilter.contract.FilterLogs(opts, "TokenForbidden", tokenRule)
	if err != nil {
		return nil, err
	}
	return &CreditFilterTokenForbiddenIterator{contract: _CreditFilter.contract, event: "TokenForbidden", logs: logs, sub: sub}, nil
}

// WatchTokenForbidden is a free log subscription operation binding the contract event 0xf17b849746e74d7186170c9553d4bbf60b4f8bb1ed81fe50c099b934fb078f05.
//
// Solidity: event TokenForbidden(address indexed token)
func (_CreditFilter *CreditFilterFilterer) WatchTokenForbidden(opts *bind.WatchOpts, sink chan<- *CreditFilterTokenForbidden, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CreditFilter.contract.WatchLogs(opts, "TokenForbidden", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditFilterTokenForbidden)
				if err := _CreditFilter.contract.UnpackLog(event, "TokenForbidden", log); err != nil {
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
func (_CreditFilter *CreditFilterFilterer) ParseTokenForbidden(log types.Log) (*CreditFilterTokenForbidden, error) {
	event := new(CreditFilterTokenForbidden)
	if err := _CreditFilter.contract.UnpackLog(event, "TokenForbidden", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CreditFilterTransferAccountAllowedIterator is returned from FilterTransferAccountAllowed and is used to iterate over the raw logs and unpacked data for TransferAccountAllowed events raised by the CreditFilter contract.
type CreditFilterTransferAccountAllowedIterator struct {
	Event *CreditFilterTransferAccountAllowed // Event containing the contract specifics and raw log

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
func (it *CreditFilterTransferAccountAllowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditFilterTransferAccountAllowed)
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
		it.Event = new(CreditFilterTransferAccountAllowed)
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
func (it *CreditFilterTransferAccountAllowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditFilterTransferAccountAllowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditFilterTransferAccountAllowed represents a TransferAccountAllowed event raised by the CreditFilter contract.
type CreditFilterTransferAccountAllowed struct {
	From  common.Address
	To    common.Address
	State bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransferAccountAllowed is a free log retrieval operation binding the contract event 0x9b3258bc4904fd6426b99843e206c6c7cdb1fd0f040121c25b71dafbb3851ee0.
//
// Solidity: event TransferAccountAllowed(address indexed from, address indexed to, bool state)
func (_CreditFilter *CreditFilterFilterer) FilterTransferAccountAllowed(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CreditFilterTransferAccountAllowedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CreditFilter.contract.FilterLogs(opts, "TransferAccountAllowed", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CreditFilterTransferAccountAllowedIterator{contract: _CreditFilter.contract, event: "TransferAccountAllowed", logs: logs, sub: sub}, nil
}

// WatchTransferAccountAllowed is a free log subscription operation binding the contract event 0x9b3258bc4904fd6426b99843e206c6c7cdb1fd0f040121c25b71dafbb3851ee0.
//
// Solidity: event TransferAccountAllowed(address indexed from, address indexed to, bool state)
func (_CreditFilter *CreditFilterFilterer) WatchTransferAccountAllowed(opts *bind.WatchOpts, sink chan<- *CreditFilterTransferAccountAllowed, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CreditFilter.contract.WatchLogs(opts, "TransferAccountAllowed", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditFilterTransferAccountAllowed)
				if err := _CreditFilter.contract.UnpackLog(event, "TransferAccountAllowed", log); err != nil {
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
func (_CreditFilter *CreditFilterFilterer) ParseTransferAccountAllowed(log types.Log) (*CreditFilterTransferAccountAllowed, error) {
	event := new(CreditFilterTransferAccountAllowed)
	if err := _CreditFilter.contract.UnpackLog(event, "TransferAccountAllowed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CreditFilterTransferPluginAllowedIterator is returned from FilterTransferPluginAllowed and is used to iterate over the raw logs and unpacked data for TransferPluginAllowed events raised by the CreditFilter contract.
type CreditFilterTransferPluginAllowedIterator struct {
	Event *CreditFilterTransferPluginAllowed // Event containing the contract specifics and raw log

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
func (it *CreditFilterTransferPluginAllowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditFilterTransferPluginAllowed)
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
		it.Event = new(CreditFilterTransferPluginAllowed)
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
func (it *CreditFilterTransferPluginAllowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditFilterTransferPluginAllowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditFilterTransferPluginAllowed represents a TransferPluginAllowed event raised by the CreditFilter contract.
type CreditFilterTransferPluginAllowed struct {
	Pugin common.Address
	State bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransferPluginAllowed is a free log retrieval operation binding the contract event 0xc7d2592986c53f858769b011e8ce6298936f8609789988e9f5ad4f0a20798897.
//
// Solidity: event TransferPluginAllowed(address indexed pugin, bool state)
func (_CreditFilter *CreditFilterFilterer) FilterTransferPluginAllowed(opts *bind.FilterOpts, pugin []common.Address) (*CreditFilterTransferPluginAllowedIterator, error) {

	var puginRule []interface{}
	for _, puginItem := range pugin {
		puginRule = append(puginRule, puginItem)
	}

	logs, sub, err := _CreditFilter.contract.FilterLogs(opts, "TransferPluginAllowed", puginRule)
	if err != nil {
		return nil, err
	}
	return &CreditFilterTransferPluginAllowedIterator{contract: _CreditFilter.contract, event: "TransferPluginAllowed", logs: logs, sub: sub}, nil
}

// WatchTransferPluginAllowed is a free log subscription operation binding the contract event 0xc7d2592986c53f858769b011e8ce6298936f8609789988e9f5ad4f0a20798897.
//
// Solidity: event TransferPluginAllowed(address indexed pugin, bool state)
func (_CreditFilter *CreditFilterFilterer) WatchTransferPluginAllowed(opts *bind.WatchOpts, sink chan<- *CreditFilterTransferPluginAllowed, pugin []common.Address) (event.Subscription, error) {

	var puginRule []interface{}
	for _, puginItem := range pugin {
		puginRule = append(puginRule, puginItem)
	}

	logs, sub, err := _CreditFilter.contract.WatchLogs(opts, "TransferPluginAllowed", puginRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditFilterTransferPluginAllowed)
				if err := _CreditFilter.contract.UnpackLog(event, "TransferPluginAllowed", log); err != nil {
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
func (_CreditFilter *CreditFilterFilterer) ParseTransferPluginAllowed(log types.Log) (*CreditFilterTransferPluginAllowed, error) {
	event := new(CreditFilterTransferPluginAllowed)
	if err := _CreditFilter.contract.UnpackLog(event, "TransferPluginAllowed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CreditFilterUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the CreditFilter contract.
type CreditFilterUnpausedIterator struct {
	Event *CreditFilterUnpaused // Event containing the contract specifics and raw log

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
func (it *CreditFilterUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditFilterUnpaused)
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
		it.Event = new(CreditFilterUnpaused)
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
func (it *CreditFilterUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditFilterUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditFilterUnpaused represents a Unpaused event raised by the CreditFilter contract.
type CreditFilterUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_CreditFilter *CreditFilterFilterer) FilterUnpaused(opts *bind.FilterOpts) (*CreditFilterUnpausedIterator, error) {

	logs, sub, err := _CreditFilter.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &CreditFilterUnpausedIterator{contract: _CreditFilter.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_CreditFilter *CreditFilterFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *CreditFilterUnpaused) (event.Subscription, error) {

	logs, sub, err := _CreditFilter.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditFilterUnpaused)
				if err := _CreditFilter.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_CreditFilter *CreditFilterFilterer) ParseUnpaused(log types.Log) (*CreditFilterUnpaused, error) {
	event := new(CreditFilterUnpaused)
	if err := _CreditFilter.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	return event, nil
}
