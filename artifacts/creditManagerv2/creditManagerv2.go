// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package creditManagerv2

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

// CreditManagerv2ABI is the input ABI used to generate the binding from.
const CreditManagerv2ABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_poolService\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"ExecuteOrder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newConfigurator\",\"type\":\"address\"}],\"name\":\"NewConfigurator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"adapterToContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"addCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"addToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allowedTokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allowedTokensCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"targetContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"approveCreditAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"}],\"name\":\"calcCreditAccountAccruedInterest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"borrowedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowedAmountWithInterest\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"adapter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"targetContract\",\"type\":\"address\"}],\"name\":\"changeContractAllowance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"}],\"name\":\"checkAndEnableToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chiThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isLiquidated\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"totalValue\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"sendAllAssets\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"convertWETH\",\"type\":\"bool\"}],\"name\":\"closeCreditAccount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"remainingFunds\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"creditAccounts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creditConfigurator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creditFacade\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"enabledTokensMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"targetContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeOrder\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"fastCheckCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balanceInBefore\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balanceOutBefore\",\"type\":\"uint256\"}],\"name\":\"fastCollateralCheck\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeInterest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeLiquidation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forbidenTokenMask\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"}],\"name\":\"fullCollateralCheck\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"}],\"name\":\"getCreditAccountOrRevert\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"result\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hfCheckInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidationDiscount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"liquidationThresholds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"increase\",\"type\":\"bool\"}],\"name\":\"manageDebt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"newBorrowedAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxBorrowedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minBorrowedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"borrowedAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"}],\"name\":\"openCreditAccount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolService\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceOracle\",\"outputs\":[{\"internalType\":\"contractIPriceOracle\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_creditConfigurator\",\"type\":\"address\"}],\"name\":\"setConfigurator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_forbidMask\",\"type\":\"uint256\"}],\"name\":\"setForbidMask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidationThreshold\",\"type\":\"uint256\"}],\"name\":\"setLiquidationThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minBorrowedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxBorrowedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_feeInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_feeLiquidation\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_liquidationDiscount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_chiThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_hfCheckInterval\",\"type\":\"uint256\"}],\"name\":\"setParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenMasksMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferAccountOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlyingToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_creditFacade\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_priceOracle\",\"type\":\"address\"}],\"name\":\"upgradeContracts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wethAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wethGateway\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// CreditManagerv2 is an auto generated Go binding around an Ethereum contract.
type CreditManagerv2 struct {
	CreditManagerv2Caller     // Read-only binding to the contract
	CreditManagerv2Transactor // Write-only binding to the contract
	CreditManagerv2Filterer   // Log filterer for contract events
}

// CreditManagerv2Caller is an auto generated read-only Go binding around an Ethereum contract.
type CreditManagerv2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditManagerv2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type CreditManagerv2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditManagerv2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CreditManagerv2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditManagerv2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CreditManagerv2Session struct {
	Contract     *CreditManagerv2  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CreditManagerv2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CreditManagerv2CallerSession struct {
	Contract *CreditManagerv2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// CreditManagerv2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CreditManagerv2TransactorSession struct {
	Contract     *CreditManagerv2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// CreditManagerv2Raw is an auto generated low-level Go binding around an Ethereum contract.
type CreditManagerv2Raw struct {
	Contract *CreditManagerv2 // Generic contract binding to access the raw methods on
}

// CreditManagerv2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CreditManagerv2CallerRaw struct {
	Contract *CreditManagerv2Caller // Generic read-only contract binding to access the raw methods on
}

// CreditManagerv2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CreditManagerv2TransactorRaw struct {
	Contract *CreditManagerv2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewCreditManagerv2 creates a new instance of CreditManagerv2, bound to a specific deployed contract.
func NewCreditManagerv2(address common.Address, backend bind.ContractBackend) (*CreditManagerv2, error) {
	contract, err := bindCreditManagerv2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CreditManagerv2{CreditManagerv2Caller: CreditManagerv2Caller{contract: contract}, CreditManagerv2Transactor: CreditManagerv2Transactor{contract: contract}, CreditManagerv2Filterer: CreditManagerv2Filterer{contract: contract}}, nil
}

// NewCreditManagerv2Caller creates a new read-only instance of CreditManagerv2, bound to a specific deployed contract.
func NewCreditManagerv2Caller(address common.Address, caller bind.ContractCaller) (*CreditManagerv2Caller, error) {
	contract, err := bindCreditManagerv2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CreditManagerv2Caller{contract: contract}, nil
}

// NewCreditManagerv2Transactor creates a new write-only instance of CreditManagerv2, bound to a specific deployed contract.
func NewCreditManagerv2Transactor(address common.Address, transactor bind.ContractTransactor) (*CreditManagerv2Transactor, error) {
	contract, err := bindCreditManagerv2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CreditManagerv2Transactor{contract: contract}, nil
}

// NewCreditManagerv2Filterer creates a new log filterer instance of CreditManagerv2, bound to a specific deployed contract.
func NewCreditManagerv2Filterer(address common.Address, filterer bind.ContractFilterer) (*CreditManagerv2Filterer, error) {
	contract, err := bindCreditManagerv2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CreditManagerv2Filterer{contract: contract}, nil
}

// bindCreditManagerv2 binds a generic wrapper to an already deployed contract.
func bindCreditManagerv2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CreditManagerv2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CreditManagerv2 *CreditManagerv2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CreditManagerv2.Contract.CreditManagerv2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CreditManagerv2 *CreditManagerv2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.CreditManagerv2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CreditManagerv2 *CreditManagerv2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.CreditManagerv2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CreditManagerv2 *CreditManagerv2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CreditManagerv2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CreditManagerv2 *CreditManagerv2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CreditManagerv2 *CreditManagerv2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.contract.Transact(opts, method, params...)
}

// AdapterToContract is a free data retrieval call binding the contract method 0xff687543.
//
// Solidity: function adapterToContract(address ) view returns(address)
func (_CreditManagerv2 *CreditManagerv2Caller) AdapterToContract(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _CreditManagerv2.contract.Call(opts, &out, "adapterToContract", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AdapterToContract is a free data retrieval call binding the contract method 0xff687543.
//
// Solidity: function adapterToContract(address ) view returns(address)
func (_CreditManagerv2 *CreditManagerv2Session) AdapterToContract(arg0 common.Address) (common.Address, error) {
	return _CreditManagerv2.Contract.AdapterToContract(&_CreditManagerv2.CallOpts, arg0)
}

// AdapterToContract is a free data retrieval call binding the contract method 0xff687543.
//
// Solidity: function adapterToContract(address ) view returns(address)
func (_CreditManagerv2 *CreditManagerv2CallerSession) AdapterToContract(arg0 common.Address) (common.Address, error) {
	return _CreditManagerv2.Contract.AdapterToContract(&_CreditManagerv2.CallOpts, arg0)
}

// AllowedTokens is a free data retrieval call binding the contract method 0x5e5f2e26.
//
// Solidity: function allowedTokens(uint256 ) view returns(address)
func (_CreditManagerv2 *CreditManagerv2Caller) AllowedTokens(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CreditManagerv2.contract.Call(opts, &out, "allowedTokens", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllowedTokens is a free data retrieval call binding the contract method 0x5e5f2e26.
//
// Solidity: function allowedTokens(uint256 ) view returns(address)
func (_CreditManagerv2 *CreditManagerv2Session) AllowedTokens(arg0 *big.Int) (common.Address, error) {
	return _CreditManagerv2.Contract.AllowedTokens(&_CreditManagerv2.CallOpts, arg0)
}

// AllowedTokens is a free data retrieval call binding the contract method 0x5e5f2e26.
//
// Solidity: function allowedTokens(uint256 ) view returns(address)
func (_CreditManagerv2 *CreditManagerv2CallerSession) AllowedTokens(arg0 *big.Int) (common.Address, error) {
	return _CreditManagerv2.Contract.AllowedTokens(&_CreditManagerv2.CallOpts, arg0)
}

// AllowedTokensCount is a free data retrieval call binding the contract method 0x20a05ff7.
//
// Solidity: function allowedTokensCount() view returns(uint256)
func (_CreditManagerv2 *CreditManagerv2Caller) AllowedTokensCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CreditManagerv2.contract.Call(opts, &out, "allowedTokensCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllowedTokensCount is a free data retrieval call binding the contract method 0x20a05ff7.
//
// Solidity: function allowedTokensCount() view returns(uint256)
func (_CreditManagerv2 *CreditManagerv2Session) AllowedTokensCount() (*big.Int, error) {
	return _CreditManagerv2.Contract.AllowedTokensCount(&_CreditManagerv2.CallOpts)
}

// AllowedTokensCount is a free data retrieval call binding the contract method 0x20a05ff7.
//
// Solidity: function allowedTokensCount() view returns(uint256)
func (_CreditManagerv2 *CreditManagerv2CallerSession) AllowedTokensCount() (*big.Int, error) {
	return _CreditManagerv2.Contract.AllowedTokensCount(&_CreditManagerv2.CallOpts)
}

// CalcCreditAccountAccruedInterest is a free data retrieval call binding the contract method 0x3192195c.
//
// Solidity: function calcCreditAccountAccruedInterest(address creditAccount) view returns(uint256 borrowedAmount, uint256 borrowedAmountWithInterest)
func (_CreditManagerv2 *CreditManagerv2Caller) CalcCreditAccountAccruedInterest(opts *bind.CallOpts, creditAccount common.Address) (struct {
	BorrowedAmount             *big.Int
	BorrowedAmountWithInterest *big.Int
}, error) {
	var out []interface{}
	err := _CreditManagerv2.contract.Call(opts, &out, "calcCreditAccountAccruedInterest", creditAccount)

	outstruct := new(struct {
		BorrowedAmount             *big.Int
		BorrowedAmountWithInterest *big.Int
	})

	outstruct.BorrowedAmount = out[0].(*big.Int)
	outstruct.BorrowedAmountWithInterest = out[1].(*big.Int)

	return *outstruct, err

}

// CalcCreditAccountAccruedInterest is a free data retrieval call binding the contract method 0x3192195c.
//
// Solidity: function calcCreditAccountAccruedInterest(address creditAccount) view returns(uint256 borrowedAmount, uint256 borrowedAmountWithInterest)
func (_CreditManagerv2 *CreditManagerv2Session) CalcCreditAccountAccruedInterest(creditAccount common.Address) (struct {
	BorrowedAmount             *big.Int
	BorrowedAmountWithInterest *big.Int
}, error) {
	return _CreditManagerv2.Contract.CalcCreditAccountAccruedInterest(&_CreditManagerv2.CallOpts, creditAccount)
}

// CalcCreditAccountAccruedInterest is a free data retrieval call binding the contract method 0x3192195c.
//
// Solidity: function calcCreditAccountAccruedInterest(address creditAccount) view returns(uint256 borrowedAmount, uint256 borrowedAmountWithInterest)
func (_CreditManagerv2 *CreditManagerv2CallerSession) CalcCreditAccountAccruedInterest(creditAccount common.Address) (struct {
	BorrowedAmount             *big.Int
	BorrowedAmountWithInterest *big.Int
}, error) {
	return _CreditManagerv2.Contract.CalcCreditAccountAccruedInterest(&_CreditManagerv2.CallOpts, creditAccount)
}

// ChiThreshold is a free data retrieval call binding the contract method 0x47dedfc9.
//
// Solidity: function chiThreshold() view returns(uint256)
func (_CreditManagerv2 *CreditManagerv2Caller) ChiThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CreditManagerv2.contract.Call(opts, &out, "chiThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChiThreshold is a free data retrieval call binding the contract method 0x47dedfc9.
//
// Solidity: function chiThreshold() view returns(uint256)
func (_CreditManagerv2 *CreditManagerv2Session) ChiThreshold() (*big.Int, error) {
	return _CreditManagerv2.Contract.ChiThreshold(&_CreditManagerv2.CallOpts)
}

// ChiThreshold is a free data retrieval call binding the contract method 0x47dedfc9.
//
// Solidity: function chiThreshold() view returns(uint256)
func (_CreditManagerv2 *CreditManagerv2CallerSession) ChiThreshold() (*big.Int, error) {
	return _CreditManagerv2.Contract.ChiThreshold(&_CreditManagerv2.CallOpts)
}

// CreditAccounts is a free data retrieval call binding the contract method 0x055ee9b5.
//
// Solidity: function creditAccounts(address ) view returns(address)
func (_CreditManagerv2 *CreditManagerv2Caller) CreditAccounts(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _CreditManagerv2.contract.Call(opts, &out, "creditAccounts", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditAccounts is a free data retrieval call binding the contract method 0x055ee9b5.
//
// Solidity: function creditAccounts(address ) view returns(address)
func (_CreditManagerv2 *CreditManagerv2Session) CreditAccounts(arg0 common.Address) (common.Address, error) {
	return _CreditManagerv2.Contract.CreditAccounts(&_CreditManagerv2.CallOpts, arg0)
}

// CreditAccounts is a free data retrieval call binding the contract method 0x055ee9b5.
//
// Solidity: function creditAccounts(address ) view returns(address)
func (_CreditManagerv2 *CreditManagerv2CallerSession) CreditAccounts(arg0 common.Address) (common.Address, error) {
	return _CreditManagerv2.Contract.CreditAccounts(&_CreditManagerv2.CallOpts, arg0)
}

// CreditConfigurator is a free data retrieval call binding the contract method 0xf9aa028a.
//
// Solidity: function creditConfigurator() view returns(address)
func (_CreditManagerv2 *CreditManagerv2Caller) CreditConfigurator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditManagerv2.contract.Call(opts, &out, "creditConfigurator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditConfigurator is a free data retrieval call binding the contract method 0xf9aa028a.
//
// Solidity: function creditConfigurator() view returns(address)
func (_CreditManagerv2 *CreditManagerv2Session) CreditConfigurator() (common.Address, error) {
	return _CreditManagerv2.Contract.CreditConfigurator(&_CreditManagerv2.CallOpts)
}

// CreditConfigurator is a free data retrieval call binding the contract method 0xf9aa028a.
//
// Solidity: function creditConfigurator() view returns(address)
func (_CreditManagerv2 *CreditManagerv2CallerSession) CreditConfigurator() (common.Address, error) {
	return _CreditManagerv2.Contract.CreditConfigurator(&_CreditManagerv2.CallOpts)
}

// CreditFacade is a free data retrieval call binding the contract method 0x2f7a1881.
//
// Solidity: function creditFacade() view returns(address)
func (_CreditManagerv2 *CreditManagerv2Caller) CreditFacade(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditManagerv2.contract.Call(opts, &out, "creditFacade")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditFacade is a free data retrieval call binding the contract method 0x2f7a1881.
//
// Solidity: function creditFacade() view returns(address)
func (_CreditManagerv2 *CreditManagerv2Session) CreditFacade() (common.Address, error) {
	return _CreditManagerv2.Contract.CreditFacade(&_CreditManagerv2.CallOpts)
}

// CreditFacade is a free data retrieval call binding the contract method 0x2f7a1881.
//
// Solidity: function creditFacade() view returns(address)
func (_CreditManagerv2 *CreditManagerv2CallerSession) CreditFacade() (common.Address, error) {
	return _CreditManagerv2.Contract.CreditFacade(&_CreditManagerv2.CallOpts)
}

// EnabledTokensMap is a free data retrieval call binding the contract method 0x8991b2f1.
//
// Solidity: function enabledTokensMap(address ) view returns(uint256)
func (_CreditManagerv2 *CreditManagerv2Caller) EnabledTokensMap(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CreditManagerv2.contract.Call(opts, &out, "enabledTokensMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EnabledTokensMap is a free data retrieval call binding the contract method 0x8991b2f1.
//
// Solidity: function enabledTokensMap(address ) view returns(uint256)
func (_CreditManagerv2 *CreditManagerv2Session) EnabledTokensMap(arg0 common.Address) (*big.Int, error) {
	return _CreditManagerv2.Contract.EnabledTokensMap(&_CreditManagerv2.CallOpts, arg0)
}

// EnabledTokensMap is a free data retrieval call binding the contract method 0x8991b2f1.
//
// Solidity: function enabledTokensMap(address ) view returns(uint256)
func (_CreditManagerv2 *CreditManagerv2CallerSession) EnabledTokensMap(arg0 common.Address) (*big.Int, error) {
	return _CreditManagerv2.Contract.EnabledTokensMap(&_CreditManagerv2.CallOpts, arg0)
}

// FastCheckCounter is a free data retrieval call binding the contract method 0x4cba294a.
//
// Solidity: function fastCheckCounter(address ) view returns(uint256)
func (_CreditManagerv2 *CreditManagerv2Caller) FastCheckCounter(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CreditManagerv2.contract.Call(opts, &out, "fastCheckCounter", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FastCheckCounter is a free data retrieval call binding the contract method 0x4cba294a.
//
// Solidity: function fastCheckCounter(address ) view returns(uint256)
func (_CreditManagerv2 *CreditManagerv2Session) FastCheckCounter(arg0 common.Address) (*big.Int, error) {
	return _CreditManagerv2.Contract.FastCheckCounter(&_CreditManagerv2.CallOpts, arg0)
}

// FastCheckCounter is a free data retrieval call binding the contract method 0x4cba294a.
//
// Solidity: function fastCheckCounter(address ) view returns(uint256)
func (_CreditManagerv2 *CreditManagerv2CallerSession) FastCheckCounter(arg0 common.Address) (*big.Int, error) {
	return _CreditManagerv2.Contract.FastCheckCounter(&_CreditManagerv2.CallOpts, arg0)
}

// FeeInterest is a free data retrieval call binding the contract method 0x5e0b63d3.
//
// Solidity: function feeInterest() view returns(uint256)
func (_CreditManagerv2 *CreditManagerv2Caller) FeeInterest(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CreditManagerv2.contract.Call(opts, &out, "feeInterest")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeInterest is a free data retrieval call binding the contract method 0x5e0b63d3.
//
// Solidity: function feeInterest() view returns(uint256)
func (_CreditManagerv2 *CreditManagerv2Session) FeeInterest() (*big.Int, error) {
	return _CreditManagerv2.Contract.FeeInterest(&_CreditManagerv2.CallOpts)
}

// FeeInterest is a free data retrieval call binding the contract method 0x5e0b63d3.
//
// Solidity: function feeInterest() view returns(uint256)
func (_CreditManagerv2 *CreditManagerv2CallerSession) FeeInterest() (*big.Int, error) {
	return _CreditManagerv2.Contract.FeeInterest(&_CreditManagerv2.CallOpts)
}

// FeeLiquidation is a free data retrieval call binding the contract method 0x3915ffaa.
//
// Solidity: function feeLiquidation() view returns(uint256)
func (_CreditManagerv2 *CreditManagerv2Caller) FeeLiquidation(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CreditManagerv2.contract.Call(opts, &out, "feeLiquidation")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeLiquidation is a free data retrieval call binding the contract method 0x3915ffaa.
//
// Solidity: function feeLiquidation() view returns(uint256)
func (_CreditManagerv2 *CreditManagerv2Session) FeeLiquidation() (*big.Int, error) {
	return _CreditManagerv2.Contract.FeeLiquidation(&_CreditManagerv2.CallOpts)
}

// FeeLiquidation is a free data retrieval call binding the contract method 0x3915ffaa.
//
// Solidity: function feeLiquidation() view returns(uint256)
func (_CreditManagerv2 *CreditManagerv2CallerSession) FeeLiquidation() (*big.Int, error) {
	return _CreditManagerv2.Contract.FeeLiquidation(&_CreditManagerv2.CallOpts)
}

// ForbidenTokenMask is a free data retrieval call binding the contract method 0xe46c95bc.
//
// Solidity: function forbidenTokenMask() view returns(uint256)
func (_CreditManagerv2 *CreditManagerv2Caller) ForbidenTokenMask(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CreditManagerv2.contract.Call(opts, &out, "forbidenTokenMask")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ForbidenTokenMask is a free data retrieval call binding the contract method 0xe46c95bc.
//
// Solidity: function forbidenTokenMask() view returns(uint256)
func (_CreditManagerv2 *CreditManagerv2Session) ForbidenTokenMask() (*big.Int, error) {
	return _CreditManagerv2.Contract.ForbidenTokenMask(&_CreditManagerv2.CallOpts)
}

// ForbidenTokenMask is a free data retrieval call binding the contract method 0xe46c95bc.
//
// Solidity: function forbidenTokenMask() view returns(uint256)
func (_CreditManagerv2 *CreditManagerv2CallerSession) ForbidenTokenMask() (*big.Int, error) {
	return _CreditManagerv2.Contract.ForbidenTokenMask(&_CreditManagerv2.CallOpts)
}

// GetCreditAccountOrRevert is a free data retrieval call binding the contract method 0xe958b704.
//
// Solidity: function getCreditAccountOrRevert(address borrower) view returns(address result)
func (_CreditManagerv2 *CreditManagerv2Caller) GetCreditAccountOrRevert(opts *bind.CallOpts, borrower common.Address) (common.Address, error) {
	var out []interface{}
	err := _CreditManagerv2.contract.Call(opts, &out, "getCreditAccountOrRevert", borrower)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCreditAccountOrRevert is a free data retrieval call binding the contract method 0xe958b704.
//
// Solidity: function getCreditAccountOrRevert(address borrower) view returns(address result)
func (_CreditManagerv2 *CreditManagerv2Session) GetCreditAccountOrRevert(borrower common.Address) (common.Address, error) {
	return _CreditManagerv2.Contract.GetCreditAccountOrRevert(&_CreditManagerv2.CallOpts, borrower)
}

// GetCreditAccountOrRevert is a free data retrieval call binding the contract method 0xe958b704.
//
// Solidity: function getCreditAccountOrRevert(address borrower) view returns(address result)
func (_CreditManagerv2 *CreditManagerv2CallerSession) GetCreditAccountOrRevert(borrower common.Address) (common.Address, error) {
	return _CreditManagerv2.Contract.GetCreditAccountOrRevert(&_CreditManagerv2.CallOpts, borrower)
}

// HfCheckInterval is a free data retrieval call binding the contract method 0xe6dee2cc.
//
// Solidity: function hfCheckInterval() view returns(uint256)
func (_CreditManagerv2 *CreditManagerv2Caller) HfCheckInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CreditManagerv2.contract.Call(opts, &out, "hfCheckInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HfCheckInterval is a free data retrieval call binding the contract method 0xe6dee2cc.
//
// Solidity: function hfCheckInterval() view returns(uint256)
func (_CreditManagerv2 *CreditManagerv2Session) HfCheckInterval() (*big.Int, error) {
	return _CreditManagerv2.Contract.HfCheckInterval(&_CreditManagerv2.CallOpts)
}

// HfCheckInterval is a free data retrieval call binding the contract method 0xe6dee2cc.
//
// Solidity: function hfCheckInterval() view returns(uint256)
func (_CreditManagerv2 *CreditManagerv2CallerSession) HfCheckInterval() (*big.Int, error) {
	return _CreditManagerv2.Contract.HfCheckInterval(&_CreditManagerv2.CallOpts)
}

// LiquidationDiscount is a free data retrieval call binding the contract method 0x8053fcbe.
//
// Solidity: function liquidationDiscount() view returns(uint256)
func (_CreditManagerv2 *CreditManagerv2Caller) LiquidationDiscount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CreditManagerv2.contract.Call(opts, &out, "liquidationDiscount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LiquidationDiscount is a free data retrieval call binding the contract method 0x8053fcbe.
//
// Solidity: function liquidationDiscount() view returns(uint256)
func (_CreditManagerv2 *CreditManagerv2Session) LiquidationDiscount() (*big.Int, error) {
	return _CreditManagerv2.Contract.LiquidationDiscount(&_CreditManagerv2.CallOpts)
}

// LiquidationDiscount is a free data retrieval call binding the contract method 0x8053fcbe.
//
// Solidity: function liquidationDiscount() view returns(uint256)
func (_CreditManagerv2 *CreditManagerv2CallerSession) LiquidationDiscount() (*big.Int, error) {
	return _CreditManagerv2.Contract.LiquidationDiscount(&_CreditManagerv2.CallOpts)
}

// LiquidationThresholds is a free data retrieval call binding the contract method 0x78327438.
//
// Solidity: function liquidationThresholds(address ) view returns(uint256)
func (_CreditManagerv2 *CreditManagerv2Caller) LiquidationThresholds(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CreditManagerv2.contract.Call(opts, &out, "liquidationThresholds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LiquidationThresholds is a free data retrieval call binding the contract method 0x78327438.
//
// Solidity: function liquidationThresholds(address ) view returns(uint256)
func (_CreditManagerv2 *CreditManagerv2Session) LiquidationThresholds(arg0 common.Address) (*big.Int, error) {
	return _CreditManagerv2.Contract.LiquidationThresholds(&_CreditManagerv2.CallOpts, arg0)
}

// LiquidationThresholds is a free data retrieval call binding the contract method 0x78327438.
//
// Solidity: function liquidationThresholds(address ) view returns(uint256)
func (_CreditManagerv2 *CreditManagerv2CallerSession) LiquidationThresholds(arg0 common.Address) (*big.Int, error) {
	return _CreditManagerv2.Contract.LiquidationThresholds(&_CreditManagerv2.CallOpts, arg0)
}

// MaxBorrowedAmount is a free data retrieval call binding the contract method 0x62186905.
//
// Solidity: function maxBorrowedAmount() view returns(uint256)
func (_CreditManagerv2 *CreditManagerv2Caller) MaxBorrowedAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CreditManagerv2.contract.Call(opts, &out, "maxBorrowedAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxBorrowedAmount is a free data retrieval call binding the contract method 0x62186905.
//
// Solidity: function maxBorrowedAmount() view returns(uint256)
func (_CreditManagerv2 *CreditManagerv2Session) MaxBorrowedAmount() (*big.Int, error) {
	return _CreditManagerv2.Contract.MaxBorrowedAmount(&_CreditManagerv2.CallOpts)
}

// MaxBorrowedAmount is a free data retrieval call binding the contract method 0x62186905.
//
// Solidity: function maxBorrowedAmount() view returns(uint256)
func (_CreditManagerv2 *CreditManagerv2CallerSession) MaxBorrowedAmount() (*big.Int, error) {
	return _CreditManagerv2.Contract.MaxBorrowedAmount(&_CreditManagerv2.CallOpts)
}

// MinBorrowedAmount is a free data retrieval call binding the contract method 0x0bc772da.
//
// Solidity: function minBorrowedAmount() view returns(uint256)
func (_CreditManagerv2 *CreditManagerv2Caller) MinBorrowedAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CreditManagerv2.contract.Call(opts, &out, "minBorrowedAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinBorrowedAmount is a free data retrieval call binding the contract method 0x0bc772da.
//
// Solidity: function minBorrowedAmount() view returns(uint256)
func (_CreditManagerv2 *CreditManagerv2Session) MinBorrowedAmount() (*big.Int, error) {
	return _CreditManagerv2.Contract.MinBorrowedAmount(&_CreditManagerv2.CallOpts)
}

// MinBorrowedAmount is a free data retrieval call binding the contract method 0x0bc772da.
//
// Solidity: function minBorrowedAmount() view returns(uint256)
func (_CreditManagerv2 *CreditManagerv2CallerSession) MinBorrowedAmount() (*big.Int, error) {
	return _CreditManagerv2.Contract.MinBorrowedAmount(&_CreditManagerv2.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CreditManagerv2 *CreditManagerv2Caller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CreditManagerv2.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CreditManagerv2 *CreditManagerv2Session) Paused() (bool, error) {
	return _CreditManagerv2.Contract.Paused(&_CreditManagerv2.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CreditManagerv2 *CreditManagerv2CallerSession) Paused() (bool, error) {
	return _CreditManagerv2.Contract.Paused(&_CreditManagerv2.CallOpts)
}

// PoolService is a free data retrieval call binding the contract method 0x570a7af2.
//
// Solidity: function poolService() view returns(address)
func (_CreditManagerv2 *CreditManagerv2Caller) PoolService(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditManagerv2.contract.Call(opts, &out, "poolService")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolService is a free data retrieval call binding the contract method 0x570a7af2.
//
// Solidity: function poolService() view returns(address)
func (_CreditManagerv2 *CreditManagerv2Session) PoolService() (common.Address, error) {
	return _CreditManagerv2.Contract.PoolService(&_CreditManagerv2.CallOpts)
}

// PoolService is a free data retrieval call binding the contract method 0x570a7af2.
//
// Solidity: function poolService() view returns(address)
func (_CreditManagerv2 *CreditManagerv2CallerSession) PoolService() (common.Address, error) {
	return _CreditManagerv2.Contract.PoolService(&_CreditManagerv2.CallOpts)
}

// PriceOracle is a free data retrieval call binding the contract method 0x2630c12f.
//
// Solidity: function priceOracle() view returns(address)
func (_CreditManagerv2 *CreditManagerv2Caller) PriceOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditManagerv2.contract.Call(opts, &out, "priceOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PriceOracle is a free data retrieval call binding the contract method 0x2630c12f.
//
// Solidity: function priceOracle() view returns(address)
func (_CreditManagerv2 *CreditManagerv2Session) PriceOracle() (common.Address, error) {
	return _CreditManagerv2.Contract.PriceOracle(&_CreditManagerv2.CallOpts)
}

// PriceOracle is a free data retrieval call binding the contract method 0x2630c12f.
//
// Solidity: function priceOracle() view returns(address)
func (_CreditManagerv2 *CreditManagerv2CallerSession) PriceOracle() (common.Address, error) {
	return _CreditManagerv2.Contract.PriceOracle(&_CreditManagerv2.CallOpts)
}

// TokenMasksMap is a free data retrieval call binding the contract method 0xf67c5bd0.
//
// Solidity: function tokenMasksMap(address ) view returns(uint256)
func (_CreditManagerv2 *CreditManagerv2Caller) TokenMasksMap(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CreditManagerv2.contract.Call(opts, &out, "tokenMasksMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenMasksMap is a free data retrieval call binding the contract method 0xf67c5bd0.
//
// Solidity: function tokenMasksMap(address ) view returns(uint256)
func (_CreditManagerv2 *CreditManagerv2Session) TokenMasksMap(arg0 common.Address) (*big.Int, error) {
	return _CreditManagerv2.Contract.TokenMasksMap(&_CreditManagerv2.CallOpts, arg0)
}

// TokenMasksMap is a free data retrieval call binding the contract method 0xf67c5bd0.
//
// Solidity: function tokenMasksMap(address ) view returns(uint256)
func (_CreditManagerv2 *CreditManagerv2CallerSession) TokenMasksMap(arg0 common.Address) (*big.Int, error) {
	return _CreditManagerv2.Contract.TokenMasksMap(&_CreditManagerv2.CallOpts, arg0)
}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_CreditManagerv2 *CreditManagerv2Caller) UnderlyingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditManagerv2.contract.Call(opts, &out, "underlyingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_CreditManagerv2 *CreditManagerv2Session) UnderlyingToken() (common.Address, error) {
	return _CreditManagerv2.Contract.UnderlyingToken(&_CreditManagerv2.CallOpts)
}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_CreditManagerv2 *CreditManagerv2CallerSession) UnderlyingToken() (common.Address, error) {
	return _CreditManagerv2.Contract.UnderlyingToken(&_CreditManagerv2.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_CreditManagerv2 *CreditManagerv2Caller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CreditManagerv2.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_CreditManagerv2 *CreditManagerv2Session) Version() (*big.Int, error) {
	return _CreditManagerv2.Contract.Version(&_CreditManagerv2.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_CreditManagerv2 *CreditManagerv2CallerSession) Version() (*big.Int, error) {
	return _CreditManagerv2.Contract.Version(&_CreditManagerv2.CallOpts)
}

// WethAddress is a free data retrieval call binding the contract method 0x4f0e0ef3.
//
// Solidity: function wethAddress() view returns(address)
func (_CreditManagerv2 *CreditManagerv2Caller) WethAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditManagerv2.contract.Call(opts, &out, "wethAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WethAddress is a free data retrieval call binding the contract method 0x4f0e0ef3.
//
// Solidity: function wethAddress() view returns(address)
func (_CreditManagerv2 *CreditManagerv2Session) WethAddress() (common.Address, error) {
	return _CreditManagerv2.Contract.WethAddress(&_CreditManagerv2.CallOpts)
}

// WethAddress is a free data retrieval call binding the contract method 0x4f0e0ef3.
//
// Solidity: function wethAddress() view returns(address)
func (_CreditManagerv2 *CreditManagerv2CallerSession) WethAddress() (common.Address, error) {
	return _CreditManagerv2.Contract.WethAddress(&_CreditManagerv2.CallOpts)
}

// WethGateway is a free data retrieval call binding the contract method 0xc5e10eef.
//
// Solidity: function wethGateway() view returns(address)
func (_CreditManagerv2 *CreditManagerv2Caller) WethGateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditManagerv2.contract.Call(opts, &out, "wethGateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WethGateway is a free data retrieval call binding the contract method 0xc5e10eef.
//
// Solidity: function wethGateway() view returns(address)
func (_CreditManagerv2 *CreditManagerv2Session) WethGateway() (common.Address, error) {
	return _CreditManagerv2.Contract.WethGateway(&_CreditManagerv2.CallOpts)
}

// WethGateway is a free data retrieval call binding the contract method 0xc5e10eef.
//
// Solidity: function wethGateway() view returns(address)
func (_CreditManagerv2 *CreditManagerv2CallerSession) WethGateway() (common.Address, error) {
	return _CreditManagerv2.Contract.WethGateway(&_CreditManagerv2.CallOpts)
}

// AddCollateral is a paid mutator transaction binding the contract method 0x830aa745.
//
// Solidity: function addCollateral(address payer, address onBehalfOf, address token, uint256 amount) returns()
func (_CreditManagerv2 *CreditManagerv2Transactor) AddCollateral(opts *bind.TransactOpts, payer common.Address, onBehalfOf common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CreditManagerv2.contract.Transact(opts, "addCollateral", payer, onBehalfOf, token, amount)
}

// AddCollateral is a paid mutator transaction binding the contract method 0x830aa745.
//
// Solidity: function addCollateral(address payer, address onBehalfOf, address token, uint256 amount) returns()
func (_CreditManagerv2 *CreditManagerv2Session) AddCollateral(payer common.Address, onBehalfOf common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.AddCollateral(&_CreditManagerv2.TransactOpts, payer, onBehalfOf, token, amount)
}

// AddCollateral is a paid mutator transaction binding the contract method 0x830aa745.
//
// Solidity: function addCollateral(address payer, address onBehalfOf, address token, uint256 amount) returns()
func (_CreditManagerv2 *CreditManagerv2TransactorSession) AddCollateral(payer common.Address, onBehalfOf common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.AddCollateral(&_CreditManagerv2.TransactOpts, payer, onBehalfOf, token, amount)
}

// AddToken is a paid mutator transaction binding the contract method 0xd48bfca7.
//
// Solidity: function addToken(address token) returns()
func (_CreditManagerv2 *CreditManagerv2Transactor) AddToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _CreditManagerv2.contract.Transact(opts, "addToken", token)
}

// AddToken is a paid mutator transaction binding the contract method 0xd48bfca7.
//
// Solidity: function addToken(address token) returns()
func (_CreditManagerv2 *CreditManagerv2Session) AddToken(token common.Address) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.AddToken(&_CreditManagerv2.TransactOpts, token)
}

// AddToken is a paid mutator transaction binding the contract method 0xd48bfca7.
//
// Solidity: function addToken(address token) returns()
func (_CreditManagerv2 *CreditManagerv2TransactorSession) AddToken(token common.Address) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.AddToken(&_CreditManagerv2.TransactOpts, token)
}

// ApproveCreditAccount is a paid mutator transaction binding the contract method 0x32a4e67c.
//
// Solidity: function approveCreditAccount(address borrower, address targetContract, address token) returns()
func (_CreditManagerv2 *CreditManagerv2Transactor) ApproveCreditAccount(opts *bind.TransactOpts, borrower common.Address, targetContract common.Address, token common.Address) (*types.Transaction, error) {
	return _CreditManagerv2.contract.Transact(opts, "approveCreditAccount", borrower, targetContract, token)
}

// ApproveCreditAccount is a paid mutator transaction binding the contract method 0x32a4e67c.
//
// Solidity: function approveCreditAccount(address borrower, address targetContract, address token) returns()
func (_CreditManagerv2 *CreditManagerv2Session) ApproveCreditAccount(borrower common.Address, targetContract common.Address, token common.Address) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.ApproveCreditAccount(&_CreditManagerv2.TransactOpts, borrower, targetContract, token)
}

// ApproveCreditAccount is a paid mutator transaction binding the contract method 0x32a4e67c.
//
// Solidity: function approveCreditAccount(address borrower, address targetContract, address token) returns()
func (_CreditManagerv2 *CreditManagerv2TransactorSession) ApproveCreditAccount(borrower common.Address, targetContract common.Address, token common.Address) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.ApproveCreditAccount(&_CreditManagerv2.TransactOpts, borrower, targetContract, token)
}

// ChangeContractAllowance is a paid mutator transaction binding the contract method 0x6e98e5e4.
//
// Solidity: function changeContractAllowance(address adapter, address targetContract) returns()
func (_CreditManagerv2 *CreditManagerv2Transactor) ChangeContractAllowance(opts *bind.TransactOpts, adapter common.Address, targetContract common.Address) (*types.Transaction, error) {
	return _CreditManagerv2.contract.Transact(opts, "changeContractAllowance", adapter, targetContract)
}

// ChangeContractAllowance is a paid mutator transaction binding the contract method 0x6e98e5e4.
//
// Solidity: function changeContractAllowance(address adapter, address targetContract) returns()
func (_CreditManagerv2 *CreditManagerv2Session) ChangeContractAllowance(adapter common.Address, targetContract common.Address) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.ChangeContractAllowance(&_CreditManagerv2.TransactOpts, adapter, targetContract)
}

// ChangeContractAllowance is a paid mutator transaction binding the contract method 0x6e98e5e4.
//
// Solidity: function changeContractAllowance(address adapter, address targetContract) returns()
func (_CreditManagerv2 *CreditManagerv2TransactorSession) ChangeContractAllowance(adapter common.Address, targetContract common.Address) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.ChangeContractAllowance(&_CreditManagerv2.TransactOpts, adapter, targetContract)
}

// CheckAndEnableToken is a paid mutator transaction binding the contract method 0x51e3f160.
//
// Solidity: function checkAndEnableToken(address creditAccount, address tokenOut) returns()
func (_CreditManagerv2 *CreditManagerv2Transactor) CheckAndEnableToken(opts *bind.TransactOpts, creditAccount common.Address, tokenOut common.Address) (*types.Transaction, error) {
	return _CreditManagerv2.contract.Transact(opts, "checkAndEnableToken", creditAccount, tokenOut)
}

// CheckAndEnableToken is a paid mutator transaction binding the contract method 0x51e3f160.
//
// Solidity: function checkAndEnableToken(address creditAccount, address tokenOut) returns()
func (_CreditManagerv2 *CreditManagerv2Session) CheckAndEnableToken(creditAccount common.Address, tokenOut common.Address) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.CheckAndEnableToken(&_CreditManagerv2.TransactOpts, creditAccount, tokenOut)
}

// CheckAndEnableToken is a paid mutator transaction binding the contract method 0x51e3f160.
//
// Solidity: function checkAndEnableToken(address creditAccount, address tokenOut) returns()
func (_CreditManagerv2 *CreditManagerv2TransactorSession) CheckAndEnableToken(creditAccount common.Address, tokenOut common.Address) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.CheckAndEnableToken(&_CreditManagerv2.TransactOpts, creditAccount, tokenOut)
}

// CloseCreditAccount is a paid mutator transaction binding the contract method 0x4f742e25.
//
// Solidity: function closeCreditAccount(address borrower, bool isLiquidated, uint256 totalValue, address payer, address to, bool sendAllAssets, bool convertWETH) returns(uint256 remainingFunds)
func (_CreditManagerv2 *CreditManagerv2Transactor) CloseCreditAccount(opts *bind.TransactOpts, borrower common.Address, isLiquidated bool, totalValue *big.Int, payer common.Address, to common.Address, sendAllAssets bool, convertWETH bool) (*types.Transaction, error) {
	return _CreditManagerv2.contract.Transact(opts, "closeCreditAccount", borrower, isLiquidated, totalValue, payer, to, sendAllAssets, convertWETH)
}

// CloseCreditAccount is a paid mutator transaction binding the contract method 0x4f742e25.
//
// Solidity: function closeCreditAccount(address borrower, bool isLiquidated, uint256 totalValue, address payer, address to, bool sendAllAssets, bool convertWETH) returns(uint256 remainingFunds)
func (_CreditManagerv2 *CreditManagerv2Session) CloseCreditAccount(borrower common.Address, isLiquidated bool, totalValue *big.Int, payer common.Address, to common.Address, sendAllAssets bool, convertWETH bool) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.CloseCreditAccount(&_CreditManagerv2.TransactOpts, borrower, isLiquidated, totalValue, payer, to, sendAllAssets, convertWETH)
}

// CloseCreditAccount is a paid mutator transaction binding the contract method 0x4f742e25.
//
// Solidity: function closeCreditAccount(address borrower, bool isLiquidated, uint256 totalValue, address payer, address to, bool sendAllAssets, bool convertWETH) returns(uint256 remainingFunds)
func (_CreditManagerv2 *CreditManagerv2TransactorSession) CloseCreditAccount(borrower common.Address, isLiquidated bool, totalValue *big.Int, payer common.Address, to common.Address, sendAllAssets bool, convertWETH bool) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.CloseCreditAccount(&_CreditManagerv2.TransactOpts, borrower, isLiquidated, totalValue, payer, to, sendAllAssets, convertWETH)
}

// ExecuteOrder is a paid mutator transaction binding the contract method 0x6ce4074a.
//
// Solidity: function executeOrder(address borrower, address targetContract, bytes data) returns(bytes)
func (_CreditManagerv2 *CreditManagerv2Transactor) ExecuteOrder(opts *bind.TransactOpts, borrower common.Address, targetContract common.Address, data []byte) (*types.Transaction, error) {
	return _CreditManagerv2.contract.Transact(opts, "executeOrder", borrower, targetContract, data)
}

// ExecuteOrder is a paid mutator transaction binding the contract method 0x6ce4074a.
//
// Solidity: function executeOrder(address borrower, address targetContract, bytes data) returns(bytes)
func (_CreditManagerv2 *CreditManagerv2Session) ExecuteOrder(borrower common.Address, targetContract common.Address, data []byte) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.ExecuteOrder(&_CreditManagerv2.TransactOpts, borrower, targetContract, data)
}

// ExecuteOrder is a paid mutator transaction binding the contract method 0x6ce4074a.
//
// Solidity: function executeOrder(address borrower, address targetContract, bytes data) returns(bytes)
func (_CreditManagerv2 *CreditManagerv2TransactorSession) ExecuteOrder(borrower common.Address, targetContract common.Address, data []byte) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.ExecuteOrder(&_CreditManagerv2.TransactOpts, borrower, targetContract, data)
}

// FastCollateralCheck is a paid mutator transaction binding the contract method 0x654a9eda.
//
// Solidity: function fastCollateralCheck(address creditAccount, address tokenIn, address tokenOut, uint256 balanceInBefore, uint256 balanceOutBefore) returns()
func (_CreditManagerv2 *CreditManagerv2Transactor) FastCollateralCheck(opts *bind.TransactOpts, creditAccount common.Address, tokenIn common.Address, tokenOut common.Address, balanceInBefore *big.Int, balanceOutBefore *big.Int) (*types.Transaction, error) {
	return _CreditManagerv2.contract.Transact(opts, "fastCollateralCheck", creditAccount, tokenIn, tokenOut, balanceInBefore, balanceOutBefore)
}

// FastCollateralCheck is a paid mutator transaction binding the contract method 0x654a9eda.
//
// Solidity: function fastCollateralCheck(address creditAccount, address tokenIn, address tokenOut, uint256 balanceInBefore, uint256 balanceOutBefore) returns()
func (_CreditManagerv2 *CreditManagerv2Session) FastCollateralCheck(creditAccount common.Address, tokenIn common.Address, tokenOut common.Address, balanceInBefore *big.Int, balanceOutBefore *big.Int) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.FastCollateralCheck(&_CreditManagerv2.TransactOpts, creditAccount, tokenIn, tokenOut, balanceInBefore, balanceOutBefore)
}

// FastCollateralCheck is a paid mutator transaction binding the contract method 0x654a9eda.
//
// Solidity: function fastCollateralCheck(address creditAccount, address tokenIn, address tokenOut, uint256 balanceInBefore, uint256 balanceOutBefore) returns()
func (_CreditManagerv2 *CreditManagerv2TransactorSession) FastCollateralCheck(creditAccount common.Address, tokenIn common.Address, tokenOut common.Address, balanceInBefore *big.Int, balanceOutBefore *big.Int) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.FastCollateralCheck(&_CreditManagerv2.TransactOpts, creditAccount, tokenIn, tokenOut, balanceInBefore, balanceOutBefore)
}

// FullCollateralCheck is a paid mutator transaction binding the contract method 0x95373018.
//
// Solidity: function fullCollateralCheck(address creditAccount) returns()
func (_CreditManagerv2 *CreditManagerv2Transactor) FullCollateralCheck(opts *bind.TransactOpts, creditAccount common.Address) (*types.Transaction, error) {
	return _CreditManagerv2.contract.Transact(opts, "fullCollateralCheck", creditAccount)
}

// FullCollateralCheck is a paid mutator transaction binding the contract method 0x95373018.
//
// Solidity: function fullCollateralCheck(address creditAccount) returns()
func (_CreditManagerv2 *CreditManagerv2Session) FullCollateralCheck(creditAccount common.Address) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.FullCollateralCheck(&_CreditManagerv2.TransactOpts, creditAccount)
}

// FullCollateralCheck is a paid mutator transaction binding the contract method 0x95373018.
//
// Solidity: function fullCollateralCheck(address creditAccount) returns()
func (_CreditManagerv2 *CreditManagerv2TransactorSession) FullCollateralCheck(creditAccount common.Address) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.FullCollateralCheck(&_CreditManagerv2.TransactOpts, creditAccount)
}

// ManageDebt is a paid mutator transaction binding the contract method 0x94cf073a.
//
// Solidity: function manageDebt(address borrower, uint256 amount, bool increase) returns(uint256 newBorrowedAmount)
func (_CreditManagerv2 *CreditManagerv2Transactor) ManageDebt(opts *bind.TransactOpts, borrower common.Address, amount *big.Int, increase bool) (*types.Transaction, error) {
	return _CreditManagerv2.contract.Transact(opts, "manageDebt", borrower, amount, increase)
}

// ManageDebt is a paid mutator transaction binding the contract method 0x94cf073a.
//
// Solidity: function manageDebt(address borrower, uint256 amount, bool increase) returns(uint256 newBorrowedAmount)
func (_CreditManagerv2 *CreditManagerv2Session) ManageDebt(borrower common.Address, amount *big.Int, increase bool) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.ManageDebt(&_CreditManagerv2.TransactOpts, borrower, amount, increase)
}

// ManageDebt is a paid mutator transaction binding the contract method 0x94cf073a.
//
// Solidity: function manageDebt(address borrower, uint256 amount, bool increase) returns(uint256 newBorrowedAmount)
func (_CreditManagerv2 *CreditManagerv2TransactorSession) ManageDebt(borrower common.Address, amount *big.Int, increase bool) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.ManageDebt(&_CreditManagerv2.TransactOpts, borrower, amount, increase)
}

// OpenCreditAccount is a paid mutator transaction binding the contract method 0x8fe3f93f.
//
// Solidity: function openCreditAccount(uint256 borrowedAmount, address onBehalfOf) returns(address)
func (_CreditManagerv2 *CreditManagerv2Transactor) OpenCreditAccount(opts *bind.TransactOpts, borrowedAmount *big.Int, onBehalfOf common.Address) (*types.Transaction, error) {
	return _CreditManagerv2.contract.Transact(opts, "openCreditAccount", borrowedAmount, onBehalfOf)
}

// OpenCreditAccount is a paid mutator transaction binding the contract method 0x8fe3f93f.
//
// Solidity: function openCreditAccount(uint256 borrowedAmount, address onBehalfOf) returns(address)
func (_CreditManagerv2 *CreditManagerv2Session) OpenCreditAccount(borrowedAmount *big.Int, onBehalfOf common.Address) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.OpenCreditAccount(&_CreditManagerv2.TransactOpts, borrowedAmount, onBehalfOf)
}

// OpenCreditAccount is a paid mutator transaction binding the contract method 0x8fe3f93f.
//
// Solidity: function openCreditAccount(uint256 borrowedAmount, address onBehalfOf) returns(address)
func (_CreditManagerv2 *CreditManagerv2TransactorSession) OpenCreditAccount(borrowedAmount *big.Int, onBehalfOf common.Address) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.OpenCreditAccount(&_CreditManagerv2.TransactOpts, borrowedAmount, onBehalfOf)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CreditManagerv2 *CreditManagerv2Transactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditManagerv2.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CreditManagerv2 *CreditManagerv2Session) Pause() (*types.Transaction, error) {
	return _CreditManagerv2.Contract.Pause(&_CreditManagerv2.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CreditManagerv2 *CreditManagerv2TransactorSession) Pause() (*types.Transaction, error) {
	return _CreditManagerv2.Contract.Pause(&_CreditManagerv2.TransactOpts)
}

// SetConfigurator is a paid mutator transaction binding the contract method 0x9f5f86ae.
//
// Solidity: function setConfigurator(address _creditConfigurator) returns()
func (_CreditManagerv2 *CreditManagerv2Transactor) SetConfigurator(opts *bind.TransactOpts, _creditConfigurator common.Address) (*types.Transaction, error) {
	return _CreditManagerv2.contract.Transact(opts, "setConfigurator", _creditConfigurator)
}

// SetConfigurator is a paid mutator transaction binding the contract method 0x9f5f86ae.
//
// Solidity: function setConfigurator(address _creditConfigurator) returns()
func (_CreditManagerv2 *CreditManagerv2Session) SetConfigurator(_creditConfigurator common.Address) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.SetConfigurator(&_CreditManagerv2.TransactOpts, _creditConfigurator)
}

// SetConfigurator is a paid mutator transaction binding the contract method 0x9f5f86ae.
//
// Solidity: function setConfigurator(address _creditConfigurator) returns()
func (_CreditManagerv2 *CreditManagerv2TransactorSession) SetConfigurator(_creditConfigurator common.Address) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.SetConfigurator(&_CreditManagerv2.TransactOpts, _creditConfigurator)
}

// SetForbidMask is a paid mutator transaction binding the contract method 0xa366f496.
//
// Solidity: function setForbidMask(uint256 _forbidMask) returns()
func (_CreditManagerv2 *CreditManagerv2Transactor) SetForbidMask(opts *bind.TransactOpts, _forbidMask *big.Int) (*types.Transaction, error) {
	return _CreditManagerv2.contract.Transact(opts, "setForbidMask", _forbidMask)
}

// SetForbidMask is a paid mutator transaction binding the contract method 0xa366f496.
//
// Solidity: function setForbidMask(uint256 _forbidMask) returns()
func (_CreditManagerv2 *CreditManagerv2Session) SetForbidMask(_forbidMask *big.Int) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.SetForbidMask(&_CreditManagerv2.TransactOpts, _forbidMask)
}

// SetForbidMask is a paid mutator transaction binding the contract method 0xa366f496.
//
// Solidity: function setForbidMask(uint256 _forbidMask) returns()
func (_CreditManagerv2 *CreditManagerv2TransactorSession) SetForbidMask(_forbidMask *big.Int) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.SetForbidMask(&_CreditManagerv2.TransactOpts, _forbidMask)
}

// SetLiquidationThreshold is a paid mutator transaction binding the contract method 0x0e30428d.
//
// Solidity: function setLiquidationThreshold(address token, uint256 liquidationThreshold) returns()
func (_CreditManagerv2 *CreditManagerv2Transactor) SetLiquidationThreshold(opts *bind.TransactOpts, token common.Address, liquidationThreshold *big.Int) (*types.Transaction, error) {
	return _CreditManagerv2.contract.Transact(opts, "setLiquidationThreshold", token, liquidationThreshold)
}

// SetLiquidationThreshold is a paid mutator transaction binding the contract method 0x0e30428d.
//
// Solidity: function setLiquidationThreshold(address token, uint256 liquidationThreshold) returns()
func (_CreditManagerv2 *CreditManagerv2Session) SetLiquidationThreshold(token common.Address, liquidationThreshold *big.Int) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.SetLiquidationThreshold(&_CreditManagerv2.TransactOpts, token, liquidationThreshold)
}

// SetLiquidationThreshold is a paid mutator transaction binding the contract method 0x0e30428d.
//
// Solidity: function setLiquidationThreshold(address token, uint256 liquidationThreshold) returns()
func (_CreditManagerv2 *CreditManagerv2TransactorSession) SetLiquidationThreshold(token common.Address, liquidationThreshold *big.Int) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.SetLiquidationThreshold(&_CreditManagerv2.TransactOpts, token, liquidationThreshold)
}

// SetParams is a paid mutator transaction binding the contract method 0xce1c4556.
//
// Solidity: function setParams(uint256 _minBorrowedAmount, uint256 _maxBorrowedAmount, uint256 _feeInterest, uint256 _feeLiquidation, uint256 _liquidationDiscount, uint256 _chiThreshold, uint256 _hfCheckInterval) returns()
func (_CreditManagerv2 *CreditManagerv2Transactor) SetParams(opts *bind.TransactOpts, _minBorrowedAmount *big.Int, _maxBorrowedAmount *big.Int, _feeInterest *big.Int, _feeLiquidation *big.Int, _liquidationDiscount *big.Int, _chiThreshold *big.Int, _hfCheckInterval *big.Int) (*types.Transaction, error) {
	return _CreditManagerv2.contract.Transact(opts, "setParams", _minBorrowedAmount, _maxBorrowedAmount, _feeInterest, _feeLiquidation, _liquidationDiscount, _chiThreshold, _hfCheckInterval)
}

// SetParams is a paid mutator transaction binding the contract method 0xce1c4556.
//
// Solidity: function setParams(uint256 _minBorrowedAmount, uint256 _maxBorrowedAmount, uint256 _feeInterest, uint256 _feeLiquidation, uint256 _liquidationDiscount, uint256 _chiThreshold, uint256 _hfCheckInterval) returns()
func (_CreditManagerv2 *CreditManagerv2Session) SetParams(_minBorrowedAmount *big.Int, _maxBorrowedAmount *big.Int, _feeInterest *big.Int, _feeLiquidation *big.Int, _liquidationDiscount *big.Int, _chiThreshold *big.Int, _hfCheckInterval *big.Int) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.SetParams(&_CreditManagerv2.TransactOpts, _minBorrowedAmount, _maxBorrowedAmount, _feeInterest, _feeLiquidation, _liquidationDiscount, _chiThreshold, _hfCheckInterval)
}

// SetParams is a paid mutator transaction binding the contract method 0xce1c4556.
//
// Solidity: function setParams(uint256 _minBorrowedAmount, uint256 _maxBorrowedAmount, uint256 _feeInterest, uint256 _feeLiquidation, uint256 _liquidationDiscount, uint256 _chiThreshold, uint256 _hfCheckInterval) returns()
func (_CreditManagerv2 *CreditManagerv2TransactorSession) SetParams(_minBorrowedAmount *big.Int, _maxBorrowedAmount *big.Int, _feeInterest *big.Int, _feeLiquidation *big.Int, _liquidationDiscount *big.Int, _chiThreshold *big.Int, _hfCheckInterval *big.Int) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.SetParams(&_CreditManagerv2.TransactOpts, _minBorrowedAmount, _maxBorrowedAmount, _feeInterest, _feeLiquidation, _liquidationDiscount, _chiThreshold, _hfCheckInterval)
}

// TransferAccountOwnership is a paid mutator transaction binding the contract method 0xe1998cf9.
//
// Solidity: function transferAccountOwnership(address from, address to) returns()
func (_CreditManagerv2 *CreditManagerv2Transactor) TransferAccountOwnership(opts *bind.TransactOpts, from common.Address, to common.Address) (*types.Transaction, error) {
	return _CreditManagerv2.contract.Transact(opts, "transferAccountOwnership", from, to)
}

// TransferAccountOwnership is a paid mutator transaction binding the contract method 0xe1998cf9.
//
// Solidity: function transferAccountOwnership(address from, address to) returns()
func (_CreditManagerv2 *CreditManagerv2Session) TransferAccountOwnership(from common.Address, to common.Address) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.TransferAccountOwnership(&_CreditManagerv2.TransactOpts, from, to)
}

// TransferAccountOwnership is a paid mutator transaction binding the contract method 0xe1998cf9.
//
// Solidity: function transferAccountOwnership(address from, address to) returns()
func (_CreditManagerv2 *CreditManagerv2TransactorSession) TransferAccountOwnership(from common.Address, to common.Address) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.TransferAccountOwnership(&_CreditManagerv2.TransactOpts, from, to)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CreditManagerv2 *CreditManagerv2Transactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditManagerv2.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CreditManagerv2 *CreditManagerv2Session) Unpause() (*types.Transaction, error) {
	return _CreditManagerv2.Contract.Unpause(&_CreditManagerv2.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CreditManagerv2 *CreditManagerv2TransactorSession) Unpause() (*types.Transaction, error) {
	return _CreditManagerv2.Contract.Unpause(&_CreditManagerv2.TransactOpts)
}

// UpgradeContracts is a paid mutator transaction binding the contract method 0x11ca4fc2.
//
// Solidity: function upgradeContracts(address _creditFacade, address _priceOracle) returns()
func (_CreditManagerv2 *CreditManagerv2Transactor) UpgradeContracts(opts *bind.TransactOpts, _creditFacade common.Address, _priceOracle common.Address) (*types.Transaction, error) {
	return _CreditManagerv2.contract.Transact(opts, "upgradeContracts", _creditFacade, _priceOracle)
}

// UpgradeContracts is a paid mutator transaction binding the contract method 0x11ca4fc2.
//
// Solidity: function upgradeContracts(address _creditFacade, address _priceOracle) returns()
func (_CreditManagerv2 *CreditManagerv2Session) UpgradeContracts(_creditFacade common.Address, _priceOracle common.Address) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.UpgradeContracts(&_CreditManagerv2.TransactOpts, _creditFacade, _priceOracle)
}

// UpgradeContracts is a paid mutator transaction binding the contract method 0x11ca4fc2.
//
// Solidity: function upgradeContracts(address _creditFacade, address _priceOracle) returns()
func (_CreditManagerv2 *CreditManagerv2TransactorSession) UpgradeContracts(_creditFacade common.Address, _priceOracle common.Address) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.UpgradeContracts(&_CreditManagerv2.TransactOpts, _creditFacade, _priceOracle)
}

// CreditManagerv2ExecuteOrderIterator is returned from FilterExecuteOrder and is used to iterate over the raw logs and unpacked data for ExecuteOrder events raised by the CreditManagerv2 contract.
type CreditManagerv2ExecuteOrderIterator struct {
	Event *CreditManagerv2ExecuteOrder // Event containing the contract specifics and raw log

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
func (it *CreditManagerv2ExecuteOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditManagerv2ExecuteOrder)
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
		it.Event = new(CreditManagerv2ExecuteOrder)
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
func (it *CreditManagerv2ExecuteOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditManagerv2ExecuteOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditManagerv2ExecuteOrder represents a ExecuteOrder event raised by the CreditManagerv2 contract.
type CreditManagerv2ExecuteOrder struct {
	Borrower common.Address
	Target   common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterExecuteOrder is a free log retrieval operation binding the contract event 0xaed1eb34af6acd8c1e3911fb2ebb875a66324b03957886bd002227b17f52ab03.
//
// Solidity: event ExecuteOrder(address indexed borrower, address indexed target)
func (_CreditManagerv2 *CreditManagerv2Filterer) FilterExecuteOrder(opts *bind.FilterOpts, borrower []common.Address, target []common.Address) (*CreditManagerv2ExecuteOrderIterator, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _CreditManagerv2.contract.FilterLogs(opts, "ExecuteOrder", borrowerRule, targetRule)
	if err != nil {
		return nil, err
	}
	return &CreditManagerv2ExecuteOrderIterator{contract: _CreditManagerv2.contract, event: "ExecuteOrder", logs: logs, sub: sub}, nil
}

// WatchExecuteOrder is a free log subscription operation binding the contract event 0xaed1eb34af6acd8c1e3911fb2ebb875a66324b03957886bd002227b17f52ab03.
//
// Solidity: event ExecuteOrder(address indexed borrower, address indexed target)
func (_CreditManagerv2 *CreditManagerv2Filterer) WatchExecuteOrder(opts *bind.WatchOpts, sink chan<- *CreditManagerv2ExecuteOrder, borrower []common.Address, target []common.Address) (event.Subscription, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _CreditManagerv2.contract.WatchLogs(opts, "ExecuteOrder", borrowerRule, targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditManagerv2ExecuteOrder)
				if err := _CreditManagerv2.contract.UnpackLog(event, "ExecuteOrder", log); err != nil {
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
func (_CreditManagerv2 *CreditManagerv2Filterer) ParseExecuteOrder(log types.Log) (*CreditManagerv2ExecuteOrder, error) {
	event := new(CreditManagerv2ExecuteOrder)
	if err := _CreditManagerv2.contract.UnpackLog(event, "ExecuteOrder", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CreditManagerv2NewConfiguratorIterator is returned from FilterNewConfigurator and is used to iterate over the raw logs and unpacked data for NewConfigurator events raised by the CreditManagerv2 contract.
type CreditManagerv2NewConfiguratorIterator struct {
	Event *CreditManagerv2NewConfigurator // Event containing the contract specifics and raw log

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
func (it *CreditManagerv2NewConfiguratorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditManagerv2NewConfigurator)
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
		it.Event = new(CreditManagerv2NewConfigurator)
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
func (it *CreditManagerv2NewConfiguratorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditManagerv2NewConfiguratorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditManagerv2NewConfigurator represents a NewConfigurator event raised by the CreditManagerv2 contract.
type CreditManagerv2NewConfigurator struct {
	NewConfigurator common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewConfigurator is a free log retrieval operation binding the contract event 0xf62005acebe9b616aefb5f248b48f5e89f28437b27d1eebc0b2d911209f297af.
//
// Solidity: event NewConfigurator(address indexed newConfigurator)
func (_CreditManagerv2 *CreditManagerv2Filterer) FilterNewConfigurator(opts *bind.FilterOpts, newConfigurator []common.Address) (*CreditManagerv2NewConfiguratorIterator, error) {

	var newConfiguratorRule []interface{}
	for _, newConfiguratorItem := range newConfigurator {
		newConfiguratorRule = append(newConfiguratorRule, newConfiguratorItem)
	}

	logs, sub, err := _CreditManagerv2.contract.FilterLogs(opts, "NewConfigurator", newConfiguratorRule)
	if err != nil {
		return nil, err
	}
	return &CreditManagerv2NewConfiguratorIterator{contract: _CreditManagerv2.contract, event: "NewConfigurator", logs: logs, sub: sub}, nil
}

// WatchNewConfigurator is a free log subscription operation binding the contract event 0xf62005acebe9b616aefb5f248b48f5e89f28437b27d1eebc0b2d911209f297af.
//
// Solidity: event NewConfigurator(address indexed newConfigurator)
func (_CreditManagerv2 *CreditManagerv2Filterer) WatchNewConfigurator(opts *bind.WatchOpts, sink chan<- *CreditManagerv2NewConfigurator, newConfigurator []common.Address) (event.Subscription, error) {

	var newConfiguratorRule []interface{}
	for _, newConfiguratorItem := range newConfigurator {
		newConfiguratorRule = append(newConfiguratorRule, newConfiguratorItem)
	}

	logs, sub, err := _CreditManagerv2.contract.WatchLogs(opts, "NewConfigurator", newConfiguratorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditManagerv2NewConfigurator)
				if err := _CreditManagerv2.contract.UnpackLog(event, "NewConfigurator", log); err != nil {
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

// ParseNewConfigurator is a log parse operation binding the contract event 0xf62005acebe9b616aefb5f248b48f5e89f28437b27d1eebc0b2d911209f297af.
//
// Solidity: event NewConfigurator(address indexed newConfigurator)
func (_CreditManagerv2 *CreditManagerv2Filterer) ParseNewConfigurator(log types.Log) (*CreditManagerv2NewConfigurator, error) {
	event := new(CreditManagerv2NewConfigurator)
	if err := _CreditManagerv2.contract.UnpackLog(event, "NewConfigurator", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CreditManagerv2PausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the CreditManagerv2 contract.
type CreditManagerv2PausedIterator struct {
	Event *CreditManagerv2Paused // Event containing the contract specifics and raw log

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
func (it *CreditManagerv2PausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditManagerv2Paused)
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
		it.Event = new(CreditManagerv2Paused)
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
func (it *CreditManagerv2PausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditManagerv2PausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditManagerv2Paused represents a Paused event raised by the CreditManagerv2 contract.
type CreditManagerv2Paused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_CreditManagerv2 *CreditManagerv2Filterer) FilterPaused(opts *bind.FilterOpts) (*CreditManagerv2PausedIterator, error) {

	logs, sub, err := _CreditManagerv2.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &CreditManagerv2PausedIterator{contract: _CreditManagerv2.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_CreditManagerv2 *CreditManagerv2Filterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *CreditManagerv2Paused) (event.Subscription, error) {

	logs, sub, err := _CreditManagerv2.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditManagerv2Paused)
				if err := _CreditManagerv2.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_CreditManagerv2 *CreditManagerv2Filterer) ParsePaused(log types.Log) (*CreditManagerv2Paused, error) {
	event := new(CreditManagerv2Paused)
	if err := _CreditManagerv2.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CreditManagerv2UnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the CreditManagerv2 contract.
type CreditManagerv2UnpausedIterator struct {
	Event *CreditManagerv2Unpaused // Event containing the contract specifics and raw log

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
func (it *CreditManagerv2UnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditManagerv2Unpaused)
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
		it.Event = new(CreditManagerv2Unpaused)
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
func (it *CreditManagerv2UnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditManagerv2UnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditManagerv2Unpaused represents a Unpaused event raised by the CreditManagerv2 contract.
type CreditManagerv2Unpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_CreditManagerv2 *CreditManagerv2Filterer) FilterUnpaused(opts *bind.FilterOpts) (*CreditManagerv2UnpausedIterator, error) {

	logs, sub, err := _CreditManagerv2.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &CreditManagerv2UnpausedIterator{contract: _CreditManagerv2.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_CreditManagerv2 *CreditManagerv2Filterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *CreditManagerv2Unpaused) (event.Subscription, error) {

	logs, sub, err := _CreditManagerv2.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditManagerv2Unpaused)
				if err := _CreditManagerv2.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_CreditManagerv2 *CreditManagerv2Filterer) ParseUnpaused(log types.Log) (*CreditManagerv2Unpaused, error) {
	event := new(CreditManagerv2Unpaused)
	if err := _CreditManagerv2.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	return event, nil
}
