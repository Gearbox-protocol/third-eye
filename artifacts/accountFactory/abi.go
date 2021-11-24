// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package accountFactory

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

// DataTypesMiningApproval is an auto generated low-level Go binding around an user-defined struct.
type DataTypesMiningApproval struct {
	Token        common.Address
	SwapContract common.Address
}

// AccountFactoryMetaData contains all meta data concerning the AccountFactory contract.
var AccountFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addressProvider\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"}],\"name\":\"AccountMinerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creditManager\",\"type\":\"address\"}],\"name\":\"InitializeCreditAccount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"NewCreditAccount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ReturnCreditAccount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"TakeForever\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_contractsRegister\",\"outputs\":[{\"internalType\":\"contractContractsRegister\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addCreditAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"swapContract\",\"type\":\"address\"}],\"internalType\":\"structDataTypes.MiningApproval[]\",\"name\":\"_miningApprovals\",\"type\":\"tuple[]\"}],\"name\":\"addMiningApprovals\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"targetContract\",\"type\":\"address\"}],\"name\":\"cancelAllowance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"countCreditAccounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"countCreditAccountsInStock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"creditAccounts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finishMining\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"}],\"name\":\"getNext\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"head\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isCreditAccount\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isMiningFinished\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"masterCreditAccount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mineCreditAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"miningApprovals\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"swapContract\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usedAccount\",\"type\":\"address\"}],\"name\":\"returnCreditAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tail\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_borrowedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_cumulativeIndexAtOpen\",\"type\":\"uint256\"}],\"name\":\"takeCreditAccount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prev\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"takeOut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AccountFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use AccountFactoryMetaData.ABI instead.
var AccountFactoryABI = AccountFactoryMetaData.ABI

// AccountFactory is an auto generated Go binding around an Ethereum contract.
type AccountFactory struct {
	AccountFactoryCaller     // Read-only binding to the contract
	AccountFactoryTransactor // Write-only binding to the contract
	AccountFactoryFilterer   // Log filterer for contract events
}

// AccountFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccountFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccountFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccountFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccountFactorySession struct {
	Contract     *AccountFactory   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AccountFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccountFactoryCallerSession struct {
	Contract *AccountFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// AccountFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccountFactoryTransactorSession struct {
	Contract     *AccountFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// AccountFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccountFactoryRaw struct {
	Contract *AccountFactory // Generic contract binding to access the raw methods on
}

// AccountFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccountFactoryCallerRaw struct {
	Contract *AccountFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// AccountFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccountFactoryTransactorRaw struct {
	Contract *AccountFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccountFactory creates a new instance of AccountFactory, bound to a specific deployed contract.
func NewAccountFactory(address common.Address, backend bind.ContractBackend) (*AccountFactory, error) {
	contract, err := bindAccountFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AccountFactory{AccountFactoryCaller: AccountFactoryCaller{contract: contract}, AccountFactoryTransactor: AccountFactoryTransactor{contract: contract}, AccountFactoryFilterer: AccountFactoryFilterer{contract: contract}}, nil
}

// NewAccountFactoryCaller creates a new read-only instance of AccountFactory, bound to a specific deployed contract.
func NewAccountFactoryCaller(address common.Address, caller bind.ContractCaller) (*AccountFactoryCaller, error) {
	contract, err := bindAccountFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccountFactoryCaller{contract: contract}, nil
}

// NewAccountFactoryTransactor creates a new write-only instance of AccountFactory, bound to a specific deployed contract.
func NewAccountFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*AccountFactoryTransactor, error) {
	contract, err := bindAccountFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccountFactoryTransactor{contract: contract}, nil
}

// NewAccountFactoryFilterer creates a new log filterer instance of AccountFactory, bound to a specific deployed contract.
func NewAccountFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*AccountFactoryFilterer, error) {
	contract, err := bindAccountFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccountFactoryFilterer{contract: contract}, nil
}

// bindAccountFactory binds a generic wrapper to an already deployed contract.
func bindAccountFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AccountFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccountFactory *AccountFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccountFactory.Contract.AccountFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccountFactory *AccountFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccountFactory.Contract.AccountFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccountFactory *AccountFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccountFactory.Contract.AccountFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccountFactory *AccountFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccountFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccountFactory *AccountFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccountFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccountFactory *AccountFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccountFactory.Contract.contract.Transact(opts, method, params...)
}

// ContractsRegister is a free data retrieval call binding the contract method 0x88f64c54.
//
// Solidity: function _contractsRegister() view returns(address)
func (_AccountFactory *AccountFactoryCaller) ContractsRegister(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AccountFactory.contract.Call(opts, &out, "_contractsRegister")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ContractsRegister is a free data retrieval call binding the contract method 0x88f64c54.
//
// Solidity: function _contractsRegister() view returns(address)
func (_AccountFactory *AccountFactorySession) ContractsRegister() (common.Address, error) {
	return _AccountFactory.Contract.ContractsRegister(&_AccountFactory.CallOpts)
}

// ContractsRegister is a free data retrieval call binding the contract method 0x88f64c54.
//
// Solidity: function _contractsRegister() view returns(address)
func (_AccountFactory *AccountFactoryCallerSession) ContractsRegister() (common.Address, error) {
	return _AccountFactory.Contract.ContractsRegister(&_AccountFactory.CallOpts)
}

// CountCreditAccounts is a free data retrieval call binding the contract method 0xb60e8518.
//
// Solidity: function countCreditAccounts() view returns(uint256)
func (_AccountFactory *AccountFactoryCaller) CountCreditAccounts(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AccountFactory.contract.Call(opts, &out, "countCreditAccounts")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CountCreditAccounts is a free data retrieval call binding the contract method 0xb60e8518.
//
// Solidity: function countCreditAccounts() view returns(uint256)
func (_AccountFactory *AccountFactorySession) CountCreditAccounts() (*big.Int, error) {
	return _AccountFactory.Contract.CountCreditAccounts(&_AccountFactory.CallOpts)
}

// CountCreditAccounts is a free data retrieval call binding the contract method 0xb60e8518.
//
// Solidity: function countCreditAccounts() view returns(uint256)
func (_AccountFactory *AccountFactoryCallerSession) CountCreditAccounts() (*big.Int, error) {
	return _AccountFactory.Contract.CountCreditAccounts(&_AccountFactory.CallOpts)
}

// CountCreditAccountsInStock is a free data retrieval call binding the contract method 0xb1939763.
//
// Solidity: function countCreditAccountsInStock() view returns(uint256)
func (_AccountFactory *AccountFactoryCaller) CountCreditAccountsInStock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AccountFactory.contract.Call(opts, &out, "countCreditAccountsInStock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CountCreditAccountsInStock is a free data retrieval call binding the contract method 0xb1939763.
//
// Solidity: function countCreditAccountsInStock() view returns(uint256)
func (_AccountFactory *AccountFactorySession) CountCreditAccountsInStock() (*big.Int, error) {
	return _AccountFactory.Contract.CountCreditAccountsInStock(&_AccountFactory.CallOpts)
}

// CountCreditAccountsInStock is a free data retrieval call binding the contract method 0xb1939763.
//
// Solidity: function countCreditAccountsInStock() view returns(uint256)
func (_AccountFactory *AccountFactoryCallerSession) CountCreditAccountsInStock() (*big.Int, error) {
	return _AccountFactory.Contract.CountCreditAccountsInStock(&_AccountFactory.CallOpts)
}

// CreditAccounts is a free data retrieval call binding the contract method 0xe3ba9ace.
//
// Solidity: function creditAccounts(uint256 id) view returns(address)
func (_AccountFactory *AccountFactoryCaller) CreditAccounts(opts *bind.CallOpts, id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AccountFactory.contract.Call(opts, &out, "creditAccounts", id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditAccounts is a free data retrieval call binding the contract method 0xe3ba9ace.
//
// Solidity: function creditAccounts(uint256 id) view returns(address)
func (_AccountFactory *AccountFactorySession) CreditAccounts(id *big.Int) (common.Address, error) {
	return _AccountFactory.Contract.CreditAccounts(&_AccountFactory.CallOpts, id)
}

// CreditAccounts is a free data retrieval call binding the contract method 0xe3ba9ace.
//
// Solidity: function creditAccounts(uint256 id) view returns(address)
func (_AccountFactory *AccountFactoryCallerSession) CreditAccounts(id *big.Int) (common.Address, error) {
	return _AccountFactory.Contract.CreditAccounts(&_AccountFactory.CallOpts, id)
}

// GetNext is a free data retrieval call binding the contract method 0x765e0159.
//
// Solidity: function getNext(address creditAccount) view returns(address)
func (_AccountFactory *AccountFactoryCaller) GetNext(opts *bind.CallOpts, creditAccount common.Address) (common.Address, error) {
	var out []interface{}
	err := _AccountFactory.contract.Call(opts, &out, "getNext", creditAccount)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetNext is a free data retrieval call binding the contract method 0x765e0159.
//
// Solidity: function getNext(address creditAccount) view returns(address)
func (_AccountFactory *AccountFactorySession) GetNext(creditAccount common.Address) (common.Address, error) {
	return _AccountFactory.Contract.GetNext(&_AccountFactory.CallOpts, creditAccount)
}

// GetNext is a free data retrieval call binding the contract method 0x765e0159.
//
// Solidity: function getNext(address creditAccount) view returns(address)
func (_AccountFactory *AccountFactoryCallerSession) GetNext(creditAccount common.Address) (common.Address, error) {
	return _AccountFactory.Contract.GetNext(&_AccountFactory.CallOpts, creditAccount)
}

// Head is a free data retrieval call binding the contract method 0x8f7dcfa3.
//
// Solidity: function head() view returns(address)
func (_AccountFactory *AccountFactoryCaller) Head(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AccountFactory.contract.Call(opts, &out, "head")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Head is a free data retrieval call binding the contract method 0x8f7dcfa3.
//
// Solidity: function head() view returns(address)
func (_AccountFactory *AccountFactorySession) Head() (common.Address, error) {
	return _AccountFactory.Contract.Head(&_AccountFactory.CallOpts)
}

// Head is a free data retrieval call binding the contract method 0x8f7dcfa3.
//
// Solidity: function head() view returns(address)
func (_AccountFactory *AccountFactoryCallerSession) Head() (common.Address, error) {
	return _AccountFactory.Contract.Head(&_AccountFactory.CallOpts)
}

// IsCreditAccount is a free data retrieval call binding the contract method 0xd82ecc48.
//
// Solidity: function isCreditAccount(address addr) view returns(bool)
func (_AccountFactory *AccountFactoryCaller) IsCreditAccount(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _AccountFactory.contract.Call(opts, &out, "isCreditAccount", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCreditAccount is a free data retrieval call binding the contract method 0xd82ecc48.
//
// Solidity: function isCreditAccount(address addr) view returns(bool)
func (_AccountFactory *AccountFactorySession) IsCreditAccount(addr common.Address) (bool, error) {
	return _AccountFactory.Contract.IsCreditAccount(&_AccountFactory.CallOpts, addr)
}

// IsCreditAccount is a free data retrieval call binding the contract method 0xd82ecc48.
//
// Solidity: function isCreditAccount(address addr) view returns(bool)
func (_AccountFactory *AccountFactoryCallerSession) IsCreditAccount(addr common.Address) (bool, error) {
	return _AccountFactory.Contract.IsCreditAccount(&_AccountFactory.CallOpts, addr)
}

// IsMiningFinished is a free data retrieval call binding the contract method 0x23428dbd.
//
// Solidity: function isMiningFinished() view returns(bool)
func (_AccountFactory *AccountFactoryCaller) IsMiningFinished(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _AccountFactory.contract.Call(opts, &out, "isMiningFinished")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMiningFinished is a free data retrieval call binding the contract method 0x23428dbd.
//
// Solidity: function isMiningFinished() view returns(bool)
func (_AccountFactory *AccountFactorySession) IsMiningFinished() (bool, error) {
	return _AccountFactory.Contract.IsMiningFinished(&_AccountFactory.CallOpts)
}

// IsMiningFinished is a free data retrieval call binding the contract method 0x23428dbd.
//
// Solidity: function isMiningFinished() view returns(bool)
func (_AccountFactory *AccountFactoryCallerSession) IsMiningFinished() (bool, error) {
	return _AccountFactory.Contract.IsMiningFinished(&_AccountFactory.CallOpts)
}

// MasterCreditAccount is a free data retrieval call binding the contract method 0x5da33c5b.
//
// Solidity: function masterCreditAccount() view returns(address)
func (_AccountFactory *AccountFactoryCaller) MasterCreditAccount(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AccountFactory.contract.Call(opts, &out, "masterCreditAccount")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MasterCreditAccount is a free data retrieval call binding the contract method 0x5da33c5b.
//
// Solidity: function masterCreditAccount() view returns(address)
func (_AccountFactory *AccountFactorySession) MasterCreditAccount() (common.Address, error) {
	return _AccountFactory.Contract.MasterCreditAccount(&_AccountFactory.CallOpts)
}

// MasterCreditAccount is a free data retrieval call binding the contract method 0x5da33c5b.
//
// Solidity: function masterCreditAccount() view returns(address)
func (_AccountFactory *AccountFactoryCallerSession) MasterCreditAccount() (common.Address, error) {
	return _AccountFactory.Contract.MasterCreditAccount(&_AccountFactory.CallOpts)
}

// MiningApprovals is a free data retrieval call binding the contract method 0x9c650789.
//
// Solidity: function miningApprovals(uint256 ) view returns(address token, address swapContract)
func (_AccountFactory *AccountFactoryCaller) MiningApprovals(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Token        common.Address
	SwapContract common.Address
}, error) {
	var out []interface{}
	err := _AccountFactory.contract.Call(opts, &out, "miningApprovals", arg0)

	outstruct := new(struct {
		Token        common.Address
		SwapContract common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Token = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.SwapContract = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// MiningApprovals is a free data retrieval call binding the contract method 0x9c650789.
//
// Solidity: function miningApprovals(uint256 ) view returns(address token, address swapContract)
func (_AccountFactory *AccountFactorySession) MiningApprovals(arg0 *big.Int) (struct {
	Token        common.Address
	SwapContract common.Address
}, error) {
	return _AccountFactory.Contract.MiningApprovals(&_AccountFactory.CallOpts, arg0)
}

// MiningApprovals is a free data retrieval call binding the contract method 0x9c650789.
//
// Solidity: function miningApprovals(uint256 ) view returns(address token, address swapContract)
func (_AccountFactory *AccountFactoryCallerSession) MiningApprovals(arg0 *big.Int) (struct {
	Token        common.Address
	SwapContract common.Address
}, error) {
	return _AccountFactory.Contract.MiningApprovals(&_AccountFactory.CallOpts, arg0)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_AccountFactory *AccountFactoryCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _AccountFactory.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_AccountFactory *AccountFactorySession) Paused() (bool, error) {
	return _AccountFactory.Contract.Paused(&_AccountFactory.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_AccountFactory *AccountFactoryCallerSession) Paused() (bool, error) {
	return _AccountFactory.Contract.Paused(&_AccountFactory.CallOpts)
}

// Tail is a free data retrieval call binding the contract method 0x13d8c840.
//
// Solidity: function tail() view returns(address)
func (_AccountFactory *AccountFactoryCaller) Tail(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AccountFactory.contract.Call(opts, &out, "tail")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Tail is a free data retrieval call binding the contract method 0x13d8c840.
//
// Solidity: function tail() view returns(address)
func (_AccountFactory *AccountFactorySession) Tail() (common.Address, error) {
	return _AccountFactory.Contract.Tail(&_AccountFactory.CallOpts)
}

// Tail is a free data retrieval call binding the contract method 0x13d8c840.
//
// Solidity: function tail() view returns(address)
func (_AccountFactory *AccountFactoryCallerSession) Tail() (common.Address, error) {
	return _AccountFactory.Contract.Tail(&_AccountFactory.CallOpts)
}

// AddCreditAccount is a paid mutator transaction binding the contract method 0xf23953ab.
//
// Solidity: function addCreditAccount() returns()
func (_AccountFactory *AccountFactoryTransactor) AddCreditAccount(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccountFactory.contract.Transact(opts, "addCreditAccount")
}

// AddCreditAccount is a paid mutator transaction binding the contract method 0xf23953ab.
//
// Solidity: function addCreditAccount() returns()
func (_AccountFactory *AccountFactorySession) AddCreditAccount() (*types.Transaction, error) {
	return _AccountFactory.Contract.AddCreditAccount(&_AccountFactory.TransactOpts)
}

// AddCreditAccount is a paid mutator transaction binding the contract method 0xf23953ab.
//
// Solidity: function addCreditAccount() returns()
func (_AccountFactory *AccountFactoryTransactorSession) AddCreditAccount() (*types.Transaction, error) {
	return _AccountFactory.Contract.AddCreditAccount(&_AccountFactory.TransactOpts)
}

// AddMiningApprovals is a paid mutator transaction binding the contract method 0x3b9c4867.
//
// Solidity: function addMiningApprovals((address,address)[] _miningApprovals) returns()
func (_AccountFactory *AccountFactoryTransactor) AddMiningApprovals(opts *bind.TransactOpts, _miningApprovals []DataTypesMiningApproval) (*types.Transaction, error) {
	return _AccountFactory.contract.Transact(opts, "addMiningApprovals", _miningApprovals)
}

// AddMiningApprovals is a paid mutator transaction binding the contract method 0x3b9c4867.
//
// Solidity: function addMiningApprovals((address,address)[] _miningApprovals) returns()
func (_AccountFactory *AccountFactorySession) AddMiningApprovals(_miningApprovals []DataTypesMiningApproval) (*types.Transaction, error) {
	return _AccountFactory.Contract.AddMiningApprovals(&_AccountFactory.TransactOpts, _miningApprovals)
}

// AddMiningApprovals is a paid mutator transaction binding the contract method 0x3b9c4867.
//
// Solidity: function addMiningApprovals((address,address)[] _miningApprovals) returns()
func (_AccountFactory *AccountFactoryTransactorSession) AddMiningApprovals(_miningApprovals []DataTypesMiningApproval) (*types.Transaction, error) {
	return _AccountFactory.Contract.AddMiningApprovals(&_AccountFactory.TransactOpts, _miningApprovals)
}

// CancelAllowance is a paid mutator transaction binding the contract method 0xa904aab6.
//
// Solidity: function cancelAllowance(address account, address token, address targetContract) returns()
func (_AccountFactory *AccountFactoryTransactor) CancelAllowance(opts *bind.TransactOpts, account common.Address, token common.Address, targetContract common.Address) (*types.Transaction, error) {
	return _AccountFactory.contract.Transact(opts, "cancelAllowance", account, token, targetContract)
}

// CancelAllowance is a paid mutator transaction binding the contract method 0xa904aab6.
//
// Solidity: function cancelAllowance(address account, address token, address targetContract) returns()
func (_AccountFactory *AccountFactorySession) CancelAllowance(account common.Address, token common.Address, targetContract common.Address) (*types.Transaction, error) {
	return _AccountFactory.Contract.CancelAllowance(&_AccountFactory.TransactOpts, account, token, targetContract)
}

// CancelAllowance is a paid mutator transaction binding the contract method 0xa904aab6.
//
// Solidity: function cancelAllowance(address account, address token, address targetContract) returns()
func (_AccountFactory *AccountFactoryTransactorSession) CancelAllowance(account common.Address, token common.Address, targetContract common.Address) (*types.Transaction, error) {
	return _AccountFactory.Contract.CancelAllowance(&_AccountFactory.TransactOpts, account, token, targetContract)
}

// FinishMining is a paid mutator transaction binding the contract method 0x3a1ed19a.
//
// Solidity: function finishMining() returns()
func (_AccountFactory *AccountFactoryTransactor) FinishMining(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccountFactory.contract.Transact(opts, "finishMining")
}

// FinishMining is a paid mutator transaction binding the contract method 0x3a1ed19a.
//
// Solidity: function finishMining() returns()
func (_AccountFactory *AccountFactorySession) FinishMining() (*types.Transaction, error) {
	return _AccountFactory.Contract.FinishMining(&_AccountFactory.TransactOpts)
}

// FinishMining is a paid mutator transaction binding the contract method 0x3a1ed19a.
//
// Solidity: function finishMining() returns()
func (_AccountFactory *AccountFactoryTransactorSession) FinishMining() (*types.Transaction, error) {
	return _AccountFactory.Contract.FinishMining(&_AccountFactory.TransactOpts)
}

// MineCreditAccount is a paid mutator transaction binding the contract method 0xb014352f.
//
// Solidity: function mineCreditAccount() returns()
func (_AccountFactory *AccountFactoryTransactor) MineCreditAccount(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccountFactory.contract.Transact(opts, "mineCreditAccount")
}

// MineCreditAccount is a paid mutator transaction binding the contract method 0xb014352f.
//
// Solidity: function mineCreditAccount() returns()
func (_AccountFactory *AccountFactorySession) MineCreditAccount() (*types.Transaction, error) {
	return _AccountFactory.Contract.MineCreditAccount(&_AccountFactory.TransactOpts)
}

// MineCreditAccount is a paid mutator transaction binding the contract method 0xb014352f.
//
// Solidity: function mineCreditAccount() returns()
func (_AccountFactory *AccountFactoryTransactorSession) MineCreditAccount() (*types.Transaction, error) {
	return _AccountFactory.Contract.MineCreditAccount(&_AccountFactory.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_AccountFactory *AccountFactoryTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccountFactory.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_AccountFactory *AccountFactorySession) Pause() (*types.Transaction, error) {
	return _AccountFactory.Contract.Pause(&_AccountFactory.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_AccountFactory *AccountFactoryTransactorSession) Pause() (*types.Transaction, error) {
	return _AccountFactory.Contract.Pause(&_AccountFactory.TransactOpts)
}

// ReturnCreditAccount is a paid mutator transaction binding the contract method 0x89b77b3e.
//
// Solidity: function returnCreditAccount(address usedAccount) returns()
func (_AccountFactory *AccountFactoryTransactor) ReturnCreditAccount(opts *bind.TransactOpts, usedAccount common.Address) (*types.Transaction, error) {
	return _AccountFactory.contract.Transact(opts, "returnCreditAccount", usedAccount)
}

// ReturnCreditAccount is a paid mutator transaction binding the contract method 0x89b77b3e.
//
// Solidity: function returnCreditAccount(address usedAccount) returns()
func (_AccountFactory *AccountFactorySession) ReturnCreditAccount(usedAccount common.Address) (*types.Transaction, error) {
	return _AccountFactory.Contract.ReturnCreditAccount(&_AccountFactory.TransactOpts, usedAccount)
}

// ReturnCreditAccount is a paid mutator transaction binding the contract method 0x89b77b3e.
//
// Solidity: function returnCreditAccount(address usedAccount) returns()
func (_AccountFactory *AccountFactoryTransactorSession) ReturnCreditAccount(usedAccount common.Address) (*types.Transaction, error) {
	return _AccountFactory.Contract.ReturnCreditAccount(&_AccountFactory.TransactOpts, usedAccount)
}

// TakeCreditAccount is a paid mutator transaction binding the contract method 0x21d18456.
//
// Solidity: function takeCreditAccount(uint256 _borrowedAmount, uint256 _cumulativeIndexAtOpen) returns(address)
func (_AccountFactory *AccountFactoryTransactor) TakeCreditAccount(opts *bind.TransactOpts, _borrowedAmount *big.Int, _cumulativeIndexAtOpen *big.Int) (*types.Transaction, error) {
	return _AccountFactory.contract.Transact(opts, "takeCreditAccount", _borrowedAmount, _cumulativeIndexAtOpen)
}

// TakeCreditAccount is a paid mutator transaction binding the contract method 0x21d18456.
//
// Solidity: function takeCreditAccount(uint256 _borrowedAmount, uint256 _cumulativeIndexAtOpen) returns(address)
func (_AccountFactory *AccountFactorySession) TakeCreditAccount(_borrowedAmount *big.Int, _cumulativeIndexAtOpen *big.Int) (*types.Transaction, error) {
	return _AccountFactory.Contract.TakeCreditAccount(&_AccountFactory.TransactOpts, _borrowedAmount, _cumulativeIndexAtOpen)
}

// TakeCreditAccount is a paid mutator transaction binding the contract method 0x21d18456.
//
// Solidity: function takeCreditAccount(uint256 _borrowedAmount, uint256 _cumulativeIndexAtOpen) returns(address)
func (_AccountFactory *AccountFactoryTransactorSession) TakeCreditAccount(_borrowedAmount *big.Int, _cumulativeIndexAtOpen *big.Int) (*types.Transaction, error) {
	return _AccountFactory.Contract.TakeCreditAccount(&_AccountFactory.TransactOpts, _borrowedAmount, _cumulativeIndexAtOpen)
}

// TakeOut is a paid mutator transaction binding the contract method 0x2932472f.
//
// Solidity: function takeOut(address prev, address creditAccount, address to) returns()
func (_AccountFactory *AccountFactoryTransactor) TakeOut(opts *bind.TransactOpts, prev common.Address, creditAccount common.Address, to common.Address) (*types.Transaction, error) {
	return _AccountFactory.contract.Transact(opts, "takeOut", prev, creditAccount, to)
}

// TakeOut is a paid mutator transaction binding the contract method 0x2932472f.
//
// Solidity: function takeOut(address prev, address creditAccount, address to) returns()
func (_AccountFactory *AccountFactorySession) TakeOut(prev common.Address, creditAccount common.Address, to common.Address) (*types.Transaction, error) {
	return _AccountFactory.Contract.TakeOut(&_AccountFactory.TransactOpts, prev, creditAccount, to)
}

// TakeOut is a paid mutator transaction binding the contract method 0x2932472f.
//
// Solidity: function takeOut(address prev, address creditAccount, address to) returns()
func (_AccountFactory *AccountFactoryTransactorSession) TakeOut(prev common.Address, creditAccount common.Address, to common.Address) (*types.Transaction, error) {
	return _AccountFactory.Contract.TakeOut(&_AccountFactory.TransactOpts, prev, creditAccount, to)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_AccountFactory *AccountFactoryTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccountFactory.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_AccountFactory *AccountFactorySession) Unpause() (*types.Transaction, error) {
	return _AccountFactory.Contract.Unpause(&_AccountFactory.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_AccountFactory *AccountFactoryTransactorSession) Unpause() (*types.Transaction, error) {
	return _AccountFactory.Contract.Unpause(&_AccountFactory.TransactOpts)
}

// AccountFactoryAccountMinerChangedIterator is returned from FilterAccountMinerChanged and is used to iterate over the raw logs and unpacked data for AccountMinerChanged events raised by the AccountFactory contract.
type AccountFactoryAccountMinerChangedIterator struct {
	Event *AccountFactoryAccountMinerChanged // Event containing the contract specifics and raw log

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
func (it *AccountFactoryAccountMinerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountFactoryAccountMinerChanged)
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
		it.Event = new(AccountFactoryAccountMinerChanged)
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
func (it *AccountFactoryAccountMinerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountFactoryAccountMinerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountFactoryAccountMinerChanged represents a AccountMinerChanged event raised by the AccountFactory contract.
type AccountFactoryAccountMinerChanged struct {
	Miner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAccountMinerChanged is a free log retrieval operation binding the contract event 0xb5b22c95380a75185488532000fd4826f19e58c5eba212f266d9861a44b671fc.
//
// Solidity: event AccountMinerChanged(address indexed miner)
func (_AccountFactory *AccountFactoryFilterer) FilterAccountMinerChanged(opts *bind.FilterOpts, miner []common.Address) (*AccountFactoryAccountMinerChangedIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _AccountFactory.contract.FilterLogs(opts, "AccountMinerChanged", minerRule)
	if err != nil {
		return nil, err
	}
	return &AccountFactoryAccountMinerChangedIterator{contract: _AccountFactory.contract, event: "AccountMinerChanged", logs: logs, sub: sub}, nil
}

// WatchAccountMinerChanged is a free log subscription operation binding the contract event 0xb5b22c95380a75185488532000fd4826f19e58c5eba212f266d9861a44b671fc.
//
// Solidity: event AccountMinerChanged(address indexed miner)
func (_AccountFactory *AccountFactoryFilterer) WatchAccountMinerChanged(opts *bind.WatchOpts, sink chan<- *AccountFactoryAccountMinerChanged, miner []common.Address) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _AccountFactory.contract.WatchLogs(opts, "AccountMinerChanged", minerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountFactoryAccountMinerChanged)
				if err := _AccountFactory.contract.UnpackLog(event, "AccountMinerChanged", log); err != nil {
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

// ParseAccountMinerChanged is a log parse operation binding the contract event 0xb5b22c95380a75185488532000fd4826f19e58c5eba212f266d9861a44b671fc.
//
// Solidity: event AccountMinerChanged(address indexed miner)
func (_AccountFactory *AccountFactoryFilterer) ParseAccountMinerChanged(log types.Log) (*AccountFactoryAccountMinerChanged, error) {
	event := new(AccountFactoryAccountMinerChanged)
	if err := _AccountFactory.contract.UnpackLog(event, "AccountMinerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountFactoryInitializeCreditAccountIterator is returned from FilterInitializeCreditAccount and is used to iterate over the raw logs and unpacked data for InitializeCreditAccount events raised by the AccountFactory contract.
type AccountFactoryInitializeCreditAccountIterator struct {
	Event *AccountFactoryInitializeCreditAccount // Event containing the contract specifics and raw log

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
func (it *AccountFactoryInitializeCreditAccountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountFactoryInitializeCreditAccount)
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
		it.Event = new(AccountFactoryInitializeCreditAccount)
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
func (it *AccountFactoryInitializeCreditAccountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountFactoryInitializeCreditAccountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountFactoryInitializeCreditAccount represents a InitializeCreditAccount event raised by the AccountFactory contract.
type AccountFactoryInitializeCreditAccount struct {
	Account       common.Address
	CreditManager common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterInitializeCreditAccount is a free log retrieval operation binding the contract event 0xf3ede7039176503a8ad1fe7cfaa29475a9dbe0cdcaf04ecf9a5c10570c47b103.
//
// Solidity: event InitializeCreditAccount(address indexed account, address indexed creditManager)
func (_AccountFactory *AccountFactoryFilterer) FilterInitializeCreditAccount(opts *bind.FilterOpts, account []common.Address, creditManager []common.Address) (*AccountFactoryInitializeCreditAccountIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var creditManagerRule []interface{}
	for _, creditManagerItem := range creditManager {
		creditManagerRule = append(creditManagerRule, creditManagerItem)
	}

	logs, sub, err := _AccountFactory.contract.FilterLogs(opts, "InitializeCreditAccount", accountRule, creditManagerRule)
	if err != nil {
		return nil, err
	}
	return &AccountFactoryInitializeCreditAccountIterator{contract: _AccountFactory.contract, event: "InitializeCreditAccount", logs: logs, sub: sub}, nil
}

// WatchInitializeCreditAccount is a free log subscription operation binding the contract event 0xf3ede7039176503a8ad1fe7cfaa29475a9dbe0cdcaf04ecf9a5c10570c47b103.
//
// Solidity: event InitializeCreditAccount(address indexed account, address indexed creditManager)
func (_AccountFactory *AccountFactoryFilterer) WatchInitializeCreditAccount(opts *bind.WatchOpts, sink chan<- *AccountFactoryInitializeCreditAccount, account []common.Address, creditManager []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var creditManagerRule []interface{}
	for _, creditManagerItem := range creditManager {
		creditManagerRule = append(creditManagerRule, creditManagerItem)
	}

	logs, sub, err := _AccountFactory.contract.WatchLogs(opts, "InitializeCreditAccount", accountRule, creditManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountFactoryInitializeCreditAccount)
				if err := _AccountFactory.contract.UnpackLog(event, "InitializeCreditAccount", log); err != nil {
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

// ParseInitializeCreditAccount is a log parse operation binding the contract event 0xf3ede7039176503a8ad1fe7cfaa29475a9dbe0cdcaf04ecf9a5c10570c47b103.
//
// Solidity: event InitializeCreditAccount(address indexed account, address indexed creditManager)
func (_AccountFactory *AccountFactoryFilterer) ParseInitializeCreditAccount(log types.Log) (*AccountFactoryInitializeCreditAccount, error) {
	event := new(AccountFactoryInitializeCreditAccount)
	if err := _AccountFactory.contract.UnpackLog(event, "InitializeCreditAccount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountFactoryNewCreditAccountIterator is returned from FilterNewCreditAccount and is used to iterate over the raw logs and unpacked data for NewCreditAccount events raised by the AccountFactory contract.
type AccountFactoryNewCreditAccountIterator struct {
	Event *AccountFactoryNewCreditAccount // Event containing the contract specifics and raw log

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
func (it *AccountFactoryNewCreditAccountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountFactoryNewCreditAccount)
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
		it.Event = new(AccountFactoryNewCreditAccount)
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
func (it *AccountFactoryNewCreditAccountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountFactoryNewCreditAccountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountFactoryNewCreditAccount represents a NewCreditAccount event raised by the AccountFactory contract.
type AccountFactoryNewCreditAccount struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewCreditAccount is a free log retrieval operation binding the contract event 0x9f69b6c10f6810213e055b0ba6bc0a4e2603f73c221aad77ea35da819cda7dc3.
//
// Solidity: event NewCreditAccount(address indexed account)
func (_AccountFactory *AccountFactoryFilterer) FilterNewCreditAccount(opts *bind.FilterOpts, account []common.Address) (*AccountFactoryNewCreditAccountIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AccountFactory.contract.FilterLogs(opts, "NewCreditAccount", accountRule)
	if err != nil {
		return nil, err
	}
	return &AccountFactoryNewCreditAccountIterator{contract: _AccountFactory.contract, event: "NewCreditAccount", logs: logs, sub: sub}, nil
}

// WatchNewCreditAccount is a free log subscription operation binding the contract event 0x9f69b6c10f6810213e055b0ba6bc0a4e2603f73c221aad77ea35da819cda7dc3.
//
// Solidity: event NewCreditAccount(address indexed account)
func (_AccountFactory *AccountFactoryFilterer) WatchNewCreditAccount(opts *bind.WatchOpts, sink chan<- *AccountFactoryNewCreditAccount, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AccountFactory.contract.WatchLogs(opts, "NewCreditAccount", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountFactoryNewCreditAccount)
				if err := _AccountFactory.contract.UnpackLog(event, "NewCreditAccount", log); err != nil {
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

// ParseNewCreditAccount is a log parse operation binding the contract event 0x9f69b6c10f6810213e055b0ba6bc0a4e2603f73c221aad77ea35da819cda7dc3.
//
// Solidity: event NewCreditAccount(address indexed account)
func (_AccountFactory *AccountFactoryFilterer) ParseNewCreditAccount(log types.Log) (*AccountFactoryNewCreditAccount, error) {
	event := new(AccountFactoryNewCreditAccount)
	if err := _AccountFactory.contract.UnpackLog(event, "NewCreditAccount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountFactoryPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the AccountFactory contract.
type AccountFactoryPausedIterator struct {
	Event *AccountFactoryPaused // Event containing the contract specifics and raw log

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
func (it *AccountFactoryPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountFactoryPaused)
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
		it.Event = new(AccountFactoryPaused)
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
func (it *AccountFactoryPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountFactoryPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountFactoryPaused represents a Paused event raised by the AccountFactory contract.
type AccountFactoryPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_AccountFactory *AccountFactoryFilterer) FilterPaused(opts *bind.FilterOpts) (*AccountFactoryPausedIterator, error) {

	logs, sub, err := _AccountFactory.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &AccountFactoryPausedIterator{contract: _AccountFactory.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_AccountFactory *AccountFactoryFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *AccountFactoryPaused) (event.Subscription, error) {

	logs, sub, err := _AccountFactory.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountFactoryPaused)
				if err := _AccountFactory.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_AccountFactory *AccountFactoryFilterer) ParsePaused(log types.Log) (*AccountFactoryPaused, error) {
	event := new(AccountFactoryPaused)
	if err := _AccountFactory.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountFactoryReturnCreditAccountIterator is returned from FilterReturnCreditAccount and is used to iterate over the raw logs and unpacked data for ReturnCreditAccount events raised by the AccountFactory contract.
type AccountFactoryReturnCreditAccountIterator struct {
	Event *AccountFactoryReturnCreditAccount // Event containing the contract specifics and raw log

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
func (it *AccountFactoryReturnCreditAccountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountFactoryReturnCreditAccount)
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
		it.Event = new(AccountFactoryReturnCreditAccount)
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
func (it *AccountFactoryReturnCreditAccountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountFactoryReturnCreditAccountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountFactoryReturnCreditAccount represents a ReturnCreditAccount event raised by the AccountFactory contract.
type AccountFactoryReturnCreditAccount struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterReturnCreditAccount is a free log retrieval operation binding the contract event 0xced6ab9afc868b3a088366f6631ae20752993b5cce5d5f0534ea5a59fcc57d56.
//
// Solidity: event ReturnCreditAccount(address indexed account)
func (_AccountFactory *AccountFactoryFilterer) FilterReturnCreditAccount(opts *bind.FilterOpts, account []common.Address) (*AccountFactoryReturnCreditAccountIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AccountFactory.contract.FilterLogs(opts, "ReturnCreditAccount", accountRule)
	if err != nil {
		return nil, err
	}
	return &AccountFactoryReturnCreditAccountIterator{contract: _AccountFactory.contract, event: "ReturnCreditAccount", logs: logs, sub: sub}, nil
}

// WatchReturnCreditAccount is a free log subscription operation binding the contract event 0xced6ab9afc868b3a088366f6631ae20752993b5cce5d5f0534ea5a59fcc57d56.
//
// Solidity: event ReturnCreditAccount(address indexed account)
func (_AccountFactory *AccountFactoryFilterer) WatchReturnCreditAccount(opts *bind.WatchOpts, sink chan<- *AccountFactoryReturnCreditAccount, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AccountFactory.contract.WatchLogs(opts, "ReturnCreditAccount", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountFactoryReturnCreditAccount)
				if err := _AccountFactory.contract.UnpackLog(event, "ReturnCreditAccount", log); err != nil {
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

// ParseReturnCreditAccount is a log parse operation binding the contract event 0xced6ab9afc868b3a088366f6631ae20752993b5cce5d5f0534ea5a59fcc57d56.
//
// Solidity: event ReturnCreditAccount(address indexed account)
func (_AccountFactory *AccountFactoryFilterer) ParseReturnCreditAccount(log types.Log) (*AccountFactoryReturnCreditAccount, error) {
	event := new(AccountFactoryReturnCreditAccount)
	if err := _AccountFactory.contract.UnpackLog(event, "ReturnCreditAccount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountFactoryTakeForeverIterator is returned from FilterTakeForever and is used to iterate over the raw logs and unpacked data for TakeForever events raised by the AccountFactory contract.
type AccountFactoryTakeForeverIterator struct {
	Event *AccountFactoryTakeForever // Event containing the contract specifics and raw log

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
func (it *AccountFactoryTakeForeverIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountFactoryTakeForever)
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
		it.Event = new(AccountFactoryTakeForever)
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
func (it *AccountFactoryTakeForeverIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountFactoryTakeForeverIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountFactoryTakeForever represents a TakeForever event raised by the AccountFactory contract.
type AccountFactoryTakeForever struct {
	CreditAccount common.Address
	To            common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTakeForever is a free log retrieval operation binding the contract event 0x25e267469ba2ae82515be7b3d45df60bf8308343f0809e8cf7319058e2255ce6.
//
// Solidity: event TakeForever(address indexed creditAccount, address indexed to)
func (_AccountFactory *AccountFactoryFilterer) FilterTakeForever(opts *bind.FilterOpts, creditAccount []common.Address, to []common.Address) (*AccountFactoryTakeForeverIterator, error) {

	var creditAccountRule []interface{}
	for _, creditAccountItem := range creditAccount {
		creditAccountRule = append(creditAccountRule, creditAccountItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AccountFactory.contract.FilterLogs(opts, "TakeForever", creditAccountRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AccountFactoryTakeForeverIterator{contract: _AccountFactory.contract, event: "TakeForever", logs: logs, sub: sub}, nil
}

// WatchTakeForever is a free log subscription operation binding the contract event 0x25e267469ba2ae82515be7b3d45df60bf8308343f0809e8cf7319058e2255ce6.
//
// Solidity: event TakeForever(address indexed creditAccount, address indexed to)
func (_AccountFactory *AccountFactoryFilterer) WatchTakeForever(opts *bind.WatchOpts, sink chan<- *AccountFactoryTakeForever, creditAccount []common.Address, to []common.Address) (event.Subscription, error) {

	var creditAccountRule []interface{}
	for _, creditAccountItem := range creditAccount {
		creditAccountRule = append(creditAccountRule, creditAccountItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AccountFactory.contract.WatchLogs(opts, "TakeForever", creditAccountRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountFactoryTakeForever)
				if err := _AccountFactory.contract.UnpackLog(event, "TakeForever", log); err != nil {
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

// ParseTakeForever is a log parse operation binding the contract event 0x25e267469ba2ae82515be7b3d45df60bf8308343f0809e8cf7319058e2255ce6.
//
// Solidity: event TakeForever(address indexed creditAccount, address indexed to)
func (_AccountFactory *AccountFactoryFilterer) ParseTakeForever(log types.Log) (*AccountFactoryTakeForever, error) {
	event := new(AccountFactoryTakeForever)
	if err := _AccountFactory.contract.UnpackLog(event, "TakeForever", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountFactoryUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the AccountFactory contract.
type AccountFactoryUnpausedIterator struct {
	Event *AccountFactoryUnpaused // Event containing the contract specifics and raw log

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
func (it *AccountFactoryUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountFactoryUnpaused)
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
		it.Event = new(AccountFactoryUnpaused)
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
func (it *AccountFactoryUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountFactoryUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountFactoryUnpaused represents a Unpaused event raised by the AccountFactory contract.
type AccountFactoryUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_AccountFactory *AccountFactoryFilterer) FilterUnpaused(opts *bind.FilterOpts) (*AccountFactoryUnpausedIterator, error) {

	logs, sub, err := _AccountFactory.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &AccountFactoryUnpausedIterator{contract: _AccountFactory.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_AccountFactory *AccountFactoryFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *AccountFactoryUnpaused) (event.Subscription, error) {

	logs, sub, err := _AccountFactory.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountFactoryUnpaused)
				if err := _AccountFactory.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_AccountFactory *AccountFactoryFilterer) ParseUnpaused(log types.Log) (*AccountFactoryUnpaused, error) {
	event := new(AccountFactoryUnpaused)
	if err := _AccountFactory.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
