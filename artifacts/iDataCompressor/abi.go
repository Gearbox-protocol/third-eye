// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iDataCompressor

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

// DataTypesContractAdapter is an auto generated low-level Go binding around an user-defined struct.
type DataTypesContractAdapter struct {
	AllowedContract common.Address
	Adapter         common.Address
}

// DataTypesCreditAccountData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesCreditAccountData struct {
	Addr                       common.Address
	Borrower                   common.Address
	InUse                      bool
	CreditManager              common.Address
	UnderlyingToken            common.Address
	BorrowedAmountPlusInterest *big.Int
	TotalValue                 *big.Int
	HealthFactor               *big.Int
	BorrowRate                 *big.Int
	Balances                   []DataTypesTokenBalance
}

// DataTypesCreditAccountDataExtended is an auto generated low-level Go binding around an user-defined struct.
type DataTypesCreditAccountDataExtended struct {
	Addr                       common.Address
	Borrower                   common.Address
	InUse                      bool
	CreditManager              common.Address
	UnderlyingToken            common.Address
	BorrowedAmountPlusInterest *big.Int
	TotalValue                 *big.Int
	HealthFactor               *big.Int
	BorrowRate                 *big.Int
	Balances                   []DataTypesTokenBalance
	RepayAmount                *big.Int
	LiquidationAmount          *big.Int
	CanBeClosed                bool
	BorrowedAmount             *big.Int
	CumulativeIndexAtOpen      *big.Int
	Since                      *big.Int
}

// DataTypesCreditManagerData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesCreditManagerData struct {
	Addr               common.Address
	HasAccount         bool
	UnderlyingToken    common.Address
	IsWETH             bool
	CanBorrow          bool
	BorrowRate         *big.Int
	MinAmount          *big.Int
	MaxAmount          *big.Int
	MaxLeverageFactor  *big.Int
	AvailableLiquidity *big.Int
	AllowedTokens      []common.Address
	Adapters           []DataTypesContractAdapter
}

// DataTypesPoolData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesPoolData struct {
	Addr                   common.Address
	IsWETH                 bool
	UnderlyingToken        common.Address
	DieselToken            common.Address
	LinearCumulativeIndex  *big.Int
	AvailableLiquidity     *big.Int
	ExpectedLiquidity      *big.Int
	ExpectedLiquidityLimit *big.Int
	TotalBorrowed          *big.Int
	DepositAPYRAY          *big.Int
	BorrowAPYRAY           *big.Int
	DieselRateRAY          *big.Int
	WithdrawFee            *big.Int
	CumulativeIndexRAY     *big.Int
	TimestampLU            *big.Int
}

// DataTypesTokenBalance is an auto generated low-level Go binding around an user-defined struct.
type DataTypesTokenBalance struct {
	Token     common.Address
	Balance   *big.Int
	IsAllowed bool
}

// DataTypesTokenInfo is an auto generated low-level Go binding around an user-defined struct.
type DataTypesTokenInfo struct {
	Addr     common.Address
	Symbol   string
	Decimals uint8
}

// IDataCompressorMetaData contains all meta data concerning the IDataCompressor contract.
var IDataCompressorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_creditManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowedAmount\",\"type\":\"uint256\"}],\"name\":\"calcExpectedAtOpenHf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"}],\"name\":\"calcExpectedHf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_creditManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"}],\"name\":\"getCreditAccountData\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"inUse\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"creditManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlyingToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"borrowedAmountPlusInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"healthFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowRate\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isAllowed\",\"type\":\"bool\"}],\"internalType\":\"structDataTypes.TokenBalance[]\",\"name\":\"balances\",\"type\":\"tuple[]\"}],\"internalType\":\"structDataTypes.CreditAccountData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"}],\"name\":\"getCreditAccountDataExtended\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"inUse\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"creditManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlyingToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"borrowedAmountPlusInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"healthFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowRate\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isAllowed\",\"type\":\"bool\"}],\"internalType\":\"structDataTypes.TokenBalance[]\",\"name\":\"balances\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"canBeClosed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"borrowedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cumulativeIndexAtOpen\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"since\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.CreditAccountDataExtended\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"}],\"name\":\"getCreditAccountList\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"inUse\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"creditManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlyingToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"borrowedAmountPlusInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"healthFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowRate\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isAllowed\",\"type\":\"bool\"}],\"internalType\":\"structDataTypes.TokenBalance[]\",\"name\":\"balances\",\"type\":\"tuple[]\"}],\"internalType\":\"structDataTypes.CreditAccountData[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_creditManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"}],\"name\":\"getCreditManagerData\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"hasAccount\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"underlyingToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isWETH\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"canBorrow\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"borrowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxLeverageFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedTokens\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"allowedContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"adapter\",\"type\":\"address\"}],\"internalType\":\"structDataTypes.ContractAdapter[]\",\"name\":\"adapters\",\"type\":\"tuple[]\"}],\"internalType\":\"structDataTypes.CreditManagerData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"}],\"name\":\"getCreditManagersList\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"hasAccount\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"underlyingToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isWETH\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"canBorrow\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"borrowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxLeverageFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedTokens\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"allowedContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"adapter\",\"type\":\"address\"}],\"internalType\":\"structDataTypes.ContractAdapter[]\",\"name\":\"adapters\",\"type\":\"tuple[]\"}],\"internalType\":\"structDataTypes.CreditManagerData[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_pool\",\"type\":\"address\"}],\"name\":\"getPoolData\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isWETH\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"underlyingToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dieselToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"linearCumulativeIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedLiquidityLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalBorrowed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositAPY_RAY\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowAPY_RAY\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dieselRate_RAY\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cumulativeIndex_RAY\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestampLU\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.PoolData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolsList\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isWETH\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"underlyingToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dieselToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"linearCumulativeIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedLiquidityLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalBorrowed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositAPY_RAY\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowAPY_RAY\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dieselRate_RAY\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cumulativeIndex_RAY\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestampLU\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.PoolData[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addr\",\"type\":\"address[]\"}],\"name\":\"getTokenData\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"internalType\":\"structDataTypes.TokenInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"}],\"name\":\"hasOpenedCreditAccount\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IDataCompressorABI is the input ABI used to generate the binding from.
// Deprecated: Use IDataCompressorMetaData.ABI instead.
var IDataCompressorABI = IDataCompressorMetaData.ABI

// IDataCompressor is an auto generated Go binding around an Ethereum contract.
type IDataCompressor struct {
	IDataCompressorCaller     // Read-only binding to the contract
	IDataCompressorTransactor // Write-only binding to the contract
	IDataCompressorFilterer   // Log filterer for contract events
}

// IDataCompressorCaller is an auto generated read-only Go binding around an Ethereum contract.
type IDataCompressorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDataCompressorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IDataCompressorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDataCompressorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IDataCompressorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDataCompressorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IDataCompressorSession struct {
	Contract     *IDataCompressor  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IDataCompressorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IDataCompressorCallerSession struct {
	Contract *IDataCompressorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IDataCompressorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IDataCompressorTransactorSession struct {
	Contract     *IDataCompressorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IDataCompressorRaw is an auto generated low-level Go binding around an Ethereum contract.
type IDataCompressorRaw struct {
	Contract *IDataCompressor // Generic contract binding to access the raw methods on
}

// IDataCompressorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IDataCompressorCallerRaw struct {
	Contract *IDataCompressorCaller // Generic read-only contract binding to access the raw methods on
}

// IDataCompressorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IDataCompressorTransactorRaw struct {
	Contract *IDataCompressorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIDataCompressor creates a new instance of IDataCompressor, bound to a specific deployed contract.
func NewIDataCompressor(address common.Address, backend bind.ContractBackend) (*IDataCompressor, error) {
	contract, err := bindIDataCompressor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IDataCompressor{IDataCompressorCaller: IDataCompressorCaller{contract: contract}, IDataCompressorTransactor: IDataCompressorTransactor{contract: contract}, IDataCompressorFilterer: IDataCompressorFilterer{contract: contract}}, nil
}

// NewIDataCompressorCaller creates a new read-only instance of IDataCompressor, bound to a specific deployed contract.
func NewIDataCompressorCaller(address common.Address, caller bind.ContractCaller) (*IDataCompressorCaller, error) {
	contract, err := bindIDataCompressor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IDataCompressorCaller{contract: contract}, nil
}

// NewIDataCompressorTransactor creates a new write-only instance of IDataCompressor, bound to a specific deployed contract.
func NewIDataCompressorTransactor(address common.Address, transactor bind.ContractTransactor) (*IDataCompressorTransactor, error) {
	contract, err := bindIDataCompressor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IDataCompressorTransactor{contract: contract}, nil
}

// NewIDataCompressorFilterer creates a new log filterer instance of IDataCompressor, bound to a specific deployed contract.
func NewIDataCompressorFilterer(address common.Address, filterer bind.ContractFilterer) (*IDataCompressorFilterer, error) {
	contract, err := bindIDataCompressor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IDataCompressorFilterer{contract: contract}, nil
}

// bindIDataCompressor binds a generic wrapper to an already deployed contract.
func bindIDataCompressor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IDataCompressorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDataCompressor *IDataCompressorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDataCompressor.Contract.IDataCompressorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDataCompressor *IDataCompressorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDataCompressor.Contract.IDataCompressorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDataCompressor *IDataCompressorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDataCompressor.Contract.IDataCompressorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDataCompressor *IDataCompressorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDataCompressor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDataCompressor *IDataCompressorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDataCompressor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDataCompressor *IDataCompressorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDataCompressor.Contract.contract.Transact(opts, method, params...)
}

// CalcExpectedAtOpenHf is a free data retrieval call binding the contract method 0x39595cf8.
//
// Solidity: function calcExpectedAtOpenHf(address _creditManager, address token, uint256 amount, uint256 borrowedAmount) view returns(uint256)
func (_IDataCompressor *IDataCompressorCaller) CalcExpectedAtOpenHf(opts *bind.CallOpts, _creditManager common.Address, token common.Address, amount *big.Int, borrowedAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IDataCompressor.contract.Call(opts, &out, "calcExpectedAtOpenHf", _creditManager, token, amount, borrowedAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcExpectedAtOpenHf is a free data retrieval call binding the contract method 0x39595cf8.
//
// Solidity: function calcExpectedAtOpenHf(address _creditManager, address token, uint256 amount, uint256 borrowedAmount) view returns(uint256)
func (_IDataCompressor *IDataCompressorSession) CalcExpectedAtOpenHf(_creditManager common.Address, token common.Address, amount *big.Int, borrowedAmount *big.Int) (*big.Int, error) {
	return _IDataCompressor.Contract.CalcExpectedAtOpenHf(&_IDataCompressor.CallOpts, _creditManager, token, amount, borrowedAmount)
}

// CalcExpectedAtOpenHf is a free data retrieval call binding the contract method 0x39595cf8.
//
// Solidity: function calcExpectedAtOpenHf(address _creditManager, address token, uint256 amount, uint256 borrowedAmount) view returns(uint256)
func (_IDataCompressor *IDataCompressorCallerSession) CalcExpectedAtOpenHf(_creditManager common.Address, token common.Address, amount *big.Int, borrowedAmount *big.Int) (*big.Int, error) {
	return _IDataCompressor.Contract.CalcExpectedAtOpenHf(&_IDataCompressor.CallOpts, _creditManager, token, amount, borrowedAmount)
}

// CalcExpectedHf is a free data retrieval call binding the contract method 0xba3b7345.
//
// Solidity: function calcExpectedHf(address creditManager, address borrower, uint256[] balances) view returns(uint256)
func (_IDataCompressor *IDataCompressorCaller) CalcExpectedHf(opts *bind.CallOpts, creditManager common.Address, borrower common.Address, balances []*big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IDataCompressor.contract.Call(opts, &out, "calcExpectedHf", creditManager, borrower, balances)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcExpectedHf is a free data retrieval call binding the contract method 0xba3b7345.
//
// Solidity: function calcExpectedHf(address creditManager, address borrower, uint256[] balances) view returns(uint256)
func (_IDataCompressor *IDataCompressorSession) CalcExpectedHf(creditManager common.Address, borrower common.Address, balances []*big.Int) (*big.Int, error) {
	return _IDataCompressor.Contract.CalcExpectedHf(&_IDataCompressor.CallOpts, creditManager, borrower, balances)
}

// CalcExpectedHf is a free data retrieval call binding the contract method 0xba3b7345.
//
// Solidity: function calcExpectedHf(address creditManager, address borrower, uint256[] balances) view returns(uint256)
func (_IDataCompressor *IDataCompressorCallerSession) CalcExpectedHf(creditManager common.Address, borrower common.Address, balances []*big.Int) (*big.Int, error) {
	return _IDataCompressor.Contract.CalcExpectedHf(&_IDataCompressor.CallOpts, creditManager, borrower, balances)
}

// GetCreditAccountData is a free data retrieval call binding the contract method 0x0dbd616d.
//
// Solidity: function getCreditAccountData(address _creditManager, address borrower) view returns((address,address,bool,address,address,uint256,uint256,uint256,uint256,(address,uint256,bool)[]))
func (_IDataCompressor *IDataCompressorCaller) GetCreditAccountData(opts *bind.CallOpts, _creditManager common.Address, borrower common.Address) (DataTypesCreditAccountData, error) {
	var out []interface{}
	err := _IDataCompressor.contract.Call(opts, &out, "getCreditAccountData", _creditManager, borrower)

	if err != nil {
		return *new(DataTypesCreditAccountData), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesCreditAccountData)).(*DataTypesCreditAccountData)

	return out0, err

}

// GetCreditAccountData is a free data retrieval call binding the contract method 0x0dbd616d.
//
// Solidity: function getCreditAccountData(address _creditManager, address borrower) view returns((address,address,bool,address,address,uint256,uint256,uint256,uint256,(address,uint256,bool)[]))
func (_IDataCompressor *IDataCompressorSession) GetCreditAccountData(_creditManager common.Address, borrower common.Address) (DataTypesCreditAccountData, error) {
	return _IDataCompressor.Contract.GetCreditAccountData(&_IDataCompressor.CallOpts, _creditManager, borrower)
}

// GetCreditAccountData is a free data retrieval call binding the contract method 0x0dbd616d.
//
// Solidity: function getCreditAccountData(address _creditManager, address borrower) view returns((address,address,bool,address,address,uint256,uint256,uint256,uint256,(address,uint256,bool)[]))
func (_IDataCompressor *IDataCompressorCallerSession) GetCreditAccountData(_creditManager common.Address, borrower common.Address) (DataTypesCreditAccountData, error) {
	return _IDataCompressor.Contract.GetCreditAccountData(&_IDataCompressor.CallOpts, _creditManager, borrower)
}

// GetCreditAccountDataExtended is a free data retrieval call binding the contract method 0x191482d4.
//
// Solidity: function getCreditAccountDataExtended(address creditManager, address borrower) view returns((address,address,bool,address,address,uint256,uint256,uint256,uint256,(address,uint256,bool)[],uint256,uint256,bool,uint256,uint256,uint256))
func (_IDataCompressor *IDataCompressorCaller) GetCreditAccountDataExtended(opts *bind.CallOpts, creditManager common.Address, borrower common.Address) (DataTypesCreditAccountDataExtended, error) {
	var out []interface{}
	err := _IDataCompressor.contract.Call(opts, &out, "getCreditAccountDataExtended", creditManager, borrower)

	if err != nil {
		return *new(DataTypesCreditAccountDataExtended), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesCreditAccountDataExtended)).(*DataTypesCreditAccountDataExtended)

	return out0, err

}

// GetCreditAccountDataExtended is a free data retrieval call binding the contract method 0x191482d4.
//
// Solidity: function getCreditAccountDataExtended(address creditManager, address borrower) view returns((address,address,bool,address,address,uint256,uint256,uint256,uint256,(address,uint256,bool)[],uint256,uint256,bool,uint256,uint256,uint256))
func (_IDataCompressor *IDataCompressorSession) GetCreditAccountDataExtended(creditManager common.Address, borrower common.Address) (DataTypesCreditAccountDataExtended, error) {
	return _IDataCompressor.Contract.GetCreditAccountDataExtended(&_IDataCompressor.CallOpts, creditManager, borrower)
}

// GetCreditAccountDataExtended is a free data retrieval call binding the contract method 0x191482d4.
//
// Solidity: function getCreditAccountDataExtended(address creditManager, address borrower) view returns((address,address,bool,address,address,uint256,uint256,uint256,uint256,(address,uint256,bool)[],uint256,uint256,bool,uint256,uint256,uint256))
func (_IDataCompressor *IDataCompressorCallerSession) GetCreditAccountDataExtended(creditManager common.Address, borrower common.Address) (DataTypesCreditAccountDataExtended, error) {
	return _IDataCompressor.Contract.GetCreditAccountDataExtended(&_IDataCompressor.CallOpts, creditManager, borrower)
}

// GetCreditAccountList is a free data retrieval call binding the contract method 0xa80deda3.
//
// Solidity: function getCreditAccountList(address borrower) view returns((address,address,bool,address,address,uint256,uint256,uint256,uint256,(address,uint256,bool)[])[])
func (_IDataCompressor *IDataCompressorCaller) GetCreditAccountList(opts *bind.CallOpts, borrower common.Address) ([]DataTypesCreditAccountData, error) {
	var out []interface{}
	err := _IDataCompressor.contract.Call(opts, &out, "getCreditAccountList", borrower)

	if err != nil {
		return *new([]DataTypesCreditAccountData), err
	}

	out0 := *abi.ConvertType(out[0], new([]DataTypesCreditAccountData)).(*[]DataTypesCreditAccountData)

	return out0, err

}

// GetCreditAccountList is a free data retrieval call binding the contract method 0xa80deda3.
//
// Solidity: function getCreditAccountList(address borrower) view returns((address,address,bool,address,address,uint256,uint256,uint256,uint256,(address,uint256,bool)[])[])
func (_IDataCompressor *IDataCompressorSession) GetCreditAccountList(borrower common.Address) ([]DataTypesCreditAccountData, error) {
	return _IDataCompressor.Contract.GetCreditAccountList(&_IDataCompressor.CallOpts, borrower)
}

// GetCreditAccountList is a free data retrieval call binding the contract method 0xa80deda3.
//
// Solidity: function getCreditAccountList(address borrower) view returns((address,address,bool,address,address,uint256,uint256,uint256,uint256,(address,uint256,bool)[])[])
func (_IDataCompressor *IDataCompressorCallerSession) GetCreditAccountList(borrower common.Address) ([]DataTypesCreditAccountData, error) {
	return _IDataCompressor.Contract.GetCreditAccountList(&_IDataCompressor.CallOpts, borrower)
}

// GetCreditManagerData is a free data retrieval call binding the contract method 0xb10b074e.
//
// Solidity: function getCreditManagerData(address _creditManager, address borrower) view returns((address,bool,address,bool,bool,uint256,uint256,uint256,uint256,uint256,address[],(address,address)[]))
func (_IDataCompressor *IDataCompressorCaller) GetCreditManagerData(opts *bind.CallOpts, _creditManager common.Address, borrower common.Address) (DataTypesCreditManagerData, error) {
	var out []interface{}
	err := _IDataCompressor.contract.Call(opts, &out, "getCreditManagerData", _creditManager, borrower)

	if err != nil {
		return *new(DataTypesCreditManagerData), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesCreditManagerData)).(*DataTypesCreditManagerData)

	return out0, err

}

// GetCreditManagerData is a free data retrieval call binding the contract method 0xb10b074e.
//
// Solidity: function getCreditManagerData(address _creditManager, address borrower) view returns((address,bool,address,bool,bool,uint256,uint256,uint256,uint256,uint256,address[],(address,address)[]))
func (_IDataCompressor *IDataCompressorSession) GetCreditManagerData(_creditManager common.Address, borrower common.Address) (DataTypesCreditManagerData, error) {
	return _IDataCompressor.Contract.GetCreditManagerData(&_IDataCompressor.CallOpts, _creditManager, borrower)
}

// GetCreditManagerData is a free data retrieval call binding the contract method 0xb10b074e.
//
// Solidity: function getCreditManagerData(address _creditManager, address borrower) view returns((address,bool,address,bool,bool,uint256,uint256,uint256,uint256,uint256,address[],(address,address)[]))
func (_IDataCompressor *IDataCompressorCallerSession) GetCreditManagerData(_creditManager common.Address, borrower common.Address) (DataTypesCreditManagerData, error) {
	return _IDataCompressor.Contract.GetCreditManagerData(&_IDataCompressor.CallOpts, _creditManager, borrower)
}

// GetCreditManagersList is a free data retrieval call binding the contract method 0xb8169039.
//
// Solidity: function getCreditManagersList(address borrower) view returns((address,bool,address,bool,bool,uint256,uint256,uint256,uint256,uint256,address[],(address,address)[])[])
func (_IDataCompressor *IDataCompressorCaller) GetCreditManagersList(opts *bind.CallOpts, borrower common.Address) ([]DataTypesCreditManagerData, error) {
	var out []interface{}
	err := _IDataCompressor.contract.Call(opts, &out, "getCreditManagersList", borrower)

	if err != nil {
		return *new([]DataTypesCreditManagerData), err
	}

	out0 := *abi.ConvertType(out[0], new([]DataTypesCreditManagerData)).(*[]DataTypesCreditManagerData)

	return out0, err

}

// GetCreditManagersList is a free data retrieval call binding the contract method 0xb8169039.
//
// Solidity: function getCreditManagersList(address borrower) view returns((address,bool,address,bool,bool,uint256,uint256,uint256,uint256,uint256,address[],(address,address)[])[])
func (_IDataCompressor *IDataCompressorSession) GetCreditManagersList(borrower common.Address) ([]DataTypesCreditManagerData, error) {
	return _IDataCompressor.Contract.GetCreditManagersList(&_IDataCompressor.CallOpts, borrower)
}

// GetCreditManagersList is a free data retrieval call binding the contract method 0xb8169039.
//
// Solidity: function getCreditManagersList(address borrower) view returns((address,bool,address,bool,bool,uint256,uint256,uint256,uint256,uint256,address[],(address,address)[])[])
func (_IDataCompressor *IDataCompressorCallerSession) GetCreditManagersList(borrower common.Address) ([]DataTypesCreditManagerData, error) {
	return _IDataCompressor.Contract.GetCreditManagersList(&_IDataCompressor.CallOpts, borrower)
}

// GetPoolData is a free data retrieval call binding the contract method 0x13d21cdf.
//
// Solidity: function getPoolData(address _pool) view returns((address,bool,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_IDataCompressor *IDataCompressorCaller) GetPoolData(opts *bind.CallOpts, _pool common.Address) (DataTypesPoolData, error) {
	var out []interface{}
	err := _IDataCompressor.contract.Call(opts, &out, "getPoolData", _pool)

	if err != nil {
		return *new(DataTypesPoolData), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesPoolData)).(*DataTypesPoolData)

	return out0, err

}

// GetPoolData is a free data retrieval call binding the contract method 0x13d21cdf.
//
// Solidity: function getPoolData(address _pool) view returns((address,bool,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_IDataCompressor *IDataCompressorSession) GetPoolData(_pool common.Address) (DataTypesPoolData, error) {
	return _IDataCompressor.Contract.GetPoolData(&_IDataCompressor.CallOpts, _pool)
}

// GetPoolData is a free data retrieval call binding the contract method 0x13d21cdf.
//
// Solidity: function getPoolData(address _pool) view returns((address,bool,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_IDataCompressor *IDataCompressorCallerSession) GetPoolData(_pool common.Address) (DataTypesPoolData, error) {
	return _IDataCompressor.Contract.GetPoolData(&_IDataCompressor.CallOpts, _pool)
}

// GetPoolsList is a free data retrieval call binding the contract method 0x1bcd8fc0.
//
// Solidity: function getPoolsList() view returns((address,bool,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256)[])
func (_IDataCompressor *IDataCompressorCaller) GetPoolsList(opts *bind.CallOpts) ([]DataTypesPoolData, error) {
	var out []interface{}
	err := _IDataCompressor.contract.Call(opts, &out, "getPoolsList")

	if err != nil {
		return *new([]DataTypesPoolData), err
	}

	out0 := *abi.ConvertType(out[0], new([]DataTypesPoolData)).(*[]DataTypesPoolData)

	return out0, err

}

// GetPoolsList is a free data retrieval call binding the contract method 0x1bcd8fc0.
//
// Solidity: function getPoolsList() view returns((address,bool,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256)[])
func (_IDataCompressor *IDataCompressorSession) GetPoolsList() ([]DataTypesPoolData, error) {
	return _IDataCompressor.Contract.GetPoolsList(&_IDataCompressor.CallOpts)
}

// GetPoolsList is a free data retrieval call binding the contract method 0x1bcd8fc0.
//
// Solidity: function getPoolsList() view returns((address,bool,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256)[])
func (_IDataCompressor *IDataCompressorCallerSession) GetPoolsList() ([]DataTypesPoolData, error) {
	return _IDataCompressor.Contract.GetPoolsList(&_IDataCompressor.CallOpts)
}

// GetTokenData is a free data retrieval call binding the contract method 0xbf2eb19e.
//
// Solidity: function getTokenData(address[] addr) view returns((address,string,uint8)[])
func (_IDataCompressor *IDataCompressorCaller) GetTokenData(opts *bind.CallOpts, addr []common.Address) ([]DataTypesTokenInfo, error) {
	var out []interface{}
	err := _IDataCompressor.contract.Call(opts, &out, "getTokenData", addr)

	if err != nil {
		return *new([]DataTypesTokenInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]DataTypesTokenInfo)).(*[]DataTypesTokenInfo)

	return out0, err

}

// GetTokenData is a free data retrieval call binding the contract method 0xbf2eb19e.
//
// Solidity: function getTokenData(address[] addr) view returns((address,string,uint8)[])
func (_IDataCompressor *IDataCompressorSession) GetTokenData(addr []common.Address) ([]DataTypesTokenInfo, error) {
	return _IDataCompressor.Contract.GetTokenData(&_IDataCompressor.CallOpts, addr)
}

// GetTokenData is a free data retrieval call binding the contract method 0xbf2eb19e.
//
// Solidity: function getTokenData(address[] addr) view returns((address,string,uint8)[])
func (_IDataCompressor *IDataCompressorCallerSession) GetTokenData(addr []common.Address) ([]DataTypesTokenInfo, error) {
	return _IDataCompressor.Contract.GetTokenData(&_IDataCompressor.CallOpts, addr)
}

// HasOpenedCreditAccount is a free data retrieval call binding the contract method 0xfc9914cb.
//
// Solidity: function hasOpenedCreditAccount(address creditManager, address borrower) view returns(bool)
func (_IDataCompressor *IDataCompressorCaller) HasOpenedCreditAccount(opts *bind.CallOpts, creditManager common.Address, borrower common.Address) (bool, error) {
	var out []interface{}
	err := _IDataCompressor.contract.Call(opts, &out, "hasOpenedCreditAccount", creditManager, borrower)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasOpenedCreditAccount is a free data retrieval call binding the contract method 0xfc9914cb.
//
// Solidity: function hasOpenedCreditAccount(address creditManager, address borrower) view returns(bool)
func (_IDataCompressor *IDataCompressorSession) HasOpenedCreditAccount(creditManager common.Address, borrower common.Address) (bool, error) {
	return _IDataCompressor.Contract.HasOpenedCreditAccount(&_IDataCompressor.CallOpts, creditManager, borrower)
}

// HasOpenedCreditAccount is a free data retrieval call binding the contract method 0xfc9914cb.
//
// Solidity: function hasOpenedCreditAccount(address creditManager, address borrower) view returns(bool)
func (_IDataCompressor *IDataCompressorCallerSession) HasOpenedCreditAccount(creditManager common.Address, borrower common.Address) (bool, error) {
	return _IDataCompressor.Contract.HasOpenedCreditAccount(&_IDataCompressor.CallOpts, creditManager, borrower)
}
