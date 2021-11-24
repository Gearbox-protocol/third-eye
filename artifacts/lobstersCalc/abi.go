// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lobstersCalc

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

// LobstersCalcMetaData contains all meta data concerning the LobstersCalc contract.
var LobstersCalcMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"maxTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"metadataOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"seed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// LobstersCalcABI is the input ABI used to generate the binding from.
// Deprecated: Use LobstersCalcMetaData.ABI instead.
var LobstersCalcABI = LobstersCalcMetaData.ABI

// LobstersCalc is an auto generated Go binding around an Ethereum contract.
type LobstersCalc struct {
	LobstersCalcCaller     // Read-only binding to the contract
	LobstersCalcTransactor // Write-only binding to the contract
	LobstersCalcFilterer   // Log filterer for contract events
}

// LobstersCalcCaller is an auto generated read-only Go binding around an Ethereum contract.
type LobstersCalcCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LobstersCalcTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LobstersCalcTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LobstersCalcFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LobstersCalcFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LobstersCalcSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LobstersCalcSession struct {
	Contract     *LobstersCalc     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LobstersCalcCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LobstersCalcCallerSession struct {
	Contract *LobstersCalcCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// LobstersCalcTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LobstersCalcTransactorSession struct {
	Contract     *LobstersCalcTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// LobstersCalcRaw is an auto generated low-level Go binding around an Ethereum contract.
type LobstersCalcRaw struct {
	Contract *LobstersCalc // Generic contract binding to access the raw methods on
}

// LobstersCalcCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LobstersCalcCallerRaw struct {
	Contract *LobstersCalcCaller // Generic read-only contract binding to access the raw methods on
}

// LobstersCalcTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LobstersCalcTransactorRaw struct {
	Contract *LobstersCalcTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLobstersCalc creates a new instance of LobstersCalc, bound to a specific deployed contract.
func NewLobstersCalc(address common.Address, backend bind.ContractBackend) (*LobstersCalc, error) {
	contract, err := bindLobstersCalc(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LobstersCalc{LobstersCalcCaller: LobstersCalcCaller{contract: contract}, LobstersCalcTransactor: LobstersCalcTransactor{contract: contract}, LobstersCalcFilterer: LobstersCalcFilterer{contract: contract}}, nil
}

// NewLobstersCalcCaller creates a new read-only instance of LobstersCalc, bound to a specific deployed contract.
func NewLobstersCalcCaller(address common.Address, caller bind.ContractCaller) (*LobstersCalcCaller, error) {
	contract, err := bindLobstersCalc(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LobstersCalcCaller{contract: contract}, nil
}

// NewLobstersCalcTransactor creates a new write-only instance of LobstersCalc, bound to a specific deployed contract.
func NewLobstersCalcTransactor(address common.Address, transactor bind.ContractTransactor) (*LobstersCalcTransactor, error) {
	contract, err := bindLobstersCalc(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LobstersCalcTransactor{contract: contract}, nil
}

// NewLobstersCalcFilterer creates a new log filterer instance of LobstersCalc, bound to a specific deployed contract.
func NewLobstersCalcFilterer(address common.Address, filterer bind.ContractFilterer) (*LobstersCalcFilterer, error) {
	contract, err := bindLobstersCalc(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LobstersCalcFilterer{contract: contract}, nil
}

// bindLobstersCalc binds a generic wrapper to an already deployed contract.
func bindLobstersCalc(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LobstersCalcABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LobstersCalc *LobstersCalcRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LobstersCalc.Contract.LobstersCalcCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LobstersCalc *LobstersCalcRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LobstersCalc.Contract.LobstersCalcTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LobstersCalc *LobstersCalcRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LobstersCalc.Contract.LobstersCalcTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LobstersCalc *LobstersCalcCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LobstersCalc.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LobstersCalc *LobstersCalcTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LobstersCalc.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LobstersCalc *LobstersCalcTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LobstersCalc.Contract.contract.Transact(opts, method, params...)
}

// MaxTokens is a free data retrieval call binding the contract method 0xe8315742.
//
// Solidity: function maxTokens() view returns(uint256)
func (_LobstersCalc *LobstersCalcCaller) MaxTokens(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LobstersCalc.contract.Call(opts, &out, "maxTokens")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxTokens is a free data retrieval call binding the contract method 0xe8315742.
//
// Solidity: function maxTokens() view returns(uint256)
func (_LobstersCalc *LobstersCalcSession) MaxTokens() (*big.Int, error) {
	return _LobstersCalc.Contract.MaxTokens(&_LobstersCalc.CallOpts)
}

// MaxTokens is a free data retrieval call binding the contract method 0xe8315742.
//
// Solidity: function maxTokens() view returns(uint256)
func (_LobstersCalc *LobstersCalcCallerSession) MaxTokens() (*big.Int, error) {
	return _LobstersCalc.Contract.MaxTokens(&_LobstersCalc.CallOpts)
}

// MetadataOf is a free data retrieval call binding the contract method 0x0ef7cc8e.
//
// Solidity: function metadataOf(uint256 _tokenId) view returns(uint256)
func (_LobstersCalc *LobstersCalcCaller) MetadataOf(opts *bind.CallOpts, _tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LobstersCalc.contract.Call(opts, &out, "metadataOf", _tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MetadataOf is a free data retrieval call binding the contract method 0x0ef7cc8e.
//
// Solidity: function metadataOf(uint256 _tokenId) view returns(uint256)
func (_LobstersCalc *LobstersCalcSession) MetadataOf(_tokenId *big.Int) (*big.Int, error) {
	return _LobstersCalc.Contract.MetadataOf(&_LobstersCalc.CallOpts, _tokenId)
}

// MetadataOf is a free data retrieval call binding the contract method 0x0ef7cc8e.
//
// Solidity: function metadataOf(uint256 _tokenId) view returns(uint256)
func (_LobstersCalc *LobstersCalcCallerSession) MetadataOf(_tokenId *big.Int) (*big.Int, error) {
	return _LobstersCalc.Contract.MetadataOf(&_LobstersCalc.CallOpts, _tokenId)
}

// Seed is a free data retrieval call binding the contract method 0x7d94792a.
//
// Solidity: function seed() view returns(uint256)
func (_LobstersCalc *LobstersCalcCaller) Seed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LobstersCalc.contract.Call(opts, &out, "seed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Seed is a free data retrieval call binding the contract method 0x7d94792a.
//
// Solidity: function seed() view returns(uint256)
func (_LobstersCalc *LobstersCalcSession) Seed() (*big.Int, error) {
	return _LobstersCalc.Contract.Seed(&_LobstersCalc.CallOpts)
}

// Seed is a free data retrieval call binding the contract method 0x7d94792a.
//
// Solidity: function seed() view returns(uint256)
func (_LobstersCalc *LobstersCalcCallerSession) Seed() (*big.Int, error) {
	return _LobstersCalc.Contract.Seed(&_LobstersCalc.CallOpts)
}
