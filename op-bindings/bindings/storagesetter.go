// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// StorageSetterMetaData contains all meta data concerning the StorageSetter contract.
var StorageSetterMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"getAddress\",\"inputs\":[{\"name\":\"_slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBool\",\"inputs\":[{\"name\":\"_slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"value_\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBytes32\",\"inputs\":[{\"name\":\"_slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUint\",\"inputs\":[{\"name\":\"_slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setAddress\",\"inputs\":[{\"name\":\"_slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_address\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setBool\",\"inputs\":[{\"name\":\"_slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_value\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setBytes32\",\"inputs\":[{\"name\":\"_slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_value\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setUint\",\"inputs\":[{\"name\":\"_slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"}]",
	Bin: "0x608060405234801561001057600080fd5b506102d6806100206000396000f3fe608060405234801561001057600080fd5b50600436106100a35760003560e01c8063a6ed563e11610076578063bd02d0f51161005b578063bd02d0f514610160578063ca446dd91461018a578063e2a4853a146100e557600080fd5b8063a6ed563e14610160578063abfdcced1461017c57600080fd5b806321f8a721146100a85780634e91db08146100e557806354fd4d50146100f95780637ae1cfca14610142575b600080fd5b6100bb6100b63660046101a8565b610198565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b6100f76100f33660046101c1565b9055565b005b6101356040518060400160405280600581526020017f312e302e3000000000000000000000000000000000000000000000000000000081525081565b6040516100dc91906101e3565b6101506100b63660046101a8565b60405190151581526020016100dc565b61016e6100b63660046101a8565b6040519081526020016100dc565b6100f76100f3366004610256565b6100f76100f336600461028b565b60006101a2825490565b92915050565b6000602082840312156101ba57600080fd5b5035919050565b600080604083850312156101d457600080fd5b50508035926020909101359150565b600060208083528351808285015260005b81811015610210578581018301518582016040015282016101f4565b81811115610222576000604083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016929092016040019392505050565b6000806040838503121561026957600080fd5b823591506020830135801515811461028057600080fd5b809150509250929050565b6000806040838503121561029e57600080fd5b82359150602083013573ffffffffffffffffffffffffffffffffffffffff8116811461028057600080fdfea164736f6c634300080f000a",
}

// StorageSetterABI is the input ABI used to generate the binding from.
// Deprecated: Use StorageSetterMetaData.ABI instead.
var StorageSetterABI = StorageSetterMetaData.ABI

// StorageSetterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StorageSetterMetaData.Bin instead.
var StorageSetterBin = StorageSetterMetaData.Bin

// DeployStorageSetter deploys a new Ethereum contract, binding an instance of StorageSetter to it.
func DeployStorageSetter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StorageSetter, error) {
	parsed, err := StorageSetterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StorageSetterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StorageSetter{StorageSetterCaller: StorageSetterCaller{contract: contract}, StorageSetterTransactor: StorageSetterTransactor{contract: contract}, StorageSetterFilterer: StorageSetterFilterer{contract: contract}}, nil
}

// StorageSetter is an auto generated Go binding around an Ethereum contract.
type StorageSetter struct {
	StorageSetterCaller     // Read-only binding to the contract
	StorageSetterTransactor // Write-only binding to the contract
	StorageSetterFilterer   // Log filterer for contract events
}

// StorageSetterCaller is an auto generated read-only Go binding around an Ethereum contract.
type StorageSetterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageSetterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StorageSetterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageSetterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StorageSetterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageSetterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StorageSetterSession struct {
	Contract     *StorageSetter    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StorageSetterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StorageSetterCallerSession struct {
	Contract *StorageSetterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// StorageSetterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StorageSetterTransactorSession struct {
	Contract     *StorageSetterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// StorageSetterRaw is an auto generated low-level Go binding around an Ethereum contract.
type StorageSetterRaw struct {
	Contract *StorageSetter // Generic contract binding to access the raw methods on
}

// StorageSetterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StorageSetterCallerRaw struct {
	Contract *StorageSetterCaller // Generic read-only contract binding to access the raw methods on
}

// StorageSetterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StorageSetterTransactorRaw struct {
	Contract *StorageSetterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStorageSetter creates a new instance of StorageSetter, bound to a specific deployed contract.
func NewStorageSetter(address common.Address, backend bind.ContractBackend) (*StorageSetter, error) {
	contract, err := bindStorageSetter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StorageSetter{StorageSetterCaller: StorageSetterCaller{contract: contract}, StorageSetterTransactor: StorageSetterTransactor{contract: contract}, StorageSetterFilterer: StorageSetterFilterer{contract: contract}}, nil
}

// NewStorageSetterCaller creates a new read-only instance of StorageSetter, bound to a specific deployed contract.
func NewStorageSetterCaller(address common.Address, caller bind.ContractCaller) (*StorageSetterCaller, error) {
	contract, err := bindStorageSetter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StorageSetterCaller{contract: contract}, nil
}

// NewStorageSetterTransactor creates a new write-only instance of StorageSetter, bound to a specific deployed contract.
func NewStorageSetterTransactor(address common.Address, transactor bind.ContractTransactor) (*StorageSetterTransactor, error) {
	contract, err := bindStorageSetter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StorageSetterTransactor{contract: contract}, nil
}

// NewStorageSetterFilterer creates a new log filterer instance of StorageSetter, bound to a specific deployed contract.
func NewStorageSetterFilterer(address common.Address, filterer bind.ContractFilterer) (*StorageSetterFilterer, error) {
	contract, err := bindStorageSetter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StorageSetterFilterer{contract: contract}, nil
}

// bindStorageSetter binds a generic wrapper to an already deployed contract.
func bindStorageSetter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StorageSetterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StorageSetter *StorageSetterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StorageSetter.Contract.StorageSetterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StorageSetter *StorageSetterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StorageSetter.Contract.StorageSetterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StorageSetter *StorageSetterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StorageSetter.Contract.StorageSetterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StorageSetter *StorageSetterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StorageSetter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StorageSetter *StorageSetterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StorageSetter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StorageSetter *StorageSetterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StorageSetter.Contract.contract.Transact(opts, method, params...)
}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 _slot) view returns(address)
func (_StorageSetter *StorageSetterCaller) GetAddress(opts *bind.CallOpts, _slot [32]byte) (common.Address, error) {
	var out []interface{}
	err := _StorageSetter.contract.Call(opts, &out, "getAddress", _slot)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 _slot) view returns(address)
func (_StorageSetter *StorageSetterSession) GetAddress(_slot [32]byte) (common.Address, error) {
	return _StorageSetter.Contract.GetAddress(&_StorageSetter.CallOpts, _slot)
}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 _slot) view returns(address)
func (_StorageSetter *StorageSetterCallerSession) GetAddress(_slot [32]byte) (common.Address, error) {
	return _StorageSetter.Contract.GetAddress(&_StorageSetter.CallOpts, _slot)
}

// GetBool is a free data retrieval call binding the contract method 0x7ae1cfca.
//
// Solidity: function getBool(bytes32 _slot) view returns(bool value_)
func (_StorageSetter *StorageSetterCaller) GetBool(opts *bind.CallOpts, _slot [32]byte) (bool, error) {
	var out []interface{}
	err := _StorageSetter.contract.Call(opts, &out, "getBool", _slot)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetBool is a free data retrieval call binding the contract method 0x7ae1cfca.
//
// Solidity: function getBool(bytes32 _slot) view returns(bool value_)
func (_StorageSetter *StorageSetterSession) GetBool(_slot [32]byte) (bool, error) {
	return _StorageSetter.Contract.GetBool(&_StorageSetter.CallOpts, _slot)
}

// GetBool is a free data retrieval call binding the contract method 0x7ae1cfca.
//
// Solidity: function getBool(bytes32 _slot) view returns(bool value_)
func (_StorageSetter *StorageSetterCallerSession) GetBool(_slot [32]byte) (bool, error) {
	return _StorageSetter.Contract.GetBool(&_StorageSetter.CallOpts, _slot)
}

// GetBytes32 is a free data retrieval call binding the contract method 0xa6ed563e.
//
// Solidity: function getBytes32(bytes32 _slot) view returns(bytes32)
func (_StorageSetter *StorageSetterCaller) GetBytes32(opts *bind.CallOpts, _slot [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _StorageSetter.contract.Call(opts, &out, "getBytes32", _slot)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetBytes32 is a free data retrieval call binding the contract method 0xa6ed563e.
//
// Solidity: function getBytes32(bytes32 _slot) view returns(bytes32)
func (_StorageSetter *StorageSetterSession) GetBytes32(_slot [32]byte) ([32]byte, error) {
	return _StorageSetter.Contract.GetBytes32(&_StorageSetter.CallOpts, _slot)
}

// GetBytes32 is a free data retrieval call binding the contract method 0xa6ed563e.
//
// Solidity: function getBytes32(bytes32 _slot) view returns(bytes32)
func (_StorageSetter *StorageSetterCallerSession) GetBytes32(_slot [32]byte) ([32]byte, error) {
	return _StorageSetter.Contract.GetBytes32(&_StorageSetter.CallOpts, _slot)
}

// GetUint is a free data retrieval call binding the contract method 0xbd02d0f5.
//
// Solidity: function getUint(bytes32 _slot) view returns(uint256)
func (_StorageSetter *StorageSetterCaller) GetUint(opts *bind.CallOpts, _slot [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _StorageSetter.contract.Call(opts, &out, "getUint", _slot)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUint is a free data retrieval call binding the contract method 0xbd02d0f5.
//
// Solidity: function getUint(bytes32 _slot) view returns(uint256)
func (_StorageSetter *StorageSetterSession) GetUint(_slot [32]byte) (*big.Int, error) {
	return _StorageSetter.Contract.GetUint(&_StorageSetter.CallOpts, _slot)
}

// GetUint is a free data retrieval call binding the contract method 0xbd02d0f5.
//
// Solidity: function getUint(bytes32 _slot) view returns(uint256)
func (_StorageSetter *StorageSetterCallerSession) GetUint(_slot [32]byte) (*big.Int, error) {
	return _StorageSetter.Contract.GetUint(&_StorageSetter.CallOpts, _slot)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_StorageSetter *StorageSetterCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StorageSetter.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_StorageSetter *StorageSetterSession) Version() (string, error) {
	return _StorageSetter.Contract.Version(&_StorageSetter.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_StorageSetter *StorageSetterCallerSession) Version() (string, error) {
	return _StorageSetter.Contract.Version(&_StorageSetter.CallOpts)
}

// SetAddress is a paid mutator transaction binding the contract method 0xca446dd9.
//
// Solidity: function setAddress(bytes32 _slot, address _address) returns()
func (_StorageSetter *StorageSetterTransactor) SetAddress(opts *bind.TransactOpts, _slot [32]byte, _address common.Address) (*types.Transaction, error) {
	return _StorageSetter.contract.Transact(opts, "setAddress", _slot, _address)
}

// SetAddress is a paid mutator transaction binding the contract method 0xca446dd9.
//
// Solidity: function setAddress(bytes32 _slot, address _address) returns()
func (_StorageSetter *StorageSetterSession) SetAddress(_slot [32]byte, _address common.Address) (*types.Transaction, error) {
	return _StorageSetter.Contract.SetAddress(&_StorageSetter.TransactOpts, _slot, _address)
}

// SetAddress is a paid mutator transaction binding the contract method 0xca446dd9.
//
// Solidity: function setAddress(bytes32 _slot, address _address) returns()
func (_StorageSetter *StorageSetterTransactorSession) SetAddress(_slot [32]byte, _address common.Address) (*types.Transaction, error) {
	return _StorageSetter.Contract.SetAddress(&_StorageSetter.TransactOpts, _slot, _address)
}

// SetBool is a paid mutator transaction binding the contract method 0xabfdcced.
//
// Solidity: function setBool(bytes32 _slot, bool _value) returns()
func (_StorageSetter *StorageSetterTransactor) SetBool(opts *bind.TransactOpts, _slot [32]byte, _value bool) (*types.Transaction, error) {
	return _StorageSetter.contract.Transact(opts, "setBool", _slot, _value)
}

// SetBool is a paid mutator transaction binding the contract method 0xabfdcced.
//
// Solidity: function setBool(bytes32 _slot, bool _value) returns()
func (_StorageSetter *StorageSetterSession) SetBool(_slot [32]byte, _value bool) (*types.Transaction, error) {
	return _StorageSetter.Contract.SetBool(&_StorageSetter.TransactOpts, _slot, _value)
}

// SetBool is a paid mutator transaction binding the contract method 0xabfdcced.
//
// Solidity: function setBool(bytes32 _slot, bool _value) returns()
func (_StorageSetter *StorageSetterTransactorSession) SetBool(_slot [32]byte, _value bool) (*types.Transaction, error) {
	return _StorageSetter.Contract.SetBool(&_StorageSetter.TransactOpts, _slot, _value)
}

// SetBytes32 is a paid mutator transaction binding the contract method 0x4e91db08.
//
// Solidity: function setBytes32(bytes32 _slot, bytes32 _value) returns()
func (_StorageSetter *StorageSetterTransactor) SetBytes32(opts *bind.TransactOpts, _slot [32]byte, _value [32]byte) (*types.Transaction, error) {
	return _StorageSetter.contract.Transact(opts, "setBytes32", _slot, _value)
}

// SetBytes32 is a paid mutator transaction binding the contract method 0x4e91db08.
//
// Solidity: function setBytes32(bytes32 _slot, bytes32 _value) returns()
func (_StorageSetter *StorageSetterSession) SetBytes32(_slot [32]byte, _value [32]byte) (*types.Transaction, error) {
	return _StorageSetter.Contract.SetBytes32(&_StorageSetter.TransactOpts, _slot, _value)
}

// SetBytes32 is a paid mutator transaction binding the contract method 0x4e91db08.
//
// Solidity: function setBytes32(bytes32 _slot, bytes32 _value) returns()
func (_StorageSetter *StorageSetterTransactorSession) SetBytes32(_slot [32]byte, _value [32]byte) (*types.Transaction, error) {
	return _StorageSetter.Contract.SetBytes32(&_StorageSetter.TransactOpts, _slot, _value)
}

// SetUint is a paid mutator transaction binding the contract method 0xe2a4853a.
//
// Solidity: function setUint(bytes32 _slot, uint256 _value) returns()
func (_StorageSetter *StorageSetterTransactor) SetUint(opts *bind.TransactOpts, _slot [32]byte, _value *big.Int) (*types.Transaction, error) {
	return _StorageSetter.contract.Transact(opts, "setUint", _slot, _value)
}

// SetUint is a paid mutator transaction binding the contract method 0xe2a4853a.
//
// Solidity: function setUint(bytes32 _slot, uint256 _value) returns()
func (_StorageSetter *StorageSetterSession) SetUint(_slot [32]byte, _value *big.Int) (*types.Transaction, error) {
	return _StorageSetter.Contract.SetUint(&_StorageSetter.TransactOpts, _slot, _value)
}

// SetUint is a paid mutator transaction binding the contract method 0xe2a4853a.
//
// Solidity: function setUint(bytes32 _slot, uint256 _value) returns()
func (_StorageSetter *StorageSetterTransactorSession) SetUint(_slot [32]byte, _value *big.Int) (*types.Transaction, error) {
	return _StorageSetter.Contract.SetUint(&_StorageSetter.TransactOpts, _slot, _value)
}
