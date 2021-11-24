// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package aCLTrait

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

// ACLTraitMetaData contains all meta data concerning the ACLTrait contract.
var ACLTraitMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ACLTraitABI is the input ABI used to generate the binding from.
// Deprecated: Use ACLTraitMetaData.ABI instead.
var ACLTraitABI = ACLTraitMetaData.ABI

// ACLTrait is an auto generated Go binding around an Ethereum contract.
type ACLTrait struct {
	ACLTraitCaller     // Read-only binding to the contract
	ACLTraitTransactor // Write-only binding to the contract
	ACLTraitFilterer   // Log filterer for contract events
}

// ACLTraitCaller is an auto generated read-only Go binding around an Ethereum contract.
type ACLTraitCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ACLTraitTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ACLTraitTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ACLTraitFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ACLTraitFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ACLTraitSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ACLTraitSession struct {
	Contract     *ACLTrait         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ACLTraitCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ACLTraitCallerSession struct {
	Contract *ACLTraitCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ACLTraitTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ACLTraitTransactorSession struct {
	Contract     *ACLTraitTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ACLTraitRaw is an auto generated low-level Go binding around an Ethereum contract.
type ACLTraitRaw struct {
	Contract *ACLTrait // Generic contract binding to access the raw methods on
}

// ACLTraitCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ACLTraitCallerRaw struct {
	Contract *ACLTraitCaller // Generic read-only contract binding to access the raw methods on
}

// ACLTraitTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ACLTraitTransactorRaw struct {
	Contract *ACLTraitTransactor // Generic write-only contract binding to access the raw methods on
}

// NewACLTrait creates a new instance of ACLTrait, bound to a specific deployed contract.
func NewACLTrait(address common.Address, backend bind.ContractBackend) (*ACLTrait, error) {
	contract, err := bindACLTrait(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ACLTrait{ACLTraitCaller: ACLTraitCaller{contract: contract}, ACLTraitTransactor: ACLTraitTransactor{contract: contract}, ACLTraitFilterer: ACLTraitFilterer{contract: contract}}, nil
}

// NewACLTraitCaller creates a new read-only instance of ACLTrait, bound to a specific deployed contract.
func NewACLTraitCaller(address common.Address, caller bind.ContractCaller) (*ACLTraitCaller, error) {
	contract, err := bindACLTrait(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ACLTraitCaller{contract: contract}, nil
}

// NewACLTraitTransactor creates a new write-only instance of ACLTrait, bound to a specific deployed contract.
func NewACLTraitTransactor(address common.Address, transactor bind.ContractTransactor) (*ACLTraitTransactor, error) {
	contract, err := bindACLTrait(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ACLTraitTransactor{contract: contract}, nil
}

// NewACLTraitFilterer creates a new log filterer instance of ACLTrait, bound to a specific deployed contract.
func NewACLTraitFilterer(address common.Address, filterer bind.ContractFilterer) (*ACLTraitFilterer, error) {
	contract, err := bindACLTrait(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ACLTraitFilterer{contract: contract}, nil
}

// bindACLTrait binds a generic wrapper to an already deployed contract.
func bindACLTrait(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ACLTraitABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ACLTrait *ACLTraitRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ACLTrait.Contract.ACLTraitCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ACLTrait *ACLTraitRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ACLTrait.Contract.ACLTraitTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ACLTrait *ACLTraitRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ACLTrait.Contract.ACLTraitTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ACLTrait *ACLTraitCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ACLTrait.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ACLTrait *ACLTraitTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ACLTrait.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ACLTrait *ACLTraitTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ACLTrait.Contract.contract.Transact(opts, method, params...)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ACLTrait *ACLTraitCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ACLTrait.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ACLTrait *ACLTraitSession) Paused() (bool, error) {
	return _ACLTrait.Contract.Paused(&_ACLTrait.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ACLTrait *ACLTraitCallerSession) Paused() (bool, error) {
	return _ACLTrait.Contract.Paused(&_ACLTrait.CallOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ACLTrait *ACLTraitTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ACLTrait.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ACLTrait *ACLTraitSession) Pause() (*types.Transaction, error) {
	return _ACLTrait.Contract.Pause(&_ACLTrait.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ACLTrait *ACLTraitTransactorSession) Pause() (*types.Transaction, error) {
	return _ACLTrait.Contract.Pause(&_ACLTrait.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ACLTrait *ACLTraitTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ACLTrait.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ACLTrait *ACLTraitSession) Unpause() (*types.Transaction, error) {
	return _ACLTrait.Contract.Unpause(&_ACLTrait.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ACLTrait *ACLTraitTransactorSession) Unpause() (*types.Transaction, error) {
	return _ACLTrait.Contract.Unpause(&_ACLTrait.TransactOpts)
}

// ACLTraitPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ACLTrait contract.
type ACLTraitPausedIterator struct {
	Event *ACLTraitPaused // Event containing the contract specifics and raw log

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
func (it *ACLTraitPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ACLTraitPaused)
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
		it.Event = new(ACLTraitPaused)
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
func (it *ACLTraitPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ACLTraitPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ACLTraitPaused represents a Paused event raised by the ACLTrait contract.
type ACLTraitPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ACLTrait *ACLTraitFilterer) FilterPaused(opts *bind.FilterOpts) (*ACLTraitPausedIterator, error) {

	logs, sub, err := _ACLTrait.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &ACLTraitPausedIterator{contract: _ACLTrait.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ACLTrait *ACLTraitFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ACLTraitPaused) (event.Subscription, error) {

	logs, sub, err := _ACLTrait.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ACLTraitPaused)
				if err := _ACLTrait.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_ACLTrait *ACLTraitFilterer) ParsePaused(log types.Log) (*ACLTraitPaused, error) {
	event := new(ACLTraitPaused)
	if err := _ACLTrait.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ACLTraitUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ACLTrait contract.
type ACLTraitUnpausedIterator struct {
	Event *ACLTraitUnpaused // Event containing the contract specifics and raw log

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
func (it *ACLTraitUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ACLTraitUnpaused)
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
		it.Event = new(ACLTraitUnpaused)
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
func (it *ACLTraitUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ACLTraitUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ACLTraitUnpaused represents a Unpaused event raised by the ACLTrait contract.
type ACLTraitUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ACLTrait *ACLTraitFilterer) FilterUnpaused(opts *bind.FilterOpts) (*ACLTraitUnpausedIterator, error) {

	logs, sub, err := _ACLTrait.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &ACLTraitUnpausedIterator{contract: _ACLTrait.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ACLTrait *ACLTraitFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ACLTraitUnpaused) (event.Subscription, error) {

	logs, sub, err := _ACLTrait.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ACLTraitUnpaused)
				if err := _ACLTrait.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_ACLTrait *ACLTraitFilterer) ParseUnpaused(log types.Log) (*ACLTraitUnpaused, error) {
	event := new(ACLTraitUnpaused)
	if err := _ACLTrait.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
