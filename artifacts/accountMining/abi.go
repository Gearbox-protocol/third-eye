// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package accountMining

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

// AccountMiningMetaData contains all meta data concerning the AccountMining contract.
var AccountMiningMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"},{\"internalType\":\"contractAddressProvider\",\"name\":\"addressProvider\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Claimed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"accountFactory\",\"outputs\":[{\"internalType\":\"contractAccountFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"amount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"isClaimed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"merkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AccountMiningABI is the input ABI used to generate the binding from.
// Deprecated: Use AccountMiningMetaData.ABI instead.
var AccountMiningABI = AccountMiningMetaData.ABI

// AccountMining is an auto generated Go binding around an Ethereum contract.
type AccountMining struct {
	AccountMiningCaller     // Read-only binding to the contract
	AccountMiningTransactor // Write-only binding to the contract
	AccountMiningFilterer   // Log filterer for contract events
}

// AccountMiningCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccountMiningCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountMiningTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccountMiningTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountMiningFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccountMiningFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountMiningSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccountMiningSession struct {
	Contract     *AccountMining    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AccountMiningCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccountMiningCallerSession struct {
	Contract *AccountMiningCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// AccountMiningTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccountMiningTransactorSession struct {
	Contract     *AccountMiningTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// AccountMiningRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccountMiningRaw struct {
	Contract *AccountMining // Generic contract binding to access the raw methods on
}

// AccountMiningCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccountMiningCallerRaw struct {
	Contract *AccountMiningCaller // Generic read-only contract binding to access the raw methods on
}

// AccountMiningTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccountMiningTransactorRaw struct {
	Contract *AccountMiningTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccountMining creates a new instance of AccountMining, bound to a specific deployed contract.
func NewAccountMining(address common.Address, backend bind.ContractBackend) (*AccountMining, error) {
	contract, err := bindAccountMining(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AccountMining{AccountMiningCaller: AccountMiningCaller{contract: contract}, AccountMiningTransactor: AccountMiningTransactor{contract: contract}, AccountMiningFilterer: AccountMiningFilterer{contract: contract}}, nil
}

// NewAccountMiningCaller creates a new read-only instance of AccountMining, bound to a specific deployed contract.
func NewAccountMiningCaller(address common.Address, caller bind.ContractCaller) (*AccountMiningCaller, error) {
	contract, err := bindAccountMining(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccountMiningCaller{contract: contract}, nil
}

// NewAccountMiningTransactor creates a new write-only instance of AccountMining, bound to a specific deployed contract.
func NewAccountMiningTransactor(address common.Address, transactor bind.ContractTransactor) (*AccountMiningTransactor, error) {
	contract, err := bindAccountMining(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccountMiningTransactor{contract: contract}, nil
}

// NewAccountMiningFilterer creates a new log filterer instance of AccountMining, bound to a specific deployed contract.
func NewAccountMiningFilterer(address common.Address, filterer bind.ContractFilterer) (*AccountMiningFilterer, error) {
	contract, err := bindAccountMining(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccountMiningFilterer{contract: contract}, nil
}

// bindAccountMining binds a generic wrapper to an already deployed contract.
func bindAccountMining(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AccountMiningABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccountMining *AccountMiningRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccountMining.Contract.AccountMiningCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccountMining *AccountMiningRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccountMining.Contract.AccountMiningTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccountMining *AccountMiningRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccountMining.Contract.AccountMiningTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccountMining *AccountMiningCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccountMining.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccountMining *AccountMiningTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccountMining.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccountMining *AccountMiningTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccountMining.Contract.contract.Transact(opts, method, params...)
}

// AccountFactory is a free data retrieval call binding the contract method 0x687cd9c1.
//
// Solidity: function accountFactory() view returns(address)
func (_AccountMining *AccountMiningCaller) AccountFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AccountMining.contract.Call(opts, &out, "accountFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AccountFactory is a free data retrieval call binding the contract method 0x687cd9c1.
//
// Solidity: function accountFactory() view returns(address)
func (_AccountMining *AccountMiningSession) AccountFactory() (common.Address, error) {
	return _AccountMining.Contract.AccountFactory(&_AccountMining.CallOpts)
}

// AccountFactory is a free data retrieval call binding the contract method 0x687cd9c1.
//
// Solidity: function accountFactory() view returns(address)
func (_AccountMining *AccountMiningCallerSession) AccountFactory() (common.Address, error) {
	return _AccountMining.Contract.AccountFactory(&_AccountMining.CallOpts)
}

// Amount is a free data retrieval call binding the contract method 0xaa8c217c.
//
// Solidity: function amount() view returns(uint256)
func (_AccountMining *AccountMiningCaller) Amount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AccountMining.contract.Call(opts, &out, "amount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Amount is a free data retrieval call binding the contract method 0xaa8c217c.
//
// Solidity: function amount() view returns(uint256)
func (_AccountMining *AccountMiningSession) Amount() (*big.Int, error) {
	return _AccountMining.Contract.Amount(&_AccountMining.CallOpts)
}

// Amount is a free data retrieval call binding the contract method 0xaa8c217c.
//
// Solidity: function amount() view returns(uint256)
func (_AccountMining *AccountMiningCallerSession) Amount() (*big.Int, error) {
	return _AccountMining.Contract.Amount(&_AccountMining.CallOpts)
}

// IsClaimed is a free data retrieval call binding the contract method 0x9e34070f.
//
// Solidity: function isClaimed(uint256 index) view returns(bool)
func (_AccountMining *AccountMiningCaller) IsClaimed(opts *bind.CallOpts, index *big.Int) (bool, error) {
	var out []interface{}
	err := _AccountMining.contract.Call(opts, &out, "isClaimed", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsClaimed is a free data retrieval call binding the contract method 0x9e34070f.
//
// Solidity: function isClaimed(uint256 index) view returns(bool)
func (_AccountMining *AccountMiningSession) IsClaimed(index *big.Int) (bool, error) {
	return _AccountMining.Contract.IsClaimed(&_AccountMining.CallOpts, index)
}

// IsClaimed is a free data retrieval call binding the contract method 0x9e34070f.
//
// Solidity: function isClaimed(uint256 index) view returns(bool)
func (_AccountMining *AccountMiningCallerSession) IsClaimed(index *big.Int) (bool, error) {
	return _AccountMining.Contract.IsClaimed(&_AccountMining.CallOpts, index)
}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_AccountMining *AccountMiningCaller) MerkleRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AccountMining.contract.Call(opts, &out, "merkleRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_AccountMining *AccountMiningSession) MerkleRoot() ([32]byte, error) {
	return _AccountMining.Contract.MerkleRoot(&_AccountMining.CallOpts)
}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_AccountMining *AccountMiningCallerSession) MerkleRoot() ([32]byte, error) {
	return _AccountMining.Contract.MerkleRoot(&_AccountMining.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_AccountMining *AccountMiningCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AccountMining.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_AccountMining *AccountMiningSession) Token() (common.Address, error) {
	return _AccountMining.Contract.Token(&_AccountMining.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_AccountMining *AccountMiningCallerSession) Token() (common.Address, error) {
	return _AccountMining.Contract.Token(&_AccountMining.CallOpts)
}

// Claim is a paid mutator transaction binding the contract method 0xae0b51df.
//
// Solidity: function claim(uint256 index, uint256 salt, bytes32[] merkleProof) returns()
func (_AccountMining *AccountMiningTransactor) Claim(opts *bind.TransactOpts, index *big.Int, salt *big.Int, merkleProof [][32]byte) (*types.Transaction, error) {
	return _AccountMining.contract.Transact(opts, "claim", index, salt, merkleProof)
}

// Claim is a paid mutator transaction binding the contract method 0xae0b51df.
//
// Solidity: function claim(uint256 index, uint256 salt, bytes32[] merkleProof) returns()
func (_AccountMining *AccountMiningSession) Claim(index *big.Int, salt *big.Int, merkleProof [][32]byte) (*types.Transaction, error) {
	return _AccountMining.Contract.Claim(&_AccountMining.TransactOpts, index, salt, merkleProof)
}

// Claim is a paid mutator transaction binding the contract method 0xae0b51df.
//
// Solidity: function claim(uint256 index, uint256 salt, bytes32[] merkleProof) returns()
func (_AccountMining *AccountMiningTransactorSession) Claim(index *big.Int, salt *big.Int, merkleProof [][32]byte) (*types.Transaction, error) {
	return _AccountMining.Contract.Claim(&_AccountMining.TransactOpts, index, salt, merkleProof)
}

// AccountMiningClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the AccountMining contract.
type AccountMiningClaimedIterator struct {
	Event *AccountMiningClaimed // Event containing the contract specifics and raw log

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
func (it *AccountMiningClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountMiningClaimed)
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
		it.Event = new(AccountMiningClaimed)
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
func (it *AccountMiningClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountMiningClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountMiningClaimed represents a Claimed event raised by the AccountMining contract.
type AccountMiningClaimed struct {
	Index   *big.Int
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0x6aa3eac93d079e5e100b1029be716caa33586c96aa4baac390669fb5c2a21212.
//
// Solidity: event Claimed(uint256 index, address account)
func (_AccountMining *AccountMiningFilterer) FilterClaimed(opts *bind.FilterOpts) (*AccountMiningClaimedIterator, error) {

	logs, sub, err := _AccountMining.contract.FilterLogs(opts, "Claimed")
	if err != nil {
		return nil, err
	}
	return &AccountMiningClaimedIterator{contract: _AccountMining.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0x6aa3eac93d079e5e100b1029be716caa33586c96aa4baac390669fb5c2a21212.
//
// Solidity: event Claimed(uint256 index, address account)
func (_AccountMining *AccountMiningFilterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *AccountMiningClaimed) (event.Subscription, error) {

	logs, sub, err := _AccountMining.contract.WatchLogs(opts, "Claimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountMiningClaimed)
				if err := _AccountMining.contract.UnpackLog(event, "Claimed", log); err != nil {
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
func (_AccountMining *AccountMiningFilterer) ParseClaimed(log types.Log) (*AccountMiningClaimed, error) {
	event := new(AccountMiningClaimed)
	if err := _AccountMining.contract.UnpackLog(event, "Claimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
