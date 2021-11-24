// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package uniswapRouterMock

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

// UniswapRouterMockMetaData contains all meta data concerning the UniswapRouterMock contract.
var UniswapRouterMockMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"addLiquidityETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"getAmountIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"getAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"name\":\"getAmountsIn\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"getAmountsOut\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"getRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"quote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"removeLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"removeLiquidityETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"removeLiquidityETHSupportingFeeOnTransferTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"removeLiquidityETHWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"removeLiquidityETHWithPermitSupportingFeeOnTransferTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"removeLiquidityWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenFrom\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenTo\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"rate_RAY\",\"type\":\"uint256\"}],\"name\":\"setRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"swapETHForExactTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"swapExactETHForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"swapExactETHForTokensSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForETH\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForETHSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokensSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"swapTokensForExactETH\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapTokensForExactTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// UniswapRouterMockABI is the input ABI used to generate the binding from.
// Deprecated: Use UniswapRouterMockMetaData.ABI instead.
var UniswapRouterMockABI = UniswapRouterMockMetaData.ABI

// UniswapRouterMock is an auto generated Go binding around an Ethereum contract.
type UniswapRouterMock struct {
	UniswapRouterMockCaller     // Read-only binding to the contract
	UniswapRouterMockTransactor // Write-only binding to the contract
	UniswapRouterMockFilterer   // Log filterer for contract events
}

// UniswapRouterMockCaller is an auto generated read-only Go binding around an Ethereum contract.
type UniswapRouterMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapRouterMockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UniswapRouterMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapRouterMockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniswapRouterMockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapRouterMockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniswapRouterMockSession struct {
	Contract     *UniswapRouterMock // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// UniswapRouterMockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniswapRouterMockCallerSession struct {
	Contract *UniswapRouterMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// UniswapRouterMockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniswapRouterMockTransactorSession struct {
	Contract     *UniswapRouterMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// UniswapRouterMockRaw is an auto generated low-level Go binding around an Ethereum contract.
type UniswapRouterMockRaw struct {
	Contract *UniswapRouterMock // Generic contract binding to access the raw methods on
}

// UniswapRouterMockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniswapRouterMockCallerRaw struct {
	Contract *UniswapRouterMockCaller // Generic read-only contract binding to access the raw methods on
}

// UniswapRouterMockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniswapRouterMockTransactorRaw struct {
	Contract *UniswapRouterMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniswapRouterMock creates a new instance of UniswapRouterMock, bound to a specific deployed contract.
func NewUniswapRouterMock(address common.Address, backend bind.ContractBackend) (*UniswapRouterMock, error) {
	contract, err := bindUniswapRouterMock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UniswapRouterMock{UniswapRouterMockCaller: UniswapRouterMockCaller{contract: contract}, UniswapRouterMockTransactor: UniswapRouterMockTransactor{contract: contract}, UniswapRouterMockFilterer: UniswapRouterMockFilterer{contract: contract}}, nil
}

// NewUniswapRouterMockCaller creates a new read-only instance of UniswapRouterMock, bound to a specific deployed contract.
func NewUniswapRouterMockCaller(address common.Address, caller bind.ContractCaller) (*UniswapRouterMockCaller, error) {
	contract, err := bindUniswapRouterMock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapRouterMockCaller{contract: contract}, nil
}

// NewUniswapRouterMockTransactor creates a new write-only instance of UniswapRouterMock, bound to a specific deployed contract.
func NewUniswapRouterMockTransactor(address common.Address, transactor bind.ContractTransactor) (*UniswapRouterMockTransactor, error) {
	contract, err := bindUniswapRouterMock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapRouterMockTransactor{contract: contract}, nil
}

// NewUniswapRouterMockFilterer creates a new log filterer instance of UniswapRouterMock, bound to a specific deployed contract.
func NewUniswapRouterMockFilterer(address common.Address, filterer bind.ContractFilterer) (*UniswapRouterMockFilterer, error) {
	contract, err := bindUniswapRouterMock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniswapRouterMockFilterer{contract: contract}, nil
}

// bindUniswapRouterMock binds a generic wrapper to an already deployed contract.
func bindUniswapRouterMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UniswapRouterMockABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapRouterMock *UniswapRouterMockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapRouterMock.Contract.UniswapRouterMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapRouterMock *UniswapRouterMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapRouterMock.Contract.UniswapRouterMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapRouterMock *UniswapRouterMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapRouterMock.Contract.UniswapRouterMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapRouterMock *UniswapRouterMockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapRouterMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapRouterMock *UniswapRouterMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapRouterMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapRouterMock *UniswapRouterMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapRouterMock.Contract.contract.Transact(opts, method, params...)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() pure returns(address)
func (_UniswapRouterMock *UniswapRouterMockCaller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniswapRouterMock.contract.Call(opts, &out, "WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() pure returns(address)
func (_UniswapRouterMock *UniswapRouterMockSession) WETH() (common.Address, error) {
	return _UniswapRouterMock.Contract.WETH(&_UniswapRouterMock.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() pure returns(address)
func (_UniswapRouterMock *UniswapRouterMockCallerSession) WETH() (common.Address, error) {
	return _UniswapRouterMock.Contract.WETH(&_UniswapRouterMock.CallOpts)
}

// AddLiquidity is a free data retrieval call binding the contract method 0xe8e33700.
//
// Solidity: function addLiquidity(address , address , uint256 , uint256 , uint256 , uint256 , address , uint256 ) pure returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_UniswapRouterMock *UniswapRouterMockCaller) AddLiquidity(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 *big.Int, arg5 *big.Int, arg6 common.Address, arg7 *big.Int) (struct {
	AmountA   *big.Int
	AmountB   *big.Int
	Liquidity *big.Int
}, error) {
	var out []interface{}
	err := _UniswapRouterMock.contract.Call(opts, &out, "addLiquidity", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)

	outstruct := new(struct {
		AmountA   *big.Int
		AmountB   *big.Int
		Liquidity *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AmountA = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.AmountB = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Liquidity = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// AddLiquidity is a free data retrieval call binding the contract method 0xe8e33700.
//
// Solidity: function addLiquidity(address , address , uint256 , uint256 , uint256 , uint256 , address , uint256 ) pure returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_UniswapRouterMock *UniswapRouterMockSession) AddLiquidity(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 *big.Int, arg5 *big.Int, arg6 common.Address, arg7 *big.Int) (struct {
	AmountA   *big.Int
	AmountB   *big.Int
	Liquidity *big.Int
}, error) {
	return _UniswapRouterMock.Contract.AddLiquidity(&_UniswapRouterMock.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

// AddLiquidity is a free data retrieval call binding the contract method 0xe8e33700.
//
// Solidity: function addLiquidity(address , address , uint256 , uint256 , uint256 , uint256 , address , uint256 ) pure returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_UniswapRouterMock *UniswapRouterMockCallerSession) AddLiquidity(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 *big.Int, arg5 *big.Int, arg6 common.Address, arg7 *big.Int) (struct {
	AmountA   *big.Int
	AmountB   *big.Int
	Liquidity *big.Int
}, error) {
	return _UniswapRouterMock.Contract.AddLiquidity(&_UniswapRouterMock.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() pure returns(address)
func (_UniswapRouterMock *UniswapRouterMockCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniswapRouterMock.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() pure returns(address)
func (_UniswapRouterMock *UniswapRouterMockSession) Factory() (common.Address, error) {
	return _UniswapRouterMock.Contract.Factory(&_UniswapRouterMock.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() pure returns(address)
func (_UniswapRouterMock *UniswapRouterMockCallerSession) Factory() (common.Address, error) {
	return _UniswapRouterMock.Contract.Factory(&_UniswapRouterMock.CallOpts)
}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 , uint256 , uint256 ) pure returns(uint256 amountIn)
func (_UniswapRouterMock *UniswapRouterMockCaller) GetAmountIn(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int, arg2 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _UniswapRouterMock.contract.Call(opts, &out, "getAmountIn", arg0, arg1, arg2)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 , uint256 , uint256 ) pure returns(uint256 amountIn)
func (_UniswapRouterMock *UniswapRouterMockSession) GetAmountIn(arg0 *big.Int, arg1 *big.Int, arg2 *big.Int) (*big.Int, error) {
	return _UniswapRouterMock.Contract.GetAmountIn(&_UniswapRouterMock.CallOpts, arg0, arg1, arg2)
}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 , uint256 , uint256 ) pure returns(uint256 amountIn)
func (_UniswapRouterMock *UniswapRouterMockCallerSession) GetAmountIn(arg0 *big.Int, arg1 *big.Int, arg2 *big.Int) (*big.Int, error) {
	return _UniswapRouterMock.Contract.GetAmountIn(&_UniswapRouterMock.CallOpts, arg0, arg1, arg2)
}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 , uint256 , uint256 ) pure returns(uint256 amountOut)
func (_UniswapRouterMock *UniswapRouterMockCaller) GetAmountOut(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int, arg2 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _UniswapRouterMock.contract.Call(opts, &out, "getAmountOut", arg0, arg1, arg2)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 , uint256 , uint256 ) pure returns(uint256 amountOut)
func (_UniswapRouterMock *UniswapRouterMockSession) GetAmountOut(arg0 *big.Int, arg1 *big.Int, arg2 *big.Int) (*big.Int, error) {
	return _UniswapRouterMock.Contract.GetAmountOut(&_UniswapRouterMock.CallOpts, arg0, arg1, arg2)
}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 , uint256 , uint256 ) pure returns(uint256 amountOut)
func (_UniswapRouterMock *UniswapRouterMockCallerSession) GetAmountOut(arg0 *big.Int, arg1 *big.Int, arg2 *big.Int) (*big.Int, error) {
	return _UniswapRouterMock.Contract.GetAmountOut(&_UniswapRouterMock.CallOpts, arg0, arg1, arg2)
}

// GetAmountsIn is a free data retrieval call binding the contract method 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 , address[] ) pure returns(uint256[])
func (_UniswapRouterMock *UniswapRouterMockCaller) GetAmountsIn(opts *bind.CallOpts, arg0 *big.Int, arg1 []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _UniswapRouterMock.contract.Call(opts, &out, "getAmountsIn", arg0, arg1)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAmountsIn is a free data retrieval call binding the contract method 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 , address[] ) pure returns(uint256[])
func (_UniswapRouterMock *UniswapRouterMockSession) GetAmountsIn(arg0 *big.Int, arg1 []common.Address) ([]*big.Int, error) {
	return _UniswapRouterMock.Contract.GetAmountsIn(&_UniswapRouterMock.CallOpts, arg0, arg1)
}

// GetAmountsIn is a free data retrieval call binding the contract method 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 , address[] ) pure returns(uint256[])
func (_UniswapRouterMock *UniswapRouterMockCallerSession) GetAmountsIn(arg0 *big.Int, arg1 []common.Address) ([]*big.Int, error) {
	return _UniswapRouterMock.Contract.GetAmountsIn(&_UniswapRouterMock.CallOpts, arg0, arg1)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[])
func (_UniswapRouterMock *UniswapRouterMockCaller) GetAmountsOut(opts *bind.CallOpts, amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _UniswapRouterMock.contract.Call(opts, &out, "getAmountsOut", amountIn, path)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[])
func (_UniswapRouterMock *UniswapRouterMockSession) GetAmountsOut(amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	return _UniswapRouterMock.Contract.GetAmountsOut(&_UniswapRouterMock.CallOpts, amountIn, path)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[])
func (_UniswapRouterMock *UniswapRouterMockCallerSession) GetAmountsOut(amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	return _UniswapRouterMock.Contract.GetAmountsOut(&_UniswapRouterMock.CallOpts, amountIn, path)
}

// GetRate is a free data retrieval call binding the contract method 0x97edd4fa.
//
// Solidity: function getRate(address[] path) view returns(uint256 rate)
func (_UniswapRouterMock *UniswapRouterMockCaller) GetRate(opts *bind.CallOpts, path []common.Address) (*big.Int, error) {
	var out []interface{}
	err := _UniswapRouterMock.contract.Call(opts, &out, "getRate", path)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRate is a free data retrieval call binding the contract method 0x97edd4fa.
//
// Solidity: function getRate(address[] path) view returns(uint256 rate)
func (_UniswapRouterMock *UniswapRouterMockSession) GetRate(path []common.Address) (*big.Int, error) {
	return _UniswapRouterMock.Contract.GetRate(&_UniswapRouterMock.CallOpts, path)
}

// GetRate is a free data retrieval call binding the contract method 0x97edd4fa.
//
// Solidity: function getRate(address[] path) view returns(uint256 rate)
func (_UniswapRouterMock *UniswapRouterMockCallerSession) GetRate(path []common.Address) (*big.Int, error) {
	return _UniswapRouterMock.Contract.GetRate(&_UniswapRouterMock.CallOpts, path)
}

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 , uint256 , uint256 ) pure returns(uint256 amountB)
func (_UniswapRouterMock *UniswapRouterMockCaller) Quote(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int, arg2 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _UniswapRouterMock.contract.Call(opts, &out, "quote", arg0, arg1, arg2)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 , uint256 , uint256 ) pure returns(uint256 amountB)
func (_UniswapRouterMock *UniswapRouterMockSession) Quote(arg0 *big.Int, arg1 *big.Int, arg2 *big.Int) (*big.Int, error) {
	return _UniswapRouterMock.Contract.Quote(&_UniswapRouterMock.CallOpts, arg0, arg1, arg2)
}

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 , uint256 , uint256 ) pure returns(uint256 amountB)
func (_UniswapRouterMock *UniswapRouterMockCallerSession) Quote(arg0 *big.Int, arg1 *big.Int, arg2 *big.Int) (*big.Int, error) {
	return _UniswapRouterMock.Contract.Quote(&_UniswapRouterMock.CallOpts, arg0, arg1, arg2)
}

// RemoveLiquidity is a free data retrieval call binding the contract method 0xbaa2abde.
//
// Solidity: function removeLiquidity(address , address , uint256 , uint256 , uint256 , address , uint256 ) pure returns(uint256 amountA, uint256 amountB)
func (_UniswapRouterMock *UniswapRouterMockCaller) RemoveLiquidity(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 *big.Int, arg5 common.Address, arg6 *big.Int) (struct {
	AmountA *big.Int
	AmountB *big.Int
}, error) {
	var out []interface{}
	err := _UniswapRouterMock.contract.Call(opts, &out, "removeLiquidity", arg0, arg1, arg2, arg3, arg4, arg5, arg6)

	outstruct := new(struct {
		AmountA *big.Int
		AmountB *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AmountA = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.AmountB = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// RemoveLiquidity is a free data retrieval call binding the contract method 0xbaa2abde.
//
// Solidity: function removeLiquidity(address , address , uint256 , uint256 , uint256 , address , uint256 ) pure returns(uint256 amountA, uint256 amountB)
func (_UniswapRouterMock *UniswapRouterMockSession) RemoveLiquidity(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 *big.Int, arg5 common.Address, arg6 *big.Int) (struct {
	AmountA *big.Int
	AmountB *big.Int
}, error) {
	return _UniswapRouterMock.Contract.RemoveLiquidity(&_UniswapRouterMock.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// RemoveLiquidity is a free data retrieval call binding the contract method 0xbaa2abde.
//
// Solidity: function removeLiquidity(address , address , uint256 , uint256 , uint256 , address , uint256 ) pure returns(uint256 amountA, uint256 amountB)
func (_UniswapRouterMock *UniswapRouterMockCallerSession) RemoveLiquidity(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 *big.Int, arg5 common.Address, arg6 *big.Int) (struct {
	AmountA *big.Int
	AmountB *big.Int
}, error) {
	return _UniswapRouterMock.Contract.RemoveLiquidity(&_UniswapRouterMock.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// RemoveLiquidityETH is a free data retrieval call binding the contract method 0x02751cec.
//
// Solidity: function removeLiquidityETH(address , uint256 , uint256 , uint256 , address , uint256 ) pure returns(uint256 amountToken, uint256 amountETH)
func (_UniswapRouterMock *UniswapRouterMockCaller) RemoveLiquidityETH(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int, arg2 *big.Int, arg3 *big.Int, arg4 common.Address, arg5 *big.Int) (struct {
	AmountToken *big.Int
	AmountETH   *big.Int
}, error) {
	var out []interface{}
	err := _UniswapRouterMock.contract.Call(opts, &out, "removeLiquidityETH", arg0, arg1, arg2, arg3, arg4, arg5)

	outstruct := new(struct {
		AmountToken *big.Int
		AmountETH   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AmountToken = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.AmountETH = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// RemoveLiquidityETH is a free data retrieval call binding the contract method 0x02751cec.
//
// Solidity: function removeLiquidityETH(address , uint256 , uint256 , uint256 , address , uint256 ) pure returns(uint256 amountToken, uint256 amountETH)
func (_UniswapRouterMock *UniswapRouterMockSession) RemoveLiquidityETH(arg0 common.Address, arg1 *big.Int, arg2 *big.Int, arg3 *big.Int, arg4 common.Address, arg5 *big.Int) (struct {
	AmountToken *big.Int
	AmountETH   *big.Int
}, error) {
	return _UniswapRouterMock.Contract.RemoveLiquidityETH(&_UniswapRouterMock.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// RemoveLiquidityETH is a free data retrieval call binding the contract method 0x02751cec.
//
// Solidity: function removeLiquidityETH(address , uint256 , uint256 , uint256 , address , uint256 ) pure returns(uint256 amountToken, uint256 amountETH)
func (_UniswapRouterMock *UniswapRouterMockCallerSession) RemoveLiquidityETH(arg0 common.Address, arg1 *big.Int, arg2 *big.Int, arg3 *big.Int, arg4 common.Address, arg5 *big.Int) (struct {
	AmountToken *big.Int
	AmountETH   *big.Int
}, error) {
	return _UniswapRouterMock.Contract.RemoveLiquidityETH(&_UniswapRouterMock.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// RemoveLiquidityETHSupportingFeeOnTransferTokens is a free data retrieval call binding the contract method 0xaf2979eb.
//
// Solidity: function removeLiquidityETHSupportingFeeOnTransferTokens(address , uint256 , uint256 , uint256 , address , uint256 ) pure returns(uint256 amountETH)
func (_UniswapRouterMock *UniswapRouterMockCaller) RemoveLiquidityETHSupportingFeeOnTransferTokens(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int, arg2 *big.Int, arg3 *big.Int, arg4 common.Address, arg5 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _UniswapRouterMock.contract.Call(opts, &out, "removeLiquidityETHSupportingFeeOnTransferTokens", arg0, arg1, arg2, arg3, arg4, arg5)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RemoveLiquidityETHSupportingFeeOnTransferTokens is a free data retrieval call binding the contract method 0xaf2979eb.
//
// Solidity: function removeLiquidityETHSupportingFeeOnTransferTokens(address , uint256 , uint256 , uint256 , address , uint256 ) pure returns(uint256 amountETH)
func (_UniswapRouterMock *UniswapRouterMockSession) RemoveLiquidityETHSupportingFeeOnTransferTokens(arg0 common.Address, arg1 *big.Int, arg2 *big.Int, arg3 *big.Int, arg4 common.Address, arg5 *big.Int) (*big.Int, error) {
	return _UniswapRouterMock.Contract.RemoveLiquidityETHSupportingFeeOnTransferTokens(&_UniswapRouterMock.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// RemoveLiquidityETHSupportingFeeOnTransferTokens is a free data retrieval call binding the contract method 0xaf2979eb.
//
// Solidity: function removeLiquidityETHSupportingFeeOnTransferTokens(address , uint256 , uint256 , uint256 , address , uint256 ) pure returns(uint256 amountETH)
func (_UniswapRouterMock *UniswapRouterMockCallerSession) RemoveLiquidityETHSupportingFeeOnTransferTokens(arg0 common.Address, arg1 *big.Int, arg2 *big.Int, arg3 *big.Int, arg4 common.Address, arg5 *big.Int) (*big.Int, error) {
	return _UniswapRouterMock.Contract.RemoveLiquidityETHSupportingFeeOnTransferTokens(&_UniswapRouterMock.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// RemoveLiquidityETHWithPermit is a free data retrieval call binding the contract method 0xded9382a.
//
// Solidity: function removeLiquidityETHWithPermit(address , uint256 , uint256 , uint256 , address , uint256 , bool , uint8 , bytes32 , bytes32 ) pure returns(uint256 amountToken, uint256 amountETH)
func (_UniswapRouterMock *UniswapRouterMockCaller) RemoveLiquidityETHWithPermit(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int, arg2 *big.Int, arg3 *big.Int, arg4 common.Address, arg5 *big.Int, arg6 bool, arg7 uint8, arg8 [32]byte, arg9 [32]byte) (struct {
	AmountToken *big.Int
	AmountETH   *big.Int
}, error) {
	var out []interface{}
	err := _UniswapRouterMock.contract.Call(opts, &out, "removeLiquidityETHWithPermit", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)

	outstruct := new(struct {
		AmountToken *big.Int
		AmountETH   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AmountToken = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.AmountETH = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// RemoveLiquidityETHWithPermit is a free data retrieval call binding the contract method 0xded9382a.
//
// Solidity: function removeLiquidityETHWithPermit(address , uint256 , uint256 , uint256 , address , uint256 , bool , uint8 , bytes32 , bytes32 ) pure returns(uint256 amountToken, uint256 amountETH)
func (_UniswapRouterMock *UniswapRouterMockSession) RemoveLiquidityETHWithPermit(arg0 common.Address, arg1 *big.Int, arg2 *big.Int, arg3 *big.Int, arg4 common.Address, arg5 *big.Int, arg6 bool, arg7 uint8, arg8 [32]byte, arg9 [32]byte) (struct {
	AmountToken *big.Int
	AmountETH   *big.Int
}, error) {
	return _UniswapRouterMock.Contract.RemoveLiquidityETHWithPermit(&_UniswapRouterMock.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
}

// RemoveLiquidityETHWithPermit is a free data retrieval call binding the contract method 0xded9382a.
//
// Solidity: function removeLiquidityETHWithPermit(address , uint256 , uint256 , uint256 , address , uint256 , bool , uint8 , bytes32 , bytes32 ) pure returns(uint256 amountToken, uint256 amountETH)
func (_UniswapRouterMock *UniswapRouterMockCallerSession) RemoveLiquidityETHWithPermit(arg0 common.Address, arg1 *big.Int, arg2 *big.Int, arg3 *big.Int, arg4 common.Address, arg5 *big.Int, arg6 bool, arg7 uint8, arg8 [32]byte, arg9 [32]byte) (struct {
	AmountToken *big.Int
	AmountETH   *big.Int
}, error) {
	return _UniswapRouterMock.Contract.RemoveLiquidityETHWithPermit(&_UniswapRouterMock.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
}

// RemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens is a free data retrieval call binding the contract method 0x5b0d5984.
//
// Solidity: function removeLiquidityETHWithPermitSupportingFeeOnTransferTokens(address , uint256 , uint256 , uint256 , address , uint256 , bool , uint8 , bytes32 , bytes32 ) pure returns(uint256 amountETH)
func (_UniswapRouterMock *UniswapRouterMockCaller) RemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int, arg2 *big.Int, arg3 *big.Int, arg4 common.Address, arg5 *big.Int, arg6 bool, arg7 uint8, arg8 [32]byte, arg9 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _UniswapRouterMock.contract.Call(opts, &out, "removeLiquidityETHWithPermitSupportingFeeOnTransferTokens", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens is a free data retrieval call binding the contract method 0x5b0d5984.
//
// Solidity: function removeLiquidityETHWithPermitSupportingFeeOnTransferTokens(address , uint256 , uint256 , uint256 , address , uint256 , bool , uint8 , bytes32 , bytes32 ) pure returns(uint256 amountETH)
func (_UniswapRouterMock *UniswapRouterMockSession) RemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens(arg0 common.Address, arg1 *big.Int, arg2 *big.Int, arg3 *big.Int, arg4 common.Address, arg5 *big.Int, arg6 bool, arg7 uint8, arg8 [32]byte, arg9 [32]byte) (*big.Int, error) {
	return _UniswapRouterMock.Contract.RemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens(&_UniswapRouterMock.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
}

// RemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens is a free data retrieval call binding the contract method 0x5b0d5984.
//
// Solidity: function removeLiquidityETHWithPermitSupportingFeeOnTransferTokens(address , uint256 , uint256 , uint256 , address , uint256 , bool , uint8 , bytes32 , bytes32 ) pure returns(uint256 amountETH)
func (_UniswapRouterMock *UniswapRouterMockCallerSession) RemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens(arg0 common.Address, arg1 *big.Int, arg2 *big.Int, arg3 *big.Int, arg4 common.Address, arg5 *big.Int, arg6 bool, arg7 uint8, arg8 [32]byte, arg9 [32]byte) (*big.Int, error) {
	return _UniswapRouterMock.Contract.RemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens(&_UniswapRouterMock.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
}

// RemoveLiquidityWithPermit is a free data retrieval call binding the contract method 0x2195995c.
//
// Solidity: function removeLiquidityWithPermit(address , address , uint256 , uint256 , uint256 , address , uint256 , bool , uint8 , bytes32 , bytes32 ) pure returns(uint256 amountA, uint256 amountB)
func (_UniswapRouterMock *UniswapRouterMockCaller) RemoveLiquidityWithPermit(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 *big.Int, arg5 common.Address, arg6 *big.Int, arg7 bool, arg8 uint8, arg9 [32]byte, arg10 [32]byte) (struct {
	AmountA *big.Int
	AmountB *big.Int
}, error) {
	var out []interface{}
	err := _UniswapRouterMock.contract.Call(opts, &out, "removeLiquidityWithPermit", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10)

	outstruct := new(struct {
		AmountA *big.Int
		AmountB *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AmountA = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.AmountB = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// RemoveLiquidityWithPermit is a free data retrieval call binding the contract method 0x2195995c.
//
// Solidity: function removeLiquidityWithPermit(address , address , uint256 , uint256 , uint256 , address , uint256 , bool , uint8 , bytes32 , bytes32 ) pure returns(uint256 amountA, uint256 amountB)
func (_UniswapRouterMock *UniswapRouterMockSession) RemoveLiquidityWithPermit(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 *big.Int, arg5 common.Address, arg6 *big.Int, arg7 bool, arg8 uint8, arg9 [32]byte, arg10 [32]byte) (struct {
	AmountA *big.Int
	AmountB *big.Int
}, error) {
	return _UniswapRouterMock.Contract.RemoveLiquidityWithPermit(&_UniswapRouterMock.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10)
}

// RemoveLiquidityWithPermit is a free data retrieval call binding the contract method 0x2195995c.
//
// Solidity: function removeLiquidityWithPermit(address , address , uint256 , uint256 , uint256 , address , uint256 , bool , uint8 , bytes32 , bytes32 ) pure returns(uint256 amountA, uint256 amountB)
func (_UniswapRouterMock *UniswapRouterMockCallerSession) RemoveLiquidityWithPermit(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 *big.Int, arg5 common.Address, arg6 *big.Int, arg7 bool, arg8 uint8, arg9 [32]byte, arg10 [32]byte) (struct {
	AmountA *big.Int
	AmountB *big.Int
}, error) {
	return _UniswapRouterMock.Contract.RemoveLiquidityWithPermit(&_UniswapRouterMock.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10)
}

// SwapExactTokensForETH is a free data retrieval call binding the contract method 0x18cbafe5.
//
// Solidity: function swapExactTokensForETH(uint256 , uint256 , address[] , address , uint256 ) pure returns(uint256[])
func (_UniswapRouterMock *UniswapRouterMockCaller) SwapExactTokensForETH(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int, arg2 []common.Address, arg3 common.Address, arg4 *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _UniswapRouterMock.contract.Call(opts, &out, "swapExactTokensForETH", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// SwapExactTokensForETH is a free data retrieval call binding the contract method 0x18cbafe5.
//
// Solidity: function swapExactTokensForETH(uint256 , uint256 , address[] , address , uint256 ) pure returns(uint256[])
func (_UniswapRouterMock *UniswapRouterMockSession) SwapExactTokensForETH(arg0 *big.Int, arg1 *big.Int, arg2 []common.Address, arg3 common.Address, arg4 *big.Int) ([]*big.Int, error) {
	return _UniswapRouterMock.Contract.SwapExactTokensForETH(&_UniswapRouterMock.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// SwapExactTokensForETH is a free data retrieval call binding the contract method 0x18cbafe5.
//
// Solidity: function swapExactTokensForETH(uint256 , uint256 , address[] , address , uint256 ) pure returns(uint256[])
func (_UniswapRouterMock *UniswapRouterMockCallerSession) SwapExactTokensForETH(arg0 *big.Int, arg1 *big.Int, arg2 []common.Address, arg3 common.Address, arg4 *big.Int) ([]*big.Int, error) {
	return _UniswapRouterMock.Contract.SwapExactTokensForETH(&_UniswapRouterMock.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// SwapExactTokensForETHSupportingFeeOnTransferTokens is a free data retrieval call binding the contract method 0x791ac947.
//
// Solidity: function swapExactTokensForETHSupportingFeeOnTransferTokens(uint256 , uint256 , address[] , address , uint256 ) pure returns()
func (_UniswapRouterMock *UniswapRouterMockCaller) SwapExactTokensForETHSupportingFeeOnTransferTokens(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int, arg2 []common.Address, arg3 common.Address, arg4 *big.Int) error {
	var out []interface{}
	err := _UniswapRouterMock.contract.Call(opts, &out, "swapExactTokensForETHSupportingFeeOnTransferTokens", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return err
	}

	return err

}

// SwapExactTokensForETHSupportingFeeOnTransferTokens is a free data retrieval call binding the contract method 0x791ac947.
//
// Solidity: function swapExactTokensForETHSupportingFeeOnTransferTokens(uint256 , uint256 , address[] , address , uint256 ) pure returns()
func (_UniswapRouterMock *UniswapRouterMockSession) SwapExactTokensForETHSupportingFeeOnTransferTokens(arg0 *big.Int, arg1 *big.Int, arg2 []common.Address, arg3 common.Address, arg4 *big.Int) error {
	return _UniswapRouterMock.Contract.SwapExactTokensForETHSupportingFeeOnTransferTokens(&_UniswapRouterMock.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// SwapExactTokensForETHSupportingFeeOnTransferTokens is a free data retrieval call binding the contract method 0x791ac947.
//
// Solidity: function swapExactTokensForETHSupportingFeeOnTransferTokens(uint256 , uint256 , address[] , address , uint256 ) pure returns()
func (_UniswapRouterMock *UniswapRouterMockCallerSession) SwapExactTokensForETHSupportingFeeOnTransferTokens(arg0 *big.Int, arg1 *big.Int, arg2 []common.Address, arg3 common.Address, arg4 *big.Int) error {
	return _UniswapRouterMock.Contract.SwapExactTokensForETHSupportingFeeOnTransferTokens(&_UniswapRouterMock.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a free data retrieval call binding the contract method 0x5c11d795.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 , uint256 , address[] , address , uint256 ) pure returns()
func (_UniswapRouterMock *UniswapRouterMockCaller) SwapExactTokensForTokensSupportingFeeOnTransferTokens(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int, arg2 []common.Address, arg3 common.Address, arg4 *big.Int) error {
	var out []interface{}
	err := _UniswapRouterMock.contract.Call(opts, &out, "swapExactTokensForTokensSupportingFeeOnTransferTokens", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return err
	}

	return err

}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a free data retrieval call binding the contract method 0x5c11d795.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 , uint256 , address[] , address , uint256 ) pure returns()
func (_UniswapRouterMock *UniswapRouterMockSession) SwapExactTokensForTokensSupportingFeeOnTransferTokens(arg0 *big.Int, arg1 *big.Int, arg2 []common.Address, arg3 common.Address, arg4 *big.Int) error {
	return _UniswapRouterMock.Contract.SwapExactTokensForTokensSupportingFeeOnTransferTokens(&_UniswapRouterMock.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a free data retrieval call binding the contract method 0x5c11d795.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 , uint256 , address[] , address , uint256 ) pure returns()
func (_UniswapRouterMock *UniswapRouterMockCallerSession) SwapExactTokensForTokensSupportingFeeOnTransferTokens(arg0 *big.Int, arg1 *big.Int, arg2 []common.Address, arg3 common.Address, arg4 *big.Int) error {
	return _UniswapRouterMock.Contract.SwapExactTokensForTokensSupportingFeeOnTransferTokens(&_UniswapRouterMock.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// SwapTokensForExactETH is a free data retrieval call binding the contract method 0x4a25d94a.
//
// Solidity: function swapTokensForExactETH(uint256 , uint256 , address[] , address , uint256 ) pure returns(uint256[])
func (_UniswapRouterMock *UniswapRouterMockCaller) SwapTokensForExactETH(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int, arg2 []common.Address, arg3 common.Address, arg4 *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _UniswapRouterMock.contract.Call(opts, &out, "swapTokensForExactETH", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// SwapTokensForExactETH is a free data retrieval call binding the contract method 0x4a25d94a.
//
// Solidity: function swapTokensForExactETH(uint256 , uint256 , address[] , address , uint256 ) pure returns(uint256[])
func (_UniswapRouterMock *UniswapRouterMockSession) SwapTokensForExactETH(arg0 *big.Int, arg1 *big.Int, arg2 []common.Address, arg3 common.Address, arg4 *big.Int) ([]*big.Int, error) {
	return _UniswapRouterMock.Contract.SwapTokensForExactETH(&_UniswapRouterMock.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// SwapTokensForExactETH is a free data retrieval call binding the contract method 0x4a25d94a.
//
// Solidity: function swapTokensForExactETH(uint256 , uint256 , address[] , address , uint256 ) pure returns(uint256[])
func (_UniswapRouterMock *UniswapRouterMockCallerSession) SwapTokensForExactETH(arg0 *big.Int, arg1 *big.Int, arg2 []common.Address, arg3 common.Address, arg4 *big.Int) ([]*big.Int, error) {
	return _UniswapRouterMock.Contract.SwapTokensForExactETH(&_UniswapRouterMock.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// AddLiquidityETH is a paid mutator transaction binding the contract method 0xf305d719.
//
// Solidity: function addLiquidityETH(address , uint256 , uint256 , uint256 , address , uint256 ) payable returns(uint256 amountToken, uint256 amountETH, uint256 liquidity)
func (_UniswapRouterMock *UniswapRouterMockTransactor) AddLiquidityETH(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int, arg2 *big.Int, arg3 *big.Int, arg4 common.Address, arg5 *big.Int) (*types.Transaction, error) {
	return _UniswapRouterMock.contract.Transact(opts, "addLiquidityETH", arg0, arg1, arg2, arg3, arg4, arg5)
}

// AddLiquidityETH is a paid mutator transaction binding the contract method 0xf305d719.
//
// Solidity: function addLiquidityETH(address , uint256 , uint256 , uint256 , address , uint256 ) payable returns(uint256 amountToken, uint256 amountETH, uint256 liquidity)
func (_UniswapRouterMock *UniswapRouterMockSession) AddLiquidityETH(arg0 common.Address, arg1 *big.Int, arg2 *big.Int, arg3 *big.Int, arg4 common.Address, arg5 *big.Int) (*types.Transaction, error) {
	return _UniswapRouterMock.Contract.AddLiquidityETH(&_UniswapRouterMock.TransactOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// AddLiquidityETH is a paid mutator transaction binding the contract method 0xf305d719.
//
// Solidity: function addLiquidityETH(address , uint256 , uint256 , uint256 , address , uint256 ) payable returns(uint256 amountToken, uint256 amountETH, uint256 liquidity)
func (_UniswapRouterMock *UniswapRouterMockTransactorSession) AddLiquidityETH(arg0 common.Address, arg1 *big.Int, arg2 *big.Int, arg3 *big.Int, arg4 common.Address, arg5 *big.Int) (*types.Transaction, error) {
	return _UniswapRouterMock.Contract.AddLiquidityETH(&_UniswapRouterMock.TransactOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// SetRate is a paid mutator transaction binding the contract method 0x5911fb9a.
//
// Solidity: function setRate(address tokenFrom, address tokenTo, uint256 rate_RAY) returns()
func (_UniswapRouterMock *UniswapRouterMockTransactor) SetRate(opts *bind.TransactOpts, tokenFrom common.Address, tokenTo common.Address, rate_RAY *big.Int) (*types.Transaction, error) {
	return _UniswapRouterMock.contract.Transact(opts, "setRate", tokenFrom, tokenTo, rate_RAY)
}

// SetRate is a paid mutator transaction binding the contract method 0x5911fb9a.
//
// Solidity: function setRate(address tokenFrom, address tokenTo, uint256 rate_RAY) returns()
func (_UniswapRouterMock *UniswapRouterMockSession) SetRate(tokenFrom common.Address, tokenTo common.Address, rate_RAY *big.Int) (*types.Transaction, error) {
	return _UniswapRouterMock.Contract.SetRate(&_UniswapRouterMock.TransactOpts, tokenFrom, tokenTo, rate_RAY)
}

// SetRate is a paid mutator transaction binding the contract method 0x5911fb9a.
//
// Solidity: function setRate(address tokenFrom, address tokenTo, uint256 rate_RAY) returns()
func (_UniswapRouterMock *UniswapRouterMockTransactorSession) SetRate(tokenFrom common.Address, tokenTo common.Address, rate_RAY *big.Int) (*types.Transaction, error) {
	return _UniswapRouterMock.Contract.SetRate(&_UniswapRouterMock.TransactOpts, tokenFrom, tokenTo, rate_RAY)
}

// SwapETHForExactTokens is a paid mutator transaction binding the contract method 0xfb3bdb41.
//
// Solidity: function swapETHForExactTokens(uint256 , address[] , address , uint256 ) payable returns(uint256[])
func (_UniswapRouterMock *UniswapRouterMockTransactor) SwapETHForExactTokens(opts *bind.TransactOpts, arg0 *big.Int, arg1 []common.Address, arg2 common.Address, arg3 *big.Int) (*types.Transaction, error) {
	return _UniswapRouterMock.contract.Transact(opts, "swapETHForExactTokens", arg0, arg1, arg2, arg3)
}

// SwapETHForExactTokens is a paid mutator transaction binding the contract method 0xfb3bdb41.
//
// Solidity: function swapETHForExactTokens(uint256 , address[] , address , uint256 ) payable returns(uint256[])
func (_UniswapRouterMock *UniswapRouterMockSession) SwapETHForExactTokens(arg0 *big.Int, arg1 []common.Address, arg2 common.Address, arg3 *big.Int) (*types.Transaction, error) {
	return _UniswapRouterMock.Contract.SwapETHForExactTokens(&_UniswapRouterMock.TransactOpts, arg0, arg1, arg2, arg3)
}

// SwapETHForExactTokens is a paid mutator transaction binding the contract method 0xfb3bdb41.
//
// Solidity: function swapETHForExactTokens(uint256 , address[] , address , uint256 ) payable returns(uint256[])
func (_UniswapRouterMock *UniswapRouterMockTransactorSession) SwapETHForExactTokens(arg0 *big.Int, arg1 []common.Address, arg2 common.Address, arg3 *big.Int) (*types.Transaction, error) {
	return _UniswapRouterMock.Contract.SwapETHForExactTokens(&_UniswapRouterMock.TransactOpts, arg0, arg1, arg2, arg3)
}

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0x7ff36ab5.
//
// Solidity: function swapExactETHForTokens(uint256 , address[] , address , uint256 ) payable returns(uint256[])
func (_UniswapRouterMock *UniswapRouterMockTransactor) SwapExactETHForTokens(opts *bind.TransactOpts, arg0 *big.Int, arg1 []common.Address, arg2 common.Address, arg3 *big.Int) (*types.Transaction, error) {
	return _UniswapRouterMock.contract.Transact(opts, "swapExactETHForTokens", arg0, arg1, arg2, arg3)
}

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0x7ff36ab5.
//
// Solidity: function swapExactETHForTokens(uint256 , address[] , address , uint256 ) payable returns(uint256[])
func (_UniswapRouterMock *UniswapRouterMockSession) SwapExactETHForTokens(arg0 *big.Int, arg1 []common.Address, arg2 common.Address, arg3 *big.Int) (*types.Transaction, error) {
	return _UniswapRouterMock.Contract.SwapExactETHForTokens(&_UniswapRouterMock.TransactOpts, arg0, arg1, arg2, arg3)
}

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0x7ff36ab5.
//
// Solidity: function swapExactETHForTokens(uint256 , address[] , address , uint256 ) payable returns(uint256[])
func (_UniswapRouterMock *UniswapRouterMockTransactorSession) SwapExactETHForTokens(arg0 *big.Int, arg1 []common.Address, arg2 common.Address, arg3 *big.Int) (*types.Transaction, error) {
	return _UniswapRouterMock.Contract.SwapExactETHForTokens(&_UniswapRouterMock.TransactOpts, arg0, arg1, arg2, arg3)
}

// SwapExactETHForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xb6f9de95.
//
// Solidity: function swapExactETHForTokensSupportingFeeOnTransferTokens(uint256 , address[] , address , uint256 ) payable returns()
func (_UniswapRouterMock *UniswapRouterMockTransactor) SwapExactETHForTokensSupportingFeeOnTransferTokens(opts *bind.TransactOpts, arg0 *big.Int, arg1 []common.Address, arg2 common.Address, arg3 *big.Int) (*types.Transaction, error) {
	return _UniswapRouterMock.contract.Transact(opts, "swapExactETHForTokensSupportingFeeOnTransferTokens", arg0, arg1, arg2, arg3)
}

// SwapExactETHForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xb6f9de95.
//
// Solidity: function swapExactETHForTokensSupportingFeeOnTransferTokens(uint256 , address[] , address , uint256 ) payable returns()
func (_UniswapRouterMock *UniswapRouterMockSession) SwapExactETHForTokensSupportingFeeOnTransferTokens(arg0 *big.Int, arg1 []common.Address, arg2 common.Address, arg3 *big.Int) (*types.Transaction, error) {
	return _UniswapRouterMock.Contract.SwapExactETHForTokensSupportingFeeOnTransferTokens(&_UniswapRouterMock.TransactOpts, arg0, arg1, arg2, arg3)
}

// SwapExactETHForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xb6f9de95.
//
// Solidity: function swapExactETHForTokensSupportingFeeOnTransferTokens(uint256 , address[] , address , uint256 ) payable returns()
func (_UniswapRouterMock *UniswapRouterMockTransactorSession) SwapExactETHForTokensSupportingFeeOnTransferTokens(arg0 *big.Int, arg1 []common.Address, arg2 common.Address, arg3 *big.Int) (*types.Transaction, error) {
	return _UniswapRouterMock.Contract.SwapExactETHForTokensSupportingFeeOnTransferTokens(&_UniswapRouterMock.TransactOpts, arg0, arg1, arg2, arg3)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[])
func (_UniswapRouterMock *UniswapRouterMockTransactor) SwapExactTokensForTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _UniswapRouterMock.contract.Transact(opts, "swapExactTokensForTokens", amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[])
func (_UniswapRouterMock *UniswapRouterMockSession) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _UniswapRouterMock.Contract.SwapExactTokensForTokens(&_UniswapRouterMock.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[])
func (_UniswapRouterMock *UniswapRouterMockTransactorSession) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _UniswapRouterMock.Contract.SwapExactTokensForTokens(&_UniswapRouterMock.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[])
func (_UniswapRouterMock *UniswapRouterMockTransactor) SwapTokensForExactTokens(opts *bind.TransactOpts, amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _UniswapRouterMock.contract.Transact(opts, "swapTokensForExactTokens", amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[])
func (_UniswapRouterMock *UniswapRouterMockSession) SwapTokensForExactTokens(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _UniswapRouterMock.Contract.SwapTokensForExactTokens(&_UniswapRouterMock.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[])
func (_UniswapRouterMock *UniswapRouterMockTransactorSession) SwapTokensForExactTokens(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _UniswapRouterMock.Contract.SwapTokensForExactTokens(&_UniswapRouterMock.TransactOpts, amountOut, amountInMax, path, to, deadline)
}
