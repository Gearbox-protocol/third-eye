// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lobsterMinter

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

// LobsterMinterMetaData contains all meta data concerning the LobsterMinter contract.
var LobsterMinterMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintCount\",\"type\":\"uint256\"}],\"name\":\"Claim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"ClaimByCollection\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_mintCount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// LobsterMinterABI is the input ABI used to generate the binding from.
// Deprecated: Use LobsterMinterMetaData.ABI instead.
var LobsterMinterABI = LobsterMinterMetaData.ABI

// LobsterMinter is an auto generated Go binding around an Ethereum contract.
type LobsterMinter struct {
	LobsterMinterCaller     // Read-only binding to the contract
	LobsterMinterTransactor // Write-only binding to the contract
	LobsterMinterFilterer   // Log filterer for contract events
}

// LobsterMinterCaller is an auto generated read-only Go binding around an Ethereum contract.
type LobsterMinterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LobsterMinterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LobsterMinterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LobsterMinterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LobsterMinterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LobsterMinterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LobsterMinterSession struct {
	Contract     *LobsterMinter    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LobsterMinterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LobsterMinterCallerSession struct {
	Contract *LobsterMinterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// LobsterMinterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LobsterMinterTransactorSession struct {
	Contract     *LobsterMinterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// LobsterMinterRaw is an auto generated low-level Go binding around an Ethereum contract.
type LobsterMinterRaw struct {
	Contract *LobsterMinter // Generic contract binding to access the raw methods on
}

// LobsterMinterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LobsterMinterCallerRaw struct {
	Contract *LobsterMinterCaller // Generic read-only contract binding to access the raw methods on
}

// LobsterMinterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LobsterMinterTransactorRaw struct {
	Contract *LobsterMinterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLobsterMinter creates a new instance of LobsterMinter, bound to a specific deployed contract.
func NewLobsterMinter(address common.Address, backend bind.ContractBackend) (*LobsterMinter, error) {
	contract, err := bindLobsterMinter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LobsterMinter{LobsterMinterCaller: LobsterMinterCaller{contract: contract}, LobsterMinterTransactor: LobsterMinterTransactor{contract: contract}, LobsterMinterFilterer: LobsterMinterFilterer{contract: contract}}, nil
}

// NewLobsterMinterCaller creates a new read-only instance of LobsterMinter, bound to a specific deployed contract.
func NewLobsterMinterCaller(address common.Address, caller bind.ContractCaller) (*LobsterMinterCaller, error) {
	contract, err := bindLobsterMinter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LobsterMinterCaller{contract: contract}, nil
}

// NewLobsterMinterTransactor creates a new write-only instance of LobsterMinter, bound to a specific deployed contract.
func NewLobsterMinterTransactor(address common.Address, transactor bind.ContractTransactor) (*LobsterMinterTransactor, error) {
	contract, err := bindLobsterMinter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LobsterMinterTransactor{contract: contract}, nil
}

// NewLobsterMinterFilterer creates a new log filterer instance of LobsterMinter, bound to a specific deployed contract.
func NewLobsterMinterFilterer(address common.Address, filterer bind.ContractFilterer) (*LobsterMinterFilterer, error) {
	contract, err := bindLobsterMinter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LobsterMinterFilterer{contract: contract}, nil
}

// bindLobsterMinter binds a generic wrapper to an already deployed contract.
func bindLobsterMinter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LobsterMinterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LobsterMinter *LobsterMinterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LobsterMinter.Contract.LobsterMinterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LobsterMinter *LobsterMinterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LobsterMinter.Contract.LobsterMinterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LobsterMinter *LobsterMinterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LobsterMinter.Contract.LobsterMinterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LobsterMinter *LobsterMinterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LobsterMinter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LobsterMinter *LobsterMinterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LobsterMinter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LobsterMinter *LobsterMinterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LobsterMinter.Contract.contract.Transact(opts, method, params...)
}

// Claim is a paid mutator transaction binding the contract method 0x172bd6de.
//
// Solidity: function claim(address _account, uint256 _count, uint256 _mintCount, bytes32[] _merkleProof) returns()
func (_LobsterMinter *LobsterMinterTransactor) Claim(opts *bind.TransactOpts, _account common.Address, _count *big.Int, _mintCount *big.Int, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _LobsterMinter.contract.Transact(opts, "claim", _account, _count, _mintCount, _merkleProof)
}

// Claim is a paid mutator transaction binding the contract method 0x172bd6de.
//
// Solidity: function claim(address _account, uint256 _count, uint256 _mintCount, bytes32[] _merkleProof) returns()
func (_LobsterMinter *LobsterMinterSession) Claim(_account common.Address, _count *big.Int, _mintCount *big.Int, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _LobsterMinter.Contract.Claim(&_LobsterMinter.TransactOpts, _account, _count, _mintCount, _merkleProof)
}

// Claim is a paid mutator transaction binding the contract method 0x172bd6de.
//
// Solidity: function claim(address _account, uint256 _count, uint256 _mintCount, bytes32[] _merkleProof) returns()
func (_LobsterMinter *LobsterMinterTransactorSession) Claim(_account common.Address, _count *big.Int, _mintCount *big.Int, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _LobsterMinter.Contract.Claim(&_LobsterMinter.TransactOpts, _account, _count, _mintCount, _merkleProof)
}

// LobsterMinterClaimIterator is returned from FilterClaim and is used to iterate over the raw logs and unpacked data for Claim events raised by the LobsterMinter contract.
type LobsterMinterClaimIterator struct {
	Event *LobsterMinterClaim // Event containing the contract specifics and raw log

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
func (it *LobsterMinterClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LobsterMinterClaim)
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
		it.Event = new(LobsterMinterClaim)
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
func (it *LobsterMinterClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LobsterMinterClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LobsterMinterClaim represents a Claim event raised by the LobsterMinter contract.
type LobsterMinterClaim struct {
	Account   common.Address
	Count     *big.Int
	MintCount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterClaim is a free log retrieval operation binding the contract event 0x34fcbac0073d7c3d388e51312faf357774904998eeb8fca628b9e6f65ee1cbf7.
//
// Solidity: event Claim(address indexed account, uint256 count, uint256 mintCount)
func (_LobsterMinter *LobsterMinterFilterer) FilterClaim(opts *bind.FilterOpts, account []common.Address) (*LobsterMinterClaimIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LobsterMinter.contract.FilterLogs(opts, "Claim", accountRule)
	if err != nil {
		return nil, err
	}
	return &LobsterMinterClaimIterator{contract: _LobsterMinter.contract, event: "Claim", logs: logs, sub: sub}, nil
}

// WatchClaim is a free log subscription operation binding the contract event 0x34fcbac0073d7c3d388e51312faf357774904998eeb8fca628b9e6f65ee1cbf7.
//
// Solidity: event Claim(address indexed account, uint256 count, uint256 mintCount)
func (_LobsterMinter *LobsterMinterFilterer) WatchClaim(opts *bind.WatchOpts, sink chan<- *LobsterMinterClaim, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LobsterMinter.contract.WatchLogs(opts, "Claim", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LobsterMinterClaim)
				if err := _LobsterMinter.contract.UnpackLog(event, "Claim", log); err != nil {
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

// ParseClaim is a log parse operation binding the contract event 0x34fcbac0073d7c3d388e51312faf357774904998eeb8fca628b9e6f65ee1cbf7.
//
// Solidity: event Claim(address indexed account, uint256 count, uint256 mintCount)
func (_LobsterMinter *LobsterMinterFilterer) ParseClaim(log types.Log) (*LobsterMinterClaim, error) {
	event := new(LobsterMinterClaim)
	if err := _LobsterMinter.contract.UnpackLog(event, "Claim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LobsterMinterClaimByCollectionIterator is returned from FilterClaimByCollection and is used to iterate over the raw logs and unpacked data for ClaimByCollection events raised by the LobsterMinter contract.
type LobsterMinterClaimByCollectionIterator struct {
	Event *LobsterMinterClaimByCollection // Event containing the contract specifics and raw log

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
func (it *LobsterMinterClaimByCollectionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LobsterMinterClaimByCollection)
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
		it.Event = new(LobsterMinterClaimByCollection)
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
func (it *LobsterMinterClaimByCollectionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LobsterMinterClaimByCollectionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LobsterMinterClaimByCollection represents a ClaimByCollection event raised by the LobsterMinter contract.
type LobsterMinterClaimByCollection struct {
	Account    common.Address
	Collection common.Address
	TokenIds   []*big.Int
	Count      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterClaimByCollection is a free log retrieval operation binding the contract event 0x83703a9d003db5fa033673642c571ad485bdd3bd0f9782d7769756ed56d2f611.
//
// Solidity: event ClaimByCollection(address indexed account, address indexed collection, uint256[] tokenIds, uint256 count)
func (_LobsterMinter *LobsterMinterFilterer) FilterClaimByCollection(opts *bind.FilterOpts, account []common.Address, collection []common.Address) (*LobsterMinterClaimByCollectionIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var collectionRule []interface{}
	for _, collectionItem := range collection {
		collectionRule = append(collectionRule, collectionItem)
	}

	logs, sub, err := _LobsterMinter.contract.FilterLogs(opts, "ClaimByCollection", accountRule, collectionRule)
	if err != nil {
		return nil, err
	}
	return &LobsterMinterClaimByCollectionIterator{contract: _LobsterMinter.contract, event: "ClaimByCollection", logs: logs, sub: sub}, nil
}

// WatchClaimByCollection is a free log subscription operation binding the contract event 0x83703a9d003db5fa033673642c571ad485bdd3bd0f9782d7769756ed56d2f611.
//
// Solidity: event ClaimByCollection(address indexed account, address indexed collection, uint256[] tokenIds, uint256 count)
func (_LobsterMinter *LobsterMinterFilterer) WatchClaimByCollection(opts *bind.WatchOpts, sink chan<- *LobsterMinterClaimByCollection, account []common.Address, collection []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var collectionRule []interface{}
	for _, collectionItem := range collection {
		collectionRule = append(collectionRule, collectionItem)
	}

	logs, sub, err := _LobsterMinter.contract.WatchLogs(opts, "ClaimByCollection", accountRule, collectionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LobsterMinterClaimByCollection)
				if err := _LobsterMinter.contract.UnpackLog(event, "ClaimByCollection", log); err != nil {
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

// ParseClaimByCollection is a log parse operation binding the contract event 0x83703a9d003db5fa033673642c571ad485bdd3bd0f9782d7769756ed56d2f611.
//
// Solidity: event ClaimByCollection(address indexed account, address indexed collection, uint256[] tokenIds, uint256 count)
func (_LobsterMinter *LobsterMinterFilterer) ParseClaimByCollection(log types.Log) (*LobsterMinterClaimByCollection, error) {
	event := new(LobsterMinterClaimByCollection)
	if err := _LobsterMinter.contract.UnpackLog(event, "ClaimByCollection", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
