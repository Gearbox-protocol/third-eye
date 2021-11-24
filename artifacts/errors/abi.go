// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package errors

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

// ErrorsMetaData contains all meta data concerning the Errors contract.
var ErrorsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"ACL_CALLER_NOT_CONFIGURATOR\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ACL_CALLER_NOT_PAUSABLE_ADMIN\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"AF_CANT_CLOSE_CREDIT_ACCOUNT_IN_THE_SAME_BLOCK\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"AF_CREDIT_ACCOUNT_NOT_IN_STOCK\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"AF_EXTERNAL_ACCOUNTS_ARE_FORBIDDEN\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"AF_MINING_IS_FINISHED\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"AS_ADDRESS_NOT_FOUND\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CA_CONNECTED_CREDIT_MANAGER_ONLY\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CA_FACTORY_ONLY\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CF_ADAPTERS_ONLY\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CF_ADAPTER_CAN_BE_USED_ONLY_ONCE\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CF_CONTRACT_IS_NOT_IN_ALLOWED_LIST\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CF_CREDIT_MANAGERS_ONLY\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CF_CREDIT_MANAGER_IS_ALREADY_SET\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CF_FAST_CHECK_NOT_COVERED_COLLATERAL_DROP\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CF_INCORRECT_CHI_THRESHOLD\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CF_INCORRECT_FAST_CHECK\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CF_INCORRECT_LIQUIDATION_THRESHOLD\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CF_INCORRECT_PRICEFEED\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CF_NON_TOKEN_CONTRACT\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CF_OPERATION_LOW_HEALTH_FACTOR\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CF_SOME_LIQUIDATION_THRESHOLD_MORE_THAN_NEW_ONE\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CF_TOKEN_IS_NOT_ALLOWED\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CF_TOO_MUCH_ALLOWED_TOKENS\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CF_TRANSFER_IS_NOT_ALLOWED\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CF_UNDERLYING_TOKEN_FILTER_CONFLICT\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CM_CANT_CLOSE_WITH_LOSS\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CM_CAN_LIQUIDATE_WITH_SUCH_HEALTH_FACTOR\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CM_CAN_UPDATE_WITH_SUCH_HEALTH_FACTOR\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CM_INCORRECT_AMOUNT\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CM_INCORRECT_FEES\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CM_INCORRECT_NEW_OWNER\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CM_INCORRECT_PARAMS\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CM_MAX_LEVERAGE_IS_TOO_HIGH\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CM_NO_OPEN_ACCOUNT\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CM_TARGET_CONTRACT_iS_NOT_ALLOWED\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CM_TRANSFER_FAILED\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CM_WETH_GATEWAY_ONLY\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CM_ZERO_ADDRESS_OR_USER_HAVE_ALREADY_OPEN_CREDIT_ACCOUNT\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CR_CREDIT_MANAGER_ALREADY_ADDED\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CR_POOL_ALREADY_ADDED\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INCORRECT_ARRAY_LENGTH\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INCORRECT_PARAMETER\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INCORRECT_PATH_LENGTH\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LA_HAS_VALUE_WITH_TOKEN_TRANSFER\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LA_INCORRECT_VALUE\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LA_LOWER_THAN_AMOUNT_MIN\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LA_TOKEN_OUT_IS_NOT_COLLATERAL\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LA_UNKNOWN_LP_INTERFACE\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LA_UNKNOWN_SWAP_INTERFACE\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MATH_ADDITION_OVERFLOW\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MATH_DIVISION_BY_ZERO\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MATH_MULTIPLICATION_OVERFLOW\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NOT_IMPLEMENTED\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"POOL_CANT_ADD_CREDIT_MANAGER_TWICE\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"POOL_CONNECTED_CREDIT_MANAGERS_ONLY\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"POOL_INCOMPATIBLE_CREDIT_ACCOUNT_MANAGER\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"POOL_INCORRECT_WITHDRAW_FEE\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"POOL_MORE_THAN_EXPECTED_LIQUIDITY_LIMIT\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PO_AGGREGATOR_DECIMALS_SHOULD_BE_18\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PO_PRICE_FEED_DOESNT_EXIST\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PO_TOKENS_WITH_DECIMALS_MORE_18_ISNT_ALLOWED\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REGISTERED_CREDIT_ACCOUNT_MANAGERS_ONLY\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REGISTERED_POOLS_ONLY\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WG_DESTINATION_IS_NOT_WETH_COMPATIBLE\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WG_NOT_ENOUGH_FUNDS\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WG_RECEIVE_IS_NOT_ALLOWED\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ZERO_ADDRESS_IS_NOT_ALLOWED\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ErrorsABI is the input ABI used to generate the binding from.
// Deprecated: Use ErrorsMetaData.ABI instead.
var ErrorsABI = ErrorsMetaData.ABI

// Errors is an auto generated Go binding around an Ethereum contract.
type Errors struct {
	ErrorsCaller     // Read-only binding to the contract
	ErrorsTransactor // Write-only binding to the contract
	ErrorsFilterer   // Log filterer for contract events
}

// ErrorsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ErrorsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ErrorsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ErrorsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ErrorsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ErrorsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ErrorsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ErrorsSession struct {
	Contract     *Errors           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ErrorsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ErrorsCallerSession struct {
	Contract *ErrorsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ErrorsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ErrorsTransactorSession struct {
	Contract     *ErrorsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ErrorsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ErrorsRaw struct {
	Contract *Errors // Generic contract binding to access the raw methods on
}

// ErrorsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ErrorsCallerRaw struct {
	Contract *ErrorsCaller // Generic read-only contract binding to access the raw methods on
}

// ErrorsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ErrorsTransactorRaw struct {
	Contract *ErrorsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewErrors creates a new instance of Errors, bound to a specific deployed contract.
func NewErrors(address common.Address, backend bind.ContractBackend) (*Errors, error) {
	contract, err := bindErrors(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Errors{ErrorsCaller: ErrorsCaller{contract: contract}, ErrorsTransactor: ErrorsTransactor{contract: contract}, ErrorsFilterer: ErrorsFilterer{contract: contract}}, nil
}

// NewErrorsCaller creates a new read-only instance of Errors, bound to a specific deployed contract.
func NewErrorsCaller(address common.Address, caller bind.ContractCaller) (*ErrorsCaller, error) {
	contract, err := bindErrors(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ErrorsCaller{contract: contract}, nil
}

// NewErrorsTransactor creates a new write-only instance of Errors, bound to a specific deployed contract.
func NewErrorsTransactor(address common.Address, transactor bind.ContractTransactor) (*ErrorsTransactor, error) {
	contract, err := bindErrors(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ErrorsTransactor{contract: contract}, nil
}

// NewErrorsFilterer creates a new log filterer instance of Errors, bound to a specific deployed contract.
func NewErrorsFilterer(address common.Address, filterer bind.ContractFilterer) (*ErrorsFilterer, error) {
	contract, err := bindErrors(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ErrorsFilterer{contract: contract}, nil
}

// bindErrors binds a generic wrapper to an already deployed contract.
func bindErrors(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ErrorsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Errors *ErrorsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Errors.Contract.ErrorsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Errors *ErrorsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Errors.Contract.ErrorsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Errors *ErrorsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Errors.Contract.ErrorsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Errors *ErrorsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Errors.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Errors *ErrorsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Errors.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Errors *ErrorsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Errors.Contract.contract.Transact(opts, method, params...)
}

// ACLCALLERNOTCONFIGURATOR is a free data retrieval call binding the contract method 0xebbd977f.
//
// Solidity: function ACL_CALLER_NOT_CONFIGURATOR() view returns(string)
func (_Errors *ErrorsCaller) ACLCALLERNOTCONFIGURATOR(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "ACL_CALLER_NOT_CONFIGURATOR")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ACLCALLERNOTCONFIGURATOR is a free data retrieval call binding the contract method 0xebbd977f.
//
// Solidity: function ACL_CALLER_NOT_CONFIGURATOR() view returns(string)
func (_Errors *ErrorsSession) ACLCALLERNOTCONFIGURATOR() (string, error) {
	return _Errors.Contract.ACLCALLERNOTCONFIGURATOR(&_Errors.CallOpts)
}

// ACLCALLERNOTCONFIGURATOR is a free data retrieval call binding the contract method 0xebbd977f.
//
// Solidity: function ACL_CALLER_NOT_CONFIGURATOR() view returns(string)
func (_Errors *ErrorsCallerSession) ACLCALLERNOTCONFIGURATOR() (string, error) {
	return _Errors.Contract.ACLCALLERNOTCONFIGURATOR(&_Errors.CallOpts)
}

// ACLCALLERNOTPAUSABLEADMIN is a free data retrieval call binding the contract method 0xa988ac60.
//
// Solidity: function ACL_CALLER_NOT_PAUSABLE_ADMIN() view returns(string)
func (_Errors *ErrorsCaller) ACLCALLERNOTPAUSABLEADMIN(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "ACL_CALLER_NOT_PAUSABLE_ADMIN")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ACLCALLERNOTPAUSABLEADMIN is a free data retrieval call binding the contract method 0xa988ac60.
//
// Solidity: function ACL_CALLER_NOT_PAUSABLE_ADMIN() view returns(string)
func (_Errors *ErrorsSession) ACLCALLERNOTPAUSABLEADMIN() (string, error) {
	return _Errors.Contract.ACLCALLERNOTPAUSABLEADMIN(&_Errors.CallOpts)
}

// ACLCALLERNOTPAUSABLEADMIN is a free data retrieval call binding the contract method 0xa988ac60.
//
// Solidity: function ACL_CALLER_NOT_PAUSABLE_ADMIN() view returns(string)
func (_Errors *ErrorsCallerSession) ACLCALLERNOTPAUSABLEADMIN() (string, error) {
	return _Errors.Contract.ACLCALLERNOTPAUSABLEADMIN(&_Errors.CallOpts)
}

// AFCANTCLOSECREDITACCOUNTINTHESAMEBLOCK is a free data retrieval call binding the contract method 0x0c9409e7.
//
// Solidity: function AF_CANT_CLOSE_CREDIT_ACCOUNT_IN_THE_SAME_BLOCK() view returns(string)
func (_Errors *ErrorsCaller) AFCANTCLOSECREDITACCOUNTINTHESAMEBLOCK(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "AF_CANT_CLOSE_CREDIT_ACCOUNT_IN_THE_SAME_BLOCK")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// AFCANTCLOSECREDITACCOUNTINTHESAMEBLOCK is a free data retrieval call binding the contract method 0x0c9409e7.
//
// Solidity: function AF_CANT_CLOSE_CREDIT_ACCOUNT_IN_THE_SAME_BLOCK() view returns(string)
func (_Errors *ErrorsSession) AFCANTCLOSECREDITACCOUNTINTHESAMEBLOCK() (string, error) {
	return _Errors.Contract.AFCANTCLOSECREDITACCOUNTINTHESAMEBLOCK(&_Errors.CallOpts)
}

// AFCANTCLOSECREDITACCOUNTINTHESAMEBLOCK is a free data retrieval call binding the contract method 0x0c9409e7.
//
// Solidity: function AF_CANT_CLOSE_CREDIT_ACCOUNT_IN_THE_SAME_BLOCK() view returns(string)
func (_Errors *ErrorsCallerSession) AFCANTCLOSECREDITACCOUNTINTHESAMEBLOCK() (string, error) {
	return _Errors.Contract.AFCANTCLOSECREDITACCOUNTINTHESAMEBLOCK(&_Errors.CallOpts)
}

// AFCREDITACCOUNTNOTINSTOCK is a free data retrieval call binding the contract method 0xac757139.
//
// Solidity: function AF_CREDIT_ACCOUNT_NOT_IN_STOCK() view returns(string)
func (_Errors *ErrorsCaller) AFCREDITACCOUNTNOTINSTOCK(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "AF_CREDIT_ACCOUNT_NOT_IN_STOCK")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// AFCREDITACCOUNTNOTINSTOCK is a free data retrieval call binding the contract method 0xac757139.
//
// Solidity: function AF_CREDIT_ACCOUNT_NOT_IN_STOCK() view returns(string)
func (_Errors *ErrorsSession) AFCREDITACCOUNTNOTINSTOCK() (string, error) {
	return _Errors.Contract.AFCREDITACCOUNTNOTINSTOCK(&_Errors.CallOpts)
}

// AFCREDITACCOUNTNOTINSTOCK is a free data retrieval call binding the contract method 0xac757139.
//
// Solidity: function AF_CREDIT_ACCOUNT_NOT_IN_STOCK() view returns(string)
func (_Errors *ErrorsCallerSession) AFCREDITACCOUNTNOTINSTOCK() (string, error) {
	return _Errors.Contract.AFCREDITACCOUNTNOTINSTOCK(&_Errors.CallOpts)
}

// AFEXTERNALACCOUNTSAREFORBIDDEN is a free data retrieval call binding the contract method 0xd1a65a38.
//
// Solidity: function AF_EXTERNAL_ACCOUNTS_ARE_FORBIDDEN() view returns(string)
func (_Errors *ErrorsCaller) AFEXTERNALACCOUNTSAREFORBIDDEN(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "AF_EXTERNAL_ACCOUNTS_ARE_FORBIDDEN")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// AFEXTERNALACCOUNTSAREFORBIDDEN is a free data retrieval call binding the contract method 0xd1a65a38.
//
// Solidity: function AF_EXTERNAL_ACCOUNTS_ARE_FORBIDDEN() view returns(string)
func (_Errors *ErrorsSession) AFEXTERNALACCOUNTSAREFORBIDDEN() (string, error) {
	return _Errors.Contract.AFEXTERNALACCOUNTSAREFORBIDDEN(&_Errors.CallOpts)
}

// AFEXTERNALACCOUNTSAREFORBIDDEN is a free data retrieval call binding the contract method 0xd1a65a38.
//
// Solidity: function AF_EXTERNAL_ACCOUNTS_ARE_FORBIDDEN() view returns(string)
func (_Errors *ErrorsCallerSession) AFEXTERNALACCOUNTSAREFORBIDDEN() (string, error) {
	return _Errors.Contract.AFEXTERNALACCOUNTSAREFORBIDDEN(&_Errors.CallOpts)
}

// AFMININGISFINISHED is a free data retrieval call binding the contract method 0x87f88ef4.
//
// Solidity: function AF_MINING_IS_FINISHED() view returns(string)
func (_Errors *ErrorsCaller) AFMININGISFINISHED(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "AF_MINING_IS_FINISHED")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// AFMININGISFINISHED is a free data retrieval call binding the contract method 0x87f88ef4.
//
// Solidity: function AF_MINING_IS_FINISHED() view returns(string)
func (_Errors *ErrorsSession) AFMININGISFINISHED() (string, error) {
	return _Errors.Contract.AFMININGISFINISHED(&_Errors.CallOpts)
}

// AFMININGISFINISHED is a free data retrieval call binding the contract method 0x87f88ef4.
//
// Solidity: function AF_MINING_IS_FINISHED() view returns(string)
func (_Errors *ErrorsCallerSession) AFMININGISFINISHED() (string, error) {
	return _Errors.Contract.AFMININGISFINISHED(&_Errors.CallOpts)
}

// ASADDRESSNOTFOUND is a free data retrieval call binding the contract method 0xde63cd40.
//
// Solidity: function AS_ADDRESS_NOT_FOUND() view returns(string)
func (_Errors *ErrorsCaller) ASADDRESSNOTFOUND(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "AS_ADDRESS_NOT_FOUND")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ASADDRESSNOTFOUND is a free data retrieval call binding the contract method 0xde63cd40.
//
// Solidity: function AS_ADDRESS_NOT_FOUND() view returns(string)
func (_Errors *ErrorsSession) ASADDRESSNOTFOUND() (string, error) {
	return _Errors.Contract.ASADDRESSNOTFOUND(&_Errors.CallOpts)
}

// ASADDRESSNOTFOUND is a free data retrieval call binding the contract method 0xde63cd40.
//
// Solidity: function AS_ADDRESS_NOT_FOUND() view returns(string)
func (_Errors *ErrorsCallerSession) ASADDRESSNOTFOUND() (string, error) {
	return _Errors.Contract.ASADDRESSNOTFOUND(&_Errors.CallOpts)
}

// CACONNECTEDCREDITMANAGERONLY is a free data retrieval call binding the contract method 0xff2a04e3.
//
// Solidity: function CA_CONNECTED_CREDIT_MANAGER_ONLY() view returns(string)
func (_Errors *ErrorsCaller) CACONNECTEDCREDITMANAGERONLY(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "CA_CONNECTED_CREDIT_MANAGER_ONLY")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CACONNECTEDCREDITMANAGERONLY is a free data retrieval call binding the contract method 0xff2a04e3.
//
// Solidity: function CA_CONNECTED_CREDIT_MANAGER_ONLY() view returns(string)
func (_Errors *ErrorsSession) CACONNECTEDCREDITMANAGERONLY() (string, error) {
	return _Errors.Contract.CACONNECTEDCREDITMANAGERONLY(&_Errors.CallOpts)
}

// CACONNECTEDCREDITMANAGERONLY is a free data retrieval call binding the contract method 0xff2a04e3.
//
// Solidity: function CA_CONNECTED_CREDIT_MANAGER_ONLY() view returns(string)
func (_Errors *ErrorsCallerSession) CACONNECTEDCREDITMANAGERONLY() (string, error) {
	return _Errors.Contract.CACONNECTEDCREDITMANAGERONLY(&_Errors.CallOpts)
}

// CAFACTORYONLY is a free data retrieval call binding the contract method 0x2357f362.
//
// Solidity: function CA_FACTORY_ONLY() view returns(string)
func (_Errors *ErrorsCaller) CAFACTORYONLY(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "CA_FACTORY_ONLY")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CAFACTORYONLY is a free data retrieval call binding the contract method 0x2357f362.
//
// Solidity: function CA_FACTORY_ONLY() view returns(string)
func (_Errors *ErrorsSession) CAFACTORYONLY() (string, error) {
	return _Errors.Contract.CAFACTORYONLY(&_Errors.CallOpts)
}

// CAFACTORYONLY is a free data retrieval call binding the contract method 0x2357f362.
//
// Solidity: function CA_FACTORY_ONLY() view returns(string)
func (_Errors *ErrorsCallerSession) CAFACTORYONLY() (string, error) {
	return _Errors.Contract.CAFACTORYONLY(&_Errors.CallOpts)
}

// CFADAPTERSONLY is a free data retrieval call binding the contract method 0x202a591b.
//
// Solidity: function CF_ADAPTERS_ONLY() view returns(string)
func (_Errors *ErrorsCaller) CFADAPTERSONLY(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "CF_ADAPTERS_ONLY")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CFADAPTERSONLY is a free data retrieval call binding the contract method 0x202a591b.
//
// Solidity: function CF_ADAPTERS_ONLY() view returns(string)
func (_Errors *ErrorsSession) CFADAPTERSONLY() (string, error) {
	return _Errors.Contract.CFADAPTERSONLY(&_Errors.CallOpts)
}

// CFADAPTERSONLY is a free data retrieval call binding the contract method 0x202a591b.
//
// Solidity: function CF_ADAPTERS_ONLY() view returns(string)
func (_Errors *ErrorsCallerSession) CFADAPTERSONLY() (string, error) {
	return _Errors.Contract.CFADAPTERSONLY(&_Errors.CallOpts)
}

// CFADAPTERCANBEUSEDONLYONCE is a free data retrieval call binding the contract method 0x1e4dfd83.
//
// Solidity: function CF_ADAPTER_CAN_BE_USED_ONLY_ONCE() view returns(string)
func (_Errors *ErrorsCaller) CFADAPTERCANBEUSEDONLYONCE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "CF_ADAPTER_CAN_BE_USED_ONLY_ONCE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CFADAPTERCANBEUSEDONLYONCE is a free data retrieval call binding the contract method 0x1e4dfd83.
//
// Solidity: function CF_ADAPTER_CAN_BE_USED_ONLY_ONCE() view returns(string)
func (_Errors *ErrorsSession) CFADAPTERCANBEUSEDONLYONCE() (string, error) {
	return _Errors.Contract.CFADAPTERCANBEUSEDONLYONCE(&_Errors.CallOpts)
}

// CFADAPTERCANBEUSEDONLYONCE is a free data retrieval call binding the contract method 0x1e4dfd83.
//
// Solidity: function CF_ADAPTER_CAN_BE_USED_ONLY_ONCE() view returns(string)
func (_Errors *ErrorsCallerSession) CFADAPTERCANBEUSEDONLYONCE() (string, error) {
	return _Errors.Contract.CFADAPTERCANBEUSEDONLYONCE(&_Errors.CallOpts)
}

// CFCONTRACTISNOTINALLOWEDLIST is a free data retrieval call binding the contract method 0xad37d10b.
//
// Solidity: function CF_CONTRACT_IS_NOT_IN_ALLOWED_LIST() view returns(string)
func (_Errors *ErrorsCaller) CFCONTRACTISNOTINALLOWEDLIST(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "CF_CONTRACT_IS_NOT_IN_ALLOWED_LIST")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CFCONTRACTISNOTINALLOWEDLIST is a free data retrieval call binding the contract method 0xad37d10b.
//
// Solidity: function CF_CONTRACT_IS_NOT_IN_ALLOWED_LIST() view returns(string)
func (_Errors *ErrorsSession) CFCONTRACTISNOTINALLOWEDLIST() (string, error) {
	return _Errors.Contract.CFCONTRACTISNOTINALLOWEDLIST(&_Errors.CallOpts)
}

// CFCONTRACTISNOTINALLOWEDLIST is a free data retrieval call binding the contract method 0xad37d10b.
//
// Solidity: function CF_CONTRACT_IS_NOT_IN_ALLOWED_LIST() view returns(string)
func (_Errors *ErrorsCallerSession) CFCONTRACTISNOTINALLOWEDLIST() (string, error) {
	return _Errors.Contract.CFCONTRACTISNOTINALLOWEDLIST(&_Errors.CallOpts)
}

// CFCREDITMANAGERSONLY is a free data retrieval call binding the contract method 0xea2c3a00.
//
// Solidity: function CF_CREDIT_MANAGERS_ONLY() view returns(string)
func (_Errors *ErrorsCaller) CFCREDITMANAGERSONLY(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "CF_CREDIT_MANAGERS_ONLY")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CFCREDITMANAGERSONLY is a free data retrieval call binding the contract method 0xea2c3a00.
//
// Solidity: function CF_CREDIT_MANAGERS_ONLY() view returns(string)
func (_Errors *ErrorsSession) CFCREDITMANAGERSONLY() (string, error) {
	return _Errors.Contract.CFCREDITMANAGERSONLY(&_Errors.CallOpts)
}

// CFCREDITMANAGERSONLY is a free data retrieval call binding the contract method 0xea2c3a00.
//
// Solidity: function CF_CREDIT_MANAGERS_ONLY() view returns(string)
func (_Errors *ErrorsCallerSession) CFCREDITMANAGERSONLY() (string, error) {
	return _Errors.Contract.CFCREDITMANAGERSONLY(&_Errors.CallOpts)
}

// CFCREDITMANAGERISALREADYSET is a free data retrieval call binding the contract method 0x1fac3a13.
//
// Solidity: function CF_CREDIT_MANAGER_IS_ALREADY_SET() view returns(string)
func (_Errors *ErrorsCaller) CFCREDITMANAGERISALREADYSET(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "CF_CREDIT_MANAGER_IS_ALREADY_SET")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CFCREDITMANAGERISALREADYSET is a free data retrieval call binding the contract method 0x1fac3a13.
//
// Solidity: function CF_CREDIT_MANAGER_IS_ALREADY_SET() view returns(string)
func (_Errors *ErrorsSession) CFCREDITMANAGERISALREADYSET() (string, error) {
	return _Errors.Contract.CFCREDITMANAGERISALREADYSET(&_Errors.CallOpts)
}

// CFCREDITMANAGERISALREADYSET is a free data retrieval call binding the contract method 0x1fac3a13.
//
// Solidity: function CF_CREDIT_MANAGER_IS_ALREADY_SET() view returns(string)
func (_Errors *ErrorsCallerSession) CFCREDITMANAGERISALREADYSET() (string, error) {
	return _Errors.Contract.CFCREDITMANAGERISALREADYSET(&_Errors.CallOpts)
}

// CFFASTCHECKNOTCOVEREDCOLLATERALDROP is a free data retrieval call binding the contract method 0x277fa84f.
//
// Solidity: function CF_FAST_CHECK_NOT_COVERED_COLLATERAL_DROP() view returns(string)
func (_Errors *ErrorsCaller) CFFASTCHECKNOTCOVEREDCOLLATERALDROP(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "CF_FAST_CHECK_NOT_COVERED_COLLATERAL_DROP")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CFFASTCHECKNOTCOVEREDCOLLATERALDROP is a free data retrieval call binding the contract method 0x277fa84f.
//
// Solidity: function CF_FAST_CHECK_NOT_COVERED_COLLATERAL_DROP() view returns(string)
func (_Errors *ErrorsSession) CFFASTCHECKNOTCOVEREDCOLLATERALDROP() (string, error) {
	return _Errors.Contract.CFFASTCHECKNOTCOVEREDCOLLATERALDROP(&_Errors.CallOpts)
}

// CFFASTCHECKNOTCOVEREDCOLLATERALDROP is a free data retrieval call binding the contract method 0x277fa84f.
//
// Solidity: function CF_FAST_CHECK_NOT_COVERED_COLLATERAL_DROP() view returns(string)
func (_Errors *ErrorsCallerSession) CFFASTCHECKNOTCOVEREDCOLLATERALDROP() (string, error) {
	return _Errors.Contract.CFFASTCHECKNOTCOVEREDCOLLATERALDROP(&_Errors.CallOpts)
}

// CFINCORRECTCHITHRESHOLD is a free data retrieval call binding the contract method 0xbdc36a02.
//
// Solidity: function CF_INCORRECT_CHI_THRESHOLD() view returns(string)
func (_Errors *ErrorsCaller) CFINCORRECTCHITHRESHOLD(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "CF_INCORRECT_CHI_THRESHOLD")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CFINCORRECTCHITHRESHOLD is a free data retrieval call binding the contract method 0xbdc36a02.
//
// Solidity: function CF_INCORRECT_CHI_THRESHOLD() view returns(string)
func (_Errors *ErrorsSession) CFINCORRECTCHITHRESHOLD() (string, error) {
	return _Errors.Contract.CFINCORRECTCHITHRESHOLD(&_Errors.CallOpts)
}

// CFINCORRECTCHITHRESHOLD is a free data retrieval call binding the contract method 0xbdc36a02.
//
// Solidity: function CF_INCORRECT_CHI_THRESHOLD() view returns(string)
func (_Errors *ErrorsCallerSession) CFINCORRECTCHITHRESHOLD() (string, error) {
	return _Errors.Contract.CFINCORRECTCHITHRESHOLD(&_Errors.CallOpts)
}

// CFINCORRECTFASTCHECK is a free data retrieval call binding the contract method 0xc02e57ef.
//
// Solidity: function CF_INCORRECT_FAST_CHECK() view returns(string)
func (_Errors *ErrorsCaller) CFINCORRECTFASTCHECK(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "CF_INCORRECT_FAST_CHECK")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CFINCORRECTFASTCHECK is a free data retrieval call binding the contract method 0xc02e57ef.
//
// Solidity: function CF_INCORRECT_FAST_CHECK() view returns(string)
func (_Errors *ErrorsSession) CFINCORRECTFASTCHECK() (string, error) {
	return _Errors.Contract.CFINCORRECTFASTCHECK(&_Errors.CallOpts)
}

// CFINCORRECTFASTCHECK is a free data retrieval call binding the contract method 0xc02e57ef.
//
// Solidity: function CF_INCORRECT_FAST_CHECK() view returns(string)
func (_Errors *ErrorsCallerSession) CFINCORRECTFASTCHECK() (string, error) {
	return _Errors.Contract.CFINCORRECTFASTCHECK(&_Errors.CallOpts)
}

// CFINCORRECTLIQUIDATIONTHRESHOLD is a free data retrieval call binding the contract method 0x6f92ed92.
//
// Solidity: function CF_INCORRECT_LIQUIDATION_THRESHOLD() view returns(string)
func (_Errors *ErrorsCaller) CFINCORRECTLIQUIDATIONTHRESHOLD(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "CF_INCORRECT_LIQUIDATION_THRESHOLD")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CFINCORRECTLIQUIDATIONTHRESHOLD is a free data retrieval call binding the contract method 0x6f92ed92.
//
// Solidity: function CF_INCORRECT_LIQUIDATION_THRESHOLD() view returns(string)
func (_Errors *ErrorsSession) CFINCORRECTLIQUIDATIONTHRESHOLD() (string, error) {
	return _Errors.Contract.CFINCORRECTLIQUIDATIONTHRESHOLD(&_Errors.CallOpts)
}

// CFINCORRECTLIQUIDATIONTHRESHOLD is a free data retrieval call binding the contract method 0x6f92ed92.
//
// Solidity: function CF_INCORRECT_LIQUIDATION_THRESHOLD() view returns(string)
func (_Errors *ErrorsCallerSession) CFINCORRECTLIQUIDATIONTHRESHOLD() (string, error) {
	return _Errors.Contract.CFINCORRECTLIQUIDATIONTHRESHOLD(&_Errors.CallOpts)
}

// CFINCORRECTPRICEFEED is a free data retrieval call binding the contract method 0x6b3c35aa.
//
// Solidity: function CF_INCORRECT_PRICEFEED() view returns(string)
func (_Errors *ErrorsCaller) CFINCORRECTPRICEFEED(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "CF_INCORRECT_PRICEFEED")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CFINCORRECTPRICEFEED is a free data retrieval call binding the contract method 0x6b3c35aa.
//
// Solidity: function CF_INCORRECT_PRICEFEED() view returns(string)
func (_Errors *ErrorsSession) CFINCORRECTPRICEFEED() (string, error) {
	return _Errors.Contract.CFINCORRECTPRICEFEED(&_Errors.CallOpts)
}

// CFINCORRECTPRICEFEED is a free data retrieval call binding the contract method 0x6b3c35aa.
//
// Solidity: function CF_INCORRECT_PRICEFEED() view returns(string)
func (_Errors *ErrorsCallerSession) CFINCORRECTPRICEFEED() (string, error) {
	return _Errors.Contract.CFINCORRECTPRICEFEED(&_Errors.CallOpts)
}

// CFNONTOKENCONTRACT is a free data retrieval call binding the contract method 0x8ea76e22.
//
// Solidity: function CF_NON_TOKEN_CONTRACT() view returns(string)
func (_Errors *ErrorsCaller) CFNONTOKENCONTRACT(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "CF_NON_TOKEN_CONTRACT")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CFNONTOKENCONTRACT is a free data retrieval call binding the contract method 0x8ea76e22.
//
// Solidity: function CF_NON_TOKEN_CONTRACT() view returns(string)
func (_Errors *ErrorsSession) CFNONTOKENCONTRACT() (string, error) {
	return _Errors.Contract.CFNONTOKENCONTRACT(&_Errors.CallOpts)
}

// CFNONTOKENCONTRACT is a free data retrieval call binding the contract method 0x8ea76e22.
//
// Solidity: function CF_NON_TOKEN_CONTRACT() view returns(string)
func (_Errors *ErrorsCallerSession) CFNONTOKENCONTRACT() (string, error) {
	return _Errors.Contract.CFNONTOKENCONTRACT(&_Errors.CallOpts)
}

// CFOPERATIONLOWHEALTHFACTOR is a free data retrieval call binding the contract method 0xb3d5b5d6.
//
// Solidity: function CF_OPERATION_LOW_HEALTH_FACTOR() view returns(string)
func (_Errors *ErrorsCaller) CFOPERATIONLOWHEALTHFACTOR(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "CF_OPERATION_LOW_HEALTH_FACTOR")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CFOPERATIONLOWHEALTHFACTOR is a free data retrieval call binding the contract method 0xb3d5b5d6.
//
// Solidity: function CF_OPERATION_LOW_HEALTH_FACTOR() view returns(string)
func (_Errors *ErrorsSession) CFOPERATIONLOWHEALTHFACTOR() (string, error) {
	return _Errors.Contract.CFOPERATIONLOWHEALTHFACTOR(&_Errors.CallOpts)
}

// CFOPERATIONLOWHEALTHFACTOR is a free data retrieval call binding the contract method 0xb3d5b5d6.
//
// Solidity: function CF_OPERATION_LOW_HEALTH_FACTOR() view returns(string)
func (_Errors *ErrorsCallerSession) CFOPERATIONLOWHEALTHFACTOR() (string, error) {
	return _Errors.Contract.CFOPERATIONLOWHEALTHFACTOR(&_Errors.CallOpts)
}

// CFSOMELIQUIDATIONTHRESHOLDMORETHANNEWONE is a free data retrieval call binding the contract method 0xe8b14757.
//
// Solidity: function CF_SOME_LIQUIDATION_THRESHOLD_MORE_THAN_NEW_ONE() view returns(string)
func (_Errors *ErrorsCaller) CFSOMELIQUIDATIONTHRESHOLDMORETHANNEWONE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "CF_SOME_LIQUIDATION_THRESHOLD_MORE_THAN_NEW_ONE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CFSOMELIQUIDATIONTHRESHOLDMORETHANNEWONE is a free data retrieval call binding the contract method 0xe8b14757.
//
// Solidity: function CF_SOME_LIQUIDATION_THRESHOLD_MORE_THAN_NEW_ONE() view returns(string)
func (_Errors *ErrorsSession) CFSOMELIQUIDATIONTHRESHOLDMORETHANNEWONE() (string, error) {
	return _Errors.Contract.CFSOMELIQUIDATIONTHRESHOLDMORETHANNEWONE(&_Errors.CallOpts)
}

// CFSOMELIQUIDATIONTHRESHOLDMORETHANNEWONE is a free data retrieval call binding the contract method 0xe8b14757.
//
// Solidity: function CF_SOME_LIQUIDATION_THRESHOLD_MORE_THAN_NEW_ONE() view returns(string)
func (_Errors *ErrorsCallerSession) CFSOMELIQUIDATIONTHRESHOLDMORETHANNEWONE() (string, error) {
	return _Errors.Contract.CFSOMELIQUIDATIONTHRESHOLDMORETHANNEWONE(&_Errors.CallOpts)
}

// CFTOKENISNOTALLOWED is a free data retrieval call binding the contract method 0x1dd371f8.
//
// Solidity: function CF_TOKEN_IS_NOT_ALLOWED() view returns(string)
func (_Errors *ErrorsCaller) CFTOKENISNOTALLOWED(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "CF_TOKEN_IS_NOT_ALLOWED")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CFTOKENISNOTALLOWED is a free data retrieval call binding the contract method 0x1dd371f8.
//
// Solidity: function CF_TOKEN_IS_NOT_ALLOWED() view returns(string)
func (_Errors *ErrorsSession) CFTOKENISNOTALLOWED() (string, error) {
	return _Errors.Contract.CFTOKENISNOTALLOWED(&_Errors.CallOpts)
}

// CFTOKENISNOTALLOWED is a free data retrieval call binding the contract method 0x1dd371f8.
//
// Solidity: function CF_TOKEN_IS_NOT_ALLOWED() view returns(string)
func (_Errors *ErrorsCallerSession) CFTOKENISNOTALLOWED() (string, error) {
	return _Errors.Contract.CFTOKENISNOTALLOWED(&_Errors.CallOpts)
}

// CFTOOMUCHALLOWEDTOKENS is a free data retrieval call binding the contract method 0x8aed8b2c.
//
// Solidity: function CF_TOO_MUCH_ALLOWED_TOKENS() view returns(string)
func (_Errors *ErrorsCaller) CFTOOMUCHALLOWEDTOKENS(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "CF_TOO_MUCH_ALLOWED_TOKENS")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CFTOOMUCHALLOWEDTOKENS is a free data retrieval call binding the contract method 0x8aed8b2c.
//
// Solidity: function CF_TOO_MUCH_ALLOWED_TOKENS() view returns(string)
func (_Errors *ErrorsSession) CFTOOMUCHALLOWEDTOKENS() (string, error) {
	return _Errors.Contract.CFTOOMUCHALLOWEDTOKENS(&_Errors.CallOpts)
}

// CFTOOMUCHALLOWEDTOKENS is a free data retrieval call binding the contract method 0x8aed8b2c.
//
// Solidity: function CF_TOO_MUCH_ALLOWED_TOKENS() view returns(string)
func (_Errors *ErrorsCallerSession) CFTOOMUCHALLOWEDTOKENS() (string, error) {
	return _Errors.Contract.CFTOOMUCHALLOWEDTOKENS(&_Errors.CallOpts)
}

// CFTRANSFERISNOTALLOWED is a free data retrieval call binding the contract method 0x83168a9b.
//
// Solidity: function CF_TRANSFER_IS_NOT_ALLOWED() view returns(string)
func (_Errors *ErrorsCaller) CFTRANSFERISNOTALLOWED(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "CF_TRANSFER_IS_NOT_ALLOWED")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CFTRANSFERISNOTALLOWED is a free data retrieval call binding the contract method 0x83168a9b.
//
// Solidity: function CF_TRANSFER_IS_NOT_ALLOWED() view returns(string)
func (_Errors *ErrorsSession) CFTRANSFERISNOTALLOWED() (string, error) {
	return _Errors.Contract.CFTRANSFERISNOTALLOWED(&_Errors.CallOpts)
}

// CFTRANSFERISNOTALLOWED is a free data retrieval call binding the contract method 0x83168a9b.
//
// Solidity: function CF_TRANSFER_IS_NOT_ALLOWED() view returns(string)
func (_Errors *ErrorsCallerSession) CFTRANSFERISNOTALLOWED() (string, error) {
	return _Errors.Contract.CFTRANSFERISNOTALLOWED(&_Errors.CallOpts)
}

// CFUNDERLYINGTOKENFILTERCONFLICT is a free data retrieval call binding the contract method 0x6c863867.
//
// Solidity: function CF_UNDERLYING_TOKEN_FILTER_CONFLICT() view returns(string)
func (_Errors *ErrorsCaller) CFUNDERLYINGTOKENFILTERCONFLICT(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "CF_UNDERLYING_TOKEN_FILTER_CONFLICT")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CFUNDERLYINGTOKENFILTERCONFLICT is a free data retrieval call binding the contract method 0x6c863867.
//
// Solidity: function CF_UNDERLYING_TOKEN_FILTER_CONFLICT() view returns(string)
func (_Errors *ErrorsSession) CFUNDERLYINGTOKENFILTERCONFLICT() (string, error) {
	return _Errors.Contract.CFUNDERLYINGTOKENFILTERCONFLICT(&_Errors.CallOpts)
}

// CFUNDERLYINGTOKENFILTERCONFLICT is a free data retrieval call binding the contract method 0x6c863867.
//
// Solidity: function CF_UNDERLYING_TOKEN_FILTER_CONFLICT() view returns(string)
func (_Errors *ErrorsCallerSession) CFUNDERLYINGTOKENFILTERCONFLICT() (string, error) {
	return _Errors.Contract.CFUNDERLYINGTOKENFILTERCONFLICT(&_Errors.CallOpts)
}

// CMCANTCLOSEWITHLOSS is a free data retrieval call binding the contract method 0x5a7afe48.
//
// Solidity: function CM_CANT_CLOSE_WITH_LOSS() view returns(string)
func (_Errors *ErrorsCaller) CMCANTCLOSEWITHLOSS(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "CM_CANT_CLOSE_WITH_LOSS")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CMCANTCLOSEWITHLOSS is a free data retrieval call binding the contract method 0x5a7afe48.
//
// Solidity: function CM_CANT_CLOSE_WITH_LOSS() view returns(string)
func (_Errors *ErrorsSession) CMCANTCLOSEWITHLOSS() (string, error) {
	return _Errors.Contract.CMCANTCLOSEWITHLOSS(&_Errors.CallOpts)
}

// CMCANTCLOSEWITHLOSS is a free data retrieval call binding the contract method 0x5a7afe48.
//
// Solidity: function CM_CANT_CLOSE_WITH_LOSS() view returns(string)
func (_Errors *ErrorsCallerSession) CMCANTCLOSEWITHLOSS() (string, error) {
	return _Errors.Contract.CMCANTCLOSEWITHLOSS(&_Errors.CallOpts)
}

// CMCANLIQUIDATEWITHSUCHHEALTHFACTOR is a free data retrieval call binding the contract method 0xa4249812.
//
// Solidity: function CM_CAN_LIQUIDATE_WITH_SUCH_HEALTH_FACTOR() view returns(string)
func (_Errors *ErrorsCaller) CMCANLIQUIDATEWITHSUCHHEALTHFACTOR(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "CM_CAN_LIQUIDATE_WITH_SUCH_HEALTH_FACTOR")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CMCANLIQUIDATEWITHSUCHHEALTHFACTOR is a free data retrieval call binding the contract method 0xa4249812.
//
// Solidity: function CM_CAN_LIQUIDATE_WITH_SUCH_HEALTH_FACTOR() view returns(string)
func (_Errors *ErrorsSession) CMCANLIQUIDATEWITHSUCHHEALTHFACTOR() (string, error) {
	return _Errors.Contract.CMCANLIQUIDATEWITHSUCHHEALTHFACTOR(&_Errors.CallOpts)
}

// CMCANLIQUIDATEWITHSUCHHEALTHFACTOR is a free data retrieval call binding the contract method 0xa4249812.
//
// Solidity: function CM_CAN_LIQUIDATE_WITH_SUCH_HEALTH_FACTOR() view returns(string)
func (_Errors *ErrorsCallerSession) CMCANLIQUIDATEWITHSUCHHEALTHFACTOR() (string, error) {
	return _Errors.Contract.CMCANLIQUIDATEWITHSUCHHEALTHFACTOR(&_Errors.CallOpts)
}

// CMCANUPDATEWITHSUCHHEALTHFACTOR is a free data retrieval call binding the contract method 0x93f7dc3c.
//
// Solidity: function CM_CAN_UPDATE_WITH_SUCH_HEALTH_FACTOR() view returns(string)
func (_Errors *ErrorsCaller) CMCANUPDATEWITHSUCHHEALTHFACTOR(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "CM_CAN_UPDATE_WITH_SUCH_HEALTH_FACTOR")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CMCANUPDATEWITHSUCHHEALTHFACTOR is a free data retrieval call binding the contract method 0x93f7dc3c.
//
// Solidity: function CM_CAN_UPDATE_WITH_SUCH_HEALTH_FACTOR() view returns(string)
func (_Errors *ErrorsSession) CMCANUPDATEWITHSUCHHEALTHFACTOR() (string, error) {
	return _Errors.Contract.CMCANUPDATEWITHSUCHHEALTHFACTOR(&_Errors.CallOpts)
}

// CMCANUPDATEWITHSUCHHEALTHFACTOR is a free data retrieval call binding the contract method 0x93f7dc3c.
//
// Solidity: function CM_CAN_UPDATE_WITH_SUCH_HEALTH_FACTOR() view returns(string)
func (_Errors *ErrorsCallerSession) CMCANUPDATEWITHSUCHHEALTHFACTOR() (string, error) {
	return _Errors.Contract.CMCANUPDATEWITHSUCHHEALTHFACTOR(&_Errors.CallOpts)
}

// CMINCORRECTAMOUNT is a free data retrieval call binding the contract method 0x69c3ae16.
//
// Solidity: function CM_INCORRECT_AMOUNT() view returns(string)
func (_Errors *ErrorsCaller) CMINCORRECTAMOUNT(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "CM_INCORRECT_AMOUNT")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CMINCORRECTAMOUNT is a free data retrieval call binding the contract method 0x69c3ae16.
//
// Solidity: function CM_INCORRECT_AMOUNT() view returns(string)
func (_Errors *ErrorsSession) CMINCORRECTAMOUNT() (string, error) {
	return _Errors.Contract.CMINCORRECTAMOUNT(&_Errors.CallOpts)
}

// CMINCORRECTAMOUNT is a free data retrieval call binding the contract method 0x69c3ae16.
//
// Solidity: function CM_INCORRECT_AMOUNT() view returns(string)
func (_Errors *ErrorsCallerSession) CMINCORRECTAMOUNT() (string, error) {
	return _Errors.Contract.CMINCORRECTAMOUNT(&_Errors.CallOpts)
}

// CMINCORRECTFEES is a free data retrieval call binding the contract method 0xde10ab9a.
//
// Solidity: function CM_INCORRECT_FEES() view returns(string)
func (_Errors *ErrorsCaller) CMINCORRECTFEES(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "CM_INCORRECT_FEES")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CMINCORRECTFEES is a free data retrieval call binding the contract method 0xde10ab9a.
//
// Solidity: function CM_INCORRECT_FEES() view returns(string)
func (_Errors *ErrorsSession) CMINCORRECTFEES() (string, error) {
	return _Errors.Contract.CMINCORRECTFEES(&_Errors.CallOpts)
}

// CMINCORRECTFEES is a free data retrieval call binding the contract method 0xde10ab9a.
//
// Solidity: function CM_INCORRECT_FEES() view returns(string)
func (_Errors *ErrorsCallerSession) CMINCORRECTFEES() (string, error) {
	return _Errors.Contract.CMINCORRECTFEES(&_Errors.CallOpts)
}

// CMINCORRECTNEWOWNER is a free data retrieval call binding the contract method 0xe13bde4c.
//
// Solidity: function CM_INCORRECT_NEW_OWNER() view returns(string)
func (_Errors *ErrorsCaller) CMINCORRECTNEWOWNER(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "CM_INCORRECT_NEW_OWNER")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CMINCORRECTNEWOWNER is a free data retrieval call binding the contract method 0xe13bde4c.
//
// Solidity: function CM_INCORRECT_NEW_OWNER() view returns(string)
func (_Errors *ErrorsSession) CMINCORRECTNEWOWNER() (string, error) {
	return _Errors.Contract.CMINCORRECTNEWOWNER(&_Errors.CallOpts)
}

// CMINCORRECTNEWOWNER is a free data retrieval call binding the contract method 0xe13bde4c.
//
// Solidity: function CM_INCORRECT_NEW_OWNER() view returns(string)
func (_Errors *ErrorsCallerSession) CMINCORRECTNEWOWNER() (string, error) {
	return _Errors.Contract.CMINCORRECTNEWOWNER(&_Errors.CallOpts)
}

// CMINCORRECTPARAMS is a free data retrieval call binding the contract method 0x2325ac93.
//
// Solidity: function CM_INCORRECT_PARAMS() view returns(string)
func (_Errors *ErrorsCaller) CMINCORRECTPARAMS(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "CM_INCORRECT_PARAMS")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CMINCORRECTPARAMS is a free data retrieval call binding the contract method 0x2325ac93.
//
// Solidity: function CM_INCORRECT_PARAMS() view returns(string)
func (_Errors *ErrorsSession) CMINCORRECTPARAMS() (string, error) {
	return _Errors.Contract.CMINCORRECTPARAMS(&_Errors.CallOpts)
}

// CMINCORRECTPARAMS is a free data retrieval call binding the contract method 0x2325ac93.
//
// Solidity: function CM_INCORRECT_PARAMS() view returns(string)
func (_Errors *ErrorsCallerSession) CMINCORRECTPARAMS() (string, error) {
	return _Errors.Contract.CMINCORRECTPARAMS(&_Errors.CallOpts)
}

// CMMAXLEVERAGEISTOOHIGH is a free data retrieval call binding the contract method 0xb5cdcbac.
//
// Solidity: function CM_MAX_LEVERAGE_IS_TOO_HIGH() view returns(string)
func (_Errors *ErrorsCaller) CMMAXLEVERAGEISTOOHIGH(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "CM_MAX_LEVERAGE_IS_TOO_HIGH")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CMMAXLEVERAGEISTOOHIGH is a free data retrieval call binding the contract method 0xb5cdcbac.
//
// Solidity: function CM_MAX_LEVERAGE_IS_TOO_HIGH() view returns(string)
func (_Errors *ErrorsSession) CMMAXLEVERAGEISTOOHIGH() (string, error) {
	return _Errors.Contract.CMMAXLEVERAGEISTOOHIGH(&_Errors.CallOpts)
}

// CMMAXLEVERAGEISTOOHIGH is a free data retrieval call binding the contract method 0xb5cdcbac.
//
// Solidity: function CM_MAX_LEVERAGE_IS_TOO_HIGH() view returns(string)
func (_Errors *ErrorsCallerSession) CMMAXLEVERAGEISTOOHIGH() (string, error) {
	return _Errors.Contract.CMMAXLEVERAGEISTOOHIGH(&_Errors.CallOpts)
}

// CMNOOPENACCOUNT is a free data retrieval call binding the contract method 0x428cf705.
//
// Solidity: function CM_NO_OPEN_ACCOUNT() view returns(string)
func (_Errors *ErrorsCaller) CMNOOPENACCOUNT(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "CM_NO_OPEN_ACCOUNT")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CMNOOPENACCOUNT is a free data retrieval call binding the contract method 0x428cf705.
//
// Solidity: function CM_NO_OPEN_ACCOUNT() view returns(string)
func (_Errors *ErrorsSession) CMNOOPENACCOUNT() (string, error) {
	return _Errors.Contract.CMNOOPENACCOUNT(&_Errors.CallOpts)
}

// CMNOOPENACCOUNT is a free data retrieval call binding the contract method 0x428cf705.
//
// Solidity: function CM_NO_OPEN_ACCOUNT() view returns(string)
func (_Errors *ErrorsCallerSession) CMNOOPENACCOUNT() (string, error) {
	return _Errors.Contract.CMNOOPENACCOUNT(&_Errors.CallOpts)
}

// CMTARGETCONTRACTISNOTALLOWED is a free data retrieval call binding the contract method 0x61b2ef2b.
//
// Solidity: function CM_TARGET_CONTRACT_iS_NOT_ALLOWED() view returns(string)
func (_Errors *ErrorsCaller) CMTARGETCONTRACTISNOTALLOWED(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "CM_TARGET_CONTRACT_iS_NOT_ALLOWED")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CMTARGETCONTRACTISNOTALLOWED is a free data retrieval call binding the contract method 0x61b2ef2b.
//
// Solidity: function CM_TARGET_CONTRACT_iS_NOT_ALLOWED() view returns(string)
func (_Errors *ErrorsSession) CMTARGETCONTRACTISNOTALLOWED() (string, error) {
	return _Errors.Contract.CMTARGETCONTRACTISNOTALLOWED(&_Errors.CallOpts)
}

// CMTARGETCONTRACTISNOTALLOWED is a free data retrieval call binding the contract method 0x61b2ef2b.
//
// Solidity: function CM_TARGET_CONTRACT_iS_NOT_ALLOWED() view returns(string)
func (_Errors *ErrorsCallerSession) CMTARGETCONTRACTISNOTALLOWED() (string, error) {
	return _Errors.Contract.CMTARGETCONTRACTISNOTALLOWED(&_Errors.CallOpts)
}

// CMTRANSFERFAILED is a free data retrieval call binding the contract method 0xb28bfe99.
//
// Solidity: function CM_TRANSFER_FAILED() view returns(string)
func (_Errors *ErrorsCaller) CMTRANSFERFAILED(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "CM_TRANSFER_FAILED")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CMTRANSFERFAILED is a free data retrieval call binding the contract method 0xb28bfe99.
//
// Solidity: function CM_TRANSFER_FAILED() view returns(string)
func (_Errors *ErrorsSession) CMTRANSFERFAILED() (string, error) {
	return _Errors.Contract.CMTRANSFERFAILED(&_Errors.CallOpts)
}

// CMTRANSFERFAILED is a free data retrieval call binding the contract method 0xb28bfe99.
//
// Solidity: function CM_TRANSFER_FAILED() view returns(string)
func (_Errors *ErrorsCallerSession) CMTRANSFERFAILED() (string, error) {
	return _Errors.Contract.CMTRANSFERFAILED(&_Errors.CallOpts)
}

// CMWETHGATEWAYONLY is a free data retrieval call binding the contract method 0x944f5d2a.
//
// Solidity: function CM_WETH_GATEWAY_ONLY() view returns(string)
func (_Errors *ErrorsCaller) CMWETHGATEWAYONLY(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "CM_WETH_GATEWAY_ONLY")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CMWETHGATEWAYONLY is a free data retrieval call binding the contract method 0x944f5d2a.
//
// Solidity: function CM_WETH_GATEWAY_ONLY() view returns(string)
func (_Errors *ErrorsSession) CMWETHGATEWAYONLY() (string, error) {
	return _Errors.Contract.CMWETHGATEWAYONLY(&_Errors.CallOpts)
}

// CMWETHGATEWAYONLY is a free data retrieval call binding the contract method 0x944f5d2a.
//
// Solidity: function CM_WETH_GATEWAY_ONLY() view returns(string)
func (_Errors *ErrorsCallerSession) CMWETHGATEWAYONLY() (string, error) {
	return _Errors.Contract.CMWETHGATEWAYONLY(&_Errors.CallOpts)
}

// CMZEROADDRESSORUSERHAVEALREADYOPENCREDITACCOUNT is a free data retrieval call binding the contract method 0x47985f2b.
//
// Solidity: function CM_ZERO_ADDRESS_OR_USER_HAVE_ALREADY_OPEN_CREDIT_ACCOUNT() view returns(string)
func (_Errors *ErrorsCaller) CMZEROADDRESSORUSERHAVEALREADYOPENCREDITACCOUNT(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "CM_ZERO_ADDRESS_OR_USER_HAVE_ALREADY_OPEN_CREDIT_ACCOUNT")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CMZEROADDRESSORUSERHAVEALREADYOPENCREDITACCOUNT is a free data retrieval call binding the contract method 0x47985f2b.
//
// Solidity: function CM_ZERO_ADDRESS_OR_USER_HAVE_ALREADY_OPEN_CREDIT_ACCOUNT() view returns(string)
func (_Errors *ErrorsSession) CMZEROADDRESSORUSERHAVEALREADYOPENCREDITACCOUNT() (string, error) {
	return _Errors.Contract.CMZEROADDRESSORUSERHAVEALREADYOPENCREDITACCOUNT(&_Errors.CallOpts)
}

// CMZEROADDRESSORUSERHAVEALREADYOPENCREDITACCOUNT is a free data retrieval call binding the contract method 0x47985f2b.
//
// Solidity: function CM_ZERO_ADDRESS_OR_USER_HAVE_ALREADY_OPEN_CREDIT_ACCOUNT() view returns(string)
func (_Errors *ErrorsCallerSession) CMZEROADDRESSORUSERHAVEALREADYOPENCREDITACCOUNT() (string, error) {
	return _Errors.Contract.CMZEROADDRESSORUSERHAVEALREADYOPENCREDITACCOUNT(&_Errors.CallOpts)
}

// CRCREDITMANAGERALREADYADDED is a free data retrieval call binding the contract method 0x99a98c99.
//
// Solidity: function CR_CREDIT_MANAGER_ALREADY_ADDED() view returns(string)
func (_Errors *ErrorsCaller) CRCREDITMANAGERALREADYADDED(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "CR_CREDIT_MANAGER_ALREADY_ADDED")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CRCREDITMANAGERALREADYADDED is a free data retrieval call binding the contract method 0x99a98c99.
//
// Solidity: function CR_CREDIT_MANAGER_ALREADY_ADDED() view returns(string)
func (_Errors *ErrorsSession) CRCREDITMANAGERALREADYADDED() (string, error) {
	return _Errors.Contract.CRCREDITMANAGERALREADYADDED(&_Errors.CallOpts)
}

// CRCREDITMANAGERALREADYADDED is a free data retrieval call binding the contract method 0x99a98c99.
//
// Solidity: function CR_CREDIT_MANAGER_ALREADY_ADDED() view returns(string)
func (_Errors *ErrorsCallerSession) CRCREDITMANAGERALREADYADDED() (string, error) {
	return _Errors.Contract.CRCREDITMANAGERALREADYADDED(&_Errors.CallOpts)
}

// CRPOOLALREADYADDED is a free data retrieval call binding the contract method 0x0a2b1d3a.
//
// Solidity: function CR_POOL_ALREADY_ADDED() view returns(string)
func (_Errors *ErrorsCaller) CRPOOLALREADYADDED(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "CR_POOL_ALREADY_ADDED")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CRPOOLALREADYADDED is a free data retrieval call binding the contract method 0x0a2b1d3a.
//
// Solidity: function CR_POOL_ALREADY_ADDED() view returns(string)
func (_Errors *ErrorsSession) CRPOOLALREADYADDED() (string, error) {
	return _Errors.Contract.CRPOOLALREADYADDED(&_Errors.CallOpts)
}

// CRPOOLALREADYADDED is a free data retrieval call binding the contract method 0x0a2b1d3a.
//
// Solidity: function CR_POOL_ALREADY_ADDED() view returns(string)
func (_Errors *ErrorsCallerSession) CRPOOLALREADYADDED() (string, error) {
	return _Errors.Contract.CRPOOLALREADYADDED(&_Errors.CallOpts)
}

// INCORRECTARRAYLENGTH is a free data retrieval call binding the contract method 0xe7f3be0c.
//
// Solidity: function INCORRECT_ARRAY_LENGTH() view returns(string)
func (_Errors *ErrorsCaller) INCORRECTARRAYLENGTH(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "INCORRECT_ARRAY_LENGTH")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// INCORRECTARRAYLENGTH is a free data retrieval call binding the contract method 0xe7f3be0c.
//
// Solidity: function INCORRECT_ARRAY_LENGTH() view returns(string)
func (_Errors *ErrorsSession) INCORRECTARRAYLENGTH() (string, error) {
	return _Errors.Contract.INCORRECTARRAYLENGTH(&_Errors.CallOpts)
}

// INCORRECTARRAYLENGTH is a free data retrieval call binding the contract method 0xe7f3be0c.
//
// Solidity: function INCORRECT_ARRAY_LENGTH() view returns(string)
func (_Errors *ErrorsCallerSession) INCORRECTARRAYLENGTH() (string, error) {
	return _Errors.Contract.INCORRECTARRAYLENGTH(&_Errors.CallOpts)
}

// INCORRECTPARAMETER is a free data retrieval call binding the contract method 0xbdcb2576.
//
// Solidity: function INCORRECT_PARAMETER() view returns(string)
func (_Errors *ErrorsCaller) INCORRECTPARAMETER(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "INCORRECT_PARAMETER")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// INCORRECTPARAMETER is a free data retrieval call binding the contract method 0xbdcb2576.
//
// Solidity: function INCORRECT_PARAMETER() view returns(string)
func (_Errors *ErrorsSession) INCORRECTPARAMETER() (string, error) {
	return _Errors.Contract.INCORRECTPARAMETER(&_Errors.CallOpts)
}

// INCORRECTPARAMETER is a free data retrieval call binding the contract method 0xbdcb2576.
//
// Solidity: function INCORRECT_PARAMETER() view returns(string)
func (_Errors *ErrorsCallerSession) INCORRECTPARAMETER() (string, error) {
	return _Errors.Contract.INCORRECTPARAMETER(&_Errors.CallOpts)
}

// INCORRECTPATHLENGTH is a free data retrieval call binding the contract method 0x3df46fe5.
//
// Solidity: function INCORRECT_PATH_LENGTH() view returns(string)
func (_Errors *ErrorsCaller) INCORRECTPATHLENGTH(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "INCORRECT_PATH_LENGTH")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// INCORRECTPATHLENGTH is a free data retrieval call binding the contract method 0x3df46fe5.
//
// Solidity: function INCORRECT_PATH_LENGTH() view returns(string)
func (_Errors *ErrorsSession) INCORRECTPATHLENGTH() (string, error) {
	return _Errors.Contract.INCORRECTPATHLENGTH(&_Errors.CallOpts)
}

// INCORRECTPATHLENGTH is a free data retrieval call binding the contract method 0x3df46fe5.
//
// Solidity: function INCORRECT_PATH_LENGTH() view returns(string)
func (_Errors *ErrorsCallerSession) INCORRECTPATHLENGTH() (string, error) {
	return _Errors.Contract.INCORRECTPATHLENGTH(&_Errors.CallOpts)
}

// LAHASVALUEWITHTOKENTRANSFER is a free data retrieval call binding the contract method 0x5fd6824d.
//
// Solidity: function LA_HAS_VALUE_WITH_TOKEN_TRANSFER() view returns(string)
func (_Errors *ErrorsCaller) LAHASVALUEWITHTOKENTRANSFER(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "LA_HAS_VALUE_WITH_TOKEN_TRANSFER")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// LAHASVALUEWITHTOKENTRANSFER is a free data retrieval call binding the contract method 0x5fd6824d.
//
// Solidity: function LA_HAS_VALUE_WITH_TOKEN_TRANSFER() view returns(string)
func (_Errors *ErrorsSession) LAHASVALUEWITHTOKENTRANSFER() (string, error) {
	return _Errors.Contract.LAHASVALUEWITHTOKENTRANSFER(&_Errors.CallOpts)
}

// LAHASVALUEWITHTOKENTRANSFER is a free data retrieval call binding the contract method 0x5fd6824d.
//
// Solidity: function LA_HAS_VALUE_WITH_TOKEN_TRANSFER() view returns(string)
func (_Errors *ErrorsCallerSession) LAHASVALUEWITHTOKENTRANSFER() (string, error) {
	return _Errors.Contract.LAHASVALUEWITHTOKENTRANSFER(&_Errors.CallOpts)
}

// LAINCORRECTVALUE is a free data retrieval call binding the contract method 0x9f4f3eeb.
//
// Solidity: function LA_INCORRECT_VALUE() view returns(string)
func (_Errors *ErrorsCaller) LAINCORRECTVALUE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "LA_INCORRECT_VALUE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// LAINCORRECTVALUE is a free data retrieval call binding the contract method 0x9f4f3eeb.
//
// Solidity: function LA_INCORRECT_VALUE() view returns(string)
func (_Errors *ErrorsSession) LAINCORRECTVALUE() (string, error) {
	return _Errors.Contract.LAINCORRECTVALUE(&_Errors.CallOpts)
}

// LAINCORRECTVALUE is a free data retrieval call binding the contract method 0x9f4f3eeb.
//
// Solidity: function LA_INCORRECT_VALUE() view returns(string)
func (_Errors *ErrorsCallerSession) LAINCORRECTVALUE() (string, error) {
	return _Errors.Contract.LAINCORRECTVALUE(&_Errors.CallOpts)
}

// LALOWERTHANAMOUNTMIN is a free data retrieval call binding the contract method 0x159eee9f.
//
// Solidity: function LA_LOWER_THAN_AMOUNT_MIN() view returns(string)
func (_Errors *ErrorsCaller) LALOWERTHANAMOUNTMIN(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "LA_LOWER_THAN_AMOUNT_MIN")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// LALOWERTHANAMOUNTMIN is a free data retrieval call binding the contract method 0x159eee9f.
//
// Solidity: function LA_LOWER_THAN_AMOUNT_MIN() view returns(string)
func (_Errors *ErrorsSession) LALOWERTHANAMOUNTMIN() (string, error) {
	return _Errors.Contract.LALOWERTHANAMOUNTMIN(&_Errors.CallOpts)
}

// LALOWERTHANAMOUNTMIN is a free data retrieval call binding the contract method 0x159eee9f.
//
// Solidity: function LA_LOWER_THAN_AMOUNT_MIN() view returns(string)
func (_Errors *ErrorsCallerSession) LALOWERTHANAMOUNTMIN() (string, error) {
	return _Errors.Contract.LALOWERTHANAMOUNTMIN(&_Errors.CallOpts)
}

// LATOKENOUTISNOTCOLLATERAL is a free data retrieval call binding the contract method 0x61fb3dc4.
//
// Solidity: function LA_TOKEN_OUT_IS_NOT_COLLATERAL() view returns(string)
func (_Errors *ErrorsCaller) LATOKENOUTISNOTCOLLATERAL(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "LA_TOKEN_OUT_IS_NOT_COLLATERAL")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// LATOKENOUTISNOTCOLLATERAL is a free data retrieval call binding the contract method 0x61fb3dc4.
//
// Solidity: function LA_TOKEN_OUT_IS_NOT_COLLATERAL() view returns(string)
func (_Errors *ErrorsSession) LATOKENOUTISNOTCOLLATERAL() (string, error) {
	return _Errors.Contract.LATOKENOUTISNOTCOLLATERAL(&_Errors.CallOpts)
}

// LATOKENOUTISNOTCOLLATERAL is a free data retrieval call binding the contract method 0x61fb3dc4.
//
// Solidity: function LA_TOKEN_OUT_IS_NOT_COLLATERAL() view returns(string)
func (_Errors *ErrorsCallerSession) LATOKENOUTISNOTCOLLATERAL() (string, error) {
	return _Errors.Contract.LATOKENOUTISNOTCOLLATERAL(&_Errors.CallOpts)
}

// LAUNKNOWNLPINTERFACE is a free data retrieval call binding the contract method 0x52be8a64.
//
// Solidity: function LA_UNKNOWN_LP_INTERFACE() view returns(string)
func (_Errors *ErrorsCaller) LAUNKNOWNLPINTERFACE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "LA_UNKNOWN_LP_INTERFACE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// LAUNKNOWNLPINTERFACE is a free data retrieval call binding the contract method 0x52be8a64.
//
// Solidity: function LA_UNKNOWN_LP_INTERFACE() view returns(string)
func (_Errors *ErrorsSession) LAUNKNOWNLPINTERFACE() (string, error) {
	return _Errors.Contract.LAUNKNOWNLPINTERFACE(&_Errors.CallOpts)
}

// LAUNKNOWNLPINTERFACE is a free data retrieval call binding the contract method 0x52be8a64.
//
// Solidity: function LA_UNKNOWN_LP_INTERFACE() view returns(string)
func (_Errors *ErrorsCallerSession) LAUNKNOWNLPINTERFACE() (string, error) {
	return _Errors.Contract.LAUNKNOWNLPINTERFACE(&_Errors.CallOpts)
}

// LAUNKNOWNSWAPINTERFACE is a free data retrieval call binding the contract method 0x012f8222.
//
// Solidity: function LA_UNKNOWN_SWAP_INTERFACE() view returns(string)
func (_Errors *ErrorsCaller) LAUNKNOWNSWAPINTERFACE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "LA_UNKNOWN_SWAP_INTERFACE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// LAUNKNOWNSWAPINTERFACE is a free data retrieval call binding the contract method 0x012f8222.
//
// Solidity: function LA_UNKNOWN_SWAP_INTERFACE() view returns(string)
func (_Errors *ErrorsSession) LAUNKNOWNSWAPINTERFACE() (string, error) {
	return _Errors.Contract.LAUNKNOWNSWAPINTERFACE(&_Errors.CallOpts)
}

// LAUNKNOWNSWAPINTERFACE is a free data retrieval call binding the contract method 0x012f8222.
//
// Solidity: function LA_UNKNOWN_SWAP_INTERFACE() view returns(string)
func (_Errors *ErrorsCallerSession) LAUNKNOWNSWAPINTERFACE() (string, error) {
	return _Errors.Contract.LAUNKNOWNSWAPINTERFACE(&_Errors.CallOpts)
}

// MATHADDITIONOVERFLOW is a free data retrieval call binding the contract method 0x0f5ee482.
//
// Solidity: function MATH_ADDITION_OVERFLOW() view returns(string)
func (_Errors *ErrorsCaller) MATHADDITIONOVERFLOW(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "MATH_ADDITION_OVERFLOW")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// MATHADDITIONOVERFLOW is a free data retrieval call binding the contract method 0x0f5ee482.
//
// Solidity: function MATH_ADDITION_OVERFLOW() view returns(string)
func (_Errors *ErrorsSession) MATHADDITIONOVERFLOW() (string, error) {
	return _Errors.Contract.MATHADDITIONOVERFLOW(&_Errors.CallOpts)
}

// MATHADDITIONOVERFLOW is a free data retrieval call binding the contract method 0x0f5ee482.
//
// Solidity: function MATH_ADDITION_OVERFLOW() view returns(string)
func (_Errors *ErrorsCallerSession) MATHADDITIONOVERFLOW() (string, error) {
	return _Errors.Contract.MATHADDITIONOVERFLOW(&_Errors.CallOpts)
}

// MATHDIVISIONBYZERO is a free data retrieval call binding the contract method 0x4349e3d8.
//
// Solidity: function MATH_DIVISION_BY_ZERO() view returns(string)
func (_Errors *ErrorsCaller) MATHDIVISIONBYZERO(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "MATH_DIVISION_BY_ZERO")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// MATHDIVISIONBYZERO is a free data retrieval call binding the contract method 0x4349e3d8.
//
// Solidity: function MATH_DIVISION_BY_ZERO() view returns(string)
func (_Errors *ErrorsSession) MATHDIVISIONBYZERO() (string, error) {
	return _Errors.Contract.MATHDIVISIONBYZERO(&_Errors.CallOpts)
}

// MATHDIVISIONBYZERO is a free data retrieval call binding the contract method 0x4349e3d8.
//
// Solidity: function MATH_DIVISION_BY_ZERO() view returns(string)
func (_Errors *ErrorsCallerSession) MATHDIVISIONBYZERO() (string, error) {
	return _Errors.Contract.MATHDIVISIONBYZERO(&_Errors.CallOpts)
}

// MATHMULTIPLICATIONOVERFLOW is a free data retrieval call binding the contract method 0x029d2344.
//
// Solidity: function MATH_MULTIPLICATION_OVERFLOW() view returns(string)
func (_Errors *ErrorsCaller) MATHMULTIPLICATIONOVERFLOW(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "MATH_MULTIPLICATION_OVERFLOW")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// MATHMULTIPLICATIONOVERFLOW is a free data retrieval call binding the contract method 0x029d2344.
//
// Solidity: function MATH_MULTIPLICATION_OVERFLOW() view returns(string)
func (_Errors *ErrorsSession) MATHMULTIPLICATIONOVERFLOW() (string, error) {
	return _Errors.Contract.MATHMULTIPLICATIONOVERFLOW(&_Errors.CallOpts)
}

// MATHMULTIPLICATIONOVERFLOW is a free data retrieval call binding the contract method 0x029d2344.
//
// Solidity: function MATH_MULTIPLICATION_OVERFLOW() view returns(string)
func (_Errors *ErrorsCallerSession) MATHMULTIPLICATIONOVERFLOW() (string, error) {
	return _Errors.Contract.MATHMULTIPLICATIONOVERFLOW(&_Errors.CallOpts)
}

// NOTIMPLEMENTED is a free data retrieval call binding the contract method 0x43f6e4ab.
//
// Solidity: function NOT_IMPLEMENTED() view returns(string)
func (_Errors *ErrorsCaller) NOTIMPLEMENTED(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "NOT_IMPLEMENTED")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NOTIMPLEMENTED is a free data retrieval call binding the contract method 0x43f6e4ab.
//
// Solidity: function NOT_IMPLEMENTED() view returns(string)
func (_Errors *ErrorsSession) NOTIMPLEMENTED() (string, error) {
	return _Errors.Contract.NOTIMPLEMENTED(&_Errors.CallOpts)
}

// NOTIMPLEMENTED is a free data retrieval call binding the contract method 0x43f6e4ab.
//
// Solidity: function NOT_IMPLEMENTED() view returns(string)
func (_Errors *ErrorsCallerSession) NOTIMPLEMENTED() (string, error) {
	return _Errors.Contract.NOTIMPLEMENTED(&_Errors.CallOpts)
}

// POOLCANTADDCREDITMANAGERTWICE is a free data retrieval call binding the contract method 0xccbf9278.
//
// Solidity: function POOL_CANT_ADD_CREDIT_MANAGER_TWICE() view returns(string)
func (_Errors *ErrorsCaller) POOLCANTADDCREDITMANAGERTWICE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "POOL_CANT_ADD_CREDIT_MANAGER_TWICE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// POOLCANTADDCREDITMANAGERTWICE is a free data retrieval call binding the contract method 0xccbf9278.
//
// Solidity: function POOL_CANT_ADD_CREDIT_MANAGER_TWICE() view returns(string)
func (_Errors *ErrorsSession) POOLCANTADDCREDITMANAGERTWICE() (string, error) {
	return _Errors.Contract.POOLCANTADDCREDITMANAGERTWICE(&_Errors.CallOpts)
}

// POOLCANTADDCREDITMANAGERTWICE is a free data retrieval call binding the contract method 0xccbf9278.
//
// Solidity: function POOL_CANT_ADD_CREDIT_MANAGER_TWICE() view returns(string)
func (_Errors *ErrorsCallerSession) POOLCANTADDCREDITMANAGERTWICE() (string, error) {
	return _Errors.Contract.POOLCANTADDCREDITMANAGERTWICE(&_Errors.CallOpts)
}

// POOLCONNECTEDCREDITMANAGERSONLY is a free data retrieval call binding the contract method 0x76d9ebb8.
//
// Solidity: function POOL_CONNECTED_CREDIT_MANAGERS_ONLY() view returns(string)
func (_Errors *ErrorsCaller) POOLCONNECTEDCREDITMANAGERSONLY(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "POOL_CONNECTED_CREDIT_MANAGERS_ONLY")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// POOLCONNECTEDCREDITMANAGERSONLY is a free data retrieval call binding the contract method 0x76d9ebb8.
//
// Solidity: function POOL_CONNECTED_CREDIT_MANAGERS_ONLY() view returns(string)
func (_Errors *ErrorsSession) POOLCONNECTEDCREDITMANAGERSONLY() (string, error) {
	return _Errors.Contract.POOLCONNECTEDCREDITMANAGERSONLY(&_Errors.CallOpts)
}

// POOLCONNECTEDCREDITMANAGERSONLY is a free data retrieval call binding the contract method 0x76d9ebb8.
//
// Solidity: function POOL_CONNECTED_CREDIT_MANAGERS_ONLY() view returns(string)
func (_Errors *ErrorsCallerSession) POOLCONNECTEDCREDITMANAGERSONLY() (string, error) {
	return _Errors.Contract.POOLCONNECTEDCREDITMANAGERSONLY(&_Errors.CallOpts)
}

// POOLINCOMPATIBLECREDITACCOUNTMANAGER is a free data retrieval call binding the contract method 0xa27c0370.
//
// Solidity: function POOL_INCOMPATIBLE_CREDIT_ACCOUNT_MANAGER() view returns(string)
func (_Errors *ErrorsCaller) POOLINCOMPATIBLECREDITACCOUNTMANAGER(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "POOL_INCOMPATIBLE_CREDIT_ACCOUNT_MANAGER")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// POOLINCOMPATIBLECREDITACCOUNTMANAGER is a free data retrieval call binding the contract method 0xa27c0370.
//
// Solidity: function POOL_INCOMPATIBLE_CREDIT_ACCOUNT_MANAGER() view returns(string)
func (_Errors *ErrorsSession) POOLINCOMPATIBLECREDITACCOUNTMANAGER() (string, error) {
	return _Errors.Contract.POOLINCOMPATIBLECREDITACCOUNTMANAGER(&_Errors.CallOpts)
}

// POOLINCOMPATIBLECREDITACCOUNTMANAGER is a free data retrieval call binding the contract method 0xa27c0370.
//
// Solidity: function POOL_INCOMPATIBLE_CREDIT_ACCOUNT_MANAGER() view returns(string)
func (_Errors *ErrorsCallerSession) POOLINCOMPATIBLECREDITACCOUNTMANAGER() (string, error) {
	return _Errors.Contract.POOLINCOMPATIBLECREDITACCOUNTMANAGER(&_Errors.CallOpts)
}

// POOLINCORRECTWITHDRAWFEE is a free data retrieval call binding the contract method 0x28432c22.
//
// Solidity: function POOL_INCORRECT_WITHDRAW_FEE() view returns(string)
func (_Errors *ErrorsCaller) POOLINCORRECTWITHDRAWFEE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "POOL_INCORRECT_WITHDRAW_FEE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// POOLINCORRECTWITHDRAWFEE is a free data retrieval call binding the contract method 0x28432c22.
//
// Solidity: function POOL_INCORRECT_WITHDRAW_FEE() view returns(string)
func (_Errors *ErrorsSession) POOLINCORRECTWITHDRAWFEE() (string, error) {
	return _Errors.Contract.POOLINCORRECTWITHDRAWFEE(&_Errors.CallOpts)
}

// POOLINCORRECTWITHDRAWFEE is a free data retrieval call binding the contract method 0x28432c22.
//
// Solidity: function POOL_INCORRECT_WITHDRAW_FEE() view returns(string)
func (_Errors *ErrorsCallerSession) POOLINCORRECTWITHDRAWFEE() (string, error) {
	return _Errors.Contract.POOLINCORRECTWITHDRAWFEE(&_Errors.CallOpts)
}

// POOLMORETHANEXPECTEDLIQUIDITYLIMIT is a free data retrieval call binding the contract method 0x119427c5.
//
// Solidity: function POOL_MORE_THAN_EXPECTED_LIQUIDITY_LIMIT() view returns(string)
func (_Errors *ErrorsCaller) POOLMORETHANEXPECTEDLIQUIDITYLIMIT(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "POOL_MORE_THAN_EXPECTED_LIQUIDITY_LIMIT")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// POOLMORETHANEXPECTEDLIQUIDITYLIMIT is a free data retrieval call binding the contract method 0x119427c5.
//
// Solidity: function POOL_MORE_THAN_EXPECTED_LIQUIDITY_LIMIT() view returns(string)
func (_Errors *ErrorsSession) POOLMORETHANEXPECTEDLIQUIDITYLIMIT() (string, error) {
	return _Errors.Contract.POOLMORETHANEXPECTEDLIQUIDITYLIMIT(&_Errors.CallOpts)
}

// POOLMORETHANEXPECTEDLIQUIDITYLIMIT is a free data retrieval call binding the contract method 0x119427c5.
//
// Solidity: function POOL_MORE_THAN_EXPECTED_LIQUIDITY_LIMIT() view returns(string)
func (_Errors *ErrorsCallerSession) POOLMORETHANEXPECTEDLIQUIDITYLIMIT() (string, error) {
	return _Errors.Contract.POOLMORETHANEXPECTEDLIQUIDITYLIMIT(&_Errors.CallOpts)
}

// POAGGREGATORDECIMALSSHOULDBE18 is a free data retrieval call binding the contract method 0xb4fcfee3.
//
// Solidity: function PO_AGGREGATOR_DECIMALS_SHOULD_BE_18() view returns(string)
func (_Errors *ErrorsCaller) POAGGREGATORDECIMALSSHOULDBE18(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "PO_AGGREGATOR_DECIMALS_SHOULD_BE_18")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// POAGGREGATORDECIMALSSHOULDBE18 is a free data retrieval call binding the contract method 0xb4fcfee3.
//
// Solidity: function PO_AGGREGATOR_DECIMALS_SHOULD_BE_18() view returns(string)
func (_Errors *ErrorsSession) POAGGREGATORDECIMALSSHOULDBE18() (string, error) {
	return _Errors.Contract.POAGGREGATORDECIMALSSHOULDBE18(&_Errors.CallOpts)
}

// POAGGREGATORDECIMALSSHOULDBE18 is a free data retrieval call binding the contract method 0xb4fcfee3.
//
// Solidity: function PO_AGGREGATOR_DECIMALS_SHOULD_BE_18() view returns(string)
func (_Errors *ErrorsCallerSession) POAGGREGATORDECIMALSSHOULDBE18() (string, error) {
	return _Errors.Contract.POAGGREGATORDECIMALSSHOULDBE18(&_Errors.CallOpts)
}

// POPRICEFEEDDOESNTEXIST is a free data retrieval call binding the contract method 0x2c944814.
//
// Solidity: function PO_PRICE_FEED_DOESNT_EXIST() view returns(string)
func (_Errors *ErrorsCaller) POPRICEFEEDDOESNTEXIST(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "PO_PRICE_FEED_DOESNT_EXIST")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// POPRICEFEEDDOESNTEXIST is a free data retrieval call binding the contract method 0x2c944814.
//
// Solidity: function PO_PRICE_FEED_DOESNT_EXIST() view returns(string)
func (_Errors *ErrorsSession) POPRICEFEEDDOESNTEXIST() (string, error) {
	return _Errors.Contract.POPRICEFEEDDOESNTEXIST(&_Errors.CallOpts)
}

// POPRICEFEEDDOESNTEXIST is a free data retrieval call binding the contract method 0x2c944814.
//
// Solidity: function PO_PRICE_FEED_DOESNT_EXIST() view returns(string)
func (_Errors *ErrorsCallerSession) POPRICEFEEDDOESNTEXIST() (string, error) {
	return _Errors.Contract.POPRICEFEEDDOESNTEXIST(&_Errors.CallOpts)
}

// POTOKENSWITHDECIMALSMORE18ISNTALLOWED is a free data retrieval call binding the contract method 0x8dcf3184.
//
// Solidity: function PO_TOKENS_WITH_DECIMALS_MORE_18_ISNT_ALLOWED() view returns(string)
func (_Errors *ErrorsCaller) POTOKENSWITHDECIMALSMORE18ISNTALLOWED(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "PO_TOKENS_WITH_DECIMALS_MORE_18_ISNT_ALLOWED")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// POTOKENSWITHDECIMALSMORE18ISNTALLOWED is a free data retrieval call binding the contract method 0x8dcf3184.
//
// Solidity: function PO_TOKENS_WITH_DECIMALS_MORE_18_ISNT_ALLOWED() view returns(string)
func (_Errors *ErrorsSession) POTOKENSWITHDECIMALSMORE18ISNTALLOWED() (string, error) {
	return _Errors.Contract.POTOKENSWITHDECIMALSMORE18ISNTALLOWED(&_Errors.CallOpts)
}

// POTOKENSWITHDECIMALSMORE18ISNTALLOWED is a free data retrieval call binding the contract method 0x8dcf3184.
//
// Solidity: function PO_TOKENS_WITH_DECIMALS_MORE_18_ISNT_ALLOWED() view returns(string)
func (_Errors *ErrorsCallerSession) POTOKENSWITHDECIMALSMORE18ISNTALLOWED() (string, error) {
	return _Errors.Contract.POTOKENSWITHDECIMALSMORE18ISNTALLOWED(&_Errors.CallOpts)
}

// REGISTEREDCREDITACCOUNTMANAGERSONLY is a free data retrieval call binding the contract method 0x94391a4a.
//
// Solidity: function REGISTERED_CREDIT_ACCOUNT_MANAGERS_ONLY() view returns(string)
func (_Errors *ErrorsCaller) REGISTEREDCREDITACCOUNTMANAGERSONLY(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "REGISTERED_CREDIT_ACCOUNT_MANAGERS_ONLY")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// REGISTEREDCREDITACCOUNTMANAGERSONLY is a free data retrieval call binding the contract method 0x94391a4a.
//
// Solidity: function REGISTERED_CREDIT_ACCOUNT_MANAGERS_ONLY() view returns(string)
func (_Errors *ErrorsSession) REGISTEREDCREDITACCOUNTMANAGERSONLY() (string, error) {
	return _Errors.Contract.REGISTEREDCREDITACCOUNTMANAGERSONLY(&_Errors.CallOpts)
}

// REGISTEREDCREDITACCOUNTMANAGERSONLY is a free data retrieval call binding the contract method 0x94391a4a.
//
// Solidity: function REGISTERED_CREDIT_ACCOUNT_MANAGERS_ONLY() view returns(string)
func (_Errors *ErrorsCallerSession) REGISTEREDCREDITACCOUNTMANAGERSONLY() (string, error) {
	return _Errors.Contract.REGISTEREDCREDITACCOUNTMANAGERSONLY(&_Errors.CallOpts)
}

// REGISTEREDPOOLSONLY is a free data retrieval call binding the contract method 0x0afeee97.
//
// Solidity: function REGISTERED_POOLS_ONLY() view returns(string)
func (_Errors *ErrorsCaller) REGISTEREDPOOLSONLY(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "REGISTERED_POOLS_ONLY")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// REGISTEREDPOOLSONLY is a free data retrieval call binding the contract method 0x0afeee97.
//
// Solidity: function REGISTERED_POOLS_ONLY() view returns(string)
func (_Errors *ErrorsSession) REGISTEREDPOOLSONLY() (string, error) {
	return _Errors.Contract.REGISTEREDPOOLSONLY(&_Errors.CallOpts)
}

// REGISTEREDPOOLSONLY is a free data retrieval call binding the contract method 0x0afeee97.
//
// Solidity: function REGISTERED_POOLS_ONLY() view returns(string)
func (_Errors *ErrorsCallerSession) REGISTEREDPOOLSONLY() (string, error) {
	return _Errors.Contract.REGISTEREDPOOLSONLY(&_Errors.CallOpts)
}

// WGDESTINATIONISNOTWETHCOMPATIBLE is a free data retrieval call binding the contract method 0x3647c9f9.
//
// Solidity: function WG_DESTINATION_IS_NOT_WETH_COMPATIBLE() view returns(string)
func (_Errors *ErrorsCaller) WGDESTINATIONISNOTWETHCOMPATIBLE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "WG_DESTINATION_IS_NOT_WETH_COMPATIBLE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// WGDESTINATIONISNOTWETHCOMPATIBLE is a free data retrieval call binding the contract method 0x3647c9f9.
//
// Solidity: function WG_DESTINATION_IS_NOT_WETH_COMPATIBLE() view returns(string)
func (_Errors *ErrorsSession) WGDESTINATIONISNOTWETHCOMPATIBLE() (string, error) {
	return _Errors.Contract.WGDESTINATIONISNOTWETHCOMPATIBLE(&_Errors.CallOpts)
}

// WGDESTINATIONISNOTWETHCOMPATIBLE is a free data retrieval call binding the contract method 0x3647c9f9.
//
// Solidity: function WG_DESTINATION_IS_NOT_WETH_COMPATIBLE() view returns(string)
func (_Errors *ErrorsCallerSession) WGDESTINATIONISNOTWETHCOMPATIBLE() (string, error) {
	return _Errors.Contract.WGDESTINATIONISNOTWETHCOMPATIBLE(&_Errors.CallOpts)
}

// WGNOTENOUGHFUNDS is a free data retrieval call binding the contract method 0xbeea5ec2.
//
// Solidity: function WG_NOT_ENOUGH_FUNDS() view returns(string)
func (_Errors *ErrorsCaller) WGNOTENOUGHFUNDS(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "WG_NOT_ENOUGH_FUNDS")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// WGNOTENOUGHFUNDS is a free data retrieval call binding the contract method 0xbeea5ec2.
//
// Solidity: function WG_NOT_ENOUGH_FUNDS() view returns(string)
func (_Errors *ErrorsSession) WGNOTENOUGHFUNDS() (string, error) {
	return _Errors.Contract.WGNOTENOUGHFUNDS(&_Errors.CallOpts)
}

// WGNOTENOUGHFUNDS is a free data retrieval call binding the contract method 0xbeea5ec2.
//
// Solidity: function WG_NOT_ENOUGH_FUNDS() view returns(string)
func (_Errors *ErrorsCallerSession) WGNOTENOUGHFUNDS() (string, error) {
	return _Errors.Contract.WGNOTENOUGHFUNDS(&_Errors.CallOpts)
}

// WGRECEIVEISNOTALLOWED is a free data retrieval call binding the contract method 0x447d8e42.
//
// Solidity: function WG_RECEIVE_IS_NOT_ALLOWED() view returns(string)
func (_Errors *ErrorsCaller) WGRECEIVEISNOTALLOWED(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "WG_RECEIVE_IS_NOT_ALLOWED")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// WGRECEIVEISNOTALLOWED is a free data retrieval call binding the contract method 0x447d8e42.
//
// Solidity: function WG_RECEIVE_IS_NOT_ALLOWED() view returns(string)
func (_Errors *ErrorsSession) WGRECEIVEISNOTALLOWED() (string, error) {
	return _Errors.Contract.WGRECEIVEISNOTALLOWED(&_Errors.CallOpts)
}

// WGRECEIVEISNOTALLOWED is a free data retrieval call binding the contract method 0x447d8e42.
//
// Solidity: function WG_RECEIVE_IS_NOT_ALLOWED() view returns(string)
func (_Errors *ErrorsCallerSession) WGRECEIVEISNOTALLOWED() (string, error) {
	return _Errors.Contract.WGRECEIVEISNOTALLOWED(&_Errors.CallOpts)
}

// ZEROADDRESSISNOTALLOWED is a free data retrieval call binding the contract method 0x3f3153b2.
//
// Solidity: function ZERO_ADDRESS_IS_NOT_ALLOWED() view returns(string)
func (_Errors *ErrorsCaller) ZEROADDRESSISNOTALLOWED(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "ZERO_ADDRESS_IS_NOT_ALLOWED")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ZEROADDRESSISNOTALLOWED is a free data retrieval call binding the contract method 0x3f3153b2.
//
// Solidity: function ZERO_ADDRESS_IS_NOT_ALLOWED() view returns(string)
func (_Errors *ErrorsSession) ZEROADDRESSISNOTALLOWED() (string, error) {
	return _Errors.Contract.ZEROADDRESSISNOTALLOWED(&_Errors.CallOpts)
}

// ZEROADDRESSISNOTALLOWED is a free data retrieval call binding the contract method 0x3f3153b2.
//
// Solidity: function ZERO_ADDRESS_IS_NOT_ALLOWED() view returns(string)
func (_Errors *ErrorsCallerSession) ZEROADDRESSISNOTALLOWED() (string, error) {
	return _Errors.Contract.ZEROADDRESSISNOTALLOWED(&_Errors.CallOpts)
}
