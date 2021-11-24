// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package aCLTraitTest

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

// ACLTraitTestMetaData contains all meta data concerning the ACLTraitTest contract.
var ACLTraitTestMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addressProvider\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"accessConfiguratorOnly\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accessWhenNotPaused\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accessWhenPaused\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ACLTraitTestABI is the input ABI used to generate the binding from.
// Deprecated: Use ACLTraitTestMetaData.ABI instead.
var ACLTraitTestABI = ACLTraitTestMetaData.ABI

// ACLTraitTest is an auto generated Go binding around an Ethereum contract.
type ACLTraitTest struct {
	ACLTraitTestCaller     // Read-only binding to the contract
	ACLTraitTestTransactor // Write-only binding to the contract
	ACLTraitTestFilterer   // Log filterer for contract events
}

// ACLTraitTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type ACLTraitTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ACLTraitTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ACLTraitTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ACLTraitTestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ACLTraitTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ACLTraitTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ACLTraitTestSession struct {
	Contract     *ACLTraitTest     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ACLTraitTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ACLTraitTestCallerSession struct {
	Contract *ACLTraitTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ACLTraitTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ACLTraitTestTransactorSession struct {
	Contract     *ACLTraitTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ACLTraitTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type ACLTraitTestRaw struct {
	Contract *ACLTraitTest // Generic contract binding to access the raw methods on
}

// ACLTraitTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ACLTraitTestCallerRaw struct {
	Contract *ACLTraitTestCaller // Generic read-only contract binding to access the raw methods on
}

// ACLTraitTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ACLTraitTestTransactorRaw struct {
	Contract *ACLTraitTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewACLTraitTest creates a new instance of ACLTraitTest, bound to a specific deployed contract.
func NewACLTraitTest(address common.Address, backend bind.ContractBackend) (*ACLTraitTest, error) {
	contract, err := bindACLTraitTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ACLTraitTest{ACLTraitTestCaller: ACLTraitTestCaller{contract: contract}, ACLTraitTestTransactor: ACLTraitTestTransactor{contract: contract}, ACLTraitTestFilterer: ACLTraitTestFilterer{contract: contract}}, nil
}

// NewACLTraitTestCaller creates a new read-only instance of ACLTraitTest, bound to a specific deployed contract.
func NewACLTraitTestCaller(address common.Address, caller bind.ContractCaller) (*ACLTraitTestCaller, error) {
	contract, err := bindACLTraitTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ACLTraitTestCaller{contract: contract}, nil
}

// NewACLTraitTestTransactor creates a new write-only instance of ACLTraitTest, bound to a specific deployed contract.
func NewACLTraitTestTransactor(address common.Address, transactor bind.ContractTransactor) (*ACLTraitTestTransactor, error) {
	contract, err := bindACLTraitTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ACLTraitTestTransactor{contract: contract}, nil
}

// NewACLTraitTestFilterer creates a new log filterer instance of ACLTraitTest, bound to a specific deployed contract.
func NewACLTraitTestFilterer(address common.Address, filterer bind.ContractFilterer) (*ACLTraitTestFilterer, error) {
	contract, err := bindACLTraitTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ACLTraitTestFilterer{contract: contract}, nil
}

// bindACLTraitTest binds a generic wrapper to an already deployed contract.
func bindACLTraitTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ACLTraitTestABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ACLTraitTest *ACLTraitTestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ACLTraitTest.Contract.ACLTraitTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ACLTraitTest *ACLTraitTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ACLTraitTest.Contract.ACLTraitTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ACLTraitTest *ACLTraitTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ACLTraitTest.Contract.ACLTraitTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ACLTraitTest *ACLTraitTestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ACLTraitTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ACLTraitTest *ACLTraitTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ACLTraitTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ACLTraitTest *ACLTraitTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ACLTraitTest.Contract.contract.Transact(opts, method, params...)
}

// AccessConfiguratorOnly is a free data retrieval call binding the contract method 0x19df69ae.
//
// Solidity: function accessConfiguratorOnly() view returns()
func (_ACLTraitTest *ACLTraitTestCaller) AccessConfiguratorOnly(opts *bind.CallOpts) error {
	var out []interface{}
	err := _ACLTraitTest.contract.Call(opts, &out, "accessConfiguratorOnly")

	if err != nil {
		return err
	}

	return err

}

// AccessConfiguratorOnly is a free data retrieval call binding the contract method 0x19df69ae.
//
// Solidity: function accessConfiguratorOnly() view returns()
func (_ACLTraitTest *ACLTraitTestSession) AccessConfiguratorOnly() error {
	return _ACLTraitTest.Contract.AccessConfiguratorOnly(&_ACLTraitTest.CallOpts)
}

// AccessConfiguratorOnly is a free data retrieval call binding the contract method 0x19df69ae.
//
// Solidity: function accessConfiguratorOnly() view returns()
func (_ACLTraitTest *ACLTraitTestCallerSession) AccessConfiguratorOnly() error {
	return _ACLTraitTest.Contract.AccessConfiguratorOnly(&_ACLTraitTest.CallOpts)
}

// AccessWhenNotPaused is a free data retrieval call binding the contract method 0x50a472ae.
//
// Solidity: function accessWhenNotPaused() view returns()
func (_ACLTraitTest *ACLTraitTestCaller) AccessWhenNotPaused(opts *bind.CallOpts) error {
	var out []interface{}
	err := _ACLTraitTest.contract.Call(opts, &out, "accessWhenNotPaused")

	if err != nil {
		return err
	}

	return err

}

// AccessWhenNotPaused is a free data retrieval call binding the contract method 0x50a472ae.
//
// Solidity: function accessWhenNotPaused() view returns()
func (_ACLTraitTest *ACLTraitTestSession) AccessWhenNotPaused() error {
	return _ACLTraitTest.Contract.AccessWhenNotPaused(&_ACLTraitTest.CallOpts)
}

// AccessWhenNotPaused is a free data retrieval call binding the contract method 0x50a472ae.
//
// Solidity: function accessWhenNotPaused() view returns()
func (_ACLTraitTest *ACLTraitTestCallerSession) AccessWhenNotPaused() error {
	return _ACLTraitTest.Contract.AccessWhenNotPaused(&_ACLTraitTest.CallOpts)
}

// AccessWhenPaused is a free data retrieval call binding the contract method 0x8def0c0b.
//
// Solidity: function accessWhenPaused() view returns()
func (_ACLTraitTest *ACLTraitTestCaller) AccessWhenPaused(opts *bind.CallOpts) error {
	var out []interface{}
	err := _ACLTraitTest.contract.Call(opts, &out, "accessWhenPaused")

	if err != nil {
		return err
	}

	return err

}

// AccessWhenPaused is a free data retrieval call binding the contract method 0x8def0c0b.
//
// Solidity: function accessWhenPaused() view returns()
func (_ACLTraitTest *ACLTraitTestSession) AccessWhenPaused() error {
	return _ACLTraitTest.Contract.AccessWhenPaused(&_ACLTraitTest.CallOpts)
}

// AccessWhenPaused is a free data retrieval call binding the contract method 0x8def0c0b.
//
// Solidity: function accessWhenPaused() view returns()
func (_ACLTraitTest *ACLTraitTestCallerSession) AccessWhenPaused() error {
	return _ACLTraitTest.Contract.AccessWhenPaused(&_ACLTraitTest.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ACLTraitTest *ACLTraitTestCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ACLTraitTest.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ACLTraitTest *ACLTraitTestSession) Paused() (bool, error) {
	return _ACLTraitTest.Contract.Paused(&_ACLTraitTest.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ACLTraitTest *ACLTraitTestCallerSession) Paused() (bool, error) {
	return _ACLTraitTest.Contract.Paused(&_ACLTraitTest.CallOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ACLTraitTest *ACLTraitTestTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ACLTraitTest.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ACLTraitTest *ACLTraitTestSession) Pause() (*types.Transaction, error) {
	return _ACLTraitTest.Contract.Pause(&_ACLTraitTest.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ACLTraitTest *ACLTraitTestTransactorSession) Pause() (*types.Transaction, error) {
	return _ACLTraitTest.Contract.Pause(&_ACLTraitTest.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ACLTraitTest *ACLTraitTestTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ACLTraitTest.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ACLTraitTest *ACLTraitTestSession) Unpause() (*types.Transaction, error) {
	return _ACLTraitTest.Contract.Unpause(&_ACLTraitTest.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ACLTraitTest *ACLTraitTestTransactorSession) Unpause() (*types.Transaction, error) {
	return _ACLTraitTest.Contract.Unpause(&_ACLTraitTest.TransactOpts)
}

// ACLTraitTestPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ACLTraitTest contract.
type ACLTraitTestPausedIterator struct {
	Event *ACLTraitTestPaused // Event containing the contract specifics and raw log

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
func (it *ACLTraitTestPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ACLTraitTestPaused)
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
		it.Event = new(ACLTraitTestPaused)
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
func (it *ACLTraitTestPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ACLTraitTestPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ACLTraitTestPaused represents a Paused event raised by the ACLTraitTest contract.
type ACLTraitTestPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ACLTraitTest *ACLTraitTestFilterer) FilterPaused(opts *bind.FilterOpts) (*ACLTraitTestPausedIterator, error) {

	logs, sub, err := _ACLTraitTest.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &ACLTraitTestPausedIterator{contract: _ACLTraitTest.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ACLTraitTest *ACLTraitTestFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ACLTraitTestPaused) (event.Subscription, error) {

	logs, sub, err := _ACLTraitTest.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ACLTraitTestPaused)
				if err := _ACLTraitTest.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_ACLTraitTest *ACLTraitTestFilterer) ParsePaused(log types.Log) (*ACLTraitTestPaused, error) {
	event := new(ACLTraitTestPaused)
	if err := _ACLTraitTest.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ACLTraitTestUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ACLTraitTest contract.
type ACLTraitTestUnpausedIterator struct {
	Event *ACLTraitTestUnpaused // Event containing the contract specifics and raw log

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
func (it *ACLTraitTestUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ACLTraitTestUnpaused)
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
		it.Event = new(ACLTraitTestUnpaused)
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
func (it *ACLTraitTestUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ACLTraitTestUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ACLTraitTestUnpaused represents a Unpaused event raised by the ACLTraitTest contract.
type ACLTraitTestUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ACLTraitTest *ACLTraitTestFilterer) FilterUnpaused(opts *bind.FilterOpts) (*ACLTraitTestUnpausedIterator, error) {

	logs, sub, err := _ACLTraitTest.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &ACLTraitTestUnpausedIterator{contract: _ACLTraitTest.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ACLTraitTest *ACLTraitTestFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ACLTraitTestUnpaused) (event.Subscription, error) {

	logs, sub, err := _ACLTraitTest.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ACLTraitTestUnpaused)
				if err := _ACLTraitTest.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_ACLTraitTest *ACLTraitTestFilterer) ParseUnpaused(log types.Log) (*ACLTraitTestUnpaused, error) {
	event := new(ACLTraitTestUnpaused)
	if err := _ACLTraitTest.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
