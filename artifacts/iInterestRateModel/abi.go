// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iInterestRateModel

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

// IInterestRateModelMetaData contains all meta data concerning the IInterestRateModel contract.
var IInterestRateModelMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expectedLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableLiquidity\",\"type\":\"uint256\"}],\"name\":\"calcBorrowRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IInterestRateModelABI is the input ABI used to generate the binding from.
// Deprecated: Use IInterestRateModelMetaData.ABI instead.
var IInterestRateModelABI = IInterestRateModelMetaData.ABI

// IInterestRateModel is an auto generated Go binding around an Ethereum contract.
type IInterestRateModel struct {
	IInterestRateModelCaller     // Read-only binding to the contract
	IInterestRateModelTransactor // Write-only binding to the contract
	IInterestRateModelFilterer   // Log filterer for contract events
}

// IInterestRateModelCaller is an auto generated read-only Go binding around an Ethereum contract.
type IInterestRateModelCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInterestRateModelTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IInterestRateModelTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInterestRateModelFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IInterestRateModelFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInterestRateModelSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IInterestRateModelSession struct {
	Contract     *IInterestRateModel // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IInterestRateModelCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IInterestRateModelCallerSession struct {
	Contract *IInterestRateModelCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// IInterestRateModelTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IInterestRateModelTransactorSession struct {
	Contract     *IInterestRateModelTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// IInterestRateModelRaw is an auto generated low-level Go binding around an Ethereum contract.
type IInterestRateModelRaw struct {
	Contract *IInterestRateModel // Generic contract binding to access the raw methods on
}

// IInterestRateModelCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IInterestRateModelCallerRaw struct {
	Contract *IInterestRateModelCaller // Generic read-only contract binding to access the raw methods on
}

// IInterestRateModelTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IInterestRateModelTransactorRaw struct {
	Contract *IInterestRateModelTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIInterestRateModel creates a new instance of IInterestRateModel, bound to a specific deployed contract.
func NewIInterestRateModel(address common.Address, backend bind.ContractBackend) (*IInterestRateModel, error) {
	contract, err := bindIInterestRateModel(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IInterestRateModel{IInterestRateModelCaller: IInterestRateModelCaller{contract: contract}, IInterestRateModelTransactor: IInterestRateModelTransactor{contract: contract}, IInterestRateModelFilterer: IInterestRateModelFilterer{contract: contract}}, nil
}

// NewIInterestRateModelCaller creates a new read-only instance of IInterestRateModel, bound to a specific deployed contract.
func NewIInterestRateModelCaller(address common.Address, caller bind.ContractCaller) (*IInterestRateModelCaller, error) {
	contract, err := bindIInterestRateModel(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IInterestRateModelCaller{contract: contract}, nil
}

// NewIInterestRateModelTransactor creates a new write-only instance of IInterestRateModel, bound to a specific deployed contract.
func NewIInterestRateModelTransactor(address common.Address, transactor bind.ContractTransactor) (*IInterestRateModelTransactor, error) {
	contract, err := bindIInterestRateModel(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IInterestRateModelTransactor{contract: contract}, nil
}

// NewIInterestRateModelFilterer creates a new log filterer instance of IInterestRateModel, bound to a specific deployed contract.
func NewIInterestRateModelFilterer(address common.Address, filterer bind.ContractFilterer) (*IInterestRateModelFilterer, error) {
	contract, err := bindIInterestRateModel(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IInterestRateModelFilterer{contract: contract}, nil
}

// bindIInterestRateModel binds a generic wrapper to an already deployed contract.
func bindIInterestRateModel(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IInterestRateModelABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IInterestRateModel *IInterestRateModelRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IInterestRateModel.Contract.IInterestRateModelCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IInterestRateModel *IInterestRateModelRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IInterestRateModel.Contract.IInterestRateModelTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IInterestRateModel *IInterestRateModelRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IInterestRateModel.Contract.IInterestRateModelTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IInterestRateModel *IInterestRateModelCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IInterestRateModel.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IInterestRateModel *IInterestRateModelTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IInterestRateModel.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IInterestRateModel *IInterestRateModelTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IInterestRateModel.Contract.contract.Transact(opts, method, params...)
}

// CalcBorrowRate is a free data retrieval call binding the contract method 0x42568d44.
//
// Solidity: function calcBorrowRate(uint256 expectedLiquidity, uint256 availableLiquidity) view returns(uint256)
func (_IInterestRateModel *IInterestRateModelCaller) CalcBorrowRate(opts *bind.CallOpts, expectedLiquidity *big.Int, availableLiquidity *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IInterestRateModel.contract.Call(opts, &out, "calcBorrowRate", expectedLiquidity, availableLiquidity)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcBorrowRate is a free data retrieval call binding the contract method 0x42568d44.
//
// Solidity: function calcBorrowRate(uint256 expectedLiquidity, uint256 availableLiquidity) view returns(uint256)
func (_IInterestRateModel *IInterestRateModelSession) CalcBorrowRate(expectedLiquidity *big.Int, availableLiquidity *big.Int) (*big.Int, error) {
	return _IInterestRateModel.Contract.CalcBorrowRate(&_IInterestRateModel.CallOpts, expectedLiquidity, availableLiquidity)
}

// CalcBorrowRate is a free data retrieval call binding the contract method 0x42568d44.
//
// Solidity: function calcBorrowRate(uint256 expectedLiquidity, uint256 availableLiquidity) view returns(uint256)
func (_IInterestRateModel *IInterestRateModelCallerSession) CalcBorrowRate(expectedLiquidity *big.Int, availableLiquidity *big.Int) (*big.Int, error) {
	return _IInterestRateModel.Contract.CalcBorrowRate(&_IInterestRateModel.CallOpts, expectedLiquidity, availableLiquidity)
}
