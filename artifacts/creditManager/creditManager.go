// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package creditManager

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// DataTypesExchange is an auto generated low-level Go binding around an user-defined struct.
type DataTypesExchange struct {
	Path         []common.Address
	AmountOutMin *big.Int
}

// CreditManagerABI is the input ABI used to generate the binding from.
const CreditManagerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addressProvider\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxLeverage\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_poolService\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_creditFilterAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_defaultSwapContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"AddCollateral\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"remainingFunds\",\"type\":\"uint256\"}],\"name\":\"CloseCreditAccount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"ExecuteOrder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"IncreaseBorrowedAmount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"remainingFunds\",\"type\":\"uint256\"}],\"name\":\"LiquidateCreditAccount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxLeverage\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeInterest\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeLiquidation\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidationDiscount\",\"type\":\"uint256\"}],\"name\":\"NewParameters\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"referralCode\",\"type\":\"uint256\"}],\"name\":\"OpenCreditAccount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"RepayCreditAccount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"TransferAccount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalValue\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isLiquidated\",\"type\":\"bool\"}],\"name\":\"_calcClosePayments\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_borrowedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountToPool\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remainingFunds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"profit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"loss\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"totalValue\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isLiquidated\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"borrowedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cumulativeIndexAtCreditAccountOpen_RAY\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cumulativeIndexNow_RAY\",\"type\":\"uint256\"}],\"name\":\"_calcClosePaymentsPure\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_borrowedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountToPool\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remainingFunds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"profit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"loss\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"addCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"targetContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isLiquidated\",\"type\":\"bool\"}],\"name\":\"calcRepayAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.Exchange[]\",\"name\":\"paths\",\"type\":\"tuple[]\"}],\"name\":\"closeCreditAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"creditAccounts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creditFilter\",\"outputs\":[{\"internalType\":\"contractICreditFilter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultSwapContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeOrder\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeInterest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeLiquidation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"}],\"name\":\"getCreditAccountOrRevert\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"}],\"name\":\"hasOpenedCreditAccount\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"increaseBorrowedAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"force\",\"type\":\"bool\"}],\"name\":\"liquidateCreditAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidationDiscount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxLeverageFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minHealthFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"leverageFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"referralCode\",\"type\":\"uint256\"}],\"name\":\"openCreditAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolService\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"targetContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"provideCreditAccountAllowance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"repayCreditAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"repayCreditAccountETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxLeverageFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_feeInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_feeLiquidation\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_liquidationDiscount\",\"type\":\"uint256\"}],\"name\":\"setParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferAccountOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlyingToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wethAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wethGateway\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// CreditManager is an auto generated Go binding around an Ethereum contract.
type CreditManager struct {
	CreditManagerCaller     // Read-only binding to the contract
	CreditManagerTransactor // Write-only binding to the contract
	CreditManagerFilterer   // Log filterer for contract events
}

// CreditManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type CreditManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CreditManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CreditManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CreditManagerSession struct {
	Contract     *CreditManager    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CreditManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CreditManagerCallerSession struct {
	Contract *CreditManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// CreditManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CreditManagerTransactorSession struct {
	Contract     *CreditManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// CreditManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type CreditManagerRaw struct {
	Contract *CreditManager // Generic contract binding to access the raw methods on
}

// CreditManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CreditManagerCallerRaw struct {
	Contract *CreditManagerCaller // Generic read-only contract binding to access the raw methods on
}

// CreditManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CreditManagerTransactorRaw struct {
	Contract *CreditManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCreditManager creates a new instance of CreditManager, bound to a specific deployed contract.
func NewCreditManager(address common.Address, backend bind.ContractBackend) (*CreditManager, error) {
	contract, err := bindCreditManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CreditManager{CreditManagerCaller: CreditManagerCaller{contract: contract}, CreditManagerTransactor: CreditManagerTransactor{contract: contract}, CreditManagerFilterer: CreditManagerFilterer{contract: contract}}, nil
}

// NewCreditManagerCaller creates a new read-only instance of CreditManager, bound to a specific deployed contract.
func NewCreditManagerCaller(address common.Address, caller bind.ContractCaller) (*CreditManagerCaller, error) {
	contract, err := bindCreditManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CreditManagerCaller{contract: contract}, nil
}

// NewCreditManagerTransactor creates a new write-only instance of CreditManager, bound to a specific deployed contract.
func NewCreditManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*CreditManagerTransactor, error) {
	contract, err := bindCreditManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CreditManagerTransactor{contract: contract}, nil
}

// NewCreditManagerFilterer creates a new log filterer instance of CreditManager, bound to a specific deployed contract.
func NewCreditManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*CreditManagerFilterer, error) {
	contract, err := bindCreditManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CreditManagerFilterer{contract: contract}, nil
}

// bindCreditManager binds a generic wrapper to an already deployed contract.
func bindCreditManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CreditManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CreditManager *CreditManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CreditManager.Contract.CreditManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CreditManager *CreditManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditManager.Contract.CreditManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CreditManager *CreditManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CreditManager.Contract.CreditManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CreditManager *CreditManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CreditManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CreditManager *CreditManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CreditManager *CreditManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CreditManager.Contract.contract.Transact(opts, method, params...)
}

// CalcClosePayments is a free data retrieval call binding the contract method 0x996329f8.
//
// Solidity: function _calcClosePayments(address creditAccount, uint256 totalValue, bool isLiquidated) view returns(uint256 _borrowedAmount, uint256 amountToPool, uint256 remainingFunds, uint256 profit, uint256 loss)
func (_CreditManager *CreditManagerCaller) CalcClosePayments(opts *bind.CallOpts, creditAccount common.Address, totalValue *big.Int, isLiquidated bool) (struct {
	BorrowedAmount *big.Int
	AmountToPool   *big.Int
	RemainingFunds *big.Int
	Profit         *big.Int
	Loss           *big.Int
}, error) {
	var out []interface{}
	err := _CreditManager.contract.Call(opts, &out, "_calcClosePayments", creditAccount, totalValue, isLiquidated)

	outstruct := new(struct {
		BorrowedAmount *big.Int
		AmountToPool   *big.Int
		RemainingFunds *big.Int
		Profit         *big.Int
		Loss           *big.Int
	})

	outstruct.BorrowedAmount = out[0].(*big.Int)
	outstruct.AmountToPool = out[1].(*big.Int)
	outstruct.RemainingFunds = out[2].(*big.Int)
	outstruct.Profit = out[3].(*big.Int)
	outstruct.Loss = out[4].(*big.Int)

	return *outstruct, err

}

// CalcClosePayments is a free data retrieval call binding the contract method 0x996329f8.
//
// Solidity: function _calcClosePayments(address creditAccount, uint256 totalValue, bool isLiquidated) view returns(uint256 _borrowedAmount, uint256 amountToPool, uint256 remainingFunds, uint256 profit, uint256 loss)
func (_CreditManager *CreditManagerSession) CalcClosePayments(creditAccount common.Address, totalValue *big.Int, isLiquidated bool) (struct {
	BorrowedAmount *big.Int
	AmountToPool   *big.Int
	RemainingFunds *big.Int
	Profit         *big.Int
	Loss           *big.Int
}, error) {
	return _CreditManager.Contract.CalcClosePayments(&_CreditManager.CallOpts, creditAccount, totalValue, isLiquidated)
}

// CalcClosePayments is a free data retrieval call binding the contract method 0x996329f8.
//
// Solidity: function _calcClosePayments(address creditAccount, uint256 totalValue, bool isLiquidated) view returns(uint256 _borrowedAmount, uint256 amountToPool, uint256 remainingFunds, uint256 profit, uint256 loss)
func (_CreditManager *CreditManagerCallerSession) CalcClosePayments(creditAccount common.Address, totalValue *big.Int, isLiquidated bool) (struct {
	BorrowedAmount *big.Int
	AmountToPool   *big.Int
	RemainingFunds *big.Int
	Profit         *big.Int
	Loss           *big.Int
}, error) {
	return _CreditManager.Contract.CalcClosePayments(&_CreditManager.CallOpts, creditAccount, totalValue, isLiquidated)
}

// CalcClosePaymentsPure is a free data retrieval call binding the contract method 0xf8dbc6b6.
//
// Solidity: function _calcClosePaymentsPure(uint256 totalValue, bool isLiquidated, uint256 borrowedAmount, uint256 cumulativeIndexAtCreditAccountOpen_RAY, uint256 cumulativeIndexNow_RAY) view returns(uint256 _borrowedAmount, uint256 amountToPool, uint256 remainingFunds, uint256 profit, uint256 loss)
func (_CreditManager *CreditManagerCaller) CalcClosePaymentsPure(opts *bind.CallOpts, totalValue *big.Int, isLiquidated bool, borrowedAmount *big.Int, cumulativeIndexAtCreditAccountOpen_RAY *big.Int, cumulativeIndexNow_RAY *big.Int) (struct {
	BorrowedAmount *big.Int
	AmountToPool   *big.Int
	RemainingFunds *big.Int
	Profit         *big.Int
	Loss           *big.Int
}, error) {
	var out []interface{}
	err := _CreditManager.contract.Call(opts, &out, "_calcClosePaymentsPure", totalValue, isLiquidated, borrowedAmount, cumulativeIndexAtCreditAccountOpen_RAY, cumulativeIndexNow_RAY)

	outstruct := new(struct {
		BorrowedAmount *big.Int
		AmountToPool   *big.Int
		RemainingFunds *big.Int
		Profit         *big.Int
		Loss           *big.Int
	})

	outstruct.BorrowedAmount = out[0].(*big.Int)
	outstruct.AmountToPool = out[1].(*big.Int)
	outstruct.RemainingFunds = out[2].(*big.Int)
	outstruct.Profit = out[3].(*big.Int)
	outstruct.Loss = out[4].(*big.Int)

	return *outstruct, err

}

// CalcClosePaymentsPure is a free data retrieval call binding the contract method 0xf8dbc6b6.
//
// Solidity: function _calcClosePaymentsPure(uint256 totalValue, bool isLiquidated, uint256 borrowedAmount, uint256 cumulativeIndexAtCreditAccountOpen_RAY, uint256 cumulativeIndexNow_RAY) view returns(uint256 _borrowedAmount, uint256 amountToPool, uint256 remainingFunds, uint256 profit, uint256 loss)
func (_CreditManager *CreditManagerSession) CalcClosePaymentsPure(totalValue *big.Int, isLiquidated bool, borrowedAmount *big.Int, cumulativeIndexAtCreditAccountOpen_RAY *big.Int, cumulativeIndexNow_RAY *big.Int) (struct {
	BorrowedAmount *big.Int
	AmountToPool   *big.Int
	RemainingFunds *big.Int
	Profit         *big.Int
	Loss           *big.Int
}, error) {
	return _CreditManager.Contract.CalcClosePaymentsPure(&_CreditManager.CallOpts, totalValue, isLiquidated, borrowedAmount, cumulativeIndexAtCreditAccountOpen_RAY, cumulativeIndexNow_RAY)
}

// CalcClosePaymentsPure is a free data retrieval call binding the contract method 0xf8dbc6b6.
//
// Solidity: function _calcClosePaymentsPure(uint256 totalValue, bool isLiquidated, uint256 borrowedAmount, uint256 cumulativeIndexAtCreditAccountOpen_RAY, uint256 cumulativeIndexNow_RAY) view returns(uint256 _borrowedAmount, uint256 amountToPool, uint256 remainingFunds, uint256 profit, uint256 loss)
func (_CreditManager *CreditManagerCallerSession) CalcClosePaymentsPure(totalValue *big.Int, isLiquidated bool, borrowedAmount *big.Int, cumulativeIndexAtCreditAccountOpen_RAY *big.Int, cumulativeIndexNow_RAY *big.Int) (struct {
	BorrowedAmount *big.Int
	AmountToPool   *big.Int
	RemainingFunds *big.Int
	Profit         *big.Int
	Loss           *big.Int
}, error) {
	return _CreditManager.Contract.CalcClosePaymentsPure(&_CreditManager.CallOpts, totalValue, isLiquidated, borrowedAmount, cumulativeIndexAtCreditAccountOpen_RAY, cumulativeIndexNow_RAY)
}

// CalcRepayAmount is a free data retrieval call binding the contract method 0x3ce07355.
//
// Solidity: function calcRepayAmount(address borrower, bool isLiquidated) view returns(uint256)
func (_CreditManager *CreditManagerCaller) CalcRepayAmount(opts *bind.CallOpts, borrower common.Address, isLiquidated bool) (*big.Int, error) {
	var out []interface{}
	err := _CreditManager.contract.Call(opts, &out, "calcRepayAmount", borrower, isLiquidated)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcRepayAmount is a free data retrieval call binding the contract method 0x3ce07355.
//
// Solidity: function calcRepayAmount(address borrower, bool isLiquidated) view returns(uint256)
func (_CreditManager *CreditManagerSession) CalcRepayAmount(borrower common.Address, isLiquidated bool) (*big.Int, error) {
	return _CreditManager.Contract.CalcRepayAmount(&_CreditManager.CallOpts, borrower, isLiquidated)
}

// CalcRepayAmount is a free data retrieval call binding the contract method 0x3ce07355.
//
// Solidity: function calcRepayAmount(address borrower, bool isLiquidated) view returns(uint256)
func (_CreditManager *CreditManagerCallerSession) CalcRepayAmount(borrower common.Address, isLiquidated bool) (*big.Int, error) {
	return _CreditManager.Contract.CalcRepayAmount(&_CreditManager.CallOpts, borrower, isLiquidated)
}

// CreditAccounts is a free data retrieval call binding the contract method 0x055ee9b5.
//
// Solidity: function creditAccounts(address ) view returns(address)
func (_CreditManager *CreditManagerCaller) CreditAccounts(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _CreditManager.contract.Call(opts, &out, "creditAccounts", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditAccounts is a free data retrieval call binding the contract method 0x055ee9b5.
//
// Solidity: function creditAccounts(address ) view returns(address)
func (_CreditManager *CreditManagerSession) CreditAccounts(arg0 common.Address) (common.Address, error) {
	return _CreditManager.Contract.CreditAccounts(&_CreditManager.CallOpts, arg0)
}

// CreditAccounts is a free data retrieval call binding the contract method 0x055ee9b5.
//
// Solidity: function creditAccounts(address ) view returns(address)
func (_CreditManager *CreditManagerCallerSession) CreditAccounts(arg0 common.Address) (common.Address, error) {
	return _CreditManager.Contract.CreditAccounts(&_CreditManager.CallOpts, arg0)
}

// CreditFilter is a free data retrieval call binding the contract method 0xf93f515b.
//
// Solidity: function creditFilter() view returns(address)
func (_CreditManager *CreditManagerCaller) CreditFilter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditManager.contract.Call(opts, &out, "creditFilter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditFilter is a free data retrieval call binding the contract method 0xf93f515b.
//
// Solidity: function creditFilter() view returns(address)
func (_CreditManager *CreditManagerSession) CreditFilter() (common.Address, error) {
	return _CreditManager.Contract.CreditFilter(&_CreditManager.CallOpts)
}

// CreditFilter is a free data retrieval call binding the contract method 0xf93f515b.
//
// Solidity: function creditFilter() view returns(address)
func (_CreditManager *CreditManagerCallerSession) CreditFilter() (common.Address, error) {
	return _CreditManager.Contract.CreditFilter(&_CreditManager.CallOpts)
}

// DefaultSwapContract is a free data retrieval call binding the contract method 0xe0c011b7.
//
// Solidity: function defaultSwapContract() view returns(address)
func (_CreditManager *CreditManagerCaller) DefaultSwapContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditManager.contract.Call(opts, &out, "defaultSwapContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DefaultSwapContract is a free data retrieval call binding the contract method 0xe0c011b7.
//
// Solidity: function defaultSwapContract() view returns(address)
func (_CreditManager *CreditManagerSession) DefaultSwapContract() (common.Address, error) {
	return _CreditManager.Contract.DefaultSwapContract(&_CreditManager.CallOpts)
}

// DefaultSwapContract is a free data retrieval call binding the contract method 0xe0c011b7.
//
// Solidity: function defaultSwapContract() view returns(address)
func (_CreditManager *CreditManagerCallerSession) DefaultSwapContract() (common.Address, error) {
	return _CreditManager.Contract.DefaultSwapContract(&_CreditManager.CallOpts)
}

// FeeInterest is a free data retrieval call binding the contract method 0x5e0b63d3.
//
// Solidity: function feeInterest() view returns(uint256)
func (_CreditManager *CreditManagerCaller) FeeInterest(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CreditManager.contract.Call(opts, &out, "feeInterest")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeInterest is a free data retrieval call binding the contract method 0x5e0b63d3.
//
// Solidity: function feeInterest() view returns(uint256)
func (_CreditManager *CreditManagerSession) FeeInterest() (*big.Int, error) {
	return _CreditManager.Contract.FeeInterest(&_CreditManager.CallOpts)
}

// FeeInterest is a free data retrieval call binding the contract method 0x5e0b63d3.
//
// Solidity: function feeInterest() view returns(uint256)
func (_CreditManager *CreditManagerCallerSession) FeeInterest() (*big.Int, error) {
	return _CreditManager.Contract.FeeInterest(&_CreditManager.CallOpts)
}

// FeeLiquidation is a free data retrieval call binding the contract method 0x3915ffaa.
//
// Solidity: function feeLiquidation() view returns(uint256)
func (_CreditManager *CreditManagerCaller) FeeLiquidation(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CreditManager.contract.Call(opts, &out, "feeLiquidation")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeLiquidation is a free data retrieval call binding the contract method 0x3915ffaa.
//
// Solidity: function feeLiquidation() view returns(uint256)
func (_CreditManager *CreditManagerSession) FeeLiquidation() (*big.Int, error) {
	return _CreditManager.Contract.FeeLiquidation(&_CreditManager.CallOpts)
}

// FeeLiquidation is a free data retrieval call binding the contract method 0x3915ffaa.
//
// Solidity: function feeLiquidation() view returns(uint256)
func (_CreditManager *CreditManagerCallerSession) FeeLiquidation() (*big.Int, error) {
	return _CreditManager.Contract.FeeLiquidation(&_CreditManager.CallOpts)
}

// GetCreditAccountOrRevert is a free data retrieval call binding the contract method 0xe958b704.
//
// Solidity: function getCreditAccountOrRevert(address borrower) view returns(address)
func (_CreditManager *CreditManagerCaller) GetCreditAccountOrRevert(opts *bind.CallOpts, borrower common.Address) (common.Address, error) {
	var out []interface{}
	err := _CreditManager.contract.Call(opts, &out, "getCreditAccountOrRevert", borrower)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCreditAccountOrRevert is a free data retrieval call binding the contract method 0xe958b704.
//
// Solidity: function getCreditAccountOrRevert(address borrower) view returns(address)
func (_CreditManager *CreditManagerSession) GetCreditAccountOrRevert(borrower common.Address) (common.Address, error) {
	return _CreditManager.Contract.GetCreditAccountOrRevert(&_CreditManager.CallOpts, borrower)
}

// GetCreditAccountOrRevert is a free data retrieval call binding the contract method 0xe958b704.
//
// Solidity: function getCreditAccountOrRevert(address borrower) view returns(address)
func (_CreditManager *CreditManagerCallerSession) GetCreditAccountOrRevert(borrower common.Address) (common.Address, error) {
	return _CreditManager.Contract.GetCreditAccountOrRevert(&_CreditManager.CallOpts, borrower)
}

// HasOpenedCreditAccount is a free data retrieval call binding the contract method 0x256ac915.
//
// Solidity: function hasOpenedCreditAccount(address borrower) view returns(bool)
func (_CreditManager *CreditManagerCaller) HasOpenedCreditAccount(opts *bind.CallOpts, borrower common.Address) (bool, error) {
	var out []interface{}
	err := _CreditManager.contract.Call(opts, &out, "hasOpenedCreditAccount", borrower)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasOpenedCreditAccount is a free data retrieval call binding the contract method 0x256ac915.
//
// Solidity: function hasOpenedCreditAccount(address borrower) view returns(bool)
func (_CreditManager *CreditManagerSession) HasOpenedCreditAccount(borrower common.Address) (bool, error) {
	return _CreditManager.Contract.HasOpenedCreditAccount(&_CreditManager.CallOpts, borrower)
}

// HasOpenedCreditAccount is a free data retrieval call binding the contract method 0x256ac915.
//
// Solidity: function hasOpenedCreditAccount(address borrower) view returns(bool)
func (_CreditManager *CreditManagerCallerSession) HasOpenedCreditAccount(borrower common.Address) (bool, error) {
	return _CreditManager.Contract.HasOpenedCreditAccount(&_CreditManager.CallOpts, borrower)
}

// LiquidationDiscount is a free data retrieval call binding the contract method 0x8053fcbe.
//
// Solidity: function liquidationDiscount() view returns(uint256)
func (_CreditManager *CreditManagerCaller) LiquidationDiscount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CreditManager.contract.Call(opts, &out, "liquidationDiscount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LiquidationDiscount is a free data retrieval call binding the contract method 0x8053fcbe.
//
// Solidity: function liquidationDiscount() view returns(uint256)
func (_CreditManager *CreditManagerSession) LiquidationDiscount() (*big.Int, error) {
	return _CreditManager.Contract.LiquidationDiscount(&_CreditManager.CallOpts)
}

// LiquidationDiscount is a free data retrieval call binding the contract method 0x8053fcbe.
//
// Solidity: function liquidationDiscount() view returns(uint256)
func (_CreditManager *CreditManagerCallerSession) LiquidationDiscount() (*big.Int, error) {
	return _CreditManager.Contract.LiquidationDiscount(&_CreditManager.CallOpts)
}

// MaxAmount is a free data retrieval call binding the contract method 0x5f48f393.
//
// Solidity: function maxAmount() view returns(uint256)
func (_CreditManager *CreditManagerCaller) MaxAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CreditManager.contract.Call(opts, &out, "maxAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxAmount is a free data retrieval call binding the contract method 0x5f48f393.
//
// Solidity: function maxAmount() view returns(uint256)
func (_CreditManager *CreditManagerSession) MaxAmount() (*big.Int, error) {
	return _CreditManager.Contract.MaxAmount(&_CreditManager.CallOpts)
}

// MaxAmount is a free data retrieval call binding the contract method 0x5f48f393.
//
// Solidity: function maxAmount() view returns(uint256)
func (_CreditManager *CreditManagerCallerSession) MaxAmount() (*big.Int, error) {
	return _CreditManager.Contract.MaxAmount(&_CreditManager.CallOpts)
}

// MaxLeverageFactor is a free data retrieval call binding the contract method 0xb2c53a6c.
//
// Solidity: function maxLeverageFactor() view returns(uint256)
func (_CreditManager *CreditManagerCaller) MaxLeverageFactor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CreditManager.contract.Call(opts, &out, "maxLeverageFactor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxLeverageFactor is a free data retrieval call binding the contract method 0xb2c53a6c.
//
// Solidity: function maxLeverageFactor() view returns(uint256)
func (_CreditManager *CreditManagerSession) MaxLeverageFactor() (*big.Int, error) {
	return _CreditManager.Contract.MaxLeverageFactor(&_CreditManager.CallOpts)
}

// MaxLeverageFactor is a free data retrieval call binding the contract method 0xb2c53a6c.
//
// Solidity: function maxLeverageFactor() view returns(uint256)
func (_CreditManager *CreditManagerCallerSession) MaxLeverageFactor() (*big.Int, error) {
	return _CreditManager.Contract.MaxLeverageFactor(&_CreditManager.CallOpts)
}

// MinAmount is a free data retrieval call binding the contract method 0x9b2cb5d8.
//
// Solidity: function minAmount() view returns(uint256)
func (_CreditManager *CreditManagerCaller) MinAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CreditManager.contract.Call(opts, &out, "minAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinAmount is a free data retrieval call binding the contract method 0x9b2cb5d8.
//
// Solidity: function minAmount() view returns(uint256)
func (_CreditManager *CreditManagerSession) MinAmount() (*big.Int, error) {
	return _CreditManager.Contract.MinAmount(&_CreditManager.CallOpts)
}

// MinAmount is a free data retrieval call binding the contract method 0x9b2cb5d8.
//
// Solidity: function minAmount() view returns(uint256)
func (_CreditManager *CreditManagerCallerSession) MinAmount() (*big.Int, error) {
	return _CreditManager.Contract.MinAmount(&_CreditManager.CallOpts)
}

// MinHealthFactor is a free data retrieval call binding the contract method 0xe1b4264c.
//
// Solidity: function minHealthFactor() view returns(uint256)
func (_CreditManager *CreditManagerCaller) MinHealthFactor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CreditManager.contract.Call(opts, &out, "minHealthFactor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinHealthFactor is a free data retrieval call binding the contract method 0xe1b4264c.
//
// Solidity: function minHealthFactor() view returns(uint256)
func (_CreditManager *CreditManagerSession) MinHealthFactor() (*big.Int, error) {
	return _CreditManager.Contract.MinHealthFactor(&_CreditManager.CallOpts)
}

// MinHealthFactor is a free data retrieval call binding the contract method 0xe1b4264c.
//
// Solidity: function minHealthFactor() view returns(uint256)
func (_CreditManager *CreditManagerCallerSession) MinHealthFactor() (*big.Int, error) {
	return _CreditManager.Contract.MinHealthFactor(&_CreditManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CreditManager *CreditManagerCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CreditManager.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CreditManager *CreditManagerSession) Paused() (bool, error) {
	return _CreditManager.Contract.Paused(&_CreditManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CreditManager *CreditManagerCallerSession) Paused() (bool, error) {
	return _CreditManager.Contract.Paused(&_CreditManager.CallOpts)
}

// PoolService is a free data retrieval call binding the contract method 0x570a7af2.
//
// Solidity: function poolService() view returns(address)
func (_CreditManager *CreditManagerCaller) PoolService(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditManager.contract.Call(opts, &out, "poolService")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolService is a free data retrieval call binding the contract method 0x570a7af2.
//
// Solidity: function poolService() view returns(address)
func (_CreditManager *CreditManagerSession) PoolService() (common.Address, error) {
	return _CreditManager.Contract.PoolService(&_CreditManager.CallOpts)
}

// PoolService is a free data retrieval call binding the contract method 0x570a7af2.
//
// Solidity: function poolService() view returns(address)
func (_CreditManager *CreditManagerCallerSession) PoolService() (common.Address, error) {
	return _CreditManager.Contract.PoolService(&_CreditManager.CallOpts)
}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_CreditManager *CreditManagerCaller) UnderlyingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditManager.contract.Call(opts, &out, "underlyingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_CreditManager *CreditManagerSession) UnderlyingToken() (common.Address, error) {
	return _CreditManager.Contract.UnderlyingToken(&_CreditManager.CallOpts)
}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_CreditManager *CreditManagerCallerSession) UnderlyingToken() (common.Address, error) {
	return _CreditManager.Contract.UnderlyingToken(&_CreditManager.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_CreditManager *CreditManagerCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CreditManager.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_CreditManager *CreditManagerSession) Version() (*big.Int, error) {
	return _CreditManager.Contract.Version(&_CreditManager.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_CreditManager *CreditManagerCallerSession) Version() (*big.Int, error) {
	return _CreditManager.Contract.Version(&_CreditManager.CallOpts)
}

// WethAddress is a free data retrieval call binding the contract method 0x4f0e0ef3.
//
// Solidity: function wethAddress() view returns(address)
func (_CreditManager *CreditManagerCaller) WethAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditManager.contract.Call(opts, &out, "wethAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WethAddress is a free data retrieval call binding the contract method 0x4f0e0ef3.
//
// Solidity: function wethAddress() view returns(address)
func (_CreditManager *CreditManagerSession) WethAddress() (common.Address, error) {
	return _CreditManager.Contract.WethAddress(&_CreditManager.CallOpts)
}

// WethAddress is a free data retrieval call binding the contract method 0x4f0e0ef3.
//
// Solidity: function wethAddress() view returns(address)
func (_CreditManager *CreditManagerCallerSession) WethAddress() (common.Address, error) {
	return _CreditManager.Contract.WethAddress(&_CreditManager.CallOpts)
}

// WethGateway is a free data retrieval call binding the contract method 0xc5e10eef.
//
// Solidity: function wethGateway() view returns(address)
func (_CreditManager *CreditManagerCaller) WethGateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditManager.contract.Call(opts, &out, "wethGateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WethGateway is a free data retrieval call binding the contract method 0xc5e10eef.
//
// Solidity: function wethGateway() view returns(address)
func (_CreditManager *CreditManagerSession) WethGateway() (common.Address, error) {
	return _CreditManager.Contract.WethGateway(&_CreditManager.CallOpts)
}

// WethGateway is a free data retrieval call binding the contract method 0xc5e10eef.
//
// Solidity: function wethGateway() view returns(address)
func (_CreditManager *CreditManagerCallerSession) WethGateway() (common.Address, error) {
	return _CreditManager.Contract.WethGateway(&_CreditManager.CallOpts)
}

// AddCollateral is a paid mutator transaction binding the contract method 0x59781034.
//
// Solidity: function addCollateral(address onBehalfOf, address token, uint256 amount) returns()
func (_CreditManager *CreditManagerTransactor) AddCollateral(opts *bind.TransactOpts, onBehalfOf common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CreditManager.contract.Transact(opts, "addCollateral", onBehalfOf, token, amount)
}

// AddCollateral is a paid mutator transaction binding the contract method 0x59781034.
//
// Solidity: function addCollateral(address onBehalfOf, address token, uint256 amount) returns()
func (_CreditManager *CreditManagerSession) AddCollateral(onBehalfOf common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CreditManager.Contract.AddCollateral(&_CreditManager.TransactOpts, onBehalfOf, token, amount)
}

// AddCollateral is a paid mutator transaction binding the contract method 0x59781034.
//
// Solidity: function addCollateral(address onBehalfOf, address token, uint256 amount) returns()
func (_CreditManager *CreditManagerTransactorSession) AddCollateral(onBehalfOf common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CreditManager.Contract.AddCollateral(&_CreditManager.TransactOpts, onBehalfOf, token, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x7e5465ba.
//
// Solidity: function approve(address targetContract, address token) returns()
func (_CreditManager *CreditManagerTransactor) Approve(opts *bind.TransactOpts, targetContract common.Address, token common.Address) (*types.Transaction, error) {
	return _CreditManager.contract.Transact(opts, "approve", targetContract, token)
}

// Approve is a paid mutator transaction binding the contract method 0x7e5465ba.
//
// Solidity: function approve(address targetContract, address token) returns()
func (_CreditManager *CreditManagerSession) Approve(targetContract common.Address, token common.Address) (*types.Transaction, error) {
	return _CreditManager.Contract.Approve(&_CreditManager.TransactOpts, targetContract, token)
}

// Approve is a paid mutator transaction binding the contract method 0x7e5465ba.
//
// Solidity: function approve(address targetContract, address token) returns()
func (_CreditManager *CreditManagerTransactorSession) Approve(targetContract common.Address, token common.Address) (*types.Transaction, error) {
	return _CreditManager.Contract.Approve(&_CreditManager.TransactOpts, targetContract, token)
}

// CloseCreditAccount is a paid mutator transaction binding the contract method 0xab114805.
//
// Solidity: function closeCreditAccount(address to, (address[],uint256)[] paths) returns()
func (_CreditManager *CreditManagerTransactor) CloseCreditAccount(opts *bind.TransactOpts, to common.Address, paths []DataTypesExchange) (*types.Transaction, error) {
	return _CreditManager.contract.Transact(opts, "closeCreditAccount", to, paths)
}

// CloseCreditAccount is a paid mutator transaction binding the contract method 0xab114805.
//
// Solidity: function closeCreditAccount(address to, (address[],uint256)[] paths) returns()
func (_CreditManager *CreditManagerSession) CloseCreditAccount(to common.Address, paths []DataTypesExchange) (*types.Transaction, error) {
	return _CreditManager.Contract.CloseCreditAccount(&_CreditManager.TransactOpts, to, paths)
}

// CloseCreditAccount is a paid mutator transaction binding the contract method 0xab114805.
//
// Solidity: function closeCreditAccount(address to, (address[],uint256)[] paths) returns()
func (_CreditManager *CreditManagerTransactorSession) CloseCreditAccount(to common.Address, paths []DataTypesExchange) (*types.Transaction, error) {
	return _CreditManager.Contract.CloseCreditAccount(&_CreditManager.TransactOpts, to, paths)
}

// ExecuteOrder is a paid mutator transaction binding the contract method 0x6ce4074a.
//
// Solidity: function executeOrder(address borrower, address target, bytes data) returns(bytes)
func (_CreditManager *CreditManagerTransactor) ExecuteOrder(opts *bind.TransactOpts, borrower common.Address, target common.Address, data []byte) (*types.Transaction, error) {
	return _CreditManager.contract.Transact(opts, "executeOrder", borrower, target, data)
}

// ExecuteOrder is a paid mutator transaction binding the contract method 0x6ce4074a.
//
// Solidity: function executeOrder(address borrower, address target, bytes data) returns(bytes)
func (_CreditManager *CreditManagerSession) ExecuteOrder(borrower common.Address, target common.Address, data []byte) (*types.Transaction, error) {
	return _CreditManager.Contract.ExecuteOrder(&_CreditManager.TransactOpts, borrower, target, data)
}

// ExecuteOrder is a paid mutator transaction binding the contract method 0x6ce4074a.
//
// Solidity: function executeOrder(address borrower, address target, bytes data) returns(bytes)
func (_CreditManager *CreditManagerTransactorSession) ExecuteOrder(borrower common.Address, target common.Address, data []byte) (*types.Transaction, error) {
	return _CreditManager.Contract.ExecuteOrder(&_CreditManager.TransactOpts, borrower, target, data)
}

// IncreaseBorrowedAmount is a paid mutator transaction binding the contract method 0x9efc60d0.
//
// Solidity: function increaseBorrowedAmount(uint256 amount) returns()
func (_CreditManager *CreditManagerTransactor) IncreaseBorrowedAmount(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _CreditManager.contract.Transact(opts, "increaseBorrowedAmount", amount)
}

// IncreaseBorrowedAmount is a paid mutator transaction binding the contract method 0x9efc60d0.
//
// Solidity: function increaseBorrowedAmount(uint256 amount) returns()
func (_CreditManager *CreditManagerSession) IncreaseBorrowedAmount(amount *big.Int) (*types.Transaction, error) {
	return _CreditManager.Contract.IncreaseBorrowedAmount(&_CreditManager.TransactOpts, amount)
}

// IncreaseBorrowedAmount is a paid mutator transaction binding the contract method 0x9efc60d0.
//
// Solidity: function increaseBorrowedAmount(uint256 amount) returns()
func (_CreditManager *CreditManagerTransactorSession) IncreaseBorrowedAmount(amount *big.Int) (*types.Transaction, error) {
	return _CreditManager.Contract.IncreaseBorrowedAmount(&_CreditManager.TransactOpts, amount)
}

// LiquidateCreditAccount is a paid mutator transaction binding the contract method 0xa69a7dd6.
//
// Solidity: function liquidateCreditAccount(address borrower, address to, bool force) returns()
func (_CreditManager *CreditManagerTransactor) LiquidateCreditAccount(opts *bind.TransactOpts, borrower common.Address, to common.Address, force bool) (*types.Transaction, error) {
	return _CreditManager.contract.Transact(opts, "liquidateCreditAccount", borrower, to, force)
}

// LiquidateCreditAccount is a paid mutator transaction binding the contract method 0xa69a7dd6.
//
// Solidity: function liquidateCreditAccount(address borrower, address to, bool force) returns()
func (_CreditManager *CreditManagerSession) LiquidateCreditAccount(borrower common.Address, to common.Address, force bool) (*types.Transaction, error) {
	return _CreditManager.Contract.LiquidateCreditAccount(&_CreditManager.TransactOpts, borrower, to, force)
}

// LiquidateCreditAccount is a paid mutator transaction binding the contract method 0xa69a7dd6.
//
// Solidity: function liquidateCreditAccount(address borrower, address to, bool force) returns()
func (_CreditManager *CreditManagerTransactorSession) LiquidateCreditAccount(borrower common.Address, to common.Address, force bool) (*types.Transaction, error) {
	return _CreditManager.Contract.LiquidateCreditAccount(&_CreditManager.TransactOpts, borrower, to, force)
}

// OpenCreditAccount is a paid mutator transaction binding the contract method 0x5288ba4b.
//
// Solidity: function openCreditAccount(uint256 amount, address onBehalfOf, uint256 leverageFactor, uint256 referralCode) returns()
func (_CreditManager *CreditManagerTransactor) OpenCreditAccount(opts *bind.TransactOpts, amount *big.Int, onBehalfOf common.Address, leverageFactor *big.Int, referralCode *big.Int) (*types.Transaction, error) {
	return _CreditManager.contract.Transact(opts, "openCreditAccount", amount, onBehalfOf, leverageFactor, referralCode)
}

// OpenCreditAccount is a paid mutator transaction binding the contract method 0x5288ba4b.
//
// Solidity: function openCreditAccount(uint256 amount, address onBehalfOf, uint256 leverageFactor, uint256 referralCode) returns()
func (_CreditManager *CreditManagerSession) OpenCreditAccount(amount *big.Int, onBehalfOf common.Address, leverageFactor *big.Int, referralCode *big.Int) (*types.Transaction, error) {
	return _CreditManager.Contract.OpenCreditAccount(&_CreditManager.TransactOpts, amount, onBehalfOf, leverageFactor, referralCode)
}

// OpenCreditAccount is a paid mutator transaction binding the contract method 0x5288ba4b.
//
// Solidity: function openCreditAccount(uint256 amount, address onBehalfOf, uint256 leverageFactor, uint256 referralCode) returns()
func (_CreditManager *CreditManagerTransactorSession) OpenCreditAccount(amount *big.Int, onBehalfOf common.Address, leverageFactor *big.Int, referralCode *big.Int) (*types.Transaction, error) {
	return _CreditManager.Contract.OpenCreditAccount(&_CreditManager.TransactOpts, amount, onBehalfOf, leverageFactor, referralCode)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CreditManager *CreditManagerTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditManager.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CreditManager *CreditManagerSession) Pause() (*types.Transaction, error) {
	return _CreditManager.Contract.Pause(&_CreditManager.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CreditManager *CreditManagerTransactorSession) Pause() (*types.Transaction, error) {
	return _CreditManager.Contract.Pause(&_CreditManager.TransactOpts)
}

// ProvideCreditAccountAllowance is a paid mutator transaction binding the contract method 0x579122ab.
//
// Solidity: function provideCreditAccountAllowance(address creditAccount, address targetContract, address token) returns()
func (_CreditManager *CreditManagerTransactor) ProvideCreditAccountAllowance(opts *bind.TransactOpts, creditAccount common.Address, targetContract common.Address, token common.Address) (*types.Transaction, error) {
	return _CreditManager.contract.Transact(opts, "provideCreditAccountAllowance", creditAccount, targetContract, token)
}

// ProvideCreditAccountAllowance is a paid mutator transaction binding the contract method 0x579122ab.
//
// Solidity: function provideCreditAccountAllowance(address creditAccount, address targetContract, address token) returns()
func (_CreditManager *CreditManagerSession) ProvideCreditAccountAllowance(creditAccount common.Address, targetContract common.Address, token common.Address) (*types.Transaction, error) {
	return _CreditManager.Contract.ProvideCreditAccountAllowance(&_CreditManager.TransactOpts, creditAccount, targetContract, token)
}

// ProvideCreditAccountAllowance is a paid mutator transaction binding the contract method 0x579122ab.
//
// Solidity: function provideCreditAccountAllowance(address creditAccount, address targetContract, address token) returns()
func (_CreditManager *CreditManagerTransactorSession) ProvideCreditAccountAllowance(creditAccount common.Address, targetContract common.Address, token common.Address) (*types.Transaction, error) {
	return _CreditManager.Contract.ProvideCreditAccountAllowance(&_CreditManager.TransactOpts, creditAccount, targetContract, token)
}

// RepayCreditAccount is a paid mutator transaction binding the contract method 0xd692ba33.
//
// Solidity: function repayCreditAccount(address to) returns()
func (_CreditManager *CreditManagerTransactor) RepayCreditAccount(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _CreditManager.contract.Transact(opts, "repayCreditAccount", to)
}

// RepayCreditAccount is a paid mutator transaction binding the contract method 0xd692ba33.
//
// Solidity: function repayCreditAccount(address to) returns()
func (_CreditManager *CreditManagerSession) RepayCreditAccount(to common.Address) (*types.Transaction, error) {
	return _CreditManager.Contract.RepayCreditAccount(&_CreditManager.TransactOpts, to)
}

// RepayCreditAccount is a paid mutator transaction binding the contract method 0xd692ba33.
//
// Solidity: function repayCreditAccount(address to) returns()
func (_CreditManager *CreditManagerTransactorSession) RepayCreditAccount(to common.Address) (*types.Transaction, error) {
	return _CreditManager.Contract.RepayCreditAccount(&_CreditManager.TransactOpts, to)
}

// RepayCreditAccountETH is a paid mutator transaction binding the contract method 0xa6eab5c2.
//
// Solidity: function repayCreditAccountETH(address borrower, address to) returns(uint256)
func (_CreditManager *CreditManagerTransactor) RepayCreditAccountETH(opts *bind.TransactOpts, borrower common.Address, to common.Address) (*types.Transaction, error) {
	return _CreditManager.contract.Transact(opts, "repayCreditAccountETH", borrower, to)
}

// RepayCreditAccountETH is a paid mutator transaction binding the contract method 0xa6eab5c2.
//
// Solidity: function repayCreditAccountETH(address borrower, address to) returns(uint256)
func (_CreditManager *CreditManagerSession) RepayCreditAccountETH(borrower common.Address, to common.Address) (*types.Transaction, error) {
	return _CreditManager.Contract.RepayCreditAccountETH(&_CreditManager.TransactOpts, borrower, to)
}

// RepayCreditAccountETH is a paid mutator transaction binding the contract method 0xa6eab5c2.
//
// Solidity: function repayCreditAccountETH(address borrower, address to) returns(uint256)
func (_CreditManager *CreditManagerTransactorSession) RepayCreditAccountETH(borrower common.Address, to common.Address) (*types.Transaction, error) {
	return _CreditManager.Contract.RepayCreditAccountETH(&_CreditManager.TransactOpts, borrower, to)
}

// SetParams is a paid mutator transaction binding the contract method 0xebb39512.
//
// Solidity: function setParams(uint256 _minAmount, uint256 _maxAmount, uint256 _maxLeverageFactor, uint256 _feeInterest, uint256 _feeLiquidation, uint256 _liquidationDiscount) returns()
func (_CreditManager *CreditManagerTransactor) SetParams(opts *bind.TransactOpts, _minAmount *big.Int, _maxAmount *big.Int, _maxLeverageFactor *big.Int, _feeInterest *big.Int, _feeLiquidation *big.Int, _liquidationDiscount *big.Int) (*types.Transaction, error) {
	return _CreditManager.contract.Transact(opts, "setParams", _minAmount, _maxAmount, _maxLeverageFactor, _feeInterest, _feeLiquidation, _liquidationDiscount)
}

// SetParams is a paid mutator transaction binding the contract method 0xebb39512.
//
// Solidity: function setParams(uint256 _minAmount, uint256 _maxAmount, uint256 _maxLeverageFactor, uint256 _feeInterest, uint256 _feeLiquidation, uint256 _liquidationDiscount) returns()
func (_CreditManager *CreditManagerSession) SetParams(_minAmount *big.Int, _maxAmount *big.Int, _maxLeverageFactor *big.Int, _feeInterest *big.Int, _feeLiquidation *big.Int, _liquidationDiscount *big.Int) (*types.Transaction, error) {
	return _CreditManager.Contract.SetParams(&_CreditManager.TransactOpts, _minAmount, _maxAmount, _maxLeverageFactor, _feeInterest, _feeLiquidation, _liquidationDiscount)
}

// SetParams is a paid mutator transaction binding the contract method 0xebb39512.
//
// Solidity: function setParams(uint256 _minAmount, uint256 _maxAmount, uint256 _maxLeverageFactor, uint256 _feeInterest, uint256 _feeLiquidation, uint256 _liquidationDiscount) returns()
func (_CreditManager *CreditManagerTransactorSession) SetParams(_minAmount *big.Int, _maxAmount *big.Int, _maxLeverageFactor *big.Int, _feeInterest *big.Int, _feeLiquidation *big.Int, _liquidationDiscount *big.Int) (*types.Transaction, error) {
	return _CreditManager.Contract.SetParams(&_CreditManager.TransactOpts, _minAmount, _maxAmount, _maxLeverageFactor, _feeInterest, _feeLiquidation, _liquidationDiscount)
}

// TransferAccountOwnership is a paid mutator transaction binding the contract method 0x5019e20a.
//
// Solidity: function transferAccountOwnership(address newOwner) returns()
func (_CreditManager *CreditManagerTransactor) TransferAccountOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CreditManager.contract.Transact(opts, "transferAccountOwnership", newOwner)
}

// TransferAccountOwnership is a paid mutator transaction binding the contract method 0x5019e20a.
//
// Solidity: function transferAccountOwnership(address newOwner) returns()
func (_CreditManager *CreditManagerSession) TransferAccountOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CreditManager.Contract.TransferAccountOwnership(&_CreditManager.TransactOpts, newOwner)
}

// TransferAccountOwnership is a paid mutator transaction binding the contract method 0x5019e20a.
//
// Solidity: function transferAccountOwnership(address newOwner) returns()
func (_CreditManager *CreditManagerTransactorSession) TransferAccountOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CreditManager.Contract.TransferAccountOwnership(&_CreditManager.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CreditManager *CreditManagerTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditManager.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CreditManager *CreditManagerSession) Unpause() (*types.Transaction, error) {
	return _CreditManager.Contract.Unpause(&_CreditManager.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CreditManager *CreditManagerTransactorSession) Unpause() (*types.Transaction, error) {
	return _CreditManager.Contract.Unpause(&_CreditManager.TransactOpts)
}

// CreditManagerAddCollateralIterator is returned from FilterAddCollateral and is used to iterate over the raw logs and unpacked data for AddCollateral events raised by the CreditManager contract.
type CreditManagerAddCollateralIterator struct {
	Event *CreditManagerAddCollateral // Event containing the contract specifics and raw log

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
func (it *CreditManagerAddCollateralIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditManagerAddCollateral)
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
		it.Event = new(CreditManagerAddCollateral)
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
func (it *CreditManagerAddCollateralIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditManagerAddCollateralIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditManagerAddCollateral represents a AddCollateral event raised by the CreditManager contract.
type CreditManagerAddCollateral struct {
	OnBehalfOf common.Address
	Token      common.Address
	Value      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAddCollateral is a free log retrieval operation binding the contract event 0xa32435755c235de2976ed44a75a2f85cb01faf0c894f639fe0c32bb9455fea8f.
//
// Solidity: event AddCollateral(address indexed onBehalfOf, address indexed token, uint256 value)
func (_CreditManager *CreditManagerFilterer) FilterAddCollateral(opts *bind.FilterOpts, onBehalfOf []common.Address, token []common.Address) (*CreditManagerAddCollateralIterator, error) {

	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CreditManager.contract.FilterLogs(opts, "AddCollateral", onBehalfOfRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &CreditManagerAddCollateralIterator{contract: _CreditManager.contract, event: "AddCollateral", logs: logs, sub: sub}, nil
}

// WatchAddCollateral is a free log subscription operation binding the contract event 0xa32435755c235de2976ed44a75a2f85cb01faf0c894f639fe0c32bb9455fea8f.
//
// Solidity: event AddCollateral(address indexed onBehalfOf, address indexed token, uint256 value)
func (_CreditManager *CreditManagerFilterer) WatchAddCollateral(opts *bind.WatchOpts, sink chan<- *CreditManagerAddCollateral, onBehalfOf []common.Address, token []common.Address) (event.Subscription, error) {

	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CreditManager.contract.WatchLogs(opts, "AddCollateral", onBehalfOfRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditManagerAddCollateral)
				if err := _CreditManager.contract.UnpackLog(event, "AddCollateral", log); err != nil {
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

// ParseAddCollateral is a log parse operation binding the contract event 0xa32435755c235de2976ed44a75a2f85cb01faf0c894f639fe0c32bb9455fea8f.
//
// Solidity: event AddCollateral(address indexed onBehalfOf, address indexed token, uint256 value)
func (_CreditManager *CreditManagerFilterer) ParseAddCollateral(log types.Log) (*CreditManagerAddCollateral, error) {
	event := new(CreditManagerAddCollateral)
	if err := _CreditManager.contract.UnpackLog(event, "AddCollateral", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CreditManagerCloseCreditAccountIterator is returned from FilterCloseCreditAccount and is used to iterate over the raw logs and unpacked data for CloseCreditAccount events raised by the CreditManager contract.
type CreditManagerCloseCreditAccountIterator struct {
	Event *CreditManagerCloseCreditAccount // Event containing the contract specifics and raw log

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
func (it *CreditManagerCloseCreditAccountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditManagerCloseCreditAccount)
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
		it.Event = new(CreditManagerCloseCreditAccount)
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
func (it *CreditManagerCloseCreditAccountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditManagerCloseCreditAccountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditManagerCloseCreditAccount represents a CloseCreditAccount event raised by the CreditManager contract.
type CreditManagerCloseCreditAccount struct {
	Owner          common.Address
	To             common.Address
	RemainingFunds *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterCloseCreditAccount is a free log retrieval operation binding the contract event 0xca05b632388199c23de1352b2e96fd72a0ec71611683330b38060c004bbf0a76.
//
// Solidity: event CloseCreditAccount(address indexed owner, address indexed to, uint256 remainingFunds)
func (_CreditManager *CreditManagerFilterer) FilterCloseCreditAccount(opts *bind.FilterOpts, owner []common.Address, to []common.Address) (*CreditManagerCloseCreditAccountIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CreditManager.contract.FilterLogs(opts, "CloseCreditAccount", ownerRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CreditManagerCloseCreditAccountIterator{contract: _CreditManager.contract, event: "CloseCreditAccount", logs: logs, sub: sub}, nil
}

// WatchCloseCreditAccount is a free log subscription operation binding the contract event 0xca05b632388199c23de1352b2e96fd72a0ec71611683330b38060c004bbf0a76.
//
// Solidity: event CloseCreditAccount(address indexed owner, address indexed to, uint256 remainingFunds)
func (_CreditManager *CreditManagerFilterer) WatchCloseCreditAccount(opts *bind.WatchOpts, sink chan<- *CreditManagerCloseCreditAccount, owner []common.Address, to []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CreditManager.contract.WatchLogs(opts, "CloseCreditAccount", ownerRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditManagerCloseCreditAccount)
				if err := _CreditManager.contract.UnpackLog(event, "CloseCreditAccount", log); err != nil {
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

// ParseCloseCreditAccount is a log parse operation binding the contract event 0xca05b632388199c23de1352b2e96fd72a0ec71611683330b38060c004bbf0a76.
//
// Solidity: event CloseCreditAccount(address indexed owner, address indexed to, uint256 remainingFunds)
func (_CreditManager *CreditManagerFilterer) ParseCloseCreditAccount(log types.Log) (*CreditManagerCloseCreditAccount, error) {
	event := new(CreditManagerCloseCreditAccount)
	if err := _CreditManager.contract.UnpackLog(event, "CloseCreditAccount", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CreditManagerExecuteOrderIterator is returned from FilterExecuteOrder and is used to iterate over the raw logs and unpacked data for ExecuteOrder events raised by the CreditManager contract.
type CreditManagerExecuteOrderIterator struct {
	Event *CreditManagerExecuteOrder // Event containing the contract specifics and raw log

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
func (it *CreditManagerExecuteOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditManagerExecuteOrder)
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
		it.Event = new(CreditManagerExecuteOrder)
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
func (it *CreditManagerExecuteOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditManagerExecuteOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditManagerExecuteOrder represents a ExecuteOrder event raised by the CreditManager contract.
type CreditManagerExecuteOrder struct {
	Borrower common.Address
	Target   common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterExecuteOrder is a free log retrieval operation binding the contract event 0xaed1eb34af6acd8c1e3911fb2ebb875a66324b03957886bd002227b17f52ab03.
//
// Solidity: event ExecuteOrder(address indexed borrower, address indexed target)
func (_CreditManager *CreditManagerFilterer) FilterExecuteOrder(opts *bind.FilterOpts, borrower []common.Address, target []common.Address) (*CreditManagerExecuteOrderIterator, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _CreditManager.contract.FilterLogs(opts, "ExecuteOrder", borrowerRule, targetRule)
	if err != nil {
		return nil, err
	}
	return &CreditManagerExecuteOrderIterator{contract: _CreditManager.contract, event: "ExecuteOrder", logs: logs, sub: sub}, nil
}

// WatchExecuteOrder is a free log subscription operation binding the contract event 0xaed1eb34af6acd8c1e3911fb2ebb875a66324b03957886bd002227b17f52ab03.
//
// Solidity: event ExecuteOrder(address indexed borrower, address indexed target)
func (_CreditManager *CreditManagerFilterer) WatchExecuteOrder(opts *bind.WatchOpts, sink chan<- *CreditManagerExecuteOrder, borrower []common.Address, target []common.Address) (event.Subscription, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _CreditManager.contract.WatchLogs(opts, "ExecuteOrder", borrowerRule, targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditManagerExecuteOrder)
				if err := _CreditManager.contract.UnpackLog(event, "ExecuteOrder", log); err != nil {
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

// ParseExecuteOrder is a log parse operation binding the contract event 0xaed1eb34af6acd8c1e3911fb2ebb875a66324b03957886bd002227b17f52ab03.
//
// Solidity: event ExecuteOrder(address indexed borrower, address indexed target)
func (_CreditManager *CreditManagerFilterer) ParseExecuteOrder(log types.Log) (*CreditManagerExecuteOrder, error) {
	event := new(CreditManagerExecuteOrder)
	if err := _CreditManager.contract.UnpackLog(event, "ExecuteOrder", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CreditManagerIncreaseBorrowedAmountIterator is returned from FilterIncreaseBorrowedAmount and is used to iterate over the raw logs and unpacked data for IncreaseBorrowedAmount events raised by the CreditManager contract.
type CreditManagerIncreaseBorrowedAmountIterator struct {
	Event *CreditManagerIncreaseBorrowedAmount // Event containing the contract specifics and raw log

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
func (it *CreditManagerIncreaseBorrowedAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditManagerIncreaseBorrowedAmount)
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
		it.Event = new(CreditManagerIncreaseBorrowedAmount)
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
func (it *CreditManagerIncreaseBorrowedAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditManagerIncreaseBorrowedAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditManagerIncreaseBorrowedAmount represents a IncreaseBorrowedAmount event raised by the CreditManager contract.
type CreditManagerIncreaseBorrowedAmount struct {
	Borrower common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterIncreaseBorrowedAmount is a free log retrieval operation binding the contract event 0x9cac51154cc0d835e2f9c9d1f59a9344588cee107f4203bf58a8c797e3a58c45.
//
// Solidity: event IncreaseBorrowedAmount(address indexed borrower, uint256 amount)
func (_CreditManager *CreditManagerFilterer) FilterIncreaseBorrowedAmount(opts *bind.FilterOpts, borrower []common.Address) (*CreditManagerIncreaseBorrowedAmountIterator, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _CreditManager.contract.FilterLogs(opts, "IncreaseBorrowedAmount", borrowerRule)
	if err != nil {
		return nil, err
	}
	return &CreditManagerIncreaseBorrowedAmountIterator{contract: _CreditManager.contract, event: "IncreaseBorrowedAmount", logs: logs, sub: sub}, nil
}

// WatchIncreaseBorrowedAmount is a free log subscription operation binding the contract event 0x9cac51154cc0d835e2f9c9d1f59a9344588cee107f4203bf58a8c797e3a58c45.
//
// Solidity: event IncreaseBorrowedAmount(address indexed borrower, uint256 amount)
func (_CreditManager *CreditManagerFilterer) WatchIncreaseBorrowedAmount(opts *bind.WatchOpts, sink chan<- *CreditManagerIncreaseBorrowedAmount, borrower []common.Address) (event.Subscription, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _CreditManager.contract.WatchLogs(opts, "IncreaseBorrowedAmount", borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditManagerIncreaseBorrowedAmount)
				if err := _CreditManager.contract.UnpackLog(event, "IncreaseBorrowedAmount", log); err != nil {
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

// ParseIncreaseBorrowedAmount is a log parse operation binding the contract event 0x9cac51154cc0d835e2f9c9d1f59a9344588cee107f4203bf58a8c797e3a58c45.
//
// Solidity: event IncreaseBorrowedAmount(address indexed borrower, uint256 amount)
func (_CreditManager *CreditManagerFilterer) ParseIncreaseBorrowedAmount(log types.Log) (*CreditManagerIncreaseBorrowedAmount, error) {
	event := new(CreditManagerIncreaseBorrowedAmount)
	if err := _CreditManager.contract.UnpackLog(event, "IncreaseBorrowedAmount", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CreditManagerLiquidateCreditAccountIterator is returned from FilterLiquidateCreditAccount and is used to iterate over the raw logs and unpacked data for LiquidateCreditAccount events raised by the CreditManager contract.
type CreditManagerLiquidateCreditAccountIterator struct {
	Event *CreditManagerLiquidateCreditAccount // Event containing the contract specifics and raw log

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
func (it *CreditManagerLiquidateCreditAccountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditManagerLiquidateCreditAccount)
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
		it.Event = new(CreditManagerLiquidateCreditAccount)
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
func (it *CreditManagerLiquidateCreditAccountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditManagerLiquidateCreditAccountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditManagerLiquidateCreditAccount represents a LiquidateCreditAccount event raised by the CreditManager contract.
type CreditManagerLiquidateCreditAccount struct {
	Owner          common.Address
	Liquidator     common.Address
	RemainingFunds *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLiquidateCreditAccount is a free log retrieval operation binding the contract event 0x5e5da6c348e62989f9cfe029252433fc99009b7d28fa3c20d675520a10ff5896.
//
// Solidity: event LiquidateCreditAccount(address indexed owner, address indexed liquidator, uint256 remainingFunds)
func (_CreditManager *CreditManagerFilterer) FilterLiquidateCreditAccount(opts *bind.FilterOpts, owner []common.Address, liquidator []common.Address) (*CreditManagerLiquidateCreditAccountIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var liquidatorRule []interface{}
	for _, liquidatorItem := range liquidator {
		liquidatorRule = append(liquidatorRule, liquidatorItem)
	}

	logs, sub, err := _CreditManager.contract.FilterLogs(opts, "LiquidateCreditAccount", ownerRule, liquidatorRule)
	if err != nil {
		return nil, err
	}
	return &CreditManagerLiquidateCreditAccountIterator{contract: _CreditManager.contract, event: "LiquidateCreditAccount", logs: logs, sub: sub}, nil
}

// WatchLiquidateCreditAccount is a free log subscription operation binding the contract event 0x5e5da6c348e62989f9cfe029252433fc99009b7d28fa3c20d675520a10ff5896.
//
// Solidity: event LiquidateCreditAccount(address indexed owner, address indexed liquidator, uint256 remainingFunds)
func (_CreditManager *CreditManagerFilterer) WatchLiquidateCreditAccount(opts *bind.WatchOpts, sink chan<- *CreditManagerLiquidateCreditAccount, owner []common.Address, liquidator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var liquidatorRule []interface{}
	for _, liquidatorItem := range liquidator {
		liquidatorRule = append(liquidatorRule, liquidatorItem)
	}

	logs, sub, err := _CreditManager.contract.WatchLogs(opts, "LiquidateCreditAccount", ownerRule, liquidatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditManagerLiquidateCreditAccount)
				if err := _CreditManager.contract.UnpackLog(event, "LiquidateCreditAccount", log); err != nil {
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

// ParseLiquidateCreditAccount is a log parse operation binding the contract event 0x5e5da6c348e62989f9cfe029252433fc99009b7d28fa3c20d675520a10ff5896.
//
// Solidity: event LiquidateCreditAccount(address indexed owner, address indexed liquidator, uint256 remainingFunds)
func (_CreditManager *CreditManagerFilterer) ParseLiquidateCreditAccount(log types.Log) (*CreditManagerLiquidateCreditAccount, error) {
	event := new(CreditManagerLiquidateCreditAccount)
	if err := _CreditManager.contract.UnpackLog(event, "LiquidateCreditAccount", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CreditManagerNewParametersIterator is returned from FilterNewParameters and is used to iterate over the raw logs and unpacked data for NewParameters events raised by the CreditManager contract.
type CreditManagerNewParametersIterator struct {
	Event *CreditManagerNewParameters // Event containing the contract specifics and raw log

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
func (it *CreditManagerNewParametersIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditManagerNewParameters)
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
		it.Event = new(CreditManagerNewParameters)
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
func (it *CreditManagerNewParametersIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditManagerNewParametersIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditManagerNewParameters represents a NewParameters event raised by the CreditManager contract.
type CreditManagerNewParameters struct {
	MinAmount           *big.Int
	MaxAmount           *big.Int
	MaxLeverage         *big.Int
	FeeInterest         *big.Int
	FeeLiquidation      *big.Int
	LiquidationDiscount *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterNewParameters is a free log retrieval operation binding the contract event 0xa32137411fc7c20db359079cd84af0e2cad58cd7a182a8a5e23e08e554e88bf0.
//
// Solidity: event NewParameters(uint256 minAmount, uint256 maxAmount, uint256 maxLeverage, uint256 feeInterest, uint256 feeLiquidation, uint256 liquidationDiscount)
func (_CreditManager *CreditManagerFilterer) FilterNewParameters(opts *bind.FilterOpts) (*CreditManagerNewParametersIterator, error) {

	logs, sub, err := _CreditManager.contract.FilterLogs(opts, "NewParameters")
	if err != nil {
		return nil, err
	}
	return &CreditManagerNewParametersIterator{contract: _CreditManager.contract, event: "NewParameters", logs: logs, sub: sub}, nil
}

// WatchNewParameters is a free log subscription operation binding the contract event 0xa32137411fc7c20db359079cd84af0e2cad58cd7a182a8a5e23e08e554e88bf0.
//
// Solidity: event NewParameters(uint256 minAmount, uint256 maxAmount, uint256 maxLeverage, uint256 feeInterest, uint256 feeLiquidation, uint256 liquidationDiscount)
func (_CreditManager *CreditManagerFilterer) WatchNewParameters(opts *bind.WatchOpts, sink chan<- *CreditManagerNewParameters) (event.Subscription, error) {

	logs, sub, err := _CreditManager.contract.WatchLogs(opts, "NewParameters")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditManagerNewParameters)
				if err := _CreditManager.contract.UnpackLog(event, "NewParameters", log); err != nil {
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

// ParseNewParameters is a log parse operation binding the contract event 0xa32137411fc7c20db359079cd84af0e2cad58cd7a182a8a5e23e08e554e88bf0.
//
// Solidity: event NewParameters(uint256 minAmount, uint256 maxAmount, uint256 maxLeverage, uint256 feeInterest, uint256 feeLiquidation, uint256 liquidationDiscount)
func (_CreditManager *CreditManagerFilterer) ParseNewParameters(log types.Log) (*CreditManagerNewParameters, error) {
	event := new(CreditManagerNewParameters)
	if err := _CreditManager.contract.UnpackLog(event, "NewParameters", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CreditManagerOpenCreditAccountIterator is returned from FilterOpenCreditAccount and is used to iterate over the raw logs and unpacked data for OpenCreditAccount events raised by the CreditManager contract.
type CreditManagerOpenCreditAccountIterator struct {
	Event *CreditManagerOpenCreditAccount // Event containing the contract specifics and raw log

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
func (it *CreditManagerOpenCreditAccountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditManagerOpenCreditAccount)
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
		it.Event = new(CreditManagerOpenCreditAccount)
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
func (it *CreditManagerOpenCreditAccountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditManagerOpenCreditAccountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditManagerOpenCreditAccount represents a OpenCreditAccount event raised by the CreditManager contract.
type CreditManagerOpenCreditAccount struct {
	Sender        common.Address
	OnBehalfOf    common.Address
	CreditAccount common.Address
	Amount        *big.Int
	BorrowAmount  *big.Int
	ReferralCode  *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOpenCreditAccount is a free log retrieval operation binding the contract event 0x7b20ae77867a263a1074203a2da261ef0d096c99395c59c9d4a0104b9f334a27.
//
// Solidity: event OpenCreditAccount(address indexed sender, address indexed onBehalfOf, address indexed creditAccount, uint256 amount, uint256 borrowAmount, uint256 referralCode)
func (_CreditManager *CreditManagerFilterer) FilterOpenCreditAccount(opts *bind.FilterOpts, sender []common.Address, onBehalfOf []common.Address, creditAccount []common.Address) (*CreditManagerOpenCreditAccountIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}
	var creditAccountRule []interface{}
	for _, creditAccountItem := range creditAccount {
		creditAccountRule = append(creditAccountRule, creditAccountItem)
	}

	logs, sub, err := _CreditManager.contract.FilterLogs(opts, "OpenCreditAccount", senderRule, onBehalfOfRule, creditAccountRule)
	if err != nil {
		return nil, err
	}
	return &CreditManagerOpenCreditAccountIterator{contract: _CreditManager.contract, event: "OpenCreditAccount", logs: logs, sub: sub}, nil
}

// WatchOpenCreditAccount is a free log subscription operation binding the contract event 0x7b20ae77867a263a1074203a2da261ef0d096c99395c59c9d4a0104b9f334a27.
//
// Solidity: event OpenCreditAccount(address indexed sender, address indexed onBehalfOf, address indexed creditAccount, uint256 amount, uint256 borrowAmount, uint256 referralCode)
func (_CreditManager *CreditManagerFilterer) WatchOpenCreditAccount(opts *bind.WatchOpts, sink chan<- *CreditManagerOpenCreditAccount, sender []common.Address, onBehalfOf []common.Address, creditAccount []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}
	var creditAccountRule []interface{}
	for _, creditAccountItem := range creditAccount {
		creditAccountRule = append(creditAccountRule, creditAccountItem)
	}

	logs, sub, err := _CreditManager.contract.WatchLogs(opts, "OpenCreditAccount", senderRule, onBehalfOfRule, creditAccountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditManagerOpenCreditAccount)
				if err := _CreditManager.contract.UnpackLog(event, "OpenCreditAccount", log); err != nil {
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

// ParseOpenCreditAccount is a log parse operation binding the contract event 0x7b20ae77867a263a1074203a2da261ef0d096c99395c59c9d4a0104b9f334a27.
//
// Solidity: event OpenCreditAccount(address indexed sender, address indexed onBehalfOf, address indexed creditAccount, uint256 amount, uint256 borrowAmount, uint256 referralCode)
func (_CreditManager *CreditManagerFilterer) ParseOpenCreditAccount(log types.Log) (*CreditManagerOpenCreditAccount, error) {
	event := new(CreditManagerOpenCreditAccount)
	if err := _CreditManager.contract.UnpackLog(event, "OpenCreditAccount", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CreditManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the CreditManager contract.
type CreditManagerPausedIterator struct {
	Event *CreditManagerPaused // Event containing the contract specifics and raw log

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
func (it *CreditManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditManagerPaused)
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
		it.Event = new(CreditManagerPaused)
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
func (it *CreditManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditManagerPaused represents a Paused event raised by the CreditManager contract.
type CreditManagerPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_CreditManager *CreditManagerFilterer) FilterPaused(opts *bind.FilterOpts) (*CreditManagerPausedIterator, error) {

	logs, sub, err := _CreditManager.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &CreditManagerPausedIterator{contract: _CreditManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_CreditManager *CreditManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *CreditManagerPaused) (event.Subscription, error) {

	logs, sub, err := _CreditManager.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditManagerPaused)
				if err := _CreditManager.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_CreditManager *CreditManagerFilterer) ParsePaused(log types.Log) (*CreditManagerPaused, error) {
	event := new(CreditManagerPaused)
	if err := _CreditManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CreditManagerRepayCreditAccountIterator is returned from FilterRepayCreditAccount and is used to iterate over the raw logs and unpacked data for RepayCreditAccount events raised by the CreditManager contract.
type CreditManagerRepayCreditAccountIterator struct {
	Event *CreditManagerRepayCreditAccount // Event containing the contract specifics and raw log

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
func (it *CreditManagerRepayCreditAccountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditManagerRepayCreditAccount)
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
		it.Event = new(CreditManagerRepayCreditAccount)
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
func (it *CreditManagerRepayCreditAccountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditManagerRepayCreditAccountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditManagerRepayCreditAccount represents a RepayCreditAccount event raised by the CreditManager contract.
type CreditManagerRepayCreditAccount struct {
	Owner common.Address
	To    common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRepayCreditAccount is a free log retrieval operation binding the contract event 0xe7c7987373a0cc4913d307f23ab8ef02e0333a2af445065e2ef7636cffc6daa7.
//
// Solidity: event RepayCreditAccount(address indexed owner, address indexed to)
func (_CreditManager *CreditManagerFilterer) FilterRepayCreditAccount(opts *bind.FilterOpts, owner []common.Address, to []common.Address) (*CreditManagerRepayCreditAccountIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CreditManager.contract.FilterLogs(opts, "RepayCreditAccount", ownerRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CreditManagerRepayCreditAccountIterator{contract: _CreditManager.contract, event: "RepayCreditAccount", logs: logs, sub: sub}, nil
}

// WatchRepayCreditAccount is a free log subscription operation binding the contract event 0xe7c7987373a0cc4913d307f23ab8ef02e0333a2af445065e2ef7636cffc6daa7.
//
// Solidity: event RepayCreditAccount(address indexed owner, address indexed to)
func (_CreditManager *CreditManagerFilterer) WatchRepayCreditAccount(opts *bind.WatchOpts, sink chan<- *CreditManagerRepayCreditAccount, owner []common.Address, to []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CreditManager.contract.WatchLogs(opts, "RepayCreditAccount", ownerRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditManagerRepayCreditAccount)
				if err := _CreditManager.contract.UnpackLog(event, "RepayCreditAccount", log); err != nil {
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

// ParseRepayCreditAccount is a log parse operation binding the contract event 0xe7c7987373a0cc4913d307f23ab8ef02e0333a2af445065e2ef7636cffc6daa7.
//
// Solidity: event RepayCreditAccount(address indexed owner, address indexed to)
func (_CreditManager *CreditManagerFilterer) ParseRepayCreditAccount(log types.Log) (*CreditManagerRepayCreditAccount, error) {
	event := new(CreditManagerRepayCreditAccount)
	if err := _CreditManager.contract.UnpackLog(event, "RepayCreditAccount", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CreditManagerTransferAccountIterator is returned from FilterTransferAccount and is used to iterate over the raw logs and unpacked data for TransferAccount events raised by the CreditManager contract.
type CreditManagerTransferAccountIterator struct {
	Event *CreditManagerTransferAccount // Event containing the contract specifics and raw log

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
func (it *CreditManagerTransferAccountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditManagerTransferAccount)
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
		it.Event = new(CreditManagerTransferAccount)
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
func (it *CreditManagerTransferAccountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditManagerTransferAccountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditManagerTransferAccount represents a TransferAccount event raised by the CreditManager contract.
type CreditManagerTransferAccount struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferAccount is a free log retrieval operation binding the contract event 0x93c70cc9715bef0d83edf2095f3595402279d274f402a73ffc17f1bcb19d863d.
//
// Solidity: event TransferAccount(address indexed oldOwner, address indexed newOwner)
func (_CreditManager *CreditManagerFilterer) FilterTransferAccount(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*CreditManagerTransferAccountIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CreditManager.contract.FilterLogs(opts, "TransferAccount", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CreditManagerTransferAccountIterator{contract: _CreditManager.contract, event: "TransferAccount", logs: logs, sub: sub}, nil
}

// WatchTransferAccount is a free log subscription operation binding the contract event 0x93c70cc9715bef0d83edf2095f3595402279d274f402a73ffc17f1bcb19d863d.
//
// Solidity: event TransferAccount(address indexed oldOwner, address indexed newOwner)
func (_CreditManager *CreditManagerFilterer) WatchTransferAccount(opts *bind.WatchOpts, sink chan<- *CreditManagerTransferAccount, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CreditManager.contract.WatchLogs(opts, "TransferAccount", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditManagerTransferAccount)
				if err := _CreditManager.contract.UnpackLog(event, "TransferAccount", log); err != nil {
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

// ParseTransferAccount is a log parse operation binding the contract event 0x93c70cc9715bef0d83edf2095f3595402279d274f402a73ffc17f1bcb19d863d.
//
// Solidity: event TransferAccount(address indexed oldOwner, address indexed newOwner)
func (_CreditManager *CreditManagerFilterer) ParseTransferAccount(log types.Log) (*CreditManagerTransferAccount, error) {
	event := new(CreditManagerTransferAccount)
	if err := _CreditManager.contract.UnpackLog(event, "TransferAccount", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CreditManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the CreditManager contract.
type CreditManagerUnpausedIterator struct {
	Event *CreditManagerUnpaused // Event containing the contract specifics and raw log

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
func (it *CreditManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditManagerUnpaused)
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
		it.Event = new(CreditManagerUnpaused)
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
func (it *CreditManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditManagerUnpaused represents a Unpaused event raised by the CreditManager contract.
type CreditManagerUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_CreditManager *CreditManagerFilterer) FilterUnpaused(opts *bind.FilterOpts) (*CreditManagerUnpausedIterator, error) {

	logs, sub, err := _CreditManager.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &CreditManagerUnpausedIterator{contract: _CreditManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_CreditManager *CreditManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *CreditManagerUnpaused) (event.Subscription, error) {

	logs, sub, err := _CreditManager.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditManagerUnpaused)
				if err := _CreditManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_CreditManager *CreditManagerFilterer) ParseUnpaused(log types.Log) (*CreditManagerUnpaused, error) {
	event := new(CreditManagerUnpaused)
	if err := _CreditManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	return event, nil
}
