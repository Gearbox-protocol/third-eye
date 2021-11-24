// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iAppAddressProvider

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

// IAppAddressProviderMetaData contains all meta data concerning the IAppAddressProvider contract.
var IAppAddressProviderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"getDataCompressor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGearToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLeveragedActions\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPriceOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWETHGateway\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWethToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IAppAddressProviderABI is the input ABI used to generate the binding from.
// Deprecated: Use IAppAddressProviderMetaData.ABI instead.
var IAppAddressProviderABI = IAppAddressProviderMetaData.ABI

// IAppAddressProvider is an auto generated Go binding around an Ethereum contract.
type IAppAddressProvider struct {
	IAppAddressProviderCaller     // Read-only binding to the contract
	IAppAddressProviderTransactor // Write-only binding to the contract
	IAppAddressProviderFilterer   // Log filterer for contract events
}

// IAppAddressProviderCaller is an auto generated read-only Go binding around an Ethereum contract.
type IAppAddressProviderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAppAddressProviderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IAppAddressProviderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAppAddressProviderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAppAddressProviderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAppAddressProviderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IAppAddressProviderSession struct {
	Contract     *IAppAddressProvider // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IAppAddressProviderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IAppAddressProviderCallerSession struct {
	Contract *IAppAddressProviderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// IAppAddressProviderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IAppAddressProviderTransactorSession struct {
	Contract     *IAppAddressProviderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// IAppAddressProviderRaw is an auto generated low-level Go binding around an Ethereum contract.
type IAppAddressProviderRaw struct {
	Contract *IAppAddressProvider // Generic contract binding to access the raw methods on
}

// IAppAddressProviderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IAppAddressProviderCallerRaw struct {
	Contract *IAppAddressProviderCaller // Generic read-only contract binding to access the raw methods on
}

// IAppAddressProviderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IAppAddressProviderTransactorRaw struct {
	Contract *IAppAddressProviderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIAppAddressProvider creates a new instance of IAppAddressProvider, bound to a specific deployed contract.
func NewIAppAddressProvider(address common.Address, backend bind.ContractBackend) (*IAppAddressProvider, error) {
	contract, err := bindIAppAddressProvider(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAppAddressProvider{IAppAddressProviderCaller: IAppAddressProviderCaller{contract: contract}, IAppAddressProviderTransactor: IAppAddressProviderTransactor{contract: contract}, IAppAddressProviderFilterer: IAppAddressProviderFilterer{contract: contract}}, nil
}

// NewIAppAddressProviderCaller creates a new read-only instance of IAppAddressProvider, bound to a specific deployed contract.
func NewIAppAddressProviderCaller(address common.Address, caller bind.ContractCaller) (*IAppAddressProviderCaller, error) {
	contract, err := bindIAppAddressProvider(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAppAddressProviderCaller{contract: contract}, nil
}

// NewIAppAddressProviderTransactor creates a new write-only instance of IAppAddressProvider, bound to a specific deployed contract.
func NewIAppAddressProviderTransactor(address common.Address, transactor bind.ContractTransactor) (*IAppAddressProviderTransactor, error) {
	contract, err := bindIAppAddressProvider(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAppAddressProviderTransactor{contract: contract}, nil
}

// NewIAppAddressProviderFilterer creates a new log filterer instance of IAppAddressProvider, bound to a specific deployed contract.
func NewIAppAddressProviderFilterer(address common.Address, filterer bind.ContractFilterer) (*IAppAddressProviderFilterer, error) {
	contract, err := bindIAppAddressProvider(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAppAddressProviderFilterer{contract: contract}, nil
}

// bindIAppAddressProvider binds a generic wrapper to an already deployed contract.
func bindIAppAddressProvider(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IAppAddressProviderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAppAddressProvider *IAppAddressProviderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAppAddressProvider.Contract.IAppAddressProviderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAppAddressProvider *IAppAddressProviderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAppAddressProvider.Contract.IAppAddressProviderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAppAddressProvider *IAppAddressProviderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAppAddressProvider.Contract.IAppAddressProviderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAppAddressProvider *IAppAddressProviderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAppAddressProvider.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAppAddressProvider *IAppAddressProviderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAppAddressProvider.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAppAddressProvider *IAppAddressProviderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAppAddressProvider.Contract.contract.Transact(opts, method, params...)
}

// GetDataCompressor is a free data retrieval call binding the contract method 0x060678c2.
//
// Solidity: function getDataCompressor() view returns(address)
func (_IAppAddressProvider *IAppAddressProviderCaller) GetDataCompressor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IAppAddressProvider.contract.Call(opts, &out, "getDataCompressor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetDataCompressor is a free data retrieval call binding the contract method 0x060678c2.
//
// Solidity: function getDataCompressor() view returns(address)
func (_IAppAddressProvider *IAppAddressProviderSession) GetDataCompressor() (common.Address, error) {
	return _IAppAddressProvider.Contract.GetDataCompressor(&_IAppAddressProvider.CallOpts)
}

// GetDataCompressor is a free data retrieval call binding the contract method 0x060678c2.
//
// Solidity: function getDataCompressor() view returns(address)
func (_IAppAddressProvider *IAppAddressProviderCallerSession) GetDataCompressor() (common.Address, error) {
	return _IAppAddressProvider.Contract.GetDataCompressor(&_IAppAddressProvider.CallOpts)
}

// GetGearToken is a free data retrieval call binding the contract method 0xaffd9243.
//
// Solidity: function getGearToken() view returns(address)
func (_IAppAddressProvider *IAppAddressProviderCaller) GetGearToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IAppAddressProvider.contract.Call(opts, &out, "getGearToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetGearToken is a free data retrieval call binding the contract method 0xaffd9243.
//
// Solidity: function getGearToken() view returns(address)
func (_IAppAddressProvider *IAppAddressProviderSession) GetGearToken() (common.Address, error) {
	return _IAppAddressProvider.Contract.GetGearToken(&_IAppAddressProvider.CallOpts)
}

// GetGearToken is a free data retrieval call binding the contract method 0xaffd9243.
//
// Solidity: function getGearToken() view returns(address)
func (_IAppAddressProvider *IAppAddressProviderCallerSession) GetGearToken() (common.Address, error) {
	return _IAppAddressProvider.Contract.GetGearToken(&_IAppAddressProvider.CallOpts)
}

// GetLeveragedActions is a free data retrieval call binding the contract method 0x44b88563.
//
// Solidity: function getLeveragedActions() view returns(address)
func (_IAppAddressProvider *IAppAddressProviderCaller) GetLeveragedActions(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IAppAddressProvider.contract.Call(opts, &out, "getLeveragedActions")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLeveragedActions is a free data retrieval call binding the contract method 0x44b88563.
//
// Solidity: function getLeveragedActions() view returns(address)
func (_IAppAddressProvider *IAppAddressProviderSession) GetLeveragedActions() (common.Address, error) {
	return _IAppAddressProvider.Contract.GetLeveragedActions(&_IAppAddressProvider.CallOpts)
}

// GetLeveragedActions is a free data retrieval call binding the contract method 0x44b88563.
//
// Solidity: function getLeveragedActions() view returns(address)
func (_IAppAddressProvider *IAppAddressProviderCallerSession) GetLeveragedActions() (common.Address, error) {
	return _IAppAddressProvider.Contract.GetLeveragedActions(&_IAppAddressProvider.CallOpts)
}

// GetPriceOracle is a free data retrieval call binding the contract method 0xfca513a8.
//
// Solidity: function getPriceOracle() view returns(address)
func (_IAppAddressProvider *IAppAddressProviderCaller) GetPriceOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IAppAddressProvider.contract.Call(opts, &out, "getPriceOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPriceOracle is a free data retrieval call binding the contract method 0xfca513a8.
//
// Solidity: function getPriceOracle() view returns(address)
func (_IAppAddressProvider *IAppAddressProviderSession) GetPriceOracle() (common.Address, error) {
	return _IAppAddressProvider.Contract.GetPriceOracle(&_IAppAddressProvider.CallOpts)
}

// GetPriceOracle is a free data retrieval call binding the contract method 0xfca513a8.
//
// Solidity: function getPriceOracle() view returns(address)
func (_IAppAddressProvider *IAppAddressProviderCallerSession) GetPriceOracle() (common.Address, error) {
	return _IAppAddressProvider.Contract.GetPriceOracle(&_IAppAddressProvider.CallOpts)
}

// GetWETHGateway is a free data retrieval call binding the contract method 0x77532ed9.
//
// Solidity: function getWETHGateway() view returns(address)
func (_IAppAddressProvider *IAppAddressProviderCaller) GetWETHGateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IAppAddressProvider.contract.Call(opts, &out, "getWETHGateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetWETHGateway is a free data retrieval call binding the contract method 0x77532ed9.
//
// Solidity: function getWETHGateway() view returns(address)
func (_IAppAddressProvider *IAppAddressProviderSession) GetWETHGateway() (common.Address, error) {
	return _IAppAddressProvider.Contract.GetWETHGateway(&_IAppAddressProvider.CallOpts)
}

// GetWETHGateway is a free data retrieval call binding the contract method 0x77532ed9.
//
// Solidity: function getWETHGateway() view returns(address)
func (_IAppAddressProvider *IAppAddressProviderCallerSession) GetWETHGateway() (common.Address, error) {
	return _IAppAddressProvider.Contract.GetWETHGateway(&_IAppAddressProvider.CallOpts)
}

// GetWethToken is a free data retrieval call binding the contract method 0x4c252f91.
//
// Solidity: function getWethToken() view returns(address)
func (_IAppAddressProvider *IAppAddressProviderCaller) GetWethToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IAppAddressProvider.contract.Call(opts, &out, "getWethToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetWethToken is a free data retrieval call binding the contract method 0x4c252f91.
//
// Solidity: function getWethToken() view returns(address)
func (_IAppAddressProvider *IAppAddressProviderSession) GetWethToken() (common.Address, error) {
	return _IAppAddressProvider.Contract.GetWethToken(&_IAppAddressProvider.CallOpts)
}

// GetWethToken is a free data retrieval call binding the contract method 0x4c252f91.
//
// Solidity: function getWethToken() view returns(address)
func (_IAppAddressProvider *IAppAddressProviderCallerSession) GetWethToken() (common.Address, error) {
	return _IAppAddressProvider.Contract.GetWethToken(&_IAppAddressProvider.CallOpts)
}
