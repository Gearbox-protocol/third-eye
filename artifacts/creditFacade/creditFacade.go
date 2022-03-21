// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package creditFacade

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

// MultiCall is an auto generated low-level Go binding around an user-defined struct.
type MultiCall struct {
	Target   common.Address
	CallData []byte
}

// CreditFacadeABI is the input ABI used to generate the binding from.
const CreditFacadeABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_creditManager\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"AddCollateral\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"CloseCreditAccount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DecreaseBorrowedAmount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"IncreaseBorrowedAmount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"remainingFunds\",\"type\":\"uint256\"}],\"name\":\"LiquidateCreditAccount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"MultiCallFinished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"}],\"name\":\"MultiCallStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"referralCode\",\"type\":\"uint256\"}],\"name\":\"OpenCreditAccount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"TransferAccount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"state\",\"type\":\"bool\"}],\"name\":\"TransferAccountAllowed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"addCollateral\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"targetContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"state\",\"type\":\"bool\"}],\"name\":\"approveAccountTransfers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"}],\"name\":\"calcCreditAccountHealthFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"hf\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"}],\"name\":\"calcTotalValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"twv\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"skipTokenMask\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"convertWETH\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMultiCall[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"closeCreditAccount\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"contractToAdapter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creditManager\",\"outputs\":[{\"internalType\":\"contractICreditManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"decreaseDebt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"}],\"name\":\"hasOpenedCreditAccount\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"increaseDebt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"isTokenAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"skipTokenMask\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"convertWETH\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMultiCall[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"liquidateCreditAccount\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMultiCall[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"multicall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"leverageFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"referralCode\",\"type\":\"uint256\"}],\"name\":\"openCreditAccount\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"borrowedAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMultiCall[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"referralCode\",\"type\":\"uint256\"}],\"name\":\"openCreditAccountMulticall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_adapter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"}],\"name\":\"setContractToAdapter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferAccountOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"transfersAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlyingToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wethAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// CreditFacade is an auto generated Go binding around an Ethereum contract.
type CreditFacade struct {
	CreditFacadeCaller     // Read-only binding to the contract
	CreditFacadeTransactor // Write-only binding to the contract
	CreditFacadeFilterer   // Log filterer for contract events
}

// CreditFacadeCaller is an auto generated read-only Go binding around an Ethereum contract.
type CreditFacadeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditFacadeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CreditFacadeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditFacadeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CreditFacadeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditFacadeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CreditFacadeSession struct {
	Contract     *CreditFacade     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CreditFacadeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CreditFacadeCallerSession struct {
	Contract *CreditFacadeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// CreditFacadeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CreditFacadeTransactorSession struct {
	Contract     *CreditFacadeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// CreditFacadeRaw is an auto generated low-level Go binding around an Ethereum contract.
type CreditFacadeRaw struct {
	Contract *CreditFacade // Generic contract binding to access the raw methods on
}

// CreditFacadeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CreditFacadeCallerRaw struct {
	Contract *CreditFacadeCaller // Generic read-only contract binding to access the raw methods on
}

// CreditFacadeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CreditFacadeTransactorRaw struct {
	Contract *CreditFacadeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCreditFacade creates a new instance of CreditFacade, bound to a specific deployed contract.
func NewCreditFacade(address common.Address, backend bind.ContractBackend) (*CreditFacade, error) {
	contract, err := bindCreditFacade(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CreditFacade{CreditFacadeCaller: CreditFacadeCaller{contract: contract}, CreditFacadeTransactor: CreditFacadeTransactor{contract: contract}, CreditFacadeFilterer: CreditFacadeFilterer{contract: contract}}, nil
}

// NewCreditFacadeCaller creates a new read-only instance of CreditFacade, bound to a specific deployed contract.
func NewCreditFacadeCaller(address common.Address, caller bind.ContractCaller) (*CreditFacadeCaller, error) {
	contract, err := bindCreditFacade(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CreditFacadeCaller{contract: contract}, nil
}

// NewCreditFacadeTransactor creates a new write-only instance of CreditFacade, bound to a specific deployed contract.
func NewCreditFacadeTransactor(address common.Address, transactor bind.ContractTransactor) (*CreditFacadeTransactor, error) {
	contract, err := bindCreditFacade(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CreditFacadeTransactor{contract: contract}, nil
}

// NewCreditFacadeFilterer creates a new log filterer instance of CreditFacade, bound to a specific deployed contract.
func NewCreditFacadeFilterer(address common.Address, filterer bind.ContractFilterer) (*CreditFacadeFilterer, error) {
	contract, err := bindCreditFacade(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CreditFacadeFilterer{contract: contract}, nil
}

// bindCreditFacade binds a generic wrapper to an already deployed contract.
func bindCreditFacade(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CreditFacadeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CreditFacade *CreditFacadeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CreditFacade.Contract.CreditFacadeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CreditFacade *CreditFacadeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditFacade.Contract.CreditFacadeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CreditFacade *CreditFacadeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CreditFacade.Contract.CreditFacadeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CreditFacade *CreditFacadeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CreditFacade.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CreditFacade *CreditFacadeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditFacade.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CreditFacade *CreditFacadeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CreditFacade.Contract.contract.Transact(opts, method, params...)
}

// CalcCreditAccountHealthFactor is a free data retrieval call binding the contract method 0xdfd59465.
//
// Solidity: function calcCreditAccountHealthFactor(address creditAccount) view returns(uint256 hf)
func (_CreditFacade *CreditFacadeCaller) CalcCreditAccountHealthFactor(opts *bind.CallOpts, creditAccount common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CreditFacade.contract.Call(opts, &out, "calcCreditAccountHealthFactor", creditAccount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcCreditAccountHealthFactor is a free data retrieval call binding the contract method 0xdfd59465.
//
// Solidity: function calcCreditAccountHealthFactor(address creditAccount) view returns(uint256 hf)
func (_CreditFacade *CreditFacadeSession) CalcCreditAccountHealthFactor(creditAccount common.Address) (*big.Int, error) {
	return _CreditFacade.Contract.CalcCreditAccountHealthFactor(&_CreditFacade.CallOpts, creditAccount)
}

// CalcCreditAccountHealthFactor is a free data retrieval call binding the contract method 0xdfd59465.
//
// Solidity: function calcCreditAccountHealthFactor(address creditAccount) view returns(uint256 hf)
func (_CreditFacade *CreditFacadeCallerSession) CalcCreditAccountHealthFactor(creditAccount common.Address) (*big.Int, error) {
	return _CreditFacade.Contract.CalcCreditAccountHealthFactor(&_CreditFacade.CallOpts, creditAccount)
}

// CalcTotalValue is a free data retrieval call binding the contract method 0xc7de38a6.
//
// Solidity: function calcTotalValue(address creditAccount) view returns(uint256 total, uint256 twv)
func (_CreditFacade *CreditFacadeCaller) CalcTotalValue(opts *bind.CallOpts, creditAccount common.Address) (struct {
	Total *big.Int
	Twv   *big.Int
}, error) {
	var out []interface{}
	err := _CreditFacade.contract.Call(opts, &out, "calcTotalValue", creditAccount)

	outstruct := new(struct {
		Total *big.Int
		Twv   *big.Int
	})

	outstruct.Total = out[0].(*big.Int)
	outstruct.Twv = out[1].(*big.Int)

	return *outstruct, err

}

// CalcTotalValue is a free data retrieval call binding the contract method 0xc7de38a6.
//
// Solidity: function calcTotalValue(address creditAccount) view returns(uint256 total, uint256 twv)
func (_CreditFacade *CreditFacadeSession) CalcTotalValue(creditAccount common.Address) (struct {
	Total *big.Int
	Twv   *big.Int
}, error) {
	return _CreditFacade.Contract.CalcTotalValue(&_CreditFacade.CallOpts, creditAccount)
}

// CalcTotalValue is a free data retrieval call binding the contract method 0xc7de38a6.
//
// Solidity: function calcTotalValue(address creditAccount) view returns(uint256 total, uint256 twv)
func (_CreditFacade *CreditFacadeCallerSession) CalcTotalValue(creditAccount common.Address) (struct {
	Total *big.Int
	Twv   *big.Int
}, error) {
	return _CreditFacade.Contract.CalcTotalValue(&_CreditFacade.CallOpts, creditAccount)
}

// ContractToAdapter is a free data retrieval call binding the contract method 0xfdd57645.
//
// Solidity: function contractToAdapter(address ) view returns(address)
func (_CreditFacade *CreditFacadeCaller) ContractToAdapter(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _CreditFacade.contract.Call(opts, &out, "contractToAdapter", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ContractToAdapter is a free data retrieval call binding the contract method 0xfdd57645.
//
// Solidity: function contractToAdapter(address ) view returns(address)
func (_CreditFacade *CreditFacadeSession) ContractToAdapter(arg0 common.Address) (common.Address, error) {
	return _CreditFacade.Contract.ContractToAdapter(&_CreditFacade.CallOpts, arg0)
}

// ContractToAdapter is a free data retrieval call binding the contract method 0xfdd57645.
//
// Solidity: function contractToAdapter(address ) view returns(address)
func (_CreditFacade *CreditFacadeCallerSession) ContractToAdapter(arg0 common.Address) (common.Address, error) {
	return _CreditFacade.Contract.ContractToAdapter(&_CreditFacade.CallOpts, arg0)
}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_CreditFacade *CreditFacadeCaller) CreditManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditFacade.contract.Call(opts, &out, "creditManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_CreditFacade *CreditFacadeSession) CreditManager() (common.Address, error) {
	return _CreditFacade.Contract.CreditManager(&_CreditFacade.CallOpts)
}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_CreditFacade *CreditFacadeCallerSession) CreditManager() (common.Address, error) {
	return _CreditFacade.Contract.CreditManager(&_CreditFacade.CallOpts)
}

// HasOpenedCreditAccount is a free data retrieval call binding the contract method 0x256ac915.
//
// Solidity: function hasOpenedCreditAccount(address borrower) view returns(bool)
func (_CreditFacade *CreditFacadeCaller) HasOpenedCreditAccount(opts *bind.CallOpts, borrower common.Address) (bool, error) {
	var out []interface{}
	err := _CreditFacade.contract.Call(opts, &out, "hasOpenedCreditAccount", borrower)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasOpenedCreditAccount is a free data retrieval call binding the contract method 0x256ac915.
//
// Solidity: function hasOpenedCreditAccount(address borrower) view returns(bool)
func (_CreditFacade *CreditFacadeSession) HasOpenedCreditAccount(borrower common.Address) (bool, error) {
	return _CreditFacade.Contract.HasOpenedCreditAccount(&_CreditFacade.CallOpts, borrower)
}

// HasOpenedCreditAccount is a free data retrieval call binding the contract method 0x256ac915.
//
// Solidity: function hasOpenedCreditAccount(address borrower) view returns(bool)
func (_CreditFacade *CreditFacadeCallerSession) HasOpenedCreditAccount(borrower common.Address) (bool, error) {
	return _CreditFacade.Contract.HasOpenedCreditAccount(&_CreditFacade.CallOpts, borrower)
}

// IsTokenAllowed is a free data retrieval call binding the contract method 0xf9eaee0d.
//
// Solidity: function isTokenAllowed(address token) view returns(bool allowed)
func (_CreditFacade *CreditFacadeCaller) IsTokenAllowed(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _CreditFacade.contract.Call(opts, &out, "isTokenAllowed", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTokenAllowed is a free data retrieval call binding the contract method 0xf9eaee0d.
//
// Solidity: function isTokenAllowed(address token) view returns(bool allowed)
func (_CreditFacade *CreditFacadeSession) IsTokenAllowed(token common.Address) (bool, error) {
	return _CreditFacade.Contract.IsTokenAllowed(&_CreditFacade.CallOpts, token)
}

// IsTokenAllowed is a free data retrieval call binding the contract method 0xf9eaee0d.
//
// Solidity: function isTokenAllowed(address token) view returns(bool allowed)
func (_CreditFacade *CreditFacadeCallerSession) IsTokenAllowed(token common.Address) (bool, error) {
	return _CreditFacade.Contract.IsTokenAllowed(&_CreditFacade.CallOpts, token)
}

// TransfersAllowed is a free data retrieval call binding the contract method 0xd9ccbec1.
//
// Solidity: function transfersAllowed(address , address ) view returns(bool)
func (_CreditFacade *CreditFacadeCaller) TransfersAllowed(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _CreditFacade.contract.Call(opts, &out, "transfersAllowed", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TransfersAllowed is a free data retrieval call binding the contract method 0xd9ccbec1.
//
// Solidity: function transfersAllowed(address , address ) view returns(bool)
func (_CreditFacade *CreditFacadeSession) TransfersAllowed(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _CreditFacade.Contract.TransfersAllowed(&_CreditFacade.CallOpts, arg0, arg1)
}

// TransfersAllowed is a free data retrieval call binding the contract method 0xd9ccbec1.
//
// Solidity: function transfersAllowed(address , address ) view returns(bool)
func (_CreditFacade *CreditFacadeCallerSession) TransfersAllowed(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _CreditFacade.Contract.TransfersAllowed(&_CreditFacade.CallOpts, arg0, arg1)
}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_CreditFacade *CreditFacadeCaller) UnderlyingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditFacade.contract.Call(opts, &out, "underlyingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_CreditFacade *CreditFacadeSession) UnderlyingToken() (common.Address, error) {
	return _CreditFacade.Contract.UnderlyingToken(&_CreditFacade.CallOpts)
}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_CreditFacade *CreditFacadeCallerSession) UnderlyingToken() (common.Address, error) {
	return _CreditFacade.Contract.UnderlyingToken(&_CreditFacade.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_CreditFacade *CreditFacadeCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CreditFacade.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_CreditFacade *CreditFacadeSession) Version() (*big.Int, error) {
	return _CreditFacade.Contract.Version(&_CreditFacade.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_CreditFacade *CreditFacadeCallerSession) Version() (*big.Int, error) {
	return _CreditFacade.Contract.Version(&_CreditFacade.CallOpts)
}

// WethAddress is a free data retrieval call binding the contract method 0x4f0e0ef3.
//
// Solidity: function wethAddress() view returns(address)
func (_CreditFacade *CreditFacadeCaller) WethAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditFacade.contract.Call(opts, &out, "wethAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WethAddress is a free data retrieval call binding the contract method 0x4f0e0ef3.
//
// Solidity: function wethAddress() view returns(address)
func (_CreditFacade *CreditFacadeSession) WethAddress() (common.Address, error) {
	return _CreditFacade.Contract.WethAddress(&_CreditFacade.CallOpts)
}

// WethAddress is a free data retrieval call binding the contract method 0x4f0e0ef3.
//
// Solidity: function wethAddress() view returns(address)
func (_CreditFacade *CreditFacadeCallerSession) WethAddress() (common.Address, error) {
	return _CreditFacade.Contract.WethAddress(&_CreditFacade.CallOpts)
}

// AddCollateral is a paid mutator transaction binding the contract method 0x59781034.
//
// Solidity: function addCollateral(address onBehalfOf, address token, uint256 amount) payable returns()
func (_CreditFacade *CreditFacadeTransactor) AddCollateral(opts *bind.TransactOpts, onBehalfOf common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CreditFacade.contract.Transact(opts, "addCollateral", onBehalfOf, token, amount)
}

// AddCollateral is a paid mutator transaction binding the contract method 0x59781034.
//
// Solidity: function addCollateral(address onBehalfOf, address token, uint256 amount) payable returns()
func (_CreditFacade *CreditFacadeSession) AddCollateral(onBehalfOf common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CreditFacade.Contract.AddCollateral(&_CreditFacade.TransactOpts, onBehalfOf, token, amount)
}

// AddCollateral is a paid mutator transaction binding the contract method 0x59781034.
//
// Solidity: function addCollateral(address onBehalfOf, address token, uint256 amount) payable returns()
func (_CreditFacade *CreditFacadeTransactorSession) AddCollateral(onBehalfOf common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CreditFacade.Contract.AddCollateral(&_CreditFacade.TransactOpts, onBehalfOf, token, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x7e5465ba.
//
// Solidity: function approve(address targetContract, address token) returns()
func (_CreditFacade *CreditFacadeTransactor) Approve(opts *bind.TransactOpts, targetContract common.Address, token common.Address) (*types.Transaction, error) {
	return _CreditFacade.contract.Transact(opts, "approve", targetContract, token)
}

// Approve is a paid mutator transaction binding the contract method 0x7e5465ba.
//
// Solidity: function approve(address targetContract, address token) returns()
func (_CreditFacade *CreditFacadeSession) Approve(targetContract common.Address, token common.Address) (*types.Transaction, error) {
	return _CreditFacade.Contract.Approve(&_CreditFacade.TransactOpts, targetContract, token)
}

// Approve is a paid mutator transaction binding the contract method 0x7e5465ba.
//
// Solidity: function approve(address targetContract, address token) returns()
func (_CreditFacade *CreditFacadeTransactorSession) Approve(targetContract common.Address, token common.Address) (*types.Transaction, error) {
	return _CreditFacade.Contract.Approve(&_CreditFacade.TransactOpts, targetContract, token)
}

// ApproveAccountTransfers is a paid mutator transaction binding the contract method 0x5f27212a.
//
// Solidity: function approveAccountTransfers(address from, bool state) returns()
func (_CreditFacade *CreditFacadeTransactor) ApproveAccountTransfers(opts *bind.TransactOpts, from common.Address, state bool) (*types.Transaction, error) {
	return _CreditFacade.contract.Transact(opts, "approveAccountTransfers", from, state)
}

// ApproveAccountTransfers is a paid mutator transaction binding the contract method 0x5f27212a.
//
// Solidity: function approveAccountTransfers(address from, bool state) returns()
func (_CreditFacade *CreditFacadeSession) ApproveAccountTransfers(from common.Address, state bool) (*types.Transaction, error) {
	return _CreditFacade.Contract.ApproveAccountTransfers(&_CreditFacade.TransactOpts, from, state)
}

// ApproveAccountTransfers is a paid mutator transaction binding the contract method 0x5f27212a.
//
// Solidity: function approveAccountTransfers(address from, bool state) returns()
func (_CreditFacade *CreditFacadeTransactorSession) ApproveAccountTransfers(from common.Address, state bool) (*types.Transaction, error) {
	return _CreditFacade.Contract.ApproveAccountTransfers(&_CreditFacade.TransactOpts, from, state)
}

// CloseCreditAccount is a paid mutator transaction binding the contract method 0x5f73fbec.
//
// Solidity: function closeCreditAccount(address to, uint256 skipTokenMask, bool convertWETH, (address,bytes)[] calls) payable returns()
func (_CreditFacade *CreditFacadeTransactor) CloseCreditAccount(opts *bind.TransactOpts, to common.Address, skipTokenMask *big.Int, convertWETH bool, calls []MultiCall) (*types.Transaction, error) {
	return _CreditFacade.contract.Transact(opts, "closeCreditAccount", to, skipTokenMask, convertWETH, calls)
}

// CloseCreditAccount is a paid mutator transaction binding the contract method 0x5f73fbec.
//
// Solidity: function closeCreditAccount(address to, uint256 skipTokenMask, bool convertWETH, (address,bytes)[] calls) payable returns()
func (_CreditFacade *CreditFacadeSession) CloseCreditAccount(to common.Address, skipTokenMask *big.Int, convertWETH bool, calls []MultiCall) (*types.Transaction, error) {
	return _CreditFacade.Contract.CloseCreditAccount(&_CreditFacade.TransactOpts, to, skipTokenMask, convertWETH, calls)
}

// CloseCreditAccount is a paid mutator transaction binding the contract method 0x5f73fbec.
//
// Solidity: function closeCreditAccount(address to, uint256 skipTokenMask, bool convertWETH, (address,bytes)[] calls) payable returns()
func (_CreditFacade *CreditFacadeTransactorSession) CloseCreditAccount(to common.Address, skipTokenMask *big.Int, convertWETH bool, calls []MultiCall) (*types.Transaction, error) {
	return _CreditFacade.Contract.CloseCreditAccount(&_CreditFacade.TransactOpts, to, skipTokenMask, convertWETH, calls)
}

// DecreaseDebt is a paid mutator transaction binding the contract method 0x2a7ba1f7.
//
// Solidity: function decreaseDebt(uint256 amount) returns()
func (_CreditFacade *CreditFacadeTransactor) DecreaseDebt(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _CreditFacade.contract.Transact(opts, "decreaseDebt", amount)
}

// DecreaseDebt is a paid mutator transaction binding the contract method 0x2a7ba1f7.
//
// Solidity: function decreaseDebt(uint256 amount) returns()
func (_CreditFacade *CreditFacadeSession) DecreaseDebt(amount *big.Int) (*types.Transaction, error) {
	return _CreditFacade.Contract.DecreaseDebt(&_CreditFacade.TransactOpts, amount)
}

// DecreaseDebt is a paid mutator transaction binding the contract method 0x2a7ba1f7.
//
// Solidity: function decreaseDebt(uint256 amount) returns()
func (_CreditFacade *CreditFacadeTransactorSession) DecreaseDebt(amount *big.Int) (*types.Transaction, error) {
	return _CreditFacade.Contract.DecreaseDebt(&_CreditFacade.TransactOpts, amount)
}

// IncreaseDebt is a paid mutator transaction binding the contract method 0x2b7c7b11.
//
// Solidity: function increaseDebt(uint256 amount) returns()
func (_CreditFacade *CreditFacadeTransactor) IncreaseDebt(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _CreditFacade.contract.Transact(opts, "increaseDebt", amount)
}

// IncreaseDebt is a paid mutator transaction binding the contract method 0x2b7c7b11.
//
// Solidity: function increaseDebt(uint256 amount) returns()
func (_CreditFacade *CreditFacadeSession) IncreaseDebt(amount *big.Int) (*types.Transaction, error) {
	return _CreditFacade.Contract.IncreaseDebt(&_CreditFacade.TransactOpts, amount)
}

// IncreaseDebt is a paid mutator transaction binding the contract method 0x2b7c7b11.
//
// Solidity: function increaseDebt(uint256 amount) returns()
func (_CreditFacade *CreditFacadeTransactorSession) IncreaseDebt(amount *big.Int) (*types.Transaction, error) {
	return _CreditFacade.Contract.IncreaseDebt(&_CreditFacade.TransactOpts, amount)
}

// LiquidateCreditAccount is a paid mutator transaction binding the contract method 0x5d91a0e0.
//
// Solidity: function liquidateCreditAccount(address borrower, address to, uint256 skipTokenMask, bool convertWETH, (address,bytes)[] calls) payable returns()
func (_CreditFacade *CreditFacadeTransactor) LiquidateCreditAccount(opts *bind.TransactOpts, borrower common.Address, to common.Address, skipTokenMask *big.Int, convertWETH bool, calls []MultiCall) (*types.Transaction, error) {
	return _CreditFacade.contract.Transact(opts, "liquidateCreditAccount", borrower, to, skipTokenMask, convertWETH, calls)
}

// LiquidateCreditAccount is a paid mutator transaction binding the contract method 0x5d91a0e0.
//
// Solidity: function liquidateCreditAccount(address borrower, address to, uint256 skipTokenMask, bool convertWETH, (address,bytes)[] calls) payable returns()
func (_CreditFacade *CreditFacadeSession) LiquidateCreditAccount(borrower common.Address, to common.Address, skipTokenMask *big.Int, convertWETH bool, calls []MultiCall) (*types.Transaction, error) {
	return _CreditFacade.Contract.LiquidateCreditAccount(&_CreditFacade.TransactOpts, borrower, to, skipTokenMask, convertWETH, calls)
}

// LiquidateCreditAccount is a paid mutator transaction binding the contract method 0x5d91a0e0.
//
// Solidity: function liquidateCreditAccount(address borrower, address to, uint256 skipTokenMask, bool convertWETH, (address,bytes)[] calls) payable returns()
func (_CreditFacade *CreditFacadeTransactorSession) LiquidateCreditAccount(borrower common.Address, to common.Address, skipTokenMask *big.Int, convertWETH bool, calls []MultiCall) (*types.Transaction, error) {
	return _CreditFacade.Contract.LiquidateCreditAccount(&_CreditFacade.TransactOpts, borrower, to, skipTokenMask, convertWETH, calls)
}

// Multicall is a paid mutator transaction binding the contract method 0xcaa5c23f.
//
// Solidity: function multicall((address,bytes)[] calls) payable returns()
func (_CreditFacade *CreditFacadeTransactor) Multicall(opts *bind.TransactOpts, calls []MultiCall) (*types.Transaction, error) {
	return _CreditFacade.contract.Transact(opts, "multicall", calls)
}

// Multicall is a paid mutator transaction binding the contract method 0xcaa5c23f.
//
// Solidity: function multicall((address,bytes)[] calls) payable returns()
func (_CreditFacade *CreditFacadeSession) Multicall(calls []MultiCall) (*types.Transaction, error) {
	return _CreditFacade.Contract.Multicall(&_CreditFacade.TransactOpts, calls)
}

// Multicall is a paid mutator transaction binding the contract method 0xcaa5c23f.
//
// Solidity: function multicall((address,bytes)[] calls) payable returns()
func (_CreditFacade *CreditFacadeTransactorSession) Multicall(calls []MultiCall) (*types.Transaction, error) {
	return _CreditFacade.Contract.Multicall(&_CreditFacade.TransactOpts, calls)
}

// OpenCreditAccount is a paid mutator transaction binding the contract method 0x5288ba4b.
//
// Solidity: function openCreditAccount(uint256 amount, address onBehalfOf, uint256 leverageFactor, uint256 referralCode) payable returns()
func (_CreditFacade *CreditFacadeTransactor) OpenCreditAccount(opts *bind.TransactOpts, amount *big.Int, onBehalfOf common.Address, leverageFactor *big.Int, referralCode *big.Int) (*types.Transaction, error) {
	return _CreditFacade.contract.Transact(opts, "openCreditAccount", amount, onBehalfOf, leverageFactor, referralCode)
}

// OpenCreditAccount is a paid mutator transaction binding the contract method 0x5288ba4b.
//
// Solidity: function openCreditAccount(uint256 amount, address onBehalfOf, uint256 leverageFactor, uint256 referralCode) payable returns()
func (_CreditFacade *CreditFacadeSession) OpenCreditAccount(amount *big.Int, onBehalfOf common.Address, leverageFactor *big.Int, referralCode *big.Int) (*types.Transaction, error) {
	return _CreditFacade.Contract.OpenCreditAccount(&_CreditFacade.TransactOpts, amount, onBehalfOf, leverageFactor, referralCode)
}

// OpenCreditAccount is a paid mutator transaction binding the contract method 0x5288ba4b.
//
// Solidity: function openCreditAccount(uint256 amount, address onBehalfOf, uint256 leverageFactor, uint256 referralCode) payable returns()
func (_CreditFacade *CreditFacadeTransactorSession) OpenCreditAccount(amount *big.Int, onBehalfOf common.Address, leverageFactor *big.Int, referralCode *big.Int) (*types.Transaction, error) {
	return _CreditFacade.Contract.OpenCreditAccount(&_CreditFacade.TransactOpts, amount, onBehalfOf, leverageFactor, referralCode)
}

// OpenCreditAccountMulticall is a paid mutator transaction binding the contract method 0x47639fa8.
//
// Solidity: function openCreditAccountMulticall(uint256 borrowedAmount, address onBehalfOf, (address,bytes)[] calls, uint256 referralCode) payable returns()
func (_CreditFacade *CreditFacadeTransactor) OpenCreditAccountMulticall(opts *bind.TransactOpts, borrowedAmount *big.Int, onBehalfOf common.Address, calls []MultiCall, referralCode *big.Int) (*types.Transaction, error) {
	return _CreditFacade.contract.Transact(opts, "openCreditAccountMulticall", borrowedAmount, onBehalfOf, calls, referralCode)
}

// OpenCreditAccountMulticall is a paid mutator transaction binding the contract method 0x47639fa8.
//
// Solidity: function openCreditAccountMulticall(uint256 borrowedAmount, address onBehalfOf, (address,bytes)[] calls, uint256 referralCode) payable returns()
func (_CreditFacade *CreditFacadeSession) OpenCreditAccountMulticall(borrowedAmount *big.Int, onBehalfOf common.Address, calls []MultiCall, referralCode *big.Int) (*types.Transaction, error) {
	return _CreditFacade.Contract.OpenCreditAccountMulticall(&_CreditFacade.TransactOpts, borrowedAmount, onBehalfOf, calls, referralCode)
}

// OpenCreditAccountMulticall is a paid mutator transaction binding the contract method 0x47639fa8.
//
// Solidity: function openCreditAccountMulticall(uint256 borrowedAmount, address onBehalfOf, (address,bytes)[] calls, uint256 referralCode) payable returns()
func (_CreditFacade *CreditFacadeTransactorSession) OpenCreditAccountMulticall(borrowedAmount *big.Int, onBehalfOf common.Address, calls []MultiCall, referralCode *big.Int) (*types.Transaction, error) {
	return _CreditFacade.Contract.OpenCreditAccountMulticall(&_CreditFacade.TransactOpts, borrowedAmount, onBehalfOf, calls, referralCode)
}

// SetContractToAdapter is a paid mutator transaction binding the contract method 0x28a144f2.
//
// Solidity: function setContractToAdapter(address _adapter, address _contract) returns()
func (_CreditFacade *CreditFacadeTransactor) SetContractToAdapter(opts *bind.TransactOpts, _adapter common.Address, _contract common.Address) (*types.Transaction, error) {
	return _CreditFacade.contract.Transact(opts, "setContractToAdapter", _adapter, _contract)
}

// SetContractToAdapter is a paid mutator transaction binding the contract method 0x28a144f2.
//
// Solidity: function setContractToAdapter(address _adapter, address _contract) returns()
func (_CreditFacade *CreditFacadeSession) SetContractToAdapter(_adapter common.Address, _contract common.Address) (*types.Transaction, error) {
	return _CreditFacade.Contract.SetContractToAdapter(&_CreditFacade.TransactOpts, _adapter, _contract)
}

// SetContractToAdapter is a paid mutator transaction binding the contract method 0x28a144f2.
//
// Solidity: function setContractToAdapter(address _adapter, address _contract) returns()
func (_CreditFacade *CreditFacadeTransactorSession) SetContractToAdapter(_adapter common.Address, _contract common.Address) (*types.Transaction, error) {
	return _CreditFacade.Contract.SetContractToAdapter(&_CreditFacade.TransactOpts, _adapter, _contract)
}

// TransferAccountOwnership is a paid mutator transaction binding the contract method 0x5019e20a.
//
// Solidity: function transferAccountOwnership(address to) returns()
func (_CreditFacade *CreditFacadeTransactor) TransferAccountOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _CreditFacade.contract.Transact(opts, "transferAccountOwnership", to)
}

// TransferAccountOwnership is a paid mutator transaction binding the contract method 0x5019e20a.
//
// Solidity: function transferAccountOwnership(address to) returns()
func (_CreditFacade *CreditFacadeSession) TransferAccountOwnership(to common.Address) (*types.Transaction, error) {
	return _CreditFacade.Contract.TransferAccountOwnership(&_CreditFacade.TransactOpts, to)
}

// TransferAccountOwnership is a paid mutator transaction binding the contract method 0x5019e20a.
//
// Solidity: function transferAccountOwnership(address to) returns()
func (_CreditFacade *CreditFacadeTransactorSession) TransferAccountOwnership(to common.Address) (*types.Transaction, error) {
	return _CreditFacade.Contract.TransferAccountOwnership(&_CreditFacade.TransactOpts, to)
}

// CreditFacadeAddCollateralIterator is returned from FilterAddCollateral and is used to iterate over the raw logs and unpacked data for AddCollateral events raised by the CreditFacade contract.
type CreditFacadeAddCollateralIterator struct {
	Event *CreditFacadeAddCollateral // Event containing the contract specifics and raw log

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
func (it *CreditFacadeAddCollateralIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditFacadeAddCollateral)
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
		it.Event = new(CreditFacadeAddCollateral)
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
func (it *CreditFacadeAddCollateralIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditFacadeAddCollateralIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditFacadeAddCollateral represents a AddCollateral event raised by the CreditFacade contract.
type CreditFacadeAddCollateral struct {
	OnBehalfOf common.Address
	Token      common.Address
	Value      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAddCollateral is a free log retrieval operation binding the contract event 0xa32435755c235de2976ed44a75a2f85cb01faf0c894f639fe0c32bb9455fea8f.
//
// Solidity: event AddCollateral(address indexed onBehalfOf, address indexed token, uint256 value)
func (_CreditFacade *CreditFacadeFilterer) FilterAddCollateral(opts *bind.FilterOpts, onBehalfOf []common.Address, token []common.Address) (*CreditFacadeAddCollateralIterator, error) {

	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CreditFacade.contract.FilterLogs(opts, "AddCollateral", onBehalfOfRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &CreditFacadeAddCollateralIterator{contract: _CreditFacade.contract, event: "AddCollateral", logs: logs, sub: sub}, nil
}

// WatchAddCollateral is a free log subscription operation binding the contract event 0xa32435755c235de2976ed44a75a2f85cb01faf0c894f639fe0c32bb9455fea8f.
//
// Solidity: event AddCollateral(address indexed onBehalfOf, address indexed token, uint256 value)
func (_CreditFacade *CreditFacadeFilterer) WatchAddCollateral(opts *bind.WatchOpts, sink chan<- *CreditFacadeAddCollateral, onBehalfOf []common.Address, token []common.Address) (event.Subscription, error) {

	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CreditFacade.contract.WatchLogs(opts, "AddCollateral", onBehalfOfRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditFacadeAddCollateral)
				if err := _CreditFacade.contract.UnpackLog(event, "AddCollateral", log); err != nil {
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
func (_CreditFacade *CreditFacadeFilterer) ParseAddCollateral(log types.Log) (*CreditFacadeAddCollateral, error) {
	event := new(CreditFacadeAddCollateral)
	if err := _CreditFacade.contract.UnpackLog(event, "AddCollateral", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CreditFacadeCloseCreditAccountIterator is returned from FilterCloseCreditAccount and is used to iterate over the raw logs and unpacked data for CloseCreditAccount events raised by the CreditFacade contract.
type CreditFacadeCloseCreditAccountIterator struct {
	Event *CreditFacadeCloseCreditAccount // Event containing the contract specifics and raw log

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
func (it *CreditFacadeCloseCreditAccountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditFacadeCloseCreditAccount)
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
		it.Event = new(CreditFacadeCloseCreditAccount)
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
func (it *CreditFacadeCloseCreditAccountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditFacadeCloseCreditAccountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditFacadeCloseCreditAccount represents a CloseCreditAccount event raised by the CreditFacade contract.
type CreditFacadeCloseCreditAccount struct {
	Owner common.Address
	To    common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterCloseCreditAccount is a free log retrieval operation binding the contract event 0x460ad03b1cf79b1d64d3aefa28475f110ab66e84649c52bb41ed796b9b391981.
//
// Solidity: event CloseCreditAccount(address indexed owner, address indexed to)
func (_CreditFacade *CreditFacadeFilterer) FilterCloseCreditAccount(opts *bind.FilterOpts, owner []common.Address, to []common.Address) (*CreditFacadeCloseCreditAccountIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CreditFacade.contract.FilterLogs(opts, "CloseCreditAccount", ownerRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CreditFacadeCloseCreditAccountIterator{contract: _CreditFacade.contract, event: "CloseCreditAccount", logs: logs, sub: sub}, nil
}

// WatchCloseCreditAccount is a free log subscription operation binding the contract event 0x460ad03b1cf79b1d64d3aefa28475f110ab66e84649c52bb41ed796b9b391981.
//
// Solidity: event CloseCreditAccount(address indexed owner, address indexed to)
func (_CreditFacade *CreditFacadeFilterer) WatchCloseCreditAccount(opts *bind.WatchOpts, sink chan<- *CreditFacadeCloseCreditAccount, owner []common.Address, to []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CreditFacade.contract.WatchLogs(opts, "CloseCreditAccount", ownerRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditFacadeCloseCreditAccount)
				if err := _CreditFacade.contract.UnpackLog(event, "CloseCreditAccount", log); err != nil {
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

// ParseCloseCreditAccount is a log parse operation binding the contract event 0x460ad03b1cf79b1d64d3aefa28475f110ab66e84649c52bb41ed796b9b391981.
//
// Solidity: event CloseCreditAccount(address indexed owner, address indexed to)
func (_CreditFacade *CreditFacadeFilterer) ParseCloseCreditAccount(log types.Log) (*CreditFacadeCloseCreditAccount, error) {
	event := new(CreditFacadeCloseCreditAccount)
	if err := _CreditFacade.contract.UnpackLog(event, "CloseCreditAccount", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CreditFacadeDecreaseBorrowedAmountIterator is returned from FilterDecreaseBorrowedAmount and is used to iterate over the raw logs and unpacked data for DecreaseBorrowedAmount events raised by the CreditFacade contract.
type CreditFacadeDecreaseBorrowedAmountIterator struct {
	Event *CreditFacadeDecreaseBorrowedAmount // Event containing the contract specifics and raw log

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
func (it *CreditFacadeDecreaseBorrowedAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditFacadeDecreaseBorrowedAmount)
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
		it.Event = new(CreditFacadeDecreaseBorrowedAmount)
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
func (it *CreditFacadeDecreaseBorrowedAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditFacadeDecreaseBorrowedAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditFacadeDecreaseBorrowedAmount represents a DecreaseBorrowedAmount event raised by the CreditFacade contract.
type CreditFacadeDecreaseBorrowedAmount struct {
	Borrower common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDecreaseBorrowedAmount is a free log retrieval operation binding the contract event 0x98274bf834d179ee08dc0604071b0dc90b54731bd5f725a5a96a39a86bce025a.
//
// Solidity: event DecreaseBorrowedAmount(address indexed borrower, uint256 amount)
func (_CreditFacade *CreditFacadeFilterer) FilterDecreaseBorrowedAmount(opts *bind.FilterOpts, borrower []common.Address) (*CreditFacadeDecreaseBorrowedAmountIterator, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _CreditFacade.contract.FilterLogs(opts, "DecreaseBorrowedAmount", borrowerRule)
	if err != nil {
		return nil, err
	}
	return &CreditFacadeDecreaseBorrowedAmountIterator{contract: _CreditFacade.contract, event: "DecreaseBorrowedAmount", logs: logs, sub: sub}, nil
}

// WatchDecreaseBorrowedAmount is a free log subscription operation binding the contract event 0x98274bf834d179ee08dc0604071b0dc90b54731bd5f725a5a96a39a86bce025a.
//
// Solidity: event DecreaseBorrowedAmount(address indexed borrower, uint256 amount)
func (_CreditFacade *CreditFacadeFilterer) WatchDecreaseBorrowedAmount(opts *bind.WatchOpts, sink chan<- *CreditFacadeDecreaseBorrowedAmount, borrower []common.Address) (event.Subscription, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _CreditFacade.contract.WatchLogs(opts, "DecreaseBorrowedAmount", borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditFacadeDecreaseBorrowedAmount)
				if err := _CreditFacade.contract.UnpackLog(event, "DecreaseBorrowedAmount", log); err != nil {
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

// ParseDecreaseBorrowedAmount is a log parse operation binding the contract event 0x98274bf834d179ee08dc0604071b0dc90b54731bd5f725a5a96a39a86bce025a.
//
// Solidity: event DecreaseBorrowedAmount(address indexed borrower, uint256 amount)
func (_CreditFacade *CreditFacadeFilterer) ParseDecreaseBorrowedAmount(log types.Log) (*CreditFacadeDecreaseBorrowedAmount, error) {
	event := new(CreditFacadeDecreaseBorrowedAmount)
	if err := _CreditFacade.contract.UnpackLog(event, "DecreaseBorrowedAmount", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CreditFacadeIncreaseBorrowedAmountIterator is returned from FilterIncreaseBorrowedAmount and is used to iterate over the raw logs and unpacked data for IncreaseBorrowedAmount events raised by the CreditFacade contract.
type CreditFacadeIncreaseBorrowedAmountIterator struct {
	Event *CreditFacadeIncreaseBorrowedAmount // Event containing the contract specifics and raw log

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
func (it *CreditFacadeIncreaseBorrowedAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditFacadeIncreaseBorrowedAmount)
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
		it.Event = new(CreditFacadeIncreaseBorrowedAmount)
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
func (it *CreditFacadeIncreaseBorrowedAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditFacadeIncreaseBorrowedAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditFacadeIncreaseBorrowedAmount represents a IncreaseBorrowedAmount event raised by the CreditFacade contract.
type CreditFacadeIncreaseBorrowedAmount struct {
	Borrower common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterIncreaseBorrowedAmount is a free log retrieval operation binding the contract event 0x9cac51154cc0d835e2f9c9d1f59a9344588cee107f4203bf58a8c797e3a58c45.
//
// Solidity: event IncreaseBorrowedAmount(address indexed borrower, uint256 amount)
func (_CreditFacade *CreditFacadeFilterer) FilterIncreaseBorrowedAmount(opts *bind.FilterOpts, borrower []common.Address) (*CreditFacadeIncreaseBorrowedAmountIterator, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _CreditFacade.contract.FilterLogs(opts, "IncreaseBorrowedAmount", borrowerRule)
	if err != nil {
		return nil, err
	}
	return &CreditFacadeIncreaseBorrowedAmountIterator{contract: _CreditFacade.contract, event: "IncreaseBorrowedAmount", logs: logs, sub: sub}, nil
}

// WatchIncreaseBorrowedAmount is a free log subscription operation binding the contract event 0x9cac51154cc0d835e2f9c9d1f59a9344588cee107f4203bf58a8c797e3a58c45.
//
// Solidity: event IncreaseBorrowedAmount(address indexed borrower, uint256 amount)
func (_CreditFacade *CreditFacadeFilterer) WatchIncreaseBorrowedAmount(opts *bind.WatchOpts, sink chan<- *CreditFacadeIncreaseBorrowedAmount, borrower []common.Address) (event.Subscription, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _CreditFacade.contract.WatchLogs(opts, "IncreaseBorrowedAmount", borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditFacadeIncreaseBorrowedAmount)
				if err := _CreditFacade.contract.UnpackLog(event, "IncreaseBorrowedAmount", log); err != nil {
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
func (_CreditFacade *CreditFacadeFilterer) ParseIncreaseBorrowedAmount(log types.Log) (*CreditFacadeIncreaseBorrowedAmount, error) {
	event := new(CreditFacadeIncreaseBorrowedAmount)
	if err := _CreditFacade.contract.UnpackLog(event, "IncreaseBorrowedAmount", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CreditFacadeLiquidateCreditAccountIterator is returned from FilterLiquidateCreditAccount and is used to iterate over the raw logs and unpacked data for LiquidateCreditAccount events raised by the CreditFacade contract.
type CreditFacadeLiquidateCreditAccountIterator struct {
	Event *CreditFacadeLiquidateCreditAccount // Event containing the contract specifics and raw log

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
func (it *CreditFacadeLiquidateCreditAccountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditFacadeLiquidateCreditAccount)
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
		it.Event = new(CreditFacadeLiquidateCreditAccount)
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
func (it *CreditFacadeLiquidateCreditAccountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditFacadeLiquidateCreditAccountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditFacadeLiquidateCreditAccount represents a LiquidateCreditAccount event raised by the CreditFacade contract.
type CreditFacadeLiquidateCreditAccount struct {
	Owner          common.Address
	Liquidator     common.Address
	To             common.Address
	RemainingFunds *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLiquidateCreditAccount is a free log retrieval operation binding the contract event 0x7dfecd8419723a9d3954585a30c2a270165d70aafa146c11c1e1b88ae1439064.
//
// Solidity: event LiquidateCreditAccount(address indexed owner, address indexed liquidator, address indexed to, uint256 remainingFunds)
func (_CreditFacade *CreditFacadeFilterer) FilterLiquidateCreditAccount(opts *bind.FilterOpts, owner []common.Address, liquidator []common.Address, to []common.Address) (*CreditFacadeLiquidateCreditAccountIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var liquidatorRule []interface{}
	for _, liquidatorItem := range liquidator {
		liquidatorRule = append(liquidatorRule, liquidatorItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CreditFacade.contract.FilterLogs(opts, "LiquidateCreditAccount", ownerRule, liquidatorRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CreditFacadeLiquidateCreditAccountIterator{contract: _CreditFacade.contract, event: "LiquidateCreditAccount", logs: logs, sub: sub}, nil
}

// WatchLiquidateCreditAccount is a free log subscription operation binding the contract event 0x7dfecd8419723a9d3954585a30c2a270165d70aafa146c11c1e1b88ae1439064.
//
// Solidity: event LiquidateCreditAccount(address indexed owner, address indexed liquidator, address indexed to, uint256 remainingFunds)
func (_CreditFacade *CreditFacadeFilterer) WatchLiquidateCreditAccount(opts *bind.WatchOpts, sink chan<- *CreditFacadeLiquidateCreditAccount, owner []common.Address, liquidator []common.Address, to []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var liquidatorRule []interface{}
	for _, liquidatorItem := range liquidator {
		liquidatorRule = append(liquidatorRule, liquidatorItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CreditFacade.contract.WatchLogs(opts, "LiquidateCreditAccount", ownerRule, liquidatorRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditFacadeLiquidateCreditAccount)
				if err := _CreditFacade.contract.UnpackLog(event, "LiquidateCreditAccount", log); err != nil {
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

// ParseLiquidateCreditAccount is a log parse operation binding the contract event 0x7dfecd8419723a9d3954585a30c2a270165d70aafa146c11c1e1b88ae1439064.
//
// Solidity: event LiquidateCreditAccount(address indexed owner, address indexed liquidator, address indexed to, uint256 remainingFunds)
func (_CreditFacade *CreditFacadeFilterer) ParseLiquidateCreditAccount(log types.Log) (*CreditFacadeLiquidateCreditAccount, error) {
	event := new(CreditFacadeLiquidateCreditAccount)
	if err := _CreditFacade.contract.UnpackLog(event, "LiquidateCreditAccount", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CreditFacadeMultiCallFinishedIterator is returned from FilterMultiCallFinished and is used to iterate over the raw logs and unpacked data for MultiCallFinished events raised by the CreditFacade contract.
type CreditFacadeMultiCallFinishedIterator struct {
	Event *CreditFacadeMultiCallFinished // Event containing the contract specifics and raw log

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
func (it *CreditFacadeMultiCallFinishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditFacadeMultiCallFinished)
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
		it.Event = new(CreditFacadeMultiCallFinished)
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
func (it *CreditFacadeMultiCallFinishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditFacadeMultiCallFinishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditFacadeMultiCallFinished represents a MultiCallFinished event raised by the CreditFacade contract.
type CreditFacadeMultiCallFinished struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterMultiCallFinished is a free log retrieval operation binding the contract event 0x60c8e8938c9a0c0d88a98d6f1c562ce68077e12bf3edb8047378f2f736cb45b4.
//
// Solidity: event MultiCallFinished()
func (_CreditFacade *CreditFacadeFilterer) FilterMultiCallFinished(opts *bind.FilterOpts) (*CreditFacadeMultiCallFinishedIterator, error) {

	logs, sub, err := _CreditFacade.contract.FilterLogs(opts, "MultiCallFinished")
	if err != nil {
		return nil, err
	}
	return &CreditFacadeMultiCallFinishedIterator{contract: _CreditFacade.contract, event: "MultiCallFinished", logs: logs, sub: sub}, nil
}

// WatchMultiCallFinished is a free log subscription operation binding the contract event 0x60c8e8938c9a0c0d88a98d6f1c562ce68077e12bf3edb8047378f2f736cb45b4.
//
// Solidity: event MultiCallFinished()
func (_CreditFacade *CreditFacadeFilterer) WatchMultiCallFinished(opts *bind.WatchOpts, sink chan<- *CreditFacadeMultiCallFinished) (event.Subscription, error) {

	logs, sub, err := _CreditFacade.contract.WatchLogs(opts, "MultiCallFinished")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditFacadeMultiCallFinished)
				if err := _CreditFacade.contract.UnpackLog(event, "MultiCallFinished", log); err != nil {
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

// ParseMultiCallFinished is a log parse operation binding the contract event 0x60c8e8938c9a0c0d88a98d6f1c562ce68077e12bf3edb8047378f2f736cb45b4.
//
// Solidity: event MultiCallFinished()
func (_CreditFacade *CreditFacadeFilterer) ParseMultiCallFinished(log types.Log) (*CreditFacadeMultiCallFinished, error) {
	event := new(CreditFacadeMultiCallFinished)
	if err := _CreditFacade.contract.UnpackLog(event, "MultiCallFinished", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CreditFacadeMultiCallStartedIterator is returned from FilterMultiCallStarted and is used to iterate over the raw logs and unpacked data for MultiCallStarted events raised by the CreditFacade contract.
type CreditFacadeMultiCallStartedIterator struct {
	Event *CreditFacadeMultiCallStarted // Event containing the contract specifics and raw log

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
func (it *CreditFacadeMultiCallStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditFacadeMultiCallStarted)
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
		it.Event = new(CreditFacadeMultiCallStarted)
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
func (it *CreditFacadeMultiCallStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditFacadeMultiCallStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditFacadeMultiCallStarted represents a MultiCallStarted event raised by the CreditFacade contract.
type CreditFacadeMultiCallStarted struct {
	Borrower common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMultiCallStarted is a free log retrieval operation binding the contract event 0x4ad424605b950d17d87835716d98c0cac1f6ff9c38114e63304902188a690811.
//
// Solidity: event MultiCallStarted(address indexed borrower)
func (_CreditFacade *CreditFacadeFilterer) FilterMultiCallStarted(opts *bind.FilterOpts, borrower []common.Address) (*CreditFacadeMultiCallStartedIterator, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _CreditFacade.contract.FilterLogs(opts, "MultiCallStarted", borrowerRule)
	if err != nil {
		return nil, err
	}
	return &CreditFacadeMultiCallStartedIterator{contract: _CreditFacade.contract, event: "MultiCallStarted", logs: logs, sub: sub}, nil
}

// WatchMultiCallStarted is a free log subscription operation binding the contract event 0x4ad424605b950d17d87835716d98c0cac1f6ff9c38114e63304902188a690811.
//
// Solidity: event MultiCallStarted(address indexed borrower)
func (_CreditFacade *CreditFacadeFilterer) WatchMultiCallStarted(opts *bind.WatchOpts, sink chan<- *CreditFacadeMultiCallStarted, borrower []common.Address) (event.Subscription, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _CreditFacade.contract.WatchLogs(opts, "MultiCallStarted", borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditFacadeMultiCallStarted)
				if err := _CreditFacade.contract.UnpackLog(event, "MultiCallStarted", log); err != nil {
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

// ParseMultiCallStarted is a log parse operation binding the contract event 0x4ad424605b950d17d87835716d98c0cac1f6ff9c38114e63304902188a690811.
//
// Solidity: event MultiCallStarted(address indexed borrower)
func (_CreditFacade *CreditFacadeFilterer) ParseMultiCallStarted(log types.Log) (*CreditFacadeMultiCallStarted, error) {
	event := new(CreditFacadeMultiCallStarted)
	if err := _CreditFacade.contract.UnpackLog(event, "MultiCallStarted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CreditFacadeOpenCreditAccountIterator is returned from FilterOpenCreditAccount and is used to iterate over the raw logs and unpacked data for OpenCreditAccount events raised by the CreditFacade contract.
type CreditFacadeOpenCreditAccountIterator struct {
	Event *CreditFacadeOpenCreditAccount // Event containing the contract specifics and raw log

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
func (it *CreditFacadeOpenCreditAccountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditFacadeOpenCreditAccount)
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
		it.Event = new(CreditFacadeOpenCreditAccount)
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
func (it *CreditFacadeOpenCreditAccountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditFacadeOpenCreditAccountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditFacadeOpenCreditAccount represents a OpenCreditAccount event raised by the CreditFacade contract.
type CreditFacadeOpenCreditAccount struct {
	OnBehalfOf    common.Address
	CreditAccount common.Address
	BorrowAmount  *big.Int
	ReferralCode  *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOpenCreditAccount is a free log retrieval operation binding the contract event 0xb74d1a34ced5680126a90a8c87ca8dfb02bf84b1921cc56349df6501fad30326.
//
// Solidity: event OpenCreditAccount(address indexed onBehalfOf, address indexed creditAccount, uint256 borrowAmount, uint256 referralCode)
func (_CreditFacade *CreditFacadeFilterer) FilterOpenCreditAccount(opts *bind.FilterOpts, onBehalfOf []common.Address, creditAccount []common.Address) (*CreditFacadeOpenCreditAccountIterator, error) {

	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}
	var creditAccountRule []interface{}
	for _, creditAccountItem := range creditAccount {
		creditAccountRule = append(creditAccountRule, creditAccountItem)
	}

	logs, sub, err := _CreditFacade.contract.FilterLogs(opts, "OpenCreditAccount", onBehalfOfRule, creditAccountRule)
	if err != nil {
		return nil, err
	}
	return &CreditFacadeOpenCreditAccountIterator{contract: _CreditFacade.contract, event: "OpenCreditAccount", logs: logs, sub: sub}, nil
}

// WatchOpenCreditAccount is a free log subscription operation binding the contract event 0xb74d1a34ced5680126a90a8c87ca8dfb02bf84b1921cc56349df6501fad30326.
//
// Solidity: event OpenCreditAccount(address indexed onBehalfOf, address indexed creditAccount, uint256 borrowAmount, uint256 referralCode)
func (_CreditFacade *CreditFacadeFilterer) WatchOpenCreditAccount(opts *bind.WatchOpts, sink chan<- *CreditFacadeOpenCreditAccount, onBehalfOf []common.Address, creditAccount []common.Address) (event.Subscription, error) {

	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}
	var creditAccountRule []interface{}
	for _, creditAccountItem := range creditAccount {
		creditAccountRule = append(creditAccountRule, creditAccountItem)
	}

	logs, sub, err := _CreditFacade.contract.WatchLogs(opts, "OpenCreditAccount", onBehalfOfRule, creditAccountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditFacadeOpenCreditAccount)
				if err := _CreditFacade.contract.UnpackLog(event, "OpenCreditAccount", log); err != nil {
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

// ParseOpenCreditAccount is a log parse operation binding the contract event 0xb74d1a34ced5680126a90a8c87ca8dfb02bf84b1921cc56349df6501fad30326.
//
// Solidity: event OpenCreditAccount(address indexed onBehalfOf, address indexed creditAccount, uint256 borrowAmount, uint256 referralCode)
func (_CreditFacade *CreditFacadeFilterer) ParseOpenCreditAccount(log types.Log) (*CreditFacadeOpenCreditAccount, error) {
	event := new(CreditFacadeOpenCreditAccount)
	if err := _CreditFacade.contract.UnpackLog(event, "OpenCreditAccount", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CreditFacadeTransferAccountIterator is returned from FilterTransferAccount and is used to iterate over the raw logs and unpacked data for TransferAccount events raised by the CreditFacade contract.
type CreditFacadeTransferAccountIterator struct {
	Event *CreditFacadeTransferAccount // Event containing the contract specifics and raw log

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
func (it *CreditFacadeTransferAccountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditFacadeTransferAccount)
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
		it.Event = new(CreditFacadeTransferAccount)
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
func (it *CreditFacadeTransferAccountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditFacadeTransferAccountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditFacadeTransferAccount represents a TransferAccount event raised by the CreditFacade contract.
type CreditFacadeTransferAccount struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferAccount is a free log retrieval operation binding the contract event 0x93c70cc9715bef0d83edf2095f3595402279d274f402a73ffc17f1bcb19d863d.
//
// Solidity: event TransferAccount(address indexed oldOwner, address indexed newOwner)
func (_CreditFacade *CreditFacadeFilterer) FilterTransferAccount(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*CreditFacadeTransferAccountIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CreditFacade.contract.FilterLogs(opts, "TransferAccount", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CreditFacadeTransferAccountIterator{contract: _CreditFacade.contract, event: "TransferAccount", logs: logs, sub: sub}, nil
}

// WatchTransferAccount is a free log subscription operation binding the contract event 0x93c70cc9715bef0d83edf2095f3595402279d274f402a73ffc17f1bcb19d863d.
//
// Solidity: event TransferAccount(address indexed oldOwner, address indexed newOwner)
func (_CreditFacade *CreditFacadeFilterer) WatchTransferAccount(opts *bind.WatchOpts, sink chan<- *CreditFacadeTransferAccount, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CreditFacade.contract.WatchLogs(opts, "TransferAccount", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditFacadeTransferAccount)
				if err := _CreditFacade.contract.UnpackLog(event, "TransferAccount", log); err != nil {
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
func (_CreditFacade *CreditFacadeFilterer) ParseTransferAccount(log types.Log) (*CreditFacadeTransferAccount, error) {
	event := new(CreditFacadeTransferAccount)
	if err := _CreditFacade.contract.UnpackLog(event, "TransferAccount", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CreditFacadeTransferAccountAllowedIterator is returned from FilterTransferAccountAllowed and is used to iterate over the raw logs and unpacked data for TransferAccountAllowed events raised by the CreditFacade contract.
type CreditFacadeTransferAccountAllowedIterator struct {
	Event *CreditFacadeTransferAccountAllowed // Event containing the contract specifics and raw log

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
func (it *CreditFacadeTransferAccountAllowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditFacadeTransferAccountAllowed)
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
		it.Event = new(CreditFacadeTransferAccountAllowed)
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
func (it *CreditFacadeTransferAccountAllowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditFacadeTransferAccountAllowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditFacadeTransferAccountAllowed represents a TransferAccountAllowed event raised by the CreditFacade contract.
type CreditFacadeTransferAccountAllowed struct {
	From  common.Address
	To    common.Address
	State bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransferAccountAllowed is a free log retrieval operation binding the contract event 0x9b3258bc4904fd6426b99843e206c6c7cdb1fd0f040121c25b71dafbb3851ee0.
//
// Solidity: event TransferAccountAllowed(address indexed from, address indexed to, bool state)
func (_CreditFacade *CreditFacadeFilterer) FilterTransferAccountAllowed(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CreditFacadeTransferAccountAllowedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CreditFacade.contract.FilterLogs(opts, "TransferAccountAllowed", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CreditFacadeTransferAccountAllowedIterator{contract: _CreditFacade.contract, event: "TransferAccountAllowed", logs: logs, sub: sub}, nil
}

// WatchTransferAccountAllowed is a free log subscription operation binding the contract event 0x9b3258bc4904fd6426b99843e206c6c7cdb1fd0f040121c25b71dafbb3851ee0.
//
// Solidity: event TransferAccountAllowed(address indexed from, address indexed to, bool state)
func (_CreditFacade *CreditFacadeFilterer) WatchTransferAccountAllowed(opts *bind.WatchOpts, sink chan<- *CreditFacadeTransferAccountAllowed, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CreditFacade.contract.WatchLogs(opts, "TransferAccountAllowed", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditFacadeTransferAccountAllowed)
				if err := _CreditFacade.contract.UnpackLog(event, "TransferAccountAllowed", log); err != nil {
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

// ParseTransferAccountAllowed is a log parse operation binding the contract event 0x9b3258bc4904fd6426b99843e206c6c7cdb1fd0f040121c25b71dafbb3851ee0.
//
// Solidity: event TransferAccountAllowed(address indexed from, address indexed to, bool state)
func (_CreditFacade *CreditFacadeFilterer) ParseTransferAccountAllowed(log types.Log) (*CreditFacadeTransferAccountAllowed, error) {
	event := new(CreditFacadeTransferAccountAllowed)
	if err := _CreditFacade.contract.UnpackLog(event, "TransferAccountAllowed", log); err != nil {
		return nil, err
	}
	return event, nil
}
