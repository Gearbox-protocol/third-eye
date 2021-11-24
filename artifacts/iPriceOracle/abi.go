// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iPriceOracle

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

// IPriceOracleMetaData contains all meta data concerning the IPriceOracle contract.
var IPriceOracleMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"priceFeed\",\"type\":\"address\"}],\"name\":\"NewPriceFeed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceFeedToken\",\"type\":\"address\"}],\"name\":\"addPriceFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenFrom\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenTo\",\"type\":\"address\"}],\"name\":\"convert\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenFrom\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenTo\",\"type\":\"address\"}],\"name\":\"getLastPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IPriceOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use IPriceOracleMetaData.ABI instead.
var IPriceOracleABI = IPriceOracleMetaData.ABI

// IPriceOracle is an auto generated Go binding around an Ethereum contract.
type IPriceOracle struct {
	IPriceOracleCaller     // Read-only binding to the contract
	IPriceOracleTransactor // Write-only binding to the contract
	IPriceOracleFilterer   // Log filterer for contract events
}

// IPriceOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type IPriceOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPriceOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IPriceOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPriceOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPriceOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPriceOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPriceOracleSession struct {
	Contract     *IPriceOracle     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPriceOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPriceOracleCallerSession struct {
	Contract *IPriceOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IPriceOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPriceOracleTransactorSession struct {
	Contract     *IPriceOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IPriceOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type IPriceOracleRaw struct {
	Contract *IPriceOracle // Generic contract binding to access the raw methods on
}

// IPriceOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPriceOracleCallerRaw struct {
	Contract *IPriceOracleCaller // Generic read-only contract binding to access the raw methods on
}

// IPriceOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPriceOracleTransactorRaw struct {
	Contract *IPriceOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIPriceOracle creates a new instance of IPriceOracle, bound to a specific deployed contract.
func NewIPriceOracle(address common.Address, backend bind.ContractBackend) (*IPriceOracle, error) {
	contract, err := bindIPriceOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPriceOracle{IPriceOracleCaller: IPriceOracleCaller{contract: contract}, IPriceOracleTransactor: IPriceOracleTransactor{contract: contract}, IPriceOracleFilterer: IPriceOracleFilterer{contract: contract}}, nil
}

// NewIPriceOracleCaller creates a new read-only instance of IPriceOracle, bound to a specific deployed contract.
func NewIPriceOracleCaller(address common.Address, caller bind.ContractCaller) (*IPriceOracleCaller, error) {
	contract, err := bindIPriceOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPriceOracleCaller{contract: contract}, nil
}

// NewIPriceOracleTransactor creates a new write-only instance of IPriceOracle, bound to a specific deployed contract.
func NewIPriceOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*IPriceOracleTransactor, error) {
	contract, err := bindIPriceOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPriceOracleTransactor{contract: contract}, nil
}

// NewIPriceOracleFilterer creates a new log filterer instance of IPriceOracle, bound to a specific deployed contract.
func NewIPriceOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*IPriceOracleFilterer, error) {
	contract, err := bindIPriceOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPriceOracleFilterer{contract: contract}, nil
}

// bindIPriceOracle binds a generic wrapper to an already deployed contract.
func bindIPriceOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IPriceOracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPriceOracle *IPriceOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPriceOracle.Contract.IPriceOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPriceOracle *IPriceOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPriceOracle.Contract.IPriceOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPriceOracle *IPriceOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPriceOracle.Contract.IPriceOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPriceOracle *IPriceOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPriceOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPriceOracle *IPriceOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPriceOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPriceOracle *IPriceOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPriceOracle.Contract.contract.Transact(opts, method, params...)
}

// Convert is a free data retrieval call binding the contract method 0xb66102df.
//
// Solidity: function convert(uint256 amount, address tokenFrom, address tokenTo) view returns(uint256)
func (_IPriceOracle *IPriceOracleCaller) Convert(opts *bind.CallOpts, amount *big.Int, tokenFrom common.Address, tokenTo common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IPriceOracle.contract.Call(opts, &out, "convert", amount, tokenFrom, tokenTo)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Convert is a free data retrieval call binding the contract method 0xb66102df.
//
// Solidity: function convert(uint256 amount, address tokenFrom, address tokenTo) view returns(uint256)
func (_IPriceOracle *IPriceOracleSession) Convert(amount *big.Int, tokenFrom common.Address, tokenTo common.Address) (*big.Int, error) {
	return _IPriceOracle.Contract.Convert(&_IPriceOracle.CallOpts, amount, tokenFrom, tokenTo)
}

// Convert is a free data retrieval call binding the contract method 0xb66102df.
//
// Solidity: function convert(uint256 amount, address tokenFrom, address tokenTo) view returns(uint256)
func (_IPriceOracle *IPriceOracleCallerSession) Convert(amount *big.Int, tokenFrom common.Address, tokenTo common.Address) (*big.Int, error) {
	return _IPriceOracle.Contract.Convert(&_IPriceOracle.CallOpts, amount, tokenFrom, tokenTo)
}

// GetLastPrice is a free data retrieval call binding the contract method 0x743b9086.
//
// Solidity: function getLastPrice(address tokenFrom, address tokenTo) view returns(uint256)
func (_IPriceOracle *IPriceOracleCaller) GetLastPrice(opts *bind.CallOpts, tokenFrom common.Address, tokenTo common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IPriceOracle.contract.Call(opts, &out, "getLastPrice", tokenFrom, tokenTo)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastPrice is a free data retrieval call binding the contract method 0x743b9086.
//
// Solidity: function getLastPrice(address tokenFrom, address tokenTo) view returns(uint256)
func (_IPriceOracle *IPriceOracleSession) GetLastPrice(tokenFrom common.Address, tokenTo common.Address) (*big.Int, error) {
	return _IPriceOracle.Contract.GetLastPrice(&_IPriceOracle.CallOpts, tokenFrom, tokenTo)
}

// GetLastPrice is a free data retrieval call binding the contract method 0x743b9086.
//
// Solidity: function getLastPrice(address tokenFrom, address tokenTo) view returns(uint256)
func (_IPriceOracle *IPriceOracleCallerSession) GetLastPrice(tokenFrom common.Address, tokenTo common.Address) (*big.Int, error) {
	return _IPriceOracle.Contract.GetLastPrice(&_IPriceOracle.CallOpts, tokenFrom, tokenTo)
}

// AddPriceFeed is a paid mutator transaction binding the contract method 0xe8a97a3e.
//
// Solidity: function addPriceFeed(address token, address priceFeedToken) returns()
func (_IPriceOracle *IPriceOracleTransactor) AddPriceFeed(opts *bind.TransactOpts, token common.Address, priceFeedToken common.Address) (*types.Transaction, error) {
	return _IPriceOracle.contract.Transact(opts, "addPriceFeed", token, priceFeedToken)
}

// AddPriceFeed is a paid mutator transaction binding the contract method 0xe8a97a3e.
//
// Solidity: function addPriceFeed(address token, address priceFeedToken) returns()
func (_IPriceOracle *IPriceOracleSession) AddPriceFeed(token common.Address, priceFeedToken common.Address) (*types.Transaction, error) {
	return _IPriceOracle.Contract.AddPriceFeed(&_IPriceOracle.TransactOpts, token, priceFeedToken)
}

// AddPriceFeed is a paid mutator transaction binding the contract method 0xe8a97a3e.
//
// Solidity: function addPriceFeed(address token, address priceFeedToken) returns()
func (_IPriceOracle *IPriceOracleTransactorSession) AddPriceFeed(token common.Address, priceFeedToken common.Address) (*types.Transaction, error) {
	return _IPriceOracle.Contract.AddPriceFeed(&_IPriceOracle.TransactOpts, token, priceFeedToken)
}

// IPriceOracleNewPriceFeedIterator is returned from FilterNewPriceFeed and is used to iterate over the raw logs and unpacked data for NewPriceFeed events raised by the IPriceOracle contract.
type IPriceOracleNewPriceFeedIterator struct {
	Event *IPriceOracleNewPriceFeed // Event containing the contract specifics and raw log

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
func (it *IPriceOracleNewPriceFeedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPriceOracleNewPriceFeed)
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
		it.Event = new(IPriceOracleNewPriceFeed)
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
func (it *IPriceOracleNewPriceFeedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPriceOracleNewPriceFeedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPriceOracleNewPriceFeed represents a NewPriceFeed event raised by the IPriceOracle contract.
type IPriceOracleNewPriceFeed struct {
	Token     common.Address
	PriceFeed common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewPriceFeed is a free log retrieval operation binding the contract event 0xe263805b03657ab13064915d0723c5ce14981547e7cba5283f66b9e5d81f6e6e.
//
// Solidity: event NewPriceFeed(address indexed token, address indexed priceFeed)
func (_IPriceOracle *IPriceOracleFilterer) FilterNewPriceFeed(opts *bind.FilterOpts, token []common.Address, priceFeed []common.Address) (*IPriceOracleNewPriceFeedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var priceFeedRule []interface{}
	for _, priceFeedItem := range priceFeed {
		priceFeedRule = append(priceFeedRule, priceFeedItem)
	}

	logs, sub, err := _IPriceOracle.contract.FilterLogs(opts, "NewPriceFeed", tokenRule, priceFeedRule)
	if err != nil {
		return nil, err
	}
	return &IPriceOracleNewPriceFeedIterator{contract: _IPriceOracle.contract, event: "NewPriceFeed", logs: logs, sub: sub}, nil
}

// WatchNewPriceFeed is a free log subscription operation binding the contract event 0xe263805b03657ab13064915d0723c5ce14981547e7cba5283f66b9e5d81f6e6e.
//
// Solidity: event NewPriceFeed(address indexed token, address indexed priceFeed)
func (_IPriceOracle *IPriceOracleFilterer) WatchNewPriceFeed(opts *bind.WatchOpts, sink chan<- *IPriceOracleNewPriceFeed, token []common.Address, priceFeed []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var priceFeedRule []interface{}
	for _, priceFeedItem := range priceFeed {
		priceFeedRule = append(priceFeedRule, priceFeedItem)
	}

	logs, sub, err := _IPriceOracle.contract.WatchLogs(opts, "NewPriceFeed", tokenRule, priceFeedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPriceOracleNewPriceFeed)
				if err := _IPriceOracle.contract.UnpackLog(event, "NewPriceFeed", log); err != nil {
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
func (_IPriceOracle *IPriceOracleFilterer) ParseNewPriceFeed(log types.Log) (*IPriceOracleNewPriceFeed, error) {
	event := new(IPriceOracleNewPriceFeed)
	if err := _IPriceOracle.contract.UnpackLog(event, "NewPriceFeed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
