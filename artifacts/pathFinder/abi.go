// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pathFinder

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

// PathFinderTradePath is an auto generated low-level Go binding around an user-defined struct.
type PathFinderTradePath struct {
	Path           []common.Address
	Rate           *big.Int
	ExpectedAmount *big.Int
}

// PathFinderMetaData contains all meta data concerning the PathFinder contract.
var PathFinderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addressProvider\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"addressProvider\",\"outputs\":[{\"internalType\":\"contractAddressProvider\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"swapInterface\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"swapType\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"bestUniPath\",\"outputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedAmount\",\"type\":\"uint256\"}],\"internalType\":\"structPathFinder.TradePath\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractsRegister\",\"outputs\":[{\"internalType\":\"contractContractsRegister\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"swapType\",\"type\":\"uint256\"}],\"name\":\"convertPathToPathV3\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_creditManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"connectorTokens\",\"type\":\"address[]\"}],\"name\":\"getClosurePaths\",\"outputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedAmount\",\"type\":\"uint256\"}],\"internalType\":\"structPathFinder.TradePath[]\",\"name\":\"result\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PathFinderABI is the input ABI used to generate the binding from.
// Deprecated: Use PathFinderMetaData.ABI instead.
var PathFinderABI = PathFinderMetaData.ABI

// PathFinder is an auto generated Go binding around an Ethereum contract.
type PathFinder struct {
	PathFinderCaller     // Read-only binding to the contract
	PathFinderTransactor // Write-only binding to the contract
	PathFinderFilterer   // Log filterer for contract events
}

// PathFinderCaller is an auto generated read-only Go binding around an Ethereum contract.
type PathFinderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PathFinderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PathFinderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PathFinderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PathFinderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PathFinderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PathFinderSession struct {
	Contract     *PathFinder       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PathFinderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PathFinderCallerSession struct {
	Contract *PathFinderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// PathFinderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PathFinderTransactorSession struct {
	Contract     *PathFinderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// PathFinderRaw is an auto generated low-level Go binding around an Ethereum contract.
type PathFinderRaw struct {
	Contract *PathFinder // Generic contract binding to access the raw methods on
}

// PathFinderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PathFinderCallerRaw struct {
	Contract *PathFinderCaller // Generic read-only contract binding to access the raw methods on
}

// PathFinderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PathFinderTransactorRaw struct {
	Contract *PathFinderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPathFinder creates a new instance of PathFinder, bound to a specific deployed contract.
func NewPathFinder(address common.Address, backend bind.ContractBackend) (*PathFinder, error) {
	contract, err := bindPathFinder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PathFinder{PathFinderCaller: PathFinderCaller{contract: contract}, PathFinderTransactor: PathFinderTransactor{contract: contract}, PathFinderFilterer: PathFinderFilterer{contract: contract}}, nil
}

// NewPathFinderCaller creates a new read-only instance of PathFinder, bound to a specific deployed contract.
func NewPathFinderCaller(address common.Address, caller bind.ContractCaller) (*PathFinderCaller, error) {
	contract, err := bindPathFinder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PathFinderCaller{contract: contract}, nil
}

// NewPathFinderTransactor creates a new write-only instance of PathFinder, bound to a specific deployed contract.
func NewPathFinderTransactor(address common.Address, transactor bind.ContractTransactor) (*PathFinderTransactor, error) {
	contract, err := bindPathFinder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PathFinderTransactor{contract: contract}, nil
}

// NewPathFinderFilterer creates a new log filterer instance of PathFinder, bound to a specific deployed contract.
func NewPathFinderFilterer(address common.Address, filterer bind.ContractFilterer) (*PathFinderFilterer, error) {
	contract, err := bindPathFinder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PathFinderFilterer{contract: contract}, nil
}

// bindPathFinder binds a generic wrapper to an already deployed contract.
func bindPathFinder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PathFinderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PathFinder *PathFinderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PathFinder.Contract.PathFinderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PathFinder *PathFinderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PathFinder.Contract.PathFinderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PathFinder *PathFinderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PathFinder.Contract.PathFinderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PathFinder *PathFinderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PathFinder.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PathFinder *PathFinderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PathFinder.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PathFinder *PathFinderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PathFinder.Contract.contract.Transact(opts, method, params...)
}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_PathFinder *PathFinderCaller) AddressProvider(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PathFinder.contract.Call(opts, &out, "addressProvider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_PathFinder *PathFinderSession) AddressProvider() (common.Address, error) {
	return _PathFinder.Contract.AddressProvider(&_PathFinder.CallOpts)
}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_PathFinder *PathFinderCallerSession) AddressProvider() (common.Address, error) {
	return _PathFinder.Contract.AddressProvider(&_PathFinder.CallOpts)
}

// ContractsRegister is a free data retrieval call binding the contract method 0x7a0c7b21.
//
// Solidity: function contractsRegister() view returns(address)
func (_PathFinder *PathFinderCaller) ContractsRegister(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PathFinder.contract.Call(opts, &out, "contractsRegister")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ContractsRegister is a free data retrieval call binding the contract method 0x7a0c7b21.
//
// Solidity: function contractsRegister() view returns(address)
func (_PathFinder *PathFinderSession) ContractsRegister() (common.Address, error) {
	return _PathFinder.Contract.ContractsRegister(&_PathFinder.CallOpts)
}

// ContractsRegister is a free data retrieval call binding the contract method 0x7a0c7b21.
//
// Solidity: function contractsRegister() view returns(address)
func (_PathFinder *PathFinderCallerSession) ContractsRegister() (common.Address, error) {
	return _PathFinder.Contract.ContractsRegister(&_PathFinder.CallOpts)
}

// ConvertPathToPathV3 is a free data retrieval call binding the contract method 0xe2430f93.
//
// Solidity: function convertPathToPathV3(address[] path, uint256 swapType) pure returns(bytes result)
func (_PathFinder *PathFinderCaller) ConvertPathToPathV3(opts *bind.CallOpts, path []common.Address, swapType *big.Int) ([]byte, error) {
	var out []interface{}
	err := _PathFinder.contract.Call(opts, &out, "convertPathToPathV3", path, swapType)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ConvertPathToPathV3 is a free data retrieval call binding the contract method 0xe2430f93.
//
// Solidity: function convertPathToPathV3(address[] path, uint256 swapType) pure returns(bytes result)
func (_PathFinder *PathFinderSession) ConvertPathToPathV3(path []common.Address, swapType *big.Int) ([]byte, error) {
	return _PathFinder.Contract.ConvertPathToPathV3(&_PathFinder.CallOpts, path, swapType)
}

// ConvertPathToPathV3 is a free data retrieval call binding the contract method 0xe2430f93.
//
// Solidity: function convertPathToPathV3(address[] path, uint256 swapType) pure returns(bytes result)
func (_PathFinder *PathFinderCallerSession) ConvertPathToPathV3(path []common.Address, swapType *big.Int) ([]byte, error) {
	return _PathFinder.Contract.ConvertPathToPathV3(&_PathFinder.CallOpts, path, swapType)
}

// BestUniPath is a paid mutator transaction binding the contract method 0x457b2030.
//
// Solidity: function bestUniPath(uint256 swapInterface, address router, uint256 swapType, address from, address to, uint256 amount, address[] tokens) returns((address[],uint256,uint256))
func (_PathFinder *PathFinderTransactor) BestUniPath(opts *bind.TransactOpts, swapInterface *big.Int, router common.Address, swapType *big.Int, from common.Address, to common.Address, amount *big.Int, tokens []common.Address) (*types.Transaction, error) {
	return _PathFinder.contract.Transact(opts, "bestUniPath", swapInterface, router, swapType, from, to, amount, tokens)
}

// BestUniPath is a paid mutator transaction binding the contract method 0x457b2030.
//
// Solidity: function bestUniPath(uint256 swapInterface, address router, uint256 swapType, address from, address to, uint256 amount, address[] tokens) returns((address[],uint256,uint256))
func (_PathFinder *PathFinderSession) BestUniPath(swapInterface *big.Int, router common.Address, swapType *big.Int, from common.Address, to common.Address, amount *big.Int, tokens []common.Address) (*types.Transaction, error) {
	return _PathFinder.Contract.BestUniPath(&_PathFinder.TransactOpts, swapInterface, router, swapType, from, to, amount, tokens)
}

// BestUniPath is a paid mutator transaction binding the contract method 0x457b2030.
//
// Solidity: function bestUniPath(uint256 swapInterface, address router, uint256 swapType, address from, address to, uint256 amount, address[] tokens) returns((address[],uint256,uint256))
func (_PathFinder *PathFinderTransactorSession) BestUniPath(swapInterface *big.Int, router common.Address, swapType *big.Int, from common.Address, to common.Address, amount *big.Int, tokens []common.Address) (*types.Transaction, error) {
	return _PathFinder.Contract.BestUniPath(&_PathFinder.TransactOpts, swapInterface, router, swapType, from, to, amount, tokens)
}

// GetClosurePaths is a paid mutator transaction binding the contract method 0x8b7ce872.
//
// Solidity: function getClosurePaths(address router, address _creditManager, address borrower, address[] connectorTokens) returns((address[],uint256,uint256)[] result)
func (_PathFinder *PathFinderTransactor) GetClosurePaths(opts *bind.TransactOpts, router common.Address, _creditManager common.Address, borrower common.Address, connectorTokens []common.Address) (*types.Transaction, error) {
	return _PathFinder.contract.Transact(opts, "getClosurePaths", router, _creditManager, borrower, connectorTokens)
}

// GetClosurePaths is a paid mutator transaction binding the contract method 0x8b7ce872.
//
// Solidity: function getClosurePaths(address router, address _creditManager, address borrower, address[] connectorTokens) returns((address[],uint256,uint256)[] result)
func (_PathFinder *PathFinderSession) GetClosurePaths(router common.Address, _creditManager common.Address, borrower common.Address, connectorTokens []common.Address) (*types.Transaction, error) {
	return _PathFinder.Contract.GetClosurePaths(&_PathFinder.TransactOpts, router, _creditManager, borrower, connectorTokens)
}

// GetClosurePaths is a paid mutator transaction binding the contract method 0x8b7ce872.
//
// Solidity: function getClosurePaths(address router, address _creditManager, address borrower, address[] connectorTokens) returns((address[],uint256,uint256)[] result)
func (_PathFinder *PathFinderTransactorSession) GetClosurePaths(router common.Address, _creditManager common.Address, borrower common.Address, connectorTokens []common.Address) (*types.Transaction, error) {
	return _PathFinder.Contract.GetClosurePaths(&_PathFinder.TransactOpts, router, _creditManager, borrower, connectorTokens)
}
