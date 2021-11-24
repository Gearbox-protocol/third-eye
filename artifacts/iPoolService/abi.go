// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iPoolService

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

// IPoolServiceMetaData contains all meta data concerning the IPoolService contract.
var IPoolServiceMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"referralCode\",\"type\":\"uint256\"}],\"name\":\"AddLiquidity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creditManager\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Borrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creditManager\",\"type\":\"address\"}],\"name\":\"BorrowForbidden\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creditManager\",\"type\":\"address\"}],\"name\":\"NewCreditManagerConnected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newLimit\",\"type\":\"uint256\"}],\"name\":\"NewExpectedLiquidityLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newInterestRateModel\",\"type\":\"address\"}],\"name\":\"NewInterestRateModel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"NewWithdrawFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RemoveLiquidity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creditManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"borrowedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"profit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"loss\",\"type\":\"uint256\"}],\"name\":\"Repay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creditManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"loss\",\"type\":\"uint256\"}],\"name\":\"UncoveredLoss\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_cumulativeIndex_RAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_timestampLU\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"referralCode\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"availableLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"borrowAPY_RAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cumulativeIndexAtOpen\",\"type\":\"uint256\"}],\"name\":\"calcCumulativeIndexAtBorrowMore\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calcLinearCumulative_RAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"creditManagers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"id\",\"type\":\"address\"}],\"name\":\"creditManagersCanBorrow\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creditManagersCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dieselToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expectedLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expectedLiquidityLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"fromDiesel\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDieselRate_RAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"borrowedAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"}],\"name\":\"lendCreditAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"removeLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"borrowedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"profit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"loss\",\"type\":\"uint256\"}],\"name\":\"repayCreditAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"toDiesel\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBorrowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlyingToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IPoolServiceABI is the input ABI used to generate the binding from.
// Deprecated: Use IPoolServiceMetaData.ABI instead.
var IPoolServiceABI = IPoolServiceMetaData.ABI

// IPoolService is an auto generated Go binding around an Ethereum contract.
type IPoolService struct {
	IPoolServiceCaller     // Read-only binding to the contract
	IPoolServiceTransactor // Write-only binding to the contract
	IPoolServiceFilterer   // Log filterer for contract events
}

// IPoolServiceCaller is an auto generated read-only Go binding around an Ethereum contract.
type IPoolServiceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPoolServiceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IPoolServiceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPoolServiceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPoolServiceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPoolServiceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPoolServiceSession struct {
	Contract     *IPoolService     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPoolServiceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPoolServiceCallerSession struct {
	Contract *IPoolServiceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IPoolServiceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPoolServiceTransactorSession struct {
	Contract     *IPoolServiceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IPoolServiceRaw is an auto generated low-level Go binding around an Ethereum contract.
type IPoolServiceRaw struct {
	Contract *IPoolService // Generic contract binding to access the raw methods on
}

// IPoolServiceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPoolServiceCallerRaw struct {
	Contract *IPoolServiceCaller // Generic read-only contract binding to access the raw methods on
}

// IPoolServiceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPoolServiceTransactorRaw struct {
	Contract *IPoolServiceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIPoolService creates a new instance of IPoolService, bound to a specific deployed contract.
func NewIPoolService(address common.Address, backend bind.ContractBackend) (*IPoolService, error) {
	contract, err := bindIPoolService(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPoolService{IPoolServiceCaller: IPoolServiceCaller{contract: contract}, IPoolServiceTransactor: IPoolServiceTransactor{contract: contract}, IPoolServiceFilterer: IPoolServiceFilterer{contract: contract}}, nil
}

// NewIPoolServiceCaller creates a new read-only instance of IPoolService, bound to a specific deployed contract.
func NewIPoolServiceCaller(address common.Address, caller bind.ContractCaller) (*IPoolServiceCaller, error) {
	contract, err := bindIPoolService(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPoolServiceCaller{contract: contract}, nil
}

// NewIPoolServiceTransactor creates a new write-only instance of IPoolService, bound to a specific deployed contract.
func NewIPoolServiceTransactor(address common.Address, transactor bind.ContractTransactor) (*IPoolServiceTransactor, error) {
	contract, err := bindIPoolService(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPoolServiceTransactor{contract: contract}, nil
}

// NewIPoolServiceFilterer creates a new log filterer instance of IPoolService, bound to a specific deployed contract.
func NewIPoolServiceFilterer(address common.Address, filterer bind.ContractFilterer) (*IPoolServiceFilterer, error) {
	contract, err := bindIPoolService(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPoolServiceFilterer{contract: contract}, nil
}

// bindIPoolService binds a generic wrapper to an already deployed contract.
func bindIPoolService(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IPoolServiceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPoolService *IPoolServiceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPoolService.Contract.IPoolServiceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPoolService *IPoolServiceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPoolService.Contract.IPoolServiceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPoolService *IPoolServiceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPoolService.Contract.IPoolServiceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPoolService *IPoolServiceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPoolService.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPoolService *IPoolServiceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPoolService.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPoolService *IPoolServiceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPoolService.Contract.contract.Transact(opts, method, params...)
}

// CumulativeIndexRAY is a free data retrieval call binding the contract method 0xdbcb313b.
//
// Solidity: function _cumulativeIndex_RAY() view returns(uint256)
func (_IPoolService *IPoolServiceCaller) CumulativeIndexRAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IPoolService.contract.Call(opts, &out, "_cumulativeIndex_RAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CumulativeIndexRAY is a free data retrieval call binding the contract method 0xdbcb313b.
//
// Solidity: function _cumulativeIndex_RAY() view returns(uint256)
func (_IPoolService *IPoolServiceSession) CumulativeIndexRAY() (*big.Int, error) {
	return _IPoolService.Contract.CumulativeIndexRAY(&_IPoolService.CallOpts)
}

// CumulativeIndexRAY is a free data retrieval call binding the contract method 0xdbcb313b.
//
// Solidity: function _cumulativeIndex_RAY() view returns(uint256)
func (_IPoolService *IPoolServiceCallerSession) CumulativeIndexRAY() (*big.Int, error) {
	return _IPoolService.Contract.CumulativeIndexRAY(&_IPoolService.CallOpts)
}

// TimestampLU is a free data retrieval call binding the contract method 0x609ae317.
//
// Solidity: function _timestampLU() view returns(uint256)
func (_IPoolService *IPoolServiceCaller) TimestampLU(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IPoolService.contract.Call(opts, &out, "_timestampLU")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TimestampLU is a free data retrieval call binding the contract method 0x609ae317.
//
// Solidity: function _timestampLU() view returns(uint256)
func (_IPoolService *IPoolServiceSession) TimestampLU() (*big.Int, error) {
	return _IPoolService.Contract.TimestampLU(&_IPoolService.CallOpts)
}

// TimestampLU is a free data retrieval call binding the contract method 0x609ae317.
//
// Solidity: function _timestampLU() view returns(uint256)
func (_IPoolService *IPoolServiceCallerSession) TimestampLU() (*big.Int, error) {
	return _IPoolService.Contract.TimestampLU(&_IPoolService.CallOpts)
}

// AvailableLiquidity is a free data retrieval call binding the contract method 0x74375359.
//
// Solidity: function availableLiquidity() view returns(uint256)
func (_IPoolService *IPoolServiceCaller) AvailableLiquidity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IPoolService.contract.Call(opts, &out, "availableLiquidity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AvailableLiquidity is a free data retrieval call binding the contract method 0x74375359.
//
// Solidity: function availableLiquidity() view returns(uint256)
func (_IPoolService *IPoolServiceSession) AvailableLiquidity() (*big.Int, error) {
	return _IPoolService.Contract.AvailableLiquidity(&_IPoolService.CallOpts)
}

// AvailableLiquidity is a free data retrieval call binding the contract method 0x74375359.
//
// Solidity: function availableLiquidity() view returns(uint256)
func (_IPoolService *IPoolServiceCallerSession) AvailableLiquidity() (*big.Int, error) {
	return _IPoolService.Contract.AvailableLiquidity(&_IPoolService.CallOpts)
}

// BorrowAPYRAY is a free data retrieval call binding the contract method 0x45d31f9d.
//
// Solidity: function borrowAPY_RAY() view returns(uint256)
func (_IPoolService *IPoolServiceCaller) BorrowAPYRAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IPoolService.contract.Call(opts, &out, "borrowAPY_RAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowAPYRAY is a free data retrieval call binding the contract method 0x45d31f9d.
//
// Solidity: function borrowAPY_RAY() view returns(uint256)
func (_IPoolService *IPoolServiceSession) BorrowAPYRAY() (*big.Int, error) {
	return _IPoolService.Contract.BorrowAPYRAY(&_IPoolService.CallOpts)
}

// BorrowAPYRAY is a free data retrieval call binding the contract method 0x45d31f9d.
//
// Solidity: function borrowAPY_RAY() view returns(uint256)
func (_IPoolService *IPoolServiceCallerSession) BorrowAPYRAY() (*big.Int, error) {
	return _IPoolService.Contract.BorrowAPYRAY(&_IPoolService.CallOpts)
}

// CalcCumulativeIndexAtBorrowMore is a free data retrieval call binding the contract method 0xc00495a1.
//
// Solidity: function calcCumulativeIndexAtBorrowMore(uint256 amount, uint256 dAmount, uint256 cumulativeIndexAtOpen) view returns(uint256)
func (_IPoolService *IPoolServiceCaller) CalcCumulativeIndexAtBorrowMore(opts *bind.CallOpts, amount *big.Int, dAmount *big.Int, cumulativeIndexAtOpen *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IPoolService.contract.Call(opts, &out, "calcCumulativeIndexAtBorrowMore", amount, dAmount, cumulativeIndexAtOpen)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcCumulativeIndexAtBorrowMore is a free data retrieval call binding the contract method 0xc00495a1.
//
// Solidity: function calcCumulativeIndexAtBorrowMore(uint256 amount, uint256 dAmount, uint256 cumulativeIndexAtOpen) view returns(uint256)
func (_IPoolService *IPoolServiceSession) CalcCumulativeIndexAtBorrowMore(amount *big.Int, dAmount *big.Int, cumulativeIndexAtOpen *big.Int) (*big.Int, error) {
	return _IPoolService.Contract.CalcCumulativeIndexAtBorrowMore(&_IPoolService.CallOpts, amount, dAmount, cumulativeIndexAtOpen)
}

// CalcCumulativeIndexAtBorrowMore is a free data retrieval call binding the contract method 0xc00495a1.
//
// Solidity: function calcCumulativeIndexAtBorrowMore(uint256 amount, uint256 dAmount, uint256 cumulativeIndexAtOpen) view returns(uint256)
func (_IPoolService *IPoolServiceCallerSession) CalcCumulativeIndexAtBorrowMore(amount *big.Int, dAmount *big.Int, cumulativeIndexAtOpen *big.Int) (*big.Int, error) {
	return _IPoolService.Contract.CalcCumulativeIndexAtBorrowMore(&_IPoolService.CallOpts, amount, dAmount, cumulativeIndexAtOpen)
}

// CalcLinearCumulativeRAY is a free data retrieval call binding the contract method 0x0fce70fb.
//
// Solidity: function calcLinearCumulative_RAY() view returns(uint256)
func (_IPoolService *IPoolServiceCaller) CalcLinearCumulativeRAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IPoolService.contract.Call(opts, &out, "calcLinearCumulative_RAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcLinearCumulativeRAY is a free data retrieval call binding the contract method 0x0fce70fb.
//
// Solidity: function calcLinearCumulative_RAY() view returns(uint256)
func (_IPoolService *IPoolServiceSession) CalcLinearCumulativeRAY() (*big.Int, error) {
	return _IPoolService.Contract.CalcLinearCumulativeRAY(&_IPoolService.CallOpts)
}

// CalcLinearCumulativeRAY is a free data retrieval call binding the contract method 0x0fce70fb.
//
// Solidity: function calcLinearCumulative_RAY() view returns(uint256)
func (_IPoolService *IPoolServiceCallerSession) CalcLinearCumulativeRAY() (*big.Int, error) {
	return _IPoolService.Contract.CalcLinearCumulativeRAY(&_IPoolService.CallOpts)
}

// CreditManagers is a free data retrieval call binding the contract method 0x1e16e4fc.
//
// Solidity: function creditManagers(uint256 id) view returns(address)
func (_IPoolService *IPoolServiceCaller) CreditManagers(opts *bind.CallOpts, id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IPoolService.contract.Call(opts, &out, "creditManagers", id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditManagers is a free data retrieval call binding the contract method 0x1e16e4fc.
//
// Solidity: function creditManagers(uint256 id) view returns(address)
func (_IPoolService *IPoolServiceSession) CreditManagers(id *big.Int) (common.Address, error) {
	return _IPoolService.Contract.CreditManagers(&_IPoolService.CallOpts, id)
}

// CreditManagers is a free data retrieval call binding the contract method 0x1e16e4fc.
//
// Solidity: function creditManagers(uint256 id) view returns(address)
func (_IPoolService *IPoolServiceCallerSession) CreditManagers(id *big.Int) (common.Address, error) {
	return _IPoolService.Contract.CreditManagers(&_IPoolService.CallOpts, id)
}

// CreditManagersCanBorrow is a free data retrieval call binding the contract method 0x2e97ca21.
//
// Solidity: function creditManagersCanBorrow(address id) view returns(bool)
func (_IPoolService *IPoolServiceCaller) CreditManagersCanBorrow(opts *bind.CallOpts, id common.Address) (bool, error) {
	var out []interface{}
	err := _IPoolService.contract.Call(opts, &out, "creditManagersCanBorrow", id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CreditManagersCanBorrow is a free data retrieval call binding the contract method 0x2e97ca21.
//
// Solidity: function creditManagersCanBorrow(address id) view returns(bool)
func (_IPoolService *IPoolServiceSession) CreditManagersCanBorrow(id common.Address) (bool, error) {
	return _IPoolService.Contract.CreditManagersCanBorrow(&_IPoolService.CallOpts, id)
}

// CreditManagersCanBorrow is a free data retrieval call binding the contract method 0x2e97ca21.
//
// Solidity: function creditManagersCanBorrow(address id) view returns(bool)
func (_IPoolService *IPoolServiceCallerSession) CreditManagersCanBorrow(id common.Address) (bool, error) {
	return _IPoolService.Contract.CreditManagersCanBorrow(&_IPoolService.CallOpts, id)
}

// CreditManagersCount is a free data retrieval call binding the contract method 0xa4e8273e.
//
// Solidity: function creditManagersCount() view returns(uint256)
func (_IPoolService *IPoolServiceCaller) CreditManagersCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IPoolService.contract.Call(opts, &out, "creditManagersCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CreditManagersCount is a free data retrieval call binding the contract method 0xa4e8273e.
//
// Solidity: function creditManagersCount() view returns(uint256)
func (_IPoolService *IPoolServiceSession) CreditManagersCount() (*big.Int, error) {
	return _IPoolService.Contract.CreditManagersCount(&_IPoolService.CallOpts)
}

// CreditManagersCount is a free data retrieval call binding the contract method 0xa4e8273e.
//
// Solidity: function creditManagersCount() view returns(uint256)
func (_IPoolService *IPoolServiceCallerSession) CreditManagersCount() (*big.Int, error) {
	return _IPoolService.Contract.CreditManagersCount(&_IPoolService.CallOpts)
}

// DieselToken is a free data retrieval call binding the contract method 0x36dda7d5.
//
// Solidity: function dieselToken() view returns(address)
func (_IPoolService *IPoolServiceCaller) DieselToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IPoolService.contract.Call(opts, &out, "dieselToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DieselToken is a free data retrieval call binding the contract method 0x36dda7d5.
//
// Solidity: function dieselToken() view returns(address)
func (_IPoolService *IPoolServiceSession) DieselToken() (common.Address, error) {
	return _IPoolService.Contract.DieselToken(&_IPoolService.CallOpts)
}

// DieselToken is a free data retrieval call binding the contract method 0x36dda7d5.
//
// Solidity: function dieselToken() view returns(address)
func (_IPoolService *IPoolServiceCallerSession) DieselToken() (common.Address, error) {
	return _IPoolService.Contract.DieselToken(&_IPoolService.CallOpts)
}

// ExpectedLiquidity is a free data retrieval call binding the contract method 0xfe14112d.
//
// Solidity: function expectedLiquidity() view returns(uint256)
func (_IPoolService *IPoolServiceCaller) ExpectedLiquidity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IPoolService.contract.Call(opts, &out, "expectedLiquidity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExpectedLiquidity is a free data retrieval call binding the contract method 0xfe14112d.
//
// Solidity: function expectedLiquidity() view returns(uint256)
func (_IPoolService *IPoolServiceSession) ExpectedLiquidity() (*big.Int, error) {
	return _IPoolService.Contract.ExpectedLiquidity(&_IPoolService.CallOpts)
}

// ExpectedLiquidity is a free data retrieval call binding the contract method 0xfe14112d.
//
// Solidity: function expectedLiquidity() view returns(uint256)
func (_IPoolService *IPoolServiceCallerSession) ExpectedLiquidity() (*big.Int, error) {
	return _IPoolService.Contract.ExpectedLiquidity(&_IPoolService.CallOpts)
}

// ExpectedLiquidityLimit is a free data retrieval call binding the contract method 0xef8d9603.
//
// Solidity: function expectedLiquidityLimit() view returns(uint256)
func (_IPoolService *IPoolServiceCaller) ExpectedLiquidityLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IPoolService.contract.Call(opts, &out, "expectedLiquidityLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExpectedLiquidityLimit is a free data retrieval call binding the contract method 0xef8d9603.
//
// Solidity: function expectedLiquidityLimit() view returns(uint256)
func (_IPoolService *IPoolServiceSession) ExpectedLiquidityLimit() (*big.Int, error) {
	return _IPoolService.Contract.ExpectedLiquidityLimit(&_IPoolService.CallOpts)
}

// ExpectedLiquidityLimit is a free data retrieval call binding the contract method 0xef8d9603.
//
// Solidity: function expectedLiquidityLimit() view returns(uint256)
func (_IPoolService *IPoolServiceCallerSession) ExpectedLiquidityLimit() (*big.Int, error) {
	return _IPoolService.Contract.ExpectedLiquidityLimit(&_IPoolService.CallOpts)
}

// FromDiesel is a free data retrieval call binding the contract method 0x5427c938.
//
// Solidity: function fromDiesel(uint256 amount) view returns(uint256)
func (_IPoolService *IPoolServiceCaller) FromDiesel(opts *bind.CallOpts, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IPoolService.contract.Call(opts, &out, "fromDiesel", amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FromDiesel is a free data retrieval call binding the contract method 0x5427c938.
//
// Solidity: function fromDiesel(uint256 amount) view returns(uint256)
func (_IPoolService *IPoolServiceSession) FromDiesel(amount *big.Int) (*big.Int, error) {
	return _IPoolService.Contract.FromDiesel(&_IPoolService.CallOpts, amount)
}

// FromDiesel is a free data retrieval call binding the contract method 0x5427c938.
//
// Solidity: function fromDiesel(uint256 amount) view returns(uint256)
func (_IPoolService *IPoolServiceCallerSession) FromDiesel(amount *big.Int) (*big.Int, error) {
	return _IPoolService.Contract.FromDiesel(&_IPoolService.CallOpts, amount)
}

// GetDieselRateRAY is a free data retrieval call binding the contract method 0x788c6bfe.
//
// Solidity: function getDieselRate_RAY() view returns(uint256)
func (_IPoolService *IPoolServiceCaller) GetDieselRateRAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IPoolService.contract.Call(opts, &out, "getDieselRate_RAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDieselRateRAY is a free data retrieval call binding the contract method 0x788c6bfe.
//
// Solidity: function getDieselRate_RAY() view returns(uint256)
func (_IPoolService *IPoolServiceSession) GetDieselRateRAY() (*big.Int, error) {
	return _IPoolService.Contract.GetDieselRateRAY(&_IPoolService.CallOpts)
}

// GetDieselRateRAY is a free data retrieval call binding the contract method 0x788c6bfe.
//
// Solidity: function getDieselRate_RAY() view returns(uint256)
func (_IPoolService *IPoolServiceCallerSession) GetDieselRateRAY() (*big.Int, error) {
	return _IPoolService.Contract.GetDieselRateRAY(&_IPoolService.CallOpts)
}

// ToDiesel is a free data retrieval call binding the contract method 0x4d778ad1.
//
// Solidity: function toDiesel(uint256 amount) view returns(uint256)
func (_IPoolService *IPoolServiceCaller) ToDiesel(opts *bind.CallOpts, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IPoolService.contract.Call(opts, &out, "toDiesel", amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ToDiesel is a free data retrieval call binding the contract method 0x4d778ad1.
//
// Solidity: function toDiesel(uint256 amount) view returns(uint256)
func (_IPoolService *IPoolServiceSession) ToDiesel(amount *big.Int) (*big.Int, error) {
	return _IPoolService.Contract.ToDiesel(&_IPoolService.CallOpts, amount)
}

// ToDiesel is a free data retrieval call binding the contract method 0x4d778ad1.
//
// Solidity: function toDiesel(uint256 amount) view returns(uint256)
func (_IPoolService *IPoolServiceCallerSession) ToDiesel(amount *big.Int) (*big.Int, error) {
	return _IPoolService.Contract.ToDiesel(&_IPoolService.CallOpts, amount)
}

// TotalBorrowed is a free data retrieval call binding the contract method 0x4c19386c.
//
// Solidity: function totalBorrowed() view returns(uint256)
func (_IPoolService *IPoolServiceCaller) TotalBorrowed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IPoolService.contract.Call(opts, &out, "totalBorrowed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBorrowed is a free data retrieval call binding the contract method 0x4c19386c.
//
// Solidity: function totalBorrowed() view returns(uint256)
func (_IPoolService *IPoolServiceSession) TotalBorrowed() (*big.Int, error) {
	return _IPoolService.Contract.TotalBorrowed(&_IPoolService.CallOpts)
}

// TotalBorrowed is a free data retrieval call binding the contract method 0x4c19386c.
//
// Solidity: function totalBorrowed() view returns(uint256)
func (_IPoolService *IPoolServiceCallerSession) TotalBorrowed() (*big.Int, error) {
	return _IPoolService.Contract.TotalBorrowed(&_IPoolService.CallOpts)
}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_IPoolService *IPoolServiceCaller) UnderlyingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IPoolService.contract.Call(opts, &out, "underlyingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_IPoolService *IPoolServiceSession) UnderlyingToken() (common.Address, error) {
	return _IPoolService.Contract.UnderlyingToken(&_IPoolService.CallOpts)
}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_IPoolService *IPoolServiceCallerSession) UnderlyingToken() (common.Address, error) {
	return _IPoolService.Contract.UnderlyingToken(&_IPoolService.CallOpts)
}

// WithdrawFee is a free data retrieval call binding the contract method 0xe941fa78.
//
// Solidity: function withdrawFee() view returns(uint256)
func (_IPoolService *IPoolServiceCaller) WithdrawFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IPoolService.contract.Call(opts, &out, "withdrawFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawFee is a free data retrieval call binding the contract method 0xe941fa78.
//
// Solidity: function withdrawFee() view returns(uint256)
func (_IPoolService *IPoolServiceSession) WithdrawFee() (*big.Int, error) {
	return _IPoolService.Contract.WithdrawFee(&_IPoolService.CallOpts)
}

// WithdrawFee is a free data retrieval call binding the contract method 0xe941fa78.
//
// Solidity: function withdrawFee() view returns(uint256)
func (_IPoolService *IPoolServiceCallerSession) WithdrawFee() (*big.Int, error) {
	return _IPoolService.Contract.WithdrawFee(&_IPoolService.CallOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x9aa5d462.
//
// Solidity: function addLiquidity(uint256 amount, address onBehalfOf, uint256 referralCode) returns()
func (_IPoolService *IPoolServiceTransactor) AddLiquidity(opts *bind.TransactOpts, amount *big.Int, onBehalfOf common.Address, referralCode *big.Int) (*types.Transaction, error) {
	return _IPoolService.contract.Transact(opts, "addLiquidity", amount, onBehalfOf, referralCode)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x9aa5d462.
//
// Solidity: function addLiquidity(uint256 amount, address onBehalfOf, uint256 referralCode) returns()
func (_IPoolService *IPoolServiceSession) AddLiquidity(amount *big.Int, onBehalfOf common.Address, referralCode *big.Int) (*types.Transaction, error) {
	return _IPoolService.Contract.AddLiquidity(&_IPoolService.TransactOpts, amount, onBehalfOf, referralCode)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x9aa5d462.
//
// Solidity: function addLiquidity(uint256 amount, address onBehalfOf, uint256 referralCode) returns()
func (_IPoolService *IPoolServiceTransactorSession) AddLiquidity(amount *big.Int, onBehalfOf common.Address, referralCode *big.Int) (*types.Transaction, error) {
	return _IPoolService.Contract.AddLiquidity(&_IPoolService.TransactOpts, amount, onBehalfOf, referralCode)
}

// LendCreditAccount is a paid mutator transaction binding the contract method 0xbf28068b.
//
// Solidity: function lendCreditAccount(uint256 borrowedAmount, address creditAccount) returns()
func (_IPoolService *IPoolServiceTransactor) LendCreditAccount(opts *bind.TransactOpts, borrowedAmount *big.Int, creditAccount common.Address) (*types.Transaction, error) {
	return _IPoolService.contract.Transact(opts, "lendCreditAccount", borrowedAmount, creditAccount)
}

// LendCreditAccount is a paid mutator transaction binding the contract method 0xbf28068b.
//
// Solidity: function lendCreditAccount(uint256 borrowedAmount, address creditAccount) returns()
func (_IPoolService *IPoolServiceSession) LendCreditAccount(borrowedAmount *big.Int, creditAccount common.Address) (*types.Transaction, error) {
	return _IPoolService.Contract.LendCreditAccount(&_IPoolService.TransactOpts, borrowedAmount, creditAccount)
}

// LendCreditAccount is a paid mutator transaction binding the contract method 0xbf28068b.
//
// Solidity: function lendCreditAccount(uint256 borrowedAmount, address creditAccount) returns()
func (_IPoolService *IPoolServiceTransactorSession) LendCreditAccount(borrowedAmount *big.Int, creditAccount common.Address) (*types.Transaction, error) {
	return _IPoolService.Contract.LendCreditAccount(&_IPoolService.TransactOpts, borrowedAmount, creditAccount)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x05fe138b.
//
// Solidity: function removeLiquidity(uint256 amount, address to) returns(uint256)
func (_IPoolService *IPoolServiceTransactor) RemoveLiquidity(opts *bind.TransactOpts, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _IPoolService.contract.Transact(opts, "removeLiquidity", amount, to)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x05fe138b.
//
// Solidity: function removeLiquidity(uint256 amount, address to) returns(uint256)
func (_IPoolService *IPoolServiceSession) RemoveLiquidity(amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _IPoolService.Contract.RemoveLiquidity(&_IPoolService.TransactOpts, amount, to)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x05fe138b.
//
// Solidity: function removeLiquidity(uint256 amount, address to) returns(uint256)
func (_IPoolService *IPoolServiceTransactorSession) RemoveLiquidity(amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _IPoolService.Contract.RemoveLiquidity(&_IPoolService.TransactOpts, amount, to)
}

// RepayCreditAccount is a paid mutator transaction binding the contract method 0xca9505e4.
//
// Solidity: function repayCreditAccount(uint256 borrowedAmount, uint256 profit, uint256 loss) returns()
func (_IPoolService *IPoolServiceTransactor) RepayCreditAccount(opts *bind.TransactOpts, borrowedAmount *big.Int, profit *big.Int, loss *big.Int) (*types.Transaction, error) {
	return _IPoolService.contract.Transact(opts, "repayCreditAccount", borrowedAmount, profit, loss)
}

// RepayCreditAccount is a paid mutator transaction binding the contract method 0xca9505e4.
//
// Solidity: function repayCreditAccount(uint256 borrowedAmount, uint256 profit, uint256 loss) returns()
func (_IPoolService *IPoolServiceSession) RepayCreditAccount(borrowedAmount *big.Int, profit *big.Int, loss *big.Int) (*types.Transaction, error) {
	return _IPoolService.Contract.RepayCreditAccount(&_IPoolService.TransactOpts, borrowedAmount, profit, loss)
}

// RepayCreditAccount is a paid mutator transaction binding the contract method 0xca9505e4.
//
// Solidity: function repayCreditAccount(uint256 borrowedAmount, uint256 profit, uint256 loss) returns()
func (_IPoolService *IPoolServiceTransactorSession) RepayCreditAccount(borrowedAmount *big.Int, profit *big.Int, loss *big.Int) (*types.Transaction, error) {
	return _IPoolService.Contract.RepayCreditAccount(&_IPoolService.TransactOpts, borrowedAmount, profit, loss)
}

// IPoolServiceAddLiquidityIterator is returned from FilterAddLiquidity and is used to iterate over the raw logs and unpacked data for AddLiquidity events raised by the IPoolService contract.
type IPoolServiceAddLiquidityIterator struct {
	Event *IPoolServiceAddLiquidity // Event containing the contract specifics and raw log

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
func (it *IPoolServiceAddLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPoolServiceAddLiquidity)
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
		it.Event = new(IPoolServiceAddLiquidity)
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
func (it *IPoolServiceAddLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPoolServiceAddLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPoolServiceAddLiquidity represents a AddLiquidity event raised by the IPoolService contract.
type IPoolServiceAddLiquidity struct {
	Sender       common.Address
	OnBehalfOf   common.Address
	Amount       *big.Int
	ReferralCode *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAddLiquidity is a free log retrieval operation binding the contract event 0xd2491a9b4fe81a7cd4511e8b7b7743951b061dad5bed7da8a7795b080ee08c7e.
//
// Solidity: event AddLiquidity(address indexed sender, address indexed onBehalfOf, uint256 amount, uint256 referralCode)
func (_IPoolService *IPoolServiceFilterer) FilterAddLiquidity(opts *bind.FilterOpts, sender []common.Address, onBehalfOf []common.Address) (*IPoolServiceAddLiquidityIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	logs, sub, err := _IPoolService.contract.FilterLogs(opts, "AddLiquidity", senderRule, onBehalfOfRule)
	if err != nil {
		return nil, err
	}
	return &IPoolServiceAddLiquidityIterator{contract: _IPoolService.contract, event: "AddLiquidity", logs: logs, sub: sub}, nil
}

// WatchAddLiquidity is a free log subscription operation binding the contract event 0xd2491a9b4fe81a7cd4511e8b7b7743951b061dad5bed7da8a7795b080ee08c7e.
//
// Solidity: event AddLiquidity(address indexed sender, address indexed onBehalfOf, uint256 amount, uint256 referralCode)
func (_IPoolService *IPoolServiceFilterer) WatchAddLiquidity(opts *bind.WatchOpts, sink chan<- *IPoolServiceAddLiquidity, sender []common.Address, onBehalfOf []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	logs, sub, err := _IPoolService.contract.WatchLogs(opts, "AddLiquidity", senderRule, onBehalfOfRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPoolServiceAddLiquidity)
				if err := _IPoolService.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
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

// ParseAddLiquidity is a log parse operation binding the contract event 0xd2491a9b4fe81a7cd4511e8b7b7743951b061dad5bed7da8a7795b080ee08c7e.
//
// Solidity: event AddLiquidity(address indexed sender, address indexed onBehalfOf, uint256 amount, uint256 referralCode)
func (_IPoolService *IPoolServiceFilterer) ParseAddLiquidity(log types.Log) (*IPoolServiceAddLiquidity, error) {
	event := new(IPoolServiceAddLiquidity)
	if err := _IPoolService.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPoolServiceBorrowIterator is returned from FilterBorrow and is used to iterate over the raw logs and unpacked data for Borrow events raised by the IPoolService contract.
type IPoolServiceBorrowIterator struct {
	Event *IPoolServiceBorrow // Event containing the contract specifics and raw log

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
func (it *IPoolServiceBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPoolServiceBorrow)
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
		it.Event = new(IPoolServiceBorrow)
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
func (it *IPoolServiceBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPoolServiceBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPoolServiceBorrow represents a Borrow event raised by the IPoolService contract.
type IPoolServiceBorrow struct {
	CreditManager common.Address
	CreditAccount common.Address
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBorrow is a free log retrieval operation binding the contract event 0x312a5e5e1079f5dda4e95dbbd0b908b291fd5b992ef22073643ab691572c5b52.
//
// Solidity: event Borrow(address indexed creditManager, address indexed creditAccount, uint256 amount)
func (_IPoolService *IPoolServiceFilterer) FilterBorrow(opts *bind.FilterOpts, creditManager []common.Address, creditAccount []common.Address) (*IPoolServiceBorrowIterator, error) {

	var creditManagerRule []interface{}
	for _, creditManagerItem := range creditManager {
		creditManagerRule = append(creditManagerRule, creditManagerItem)
	}
	var creditAccountRule []interface{}
	for _, creditAccountItem := range creditAccount {
		creditAccountRule = append(creditAccountRule, creditAccountItem)
	}

	logs, sub, err := _IPoolService.contract.FilterLogs(opts, "Borrow", creditManagerRule, creditAccountRule)
	if err != nil {
		return nil, err
	}
	return &IPoolServiceBorrowIterator{contract: _IPoolService.contract, event: "Borrow", logs: logs, sub: sub}, nil
}

// WatchBorrow is a free log subscription operation binding the contract event 0x312a5e5e1079f5dda4e95dbbd0b908b291fd5b992ef22073643ab691572c5b52.
//
// Solidity: event Borrow(address indexed creditManager, address indexed creditAccount, uint256 amount)
func (_IPoolService *IPoolServiceFilterer) WatchBorrow(opts *bind.WatchOpts, sink chan<- *IPoolServiceBorrow, creditManager []common.Address, creditAccount []common.Address) (event.Subscription, error) {

	var creditManagerRule []interface{}
	for _, creditManagerItem := range creditManager {
		creditManagerRule = append(creditManagerRule, creditManagerItem)
	}
	var creditAccountRule []interface{}
	for _, creditAccountItem := range creditAccount {
		creditAccountRule = append(creditAccountRule, creditAccountItem)
	}

	logs, sub, err := _IPoolService.contract.WatchLogs(opts, "Borrow", creditManagerRule, creditAccountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPoolServiceBorrow)
				if err := _IPoolService.contract.UnpackLog(event, "Borrow", log); err != nil {
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

// ParseBorrow is a log parse operation binding the contract event 0x312a5e5e1079f5dda4e95dbbd0b908b291fd5b992ef22073643ab691572c5b52.
//
// Solidity: event Borrow(address indexed creditManager, address indexed creditAccount, uint256 amount)
func (_IPoolService *IPoolServiceFilterer) ParseBorrow(log types.Log) (*IPoolServiceBorrow, error) {
	event := new(IPoolServiceBorrow)
	if err := _IPoolService.contract.UnpackLog(event, "Borrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPoolServiceBorrowForbiddenIterator is returned from FilterBorrowForbidden and is used to iterate over the raw logs and unpacked data for BorrowForbidden events raised by the IPoolService contract.
type IPoolServiceBorrowForbiddenIterator struct {
	Event *IPoolServiceBorrowForbidden // Event containing the contract specifics and raw log

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
func (it *IPoolServiceBorrowForbiddenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPoolServiceBorrowForbidden)
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
		it.Event = new(IPoolServiceBorrowForbidden)
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
func (it *IPoolServiceBorrowForbiddenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPoolServiceBorrowForbiddenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPoolServiceBorrowForbidden represents a BorrowForbidden event raised by the IPoolService contract.
type IPoolServiceBorrowForbidden struct {
	CreditManager common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBorrowForbidden is a free log retrieval operation binding the contract event 0x9181736fce85d2d4cca2e4406f10679302ae5c387180fdb62963af3cd9a24fd6.
//
// Solidity: event BorrowForbidden(address indexed creditManager)
func (_IPoolService *IPoolServiceFilterer) FilterBorrowForbidden(opts *bind.FilterOpts, creditManager []common.Address) (*IPoolServiceBorrowForbiddenIterator, error) {

	var creditManagerRule []interface{}
	for _, creditManagerItem := range creditManager {
		creditManagerRule = append(creditManagerRule, creditManagerItem)
	}

	logs, sub, err := _IPoolService.contract.FilterLogs(opts, "BorrowForbidden", creditManagerRule)
	if err != nil {
		return nil, err
	}
	return &IPoolServiceBorrowForbiddenIterator{contract: _IPoolService.contract, event: "BorrowForbidden", logs: logs, sub: sub}, nil
}

// WatchBorrowForbidden is a free log subscription operation binding the contract event 0x9181736fce85d2d4cca2e4406f10679302ae5c387180fdb62963af3cd9a24fd6.
//
// Solidity: event BorrowForbidden(address indexed creditManager)
func (_IPoolService *IPoolServiceFilterer) WatchBorrowForbidden(opts *bind.WatchOpts, sink chan<- *IPoolServiceBorrowForbidden, creditManager []common.Address) (event.Subscription, error) {

	var creditManagerRule []interface{}
	for _, creditManagerItem := range creditManager {
		creditManagerRule = append(creditManagerRule, creditManagerItem)
	}

	logs, sub, err := _IPoolService.contract.WatchLogs(opts, "BorrowForbidden", creditManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPoolServiceBorrowForbidden)
				if err := _IPoolService.contract.UnpackLog(event, "BorrowForbidden", log); err != nil {
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

// ParseBorrowForbidden is a log parse operation binding the contract event 0x9181736fce85d2d4cca2e4406f10679302ae5c387180fdb62963af3cd9a24fd6.
//
// Solidity: event BorrowForbidden(address indexed creditManager)
func (_IPoolService *IPoolServiceFilterer) ParseBorrowForbidden(log types.Log) (*IPoolServiceBorrowForbidden, error) {
	event := new(IPoolServiceBorrowForbidden)
	if err := _IPoolService.contract.UnpackLog(event, "BorrowForbidden", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPoolServiceNewCreditManagerConnectedIterator is returned from FilterNewCreditManagerConnected and is used to iterate over the raw logs and unpacked data for NewCreditManagerConnected events raised by the IPoolService contract.
type IPoolServiceNewCreditManagerConnectedIterator struct {
	Event *IPoolServiceNewCreditManagerConnected // Event containing the contract specifics and raw log

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
func (it *IPoolServiceNewCreditManagerConnectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPoolServiceNewCreditManagerConnected)
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
		it.Event = new(IPoolServiceNewCreditManagerConnected)
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
func (it *IPoolServiceNewCreditManagerConnectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPoolServiceNewCreditManagerConnectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPoolServiceNewCreditManagerConnected represents a NewCreditManagerConnected event raised by the IPoolService contract.
type IPoolServiceNewCreditManagerConnected struct {
	CreditManager common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNewCreditManagerConnected is a free log retrieval operation binding the contract event 0xe076020e7eac3915d33aec40c24f95e73eb6c9921ff89747d50aa8fd934d2c01.
//
// Solidity: event NewCreditManagerConnected(address indexed creditManager)
func (_IPoolService *IPoolServiceFilterer) FilterNewCreditManagerConnected(opts *bind.FilterOpts, creditManager []common.Address) (*IPoolServiceNewCreditManagerConnectedIterator, error) {

	var creditManagerRule []interface{}
	for _, creditManagerItem := range creditManager {
		creditManagerRule = append(creditManagerRule, creditManagerItem)
	}

	logs, sub, err := _IPoolService.contract.FilterLogs(opts, "NewCreditManagerConnected", creditManagerRule)
	if err != nil {
		return nil, err
	}
	return &IPoolServiceNewCreditManagerConnectedIterator{contract: _IPoolService.contract, event: "NewCreditManagerConnected", logs: logs, sub: sub}, nil
}

// WatchNewCreditManagerConnected is a free log subscription operation binding the contract event 0xe076020e7eac3915d33aec40c24f95e73eb6c9921ff89747d50aa8fd934d2c01.
//
// Solidity: event NewCreditManagerConnected(address indexed creditManager)
func (_IPoolService *IPoolServiceFilterer) WatchNewCreditManagerConnected(opts *bind.WatchOpts, sink chan<- *IPoolServiceNewCreditManagerConnected, creditManager []common.Address) (event.Subscription, error) {

	var creditManagerRule []interface{}
	for _, creditManagerItem := range creditManager {
		creditManagerRule = append(creditManagerRule, creditManagerItem)
	}

	logs, sub, err := _IPoolService.contract.WatchLogs(opts, "NewCreditManagerConnected", creditManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPoolServiceNewCreditManagerConnected)
				if err := _IPoolService.contract.UnpackLog(event, "NewCreditManagerConnected", log); err != nil {
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

// ParseNewCreditManagerConnected is a log parse operation binding the contract event 0xe076020e7eac3915d33aec40c24f95e73eb6c9921ff89747d50aa8fd934d2c01.
//
// Solidity: event NewCreditManagerConnected(address indexed creditManager)
func (_IPoolService *IPoolServiceFilterer) ParseNewCreditManagerConnected(log types.Log) (*IPoolServiceNewCreditManagerConnected, error) {
	event := new(IPoolServiceNewCreditManagerConnected)
	if err := _IPoolService.contract.UnpackLog(event, "NewCreditManagerConnected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPoolServiceNewExpectedLiquidityLimitIterator is returned from FilterNewExpectedLiquidityLimit and is used to iterate over the raw logs and unpacked data for NewExpectedLiquidityLimit events raised by the IPoolService contract.
type IPoolServiceNewExpectedLiquidityLimitIterator struct {
	Event *IPoolServiceNewExpectedLiquidityLimit // Event containing the contract specifics and raw log

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
func (it *IPoolServiceNewExpectedLiquidityLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPoolServiceNewExpectedLiquidityLimit)
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
		it.Event = new(IPoolServiceNewExpectedLiquidityLimit)
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
func (it *IPoolServiceNewExpectedLiquidityLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPoolServiceNewExpectedLiquidityLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPoolServiceNewExpectedLiquidityLimit represents a NewExpectedLiquidityLimit event raised by the IPoolService contract.
type IPoolServiceNewExpectedLiquidityLimit struct {
	NewLimit *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewExpectedLiquidityLimit is a free log retrieval operation binding the contract event 0xd7a183c9fe85b604c25d54bd676e0866f6c13bcca9fb9b0850213de118fdc99c.
//
// Solidity: event NewExpectedLiquidityLimit(uint256 newLimit)
func (_IPoolService *IPoolServiceFilterer) FilterNewExpectedLiquidityLimit(opts *bind.FilterOpts) (*IPoolServiceNewExpectedLiquidityLimitIterator, error) {

	logs, sub, err := _IPoolService.contract.FilterLogs(opts, "NewExpectedLiquidityLimit")
	if err != nil {
		return nil, err
	}
	return &IPoolServiceNewExpectedLiquidityLimitIterator{contract: _IPoolService.contract, event: "NewExpectedLiquidityLimit", logs: logs, sub: sub}, nil
}

// WatchNewExpectedLiquidityLimit is a free log subscription operation binding the contract event 0xd7a183c9fe85b604c25d54bd676e0866f6c13bcca9fb9b0850213de118fdc99c.
//
// Solidity: event NewExpectedLiquidityLimit(uint256 newLimit)
func (_IPoolService *IPoolServiceFilterer) WatchNewExpectedLiquidityLimit(opts *bind.WatchOpts, sink chan<- *IPoolServiceNewExpectedLiquidityLimit) (event.Subscription, error) {

	logs, sub, err := _IPoolService.contract.WatchLogs(opts, "NewExpectedLiquidityLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPoolServiceNewExpectedLiquidityLimit)
				if err := _IPoolService.contract.UnpackLog(event, "NewExpectedLiquidityLimit", log); err != nil {
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

// ParseNewExpectedLiquidityLimit is a log parse operation binding the contract event 0xd7a183c9fe85b604c25d54bd676e0866f6c13bcca9fb9b0850213de118fdc99c.
//
// Solidity: event NewExpectedLiquidityLimit(uint256 newLimit)
func (_IPoolService *IPoolServiceFilterer) ParseNewExpectedLiquidityLimit(log types.Log) (*IPoolServiceNewExpectedLiquidityLimit, error) {
	event := new(IPoolServiceNewExpectedLiquidityLimit)
	if err := _IPoolService.contract.UnpackLog(event, "NewExpectedLiquidityLimit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPoolServiceNewInterestRateModelIterator is returned from FilterNewInterestRateModel and is used to iterate over the raw logs and unpacked data for NewInterestRateModel events raised by the IPoolService contract.
type IPoolServiceNewInterestRateModelIterator struct {
	Event *IPoolServiceNewInterestRateModel // Event containing the contract specifics and raw log

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
func (it *IPoolServiceNewInterestRateModelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPoolServiceNewInterestRateModel)
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
		it.Event = new(IPoolServiceNewInterestRateModel)
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
func (it *IPoolServiceNewInterestRateModelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPoolServiceNewInterestRateModelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPoolServiceNewInterestRateModel represents a NewInterestRateModel event raised by the IPoolService contract.
type IPoolServiceNewInterestRateModel struct {
	NewInterestRateModel common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterNewInterestRateModel is a free log retrieval operation binding the contract event 0x0ec6cb7631d36954a05ffd646135bfd9995c71e7fa36d26abb1ad9f24a040ea1.
//
// Solidity: event NewInterestRateModel(address indexed newInterestRateModel)
func (_IPoolService *IPoolServiceFilterer) FilterNewInterestRateModel(opts *bind.FilterOpts, newInterestRateModel []common.Address) (*IPoolServiceNewInterestRateModelIterator, error) {

	var newInterestRateModelRule []interface{}
	for _, newInterestRateModelItem := range newInterestRateModel {
		newInterestRateModelRule = append(newInterestRateModelRule, newInterestRateModelItem)
	}

	logs, sub, err := _IPoolService.contract.FilterLogs(opts, "NewInterestRateModel", newInterestRateModelRule)
	if err != nil {
		return nil, err
	}
	return &IPoolServiceNewInterestRateModelIterator{contract: _IPoolService.contract, event: "NewInterestRateModel", logs: logs, sub: sub}, nil
}

// WatchNewInterestRateModel is a free log subscription operation binding the contract event 0x0ec6cb7631d36954a05ffd646135bfd9995c71e7fa36d26abb1ad9f24a040ea1.
//
// Solidity: event NewInterestRateModel(address indexed newInterestRateModel)
func (_IPoolService *IPoolServiceFilterer) WatchNewInterestRateModel(opts *bind.WatchOpts, sink chan<- *IPoolServiceNewInterestRateModel, newInterestRateModel []common.Address) (event.Subscription, error) {

	var newInterestRateModelRule []interface{}
	for _, newInterestRateModelItem := range newInterestRateModel {
		newInterestRateModelRule = append(newInterestRateModelRule, newInterestRateModelItem)
	}

	logs, sub, err := _IPoolService.contract.WatchLogs(opts, "NewInterestRateModel", newInterestRateModelRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPoolServiceNewInterestRateModel)
				if err := _IPoolService.contract.UnpackLog(event, "NewInterestRateModel", log); err != nil {
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

// ParseNewInterestRateModel is a log parse operation binding the contract event 0x0ec6cb7631d36954a05ffd646135bfd9995c71e7fa36d26abb1ad9f24a040ea1.
//
// Solidity: event NewInterestRateModel(address indexed newInterestRateModel)
func (_IPoolService *IPoolServiceFilterer) ParseNewInterestRateModel(log types.Log) (*IPoolServiceNewInterestRateModel, error) {
	event := new(IPoolServiceNewInterestRateModel)
	if err := _IPoolService.contract.UnpackLog(event, "NewInterestRateModel", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPoolServiceNewWithdrawFeeIterator is returned from FilterNewWithdrawFee and is used to iterate over the raw logs and unpacked data for NewWithdrawFee events raised by the IPoolService contract.
type IPoolServiceNewWithdrawFeeIterator struct {
	Event *IPoolServiceNewWithdrawFee // Event containing the contract specifics and raw log

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
func (it *IPoolServiceNewWithdrawFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPoolServiceNewWithdrawFee)
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
		it.Event = new(IPoolServiceNewWithdrawFee)
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
func (it *IPoolServiceNewWithdrawFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPoolServiceNewWithdrawFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPoolServiceNewWithdrawFee represents a NewWithdrawFee event raised by the IPoolService contract.
type IPoolServiceNewWithdrawFee struct {
	Fee *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterNewWithdrawFee is a free log retrieval operation binding the contract event 0xd5fe46099fa396290a7f57e36c3c3c8774e2562c18ed5d1dcc0fa75071e03f1d.
//
// Solidity: event NewWithdrawFee(uint256 fee)
func (_IPoolService *IPoolServiceFilterer) FilterNewWithdrawFee(opts *bind.FilterOpts) (*IPoolServiceNewWithdrawFeeIterator, error) {

	logs, sub, err := _IPoolService.contract.FilterLogs(opts, "NewWithdrawFee")
	if err != nil {
		return nil, err
	}
	return &IPoolServiceNewWithdrawFeeIterator{contract: _IPoolService.contract, event: "NewWithdrawFee", logs: logs, sub: sub}, nil
}

// WatchNewWithdrawFee is a free log subscription operation binding the contract event 0xd5fe46099fa396290a7f57e36c3c3c8774e2562c18ed5d1dcc0fa75071e03f1d.
//
// Solidity: event NewWithdrawFee(uint256 fee)
func (_IPoolService *IPoolServiceFilterer) WatchNewWithdrawFee(opts *bind.WatchOpts, sink chan<- *IPoolServiceNewWithdrawFee) (event.Subscription, error) {

	logs, sub, err := _IPoolService.contract.WatchLogs(opts, "NewWithdrawFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPoolServiceNewWithdrawFee)
				if err := _IPoolService.contract.UnpackLog(event, "NewWithdrawFee", log); err != nil {
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

// ParseNewWithdrawFee is a log parse operation binding the contract event 0xd5fe46099fa396290a7f57e36c3c3c8774e2562c18ed5d1dcc0fa75071e03f1d.
//
// Solidity: event NewWithdrawFee(uint256 fee)
func (_IPoolService *IPoolServiceFilterer) ParseNewWithdrawFee(log types.Log) (*IPoolServiceNewWithdrawFee, error) {
	event := new(IPoolServiceNewWithdrawFee)
	if err := _IPoolService.contract.UnpackLog(event, "NewWithdrawFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPoolServiceRemoveLiquidityIterator is returned from FilterRemoveLiquidity and is used to iterate over the raw logs and unpacked data for RemoveLiquidity events raised by the IPoolService contract.
type IPoolServiceRemoveLiquidityIterator struct {
	Event *IPoolServiceRemoveLiquidity // Event containing the contract specifics and raw log

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
func (it *IPoolServiceRemoveLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPoolServiceRemoveLiquidity)
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
		it.Event = new(IPoolServiceRemoveLiquidity)
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
func (it *IPoolServiceRemoveLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPoolServiceRemoveLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPoolServiceRemoveLiquidity represents a RemoveLiquidity event raised by the IPoolService contract.
type IPoolServiceRemoveLiquidity struct {
	Sender common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidity is a free log retrieval operation binding the contract event 0xd8ae9b9ba89e637bcb66a69ac91e8f688018e81d6f92c57e02226425c8efbdf6.
//
// Solidity: event RemoveLiquidity(address indexed sender, address indexed to, uint256 amount)
func (_IPoolService *IPoolServiceFilterer) FilterRemoveLiquidity(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*IPoolServiceRemoveLiquidityIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IPoolService.contract.FilterLogs(opts, "RemoveLiquidity", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IPoolServiceRemoveLiquidityIterator{contract: _IPoolService.contract, event: "RemoveLiquidity", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidity is a free log subscription operation binding the contract event 0xd8ae9b9ba89e637bcb66a69ac91e8f688018e81d6f92c57e02226425c8efbdf6.
//
// Solidity: event RemoveLiquidity(address indexed sender, address indexed to, uint256 amount)
func (_IPoolService *IPoolServiceFilterer) WatchRemoveLiquidity(opts *bind.WatchOpts, sink chan<- *IPoolServiceRemoveLiquidity, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IPoolService.contract.WatchLogs(opts, "RemoveLiquidity", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPoolServiceRemoveLiquidity)
				if err := _IPoolService.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
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

// ParseRemoveLiquidity is a log parse operation binding the contract event 0xd8ae9b9ba89e637bcb66a69ac91e8f688018e81d6f92c57e02226425c8efbdf6.
//
// Solidity: event RemoveLiquidity(address indexed sender, address indexed to, uint256 amount)
func (_IPoolService *IPoolServiceFilterer) ParseRemoveLiquidity(log types.Log) (*IPoolServiceRemoveLiquidity, error) {
	event := new(IPoolServiceRemoveLiquidity)
	if err := _IPoolService.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPoolServiceRepayIterator is returned from FilterRepay and is used to iterate over the raw logs and unpacked data for Repay events raised by the IPoolService contract.
type IPoolServiceRepayIterator struct {
	Event *IPoolServiceRepay // Event containing the contract specifics and raw log

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
func (it *IPoolServiceRepayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPoolServiceRepay)
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
		it.Event = new(IPoolServiceRepay)
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
func (it *IPoolServiceRepayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPoolServiceRepayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPoolServiceRepay represents a Repay event raised by the IPoolService contract.
type IPoolServiceRepay struct {
	CreditManager  common.Address
	BorrowedAmount *big.Int
	Profit         *big.Int
	Loss           *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRepay is a free log retrieval operation binding the contract event 0x2fe77b1c99aca6b022b8efc6e3e8dd1b48b30748709339b65c50ef3263443e09.
//
// Solidity: event Repay(address indexed creditManager, uint256 borrowedAmount, uint256 profit, uint256 loss)
func (_IPoolService *IPoolServiceFilterer) FilterRepay(opts *bind.FilterOpts, creditManager []common.Address) (*IPoolServiceRepayIterator, error) {

	var creditManagerRule []interface{}
	for _, creditManagerItem := range creditManager {
		creditManagerRule = append(creditManagerRule, creditManagerItem)
	}

	logs, sub, err := _IPoolService.contract.FilterLogs(opts, "Repay", creditManagerRule)
	if err != nil {
		return nil, err
	}
	return &IPoolServiceRepayIterator{contract: _IPoolService.contract, event: "Repay", logs: logs, sub: sub}, nil
}

// WatchRepay is a free log subscription operation binding the contract event 0x2fe77b1c99aca6b022b8efc6e3e8dd1b48b30748709339b65c50ef3263443e09.
//
// Solidity: event Repay(address indexed creditManager, uint256 borrowedAmount, uint256 profit, uint256 loss)
func (_IPoolService *IPoolServiceFilterer) WatchRepay(opts *bind.WatchOpts, sink chan<- *IPoolServiceRepay, creditManager []common.Address) (event.Subscription, error) {

	var creditManagerRule []interface{}
	for _, creditManagerItem := range creditManager {
		creditManagerRule = append(creditManagerRule, creditManagerItem)
	}

	logs, sub, err := _IPoolService.contract.WatchLogs(opts, "Repay", creditManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPoolServiceRepay)
				if err := _IPoolService.contract.UnpackLog(event, "Repay", log); err != nil {
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

// ParseRepay is a log parse operation binding the contract event 0x2fe77b1c99aca6b022b8efc6e3e8dd1b48b30748709339b65c50ef3263443e09.
//
// Solidity: event Repay(address indexed creditManager, uint256 borrowedAmount, uint256 profit, uint256 loss)
func (_IPoolService *IPoolServiceFilterer) ParseRepay(log types.Log) (*IPoolServiceRepay, error) {
	event := new(IPoolServiceRepay)
	if err := _IPoolService.contract.UnpackLog(event, "Repay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPoolServiceUncoveredLossIterator is returned from FilterUncoveredLoss and is used to iterate over the raw logs and unpacked data for UncoveredLoss events raised by the IPoolService contract.
type IPoolServiceUncoveredLossIterator struct {
	Event *IPoolServiceUncoveredLoss // Event containing the contract specifics and raw log

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
func (it *IPoolServiceUncoveredLossIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPoolServiceUncoveredLoss)
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
		it.Event = new(IPoolServiceUncoveredLoss)
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
func (it *IPoolServiceUncoveredLossIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPoolServiceUncoveredLossIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPoolServiceUncoveredLoss represents a UncoveredLoss event raised by the IPoolService contract.
type IPoolServiceUncoveredLoss struct {
	CreditManager common.Address
	Loss          *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUncoveredLoss is a free log retrieval operation binding the contract event 0xef3653ded679720ab04913b6f3820be7cedc8286d42ff5dd8dff17e91bd2964c.
//
// Solidity: event UncoveredLoss(address indexed creditManager, uint256 loss)
func (_IPoolService *IPoolServiceFilterer) FilterUncoveredLoss(opts *bind.FilterOpts, creditManager []common.Address) (*IPoolServiceUncoveredLossIterator, error) {

	var creditManagerRule []interface{}
	for _, creditManagerItem := range creditManager {
		creditManagerRule = append(creditManagerRule, creditManagerItem)
	}

	logs, sub, err := _IPoolService.contract.FilterLogs(opts, "UncoveredLoss", creditManagerRule)
	if err != nil {
		return nil, err
	}
	return &IPoolServiceUncoveredLossIterator{contract: _IPoolService.contract, event: "UncoveredLoss", logs: logs, sub: sub}, nil
}

// WatchUncoveredLoss is a free log subscription operation binding the contract event 0xef3653ded679720ab04913b6f3820be7cedc8286d42ff5dd8dff17e91bd2964c.
//
// Solidity: event UncoveredLoss(address indexed creditManager, uint256 loss)
func (_IPoolService *IPoolServiceFilterer) WatchUncoveredLoss(opts *bind.WatchOpts, sink chan<- *IPoolServiceUncoveredLoss, creditManager []common.Address) (event.Subscription, error) {

	var creditManagerRule []interface{}
	for _, creditManagerItem := range creditManager {
		creditManagerRule = append(creditManagerRule, creditManagerItem)
	}

	logs, sub, err := _IPoolService.contract.WatchLogs(opts, "UncoveredLoss", creditManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPoolServiceUncoveredLoss)
				if err := _IPoolService.contract.UnpackLog(event, "UncoveredLoss", log); err != nil {
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

// ParseUncoveredLoss is a log parse operation binding the contract event 0xef3653ded679720ab04913b6f3820be7cedc8286d42ff5dd8dff17e91bd2964c.
//
// Solidity: event UncoveredLoss(address indexed creditManager, uint256 loss)
func (_IPoolService *IPoolServiceFilterer) ParseUncoveredLoss(log types.Log) (*IPoolServiceUncoveredLoss, error) {
	event := new(IPoolServiceUncoveredLoss)
	if err := _IPoolService.contract.UnpackLog(event, "UncoveredLoss", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
