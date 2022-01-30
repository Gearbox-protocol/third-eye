// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package yearnPriceFeed

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

// YearnPriceFeedABI is the input ABI used to generate the binding from.
const YearnPriceFeedABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addressProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_yVault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_priceFeed\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_lowerBound\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_upperBound\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lowerBound\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"upperBound\",\"type\":\"uint256\"}],\"name\":\"NewLimiterParams\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimalsDivider\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"description\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint80\",\"name\":\"\",\"type\":\"uint80\"}],\"name\":\"getRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"\",\"type\":\"uint80\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lowerBound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceFeed\",\"outputs\":[{\"internalType\":\"contractAggregatorV3Interface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_lowerBound\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_upperBound\",\"type\":\"uint256\"}],\"name\":\"setLimiter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timestampLimiter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"upperBound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"yVault\",\"outputs\":[{\"internalType\":\"contractIYVault\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

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

// DecimalsDivider is a free data retrieval call binding the contract method 0xa834559e.
//
// Solidity: function decimalsDivider() view returns(uint256)
func (_YearnPriceFeed *YearnPriceFeedCaller) DecimalsDivider(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YearnPriceFeed.contract.Call(opts, &out, "decimalsDivider")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DecimalsDivider is a free data retrieval call binding the contract method 0xa834559e.
//
// Solidity: function decimalsDivider() view returns(uint256)
func (_YearnPriceFeed *YearnPriceFeedSession) DecimalsDivider() (*big.Int, error) {
	return _YearnPriceFeed.Contract.DecimalsDivider(&_YearnPriceFeed.CallOpts)
}

// DecimalsDivider is a free data retrieval call binding the contract method 0xa834559e.
//
// Solidity: function decimalsDivider() view returns(uint256)
func (_YearnPriceFeed *YearnPriceFeedCallerSession) DecimalsDivider() (*big.Int, error) {
	return _YearnPriceFeed.Contract.DecimalsDivider(&_YearnPriceFeed.CallOpts)
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
// Solidity: function getRoundData(uint80 ) pure returns(uint80, int256, uint256, uint256, uint80)
func (_YearnPriceFeed *YearnPriceFeedCaller) GetRoundData(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _YearnPriceFeed.contract.Call(opts, &out, "getRoundData", arg0)

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
// Solidity: function getRoundData(uint80 ) pure returns(uint80, int256, uint256, uint256, uint80)
func (_YearnPriceFeed *YearnPriceFeedSession) GetRoundData(arg0 *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _YearnPriceFeed.Contract.GetRoundData(&_YearnPriceFeed.CallOpts, arg0)
}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 ) pure returns(uint80, int256, uint256, uint256, uint80)
func (_YearnPriceFeed *YearnPriceFeedCallerSession) GetRoundData(arg0 *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _YearnPriceFeed.Contract.GetRoundData(&_YearnPriceFeed.CallOpts, arg0)
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

	outstruct.RoundId = out[0].(*big.Int)
	outstruct.Answer = out[1].(*big.Int)
	outstruct.StartedAt = out[2].(*big.Int)
	outstruct.UpdatedAt = out[3].(*big.Int)
	outstruct.AnsweredInRound = out[4].(*big.Int)

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

// LowerBound is a free data retrieval call binding the contract method 0xa384d6ff.
//
// Solidity: function lowerBound() view returns(uint256)
func (_YearnPriceFeed *YearnPriceFeedCaller) LowerBound(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YearnPriceFeed.contract.Call(opts, &out, "lowerBound")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LowerBound is a free data retrieval call binding the contract method 0xa384d6ff.
//
// Solidity: function lowerBound() view returns(uint256)
func (_YearnPriceFeed *YearnPriceFeedSession) LowerBound() (*big.Int, error) {
	return _YearnPriceFeed.Contract.LowerBound(&_YearnPriceFeed.CallOpts)
}

// LowerBound is a free data retrieval call binding the contract method 0xa384d6ff.
//
// Solidity: function lowerBound() view returns(uint256)
func (_YearnPriceFeed *YearnPriceFeedCallerSession) LowerBound() (*big.Int, error) {
	return _YearnPriceFeed.Contract.LowerBound(&_YearnPriceFeed.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_YearnPriceFeed *YearnPriceFeedCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _YearnPriceFeed.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_YearnPriceFeed *YearnPriceFeedSession) Paused() (bool, error) {
	return _YearnPriceFeed.Contract.Paused(&_YearnPriceFeed.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_YearnPriceFeed *YearnPriceFeedCallerSession) Paused() (bool, error) {
	return _YearnPriceFeed.Contract.Paused(&_YearnPriceFeed.CallOpts)
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

// TimestampLimiter is a free data retrieval call binding the contract method 0x2c51298c.
//
// Solidity: function timestampLimiter() view returns(uint256)
func (_YearnPriceFeed *YearnPriceFeedCaller) TimestampLimiter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YearnPriceFeed.contract.Call(opts, &out, "timestampLimiter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TimestampLimiter is a free data retrieval call binding the contract method 0x2c51298c.
//
// Solidity: function timestampLimiter() view returns(uint256)
func (_YearnPriceFeed *YearnPriceFeedSession) TimestampLimiter() (*big.Int, error) {
	return _YearnPriceFeed.Contract.TimestampLimiter(&_YearnPriceFeed.CallOpts)
}

// TimestampLimiter is a free data retrieval call binding the contract method 0x2c51298c.
//
// Solidity: function timestampLimiter() view returns(uint256)
func (_YearnPriceFeed *YearnPriceFeedCallerSession) TimestampLimiter() (*big.Int, error) {
	return _YearnPriceFeed.Contract.TimestampLimiter(&_YearnPriceFeed.CallOpts)
}

// UpperBound is a free data retrieval call binding the contract method 0xb09ad8a0.
//
// Solidity: function upperBound() view returns(uint256)
func (_YearnPriceFeed *YearnPriceFeedCaller) UpperBound(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YearnPriceFeed.contract.Call(opts, &out, "upperBound")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UpperBound is a free data retrieval call binding the contract method 0xb09ad8a0.
//
// Solidity: function upperBound() view returns(uint256)
func (_YearnPriceFeed *YearnPriceFeedSession) UpperBound() (*big.Int, error) {
	return _YearnPriceFeed.Contract.UpperBound(&_YearnPriceFeed.CallOpts)
}

// UpperBound is a free data retrieval call binding the contract method 0xb09ad8a0.
//
// Solidity: function upperBound() view returns(uint256)
func (_YearnPriceFeed *YearnPriceFeedCallerSession) UpperBound() (*big.Int, error) {
	return _YearnPriceFeed.Contract.UpperBound(&_YearnPriceFeed.CallOpts)
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

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_YearnPriceFeed *YearnPriceFeedTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YearnPriceFeed.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_YearnPriceFeed *YearnPriceFeedSession) Pause() (*types.Transaction, error) {
	return _YearnPriceFeed.Contract.Pause(&_YearnPriceFeed.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_YearnPriceFeed *YearnPriceFeedTransactorSession) Pause() (*types.Transaction, error) {
	return _YearnPriceFeed.Contract.Pause(&_YearnPriceFeed.TransactOpts)
}

// SetLimiter is a paid mutator transaction binding the contract method 0x0bdea781.
//
// Solidity: function setLimiter(uint256 _lowerBound, uint256 _upperBound) returns()
func (_YearnPriceFeed *YearnPriceFeedTransactor) SetLimiter(opts *bind.TransactOpts, _lowerBound *big.Int, _upperBound *big.Int) (*types.Transaction, error) {
	return _YearnPriceFeed.contract.Transact(opts, "setLimiter", _lowerBound, _upperBound)
}

// SetLimiter is a paid mutator transaction binding the contract method 0x0bdea781.
//
// Solidity: function setLimiter(uint256 _lowerBound, uint256 _upperBound) returns()
func (_YearnPriceFeed *YearnPriceFeedSession) SetLimiter(_lowerBound *big.Int, _upperBound *big.Int) (*types.Transaction, error) {
	return _YearnPriceFeed.Contract.SetLimiter(&_YearnPriceFeed.TransactOpts, _lowerBound, _upperBound)
}

// SetLimiter is a paid mutator transaction binding the contract method 0x0bdea781.
//
// Solidity: function setLimiter(uint256 _lowerBound, uint256 _upperBound) returns()
func (_YearnPriceFeed *YearnPriceFeedTransactorSession) SetLimiter(_lowerBound *big.Int, _upperBound *big.Int) (*types.Transaction, error) {
	return _YearnPriceFeed.Contract.SetLimiter(&_YearnPriceFeed.TransactOpts, _lowerBound, _upperBound)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_YearnPriceFeed *YearnPriceFeedTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YearnPriceFeed.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_YearnPriceFeed *YearnPriceFeedSession) Unpause() (*types.Transaction, error) {
	return _YearnPriceFeed.Contract.Unpause(&_YearnPriceFeed.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_YearnPriceFeed *YearnPriceFeedTransactorSession) Unpause() (*types.Transaction, error) {
	return _YearnPriceFeed.Contract.Unpause(&_YearnPriceFeed.TransactOpts)
}

// YearnPriceFeedNewLimiterParamsIterator is returned from FilterNewLimiterParams and is used to iterate over the raw logs and unpacked data for NewLimiterParams events raised by the YearnPriceFeed contract.
type YearnPriceFeedNewLimiterParamsIterator struct {
	Event *YearnPriceFeedNewLimiterParams // Event containing the contract specifics and raw log

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
func (it *YearnPriceFeedNewLimiterParamsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnPriceFeedNewLimiterParams)
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
		it.Event = new(YearnPriceFeedNewLimiterParams)
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
func (it *YearnPriceFeedNewLimiterParamsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnPriceFeedNewLimiterParamsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnPriceFeedNewLimiterParams represents a NewLimiterParams event raised by the YearnPriceFeed contract.
type YearnPriceFeedNewLimiterParams struct {
	LowerBound *big.Int
	UpperBound *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewLimiterParams is a free log retrieval operation binding the contract event 0x82e7ee47180a631312683eeb2a85ad264c9af490d54de5a75bbdb95b968c6de2.
//
// Solidity: event NewLimiterParams(uint256 lowerBound, uint256 upperBound)
func (_YearnPriceFeed *YearnPriceFeedFilterer) FilterNewLimiterParams(opts *bind.FilterOpts) (*YearnPriceFeedNewLimiterParamsIterator, error) {

	logs, sub, err := _YearnPriceFeed.contract.FilterLogs(opts, "NewLimiterParams")
	if err != nil {
		return nil, err
	}
	return &YearnPriceFeedNewLimiterParamsIterator{contract: _YearnPriceFeed.contract, event: "NewLimiterParams", logs: logs, sub: sub}, nil
}

// WatchNewLimiterParams is a free log subscription operation binding the contract event 0x82e7ee47180a631312683eeb2a85ad264c9af490d54de5a75bbdb95b968c6de2.
//
// Solidity: event NewLimiterParams(uint256 lowerBound, uint256 upperBound)
func (_YearnPriceFeed *YearnPriceFeedFilterer) WatchNewLimiterParams(opts *bind.WatchOpts, sink chan<- *YearnPriceFeedNewLimiterParams) (event.Subscription, error) {

	logs, sub, err := _YearnPriceFeed.contract.WatchLogs(opts, "NewLimiterParams")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnPriceFeedNewLimiterParams)
				if err := _YearnPriceFeed.contract.UnpackLog(event, "NewLimiterParams", log); err != nil {
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

// ParseNewLimiterParams is a log parse operation binding the contract event 0x82e7ee47180a631312683eeb2a85ad264c9af490d54de5a75bbdb95b968c6de2.
//
// Solidity: event NewLimiterParams(uint256 lowerBound, uint256 upperBound)
func (_YearnPriceFeed *YearnPriceFeedFilterer) ParseNewLimiterParams(log types.Log) (*YearnPriceFeedNewLimiterParams, error) {
	event := new(YearnPriceFeedNewLimiterParams)
	if err := _YearnPriceFeed.contract.UnpackLog(event, "NewLimiterParams", log); err != nil {
		return nil, err
	}
	return event, nil
}

// YearnPriceFeedPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the YearnPriceFeed contract.
type YearnPriceFeedPausedIterator struct {
	Event *YearnPriceFeedPaused // Event containing the contract specifics and raw log

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
func (it *YearnPriceFeedPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnPriceFeedPaused)
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
		it.Event = new(YearnPriceFeedPaused)
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
func (it *YearnPriceFeedPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnPriceFeedPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnPriceFeedPaused represents a Paused event raised by the YearnPriceFeed contract.
type YearnPriceFeedPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_YearnPriceFeed *YearnPriceFeedFilterer) FilterPaused(opts *bind.FilterOpts) (*YearnPriceFeedPausedIterator, error) {

	logs, sub, err := _YearnPriceFeed.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &YearnPriceFeedPausedIterator{contract: _YearnPriceFeed.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_YearnPriceFeed *YearnPriceFeedFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *YearnPriceFeedPaused) (event.Subscription, error) {

	logs, sub, err := _YearnPriceFeed.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnPriceFeedPaused)
				if err := _YearnPriceFeed.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_YearnPriceFeed *YearnPriceFeedFilterer) ParsePaused(log types.Log) (*YearnPriceFeedPaused, error) {
	event := new(YearnPriceFeedPaused)
	if err := _YearnPriceFeed.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	return event, nil
}

// YearnPriceFeedUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the YearnPriceFeed contract.
type YearnPriceFeedUnpausedIterator struct {
	Event *YearnPriceFeedUnpaused // Event containing the contract specifics and raw log

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
func (it *YearnPriceFeedUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnPriceFeedUnpaused)
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
		it.Event = new(YearnPriceFeedUnpaused)
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
func (it *YearnPriceFeedUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnPriceFeedUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnPriceFeedUnpaused represents a Unpaused event raised by the YearnPriceFeed contract.
type YearnPriceFeedUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_YearnPriceFeed *YearnPriceFeedFilterer) FilterUnpaused(opts *bind.FilterOpts) (*YearnPriceFeedUnpausedIterator, error) {

	logs, sub, err := _YearnPriceFeed.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &YearnPriceFeedUnpausedIterator{contract: _YearnPriceFeed.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_YearnPriceFeed *YearnPriceFeedFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *YearnPriceFeedUnpaused) (event.Subscription, error) {

	logs, sub, err := _YearnPriceFeed.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnPriceFeedUnpaused)
				if err := _YearnPriceFeed.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_YearnPriceFeed *YearnPriceFeedFilterer) ParseUnpaused(log types.Log) (*YearnPriceFeedUnpaused, error) {
	event := new(YearnPriceFeedUnpaused)
	if err := _YearnPriceFeed.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	return event, nil
}
