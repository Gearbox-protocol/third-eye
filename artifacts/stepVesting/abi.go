// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stepVesting

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

// StepVestingMetaData contains all meta data concerning the StepVesting contract.
var StepVestingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIGearToken\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_started\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_cliffDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_stepDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_cliffAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_stepAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_numOfSteps\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldWallet\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newWallet\",\"type\":\"address\"}],\"name\":\"ReceiverChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"available\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cliffAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cliffDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"}],\"name\":\"delegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numOfSteps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"receiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"setReceiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"started\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stepAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stepDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIGearToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// StepVestingABI is the input ABI used to generate the binding from.
// Deprecated: Use StepVestingMetaData.ABI instead.
var StepVestingABI = StepVestingMetaData.ABI

// StepVesting is an auto generated Go binding around an Ethereum contract.
type StepVesting struct {
	StepVestingCaller     // Read-only binding to the contract
	StepVestingTransactor // Write-only binding to the contract
	StepVestingFilterer   // Log filterer for contract events
}

// StepVestingCaller is an auto generated read-only Go binding around an Ethereum contract.
type StepVestingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StepVestingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StepVestingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StepVestingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StepVestingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StepVestingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StepVestingSession struct {
	Contract     *StepVesting      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StepVestingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StepVestingCallerSession struct {
	Contract *StepVestingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// StepVestingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StepVestingTransactorSession struct {
	Contract     *StepVestingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// StepVestingRaw is an auto generated low-level Go binding around an Ethereum contract.
type StepVestingRaw struct {
	Contract *StepVesting // Generic contract binding to access the raw methods on
}

// StepVestingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StepVestingCallerRaw struct {
	Contract *StepVestingCaller // Generic read-only contract binding to access the raw methods on
}

// StepVestingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StepVestingTransactorRaw struct {
	Contract *StepVestingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStepVesting creates a new instance of StepVesting, bound to a specific deployed contract.
func NewStepVesting(address common.Address, backend bind.ContractBackend) (*StepVesting, error) {
	contract, err := bindStepVesting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StepVesting{StepVestingCaller: StepVestingCaller{contract: contract}, StepVestingTransactor: StepVestingTransactor{contract: contract}, StepVestingFilterer: StepVestingFilterer{contract: contract}}, nil
}

// NewStepVestingCaller creates a new read-only instance of StepVesting, bound to a specific deployed contract.
func NewStepVestingCaller(address common.Address, caller bind.ContractCaller) (*StepVestingCaller, error) {
	contract, err := bindStepVesting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StepVestingCaller{contract: contract}, nil
}

// NewStepVestingTransactor creates a new write-only instance of StepVesting, bound to a specific deployed contract.
func NewStepVestingTransactor(address common.Address, transactor bind.ContractTransactor) (*StepVestingTransactor, error) {
	contract, err := bindStepVesting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StepVestingTransactor{contract: contract}, nil
}

// NewStepVestingFilterer creates a new log filterer instance of StepVesting, bound to a specific deployed contract.
func NewStepVestingFilterer(address common.Address, filterer bind.ContractFilterer) (*StepVestingFilterer, error) {
	contract, err := bindStepVesting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StepVestingFilterer{contract: contract}, nil
}

// bindStepVesting binds a generic wrapper to an already deployed contract.
func bindStepVesting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StepVestingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StepVesting *StepVestingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StepVesting.Contract.StepVestingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StepVesting *StepVestingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StepVesting.Contract.StepVestingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StepVesting *StepVestingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StepVesting.Contract.StepVestingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StepVesting *StepVestingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StepVesting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StepVesting *StepVestingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StepVesting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StepVesting *StepVestingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StepVesting.Contract.contract.Transact(opts, method, params...)
}

// Available is a free data retrieval call binding the contract method 0x48a0d754.
//
// Solidity: function available() view returns(uint256)
func (_StepVesting *StepVestingCaller) Available(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StepVesting.contract.Call(opts, &out, "available")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Available is a free data retrieval call binding the contract method 0x48a0d754.
//
// Solidity: function available() view returns(uint256)
func (_StepVesting *StepVestingSession) Available() (*big.Int, error) {
	return _StepVesting.Contract.Available(&_StepVesting.CallOpts)
}

// Available is a free data retrieval call binding the contract method 0x48a0d754.
//
// Solidity: function available() view returns(uint256)
func (_StepVesting *StepVestingCallerSession) Available() (*big.Int, error) {
	return _StepVesting.Contract.Available(&_StepVesting.CallOpts)
}

// Claimable is a free data retrieval call binding the contract method 0xaf38d757.
//
// Solidity: function claimable() view returns(uint256)
func (_StepVesting *StepVestingCaller) Claimable(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StepVesting.contract.Call(opts, &out, "claimable")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Claimable is a free data retrieval call binding the contract method 0xaf38d757.
//
// Solidity: function claimable() view returns(uint256)
func (_StepVesting *StepVestingSession) Claimable() (*big.Int, error) {
	return _StepVesting.Contract.Claimable(&_StepVesting.CallOpts)
}

// Claimable is a free data retrieval call binding the contract method 0xaf38d757.
//
// Solidity: function claimable() view returns(uint256)
func (_StepVesting *StepVestingCallerSession) Claimable() (*big.Int, error) {
	return _StepVesting.Contract.Claimable(&_StepVesting.CallOpts)
}

// Claimed is a free data retrieval call binding the contract method 0xe834a834.
//
// Solidity: function claimed() view returns(uint256)
func (_StepVesting *StepVestingCaller) Claimed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StepVesting.contract.Call(opts, &out, "claimed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Claimed is a free data retrieval call binding the contract method 0xe834a834.
//
// Solidity: function claimed() view returns(uint256)
func (_StepVesting *StepVestingSession) Claimed() (*big.Int, error) {
	return _StepVesting.Contract.Claimed(&_StepVesting.CallOpts)
}

// Claimed is a free data retrieval call binding the contract method 0xe834a834.
//
// Solidity: function claimed() view returns(uint256)
func (_StepVesting *StepVestingCallerSession) Claimed() (*big.Int, error) {
	return _StepVesting.Contract.Claimed(&_StepVesting.CallOpts)
}

// CliffAmount is a free data retrieval call binding the contract method 0x460ad439.
//
// Solidity: function cliffAmount() view returns(uint256)
func (_StepVesting *StepVestingCaller) CliffAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StepVesting.contract.Call(opts, &out, "cliffAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CliffAmount is a free data retrieval call binding the contract method 0x460ad439.
//
// Solidity: function cliffAmount() view returns(uint256)
func (_StepVesting *StepVestingSession) CliffAmount() (*big.Int, error) {
	return _StepVesting.Contract.CliffAmount(&_StepVesting.CallOpts)
}

// CliffAmount is a free data retrieval call binding the contract method 0x460ad439.
//
// Solidity: function cliffAmount() view returns(uint256)
func (_StepVesting *StepVestingCallerSession) CliffAmount() (*big.Int, error) {
	return _StepVesting.Contract.CliffAmount(&_StepVesting.CallOpts)
}

// CliffDuration is a free data retrieval call binding the contract method 0xd85349f7.
//
// Solidity: function cliffDuration() view returns(uint256)
func (_StepVesting *StepVestingCaller) CliffDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StepVesting.contract.Call(opts, &out, "cliffDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CliffDuration is a free data retrieval call binding the contract method 0xd85349f7.
//
// Solidity: function cliffDuration() view returns(uint256)
func (_StepVesting *StepVestingSession) CliffDuration() (*big.Int, error) {
	return _StepVesting.Contract.CliffDuration(&_StepVesting.CallOpts)
}

// CliffDuration is a free data retrieval call binding the contract method 0xd85349f7.
//
// Solidity: function cliffDuration() view returns(uint256)
func (_StepVesting *StepVestingCallerSession) CliffDuration() (*big.Int, error) {
	return _StepVesting.Contract.CliffDuration(&_StepVesting.CallOpts)
}

// NumOfSteps is a free data retrieval call binding the contract method 0x5d1fbf54.
//
// Solidity: function numOfSteps() view returns(uint256)
func (_StepVesting *StepVestingCaller) NumOfSteps(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StepVesting.contract.Call(opts, &out, "numOfSteps")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumOfSteps is a free data retrieval call binding the contract method 0x5d1fbf54.
//
// Solidity: function numOfSteps() view returns(uint256)
func (_StepVesting *StepVestingSession) NumOfSteps() (*big.Int, error) {
	return _StepVesting.Contract.NumOfSteps(&_StepVesting.CallOpts)
}

// NumOfSteps is a free data retrieval call binding the contract method 0x5d1fbf54.
//
// Solidity: function numOfSteps() view returns(uint256)
func (_StepVesting *StepVestingCallerSession) NumOfSteps() (*big.Int, error) {
	return _StepVesting.Contract.NumOfSteps(&_StepVesting.CallOpts)
}

// Receiver is a free data retrieval call binding the contract method 0xf7260d3e.
//
// Solidity: function receiver() view returns(address)
func (_StepVesting *StepVestingCaller) Receiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StepVesting.contract.Call(opts, &out, "receiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Receiver is a free data retrieval call binding the contract method 0xf7260d3e.
//
// Solidity: function receiver() view returns(address)
func (_StepVesting *StepVestingSession) Receiver() (common.Address, error) {
	return _StepVesting.Contract.Receiver(&_StepVesting.CallOpts)
}

// Receiver is a free data retrieval call binding the contract method 0xf7260d3e.
//
// Solidity: function receiver() view returns(address)
func (_StepVesting *StepVestingCallerSession) Receiver() (common.Address, error) {
	return _StepVesting.Contract.Receiver(&_StepVesting.CallOpts)
}

// Started is a free data retrieval call binding the contract method 0x1f2698ab.
//
// Solidity: function started() view returns(uint256)
func (_StepVesting *StepVestingCaller) Started(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StepVesting.contract.Call(opts, &out, "started")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Started is a free data retrieval call binding the contract method 0x1f2698ab.
//
// Solidity: function started() view returns(uint256)
func (_StepVesting *StepVestingSession) Started() (*big.Int, error) {
	return _StepVesting.Contract.Started(&_StepVesting.CallOpts)
}

// Started is a free data retrieval call binding the contract method 0x1f2698ab.
//
// Solidity: function started() view returns(uint256)
func (_StepVesting *StepVestingCallerSession) Started() (*big.Int, error) {
	return _StepVesting.Contract.Started(&_StepVesting.CallOpts)
}

// StepAmount is a free data retrieval call binding the contract method 0x1989488b.
//
// Solidity: function stepAmount() view returns(uint256)
func (_StepVesting *StepVestingCaller) StepAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StepVesting.contract.Call(opts, &out, "stepAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StepAmount is a free data retrieval call binding the contract method 0x1989488b.
//
// Solidity: function stepAmount() view returns(uint256)
func (_StepVesting *StepVestingSession) StepAmount() (*big.Int, error) {
	return _StepVesting.Contract.StepAmount(&_StepVesting.CallOpts)
}

// StepAmount is a free data retrieval call binding the contract method 0x1989488b.
//
// Solidity: function stepAmount() view returns(uint256)
func (_StepVesting *StepVestingCallerSession) StepAmount() (*big.Int, error) {
	return _StepVesting.Contract.StepAmount(&_StepVesting.CallOpts)
}

// StepDuration is a free data retrieval call binding the contract method 0x4a4e5776.
//
// Solidity: function stepDuration() view returns(uint256)
func (_StepVesting *StepVestingCaller) StepDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StepVesting.contract.Call(opts, &out, "stepDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StepDuration is a free data retrieval call binding the contract method 0x4a4e5776.
//
// Solidity: function stepDuration() view returns(uint256)
func (_StepVesting *StepVestingSession) StepDuration() (*big.Int, error) {
	return _StepVesting.Contract.StepDuration(&_StepVesting.CallOpts)
}

// StepDuration is a free data retrieval call binding the contract method 0x4a4e5776.
//
// Solidity: function stepDuration() view returns(uint256)
func (_StepVesting *StepVestingCallerSession) StepDuration() (*big.Int, error) {
	return _StepVesting.Contract.StepDuration(&_StepVesting.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_StepVesting *StepVestingCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StepVesting.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_StepVesting *StepVestingSession) Token() (common.Address, error) {
	return _StepVesting.Contract.Token(&_StepVesting.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_StepVesting *StepVestingCallerSession) Token() (common.Address, error) {
	return _StepVesting.Contract.Token(&_StepVesting.CallOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_StepVesting *StepVestingTransactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StepVesting.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_StepVesting *StepVestingSession) Claim() (*types.Transaction, error) {
	return _StepVesting.Contract.Claim(&_StepVesting.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_StepVesting *StepVestingTransactorSession) Claim() (*types.Transaction, error) {
	return _StepVesting.Contract.Claim(&_StepVesting.TransactOpts)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address delegatee) returns()
func (_StepVesting *StepVestingTransactor) Delegate(opts *bind.TransactOpts, delegatee common.Address) (*types.Transaction, error) {
	return _StepVesting.contract.Transact(opts, "delegate", delegatee)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address delegatee) returns()
func (_StepVesting *StepVestingSession) Delegate(delegatee common.Address) (*types.Transaction, error) {
	return _StepVesting.Contract.Delegate(&_StepVesting.TransactOpts, delegatee)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address delegatee) returns()
func (_StepVesting *StepVestingTransactorSession) Delegate(delegatee common.Address) (*types.Transaction, error) {
	return _StepVesting.Contract.Delegate(&_StepVesting.TransactOpts, delegatee)
}

// SetReceiver is a paid mutator transaction binding the contract method 0x718da7ee.
//
// Solidity: function setReceiver(address _receiver) returns()
func (_StepVesting *StepVestingTransactor) SetReceiver(opts *bind.TransactOpts, _receiver common.Address) (*types.Transaction, error) {
	return _StepVesting.contract.Transact(opts, "setReceiver", _receiver)
}

// SetReceiver is a paid mutator transaction binding the contract method 0x718da7ee.
//
// Solidity: function setReceiver(address _receiver) returns()
func (_StepVesting *StepVestingSession) SetReceiver(_receiver common.Address) (*types.Transaction, error) {
	return _StepVesting.Contract.SetReceiver(&_StepVesting.TransactOpts, _receiver)
}

// SetReceiver is a paid mutator transaction binding the contract method 0x718da7ee.
//
// Solidity: function setReceiver(address _receiver) returns()
func (_StepVesting *StepVestingTransactorSession) SetReceiver(_receiver common.Address) (*types.Transaction, error) {
	return _StepVesting.Contract.SetReceiver(&_StepVesting.TransactOpts, _receiver)
}

// StepVestingReceiverChangedIterator is returned from FilterReceiverChanged and is used to iterate over the raw logs and unpacked data for ReceiverChanged events raised by the StepVesting contract.
type StepVestingReceiverChangedIterator struct {
	Event *StepVestingReceiverChanged // Event containing the contract specifics and raw log

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
func (it *StepVestingReceiverChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StepVestingReceiverChanged)
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
		it.Event = new(StepVestingReceiverChanged)
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
func (it *StepVestingReceiverChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StepVestingReceiverChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StepVestingReceiverChanged represents a ReceiverChanged event raised by the StepVesting contract.
type StepVestingReceiverChanged struct {
	OldWallet common.Address
	NewWallet common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterReceiverChanged is a free log retrieval operation binding the contract event 0xd36aafedb017e43b79d3cf6aa1987d3fbb9fff33e1738c71dbf6b2abaadbded0.
//
// Solidity: event ReceiverChanged(address oldWallet, address newWallet)
func (_StepVesting *StepVestingFilterer) FilterReceiverChanged(opts *bind.FilterOpts) (*StepVestingReceiverChangedIterator, error) {

	logs, sub, err := _StepVesting.contract.FilterLogs(opts, "ReceiverChanged")
	if err != nil {
		return nil, err
	}
	return &StepVestingReceiverChangedIterator{contract: _StepVesting.contract, event: "ReceiverChanged", logs: logs, sub: sub}, nil
}

// WatchReceiverChanged is a free log subscription operation binding the contract event 0xd36aafedb017e43b79d3cf6aa1987d3fbb9fff33e1738c71dbf6b2abaadbded0.
//
// Solidity: event ReceiverChanged(address oldWallet, address newWallet)
func (_StepVesting *StepVestingFilterer) WatchReceiverChanged(opts *bind.WatchOpts, sink chan<- *StepVestingReceiverChanged) (event.Subscription, error) {

	logs, sub, err := _StepVesting.contract.WatchLogs(opts, "ReceiverChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StepVestingReceiverChanged)
				if err := _StepVesting.contract.UnpackLog(event, "ReceiverChanged", log); err != nil {
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

// ParseReceiverChanged is a log parse operation binding the contract event 0xd36aafedb017e43b79d3cf6aa1987d3fbb9fff33e1738c71dbf6b2abaadbded0.
//
// Solidity: event ReceiverChanged(address oldWallet, address newWallet)
func (_StepVesting *StepVestingFilterer) ParseReceiverChanged(log types.Log) (*StepVestingReceiverChanged, error) {
	event := new(StepVestingReceiverChanged)
	if err := _StepVesting.contract.UnpackLog(event, "ReceiverChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
