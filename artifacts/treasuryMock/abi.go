// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package treasuryMock

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

// TreasuryMockMetaData contains all meta data concerning the TreasuryMock contract.
var TreasuryMockMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"NewDonation\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// TreasuryMockABI is the input ABI used to generate the binding from.
// Deprecated: Use TreasuryMockMetaData.ABI instead.
var TreasuryMockABI = TreasuryMockMetaData.ABI

// TreasuryMock is an auto generated Go binding around an Ethereum contract.
type TreasuryMock struct {
	TreasuryMockCaller     // Read-only binding to the contract
	TreasuryMockTransactor // Write-only binding to the contract
	TreasuryMockFilterer   // Log filterer for contract events
}

// TreasuryMockCaller is an auto generated read-only Go binding around an Ethereum contract.
type TreasuryMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TreasuryMockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TreasuryMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TreasuryMockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TreasuryMockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TreasuryMockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TreasuryMockSession struct {
	Contract     *TreasuryMock     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TreasuryMockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TreasuryMockCallerSession struct {
	Contract *TreasuryMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// TreasuryMockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TreasuryMockTransactorSession struct {
	Contract     *TreasuryMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TreasuryMockRaw is an auto generated low-level Go binding around an Ethereum contract.
type TreasuryMockRaw struct {
	Contract *TreasuryMock // Generic contract binding to access the raw methods on
}

// TreasuryMockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TreasuryMockCallerRaw struct {
	Contract *TreasuryMockCaller // Generic read-only contract binding to access the raw methods on
}

// TreasuryMockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TreasuryMockTransactorRaw struct {
	Contract *TreasuryMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTreasuryMock creates a new instance of TreasuryMock, bound to a specific deployed contract.
func NewTreasuryMock(address common.Address, backend bind.ContractBackend) (*TreasuryMock, error) {
	contract, err := bindTreasuryMock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TreasuryMock{TreasuryMockCaller: TreasuryMockCaller{contract: contract}, TreasuryMockTransactor: TreasuryMockTransactor{contract: contract}, TreasuryMockFilterer: TreasuryMockFilterer{contract: contract}}, nil
}

// NewTreasuryMockCaller creates a new read-only instance of TreasuryMock, bound to a specific deployed contract.
func NewTreasuryMockCaller(address common.Address, caller bind.ContractCaller) (*TreasuryMockCaller, error) {
	contract, err := bindTreasuryMock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TreasuryMockCaller{contract: contract}, nil
}

// NewTreasuryMockTransactor creates a new write-only instance of TreasuryMock, bound to a specific deployed contract.
func NewTreasuryMockTransactor(address common.Address, transactor bind.ContractTransactor) (*TreasuryMockTransactor, error) {
	contract, err := bindTreasuryMock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TreasuryMockTransactor{contract: contract}, nil
}

// NewTreasuryMockFilterer creates a new log filterer instance of TreasuryMock, bound to a specific deployed contract.
func NewTreasuryMockFilterer(address common.Address, filterer bind.ContractFilterer) (*TreasuryMockFilterer, error) {
	contract, err := bindTreasuryMock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TreasuryMockFilterer{contract: contract}, nil
}

// bindTreasuryMock binds a generic wrapper to an already deployed contract.
func bindTreasuryMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TreasuryMockABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TreasuryMock *TreasuryMockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TreasuryMock.Contract.TreasuryMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TreasuryMock *TreasuryMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TreasuryMock.Contract.TreasuryMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TreasuryMock *TreasuryMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TreasuryMock.Contract.TreasuryMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TreasuryMock *TreasuryMockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TreasuryMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TreasuryMock *TreasuryMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TreasuryMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TreasuryMock *TreasuryMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TreasuryMock.Contract.contract.Transact(opts, method, params...)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TreasuryMock *TreasuryMockTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TreasuryMock.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TreasuryMock *TreasuryMockSession) Receive() (*types.Transaction, error) {
	return _TreasuryMock.Contract.Receive(&_TreasuryMock.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TreasuryMock *TreasuryMockTransactorSession) Receive() (*types.Transaction, error) {
	return _TreasuryMock.Contract.Receive(&_TreasuryMock.TransactOpts)
}

// TreasuryMockNewDonationIterator is returned from FilterNewDonation and is used to iterate over the raw logs and unpacked data for NewDonation events raised by the TreasuryMock contract.
type TreasuryMockNewDonationIterator struct {
	Event *TreasuryMockNewDonation // Event containing the contract specifics and raw log

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
func (it *TreasuryMockNewDonationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TreasuryMockNewDonation)
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
		it.Event = new(TreasuryMockNewDonation)
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
func (it *TreasuryMockNewDonationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TreasuryMockNewDonationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TreasuryMockNewDonation represents a NewDonation event raised by the TreasuryMock contract.
type TreasuryMockNewDonation struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNewDonation is a free log retrieval operation binding the contract event 0x8ffa785350fa6b5fee858c4ca63eff2704b9538ff446bd673c1f6c11fc7aca16.
//
// Solidity: event NewDonation(uint256 amount)
func (_TreasuryMock *TreasuryMockFilterer) FilterNewDonation(opts *bind.FilterOpts) (*TreasuryMockNewDonationIterator, error) {

	logs, sub, err := _TreasuryMock.contract.FilterLogs(opts, "NewDonation")
	if err != nil {
		return nil, err
	}
	return &TreasuryMockNewDonationIterator{contract: _TreasuryMock.contract, event: "NewDonation", logs: logs, sub: sub}, nil
}

// WatchNewDonation is a free log subscription operation binding the contract event 0x8ffa785350fa6b5fee858c4ca63eff2704b9538ff446bd673c1f6c11fc7aca16.
//
// Solidity: event NewDonation(uint256 amount)
func (_TreasuryMock *TreasuryMockFilterer) WatchNewDonation(opts *bind.WatchOpts, sink chan<- *TreasuryMockNewDonation) (event.Subscription, error) {

	logs, sub, err := _TreasuryMock.contract.WatchLogs(opts, "NewDonation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TreasuryMockNewDonation)
				if err := _TreasuryMock.contract.UnpackLog(event, "NewDonation", log); err != nil {
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

// ParseNewDonation is a log parse operation binding the contract event 0x8ffa785350fa6b5fee858c4ca63eff2704b9538ff446bd673c1f6c11fc7aca16.
//
// Solidity: event NewDonation(uint256 amount)
func (_TreasuryMock *TreasuryMockFilterer) ParseNewDonation(log types.Log) (*TreasuryMockNewDonation, error) {
	event := new(TreasuryMockNewDonation)
	if err := _TreasuryMock.contract.UnpackLog(event, "NewDonation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
