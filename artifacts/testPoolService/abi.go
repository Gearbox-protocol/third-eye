// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testPoolService

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

// TestPoolServiceMetaData contains all meta data concerning the TestPoolService contract.
var TestPoolServiceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addressProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_underlyingToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_dieselAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"interestRateModelAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_expectedLiquidityLimit\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"referralCode\",\"type\":\"uint256\"}],\"name\":\"AddLiquidity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creditManager\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Borrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creditManager\",\"type\":\"address\"}],\"name\":\"BorrowForbidden\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creditManager\",\"type\":\"address\"}],\"name\":\"NewCreditManagerConnected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newLimit\",\"type\":\"uint256\"}],\"name\":\"NewExpectedLiquidityLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newInterestRateModel\",\"type\":\"address\"}],\"name\":\"NewInterestRateModel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"NewWithdrawFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RemoveLiquidity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creditManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"borrowedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"profit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"loss\",\"type\":\"uint256\"}],\"name\":\"Repay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creditManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"loss\",\"type\":\"uint256\"}],\"name\":\"UncoveredLoss\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_cumulativeIndex_RAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_expectedLiquidityLU\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_timestampLU\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"referralCode\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addressProvider\",\"outputs\":[{\"internalType\":\"contractAddressProvider\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"availableLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"borrowAPY_RAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cumulativeIndexAtOpen\",\"type\":\"uint256\"}],\"name\":\"calcCumulativeIndexAtBorrowMore\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calcLinearCumulative_RAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cumulativeIndex_RAY\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentBorrowRate_RAY\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeDifference\",\"type\":\"uint256\"}],\"name\":\"calcLinearIndex_RAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_creditManager\",\"type\":\"address\"}],\"name\":\"connectCreditManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"creditManagers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"creditManagersCanBorrow\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"creditManagersCanRepay\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creditManagersCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dieselToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expectedLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expectedLiquidityLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_creditManager\",\"type\":\"address\"}],\"name\":\"forbidCreditManagerToBorrow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"fromDiesel\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCumulativeIndex_RAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDieselRate_RAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExpectedLU\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTimestampLU\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"interestRateModel\",\"outputs\":[{\"internalType\":\"contractIInterestRateModel\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"borrowedAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"}],\"name\":\"lendCreditAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"removeLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"borrowedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"profit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"loss\",\"type\":\"uint256\"}],\"name\":\"repayCreditAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newExpectedLiquidity\",\"type\":\"uint256\"}],\"name\":\"setExpectedLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newLimit\",\"type\":\"uint256\"}],\"name\":\"setExpectedLiquidityLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"setWithdrawFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"toDiesel\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBorrowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasuryAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlyingToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateBorrowRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_interestRateModel\",\"type\":\"address\"}],\"name\":\"updateInterestRateModel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// TestPoolServiceABI is the input ABI used to generate the binding from.
// Deprecated: Use TestPoolServiceMetaData.ABI instead.
var TestPoolServiceABI = TestPoolServiceMetaData.ABI

// TestPoolService is an auto generated Go binding around an Ethereum contract.
type TestPoolService struct {
	TestPoolServiceCaller     // Read-only binding to the contract
	TestPoolServiceTransactor // Write-only binding to the contract
	TestPoolServiceFilterer   // Log filterer for contract events
}

// TestPoolServiceCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestPoolServiceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestPoolServiceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestPoolServiceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestPoolServiceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestPoolServiceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestPoolServiceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestPoolServiceSession struct {
	Contract     *TestPoolService  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestPoolServiceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestPoolServiceCallerSession struct {
	Contract *TestPoolServiceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// TestPoolServiceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestPoolServiceTransactorSession struct {
	Contract     *TestPoolServiceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// TestPoolServiceRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestPoolServiceRaw struct {
	Contract *TestPoolService // Generic contract binding to access the raw methods on
}

// TestPoolServiceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestPoolServiceCallerRaw struct {
	Contract *TestPoolServiceCaller // Generic read-only contract binding to access the raw methods on
}

// TestPoolServiceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestPoolServiceTransactorRaw struct {
	Contract *TestPoolServiceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestPoolService creates a new instance of TestPoolService, bound to a specific deployed contract.
func NewTestPoolService(address common.Address, backend bind.ContractBackend) (*TestPoolService, error) {
	contract, err := bindTestPoolService(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestPoolService{TestPoolServiceCaller: TestPoolServiceCaller{contract: contract}, TestPoolServiceTransactor: TestPoolServiceTransactor{contract: contract}, TestPoolServiceFilterer: TestPoolServiceFilterer{contract: contract}}, nil
}

// NewTestPoolServiceCaller creates a new read-only instance of TestPoolService, bound to a specific deployed contract.
func NewTestPoolServiceCaller(address common.Address, caller bind.ContractCaller) (*TestPoolServiceCaller, error) {
	contract, err := bindTestPoolService(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestPoolServiceCaller{contract: contract}, nil
}

// NewTestPoolServiceTransactor creates a new write-only instance of TestPoolService, bound to a specific deployed contract.
func NewTestPoolServiceTransactor(address common.Address, transactor bind.ContractTransactor) (*TestPoolServiceTransactor, error) {
	contract, err := bindTestPoolService(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestPoolServiceTransactor{contract: contract}, nil
}

// NewTestPoolServiceFilterer creates a new log filterer instance of TestPoolService, bound to a specific deployed contract.
func NewTestPoolServiceFilterer(address common.Address, filterer bind.ContractFilterer) (*TestPoolServiceFilterer, error) {
	contract, err := bindTestPoolService(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestPoolServiceFilterer{contract: contract}, nil
}

// bindTestPoolService binds a generic wrapper to an already deployed contract.
func bindTestPoolService(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestPoolServiceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestPoolService *TestPoolServiceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestPoolService.Contract.TestPoolServiceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestPoolService *TestPoolServiceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestPoolService.Contract.TestPoolServiceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestPoolService *TestPoolServiceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestPoolService.Contract.TestPoolServiceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestPoolService *TestPoolServiceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestPoolService.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestPoolService *TestPoolServiceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestPoolService.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestPoolService *TestPoolServiceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestPoolService.Contract.contract.Transact(opts, method, params...)
}

// CumulativeIndexRAY is a free data retrieval call binding the contract method 0xdbcb313b.
//
// Solidity: function _cumulativeIndex_RAY() view returns(uint256)
func (_TestPoolService *TestPoolServiceCaller) CumulativeIndexRAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestPoolService.contract.Call(opts, &out, "_cumulativeIndex_RAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CumulativeIndexRAY is a free data retrieval call binding the contract method 0xdbcb313b.
//
// Solidity: function _cumulativeIndex_RAY() view returns(uint256)
func (_TestPoolService *TestPoolServiceSession) CumulativeIndexRAY() (*big.Int, error) {
	return _TestPoolService.Contract.CumulativeIndexRAY(&_TestPoolService.CallOpts)
}

// CumulativeIndexRAY is a free data retrieval call binding the contract method 0xdbcb313b.
//
// Solidity: function _cumulativeIndex_RAY() view returns(uint256)
func (_TestPoolService *TestPoolServiceCallerSession) CumulativeIndexRAY() (*big.Int, error) {
	return _TestPoolService.Contract.CumulativeIndexRAY(&_TestPoolService.CallOpts)
}

// ExpectedLiquidityLU is a free data retrieval call binding the contract method 0x030dbb04.
//
// Solidity: function _expectedLiquidityLU() view returns(uint256)
func (_TestPoolService *TestPoolServiceCaller) ExpectedLiquidityLU(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestPoolService.contract.Call(opts, &out, "_expectedLiquidityLU")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExpectedLiquidityLU is a free data retrieval call binding the contract method 0x030dbb04.
//
// Solidity: function _expectedLiquidityLU() view returns(uint256)
func (_TestPoolService *TestPoolServiceSession) ExpectedLiquidityLU() (*big.Int, error) {
	return _TestPoolService.Contract.ExpectedLiquidityLU(&_TestPoolService.CallOpts)
}

// ExpectedLiquidityLU is a free data retrieval call binding the contract method 0x030dbb04.
//
// Solidity: function _expectedLiquidityLU() view returns(uint256)
func (_TestPoolService *TestPoolServiceCallerSession) ExpectedLiquidityLU() (*big.Int, error) {
	return _TestPoolService.Contract.ExpectedLiquidityLU(&_TestPoolService.CallOpts)
}

// TimestampLU is a free data retrieval call binding the contract method 0x609ae317.
//
// Solidity: function _timestampLU() view returns(uint256)
func (_TestPoolService *TestPoolServiceCaller) TimestampLU(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestPoolService.contract.Call(opts, &out, "_timestampLU")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TimestampLU is a free data retrieval call binding the contract method 0x609ae317.
//
// Solidity: function _timestampLU() view returns(uint256)
func (_TestPoolService *TestPoolServiceSession) TimestampLU() (*big.Int, error) {
	return _TestPoolService.Contract.TimestampLU(&_TestPoolService.CallOpts)
}

// TimestampLU is a free data retrieval call binding the contract method 0x609ae317.
//
// Solidity: function _timestampLU() view returns(uint256)
func (_TestPoolService *TestPoolServiceCallerSession) TimestampLU() (*big.Int, error) {
	return _TestPoolService.Contract.TimestampLU(&_TestPoolService.CallOpts)
}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_TestPoolService *TestPoolServiceCaller) AddressProvider(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TestPoolService.contract.Call(opts, &out, "addressProvider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_TestPoolService *TestPoolServiceSession) AddressProvider() (common.Address, error) {
	return _TestPoolService.Contract.AddressProvider(&_TestPoolService.CallOpts)
}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_TestPoolService *TestPoolServiceCallerSession) AddressProvider() (common.Address, error) {
	return _TestPoolService.Contract.AddressProvider(&_TestPoolService.CallOpts)
}

// AvailableLiquidity is a free data retrieval call binding the contract method 0x74375359.
//
// Solidity: function availableLiquidity() view returns(uint256)
func (_TestPoolService *TestPoolServiceCaller) AvailableLiquidity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestPoolService.contract.Call(opts, &out, "availableLiquidity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AvailableLiquidity is a free data retrieval call binding the contract method 0x74375359.
//
// Solidity: function availableLiquidity() view returns(uint256)
func (_TestPoolService *TestPoolServiceSession) AvailableLiquidity() (*big.Int, error) {
	return _TestPoolService.Contract.AvailableLiquidity(&_TestPoolService.CallOpts)
}

// AvailableLiquidity is a free data retrieval call binding the contract method 0x74375359.
//
// Solidity: function availableLiquidity() view returns(uint256)
func (_TestPoolService *TestPoolServiceCallerSession) AvailableLiquidity() (*big.Int, error) {
	return _TestPoolService.Contract.AvailableLiquidity(&_TestPoolService.CallOpts)
}

// BorrowAPYRAY is a free data retrieval call binding the contract method 0x45d31f9d.
//
// Solidity: function borrowAPY_RAY() view returns(uint256)
func (_TestPoolService *TestPoolServiceCaller) BorrowAPYRAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestPoolService.contract.Call(opts, &out, "borrowAPY_RAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowAPYRAY is a free data retrieval call binding the contract method 0x45d31f9d.
//
// Solidity: function borrowAPY_RAY() view returns(uint256)
func (_TestPoolService *TestPoolServiceSession) BorrowAPYRAY() (*big.Int, error) {
	return _TestPoolService.Contract.BorrowAPYRAY(&_TestPoolService.CallOpts)
}

// BorrowAPYRAY is a free data retrieval call binding the contract method 0x45d31f9d.
//
// Solidity: function borrowAPY_RAY() view returns(uint256)
func (_TestPoolService *TestPoolServiceCallerSession) BorrowAPYRAY() (*big.Int, error) {
	return _TestPoolService.Contract.BorrowAPYRAY(&_TestPoolService.CallOpts)
}

// CalcCumulativeIndexAtBorrowMore is a free data retrieval call binding the contract method 0xc00495a1.
//
// Solidity: function calcCumulativeIndexAtBorrowMore(uint256 amount, uint256 dAmount, uint256 cumulativeIndexAtOpen) view returns(uint256)
func (_TestPoolService *TestPoolServiceCaller) CalcCumulativeIndexAtBorrowMore(opts *bind.CallOpts, amount *big.Int, dAmount *big.Int, cumulativeIndexAtOpen *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TestPoolService.contract.Call(opts, &out, "calcCumulativeIndexAtBorrowMore", amount, dAmount, cumulativeIndexAtOpen)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcCumulativeIndexAtBorrowMore is a free data retrieval call binding the contract method 0xc00495a1.
//
// Solidity: function calcCumulativeIndexAtBorrowMore(uint256 amount, uint256 dAmount, uint256 cumulativeIndexAtOpen) view returns(uint256)
func (_TestPoolService *TestPoolServiceSession) CalcCumulativeIndexAtBorrowMore(amount *big.Int, dAmount *big.Int, cumulativeIndexAtOpen *big.Int) (*big.Int, error) {
	return _TestPoolService.Contract.CalcCumulativeIndexAtBorrowMore(&_TestPoolService.CallOpts, amount, dAmount, cumulativeIndexAtOpen)
}

// CalcCumulativeIndexAtBorrowMore is a free data retrieval call binding the contract method 0xc00495a1.
//
// Solidity: function calcCumulativeIndexAtBorrowMore(uint256 amount, uint256 dAmount, uint256 cumulativeIndexAtOpen) view returns(uint256)
func (_TestPoolService *TestPoolServiceCallerSession) CalcCumulativeIndexAtBorrowMore(amount *big.Int, dAmount *big.Int, cumulativeIndexAtOpen *big.Int) (*big.Int, error) {
	return _TestPoolService.Contract.CalcCumulativeIndexAtBorrowMore(&_TestPoolService.CallOpts, amount, dAmount, cumulativeIndexAtOpen)
}

// CalcLinearCumulativeRAY is a free data retrieval call binding the contract method 0x0fce70fb.
//
// Solidity: function calcLinearCumulative_RAY() view returns(uint256)
func (_TestPoolService *TestPoolServiceCaller) CalcLinearCumulativeRAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestPoolService.contract.Call(opts, &out, "calcLinearCumulative_RAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcLinearCumulativeRAY is a free data retrieval call binding the contract method 0x0fce70fb.
//
// Solidity: function calcLinearCumulative_RAY() view returns(uint256)
func (_TestPoolService *TestPoolServiceSession) CalcLinearCumulativeRAY() (*big.Int, error) {
	return _TestPoolService.Contract.CalcLinearCumulativeRAY(&_TestPoolService.CallOpts)
}

// CalcLinearCumulativeRAY is a free data retrieval call binding the contract method 0x0fce70fb.
//
// Solidity: function calcLinearCumulative_RAY() view returns(uint256)
func (_TestPoolService *TestPoolServiceCallerSession) CalcLinearCumulativeRAY() (*big.Int, error) {
	return _TestPoolService.Contract.CalcLinearCumulativeRAY(&_TestPoolService.CallOpts)
}

// CalcLinearIndexRAY is a free data retrieval call binding the contract method 0x31d8bc27.
//
// Solidity: function calcLinearIndex_RAY(uint256 cumulativeIndex_RAY, uint256 currentBorrowRate_RAY, uint256 timeDifference) pure returns(uint256)
func (_TestPoolService *TestPoolServiceCaller) CalcLinearIndexRAY(opts *bind.CallOpts, cumulativeIndex_RAY *big.Int, currentBorrowRate_RAY *big.Int, timeDifference *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TestPoolService.contract.Call(opts, &out, "calcLinearIndex_RAY", cumulativeIndex_RAY, currentBorrowRate_RAY, timeDifference)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcLinearIndexRAY is a free data retrieval call binding the contract method 0x31d8bc27.
//
// Solidity: function calcLinearIndex_RAY(uint256 cumulativeIndex_RAY, uint256 currentBorrowRate_RAY, uint256 timeDifference) pure returns(uint256)
func (_TestPoolService *TestPoolServiceSession) CalcLinearIndexRAY(cumulativeIndex_RAY *big.Int, currentBorrowRate_RAY *big.Int, timeDifference *big.Int) (*big.Int, error) {
	return _TestPoolService.Contract.CalcLinearIndexRAY(&_TestPoolService.CallOpts, cumulativeIndex_RAY, currentBorrowRate_RAY, timeDifference)
}

// CalcLinearIndexRAY is a free data retrieval call binding the contract method 0x31d8bc27.
//
// Solidity: function calcLinearIndex_RAY(uint256 cumulativeIndex_RAY, uint256 currentBorrowRate_RAY, uint256 timeDifference) pure returns(uint256)
func (_TestPoolService *TestPoolServiceCallerSession) CalcLinearIndexRAY(cumulativeIndex_RAY *big.Int, currentBorrowRate_RAY *big.Int, timeDifference *big.Int) (*big.Int, error) {
	return _TestPoolService.Contract.CalcLinearIndexRAY(&_TestPoolService.CallOpts, cumulativeIndex_RAY, currentBorrowRate_RAY, timeDifference)
}

// CreditManagers is a free data retrieval call binding the contract method 0x1e16e4fc.
//
// Solidity: function creditManagers(uint256 ) view returns(address)
func (_TestPoolService *TestPoolServiceCaller) CreditManagers(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TestPoolService.contract.Call(opts, &out, "creditManagers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditManagers is a free data retrieval call binding the contract method 0x1e16e4fc.
//
// Solidity: function creditManagers(uint256 ) view returns(address)
func (_TestPoolService *TestPoolServiceSession) CreditManagers(arg0 *big.Int) (common.Address, error) {
	return _TestPoolService.Contract.CreditManagers(&_TestPoolService.CallOpts, arg0)
}

// CreditManagers is a free data retrieval call binding the contract method 0x1e16e4fc.
//
// Solidity: function creditManagers(uint256 ) view returns(address)
func (_TestPoolService *TestPoolServiceCallerSession) CreditManagers(arg0 *big.Int) (common.Address, error) {
	return _TestPoolService.Contract.CreditManagers(&_TestPoolService.CallOpts, arg0)
}

// CreditManagersCanBorrow is a free data retrieval call binding the contract method 0x2e97ca21.
//
// Solidity: function creditManagersCanBorrow(address ) view returns(bool)
func (_TestPoolService *TestPoolServiceCaller) CreditManagersCanBorrow(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _TestPoolService.contract.Call(opts, &out, "creditManagersCanBorrow", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CreditManagersCanBorrow is a free data retrieval call binding the contract method 0x2e97ca21.
//
// Solidity: function creditManagersCanBorrow(address ) view returns(bool)
func (_TestPoolService *TestPoolServiceSession) CreditManagersCanBorrow(arg0 common.Address) (bool, error) {
	return _TestPoolService.Contract.CreditManagersCanBorrow(&_TestPoolService.CallOpts, arg0)
}

// CreditManagersCanBorrow is a free data retrieval call binding the contract method 0x2e97ca21.
//
// Solidity: function creditManagersCanBorrow(address ) view returns(bool)
func (_TestPoolService *TestPoolServiceCallerSession) CreditManagersCanBorrow(arg0 common.Address) (bool, error) {
	return _TestPoolService.Contract.CreditManagersCanBorrow(&_TestPoolService.CallOpts, arg0)
}

// CreditManagersCanRepay is a free data retrieval call binding the contract method 0x3e163df0.
//
// Solidity: function creditManagersCanRepay(address ) view returns(bool)
func (_TestPoolService *TestPoolServiceCaller) CreditManagersCanRepay(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _TestPoolService.contract.Call(opts, &out, "creditManagersCanRepay", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CreditManagersCanRepay is a free data retrieval call binding the contract method 0x3e163df0.
//
// Solidity: function creditManagersCanRepay(address ) view returns(bool)
func (_TestPoolService *TestPoolServiceSession) CreditManagersCanRepay(arg0 common.Address) (bool, error) {
	return _TestPoolService.Contract.CreditManagersCanRepay(&_TestPoolService.CallOpts, arg0)
}

// CreditManagersCanRepay is a free data retrieval call binding the contract method 0x3e163df0.
//
// Solidity: function creditManagersCanRepay(address ) view returns(bool)
func (_TestPoolService *TestPoolServiceCallerSession) CreditManagersCanRepay(arg0 common.Address) (bool, error) {
	return _TestPoolService.Contract.CreditManagersCanRepay(&_TestPoolService.CallOpts, arg0)
}

// CreditManagersCount is a free data retrieval call binding the contract method 0xa4e8273e.
//
// Solidity: function creditManagersCount() view returns(uint256)
func (_TestPoolService *TestPoolServiceCaller) CreditManagersCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestPoolService.contract.Call(opts, &out, "creditManagersCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CreditManagersCount is a free data retrieval call binding the contract method 0xa4e8273e.
//
// Solidity: function creditManagersCount() view returns(uint256)
func (_TestPoolService *TestPoolServiceSession) CreditManagersCount() (*big.Int, error) {
	return _TestPoolService.Contract.CreditManagersCount(&_TestPoolService.CallOpts)
}

// CreditManagersCount is a free data retrieval call binding the contract method 0xa4e8273e.
//
// Solidity: function creditManagersCount() view returns(uint256)
func (_TestPoolService *TestPoolServiceCallerSession) CreditManagersCount() (*big.Int, error) {
	return _TestPoolService.Contract.CreditManagersCount(&_TestPoolService.CallOpts)
}

// DieselToken is a free data retrieval call binding the contract method 0x36dda7d5.
//
// Solidity: function dieselToken() view returns(address)
func (_TestPoolService *TestPoolServiceCaller) DieselToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TestPoolService.contract.Call(opts, &out, "dieselToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DieselToken is a free data retrieval call binding the contract method 0x36dda7d5.
//
// Solidity: function dieselToken() view returns(address)
func (_TestPoolService *TestPoolServiceSession) DieselToken() (common.Address, error) {
	return _TestPoolService.Contract.DieselToken(&_TestPoolService.CallOpts)
}

// DieselToken is a free data retrieval call binding the contract method 0x36dda7d5.
//
// Solidity: function dieselToken() view returns(address)
func (_TestPoolService *TestPoolServiceCallerSession) DieselToken() (common.Address, error) {
	return _TestPoolService.Contract.DieselToken(&_TestPoolService.CallOpts)
}

// ExpectedLiquidity is a free data retrieval call binding the contract method 0xfe14112d.
//
// Solidity: function expectedLiquidity() view returns(uint256)
func (_TestPoolService *TestPoolServiceCaller) ExpectedLiquidity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestPoolService.contract.Call(opts, &out, "expectedLiquidity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExpectedLiquidity is a free data retrieval call binding the contract method 0xfe14112d.
//
// Solidity: function expectedLiquidity() view returns(uint256)
func (_TestPoolService *TestPoolServiceSession) ExpectedLiquidity() (*big.Int, error) {
	return _TestPoolService.Contract.ExpectedLiquidity(&_TestPoolService.CallOpts)
}

// ExpectedLiquidity is a free data retrieval call binding the contract method 0xfe14112d.
//
// Solidity: function expectedLiquidity() view returns(uint256)
func (_TestPoolService *TestPoolServiceCallerSession) ExpectedLiquidity() (*big.Int, error) {
	return _TestPoolService.Contract.ExpectedLiquidity(&_TestPoolService.CallOpts)
}

// ExpectedLiquidityLimit is a free data retrieval call binding the contract method 0xef8d9603.
//
// Solidity: function expectedLiquidityLimit() view returns(uint256)
func (_TestPoolService *TestPoolServiceCaller) ExpectedLiquidityLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestPoolService.contract.Call(opts, &out, "expectedLiquidityLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExpectedLiquidityLimit is a free data retrieval call binding the contract method 0xef8d9603.
//
// Solidity: function expectedLiquidityLimit() view returns(uint256)
func (_TestPoolService *TestPoolServiceSession) ExpectedLiquidityLimit() (*big.Int, error) {
	return _TestPoolService.Contract.ExpectedLiquidityLimit(&_TestPoolService.CallOpts)
}

// ExpectedLiquidityLimit is a free data retrieval call binding the contract method 0xef8d9603.
//
// Solidity: function expectedLiquidityLimit() view returns(uint256)
func (_TestPoolService *TestPoolServiceCallerSession) ExpectedLiquidityLimit() (*big.Int, error) {
	return _TestPoolService.Contract.ExpectedLiquidityLimit(&_TestPoolService.CallOpts)
}

// FromDiesel is a free data retrieval call binding the contract method 0x5427c938.
//
// Solidity: function fromDiesel(uint256 amount) view returns(uint256)
func (_TestPoolService *TestPoolServiceCaller) FromDiesel(opts *bind.CallOpts, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TestPoolService.contract.Call(opts, &out, "fromDiesel", amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FromDiesel is a free data retrieval call binding the contract method 0x5427c938.
//
// Solidity: function fromDiesel(uint256 amount) view returns(uint256)
func (_TestPoolService *TestPoolServiceSession) FromDiesel(amount *big.Int) (*big.Int, error) {
	return _TestPoolService.Contract.FromDiesel(&_TestPoolService.CallOpts, amount)
}

// FromDiesel is a free data retrieval call binding the contract method 0x5427c938.
//
// Solidity: function fromDiesel(uint256 amount) view returns(uint256)
func (_TestPoolService *TestPoolServiceCallerSession) FromDiesel(amount *big.Int) (*big.Int, error) {
	return _TestPoolService.Contract.FromDiesel(&_TestPoolService.CallOpts, amount)
}

// GetCumulativeIndexRAY is a free data retrieval call binding the contract method 0xa7a5eccc.
//
// Solidity: function getCumulativeIndex_RAY() view returns(uint256)
func (_TestPoolService *TestPoolServiceCaller) GetCumulativeIndexRAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestPoolService.contract.Call(opts, &out, "getCumulativeIndex_RAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCumulativeIndexRAY is a free data retrieval call binding the contract method 0xa7a5eccc.
//
// Solidity: function getCumulativeIndex_RAY() view returns(uint256)
func (_TestPoolService *TestPoolServiceSession) GetCumulativeIndexRAY() (*big.Int, error) {
	return _TestPoolService.Contract.GetCumulativeIndexRAY(&_TestPoolService.CallOpts)
}

// GetCumulativeIndexRAY is a free data retrieval call binding the contract method 0xa7a5eccc.
//
// Solidity: function getCumulativeIndex_RAY() view returns(uint256)
func (_TestPoolService *TestPoolServiceCallerSession) GetCumulativeIndexRAY() (*big.Int, error) {
	return _TestPoolService.Contract.GetCumulativeIndexRAY(&_TestPoolService.CallOpts)
}

// GetDieselRateRAY is a free data retrieval call binding the contract method 0x788c6bfe.
//
// Solidity: function getDieselRate_RAY() view returns(uint256)
func (_TestPoolService *TestPoolServiceCaller) GetDieselRateRAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestPoolService.contract.Call(opts, &out, "getDieselRate_RAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDieselRateRAY is a free data retrieval call binding the contract method 0x788c6bfe.
//
// Solidity: function getDieselRate_RAY() view returns(uint256)
func (_TestPoolService *TestPoolServiceSession) GetDieselRateRAY() (*big.Int, error) {
	return _TestPoolService.Contract.GetDieselRateRAY(&_TestPoolService.CallOpts)
}

// GetDieselRateRAY is a free data retrieval call binding the contract method 0x788c6bfe.
//
// Solidity: function getDieselRate_RAY() view returns(uint256)
func (_TestPoolService *TestPoolServiceCallerSession) GetDieselRateRAY() (*big.Int, error) {
	return _TestPoolService.Contract.GetDieselRateRAY(&_TestPoolService.CallOpts)
}

// GetExpectedLU is a free data retrieval call binding the contract method 0x52cb995f.
//
// Solidity: function getExpectedLU() view returns(uint256)
func (_TestPoolService *TestPoolServiceCaller) GetExpectedLU(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestPoolService.contract.Call(opts, &out, "getExpectedLU")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetExpectedLU is a free data retrieval call binding the contract method 0x52cb995f.
//
// Solidity: function getExpectedLU() view returns(uint256)
func (_TestPoolService *TestPoolServiceSession) GetExpectedLU() (*big.Int, error) {
	return _TestPoolService.Contract.GetExpectedLU(&_TestPoolService.CallOpts)
}

// GetExpectedLU is a free data retrieval call binding the contract method 0x52cb995f.
//
// Solidity: function getExpectedLU() view returns(uint256)
func (_TestPoolService *TestPoolServiceCallerSession) GetExpectedLU() (*big.Int, error) {
	return _TestPoolService.Contract.GetExpectedLU(&_TestPoolService.CallOpts)
}

// GetTimestampLU is a free data retrieval call binding the contract method 0x5b29233b.
//
// Solidity: function getTimestampLU() view returns(uint256)
func (_TestPoolService *TestPoolServiceCaller) GetTimestampLU(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestPoolService.contract.Call(opts, &out, "getTimestampLU")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTimestampLU is a free data retrieval call binding the contract method 0x5b29233b.
//
// Solidity: function getTimestampLU() view returns(uint256)
func (_TestPoolService *TestPoolServiceSession) GetTimestampLU() (*big.Int, error) {
	return _TestPoolService.Contract.GetTimestampLU(&_TestPoolService.CallOpts)
}

// GetTimestampLU is a free data retrieval call binding the contract method 0x5b29233b.
//
// Solidity: function getTimestampLU() view returns(uint256)
func (_TestPoolService *TestPoolServiceCallerSession) GetTimestampLU() (*big.Int, error) {
	return _TestPoolService.Contract.GetTimestampLU(&_TestPoolService.CallOpts)
}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_TestPoolService *TestPoolServiceCaller) InterestRateModel(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TestPoolService.contract.Call(opts, &out, "interestRateModel")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_TestPoolService *TestPoolServiceSession) InterestRateModel() (common.Address, error) {
	return _TestPoolService.Contract.InterestRateModel(&_TestPoolService.CallOpts)
}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_TestPoolService *TestPoolServiceCallerSession) InterestRateModel() (common.Address, error) {
	return _TestPoolService.Contract.InterestRateModel(&_TestPoolService.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_TestPoolService *TestPoolServiceCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TestPoolService.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_TestPoolService *TestPoolServiceSession) Paused() (bool, error) {
	return _TestPoolService.Contract.Paused(&_TestPoolService.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_TestPoolService *TestPoolServiceCallerSession) Paused() (bool, error) {
	return _TestPoolService.Contract.Paused(&_TestPoolService.CallOpts)
}

// ToDiesel is a free data retrieval call binding the contract method 0x4d778ad1.
//
// Solidity: function toDiesel(uint256 amount) view returns(uint256)
func (_TestPoolService *TestPoolServiceCaller) ToDiesel(opts *bind.CallOpts, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TestPoolService.contract.Call(opts, &out, "toDiesel", amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ToDiesel is a free data retrieval call binding the contract method 0x4d778ad1.
//
// Solidity: function toDiesel(uint256 amount) view returns(uint256)
func (_TestPoolService *TestPoolServiceSession) ToDiesel(amount *big.Int) (*big.Int, error) {
	return _TestPoolService.Contract.ToDiesel(&_TestPoolService.CallOpts, amount)
}

// ToDiesel is a free data retrieval call binding the contract method 0x4d778ad1.
//
// Solidity: function toDiesel(uint256 amount) view returns(uint256)
func (_TestPoolService *TestPoolServiceCallerSession) ToDiesel(amount *big.Int) (*big.Int, error) {
	return _TestPoolService.Contract.ToDiesel(&_TestPoolService.CallOpts, amount)
}

// TotalBorrowed is a free data retrieval call binding the contract method 0x4c19386c.
//
// Solidity: function totalBorrowed() view returns(uint256)
func (_TestPoolService *TestPoolServiceCaller) TotalBorrowed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestPoolService.contract.Call(opts, &out, "totalBorrowed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBorrowed is a free data retrieval call binding the contract method 0x4c19386c.
//
// Solidity: function totalBorrowed() view returns(uint256)
func (_TestPoolService *TestPoolServiceSession) TotalBorrowed() (*big.Int, error) {
	return _TestPoolService.Contract.TotalBorrowed(&_TestPoolService.CallOpts)
}

// TotalBorrowed is a free data retrieval call binding the contract method 0x4c19386c.
//
// Solidity: function totalBorrowed() view returns(uint256)
func (_TestPoolService *TestPoolServiceCallerSession) TotalBorrowed() (*big.Int, error) {
	return _TestPoolService.Contract.TotalBorrowed(&_TestPoolService.CallOpts)
}

// TreasuryAddress is a free data retrieval call binding the contract method 0xc5f956af.
//
// Solidity: function treasuryAddress() view returns(address)
func (_TestPoolService *TestPoolServiceCaller) TreasuryAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TestPoolService.contract.Call(opts, &out, "treasuryAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TreasuryAddress is a free data retrieval call binding the contract method 0xc5f956af.
//
// Solidity: function treasuryAddress() view returns(address)
func (_TestPoolService *TestPoolServiceSession) TreasuryAddress() (common.Address, error) {
	return _TestPoolService.Contract.TreasuryAddress(&_TestPoolService.CallOpts)
}

// TreasuryAddress is a free data retrieval call binding the contract method 0xc5f956af.
//
// Solidity: function treasuryAddress() view returns(address)
func (_TestPoolService *TestPoolServiceCallerSession) TreasuryAddress() (common.Address, error) {
	return _TestPoolService.Contract.TreasuryAddress(&_TestPoolService.CallOpts)
}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_TestPoolService *TestPoolServiceCaller) UnderlyingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TestPoolService.contract.Call(opts, &out, "underlyingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_TestPoolService *TestPoolServiceSession) UnderlyingToken() (common.Address, error) {
	return _TestPoolService.Contract.UnderlyingToken(&_TestPoolService.CallOpts)
}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_TestPoolService *TestPoolServiceCallerSession) UnderlyingToken() (common.Address, error) {
	return _TestPoolService.Contract.UnderlyingToken(&_TestPoolService.CallOpts)
}

// WithdrawFee is a free data retrieval call binding the contract method 0xe941fa78.
//
// Solidity: function withdrawFee() view returns(uint256)
func (_TestPoolService *TestPoolServiceCaller) WithdrawFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestPoolService.contract.Call(opts, &out, "withdrawFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawFee is a free data retrieval call binding the contract method 0xe941fa78.
//
// Solidity: function withdrawFee() view returns(uint256)
func (_TestPoolService *TestPoolServiceSession) WithdrawFee() (*big.Int, error) {
	return _TestPoolService.Contract.WithdrawFee(&_TestPoolService.CallOpts)
}

// WithdrawFee is a free data retrieval call binding the contract method 0xe941fa78.
//
// Solidity: function withdrawFee() view returns(uint256)
func (_TestPoolService *TestPoolServiceCallerSession) WithdrawFee() (*big.Int, error) {
	return _TestPoolService.Contract.WithdrawFee(&_TestPoolService.CallOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x9aa5d462.
//
// Solidity: function addLiquidity(uint256 amount, address onBehalfOf, uint256 referralCode) returns()
func (_TestPoolService *TestPoolServiceTransactor) AddLiquidity(opts *bind.TransactOpts, amount *big.Int, onBehalfOf common.Address, referralCode *big.Int) (*types.Transaction, error) {
	return _TestPoolService.contract.Transact(opts, "addLiquidity", amount, onBehalfOf, referralCode)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x9aa5d462.
//
// Solidity: function addLiquidity(uint256 amount, address onBehalfOf, uint256 referralCode) returns()
func (_TestPoolService *TestPoolServiceSession) AddLiquidity(amount *big.Int, onBehalfOf common.Address, referralCode *big.Int) (*types.Transaction, error) {
	return _TestPoolService.Contract.AddLiquidity(&_TestPoolService.TransactOpts, amount, onBehalfOf, referralCode)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x9aa5d462.
//
// Solidity: function addLiquidity(uint256 amount, address onBehalfOf, uint256 referralCode) returns()
func (_TestPoolService *TestPoolServiceTransactorSession) AddLiquidity(amount *big.Int, onBehalfOf common.Address, referralCode *big.Int) (*types.Transaction, error) {
	return _TestPoolService.Contract.AddLiquidity(&_TestPoolService.TransactOpts, amount, onBehalfOf, referralCode)
}

// ConnectCreditManager is a paid mutator transaction binding the contract method 0xcf33d955.
//
// Solidity: function connectCreditManager(address _creditManager) returns()
func (_TestPoolService *TestPoolServiceTransactor) ConnectCreditManager(opts *bind.TransactOpts, _creditManager common.Address) (*types.Transaction, error) {
	return _TestPoolService.contract.Transact(opts, "connectCreditManager", _creditManager)
}

// ConnectCreditManager is a paid mutator transaction binding the contract method 0xcf33d955.
//
// Solidity: function connectCreditManager(address _creditManager) returns()
func (_TestPoolService *TestPoolServiceSession) ConnectCreditManager(_creditManager common.Address) (*types.Transaction, error) {
	return _TestPoolService.Contract.ConnectCreditManager(&_TestPoolService.TransactOpts, _creditManager)
}

// ConnectCreditManager is a paid mutator transaction binding the contract method 0xcf33d955.
//
// Solidity: function connectCreditManager(address _creditManager) returns()
func (_TestPoolService *TestPoolServiceTransactorSession) ConnectCreditManager(_creditManager common.Address) (*types.Transaction, error) {
	return _TestPoolService.Contract.ConnectCreditManager(&_TestPoolService.TransactOpts, _creditManager)
}

// ForbidCreditManagerToBorrow is a paid mutator transaction binding the contract method 0x078c4781.
//
// Solidity: function forbidCreditManagerToBorrow(address _creditManager) returns()
func (_TestPoolService *TestPoolServiceTransactor) ForbidCreditManagerToBorrow(opts *bind.TransactOpts, _creditManager common.Address) (*types.Transaction, error) {
	return _TestPoolService.contract.Transact(opts, "forbidCreditManagerToBorrow", _creditManager)
}

// ForbidCreditManagerToBorrow is a paid mutator transaction binding the contract method 0x078c4781.
//
// Solidity: function forbidCreditManagerToBorrow(address _creditManager) returns()
func (_TestPoolService *TestPoolServiceSession) ForbidCreditManagerToBorrow(_creditManager common.Address) (*types.Transaction, error) {
	return _TestPoolService.Contract.ForbidCreditManagerToBorrow(&_TestPoolService.TransactOpts, _creditManager)
}

// ForbidCreditManagerToBorrow is a paid mutator transaction binding the contract method 0x078c4781.
//
// Solidity: function forbidCreditManagerToBorrow(address _creditManager) returns()
func (_TestPoolService *TestPoolServiceTransactorSession) ForbidCreditManagerToBorrow(_creditManager common.Address) (*types.Transaction, error) {
	return _TestPoolService.Contract.ForbidCreditManagerToBorrow(&_TestPoolService.TransactOpts, _creditManager)
}

// LendCreditAccount is a paid mutator transaction binding the contract method 0xbf28068b.
//
// Solidity: function lendCreditAccount(uint256 borrowedAmount, address creditAccount) returns()
func (_TestPoolService *TestPoolServiceTransactor) LendCreditAccount(opts *bind.TransactOpts, borrowedAmount *big.Int, creditAccount common.Address) (*types.Transaction, error) {
	return _TestPoolService.contract.Transact(opts, "lendCreditAccount", borrowedAmount, creditAccount)
}

// LendCreditAccount is a paid mutator transaction binding the contract method 0xbf28068b.
//
// Solidity: function lendCreditAccount(uint256 borrowedAmount, address creditAccount) returns()
func (_TestPoolService *TestPoolServiceSession) LendCreditAccount(borrowedAmount *big.Int, creditAccount common.Address) (*types.Transaction, error) {
	return _TestPoolService.Contract.LendCreditAccount(&_TestPoolService.TransactOpts, borrowedAmount, creditAccount)
}

// LendCreditAccount is a paid mutator transaction binding the contract method 0xbf28068b.
//
// Solidity: function lendCreditAccount(uint256 borrowedAmount, address creditAccount) returns()
func (_TestPoolService *TestPoolServiceTransactorSession) LendCreditAccount(borrowedAmount *big.Int, creditAccount common.Address) (*types.Transaction, error) {
	return _TestPoolService.Contract.LendCreditAccount(&_TestPoolService.TransactOpts, borrowedAmount, creditAccount)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_TestPoolService *TestPoolServiceTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestPoolService.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_TestPoolService *TestPoolServiceSession) Pause() (*types.Transaction, error) {
	return _TestPoolService.Contract.Pause(&_TestPoolService.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_TestPoolService *TestPoolServiceTransactorSession) Pause() (*types.Transaction, error) {
	return _TestPoolService.Contract.Pause(&_TestPoolService.TransactOpts)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x05fe138b.
//
// Solidity: function removeLiquidity(uint256 amount, address to) returns(uint256)
func (_TestPoolService *TestPoolServiceTransactor) RemoveLiquidity(opts *bind.TransactOpts, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _TestPoolService.contract.Transact(opts, "removeLiquidity", amount, to)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x05fe138b.
//
// Solidity: function removeLiquidity(uint256 amount, address to) returns(uint256)
func (_TestPoolService *TestPoolServiceSession) RemoveLiquidity(amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _TestPoolService.Contract.RemoveLiquidity(&_TestPoolService.TransactOpts, amount, to)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x05fe138b.
//
// Solidity: function removeLiquidity(uint256 amount, address to) returns(uint256)
func (_TestPoolService *TestPoolServiceTransactorSession) RemoveLiquidity(amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _TestPoolService.Contract.RemoveLiquidity(&_TestPoolService.TransactOpts, amount, to)
}

// RepayCreditAccount is a paid mutator transaction binding the contract method 0xca9505e4.
//
// Solidity: function repayCreditAccount(uint256 borrowedAmount, uint256 profit, uint256 loss) returns()
func (_TestPoolService *TestPoolServiceTransactor) RepayCreditAccount(opts *bind.TransactOpts, borrowedAmount *big.Int, profit *big.Int, loss *big.Int) (*types.Transaction, error) {
	return _TestPoolService.contract.Transact(opts, "repayCreditAccount", borrowedAmount, profit, loss)
}

// RepayCreditAccount is a paid mutator transaction binding the contract method 0xca9505e4.
//
// Solidity: function repayCreditAccount(uint256 borrowedAmount, uint256 profit, uint256 loss) returns()
func (_TestPoolService *TestPoolServiceSession) RepayCreditAccount(borrowedAmount *big.Int, profit *big.Int, loss *big.Int) (*types.Transaction, error) {
	return _TestPoolService.Contract.RepayCreditAccount(&_TestPoolService.TransactOpts, borrowedAmount, profit, loss)
}

// RepayCreditAccount is a paid mutator transaction binding the contract method 0xca9505e4.
//
// Solidity: function repayCreditAccount(uint256 borrowedAmount, uint256 profit, uint256 loss) returns()
func (_TestPoolService *TestPoolServiceTransactorSession) RepayCreditAccount(borrowedAmount *big.Int, profit *big.Int, loss *big.Int) (*types.Transaction, error) {
	return _TestPoolService.Contract.RepayCreditAccount(&_TestPoolService.TransactOpts, borrowedAmount, profit, loss)
}

// SetExpectedLiquidity is a paid mutator transaction binding the contract method 0x7dfc6afe.
//
// Solidity: function setExpectedLiquidity(uint256 newExpectedLiquidity) returns()
func (_TestPoolService *TestPoolServiceTransactor) SetExpectedLiquidity(opts *bind.TransactOpts, newExpectedLiquidity *big.Int) (*types.Transaction, error) {
	return _TestPoolService.contract.Transact(opts, "setExpectedLiquidity", newExpectedLiquidity)
}

// SetExpectedLiquidity is a paid mutator transaction binding the contract method 0x7dfc6afe.
//
// Solidity: function setExpectedLiquidity(uint256 newExpectedLiquidity) returns()
func (_TestPoolService *TestPoolServiceSession) SetExpectedLiquidity(newExpectedLiquidity *big.Int) (*types.Transaction, error) {
	return _TestPoolService.Contract.SetExpectedLiquidity(&_TestPoolService.TransactOpts, newExpectedLiquidity)
}

// SetExpectedLiquidity is a paid mutator transaction binding the contract method 0x7dfc6afe.
//
// Solidity: function setExpectedLiquidity(uint256 newExpectedLiquidity) returns()
func (_TestPoolService *TestPoolServiceTransactorSession) SetExpectedLiquidity(newExpectedLiquidity *big.Int) (*types.Transaction, error) {
	return _TestPoolService.Contract.SetExpectedLiquidity(&_TestPoolService.TransactOpts, newExpectedLiquidity)
}

// SetExpectedLiquidityLimit is a paid mutator transaction binding the contract method 0xbb04b193.
//
// Solidity: function setExpectedLiquidityLimit(uint256 newLimit) returns()
func (_TestPoolService *TestPoolServiceTransactor) SetExpectedLiquidityLimit(opts *bind.TransactOpts, newLimit *big.Int) (*types.Transaction, error) {
	return _TestPoolService.contract.Transact(opts, "setExpectedLiquidityLimit", newLimit)
}

// SetExpectedLiquidityLimit is a paid mutator transaction binding the contract method 0xbb04b193.
//
// Solidity: function setExpectedLiquidityLimit(uint256 newLimit) returns()
func (_TestPoolService *TestPoolServiceSession) SetExpectedLiquidityLimit(newLimit *big.Int) (*types.Transaction, error) {
	return _TestPoolService.Contract.SetExpectedLiquidityLimit(&_TestPoolService.TransactOpts, newLimit)
}

// SetExpectedLiquidityLimit is a paid mutator transaction binding the contract method 0xbb04b193.
//
// Solidity: function setExpectedLiquidityLimit(uint256 newLimit) returns()
func (_TestPoolService *TestPoolServiceTransactorSession) SetExpectedLiquidityLimit(newLimit *big.Int) (*types.Transaction, error) {
	return _TestPoolService.Contract.SetExpectedLiquidityLimit(&_TestPoolService.TransactOpts, newLimit)
}

// SetWithdrawFee is a paid mutator transaction binding the contract method 0xb6ac642a.
//
// Solidity: function setWithdrawFee(uint256 fee) returns()
func (_TestPoolService *TestPoolServiceTransactor) SetWithdrawFee(opts *bind.TransactOpts, fee *big.Int) (*types.Transaction, error) {
	return _TestPoolService.contract.Transact(opts, "setWithdrawFee", fee)
}

// SetWithdrawFee is a paid mutator transaction binding the contract method 0xb6ac642a.
//
// Solidity: function setWithdrawFee(uint256 fee) returns()
func (_TestPoolService *TestPoolServiceSession) SetWithdrawFee(fee *big.Int) (*types.Transaction, error) {
	return _TestPoolService.Contract.SetWithdrawFee(&_TestPoolService.TransactOpts, fee)
}

// SetWithdrawFee is a paid mutator transaction binding the contract method 0xb6ac642a.
//
// Solidity: function setWithdrawFee(uint256 fee) returns()
func (_TestPoolService *TestPoolServiceTransactorSession) SetWithdrawFee(fee *big.Int) (*types.Transaction, error) {
	return _TestPoolService.Contract.SetWithdrawFee(&_TestPoolService.TransactOpts, fee)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_TestPoolService *TestPoolServiceTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestPoolService.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_TestPoolService *TestPoolServiceSession) Unpause() (*types.Transaction, error) {
	return _TestPoolService.Contract.Unpause(&_TestPoolService.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_TestPoolService *TestPoolServiceTransactorSession) Unpause() (*types.Transaction, error) {
	return _TestPoolService.Contract.Unpause(&_TestPoolService.TransactOpts)
}

// UpdateBorrowRate is a paid mutator transaction binding the contract method 0xd41ee0f7.
//
// Solidity: function updateBorrowRate() returns()
func (_TestPoolService *TestPoolServiceTransactor) UpdateBorrowRate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestPoolService.contract.Transact(opts, "updateBorrowRate")
}

// UpdateBorrowRate is a paid mutator transaction binding the contract method 0xd41ee0f7.
//
// Solidity: function updateBorrowRate() returns()
func (_TestPoolService *TestPoolServiceSession) UpdateBorrowRate() (*types.Transaction, error) {
	return _TestPoolService.Contract.UpdateBorrowRate(&_TestPoolService.TransactOpts)
}

// UpdateBorrowRate is a paid mutator transaction binding the contract method 0xd41ee0f7.
//
// Solidity: function updateBorrowRate() returns()
func (_TestPoolService *TestPoolServiceTransactorSession) UpdateBorrowRate() (*types.Transaction, error) {
	return _TestPoolService.Contract.UpdateBorrowRate(&_TestPoolService.TransactOpts)
}

// UpdateInterestRateModel is a paid mutator transaction binding the contract method 0x5664cacf.
//
// Solidity: function updateInterestRateModel(address _interestRateModel) returns()
func (_TestPoolService *TestPoolServiceTransactor) UpdateInterestRateModel(opts *bind.TransactOpts, _interestRateModel common.Address) (*types.Transaction, error) {
	return _TestPoolService.contract.Transact(opts, "updateInterestRateModel", _interestRateModel)
}

// UpdateInterestRateModel is a paid mutator transaction binding the contract method 0x5664cacf.
//
// Solidity: function updateInterestRateModel(address _interestRateModel) returns()
func (_TestPoolService *TestPoolServiceSession) UpdateInterestRateModel(_interestRateModel common.Address) (*types.Transaction, error) {
	return _TestPoolService.Contract.UpdateInterestRateModel(&_TestPoolService.TransactOpts, _interestRateModel)
}

// UpdateInterestRateModel is a paid mutator transaction binding the contract method 0x5664cacf.
//
// Solidity: function updateInterestRateModel(address _interestRateModel) returns()
func (_TestPoolService *TestPoolServiceTransactorSession) UpdateInterestRateModel(_interestRateModel common.Address) (*types.Transaction, error) {
	return _TestPoolService.Contract.UpdateInterestRateModel(&_TestPoolService.TransactOpts, _interestRateModel)
}

// TestPoolServiceAddLiquidityIterator is returned from FilterAddLiquidity and is used to iterate over the raw logs and unpacked data for AddLiquidity events raised by the TestPoolService contract.
type TestPoolServiceAddLiquidityIterator struct {
	Event *TestPoolServiceAddLiquidity // Event containing the contract specifics and raw log

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
func (it *TestPoolServiceAddLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestPoolServiceAddLiquidity)
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
		it.Event = new(TestPoolServiceAddLiquidity)
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
func (it *TestPoolServiceAddLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestPoolServiceAddLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestPoolServiceAddLiquidity represents a AddLiquidity event raised by the TestPoolService contract.
type TestPoolServiceAddLiquidity struct {
	Sender       common.Address
	OnBehalfOf   common.Address
	Amount       *big.Int
	ReferralCode *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAddLiquidity is a free log retrieval operation binding the contract event 0xd2491a9b4fe81a7cd4511e8b7b7743951b061dad5bed7da8a7795b080ee08c7e.
//
// Solidity: event AddLiquidity(address indexed sender, address indexed onBehalfOf, uint256 amount, uint256 referralCode)
func (_TestPoolService *TestPoolServiceFilterer) FilterAddLiquidity(opts *bind.FilterOpts, sender []common.Address, onBehalfOf []common.Address) (*TestPoolServiceAddLiquidityIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	logs, sub, err := _TestPoolService.contract.FilterLogs(opts, "AddLiquidity", senderRule, onBehalfOfRule)
	if err != nil {
		return nil, err
	}
	return &TestPoolServiceAddLiquidityIterator{contract: _TestPoolService.contract, event: "AddLiquidity", logs: logs, sub: sub}, nil
}

// WatchAddLiquidity is a free log subscription operation binding the contract event 0xd2491a9b4fe81a7cd4511e8b7b7743951b061dad5bed7da8a7795b080ee08c7e.
//
// Solidity: event AddLiquidity(address indexed sender, address indexed onBehalfOf, uint256 amount, uint256 referralCode)
func (_TestPoolService *TestPoolServiceFilterer) WatchAddLiquidity(opts *bind.WatchOpts, sink chan<- *TestPoolServiceAddLiquidity, sender []common.Address, onBehalfOf []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	logs, sub, err := _TestPoolService.contract.WatchLogs(opts, "AddLiquidity", senderRule, onBehalfOfRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestPoolServiceAddLiquidity)
				if err := _TestPoolService.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
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
func (_TestPoolService *TestPoolServiceFilterer) ParseAddLiquidity(log types.Log) (*TestPoolServiceAddLiquidity, error) {
	event := new(TestPoolServiceAddLiquidity)
	if err := _TestPoolService.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestPoolServiceBorrowIterator is returned from FilterBorrow and is used to iterate over the raw logs and unpacked data for Borrow events raised by the TestPoolService contract.
type TestPoolServiceBorrowIterator struct {
	Event *TestPoolServiceBorrow // Event containing the contract specifics and raw log

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
func (it *TestPoolServiceBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestPoolServiceBorrow)
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
		it.Event = new(TestPoolServiceBorrow)
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
func (it *TestPoolServiceBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestPoolServiceBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestPoolServiceBorrow represents a Borrow event raised by the TestPoolService contract.
type TestPoolServiceBorrow struct {
	CreditManager common.Address
	CreditAccount common.Address
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBorrow is a free log retrieval operation binding the contract event 0x312a5e5e1079f5dda4e95dbbd0b908b291fd5b992ef22073643ab691572c5b52.
//
// Solidity: event Borrow(address indexed creditManager, address indexed creditAccount, uint256 amount)
func (_TestPoolService *TestPoolServiceFilterer) FilterBorrow(opts *bind.FilterOpts, creditManager []common.Address, creditAccount []common.Address) (*TestPoolServiceBorrowIterator, error) {

	var creditManagerRule []interface{}
	for _, creditManagerItem := range creditManager {
		creditManagerRule = append(creditManagerRule, creditManagerItem)
	}
	var creditAccountRule []interface{}
	for _, creditAccountItem := range creditAccount {
		creditAccountRule = append(creditAccountRule, creditAccountItem)
	}

	logs, sub, err := _TestPoolService.contract.FilterLogs(opts, "Borrow", creditManagerRule, creditAccountRule)
	if err != nil {
		return nil, err
	}
	return &TestPoolServiceBorrowIterator{contract: _TestPoolService.contract, event: "Borrow", logs: logs, sub: sub}, nil
}

// WatchBorrow is a free log subscription operation binding the contract event 0x312a5e5e1079f5dda4e95dbbd0b908b291fd5b992ef22073643ab691572c5b52.
//
// Solidity: event Borrow(address indexed creditManager, address indexed creditAccount, uint256 amount)
func (_TestPoolService *TestPoolServiceFilterer) WatchBorrow(opts *bind.WatchOpts, sink chan<- *TestPoolServiceBorrow, creditManager []common.Address, creditAccount []common.Address) (event.Subscription, error) {

	var creditManagerRule []interface{}
	for _, creditManagerItem := range creditManager {
		creditManagerRule = append(creditManagerRule, creditManagerItem)
	}
	var creditAccountRule []interface{}
	for _, creditAccountItem := range creditAccount {
		creditAccountRule = append(creditAccountRule, creditAccountItem)
	}

	logs, sub, err := _TestPoolService.contract.WatchLogs(opts, "Borrow", creditManagerRule, creditAccountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestPoolServiceBorrow)
				if err := _TestPoolService.contract.UnpackLog(event, "Borrow", log); err != nil {
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
func (_TestPoolService *TestPoolServiceFilterer) ParseBorrow(log types.Log) (*TestPoolServiceBorrow, error) {
	event := new(TestPoolServiceBorrow)
	if err := _TestPoolService.contract.UnpackLog(event, "Borrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestPoolServiceBorrowForbiddenIterator is returned from FilterBorrowForbidden and is used to iterate over the raw logs and unpacked data for BorrowForbidden events raised by the TestPoolService contract.
type TestPoolServiceBorrowForbiddenIterator struct {
	Event *TestPoolServiceBorrowForbidden // Event containing the contract specifics and raw log

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
func (it *TestPoolServiceBorrowForbiddenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestPoolServiceBorrowForbidden)
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
		it.Event = new(TestPoolServiceBorrowForbidden)
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
func (it *TestPoolServiceBorrowForbiddenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestPoolServiceBorrowForbiddenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestPoolServiceBorrowForbidden represents a BorrowForbidden event raised by the TestPoolService contract.
type TestPoolServiceBorrowForbidden struct {
	CreditManager common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBorrowForbidden is a free log retrieval operation binding the contract event 0x9181736fce85d2d4cca2e4406f10679302ae5c387180fdb62963af3cd9a24fd6.
//
// Solidity: event BorrowForbidden(address indexed creditManager)
func (_TestPoolService *TestPoolServiceFilterer) FilterBorrowForbidden(opts *bind.FilterOpts, creditManager []common.Address) (*TestPoolServiceBorrowForbiddenIterator, error) {

	var creditManagerRule []interface{}
	for _, creditManagerItem := range creditManager {
		creditManagerRule = append(creditManagerRule, creditManagerItem)
	}

	logs, sub, err := _TestPoolService.contract.FilterLogs(opts, "BorrowForbidden", creditManagerRule)
	if err != nil {
		return nil, err
	}
	return &TestPoolServiceBorrowForbiddenIterator{contract: _TestPoolService.contract, event: "BorrowForbidden", logs: logs, sub: sub}, nil
}

// WatchBorrowForbidden is a free log subscription operation binding the contract event 0x9181736fce85d2d4cca2e4406f10679302ae5c387180fdb62963af3cd9a24fd6.
//
// Solidity: event BorrowForbidden(address indexed creditManager)
func (_TestPoolService *TestPoolServiceFilterer) WatchBorrowForbidden(opts *bind.WatchOpts, sink chan<- *TestPoolServiceBorrowForbidden, creditManager []common.Address) (event.Subscription, error) {

	var creditManagerRule []interface{}
	for _, creditManagerItem := range creditManager {
		creditManagerRule = append(creditManagerRule, creditManagerItem)
	}

	logs, sub, err := _TestPoolService.contract.WatchLogs(opts, "BorrowForbidden", creditManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestPoolServiceBorrowForbidden)
				if err := _TestPoolService.contract.UnpackLog(event, "BorrowForbidden", log); err != nil {
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
func (_TestPoolService *TestPoolServiceFilterer) ParseBorrowForbidden(log types.Log) (*TestPoolServiceBorrowForbidden, error) {
	event := new(TestPoolServiceBorrowForbidden)
	if err := _TestPoolService.contract.UnpackLog(event, "BorrowForbidden", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestPoolServiceNewCreditManagerConnectedIterator is returned from FilterNewCreditManagerConnected and is used to iterate over the raw logs and unpacked data for NewCreditManagerConnected events raised by the TestPoolService contract.
type TestPoolServiceNewCreditManagerConnectedIterator struct {
	Event *TestPoolServiceNewCreditManagerConnected // Event containing the contract specifics and raw log

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
func (it *TestPoolServiceNewCreditManagerConnectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestPoolServiceNewCreditManagerConnected)
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
		it.Event = new(TestPoolServiceNewCreditManagerConnected)
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
func (it *TestPoolServiceNewCreditManagerConnectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestPoolServiceNewCreditManagerConnectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestPoolServiceNewCreditManagerConnected represents a NewCreditManagerConnected event raised by the TestPoolService contract.
type TestPoolServiceNewCreditManagerConnected struct {
	CreditManager common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNewCreditManagerConnected is a free log retrieval operation binding the contract event 0xe076020e7eac3915d33aec40c24f95e73eb6c9921ff89747d50aa8fd934d2c01.
//
// Solidity: event NewCreditManagerConnected(address indexed creditManager)
func (_TestPoolService *TestPoolServiceFilterer) FilterNewCreditManagerConnected(opts *bind.FilterOpts, creditManager []common.Address) (*TestPoolServiceNewCreditManagerConnectedIterator, error) {

	var creditManagerRule []interface{}
	for _, creditManagerItem := range creditManager {
		creditManagerRule = append(creditManagerRule, creditManagerItem)
	}

	logs, sub, err := _TestPoolService.contract.FilterLogs(opts, "NewCreditManagerConnected", creditManagerRule)
	if err != nil {
		return nil, err
	}
	return &TestPoolServiceNewCreditManagerConnectedIterator{contract: _TestPoolService.contract, event: "NewCreditManagerConnected", logs: logs, sub: sub}, nil
}

// WatchNewCreditManagerConnected is a free log subscription operation binding the contract event 0xe076020e7eac3915d33aec40c24f95e73eb6c9921ff89747d50aa8fd934d2c01.
//
// Solidity: event NewCreditManagerConnected(address indexed creditManager)
func (_TestPoolService *TestPoolServiceFilterer) WatchNewCreditManagerConnected(opts *bind.WatchOpts, sink chan<- *TestPoolServiceNewCreditManagerConnected, creditManager []common.Address) (event.Subscription, error) {

	var creditManagerRule []interface{}
	for _, creditManagerItem := range creditManager {
		creditManagerRule = append(creditManagerRule, creditManagerItem)
	}

	logs, sub, err := _TestPoolService.contract.WatchLogs(opts, "NewCreditManagerConnected", creditManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestPoolServiceNewCreditManagerConnected)
				if err := _TestPoolService.contract.UnpackLog(event, "NewCreditManagerConnected", log); err != nil {
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
func (_TestPoolService *TestPoolServiceFilterer) ParseNewCreditManagerConnected(log types.Log) (*TestPoolServiceNewCreditManagerConnected, error) {
	event := new(TestPoolServiceNewCreditManagerConnected)
	if err := _TestPoolService.contract.UnpackLog(event, "NewCreditManagerConnected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestPoolServiceNewExpectedLiquidityLimitIterator is returned from FilterNewExpectedLiquidityLimit and is used to iterate over the raw logs and unpacked data for NewExpectedLiquidityLimit events raised by the TestPoolService contract.
type TestPoolServiceNewExpectedLiquidityLimitIterator struct {
	Event *TestPoolServiceNewExpectedLiquidityLimit // Event containing the contract specifics and raw log

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
func (it *TestPoolServiceNewExpectedLiquidityLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestPoolServiceNewExpectedLiquidityLimit)
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
		it.Event = new(TestPoolServiceNewExpectedLiquidityLimit)
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
func (it *TestPoolServiceNewExpectedLiquidityLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestPoolServiceNewExpectedLiquidityLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestPoolServiceNewExpectedLiquidityLimit represents a NewExpectedLiquidityLimit event raised by the TestPoolService contract.
type TestPoolServiceNewExpectedLiquidityLimit struct {
	NewLimit *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewExpectedLiquidityLimit is a free log retrieval operation binding the contract event 0xd7a183c9fe85b604c25d54bd676e0866f6c13bcca9fb9b0850213de118fdc99c.
//
// Solidity: event NewExpectedLiquidityLimit(uint256 newLimit)
func (_TestPoolService *TestPoolServiceFilterer) FilterNewExpectedLiquidityLimit(opts *bind.FilterOpts) (*TestPoolServiceNewExpectedLiquidityLimitIterator, error) {

	logs, sub, err := _TestPoolService.contract.FilterLogs(opts, "NewExpectedLiquidityLimit")
	if err != nil {
		return nil, err
	}
	return &TestPoolServiceNewExpectedLiquidityLimitIterator{contract: _TestPoolService.contract, event: "NewExpectedLiquidityLimit", logs: logs, sub: sub}, nil
}

// WatchNewExpectedLiquidityLimit is a free log subscription operation binding the contract event 0xd7a183c9fe85b604c25d54bd676e0866f6c13bcca9fb9b0850213de118fdc99c.
//
// Solidity: event NewExpectedLiquidityLimit(uint256 newLimit)
func (_TestPoolService *TestPoolServiceFilterer) WatchNewExpectedLiquidityLimit(opts *bind.WatchOpts, sink chan<- *TestPoolServiceNewExpectedLiquidityLimit) (event.Subscription, error) {

	logs, sub, err := _TestPoolService.contract.WatchLogs(opts, "NewExpectedLiquidityLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestPoolServiceNewExpectedLiquidityLimit)
				if err := _TestPoolService.contract.UnpackLog(event, "NewExpectedLiquidityLimit", log); err != nil {
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
func (_TestPoolService *TestPoolServiceFilterer) ParseNewExpectedLiquidityLimit(log types.Log) (*TestPoolServiceNewExpectedLiquidityLimit, error) {
	event := new(TestPoolServiceNewExpectedLiquidityLimit)
	if err := _TestPoolService.contract.UnpackLog(event, "NewExpectedLiquidityLimit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestPoolServiceNewInterestRateModelIterator is returned from FilterNewInterestRateModel and is used to iterate over the raw logs and unpacked data for NewInterestRateModel events raised by the TestPoolService contract.
type TestPoolServiceNewInterestRateModelIterator struct {
	Event *TestPoolServiceNewInterestRateModel // Event containing the contract specifics and raw log

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
func (it *TestPoolServiceNewInterestRateModelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestPoolServiceNewInterestRateModel)
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
		it.Event = new(TestPoolServiceNewInterestRateModel)
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
func (it *TestPoolServiceNewInterestRateModelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestPoolServiceNewInterestRateModelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestPoolServiceNewInterestRateModel represents a NewInterestRateModel event raised by the TestPoolService contract.
type TestPoolServiceNewInterestRateModel struct {
	NewInterestRateModel common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterNewInterestRateModel is a free log retrieval operation binding the contract event 0x0ec6cb7631d36954a05ffd646135bfd9995c71e7fa36d26abb1ad9f24a040ea1.
//
// Solidity: event NewInterestRateModel(address indexed newInterestRateModel)
func (_TestPoolService *TestPoolServiceFilterer) FilterNewInterestRateModel(opts *bind.FilterOpts, newInterestRateModel []common.Address) (*TestPoolServiceNewInterestRateModelIterator, error) {

	var newInterestRateModelRule []interface{}
	for _, newInterestRateModelItem := range newInterestRateModel {
		newInterestRateModelRule = append(newInterestRateModelRule, newInterestRateModelItem)
	}

	logs, sub, err := _TestPoolService.contract.FilterLogs(opts, "NewInterestRateModel", newInterestRateModelRule)
	if err != nil {
		return nil, err
	}
	return &TestPoolServiceNewInterestRateModelIterator{contract: _TestPoolService.contract, event: "NewInterestRateModel", logs: logs, sub: sub}, nil
}

// WatchNewInterestRateModel is a free log subscription operation binding the contract event 0x0ec6cb7631d36954a05ffd646135bfd9995c71e7fa36d26abb1ad9f24a040ea1.
//
// Solidity: event NewInterestRateModel(address indexed newInterestRateModel)
func (_TestPoolService *TestPoolServiceFilterer) WatchNewInterestRateModel(opts *bind.WatchOpts, sink chan<- *TestPoolServiceNewInterestRateModel, newInterestRateModel []common.Address) (event.Subscription, error) {

	var newInterestRateModelRule []interface{}
	for _, newInterestRateModelItem := range newInterestRateModel {
		newInterestRateModelRule = append(newInterestRateModelRule, newInterestRateModelItem)
	}

	logs, sub, err := _TestPoolService.contract.WatchLogs(opts, "NewInterestRateModel", newInterestRateModelRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestPoolServiceNewInterestRateModel)
				if err := _TestPoolService.contract.UnpackLog(event, "NewInterestRateModel", log); err != nil {
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
func (_TestPoolService *TestPoolServiceFilterer) ParseNewInterestRateModel(log types.Log) (*TestPoolServiceNewInterestRateModel, error) {
	event := new(TestPoolServiceNewInterestRateModel)
	if err := _TestPoolService.contract.UnpackLog(event, "NewInterestRateModel", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestPoolServiceNewWithdrawFeeIterator is returned from FilterNewWithdrawFee and is used to iterate over the raw logs and unpacked data for NewWithdrawFee events raised by the TestPoolService contract.
type TestPoolServiceNewWithdrawFeeIterator struct {
	Event *TestPoolServiceNewWithdrawFee // Event containing the contract specifics and raw log

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
func (it *TestPoolServiceNewWithdrawFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestPoolServiceNewWithdrawFee)
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
		it.Event = new(TestPoolServiceNewWithdrawFee)
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
func (it *TestPoolServiceNewWithdrawFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestPoolServiceNewWithdrawFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestPoolServiceNewWithdrawFee represents a NewWithdrawFee event raised by the TestPoolService contract.
type TestPoolServiceNewWithdrawFee struct {
	Fee *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterNewWithdrawFee is a free log retrieval operation binding the contract event 0xd5fe46099fa396290a7f57e36c3c3c8774e2562c18ed5d1dcc0fa75071e03f1d.
//
// Solidity: event NewWithdrawFee(uint256 fee)
func (_TestPoolService *TestPoolServiceFilterer) FilterNewWithdrawFee(opts *bind.FilterOpts) (*TestPoolServiceNewWithdrawFeeIterator, error) {

	logs, sub, err := _TestPoolService.contract.FilterLogs(opts, "NewWithdrawFee")
	if err != nil {
		return nil, err
	}
	return &TestPoolServiceNewWithdrawFeeIterator{contract: _TestPoolService.contract, event: "NewWithdrawFee", logs: logs, sub: sub}, nil
}

// WatchNewWithdrawFee is a free log subscription operation binding the contract event 0xd5fe46099fa396290a7f57e36c3c3c8774e2562c18ed5d1dcc0fa75071e03f1d.
//
// Solidity: event NewWithdrawFee(uint256 fee)
func (_TestPoolService *TestPoolServiceFilterer) WatchNewWithdrawFee(opts *bind.WatchOpts, sink chan<- *TestPoolServiceNewWithdrawFee) (event.Subscription, error) {

	logs, sub, err := _TestPoolService.contract.WatchLogs(opts, "NewWithdrawFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestPoolServiceNewWithdrawFee)
				if err := _TestPoolService.contract.UnpackLog(event, "NewWithdrawFee", log); err != nil {
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
func (_TestPoolService *TestPoolServiceFilterer) ParseNewWithdrawFee(log types.Log) (*TestPoolServiceNewWithdrawFee, error) {
	event := new(TestPoolServiceNewWithdrawFee)
	if err := _TestPoolService.contract.UnpackLog(event, "NewWithdrawFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestPoolServicePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the TestPoolService contract.
type TestPoolServicePausedIterator struct {
	Event *TestPoolServicePaused // Event containing the contract specifics and raw log

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
func (it *TestPoolServicePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestPoolServicePaused)
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
		it.Event = new(TestPoolServicePaused)
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
func (it *TestPoolServicePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestPoolServicePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestPoolServicePaused represents a Paused event raised by the TestPoolService contract.
type TestPoolServicePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_TestPoolService *TestPoolServiceFilterer) FilterPaused(opts *bind.FilterOpts) (*TestPoolServicePausedIterator, error) {

	logs, sub, err := _TestPoolService.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &TestPoolServicePausedIterator{contract: _TestPoolService.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_TestPoolService *TestPoolServiceFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *TestPoolServicePaused) (event.Subscription, error) {

	logs, sub, err := _TestPoolService.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestPoolServicePaused)
				if err := _TestPoolService.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_TestPoolService *TestPoolServiceFilterer) ParsePaused(log types.Log) (*TestPoolServicePaused, error) {
	event := new(TestPoolServicePaused)
	if err := _TestPoolService.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestPoolServiceRemoveLiquidityIterator is returned from FilterRemoveLiquidity and is used to iterate over the raw logs and unpacked data for RemoveLiquidity events raised by the TestPoolService contract.
type TestPoolServiceRemoveLiquidityIterator struct {
	Event *TestPoolServiceRemoveLiquidity // Event containing the contract specifics and raw log

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
func (it *TestPoolServiceRemoveLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestPoolServiceRemoveLiquidity)
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
		it.Event = new(TestPoolServiceRemoveLiquidity)
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
func (it *TestPoolServiceRemoveLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestPoolServiceRemoveLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestPoolServiceRemoveLiquidity represents a RemoveLiquidity event raised by the TestPoolService contract.
type TestPoolServiceRemoveLiquidity struct {
	Sender common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidity is a free log retrieval operation binding the contract event 0xd8ae9b9ba89e637bcb66a69ac91e8f688018e81d6f92c57e02226425c8efbdf6.
//
// Solidity: event RemoveLiquidity(address indexed sender, address indexed to, uint256 amount)
func (_TestPoolService *TestPoolServiceFilterer) FilterRemoveLiquidity(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*TestPoolServiceRemoveLiquidityIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TestPoolService.contract.FilterLogs(opts, "RemoveLiquidity", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TestPoolServiceRemoveLiquidityIterator{contract: _TestPoolService.contract, event: "RemoveLiquidity", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidity is a free log subscription operation binding the contract event 0xd8ae9b9ba89e637bcb66a69ac91e8f688018e81d6f92c57e02226425c8efbdf6.
//
// Solidity: event RemoveLiquidity(address indexed sender, address indexed to, uint256 amount)
func (_TestPoolService *TestPoolServiceFilterer) WatchRemoveLiquidity(opts *bind.WatchOpts, sink chan<- *TestPoolServiceRemoveLiquidity, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TestPoolService.contract.WatchLogs(opts, "RemoveLiquidity", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestPoolServiceRemoveLiquidity)
				if err := _TestPoolService.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
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
func (_TestPoolService *TestPoolServiceFilterer) ParseRemoveLiquidity(log types.Log) (*TestPoolServiceRemoveLiquidity, error) {
	event := new(TestPoolServiceRemoveLiquidity)
	if err := _TestPoolService.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestPoolServiceRepayIterator is returned from FilterRepay and is used to iterate over the raw logs and unpacked data for Repay events raised by the TestPoolService contract.
type TestPoolServiceRepayIterator struct {
	Event *TestPoolServiceRepay // Event containing the contract specifics and raw log

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
func (it *TestPoolServiceRepayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestPoolServiceRepay)
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
		it.Event = new(TestPoolServiceRepay)
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
func (it *TestPoolServiceRepayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestPoolServiceRepayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestPoolServiceRepay represents a Repay event raised by the TestPoolService contract.
type TestPoolServiceRepay struct {
	CreditManager  common.Address
	BorrowedAmount *big.Int
	Profit         *big.Int
	Loss           *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRepay is a free log retrieval operation binding the contract event 0x2fe77b1c99aca6b022b8efc6e3e8dd1b48b30748709339b65c50ef3263443e09.
//
// Solidity: event Repay(address indexed creditManager, uint256 borrowedAmount, uint256 profit, uint256 loss)
func (_TestPoolService *TestPoolServiceFilterer) FilterRepay(opts *bind.FilterOpts, creditManager []common.Address) (*TestPoolServiceRepayIterator, error) {

	var creditManagerRule []interface{}
	for _, creditManagerItem := range creditManager {
		creditManagerRule = append(creditManagerRule, creditManagerItem)
	}

	logs, sub, err := _TestPoolService.contract.FilterLogs(opts, "Repay", creditManagerRule)
	if err != nil {
		return nil, err
	}
	return &TestPoolServiceRepayIterator{contract: _TestPoolService.contract, event: "Repay", logs: logs, sub: sub}, nil
}

// WatchRepay is a free log subscription operation binding the contract event 0x2fe77b1c99aca6b022b8efc6e3e8dd1b48b30748709339b65c50ef3263443e09.
//
// Solidity: event Repay(address indexed creditManager, uint256 borrowedAmount, uint256 profit, uint256 loss)
func (_TestPoolService *TestPoolServiceFilterer) WatchRepay(opts *bind.WatchOpts, sink chan<- *TestPoolServiceRepay, creditManager []common.Address) (event.Subscription, error) {

	var creditManagerRule []interface{}
	for _, creditManagerItem := range creditManager {
		creditManagerRule = append(creditManagerRule, creditManagerItem)
	}

	logs, sub, err := _TestPoolService.contract.WatchLogs(opts, "Repay", creditManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestPoolServiceRepay)
				if err := _TestPoolService.contract.UnpackLog(event, "Repay", log); err != nil {
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
func (_TestPoolService *TestPoolServiceFilterer) ParseRepay(log types.Log) (*TestPoolServiceRepay, error) {
	event := new(TestPoolServiceRepay)
	if err := _TestPoolService.contract.UnpackLog(event, "Repay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestPoolServiceUncoveredLossIterator is returned from FilterUncoveredLoss and is used to iterate over the raw logs and unpacked data for UncoveredLoss events raised by the TestPoolService contract.
type TestPoolServiceUncoveredLossIterator struct {
	Event *TestPoolServiceUncoveredLoss // Event containing the contract specifics and raw log

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
func (it *TestPoolServiceUncoveredLossIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestPoolServiceUncoveredLoss)
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
		it.Event = new(TestPoolServiceUncoveredLoss)
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
func (it *TestPoolServiceUncoveredLossIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestPoolServiceUncoveredLossIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestPoolServiceUncoveredLoss represents a UncoveredLoss event raised by the TestPoolService contract.
type TestPoolServiceUncoveredLoss struct {
	CreditManager common.Address
	Loss          *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUncoveredLoss is a free log retrieval operation binding the contract event 0xef3653ded679720ab04913b6f3820be7cedc8286d42ff5dd8dff17e91bd2964c.
//
// Solidity: event UncoveredLoss(address indexed creditManager, uint256 loss)
func (_TestPoolService *TestPoolServiceFilterer) FilterUncoveredLoss(opts *bind.FilterOpts, creditManager []common.Address) (*TestPoolServiceUncoveredLossIterator, error) {

	var creditManagerRule []interface{}
	for _, creditManagerItem := range creditManager {
		creditManagerRule = append(creditManagerRule, creditManagerItem)
	}

	logs, sub, err := _TestPoolService.contract.FilterLogs(opts, "UncoveredLoss", creditManagerRule)
	if err != nil {
		return nil, err
	}
	return &TestPoolServiceUncoveredLossIterator{contract: _TestPoolService.contract, event: "UncoveredLoss", logs: logs, sub: sub}, nil
}

// WatchUncoveredLoss is a free log subscription operation binding the contract event 0xef3653ded679720ab04913b6f3820be7cedc8286d42ff5dd8dff17e91bd2964c.
//
// Solidity: event UncoveredLoss(address indexed creditManager, uint256 loss)
func (_TestPoolService *TestPoolServiceFilterer) WatchUncoveredLoss(opts *bind.WatchOpts, sink chan<- *TestPoolServiceUncoveredLoss, creditManager []common.Address) (event.Subscription, error) {

	var creditManagerRule []interface{}
	for _, creditManagerItem := range creditManager {
		creditManagerRule = append(creditManagerRule, creditManagerItem)
	}

	logs, sub, err := _TestPoolService.contract.WatchLogs(opts, "UncoveredLoss", creditManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestPoolServiceUncoveredLoss)
				if err := _TestPoolService.contract.UnpackLog(event, "UncoveredLoss", log); err != nil {
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
func (_TestPoolService *TestPoolServiceFilterer) ParseUncoveredLoss(log types.Log) (*TestPoolServiceUncoveredLoss, error) {
	event := new(TestPoolServiceUncoveredLoss)
	if err := _TestPoolService.contract.UnpackLog(event, "UncoveredLoss", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestPoolServiceUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the TestPoolService contract.
type TestPoolServiceUnpausedIterator struct {
	Event *TestPoolServiceUnpaused // Event containing the contract specifics and raw log

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
func (it *TestPoolServiceUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestPoolServiceUnpaused)
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
		it.Event = new(TestPoolServiceUnpaused)
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
func (it *TestPoolServiceUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestPoolServiceUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestPoolServiceUnpaused represents a Unpaused event raised by the TestPoolService contract.
type TestPoolServiceUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_TestPoolService *TestPoolServiceFilterer) FilterUnpaused(opts *bind.FilterOpts) (*TestPoolServiceUnpausedIterator, error) {

	logs, sub, err := _TestPoolService.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &TestPoolServiceUnpausedIterator{contract: _TestPoolService.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_TestPoolService *TestPoolServiceFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *TestPoolServiceUnpaused) (event.Subscription, error) {

	logs, sub, err := _TestPoolService.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestPoolServiceUnpaused)
				if err := _TestPoolService.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_TestPoolService *TestPoolServiceFilterer) ParseUnpaused(log types.Log) (*TestPoolServiceUnpaused, error) {
	event := new(TestPoolServiceUnpaused)
	if err := _TestPoolService.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
