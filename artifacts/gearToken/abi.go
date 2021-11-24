// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gearToken

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

// GearTokenMetaData contains all meta data concerning the GearToken contract.
var GearTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fromDelegate\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"toDelegate\",\"type\":\"address\"}],\"name\":\"DelegateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBalance\",\"type\":\"uint256\"}],\"name\":\"DelegateVotesChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"}],\"name\":\"MinerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"TransferAllowed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DELEGATION_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMIT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allowTransfers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"rawAmount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"checkpoints\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"fromBlock\",\"type\":\"uint32\"},{\"internalType\":\"uint96\",\"name\":\"votes\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"}],\"name\":\"delegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"delegateBySig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"delegates\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getCurrentVotes\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getPriorVotes\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"manager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"miner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"numCheckpoints\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"rawAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_miner\",\"type\":\"address\"}],\"name\":\"setMiner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"rawAmount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"rawAmount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newManager\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transfersAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// GearTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use GearTokenMetaData.ABI instead.
var GearTokenABI = GearTokenMetaData.ABI

// GearToken is an auto generated Go binding around an Ethereum contract.
type GearToken struct {
	GearTokenCaller     // Read-only binding to the contract
	GearTokenTransactor // Write-only binding to the contract
	GearTokenFilterer   // Log filterer for contract events
}

// GearTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type GearTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GearTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GearTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GearTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GearTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GearTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GearTokenSession struct {
	Contract     *GearToken        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GearTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GearTokenCallerSession struct {
	Contract *GearTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// GearTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GearTokenTransactorSession struct {
	Contract     *GearTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// GearTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type GearTokenRaw struct {
	Contract *GearToken // Generic contract binding to access the raw methods on
}

// GearTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GearTokenCallerRaw struct {
	Contract *GearTokenCaller // Generic read-only contract binding to access the raw methods on
}

// GearTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GearTokenTransactorRaw struct {
	Contract *GearTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGearToken creates a new instance of GearToken, bound to a specific deployed contract.
func NewGearToken(address common.Address, backend bind.ContractBackend) (*GearToken, error) {
	contract, err := bindGearToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GearToken{GearTokenCaller: GearTokenCaller{contract: contract}, GearTokenTransactor: GearTokenTransactor{contract: contract}, GearTokenFilterer: GearTokenFilterer{contract: contract}}, nil
}

// NewGearTokenCaller creates a new read-only instance of GearToken, bound to a specific deployed contract.
func NewGearTokenCaller(address common.Address, caller bind.ContractCaller) (*GearTokenCaller, error) {
	contract, err := bindGearToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GearTokenCaller{contract: contract}, nil
}

// NewGearTokenTransactor creates a new write-only instance of GearToken, bound to a specific deployed contract.
func NewGearTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*GearTokenTransactor, error) {
	contract, err := bindGearToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GearTokenTransactor{contract: contract}, nil
}

// NewGearTokenFilterer creates a new log filterer instance of GearToken, bound to a specific deployed contract.
func NewGearTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*GearTokenFilterer, error) {
	contract, err := bindGearToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GearTokenFilterer{contract: contract}, nil
}

// bindGearToken binds a generic wrapper to an already deployed contract.
func bindGearToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GearTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GearToken *GearTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GearToken.Contract.GearTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GearToken *GearTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GearToken.Contract.GearTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GearToken *GearTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GearToken.Contract.GearTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GearToken *GearTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GearToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GearToken *GearTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GearToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GearToken *GearTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GearToken.Contract.contract.Transact(opts, method, params...)
}

// DELEGATIONTYPEHASH is a free data retrieval call binding the contract method 0xe7a324dc.
//
// Solidity: function DELEGATION_TYPEHASH() view returns(bytes32)
func (_GearToken *GearTokenCaller) DELEGATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GearToken.contract.Call(opts, &out, "DELEGATION_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DELEGATIONTYPEHASH is a free data retrieval call binding the contract method 0xe7a324dc.
//
// Solidity: function DELEGATION_TYPEHASH() view returns(bytes32)
func (_GearToken *GearTokenSession) DELEGATIONTYPEHASH() ([32]byte, error) {
	return _GearToken.Contract.DELEGATIONTYPEHASH(&_GearToken.CallOpts)
}

// DELEGATIONTYPEHASH is a free data retrieval call binding the contract method 0xe7a324dc.
//
// Solidity: function DELEGATION_TYPEHASH() view returns(bytes32)
func (_GearToken *GearTokenCallerSession) DELEGATIONTYPEHASH() ([32]byte, error) {
	return _GearToken.Contract.DELEGATIONTYPEHASH(&_GearToken.CallOpts)
}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_GearToken *GearTokenCaller) DOMAINTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GearToken.contract.Call(opts, &out, "DOMAIN_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_GearToken *GearTokenSession) DOMAINTYPEHASH() ([32]byte, error) {
	return _GearToken.Contract.DOMAINTYPEHASH(&_GearToken.CallOpts)
}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_GearToken *GearTokenCallerSession) DOMAINTYPEHASH() ([32]byte, error) {
	return _GearToken.Contract.DOMAINTYPEHASH(&_GearToken.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_GearToken *GearTokenCaller) PERMITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GearToken.contract.Call(opts, &out, "PERMIT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_GearToken *GearTokenSession) PERMITTYPEHASH() ([32]byte, error) {
	return _GearToken.Contract.PERMITTYPEHASH(&_GearToken.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_GearToken *GearTokenCallerSession) PERMITTYPEHASH() ([32]byte, error) {
	return _GearToken.Contract.PERMITTYPEHASH(&_GearToken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address account, address spender) view returns(uint256)
func (_GearToken *GearTokenCaller) Allowance(opts *bind.CallOpts, account common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GearToken.contract.Call(opts, &out, "allowance", account, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address account, address spender) view returns(uint256)
func (_GearToken *GearTokenSession) Allowance(account common.Address, spender common.Address) (*big.Int, error) {
	return _GearToken.Contract.Allowance(&_GearToken.CallOpts, account, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address account, address spender) view returns(uint256)
func (_GearToken *GearTokenCallerSession) Allowance(account common.Address, spender common.Address) (*big.Int, error) {
	return _GearToken.Contract.Allowance(&_GearToken.CallOpts, account, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_GearToken *GearTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GearToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_GearToken *GearTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _GearToken.Contract.BalanceOf(&_GearToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_GearToken *GearTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _GearToken.Contract.BalanceOf(&_GearToken.CallOpts, account)
}

// Checkpoints is a free data retrieval call binding the contract method 0xf1127ed8.
//
// Solidity: function checkpoints(address , uint32 ) view returns(uint32 fromBlock, uint96 votes)
func (_GearToken *GearTokenCaller) Checkpoints(opts *bind.CallOpts, arg0 common.Address, arg1 uint32) (struct {
	FromBlock uint32
	Votes     *big.Int
}, error) {
	var out []interface{}
	err := _GearToken.contract.Call(opts, &out, "checkpoints", arg0, arg1)

	outstruct := new(struct {
		FromBlock uint32
		Votes     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FromBlock = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.Votes = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Checkpoints is a free data retrieval call binding the contract method 0xf1127ed8.
//
// Solidity: function checkpoints(address , uint32 ) view returns(uint32 fromBlock, uint96 votes)
func (_GearToken *GearTokenSession) Checkpoints(arg0 common.Address, arg1 uint32) (struct {
	FromBlock uint32
	Votes     *big.Int
}, error) {
	return _GearToken.Contract.Checkpoints(&_GearToken.CallOpts, arg0, arg1)
}

// Checkpoints is a free data retrieval call binding the contract method 0xf1127ed8.
//
// Solidity: function checkpoints(address , uint32 ) view returns(uint32 fromBlock, uint96 votes)
func (_GearToken *GearTokenCallerSession) Checkpoints(arg0 common.Address, arg1 uint32) (struct {
	FromBlock uint32
	Votes     *big.Int
}, error) {
	return _GearToken.Contract.Checkpoints(&_GearToken.CallOpts, arg0, arg1)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_GearToken *GearTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _GearToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_GearToken *GearTokenSession) Decimals() (uint8, error) {
	return _GearToken.Contract.Decimals(&_GearToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_GearToken *GearTokenCallerSession) Decimals() (uint8, error) {
	return _GearToken.Contract.Decimals(&_GearToken.CallOpts)
}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address ) view returns(address)
func (_GearToken *GearTokenCaller) Delegates(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _GearToken.contract.Call(opts, &out, "delegates", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address ) view returns(address)
func (_GearToken *GearTokenSession) Delegates(arg0 common.Address) (common.Address, error) {
	return _GearToken.Contract.Delegates(&_GearToken.CallOpts, arg0)
}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address ) view returns(address)
func (_GearToken *GearTokenCallerSession) Delegates(arg0 common.Address) (common.Address, error) {
	return _GearToken.Contract.Delegates(&_GearToken.CallOpts, arg0)
}

// GetCurrentVotes is a free data retrieval call binding the contract method 0xb4b5ea57.
//
// Solidity: function getCurrentVotes(address account) view returns(uint96)
func (_GearToken *GearTokenCaller) GetCurrentVotes(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GearToken.contract.Call(opts, &out, "getCurrentVotes", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentVotes is a free data retrieval call binding the contract method 0xb4b5ea57.
//
// Solidity: function getCurrentVotes(address account) view returns(uint96)
func (_GearToken *GearTokenSession) GetCurrentVotes(account common.Address) (*big.Int, error) {
	return _GearToken.Contract.GetCurrentVotes(&_GearToken.CallOpts, account)
}

// GetCurrentVotes is a free data retrieval call binding the contract method 0xb4b5ea57.
//
// Solidity: function getCurrentVotes(address account) view returns(uint96)
func (_GearToken *GearTokenCallerSession) GetCurrentVotes(account common.Address) (*big.Int, error) {
	return _GearToken.Contract.GetCurrentVotes(&_GearToken.CallOpts, account)
}

// GetPriorVotes is a free data retrieval call binding the contract method 0x782d6fe1.
//
// Solidity: function getPriorVotes(address account, uint256 blockNumber) view returns(uint96)
func (_GearToken *GearTokenCaller) GetPriorVotes(opts *bind.CallOpts, account common.Address, blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _GearToken.contract.Call(opts, &out, "getPriorVotes", account, blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPriorVotes is a free data retrieval call binding the contract method 0x782d6fe1.
//
// Solidity: function getPriorVotes(address account, uint256 blockNumber) view returns(uint96)
func (_GearToken *GearTokenSession) GetPriorVotes(account common.Address, blockNumber *big.Int) (*big.Int, error) {
	return _GearToken.Contract.GetPriorVotes(&_GearToken.CallOpts, account, blockNumber)
}

// GetPriorVotes is a free data retrieval call binding the contract method 0x782d6fe1.
//
// Solidity: function getPriorVotes(address account, uint256 blockNumber) view returns(uint96)
func (_GearToken *GearTokenCallerSession) GetPriorVotes(account common.Address, blockNumber *big.Int) (*big.Int, error) {
	return _GearToken.Contract.GetPriorVotes(&_GearToken.CallOpts, account, blockNumber)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_GearToken *GearTokenCaller) Manager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GearToken.contract.Call(opts, &out, "manager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_GearToken *GearTokenSession) Manager() (common.Address, error) {
	return _GearToken.Contract.Manager(&_GearToken.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_GearToken *GearTokenCallerSession) Manager() (common.Address, error) {
	return _GearToken.Contract.Manager(&_GearToken.CallOpts)
}

// Miner is a free data retrieval call binding the contract method 0x349dc329.
//
// Solidity: function miner() view returns(address)
func (_GearToken *GearTokenCaller) Miner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GearToken.contract.Call(opts, &out, "miner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Miner is a free data retrieval call binding the contract method 0x349dc329.
//
// Solidity: function miner() view returns(address)
func (_GearToken *GearTokenSession) Miner() (common.Address, error) {
	return _GearToken.Contract.Miner(&_GearToken.CallOpts)
}

// Miner is a free data retrieval call binding the contract method 0x349dc329.
//
// Solidity: function miner() view returns(address)
func (_GearToken *GearTokenCallerSession) Miner() (common.Address, error) {
	return _GearToken.Contract.Miner(&_GearToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_GearToken *GearTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _GearToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_GearToken *GearTokenSession) Name() (string, error) {
	return _GearToken.Contract.Name(&_GearToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_GearToken *GearTokenCallerSession) Name() (string, error) {
	return _GearToken.Contract.Name(&_GearToken.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_GearToken *GearTokenCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GearToken.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_GearToken *GearTokenSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _GearToken.Contract.Nonces(&_GearToken.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_GearToken *GearTokenCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _GearToken.Contract.Nonces(&_GearToken.CallOpts, arg0)
}

// NumCheckpoints is a free data retrieval call binding the contract method 0x6fcfff45.
//
// Solidity: function numCheckpoints(address ) view returns(uint32)
func (_GearToken *GearTokenCaller) NumCheckpoints(opts *bind.CallOpts, arg0 common.Address) (uint32, error) {
	var out []interface{}
	err := _GearToken.contract.Call(opts, &out, "numCheckpoints", arg0)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// NumCheckpoints is a free data retrieval call binding the contract method 0x6fcfff45.
//
// Solidity: function numCheckpoints(address ) view returns(uint32)
func (_GearToken *GearTokenSession) NumCheckpoints(arg0 common.Address) (uint32, error) {
	return _GearToken.Contract.NumCheckpoints(&_GearToken.CallOpts, arg0)
}

// NumCheckpoints is a free data retrieval call binding the contract method 0x6fcfff45.
//
// Solidity: function numCheckpoints(address ) view returns(uint32)
func (_GearToken *GearTokenCallerSession) NumCheckpoints(arg0 common.Address) (uint32, error) {
	return _GearToken.Contract.NumCheckpoints(&_GearToken.CallOpts, arg0)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_GearToken *GearTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _GearToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_GearToken *GearTokenSession) Symbol() (string, error) {
	return _GearToken.Contract.Symbol(&_GearToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_GearToken *GearTokenCallerSession) Symbol() (string, error) {
	return _GearToken.Contract.Symbol(&_GearToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_GearToken *GearTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GearToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_GearToken *GearTokenSession) TotalSupply() (*big.Int, error) {
	return _GearToken.Contract.TotalSupply(&_GearToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_GearToken *GearTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _GearToken.Contract.TotalSupply(&_GearToken.CallOpts)
}

// TransfersAllowed is a free data retrieval call binding the contract method 0xb0660c3d.
//
// Solidity: function transfersAllowed() view returns(bool)
func (_GearToken *GearTokenCaller) TransfersAllowed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _GearToken.contract.Call(opts, &out, "transfersAllowed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TransfersAllowed is a free data retrieval call binding the contract method 0xb0660c3d.
//
// Solidity: function transfersAllowed() view returns(bool)
func (_GearToken *GearTokenSession) TransfersAllowed() (bool, error) {
	return _GearToken.Contract.TransfersAllowed(&_GearToken.CallOpts)
}

// TransfersAllowed is a free data retrieval call binding the contract method 0xb0660c3d.
//
// Solidity: function transfersAllowed() view returns(bool)
func (_GearToken *GearTokenCallerSession) TransfersAllowed() (bool, error) {
	return _GearToken.Contract.TransfersAllowed(&_GearToken.CallOpts)
}

// AllowTransfers is a paid mutator transaction binding the contract method 0x2185810b.
//
// Solidity: function allowTransfers() returns()
func (_GearToken *GearTokenTransactor) AllowTransfers(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GearToken.contract.Transact(opts, "allowTransfers")
}

// AllowTransfers is a paid mutator transaction binding the contract method 0x2185810b.
//
// Solidity: function allowTransfers() returns()
func (_GearToken *GearTokenSession) AllowTransfers() (*types.Transaction, error) {
	return _GearToken.Contract.AllowTransfers(&_GearToken.TransactOpts)
}

// AllowTransfers is a paid mutator transaction binding the contract method 0x2185810b.
//
// Solidity: function allowTransfers() returns()
func (_GearToken *GearTokenTransactorSession) AllowTransfers() (*types.Transaction, error) {
	return _GearToken.Contract.AllowTransfers(&_GearToken.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 rawAmount) returns(bool)
func (_GearToken *GearTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, rawAmount *big.Int) (*types.Transaction, error) {
	return _GearToken.contract.Transact(opts, "approve", spender, rawAmount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 rawAmount) returns(bool)
func (_GearToken *GearTokenSession) Approve(spender common.Address, rawAmount *big.Int) (*types.Transaction, error) {
	return _GearToken.Contract.Approve(&_GearToken.TransactOpts, spender, rawAmount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 rawAmount) returns(bool)
func (_GearToken *GearTokenTransactorSession) Approve(spender common.Address, rawAmount *big.Int) (*types.Transaction, error) {
	return _GearToken.Contract.Approve(&_GearToken.TransactOpts, spender, rawAmount)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address delegatee) returns()
func (_GearToken *GearTokenTransactor) Delegate(opts *bind.TransactOpts, delegatee common.Address) (*types.Transaction, error) {
	return _GearToken.contract.Transact(opts, "delegate", delegatee)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address delegatee) returns()
func (_GearToken *GearTokenSession) Delegate(delegatee common.Address) (*types.Transaction, error) {
	return _GearToken.Contract.Delegate(&_GearToken.TransactOpts, delegatee)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address delegatee) returns()
func (_GearToken *GearTokenTransactorSession) Delegate(delegatee common.Address) (*types.Transaction, error) {
	return _GearToken.Contract.Delegate(&_GearToken.TransactOpts, delegatee)
}

// DelegateBySig is a paid mutator transaction binding the contract method 0xc3cda520.
//
// Solidity: function delegateBySig(address delegatee, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_GearToken *GearTokenTransactor) DelegateBySig(opts *bind.TransactOpts, delegatee common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _GearToken.contract.Transact(opts, "delegateBySig", delegatee, nonce, expiry, v, r, s)
}

// DelegateBySig is a paid mutator transaction binding the contract method 0xc3cda520.
//
// Solidity: function delegateBySig(address delegatee, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_GearToken *GearTokenSession) DelegateBySig(delegatee common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _GearToken.Contract.DelegateBySig(&_GearToken.TransactOpts, delegatee, nonce, expiry, v, r, s)
}

// DelegateBySig is a paid mutator transaction binding the contract method 0xc3cda520.
//
// Solidity: function delegateBySig(address delegatee, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_GearToken *GearTokenTransactorSession) DelegateBySig(delegatee common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _GearToken.Contract.DelegateBySig(&_GearToken.TransactOpts, delegatee, nonce, expiry, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 rawAmount, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_GearToken *GearTokenTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, rawAmount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _GearToken.contract.Transact(opts, "permit", owner, spender, rawAmount, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 rawAmount, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_GearToken *GearTokenSession) Permit(owner common.Address, spender common.Address, rawAmount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _GearToken.Contract.Permit(&_GearToken.TransactOpts, owner, spender, rawAmount, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 rawAmount, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_GearToken *GearTokenTransactorSession) Permit(owner common.Address, spender common.Address, rawAmount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _GearToken.Contract.Permit(&_GearToken.TransactOpts, owner, spender, rawAmount, deadline, v, r, s)
}

// SetMiner is a paid mutator transaction binding the contract method 0x9742ca46.
//
// Solidity: function setMiner(address _miner) returns()
func (_GearToken *GearTokenTransactor) SetMiner(opts *bind.TransactOpts, _miner common.Address) (*types.Transaction, error) {
	return _GearToken.contract.Transact(opts, "setMiner", _miner)
}

// SetMiner is a paid mutator transaction binding the contract method 0x9742ca46.
//
// Solidity: function setMiner(address _miner) returns()
func (_GearToken *GearTokenSession) SetMiner(_miner common.Address) (*types.Transaction, error) {
	return _GearToken.Contract.SetMiner(&_GearToken.TransactOpts, _miner)
}

// SetMiner is a paid mutator transaction binding the contract method 0x9742ca46.
//
// Solidity: function setMiner(address _miner) returns()
func (_GearToken *GearTokenTransactorSession) SetMiner(_miner common.Address) (*types.Transaction, error) {
	return _GearToken.Contract.SetMiner(&_GearToken.TransactOpts, _miner)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 rawAmount) returns(bool)
func (_GearToken *GearTokenTransactor) Transfer(opts *bind.TransactOpts, dst common.Address, rawAmount *big.Int) (*types.Transaction, error) {
	return _GearToken.contract.Transact(opts, "transfer", dst, rawAmount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 rawAmount) returns(bool)
func (_GearToken *GearTokenSession) Transfer(dst common.Address, rawAmount *big.Int) (*types.Transaction, error) {
	return _GearToken.Contract.Transfer(&_GearToken.TransactOpts, dst, rawAmount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 rawAmount) returns(bool)
func (_GearToken *GearTokenTransactorSession) Transfer(dst common.Address, rawAmount *big.Int) (*types.Transaction, error) {
	return _GearToken.Contract.Transfer(&_GearToken.TransactOpts, dst, rawAmount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 rawAmount) returns(bool)
func (_GearToken *GearTokenTransactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, rawAmount *big.Int) (*types.Transaction, error) {
	return _GearToken.contract.Transact(opts, "transferFrom", src, dst, rawAmount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 rawAmount) returns(bool)
func (_GearToken *GearTokenSession) TransferFrom(src common.Address, dst common.Address, rawAmount *big.Int) (*types.Transaction, error) {
	return _GearToken.Contract.TransferFrom(&_GearToken.TransactOpts, src, dst, rawAmount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 rawAmount) returns(bool)
func (_GearToken *GearTokenTransactorSession) TransferFrom(src common.Address, dst common.Address, rawAmount *big.Int) (*types.Transaction, error) {
	return _GearToken.Contract.TransferFrom(&_GearToken.TransactOpts, src, dst, rawAmount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newManager) returns()
func (_GearToken *GearTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newManager common.Address) (*types.Transaction, error) {
	return _GearToken.contract.Transact(opts, "transferOwnership", newManager)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newManager) returns()
func (_GearToken *GearTokenSession) TransferOwnership(newManager common.Address) (*types.Transaction, error) {
	return _GearToken.Contract.TransferOwnership(&_GearToken.TransactOpts, newManager)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newManager) returns()
func (_GearToken *GearTokenTransactorSession) TransferOwnership(newManager common.Address) (*types.Transaction, error) {
	return _GearToken.Contract.TransferOwnership(&_GearToken.TransactOpts, newManager)
}

// GearTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the GearToken contract.
type GearTokenApprovalIterator struct {
	Event *GearTokenApproval // Event containing the contract specifics and raw log

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
func (it *GearTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GearTokenApproval)
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
		it.Event = new(GearTokenApproval)
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
func (it *GearTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GearTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GearTokenApproval represents a Approval event raised by the GearToken contract.
type GearTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_GearToken *GearTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*GearTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _GearToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &GearTokenApprovalIterator{contract: _GearToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_GearToken *GearTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *GearTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _GearToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GearTokenApproval)
				if err := _GearToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_GearToken *GearTokenFilterer) ParseApproval(log types.Log) (*GearTokenApproval, error) {
	event := new(GearTokenApproval)
	if err := _GearToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GearTokenDelegateChangedIterator is returned from FilterDelegateChanged and is used to iterate over the raw logs and unpacked data for DelegateChanged events raised by the GearToken contract.
type GearTokenDelegateChangedIterator struct {
	Event *GearTokenDelegateChanged // Event containing the contract specifics and raw log

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
func (it *GearTokenDelegateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GearTokenDelegateChanged)
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
		it.Event = new(GearTokenDelegateChanged)
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
func (it *GearTokenDelegateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GearTokenDelegateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GearTokenDelegateChanged represents a DelegateChanged event raised by the GearToken contract.
type GearTokenDelegateChanged struct {
	Delegator    common.Address
	FromDelegate common.Address
	ToDelegate   common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDelegateChanged is a free log retrieval operation binding the contract event 0x3134e8a2e6d97e929a7e54011ea5485d7d196dd5f0ba4d4ef95803e8e3fc257f.
//
// Solidity: event DelegateChanged(address indexed delegator, address indexed fromDelegate, address indexed toDelegate)
func (_GearToken *GearTokenFilterer) FilterDelegateChanged(opts *bind.FilterOpts, delegator []common.Address, fromDelegate []common.Address, toDelegate []common.Address) (*GearTokenDelegateChangedIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var fromDelegateRule []interface{}
	for _, fromDelegateItem := range fromDelegate {
		fromDelegateRule = append(fromDelegateRule, fromDelegateItem)
	}
	var toDelegateRule []interface{}
	for _, toDelegateItem := range toDelegate {
		toDelegateRule = append(toDelegateRule, toDelegateItem)
	}

	logs, sub, err := _GearToken.contract.FilterLogs(opts, "DelegateChanged", delegatorRule, fromDelegateRule, toDelegateRule)
	if err != nil {
		return nil, err
	}
	return &GearTokenDelegateChangedIterator{contract: _GearToken.contract, event: "DelegateChanged", logs: logs, sub: sub}, nil
}

// WatchDelegateChanged is a free log subscription operation binding the contract event 0x3134e8a2e6d97e929a7e54011ea5485d7d196dd5f0ba4d4ef95803e8e3fc257f.
//
// Solidity: event DelegateChanged(address indexed delegator, address indexed fromDelegate, address indexed toDelegate)
func (_GearToken *GearTokenFilterer) WatchDelegateChanged(opts *bind.WatchOpts, sink chan<- *GearTokenDelegateChanged, delegator []common.Address, fromDelegate []common.Address, toDelegate []common.Address) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var fromDelegateRule []interface{}
	for _, fromDelegateItem := range fromDelegate {
		fromDelegateRule = append(fromDelegateRule, fromDelegateItem)
	}
	var toDelegateRule []interface{}
	for _, toDelegateItem := range toDelegate {
		toDelegateRule = append(toDelegateRule, toDelegateItem)
	}

	logs, sub, err := _GearToken.contract.WatchLogs(opts, "DelegateChanged", delegatorRule, fromDelegateRule, toDelegateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GearTokenDelegateChanged)
				if err := _GearToken.contract.UnpackLog(event, "DelegateChanged", log); err != nil {
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

// ParseDelegateChanged is a log parse operation binding the contract event 0x3134e8a2e6d97e929a7e54011ea5485d7d196dd5f0ba4d4ef95803e8e3fc257f.
//
// Solidity: event DelegateChanged(address indexed delegator, address indexed fromDelegate, address indexed toDelegate)
func (_GearToken *GearTokenFilterer) ParseDelegateChanged(log types.Log) (*GearTokenDelegateChanged, error) {
	event := new(GearTokenDelegateChanged)
	if err := _GearToken.contract.UnpackLog(event, "DelegateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GearTokenDelegateVotesChangedIterator is returned from FilterDelegateVotesChanged and is used to iterate over the raw logs and unpacked data for DelegateVotesChanged events raised by the GearToken contract.
type GearTokenDelegateVotesChangedIterator struct {
	Event *GearTokenDelegateVotesChanged // Event containing the contract specifics and raw log

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
func (it *GearTokenDelegateVotesChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GearTokenDelegateVotesChanged)
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
		it.Event = new(GearTokenDelegateVotesChanged)
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
func (it *GearTokenDelegateVotesChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GearTokenDelegateVotesChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GearTokenDelegateVotesChanged represents a DelegateVotesChanged event raised by the GearToken contract.
type GearTokenDelegateVotesChanged struct {
	Delegate        common.Address
	PreviousBalance *big.Int
	NewBalance      *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDelegateVotesChanged is a free log retrieval operation binding the contract event 0xdec2bacdd2f05b59de34da9b523dff8be42e5e38e818c82fdb0bae774387a724.
//
// Solidity: event DelegateVotesChanged(address indexed delegate, uint256 previousBalance, uint256 newBalance)
func (_GearToken *GearTokenFilterer) FilterDelegateVotesChanged(opts *bind.FilterOpts, delegate []common.Address) (*GearTokenDelegateVotesChangedIterator, error) {

	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}

	logs, sub, err := _GearToken.contract.FilterLogs(opts, "DelegateVotesChanged", delegateRule)
	if err != nil {
		return nil, err
	}
	return &GearTokenDelegateVotesChangedIterator{contract: _GearToken.contract, event: "DelegateVotesChanged", logs: logs, sub: sub}, nil
}

// WatchDelegateVotesChanged is a free log subscription operation binding the contract event 0xdec2bacdd2f05b59de34da9b523dff8be42e5e38e818c82fdb0bae774387a724.
//
// Solidity: event DelegateVotesChanged(address indexed delegate, uint256 previousBalance, uint256 newBalance)
func (_GearToken *GearTokenFilterer) WatchDelegateVotesChanged(opts *bind.WatchOpts, sink chan<- *GearTokenDelegateVotesChanged, delegate []common.Address) (event.Subscription, error) {

	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}

	logs, sub, err := _GearToken.contract.WatchLogs(opts, "DelegateVotesChanged", delegateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GearTokenDelegateVotesChanged)
				if err := _GearToken.contract.UnpackLog(event, "DelegateVotesChanged", log); err != nil {
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

// ParseDelegateVotesChanged is a log parse operation binding the contract event 0xdec2bacdd2f05b59de34da9b523dff8be42e5e38e818c82fdb0bae774387a724.
//
// Solidity: event DelegateVotesChanged(address indexed delegate, uint256 previousBalance, uint256 newBalance)
func (_GearToken *GearTokenFilterer) ParseDelegateVotesChanged(log types.Log) (*GearTokenDelegateVotesChanged, error) {
	event := new(GearTokenDelegateVotesChanged)
	if err := _GearToken.contract.UnpackLog(event, "DelegateVotesChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GearTokenMinerSetIterator is returned from FilterMinerSet and is used to iterate over the raw logs and unpacked data for MinerSet events raised by the GearToken contract.
type GearTokenMinerSetIterator struct {
	Event *GearTokenMinerSet // Event containing the contract specifics and raw log

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
func (it *GearTokenMinerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GearTokenMinerSet)
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
		it.Event = new(GearTokenMinerSet)
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
func (it *GearTokenMinerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GearTokenMinerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GearTokenMinerSet represents a MinerSet event raised by the GearToken contract.
type GearTokenMinerSet struct {
	Miner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMinerSet is a free log retrieval operation binding the contract event 0x2f834d1c8c4b956018fff5faca4d99868ae635487424d9c265c257ccbc698c6a.
//
// Solidity: event MinerSet(address indexed miner)
func (_GearToken *GearTokenFilterer) FilterMinerSet(opts *bind.FilterOpts, miner []common.Address) (*GearTokenMinerSetIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _GearToken.contract.FilterLogs(opts, "MinerSet", minerRule)
	if err != nil {
		return nil, err
	}
	return &GearTokenMinerSetIterator{contract: _GearToken.contract, event: "MinerSet", logs: logs, sub: sub}, nil
}

// WatchMinerSet is a free log subscription operation binding the contract event 0x2f834d1c8c4b956018fff5faca4d99868ae635487424d9c265c257ccbc698c6a.
//
// Solidity: event MinerSet(address indexed miner)
func (_GearToken *GearTokenFilterer) WatchMinerSet(opts *bind.WatchOpts, sink chan<- *GearTokenMinerSet, miner []common.Address) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _GearToken.contract.WatchLogs(opts, "MinerSet", minerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GearTokenMinerSet)
				if err := _GearToken.contract.UnpackLog(event, "MinerSet", log); err != nil {
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

// ParseMinerSet is a log parse operation binding the contract event 0x2f834d1c8c4b956018fff5faca4d99868ae635487424d9c265c257ccbc698c6a.
//
// Solidity: event MinerSet(address indexed miner)
func (_GearToken *GearTokenFilterer) ParseMinerSet(log types.Log) (*GearTokenMinerSet, error) {
	event := new(GearTokenMinerSet)
	if err := _GearToken.contract.UnpackLog(event, "MinerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GearTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the GearToken contract.
type GearTokenOwnershipTransferredIterator struct {
	Event *GearTokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *GearTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GearTokenOwnershipTransferred)
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
		it.Event = new(GearTokenOwnershipTransferred)
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
func (it *GearTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GearTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GearTokenOwnershipTransferred represents a OwnershipTransferred event raised by the GearToken contract.
type GearTokenOwnershipTransferred struct {
	Owner    common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed owner, address indexed newOwner)
func (_GearToken *GearTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, owner []common.Address, newOwner []common.Address) (*GearTokenOwnershipTransferredIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GearToken.contract.FilterLogs(opts, "OwnershipTransferred", ownerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &GearTokenOwnershipTransferredIterator{contract: _GearToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed owner, address indexed newOwner)
func (_GearToken *GearTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *GearTokenOwnershipTransferred, owner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GearToken.contract.WatchLogs(opts, "OwnershipTransferred", ownerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GearTokenOwnershipTransferred)
				if err := _GearToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
// Solidity: event OwnershipTransferred(address indexed owner, address indexed newOwner)
func (_GearToken *GearTokenFilterer) ParseOwnershipTransferred(log types.Log) (*GearTokenOwnershipTransferred, error) {
	event := new(GearTokenOwnershipTransferred)
	if err := _GearToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GearTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the GearToken contract.
type GearTokenTransferIterator struct {
	Event *GearTokenTransfer // Event containing the contract specifics and raw log

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
func (it *GearTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GearTokenTransfer)
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
		it.Event = new(GearTokenTransfer)
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
func (it *GearTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GearTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GearTokenTransfer represents a Transfer event raised by the GearToken contract.
type GearTokenTransfer struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_GearToken *GearTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*GearTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GearToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &GearTokenTransferIterator{contract: _GearToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_GearToken *GearTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *GearTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GearToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GearTokenTransfer)
				if err := _GearToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_GearToken *GearTokenFilterer) ParseTransfer(log types.Log) (*GearTokenTransfer, error) {
	event := new(GearTokenTransfer)
	if err := _GearToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GearTokenTransferAllowedIterator is returned from FilterTransferAllowed and is used to iterate over the raw logs and unpacked data for TransferAllowed events raised by the GearToken contract.
type GearTokenTransferAllowedIterator struct {
	Event *GearTokenTransferAllowed // Event containing the contract specifics and raw log

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
func (it *GearTokenTransferAllowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GearTokenTransferAllowed)
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
		it.Event = new(GearTokenTransferAllowed)
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
func (it *GearTokenTransferAllowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GearTokenTransferAllowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GearTokenTransferAllowed represents a TransferAllowed event raised by the GearToken contract.
type GearTokenTransferAllowed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterTransferAllowed is a free log retrieval operation binding the contract event 0x795b0e16c8da9807b0a215f3749bd6dbcc49fc0472183f4e446abb7dcbd9d007.
//
// Solidity: event TransferAllowed()
func (_GearToken *GearTokenFilterer) FilterTransferAllowed(opts *bind.FilterOpts) (*GearTokenTransferAllowedIterator, error) {

	logs, sub, err := _GearToken.contract.FilterLogs(opts, "TransferAllowed")
	if err != nil {
		return nil, err
	}
	return &GearTokenTransferAllowedIterator{contract: _GearToken.contract, event: "TransferAllowed", logs: logs, sub: sub}, nil
}

// WatchTransferAllowed is a free log subscription operation binding the contract event 0x795b0e16c8da9807b0a215f3749bd6dbcc49fc0472183f4e446abb7dcbd9d007.
//
// Solidity: event TransferAllowed()
func (_GearToken *GearTokenFilterer) WatchTransferAllowed(opts *bind.WatchOpts, sink chan<- *GearTokenTransferAllowed) (event.Subscription, error) {

	logs, sub, err := _GearToken.contract.WatchLogs(opts, "TransferAllowed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GearTokenTransferAllowed)
				if err := _GearToken.contract.UnpackLog(event, "TransferAllowed", log); err != nil {
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

// ParseTransferAllowed is a log parse operation binding the contract event 0x795b0e16c8da9807b0a215f3749bd6dbcc49fc0472183f4e446abb7dcbd9d007.
//
// Solidity: event TransferAllowed()
func (_GearToken *GearTokenFilterer) ParseTransferAllowed(log types.Log) (*GearTokenTransferAllowed, error) {
	event := new(GearTokenTransferAllowed)
	if err := _GearToken.contract.UnpackLog(event, "TransferAllowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
