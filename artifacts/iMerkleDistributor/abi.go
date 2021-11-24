// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iMerkleDistributor

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

// IMerkleDistributorMetaData contains all meta data concerning the IMerkleDistributor contract.
var IMerkleDistributorMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Claimed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"isClaimed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"merkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IMerkleDistributorABI is the input ABI used to generate the binding from.
// Deprecated: Use IMerkleDistributorMetaData.ABI instead.
var IMerkleDistributorABI = IMerkleDistributorMetaData.ABI

// IMerkleDistributor is an auto generated Go binding around an Ethereum contract.
type IMerkleDistributor struct {
	IMerkleDistributorCaller     // Read-only binding to the contract
	IMerkleDistributorTransactor // Write-only binding to the contract
	IMerkleDistributorFilterer   // Log filterer for contract events
}

// IMerkleDistributorCaller is an auto generated read-only Go binding around an Ethereum contract.
type IMerkleDistributorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMerkleDistributorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IMerkleDistributorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMerkleDistributorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IMerkleDistributorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMerkleDistributorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IMerkleDistributorSession struct {
	Contract     *IMerkleDistributor // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IMerkleDistributorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IMerkleDistributorCallerSession struct {
	Contract *IMerkleDistributorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// IMerkleDistributorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IMerkleDistributorTransactorSession struct {
	Contract     *IMerkleDistributorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// IMerkleDistributorRaw is an auto generated low-level Go binding around an Ethereum contract.
type IMerkleDistributorRaw struct {
	Contract *IMerkleDistributor // Generic contract binding to access the raw methods on
}

// IMerkleDistributorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IMerkleDistributorCallerRaw struct {
	Contract *IMerkleDistributorCaller // Generic read-only contract binding to access the raw methods on
}

// IMerkleDistributorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IMerkleDistributorTransactorRaw struct {
	Contract *IMerkleDistributorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIMerkleDistributor creates a new instance of IMerkleDistributor, bound to a specific deployed contract.
func NewIMerkleDistributor(address common.Address, backend bind.ContractBackend) (*IMerkleDistributor, error) {
	contract, err := bindIMerkleDistributor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IMerkleDistributor{IMerkleDistributorCaller: IMerkleDistributorCaller{contract: contract}, IMerkleDistributorTransactor: IMerkleDistributorTransactor{contract: contract}, IMerkleDistributorFilterer: IMerkleDistributorFilterer{contract: contract}}, nil
}

// NewIMerkleDistributorCaller creates a new read-only instance of IMerkleDistributor, bound to a specific deployed contract.
func NewIMerkleDistributorCaller(address common.Address, caller bind.ContractCaller) (*IMerkleDistributorCaller, error) {
	contract, err := bindIMerkleDistributor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IMerkleDistributorCaller{contract: contract}, nil
}

// NewIMerkleDistributorTransactor creates a new write-only instance of IMerkleDistributor, bound to a specific deployed contract.
func NewIMerkleDistributorTransactor(address common.Address, transactor bind.ContractTransactor) (*IMerkleDistributorTransactor, error) {
	contract, err := bindIMerkleDistributor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IMerkleDistributorTransactor{contract: contract}, nil
}

// NewIMerkleDistributorFilterer creates a new log filterer instance of IMerkleDistributor, bound to a specific deployed contract.
func NewIMerkleDistributorFilterer(address common.Address, filterer bind.ContractFilterer) (*IMerkleDistributorFilterer, error) {
	contract, err := bindIMerkleDistributor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IMerkleDistributorFilterer{contract: contract}, nil
}

// bindIMerkleDistributor binds a generic wrapper to an already deployed contract.
func bindIMerkleDistributor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IMerkleDistributorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMerkleDistributor *IMerkleDistributorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMerkleDistributor.Contract.IMerkleDistributorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMerkleDistributor *IMerkleDistributorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMerkleDistributor.Contract.IMerkleDistributorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMerkleDistributor *IMerkleDistributorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMerkleDistributor.Contract.IMerkleDistributorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMerkleDistributor *IMerkleDistributorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMerkleDistributor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMerkleDistributor *IMerkleDistributorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMerkleDistributor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMerkleDistributor *IMerkleDistributorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMerkleDistributor.Contract.contract.Transact(opts, method, params...)
}

// IsClaimed is a free data retrieval call binding the contract method 0x9e34070f.
//
// Solidity: function isClaimed(uint256 index) view returns(bool)
func (_IMerkleDistributor *IMerkleDistributorCaller) IsClaimed(opts *bind.CallOpts, index *big.Int) (bool, error) {
	var out []interface{}
	err := _IMerkleDistributor.contract.Call(opts, &out, "isClaimed", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsClaimed is a free data retrieval call binding the contract method 0x9e34070f.
//
// Solidity: function isClaimed(uint256 index) view returns(bool)
func (_IMerkleDistributor *IMerkleDistributorSession) IsClaimed(index *big.Int) (bool, error) {
	return _IMerkleDistributor.Contract.IsClaimed(&_IMerkleDistributor.CallOpts, index)
}

// IsClaimed is a free data retrieval call binding the contract method 0x9e34070f.
//
// Solidity: function isClaimed(uint256 index) view returns(bool)
func (_IMerkleDistributor *IMerkleDistributorCallerSession) IsClaimed(index *big.Int) (bool, error) {
	return _IMerkleDistributor.Contract.IsClaimed(&_IMerkleDistributor.CallOpts, index)
}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_IMerkleDistributor *IMerkleDistributorCaller) MerkleRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IMerkleDistributor.contract.Call(opts, &out, "merkleRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_IMerkleDistributor *IMerkleDistributorSession) MerkleRoot() ([32]byte, error) {
	return _IMerkleDistributor.Contract.MerkleRoot(&_IMerkleDistributor.CallOpts)
}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_IMerkleDistributor *IMerkleDistributorCallerSession) MerkleRoot() ([32]byte, error) {
	return _IMerkleDistributor.Contract.MerkleRoot(&_IMerkleDistributor.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_IMerkleDistributor *IMerkleDistributorCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IMerkleDistributor.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_IMerkleDistributor *IMerkleDistributorSession) Token() (common.Address, error) {
	return _IMerkleDistributor.Contract.Token(&_IMerkleDistributor.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_IMerkleDistributor *IMerkleDistributorCallerSession) Token() (common.Address, error) {
	return _IMerkleDistributor.Contract.Token(&_IMerkleDistributor.CallOpts)
}

// Claim is a paid mutator transaction binding the contract method 0xae0b51df.
//
// Solidity: function claim(uint256 index, uint256 salt, bytes32[] merkleProof) returns()
func (_IMerkleDistributor *IMerkleDistributorTransactor) Claim(opts *bind.TransactOpts, index *big.Int, salt *big.Int, merkleProof [][32]byte) (*types.Transaction, error) {
	return _IMerkleDistributor.contract.Transact(opts, "claim", index, salt, merkleProof)
}

// Claim is a paid mutator transaction binding the contract method 0xae0b51df.
//
// Solidity: function claim(uint256 index, uint256 salt, bytes32[] merkleProof) returns()
func (_IMerkleDistributor *IMerkleDistributorSession) Claim(index *big.Int, salt *big.Int, merkleProof [][32]byte) (*types.Transaction, error) {
	return _IMerkleDistributor.Contract.Claim(&_IMerkleDistributor.TransactOpts, index, salt, merkleProof)
}

// Claim is a paid mutator transaction binding the contract method 0xae0b51df.
//
// Solidity: function claim(uint256 index, uint256 salt, bytes32[] merkleProof) returns()
func (_IMerkleDistributor *IMerkleDistributorTransactorSession) Claim(index *big.Int, salt *big.Int, merkleProof [][32]byte) (*types.Transaction, error) {
	return _IMerkleDistributor.Contract.Claim(&_IMerkleDistributor.TransactOpts, index, salt, merkleProof)
}

// IMerkleDistributorClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the IMerkleDistributor contract.
type IMerkleDistributorClaimedIterator struct {
	Event *IMerkleDistributorClaimed // Event containing the contract specifics and raw log

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
func (it *IMerkleDistributorClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IMerkleDistributorClaimed)
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
		it.Event = new(IMerkleDistributorClaimed)
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
func (it *IMerkleDistributorClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IMerkleDistributorClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IMerkleDistributorClaimed represents a Claimed event raised by the IMerkleDistributor contract.
type IMerkleDistributorClaimed struct {
	Index   *big.Int
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0x6aa3eac93d079e5e100b1029be716caa33586c96aa4baac390669fb5c2a21212.
//
// Solidity: event Claimed(uint256 index, address account)
func (_IMerkleDistributor *IMerkleDistributorFilterer) FilterClaimed(opts *bind.FilterOpts) (*IMerkleDistributorClaimedIterator, error) {

	logs, sub, err := _IMerkleDistributor.contract.FilterLogs(opts, "Claimed")
	if err != nil {
		return nil, err
	}
	return &IMerkleDistributorClaimedIterator{contract: _IMerkleDistributor.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0x6aa3eac93d079e5e100b1029be716caa33586c96aa4baac390669fb5c2a21212.
//
// Solidity: event Claimed(uint256 index, address account)
func (_IMerkleDistributor *IMerkleDistributorFilterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *IMerkleDistributorClaimed) (event.Subscription, error) {

	logs, sub, err := _IMerkleDistributor.contract.WatchLogs(opts, "Claimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IMerkleDistributorClaimed)
				if err := _IMerkleDistributor.contract.UnpackLog(event, "Claimed", log); err != nil {
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

// ParseClaimed is a log parse operation binding the contract event 0x6aa3eac93d079e5e100b1029be716caa33586c96aa4baac390669fb5c2a21212.
//
// Solidity: event Claimed(uint256 index, address account)
func (_IMerkleDistributor *IMerkleDistributorFilterer) ParseClaimed(log types.Log) (*IMerkleDistributorClaimed, error) {
	event := new(IMerkleDistributorClaimed)
	if err := _IMerkleDistributor.contract.UnpackLog(event, "Claimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
