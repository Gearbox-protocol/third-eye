// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package leveragedActions

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

// ISwapRouterExactInputParams is an auto generated low-level Go binding around an user-defined struct.
type ISwapRouterExactInputParams struct {
	Path             []byte
	Recipient        common.Address
	Deadline         *big.Int
	AmountIn         *big.Int
	AmountOutMinimum *big.Int
}

// LeveragedActionsLongParameters is an auto generated low-level Go binding around an user-defined struct.
type LeveragedActionsLongParameters struct {
	CreditManager  common.Address
	LeverageFactor *big.Int
	SwapInterface  *big.Int
	SwapContract   common.Address
	SwapCalldata   []byte
	LpInterface    *big.Int
	LpContract     common.Address
	AmountOutMin   *big.Int
}

// LeveragedActionsMetaData contains all meta data concerning the LeveragedActions contract.
var LeveragedActionsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addressProvider\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"shortSwapContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"longSwapContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"lpContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"referralCode\",\"type\":\"uint256\"}],\"name\":\"Action\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"contractsRegister\",\"outputs\":[{\"internalType\":\"contractContractsRegister\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditManager\",\"type\":\"address\"}],\"name\":\"isTransferAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditManager\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"leverageFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lpInterface\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"lpContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"referralCode\",\"type\":\"uint256\"}],\"name\":\"openLP\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"creditManager\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"leverageFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapInterface\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"swapContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"swapCalldata\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"lpInterface\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"lpContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"}],\"internalType\":\"structLeveragedActions.LongParameters\",\"name\":\"longParams\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"referralCode\",\"type\":\"uint256\"}],\"name\":\"openLong\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"curvePool\",\"type\":\"address\"},{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"creditManager\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"leverageFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapInterface\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"swapContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"swapCalldata\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"lpInterface\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"lpContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"}],\"internalType\":\"structLeveragedActions.LongParameters\",\"name\":\"longParams\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"referralCode\",\"type\":\"uint256\"}],\"name\":\"openShortCurve\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"creditManager\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"leverageFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapInterface\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"swapContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"swapCalldata\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"lpInterface\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"lpContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"}],\"internalType\":\"structLeveragedActions.LongParameters\",\"name\":\"longParams\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"referralCode\",\"type\":\"uint256\"}],\"name\":\"openShortUniV2\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinimum\",\"type\":\"uint256\"}],\"internalType\":\"structISwapRouter.ExactInputParams\",\"name\":\"paramsV3\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"creditManager\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"leverageFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapInterface\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"swapContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"swapCalldata\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"lpInterface\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"lpContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"}],\"internalType\":\"structLeveragedActions.LongParameters\",\"name\":\"longParams\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"referralCode\",\"type\":\"uint256\"}],\"name\":\"openShortUniV3\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wethGateway\",\"outputs\":[{\"internalType\":\"contractIWETHGateway\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wethToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// LeveragedActionsABI is the input ABI used to generate the binding from.
// Deprecated: Use LeveragedActionsMetaData.ABI instead.
var LeveragedActionsABI = LeveragedActionsMetaData.ABI

// LeveragedActions is an auto generated Go binding around an Ethereum contract.
type LeveragedActions struct {
	LeveragedActionsCaller     // Read-only binding to the contract
	LeveragedActionsTransactor // Write-only binding to the contract
	LeveragedActionsFilterer   // Log filterer for contract events
}

// LeveragedActionsCaller is an auto generated read-only Go binding around an Ethereum contract.
type LeveragedActionsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LeveragedActionsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LeveragedActionsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LeveragedActionsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LeveragedActionsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LeveragedActionsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LeveragedActionsSession struct {
	Contract     *LeveragedActions // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LeveragedActionsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LeveragedActionsCallerSession struct {
	Contract *LeveragedActionsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// LeveragedActionsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LeveragedActionsTransactorSession struct {
	Contract     *LeveragedActionsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// LeveragedActionsRaw is an auto generated low-level Go binding around an Ethereum contract.
type LeveragedActionsRaw struct {
	Contract *LeveragedActions // Generic contract binding to access the raw methods on
}

// LeveragedActionsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LeveragedActionsCallerRaw struct {
	Contract *LeveragedActionsCaller // Generic read-only contract binding to access the raw methods on
}

// LeveragedActionsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LeveragedActionsTransactorRaw struct {
	Contract *LeveragedActionsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLeveragedActions creates a new instance of LeveragedActions, bound to a specific deployed contract.
func NewLeveragedActions(address common.Address, backend bind.ContractBackend) (*LeveragedActions, error) {
	contract, err := bindLeveragedActions(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LeveragedActions{LeveragedActionsCaller: LeveragedActionsCaller{contract: contract}, LeveragedActionsTransactor: LeveragedActionsTransactor{contract: contract}, LeveragedActionsFilterer: LeveragedActionsFilterer{contract: contract}}, nil
}

// NewLeveragedActionsCaller creates a new read-only instance of LeveragedActions, bound to a specific deployed contract.
func NewLeveragedActionsCaller(address common.Address, caller bind.ContractCaller) (*LeveragedActionsCaller, error) {
	contract, err := bindLeveragedActions(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LeveragedActionsCaller{contract: contract}, nil
}

// NewLeveragedActionsTransactor creates a new write-only instance of LeveragedActions, bound to a specific deployed contract.
func NewLeveragedActionsTransactor(address common.Address, transactor bind.ContractTransactor) (*LeveragedActionsTransactor, error) {
	contract, err := bindLeveragedActions(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LeveragedActionsTransactor{contract: contract}, nil
}

// NewLeveragedActionsFilterer creates a new log filterer instance of LeveragedActions, bound to a specific deployed contract.
func NewLeveragedActionsFilterer(address common.Address, filterer bind.ContractFilterer) (*LeveragedActionsFilterer, error) {
	contract, err := bindLeveragedActions(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LeveragedActionsFilterer{contract: contract}, nil
}

// bindLeveragedActions binds a generic wrapper to an already deployed contract.
func bindLeveragedActions(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LeveragedActionsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LeveragedActions *LeveragedActionsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LeveragedActions.Contract.LeveragedActionsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LeveragedActions *LeveragedActionsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LeveragedActions.Contract.LeveragedActionsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LeveragedActions *LeveragedActionsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LeveragedActions.Contract.LeveragedActionsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LeveragedActions *LeveragedActionsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LeveragedActions.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LeveragedActions *LeveragedActionsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LeveragedActions.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LeveragedActions *LeveragedActionsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LeveragedActions.Contract.contract.Transact(opts, method, params...)
}

// ContractsRegister is a free data retrieval call binding the contract method 0x7a0c7b21.
//
// Solidity: function contractsRegister() view returns(address)
func (_LeveragedActions *LeveragedActionsCaller) ContractsRegister(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LeveragedActions.contract.Call(opts, &out, "contractsRegister")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ContractsRegister is a free data retrieval call binding the contract method 0x7a0c7b21.
//
// Solidity: function contractsRegister() view returns(address)
func (_LeveragedActions *LeveragedActionsSession) ContractsRegister() (common.Address, error) {
	return _LeveragedActions.Contract.ContractsRegister(&_LeveragedActions.CallOpts)
}

// ContractsRegister is a free data retrieval call binding the contract method 0x7a0c7b21.
//
// Solidity: function contractsRegister() view returns(address)
func (_LeveragedActions *LeveragedActionsCallerSession) ContractsRegister() (common.Address, error) {
	return _LeveragedActions.Contract.ContractsRegister(&_LeveragedActions.CallOpts)
}

// IsTransferAllowed is a free data retrieval call binding the contract method 0x8822048e.
//
// Solidity: function isTransferAllowed(address creditManager) view returns(bool)
func (_LeveragedActions *LeveragedActionsCaller) IsTransferAllowed(opts *bind.CallOpts, creditManager common.Address) (bool, error) {
	var out []interface{}
	err := _LeveragedActions.contract.Call(opts, &out, "isTransferAllowed", creditManager)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTransferAllowed is a free data retrieval call binding the contract method 0x8822048e.
//
// Solidity: function isTransferAllowed(address creditManager) view returns(bool)
func (_LeveragedActions *LeveragedActionsSession) IsTransferAllowed(creditManager common.Address) (bool, error) {
	return _LeveragedActions.Contract.IsTransferAllowed(&_LeveragedActions.CallOpts, creditManager)
}

// IsTransferAllowed is a free data retrieval call binding the contract method 0x8822048e.
//
// Solidity: function isTransferAllowed(address creditManager) view returns(bool)
func (_LeveragedActions *LeveragedActionsCallerSession) IsTransferAllowed(creditManager common.Address) (bool, error) {
	return _LeveragedActions.Contract.IsTransferAllowed(&_LeveragedActions.CallOpts, creditManager)
}

// WethGateway is a free data retrieval call binding the contract method 0xc5e10eef.
//
// Solidity: function wethGateway() view returns(address)
func (_LeveragedActions *LeveragedActionsCaller) WethGateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LeveragedActions.contract.Call(opts, &out, "wethGateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WethGateway is a free data retrieval call binding the contract method 0xc5e10eef.
//
// Solidity: function wethGateway() view returns(address)
func (_LeveragedActions *LeveragedActionsSession) WethGateway() (common.Address, error) {
	return _LeveragedActions.Contract.WethGateway(&_LeveragedActions.CallOpts)
}

// WethGateway is a free data retrieval call binding the contract method 0xc5e10eef.
//
// Solidity: function wethGateway() view returns(address)
func (_LeveragedActions *LeveragedActionsCallerSession) WethGateway() (common.Address, error) {
	return _LeveragedActions.Contract.WethGateway(&_LeveragedActions.CallOpts)
}

// WethToken is a free data retrieval call binding the contract method 0x4b57b0be.
//
// Solidity: function wethToken() view returns(address)
func (_LeveragedActions *LeveragedActionsCaller) WethToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LeveragedActions.contract.Call(opts, &out, "wethToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WethToken is a free data retrieval call binding the contract method 0x4b57b0be.
//
// Solidity: function wethToken() view returns(address)
func (_LeveragedActions *LeveragedActionsSession) WethToken() (common.Address, error) {
	return _LeveragedActions.Contract.WethToken(&_LeveragedActions.CallOpts)
}

// WethToken is a free data retrieval call binding the contract method 0x4b57b0be.
//
// Solidity: function wethToken() view returns(address)
func (_LeveragedActions *LeveragedActionsCallerSession) WethToken() (common.Address, error) {
	return _LeveragedActions.Contract.WethToken(&_LeveragedActions.CallOpts)
}

// OpenLP is a paid mutator transaction binding the contract method 0x96112676.
//
// Solidity: function openLP(address creditManager, uint256 leverageFactor, uint256 amountIn, uint256 lpInterface, address lpContract, uint256 amountOutMin, uint256 referralCode) payable returns()
func (_LeveragedActions *LeveragedActionsTransactor) OpenLP(opts *bind.TransactOpts, creditManager common.Address, leverageFactor *big.Int, amountIn *big.Int, lpInterface *big.Int, lpContract common.Address, amountOutMin *big.Int, referralCode *big.Int) (*types.Transaction, error) {
	return _LeveragedActions.contract.Transact(opts, "openLP", creditManager, leverageFactor, amountIn, lpInterface, lpContract, amountOutMin, referralCode)
}

// OpenLP is a paid mutator transaction binding the contract method 0x96112676.
//
// Solidity: function openLP(address creditManager, uint256 leverageFactor, uint256 amountIn, uint256 lpInterface, address lpContract, uint256 amountOutMin, uint256 referralCode) payable returns()
func (_LeveragedActions *LeveragedActionsSession) OpenLP(creditManager common.Address, leverageFactor *big.Int, amountIn *big.Int, lpInterface *big.Int, lpContract common.Address, amountOutMin *big.Int, referralCode *big.Int) (*types.Transaction, error) {
	return _LeveragedActions.Contract.OpenLP(&_LeveragedActions.TransactOpts, creditManager, leverageFactor, amountIn, lpInterface, lpContract, amountOutMin, referralCode)
}

// OpenLP is a paid mutator transaction binding the contract method 0x96112676.
//
// Solidity: function openLP(address creditManager, uint256 leverageFactor, uint256 amountIn, uint256 lpInterface, address lpContract, uint256 amountOutMin, uint256 referralCode) payable returns()
func (_LeveragedActions *LeveragedActionsTransactorSession) OpenLP(creditManager common.Address, leverageFactor *big.Int, amountIn *big.Int, lpInterface *big.Int, lpContract common.Address, amountOutMin *big.Int, referralCode *big.Int) (*types.Transaction, error) {
	return _LeveragedActions.Contract.OpenLP(&_LeveragedActions.TransactOpts, creditManager, leverageFactor, amountIn, lpInterface, lpContract, amountOutMin, referralCode)
}

// OpenLong is a paid mutator transaction binding the contract method 0x3e04dcb5.
//
// Solidity: function openLong(uint256 amountIn, (address,uint256,uint256,address,bytes,uint256,address,uint256) longParams, uint256 referralCode) payable returns()
func (_LeveragedActions *LeveragedActionsTransactor) OpenLong(opts *bind.TransactOpts, amountIn *big.Int, longParams LeveragedActionsLongParameters, referralCode *big.Int) (*types.Transaction, error) {
	return _LeveragedActions.contract.Transact(opts, "openLong", amountIn, longParams, referralCode)
}

// OpenLong is a paid mutator transaction binding the contract method 0x3e04dcb5.
//
// Solidity: function openLong(uint256 amountIn, (address,uint256,uint256,address,bytes,uint256,address,uint256) longParams, uint256 referralCode) payable returns()
func (_LeveragedActions *LeveragedActionsSession) OpenLong(amountIn *big.Int, longParams LeveragedActionsLongParameters, referralCode *big.Int) (*types.Transaction, error) {
	return _LeveragedActions.Contract.OpenLong(&_LeveragedActions.TransactOpts, amountIn, longParams, referralCode)
}

// OpenLong is a paid mutator transaction binding the contract method 0x3e04dcb5.
//
// Solidity: function openLong(uint256 amountIn, (address,uint256,uint256,address,bytes,uint256,address,uint256) longParams, uint256 referralCode) payable returns()
func (_LeveragedActions *LeveragedActionsTransactorSession) OpenLong(amountIn *big.Int, longParams LeveragedActionsLongParameters, referralCode *big.Int) (*types.Transaction, error) {
	return _LeveragedActions.Contract.OpenLong(&_LeveragedActions.TransactOpts, amountIn, longParams, referralCode)
}

// OpenShortCurve is a paid mutator transaction binding the contract method 0xa9f59b71.
//
// Solidity: function openShortCurve(address curvePool, int128 i, int128 j, uint256 amountIn, uint256 amountOutMin, (address,uint256,uint256,address,bytes,uint256,address,uint256) longParams, uint256 referralCode) payable returns()
func (_LeveragedActions *LeveragedActionsTransactor) OpenShortCurve(opts *bind.TransactOpts, curvePool common.Address, i *big.Int, j *big.Int, amountIn *big.Int, amountOutMin *big.Int, longParams LeveragedActionsLongParameters, referralCode *big.Int) (*types.Transaction, error) {
	return _LeveragedActions.contract.Transact(opts, "openShortCurve", curvePool, i, j, amountIn, amountOutMin, longParams, referralCode)
}

// OpenShortCurve is a paid mutator transaction binding the contract method 0xa9f59b71.
//
// Solidity: function openShortCurve(address curvePool, int128 i, int128 j, uint256 amountIn, uint256 amountOutMin, (address,uint256,uint256,address,bytes,uint256,address,uint256) longParams, uint256 referralCode) payable returns()
func (_LeveragedActions *LeveragedActionsSession) OpenShortCurve(curvePool common.Address, i *big.Int, j *big.Int, amountIn *big.Int, amountOutMin *big.Int, longParams LeveragedActionsLongParameters, referralCode *big.Int) (*types.Transaction, error) {
	return _LeveragedActions.Contract.OpenShortCurve(&_LeveragedActions.TransactOpts, curvePool, i, j, amountIn, amountOutMin, longParams, referralCode)
}

// OpenShortCurve is a paid mutator transaction binding the contract method 0xa9f59b71.
//
// Solidity: function openShortCurve(address curvePool, int128 i, int128 j, uint256 amountIn, uint256 amountOutMin, (address,uint256,uint256,address,bytes,uint256,address,uint256) longParams, uint256 referralCode) payable returns()
func (_LeveragedActions *LeveragedActionsTransactorSession) OpenShortCurve(curvePool common.Address, i *big.Int, j *big.Int, amountIn *big.Int, amountOutMin *big.Int, longParams LeveragedActionsLongParameters, referralCode *big.Int) (*types.Transaction, error) {
	return _LeveragedActions.Contract.OpenShortCurve(&_LeveragedActions.TransactOpts, curvePool, i, j, amountIn, amountOutMin, longParams, referralCode)
}

// OpenShortUniV2 is a paid mutator transaction binding the contract method 0x601ee770.
//
// Solidity: function openShortUniV2(address router, uint256 amountIn, uint256 amountOutMin, address[] path, uint256 deadline, (address,uint256,uint256,address,bytes,uint256,address,uint256) longParams, uint256 referralCode) payable returns()
func (_LeveragedActions *LeveragedActionsTransactor) OpenShortUniV2(opts *bind.TransactOpts, router common.Address, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, deadline *big.Int, longParams LeveragedActionsLongParameters, referralCode *big.Int) (*types.Transaction, error) {
	return _LeveragedActions.contract.Transact(opts, "openShortUniV2", router, amountIn, amountOutMin, path, deadline, longParams, referralCode)
}

// OpenShortUniV2 is a paid mutator transaction binding the contract method 0x601ee770.
//
// Solidity: function openShortUniV2(address router, uint256 amountIn, uint256 amountOutMin, address[] path, uint256 deadline, (address,uint256,uint256,address,bytes,uint256,address,uint256) longParams, uint256 referralCode) payable returns()
func (_LeveragedActions *LeveragedActionsSession) OpenShortUniV2(router common.Address, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, deadline *big.Int, longParams LeveragedActionsLongParameters, referralCode *big.Int) (*types.Transaction, error) {
	return _LeveragedActions.Contract.OpenShortUniV2(&_LeveragedActions.TransactOpts, router, amountIn, amountOutMin, path, deadline, longParams, referralCode)
}

// OpenShortUniV2 is a paid mutator transaction binding the contract method 0x601ee770.
//
// Solidity: function openShortUniV2(address router, uint256 amountIn, uint256 amountOutMin, address[] path, uint256 deadline, (address,uint256,uint256,address,bytes,uint256,address,uint256) longParams, uint256 referralCode) payable returns()
func (_LeveragedActions *LeveragedActionsTransactorSession) OpenShortUniV2(router common.Address, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, deadline *big.Int, longParams LeveragedActionsLongParameters, referralCode *big.Int) (*types.Transaction, error) {
	return _LeveragedActions.Contract.OpenShortUniV2(&_LeveragedActions.TransactOpts, router, amountIn, amountOutMin, path, deadline, longParams, referralCode)
}

// OpenShortUniV3 is a paid mutator transaction binding the contract method 0x1505ddd3.
//
// Solidity: function openShortUniV3(address router, (bytes,address,uint256,uint256,uint256) paramsV3, (address,uint256,uint256,address,bytes,uint256,address,uint256) longParams, uint256 referralCode) payable returns()
func (_LeveragedActions *LeveragedActionsTransactor) OpenShortUniV3(opts *bind.TransactOpts, router common.Address, paramsV3 ISwapRouterExactInputParams, longParams LeveragedActionsLongParameters, referralCode *big.Int) (*types.Transaction, error) {
	return _LeveragedActions.contract.Transact(opts, "openShortUniV3", router, paramsV3, longParams, referralCode)
}

// OpenShortUniV3 is a paid mutator transaction binding the contract method 0x1505ddd3.
//
// Solidity: function openShortUniV3(address router, (bytes,address,uint256,uint256,uint256) paramsV3, (address,uint256,uint256,address,bytes,uint256,address,uint256) longParams, uint256 referralCode) payable returns()
func (_LeveragedActions *LeveragedActionsSession) OpenShortUniV3(router common.Address, paramsV3 ISwapRouterExactInputParams, longParams LeveragedActionsLongParameters, referralCode *big.Int) (*types.Transaction, error) {
	return _LeveragedActions.Contract.OpenShortUniV3(&_LeveragedActions.TransactOpts, router, paramsV3, longParams, referralCode)
}

// OpenShortUniV3 is a paid mutator transaction binding the contract method 0x1505ddd3.
//
// Solidity: function openShortUniV3(address router, (bytes,address,uint256,uint256,uint256) paramsV3, (address,uint256,uint256,address,bytes,uint256,address,uint256) longParams, uint256 referralCode) payable returns()
func (_LeveragedActions *LeveragedActionsTransactorSession) OpenShortUniV3(router common.Address, paramsV3 ISwapRouterExactInputParams, longParams LeveragedActionsLongParameters, referralCode *big.Int) (*types.Transaction, error) {
	return _LeveragedActions.Contract.OpenShortUniV3(&_LeveragedActions.TransactOpts, router, paramsV3, longParams, referralCode)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LeveragedActions *LeveragedActionsTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LeveragedActions.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LeveragedActions *LeveragedActionsSession) Receive() (*types.Transaction, error) {
	return _LeveragedActions.Contract.Receive(&_LeveragedActions.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LeveragedActions *LeveragedActionsTransactorSession) Receive() (*types.Transaction, error) {
	return _LeveragedActions.Contract.Receive(&_LeveragedActions.TransactOpts)
}

// LeveragedActionsActionIterator is returned from FilterAction and is used to iterate over the raw logs and unpacked data for Action events raised by the LeveragedActions contract.
type LeveragedActionsActionIterator struct {
	Event *LeveragedActionsAction // Event containing the contract specifics and raw log

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
func (it *LeveragedActionsActionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LeveragedActionsAction)
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
		it.Event = new(LeveragedActionsAction)
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
func (it *LeveragedActionsActionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LeveragedActionsActionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LeveragedActionsAction represents a Action event raised by the LeveragedActions contract.
type LeveragedActionsAction struct {
	TokenIn           common.Address
	Collateral        common.Address
	Asset             common.Address
	AmountIn          *big.Int
	ShortSwapContract common.Address
	LongSwapContract  common.Address
	LpContract        common.Address
	ReferralCode      *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterAction is a free log retrieval operation binding the contract event 0xb709cbfc91402c021e4a8c00ac331f0d9c3617c05eb744ec9ee15eb4f2e6c2d5.
//
// Solidity: event Action(address indexed tokenIn, address indexed collateral, address indexed asset, uint256 amountIn, address shortSwapContract, address longSwapContract, address lpContract, uint256 referralCode)
func (_LeveragedActions *LeveragedActionsFilterer) FilterAction(opts *bind.FilterOpts, tokenIn []common.Address, collateral []common.Address, asset []common.Address) (*LeveragedActionsActionIterator, error) {

	var tokenInRule []interface{}
	for _, tokenInItem := range tokenIn {
		tokenInRule = append(tokenInRule, tokenInItem)
	}
	var collateralRule []interface{}
	for _, collateralItem := range collateral {
		collateralRule = append(collateralRule, collateralItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _LeveragedActions.contract.FilterLogs(opts, "Action", tokenInRule, collateralRule, assetRule)
	if err != nil {
		return nil, err
	}
	return &LeveragedActionsActionIterator{contract: _LeveragedActions.contract, event: "Action", logs: logs, sub: sub}, nil
}

// WatchAction is a free log subscription operation binding the contract event 0xb709cbfc91402c021e4a8c00ac331f0d9c3617c05eb744ec9ee15eb4f2e6c2d5.
//
// Solidity: event Action(address indexed tokenIn, address indexed collateral, address indexed asset, uint256 amountIn, address shortSwapContract, address longSwapContract, address lpContract, uint256 referralCode)
func (_LeveragedActions *LeveragedActionsFilterer) WatchAction(opts *bind.WatchOpts, sink chan<- *LeveragedActionsAction, tokenIn []common.Address, collateral []common.Address, asset []common.Address) (event.Subscription, error) {

	var tokenInRule []interface{}
	for _, tokenInItem := range tokenIn {
		tokenInRule = append(tokenInRule, tokenInItem)
	}
	var collateralRule []interface{}
	for _, collateralItem := range collateral {
		collateralRule = append(collateralRule, collateralItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _LeveragedActions.contract.WatchLogs(opts, "Action", tokenInRule, collateralRule, assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LeveragedActionsAction)
				if err := _LeveragedActions.contract.UnpackLog(event, "Action", log); err != nil {
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

// ParseAction is a log parse operation binding the contract event 0xb709cbfc91402c021e4a8c00ac331f0d9c3617c05eb744ec9ee15eb4f2e6c2d5.
//
// Solidity: event Action(address indexed tokenIn, address indexed collateral, address indexed asset, uint256 amountIn, address shortSwapContract, address longSwapContract, address lpContract, uint256 referralCode)
func (_LeveragedActions *LeveragedActionsFilterer) ParseAction(log types.Log) (*LeveragedActionsAction, error) {
	event := new(LeveragedActionsAction)
	if err := _LeveragedActions.contract.UnpackLog(event, "Action", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
