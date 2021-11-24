// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iAccountFactory

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

// IAccountFactoryMetaData contains all meta data concerning the IAccountFactory contract.
var IAccountFactoryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"}],\"name\":\"AccountMinerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creditManager\",\"type\":\"address\"}],\"name\":\"InitializeCreditAccount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"NewCreditAccount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ReturnCreditAccount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"TakeForever\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"countCreditAccounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"countCreditAccountsInStock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"creditAccounts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"}],\"name\":\"getNext\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"head\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usedAccount\",\"type\":\"address\"}],\"name\":\"returnCreditAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tail\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_borrowedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_cumulativeIndexAtOpen\",\"type\":\"uint256\"}],\"name\":\"takeCreditAccount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IAccountFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use IAccountFactoryMetaData.ABI instead.
var IAccountFactoryABI = IAccountFactoryMetaData.ABI

// IAccountFactory is an auto generated Go binding around an Ethereum contract.
type IAccountFactory struct {
	IAccountFactoryCaller     // Read-only binding to the contract
	IAccountFactoryTransactor // Write-only binding to the contract
	IAccountFactoryFilterer   // Log filterer for contract events
}

// IAccountFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IAccountFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAccountFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IAccountFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAccountFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAccountFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAccountFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IAccountFactorySession struct {
	Contract     *IAccountFactory  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAccountFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IAccountFactoryCallerSession struct {
	Contract *IAccountFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IAccountFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IAccountFactoryTransactorSession struct {
	Contract     *IAccountFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IAccountFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IAccountFactoryRaw struct {
	Contract *IAccountFactory // Generic contract binding to access the raw methods on
}

// IAccountFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IAccountFactoryCallerRaw struct {
	Contract *IAccountFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// IAccountFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IAccountFactoryTransactorRaw struct {
	Contract *IAccountFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIAccountFactory creates a new instance of IAccountFactory, bound to a specific deployed contract.
func NewIAccountFactory(address common.Address, backend bind.ContractBackend) (*IAccountFactory, error) {
	contract, err := bindIAccountFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAccountFactory{IAccountFactoryCaller: IAccountFactoryCaller{contract: contract}, IAccountFactoryTransactor: IAccountFactoryTransactor{contract: contract}, IAccountFactoryFilterer: IAccountFactoryFilterer{contract: contract}}, nil
}

// NewIAccountFactoryCaller creates a new read-only instance of IAccountFactory, bound to a specific deployed contract.
func NewIAccountFactoryCaller(address common.Address, caller bind.ContractCaller) (*IAccountFactoryCaller, error) {
	contract, err := bindIAccountFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAccountFactoryCaller{contract: contract}, nil
}

// NewIAccountFactoryTransactor creates a new write-only instance of IAccountFactory, bound to a specific deployed contract.
func NewIAccountFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*IAccountFactoryTransactor, error) {
	contract, err := bindIAccountFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAccountFactoryTransactor{contract: contract}, nil
}

// NewIAccountFactoryFilterer creates a new log filterer instance of IAccountFactory, bound to a specific deployed contract.
func NewIAccountFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*IAccountFactoryFilterer, error) {
	contract, err := bindIAccountFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAccountFactoryFilterer{contract: contract}, nil
}

// bindIAccountFactory binds a generic wrapper to an already deployed contract.
func bindIAccountFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IAccountFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAccountFactory *IAccountFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAccountFactory.Contract.IAccountFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAccountFactory *IAccountFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAccountFactory.Contract.IAccountFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAccountFactory *IAccountFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAccountFactory.Contract.IAccountFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAccountFactory *IAccountFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAccountFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAccountFactory *IAccountFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAccountFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAccountFactory *IAccountFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAccountFactory.Contract.contract.Transact(opts, method, params...)
}

// CountCreditAccounts is a free data retrieval call binding the contract method 0xb60e8518.
//
// Solidity: function countCreditAccounts() view returns(uint256)
func (_IAccountFactory *IAccountFactoryCaller) CountCreditAccounts(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IAccountFactory.contract.Call(opts, &out, "countCreditAccounts")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CountCreditAccounts is a free data retrieval call binding the contract method 0xb60e8518.
//
// Solidity: function countCreditAccounts() view returns(uint256)
func (_IAccountFactory *IAccountFactorySession) CountCreditAccounts() (*big.Int, error) {
	return _IAccountFactory.Contract.CountCreditAccounts(&_IAccountFactory.CallOpts)
}

// CountCreditAccounts is a free data retrieval call binding the contract method 0xb60e8518.
//
// Solidity: function countCreditAccounts() view returns(uint256)
func (_IAccountFactory *IAccountFactoryCallerSession) CountCreditAccounts() (*big.Int, error) {
	return _IAccountFactory.Contract.CountCreditAccounts(&_IAccountFactory.CallOpts)
}

// CountCreditAccountsInStock is a free data retrieval call binding the contract method 0xb1939763.
//
// Solidity: function countCreditAccountsInStock() view returns(uint256)
func (_IAccountFactory *IAccountFactoryCaller) CountCreditAccountsInStock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IAccountFactory.contract.Call(opts, &out, "countCreditAccountsInStock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CountCreditAccountsInStock is a free data retrieval call binding the contract method 0xb1939763.
//
// Solidity: function countCreditAccountsInStock() view returns(uint256)
func (_IAccountFactory *IAccountFactorySession) CountCreditAccountsInStock() (*big.Int, error) {
	return _IAccountFactory.Contract.CountCreditAccountsInStock(&_IAccountFactory.CallOpts)
}

// CountCreditAccountsInStock is a free data retrieval call binding the contract method 0xb1939763.
//
// Solidity: function countCreditAccountsInStock() view returns(uint256)
func (_IAccountFactory *IAccountFactoryCallerSession) CountCreditAccountsInStock() (*big.Int, error) {
	return _IAccountFactory.Contract.CountCreditAccountsInStock(&_IAccountFactory.CallOpts)
}

// CreditAccounts is a free data retrieval call binding the contract method 0xe3ba9ace.
//
// Solidity: function creditAccounts(uint256 id) view returns(address)
func (_IAccountFactory *IAccountFactoryCaller) CreditAccounts(opts *bind.CallOpts, id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IAccountFactory.contract.Call(opts, &out, "creditAccounts", id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditAccounts is a free data retrieval call binding the contract method 0xe3ba9ace.
//
// Solidity: function creditAccounts(uint256 id) view returns(address)
func (_IAccountFactory *IAccountFactorySession) CreditAccounts(id *big.Int) (common.Address, error) {
	return _IAccountFactory.Contract.CreditAccounts(&_IAccountFactory.CallOpts, id)
}

// CreditAccounts is a free data retrieval call binding the contract method 0xe3ba9ace.
//
// Solidity: function creditAccounts(uint256 id) view returns(address)
func (_IAccountFactory *IAccountFactoryCallerSession) CreditAccounts(id *big.Int) (common.Address, error) {
	return _IAccountFactory.Contract.CreditAccounts(&_IAccountFactory.CallOpts, id)
}

// GetNext is a free data retrieval call binding the contract method 0x765e0159.
//
// Solidity: function getNext(address creditAccount) view returns(address)
func (_IAccountFactory *IAccountFactoryCaller) GetNext(opts *bind.CallOpts, creditAccount common.Address) (common.Address, error) {
	var out []interface{}
	err := _IAccountFactory.contract.Call(opts, &out, "getNext", creditAccount)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetNext is a free data retrieval call binding the contract method 0x765e0159.
//
// Solidity: function getNext(address creditAccount) view returns(address)
func (_IAccountFactory *IAccountFactorySession) GetNext(creditAccount common.Address) (common.Address, error) {
	return _IAccountFactory.Contract.GetNext(&_IAccountFactory.CallOpts, creditAccount)
}

// GetNext is a free data retrieval call binding the contract method 0x765e0159.
//
// Solidity: function getNext(address creditAccount) view returns(address)
func (_IAccountFactory *IAccountFactoryCallerSession) GetNext(creditAccount common.Address) (common.Address, error) {
	return _IAccountFactory.Contract.GetNext(&_IAccountFactory.CallOpts, creditAccount)
}

// Head is a free data retrieval call binding the contract method 0x8f7dcfa3.
//
// Solidity: function head() view returns(address)
func (_IAccountFactory *IAccountFactoryCaller) Head(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IAccountFactory.contract.Call(opts, &out, "head")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Head is a free data retrieval call binding the contract method 0x8f7dcfa3.
//
// Solidity: function head() view returns(address)
func (_IAccountFactory *IAccountFactorySession) Head() (common.Address, error) {
	return _IAccountFactory.Contract.Head(&_IAccountFactory.CallOpts)
}

// Head is a free data retrieval call binding the contract method 0x8f7dcfa3.
//
// Solidity: function head() view returns(address)
func (_IAccountFactory *IAccountFactoryCallerSession) Head() (common.Address, error) {
	return _IAccountFactory.Contract.Head(&_IAccountFactory.CallOpts)
}

// Tail is a free data retrieval call binding the contract method 0x13d8c840.
//
// Solidity: function tail() view returns(address)
func (_IAccountFactory *IAccountFactoryCaller) Tail(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IAccountFactory.contract.Call(opts, &out, "tail")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Tail is a free data retrieval call binding the contract method 0x13d8c840.
//
// Solidity: function tail() view returns(address)
func (_IAccountFactory *IAccountFactorySession) Tail() (common.Address, error) {
	return _IAccountFactory.Contract.Tail(&_IAccountFactory.CallOpts)
}

// Tail is a free data retrieval call binding the contract method 0x13d8c840.
//
// Solidity: function tail() view returns(address)
func (_IAccountFactory *IAccountFactoryCallerSession) Tail() (common.Address, error) {
	return _IAccountFactory.Contract.Tail(&_IAccountFactory.CallOpts)
}

// ReturnCreditAccount is a paid mutator transaction binding the contract method 0x89b77b3e.
//
// Solidity: function returnCreditAccount(address usedAccount) returns()
func (_IAccountFactory *IAccountFactoryTransactor) ReturnCreditAccount(opts *bind.TransactOpts, usedAccount common.Address) (*types.Transaction, error) {
	return _IAccountFactory.contract.Transact(opts, "returnCreditAccount", usedAccount)
}

// ReturnCreditAccount is a paid mutator transaction binding the contract method 0x89b77b3e.
//
// Solidity: function returnCreditAccount(address usedAccount) returns()
func (_IAccountFactory *IAccountFactorySession) ReturnCreditAccount(usedAccount common.Address) (*types.Transaction, error) {
	return _IAccountFactory.Contract.ReturnCreditAccount(&_IAccountFactory.TransactOpts, usedAccount)
}

// ReturnCreditAccount is a paid mutator transaction binding the contract method 0x89b77b3e.
//
// Solidity: function returnCreditAccount(address usedAccount) returns()
func (_IAccountFactory *IAccountFactoryTransactorSession) ReturnCreditAccount(usedAccount common.Address) (*types.Transaction, error) {
	return _IAccountFactory.Contract.ReturnCreditAccount(&_IAccountFactory.TransactOpts, usedAccount)
}

// TakeCreditAccount is a paid mutator transaction binding the contract method 0x21d18456.
//
// Solidity: function takeCreditAccount(uint256 _borrowedAmount, uint256 _cumulativeIndexAtOpen) returns(address)
func (_IAccountFactory *IAccountFactoryTransactor) TakeCreditAccount(opts *bind.TransactOpts, _borrowedAmount *big.Int, _cumulativeIndexAtOpen *big.Int) (*types.Transaction, error) {
	return _IAccountFactory.contract.Transact(opts, "takeCreditAccount", _borrowedAmount, _cumulativeIndexAtOpen)
}

// TakeCreditAccount is a paid mutator transaction binding the contract method 0x21d18456.
//
// Solidity: function takeCreditAccount(uint256 _borrowedAmount, uint256 _cumulativeIndexAtOpen) returns(address)
func (_IAccountFactory *IAccountFactorySession) TakeCreditAccount(_borrowedAmount *big.Int, _cumulativeIndexAtOpen *big.Int) (*types.Transaction, error) {
	return _IAccountFactory.Contract.TakeCreditAccount(&_IAccountFactory.TransactOpts, _borrowedAmount, _cumulativeIndexAtOpen)
}

// TakeCreditAccount is a paid mutator transaction binding the contract method 0x21d18456.
//
// Solidity: function takeCreditAccount(uint256 _borrowedAmount, uint256 _cumulativeIndexAtOpen) returns(address)
func (_IAccountFactory *IAccountFactoryTransactorSession) TakeCreditAccount(_borrowedAmount *big.Int, _cumulativeIndexAtOpen *big.Int) (*types.Transaction, error) {
	return _IAccountFactory.Contract.TakeCreditAccount(&_IAccountFactory.TransactOpts, _borrowedAmount, _cumulativeIndexAtOpen)
}

// IAccountFactoryAccountMinerChangedIterator is returned from FilterAccountMinerChanged and is used to iterate over the raw logs and unpacked data for AccountMinerChanged events raised by the IAccountFactory contract.
type IAccountFactoryAccountMinerChangedIterator struct {
	Event *IAccountFactoryAccountMinerChanged // Event containing the contract specifics and raw log

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
func (it *IAccountFactoryAccountMinerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAccountFactoryAccountMinerChanged)
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
		it.Event = new(IAccountFactoryAccountMinerChanged)
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
func (it *IAccountFactoryAccountMinerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAccountFactoryAccountMinerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAccountFactoryAccountMinerChanged represents a AccountMinerChanged event raised by the IAccountFactory contract.
type IAccountFactoryAccountMinerChanged struct {
	Miner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAccountMinerChanged is a free log retrieval operation binding the contract event 0xb5b22c95380a75185488532000fd4826f19e58c5eba212f266d9861a44b671fc.
//
// Solidity: event AccountMinerChanged(address indexed miner)
func (_IAccountFactory *IAccountFactoryFilterer) FilterAccountMinerChanged(opts *bind.FilterOpts, miner []common.Address) (*IAccountFactoryAccountMinerChangedIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _IAccountFactory.contract.FilterLogs(opts, "AccountMinerChanged", minerRule)
	if err != nil {
		return nil, err
	}
	return &IAccountFactoryAccountMinerChangedIterator{contract: _IAccountFactory.contract, event: "AccountMinerChanged", logs: logs, sub: sub}, nil
}

// WatchAccountMinerChanged is a free log subscription operation binding the contract event 0xb5b22c95380a75185488532000fd4826f19e58c5eba212f266d9861a44b671fc.
//
// Solidity: event AccountMinerChanged(address indexed miner)
func (_IAccountFactory *IAccountFactoryFilterer) WatchAccountMinerChanged(opts *bind.WatchOpts, sink chan<- *IAccountFactoryAccountMinerChanged, miner []common.Address) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _IAccountFactory.contract.WatchLogs(opts, "AccountMinerChanged", minerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAccountFactoryAccountMinerChanged)
				if err := _IAccountFactory.contract.UnpackLog(event, "AccountMinerChanged", log); err != nil {
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

// ParseAccountMinerChanged is a log parse operation binding the contract event 0xb5b22c95380a75185488532000fd4826f19e58c5eba212f266d9861a44b671fc.
//
// Solidity: event AccountMinerChanged(address indexed miner)
func (_IAccountFactory *IAccountFactoryFilterer) ParseAccountMinerChanged(log types.Log) (*IAccountFactoryAccountMinerChanged, error) {
	event := new(IAccountFactoryAccountMinerChanged)
	if err := _IAccountFactory.contract.UnpackLog(event, "AccountMinerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAccountFactoryInitializeCreditAccountIterator is returned from FilterInitializeCreditAccount and is used to iterate over the raw logs and unpacked data for InitializeCreditAccount events raised by the IAccountFactory contract.
type IAccountFactoryInitializeCreditAccountIterator struct {
	Event *IAccountFactoryInitializeCreditAccount // Event containing the contract specifics and raw log

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
func (it *IAccountFactoryInitializeCreditAccountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAccountFactoryInitializeCreditAccount)
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
		it.Event = new(IAccountFactoryInitializeCreditAccount)
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
func (it *IAccountFactoryInitializeCreditAccountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAccountFactoryInitializeCreditAccountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAccountFactoryInitializeCreditAccount represents a InitializeCreditAccount event raised by the IAccountFactory contract.
type IAccountFactoryInitializeCreditAccount struct {
	Account       common.Address
	CreditManager common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterInitializeCreditAccount is a free log retrieval operation binding the contract event 0xf3ede7039176503a8ad1fe7cfaa29475a9dbe0cdcaf04ecf9a5c10570c47b103.
//
// Solidity: event InitializeCreditAccount(address indexed account, address indexed creditManager)
func (_IAccountFactory *IAccountFactoryFilterer) FilterInitializeCreditAccount(opts *bind.FilterOpts, account []common.Address, creditManager []common.Address) (*IAccountFactoryInitializeCreditAccountIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var creditManagerRule []interface{}
	for _, creditManagerItem := range creditManager {
		creditManagerRule = append(creditManagerRule, creditManagerItem)
	}

	logs, sub, err := _IAccountFactory.contract.FilterLogs(opts, "InitializeCreditAccount", accountRule, creditManagerRule)
	if err != nil {
		return nil, err
	}
	return &IAccountFactoryInitializeCreditAccountIterator{contract: _IAccountFactory.contract, event: "InitializeCreditAccount", logs: logs, sub: sub}, nil
}

// WatchInitializeCreditAccount is a free log subscription operation binding the contract event 0xf3ede7039176503a8ad1fe7cfaa29475a9dbe0cdcaf04ecf9a5c10570c47b103.
//
// Solidity: event InitializeCreditAccount(address indexed account, address indexed creditManager)
func (_IAccountFactory *IAccountFactoryFilterer) WatchInitializeCreditAccount(opts *bind.WatchOpts, sink chan<- *IAccountFactoryInitializeCreditAccount, account []common.Address, creditManager []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var creditManagerRule []interface{}
	for _, creditManagerItem := range creditManager {
		creditManagerRule = append(creditManagerRule, creditManagerItem)
	}

	logs, sub, err := _IAccountFactory.contract.WatchLogs(opts, "InitializeCreditAccount", accountRule, creditManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAccountFactoryInitializeCreditAccount)
				if err := _IAccountFactory.contract.UnpackLog(event, "InitializeCreditAccount", log); err != nil {
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

// ParseInitializeCreditAccount is a log parse operation binding the contract event 0xf3ede7039176503a8ad1fe7cfaa29475a9dbe0cdcaf04ecf9a5c10570c47b103.
//
// Solidity: event InitializeCreditAccount(address indexed account, address indexed creditManager)
func (_IAccountFactory *IAccountFactoryFilterer) ParseInitializeCreditAccount(log types.Log) (*IAccountFactoryInitializeCreditAccount, error) {
	event := new(IAccountFactoryInitializeCreditAccount)
	if err := _IAccountFactory.contract.UnpackLog(event, "InitializeCreditAccount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAccountFactoryNewCreditAccountIterator is returned from FilterNewCreditAccount and is used to iterate over the raw logs and unpacked data for NewCreditAccount events raised by the IAccountFactory contract.
type IAccountFactoryNewCreditAccountIterator struct {
	Event *IAccountFactoryNewCreditAccount // Event containing the contract specifics and raw log

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
func (it *IAccountFactoryNewCreditAccountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAccountFactoryNewCreditAccount)
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
		it.Event = new(IAccountFactoryNewCreditAccount)
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
func (it *IAccountFactoryNewCreditAccountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAccountFactoryNewCreditAccountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAccountFactoryNewCreditAccount represents a NewCreditAccount event raised by the IAccountFactory contract.
type IAccountFactoryNewCreditAccount struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewCreditAccount is a free log retrieval operation binding the contract event 0x9f69b6c10f6810213e055b0ba6bc0a4e2603f73c221aad77ea35da819cda7dc3.
//
// Solidity: event NewCreditAccount(address indexed account)
func (_IAccountFactory *IAccountFactoryFilterer) FilterNewCreditAccount(opts *bind.FilterOpts, account []common.Address) (*IAccountFactoryNewCreditAccountIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IAccountFactory.contract.FilterLogs(opts, "NewCreditAccount", accountRule)
	if err != nil {
		return nil, err
	}
	return &IAccountFactoryNewCreditAccountIterator{contract: _IAccountFactory.contract, event: "NewCreditAccount", logs: logs, sub: sub}, nil
}

// WatchNewCreditAccount is a free log subscription operation binding the contract event 0x9f69b6c10f6810213e055b0ba6bc0a4e2603f73c221aad77ea35da819cda7dc3.
//
// Solidity: event NewCreditAccount(address indexed account)
func (_IAccountFactory *IAccountFactoryFilterer) WatchNewCreditAccount(opts *bind.WatchOpts, sink chan<- *IAccountFactoryNewCreditAccount, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IAccountFactory.contract.WatchLogs(opts, "NewCreditAccount", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAccountFactoryNewCreditAccount)
				if err := _IAccountFactory.contract.UnpackLog(event, "NewCreditAccount", log); err != nil {
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

// ParseNewCreditAccount is a log parse operation binding the contract event 0x9f69b6c10f6810213e055b0ba6bc0a4e2603f73c221aad77ea35da819cda7dc3.
//
// Solidity: event NewCreditAccount(address indexed account)
func (_IAccountFactory *IAccountFactoryFilterer) ParseNewCreditAccount(log types.Log) (*IAccountFactoryNewCreditAccount, error) {
	event := new(IAccountFactoryNewCreditAccount)
	if err := _IAccountFactory.contract.UnpackLog(event, "NewCreditAccount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAccountFactoryReturnCreditAccountIterator is returned from FilterReturnCreditAccount and is used to iterate over the raw logs and unpacked data for ReturnCreditAccount events raised by the IAccountFactory contract.
type IAccountFactoryReturnCreditAccountIterator struct {
	Event *IAccountFactoryReturnCreditAccount // Event containing the contract specifics and raw log

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
func (it *IAccountFactoryReturnCreditAccountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAccountFactoryReturnCreditAccount)
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
		it.Event = new(IAccountFactoryReturnCreditAccount)
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
func (it *IAccountFactoryReturnCreditAccountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAccountFactoryReturnCreditAccountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAccountFactoryReturnCreditAccount represents a ReturnCreditAccount event raised by the IAccountFactory contract.
type IAccountFactoryReturnCreditAccount struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterReturnCreditAccount is a free log retrieval operation binding the contract event 0xced6ab9afc868b3a088366f6631ae20752993b5cce5d5f0534ea5a59fcc57d56.
//
// Solidity: event ReturnCreditAccount(address indexed account)
func (_IAccountFactory *IAccountFactoryFilterer) FilterReturnCreditAccount(opts *bind.FilterOpts, account []common.Address) (*IAccountFactoryReturnCreditAccountIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IAccountFactory.contract.FilterLogs(opts, "ReturnCreditAccount", accountRule)
	if err != nil {
		return nil, err
	}
	return &IAccountFactoryReturnCreditAccountIterator{contract: _IAccountFactory.contract, event: "ReturnCreditAccount", logs: logs, sub: sub}, nil
}

// WatchReturnCreditAccount is a free log subscription operation binding the contract event 0xced6ab9afc868b3a088366f6631ae20752993b5cce5d5f0534ea5a59fcc57d56.
//
// Solidity: event ReturnCreditAccount(address indexed account)
func (_IAccountFactory *IAccountFactoryFilterer) WatchReturnCreditAccount(opts *bind.WatchOpts, sink chan<- *IAccountFactoryReturnCreditAccount, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IAccountFactory.contract.WatchLogs(opts, "ReturnCreditAccount", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAccountFactoryReturnCreditAccount)
				if err := _IAccountFactory.contract.UnpackLog(event, "ReturnCreditAccount", log); err != nil {
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

// ParseReturnCreditAccount is a log parse operation binding the contract event 0xced6ab9afc868b3a088366f6631ae20752993b5cce5d5f0534ea5a59fcc57d56.
//
// Solidity: event ReturnCreditAccount(address indexed account)
func (_IAccountFactory *IAccountFactoryFilterer) ParseReturnCreditAccount(log types.Log) (*IAccountFactoryReturnCreditAccount, error) {
	event := new(IAccountFactoryReturnCreditAccount)
	if err := _IAccountFactory.contract.UnpackLog(event, "ReturnCreditAccount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAccountFactoryTakeForeverIterator is returned from FilterTakeForever and is used to iterate over the raw logs and unpacked data for TakeForever events raised by the IAccountFactory contract.
type IAccountFactoryTakeForeverIterator struct {
	Event *IAccountFactoryTakeForever // Event containing the contract specifics and raw log

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
func (it *IAccountFactoryTakeForeverIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAccountFactoryTakeForever)
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
		it.Event = new(IAccountFactoryTakeForever)
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
func (it *IAccountFactoryTakeForeverIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAccountFactoryTakeForeverIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAccountFactoryTakeForever represents a TakeForever event raised by the IAccountFactory contract.
type IAccountFactoryTakeForever struct {
	CreditAccount common.Address
	To            common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTakeForever is a free log retrieval operation binding the contract event 0x25e267469ba2ae82515be7b3d45df60bf8308343f0809e8cf7319058e2255ce6.
//
// Solidity: event TakeForever(address indexed creditAccount, address indexed to)
func (_IAccountFactory *IAccountFactoryFilterer) FilterTakeForever(opts *bind.FilterOpts, creditAccount []common.Address, to []common.Address) (*IAccountFactoryTakeForeverIterator, error) {

	var creditAccountRule []interface{}
	for _, creditAccountItem := range creditAccount {
		creditAccountRule = append(creditAccountRule, creditAccountItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IAccountFactory.contract.FilterLogs(opts, "TakeForever", creditAccountRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IAccountFactoryTakeForeverIterator{contract: _IAccountFactory.contract, event: "TakeForever", logs: logs, sub: sub}, nil
}

// WatchTakeForever is a free log subscription operation binding the contract event 0x25e267469ba2ae82515be7b3d45df60bf8308343f0809e8cf7319058e2255ce6.
//
// Solidity: event TakeForever(address indexed creditAccount, address indexed to)
func (_IAccountFactory *IAccountFactoryFilterer) WatchTakeForever(opts *bind.WatchOpts, sink chan<- *IAccountFactoryTakeForever, creditAccount []common.Address, to []common.Address) (event.Subscription, error) {

	var creditAccountRule []interface{}
	for _, creditAccountItem := range creditAccount {
		creditAccountRule = append(creditAccountRule, creditAccountItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IAccountFactory.contract.WatchLogs(opts, "TakeForever", creditAccountRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAccountFactoryTakeForever)
				if err := _IAccountFactory.contract.UnpackLog(event, "TakeForever", log); err != nil {
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

// ParseTakeForever is a log parse operation binding the contract event 0x25e267469ba2ae82515be7b3d45df60bf8308343f0809e8cf7319058e2255ce6.
//
// Solidity: event TakeForever(address indexed creditAccount, address indexed to)
func (_IAccountFactory *IAccountFactoryFilterer) ParseTakeForever(log types.Log) (*IAccountFactoryTakeForever, error) {
	event := new(IAccountFactoryTakeForever)
	if err := _IAccountFactory.contract.UnpackLog(event, "TakeForever", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
