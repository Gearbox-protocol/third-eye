// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package uniswapV3Adapter

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

// ISwapRouterExactInputSingleParams is an auto generated low-level Go binding around an user-defined struct.
type ISwapRouterExactInputSingleParams struct {
	TokenIn           common.Address
	TokenOut          common.Address
	Fee               *big.Int
	Recipient         common.Address
	Deadline          *big.Int
	AmountIn          *big.Int
	AmountOutMinimum  *big.Int
	SqrtPriceLimitX96 *big.Int
}

// ISwapRouterExactOutputParams is an auto generated low-level Go binding around an user-defined struct.
type ISwapRouterExactOutputParams struct {
	Path            []byte
	Recipient       common.Address
	Deadline        *big.Int
	AmountOut       *big.Int
	AmountInMaximum *big.Int
}

// ISwapRouterExactOutputSingleParams is an auto generated low-level Go binding around an user-defined struct.
type ISwapRouterExactOutputSingleParams struct {
	TokenIn           common.Address
	TokenOut          common.Address
	Fee               *big.Int
	Recipient         common.Address
	Deadline          *big.Int
	AmountOut         *big.Int
	AmountInMaximum   *big.Int
	SqrtPriceLimitX96 *big.Int
}

// UniswapV3AdapterMetaData contains all meta data concerning the UniswapV3Adapter contract.
var UniswapV3AdapterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_creditManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"creditFilter\",\"outputs\":[{\"internalType\":\"contractICreditFilter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creditManager\",\"outputs\":[{\"internalType\":\"contractICreditManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinimum\",\"type\":\"uint256\"}],\"internalType\":\"structISwapRouter.ExactInputParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactInput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinimum\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"internalType\":\"structISwapRouter.ExactInputSingleParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactInputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMaximum\",\"type\":\"uint256\"}],\"internalType\":\"structISwapRouter.ExactOutputParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactOutput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMaximum\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"internalType\":\"structISwapRouter.ExactOutputSingleParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactOutputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"router\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// UniswapV3AdapterABI is the input ABI used to generate the binding from.
// Deprecated: Use UniswapV3AdapterMetaData.ABI instead.
var UniswapV3AdapterABI = UniswapV3AdapterMetaData.ABI

// UniswapV3Adapter is an auto generated Go binding around an Ethereum contract.
type UniswapV3Adapter struct {
	UniswapV3AdapterCaller     // Read-only binding to the contract
	UniswapV3AdapterTransactor // Write-only binding to the contract
	UniswapV3AdapterFilterer   // Log filterer for contract events
}

// UniswapV3AdapterCaller is an auto generated read-only Go binding around an Ethereum contract.
type UniswapV3AdapterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV3AdapterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UniswapV3AdapterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV3AdapterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniswapV3AdapterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV3AdapterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniswapV3AdapterSession struct {
	Contract     *UniswapV3Adapter // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UniswapV3AdapterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniswapV3AdapterCallerSession struct {
	Contract *UniswapV3AdapterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// UniswapV3AdapterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniswapV3AdapterTransactorSession struct {
	Contract     *UniswapV3AdapterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// UniswapV3AdapterRaw is an auto generated low-level Go binding around an Ethereum contract.
type UniswapV3AdapterRaw struct {
	Contract *UniswapV3Adapter // Generic contract binding to access the raw methods on
}

// UniswapV3AdapterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniswapV3AdapterCallerRaw struct {
	Contract *UniswapV3AdapterCaller // Generic read-only contract binding to access the raw methods on
}

// UniswapV3AdapterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniswapV3AdapterTransactorRaw struct {
	Contract *UniswapV3AdapterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniswapV3Adapter creates a new instance of UniswapV3Adapter, bound to a specific deployed contract.
func NewUniswapV3Adapter(address common.Address, backend bind.ContractBackend) (*UniswapV3Adapter, error) {
	contract, err := bindUniswapV3Adapter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UniswapV3Adapter{UniswapV3AdapterCaller: UniswapV3AdapterCaller{contract: contract}, UniswapV3AdapterTransactor: UniswapV3AdapterTransactor{contract: contract}, UniswapV3AdapterFilterer: UniswapV3AdapterFilterer{contract: contract}}, nil
}

// NewUniswapV3AdapterCaller creates a new read-only instance of UniswapV3Adapter, bound to a specific deployed contract.
func NewUniswapV3AdapterCaller(address common.Address, caller bind.ContractCaller) (*UniswapV3AdapterCaller, error) {
	contract, err := bindUniswapV3Adapter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapV3AdapterCaller{contract: contract}, nil
}

// NewUniswapV3AdapterTransactor creates a new write-only instance of UniswapV3Adapter, bound to a specific deployed contract.
func NewUniswapV3AdapterTransactor(address common.Address, transactor bind.ContractTransactor) (*UniswapV3AdapterTransactor, error) {
	contract, err := bindUniswapV3Adapter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapV3AdapterTransactor{contract: contract}, nil
}

// NewUniswapV3AdapterFilterer creates a new log filterer instance of UniswapV3Adapter, bound to a specific deployed contract.
func NewUniswapV3AdapterFilterer(address common.Address, filterer bind.ContractFilterer) (*UniswapV3AdapterFilterer, error) {
	contract, err := bindUniswapV3Adapter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniswapV3AdapterFilterer{contract: contract}, nil
}

// bindUniswapV3Adapter binds a generic wrapper to an already deployed contract.
func bindUniswapV3Adapter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UniswapV3AdapterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapV3Adapter *UniswapV3AdapterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapV3Adapter.Contract.UniswapV3AdapterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapV3Adapter *UniswapV3AdapterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV3Adapter.Contract.UniswapV3AdapterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapV3Adapter *UniswapV3AdapterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapV3Adapter.Contract.UniswapV3AdapterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapV3Adapter *UniswapV3AdapterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapV3Adapter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapV3Adapter *UniswapV3AdapterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV3Adapter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapV3Adapter *UniswapV3AdapterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapV3Adapter.Contract.contract.Transact(opts, method, params...)
}

// CreditFilter is a free data retrieval call binding the contract method 0xf93f515b.
//
// Solidity: function creditFilter() view returns(address)
func (_UniswapV3Adapter *UniswapV3AdapterCaller) CreditFilter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniswapV3Adapter.contract.Call(opts, &out, "creditFilter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditFilter is a free data retrieval call binding the contract method 0xf93f515b.
//
// Solidity: function creditFilter() view returns(address)
func (_UniswapV3Adapter *UniswapV3AdapterSession) CreditFilter() (common.Address, error) {
	return _UniswapV3Adapter.Contract.CreditFilter(&_UniswapV3Adapter.CallOpts)
}

// CreditFilter is a free data retrieval call binding the contract method 0xf93f515b.
//
// Solidity: function creditFilter() view returns(address)
func (_UniswapV3Adapter *UniswapV3AdapterCallerSession) CreditFilter() (common.Address, error) {
	return _UniswapV3Adapter.Contract.CreditFilter(&_UniswapV3Adapter.CallOpts)
}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_UniswapV3Adapter *UniswapV3AdapterCaller) CreditManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniswapV3Adapter.contract.Call(opts, &out, "creditManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_UniswapV3Adapter *UniswapV3AdapterSession) CreditManager() (common.Address, error) {
	return _UniswapV3Adapter.Contract.CreditManager(&_UniswapV3Adapter.CallOpts)
}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_UniswapV3Adapter *UniswapV3AdapterCallerSession) CreditManager() (common.Address, error) {
	return _UniswapV3Adapter.Contract.CreditManager(&_UniswapV3Adapter.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_UniswapV3Adapter *UniswapV3AdapterCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniswapV3Adapter.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_UniswapV3Adapter *UniswapV3AdapterSession) Router() (common.Address, error) {
	return _UniswapV3Adapter.Contract.Router(&_UniswapV3Adapter.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_UniswapV3Adapter *UniswapV3AdapterCallerSession) Router() (common.Address, error) {
	return _UniswapV3Adapter.Contract.Router(&_UniswapV3Adapter.CallOpts)
}

// ExactInput is a paid mutator transaction binding the contract method 0xc04b8d59.
//
// Solidity: function exactInput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountOut)
func (_UniswapV3Adapter *UniswapV3AdapterTransactor) ExactInput(opts *bind.TransactOpts, params ISwapRouterExactInputParams) (*types.Transaction, error) {
	return _UniswapV3Adapter.contract.Transact(opts, "exactInput", params)
}

// ExactInput is a paid mutator transaction binding the contract method 0xc04b8d59.
//
// Solidity: function exactInput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountOut)
func (_UniswapV3Adapter *UniswapV3AdapterSession) ExactInput(params ISwapRouterExactInputParams) (*types.Transaction, error) {
	return _UniswapV3Adapter.Contract.ExactInput(&_UniswapV3Adapter.TransactOpts, params)
}

// ExactInput is a paid mutator transaction binding the contract method 0xc04b8d59.
//
// Solidity: function exactInput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountOut)
func (_UniswapV3Adapter *UniswapV3AdapterTransactorSession) ExactInput(params ISwapRouterExactInputParams) (*types.Transaction, error) {
	return _UniswapV3Adapter.Contract.ExactInput(&_UniswapV3Adapter.TransactOpts, params)
}

// ExactInputSingle is a paid mutator transaction binding the contract method 0x414bf389.
//
// Solidity: function exactInputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountOut)
func (_UniswapV3Adapter *UniswapV3AdapterTransactor) ExactInputSingle(opts *bind.TransactOpts, params ISwapRouterExactInputSingleParams) (*types.Transaction, error) {
	return _UniswapV3Adapter.contract.Transact(opts, "exactInputSingle", params)
}

// ExactInputSingle is a paid mutator transaction binding the contract method 0x414bf389.
//
// Solidity: function exactInputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountOut)
func (_UniswapV3Adapter *UniswapV3AdapterSession) ExactInputSingle(params ISwapRouterExactInputSingleParams) (*types.Transaction, error) {
	return _UniswapV3Adapter.Contract.ExactInputSingle(&_UniswapV3Adapter.TransactOpts, params)
}

// ExactInputSingle is a paid mutator transaction binding the contract method 0x414bf389.
//
// Solidity: function exactInputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountOut)
func (_UniswapV3Adapter *UniswapV3AdapterTransactorSession) ExactInputSingle(params ISwapRouterExactInputSingleParams) (*types.Transaction, error) {
	return _UniswapV3Adapter.Contract.ExactInputSingle(&_UniswapV3Adapter.TransactOpts, params)
}

// ExactOutput is a paid mutator transaction binding the contract method 0xf28c0498.
//
// Solidity: function exactOutput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountIn)
func (_UniswapV3Adapter *UniswapV3AdapterTransactor) ExactOutput(opts *bind.TransactOpts, params ISwapRouterExactOutputParams) (*types.Transaction, error) {
	return _UniswapV3Adapter.contract.Transact(opts, "exactOutput", params)
}

// ExactOutput is a paid mutator transaction binding the contract method 0xf28c0498.
//
// Solidity: function exactOutput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountIn)
func (_UniswapV3Adapter *UniswapV3AdapterSession) ExactOutput(params ISwapRouterExactOutputParams) (*types.Transaction, error) {
	return _UniswapV3Adapter.Contract.ExactOutput(&_UniswapV3Adapter.TransactOpts, params)
}

// ExactOutput is a paid mutator transaction binding the contract method 0xf28c0498.
//
// Solidity: function exactOutput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountIn)
func (_UniswapV3Adapter *UniswapV3AdapterTransactorSession) ExactOutput(params ISwapRouterExactOutputParams) (*types.Transaction, error) {
	return _UniswapV3Adapter.Contract.ExactOutput(&_UniswapV3Adapter.TransactOpts, params)
}

// ExactOutputSingle is a paid mutator transaction binding the contract method 0xdb3e2198.
//
// Solidity: function exactOutputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountIn)
func (_UniswapV3Adapter *UniswapV3AdapterTransactor) ExactOutputSingle(opts *bind.TransactOpts, params ISwapRouterExactOutputSingleParams) (*types.Transaction, error) {
	return _UniswapV3Adapter.contract.Transact(opts, "exactOutputSingle", params)
}

// ExactOutputSingle is a paid mutator transaction binding the contract method 0xdb3e2198.
//
// Solidity: function exactOutputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountIn)
func (_UniswapV3Adapter *UniswapV3AdapterSession) ExactOutputSingle(params ISwapRouterExactOutputSingleParams) (*types.Transaction, error) {
	return _UniswapV3Adapter.Contract.ExactOutputSingle(&_UniswapV3Adapter.TransactOpts, params)
}

// ExactOutputSingle is a paid mutator transaction binding the contract method 0xdb3e2198.
//
// Solidity: function exactOutputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountIn)
func (_UniswapV3Adapter *UniswapV3AdapterTransactorSession) ExactOutputSingle(params ISwapRouterExactOutputSingleParams) (*types.Transaction, error) {
	return _UniswapV3Adapter.Contract.ExactOutputSingle(&_UniswapV3Adapter.TransactOpts, params)
}
