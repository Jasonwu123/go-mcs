// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package swanpayment

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

// IPaymentMinimallockPaymentParam is an auto generated low-level Go binding around an user-defined structs.
type IPaymentMinimallockPaymentParam struct {
	Id         string
	MinPayment *big.Int
	Amount     *big.Int
	LockTime   *big.Int
	Recipient  common.Address
	Size       *big.Int
	CopyLimit  uint8
}

// SwanpaymentMetaData contains all meta data concerning the Swanpayment contract.
var SwanpaymentMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"minPayment\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockTime\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"copyLimit\",\"type\":\"uint8\"}],\"internalType\":\"structIPaymentMinimal.lockPaymentParam\",\"name\":\"param\",\"type\":\"tuple\"}],\"name\":\"lockTokenPayment\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// SwanpaymentABI is the input ABI used to generate the binding from.
// Deprecated: Use SwanpaymentMetaData.ABI instead.
var SwanpaymentABI = SwanpaymentMetaData.ABI

// Swanpayment is an auto generated Go binding around an Ethereum contract.
type Swanpayment struct {
	SwanpaymentCaller     // Read-only binding to the contract
	SwanpaymentTransactor // Write-only binding to the contract
	SwanpaymentFilterer   // Log filterer for contract events
}

// SwanpaymentCaller is an auto generated read-only Go binding around an Ethereum contract.
type SwanpaymentCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwanpaymentTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SwanpaymentTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwanpaymentFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SwanpaymentFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwanpaymentSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SwanpaymentSession struct {
	Contract     *Swanpayment      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwanpaymentCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SwanpaymentCallerSession struct {
	Contract *SwanpaymentCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// SwanpaymentTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SwanpaymentTransactorSession struct {
	Contract     *SwanpaymentTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SwanpaymentRaw is an auto generated low-level Go binding around an Ethereum contract.
type SwanpaymentRaw struct {
	Contract *Swanpayment // Generic contract binding to access the raw methods on
}

// SwanpaymentCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SwanpaymentCallerRaw struct {
	Contract *SwanpaymentCaller // Generic read-only contract binding to access the raw methods on
}

// SwanpaymentTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SwanpaymentTransactorRaw struct {
	Contract *SwanpaymentTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSwanpayment creates a new instance of Swanpayment, bound to a specific deployed contract.
func NewSwanpayment(address common.Address, backend bind.ContractBackend) (*Swanpayment, error) {
	contract, err := bindSwanpayment(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Swanpayment{SwanpaymentCaller: SwanpaymentCaller{contract: contract}, SwanpaymentTransactor: SwanpaymentTransactor{contract: contract}, SwanpaymentFilterer: SwanpaymentFilterer{contract: contract}}, nil
}

// NewSwanpaymentCaller creates a new read-only instance of Swanpayment, bound to a specific deployed contract.
func NewSwanpaymentCaller(address common.Address, caller bind.ContractCaller) (*SwanpaymentCaller, error) {
	contract, err := bindSwanpayment(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SwanpaymentCaller{contract: contract}, nil
}

// NewSwanpaymentTransactor creates a new write-only instance of Swanpayment, bound to a specific deployed contract.
func NewSwanpaymentTransactor(address common.Address, transactor bind.ContractTransactor) (*SwanpaymentTransactor, error) {
	contract, err := bindSwanpayment(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SwanpaymentTransactor{contract: contract}, nil
}

// NewSwanpaymentFilterer creates a new log filterer instance of Swanpayment, bound to a specific deployed contract.
func NewSwanpaymentFilterer(address common.Address, filterer bind.ContractFilterer) (*SwanpaymentFilterer, error) {
	contract, err := bindSwanpayment(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SwanpaymentFilterer{contract: contract}, nil
}

// bindSwanpayment binds a generic wrapper to an already deployed contract.
func bindSwanpayment(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SwanpaymentABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a structs for named
// returns.
func (_Swanpayment *SwanpaymentRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Swanpayment.Contract.SwanpaymentCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Swanpayment *SwanpaymentRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swanpayment.Contract.SwanpaymentTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Swanpayment *SwanpaymentRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Swanpayment.Contract.SwanpaymentTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a structs for named
// returns.
func (_Swanpayment *SwanpaymentCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Swanpayment.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Swanpayment *SwanpaymentTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swanpayment.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Swanpayment *SwanpaymentTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Swanpayment.Contract.contract.Transact(opts, method, params...)
}

// LockTokenPayment is a paid mutator transaction binding the contract method 0xf4d98717.
//
// Solidity: function lockTokenPayment((string,uint256,uint256,uint256,address,uint256,uint8) param) returns(bool)
func (_Swanpayment *SwanpaymentTransactor) LockTokenPayment(opts *bind.TransactOpts, param IPaymentMinimallockPaymentParam) (*types.Transaction, error) {
	return _Swanpayment.contract.Transact(opts, "lockTokenPayment", param)
}

// LockTokenPayment is a paid mutator transaction binding the contract method 0xf4d98717.
//
// Solidity: function lockTokenPayment((string,uint256,uint256,uint256,address,uint256,uint8) param) returns(bool)
func (_Swanpayment *SwanpaymentSession) LockTokenPayment(param IPaymentMinimallockPaymentParam) (*types.Transaction, error) {
	return _Swanpayment.Contract.LockTokenPayment(&_Swanpayment.TransactOpts, param)
}

// LockTokenPayment is a paid mutator transaction binding the contract method 0xf4d98717.
//
// Solidity: function lockTokenPayment((string,uint256,uint256,uint256,address,uint256,uint8) param) returns(bool)
func (_Swanpayment *SwanpaymentTransactorSession) LockTokenPayment(param IPaymentMinimallockPaymentParam) (*types.Transaction, error) {
	return _Swanpayment.Contract.LockTokenPayment(&_Swanpayment.TransactOpts, param)
}
