// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractsRegister

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

// ContractsRegisterMetaData contains all meta data concerning the ContractsRegister contract.
var ContractsRegisterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addressProvider\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creditManager\",\"type\":\"address\"}],\"name\":\"NewCreditManagerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"NewPoolAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newCreditManager\",\"type\":\"address\"}],\"name\":\"addCreditManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPoolAddress\",\"type\":\"address\"}],\"name\":\"addPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"creditManagers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCreditManagers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCreditManagersCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPools\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isCreditManager\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isPool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pools\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ContractsRegisterABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractsRegisterMetaData.ABI instead.
var ContractsRegisterABI = ContractsRegisterMetaData.ABI

// ContractsRegister is an auto generated Go binding around an Ethereum contract.
type ContractsRegister struct {
	ContractsRegisterCaller     // Read-only binding to the contract
	ContractsRegisterTransactor // Write-only binding to the contract
	ContractsRegisterFilterer   // Log filterer for contract events
}

// ContractsRegisterCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractsRegisterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsRegisterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractsRegisterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsRegisterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractsRegisterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsRegisterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractsRegisterSession struct {
	Contract     *ContractsRegister // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ContractsRegisterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractsRegisterCallerSession struct {
	Contract *ContractsRegisterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ContractsRegisterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractsRegisterTransactorSession struct {
	Contract     *ContractsRegisterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ContractsRegisterRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractsRegisterRaw struct {
	Contract *ContractsRegister // Generic contract binding to access the raw methods on
}

// ContractsRegisterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractsRegisterCallerRaw struct {
	Contract *ContractsRegisterCaller // Generic read-only contract binding to access the raw methods on
}

// ContractsRegisterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractsRegisterTransactorRaw struct {
	Contract *ContractsRegisterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractsRegister creates a new instance of ContractsRegister, bound to a specific deployed contract.
func NewContractsRegister(address common.Address, backend bind.ContractBackend) (*ContractsRegister, error) {
	contract, err := bindContractsRegister(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractsRegister{ContractsRegisterCaller: ContractsRegisterCaller{contract: contract}, ContractsRegisterTransactor: ContractsRegisterTransactor{contract: contract}, ContractsRegisterFilterer: ContractsRegisterFilterer{contract: contract}}, nil
}

// NewContractsRegisterCaller creates a new read-only instance of ContractsRegister, bound to a specific deployed contract.
func NewContractsRegisterCaller(address common.Address, caller bind.ContractCaller) (*ContractsRegisterCaller, error) {
	contract, err := bindContractsRegister(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsRegisterCaller{contract: contract}, nil
}

// NewContractsRegisterTransactor creates a new write-only instance of ContractsRegister, bound to a specific deployed contract.
func NewContractsRegisterTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractsRegisterTransactor, error) {
	contract, err := bindContractsRegister(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsRegisterTransactor{contract: contract}, nil
}

// NewContractsRegisterFilterer creates a new log filterer instance of ContractsRegister, bound to a specific deployed contract.
func NewContractsRegisterFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractsRegisterFilterer, error) {
	contract, err := bindContractsRegister(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractsRegisterFilterer{contract: contract}, nil
}

// bindContractsRegister binds a generic wrapper to an already deployed contract.
func bindContractsRegister(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractsRegisterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractsRegister *ContractsRegisterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractsRegister.Contract.ContractsRegisterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractsRegister *ContractsRegisterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractsRegister.Contract.ContractsRegisterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractsRegister *ContractsRegisterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractsRegister.Contract.ContractsRegisterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractsRegister *ContractsRegisterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractsRegister.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractsRegister *ContractsRegisterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractsRegister.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractsRegister *ContractsRegisterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractsRegister.Contract.contract.Transact(opts, method, params...)
}

// CreditManagers is a free data retrieval call binding the contract method 0x1e16e4fc.
//
// Solidity: function creditManagers(uint256 ) view returns(address)
func (_ContractsRegister *ContractsRegisterCaller) CreditManagers(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ContractsRegister.contract.Call(opts, &out, "creditManagers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditManagers is a free data retrieval call binding the contract method 0x1e16e4fc.
//
// Solidity: function creditManagers(uint256 ) view returns(address)
func (_ContractsRegister *ContractsRegisterSession) CreditManagers(arg0 *big.Int) (common.Address, error) {
	return _ContractsRegister.Contract.CreditManagers(&_ContractsRegister.CallOpts, arg0)
}

// CreditManagers is a free data retrieval call binding the contract method 0x1e16e4fc.
//
// Solidity: function creditManagers(uint256 ) view returns(address)
func (_ContractsRegister *ContractsRegisterCallerSession) CreditManagers(arg0 *big.Int) (common.Address, error) {
	return _ContractsRegister.Contract.CreditManagers(&_ContractsRegister.CallOpts, arg0)
}

// GetCreditManagers is a free data retrieval call binding the contract method 0x94144856.
//
// Solidity: function getCreditManagers() view returns(address[])
func (_ContractsRegister *ContractsRegisterCaller) GetCreditManagers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ContractsRegister.contract.Call(opts, &out, "getCreditManagers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetCreditManagers is a free data retrieval call binding the contract method 0x94144856.
//
// Solidity: function getCreditManagers() view returns(address[])
func (_ContractsRegister *ContractsRegisterSession) GetCreditManagers() ([]common.Address, error) {
	return _ContractsRegister.Contract.GetCreditManagers(&_ContractsRegister.CallOpts)
}

// GetCreditManagers is a free data retrieval call binding the contract method 0x94144856.
//
// Solidity: function getCreditManagers() view returns(address[])
func (_ContractsRegister *ContractsRegisterCallerSession) GetCreditManagers() ([]common.Address, error) {
	return _ContractsRegister.Contract.GetCreditManagers(&_ContractsRegister.CallOpts)
}

// GetCreditManagersCount is a free data retrieval call binding the contract method 0xc29277cd.
//
// Solidity: function getCreditManagersCount() view returns(uint256)
func (_ContractsRegister *ContractsRegisterCaller) GetCreditManagersCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractsRegister.contract.Call(opts, &out, "getCreditManagersCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCreditManagersCount is a free data retrieval call binding the contract method 0xc29277cd.
//
// Solidity: function getCreditManagersCount() view returns(uint256)
func (_ContractsRegister *ContractsRegisterSession) GetCreditManagersCount() (*big.Int, error) {
	return _ContractsRegister.Contract.GetCreditManagersCount(&_ContractsRegister.CallOpts)
}

// GetCreditManagersCount is a free data retrieval call binding the contract method 0xc29277cd.
//
// Solidity: function getCreditManagersCount() view returns(uint256)
func (_ContractsRegister *ContractsRegisterCallerSession) GetCreditManagersCount() (*big.Int, error) {
	return _ContractsRegister.Contract.GetCreditManagersCount(&_ContractsRegister.CallOpts)
}

// GetPools is a free data retrieval call binding the contract method 0x673a2a1f.
//
// Solidity: function getPools() view returns(address[])
func (_ContractsRegister *ContractsRegisterCaller) GetPools(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ContractsRegister.contract.Call(opts, &out, "getPools")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetPools is a free data retrieval call binding the contract method 0x673a2a1f.
//
// Solidity: function getPools() view returns(address[])
func (_ContractsRegister *ContractsRegisterSession) GetPools() ([]common.Address, error) {
	return _ContractsRegister.Contract.GetPools(&_ContractsRegister.CallOpts)
}

// GetPools is a free data retrieval call binding the contract method 0x673a2a1f.
//
// Solidity: function getPools() view returns(address[])
func (_ContractsRegister *ContractsRegisterCallerSession) GetPools() ([]common.Address, error) {
	return _ContractsRegister.Contract.GetPools(&_ContractsRegister.CallOpts)
}

// GetPoolsCount is a free data retrieval call binding the contract method 0xb4ac6860.
//
// Solidity: function getPoolsCount() view returns(uint256)
func (_ContractsRegister *ContractsRegisterCaller) GetPoolsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractsRegister.contract.Call(opts, &out, "getPoolsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPoolsCount is a free data retrieval call binding the contract method 0xb4ac6860.
//
// Solidity: function getPoolsCount() view returns(uint256)
func (_ContractsRegister *ContractsRegisterSession) GetPoolsCount() (*big.Int, error) {
	return _ContractsRegister.Contract.GetPoolsCount(&_ContractsRegister.CallOpts)
}

// GetPoolsCount is a free data retrieval call binding the contract method 0xb4ac6860.
//
// Solidity: function getPoolsCount() view returns(uint256)
func (_ContractsRegister *ContractsRegisterCallerSession) GetPoolsCount() (*big.Int, error) {
	return _ContractsRegister.Contract.GetPoolsCount(&_ContractsRegister.CallOpts)
}

// IsCreditManager is a free data retrieval call binding the contract method 0x6fbc6f6b.
//
// Solidity: function isCreditManager(address ) view returns(bool)
func (_ContractsRegister *ContractsRegisterCaller) IsCreditManager(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ContractsRegister.contract.Call(opts, &out, "isCreditManager", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCreditManager is a free data retrieval call binding the contract method 0x6fbc6f6b.
//
// Solidity: function isCreditManager(address ) view returns(bool)
func (_ContractsRegister *ContractsRegisterSession) IsCreditManager(arg0 common.Address) (bool, error) {
	return _ContractsRegister.Contract.IsCreditManager(&_ContractsRegister.CallOpts, arg0)
}

// IsCreditManager is a free data retrieval call binding the contract method 0x6fbc6f6b.
//
// Solidity: function isCreditManager(address ) view returns(bool)
func (_ContractsRegister *ContractsRegisterCallerSession) IsCreditManager(arg0 common.Address) (bool, error) {
	return _ContractsRegister.Contract.IsCreditManager(&_ContractsRegister.CallOpts, arg0)
}

// IsPool is a free data retrieval call binding the contract method 0x5b16ebb7.
//
// Solidity: function isPool(address ) view returns(bool)
func (_ContractsRegister *ContractsRegisterCaller) IsPool(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ContractsRegister.contract.Call(opts, &out, "isPool", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPool is a free data retrieval call binding the contract method 0x5b16ebb7.
//
// Solidity: function isPool(address ) view returns(bool)
func (_ContractsRegister *ContractsRegisterSession) IsPool(arg0 common.Address) (bool, error) {
	return _ContractsRegister.Contract.IsPool(&_ContractsRegister.CallOpts, arg0)
}

// IsPool is a free data retrieval call binding the contract method 0x5b16ebb7.
//
// Solidity: function isPool(address ) view returns(bool)
func (_ContractsRegister *ContractsRegisterCallerSession) IsPool(arg0 common.Address) (bool, error) {
	return _ContractsRegister.Contract.IsPool(&_ContractsRegister.CallOpts, arg0)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ContractsRegister *ContractsRegisterCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ContractsRegister.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ContractsRegister *ContractsRegisterSession) Paused() (bool, error) {
	return _ContractsRegister.Contract.Paused(&_ContractsRegister.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ContractsRegister *ContractsRegisterCallerSession) Paused() (bool, error) {
	return _ContractsRegister.Contract.Paused(&_ContractsRegister.CallOpts)
}

// Pools is a free data retrieval call binding the contract method 0xac4afa38.
//
// Solidity: function pools(uint256 ) view returns(address)
func (_ContractsRegister *ContractsRegisterCaller) Pools(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ContractsRegister.contract.Call(opts, &out, "pools", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pools is a free data retrieval call binding the contract method 0xac4afa38.
//
// Solidity: function pools(uint256 ) view returns(address)
func (_ContractsRegister *ContractsRegisterSession) Pools(arg0 *big.Int) (common.Address, error) {
	return _ContractsRegister.Contract.Pools(&_ContractsRegister.CallOpts, arg0)
}

// Pools is a free data retrieval call binding the contract method 0xac4afa38.
//
// Solidity: function pools(uint256 ) view returns(address)
func (_ContractsRegister *ContractsRegisterCallerSession) Pools(arg0 *big.Int) (common.Address, error) {
	return _ContractsRegister.Contract.Pools(&_ContractsRegister.CallOpts, arg0)
}

// AddCreditManager is a paid mutator transaction binding the contract method 0xe26b2f63.
//
// Solidity: function addCreditManager(address newCreditManager) returns()
func (_ContractsRegister *ContractsRegisterTransactor) AddCreditManager(opts *bind.TransactOpts, newCreditManager common.Address) (*types.Transaction, error) {
	return _ContractsRegister.contract.Transact(opts, "addCreditManager", newCreditManager)
}

// AddCreditManager is a paid mutator transaction binding the contract method 0xe26b2f63.
//
// Solidity: function addCreditManager(address newCreditManager) returns()
func (_ContractsRegister *ContractsRegisterSession) AddCreditManager(newCreditManager common.Address) (*types.Transaction, error) {
	return _ContractsRegister.Contract.AddCreditManager(&_ContractsRegister.TransactOpts, newCreditManager)
}

// AddCreditManager is a paid mutator transaction binding the contract method 0xe26b2f63.
//
// Solidity: function addCreditManager(address newCreditManager) returns()
func (_ContractsRegister *ContractsRegisterTransactorSession) AddCreditManager(newCreditManager common.Address) (*types.Transaction, error) {
	return _ContractsRegister.Contract.AddCreditManager(&_ContractsRegister.TransactOpts, newCreditManager)
}

// AddPool is a paid mutator transaction binding the contract method 0xd914cd4b.
//
// Solidity: function addPool(address newPoolAddress) returns()
func (_ContractsRegister *ContractsRegisterTransactor) AddPool(opts *bind.TransactOpts, newPoolAddress common.Address) (*types.Transaction, error) {
	return _ContractsRegister.contract.Transact(opts, "addPool", newPoolAddress)
}

// AddPool is a paid mutator transaction binding the contract method 0xd914cd4b.
//
// Solidity: function addPool(address newPoolAddress) returns()
func (_ContractsRegister *ContractsRegisterSession) AddPool(newPoolAddress common.Address) (*types.Transaction, error) {
	return _ContractsRegister.Contract.AddPool(&_ContractsRegister.TransactOpts, newPoolAddress)
}

// AddPool is a paid mutator transaction binding the contract method 0xd914cd4b.
//
// Solidity: function addPool(address newPoolAddress) returns()
func (_ContractsRegister *ContractsRegisterTransactorSession) AddPool(newPoolAddress common.Address) (*types.Transaction, error) {
	return _ContractsRegister.Contract.AddPool(&_ContractsRegister.TransactOpts, newPoolAddress)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ContractsRegister *ContractsRegisterTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractsRegister.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ContractsRegister *ContractsRegisterSession) Pause() (*types.Transaction, error) {
	return _ContractsRegister.Contract.Pause(&_ContractsRegister.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ContractsRegister *ContractsRegisterTransactorSession) Pause() (*types.Transaction, error) {
	return _ContractsRegister.Contract.Pause(&_ContractsRegister.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ContractsRegister *ContractsRegisterTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractsRegister.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ContractsRegister *ContractsRegisterSession) Unpause() (*types.Transaction, error) {
	return _ContractsRegister.Contract.Unpause(&_ContractsRegister.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ContractsRegister *ContractsRegisterTransactorSession) Unpause() (*types.Transaction, error) {
	return _ContractsRegister.Contract.Unpause(&_ContractsRegister.TransactOpts)
}

// ContractsRegisterNewCreditManagerAddedIterator is returned from FilterNewCreditManagerAdded and is used to iterate over the raw logs and unpacked data for NewCreditManagerAdded events raised by the ContractsRegister contract.
type ContractsRegisterNewCreditManagerAddedIterator struct {
	Event *ContractsRegisterNewCreditManagerAdded // Event containing the contract specifics and raw log

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
func (it *ContractsRegisterNewCreditManagerAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsRegisterNewCreditManagerAdded)
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
		it.Event = new(ContractsRegisterNewCreditManagerAdded)
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
func (it *ContractsRegisterNewCreditManagerAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsRegisterNewCreditManagerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsRegisterNewCreditManagerAdded represents a NewCreditManagerAdded event raised by the ContractsRegister contract.
type ContractsRegisterNewCreditManagerAdded struct {
	CreditManager common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNewCreditManagerAdded is a free log retrieval operation binding the contract event 0x58ad3cfc4b6552a53c8c4128ae9b080e14b4378a159280643a62c6f709cee24f.
//
// Solidity: event NewCreditManagerAdded(address indexed creditManager)
func (_ContractsRegister *ContractsRegisterFilterer) FilterNewCreditManagerAdded(opts *bind.FilterOpts, creditManager []common.Address) (*ContractsRegisterNewCreditManagerAddedIterator, error) {

	var creditManagerRule []interface{}
	for _, creditManagerItem := range creditManager {
		creditManagerRule = append(creditManagerRule, creditManagerItem)
	}

	logs, sub, err := _ContractsRegister.contract.FilterLogs(opts, "NewCreditManagerAdded", creditManagerRule)
	if err != nil {
		return nil, err
	}
	return &ContractsRegisterNewCreditManagerAddedIterator{contract: _ContractsRegister.contract, event: "NewCreditManagerAdded", logs: logs, sub: sub}, nil
}

// WatchNewCreditManagerAdded is a free log subscription operation binding the contract event 0x58ad3cfc4b6552a53c8c4128ae9b080e14b4378a159280643a62c6f709cee24f.
//
// Solidity: event NewCreditManagerAdded(address indexed creditManager)
func (_ContractsRegister *ContractsRegisterFilterer) WatchNewCreditManagerAdded(opts *bind.WatchOpts, sink chan<- *ContractsRegisterNewCreditManagerAdded, creditManager []common.Address) (event.Subscription, error) {

	var creditManagerRule []interface{}
	for _, creditManagerItem := range creditManager {
		creditManagerRule = append(creditManagerRule, creditManagerItem)
	}

	logs, sub, err := _ContractsRegister.contract.WatchLogs(opts, "NewCreditManagerAdded", creditManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsRegisterNewCreditManagerAdded)
				if err := _ContractsRegister.contract.UnpackLog(event, "NewCreditManagerAdded", log); err != nil {
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

// ParseNewCreditManagerAdded is a log parse operation binding the contract event 0x58ad3cfc4b6552a53c8c4128ae9b080e14b4378a159280643a62c6f709cee24f.
//
// Solidity: event NewCreditManagerAdded(address indexed creditManager)
func (_ContractsRegister *ContractsRegisterFilterer) ParseNewCreditManagerAdded(log types.Log) (*ContractsRegisterNewCreditManagerAdded, error) {
	event := new(ContractsRegisterNewCreditManagerAdded)
	if err := _ContractsRegister.contract.UnpackLog(event, "NewCreditManagerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsRegisterNewPoolAddedIterator is returned from FilterNewPoolAdded and is used to iterate over the raw logs and unpacked data for NewPoolAdded events raised by the ContractsRegister contract.
type ContractsRegisterNewPoolAddedIterator struct {
	Event *ContractsRegisterNewPoolAdded // Event containing the contract specifics and raw log

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
func (it *ContractsRegisterNewPoolAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsRegisterNewPoolAdded)
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
		it.Event = new(ContractsRegisterNewPoolAdded)
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
func (it *ContractsRegisterNewPoolAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsRegisterNewPoolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsRegisterNewPoolAdded represents a NewPoolAdded event raised by the ContractsRegister contract.
type ContractsRegisterNewPoolAdded struct {
	Pool common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterNewPoolAdded is a free log retrieval operation binding the contract event 0xf816b5143086c89d103a0683286be86c2b741e83ebfa75135aae606e2f5c6e53.
//
// Solidity: event NewPoolAdded(address indexed pool)
func (_ContractsRegister *ContractsRegisterFilterer) FilterNewPoolAdded(opts *bind.FilterOpts, pool []common.Address) (*ContractsRegisterNewPoolAddedIterator, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _ContractsRegister.contract.FilterLogs(opts, "NewPoolAdded", poolRule)
	if err != nil {
		return nil, err
	}
	return &ContractsRegisterNewPoolAddedIterator{contract: _ContractsRegister.contract, event: "NewPoolAdded", logs: logs, sub: sub}, nil
}

// WatchNewPoolAdded is a free log subscription operation binding the contract event 0xf816b5143086c89d103a0683286be86c2b741e83ebfa75135aae606e2f5c6e53.
//
// Solidity: event NewPoolAdded(address indexed pool)
func (_ContractsRegister *ContractsRegisterFilterer) WatchNewPoolAdded(opts *bind.WatchOpts, sink chan<- *ContractsRegisterNewPoolAdded, pool []common.Address) (event.Subscription, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _ContractsRegister.contract.WatchLogs(opts, "NewPoolAdded", poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsRegisterNewPoolAdded)
				if err := _ContractsRegister.contract.UnpackLog(event, "NewPoolAdded", log); err != nil {
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

// ParseNewPoolAdded is a log parse operation binding the contract event 0xf816b5143086c89d103a0683286be86c2b741e83ebfa75135aae606e2f5c6e53.
//
// Solidity: event NewPoolAdded(address indexed pool)
func (_ContractsRegister *ContractsRegisterFilterer) ParseNewPoolAdded(log types.Log) (*ContractsRegisterNewPoolAdded, error) {
	event := new(ContractsRegisterNewPoolAdded)
	if err := _ContractsRegister.contract.UnpackLog(event, "NewPoolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsRegisterPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ContractsRegister contract.
type ContractsRegisterPausedIterator struct {
	Event *ContractsRegisterPaused // Event containing the contract specifics and raw log

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
func (it *ContractsRegisterPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsRegisterPaused)
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
		it.Event = new(ContractsRegisterPaused)
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
func (it *ContractsRegisterPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsRegisterPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsRegisterPaused represents a Paused event raised by the ContractsRegister contract.
type ContractsRegisterPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ContractsRegister *ContractsRegisterFilterer) FilterPaused(opts *bind.FilterOpts) (*ContractsRegisterPausedIterator, error) {

	logs, sub, err := _ContractsRegister.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &ContractsRegisterPausedIterator{contract: _ContractsRegister.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ContractsRegister *ContractsRegisterFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractsRegisterPaused) (event.Subscription, error) {

	logs, sub, err := _ContractsRegister.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsRegisterPaused)
				if err := _ContractsRegister.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_ContractsRegister *ContractsRegisterFilterer) ParsePaused(log types.Log) (*ContractsRegisterPaused, error) {
	event := new(ContractsRegisterPaused)
	if err := _ContractsRegister.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsRegisterUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ContractsRegister contract.
type ContractsRegisterUnpausedIterator struct {
	Event *ContractsRegisterUnpaused // Event containing the contract specifics and raw log

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
func (it *ContractsRegisterUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsRegisterUnpaused)
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
		it.Event = new(ContractsRegisterUnpaused)
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
func (it *ContractsRegisterUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsRegisterUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsRegisterUnpaused represents a Unpaused event raised by the ContractsRegister contract.
type ContractsRegisterUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ContractsRegister *ContractsRegisterFilterer) FilterUnpaused(opts *bind.FilterOpts) (*ContractsRegisterUnpausedIterator, error) {

	logs, sub, err := _ContractsRegister.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &ContractsRegisterUnpausedIterator{contract: _ContractsRegister.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ContractsRegister *ContractsRegisterFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractsRegisterUnpaused) (event.Subscription, error) {

	logs, sub, err := _ContractsRegister.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsRegisterUnpaused)
				if err := _ContractsRegister.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_ContractsRegister *ContractsRegisterFilterer) ParseUnpaused(log types.Log) (*ContractsRegisterUnpaused, error) {
	event := new(ContractsRegisterUnpaused)
	if err := _ContractsRegister.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
