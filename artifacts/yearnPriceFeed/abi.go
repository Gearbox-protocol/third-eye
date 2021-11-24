// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package yearnPriceFeed

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

// YearnPriceFeedMetaData contains all meta data concerning the YearnPriceFeed contract.
var YearnPriceFeedMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_yVault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_priceFeed\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"description\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint80\",\"name\":\"_roundId\",\"type\":\"uint80\"}],\"name\":\"getRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceFeed\",\"outputs\":[{\"internalType\":\"contractAggregatorV3Interface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"yVault\",\"outputs\":[{\"internalType\":\"contractIYVault\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// YearnPriceFeedABI is the input ABI used to generate the binding from.
// Deprecated: Use YearnPriceFeedMetaData.ABI instead.
var YearnPriceFeedABI = YearnPriceFeedMetaData.ABI

// YearnPriceFeed is an auto generated Go binding around an Ethereum contract.
type YearnPriceFeed struct {
	YearnPriceFeedCaller     // Read-only binding to the contract
	YearnPriceFeedTransactor // Write-only binding to the contract
	YearnPriceFeedFilterer   // Log filterer for contract events
}

// YearnPriceFeedCaller is an auto generated read-only Go binding around an Ethereum contract.
type YearnPriceFeedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YearnPriceFeedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type YearnPriceFeedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YearnPriceFeedFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type YearnPriceFeedFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YearnPriceFeedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type YearnPriceFeedSession struct {
	Contract     *YearnPriceFeed   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// YearnPriceFeedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type YearnPriceFeedCallerSession struct {
	Contract *YearnPriceFeedCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// YearnPriceFeedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type YearnPriceFeedTransactorSession struct {
	Contract     *YearnPriceFeedTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// YearnPriceFeedRaw is an auto generated low-level Go binding around an Ethereum contract.
type YearnPriceFeedRaw struct {
	Contract *YearnPriceFeed // Generic contract binding to access the raw methods on
}

// YearnPriceFeedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type YearnPriceFeedCallerRaw struct {
	Contract *YearnPriceFeedCaller // Generic read-only contract binding to access the raw methods on
}

// YearnPriceFeedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type YearnPriceFeedTransactorRaw struct {
	Contract *YearnPriceFeedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewYearnPriceFeed creates a new instance of YearnPriceFeed, bound to a specific deployed contract.
func NewYearnPriceFeed(address common.Address, backend bind.ContractBackend) (*YearnPriceFeed, error) {
	contract, err := bindYearnPriceFeed(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &YearnPriceFeed{YearnPriceFeedCaller: YearnPriceFeedCaller{contract: contract}, YearnPriceFeedTransactor: YearnPriceFeedTransactor{contract: contract}, YearnPriceFeedFilterer: YearnPriceFeedFilterer{contract: contract}}, nil
}

// NewYearnPriceFeedCaller creates a new read-only instance of YearnPriceFeed, bound to a specific deployed contract.
func NewYearnPriceFeedCaller(address common.Address, caller bind.ContractCaller) (*YearnPriceFeedCaller, error) {
	contract, err := bindYearnPriceFeed(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &YearnPriceFeedCaller{contract: contract}, nil
}

// NewYearnPriceFeedTransactor creates a new write-only instance of YearnPriceFeed, bound to a specific deployed contract.
func NewYearnPriceFeedTransactor(address common.Address, transactor bind.ContractTransactor) (*YearnPriceFeedTransactor, error) {
	contract, err := bindYearnPriceFeed(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &YearnPriceFeedTransactor{contract: contract}, nil
}

// NewYearnPriceFeedFilterer creates a new log filterer instance of YearnPriceFeed, bound to a specific deployed contract.
func NewYearnPriceFeedFilterer(address common.Address, filterer bind.ContractFilterer) (*YearnPriceFeedFilterer, error) {
	contract, err := bindYearnPriceFeed(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &YearnPriceFeedFilterer{contract: contract}, nil
}

// bindYearnPriceFeed binds a generic wrapper to an already deployed contract.
func bindYearnPriceFeed(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(YearnPriceFeedABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YearnPriceFeed *YearnPriceFeedRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YearnPriceFeed.Contract.YearnPriceFeedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YearnPriceFeed *YearnPriceFeedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YearnPriceFeed.Contract.YearnPriceFeedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YearnPriceFeed *YearnPriceFeedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YearnPriceFeed.Contract.YearnPriceFeedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YearnPriceFeed *YearnPriceFeedCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YearnPriceFeed.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YearnPriceFeed *YearnPriceFeedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YearnPriceFeed.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YearnPriceFeed *YearnPriceFeedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YearnPriceFeed.Contract.contract.Transact(opts, method, params...)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_YearnPriceFeed *YearnPriceFeedCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _YearnPriceFeed.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_YearnPriceFeed *YearnPriceFeedSession) Decimals() (uint8, error) {
	return _YearnPriceFeed.Contract.Decimals(&_YearnPriceFeed.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_YearnPriceFeed *YearnPriceFeedCallerSession) Decimals() (uint8, error) {
	return _YearnPriceFeed.Contract.Decimals(&_YearnPriceFeed.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_YearnPriceFeed *YearnPriceFeedCaller) Description(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _YearnPriceFeed.contract.Call(opts, &out, "description")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_YearnPriceFeed *YearnPriceFeedSession) Description() (string, error) {
	return _YearnPriceFeed.Contract.Description(&_YearnPriceFeed.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_YearnPriceFeed *YearnPriceFeedCallerSession) Description() (string, error) {
	return _YearnPriceFeed.Contract.Description(&_YearnPriceFeed.CallOpts)
}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 _roundId) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_YearnPriceFeed *YearnPriceFeedCaller) GetRoundData(opts *bind.CallOpts, _roundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	var out []interface{}
	err := _YearnPriceFeed.contract.Call(opts, &out, "getRoundData", _roundId)

	outstruct := new(struct {
		RoundId         *big.Int
		Answer          *big.Int
		StartedAt       *big.Int
		UpdatedAt       *big.Int
		AnsweredInRound *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RoundId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Answer = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StartedAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UpdatedAt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.AnsweredInRound = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 _roundId) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_YearnPriceFeed *YearnPriceFeedSession) GetRoundData(_roundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _YearnPriceFeed.Contract.GetRoundData(&_YearnPriceFeed.CallOpts, _roundId)
}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 _roundId) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_YearnPriceFeed *YearnPriceFeedCallerSession) GetRoundData(_roundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _YearnPriceFeed.Contract.GetRoundData(&_YearnPriceFeed.CallOpts, _roundId)
}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_YearnPriceFeed *YearnPriceFeedCaller) LatestRoundData(opts *bind.CallOpts) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	var out []interface{}
	err := _YearnPriceFeed.contract.Call(opts, &out, "latestRoundData")

	outstruct := new(struct {
		RoundId         *big.Int
		Answer          *big.Int
		StartedAt       *big.Int
		UpdatedAt       *big.Int
		AnsweredInRound *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RoundId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Answer = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StartedAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UpdatedAt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.AnsweredInRound = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_YearnPriceFeed *YearnPriceFeedSession) LatestRoundData() (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _YearnPriceFeed.Contract.LatestRoundData(&_YearnPriceFeed.CallOpts)
}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_YearnPriceFeed *YearnPriceFeedCallerSession) LatestRoundData() (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _YearnPriceFeed.Contract.LatestRoundData(&_YearnPriceFeed.CallOpts)
}

// PriceFeed is a free data retrieval call binding the contract method 0x741bef1a.
//
// Solidity: function priceFeed() view returns(address)
func (_YearnPriceFeed *YearnPriceFeedCaller) PriceFeed(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YearnPriceFeed.contract.Call(opts, &out, "priceFeed")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PriceFeed is a free data retrieval call binding the contract method 0x741bef1a.
//
// Solidity: function priceFeed() view returns(address)
func (_YearnPriceFeed *YearnPriceFeedSession) PriceFeed() (common.Address, error) {
	return _YearnPriceFeed.Contract.PriceFeed(&_YearnPriceFeed.CallOpts)
}

// PriceFeed is a free data retrieval call binding the contract method 0x741bef1a.
//
// Solidity: function priceFeed() view returns(address)
func (_YearnPriceFeed *YearnPriceFeedCallerSession) PriceFeed() (common.Address, error) {
	return _YearnPriceFeed.Contract.PriceFeed(&_YearnPriceFeed.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_YearnPriceFeed *YearnPriceFeedCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YearnPriceFeed.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_YearnPriceFeed *YearnPriceFeedSession) Version() (*big.Int, error) {
	return _YearnPriceFeed.Contract.Version(&_YearnPriceFeed.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_YearnPriceFeed *YearnPriceFeedCallerSession) Version() (*big.Int, error) {
	return _YearnPriceFeed.Contract.Version(&_YearnPriceFeed.CallOpts)
}

// YVault is a free data retrieval call binding the contract method 0x33303f8e.
//
// Solidity: function yVault() view returns(address)
func (_YearnPriceFeed *YearnPriceFeedCaller) YVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YearnPriceFeed.contract.Call(opts, &out, "yVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// YVault is a free data retrieval call binding the contract method 0x33303f8e.
//
// Solidity: function yVault() view returns(address)
func (_YearnPriceFeed *YearnPriceFeedSession) YVault() (common.Address, error) {
	return _YearnPriceFeed.Contract.YVault(&_YearnPriceFeed.CallOpts)
}

// YVault is a free data retrieval call binding the contract method 0x33303f8e.
//
// Solidity: function yVault() view returns(address)
func (_YearnPriceFeed *YearnPriceFeedCallerSession) YVault() (common.Address, error) {
	return _YearnPriceFeed.Contract.YVault(&_YearnPriceFeed.CallOpts)
}
