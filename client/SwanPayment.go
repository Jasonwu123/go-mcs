// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package client

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

// IPaymentMinimallockPaymentParam is an auto generated low-level Go binding around an user-defined struct.
type IPaymentMinimallockPaymentParam struct {
	Id         string
	MinPayment *big.Int
	Amount     *big.Int
	LockTime   *big.Int
	Recipient  common.Address
	Size       *big.Int
	CopyLimit  uint8
}

// ClientMetaData contains all meta data concerning the Client contract.
var ClientMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"minPayment\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockTime\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"copyLimit\",\"type\":\"uint8\"}],\"internalType\":\"structIPaymentMinimal.lockPaymentParam\",\"name\":\"param\",\"type\":\"tuple\"}],\"name\":\"lockTokenPayment\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ClientABI is the input ABI used to generate the binding from.
// Deprecated: Use ClientMetaData.ABI instead.
var ClientABI = ClientMetaData.ABI

// Client is an auto generated Go binding around an Ethereum contract.
type Client struct {
	ClientCaller     // Read-only binding to the contract
	ClientTransactor // Write-only binding to the contract
	ClientFilterer   // Log filterer for contract events
}

// ClientCaller is an auto generated read-only Go binding around an Ethereum contract.
type ClientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ClientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClientFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ClientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ClientSession struct {
	Contract     *Client           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ClientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ClientCallerSession struct {
	Contract *ClientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ClientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ClientTransactorSession struct {
	Contract     *ClientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ClientRaw is an auto generated low-level Go binding around an Ethereum contract.
type ClientRaw struct {
	Contract *Client // Generic contract binding to access the raw methods on
}

// ClientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ClientCallerRaw struct {
	Contract *ClientCaller // Generic read-only contract binding to access the raw methods on
}

// ClientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ClientTransactorRaw struct {
	Contract *ClientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewClient creates a new instance of Client, bound to a specific deployed contract.
func NewClient(address common.Address, backend bind.ContractBackend) (*Client, error) {
	contract, err := bindClient(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Client{ClientCaller: ClientCaller{contract: contract}, ClientTransactor: ClientTransactor{contract: contract}, ClientFilterer: ClientFilterer{contract: contract}}, nil
}

// NewClientCaller creates a new read-only instance of Client, bound to a specific deployed contract.
func NewClientCaller(address common.Address, caller bind.ContractCaller) (*ClientCaller, error) {
	contract, err := bindClient(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ClientCaller{contract: contract}, nil
}

// NewClientTransactor creates a new write-only instance of Client, bound to a specific deployed contract.
func NewClientTransactor(address common.Address, transactor bind.ContractTransactor) (*ClientTransactor, error) {
	contract, err := bindClient(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ClientTransactor{contract: contract}, nil
}

// NewClientFilterer creates a new log filterer instance of Client, bound to a specific deployed contract.
func NewClientFilterer(address common.Address, filterer bind.ContractFilterer) (*ClientFilterer, error) {
	contract, err := bindClient(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ClientFilterer{contract: contract}, nil
}

// bindClient binds a generic wrapper to an already deployed contract.
func bindClient(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ClientABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Client *ClientRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Client.Contract.ClientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Client *ClientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Client.Contract.ClientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Client *ClientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Client.Contract.ClientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Client *ClientCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Client.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Client *ClientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Client.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Client *ClientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Client.Contract.contract.Transact(opts, method, params...)
}

// LockTokenPayment is a paid mutator transaction binding the contract method 0xf4d98717.
//
// Solidity: function lockTokenPayment((string,uint256,uint256,uint256,address,uint256,uint8) param) returns(bool)
func (_Client *ClientTransactor) LockTokenPayment(opts *bind.TransactOpts, param IPaymentMinimallockPaymentParam) (*types.Transaction, error) {
	return _Client.contract.Transact(opts, "lockTokenPayment", param)
}

// LockTokenPayment is a paid mutator transaction binding the contract method 0xf4d98717.
//
// Solidity: function lockTokenPayment((string,uint256,uint256,uint256,address,uint256,uint8) param) returns(bool)
func (_Client *ClientSession) LockTokenPayment(param IPaymentMinimallockPaymentParam) (*types.Transaction, error) {
	return _Client.Contract.LockTokenPayment(&_Client.TransactOpts, param)
}

// LockTokenPayment is a paid mutator transaction binding the contract method 0xf4d98717.
//
// Solidity: function lockTokenPayment((string,uint256,uint256,uint256,address,uint256,uint8) param) returns(bool)
func (_Client *ClientTransactorSession) LockTokenPayment(param IPaymentMinimallockPaymentParam) (*types.Transaction, error) {
	return _Client.Contract.LockTokenPayment(&_Client.TransactOpts, param)
}
