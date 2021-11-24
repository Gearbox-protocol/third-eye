// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package wadRayMathTest

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

// WadRayMathTestMetaData contains all meta data concerning the WadRayMathTest contract.
var WadRayMathTestMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"halfRay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"halfWad\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ray\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"rayDiv\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"rayMul\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"rayToWad\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wad\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"wadDiv\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"wadMul\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"wadToRay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// WadRayMathTestABI is the input ABI used to generate the binding from.
// Deprecated: Use WadRayMathTestMetaData.ABI instead.
var WadRayMathTestABI = WadRayMathTestMetaData.ABI

// WadRayMathTest is an auto generated Go binding around an Ethereum contract.
type WadRayMathTest struct {
	WadRayMathTestCaller     // Read-only binding to the contract
	WadRayMathTestTransactor // Write-only binding to the contract
	WadRayMathTestFilterer   // Log filterer for contract events
}

// WadRayMathTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type WadRayMathTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WadRayMathTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WadRayMathTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WadRayMathTestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WadRayMathTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WadRayMathTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WadRayMathTestSession struct {
	Contract     *WadRayMathTest   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WadRayMathTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WadRayMathTestCallerSession struct {
	Contract *WadRayMathTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// WadRayMathTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WadRayMathTestTransactorSession struct {
	Contract     *WadRayMathTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// WadRayMathTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type WadRayMathTestRaw struct {
	Contract *WadRayMathTest // Generic contract binding to access the raw methods on
}

// WadRayMathTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WadRayMathTestCallerRaw struct {
	Contract *WadRayMathTestCaller // Generic read-only contract binding to access the raw methods on
}

// WadRayMathTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WadRayMathTestTransactorRaw struct {
	Contract *WadRayMathTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWadRayMathTest creates a new instance of WadRayMathTest, bound to a specific deployed contract.
func NewWadRayMathTest(address common.Address, backend bind.ContractBackend) (*WadRayMathTest, error) {
	contract, err := bindWadRayMathTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WadRayMathTest{WadRayMathTestCaller: WadRayMathTestCaller{contract: contract}, WadRayMathTestTransactor: WadRayMathTestTransactor{contract: contract}, WadRayMathTestFilterer: WadRayMathTestFilterer{contract: contract}}, nil
}

// NewWadRayMathTestCaller creates a new read-only instance of WadRayMathTest, bound to a specific deployed contract.
func NewWadRayMathTestCaller(address common.Address, caller bind.ContractCaller) (*WadRayMathTestCaller, error) {
	contract, err := bindWadRayMathTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WadRayMathTestCaller{contract: contract}, nil
}

// NewWadRayMathTestTransactor creates a new write-only instance of WadRayMathTest, bound to a specific deployed contract.
func NewWadRayMathTestTransactor(address common.Address, transactor bind.ContractTransactor) (*WadRayMathTestTransactor, error) {
	contract, err := bindWadRayMathTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WadRayMathTestTransactor{contract: contract}, nil
}

// NewWadRayMathTestFilterer creates a new log filterer instance of WadRayMathTest, bound to a specific deployed contract.
func NewWadRayMathTestFilterer(address common.Address, filterer bind.ContractFilterer) (*WadRayMathTestFilterer, error) {
	contract, err := bindWadRayMathTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WadRayMathTestFilterer{contract: contract}, nil
}

// bindWadRayMathTest binds a generic wrapper to an already deployed contract.
func bindWadRayMathTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WadRayMathTestABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WadRayMathTest *WadRayMathTestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WadRayMathTest.Contract.WadRayMathTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WadRayMathTest *WadRayMathTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WadRayMathTest.Contract.WadRayMathTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WadRayMathTest *WadRayMathTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WadRayMathTest.Contract.WadRayMathTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WadRayMathTest *WadRayMathTestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WadRayMathTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WadRayMathTest *WadRayMathTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WadRayMathTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WadRayMathTest *WadRayMathTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WadRayMathTest.Contract.contract.Transact(opts, method, params...)
}

// HalfRay is a free data retrieval call binding the contract method 0x1fa89fc6.
//
// Solidity: function halfRay() pure returns(uint256)
func (_WadRayMathTest *WadRayMathTestCaller) HalfRay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WadRayMathTest.contract.Call(opts, &out, "halfRay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HalfRay is a free data retrieval call binding the contract method 0x1fa89fc6.
//
// Solidity: function halfRay() pure returns(uint256)
func (_WadRayMathTest *WadRayMathTestSession) HalfRay() (*big.Int, error) {
	return _WadRayMathTest.Contract.HalfRay(&_WadRayMathTest.CallOpts)
}

// HalfRay is a free data retrieval call binding the contract method 0x1fa89fc6.
//
// Solidity: function halfRay() pure returns(uint256)
func (_WadRayMathTest *WadRayMathTestCallerSession) HalfRay() (*big.Int, error) {
	return _WadRayMathTest.Contract.HalfRay(&_WadRayMathTest.CallOpts)
}

// HalfWad is a free data retrieval call binding the contract method 0xe304e1d3.
//
// Solidity: function halfWad() pure returns(uint256)
func (_WadRayMathTest *WadRayMathTestCaller) HalfWad(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WadRayMathTest.contract.Call(opts, &out, "halfWad")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HalfWad is a free data retrieval call binding the contract method 0xe304e1d3.
//
// Solidity: function halfWad() pure returns(uint256)
func (_WadRayMathTest *WadRayMathTestSession) HalfWad() (*big.Int, error) {
	return _WadRayMathTest.Contract.HalfWad(&_WadRayMathTest.CallOpts)
}

// HalfWad is a free data retrieval call binding the contract method 0xe304e1d3.
//
// Solidity: function halfWad() pure returns(uint256)
func (_WadRayMathTest *WadRayMathTestCallerSession) HalfWad() (*big.Int, error) {
	return _WadRayMathTest.Contract.HalfWad(&_WadRayMathTest.CallOpts)
}

// Ray is a free data retrieval call binding the contract method 0x416a8b20.
//
// Solidity: function ray() pure returns(uint256)
func (_WadRayMathTest *WadRayMathTestCaller) Ray(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WadRayMathTest.contract.Call(opts, &out, "ray")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Ray is a free data retrieval call binding the contract method 0x416a8b20.
//
// Solidity: function ray() pure returns(uint256)
func (_WadRayMathTest *WadRayMathTestSession) Ray() (*big.Int, error) {
	return _WadRayMathTest.Contract.Ray(&_WadRayMathTest.CallOpts)
}

// Ray is a free data retrieval call binding the contract method 0x416a8b20.
//
// Solidity: function ray() pure returns(uint256)
func (_WadRayMathTest *WadRayMathTestCallerSession) Ray() (*big.Int, error) {
	return _WadRayMathTest.Contract.Ray(&_WadRayMathTest.CallOpts)
}

// RayDiv is a free data retrieval call binding the contract method 0x9c34d880.
//
// Solidity: function rayDiv(uint256 a, uint256 b) pure returns(uint256)
func (_WadRayMathTest *WadRayMathTestCaller) RayDiv(opts *bind.CallOpts, a *big.Int, b *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _WadRayMathTest.contract.Call(opts, &out, "rayDiv", a, b)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RayDiv is a free data retrieval call binding the contract method 0x9c34d880.
//
// Solidity: function rayDiv(uint256 a, uint256 b) pure returns(uint256)
func (_WadRayMathTest *WadRayMathTestSession) RayDiv(a *big.Int, b *big.Int) (*big.Int, error) {
	return _WadRayMathTest.Contract.RayDiv(&_WadRayMathTest.CallOpts, a, b)
}

// RayDiv is a free data retrieval call binding the contract method 0x9c34d880.
//
// Solidity: function rayDiv(uint256 a, uint256 b) pure returns(uint256)
func (_WadRayMathTest *WadRayMathTestCallerSession) RayDiv(a *big.Int, b *big.Int) (*big.Int, error) {
	return _WadRayMathTest.Contract.RayDiv(&_WadRayMathTest.CallOpts, a, b)
}

// RayMul is a free data retrieval call binding the contract method 0xd2e30585.
//
// Solidity: function rayMul(uint256 a, uint256 b) pure returns(uint256)
func (_WadRayMathTest *WadRayMathTestCaller) RayMul(opts *bind.CallOpts, a *big.Int, b *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _WadRayMathTest.contract.Call(opts, &out, "rayMul", a, b)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RayMul is a free data retrieval call binding the contract method 0xd2e30585.
//
// Solidity: function rayMul(uint256 a, uint256 b) pure returns(uint256)
func (_WadRayMathTest *WadRayMathTestSession) RayMul(a *big.Int, b *big.Int) (*big.Int, error) {
	return _WadRayMathTest.Contract.RayMul(&_WadRayMathTest.CallOpts, a, b)
}

// RayMul is a free data retrieval call binding the contract method 0xd2e30585.
//
// Solidity: function rayMul(uint256 a, uint256 b) pure returns(uint256)
func (_WadRayMathTest *WadRayMathTestCallerSession) RayMul(a *big.Int, b *big.Int) (*big.Int, error) {
	return _WadRayMathTest.Contract.RayMul(&_WadRayMathTest.CallOpts, a, b)
}

// RayToWad is a free data retrieval call binding the contract method 0x29cb5aa4.
//
// Solidity: function rayToWad(uint256 a) pure returns(uint256)
func (_WadRayMathTest *WadRayMathTestCaller) RayToWad(opts *bind.CallOpts, a *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _WadRayMathTest.contract.Call(opts, &out, "rayToWad", a)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RayToWad is a free data retrieval call binding the contract method 0x29cb5aa4.
//
// Solidity: function rayToWad(uint256 a) pure returns(uint256)
func (_WadRayMathTest *WadRayMathTestSession) RayToWad(a *big.Int) (*big.Int, error) {
	return _WadRayMathTest.Contract.RayToWad(&_WadRayMathTest.CallOpts, a)
}

// RayToWad is a free data retrieval call binding the contract method 0x29cb5aa4.
//
// Solidity: function rayToWad(uint256 a) pure returns(uint256)
func (_WadRayMathTest *WadRayMathTestCallerSession) RayToWad(a *big.Int) (*big.Int, error) {
	return _WadRayMathTest.Contract.RayToWad(&_WadRayMathTest.CallOpts, a)
}

// Wad is a free data retrieval call binding the contract method 0x7df38c5b.
//
// Solidity: function wad() pure returns(uint256)
func (_WadRayMathTest *WadRayMathTestCaller) Wad(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WadRayMathTest.contract.Call(opts, &out, "wad")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wad is a free data retrieval call binding the contract method 0x7df38c5b.
//
// Solidity: function wad() pure returns(uint256)
func (_WadRayMathTest *WadRayMathTestSession) Wad() (*big.Int, error) {
	return _WadRayMathTest.Contract.Wad(&_WadRayMathTest.CallOpts)
}

// Wad is a free data retrieval call binding the contract method 0x7df38c5b.
//
// Solidity: function wad() pure returns(uint256)
func (_WadRayMathTest *WadRayMathTestCallerSession) Wad() (*big.Int, error) {
	return _WadRayMathTest.Contract.Wad(&_WadRayMathTest.CallOpts)
}

// WadDiv is a free data retrieval call binding the contract method 0xe57b6d3b.
//
// Solidity: function wadDiv(uint256 a, uint256 b) pure returns(uint256)
func (_WadRayMathTest *WadRayMathTestCaller) WadDiv(opts *bind.CallOpts, a *big.Int, b *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _WadRayMathTest.contract.Call(opts, &out, "wadDiv", a, b)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WadDiv is a free data retrieval call binding the contract method 0xe57b6d3b.
//
// Solidity: function wadDiv(uint256 a, uint256 b) pure returns(uint256)
func (_WadRayMathTest *WadRayMathTestSession) WadDiv(a *big.Int, b *big.Int) (*big.Int, error) {
	return _WadRayMathTest.Contract.WadDiv(&_WadRayMathTest.CallOpts, a, b)
}

// WadDiv is a free data retrieval call binding the contract method 0xe57b6d3b.
//
// Solidity: function wadDiv(uint256 a, uint256 b) pure returns(uint256)
func (_WadRayMathTest *WadRayMathTestCallerSession) WadDiv(a *big.Int, b *big.Int) (*big.Int, error) {
	return _WadRayMathTest.Contract.WadDiv(&_WadRayMathTest.CallOpts, a, b)
}

// WadMul is a free data retrieval call binding the contract method 0x761fdad6.
//
// Solidity: function wadMul(uint256 a, uint256 b) pure returns(uint256)
func (_WadRayMathTest *WadRayMathTestCaller) WadMul(opts *bind.CallOpts, a *big.Int, b *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _WadRayMathTest.contract.Call(opts, &out, "wadMul", a, b)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WadMul is a free data retrieval call binding the contract method 0x761fdad6.
//
// Solidity: function wadMul(uint256 a, uint256 b) pure returns(uint256)
func (_WadRayMathTest *WadRayMathTestSession) WadMul(a *big.Int, b *big.Int) (*big.Int, error) {
	return _WadRayMathTest.Contract.WadMul(&_WadRayMathTest.CallOpts, a, b)
}

// WadMul is a free data retrieval call binding the contract method 0x761fdad6.
//
// Solidity: function wadMul(uint256 a, uint256 b) pure returns(uint256)
func (_WadRayMathTest *WadRayMathTestCallerSession) WadMul(a *big.Int, b *big.Int) (*big.Int, error) {
	return _WadRayMathTest.Contract.WadMul(&_WadRayMathTest.CallOpts, a, b)
}

// WadToRay is a free data retrieval call binding the contract method 0x10de27b9.
//
// Solidity: function wadToRay(uint256 a) pure returns(uint256)
func (_WadRayMathTest *WadRayMathTestCaller) WadToRay(opts *bind.CallOpts, a *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _WadRayMathTest.contract.Call(opts, &out, "wadToRay", a)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WadToRay is a free data retrieval call binding the contract method 0x10de27b9.
//
// Solidity: function wadToRay(uint256 a) pure returns(uint256)
func (_WadRayMathTest *WadRayMathTestSession) WadToRay(a *big.Int) (*big.Int, error) {
	return _WadRayMathTest.Contract.WadToRay(&_WadRayMathTest.CallOpts, a)
}

// WadToRay is a free data retrieval call binding the contract method 0x10de27b9.
//
// Solidity: function wadToRay(uint256 a) pure returns(uint256)
func (_WadRayMathTest *WadRayMathTestCallerSession) WadToRay(a *big.Int) (*big.Int, error) {
	return _WadRayMathTest.Contract.WadToRay(&_WadRayMathTest.CallOpts, a)
}
