// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package linearInterestRateModel

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

// LinearInterestRateModelMetaData contains all meta data concerning the LinearInterestRateModel contract.
var LinearInterestRateModelMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"U_optimal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"R_base\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"R_slope1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"R_slope2\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"_R_base_RAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_R_slope1_RAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_R_slope2_RAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_U_Optimal_WAD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_U_Optimal_inverted_WAD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expectedLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableLiquidity\",\"type\":\"uint256\"}],\"name\":\"calcBorrowRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getModelParameters\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"U_optimal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"R_base\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"R_slope1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"R_slope2\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// LinearInterestRateModelABI is the input ABI used to generate the binding from.
// Deprecated: Use LinearInterestRateModelMetaData.ABI instead.
var LinearInterestRateModelABI = LinearInterestRateModelMetaData.ABI

// LinearInterestRateModel is an auto generated Go binding around an Ethereum contract.
type LinearInterestRateModel struct {
	LinearInterestRateModelCaller     // Read-only binding to the contract
	LinearInterestRateModelTransactor // Write-only binding to the contract
	LinearInterestRateModelFilterer   // Log filterer for contract events
}

// LinearInterestRateModelCaller is an auto generated read-only Go binding around an Ethereum contract.
type LinearInterestRateModelCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LinearInterestRateModelTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LinearInterestRateModelTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LinearInterestRateModelFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LinearInterestRateModelFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LinearInterestRateModelSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LinearInterestRateModelSession struct {
	Contract     *LinearInterestRateModel // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// LinearInterestRateModelCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LinearInterestRateModelCallerSession struct {
	Contract *LinearInterestRateModelCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// LinearInterestRateModelTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LinearInterestRateModelTransactorSession struct {
	Contract     *LinearInterestRateModelTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// LinearInterestRateModelRaw is an auto generated low-level Go binding around an Ethereum contract.
type LinearInterestRateModelRaw struct {
	Contract *LinearInterestRateModel // Generic contract binding to access the raw methods on
}

// LinearInterestRateModelCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LinearInterestRateModelCallerRaw struct {
	Contract *LinearInterestRateModelCaller // Generic read-only contract binding to access the raw methods on
}

// LinearInterestRateModelTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LinearInterestRateModelTransactorRaw struct {
	Contract *LinearInterestRateModelTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLinearInterestRateModel creates a new instance of LinearInterestRateModel, bound to a specific deployed contract.
func NewLinearInterestRateModel(address common.Address, backend bind.ContractBackend) (*LinearInterestRateModel, error) {
	contract, err := bindLinearInterestRateModel(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LinearInterestRateModel{LinearInterestRateModelCaller: LinearInterestRateModelCaller{contract: contract}, LinearInterestRateModelTransactor: LinearInterestRateModelTransactor{contract: contract}, LinearInterestRateModelFilterer: LinearInterestRateModelFilterer{contract: contract}}, nil
}

// NewLinearInterestRateModelCaller creates a new read-only instance of LinearInterestRateModel, bound to a specific deployed contract.
func NewLinearInterestRateModelCaller(address common.Address, caller bind.ContractCaller) (*LinearInterestRateModelCaller, error) {
	contract, err := bindLinearInterestRateModel(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LinearInterestRateModelCaller{contract: contract}, nil
}

// NewLinearInterestRateModelTransactor creates a new write-only instance of LinearInterestRateModel, bound to a specific deployed contract.
func NewLinearInterestRateModelTransactor(address common.Address, transactor bind.ContractTransactor) (*LinearInterestRateModelTransactor, error) {
	contract, err := bindLinearInterestRateModel(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LinearInterestRateModelTransactor{contract: contract}, nil
}

// NewLinearInterestRateModelFilterer creates a new log filterer instance of LinearInterestRateModel, bound to a specific deployed contract.
func NewLinearInterestRateModelFilterer(address common.Address, filterer bind.ContractFilterer) (*LinearInterestRateModelFilterer, error) {
	contract, err := bindLinearInterestRateModel(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LinearInterestRateModelFilterer{contract: contract}, nil
}

// bindLinearInterestRateModel binds a generic wrapper to an already deployed contract.
func bindLinearInterestRateModel(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LinearInterestRateModelABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LinearInterestRateModel *LinearInterestRateModelRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LinearInterestRateModel.Contract.LinearInterestRateModelCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LinearInterestRateModel *LinearInterestRateModelRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LinearInterestRateModel.Contract.LinearInterestRateModelTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LinearInterestRateModel *LinearInterestRateModelRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LinearInterestRateModel.Contract.LinearInterestRateModelTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LinearInterestRateModel *LinearInterestRateModelCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LinearInterestRateModel.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LinearInterestRateModel *LinearInterestRateModelTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LinearInterestRateModel.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LinearInterestRateModel *LinearInterestRateModelTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LinearInterestRateModel.Contract.contract.Transact(opts, method, params...)
}

// RBaseRAY is a free data retrieval call binding the contract method 0x9cd3fdb5.
//
// Solidity: function _R_base_RAY() view returns(uint256)
func (_LinearInterestRateModel *LinearInterestRateModelCaller) RBaseRAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LinearInterestRateModel.contract.Call(opts, &out, "_R_base_RAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RBaseRAY is a free data retrieval call binding the contract method 0x9cd3fdb5.
//
// Solidity: function _R_base_RAY() view returns(uint256)
func (_LinearInterestRateModel *LinearInterestRateModelSession) RBaseRAY() (*big.Int, error) {
	return _LinearInterestRateModel.Contract.RBaseRAY(&_LinearInterestRateModel.CallOpts)
}

// RBaseRAY is a free data retrieval call binding the contract method 0x9cd3fdb5.
//
// Solidity: function _R_base_RAY() view returns(uint256)
func (_LinearInterestRateModel *LinearInterestRateModelCallerSession) RBaseRAY() (*big.Int, error) {
	return _LinearInterestRateModel.Contract.RBaseRAY(&_LinearInterestRateModel.CallOpts)
}

// RSlope1RAY is a free data retrieval call binding the contract method 0x9aec06ea.
//
// Solidity: function _R_slope1_RAY() view returns(uint256)
func (_LinearInterestRateModel *LinearInterestRateModelCaller) RSlope1RAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LinearInterestRateModel.contract.Call(opts, &out, "_R_slope1_RAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RSlope1RAY is a free data retrieval call binding the contract method 0x9aec06ea.
//
// Solidity: function _R_slope1_RAY() view returns(uint256)
func (_LinearInterestRateModel *LinearInterestRateModelSession) RSlope1RAY() (*big.Int, error) {
	return _LinearInterestRateModel.Contract.RSlope1RAY(&_LinearInterestRateModel.CallOpts)
}

// RSlope1RAY is a free data retrieval call binding the contract method 0x9aec06ea.
//
// Solidity: function _R_slope1_RAY() view returns(uint256)
func (_LinearInterestRateModel *LinearInterestRateModelCallerSession) RSlope1RAY() (*big.Int, error) {
	return _LinearInterestRateModel.Contract.RSlope1RAY(&_LinearInterestRateModel.CallOpts)
}

// RSlope2RAY is a free data retrieval call binding the contract method 0xfc4b2b78.
//
// Solidity: function _R_slope2_RAY() view returns(uint256)
func (_LinearInterestRateModel *LinearInterestRateModelCaller) RSlope2RAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LinearInterestRateModel.contract.Call(opts, &out, "_R_slope2_RAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RSlope2RAY is a free data retrieval call binding the contract method 0xfc4b2b78.
//
// Solidity: function _R_slope2_RAY() view returns(uint256)
func (_LinearInterestRateModel *LinearInterestRateModelSession) RSlope2RAY() (*big.Int, error) {
	return _LinearInterestRateModel.Contract.RSlope2RAY(&_LinearInterestRateModel.CallOpts)
}

// RSlope2RAY is a free data retrieval call binding the contract method 0xfc4b2b78.
//
// Solidity: function _R_slope2_RAY() view returns(uint256)
func (_LinearInterestRateModel *LinearInterestRateModelCallerSession) RSlope2RAY() (*big.Int, error) {
	return _LinearInterestRateModel.Contract.RSlope2RAY(&_LinearInterestRateModel.CallOpts)
}

// UOptimalWAD is a free data retrieval call binding the contract method 0x50ced104.
//
// Solidity: function _U_Optimal_WAD() view returns(uint256)
func (_LinearInterestRateModel *LinearInterestRateModelCaller) UOptimalWAD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LinearInterestRateModel.contract.Call(opts, &out, "_U_Optimal_WAD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UOptimalWAD is a free data retrieval call binding the contract method 0x50ced104.
//
// Solidity: function _U_Optimal_WAD() view returns(uint256)
func (_LinearInterestRateModel *LinearInterestRateModelSession) UOptimalWAD() (*big.Int, error) {
	return _LinearInterestRateModel.Contract.UOptimalWAD(&_LinearInterestRateModel.CallOpts)
}

// UOptimalWAD is a free data retrieval call binding the contract method 0x50ced104.
//
// Solidity: function _U_Optimal_WAD() view returns(uint256)
func (_LinearInterestRateModel *LinearInterestRateModelCallerSession) UOptimalWAD() (*big.Int, error) {
	return _LinearInterestRateModel.Contract.UOptimalWAD(&_LinearInterestRateModel.CallOpts)
}

// UOptimalInvertedWAD is a free data retrieval call binding the contract method 0xf81d4381.
//
// Solidity: function _U_Optimal_inverted_WAD() view returns(uint256)
func (_LinearInterestRateModel *LinearInterestRateModelCaller) UOptimalInvertedWAD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LinearInterestRateModel.contract.Call(opts, &out, "_U_Optimal_inverted_WAD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UOptimalInvertedWAD is a free data retrieval call binding the contract method 0xf81d4381.
//
// Solidity: function _U_Optimal_inverted_WAD() view returns(uint256)
func (_LinearInterestRateModel *LinearInterestRateModelSession) UOptimalInvertedWAD() (*big.Int, error) {
	return _LinearInterestRateModel.Contract.UOptimalInvertedWAD(&_LinearInterestRateModel.CallOpts)
}

// UOptimalInvertedWAD is a free data retrieval call binding the contract method 0xf81d4381.
//
// Solidity: function _U_Optimal_inverted_WAD() view returns(uint256)
func (_LinearInterestRateModel *LinearInterestRateModelCallerSession) UOptimalInvertedWAD() (*big.Int, error) {
	return _LinearInterestRateModel.Contract.UOptimalInvertedWAD(&_LinearInterestRateModel.CallOpts)
}

// CalcBorrowRate is a free data retrieval call binding the contract method 0x42568d44.
//
// Solidity: function calcBorrowRate(uint256 expectedLiquidity, uint256 availableLiquidity) view returns(uint256)
func (_LinearInterestRateModel *LinearInterestRateModelCaller) CalcBorrowRate(opts *bind.CallOpts, expectedLiquidity *big.Int, availableLiquidity *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LinearInterestRateModel.contract.Call(opts, &out, "calcBorrowRate", expectedLiquidity, availableLiquidity)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcBorrowRate is a free data retrieval call binding the contract method 0x42568d44.
//
// Solidity: function calcBorrowRate(uint256 expectedLiquidity, uint256 availableLiquidity) view returns(uint256)
func (_LinearInterestRateModel *LinearInterestRateModelSession) CalcBorrowRate(expectedLiquidity *big.Int, availableLiquidity *big.Int) (*big.Int, error) {
	return _LinearInterestRateModel.Contract.CalcBorrowRate(&_LinearInterestRateModel.CallOpts, expectedLiquidity, availableLiquidity)
}

// CalcBorrowRate is a free data retrieval call binding the contract method 0x42568d44.
//
// Solidity: function calcBorrowRate(uint256 expectedLiquidity, uint256 availableLiquidity) view returns(uint256)
func (_LinearInterestRateModel *LinearInterestRateModelCallerSession) CalcBorrowRate(expectedLiquidity *big.Int, availableLiquidity *big.Int) (*big.Int, error) {
	return _LinearInterestRateModel.Contract.CalcBorrowRate(&_LinearInterestRateModel.CallOpts, expectedLiquidity, availableLiquidity)
}

// GetModelParameters is a free data retrieval call binding the contract method 0xc8284e6d.
//
// Solidity: function getModelParameters() view returns(uint256 U_optimal, uint256 R_base, uint256 R_slope1, uint256 R_slope2)
func (_LinearInterestRateModel *LinearInterestRateModelCaller) GetModelParameters(opts *bind.CallOpts) (struct {
	UOptimal *big.Int
	RBase    *big.Int
	RSlope1  *big.Int
	RSlope2  *big.Int
}, error) {
	var out []interface{}
	err := _LinearInterestRateModel.contract.Call(opts, &out, "getModelParameters")

	outstruct := new(struct {
		UOptimal *big.Int
		RBase    *big.Int
		RSlope1  *big.Int
		RSlope2  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.UOptimal = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RBase = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.RSlope1 = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.RSlope2 = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetModelParameters is a free data retrieval call binding the contract method 0xc8284e6d.
//
// Solidity: function getModelParameters() view returns(uint256 U_optimal, uint256 R_base, uint256 R_slope1, uint256 R_slope2)
func (_LinearInterestRateModel *LinearInterestRateModelSession) GetModelParameters() (struct {
	UOptimal *big.Int
	RBase    *big.Int
	RSlope1  *big.Int
	RSlope2  *big.Int
}, error) {
	return _LinearInterestRateModel.Contract.GetModelParameters(&_LinearInterestRateModel.CallOpts)
}

// GetModelParameters is a free data retrieval call binding the contract method 0xc8284e6d.
//
// Solidity: function getModelParameters() view returns(uint256 U_optimal, uint256 R_base, uint256 R_slope1, uint256 R_slope2)
func (_LinearInterestRateModel *LinearInterestRateModelCallerSession) GetModelParameters() (struct {
	UOptimal *big.Int
	RBase    *big.Int
	RSlope1  *big.Int
	RSlope2  *big.Int
}, error) {
	return _LinearInterestRateModel.Contract.GetModelParameters(&_LinearInterestRateModel.CallOpts)
}
