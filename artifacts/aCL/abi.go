// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package aCL

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

// ACLMetaData contains all meta data concerning the ACL contract.
var ACLMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"PausableAdminAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"PausableAdminRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"UnpausableAdminAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"UnpausableAdminRemoved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"addPausableAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"addUnpausableAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isConfigurator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isPausableAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isUnpausableAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pausableAdminSet\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"removePausableAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"removeUnpausableAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"unpausableAdminSet\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ACLABI is the input ABI used to generate the binding from.
// Deprecated: Use ACLMetaData.ABI instead.
var ACLABI = ACLMetaData.ABI

// ACL is an auto generated Go binding around an Ethereum contract.
type ACL struct {
	ACLCaller     // Read-only binding to the contract
	ACLTransactor // Write-only binding to the contract
	ACLFilterer   // Log filterer for contract events
}

// ACLCaller is an auto generated read-only Go binding around an Ethereum contract.
type ACLCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ACLTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ACLTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ACLFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ACLFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ACLSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ACLSession struct {
	Contract     *ACL              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ACLCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ACLCallerSession struct {
	Contract *ACLCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ACLTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ACLTransactorSession struct {
	Contract     *ACLTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ACLRaw is an auto generated low-level Go binding around an Ethereum contract.
type ACLRaw struct {
	Contract *ACL // Generic contract binding to access the raw methods on
}

// ACLCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ACLCallerRaw struct {
	Contract *ACLCaller // Generic read-only contract binding to access the raw methods on
}

// ACLTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ACLTransactorRaw struct {
	Contract *ACLTransactor // Generic write-only contract binding to access the raw methods on
}

// NewACL creates a new instance of ACL, bound to a specific deployed contract.
func NewACL(address common.Address, backend bind.ContractBackend) (*ACL, error) {
	contract, err := bindACL(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ACL{ACLCaller: ACLCaller{contract: contract}, ACLTransactor: ACLTransactor{contract: contract}, ACLFilterer: ACLFilterer{contract: contract}}, nil
}

// NewACLCaller creates a new read-only instance of ACL, bound to a specific deployed contract.
func NewACLCaller(address common.Address, caller bind.ContractCaller) (*ACLCaller, error) {
	contract, err := bindACL(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ACLCaller{contract: contract}, nil
}

// NewACLTransactor creates a new write-only instance of ACL, bound to a specific deployed contract.
func NewACLTransactor(address common.Address, transactor bind.ContractTransactor) (*ACLTransactor, error) {
	contract, err := bindACL(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ACLTransactor{contract: contract}, nil
}

// NewACLFilterer creates a new log filterer instance of ACL, bound to a specific deployed contract.
func NewACLFilterer(address common.Address, filterer bind.ContractFilterer) (*ACLFilterer, error) {
	contract, err := bindACL(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ACLFilterer{contract: contract}, nil
}

// bindACL binds a generic wrapper to an already deployed contract.
func bindACL(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ACLABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ACL *ACLRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ACL.Contract.ACLCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ACL *ACLRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ACL.Contract.ACLTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ACL *ACLRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ACL.Contract.ACLTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ACL *ACLCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ACL.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ACL *ACLTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ACL.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ACL *ACLTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ACL.Contract.contract.Transact(opts, method, params...)
}

// IsConfigurator is a free data retrieval call binding the contract method 0x5f259aba.
//
// Solidity: function isConfigurator(address account) view returns(bool)
func (_ACL *ACLCaller) IsConfigurator(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _ACL.contract.Call(opts, &out, "isConfigurator", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsConfigurator is a free data retrieval call binding the contract method 0x5f259aba.
//
// Solidity: function isConfigurator(address account) view returns(bool)
func (_ACL *ACLSession) IsConfigurator(account common.Address) (bool, error) {
	return _ACL.Contract.IsConfigurator(&_ACL.CallOpts, account)
}

// IsConfigurator is a free data retrieval call binding the contract method 0x5f259aba.
//
// Solidity: function isConfigurator(address account) view returns(bool)
func (_ACL *ACLCallerSession) IsConfigurator(account common.Address) (bool, error) {
	return _ACL.Contract.IsConfigurator(&_ACL.CallOpts, account)
}

// IsPausableAdmin is a free data retrieval call binding the contract method 0x3a41ec64.
//
// Solidity: function isPausableAdmin(address addr) view returns(bool)
func (_ACL *ACLCaller) IsPausableAdmin(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _ACL.contract.Call(opts, &out, "isPausableAdmin", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPausableAdmin is a free data retrieval call binding the contract method 0x3a41ec64.
//
// Solidity: function isPausableAdmin(address addr) view returns(bool)
func (_ACL *ACLSession) IsPausableAdmin(addr common.Address) (bool, error) {
	return _ACL.Contract.IsPausableAdmin(&_ACL.CallOpts, addr)
}

// IsPausableAdmin is a free data retrieval call binding the contract method 0x3a41ec64.
//
// Solidity: function isPausableAdmin(address addr) view returns(bool)
func (_ACL *ACLCallerSession) IsPausableAdmin(addr common.Address) (bool, error) {
	return _ACL.Contract.IsPausableAdmin(&_ACL.CallOpts, addr)
}

// IsUnpausableAdmin is a free data retrieval call binding the contract method 0xd4eb5db0.
//
// Solidity: function isUnpausableAdmin(address addr) view returns(bool)
func (_ACL *ACLCaller) IsUnpausableAdmin(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _ACL.contract.Call(opts, &out, "isUnpausableAdmin", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsUnpausableAdmin is a free data retrieval call binding the contract method 0xd4eb5db0.
//
// Solidity: function isUnpausableAdmin(address addr) view returns(bool)
func (_ACL *ACLSession) IsUnpausableAdmin(addr common.Address) (bool, error) {
	return _ACL.Contract.IsUnpausableAdmin(&_ACL.CallOpts, addr)
}

// IsUnpausableAdmin is a free data retrieval call binding the contract method 0xd4eb5db0.
//
// Solidity: function isUnpausableAdmin(address addr) view returns(bool)
func (_ACL *ACLCallerSession) IsUnpausableAdmin(addr common.Address) (bool, error) {
	return _ACL.Contract.IsUnpausableAdmin(&_ACL.CallOpts, addr)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ACL *ACLCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ACL.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ACL *ACLSession) Owner() (common.Address, error) {
	return _ACL.Contract.Owner(&_ACL.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ACL *ACLCallerSession) Owner() (common.Address, error) {
	return _ACL.Contract.Owner(&_ACL.CallOpts)
}

// PausableAdminSet is a free data retrieval call binding the contract method 0x35914829.
//
// Solidity: function pausableAdminSet(address ) view returns(bool)
func (_ACL *ACLCaller) PausableAdminSet(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ACL.contract.Call(opts, &out, "pausableAdminSet", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PausableAdminSet is a free data retrieval call binding the contract method 0x35914829.
//
// Solidity: function pausableAdminSet(address ) view returns(bool)
func (_ACL *ACLSession) PausableAdminSet(arg0 common.Address) (bool, error) {
	return _ACL.Contract.PausableAdminSet(&_ACL.CallOpts, arg0)
}

// PausableAdminSet is a free data retrieval call binding the contract method 0x35914829.
//
// Solidity: function pausableAdminSet(address ) view returns(bool)
func (_ACL *ACLCallerSession) PausableAdminSet(arg0 common.Address) (bool, error) {
	return _ACL.Contract.PausableAdminSet(&_ACL.CallOpts, arg0)
}

// UnpausableAdminSet is a free data retrieval call binding the contract method 0x73281819.
//
// Solidity: function unpausableAdminSet(address ) view returns(bool)
func (_ACL *ACLCaller) UnpausableAdminSet(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ACL.contract.Call(opts, &out, "unpausableAdminSet", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UnpausableAdminSet is a free data retrieval call binding the contract method 0x73281819.
//
// Solidity: function unpausableAdminSet(address ) view returns(bool)
func (_ACL *ACLSession) UnpausableAdminSet(arg0 common.Address) (bool, error) {
	return _ACL.Contract.UnpausableAdminSet(&_ACL.CallOpts, arg0)
}

// UnpausableAdminSet is a free data retrieval call binding the contract method 0x73281819.
//
// Solidity: function unpausableAdminSet(address ) view returns(bool)
func (_ACL *ACLCallerSession) UnpausableAdminSet(arg0 common.Address) (bool, error) {
	return _ACL.Contract.UnpausableAdminSet(&_ACL.CallOpts, arg0)
}

// AddPausableAdmin is a paid mutator transaction binding the contract method 0x4910832f.
//
// Solidity: function addPausableAdmin(address newAdmin) returns()
func (_ACL *ACLTransactor) AddPausableAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _ACL.contract.Transact(opts, "addPausableAdmin", newAdmin)
}

// AddPausableAdmin is a paid mutator transaction binding the contract method 0x4910832f.
//
// Solidity: function addPausableAdmin(address newAdmin) returns()
func (_ACL *ACLSession) AddPausableAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _ACL.Contract.AddPausableAdmin(&_ACL.TransactOpts, newAdmin)
}

// AddPausableAdmin is a paid mutator transaction binding the contract method 0x4910832f.
//
// Solidity: function addPausableAdmin(address newAdmin) returns()
func (_ACL *ACLTransactorSession) AddPausableAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _ACL.Contract.AddPausableAdmin(&_ACL.TransactOpts, newAdmin)
}

// AddUnpausableAdmin is a paid mutator transaction binding the contract method 0x819ad68e.
//
// Solidity: function addUnpausableAdmin(address newAdmin) returns()
func (_ACL *ACLTransactor) AddUnpausableAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _ACL.contract.Transact(opts, "addUnpausableAdmin", newAdmin)
}

// AddUnpausableAdmin is a paid mutator transaction binding the contract method 0x819ad68e.
//
// Solidity: function addUnpausableAdmin(address newAdmin) returns()
func (_ACL *ACLSession) AddUnpausableAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _ACL.Contract.AddUnpausableAdmin(&_ACL.TransactOpts, newAdmin)
}

// AddUnpausableAdmin is a paid mutator transaction binding the contract method 0x819ad68e.
//
// Solidity: function addUnpausableAdmin(address newAdmin) returns()
func (_ACL *ACLTransactorSession) AddUnpausableAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _ACL.Contract.AddUnpausableAdmin(&_ACL.TransactOpts, newAdmin)
}

// RemovePausableAdmin is a paid mutator transaction binding the contract method 0xba306df1.
//
// Solidity: function removePausableAdmin(address admin) returns()
func (_ACL *ACLTransactor) RemovePausableAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _ACL.contract.Transact(opts, "removePausableAdmin", admin)
}

// RemovePausableAdmin is a paid mutator transaction binding the contract method 0xba306df1.
//
// Solidity: function removePausableAdmin(address admin) returns()
func (_ACL *ACLSession) RemovePausableAdmin(admin common.Address) (*types.Transaction, error) {
	return _ACL.Contract.RemovePausableAdmin(&_ACL.TransactOpts, admin)
}

// RemovePausableAdmin is a paid mutator transaction binding the contract method 0xba306df1.
//
// Solidity: function removePausableAdmin(address admin) returns()
func (_ACL *ACLTransactorSession) RemovePausableAdmin(admin common.Address) (*types.Transaction, error) {
	return _ACL.Contract.RemovePausableAdmin(&_ACL.TransactOpts, admin)
}

// RemoveUnpausableAdmin is a paid mutator transaction binding the contract method 0xadce758d.
//
// Solidity: function removeUnpausableAdmin(address admin) returns()
func (_ACL *ACLTransactor) RemoveUnpausableAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _ACL.contract.Transact(opts, "removeUnpausableAdmin", admin)
}

// RemoveUnpausableAdmin is a paid mutator transaction binding the contract method 0xadce758d.
//
// Solidity: function removeUnpausableAdmin(address admin) returns()
func (_ACL *ACLSession) RemoveUnpausableAdmin(admin common.Address) (*types.Transaction, error) {
	return _ACL.Contract.RemoveUnpausableAdmin(&_ACL.TransactOpts, admin)
}

// RemoveUnpausableAdmin is a paid mutator transaction binding the contract method 0xadce758d.
//
// Solidity: function removeUnpausableAdmin(address admin) returns()
func (_ACL *ACLTransactorSession) RemoveUnpausableAdmin(admin common.Address) (*types.Transaction, error) {
	return _ACL.Contract.RemoveUnpausableAdmin(&_ACL.TransactOpts, admin)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ACL *ACLTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ACL.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ACL *ACLSession) RenounceOwnership() (*types.Transaction, error) {
	return _ACL.Contract.RenounceOwnership(&_ACL.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ACL *ACLTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ACL.Contract.RenounceOwnership(&_ACL.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ACL *ACLTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ACL.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ACL *ACLSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ACL.Contract.TransferOwnership(&_ACL.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ACL *ACLTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ACL.Contract.TransferOwnership(&_ACL.TransactOpts, newOwner)
}

// ACLOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ACL contract.
type ACLOwnershipTransferredIterator struct {
	Event *ACLOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ACLOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ACLOwnershipTransferred)
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
		it.Event = new(ACLOwnershipTransferred)
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
func (it *ACLOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ACLOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ACLOwnershipTransferred represents a OwnershipTransferred event raised by the ACL contract.
type ACLOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ACL *ACLFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ACLOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ACL.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ACLOwnershipTransferredIterator{contract: _ACL.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ACL *ACLFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ACLOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ACL.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ACLOwnershipTransferred)
				if err := _ACL.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ACL *ACLFilterer) ParseOwnershipTransferred(log types.Log) (*ACLOwnershipTransferred, error) {
	event := new(ACLOwnershipTransferred)
	if err := _ACL.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ACLPausableAdminAddedIterator is returned from FilterPausableAdminAdded and is used to iterate over the raw logs and unpacked data for PausableAdminAdded events raised by the ACL contract.
type ACLPausableAdminAddedIterator struct {
	Event *ACLPausableAdminAdded // Event containing the contract specifics and raw log

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
func (it *ACLPausableAdminAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ACLPausableAdminAdded)
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
		it.Event = new(ACLPausableAdminAdded)
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
func (it *ACLPausableAdminAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ACLPausableAdminAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ACLPausableAdminAdded represents a PausableAdminAdded event raised by the ACL contract.
type ACLPausableAdminAdded struct {
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPausableAdminAdded is a free log retrieval operation binding the contract event 0xae26b1cfe9454ba87274a4e8330b6654684362d0f3d7bbd17f7449a1d38387c6.
//
// Solidity: event PausableAdminAdded(address indexed newAdmin)
func (_ACL *ACLFilterer) FilterPausableAdminAdded(opts *bind.FilterOpts, newAdmin []common.Address) (*ACLPausableAdminAddedIterator, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _ACL.contract.FilterLogs(opts, "PausableAdminAdded", newAdminRule)
	if err != nil {
		return nil, err
	}
	return &ACLPausableAdminAddedIterator{contract: _ACL.contract, event: "PausableAdminAdded", logs: logs, sub: sub}, nil
}

// WatchPausableAdminAdded is a free log subscription operation binding the contract event 0xae26b1cfe9454ba87274a4e8330b6654684362d0f3d7bbd17f7449a1d38387c6.
//
// Solidity: event PausableAdminAdded(address indexed newAdmin)
func (_ACL *ACLFilterer) WatchPausableAdminAdded(opts *bind.WatchOpts, sink chan<- *ACLPausableAdminAdded, newAdmin []common.Address) (event.Subscription, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _ACL.contract.WatchLogs(opts, "PausableAdminAdded", newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ACLPausableAdminAdded)
				if err := _ACL.contract.UnpackLog(event, "PausableAdminAdded", log); err != nil {
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

// ParsePausableAdminAdded is a log parse operation binding the contract event 0xae26b1cfe9454ba87274a4e8330b6654684362d0f3d7bbd17f7449a1d38387c6.
//
// Solidity: event PausableAdminAdded(address indexed newAdmin)
func (_ACL *ACLFilterer) ParsePausableAdminAdded(log types.Log) (*ACLPausableAdminAdded, error) {
	event := new(ACLPausableAdminAdded)
	if err := _ACL.contract.UnpackLog(event, "PausableAdminAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ACLPausableAdminRemovedIterator is returned from FilterPausableAdminRemoved and is used to iterate over the raw logs and unpacked data for PausableAdminRemoved events raised by the ACL contract.
type ACLPausableAdminRemovedIterator struct {
	Event *ACLPausableAdminRemoved // Event containing the contract specifics and raw log

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
func (it *ACLPausableAdminRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ACLPausableAdminRemoved)
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
		it.Event = new(ACLPausableAdminRemoved)
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
func (it *ACLPausableAdminRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ACLPausableAdminRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ACLPausableAdminRemoved represents a PausableAdminRemoved event raised by the ACL contract.
type ACLPausableAdminRemoved struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPausableAdminRemoved is a free log retrieval operation binding the contract event 0x28b01395b7e25d20552a0c8dc8ecd3b1d4abc986f14dad7885fd45b6fd73c8d9.
//
// Solidity: event PausableAdminRemoved(address indexed admin)
func (_ACL *ACLFilterer) FilterPausableAdminRemoved(opts *bind.FilterOpts, admin []common.Address) (*ACLPausableAdminRemovedIterator, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _ACL.contract.FilterLogs(opts, "PausableAdminRemoved", adminRule)
	if err != nil {
		return nil, err
	}
	return &ACLPausableAdminRemovedIterator{contract: _ACL.contract, event: "PausableAdminRemoved", logs: logs, sub: sub}, nil
}

// WatchPausableAdminRemoved is a free log subscription operation binding the contract event 0x28b01395b7e25d20552a0c8dc8ecd3b1d4abc986f14dad7885fd45b6fd73c8d9.
//
// Solidity: event PausableAdminRemoved(address indexed admin)
func (_ACL *ACLFilterer) WatchPausableAdminRemoved(opts *bind.WatchOpts, sink chan<- *ACLPausableAdminRemoved, admin []common.Address) (event.Subscription, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _ACL.contract.WatchLogs(opts, "PausableAdminRemoved", adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ACLPausableAdminRemoved)
				if err := _ACL.contract.UnpackLog(event, "PausableAdminRemoved", log); err != nil {
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

// ParsePausableAdminRemoved is a log parse operation binding the contract event 0x28b01395b7e25d20552a0c8dc8ecd3b1d4abc986f14dad7885fd45b6fd73c8d9.
//
// Solidity: event PausableAdminRemoved(address indexed admin)
func (_ACL *ACLFilterer) ParsePausableAdminRemoved(log types.Log) (*ACLPausableAdminRemoved, error) {
	event := new(ACLPausableAdminRemoved)
	if err := _ACL.contract.UnpackLog(event, "PausableAdminRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ACLUnpausableAdminAddedIterator is returned from FilterUnpausableAdminAdded and is used to iterate over the raw logs and unpacked data for UnpausableAdminAdded events raised by the ACL contract.
type ACLUnpausableAdminAddedIterator struct {
	Event *ACLUnpausableAdminAdded // Event containing the contract specifics and raw log

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
func (it *ACLUnpausableAdminAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ACLUnpausableAdminAdded)
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
		it.Event = new(ACLUnpausableAdminAdded)
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
func (it *ACLUnpausableAdminAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ACLUnpausableAdminAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ACLUnpausableAdminAdded represents a UnpausableAdminAdded event raised by the ACL contract.
type ACLUnpausableAdminAdded struct {
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUnpausableAdminAdded is a free log retrieval operation binding the contract event 0xd400da6c0c0a894dacc0981730b88af0545d00272ee8fff1437bf560ff245fc4.
//
// Solidity: event UnpausableAdminAdded(address indexed newAdmin)
func (_ACL *ACLFilterer) FilterUnpausableAdminAdded(opts *bind.FilterOpts, newAdmin []common.Address) (*ACLUnpausableAdminAddedIterator, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _ACL.contract.FilterLogs(opts, "UnpausableAdminAdded", newAdminRule)
	if err != nil {
		return nil, err
	}
	return &ACLUnpausableAdminAddedIterator{contract: _ACL.contract, event: "UnpausableAdminAdded", logs: logs, sub: sub}, nil
}

// WatchUnpausableAdminAdded is a free log subscription operation binding the contract event 0xd400da6c0c0a894dacc0981730b88af0545d00272ee8fff1437bf560ff245fc4.
//
// Solidity: event UnpausableAdminAdded(address indexed newAdmin)
func (_ACL *ACLFilterer) WatchUnpausableAdminAdded(opts *bind.WatchOpts, sink chan<- *ACLUnpausableAdminAdded, newAdmin []common.Address) (event.Subscription, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _ACL.contract.WatchLogs(opts, "UnpausableAdminAdded", newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ACLUnpausableAdminAdded)
				if err := _ACL.contract.UnpackLog(event, "UnpausableAdminAdded", log); err != nil {
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

// ParseUnpausableAdminAdded is a log parse operation binding the contract event 0xd400da6c0c0a894dacc0981730b88af0545d00272ee8fff1437bf560ff245fc4.
//
// Solidity: event UnpausableAdminAdded(address indexed newAdmin)
func (_ACL *ACLFilterer) ParseUnpausableAdminAdded(log types.Log) (*ACLUnpausableAdminAdded, error) {
	event := new(ACLUnpausableAdminAdded)
	if err := _ACL.contract.UnpackLog(event, "UnpausableAdminAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ACLUnpausableAdminRemovedIterator is returned from FilterUnpausableAdminRemoved and is used to iterate over the raw logs and unpacked data for UnpausableAdminRemoved events raised by the ACL contract.
type ACLUnpausableAdminRemovedIterator struct {
	Event *ACLUnpausableAdminRemoved // Event containing the contract specifics and raw log

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
func (it *ACLUnpausableAdminRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ACLUnpausableAdminRemoved)
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
		it.Event = new(ACLUnpausableAdminRemoved)
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
func (it *ACLUnpausableAdminRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ACLUnpausableAdminRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ACLUnpausableAdminRemoved represents a UnpausableAdminRemoved event raised by the ACL contract.
type ACLUnpausableAdminRemoved struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUnpausableAdminRemoved is a free log retrieval operation binding the contract event 0x1998397e7203f7baca9d6f41b9e4da6e768daac5caad4234fb9bf5869d271545.
//
// Solidity: event UnpausableAdminRemoved(address indexed admin)
func (_ACL *ACLFilterer) FilterUnpausableAdminRemoved(opts *bind.FilterOpts, admin []common.Address) (*ACLUnpausableAdminRemovedIterator, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _ACL.contract.FilterLogs(opts, "UnpausableAdminRemoved", adminRule)
	if err != nil {
		return nil, err
	}
	return &ACLUnpausableAdminRemovedIterator{contract: _ACL.contract, event: "UnpausableAdminRemoved", logs: logs, sub: sub}, nil
}

// WatchUnpausableAdminRemoved is a free log subscription operation binding the contract event 0x1998397e7203f7baca9d6f41b9e4da6e768daac5caad4234fb9bf5869d271545.
//
// Solidity: event UnpausableAdminRemoved(address indexed admin)
func (_ACL *ACLFilterer) WatchUnpausableAdminRemoved(opts *bind.WatchOpts, sink chan<- *ACLUnpausableAdminRemoved, admin []common.Address) (event.Subscription, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _ACL.contract.WatchLogs(opts, "UnpausableAdminRemoved", adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ACLUnpausableAdminRemoved)
				if err := _ACL.contract.UnpackLog(event, "UnpausableAdminRemoved", log); err != nil {
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

// ParseUnpausableAdminRemoved is a log parse operation binding the contract event 0x1998397e7203f7baca9d6f41b9e4da6e768daac5caad4234fb9bf5869d271545.
//
// Solidity: event UnpausableAdminRemoved(address indexed admin)
func (_ACL *ACLFilterer) ParseUnpausableAdminRemoved(log types.Log) (*ACLUnpausableAdminRemoved, error) {
	event := new(ACLUnpausableAdminRemoved)
	if err := _ACL.contract.UnpackLog(event, "UnpausableAdminRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
