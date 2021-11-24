// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mockPoolService

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

// MockPoolServiceMetaData contains all meta data concerning the MockPoolService contract.
var MockPoolServiceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_underlyingToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"referralCode\",\"type\":\"uint256\"}],\"name\":\"AddLiquidity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creditManager\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Borrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creditManager\",\"type\":\"address\"}],\"name\":\"BorrowForbidden\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creditManager\",\"type\":\"address\"}],\"name\":\"NewCreditManagerConnected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newLimit\",\"type\":\"uint256\"}],\"name\":\"NewExpectedLiquidityLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newInterestRateModel\",\"type\":\"address\"}],\"name\":\"NewInterestRateModel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"NewWithdrawFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RemoveLiquidity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creditManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"borrowedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"profit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"loss\",\"type\":\"uint256\"}],\"name\":\"Repay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creditManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"loss\",\"type\":\"uint256\"}],\"name\":\"UncoveredLoss\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_cumulativeIndex_RAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_expectedLiquidityLU\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_timestampLU\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"referralCode\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addressProvider\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"availableLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"borrowAPY_RAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cumulativeIndexAtOpen\",\"type\":\"uint256\"}],\"name\":\"calcCumulativeIndexAtBorrowMore\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calcLinearCumulative_RAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calcLinearIndex_RAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_creditManager\",\"type\":\"address\"}],\"name\":\"connectCreditManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"creditManagers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"creditManagersCanBorrow\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"creditManagersCanRepay\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creditManagersCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dieselToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expectedLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expectedLiquidityLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_creditManager\",\"type\":\"address\"}],\"name\":\"forbidCreditManagerToBorrow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"fromDiesel\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDieselRate_RAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"interestRateModel\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lendAccount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lendAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"borrowedAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"}],\"name\":\"lendCreditAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_interestRateModel\",\"type\":\"address\"}],\"name\":\"newInterestRateModel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"removeLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"repayAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"borrowedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"profit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"loss\",\"type\":\"uint256\"}],\"name\":\"repayCreditAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"repayLoss\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"repayProfit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_cumulativeIndex_RAY\",\"type\":\"uint256\"}],\"name\":\"setCumulative_RAY\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"setExpectedLiquidityLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"setWithdrawFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"toDiesel\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBorrowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasuryAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlyingToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MockPoolServiceABI is the input ABI used to generate the binding from.
// Deprecated: Use MockPoolServiceMetaData.ABI instead.
var MockPoolServiceABI = MockPoolServiceMetaData.ABI

// MockPoolService is an auto generated Go binding around an Ethereum contract.
type MockPoolService struct {
	MockPoolServiceCaller     // Read-only binding to the contract
	MockPoolServiceTransactor // Write-only binding to the contract
	MockPoolServiceFilterer   // Log filterer for contract events
}

// MockPoolServiceCaller is an auto generated read-only Go binding around an Ethereum contract.
type MockPoolServiceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockPoolServiceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MockPoolServiceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockPoolServiceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MockPoolServiceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockPoolServiceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MockPoolServiceSession struct {
	Contract     *MockPoolService  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MockPoolServiceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MockPoolServiceCallerSession struct {
	Contract *MockPoolServiceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// MockPoolServiceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MockPoolServiceTransactorSession struct {
	Contract     *MockPoolServiceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// MockPoolServiceRaw is an auto generated low-level Go binding around an Ethereum contract.
type MockPoolServiceRaw struct {
	Contract *MockPoolService // Generic contract binding to access the raw methods on
}

// MockPoolServiceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MockPoolServiceCallerRaw struct {
	Contract *MockPoolServiceCaller // Generic read-only contract binding to access the raw methods on
}

// MockPoolServiceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MockPoolServiceTransactorRaw struct {
	Contract *MockPoolServiceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMockPoolService creates a new instance of MockPoolService, bound to a specific deployed contract.
func NewMockPoolService(address common.Address, backend bind.ContractBackend) (*MockPoolService, error) {
	contract, err := bindMockPoolService(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockPoolService{MockPoolServiceCaller: MockPoolServiceCaller{contract: contract}, MockPoolServiceTransactor: MockPoolServiceTransactor{contract: contract}, MockPoolServiceFilterer: MockPoolServiceFilterer{contract: contract}}, nil
}

// NewMockPoolServiceCaller creates a new read-only instance of MockPoolService, bound to a specific deployed contract.
func NewMockPoolServiceCaller(address common.Address, caller bind.ContractCaller) (*MockPoolServiceCaller, error) {
	contract, err := bindMockPoolService(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockPoolServiceCaller{contract: contract}, nil
}

// NewMockPoolServiceTransactor creates a new write-only instance of MockPoolService, bound to a specific deployed contract.
func NewMockPoolServiceTransactor(address common.Address, transactor bind.ContractTransactor) (*MockPoolServiceTransactor, error) {
	contract, err := bindMockPoolService(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockPoolServiceTransactor{contract: contract}, nil
}

// NewMockPoolServiceFilterer creates a new log filterer instance of MockPoolService, bound to a specific deployed contract.
func NewMockPoolServiceFilterer(address common.Address, filterer bind.ContractFilterer) (*MockPoolServiceFilterer, error) {
	contract, err := bindMockPoolService(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockPoolServiceFilterer{contract: contract}, nil
}

// bindMockPoolService binds a generic wrapper to an already deployed contract.
func bindMockPoolService(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MockPoolServiceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockPoolService *MockPoolServiceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockPoolService.Contract.MockPoolServiceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockPoolService *MockPoolServiceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockPoolService.Contract.MockPoolServiceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockPoolService *MockPoolServiceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockPoolService.Contract.MockPoolServiceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockPoolService *MockPoolServiceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockPoolService.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockPoolService *MockPoolServiceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockPoolService.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockPoolService *MockPoolServiceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockPoolService.Contract.contract.Transact(opts, method, params...)
}

// CumulativeIndexRAY is a free data retrieval call binding the contract method 0xdbcb313b.
//
// Solidity: function _cumulativeIndex_RAY() view returns(uint256)
func (_MockPoolService *MockPoolServiceCaller) CumulativeIndexRAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MockPoolService.contract.Call(opts, &out, "_cumulativeIndex_RAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CumulativeIndexRAY is a free data retrieval call binding the contract method 0xdbcb313b.
//
// Solidity: function _cumulativeIndex_RAY() view returns(uint256)
func (_MockPoolService *MockPoolServiceSession) CumulativeIndexRAY() (*big.Int, error) {
	return _MockPoolService.Contract.CumulativeIndexRAY(&_MockPoolService.CallOpts)
}

// CumulativeIndexRAY is a free data retrieval call binding the contract method 0xdbcb313b.
//
// Solidity: function _cumulativeIndex_RAY() view returns(uint256)
func (_MockPoolService *MockPoolServiceCallerSession) CumulativeIndexRAY() (*big.Int, error) {
	return _MockPoolService.Contract.CumulativeIndexRAY(&_MockPoolService.CallOpts)
}

// ExpectedLiquidityLU is a free data retrieval call binding the contract method 0x030dbb04.
//
// Solidity: function _expectedLiquidityLU() view returns(uint256)
func (_MockPoolService *MockPoolServiceCaller) ExpectedLiquidityLU(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MockPoolService.contract.Call(opts, &out, "_expectedLiquidityLU")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExpectedLiquidityLU is a free data retrieval call binding the contract method 0x030dbb04.
//
// Solidity: function _expectedLiquidityLU() view returns(uint256)
func (_MockPoolService *MockPoolServiceSession) ExpectedLiquidityLU() (*big.Int, error) {
	return _MockPoolService.Contract.ExpectedLiquidityLU(&_MockPoolService.CallOpts)
}

// ExpectedLiquidityLU is a free data retrieval call binding the contract method 0x030dbb04.
//
// Solidity: function _expectedLiquidityLU() view returns(uint256)
func (_MockPoolService *MockPoolServiceCallerSession) ExpectedLiquidityLU() (*big.Int, error) {
	return _MockPoolService.Contract.ExpectedLiquidityLU(&_MockPoolService.CallOpts)
}

// TimestampLU is a free data retrieval call binding the contract method 0x609ae317.
//
// Solidity: function _timestampLU() view returns(uint256)
func (_MockPoolService *MockPoolServiceCaller) TimestampLU(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MockPoolService.contract.Call(opts, &out, "_timestampLU")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TimestampLU is a free data retrieval call binding the contract method 0x609ae317.
//
// Solidity: function _timestampLU() view returns(uint256)
func (_MockPoolService *MockPoolServiceSession) TimestampLU() (*big.Int, error) {
	return _MockPoolService.Contract.TimestampLU(&_MockPoolService.CallOpts)
}

// TimestampLU is a free data retrieval call binding the contract method 0x609ae317.
//
// Solidity: function _timestampLU() view returns(uint256)
func (_MockPoolService *MockPoolServiceCallerSession) TimestampLU() (*big.Int, error) {
	return _MockPoolService.Contract.TimestampLU(&_MockPoolService.CallOpts)
}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_MockPoolService *MockPoolServiceCaller) AddressProvider(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockPoolService.contract.Call(opts, &out, "addressProvider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_MockPoolService *MockPoolServiceSession) AddressProvider() (common.Address, error) {
	return _MockPoolService.Contract.AddressProvider(&_MockPoolService.CallOpts)
}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_MockPoolService *MockPoolServiceCallerSession) AddressProvider() (common.Address, error) {
	return _MockPoolService.Contract.AddressProvider(&_MockPoolService.CallOpts)
}

// AvailableLiquidity is a free data retrieval call binding the contract method 0x74375359.
//
// Solidity: function availableLiquidity() view returns(uint256)
func (_MockPoolService *MockPoolServiceCaller) AvailableLiquidity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MockPoolService.contract.Call(opts, &out, "availableLiquidity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AvailableLiquidity is a free data retrieval call binding the contract method 0x74375359.
//
// Solidity: function availableLiquidity() view returns(uint256)
func (_MockPoolService *MockPoolServiceSession) AvailableLiquidity() (*big.Int, error) {
	return _MockPoolService.Contract.AvailableLiquidity(&_MockPoolService.CallOpts)
}

// AvailableLiquidity is a free data retrieval call binding the contract method 0x74375359.
//
// Solidity: function availableLiquidity() view returns(uint256)
func (_MockPoolService *MockPoolServiceCallerSession) AvailableLiquidity() (*big.Int, error) {
	return _MockPoolService.Contract.AvailableLiquidity(&_MockPoolService.CallOpts)
}

// BorrowAPYRAY is a free data retrieval call binding the contract method 0x45d31f9d.
//
// Solidity: function borrowAPY_RAY() view returns(uint256)
func (_MockPoolService *MockPoolServiceCaller) BorrowAPYRAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MockPoolService.contract.Call(opts, &out, "borrowAPY_RAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowAPYRAY is a free data retrieval call binding the contract method 0x45d31f9d.
//
// Solidity: function borrowAPY_RAY() view returns(uint256)
func (_MockPoolService *MockPoolServiceSession) BorrowAPYRAY() (*big.Int, error) {
	return _MockPoolService.Contract.BorrowAPYRAY(&_MockPoolService.CallOpts)
}

// BorrowAPYRAY is a free data retrieval call binding the contract method 0x45d31f9d.
//
// Solidity: function borrowAPY_RAY() view returns(uint256)
func (_MockPoolService *MockPoolServiceCallerSession) BorrowAPYRAY() (*big.Int, error) {
	return _MockPoolService.Contract.BorrowAPYRAY(&_MockPoolService.CallOpts)
}

// CalcCumulativeIndexAtBorrowMore is a free data retrieval call binding the contract method 0xc00495a1.
//
// Solidity: function calcCumulativeIndexAtBorrowMore(uint256 amount, uint256 dAmount, uint256 cumulativeIndexAtOpen) view returns(uint256)
func (_MockPoolService *MockPoolServiceCaller) CalcCumulativeIndexAtBorrowMore(opts *bind.CallOpts, amount *big.Int, dAmount *big.Int, cumulativeIndexAtOpen *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MockPoolService.contract.Call(opts, &out, "calcCumulativeIndexAtBorrowMore", amount, dAmount, cumulativeIndexAtOpen)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcCumulativeIndexAtBorrowMore is a free data retrieval call binding the contract method 0xc00495a1.
//
// Solidity: function calcCumulativeIndexAtBorrowMore(uint256 amount, uint256 dAmount, uint256 cumulativeIndexAtOpen) view returns(uint256)
func (_MockPoolService *MockPoolServiceSession) CalcCumulativeIndexAtBorrowMore(amount *big.Int, dAmount *big.Int, cumulativeIndexAtOpen *big.Int) (*big.Int, error) {
	return _MockPoolService.Contract.CalcCumulativeIndexAtBorrowMore(&_MockPoolService.CallOpts, amount, dAmount, cumulativeIndexAtOpen)
}

// CalcCumulativeIndexAtBorrowMore is a free data retrieval call binding the contract method 0xc00495a1.
//
// Solidity: function calcCumulativeIndexAtBorrowMore(uint256 amount, uint256 dAmount, uint256 cumulativeIndexAtOpen) view returns(uint256)
func (_MockPoolService *MockPoolServiceCallerSession) CalcCumulativeIndexAtBorrowMore(amount *big.Int, dAmount *big.Int, cumulativeIndexAtOpen *big.Int) (*big.Int, error) {
	return _MockPoolService.Contract.CalcCumulativeIndexAtBorrowMore(&_MockPoolService.CallOpts, amount, dAmount, cumulativeIndexAtOpen)
}

// CalcLinearCumulativeRAY is a free data retrieval call binding the contract method 0x0fce70fb.
//
// Solidity: function calcLinearCumulative_RAY() view returns(uint256)
func (_MockPoolService *MockPoolServiceCaller) CalcLinearCumulativeRAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MockPoolService.contract.Call(opts, &out, "calcLinearCumulative_RAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcLinearCumulativeRAY is a free data retrieval call binding the contract method 0x0fce70fb.
//
// Solidity: function calcLinearCumulative_RAY() view returns(uint256)
func (_MockPoolService *MockPoolServiceSession) CalcLinearCumulativeRAY() (*big.Int, error) {
	return _MockPoolService.Contract.CalcLinearCumulativeRAY(&_MockPoolService.CallOpts)
}

// CalcLinearCumulativeRAY is a free data retrieval call binding the contract method 0x0fce70fb.
//
// Solidity: function calcLinearCumulative_RAY() view returns(uint256)
func (_MockPoolService *MockPoolServiceCallerSession) CalcLinearCumulativeRAY() (*big.Int, error) {
	return _MockPoolService.Contract.CalcLinearCumulativeRAY(&_MockPoolService.CallOpts)
}

// CalcLinearIndexRAY is a free data retrieval call binding the contract method 0x532fe502.
//
// Solidity: function calcLinearIndex_RAY() view returns(uint256)
func (_MockPoolService *MockPoolServiceCaller) CalcLinearIndexRAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MockPoolService.contract.Call(opts, &out, "calcLinearIndex_RAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcLinearIndexRAY is a free data retrieval call binding the contract method 0x532fe502.
//
// Solidity: function calcLinearIndex_RAY() view returns(uint256)
func (_MockPoolService *MockPoolServiceSession) CalcLinearIndexRAY() (*big.Int, error) {
	return _MockPoolService.Contract.CalcLinearIndexRAY(&_MockPoolService.CallOpts)
}

// CalcLinearIndexRAY is a free data retrieval call binding the contract method 0x532fe502.
//
// Solidity: function calcLinearIndex_RAY() view returns(uint256)
func (_MockPoolService *MockPoolServiceCallerSession) CalcLinearIndexRAY() (*big.Int, error) {
	return _MockPoolService.Contract.CalcLinearIndexRAY(&_MockPoolService.CallOpts)
}

// CreditManagers is a free data retrieval call binding the contract method 0x1e16e4fc.
//
// Solidity: function creditManagers(uint256 ) view returns(address)
func (_MockPoolService *MockPoolServiceCaller) CreditManagers(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MockPoolService.contract.Call(opts, &out, "creditManagers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditManagers is a free data retrieval call binding the contract method 0x1e16e4fc.
//
// Solidity: function creditManagers(uint256 ) view returns(address)
func (_MockPoolService *MockPoolServiceSession) CreditManagers(arg0 *big.Int) (common.Address, error) {
	return _MockPoolService.Contract.CreditManagers(&_MockPoolService.CallOpts, arg0)
}

// CreditManagers is a free data retrieval call binding the contract method 0x1e16e4fc.
//
// Solidity: function creditManagers(uint256 ) view returns(address)
func (_MockPoolService *MockPoolServiceCallerSession) CreditManagers(arg0 *big.Int) (common.Address, error) {
	return _MockPoolService.Contract.CreditManagers(&_MockPoolService.CallOpts, arg0)
}

// CreditManagersCanBorrow is a free data retrieval call binding the contract method 0x2e97ca21.
//
// Solidity: function creditManagersCanBorrow(address ) view returns(bool)
func (_MockPoolService *MockPoolServiceCaller) CreditManagersCanBorrow(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _MockPoolService.contract.Call(opts, &out, "creditManagersCanBorrow", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CreditManagersCanBorrow is a free data retrieval call binding the contract method 0x2e97ca21.
//
// Solidity: function creditManagersCanBorrow(address ) view returns(bool)
func (_MockPoolService *MockPoolServiceSession) CreditManagersCanBorrow(arg0 common.Address) (bool, error) {
	return _MockPoolService.Contract.CreditManagersCanBorrow(&_MockPoolService.CallOpts, arg0)
}

// CreditManagersCanBorrow is a free data retrieval call binding the contract method 0x2e97ca21.
//
// Solidity: function creditManagersCanBorrow(address ) view returns(bool)
func (_MockPoolService *MockPoolServiceCallerSession) CreditManagersCanBorrow(arg0 common.Address) (bool, error) {
	return _MockPoolService.Contract.CreditManagersCanBorrow(&_MockPoolService.CallOpts, arg0)
}

// CreditManagersCanRepay is a free data retrieval call binding the contract method 0x3e163df0.
//
// Solidity: function creditManagersCanRepay(address ) view returns(bool)
func (_MockPoolService *MockPoolServiceCaller) CreditManagersCanRepay(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _MockPoolService.contract.Call(opts, &out, "creditManagersCanRepay", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CreditManagersCanRepay is a free data retrieval call binding the contract method 0x3e163df0.
//
// Solidity: function creditManagersCanRepay(address ) view returns(bool)
func (_MockPoolService *MockPoolServiceSession) CreditManagersCanRepay(arg0 common.Address) (bool, error) {
	return _MockPoolService.Contract.CreditManagersCanRepay(&_MockPoolService.CallOpts, arg0)
}

// CreditManagersCanRepay is a free data retrieval call binding the contract method 0x3e163df0.
//
// Solidity: function creditManagersCanRepay(address ) view returns(bool)
func (_MockPoolService *MockPoolServiceCallerSession) CreditManagersCanRepay(arg0 common.Address) (bool, error) {
	return _MockPoolService.Contract.CreditManagersCanRepay(&_MockPoolService.CallOpts, arg0)
}

// CreditManagersCount is a free data retrieval call binding the contract method 0xa4e8273e.
//
// Solidity: function creditManagersCount() pure returns(uint256)
func (_MockPoolService *MockPoolServiceCaller) CreditManagersCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MockPoolService.contract.Call(opts, &out, "creditManagersCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CreditManagersCount is a free data retrieval call binding the contract method 0xa4e8273e.
//
// Solidity: function creditManagersCount() pure returns(uint256)
func (_MockPoolService *MockPoolServiceSession) CreditManagersCount() (*big.Int, error) {
	return _MockPoolService.Contract.CreditManagersCount(&_MockPoolService.CallOpts)
}

// CreditManagersCount is a free data retrieval call binding the contract method 0xa4e8273e.
//
// Solidity: function creditManagersCount() pure returns(uint256)
func (_MockPoolService *MockPoolServiceCallerSession) CreditManagersCount() (*big.Int, error) {
	return _MockPoolService.Contract.CreditManagersCount(&_MockPoolService.CallOpts)
}

// DieselToken is a free data retrieval call binding the contract method 0x36dda7d5.
//
// Solidity: function dieselToken() view returns(address)
func (_MockPoolService *MockPoolServiceCaller) DieselToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockPoolService.contract.Call(opts, &out, "dieselToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DieselToken is a free data retrieval call binding the contract method 0x36dda7d5.
//
// Solidity: function dieselToken() view returns(address)
func (_MockPoolService *MockPoolServiceSession) DieselToken() (common.Address, error) {
	return _MockPoolService.Contract.DieselToken(&_MockPoolService.CallOpts)
}

// DieselToken is a free data retrieval call binding the contract method 0x36dda7d5.
//
// Solidity: function dieselToken() view returns(address)
func (_MockPoolService *MockPoolServiceCallerSession) DieselToken() (common.Address, error) {
	return _MockPoolService.Contract.DieselToken(&_MockPoolService.CallOpts)
}

// ExpectedLiquidity is a free data retrieval call binding the contract method 0xfe14112d.
//
// Solidity: function expectedLiquidity() pure returns(uint256)
func (_MockPoolService *MockPoolServiceCaller) ExpectedLiquidity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MockPoolService.contract.Call(opts, &out, "expectedLiquidity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExpectedLiquidity is a free data retrieval call binding the contract method 0xfe14112d.
//
// Solidity: function expectedLiquidity() pure returns(uint256)
func (_MockPoolService *MockPoolServiceSession) ExpectedLiquidity() (*big.Int, error) {
	return _MockPoolService.Contract.ExpectedLiquidity(&_MockPoolService.CallOpts)
}

// ExpectedLiquidity is a free data retrieval call binding the contract method 0xfe14112d.
//
// Solidity: function expectedLiquidity() pure returns(uint256)
func (_MockPoolService *MockPoolServiceCallerSession) ExpectedLiquidity() (*big.Int, error) {
	return _MockPoolService.Contract.ExpectedLiquidity(&_MockPoolService.CallOpts)
}

// ExpectedLiquidityLimit is a free data retrieval call binding the contract method 0xef8d9603.
//
// Solidity: function expectedLiquidityLimit() view returns(uint256)
func (_MockPoolService *MockPoolServiceCaller) ExpectedLiquidityLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MockPoolService.contract.Call(opts, &out, "expectedLiquidityLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExpectedLiquidityLimit is a free data retrieval call binding the contract method 0xef8d9603.
//
// Solidity: function expectedLiquidityLimit() view returns(uint256)
func (_MockPoolService *MockPoolServiceSession) ExpectedLiquidityLimit() (*big.Int, error) {
	return _MockPoolService.Contract.ExpectedLiquidityLimit(&_MockPoolService.CallOpts)
}

// ExpectedLiquidityLimit is a free data retrieval call binding the contract method 0xef8d9603.
//
// Solidity: function expectedLiquidityLimit() view returns(uint256)
func (_MockPoolService *MockPoolServiceCallerSession) ExpectedLiquidityLimit() (*big.Int, error) {
	return _MockPoolService.Contract.ExpectedLiquidityLimit(&_MockPoolService.CallOpts)
}

// FromDiesel is a free data retrieval call binding the contract method 0x5427c938.
//
// Solidity: function fromDiesel(uint256 amount) view returns(uint256)
func (_MockPoolService *MockPoolServiceCaller) FromDiesel(opts *bind.CallOpts, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MockPoolService.contract.Call(opts, &out, "fromDiesel", amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FromDiesel is a free data retrieval call binding the contract method 0x5427c938.
//
// Solidity: function fromDiesel(uint256 amount) view returns(uint256)
func (_MockPoolService *MockPoolServiceSession) FromDiesel(amount *big.Int) (*big.Int, error) {
	return _MockPoolService.Contract.FromDiesel(&_MockPoolService.CallOpts, amount)
}

// FromDiesel is a free data retrieval call binding the contract method 0x5427c938.
//
// Solidity: function fromDiesel(uint256 amount) view returns(uint256)
func (_MockPoolService *MockPoolServiceCallerSession) FromDiesel(amount *big.Int) (*big.Int, error) {
	return _MockPoolService.Contract.FromDiesel(&_MockPoolService.CallOpts, amount)
}

// GetDieselRateRAY is a free data retrieval call binding the contract method 0x788c6bfe.
//
// Solidity: function getDieselRate_RAY() pure returns(uint256)
func (_MockPoolService *MockPoolServiceCaller) GetDieselRateRAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MockPoolService.contract.Call(opts, &out, "getDieselRate_RAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDieselRateRAY is a free data retrieval call binding the contract method 0x788c6bfe.
//
// Solidity: function getDieselRate_RAY() pure returns(uint256)
func (_MockPoolService *MockPoolServiceSession) GetDieselRateRAY() (*big.Int, error) {
	return _MockPoolService.Contract.GetDieselRateRAY(&_MockPoolService.CallOpts)
}

// GetDieselRateRAY is a free data retrieval call binding the contract method 0x788c6bfe.
//
// Solidity: function getDieselRate_RAY() pure returns(uint256)
func (_MockPoolService *MockPoolServiceCallerSession) GetDieselRateRAY() (*big.Int, error) {
	return _MockPoolService.Contract.GetDieselRateRAY(&_MockPoolService.CallOpts)
}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_MockPoolService *MockPoolServiceCaller) InterestRateModel(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockPoolService.contract.Call(opts, &out, "interestRateModel")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_MockPoolService *MockPoolServiceSession) InterestRateModel() (common.Address, error) {
	return _MockPoolService.Contract.InterestRateModel(&_MockPoolService.CallOpts)
}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_MockPoolService *MockPoolServiceCallerSession) InterestRateModel() (common.Address, error) {
	return _MockPoolService.Contract.InterestRateModel(&_MockPoolService.CallOpts)
}

// LendAccount is a free data retrieval call binding the contract method 0x2a3354c9.
//
// Solidity: function lendAccount() view returns(address)
func (_MockPoolService *MockPoolServiceCaller) LendAccount(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockPoolService.contract.Call(opts, &out, "lendAccount")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LendAccount is a free data retrieval call binding the contract method 0x2a3354c9.
//
// Solidity: function lendAccount() view returns(address)
func (_MockPoolService *MockPoolServiceSession) LendAccount() (common.Address, error) {
	return _MockPoolService.Contract.LendAccount(&_MockPoolService.CallOpts)
}

// LendAccount is a free data retrieval call binding the contract method 0x2a3354c9.
//
// Solidity: function lendAccount() view returns(address)
func (_MockPoolService *MockPoolServiceCallerSession) LendAccount() (common.Address, error) {
	return _MockPoolService.Contract.LendAccount(&_MockPoolService.CallOpts)
}

// LendAmount is a free data retrieval call binding the contract method 0x29f3d3b6.
//
// Solidity: function lendAmount() view returns(uint256)
func (_MockPoolService *MockPoolServiceCaller) LendAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MockPoolService.contract.Call(opts, &out, "lendAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LendAmount is a free data retrieval call binding the contract method 0x29f3d3b6.
//
// Solidity: function lendAmount() view returns(uint256)
func (_MockPoolService *MockPoolServiceSession) LendAmount() (*big.Int, error) {
	return _MockPoolService.Contract.LendAmount(&_MockPoolService.CallOpts)
}

// LendAmount is a free data retrieval call binding the contract method 0x29f3d3b6.
//
// Solidity: function lendAmount() view returns(uint256)
func (_MockPoolService *MockPoolServiceCallerSession) LendAmount() (*big.Int, error) {
	return _MockPoolService.Contract.LendAmount(&_MockPoolService.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MockPoolService *MockPoolServiceCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MockPoolService.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MockPoolService *MockPoolServiceSession) Paused() (bool, error) {
	return _MockPoolService.Contract.Paused(&_MockPoolService.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MockPoolService *MockPoolServiceCallerSession) Paused() (bool, error) {
	return _MockPoolService.Contract.Paused(&_MockPoolService.CallOpts)
}

// RepayAmount is a free data retrieval call binding the contract method 0xd0efe753.
//
// Solidity: function repayAmount() view returns(uint256)
func (_MockPoolService *MockPoolServiceCaller) RepayAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MockPoolService.contract.Call(opts, &out, "repayAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RepayAmount is a free data retrieval call binding the contract method 0xd0efe753.
//
// Solidity: function repayAmount() view returns(uint256)
func (_MockPoolService *MockPoolServiceSession) RepayAmount() (*big.Int, error) {
	return _MockPoolService.Contract.RepayAmount(&_MockPoolService.CallOpts)
}

// RepayAmount is a free data retrieval call binding the contract method 0xd0efe753.
//
// Solidity: function repayAmount() view returns(uint256)
func (_MockPoolService *MockPoolServiceCallerSession) RepayAmount() (*big.Int, error) {
	return _MockPoolService.Contract.RepayAmount(&_MockPoolService.CallOpts)
}

// RepayLoss is a free data retrieval call binding the contract method 0x67c99d58.
//
// Solidity: function repayLoss() view returns(uint256)
func (_MockPoolService *MockPoolServiceCaller) RepayLoss(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MockPoolService.contract.Call(opts, &out, "repayLoss")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RepayLoss is a free data retrieval call binding the contract method 0x67c99d58.
//
// Solidity: function repayLoss() view returns(uint256)
func (_MockPoolService *MockPoolServiceSession) RepayLoss() (*big.Int, error) {
	return _MockPoolService.Contract.RepayLoss(&_MockPoolService.CallOpts)
}

// RepayLoss is a free data retrieval call binding the contract method 0x67c99d58.
//
// Solidity: function repayLoss() view returns(uint256)
func (_MockPoolService *MockPoolServiceCallerSession) RepayLoss() (*big.Int, error) {
	return _MockPoolService.Contract.RepayLoss(&_MockPoolService.CallOpts)
}

// RepayProfit is a free data retrieval call binding the contract method 0xcb3905e1.
//
// Solidity: function repayProfit() view returns(uint256)
func (_MockPoolService *MockPoolServiceCaller) RepayProfit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MockPoolService.contract.Call(opts, &out, "repayProfit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RepayProfit is a free data retrieval call binding the contract method 0xcb3905e1.
//
// Solidity: function repayProfit() view returns(uint256)
func (_MockPoolService *MockPoolServiceSession) RepayProfit() (*big.Int, error) {
	return _MockPoolService.Contract.RepayProfit(&_MockPoolService.CallOpts)
}

// RepayProfit is a free data retrieval call binding the contract method 0xcb3905e1.
//
// Solidity: function repayProfit() view returns(uint256)
func (_MockPoolService *MockPoolServiceCallerSession) RepayProfit() (*big.Int, error) {
	return _MockPoolService.Contract.RepayProfit(&_MockPoolService.CallOpts)
}

// ToDiesel is a free data retrieval call binding the contract method 0x4d778ad1.
//
// Solidity: function toDiesel(uint256 amount) view returns(uint256)
func (_MockPoolService *MockPoolServiceCaller) ToDiesel(opts *bind.CallOpts, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MockPoolService.contract.Call(opts, &out, "toDiesel", amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ToDiesel is a free data retrieval call binding the contract method 0x4d778ad1.
//
// Solidity: function toDiesel(uint256 amount) view returns(uint256)
func (_MockPoolService *MockPoolServiceSession) ToDiesel(amount *big.Int) (*big.Int, error) {
	return _MockPoolService.Contract.ToDiesel(&_MockPoolService.CallOpts, amount)
}

// ToDiesel is a free data retrieval call binding the contract method 0x4d778ad1.
//
// Solidity: function toDiesel(uint256 amount) view returns(uint256)
func (_MockPoolService *MockPoolServiceCallerSession) ToDiesel(amount *big.Int) (*big.Int, error) {
	return _MockPoolService.Contract.ToDiesel(&_MockPoolService.CallOpts, amount)
}

// TotalBorrowed is a free data retrieval call binding the contract method 0x4c19386c.
//
// Solidity: function totalBorrowed() view returns(uint256)
func (_MockPoolService *MockPoolServiceCaller) TotalBorrowed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MockPoolService.contract.Call(opts, &out, "totalBorrowed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBorrowed is a free data retrieval call binding the contract method 0x4c19386c.
//
// Solidity: function totalBorrowed() view returns(uint256)
func (_MockPoolService *MockPoolServiceSession) TotalBorrowed() (*big.Int, error) {
	return _MockPoolService.Contract.TotalBorrowed(&_MockPoolService.CallOpts)
}

// TotalBorrowed is a free data retrieval call binding the contract method 0x4c19386c.
//
// Solidity: function totalBorrowed() view returns(uint256)
func (_MockPoolService *MockPoolServiceCallerSession) TotalBorrowed() (*big.Int, error) {
	return _MockPoolService.Contract.TotalBorrowed(&_MockPoolService.CallOpts)
}

// TreasuryAddress is a free data retrieval call binding the contract method 0xc5f956af.
//
// Solidity: function treasuryAddress() view returns(address)
func (_MockPoolService *MockPoolServiceCaller) TreasuryAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockPoolService.contract.Call(opts, &out, "treasuryAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TreasuryAddress is a free data retrieval call binding the contract method 0xc5f956af.
//
// Solidity: function treasuryAddress() view returns(address)
func (_MockPoolService *MockPoolServiceSession) TreasuryAddress() (common.Address, error) {
	return _MockPoolService.Contract.TreasuryAddress(&_MockPoolService.CallOpts)
}

// TreasuryAddress is a free data retrieval call binding the contract method 0xc5f956af.
//
// Solidity: function treasuryAddress() view returns(address)
func (_MockPoolService *MockPoolServiceCallerSession) TreasuryAddress() (common.Address, error) {
	return _MockPoolService.Contract.TreasuryAddress(&_MockPoolService.CallOpts)
}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_MockPoolService *MockPoolServiceCaller) UnderlyingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockPoolService.contract.Call(opts, &out, "underlyingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_MockPoolService *MockPoolServiceSession) UnderlyingToken() (common.Address, error) {
	return _MockPoolService.Contract.UnderlyingToken(&_MockPoolService.CallOpts)
}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_MockPoolService *MockPoolServiceCallerSession) UnderlyingToken() (common.Address, error) {
	return _MockPoolService.Contract.UnderlyingToken(&_MockPoolService.CallOpts)
}

// WithdrawFee is a free data retrieval call binding the contract method 0xe941fa78.
//
// Solidity: function withdrawFee() view returns(uint256)
func (_MockPoolService *MockPoolServiceCaller) WithdrawFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MockPoolService.contract.Call(opts, &out, "withdrawFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawFee is a free data retrieval call binding the contract method 0xe941fa78.
//
// Solidity: function withdrawFee() view returns(uint256)
func (_MockPoolService *MockPoolServiceSession) WithdrawFee() (*big.Int, error) {
	return _MockPoolService.Contract.WithdrawFee(&_MockPoolService.CallOpts)
}

// WithdrawFee is a free data retrieval call binding the contract method 0xe941fa78.
//
// Solidity: function withdrawFee() view returns(uint256)
func (_MockPoolService *MockPoolServiceCallerSession) WithdrawFee() (*big.Int, error) {
	return _MockPoolService.Contract.WithdrawFee(&_MockPoolService.CallOpts)
}

// WithdrawMultiplier is a free data retrieval call binding the contract method 0xb3554a0a.
//
// Solidity: function withdrawMultiplier() view returns(uint256)
func (_MockPoolService *MockPoolServiceCaller) WithdrawMultiplier(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MockPoolService.contract.Call(opts, &out, "withdrawMultiplier")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawMultiplier is a free data retrieval call binding the contract method 0xb3554a0a.
//
// Solidity: function withdrawMultiplier() view returns(uint256)
func (_MockPoolService *MockPoolServiceSession) WithdrawMultiplier() (*big.Int, error) {
	return _MockPoolService.Contract.WithdrawMultiplier(&_MockPoolService.CallOpts)
}

// WithdrawMultiplier is a free data retrieval call binding the contract method 0xb3554a0a.
//
// Solidity: function withdrawMultiplier() view returns(uint256)
func (_MockPoolService *MockPoolServiceCallerSession) WithdrawMultiplier() (*big.Int, error) {
	return _MockPoolService.Contract.WithdrawMultiplier(&_MockPoolService.CallOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x9aa5d462.
//
// Solidity: function addLiquidity(uint256 amount, address onBehalfOf, uint256 referralCode) returns()
func (_MockPoolService *MockPoolServiceTransactor) AddLiquidity(opts *bind.TransactOpts, amount *big.Int, onBehalfOf common.Address, referralCode *big.Int) (*types.Transaction, error) {
	return _MockPoolService.contract.Transact(opts, "addLiquidity", amount, onBehalfOf, referralCode)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x9aa5d462.
//
// Solidity: function addLiquidity(uint256 amount, address onBehalfOf, uint256 referralCode) returns()
func (_MockPoolService *MockPoolServiceSession) AddLiquidity(amount *big.Int, onBehalfOf common.Address, referralCode *big.Int) (*types.Transaction, error) {
	return _MockPoolService.Contract.AddLiquidity(&_MockPoolService.TransactOpts, amount, onBehalfOf, referralCode)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x9aa5d462.
//
// Solidity: function addLiquidity(uint256 amount, address onBehalfOf, uint256 referralCode) returns()
func (_MockPoolService *MockPoolServiceTransactorSession) AddLiquidity(amount *big.Int, onBehalfOf common.Address, referralCode *big.Int) (*types.Transaction, error) {
	return _MockPoolService.Contract.AddLiquidity(&_MockPoolService.TransactOpts, amount, onBehalfOf, referralCode)
}

// ConnectCreditManager is a paid mutator transaction binding the contract method 0xcf33d955.
//
// Solidity: function connectCreditManager(address _creditManager) returns()
func (_MockPoolService *MockPoolServiceTransactor) ConnectCreditManager(opts *bind.TransactOpts, _creditManager common.Address) (*types.Transaction, error) {
	return _MockPoolService.contract.Transact(opts, "connectCreditManager", _creditManager)
}

// ConnectCreditManager is a paid mutator transaction binding the contract method 0xcf33d955.
//
// Solidity: function connectCreditManager(address _creditManager) returns()
func (_MockPoolService *MockPoolServiceSession) ConnectCreditManager(_creditManager common.Address) (*types.Transaction, error) {
	return _MockPoolService.Contract.ConnectCreditManager(&_MockPoolService.TransactOpts, _creditManager)
}

// ConnectCreditManager is a paid mutator transaction binding the contract method 0xcf33d955.
//
// Solidity: function connectCreditManager(address _creditManager) returns()
func (_MockPoolService *MockPoolServiceTransactorSession) ConnectCreditManager(_creditManager common.Address) (*types.Transaction, error) {
	return _MockPoolService.Contract.ConnectCreditManager(&_MockPoolService.TransactOpts, _creditManager)
}

// ForbidCreditManagerToBorrow is a paid mutator transaction binding the contract method 0x078c4781.
//
// Solidity: function forbidCreditManagerToBorrow(address _creditManager) returns()
func (_MockPoolService *MockPoolServiceTransactor) ForbidCreditManagerToBorrow(opts *bind.TransactOpts, _creditManager common.Address) (*types.Transaction, error) {
	return _MockPoolService.contract.Transact(opts, "forbidCreditManagerToBorrow", _creditManager)
}

// ForbidCreditManagerToBorrow is a paid mutator transaction binding the contract method 0x078c4781.
//
// Solidity: function forbidCreditManagerToBorrow(address _creditManager) returns()
func (_MockPoolService *MockPoolServiceSession) ForbidCreditManagerToBorrow(_creditManager common.Address) (*types.Transaction, error) {
	return _MockPoolService.Contract.ForbidCreditManagerToBorrow(&_MockPoolService.TransactOpts, _creditManager)
}

// ForbidCreditManagerToBorrow is a paid mutator transaction binding the contract method 0x078c4781.
//
// Solidity: function forbidCreditManagerToBorrow(address _creditManager) returns()
func (_MockPoolService *MockPoolServiceTransactorSession) ForbidCreditManagerToBorrow(_creditManager common.Address) (*types.Transaction, error) {
	return _MockPoolService.Contract.ForbidCreditManagerToBorrow(&_MockPoolService.TransactOpts, _creditManager)
}

// LendCreditAccount is a paid mutator transaction binding the contract method 0xbf28068b.
//
// Solidity: function lendCreditAccount(uint256 borrowedAmount, address creditAccount) returns()
func (_MockPoolService *MockPoolServiceTransactor) LendCreditAccount(opts *bind.TransactOpts, borrowedAmount *big.Int, creditAccount common.Address) (*types.Transaction, error) {
	return _MockPoolService.contract.Transact(opts, "lendCreditAccount", borrowedAmount, creditAccount)
}

// LendCreditAccount is a paid mutator transaction binding the contract method 0xbf28068b.
//
// Solidity: function lendCreditAccount(uint256 borrowedAmount, address creditAccount) returns()
func (_MockPoolService *MockPoolServiceSession) LendCreditAccount(borrowedAmount *big.Int, creditAccount common.Address) (*types.Transaction, error) {
	return _MockPoolService.Contract.LendCreditAccount(&_MockPoolService.TransactOpts, borrowedAmount, creditAccount)
}

// LendCreditAccount is a paid mutator transaction binding the contract method 0xbf28068b.
//
// Solidity: function lendCreditAccount(uint256 borrowedAmount, address creditAccount) returns()
func (_MockPoolService *MockPoolServiceTransactorSession) LendCreditAccount(borrowedAmount *big.Int, creditAccount common.Address) (*types.Transaction, error) {
	return _MockPoolService.Contract.LendCreditAccount(&_MockPoolService.TransactOpts, borrowedAmount, creditAccount)
}

// NewInterestRateModel is a paid mutator transaction binding the contract method 0xf11a6487.
//
// Solidity: function newInterestRateModel(address _interestRateModel) returns()
func (_MockPoolService *MockPoolServiceTransactor) NewInterestRateModel(opts *bind.TransactOpts, _interestRateModel common.Address) (*types.Transaction, error) {
	return _MockPoolService.contract.Transact(opts, "newInterestRateModel", _interestRateModel)
}

// NewInterestRateModel is a paid mutator transaction binding the contract method 0xf11a6487.
//
// Solidity: function newInterestRateModel(address _interestRateModel) returns()
func (_MockPoolService *MockPoolServiceSession) NewInterestRateModel(_interestRateModel common.Address) (*types.Transaction, error) {
	return _MockPoolService.Contract.NewInterestRateModel(&_MockPoolService.TransactOpts, _interestRateModel)
}

// NewInterestRateModel is a paid mutator transaction binding the contract method 0xf11a6487.
//
// Solidity: function newInterestRateModel(address _interestRateModel) returns()
func (_MockPoolService *MockPoolServiceTransactorSession) NewInterestRateModel(_interestRateModel common.Address) (*types.Transaction, error) {
	return _MockPoolService.Contract.NewInterestRateModel(&_MockPoolService.TransactOpts, _interestRateModel)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MockPoolService *MockPoolServiceTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockPoolService.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MockPoolService *MockPoolServiceSession) Pause() (*types.Transaction, error) {
	return _MockPoolService.Contract.Pause(&_MockPoolService.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MockPoolService *MockPoolServiceTransactorSession) Pause() (*types.Transaction, error) {
	return _MockPoolService.Contract.Pause(&_MockPoolService.TransactOpts)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x05fe138b.
//
// Solidity: function removeLiquidity(uint256 amount, address to) returns(uint256)
func (_MockPoolService *MockPoolServiceTransactor) RemoveLiquidity(opts *bind.TransactOpts, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _MockPoolService.contract.Transact(opts, "removeLiquidity", amount, to)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x05fe138b.
//
// Solidity: function removeLiquidity(uint256 amount, address to) returns(uint256)
func (_MockPoolService *MockPoolServiceSession) RemoveLiquidity(amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _MockPoolService.Contract.RemoveLiquidity(&_MockPoolService.TransactOpts, amount, to)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x05fe138b.
//
// Solidity: function removeLiquidity(uint256 amount, address to) returns(uint256)
func (_MockPoolService *MockPoolServiceTransactorSession) RemoveLiquidity(amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _MockPoolService.Contract.RemoveLiquidity(&_MockPoolService.TransactOpts, amount, to)
}

// RepayCreditAccount is a paid mutator transaction binding the contract method 0xca9505e4.
//
// Solidity: function repayCreditAccount(uint256 borrowedAmount, uint256 profit, uint256 loss) returns()
func (_MockPoolService *MockPoolServiceTransactor) RepayCreditAccount(opts *bind.TransactOpts, borrowedAmount *big.Int, profit *big.Int, loss *big.Int) (*types.Transaction, error) {
	return _MockPoolService.contract.Transact(opts, "repayCreditAccount", borrowedAmount, profit, loss)
}

// RepayCreditAccount is a paid mutator transaction binding the contract method 0xca9505e4.
//
// Solidity: function repayCreditAccount(uint256 borrowedAmount, uint256 profit, uint256 loss) returns()
func (_MockPoolService *MockPoolServiceSession) RepayCreditAccount(borrowedAmount *big.Int, profit *big.Int, loss *big.Int) (*types.Transaction, error) {
	return _MockPoolService.Contract.RepayCreditAccount(&_MockPoolService.TransactOpts, borrowedAmount, profit, loss)
}

// RepayCreditAccount is a paid mutator transaction binding the contract method 0xca9505e4.
//
// Solidity: function repayCreditAccount(uint256 borrowedAmount, uint256 profit, uint256 loss) returns()
func (_MockPoolService *MockPoolServiceTransactorSession) RepayCreditAccount(borrowedAmount *big.Int, profit *big.Int, loss *big.Int) (*types.Transaction, error) {
	return _MockPoolService.Contract.RepayCreditAccount(&_MockPoolService.TransactOpts, borrowedAmount, profit, loss)
}

// SetCumulativeRAY is a paid mutator transaction binding the contract method 0x96df5dc0.
//
// Solidity: function setCumulative_RAY(uint256 _cumulativeIndex_RAY) returns()
func (_MockPoolService *MockPoolServiceTransactor) SetCumulativeRAY(opts *bind.TransactOpts, _cumulativeIndex_RAY *big.Int) (*types.Transaction, error) {
	return _MockPoolService.contract.Transact(opts, "setCumulative_RAY", _cumulativeIndex_RAY)
}

// SetCumulativeRAY is a paid mutator transaction binding the contract method 0x96df5dc0.
//
// Solidity: function setCumulative_RAY(uint256 _cumulativeIndex_RAY) returns()
func (_MockPoolService *MockPoolServiceSession) SetCumulativeRAY(_cumulativeIndex_RAY *big.Int) (*types.Transaction, error) {
	return _MockPoolService.Contract.SetCumulativeRAY(&_MockPoolService.TransactOpts, _cumulativeIndex_RAY)
}

// SetCumulativeRAY is a paid mutator transaction binding the contract method 0x96df5dc0.
//
// Solidity: function setCumulative_RAY(uint256 _cumulativeIndex_RAY) returns()
func (_MockPoolService *MockPoolServiceTransactorSession) SetCumulativeRAY(_cumulativeIndex_RAY *big.Int) (*types.Transaction, error) {
	return _MockPoolService.Contract.SetCumulativeRAY(&_MockPoolService.TransactOpts, _cumulativeIndex_RAY)
}

// SetExpectedLiquidityLimit is a paid mutator transaction binding the contract method 0xbb04b193.
//
// Solidity: function setExpectedLiquidityLimit(uint256 num) returns()
func (_MockPoolService *MockPoolServiceTransactor) SetExpectedLiquidityLimit(opts *bind.TransactOpts, num *big.Int) (*types.Transaction, error) {
	return _MockPoolService.contract.Transact(opts, "setExpectedLiquidityLimit", num)
}

// SetExpectedLiquidityLimit is a paid mutator transaction binding the contract method 0xbb04b193.
//
// Solidity: function setExpectedLiquidityLimit(uint256 num) returns()
func (_MockPoolService *MockPoolServiceSession) SetExpectedLiquidityLimit(num *big.Int) (*types.Transaction, error) {
	return _MockPoolService.Contract.SetExpectedLiquidityLimit(&_MockPoolService.TransactOpts, num)
}

// SetExpectedLiquidityLimit is a paid mutator transaction binding the contract method 0xbb04b193.
//
// Solidity: function setExpectedLiquidityLimit(uint256 num) returns()
func (_MockPoolService *MockPoolServiceTransactorSession) SetExpectedLiquidityLimit(num *big.Int) (*types.Transaction, error) {
	return _MockPoolService.Contract.SetExpectedLiquidityLimit(&_MockPoolService.TransactOpts, num)
}

// SetWithdrawFee is a paid mutator transaction binding the contract method 0xb6ac642a.
//
// Solidity: function setWithdrawFee(uint256 num) returns()
func (_MockPoolService *MockPoolServiceTransactor) SetWithdrawFee(opts *bind.TransactOpts, num *big.Int) (*types.Transaction, error) {
	return _MockPoolService.contract.Transact(opts, "setWithdrawFee", num)
}

// SetWithdrawFee is a paid mutator transaction binding the contract method 0xb6ac642a.
//
// Solidity: function setWithdrawFee(uint256 num) returns()
func (_MockPoolService *MockPoolServiceSession) SetWithdrawFee(num *big.Int) (*types.Transaction, error) {
	return _MockPoolService.Contract.SetWithdrawFee(&_MockPoolService.TransactOpts, num)
}

// SetWithdrawFee is a paid mutator transaction binding the contract method 0xb6ac642a.
//
// Solidity: function setWithdrawFee(uint256 num) returns()
func (_MockPoolService *MockPoolServiceTransactorSession) SetWithdrawFee(num *big.Int) (*types.Transaction, error) {
	return _MockPoolService.Contract.SetWithdrawFee(&_MockPoolService.TransactOpts, num)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MockPoolService *MockPoolServiceTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockPoolService.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MockPoolService *MockPoolServiceSession) Unpause() (*types.Transaction, error) {
	return _MockPoolService.Contract.Unpause(&_MockPoolService.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MockPoolService *MockPoolServiceTransactorSession) Unpause() (*types.Transaction, error) {
	return _MockPoolService.Contract.Unpause(&_MockPoolService.TransactOpts)
}

// MockPoolServiceAddLiquidityIterator is returned from FilterAddLiquidity and is used to iterate over the raw logs and unpacked data for AddLiquidity events raised by the MockPoolService contract.
type MockPoolServiceAddLiquidityIterator struct {
	Event *MockPoolServiceAddLiquidity // Event containing the contract specifics and raw log

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
func (it *MockPoolServiceAddLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockPoolServiceAddLiquidity)
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
		it.Event = new(MockPoolServiceAddLiquidity)
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
func (it *MockPoolServiceAddLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockPoolServiceAddLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockPoolServiceAddLiquidity represents a AddLiquidity event raised by the MockPoolService contract.
type MockPoolServiceAddLiquidity struct {
	Sender       common.Address
	OnBehalfOf   common.Address
	Amount       *big.Int
	ReferralCode *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAddLiquidity is a free log retrieval operation binding the contract event 0xd2491a9b4fe81a7cd4511e8b7b7743951b061dad5bed7da8a7795b080ee08c7e.
//
// Solidity: event AddLiquidity(address indexed sender, address indexed onBehalfOf, uint256 amount, uint256 referralCode)
func (_MockPoolService *MockPoolServiceFilterer) FilterAddLiquidity(opts *bind.FilterOpts, sender []common.Address, onBehalfOf []common.Address) (*MockPoolServiceAddLiquidityIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	logs, sub, err := _MockPoolService.contract.FilterLogs(opts, "AddLiquidity", senderRule, onBehalfOfRule)
	if err != nil {
		return nil, err
	}
	return &MockPoolServiceAddLiquidityIterator{contract: _MockPoolService.contract, event: "AddLiquidity", logs: logs, sub: sub}, nil
}

// WatchAddLiquidity is a free log subscription operation binding the contract event 0xd2491a9b4fe81a7cd4511e8b7b7743951b061dad5bed7da8a7795b080ee08c7e.
//
// Solidity: event AddLiquidity(address indexed sender, address indexed onBehalfOf, uint256 amount, uint256 referralCode)
func (_MockPoolService *MockPoolServiceFilterer) WatchAddLiquidity(opts *bind.WatchOpts, sink chan<- *MockPoolServiceAddLiquidity, sender []common.Address, onBehalfOf []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	logs, sub, err := _MockPoolService.contract.WatchLogs(opts, "AddLiquidity", senderRule, onBehalfOfRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockPoolServiceAddLiquidity)
				if err := _MockPoolService.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
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
func (_MockPoolService *MockPoolServiceFilterer) ParseAddLiquidity(log types.Log) (*MockPoolServiceAddLiquidity, error) {
	event := new(MockPoolServiceAddLiquidity)
	if err := _MockPoolService.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockPoolServiceBorrowIterator is returned from FilterBorrow and is used to iterate over the raw logs and unpacked data for Borrow events raised by the MockPoolService contract.
type MockPoolServiceBorrowIterator struct {
	Event *MockPoolServiceBorrow // Event containing the contract specifics and raw log

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
func (it *MockPoolServiceBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockPoolServiceBorrow)
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
		it.Event = new(MockPoolServiceBorrow)
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
func (it *MockPoolServiceBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockPoolServiceBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockPoolServiceBorrow represents a Borrow event raised by the MockPoolService contract.
type MockPoolServiceBorrow struct {
	CreditManager common.Address
	CreditAccount common.Address
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBorrow is a free log retrieval operation binding the contract event 0x312a5e5e1079f5dda4e95dbbd0b908b291fd5b992ef22073643ab691572c5b52.
//
// Solidity: event Borrow(address indexed creditManager, address indexed creditAccount, uint256 amount)
func (_MockPoolService *MockPoolServiceFilterer) FilterBorrow(opts *bind.FilterOpts, creditManager []common.Address, creditAccount []common.Address) (*MockPoolServiceBorrowIterator, error) {

	var creditManagerRule []interface{}
	for _, creditManagerItem := range creditManager {
		creditManagerRule = append(creditManagerRule, creditManagerItem)
	}
	var creditAccountRule []interface{}
	for _, creditAccountItem := range creditAccount {
		creditAccountRule = append(creditAccountRule, creditAccountItem)
	}

	logs, sub, err := _MockPoolService.contract.FilterLogs(opts, "Borrow", creditManagerRule, creditAccountRule)
	if err != nil {
		return nil, err
	}
	return &MockPoolServiceBorrowIterator{contract: _MockPoolService.contract, event: "Borrow", logs: logs, sub: sub}, nil
}

// WatchBorrow is a free log subscription operation binding the contract event 0x312a5e5e1079f5dda4e95dbbd0b908b291fd5b992ef22073643ab691572c5b52.
//
// Solidity: event Borrow(address indexed creditManager, address indexed creditAccount, uint256 amount)
func (_MockPoolService *MockPoolServiceFilterer) WatchBorrow(opts *bind.WatchOpts, sink chan<- *MockPoolServiceBorrow, creditManager []common.Address, creditAccount []common.Address) (event.Subscription, error) {

	var creditManagerRule []interface{}
	for _, creditManagerItem := range creditManager {
		creditManagerRule = append(creditManagerRule, creditManagerItem)
	}
	var creditAccountRule []interface{}
	for _, creditAccountItem := range creditAccount {
		creditAccountRule = append(creditAccountRule, creditAccountItem)
	}

	logs, sub, err := _MockPoolService.contract.WatchLogs(opts, "Borrow", creditManagerRule, creditAccountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockPoolServiceBorrow)
				if err := _MockPoolService.contract.UnpackLog(event, "Borrow", log); err != nil {
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
func (_MockPoolService *MockPoolServiceFilterer) ParseBorrow(log types.Log) (*MockPoolServiceBorrow, error) {
	event := new(MockPoolServiceBorrow)
	if err := _MockPoolService.contract.UnpackLog(event, "Borrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockPoolServiceBorrowForbiddenIterator is returned from FilterBorrowForbidden and is used to iterate over the raw logs and unpacked data for BorrowForbidden events raised by the MockPoolService contract.
type MockPoolServiceBorrowForbiddenIterator struct {
	Event *MockPoolServiceBorrowForbidden // Event containing the contract specifics and raw log

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
func (it *MockPoolServiceBorrowForbiddenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockPoolServiceBorrowForbidden)
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
		it.Event = new(MockPoolServiceBorrowForbidden)
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
func (it *MockPoolServiceBorrowForbiddenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockPoolServiceBorrowForbiddenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockPoolServiceBorrowForbidden represents a BorrowForbidden event raised by the MockPoolService contract.
type MockPoolServiceBorrowForbidden struct {
	CreditManager common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBorrowForbidden is a free log retrieval operation binding the contract event 0x9181736fce85d2d4cca2e4406f10679302ae5c387180fdb62963af3cd9a24fd6.
//
// Solidity: event BorrowForbidden(address indexed creditManager)
func (_MockPoolService *MockPoolServiceFilterer) FilterBorrowForbidden(opts *bind.FilterOpts, creditManager []common.Address) (*MockPoolServiceBorrowForbiddenIterator, error) {

	var creditManagerRule []interface{}
	for _, creditManagerItem := range creditManager {
		creditManagerRule = append(creditManagerRule, creditManagerItem)
	}

	logs, sub, err := _MockPoolService.contract.FilterLogs(opts, "BorrowForbidden", creditManagerRule)
	if err != nil {
		return nil, err
	}
	return &MockPoolServiceBorrowForbiddenIterator{contract: _MockPoolService.contract, event: "BorrowForbidden", logs: logs, sub: sub}, nil
}

// WatchBorrowForbidden is a free log subscription operation binding the contract event 0x9181736fce85d2d4cca2e4406f10679302ae5c387180fdb62963af3cd9a24fd6.
//
// Solidity: event BorrowForbidden(address indexed creditManager)
func (_MockPoolService *MockPoolServiceFilterer) WatchBorrowForbidden(opts *bind.WatchOpts, sink chan<- *MockPoolServiceBorrowForbidden, creditManager []common.Address) (event.Subscription, error) {

	var creditManagerRule []interface{}
	for _, creditManagerItem := range creditManager {
		creditManagerRule = append(creditManagerRule, creditManagerItem)
	}

	logs, sub, err := _MockPoolService.contract.WatchLogs(opts, "BorrowForbidden", creditManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockPoolServiceBorrowForbidden)
				if err := _MockPoolService.contract.UnpackLog(event, "BorrowForbidden", log); err != nil {
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
func (_MockPoolService *MockPoolServiceFilterer) ParseBorrowForbidden(log types.Log) (*MockPoolServiceBorrowForbidden, error) {
	event := new(MockPoolServiceBorrowForbidden)
	if err := _MockPoolService.contract.UnpackLog(event, "BorrowForbidden", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockPoolServiceNewCreditManagerConnectedIterator is returned from FilterNewCreditManagerConnected and is used to iterate over the raw logs and unpacked data for NewCreditManagerConnected events raised by the MockPoolService contract.
type MockPoolServiceNewCreditManagerConnectedIterator struct {
	Event *MockPoolServiceNewCreditManagerConnected // Event containing the contract specifics and raw log

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
func (it *MockPoolServiceNewCreditManagerConnectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockPoolServiceNewCreditManagerConnected)
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
		it.Event = new(MockPoolServiceNewCreditManagerConnected)
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
func (it *MockPoolServiceNewCreditManagerConnectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockPoolServiceNewCreditManagerConnectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockPoolServiceNewCreditManagerConnected represents a NewCreditManagerConnected event raised by the MockPoolService contract.
type MockPoolServiceNewCreditManagerConnected struct {
	CreditManager common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNewCreditManagerConnected is a free log retrieval operation binding the contract event 0xe076020e7eac3915d33aec40c24f95e73eb6c9921ff89747d50aa8fd934d2c01.
//
// Solidity: event NewCreditManagerConnected(address indexed creditManager)
func (_MockPoolService *MockPoolServiceFilterer) FilterNewCreditManagerConnected(opts *bind.FilterOpts, creditManager []common.Address) (*MockPoolServiceNewCreditManagerConnectedIterator, error) {

	var creditManagerRule []interface{}
	for _, creditManagerItem := range creditManager {
		creditManagerRule = append(creditManagerRule, creditManagerItem)
	}

	logs, sub, err := _MockPoolService.contract.FilterLogs(opts, "NewCreditManagerConnected", creditManagerRule)
	if err != nil {
		return nil, err
	}
	return &MockPoolServiceNewCreditManagerConnectedIterator{contract: _MockPoolService.contract, event: "NewCreditManagerConnected", logs: logs, sub: sub}, nil
}

// WatchNewCreditManagerConnected is a free log subscription operation binding the contract event 0xe076020e7eac3915d33aec40c24f95e73eb6c9921ff89747d50aa8fd934d2c01.
//
// Solidity: event NewCreditManagerConnected(address indexed creditManager)
func (_MockPoolService *MockPoolServiceFilterer) WatchNewCreditManagerConnected(opts *bind.WatchOpts, sink chan<- *MockPoolServiceNewCreditManagerConnected, creditManager []common.Address) (event.Subscription, error) {

	var creditManagerRule []interface{}
	for _, creditManagerItem := range creditManager {
		creditManagerRule = append(creditManagerRule, creditManagerItem)
	}

	logs, sub, err := _MockPoolService.contract.WatchLogs(opts, "NewCreditManagerConnected", creditManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockPoolServiceNewCreditManagerConnected)
				if err := _MockPoolService.contract.UnpackLog(event, "NewCreditManagerConnected", log); err != nil {
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
func (_MockPoolService *MockPoolServiceFilterer) ParseNewCreditManagerConnected(log types.Log) (*MockPoolServiceNewCreditManagerConnected, error) {
	event := new(MockPoolServiceNewCreditManagerConnected)
	if err := _MockPoolService.contract.UnpackLog(event, "NewCreditManagerConnected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockPoolServiceNewExpectedLiquidityLimitIterator is returned from FilterNewExpectedLiquidityLimit and is used to iterate over the raw logs and unpacked data for NewExpectedLiquidityLimit events raised by the MockPoolService contract.
type MockPoolServiceNewExpectedLiquidityLimitIterator struct {
	Event *MockPoolServiceNewExpectedLiquidityLimit // Event containing the contract specifics and raw log

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
func (it *MockPoolServiceNewExpectedLiquidityLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockPoolServiceNewExpectedLiquidityLimit)
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
		it.Event = new(MockPoolServiceNewExpectedLiquidityLimit)
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
func (it *MockPoolServiceNewExpectedLiquidityLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockPoolServiceNewExpectedLiquidityLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockPoolServiceNewExpectedLiquidityLimit represents a NewExpectedLiquidityLimit event raised by the MockPoolService contract.
type MockPoolServiceNewExpectedLiquidityLimit struct {
	NewLimit *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewExpectedLiquidityLimit is a free log retrieval operation binding the contract event 0xd7a183c9fe85b604c25d54bd676e0866f6c13bcca9fb9b0850213de118fdc99c.
//
// Solidity: event NewExpectedLiquidityLimit(uint256 newLimit)
func (_MockPoolService *MockPoolServiceFilterer) FilterNewExpectedLiquidityLimit(opts *bind.FilterOpts) (*MockPoolServiceNewExpectedLiquidityLimitIterator, error) {

	logs, sub, err := _MockPoolService.contract.FilterLogs(opts, "NewExpectedLiquidityLimit")
	if err != nil {
		return nil, err
	}
	return &MockPoolServiceNewExpectedLiquidityLimitIterator{contract: _MockPoolService.contract, event: "NewExpectedLiquidityLimit", logs: logs, sub: sub}, nil
}

// WatchNewExpectedLiquidityLimit is a free log subscription operation binding the contract event 0xd7a183c9fe85b604c25d54bd676e0866f6c13bcca9fb9b0850213de118fdc99c.
//
// Solidity: event NewExpectedLiquidityLimit(uint256 newLimit)
func (_MockPoolService *MockPoolServiceFilterer) WatchNewExpectedLiquidityLimit(opts *bind.WatchOpts, sink chan<- *MockPoolServiceNewExpectedLiquidityLimit) (event.Subscription, error) {

	logs, sub, err := _MockPoolService.contract.WatchLogs(opts, "NewExpectedLiquidityLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockPoolServiceNewExpectedLiquidityLimit)
				if err := _MockPoolService.contract.UnpackLog(event, "NewExpectedLiquidityLimit", log); err != nil {
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
func (_MockPoolService *MockPoolServiceFilterer) ParseNewExpectedLiquidityLimit(log types.Log) (*MockPoolServiceNewExpectedLiquidityLimit, error) {
	event := new(MockPoolServiceNewExpectedLiquidityLimit)
	if err := _MockPoolService.contract.UnpackLog(event, "NewExpectedLiquidityLimit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockPoolServiceNewInterestRateModelIterator is returned from FilterNewInterestRateModel and is used to iterate over the raw logs and unpacked data for NewInterestRateModel events raised by the MockPoolService contract.
type MockPoolServiceNewInterestRateModelIterator struct {
	Event *MockPoolServiceNewInterestRateModel // Event containing the contract specifics and raw log

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
func (it *MockPoolServiceNewInterestRateModelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockPoolServiceNewInterestRateModel)
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
		it.Event = new(MockPoolServiceNewInterestRateModel)
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
func (it *MockPoolServiceNewInterestRateModelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockPoolServiceNewInterestRateModelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockPoolServiceNewInterestRateModel represents a NewInterestRateModel event raised by the MockPoolService contract.
type MockPoolServiceNewInterestRateModel struct {
	NewInterestRateModel common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterNewInterestRateModel is a free log retrieval operation binding the contract event 0x0ec6cb7631d36954a05ffd646135bfd9995c71e7fa36d26abb1ad9f24a040ea1.
//
// Solidity: event NewInterestRateModel(address indexed newInterestRateModel)
func (_MockPoolService *MockPoolServiceFilterer) FilterNewInterestRateModel(opts *bind.FilterOpts, newInterestRateModel []common.Address) (*MockPoolServiceNewInterestRateModelIterator, error) {

	var newInterestRateModelRule []interface{}
	for _, newInterestRateModelItem := range newInterestRateModel {
		newInterestRateModelRule = append(newInterestRateModelRule, newInterestRateModelItem)
	}

	logs, sub, err := _MockPoolService.contract.FilterLogs(opts, "NewInterestRateModel", newInterestRateModelRule)
	if err != nil {
		return nil, err
	}
	return &MockPoolServiceNewInterestRateModelIterator{contract: _MockPoolService.contract, event: "NewInterestRateModel", logs: logs, sub: sub}, nil
}

// WatchNewInterestRateModel is a free log subscription operation binding the contract event 0x0ec6cb7631d36954a05ffd646135bfd9995c71e7fa36d26abb1ad9f24a040ea1.
//
// Solidity: event NewInterestRateModel(address indexed newInterestRateModel)
func (_MockPoolService *MockPoolServiceFilterer) WatchNewInterestRateModel(opts *bind.WatchOpts, sink chan<- *MockPoolServiceNewInterestRateModel, newInterestRateModel []common.Address) (event.Subscription, error) {

	var newInterestRateModelRule []interface{}
	for _, newInterestRateModelItem := range newInterestRateModel {
		newInterestRateModelRule = append(newInterestRateModelRule, newInterestRateModelItem)
	}

	logs, sub, err := _MockPoolService.contract.WatchLogs(opts, "NewInterestRateModel", newInterestRateModelRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockPoolServiceNewInterestRateModel)
				if err := _MockPoolService.contract.UnpackLog(event, "NewInterestRateModel", log); err != nil {
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
func (_MockPoolService *MockPoolServiceFilterer) ParseNewInterestRateModel(log types.Log) (*MockPoolServiceNewInterestRateModel, error) {
	event := new(MockPoolServiceNewInterestRateModel)
	if err := _MockPoolService.contract.UnpackLog(event, "NewInterestRateModel", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockPoolServiceNewWithdrawFeeIterator is returned from FilterNewWithdrawFee and is used to iterate over the raw logs and unpacked data for NewWithdrawFee events raised by the MockPoolService contract.
type MockPoolServiceNewWithdrawFeeIterator struct {
	Event *MockPoolServiceNewWithdrawFee // Event containing the contract specifics and raw log

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
func (it *MockPoolServiceNewWithdrawFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockPoolServiceNewWithdrawFee)
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
		it.Event = new(MockPoolServiceNewWithdrawFee)
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
func (it *MockPoolServiceNewWithdrawFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockPoolServiceNewWithdrawFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockPoolServiceNewWithdrawFee represents a NewWithdrawFee event raised by the MockPoolService contract.
type MockPoolServiceNewWithdrawFee struct {
	Fee *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterNewWithdrawFee is a free log retrieval operation binding the contract event 0xd5fe46099fa396290a7f57e36c3c3c8774e2562c18ed5d1dcc0fa75071e03f1d.
//
// Solidity: event NewWithdrawFee(uint256 fee)
func (_MockPoolService *MockPoolServiceFilterer) FilterNewWithdrawFee(opts *bind.FilterOpts) (*MockPoolServiceNewWithdrawFeeIterator, error) {

	logs, sub, err := _MockPoolService.contract.FilterLogs(opts, "NewWithdrawFee")
	if err != nil {
		return nil, err
	}
	return &MockPoolServiceNewWithdrawFeeIterator{contract: _MockPoolService.contract, event: "NewWithdrawFee", logs: logs, sub: sub}, nil
}

// WatchNewWithdrawFee is a free log subscription operation binding the contract event 0xd5fe46099fa396290a7f57e36c3c3c8774e2562c18ed5d1dcc0fa75071e03f1d.
//
// Solidity: event NewWithdrawFee(uint256 fee)
func (_MockPoolService *MockPoolServiceFilterer) WatchNewWithdrawFee(opts *bind.WatchOpts, sink chan<- *MockPoolServiceNewWithdrawFee) (event.Subscription, error) {

	logs, sub, err := _MockPoolService.contract.WatchLogs(opts, "NewWithdrawFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockPoolServiceNewWithdrawFee)
				if err := _MockPoolService.contract.UnpackLog(event, "NewWithdrawFee", log); err != nil {
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
func (_MockPoolService *MockPoolServiceFilterer) ParseNewWithdrawFee(log types.Log) (*MockPoolServiceNewWithdrawFee, error) {
	event := new(MockPoolServiceNewWithdrawFee)
	if err := _MockPoolService.contract.UnpackLog(event, "NewWithdrawFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockPoolServiceRemoveLiquidityIterator is returned from FilterRemoveLiquidity and is used to iterate over the raw logs and unpacked data for RemoveLiquidity events raised by the MockPoolService contract.
type MockPoolServiceRemoveLiquidityIterator struct {
	Event *MockPoolServiceRemoveLiquidity // Event containing the contract specifics and raw log

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
func (it *MockPoolServiceRemoveLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockPoolServiceRemoveLiquidity)
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
		it.Event = new(MockPoolServiceRemoveLiquidity)
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
func (it *MockPoolServiceRemoveLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockPoolServiceRemoveLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockPoolServiceRemoveLiquidity represents a RemoveLiquidity event raised by the MockPoolService contract.
type MockPoolServiceRemoveLiquidity struct {
	Sender common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidity is a free log retrieval operation binding the contract event 0xd8ae9b9ba89e637bcb66a69ac91e8f688018e81d6f92c57e02226425c8efbdf6.
//
// Solidity: event RemoveLiquidity(address indexed sender, address indexed to, uint256 amount)
func (_MockPoolService *MockPoolServiceFilterer) FilterRemoveLiquidity(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*MockPoolServiceRemoveLiquidityIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MockPoolService.contract.FilterLogs(opts, "RemoveLiquidity", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MockPoolServiceRemoveLiquidityIterator{contract: _MockPoolService.contract, event: "RemoveLiquidity", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidity is a free log subscription operation binding the contract event 0xd8ae9b9ba89e637bcb66a69ac91e8f688018e81d6f92c57e02226425c8efbdf6.
//
// Solidity: event RemoveLiquidity(address indexed sender, address indexed to, uint256 amount)
func (_MockPoolService *MockPoolServiceFilterer) WatchRemoveLiquidity(opts *bind.WatchOpts, sink chan<- *MockPoolServiceRemoveLiquidity, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MockPoolService.contract.WatchLogs(opts, "RemoveLiquidity", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockPoolServiceRemoveLiquidity)
				if err := _MockPoolService.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
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
func (_MockPoolService *MockPoolServiceFilterer) ParseRemoveLiquidity(log types.Log) (*MockPoolServiceRemoveLiquidity, error) {
	event := new(MockPoolServiceRemoveLiquidity)
	if err := _MockPoolService.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockPoolServiceRepayIterator is returned from FilterRepay and is used to iterate over the raw logs and unpacked data for Repay events raised by the MockPoolService contract.
type MockPoolServiceRepayIterator struct {
	Event *MockPoolServiceRepay // Event containing the contract specifics and raw log

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
func (it *MockPoolServiceRepayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockPoolServiceRepay)
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
		it.Event = new(MockPoolServiceRepay)
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
func (it *MockPoolServiceRepayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockPoolServiceRepayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockPoolServiceRepay represents a Repay event raised by the MockPoolService contract.
type MockPoolServiceRepay struct {
	CreditManager  common.Address
	BorrowedAmount *big.Int
	Profit         *big.Int
	Loss           *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRepay is a free log retrieval operation binding the contract event 0x2fe77b1c99aca6b022b8efc6e3e8dd1b48b30748709339b65c50ef3263443e09.
//
// Solidity: event Repay(address indexed creditManager, uint256 borrowedAmount, uint256 profit, uint256 loss)
func (_MockPoolService *MockPoolServiceFilterer) FilterRepay(opts *bind.FilterOpts, creditManager []common.Address) (*MockPoolServiceRepayIterator, error) {

	var creditManagerRule []interface{}
	for _, creditManagerItem := range creditManager {
		creditManagerRule = append(creditManagerRule, creditManagerItem)
	}

	logs, sub, err := _MockPoolService.contract.FilterLogs(opts, "Repay", creditManagerRule)
	if err != nil {
		return nil, err
	}
	return &MockPoolServiceRepayIterator{contract: _MockPoolService.contract, event: "Repay", logs: logs, sub: sub}, nil
}

// WatchRepay is a free log subscription operation binding the contract event 0x2fe77b1c99aca6b022b8efc6e3e8dd1b48b30748709339b65c50ef3263443e09.
//
// Solidity: event Repay(address indexed creditManager, uint256 borrowedAmount, uint256 profit, uint256 loss)
func (_MockPoolService *MockPoolServiceFilterer) WatchRepay(opts *bind.WatchOpts, sink chan<- *MockPoolServiceRepay, creditManager []common.Address) (event.Subscription, error) {

	var creditManagerRule []interface{}
	for _, creditManagerItem := range creditManager {
		creditManagerRule = append(creditManagerRule, creditManagerItem)
	}

	logs, sub, err := _MockPoolService.contract.WatchLogs(opts, "Repay", creditManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockPoolServiceRepay)
				if err := _MockPoolService.contract.UnpackLog(event, "Repay", log); err != nil {
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
func (_MockPoolService *MockPoolServiceFilterer) ParseRepay(log types.Log) (*MockPoolServiceRepay, error) {
	event := new(MockPoolServiceRepay)
	if err := _MockPoolService.contract.UnpackLog(event, "Repay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockPoolServiceUncoveredLossIterator is returned from FilterUncoveredLoss and is used to iterate over the raw logs and unpacked data for UncoveredLoss events raised by the MockPoolService contract.
type MockPoolServiceUncoveredLossIterator struct {
	Event *MockPoolServiceUncoveredLoss // Event containing the contract specifics and raw log

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
func (it *MockPoolServiceUncoveredLossIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockPoolServiceUncoveredLoss)
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
		it.Event = new(MockPoolServiceUncoveredLoss)
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
func (it *MockPoolServiceUncoveredLossIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockPoolServiceUncoveredLossIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockPoolServiceUncoveredLoss represents a UncoveredLoss event raised by the MockPoolService contract.
type MockPoolServiceUncoveredLoss struct {
	CreditManager common.Address
	Loss          *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUncoveredLoss is a free log retrieval operation binding the contract event 0xef3653ded679720ab04913b6f3820be7cedc8286d42ff5dd8dff17e91bd2964c.
//
// Solidity: event UncoveredLoss(address indexed creditManager, uint256 loss)
func (_MockPoolService *MockPoolServiceFilterer) FilterUncoveredLoss(opts *bind.FilterOpts, creditManager []common.Address) (*MockPoolServiceUncoveredLossIterator, error) {

	var creditManagerRule []interface{}
	for _, creditManagerItem := range creditManager {
		creditManagerRule = append(creditManagerRule, creditManagerItem)
	}

	logs, sub, err := _MockPoolService.contract.FilterLogs(opts, "UncoveredLoss", creditManagerRule)
	if err != nil {
		return nil, err
	}
	return &MockPoolServiceUncoveredLossIterator{contract: _MockPoolService.contract, event: "UncoveredLoss", logs: logs, sub: sub}, nil
}

// WatchUncoveredLoss is a free log subscription operation binding the contract event 0xef3653ded679720ab04913b6f3820be7cedc8286d42ff5dd8dff17e91bd2964c.
//
// Solidity: event UncoveredLoss(address indexed creditManager, uint256 loss)
func (_MockPoolService *MockPoolServiceFilterer) WatchUncoveredLoss(opts *bind.WatchOpts, sink chan<- *MockPoolServiceUncoveredLoss, creditManager []common.Address) (event.Subscription, error) {

	var creditManagerRule []interface{}
	for _, creditManagerItem := range creditManager {
		creditManagerRule = append(creditManagerRule, creditManagerItem)
	}

	logs, sub, err := _MockPoolService.contract.WatchLogs(opts, "UncoveredLoss", creditManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockPoolServiceUncoveredLoss)
				if err := _MockPoolService.contract.UnpackLog(event, "UncoveredLoss", log); err != nil {
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
func (_MockPoolService *MockPoolServiceFilterer) ParseUncoveredLoss(log types.Log) (*MockPoolServiceUncoveredLoss, error) {
	event := new(MockPoolServiceUncoveredLoss)
	if err := _MockPoolService.contract.UnpackLog(event, "UncoveredLoss", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
