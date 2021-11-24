// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package chainlinkPriceFeedMock

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

// ChainlinkPriceFeedMockMetaData contains all meta data concerning the ChainlinkPriceFeedMock contract.
var ChainlinkPriceFeedMockMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"_price\",\"type\":\"int256\"},{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"description\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint80\",\"name\":\"\",\"type\":\"uint80\"}],\"name\":\"getRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"newPrice\",\"type\":\"int256\"}],\"name\":\"setPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// ChainlinkPriceFeedMockABI is the input ABI used to generate the binding from.
// Deprecated: Use ChainlinkPriceFeedMockMetaData.ABI instead.
var ChainlinkPriceFeedMockABI = ChainlinkPriceFeedMockMetaData.ABI

// ChainlinkPriceFeedMock is an auto generated Go binding around an Ethereum contract.
type ChainlinkPriceFeedMock struct {
	ChainlinkPriceFeedMockCaller     // Read-only binding to the contract
	ChainlinkPriceFeedMockTransactor // Write-only binding to the contract
	ChainlinkPriceFeedMockFilterer   // Log filterer for contract events
}

// ChainlinkPriceFeedMockCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChainlinkPriceFeedMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChainlinkPriceFeedMockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChainlinkPriceFeedMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChainlinkPriceFeedMockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChainlinkPriceFeedMockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChainlinkPriceFeedMockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChainlinkPriceFeedMockSession struct {
	Contract     *ChainlinkPriceFeedMock // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ChainlinkPriceFeedMockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChainlinkPriceFeedMockCallerSession struct {
	Contract *ChainlinkPriceFeedMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// ChainlinkPriceFeedMockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChainlinkPriceFeedMockTransactorSession struct {
	Contract     *ChainlinkPriceFeedMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// ChainlinkPriceFeedMockRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChainlinkPriceFeedMockRaw struct {
	Contract *ChainlinkPriceFeedMock // Generic contract binding to access the raw methods on
}

// ChainlinkPriceFeedMockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChainlinkPriceFeedMockCallerRaw struct {
	Contract *ChainlinkPriceFeedMockCaller // Generic read-only contract binding to access the raw methods on
}

// ChainlinkPriceFeedMockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChainlinkPriceFeedMockTransactorRaw struct {
	Contract *ChainlinkPriceFeedMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChainlinkPriceFeedMock creates a new instance of ChainlinkPriceFeedMock, bound to a specific deployed contract.
func NewChainlinkPriceFeedMock(address common.Address, backend bind.ContractBackend) (*ChainlinkPriceFeedMock, error) {
	contract, err := bindChainlinkPriceFeedMock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChainlinkPriceFeedMock{ChainlinkPriceFeedMockCaller: ChainlinkPriceFeedMockCaller{contract: contract}, ChainlinkPriceFeedMockTransactor: ChainlinkPriceFeedMockTransactor{contract: contract}, ChainlinkPriceFeedMockFilterer: ChainlinkPriceFeedMockFilterer{contract: contract}}, nil
}

// NewChainlinkPriceFeedMockCaller creates a new read-only instance of ChainlinkPriceFeedMock, bound to a specific deployed contract.
func NewChainlinkPriceFeedMockCaller(address common.Address, caller bind.ContractCaller) (*ChainlinkPriceFeedMockCaller, error) {
	contract, err := bindChainlinkPriceFeedMock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChainlinkPriceFeedMockCaller{contract: contract}, nil
}

// NewChainlinkPriceFeedMockTransactor creates a new write-only instance of ChainlinkPriceFeedMock, bound to a specific deployed contract.
func NewChainlinkPriceFeedMockTransactor(address common.Address, transactor bind.ContractTransactor) (*ChainlinkPriceFeedMockTransactor, error) {
	contract, err := bindChainlinkPriceFeedMock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChainlinkPriceFeedMockTransactor{contract: contract}, nil
}

// NewChainlinkPriceFeedMockFilterer creates a new log filterer instance of ChainlinkPriceFeedMock, bound to a specific deployed contract.
func NewChainlinkPriceFeedMockFilterer(address common.Address, filterer bind.ContractFilterer) (*ChainlinkPriceFeedMockFilterer, error) {
	contract, err := bindChainlinkPriceFeedMock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChainlinkPriceFeedMockFilterer{contract: contract}, nil
}

// bindChainlinkPriceFeedMock binds a generic wrapper to an already deployed contract.
func bindChainlinkPriceFeedMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChainlinkPriceFeedMockABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChainlinkPriceFeedMock *ChainlinkPriceFeedMockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChainlinkPriceFeedMock.Contract.ChainlinkPriceFeedMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChainlinkPriceFeedMock *ChainlinkPriceFeedMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChainlinkPriceFeedMock.Contract.ChainlinkPriceFeedMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChainlinkPriceFeedMock *ChainlinkPriceFeedMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChainlinkPriceFeedMock.Contract.ChainlinkPriceFeedMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChainlinkPriceFeedMock *ChainlinkPriceFeedMockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChainlinkPriceFeedMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChainlinkPriceFeedMock *ChainlinkPriceFeedMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChainlinkPriceFeedMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChainlinkPriceFeedMock *ChainlinkPriceFeedMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChainlinkPriceFeedMock.Contract.contract.Transact(opts, method, params...)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ChainlinkPriceFeedMock *ChainlinkPriceFeedMockCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ChainlinkPriceFeedMock.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ChainlinkPriceFeedMock *ChainlinkPriceFeedMockSession) Decimals() (uint8, error) {
	return _ChainlinkPriceFeedMock.Contract.Decimals(&_ChainlinkPriceFeedMock.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ChainlinkPriceFeedMock *ChainlinkPriceFeedMockCallerSession) Decimals() (uint8, error) {
	return _ChainlinkPriceFeedMock.Contract.Decimals(&_ChainlinkPriceFeedMock.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() pure returns(string)
func (_ChainlinkPriceFeedMock *ChainlinkPriceFeedMockCaller) Description(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ChainlinkPriceFeedMock.contract.Call(opts, &out, "description")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() pure returns(string)
func (_ChainlinkPriceFeedMock *ChainlinkPriceFeedMockSession) Description() (string, error) {
	return _ChainlinkPriceFeedMock.Contract.Description(&_ChainlinkPriceFeedMock.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() pure returns(string)
func (_ChainlinkPriceFeedMock *ChainlinkPriceFeedMockCallerSession) Description() (string, error) {
	return _ChainlinkPriceFeedMock.Contract.Description(&_ChainlinkPriceFeedMock.CallOpts)
}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 ) view returns(uint80, int256, uint256, uint256, uint80)
func (_ChainlinkPriceFeedMock *ChainlinkPriceFeedMockCaller) GetRoundData(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _ChainlinkPriceFeedMock.contract.Call(opts, &out, "getRoundData", arg0)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, err

}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 ) view returns(uint80, int256, uint256, uint256, uint80)
func (_ChainlinkPriceFeedMock *ChainlinkPriceFeedMockSession) GetRoundData(arg0 *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _ChainlinkPriceFeedMock.Contract.GetRoundData(&_ChainlinkPriceFeedMock.CallOpts, arg0)
}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 ) view returns(uint80, int256, uint256, uint256, uint80)
func (_ChainlinkPriceFeedMock *ChainlinkPriceFeedMockCallerSession) GetRoundData(arg0 *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _ChainlinkPriceFeedMock.Contract.GetRoundData(&_ChainlinkPriceFeedMock.CallOpts, arg0)
}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80, int256, uint256, uint256, uint80)
func (_ChainlinkPriceFeedMock *ChainlinkPriceFeedMockCaller) LatestRoundData(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _ChainlinkPriceFeedMock.contract.Call(opts, &out, "latestRoundData")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, err

}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80, int256, uint256, uint256, uint80)
func (_ChainlinkPriceFeedMock *ChainlinkPriceFeedMockSession) LatestRoundData() (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _ChainlinkPriceFeedMock.Contract.LatestRoundData(&_ChainlinkPriceFeedMock.CallOpts)
}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80, int256, uint256, uint256, uint80)
func (_ChainlinkPriceFeedMock *ChainlinkPriceFeedMockCallerSession) LatestRoundData() (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _ChainlinkPriceFeedMock.Contract.LatestRoundData(&_ChainlinkPriceFeedMock.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint256)
func (_ChainlinkPriceFeedMock *ChainlinkPriceFeedMockCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ChainlinkPriceFeedMock.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint256)
func (_ChainlinkPriceFeedMock *ChainlinkPriceFeedMockSession) Version() (*big.Int, error) {
	return _ChainlinkPriceFeedMock.Contract.Version(&_ChainlinkPriceFeedMock.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint256)
func (_ChainlinkPriceFeedMock *ChainlinkPriceFeedMockCallerSession) Version() (*big.Int, error) {
	return _ChainlinkPriceFeedMock.Contract.Version(&_ChainlinkPriceFeedMock.CallOpts)
}

// SetPrice is a paid mutator transaction binding the contract method 0xf7a30806.
//
// Solidity: function setPrice(int256 newPrice) returns()
func (_ChainlinkPriceFeedMock *ChainlinkPriceFeedMockTransactor) SetPrice(opts *bind.TransactOpts, newPrice *big.Int) (*types.Transaction, error) {
	return _ChainlinkPriceFeedMock.contract.Transact(opts, "setPrice", newPrice)
}

// SetPrice is a paid mutator transaction binding the contract method 0xf7a30806.
//
// Solidity: function setPrice(int256 newPrice) returns()
func (_ChainlinkPriceFeedMock *ChainlinkPriceFeedMockSession) SetPrice(newPrice *big.Int) (*types.Transaction, error) {
	return _ChainlinkPriceFeedMock.Contract.SetPrice(&_ChainlinkPriceFeedMock.TransactOpts, newPrice)
}

// SetPrice is a paid mutator transaction binding the contract method 0xf7a30806.
//
// Solidity: function setPrice(int256 newPrice) returns()
func (_ChainlinkPriceFeedMock *ChainlinkPriceFeedMockTransactorSession) SetPrice(newPrice *big.Int) (*types.Transaction, error) {
	return _ChainlinkPriceFeedMock.Contract.SetPrice(&_ChainlinkPriceFeedMock.TransactOpts, newPrice)
}
