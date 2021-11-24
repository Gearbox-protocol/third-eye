// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package poolService

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

// PoolServiceMetaData contains all meta data concerning the PoolService contract.
var PoolServiceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addressProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_underlyingToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_dieselAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_interestRateModelAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_expectedLiquidityLimit\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"referralCode\",\"type\":\"uint256\"}],\"name\":\"AddLiquidity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creditManager\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Borrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creditManager\",\"type\":\"address\"}],\"name\":\"BorrowForbidden\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creditManager\",\"type\":\"address\"}],\"name\":\"NewCreditManagerConnected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newLimit\",\"type\":\"uint256\"}],\"name\":\"NewExpectedLiquidityLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newInterestRateModel\",\"type\":\"address\"}],\"name\":\"NewInterestRateModel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"NewWithdrawFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RemoveLiquidity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creditManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"borrowedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"profit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"loss\",\"type\":\"uint256\"}],\"name\":\"Repay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creditManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"loss\",\"type\":\"uint256\"}],\"name\":\"UncoveredLoss\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_cumulativeIndex_RAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_expectedLiquidityLU\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_timestampLU\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"referralCode\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addressProvider\",\"outputs\":[{\"internalType\":\"contractAddressProvider\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"availableLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"borrowAPY_RAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cumulativeIndexAtOpen\",\"type\":\"uint256\"}],\"name\":\"calcCumulativeIndexAtBorrowMore\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calcLinearCumulative_RAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cumulativeIndex_RAY\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentBorrowRate_RAY\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeDifference\",\"type\":\"uint256\"}],\"name\":\"calcLinearIndex_RAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_creditManager\",\"type\":\"address\"}],\"name\":\"connectCreditManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"creditManagers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"creditManagersCanBorrow\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"creditManagersCanRepay\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creditManagersCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dieselToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expectedLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expectedLiquidityLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_creditManager\",\"type\":\"address\"}],\"name\":\"forbidCreditManagerToBorrow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"fromDiesel\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDieselRate_RAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"interestRateModel\",\"outputs\":[{\"internalType\":\"contractIInterestRateModel\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"borrowedAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"}],\"name\":\"lendCreditAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"removeLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"borrowedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"profit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"loss\",\"type\":\"uint256\"}],\"name\":\"repayCreditAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newLimit\",\"type\":\"uint256\"}],\"name\":\"setExpectedLiquidityLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"setWithdrawFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"toDiesel\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBorrowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasuryAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlyingToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_interestRateModel\",\"type\":\"address\"}],\"name\":\"updateInterestRateModel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// PoolServiceABI is the input ABI used to generate the binding from.
// Deprecated: Use PoolServiceMetaData.ABI instead.
var PoolServiceABI = PoolServiceMetaData.ABI

// PoolService is an auto generated Go binding around an Ethereum contract.
type PoolService struct {
	PoolServiceCaller     // Read-only binding to the contract
	PoolServiceTransactor // Write-only binding to the contract
	PoolServiceFilterer   // Log filterer for contract events
}

// PoolServiceCaller is an auto generated read-only Go binding around an Ethereum contract.
type PoolServiceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolServiceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PoolServiceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolServiceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PoolServiceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolServiceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PoolServiceSession struct {
	Contract     *PoolService      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoolServiceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PoolServiceCallerSession struct {
	Contract *PoolServiceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// PoolServiceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PoolServiceTransactorSession struct {
	Contract     *PoolServiceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// PoolServiceRaw is an auto generated low-level Go binding around an Ethereum contract.
type PoolServiceRaw struct {
	Contract *PoolService // Generic contract binding to access the raw methods on
}

// PoolServiceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PoolServiceCallerRaw struct {
	Contract *PoolServiceCaller // Generic read-only contract binding to access the raw methods on
}

// PoolServiceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PoolServiceTransactorRaw struct {
	Contract *PoolServiceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPoolService creates a new instance of PoolService, bound to a specific deployed contract.
func NewPoolService(address common.Address, backend bind.ContractBackend) (*PoolService, error) {
	contract, err := bindPoolService(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PoolService{PoolServiceCaller: PoolServiceCaller{contract: contract}, PoolServiceTransactor: PoolServiceTransactor{contract: contract}, PoolServiceFilterer: PoolServiceFilterer{contract: contract}}, nil
}

// NewPoolServiceCaller creates a new read-only instance of PoolService, bound to a specific deployed contract.
func NewPoolServiceCaller(address common.Address, caller bind.ContractCaller) (*PoolServiceCaller, error) {
	contract, err := bindPoolService(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PoolServiceCaller{contract: contract}, nil
}

// NewPoolServiceTransactor creates a new write-only instance of PoolService, bound to a specific deployed contract.
func NewPoolServiceTransactor(address common.Address, transactor bind.ContractTransactor) (*PoolServiceTransactor, error) {
	contract, err := bindPoolService(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PoolServiceTransactor{contract: contract}, nil
}

// NewPoolServiceFilterer creates a new log filterer instance of PoolService, bound to a specific deployed contract.
func NewPoolServiceFilterer(address common.Address, filterer bind.ContractFilterer) (*PoolServiceFilterer, error) {
	contract, err := bindPoolService(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PoolServiceFilterer{contract: contract}, nil
}

// bindPoolService binds a generic wrapper to an already deployed contract.
func bindPoolService(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PoolServiceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoolService *PoolServiceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PoolService.Contract.PoolServiceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoolService *PoolServiceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolService.Contract.PoolServiceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoolService *PoolServiceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoolService.Contract.PoolServiceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoolService *PoolServiceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PoolService.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoolService *PoolServiceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolService.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoolService *PoolServiceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoolService.Contract.contract.Transact(opts, method, params...)
}

// CumulativeIndexRAY is a free data retrieval call binding the contract method 0xdbcb313b.
//
// Solidity: function _cumulativeIndex_RAY() view returns(uint256)
func (_PoolService *PoolServiceCaller) CumulativeIndexRAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PoolService.contract.Call(opts, &out, "_cumulativeIndex_RAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CumulativeIndexRAY is a free data retrieval call binding the contract method 0xdbcb313b.
//
// Solidity: function _cumulativeIndex_RAY() view returns(uint256)
func (_PoolService *PoolServiceSession) CumulativeIndexRAY() (*big.Int, error) {
	return _PoolService.Contract.CumulativeIndexRAY(&_PoolService.CallOpts)
}

// CumulativeIndexRAY is a free data retrieval call binding the contract method 0xdbcb313b.
//
// Solidity: function _cumulativeIndex_RAY() view returns(uint256)
func (_PoolService *PoolServiceCallerSession) CumulativeIndexRAY() (*big.Int, error) {
	return _PoolService.Contract.CumulativeIndexRAY(&_PoolService.CallOpts)
}

// ExpectedLiquidityLU is a free data retrieval call binding the contract method 0x030dbb04.
//
// Solidity: function _expectedLiquidityLU() view returns(uint256)
func (_PoolService *PoolServiceCaller) ExpectedLiquidityLU(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PoolService.contract.Call(opts, &out, "_expectedLiquidityLU")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExpectedLiquidityLU is a free data retrieval call binding the contract method 0x030dbb04.
//
// Solidity: function _expectedLiquidityLU() view returns(uint256)
func (_PoolService *PoolServiceSession) ExpectedLiquidityLU() (*big.Int, error) {
	return _PoolService.Contract.ExpectedLiquidityLU(&_PoolService.CallOpts)
}

// ExpectedLiquidityLU is a free data retrieval call binding the contract method 0x030dbb04.
//
// Solidity: function _expectedLiquidityLU() view returns(uint256)
func (_PoolService *PoolServiceCallerSession) ExpectedLiquidityLU() (*big.Int, error) {
	return _PoolService.Contract.ExpectedLiquidityLU(&_PoolService.CallOpts)
}

// TimestampLU is a free data retrieval call binding the contract method 0x609ae317.
//
// Solidity: function _timestampLU() view returns(uint256)
func (_PoolService *PoolServiceCaller) TimestampLU(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PoolService.contract.Call(opts, &out, "_timestampLU")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TimestampLU is a free data retrieval call binding the contract method 0x609ae317.
//
// Solidity: function _timestampLU() view returns(uint256)
func (_PoolService *PoolServiceSession) TimestampLU() (*big.Int, error) {
	return _PoolService.Contract.TimestampLU(&_PoolService.CallOpts)
}

// TimestampLU is a free data retrieval call binding the contract method 0x609ae317.
//
// Solidity: function _timestampLU() view returns(uint256)
func (_PoolService *PoolServiceCallerSession) TimestampLU() (*big.Int, error) {
	return _PoolService.Contract.TimestampLU(&_PoolService.CallOpts)
}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_PoolService *PoolServiceCaller) AddressProvider(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoolService.contract.Call(opts, &out, "addressProvider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_PoolService *PoolServiceSession) AddressProvider() (common.Address, error) {
	return _PoolService.Contract.AddressProvider(&_PoolService.CallOpts)
}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_PoolService *PoolServiceCallerSession) AddressProvider() (common.Address, error) {
	return _PoolService.Contract.AddressProvider(&_PoolService.CallOpts)
}

// AvailableLiquidity is a free data retrieval call binding the contract method 0x74375359.
//
// Solidity: function availableLiquidity() view returns(uint256)
func (_PoolService *PoolServiceCaller) AvailableLiquidity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PoolService.contract.Call(opts, &out, "availableLiquidity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AvailableLiquidity is a free data retrieval call binding the contract method 0x74375359.
//
// Solidity: function availableLiquidity() view returns(uint256)
func (_PoolService *PoolServiceSession) AvailableLiquidity() (*big.Int, error) {
	return _PoolService.Contract.AvailableLiquidity(&_PoolService.CallOpts)
}

// AvailableLiquidity is a free data retrieval call binding the contract method 0x74375359.
//
// Solidity: function availableLiquidity() view returns(uint256)
func (_PoolService *PoolServiceCallerSession) AvailableLiquidity() (*big.Int, error) {
	return _PoolService.Contract.AvailableLiquidity(&_PoolService.CallOpts)
}

// BorrowAPYRAY is a free data retrieval call binding the contract method 0x45d31f9d.
//
// Solidity: function borrowAPY_RAY() view returns(uint256)
func (_PoolService *PoolServiceCaller) BorrowAPYRAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PoolService.contract.Call(opts, &out, "borrowAPY_RAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowAPYRAY is a free data retrieval call binding the contract method 0x45d31f9d.
//
// Solidity: function borrowAPY_RAY() view returns(uint256)
func (_PoolService *PoolServiceSession) BorrowAPYRAY() (*big.Int, error) {
	return _PoolService.Contract.BorrowAPYRAY(&_PoolService.CallOpts)
}

// BorrowAPYRAY is a free data retrieval call binding the contract method 0x45d31f9d.
//
// Solidity: function borrowAPY_RAY() view returns(uint256)
func (_PoolService *PoolServiceCallerSession) BorrowAPYRAY() (*big.Int, error) {
	return _PoolService.Contract.BorrowAPYRAY(&_PoolService.CallOpts)
}

// CalcCumulativeIndexAtBorrowMore is a free data retrieval call binding the contract method 0xc00495a1.
//
// Solidity: function calcCumulativeIndexAtBorrowMore(uint256 amount, uint256 dAmount, uint256 cumulativeIndexAtOpen) view returns(uint256)
func (_PoolService *PoolServiceCaller) CalcCumulativeIndexAtBorrowMore(opts *bind.CallOpts, amount *big.Int, dAmount *big.Int, cumulativeIndexAtOpen *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PoolService.contract.Call(opts, &out, "calcCumulativeIndexAtBorrowMore", amount, dAmount, cumulativeIndexAtOpen)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcCumulativeIndexAtBorrowMore is a free data retrieval call binding the contract method 0xc00495a1.
//
// Solidity: function calcCumulativeIndexAtBorrowMore(uint256 amount, uint256 dAmount, uint256 cumulativeIndexAtOpen) view returns(uint256)
func (_PoolService *PoolServiceSession) CalcCumulativeIndexAtBorrowMore(amount *big.Int, dAmount *big.Int, cumulativeIndexAtOpen *big.Int) (*big.Int, error) {
	return _PoolService.Contract.CalcCumulativeIndexAtBorrowMore(&_PoolService.CallOpts, amount, dAmount, cumulativeIndexAtOpen)
}

// CalcCumulativeIndexAtBorrowMore is a free data retrieval call binding the contract method 0xc00495a1.
//
// Solidity: function calcCumulativeIndexAtBorrowMore(uint256 amount, uint256 dAmount, uint256 cumulativeIndexAtOpen) view returns(uint256)
func (_PoolService *PoolServiceCallerSession) CalcCumulativeIndexAtBorrowMore(amount *big.Int, dAmount *big.Int, cumulativeIndexAtOpen *big.Int) (*big.Int, error) {
	return _PoolService.Contract.CalcCumulativeIndexAtBorrowMore(&_PoolService.CallOpts, amount, dAmount, cumulativeIndexAtOpen)
}

// CalcLinearCumulativeRAY is a free data retrieval call binding the contract method 0x0fce70fb.
//
// Solidity: function calcLinearCumulative_RAY() view returns(uint256)
func (_PoolService *PoolServiceCaller) CalcLinearCumulativeRAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PoolService.contract.Call(opts, &out, "calcLinearCumulative_RAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcLinearCumulativeRAY is a free data retrieval call binding the contract method 0x0fce70fb.
//
// Solidity: function calcLinearCumulative_RAY() view returns(uint256)
func (_PoolService *PoolServiceSession) CalcLinearCumulativeRAY() (*big.Int, error) {
	return _PoolService.Contract.CalcLinearCumulativeRAY(&_PoolService.CallOpts)
}

// CalcLinearCumulativeRAY is a free data retrieval call binding the contract method 0x0fce70fb.
//
// Solidity: function calcLinearCumulative_RAY() view returns(uint256)
func (_PoolService *PoolServiceCallerSession) CalcLinearCumulativeRAY() (*big.Int, error) {
	return _PoolService.Contract.CalcLinearCumulativeRAY(&_PoolService.CallOpts)
}

// CalcLinearIndexRAY is a free data retrieval call binding the contract method 0x31d8bc27.
//
// Solidity: function calcLinearIndex_RAY(uint256 cumulativeIndex_RAY, uint256 currentBorrowRate_RAY, uint256 timeDifference) pure returns(uint256)
func (_PoolService *PoolServiceCaller) CalcLinearIndexRAY(opts *bind.CallOpts, cumulativeIndex_RAY *big.Int, currentBorrowRate_RAY *big.Int, timeDifference *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PoolService.contract.Call(opts, &out, "calcLinearIndex_RAY", cumulativeIndex_RAY, currentBorrowRate_RAY, timeDifference)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcLinearIndexRAY is a free data retrieval call binding the contract method 0x31d8bc27.
//
// Solidity: function calcLinearIndex_RAY(uint256 cumulativeIndex_RAY, uint256 currentBorrowRate_RAY, uint256 timeDifference) pure returns(uint256)
func (_PoolService *PoolServiceSession) CalcLinearIndexRAY(cumulativeIndex_RAY *big.Int, currentBorrowRate_RAY *big.Int, timeDifference *big.Int) (*big.Int, error) {
	return _PoolService.Contract.CalcLinearIndexRAY(&_PoolService.CallOpts, cumulativeIndex_RAY, currentBorrowRate_RAY, timeDifference)
}

// CalcLinearIndexRAY is a free data retrieval call binding the contract method 0x31d8bc27.
//
// Solidity: function calcLinearIndex_RAY(uint256 cumulativeIndex_RAY, uint256 currentBorrowRate_RAY, uint256 timeDifference) pure returns(uint256)
func (_PoolService *PoolServiceCallerSession) CalcLinearIndexRAY(cumulativeIndex_RAY *big.Int, currentBorrowRate_RAY *big.Int, timeDifference *big.Int) (*big.Int, error) {
	return _PoolService.Contract.CalcLinearIndexRAY(&_PoolService.CallOpts, cumulativeIndex_RAY, currentBorrowRate_RAY, timeDifference)
}

// CreditManagers is a free data retrieval call binding the contract method 0x1e16e4fc.
//
// Solidity: function creditManagers(uint256 ) view returns(address)
func (_PoolService *PoolServiceCaller) CreditManagers(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _PoolService.contract.Call(opts, &out, "creditManagers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditManagers is a free data retrieval call binding the contract method 0x1e16e4fc.
//
// Solidity: function creditManagers(uint256 ) view returns(address)
func (_PoolService *PoolServiceSession) CreditManagers(arg0 *big.Int) (common.Address, error) {
	return _PoolService.Contract.CreditManagers(&_PoolService.CallOpts, arg0)
}

// CreditManagers is a free data retrieval call binding the contract method 0x1e16e4fc.
//
// Solidity: function creditManagers(uint256 ) view returns(address)
func (_PoolService *PoolServiceCallerSession) CreditManagers(arg0 *big.Int) (common.Address, error) {
	return _PoolService.Contract.CreditManagers(&_PoolService.CallOpts, arg0)
}

// CreditManagersCanBorrow is a free data retrieval call binding the contract method 0x2e97ca21.
//
// Solidity: function creditManagersCanBorrow(address ) view returns(bool)
func (_PoolService *PoolServiceCaller) CreditManagersCanBorrow(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _PoolService.contract.Call(opts, &out, "creditManagersCanBorrow", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CreditManagersCanBorrow is a free data retrieval call binding the contract method 0x2e97ca21.
//
// Solidity: function creditManagersCanBorrow(address ) view returns(bool)
func (_PoolService *PoolServiceSession) CreditManagersCanBorrow(arg0 common.Address) (bool, error) {
	return _PoolService.Contract.CreditManagersCanBorrow(&_PoolService.CallOpts, arg0)
}

// CreditManagersCanBorrow is a free data retrieval call binding the contract method 0x2e97ca21.
//
// Solidity: function creditManagersCanBorrow(address ) view returns(bool)
func (_PoolService *PoolServiceCallerSession) CreditManagersCanBorrow(arg0 common.Address) (bool, error) {
	return _PoolService.Contract.CreditManagersCanBorrow(&_PoolService.CallOpts, arg0)
}

// CreditManagersCanRepay is a free data retrieval call binding the contract method 0x3e163df0.
//
// Solidity: function creditManagersCanRepay(address ) view returns(bool)
func (_PoolService *PoolServiceCaller) CreditManagersCanRepay(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _PoolService.contract.Call(opts, &out, "creditManagersCanRepay", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CreditManagersCanRepay is a free data retrieval call binding the contract method 0x3e163df0.
//
// Solidity: function creditManagersCanRepay(address ) view returns(bool)
func (_PoolService *PoolServiceSession) CreditManagersCanRepay(arg0 common.Address) (bool, error) {
	return _PoolService.Contract.CreditManagersCanRepay(&_PoolService.CallOpts, arg0)
}

// CreditManagersCanRepay is a free data retrieval call binding the contract method 0x3e163df0.
//
// Solidity: function creditManagersCanRepay(address ) view returns(bool)
func (_PoolService *PoolServiceCallerSession) CreditManagersCanRepay(arg0 common.Address) (bool, error) {
	return _PoolService.Contract.CreditManagersCanRepay(&_PoolService.CallOpts, arg0)
}

// CreditManagersCount is a free data retrieval call binding the contract method 0xa4e8273e.
//
// Solidity: function creditManagersCount() view returns(uint256)
func (_PoolService *PoolServiceCaller) CreditManagersCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PoolService.contract.Call(opts, &out, "creditManagersCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CreditManagersCount is a free data retrieval call binding the contract method 0xa4e8273e.
//
// Solidity: function creditManagersCount() view returns(uint256)
func (_PoolService *PoolServiceSession) CreditManagersCount() (*big.Int, error) {
	return _PoolService.Contract.CreditManagersCount(&_PoolService.CallOpts)
}

// CreditManagersCount is a free data retrieval call binding the contract method 0xa4e8273e.
//
// Solidity: function creditManagersCount() view returns(uint256)
func (_PoolService *PoolServiceCallerSession) CreditManagersCount() (*big.Int, error) {
	return _PoolService.Contract.CreditManagersCount(&_PoolService.CallOpts)
}

// DieselToken is a free data retrieval call binding the contract method 0x36dda7d5.
//
// Solidity: function dieselToken() view returns(address)
func (_PoolService *PoolServiceCaller) DieselToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoolService.contract.Call(opts, &out, "dieselToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DieselToken is a free data retrieval call binding the contract method 0x36dda7d5.
//
// Solidity: function dieselToken() view returns(address)
func (_PoolService *PoolServiceSession) DieselToken() (common.Address, error) {
	return _PoolService.Contract.DieselToken(&_PoolService.CallOpts)
}

// DieselToken is a free data retrieval call binding the contract method 0x36dda7d5.
//
// Solidity: function dieselToken() view returns(address)
func (_PoolService *PoolServiceCallerSession) DieselToken() (common.Address, error) {
	return _PoolService.Contract.DieselToken(&_PoolService.CallOpts)
}

// ExpectedLiquidity is a free data retrieval call binding the contract method 0xfe14112d.
//
// Solidity: function expectedLiquidity() view returns(uint256)
func (_PoolService *PoolServiceCaller) ExpectedLiquidity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PoolService.contract.Call(opts, &out, "expectedLiquidity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExpectedLiquidity is a free data retrieval call binding the contract method 0xfe14112d.
//
// Solidity: function expectedLiquidity() view returns(uint256)
func (_PoolService *PoolServiceSession) ExpectedLiquidity() (*big.Int, error) {
	return _PoolService.Contract.ExpectedLiquidity(&_PoolService.CallOpts)
}

// ExpectedLiquidity is a free data retrieval call binding the contract method 0xfe14112d.
//
// Solidity: function expectedLiquidity() view returns(uint256)
func (_PoolService *PoolServiceCallerSession) ExpectedLiquidity() (*big.Int, error) {
	return _PoolService.Contract.ExpectedLiquidity(&_PoolService.CallOpts)
}

// ExpectedLiquidityLimit is a free data retrieval call binding the contract method 0xef8d9603.
//
// Solidity: function expectedLiquidityLimit() view returns(uint256)
func (_PoolService *PoolServiceCaller) ExpectedLiquidityLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PoolService.contract.Call(opts, &out, "expectedLiquidityLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExpectedLiquidityLimit is a free data retrieval call binding the contract method 0xef8d9603.
//
// Solidity: function expectedLiquidityLimit() view returns(uint256)
func (_PoolService *PoolServiceSession) ExpectedLiquidityLimit() (*big.Int, error) {
	return _PoolService.Contract.ExpectedLiquidityLimit(&_PoolService.CallOpts)
}

// ExpectedLiquidityLimit is a free data retrieval call binding the contract method 0xef8d9603.
//
// Solidity: function expectedLiquidityLimit() view returns(uint256)
func (_PoolService *PoolServiceCallerSession) ExpectedLiquidityLimit() (*big.Int, error) {
	return _PoolService.Contract.ExpectedLiquidityLimit(&_PoolService.CallOpts)
}

// FromDiesel is a free data retrieval call binding the contract method 0x5427c938.
//
// Solidity: function fromDiesel(uint256 amount) view returns(uint256)
func (_PoolService *PoolServiceCaller) FromDiesel(opts *bind.CallOpts, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PoolService.contract.Call(opts, &out, "fromDiesel", amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FromDiesel is a free data retrieval call binding the contract method 0x5427c938.
//
// Solidity: function fromDiesel(uint256 amount) view returns(uint256)
func (_PoolService *PoolServiceSession) FromDiesel(amount *big.Int) (*big.Int, error) {
	return _PoolService.Contract.FromDiesel(&_PoolService.CallOpts, amount)
}

// FromDiesel is a free data retrieval call binding the contract method 0x5427c938.
//
// Solidity: function fromDiesel(uint256 amount) view returns(uint256)
func (_PoolService *PoolServiceCallerSession) FromDiesel(amount *big.Int) (*big.Int, error) {
	return _PoolService.Contract.FromDiesel(&_PoolService.CallOpts, amount)
}

// GetDieselRateRAY is a free data retrieval call binding the contract method 0x788c6bfe.
//
// Solidity: function getDieselRate_RAY() view returns(uint256)
func (_PoolService *PoolServiceCaller) GetDieselRateRAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PoolService.contract.Call(opts, &out, "getDieselRate_RAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDieselRateRAY is a free data retrieval call binding the contract method 0x788c6bfe.
//
// Solidity: function getDieselRate_RAY() view returns(uint256)
func (_PoolService *PoolServiceSession) GetDieselRateRAY() (*big.Int, error) {
	return _PoolService.Contract.GetDieselRateRAY(&_PoolService.CallOpts)
}

// GetDieselRateRAY is a free data retrieval call binding the contract method 0x788c6bfe.
//
// Solidity: function getDieselRate_RAY() view returns(uint256)
func (_PoolService *PoolServiceCallerSession) GetDieselRateRAY() (*big.Int, error) {
	return _PoolService.Contract.GetDieselRateRAY(&_PoolService.CallOpts)
}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_PoolService *PoolServiceCaller) InterestRateModel(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoolService.contract.Call(opts, &out, "interestRateModel")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_PoolService *PoolServiceSession) InterestRateModel() (common.Address, error) {
	return _PoolService.Contract.InterestRateModel(&_PoolService.CallOpts)
}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_PoolService *PoolServiceCallerSession) InterestRateModel() (common.Address, error) {
	return _PoolService.Contract.InterestRateModel(&_PoolService.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PoolService *PoolServiceCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PoolService.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PoolService *PoolServiceSession) Paused() (bool, error) {
	return _PoolService.Contract.Paused(&_PoolService.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PoolService *PoolServiceCallerSession) Paused() (bool, error) {
	return _PoolService.Contract.Paused(&_PoolService.CallOpts)
}

// ToDiesel is a free data retrieval call binding the contract method 0x4d778ad1.
//
// Solidity: function toDiesel(uint256 amount) view returns(uint256)
func (_PoolService *PoolServiceCaller) ToDiesel(opts *bind.CallOpts, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PoolService.contract.Call(opts, &out, "toDiesel", amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ToDiesel is a free data retrieval call binding the contract method 0x4d778ad1.
//
// Solidity: function toDiesel(uint256 amount) view returns(uint256)
func (_PoolService *PoolServiceSession) ToDiesel(amount *big.Int) (*big.Int, error) {
	return _PoolService.Contract.ToDiesel(&_PoolService.CallOpts, amount)
}

// ToDiesel is a free data retrieval call binding the contract method 0x4d778ad1.
//
// Solidity: function toDiesel(uint256 amount) view returns(uint256)
func (_PoolService *PoolServiceCallerSession) ToDiesel(amount *big.Int) (*big.Int, error) {
	return _PoolService.Contract.ToDiesel(&_PoolService.CallOpts, amount)
}

// TotalBorrowed is a free data retrieval call binding the contract method 0x4c19386c.
//
// Solidity: function totalBorrowed() view returns(uint256)
func (_PoolService *PoolServiceCaller) TotalBorrowed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PoolService.contract.Call(opts, &out, "totalBorrowed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBorrowed is a free data retrieval call binding the contract method 0x4c19386c.
//
// Solidity: function totalBorrowed() view returns(uint256)
func (_PoolService *PoolServiceSession) TotalBorrowed() (*big.Int, error) {
	return _PoolService.Contract.TotalBorrowed(&_PoolService.CallOpts)
}

// TotalBorrowed is a free data retrieval call binding the contract method 0x4c19386c.
//
// Solidity: function totalBorrowed() view returns(uint256)
func (_PoolService *PoolServiceCallerSession) TotalBorrowed() (*big.Int, error) {
	return _PoolService.Contract.TotalBorrowed(&_PoolService.CallOpts)
}

// TreasuryAddress is a free data retrieval call binding the contract method 0xc5f956af.
//
// Solidity: function treasuryAddress() view returns(address)
func (_PoolService *PoolServiceCaller) TreasuryAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoolService.contract.Call(opts, &out, "treasuryAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TreasuryAddress is a free data retrieval call binding the contract method 0xc5f956af.
//
// Solidity: function treasuryAddress() view returns(address)
func (_PoolService *PoolServiceSession) TreasuryAddress() (common.Address, error) {
	return _PoolService.Contract.TreasuryAddress(&_PoolService.CallOpts)
}

// TreasuryAddress is a free data retrieval call binding the contract method 0xc5f956af.
//
// Solidity: function treasuryAddress() view returns(address)
func (_PoolService *PoolServiceCallerSession) TreasuryAddress() (common.Address, error) {
	return _PoolService.Contract.TreasuryAddress(&_PoolService.CallOpts)
}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_PoolService *PoolServiceCaller) UnderlyingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoolService.contract.Call(opts, &out, "underlyingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_PoolService *PoolServiceSession) UnderlyingToken() (common.Address, error) {
	return _PoolService.Contract.UnderlyingToken(&_PoolService.CallOpts)
}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_PoolService *PoolServiceCallerSession) UnderlyingToken() (common.Address, error) {
	return _PoolService.Contract.UnderlyingToken(&_PoolService.CallOpts)
}

// WithdrawFee is a free data retrieval call binding the contract method 0xe941fa78.
//
// Solidity: function withdrawFee() view returns(uint256)
func (_PoolService *PoolServiceCaller) WithdrawFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PoolService.contract.Call(opts, &out, "withdrawFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawFee is a free data retrieval call binding the contract method 0xe941fa78.
//
// Solidity: function withdrawFee() view returns(uint256)
func (_PoolService *PoolServiceSession) WithdrawFee() (*big.Int, error) {
	return _PoolService.Contract.WithdrawFee(&_PoolService.CallOpts)
}

// WithdrawFee is a free data retrieval call binding the contract method 0xe941fa78.
//
// Solidity: function withdrawFee() view returns(uint256)
func (_PoolService *PoolServiceCallerSession) WithdrawFee() (*big.Int, error) {
	return _PoolService.Contract.WithdrawFee(&_PoolService.CallOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x9aa5d462.
//
// Solidity: function addLiquidity(uint256 amount, address onBehalfOf, uint256 referralCode) returns()
func (_PoolService *PoolServiceTransactor) AddLiquidity(opts *bind.TransactOpts, amount *big.Int, onBehalfOf common.Address, referralCode *big.Int) (*types.Transaction, error) {
	return _PoolService.contract.Transact(opts, "addLiquidity", amount, onBehalfOf, referralCode)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x9aa5d462.
//
// Solidity: function addLiquidity(uint256 amount, address onBehalfOf, uint256 referralCode) returns()
func (_PoolService *PoolServiceSession) AddLiquidity(amount *big.Int, onBehalfOf common.Address, referralCode *big.Int) (*types.Transaction, error) {
	return _PoolService.Contract.AddLiquidity(&_PoolService.TransactOpts, amount, onBehalfOf, referralCode)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x9aa5d462.
//
// Solidity: function addLiquidity(uint256 amount, address onBehalfOf, uint256 referralCode) returns()
func (_PoolService *PoolServiceTransactorSession) AddLiquidity(amount *big.Int, onBehalfOf common.Address, referralCode *big.Int) (*types.Transaction, error) {
	return _PoolService.Contract.AddLiquidity(&_PoolService.TransactOpts, amount, onBehalfOf, referralCode)
}

// ConnectCreditManager is a paid mutator transaction binding the contract method 0xcf33d955.
//
// Solidity: function connectCreditManager(address _creditManager) returns()
func (_PoolService *PoolServiceTransactor) ConnectCreditManager(opts *bind.TransactOpts, _creditManager common.Address) (*types.Transaction, error) {
	return _PoolService.contract.Transact(opts, "connectCreditManager", _creditManager)
}

// ConnectCreditManager is a paid mutator transaction binding the contract method 0xcf33d955.
//
// Solidity: function connectCreditManager(address _creditManager) returns()
func (_PoolService *PoolServiceSession) ConnectCreditManager(_creditManager common.Address) (*types.Transaction, error) {
	return _PoolService.Contract.ConnectCreditManager(&_PoolService.TransactOpts, _creditManager)
}

// ConnectCreditManager is a paid mutator transaction binding the contract method 0xcf33d955.
//
// Solidity: function connectCreditManager(address _creditManager) returns()
func (_PoolService *PoolServiceTransactorSession) ConnectCreditManager(_creditManager common.Address) (*types.Transaction, error) {
	return _PoolService.Contract.ConnectCreditManager(&_PoolService.TransactOpts, _creditManager)
}

// ForbidCreditManagerToBorrow is a paid mutator transaction binding the contract method 0x078c4781.
//
// Solidity: function forbidCreditManagerToBorrow(address _creditManager) returns()
func (_PoolService *PoolServiceTransactor) ForbidCreditManagerToBorrow(opts *bind.TransactOpts, _creditManager common.Address) (*types.Transaction, error) {
	return _PoolService.contract.Transact(opts, "forbidCreditManagerToBorrow", _creditManager)
}

// ForbidCreditManagerToBorrow is a paid mutator transaction binding the contract method 0x078c4781.
//
// Solidity: function forbidCreditManagerToBorrow(address _creditManager) returns()
func (_PoolService *PoolServiceSession) ForbidCreditManagerToBorrow(_creditManager common.Address) (*types.Transaction, error) {
	return _PoolService.Contract.ForbidCreditManagerToBorrow(&_PoolService.TransactOpts, _creditManager)
}

// ForbidCreditManagerToBorrow is a paid mutator transaction binding the contract method 0x078c4781.
//
// Solidity: function forbidCreditManagerToBorrow(address _creditManager) returns()
func (_PoolService *PoolServiceTransactorSession) ForbidCreditManagerToBorrow(_creditManager common.Address) (*types.Transaction, error) {
	return _PoolService.Contract.ForbidCreditManagerToBorrow(&_PoolService.TransactOpts, _creditManager)
}

// LendCreditAccount is a paid mutator transaction binding the contract method 0xbf28068b.
//
// Solidity: function lendCreditAccount(uint256 borrowedAmount, address creditAccount) returns()
func (_PoolService *PoolServiceTransactor) LendCreditAccount(opts *bind.TransactOpts, borrowedAmount *big.Int, creditAccount common.Address) (*types.Transaction, error) {
	return _PoolService.contract.Transact(opts, "lendCreditAccount", borrowedAmount, creditAccount)
}

// LendCreditAccount is a paid mutator transaction binding the contract method 0xbf28068b.
//
// Solidity: function lendCreditAccount(uint256 borrowedAmount, address creditAccount) returns()
func (_PoolService *PoolServiceSession) LendCreditAccount(borrowedAmount *big.Int, creditAccount common.Address) (*types.Transaction, error) {
	return _PoolService.Contract.LendCreditAccount(&_PoolService.TransactOpts, borrowedAmount, creditAccount)
}

// LendCreditAccount is a paid mutator transaction binding the contract method 0xbf28068b.
//
// Solidity: function lendCreditAccount(uint256 borrowedAmount, address creditAccount) returns()
func (_PoolService *PoolServiceTransactorSession) LendCreditAccount(borrowedAmount *big.Int, creditAccount common.Address) (*types.Transaction, error) {
	return _PoolService.Contract.LendCreditAccount(&_PoolService.TransactOpts, borrowedAmount, creditAccount)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PoolService *PoolServiceTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolService.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PoolService *PoolServiceSession) Pause() (*types.Transaction, error) {
	return _PoolService.Contract.Pause(&_PoolService.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PoolService *PoolServiceTransactorSession) Pause() (*types.Transaction, error) {
	return _PoolService.Contract.Pause(&_PoolService.TransactOpts)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x05fe138b.
//
// Solidity: function removeLiquidity(uint256 amount, address to) returns(uint256)
func (_PoolService *PoolServiceTransactor) RemoveLiquidity(opts *bind.TransactOpts, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _PoolService.contract.Transact(opts, "removeLiquidity", amount, to)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x05fe138b.
//
// Solidity: function removeLiquidity(uint256 amount, address to) returns(uint256)
func (_PoolService *PoolServiceSession) RemoveLiquidity(amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _PoolService.Contract.RemoveLiquidity(&_PoolService.TransactOpts, amount, to)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x05fe138b.
//
// Solidity: function removeLiquidity(uint256 amount, address to) returns(uint256)
func (_PoolService *PoolServiceTransactorSession) RemoveLiquidity(amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _PoolService.Contract.RemoveLiquidity(&_PoolService.TransactOpts, amount, to)
}

// RepayCreditAccount is a paid mutator transaction binding the contract method 0xca9505e4.
//
// Solidity: function repayCreditAccount(uint256 borrowedAmount, uint256 profit, uint256 loss) returns()
func (_PoolService *PoolServiceTransactor) RepayCreditAccount(opts *bind.TransactOpts, borrowedAmount *big.Int, profit *big.Int, loss *big.Int) (*types.Transaction, error) {
	return _PoolService.contract.Transact(opts, "repayCreditAccount", borrowedAmount, profit, loss)
}

// RepayCreditAccount is a paid mutator transaction binding the contract method 0xca9505e4.
//
// Solidity: function repayCreditAccount(uint256 borrowedAmount, uint256 profit, uint256 loss) returns()
func (_PoolService *PoolServiceSession) RepayCreditAccount(borrowedAmount *big.Int, profit *big.Int, loss *big.Int) (*types.Transaction, error) {
	return _PoolService.Contract.RepayCreditAccount(&_PoolService.TransactOpts, borrowedAmount, profit, loss)
}

// RepayCreditAccount is a paid mutator transaction binding the contract method 0xca9505e4.
//
// Solidity: function repayCreditAccount(uint256 borrowedAmount, uint256 profit, uint256 loss) returns()
func (_PoolService *PoolServiceTransactorSession) RepayCreditAccount(borrowedAmount *big.Int, profit *big.Int, loss *big.Int) (*types.Transaction, error) {
	return _PoolService.Contract.RepayCreditAccount(&_PoolService.TransactOpts, borrowedAmount, profit, loss)
}

// SetExpectedLiquidityLimit is a paid mutator transaction binding the contract method 0xbb04b193.
//
// Solidity: function setExpectedLiquidityLimit(uint256 newLimit) returns()
func (_PoolService *PoolServiceTransactor) SetExpectedLiquidityLimit(opts *bind.TransactOpts, newLimit *big.Int) (*types.Transaction, error) {
	return _PoolService.contract.Transact(opts, "setExpectedLiquidityLimit", newLimit)
}

// SetExpectedLiquidityLimit is a paid mutator transaction binding the contract method 0xbb04b193.
//
// Solidity: function setExpectedLiquidityLimit(uint256 newLimit) returns()
func (_PoolService *PoolServiceSession) SetExpectedLiquidityLimit(newLimit *big.Int) (*types.Transaction, error) {
	return _PoolService.Contract.SetExpectedLiquidityLimit(&_PoolService.TransactOpts, newLimit)
}

// SetExpectedLiquidityLimit is a paid mutator transaction binding the contract method 0xbb04b193.
//
// Solidity: function setExpectedLiquidityLimit(uint256 newLimit) returns()
func (_PoolService *PoolServiceTransactorSession) SetExpectedLiquidityLimit(newLimit *big.Int) (*types.Transaction, error) {
	return _PoolService.Contract.SetExpectedLiquidityLimit(&_PoolService.TransactOpts, newLimit)
}

// SetWithdrawFee is a paid mutator transaction binding the contract method 0xb6ac642a.
//
// Solidity: function setWithdrawFee(uint256 fee) returns()
func (_PoolService *PoolServiceTransactor) SetWithdrawFee(opts *bind.TransactOpts, fee *big.Int) (*types.Transaction, error) {
	return _PoolService.contract.Transact(opts, "setWithdrawFee", fee)
}

// SetWithdrawFee is a paid mutator transaction binding the contract method 0xb6ac642a.
//
// Solidity: function setWithdrawFee(uint256 fee) returns()
func (_PoolService *PoolServiceSession) SetWithdrawFee(fee *big.Int) (*types.Transaction, error) {
	return _PoolService.Contract.SetWithdrawFee(&_PoolService.TransactOpts, fee)
}

// SetWithdrawFee is a paid mutator transaction binding the contract method 0xb6ac642a.
//
// Solidity: function setWithdrawFee(uint256 fee) returns()
func (_PoolService *PoolServiceTransactorSession) SetWithdrawFee(fee *big.Int) (*types.Transaction, error) {
	return _PoolService.Contract.SetWithdrawFee(&_PoolService.TransactOpts, fee)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PoolService *PoolServiceTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolService.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PoolService *PoolServiceSession) Unpause() (*types.Transaction, error) {
	return _PoolService.Contract.Unpause(&_PoolService.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PoolService *PoolServiceTransactorSession) Unpause() (*types.Transaction, error) {
	return _PoolService.Contract.Unpause(&_PoolService.TransactOpts)
}

// UpdateInterestRateModel is a paid mutator transaction binding the contract method 0x5664cacf.
//
// Solidity: function updateInterestRateModel(address _interestRateModel) returns()
func (_PoolService *PoolServiceTransactor) UpdateInterestRateModel(opts *bind.TransactOpts, _interestRateModel common.Address) (*types.Transaction, error) {
	return _PoolService.contract.Transact(opts, "updateInterestRateModel", _interestRateModel)
}

// UpdateInterestRateModel is a paid mutator transaction binding the contract method 0x5664cacf.
//
// Solidity: function updateInterestRateModel(address _interestRateModel) returns()
func (_PoolService *PoolServiceSession) UpdateInterestRateModel(_interestRateModel common.Address) (*types.Transaction, error) {
	return _PoolService.Contract.UpdateInterestRateModel(&_PoolService.TransactOpts, _interestRateModel)
}

// UpdateInterestRateModel is a paid mutator transaction binding the contract method 0x5664cacf.
//
// Solidity: function updateInterestRateModel(address _interestRateModel) returns()
func (_PoolService *PoolServiceTransactorSession) UpdateInterestRateModel(_interestRateModel common.Address) (*types.Transaction, error) {
	return _PoolService.Contract.UpdateInterestRateModel(&_PoolService.TransactOpts, _interestRateModel)
}

// PoolServiceAddLiquidityIterator is returned from FilterAddLiquidity and is used to iterate over the raw logs and unpacked data for AddLiquidity events raised by the PoolService contract.
type PoolServiceAddLiquidityIterator struct {
	Event *PoolServiceAddLiquidity // Event containing the contract specifics and raw log

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
func (it *PoolServiceAddLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolServiceAddLiquidity)
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
		it.Event = new(PoolServiceAddLiquidity)
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
func (it *PoolServiceAddLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolServiceAddLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolServiceAddLiquidity represents a AddLiquidity event raised by the PoolService contract.
type PoolServiceAddLiquidity struct {
	Sender       common.Address
	OnBehalfOf   common.Address
	Amount       *big.Int
	ReferralCode *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAddLiquidity is a free log retrieval operation binding the contract event 0xd2491a9b4fe81a7cd4511e8b7b7743951b061dad5bed7da8a7795b080ee08c7e.
//
// Solidity: event AddLiquidity(address indexed sender, address indexed onBehalfOf, uint256 amount, uint256 referralCode)
func (_PoolService *PoolServiceFilterer) FilterAddLiquidity(opts *bind.FilterOpts, sender []common.Address, onBehalfOf []common.Address) (*PoolServiceAddLiquidityIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	logs, sub, err := _PoolService.contract.FilterLogs(opts, "AddLiquidity", senderRule, onBehalfOfRule)
	if err != nil {
		return nil, err
	}
	return &PoolServiceAddLiquidityIterator{contract: _PoolService.contract, event: "AddLiquidity", logs: logs, sub: sub}, nil
}

// WatchAddLiquidity is a free log subscription operation binding the contract event 0xd2491a9b4fe81a7cd4511e8b7b7743951b061dad5bed7da8a7795b080ee08c7e.
//
// Solidity: event AddLiquidity(address indexed sender, address indexed onBehalfOf, uint256 amount, uint256 referralCode)
func (_PoolService *PoolServiceFilterer) WatchAddLiquidity(opts *bind.WatchOpts, sink chan<- *PoolServiceAddLiquidity, sender []common.Address, onBehalfOf []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	logs, sub, err := _PoolService.contract.WatchLogs(opts, "AddLiquidity", senderRule, onBehalfOfRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolServiceAddLiquidity)
				if err := _PoolService.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
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
func (_PoolService *PoolServiceFilterer) ParseAddLiquidity(log types.Log) (*PoolServiceAddLiquidity, error) {
	event := new(PoolServiceAddLiquidity)
	if err := _PoolService.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolServiceBorrowIterator is returned from FilterBorrow and is used to iterate over the raw logs and unpacked data for Borrow events raised by the PoolService contract.
type PoolServiceBorrowIterator struct {
	Event *PoolServiceBorrow // Event containing the contract specifics and raw log

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
func (it *PoolServiceBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolServiceBorrow)
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
		it.Event = new(PoolServiceBorrow)
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
func (it *PoolServiceBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolServiceBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolServiceBorrow represents a Borrow event raised by the PoolService contract.
type PoolServiceBorrow struct {
	CreditManager common.Address
	CreditAccount common.Address
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBorrow is a free log retrieval operation binding the contract event 0x312a5e5e1079f5dda4e95dbbd0b908b291fd5b992ef22073643ab691572c5b52.
//
// Solidity: event Borrow(address indexed creditManager, address indexed creditAccount, uint256 amount)
func (_PoolService *PoolServiceFilterer) FilterBorrow(opts *bind.FilterOpts, creditManager []common.Address, creditAccount []common.Address) (*PoolServiceBorrowIterator, error) {

	var creditManagerRule []interface{}
	for _, creditManagerItem := range creditManager {
		creditManagerRule = append(creditManagerRule, creditManagerItem)
	}
	var creditAccountRule []interface{}
	for _, creditAccountItem := range creditAccount {
		creditAccountRule = append(creditAccountRule, creditAccountItem)
	}

	logs, sub, err := _PoolService.contract.FilterLogs(opts, "Borrow", creditManagerRule, creditAccountRule)
	if err != nil {
		return nil, err
	}
	return &PoolServiceBorrowIterator{contract: _PoolService.contract, event: "Borrow", logs: logs, sub: sub}, nil
}

// WatchBorrow is a free log subscription operation binding the contract event 0x312a5e5e1079f5dda4e95dbbd0b908b291fd5b992ef22073643ab691572c5b52.
//
// Solidity: event Borrow(address indexed creditManager, address indexed creditAccount, uint256 amount)
func (_PoolService *PoolServiceFilterer) WatchBorrow(opts *bind.WatchOpts, sink chan<- *PoolServiceBorrow, creditManager []common.Address, creditAccount []common.Address) (event.Subscription, error) {

	var creditManagerRule []interface{}
	for _, creditManagerItem := range creditManager {
		creditManagerRule = append(creditManagerRule, creditManagerItem)
	}
	var creditAccountRule []interface{}
	for _, creditAccountItem := range creditAccount {
		creditAccountRule = append(creditAccountRule, creditAccountItem)
	}

	logs, sub, err := _PoolService.contract.WatchLogs(opts, "Borrow", creditManagerRule, creditAccountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolServiceBorrow)
				if err := _PoolService.contract.UnpackLog(event, "Borrow", log); err != nil {
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
func (_PoolService *PoolServiceFilterer) ParseBorrow(log types.Log) (*PoolServiceBorrow, error) {
	event := new(PoolServiceBorrow)
	if err := _PoolService.contract.UnpackLog(event, "Borrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolServiceBorrowForbiddenIterator is returned from FilterBorrowForbidden and is used to iterate over the raw logs and unpacked data for BorrowForbidden events raised by the PoolService contract.
type PoolServiceBorrowForbiddenIterator struct {
	Event *PoolServiceBorrowForbidden // Event containing the contract specifics and raw log

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
func (it *PoolServiceBorrowForbiddenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolServiceBorrowForbidden)
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
		it.Event = new(PoolServiceBorrowForbidden)
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
func (it *PoolServiceBorrowForbiddenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolServiceBorrowForbiddenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolServiceBorrowForbidden represents a BorrowForbidden event raised by the PoolService contract.
type PoolServiceBorrowForbidden struct {
	CreditManager common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBorrowForbidden is a free log retrieval operation binding the contract event 0x9181736fce85d2d4cca2e4406f10679302ae5c387180fdb62963af3cd9a24fd6.
//
// Solidity: event BorrowForbidden(address indexed creditManager)
func (_PoolService *PoolServiceFilterer) FilterBorrowForbidden(opts *bind.FilterOpts, creditManager []common.Address) (*PoolServiceBorrowForbiddenIterator, error) {

	var creditManagerRule []interface{}
	for _, creditManagerItem := range creditManager {
		creditManagerRule = append(creditManagerRule, creditManagerItem)
	}

	logs, sub, err := _PoolService.contract.FilterLogs(opts, "BorrowForbidden", creditManagerRule)
	if err != nil {
		return nil, err
	}
	return &PoolServiceBorrowForbiddenIterator{contract: _PoolService.contract, event: "BorrowForbidden", logs: logs, sub: sub}, nil
}

// WatchBorrowForbidden is a free log subscription operation binding the contract event 0x9181736fce85d2d4cca2e4406f10679302ae5c387180fdb62963af3cd9a24fd6.
//
// Solidity: event BorrowForbidden(address indexed creditManager)
func (_PoolService *PoolServiceFilterer) WatchBorrowForbidden(opts *bind.WatchOpts, sink chan<- *PoolServiceBorrowForbidden, creditManager []common.Address) (event.Subscription, error) {

	var creditManagerRule []interface{}
	for _, creditManagerItem := range creditManager {
		creditManagerRule = append(creditManagerRule, creditManagerItem)
	}

	logs, sub, err := _PoolService.contract.WatchLogs(opts, "BorrowForbidden", creditManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolServiceBorrowForbidden)
				if err := _PoolService.contract.UnpackLog(event, "BorrowForbidden", log); err != nil {
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
func (_PoolService *PoolServiceFilterer) ParseBorrowForbidden(log types.Log) (*PoolServiceBorrowForbidden, error) {
	event := new(PoolServiceBorrowForbidden)
	if err := _PoolService.contract.UnpackLog(event, "BorrowForbidden", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolServiceNewCreditManagerConnectedIterator is returned from FilterNewCreditManagerConnected and is used to iterate over the raw logs and unpacked data for NewCreditManagerConnected events raised by the PoolService contract.
type PoolServiceNewCreditManagerConnectedIterator struct {
	Event *PoolServiceNewCreditManagerConnected // Event containing the contract specifics and raw log

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
func (it *PoolServiceNewCreditManagerConnectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolServiceNewCreditManagerConnected)
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
		it.Event = new(PoolServiceNewCreditManagerConnected)
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
func (it *PoolServiceNewCreditManagerConnectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolServiceNewCreditManagerConnectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolServiceNewCreditManagerConnected represents a NewCreditManagerConnected event raised by the PoolService contract.
type PoolServiceNewCreditManagerConnected struct {
	CreditManager common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNewCreditManagerConnected is a free log retrieval operation binding the contract event 0xe076020e7eac3915d33aec40c24f95e73eb6c9921ff89747d50aa8fd934d2c01.
//
// Solidity: event NewCreditManagerConnected(address indexed creditManager)
func (_PoolService *PoolServiceFilterer) FilterNewCreditManagerConnected(opts *bind.FilterOpts, creditManager []common.Address) (*PoolServiceNewCreditManagerConnectedIterator, error) {

	var creditManagerRule []interface{}
	for _, creditManagerItem := range creditManager {
		creditManagerRule = append(creditManagerRule, creditManagerItem)
	}

	logs, sub, err := _PoolService.contract.FilterLogs(opts, "NewCreditManagerConnected", creditManagerRule)
	if err != nil {
		return nil, err
	}
	return &PoolServiceNewCreditManagerConnectedIterator{contract: _PoolService.contract, event: "NewCreditManagerConnected", logs: logs, sub: sub}, nil
}

// WatchNewCreditManagerConnected is a free log subscription operation binding the contract event 0xe076020e7eac3915d33aec40c24f95e73eb6c9921ff89747d50aa8fd934d2c01.
//
// Solidity: event NewCreditManagerConnected(address indexed creditManager)
func (_PoolService *PoolServiceFilterer) WatchNewCreditManagerConnected(opts *bind.WatchOpts, sink chan<- *PoolServiceNewCreditManagerConnected, creditManager []common.Address) (event.Subscription, error) {

	var creditManagerRule []interface{}
	for _, creditManagerItem := range creditManager {
		creditManagerRule = append(creditManagerRule, creditManagerItem)
	}

	logs, sub, err := _PoolService.contract.WatchLogs(opts, "NewCreditManagerConnected", creditManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolServiceNewCreditManagerConnected)
				if err := _PoolService.contract.UnpackLog(event, "NewCreditManagerConnected", log); err != nil {
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
func (_PoolService *PoolServiceFilterer) ParseNewCreditManagerConnected(log types.Log) (*PoolServiceNewCreditManagerConnected, error) {
	event := new(PoolServiceNewCreditManagerConnected)
	if err := _PoolService.contract.UnpackLog(event, "NewCreditManagerConnected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolServiceNewExpectedLiquidityLimitIterator is returned from FilterNewExpectedLiquidityLimit and is used to iterate over the raw logs and unpacked data for NewExpectedLiquidityLimit events raised by the PoolService contract.
type PoolServiceNewExpectedLiquidityLimitIterator struct {
	Event *PoolServiceNewExpectedLiquidityLimit // Event containing the contract specifics and raw log

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
func (it *PoolServiceNewExpectedLiquidityLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolServiceNewExpectedLiquidityLimit)
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
		it.Event = new(PoolServiceNewExpectedLiquidityLimit)
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
func (it *PoolServiceNewExpectedLiquidityLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolServiceNewExpectedLiquidityLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolServiceNewExpectedLiquidityLimit represents a NewExpectedLiquidityLimit event raised by the PoolService contract.
type PoolServiceNewExpectedLiquidityLimit struct {
	NewLimit *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewExpectedLiquidityLimit is a free log retrieval operation binding the contract event 0xd7a183c9fe85b604c25d54bd676e0866f6c13bcca9fb9b0850213de118fdc99c.
//
// Solidity: event NewExpectedLiquidityLimit(uint256 newLimit)
func (_PoolService *PoolServiceFilterer) FilterNewExpectedLiquidityLimit(opts *bind.FilterOpts) (*PoolServiceNewExpectedLiquidityLimitIterator, error) {

	logs, sub, err := _PoolService.contract.FilterLogs(opts, "NewExpectedLiquidityLimit")
	if err != nil {
		return nil, err
	}
	return &PoolServiceNewExpectedLiquidityLimitIterator{contract: _PoolService.contract, event: "NewExpectedLiquidityLimit", logs: logs, sub: sub}, nil
}

// WatchNewExpectedLiquidityLimit is a free log subscription operation binding the contract event 0xd7a183c9fe85b604c25d54bd676e0866f6c13bcca9fb9b0850213de118fdc99c.
//
// Solidity: event NewExpectedLiquidityLimit(uint256 newLimit)
func (_PoolService *PoolServiceFilterer) WatchNewExpectedLiquidityLimit(opts *bind.WatchOpts, sink chan<- *PoolServiceNewExpectedLiquidityLimit) (event.Subscription, error) {

	logs, sub, err := _PoolService.contract.WatchLogs(opts, "NewExpectedLiquidityLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolServiceNewExpectedLiquidityLimit)
				if err := _PoolService.contract.UnpackLog(event, "NewExpectedLiquidityLimit", log); err != nil {
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
func (_PoolService *PoolServiceFilterer) ParseNewExpectedLiquidityLimit(log types.Log) (*PoolServiceNewExpectedLiquidityLimit, error) {
	event := new(PoolServiceNewExpectedLiquidityLimit)
	if err := _PoolService.contract.UnpackLog(event, "NewExpectedLiquidityLimit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolServiceNewInterestRateModelIterator is returned from FilterNewInterestRateModel and is used to iterate over the raw logs and unpacked data for NewInterestRateModel events raised by the PoolService contract.
type PoolServiceNewInterestRateModelIterator struct {
	Event *PoolServiceNewInterestRateModel // Event containing the contract specifics and raw log

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
func (it *PoolServiceNewInterestRateModelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolServiceNewInterestRateModel)
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
		it.Event = new(PoolServiceNewInterestRateModel)
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
func (it *PoolServiceNewInterestRateModelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolServiceNewInterestRateModelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolServiceNewInterestRateModel represents a NewInterestRateModel event raised by the PoolService contract.
type PoolServiceNewInterestRateModel struct {
	NewInterestRateModel common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterNewInterestRateModel is a free log retrieval operation binding the contract event 0x0ec6cb7631d36954a05ffd646135bfd9995c71e7fa36d26abb1ad9f24a040ea1.
//
// Solidity: event NewInterestRateModel(address indexed newInterestRateModel)
func (_PoolService *PoolServiceFilterer) FilterNewInterestRateModel(opts *bind.FilterOpts, newInterestRateModel []common.Address) (*PoolServiceNewInterestRateModelIterator, error) {

	var newInterestRateModelRule []interface{}
	for _, newInterestRateModelItem := range newInterestRateModel {
		newInterestRateModelRule = append(newInterestRateModelRule, newInterestRateModelItem)
	}

	logs, sub, err := _PoolService.contract.FilterLogs(opts, "NewInterestRateModel", newInterestRateModelRule)
	if err != nil {
		return nil, err
	}
	return &PoolServiceNewInterestRateModelIterator{contract: _PoolService.contract, event: "NewInterestRateModel", logs: logs, sub: sub}, nil
}

// WatchNewInterestRateModel is a free log subscription operation binding the contract event 0x0ec6cb7631d36954a05ffd646135bfd9995c71e7fa36d26abb1ad9f24a040ea1.
//
// Solidity: event NewInterestRateModel(address indexed newInterestRateModel)
func (_PoolService *PoolServiceFilterer) WatchNewInterestRateModel(opts *bind.WatchOpts, sink chan<- *PoolServiceNewInterestRateModel, newInterestRateModel []common.Address) (event.Subscription, error) {

	var newInterestRateModelRule []interface{}
	for _, newInterestRateModelItem := range newInterestRateModel {
		newInterestRateModelRule = append(newInterestRateModelRule, newInterestRateModelItem)
	}

	logs, sub, err := _PoolService.contract.WatchLogs(opts, "NewInterestRateModel", newInterestRateModelRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolServiceNewInterestRateModel)
				if err := _PoolService.contract.UnpackLog(event, "NewInterestRateModel", log); err != nil {
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
func (_PoolService *PoolServiceFilterer) ParseNewInterestRateModel(log types.Log) (*PoolServiceNewInterestRateModel, error) {
	event := new(PoolServiceNewInterestRateModel)
	if err := _PoolService.contract.UnpackLog(event, "NewInterestRateModel", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolServiceNewWithdrawFeeIterator is returned from FilterNewWithdrawFee and is used to iterate over the raw logs and unpacked data for NewWithdrawFee events raised by the PoolService contract.
type PoolServiceNewWithdrawFeeIterator struct {
	Event *PoolServiceNewWithdrawFee // Event containing the contract specifics and raw log

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
func (it *PoolServiceNewWithdrawFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolServiceNewWithdrawFee)
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
		it.Event = new(PoolServiceNewWithdrawFee)
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
func (it *PoolServiceNewWithdrawFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolServiceNewWithdrawFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolServiceNewWithdrawFee represents a NewWithdrawFee event raised by the PoolService contract.
type PoolServiceNewWithdrawFee struct {
	Fee *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterNewWithdrawFee is a free log retrieval operation binding the contract event 0xd5fe46099fa396290a7f57e36c3c3c8774e2562c18ed5d1dcc0fa75071e03f1d.
//
// Solidity: event NewWithdrawFee(uint256 fee)
func (_PoolService *PoolServiceFilterer) FilterNewWithdrawFee(opts *bind.FilterOpts) (*PoolServiceNewWithdrawFeeIterator, error) {

	logs, sub, err := _PoolService.contract.FilterLogs(opts, "NewWithdrawFee")
	if err != nil {
		return nil, err
	}
	return &PoolServiceNewWithdrawFeeIterator{contract: _PoolService.contract, event: "NewWithdrawFee", logs: logs, sub: sub}, nil
}

// WatchNewWithdrawFee is a free log subscription operation binding the contract event 0xd5fe46099fa396290a7f57e36c3c3c8774e2562c18ed5d1dcc0fa75071e03f1d.
//
// Solidity: event NewWithdrawFee(uint256 fee)
func (_PoolService *PoolServiceFilterer) WatchNewWithdrawFee(opts *bind.WatchOpts, sink chan<- *PoolServiceNewWithdrawFee) (event.Subscription, error) {

	logs, sub, err := _PoolService.contract.WatchLogs(opts, "NewWithdrawFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolServiceNewWithdrawFee)
				if err := _PoolService.contract.UnpackLog(event, "NewWithdrawFee", log); err != nil {
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
func (_PoolService *PoolServiceFilterer) ParseNewWithdrawFee(log types.Log) (*PoolServiceNewWithdrawFee, error) {
	event := new(PoolServiceNewWithdrawFee)
	if err := _PoolService.contract.UnpackLog(event, "NewWithdrawFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolServicePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the PoolService contract.
type PoolServicePausedIterator struct {
	Event *PoolServicePaused // Event containing the contract specifics and raw log

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
func (it *PoolServicePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolServicePaused)
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
		it.Event = new(PoolServicePaused)
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
func (it *PoolServicePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolServicePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolServicePaused represents a Paused event raised by the PoolService contract.
type PoolServicePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PoolService *PoolServiceFilterer) FilterPaused(opts *bind.FilterOpts) (*PoolServicePausedIterator, error) {

	logs, sub, err := _PoolService.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &PoolServicePausedIterator{contract: _PoolService.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PoolService *PoolServiceFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PoolServicePaused) (event.Subscription, error) {

	logs, sub, err := _PoolService.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolServicePaused)
				if err := _PoolService.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_PoolService *PoolServiceFilterer) ParsePaused(log types.Log) (*PoolServicePaused, error) {
	event := new(PoolServicePaused)
	if err := _PoolService.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolServiceRemoveLiquidityIterator is returned from FilterRemoveLiquidity and is used to iterate over the raw logs and unpacked data for RemoveLiquidity events raised by the PoolService contract.
type PoolServiceRemoveLiquidityIterator struct {
	Event *PoolServiceRemoveLiquidity // Event containing the contract specifics and raw log

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
func (it *PoolServiceRemoveLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolServiceRemoveLiquidity)
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
		it.Event = new(PoolServiceRemoveLiquidity)
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
func (it *PoolServiceRemoveLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolServiceRemoveLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolServiceRemoveLiquidity represents a RemoveLiquidity event raised by the PoolService contract.
type PoolServiceRemoveLiquidity struct {
	Sender common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidity is a free log retrieval operation binding the contract event 0xd8ae9b9ba89e637bcb66a69ac91e8f688018e81d6f92c57e02226425c8efbdf6.
//
// Solidity: event RemoveLiquidity(address indexed sender, address indexed to, uint256 amount)
func (_PoolService *PoolServiceFilterer) FilterRemoveLiquidity(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*PoolServiceRemoveLiquidityIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PoolService.contract.FilterLogs(opts, "RemoveLiquidity", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PoolServiceRemoveLiquidityIterator{contract: _PoolService.contract, event: "RemoveLiquidity", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidity is a free log subscription operation binding the contract event 0xd8ae9b9ba89e637bcb66a69ac91e8f688018e81d6f92c57e02226425c8efbdf6.
//
// Solidity: event RemoveLiquidity(address indexed sender, address indexed to, uint256 amount)
func (_PoolService *PoolServiceFilterer) WatchRemoveLiquidity(opts *bind.WatchOpts, sink chan<- *PoolServiceRemoveLiquidity, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PoolService.contract.WatchLogs(opts, "RemoveLiquidity", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolServiceRemoveLiquidity)
				if err := _PoolService.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
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
func (_PoolService *PoolServiceFilterer) ParseRemoveLiquidity(log types.Log) (*PoolServiceRemoveLiquidity, error) {
	event := new(PoolServiceRemoveLiquidity)
	if err := _PoolService.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolServiceRepayIterator is returned from FilterRepay and is used to iterate over the raw logs and unpacked data for Repay events raised by the PoolService contract.
type PoolServiceRepayIterator struct {
	Event *PoolServiceRepay // Event containing the contract specifics and raw log

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
func (it *PoolServiceRepayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolServiceRepay)
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
		it.Event = new(PoolServiceRepay)
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
func (it *PoolServiceRepayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolServiceRepayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolServiceRepay represents a Repay event raised by the PoolService contract.
type PoolServiceRepay struct {
	CreditManager  common.Address
	BorrowedAmount *big.Int
	Profit         *big.Int
	Loss           *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRepay is a free log retrieval operation binding the contract event 0x2fe77b1c99aca6b022b8efc6e3e8dd1b48b30748709339b65c50ef3263443e09.
//
// Solidity: event Repay(address indexed creditManager, uint256 borrowedAmount, uint256 profit, uint256 loss)
func (_PoolService *PoolServiceFilterer) FilterRepay(opts *bind.FilterOpts, creditManager []common.Address) (*PoolServiceRepayIterator, error) {

	var creditManagerRule []interface{}
	for _, creditManagerItem := range creditManager {
		creditManagerRule = append(creditManagerRule, creditManagerItem)
	}

	logs, sub, err := _PoolService.contract.FilterLogs(opts, "Repay", creditManagerRule)
	if err != nil {
		return nil, err
	}
	return &PoolServiceRepayIterator{contract: _PoolService.contract, event: "Repay", logs: logs, sub: sub}, nil
}

// WatchRepay is a free log subscription operation binding the contract event 0x2fe77b1c99aca6b022b8efc6e3e8dd1b48b30748709339b65c50ef3263443e09.
//
// Solidity: event Repay(address indexed creditManager, uint256 borrowedAmount, uint256 profit, uint256 loss)
func (_PoolService *PoolServiceFilterer) WatchRepay(opts *bind.WatchOpts, sink chan<- *PoolServiceRepay, creditManager []common.Address) (event.Subscription, error) {

	var creditManagerRule []interface{}
	for _, creditManagerItem := range creditManager {
		creditManagerRule = append(creditManagerRule, creditManagerItem)
	}

	logs, sub, err := _PoolService.contract.WatchLogs(opts, "Repay", creditManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolServiceRepay)
				if err := _PoolService.contract.UnpackLog(event, "Repay", log); err != nil {
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
func (_PoolService *PoolServiceFilterer) ParseRepay(log types.Log) (*PoolServiceRepay, error) {
	event := new(PoolServiceRepay)
	if err := _PoolService.contract.UnpackLog(event, "Repay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolServiceUncoveredLossIterator is returned from FilterUncoveredLoss and is used to iterate over the raw logs and unpacked data for UncoveredLoss events raised by the PoolService contract.
type PoolServiceUncoveredLossIterator struct {
	Event *PoolServiceUncoveredLoss // Event containing the contract specifics and raw log

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
func (it *PoolServiceUncoveredLossIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolServiceUncoveredLoss)
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
		it.Event = new(PoolServiceUncoveredLoss)
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
func (it *PoolServiceUncoveredLossIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolServiceUncoveredLossIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolServiceUncoveredLoss represents a UncoveredLoss event raised by the PoolService contract.
type PoolServiceUncoveredLoss struct {
	CreditManager common.Address
	Loss          *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUncoveredLoss is a free log retrieval operation binding the contract event 0xef3653ded679720ab04913b6f3820be7cedc8286d42ff5dd8dff17e91bd2964c.
//
// Solidity: event UncoveredLoss(address indexed creditManager, uint256 loss)
func (_PoolService *PoolServiceFilterer) FilterUncoveredLoss(opts *bind.FilterOpts, creditManager []common.Address) (*PoolServiceUncoveredLossIterator, error) {

	var creditManagerRule []interface{}
	for _, creditManagerItem := range creditManager {
		creditManagerRule = append(creditManagerRule, creditManagerItem)
	}

	logs, sub, err := _PoolService.contract.FilterLogs(opts, "UncoveredLoss", creditManagerRule)
	if err != nil {
		return nil, err
	}
	return &PoolServiceUncoveredLossIterator{contract: _PoolService.contract, event: "UncoveredLoss", logs: logs, sub: sub}, nil
}

// WatchUncoveredLoss is a free log subscription operation binding the contract event 0xef3653ded679720ab04913b6f3820be7cedc8286d42ff5dd8dff17e91bd2964c.
//
// Solidity: event UncoveredLoss(address indexed creditManager, uint256 loss)
func (_PoolService *PoolServiceFilterer) WatchUncoveredLoss(opts *bind.WatchOpts, sink chan<- *PoolServiceUncoveredLoss, creditManager []common.Address) (event.Subscription, error) {

	var creditManagerRule []interface{}
	for _, creditManagerItem := range creditManager {
		creditManagerRule = append(creditManagerRule, creditManagerItem)
	}

	logs, sub, err := _PoolService.contract.WatchLogs(opts, "UncoveredLoss", creditManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolServiceUncoveredLoss)
				if err := _PoolService.contract.UnpackLog(event, "UncoveredLoss", log); err != nil {
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
func (_PoolService *PoolServiceFilterer) ParseUncoveredLoss(log types.Log) (*PoolServiceUncoveredLoss, error) {
	event := new(PoolServiceUncoveredLoss)
	if err := _PoolService.contract.UnpackLog(event, "UncoveredLoss", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolServiceUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the PoolService contract.
type PoolServiceUnpausedIterator struct {
	Event *PoolServiceUnpaused // Event containing the contract specifics and raw log

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
func (it *PoolServiceUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolServiceUnpaused)
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
		it.Event = new(PoolServiceUnpaused)
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
func (it *PoolServiceUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolServiceUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolServiceUnpaused represents a Unpaused event raised by the PoolService contract.
type PoolServiceUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PoolService *PoolServiceFilterer) FilterUnpaused(opts *bind.FilterOpts) (*PoolServiceUnpausedIterator, error) {

	logs, sub, err := _PoolService.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &PoolServiceUnpausedIterator{contract: _PoolService.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PoolService *PoolServiceFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PoolServiceUnpaused) (event.Subscription, error) {

	logs, sub, err := _PoolService.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolServiceUnpaused)
				if err := _PoolService.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_PoolService *PoolServiceFilterer) ParseUnpaused(log types.Log) (*PoolServiceUnpaused, error) {
	event := new(PoolServiceUnpaused)
	if err := _PoolService.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
