// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package priceOracle

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

// PriceOracleABI is the input ABI used to generate the binding from.
const PriceOracleABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addressProvider\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"priceFeed\",\"type\":\"address\"}],\"name\":\"NewPriceFeed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceFeed\",\"type\":\"address\"}],\"name\":\"addPriceFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenFrom\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenTo\",\"type\":\"address\"}],\"name\":\"convert\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"decimalsDividers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"decimalsMultipliers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenFrom\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenTo\",\"type\":\"address\"}],\"name\":\"getLastPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"priceFeeds\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wethAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// PriceOracle is an auto generated Go binding around an Ethereum contract.
type PriceOracle struct {
	PriceOracleCaller     // Read-only binding to the contract
	PriceOracleTransactor // Write-only binding to the contract
	PriceOracleFilterer   // Log filterer for contract events
}

// PriceOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type PriceOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PriceOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PriceOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PriceOracleSession struct {
	Contract     *PriceOracle      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PriceOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PriceOracleCallerSession struct {
	Contract *PriceOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// PriceOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PriceOracleTransactorSession struct {
	Contract     *PriceOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// PriceOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type PriceOracleRaw struct {
	Contract *PriceOracle // Generic contract binding to access the raw methods on
}

// PriceOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PriceOracleCallerRaw struct {
	Contract *PriceOracleCaller // Generic read-only contract binding to access the raw methods on
}

// PriceOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PriceOracleTransactorRaw struct {
	Contract *PriceOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPriceOracle creates a new instance of PriceOracle, bound to a specific deployed contract.
func NewPriceOracle(address common.Address, backend bind.ContractBackend) (*PriceOracle, error) {
	contract, err := bindPriceOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PriceOracle{PriceOracleCaller: PriceOracleCaller{contract: contract}, PriceOracleTransactor: PriceOracleTransactor{contract: contract}, PriceOracleFilterer: PriceOracleFilterer{contract: contract}}, nil
}

// NewPriceOracleCaller creates a new read-only instance of PriceOracle, bound to a specific deployed contract.
func NewPriceOracleCaller(address common.Address, caller bind.ContractCaller) (*PriceOracleCaller, error) {
	contract, err := bindPriceOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PriceOracleCaller{contract: contract}, nil
}

// NewPriceOracleTransactor creates a new write-only instance of PriceOracle, bound to a specific deployed contract.
func NewPriceOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*PriceOracleTransactor, error) {
	contract, err := bindPriceOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PriceOracleTransactor{contract: contract}, nil
}

// NewPriceOracleFilterer creates a new log filterer instance of PriceOracle, bound to a specific deployed contract.
func NewPriceOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*PriceOracleFilterer, error) {
	contract, err := bindPriceOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PriceOracleFilterer{contract: contract}, nil
}

// bindPriceOracle binds a generic wrapper to an already deployed contract.
func bindPriceOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PriceOracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriceOracle *PriceOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PriceOracle.Contract.PriceOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriceOracle *PriceOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceOracle.Contract.PriceOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriceOracle *PriceOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriceOracle.Contract.PriceOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriceOracle *PriceOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PriceOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriceOracle *PriceOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriceOracle *PriceOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriceOracle.Contract.contract.Transact(opts, method, params...)
}

// Convert is a free data retrieval call binding the contract method 0xb66102df.
//
// Solidity: function convert(uint256 amount, address tokenFrom, address tokenTo) view returns(uint256)
func (_PriceOracle *PriceOracleCaller) Convert(opts *bind.CallOpts, amount *big.Int, tokenFrom common.Address, tokenTo common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PriceOracle.contract.Call(opts, &out, "convert", amount, tokenFrom, tokenTo)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Convert is a free data retrieval call binding the contract method 0xb66102df.
//
// Solidity: function convert(uint256 amount, address tokenFrom, address tokenTo) view returns(uint256)
func (_PriceOracle *PriceOracleSession) Convert(amount *big.Int, tokenFrom common.Address, tokenTo common.Address) (*big.Int, error) {
	return _PriceOracle.Contract.Convert(&_PriceOracle.CallOpts, amount, tokenFrom, tokenTo)
}

// Convert is a free data retrieval call binding the contract method 0xb66102df.
//
// Solidity: function convert(uint256 amount, address tokenFrom, address tokenTo) view returns(uint256)
func (_PriceOracle *PriceOracleCallerSession) Convert(amount *big.Int, tokenFrom common.Address, tokenTo common.Address) (*big.Int, error) {
	return _PriceOracle.Contract.Convert(&_PriceOracle.CallOpts, amount, tokenFrom, tokenTo)
}

// DecimalsDividers is a free data retrieval call binding the contract method 0x21c9afa5.
//
// Solidity: function decimalsDividers(address ) view returns(uint256)
func (_PriceOracle *PriceOracleCaller) DecimalsDividers(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PriceOracle.contract.Call(opts, &out, "decimalsDividers", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DecimalsDividers is a free data retrieval call binding the contract method 0x21c9afa5.
//
// Solidity: function decimalsDividers(address ) view returns(uint256)
func (_PriceOracle *PriceOracleSession) DecimalsDividers(arg0 common.Address) (*big.Int, error) {
	return _PriceOracle.Contract.DecimalsDividers(&_PriceOracle.CallOpts, arg0)
}

// DecimalsDividers is a free data retrieval call binding the contract method 0x21c9afa5.
//
// Solidity: function decimalsDividers(address ) view returns(uint256)
func (_PriceOracle *PriceOracleCallerSession) DecimalsDividers(arg0 common.Address) (*big.Int, error) {
	return _PriceOracle.Contract.DecimalsDividers(&_PriceOracle.CallOpts, arg0)
}

// DecimalsMultipliers is a free data retrieval call binding the contract method 0xd4561137.
//
// Solidity: function decimalsMultipliers(address ) view returns(uint256)
func (_PriceOracle *PriceOracleCaller) DecimalsMultipliers(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PriceOracle.contract.Call(opts, &out, "decimalsMultipliers", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DecimalsMultipliers is a free data retrieval call binding the contract method 0xd4561137.
//
// Solidity: function decimalsMultipliers(address ) view returns(uint256)
func (_PriceOracle *PriceOracleSession) DecimalsMultipliers(arg0 common.Address) (*big.Int, error) {
	return _PriceOracle.Contract.DecimalsMultipliers(&_PriceOracle.CallOpts, arg0)
}

// DecimalsMultipliers is a free data retrieval call binding the contract method 0xd4561137.
//
// Solidity: function decimalsMultipliers(address ) view returns(uint256)
func (_PriceOracle *PriceOracleCallerSession) DecimalsMultipliers(arg0 common.Address) (*big.Int, error) {
	return _PriceOracle.Contract.DecimalsMultipliers(&_PriceOracle.CallOpts, arg0)
}

// GetLastPrice is a free data retrieval call binding the contract method 0x743b9086.
//
// Solidity: function getLastPrice(address tokenFrom, address tokenTo) view returns(uint256)
func (_PriceOracle *PriceOracleCaller) GetLastPrice(opts *bind.CallOpts, tokenFrom common.Address, tokenTo common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PriceOracle.contract.Call(opts, &out, "getLastPrice", tokenFrom, tokenTo)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastPrice is a free data retrieval call binding the contract method 0x743b9086.
//
// Solidity: function getLastPrice(address tokenFrom, address tokenTo) view returns(uint256)
func (_PriceOracle *PriceOracleSession) GetLastPrice(tokenFrom common.Address, tokenTo common.Address) (*big.Int, error) {
	return _PriceOracle.Contract.GetLastPrice(&_PriceOracle.CallOpts, tokenFrom, tokenTo)
}

// GetLastPrice is a free data retrieval call binding the contract method 0x743b9086.
//
// Solidity: function getLastPrice(address tokenFrom, address tokenTo) view returns(uint256)
func (_PriceOracle *PriceOracleCallerSession) GetLastPrice(tokenFrom common.Address, tokenTo common.Address) (*big.Int, error) {
	return _PriceOracle.Contract.GetLastPrice(&_PriceOracle.CallOpts, tokenFrom, tokenTo)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PriceOracle *PriceOracleCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PriceOracle.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PriceOracle *PriceOracleSession) Paused() (bool, error) {
	return _PriceOracle.Contract.Paused(&_PriceOracle.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PriceOracle *PriceOracleCallerSession) Paused() (bool, error) {
	return _PriceOracle.Contract.Paused(&_PriceOracle.CallOpts)
}

// PriceFeeds is a free data retrieval call binding the contract method 0x9dcb511a.
//
// Solidity: function priceFeeds(address ) view returns(address)
func (_PriceOracle *PriceOracleCaller) PriceFeeds(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _PriceOracle.contract.Call(opts, &out, "priceFeeds", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PriceFeeds is a free data retrieval call binding the contract method 0x9dcb511a.
//
// Solidity: function priceFeeds(address ) view returns(address)
func (_PriceOracle *PriceOracleSession) PriceFeeds(arg0 common.Address) (common.Address, error) {
	return _PriceOracle.Contract.PriceFeeds(&_PriceOracle.CallOpts, arg0)
}

// PriceFeeds is a free data retrieval call binding the contract method 0x9dcb511a.
//
// Solidity: function priceFeeds(address ) view returns(address)
func (_PriceOracle *PriceOracleCallerSession) PriceFeeds(arg0 common.Address) (common.Address, error) {
	return _PriceOracle.Contract.PriceFeeds(&_PriceOracle.CallOpts, arg0)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_PriceOracle *PriceOracleCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PriceOracle.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_PriceOracle *PriceOracleSession) Version() (*big.Int, error) {
	return _PriceOracle.Contract.Version(&_PriceOracle.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_PriceOracle *PriceOracleCallerSession) Version() (*big.Int, error) {
	return _PriceOracle.Contract.Version(&_PriceOracle.CallOpts)
}

// WethAddress is a free data retrieval call binding the contract method 0x4f0e0ef3.
//
// Solidity: function wethAddress() view returns(address)
func (_PriceOracle *PriceOracleCaller) WethAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PriceOracle.contract.Call(opts, &out, "wethAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WethAddress is a free data retrieval call binding the contract method 0x4f0e0ef3.
//
// Solidity: function wethAddress() view returns(address)
func (_PriceOracle *PriceOracleSession) WethAddress() (common.Address, error) {
	return _PriceOracle.Contract.WethAddress(&_PriceOracle.CallOpts)
}

// WethAddress is a free data retrieval call binding the contract method 0x4f0e0ef3.
//
// Solidity: function wethAddress() view returns(address)
func (_PriceOracle *PriceOracleCallerSession) WethAddress() (common.Address, error) {
	return _PriceOracle.Contract.WethAddress(&_PriceOracle.CallOpts)
}

// AddPriceFeed is a paid mutator transaction binding the contract method 0xe8a97a3e.
//
// Solidity: function addPriceFeed(address token, address priceFeed) returns()
func (_PriceOracle *PriceOracleTransactor) AddPriceFeed(opts *bind.TransactOpts, token common.Address, priceFeed common.Address) (*types.Transaction, error) {
	return _PriceOracle.contract.Transact(opts, "addPriceFeed", token, priceFeed)
}

// AddPriceFeed is a paid mutator transaction binding the contract method 0xe8a97a3e.
//
// Solidity: function addPriceFeed(address token, address priceFeed) returns()
func (_PriceOracle *PriceOracleSession) AddPriceFeed(token common.Address, priceFeed common.Address) (*types.Transaction, error) {
	return _PriceOracle.Contract.AddPriceFeed(&_PriceOracle.TransactOpts, token, priceFeed)
}

// AddPriceFeed is a paid mutator transaction binding the contract method 0xe8a97a3e.
//
// Solidity: function addPriceFeed(address token, address priceFeed) returns()
func (_PriceOracle *PriceOracleTransactorSession) AddPriceFeed(token common.Address, priceFeed common.Address) (*types.Transaction, error) {
	return _PriceOracle.Contract.AddPriceFeed(&_PriceOracle.TransactOpts, token, priceFeed)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PriceOracle *PriceOracleTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceOracle.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PriceOracle *PriceOracleSession) Pause() (*types.Transaction, error) {
	return _PriceOracle.Contract.Pause(&_PriceOracle.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PriceOracle *PriceOracleTransactorSession) Pause() (*types.Transaction, error) {
	return _PriceOracle.Contract.Pause(&_PriceOracle.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PriceOracle *PriceOracleTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceOracle.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PriceOracle *PriceOracleSession) Unpause() (*types.Transaction, error) {
	return _PriceOracle.Contract.Unpause(&_PriceOracle.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PriceOracle *PriceOracleTransactorSession) Unpause() (*types.Transaction, error) {
	return _PriceOracle.Contract.Unpause(&_PriceOracle.TransactOpts)
}

// PriceOracleNewPriceFeedIterator is returned from FilterNewPriceFeed and is used to iterate over the raw logs and unpacked data for NewPriceFeed events raised by the PriceOracle contract.
type PriceOracleNewPriceFeedIterator struct {
	Event *PriceOracleNewPriceFeed // Event containing the contract specifics and raw log

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
func (it *PriceOracleNewPriceFeedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceOracleNewPriceFeed)
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
		it.Event = new(PriceOracleNewPriceFeed)
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
func (it *PriceOracleNewPriceFeedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceOracleNewPriceFeedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceOracleNewPriceFeed represents a NewPriceFeed event raised by the PriceOracle contract.
type PriceOracleNewPriceFeed struct {
	Token     common.Address
	PriceFeed common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewPriceFeed is a free log retrieval operation binding the contract event 0xe263805b03657ab13064915d0723c5ce14981547e7cba5283f66b9e5d81f6e6e.
//
// Solidity: event NewPriceFeed(address indexed token, address indexed priceFeed)
func (_PriceOracle *PriceOracleFilterer) FilterNewPriceFeed(opts *bind.FilterOpts, token []common.Address, priceFeed []common.Address) (*PriceOracleNewPriceFeedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var priceFeedRule []interface{}
	for _, priceFeedItem := range priceFeed {
		priceFeedRule = append(priceFeedRule, priceFeedItem)
	}

	logs, sub, err := _PriceOracle.contract.FilterLogs(opts, "NewPriceFeed", tokenRule, priceFeedRule)
	if err != nil {
		return nil, err
	}
	return &PriceOracleNewPriceFeedIterator{contract: _PriceOracle.contract, event: "NewPriceFeed", logs: logs, sub: sub}, nil
}

// WatchNewPriceFeed is a free log subscription operation binding the contract event 0xe263805b03657ab13064915d0723c5ce14981547e7cba5283f66b9e5d81f6e6e.
//
// Solidity: event NewPriceFeed(address indexed token, address indexed priceFeed)
func (_PriceOracle *PriceOracleFilterer) WatchNewPriceFeed(opts *bind.WatchOpts, sink chan<- *PriceOracleNewPriceFeed, token []common.Address, priceFeed []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var priceFeedRule []interface{}
	for _, priceFeedItem := range priceFeed {
		priceFeedRule = append(priceFeedRule, priceFeedItem)
	}

	logs, sub, err := _PriceOracle.contract.WatchLogs(opts, "NewPriceFeed", tokenRule, priceFeedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceOracleNewPriceFeed)
				if err := _PriceOracle.contract.UnpackLog(event, "NewPriceFeed", log); err != nil {
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

// ParseNewPriceFeed is a log parse operation binding the contract event 0xe263805b03657ab13064915d0723c5ce14981547e7cba5283f66b9e5d81f6e6e.
//
// Solidity: event NewPriceFeed(address indexed token, address indexed priceFeed)
func (_PriceOracle *PriceOracleFilterer) ParseNewPriceFeed(log types.Log) (*PriceOracleNewPriceFeed, error) {
	event := new(PriceOracleNewPriceFeed)
	if err := _PriceOracle.contract.UnpackLog(event, "NewPriceFeed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PriceOraclePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the PriceOracle contract.
type PriceOraclePausedIterator struct {
	Event *PriceOraclePaused // Event containing the contract specifics and raw log

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
func (it *PriceOraclePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceOraclePaused)
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
		it.Event = new(PriceOraclePaused)
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
func (it *PriceOraclePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceOraclePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceOraclePaused represents a Paused event raised by the PriceOracle contract.
type PriceOraclePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PriceOracle *PriceOracleFilterer) FilterPaused(opts *bind.FilterOpts) (*PriceOraclePausedIterator, error) {

	logs, sub, err := _PriceOracle.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &PriceOraclePausedIterator{contract: _PriceOracle.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PriceOracle *PriceOracleFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PriceOraclePaused) (event.Subscription, error) {

	logs, sub, err := _PriceOracle.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceOraclePaused)
				if err := _PriceOracle.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_PriceOracle *PriceOracleFilterer) ParsePaused(log types.Log) (*PriceOraclePaused, error) {
	event := new(PriceOraclePaused)
	if err := _PriceOracle.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PriceOracleUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the PriceOracle contract.
type PriceOracleUnpausedIterator struct {
	Event *PriceOracleUnpaused // Event containing the contract specifics and raw log

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
func (it *PriceOracleUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceOracleUnpaused)
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
		it.Event = new(PriceOracleUnpaused)
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
func (it *PriceOracleUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceOracleUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceOracleUnpaused represents a Unpaused event raised by the PriceOracle contract.
type PriceOracleUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PriceOracle *PriceOracleFilterer) FilterUnpaused(opts *bind.FilterOpts) (*PriceOracleUnpausedIterator, error) {

	logs, sub, err := _PriceOracle.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &PriceOracleUnpausedIterator{contract: _PriceOracle.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PriceOracle *PriceOracleFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PriceOracleUnpaused) (event.Subscription, error) {

	logs, sub, err := _PriceOracle.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceOracleUnpaused)
				if err := _PriceOracle.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_PriceOracle *PriceOracleFilterer) ParseUnpaused(log types.Log) (*PriceOracleUnpaused, error) {
	event := new(PriceOracleUnpaused)
	if err := _PriceOracle.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	return event, nil
}
