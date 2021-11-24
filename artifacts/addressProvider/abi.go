// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package addressProvider

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

// AddressProviderMetaData contains all meta data concerning the AddressProvider contract.
var AddressProviderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"service\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"AddressSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ACCOUNT_FACTORY\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ACL\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CONTRACTS_REGISTER\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DATA_COMPRESSOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GEAR_TOKEN\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LEVERAGED_ACTIONS\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PRICE_ORACLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TREASURY_CONTRACT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WETH_GATEWAY\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WETH_TOKEN\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"addresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getACL\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAccountFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getContractsRegister\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDataCompressor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGearToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLeveragedActions\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPriceOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTreasuryContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWETHGateway\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWethToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setACL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setAccountFactory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setContractsRegister\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setDataCompressor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setGearToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setLeveragedActions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setPriceOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setTreasuryContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setWETHGateway\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setWethToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AddressProviderABI is the input ABI used to generate the binding from.
// Deprecated: Use AddressProviderMetaData.ABI instead.
var AddressProviderABI = AddressProviderMetaData.ABI

// AddressProvider is an auto generated Go binding around an Ethereum contract.
type AddressProvider struct {
	AddressProviderCaller     // Read-only binding to the contract
	AddressProviderTransactor // Write-only binding to the contract
	AddressProviderFilterer   // Log filterer for contract events
}

// AddressProviderCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressProviderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressProviderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressProviderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressProviderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressProviderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressProviderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressProviderSession struct {
	Contract     *AddressProvider  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AddressProviderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressProviderCallerSession struct {
	Contract *AddressProviderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// AddressProviderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressProviderTransactorSession struct {
	Contract     *AddressProviderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// AddressProviderRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddressProviderRaw struct {
	Contract *AddressProvider // Generic contract binding to access the raw methods on
}

// AddressProviderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressProviderCallerRaw struct {
	Contract *AddressProviderCaller // Generic read-only contract binding to access the raw methods on
}

// AddressProviderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressProviderTransactorRaw struct {
	Contract *AddressProviderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddressProvider creates a new instance of AddressProvider, bound to a specific deployed contract.
func NewAddressProvider(address common.Address, backend bind.ContractBackend) (*AddressProvider, error) {
	contract, err := bindAddressProvider(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AddressProvider{AddressProviderCaller: AddressProviderCaller{contract: contract}, AddressProviderTransactor: AddressProviderTransactor{contract: contract}, AddressProviderFilterer: AddressProviderFilterer{contract: contract}}, nil
}

// NewAddressProviderCaller creates a new read-only instance of AddressProvider, bound to a specific deployed contract.
func NewAddressProviderCaller(address common.Address, caller bind.ContractCaller) (*AddressProviderCaller, error) {
	contract, err := bindAddressProvider(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressProviderCaller{contract: contract}, nil
}

// NewAddressProviderTransactor creates a new write-only instance of AddressProvider, bound to a specific deployed contract.
func NewAddressProviderTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressProviderTransactor, error) {
	contract, err := bindAddressProvider(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressProviderTransactor{contract: contract}, nil
}

// NewAddressProviderFilterer creates a new log filterer instance of AddressProvider, bound to a specific deployed contract.
func NewAddressProviderFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressProviderFilterer, error) {
	contract, err := bindAddressProvider(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressProviderFilterer{contract: contract}, nil
}

// bindAddressProvider binds a generic wrapper to an already deployed contract.
func bindAddressProvider(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressProviderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressProvider *AddressProviderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AddressProvider.Contract.AddressProviderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressProvider *AddressProviderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressProvider.Contract.AddressProviderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressProvider *AddressProviderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressProvider.Contract.AddressProviderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressProvider *AddressProviderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AddressProvider.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressProvider *AddressProviderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressProvider.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressProvider *AddressProviderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressProvider.Contract.contract.Transact(opts, method, params...)
}

// ACCOUNTFACTORY is a free data retrieval call binding the contract method 0x05197d10.
//
// Solidity: function ACCOUNT_FACTORY() view returns(bytes32)
func (_AddressProvider *AddressProviderCaller) ACCOUNTFACTORY(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AddressProvider.contract.Call(opts, &out, "ACCOUNT_FACTORY")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ACCOUNTFACTORY is a free data retrieval call binding the contract method 0x05197d10.
//
// Solidity: function ACCOUNT_FACTORY() view returns(bytes32)
func (_AddressProvider *AddressProviderSession) ACCOUNTFACTORY() ([32]byte, error) {
	return _AddressProvider.Contract.ACCOUNTFACTORY(&_AddressProvider.CallOpts)
}

// ACCOUNTFACTORY is a free data retrieval call binding the contract method 0x05197d10.
//
// Solidity: function ACCOUNT_FACTORY() view returns(bytes32)
func (_AddressProvider *AddressProviderCallerSession) ACCOUNTFACTORY() ([32]byte, error) {
	return _AddressProvider.Contract.ACCOUNTFACTORY(&_AddressProvider.CallOpts)
}

// ACL is a free data retrieval call binding the contract method 0x7af53532.
//
// Solidity: function ACL() view returns(bytes32)
func (_AddressProvider *AddressProviderCaller) ACL(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AddressProvider.contract.Call(opts, &out, "ACL")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ACL is a free data retrieval call binding the contract method 0x7af53532.
//
// Solidity: function ACL() view returns(bytes32)
func (_AddressProvider *AddressProviderSession) ACL() ([32]byte, error) {
	return _AddressProvider.Contract.ACL(&_AddressProvider.CallOpts)
}

// ACL is a free data retrieval call binding the contract method 0x7af53532.
//
// Solidity: function ACL() view returns(bytes32)
func (_AddressProvider *AddressProviderCallerSession) ACL() ([32]byte, error) {
	return _AddressProvider.Contract.ACL(&_AddressProvider.CallOpts)
}

// CONTRACTSREGISTER is a free data retrieval call binding the contract method 0xf9366f47.
//
// Solidity: function CONTRACTS_REGISTER() view returns(bytes32)
func (_AddressProvider *AddressProviderCaller) CONTRACTSREGISTER(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AddressProvider.contract.Call(opts, &out, "CONTRACTS_REGISTER")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CONTRACTSREGISTER is a free data retrieval call binding the contract method 0xf9366f47.
//
// Solidity: function CONTRACTS_REGISTER() view returns(bytes32)
func (_AddressProvider *AddressProviderSession) CONTRACTSREGISTER() ([32]byte, error) {
	return _AddressProvider.Contract.CONTRACTSREGISTER(&_AddressProvider.CallOpts)
}

// CONTRACTSREGISTER is a free data retrieval call binding the contract method 0xf9366f47.
//
// Solidity: function CONTRACTS_REGISTER() view returns(bytes32)
func (_AddressProvider *AddressProviderCallerSession) CONTRACTSREGISTER() ([32]byte, error) {
	return _AddressProvider.Contract.CONTRACTSREGISTER(&_AddressProvider.CallOpts)
}

// DATACOMPRESSOR is a free data retrieval call binding the contract method 0x72788be7.
//
// Solidity: function DATA_COMPRESSOR() view returns(bytes32)
func (_AddressProvider *AddressProviderCaller) DATACOMPRESSOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AddressProvider.contract.Call(opts, &out, "DATA_COMPRESSOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DATACOMPRESSOR is a free data retrieval call binding the contract method 0x72788be7.
//
// Solidity: function DATA_COMPRESSOR() view returns(bytes32)
func (_AddressProvider *AddressProviderSession) DATACOMPRESSOR() ([32]byte, error) {
	return _AddressProvider.Contract.DATACOMPRESSOR(&_AddressProvider.CallOpts)
}

// DATACOMPRESSOR is a free data retrieval call binding the contract method 0x72788be7.
//
// Solidity: function DATA_COMPRESSOR() view returns(bytes32)
func (_AddressProvider *AddressProviderCallerSession) DATACOMPRESSOR() ([32]byte, error) {
	return _AddressProvider.Contract.DATACOMPRESSOR(&_AddressProvider.CallOpts)
}

// GEARTOKEN is a free data retrieval call binding the contract method 0x124a6462.
//
// Solidity: function GEAR_TOKEN() view returns(bytes32)
func (_AddressProvider *AddressProviderCaller) GEARTOKEN(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AddressProvider.contract.Call(opts, &out, "GEAR_TOKEN")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GEARTOKEN is a free data retrieval call binding the contract method 0x124a6462.
//
// Solidity: function GEAR_TOKEN() view returns(bytes32)
func (_AddressProvider *AddressProviderSession) GEARTOKEN() ([32]byte, error) {
	return _AddressProvider.Contract.GEARTOKEN(&_AddressProvider.CallOpts)
}

// GEARTOKEN is a free data retrieval call binding the contract method 0x124a6462.
//
// Solidity: function GEAR_TOKEN() view returns(bytes32)
func (_AddressProvider *AddressProviderCallerSession) GEARTOKEN() ([32]byte, error) {
	return _AddressProvider.Contract.GEARTOKEN(&_AddressProvider.CallOpts)
}

// LEVERAGEDACTIONS is a free data retrieval call binding the contract method 0x0e3b95dc.
//
// Solidity: function LEVERAGED_ACTIONS() view returns(bytes32)
func (_AddressProvider *AddressProviderCaller) LEVERAGEDACTIONS(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AddressProvider.contract.Call(opts, &out, "LEVERAGED_ACTIONS")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LEVERAGEDACTIONS is a free data retrieval call binding the contract method 0x0e3b95dc.
//
// Solidity: function LEVERAGED_ACTIONS() view returns(bytes32)
func (_AddressProvider *AddressProviderSession) LEVERAGEDACTIONS() ([32]byte, error) {
	return _AddressProvider.Contract.LEVERAGEDACTIONS(&_AddressProvider.CallOpts)
}

// LEVERAGEDACTIONS is a free data retrieval call binding the contract method 0x0e3b95dc.
//
// Solidity: function LEVERAGED_ACTIONS() view returns(bytes32)
func (_AddressProvider *AddressProviderCallerSession) LEVERAGEDACTIONS() ([32]byte, error) {
	return _AddressProvider.Contract.LEVERAGEDACTIONS(&_AddressProvider.CallOpts)
}

// PRICEORACLE is a free data retrieval call binding the contract method 0x0a19399a.
//
// Solidity: function PRICE_ORACLE() view returns(bytes32)
func (_AddressProvider *AddressProviderCaller) PRICEORACLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AddressProvider.contract.Call(opts, &out, "PRICE_ORACLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PRICEORACLE is a free data retrieval call binding the contract method 0x0a19399a.
//
// Solidity: function PRICE_ORACLE() view returns(bytes32)
func (_AddressProvider *AddressProviderSession) PRICEORACLE() ([32]byte, error) {
	return _AddressProvider.Contract.PRICEORACLE(&_AddressProvider.CallOpts)
}

// PRICEORACLE is a free data retrieval call binding the contract method 0x0a19399a.
//
// Solidity: function PRICE_ORACLE() view returns(bytes32)
func (_AddressProvider *AddressProviderCallerSession) PRICEORACLE() ([32]byte, error) {
	return _AddressProvider.Contract.PRICEORACLE(&_AddressProvider.CallOpts)
}

// TREASURYCONTRACT is a free data retrieval call binding the contract method 0x9e9df2b9.
//
// Solidity: function TREASURY_CONTRACT() view returns(bytes32)
func (_AddressProvider *AddressProviderCaller) TREASURYCONTRACT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AddressProvider.contract.Call(opts, &out, "TREASURY_CONTRACT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TREASURYCONTRACT is a free data retrieval call binding the contract method 0x9e9df2b9.
//
// Solidity: function TREASURY_CONTRACT() view returns(bytes32)
func (_AddressProvider *AddressProviderSession) TREASURYCONTRACT() ([32]byte, error) {
	return _AddressProvider.Contract.TREASURYCONTRACT(&_AddressProvider.CallOpts)
}

// TREASURYCONTRACT is a free data retrieval call binding the contract method 0x9e9df2b9.
//
// Solidity: function TREASURY_CONTRACT() view returns(bytes32)
func (_AddressProvider *AddressProviderCallerSession) TREASURYCONTRACT() ([32]byte, error) {
	return _AddressProvider.Contract.TREASURYCONTRACT(&_AddressProvider.CallOpts)
}

// WETHGATEWAY is a free data retrieval call binding the contract method 0xae5a98ba.
//
// Solidity: function WETH_GATEWAY() view returns(bytes32)
func (_AddressProvider *AddressProviderCaller) WETHGATEWAY(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AddressProvider.contract.Call(opts, &out, "WETH_GATEWAY")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WETHGATEWAY is a free data retrieval call binding the contract method 0xae5a98ba.
//
// Solidity: function WETH_GATEWAY() view returns(bytes32)
func (_AddressProvider *AddressProviderSession) WETHGATEWAY() ([32]byte, error) {
	return _AddressProvider.Contract.WETHGATEWAY(&_AddressProvider.CallOpts)
}

// WETHGATEWAY is a free data retrieval call binding the contract method 0xae5a98ba.
//
// Solidity: function WETH_GATEWAY() view returns(bytes32)
func (_AddressProvider *AddressProviderCallerSession) WETHGATEWAY() ([32]byte, error) {
	return _AddressProvider.Contract.WETHGATEWAY(&_AddressProvider.CallOpts)
}

// WETHTOKEN is a free data retrieval call binding the contract method 0x37d277d4.
//
// Solidity: function WETH_TOKEN() view returns(bytes32)
func (_AddressProvider *AddressProviderCaller) WETHTOKEN(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AddressProvider.contract.Call(opts, &out, "WETH_TOKEN")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WETHTOKEN is a free data retrieval call binding the contract method 0x37d277d4.
//
// Solidity: function WETH_TOKEN() view returns(bytes32)
func (_AddressProvider *AddressProviderSession) WETHTOKEN() ([32]byte, error) {
	return _AddressProvider.Contract.WETHTOKEN(&_AddressProvider.CallOpts)
}

// WETHTOKEN is a free data retrieval call binding the contract method 0x37d277d4.
//
// Solidity: function WETH_TOKEN() view returns(bytes32)
func (_AddressProvider *AddressProviderCallerSession) WETHTOKEN() ([32]byte, error) {
	return _AddressProvider.Contract.WETHTOKEN(&_AddressProvider.CallOpts)
}

// Addresses is a free data retrieval call binding the contract method 0x699f200f.
//
// Solidity: function addresses(bytes32 ) view returns(address)
func (_AddressProvider *AddressProviderCaller) Addresses(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _AddressProvider.contract.Call(opts, &out, "addresses", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Addresses is a free data retrieval call binding the contract method 0x699f200f.
//
// Solidity: function addresses(bytes32 ) view returns(address)
func (_AddressProvider *AddressProviderSession) Addresses(arg0 [32]byte) (common.Address, error) {
	return _AddressProvider.Contract.Addresses(&_AddressProvider.CallOpts, arg0)
}

// Addresses is a free data retrieval call binding the contract method 0x699f200f.
//
// Solidity: function addresses(bytes32 ) view returns(address)
func (_AddressProvider *AddressProviderCallerSession) Addresses(arg0 [32]byte) (common.Address, error) {
	return _AddressProvider.Contract.Addresses(&_AddressProvider.CallOpts, arg0)
}

// GetACL is a free data retrieval call binding the contract method 0x08737695.
//
// Solidity: function getACL() view returns(address)
func (_AddressProvider *AddressProviderCaller) GetACL(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AddressProvider.contract.Call(opts, &out, "getACL")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetACL is a free data retrieval call binding the contract method 0x08737695.
//
// Solidity: function getACL() view returns(address)
func (_AddressProvider *AddressProviderSession) GetACL() (common.Address, error) {
	return _AddressProvider.Contract.GetACL(&_AddressProvider.CallOpts)
}

// GetACL is a free data retrieval call binding the contract method 0x08737695.
//
// Solidity: function getACL() view returns(address)
func (_AddressProvider *AddressProviderCallerSession) GetACL() (common.Address, error) {
	return _AddressProvider.Contract.GetACL(&_AddressProvider.CallOpts)
}

// GetAccountFactory is a free data retrieval call binding the contract method 0x9068a868.
//
// Solidity: function getAccountFactory() view returns(address)
func (_AddressProvider *AddressProviderCaller) GetAccountFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AddressProvider.contract.Call(opts, &out, "getAccountFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAccountFactory is a free data retrieval call binding the contract method 0x9068a868.
//
// Solidity: function getAccountFactory() view returns(address)
func (_AddressProvider *AddressProviderSession) GetAccountFactory() (common.Address, error) {
	return _AddressProvider.Contract.GetAccountFactory(&_AddressProvider.CallOpts)
}

// GetAccountFactory is a free data retrieval call binding the contract method 0x9068a868.
//
// Solidity: function getAccountFactory() view returns(address)
func (_AddressProvider *AddressProviderCallerSession) GetAccountFactory() (common.Address, error) {
	return _AddressProvider.Contract.GetAccountFactory(&_AddressProvider.CallOpts)
}

// GetContractsRegister is a free data retrieval call binding the contract method 0xc513c9bb.
//
// Solidity: function getContractsRegister() view returns(address)
func (_AddressProvider *AddressProviderCaller) GetContractsRegister(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AddressProvider.contract.Call(opts, &out, "getContractsRegister")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetContractsRegister is a free data retrieval call binding the contract method 0xc513c9bb.
//
// Solidity: function getContractsRegister() view returns(address)
func (_AddressProvider *AddressProviderSession) GetContractsRegister() (common.Address, error) {
	return _AddressProvider.Contract.GetContractsRegister(&_AddressProvider.CallOpts)
}

// GetContractsRegister is a free data retrieval call binding the contract method 0xc513c9bb.
//
// Solidity: function getContractsRegister() view returns(address)
func (_AddressProvider *AddressProviderCallerSession) GetContractsRegister() (common.Address, error) {
	return _AddressProvider.Contract.GetContractsRegister(&_AddressProvider.CallOpts)
}

// GetDataCompressor is a free data retrieval call binding the contract method 0x060678c2.
//
// Solidity: function getDataCompressor() view returns(address)
func (_AddressProvider *AddressProviderCaller) GetDataCompressor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AddressProvider.contract.Call(opts, &out, "getDataCompressor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetDataCompressor is a free data retrieval call binding the contract method 0x060678c2.
//
// Solidity: function getDataCompressor() view returns(address)
func (_AddressProvider *AddressProviderSession) GetDataCompressor() (common.Address, error) {
	return _AddressProvider.Contract.GetDataCompressor(&_AddressProvider.CallOpts)
}

// GetDataCompressor is a free data retrieval call binding the contract method 0x060678c2.
//
// Solidity: function getDataCompressor() view returns(address)
func (_AddressProvider *AddressProviderCallerSession) GetDataCompressor() (common.Address, error) {
	return _AddressProvider.Contract.GetDataCompressor(&_AddressProvider.CallOpts)
}

// GetGearToken is a free data retrieval call binding the contract method 0xaffd9243.
//
// Solidity: function getGearToken() view returns(address)
func (_AddressProvider *AddressProviderCaller) GetGearToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AddressProvider.contract.Call(opts, &out, "getGearToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetGearToken is a free data retrieval call binding the contract method 0xaffd9243.
//
// Solidity: function getGearToken() view returns(address)
func (_AddressProvider *AddressProviderSession) GetGearToken() (common.Address, error) {
	return _AddressProvider.Contract.GetGearToken(&_AddressProvider.CallOpts)
}

// GetGearToken is a free data retrieval call binding the contract method 0xaffd9243.
//
// Solidity: function getGearToken() view returns(address)
func (_AddressProvider *AddressProviderCallerSession) GetGearToken() (common.Address, error) {
	return _AddressProvider.Contract.GetGearToken(&_AddressProvider.CallOpts)
}

// GetLeveragedActions is a free data retrieval call binding the contract method 0x44b88563.
//
// Solidity: function getLeveragedActions() view returns(address)
func (_AddressProvider *AddressProviderCaller) GetLeveragedActions(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AddressProvider.contract.Call(opts, &out, "getLeveragedActions")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLeveragedActions is a free data retrieval call binding the contract method 0x44b88563.
//
// Solidity: function getLeveragedActions() view returns(address)
func (_AddressProvider *AddressProviderSession) GetLeveragedActions() (common.Address, error) {
	return _AddressProvider.Contract.GetLeveragedActions(&_AddressProvider.CallOpts)
}

// GetLeveragedActions is a free data retrieval call binding the contract method 0x44b88563.
//
// Solidity: function getLeveragedActions() view returns(address)
func (_AddressProvider *AddressProviderCallerSession) GetLeveragedActions() (common.Address, error) {
	return _AddressProvider.Contract.GetLeveragedActions(&_AddressProvider.CallOpts)
}

// GetPriceOracle is a free data retrieval call binding the contract method 0xfca513a8.
//
// Solidity: function getPriceOracle() view returns(address)
func (_AddressProvider *AddressProviderCaller) GetPriceOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AddressProvider.contract.Call(opts, &out, "getPriceOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPriceOracle is a free data retrieval call binding the contract method 0xfca513a8.
//
// Solidity: function getPriceOracle() view returns(address)
func (_AddressProvider *AddressProviderSession) GetPriceOracle() (common.Address, error) {
	return _AddressProvider.Contract.GetPriceOracle(&_AddressProvider.CallOpts)
}

// GetPriceOracle is a free data retrieval call binding the contract method 0xfca513a8.
//
// Solidity: function getPriceOracle() view returns(address)
func (_AddressProvider *AddressProviderCallerSession) GetPriceOracle() (common.Address, error) {
	return _AddressProvider.Contract.GetPriceOracle(&_AddressProvider.CallOpts)
}

// GetTreasuryContract is a free data retrieval call binding the contract method 0x26c74fc3.
//
// Solidity: function getTreasuryContract() view returns(address)
func (_AddressProvider *AddressProviderCaller) GetTreasuryContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AddressProvider.contract.Call(opts, &out, "getTreasuryContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTreasuryContract is a free data retrieval call binding the contract method 0x26c74fc3.
//
// Solidity: function getTreasuryContract() view returns(address)
func (_AddressProvider *AddressProviderSession) GetTreasuryContract() (common.Address, error) {
	return _AddressProvider.Contract.GetTreasuryContract(&_AddressProvider.CallOpts)
}

// GetTreasuryContract is a free data retrieval call binding the contract method 0x26c74fc3.
//
// Solidity: function getTreasuryContract() view returns(address)
func (_AddressProvider *AddressProviderCallerSession) GetTreasuryContract() (common.Address, error) {
	return _AddressProvider.Contract.GetTreasuryContract(&_AddressProvider.CallOpts)
}

// GetWETHGateway is a free data retrieval call binding the contract method 0x77532ed9.
//
// Solidity: function getWETHGateway() view returns(address)
func (_AddressProvider *AddressProviderCaller) GetWETHGateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AddressProvider.contract.Call(opts, &out, "getWETHGateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetWETHGateway is a free data retrieval call binding the contract method 0x77532ed9.
//
// Solidity: function getWETHGateway() view returns(address)
func (_AddressProvider *AddressProviderSession) GetWETHGateway() (common.Address, error) {
	return _AddressProvider.Contract.GetWETHGateway(&_AddressProvider.CallOpts)
}

// GetWETHGateway is a free data retrieval call binding the contract method 0x77532ed9.
//
// Solidity: function getWETHGateway() view returns(address)
func (_AddressProvider *AddressProviderCallerSession) GetWETHGateway() (common.Address, error) {
	return _AddressProvider.Contract.GetWETHGateway(&_AddressProvider.CallOpts)
}

// GetWethToken is a free data retrieval call binding the contract method 0x4c252f91.
//
// Solidity: function getWethToken() view returns(address)
func (_AddressProvider *AddressProviderCaller) GetWethToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AddressProvider.contract.Call(opts, &out, "getWethToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetWethToken is a free data retrieval call binding the contract method 0x4c252f91.
//
// Solidity: function getWethToken() view returns(address)
func (_AddressProvider *AddressProviderSession) GetWethToken() (common.Address, error) {
	return _AddressProvider.Contract.GetWethToken(&_AddressProvider.CallOpts)
}

// GetWethToken is a free data retrieval call binding the contract method 0x4c252f91.
//
// Solidity: function getWethToken() view returns(address)
func (_AddressProvider *AddressProviderCallerSession) GetWethToken() (common.Address, error) {
	return _AddressProvider.Contract.GetWethToken(&_AddressProvider.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AddressProvider *AddressProviderCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AddressProvider.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AddressProvider *AddressProviderSession) Owner() (common.Address, error) {
	return _AddressProvider.Contract.Owner(&_AddressProvider.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AddressProvider *AddressProviderCallerSession) Owner() (common.Address, error) {
	return _AddressProvider.Contract.Owner(&_AddressProvider.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AddressProvider *AddressProviderTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressProvider.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AddressProvider *AddressProviderSession) RenounceOwnership() (*types.Transaction, error) {
	return _AddressProvider.Contract.RenounceOwnership(&_AddressProvider.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AddressProvider *AddressProviderTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AddressProvider.Contract.RenounceOwnership(&_AddressProvider.TransactOpts)
}

// SetACL is a paid mutator transaction binding the contract method 0x76aad605.
//
// Solidity: function setACL(address _address) returns()
func (_AddressProvider *AddressProviderTransactor) SetACL(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _AddressProvider.contract.Transact(opts, "setACL", _address)
}

// SetACL is a paid mutator transaction binding the contract method 0x76aad605.
//
// Solidity: function setACL(address _address) returns()
func (_AddressProvider *AddressProviderSession) SetACL(_address common.Address) (*types.Transaction, error) {
	return _AddressProvider.Contract.SetACL(&_AddressProvider.TransactOpts, _address)
}

// SetACL is a paid mutator transaction binding the contract method 0x76aad605.
//
// Solidity: function setACL(address _address) returns()
func (_AddressProvider *AddressProviderTransactorSession) SetACL(_address common.Address) (*types.Transaction, error) {
	return _AddressProvider.Contract.SetACL(&_AddressProvider.TransactOpts, _address)
}

// SetAccountFactory is a paid mutator transaction binding the contract method 0xaddc1a76.
//
// Solidity: function setAccountFactory(address _address) returns()
func (_AddressProvider *AddressProviderTransactor) SetAccountFactory(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _AddressProvider.contract.Transact(opts, "setAccountFactory", _address)
}

// SetAccountFactory is a paid mutator transaction binding the contract method 0xaddc1a76.
//
// Solidity: function setAccountFactory(address _address) returns()
func (_AddressProvider *AddressProviderSession) SetAccountFactory(_address common.Address) (*types.Transaction, error) {
	return _AddressProvider.Contract.SetAccountFactory(&_AddressProvider.TransactOpts, _address)
}

// SetAccountFactory is a paid mutator transaction binding the contract method 0xaddc1a76.
//
// Solidity: function setAccountFactory(address _address) returns()
func (_AddressProvider *AddressProviderTransactorSession) SetAccountFactory(_address common.Address) (*types.Transaction, error) {
	return _AddressProvider.Contract.SetAccountFactory(&_AddressProvider.TransactOpts, _address)
}

// SetContractsRegister is a paid mutator transaction binding the contract method 0xce3c4ae4.
//
// Solidity: function setContractsRegister(address _address) returns()
func (_AddressProvider *AddressProviderTransactor) SetContractsRegister(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _AddressProvider.contract.Transact(opts, "setContractsRegister", _address)
}

// SetContractsRegister is a paid mutator transaction binding the contract method 0xce3c4ae4.
//
// Solidity: function setContractsRegister(address _address) returns()
func (_AddressProvider *AddressProviderSession) SetContractsRegister(_address common.Address) (*types.Transaction, error) {
	return _AddressProvider.Contract.SetContractsRegister(&_AddressProvider.TransactOpts, _address)
}

// SetContractsRegister is a paid mutator transaction binding the contract method 0xce3c4ae4.
//
// Solidity: function setContractsRegister(address _address) returns()
func (_AddressProvider *AddressProviderTransactorSession) SetContractsRegister(_address common.Address) (*types.Transaction, error) {
	return _AddressProvider.Contract.SetContractsRegister(&_AddressProvider.TransactOpts, _address)
}

// SetDataCompressor is a paid mutator transaction binding the contract method 0xc5120b39.
//
// Solidity: function setDataCompressor(address _address) returns()
func (_AddressProvider *AddressProviderTransactor) SetDataCompressor(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _AddressProvider.contract.Transact(opts, "setDataCompressor", _address)
}

// SetDataCompressor is a paid mutator transaction binding the contract method 0xc5120b39.
//
// Solidity: function setDataCompressor(address _address) returns()
func (_AddressProvider *AddressProviderSession) SetDataCompressor(_address common.Address) (*types.Transaction, error) {
	return _AddressProvider.Contract.SetDataCompressor(&_AddressProvider.TransactOpts, _address)
}

// SetDataCompressor is a paid mutator transaction binding the contract method 0xc5120b39.
//
// Solidity: function setDataCompressor(address _address) returns()
func (_AddressProvider *AddressProviderTransactorSession) SetDataCompressor(_address common.Address) (*types.Transaction, error) {
	return _AddressProvider.Contract.SetDataCompressor(&_AddressProvider.TransactOpts, _address)
}

// SetGearToken is a paid mutator transaction binding the contract method 0xbcaead98.
//
// Solidity: function setGearToken(address _address) returns()
func (_AddressProvider *AddressProviderTransactor) SetGearToken(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _AddressProvider.contract.Transact(opts, "setGearToken", _address)
}

// SetGearToken is a paid mutator transaction binding the contract method 0xbcaead98.
//
// Solidity: function setGearToken(address _address) returns()
func (_AddressProvider *AddressProviderSession) SetGearToken(_address common.Address) (*types.Transaction, error) {
	return _AddressProvider.Contract.SetGearToken(&_AddressProvider.TransactOpts, _address)
}

// SetGearToken is a paid mutator transaction binding the contract method 0xbcaead98.
//
// Solidity: function setGearToken(address _address) returns()
func (_AddressProvider *AddressProviderTransactorSession) SetGearToken(_address common.Address) (*types.Transaction, error) {
	return _AddressProvider.Contract.SetGearToken(&_AddressProvider.TransactOpts, _address)
}

// SetLeveragedActions is a paid mutator transaction binding the contract method 0x7b6757ff.
//
// Solidity: function setLeveragedActions(address _address) returns()
func (_AddressProvider *AddressProviderTransactor) SetLeveragedActions(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _AddressProvider.contract.Transact(opts, "setLeveragedActions", _address)
}

// SetLeveragedActions is a paid mutator transaction binding the contract method 0x7b6757ff.
//
// Solidity: function setLeveragedActions(address _address) returns()
func (_AddressProvider *AddressProviderSession) SetLeveragedActions(_address common.Address) (*types.Transaction, error) {
	return _AddressProvider.Contract.SetLeveragedActions(&_AddressProvider.TransactOpts, _address)
}

// SetLeveragedActions is a paid mutator transaction binding the contract method 0x7b6757ff.
//
// Solidity: function setLeveragedActions(address _address) returns()
func (_AddressProvider *AddressProviderTransactorSession) SetLeveragedActions(_address common.Address) (*types.Transaction, error) {
	return _AddressProvider.Contract.SetLeveragedActions(&_AddressProvider.TransactOpts, _address)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address _address) returns()
func (_AddressProvider *AddressProviderTransactor) SetPriceOracle(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _AddressProvider.contract.Transact(opts, "setPriceOracle", _address)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address _address) returns()
func (_AddressProvider *AddressProviderSession) SetPriceOracle(_address common.Address) (*types.Transaction, error) {
	return _AddressProvider.Contract.SetPriceOracle(&_AddressProvider.TransactOpts, _address)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address _address) returns()
func (_AddressProvider *AddressProviderTransactorSession) SetPriceOracle(_address common.Address) (*types.Transaction, error) {
	return _AddressProvider.Contract.SetPriceOracle(&_AddressProvider.TransactOpts, _address)
}

// SetTreasuryContract is a paid mutator transaction binding the contract method 0x1ed65110.
//
// Solidity: function setTreasuryContract(address _address) returns()
func (_AddressProvider *AddressProviderTransactor) SetTreasuryContract(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _AddressProvider.contract.Transact(opts, "setTreasuryContract", _address)
}

// SetTreasuryContract is a paid mutator transaction binding the contract method 0x1ed65110.
//
// Solidity: function setTreasuryContract(address _address) returns()
func (_AddressProvider *AddressProviderSession) SetTreasuryContract(_address common.Address) (*types.Transaction, error) {
	return _AddressProvider.Contract.SetTreasuryContract(&_AddressProvider.TransactOpts, _address)
}

// SetTreasuryContract is a paid mutator transaction binding the contract method 0x1ed65110.
//
// Solidity: function setTreasuryContract(address _address) returns()
func (_AddressProvider *AddressProviderTransactorSession) SetTreasuryContract(_address common.Address) (*types.Transaction, error) {
	return _AddressProvider.Contract.SetTreasuryContract(&_AddressProvider.TransactOpts, _address)
}

// SetWETHGateway is a paid mutator transaction binding the contract method 0x21da5837.
//
// Solidity: function setWETHGateway(address _address) returns()
func (_AddressProvider *AddressProviderTransactor) SetWETHGateway(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _AddressProvider.contract.Transact(opts, "setWETHGateway", _address)
}

// SetWETHGateway is a paid mutator transaction binding the contract method 0x21da5837.
//
// Solidity: function setWETHGateway(address _address) returns()
func (_AddressProvider *AddressProviderSession) SetWETHGateway(_address common.Address) (*types.Transaction, error) {
	return _AddressProvider.Contract.SetWETHGateway(&_AddressProvider.TransactOpts, _address)
}

// SetWETHGateway is a paid mutator transaction binding the contract method 0x21da5837.
//
// Solidity: function setWETHGateway(address _address) returns()
func (_AddressProvider *AddressProviderTransactorSession) SetWETHGateway(_address common.Address) (*types.Transaction, error) {
	return _AddressProvider.Contract.SetWETHGateway(&_AddressProvider.TransactOpts, _address)
}

// SetWethToken is a paid mutator transaction binding the contract method 0x86e09c08.
//
// Solidity: function setWethToken(address _address) returns()
func (_AddressProvider *AddressProviderTransactor) SetWethToken(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _AddressProvider.contract.Transact(opts, "setWethToken", _address)
}

// SetWethToken is a paid mutator transaction binding the contract method 0x86e09c08.
//
// Solidity: function setWethToken(address _address) returns()
func (_AddressProvider *AddressProviderSession) SetWethToken(_address common.Address) (*types.Transaction, error) {
	return _AddressProvider.Contract.SetWethToken(&_AddressProvider.TransactOpts, _address)
}

// SetWethToken is a paid mutator transaction binding the contract method 0x86e09c08.
//
// Solidity: function setWethToken(address _address) returns()
func (_AddressProvider *AddressProviderTransactorSession) SetWethToken(_address common.Address) (*types.Transaction, error) {
	return _AddressProvider.Contract.SetWethToken(&_AddressProvider.TransactOpts, _address)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AddressProvider *AddressProviderTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AddressProvider.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AddressProvider *AddressProviderSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AddressProvider.Contract.TransferOwnership(&_AddressProvider.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AddressProvider *AddressProviderTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AddressProvider.Contract.TransferOwnership(&_AddressProvider.TransactOpts, newOwner)
}

// AddressProviderAddressSetIterator is returned from FilterAddressSet and is used to iterate over the raw logs and unpacked data for AddressSet events raised by the AddressProvider contract.
type AddressProviderAddressSetIterator struct {
	Event *AddressProviderAddressSet // Event containing the contract specifics and raw log

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
func (it *AddressProviderAddressSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AddressProviderAddressSet)
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
		it.Event = new(AddressProviderAddressSet)
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
func (it *AddressProviderAddressSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AddressProviderAddressSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AddressProviderAddressSet represents a AddressSet event raised by the AddressProvider contract.
type AddressProviderAddressSet struct {
	Service    [32]byte
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAddressSet is a free log retrieval operation binding the contract event 0xb37614c7d254ea8d16eb81fa11dddaeb266aa8ba4917980859c7740aff30c691.
//
// Solidity: event AddressSet(bytes32 indexed service, address indexed newAddress)
func (_AddressProvider *AddressProviderFilterer) FilterAddressSet(opts *bind.FilterOpts, service [][32]byte, newAddress []common.Address) (*AddressProviderAddressSetIterator, error) {

	var serviceRule []interface{}
	for _, serviceItem := range service {
		serviceRule = append(serviceRule, serviceItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AddressProvider.contract.FilterLogs(opts, "AddressSet", serviceRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return &AddressProviderAddressSetIterator{contract: _AddressProvider.contract, event: "AddressSet", logs: logs, sub: sub}, nil
}

// WatchAddressSet is a free log subscription operation binding the contract event 0xb37614c7d254ea8d16eb81fa11dddaeb266aa8ba4917980859c7740aff30c691.
//
// Solidity: event AddressSet(bytes32 indexed service, address indexed newAddress)
func (_AddressProvider *AddressProviderFilterer) WatchAddressSet(opts *bind.WatchOpts, sink chan<- *AddressProviderAddressSet, service [][32]byte, newAddress []common.Address) (event.Subscription, error) {

	var serviceRule []interface{}
	for _, serviceItem := range service {
		serviceRule = append(serviceRule, serviceItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AddressProvider.contract.WatchLogs(opts, "AddressSet", serviceRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AddressProviderAddressSet)
				if err := _AddressProvider.contract.UnpackLog(event, "AddressSet", log); err != nil {
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

// ParseAddressSet is a log parse operation binding the contract event 0xb37614c7d254ea8d16eb81fa11dddaeb266aa8ba4917980859c7740aff30c691.
//
// Solidity: event AddressSet(bytes32 indexed service, address indexed newAddress)
func (_AddressProvider *AddressProviderFilterer) ParseAddressSet(log types.Log) (*AddressProviderAddressSet, error) {
	event := new(AddressProviderAddressSet)
	if err := _AddressProvider.contract.UnpackLog(event, "AddressSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AddressProviderOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AddressProvider contract.
type AddressProviderOwnershipTransferredIterator struct {
	Event *AddressProviderOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AddressProviderOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AddressProviderOwnershipTransferred)
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
		it.Event = new(AddressProviderOwnershipTransferred)
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
func (it *AddressProviderOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AddressProviderOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AddressProviderOwnershipTransferred represents a OwnershipTransferred event raised by the AddressProvider contract.
type AddressProviderOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AddressProvider *AddressProviderFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AddressProviderOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AddressProvider.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AddressProviderOwnershipTransferredIterator{contract: _AddressProvider.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AddressProvider *AddressProviderFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AddressProviderOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AddressProvider.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AddressProviderOwnershipTransferred)
				if err := _AddressProvider.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AddressProvider *AddressProviderFilterer) ParseOwnershipTransferred(log types.Log) (*AddressProviderOwnershipTransferred, error) {
	event := new(AddressProviderOwnershipTransferred)
	if err := _AddressProvider.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
