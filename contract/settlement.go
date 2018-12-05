// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// MathABI is the input ABI used to generate the binding from.
const MathABI = "[]"

// MathBin is the compiled bytecode used for deploying new contracts.
const MathBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a7230582073952228720e406e0ffac900950da50cb694b781d373e1e7c4b5ce648481fad10029`

// DeployMath deploys a new Ethereum contract, binding an instance of Math to it.
func DeployMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Math, error) {
	parsed, err := abi.JSON(strings.NewReader(MathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Math{MathCaller: MathCaller{contract: contract}, MathTransactor: MathTransactor{contract: contract}, MathFilterer: MathFilterer{contract: contract}}, nil
}

// Math is an auto generated Go binding around an Ethereum contract.
type Math struct {
	MathCaller     // Read-only binding to the contract
	MathTransactor // Write-only binding to the contract
	MathFilterer   // Log filterer for contract events
}

// MathCaller is an auto generated read-only Go binding around an Ethereum contract.
type MathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MathSession struct {
	Contract     *Math             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MathCallerSession struct {
	Contract *MathCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MathTransactorSession struct {
	Contract     *MathTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MathRaw is an auto generated low-level Go binding around an Ethereum contract.
type MathRaw struct {
	Contract *Math // Generic contract binding to access the raw methods on
}

// MathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MathCallerRaw struct {
	Contract *MathCaller // Generic read-only contract binding to access the raw methods on
}

// MathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MathTransactorRaw struct {
	Contract *MathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMath creates a new instance of Math, bound to a specific deployed contract.
func NewMath(address common.Address, backend bind.ContractBackend) (*Math, error) {
	contract, err := bindMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Math{MathCaller: MathCaller{contract: contract}, MathTransactor: MathTransactor{contract: contract}, MathFilterer: MathFilterer{contract: contract}}, nil
}

// NewMathCaller creates a new read-only instance of Math, bound to a specific deployed contract.
func NewMathCaller(address common.Address, caller bind.ContractCaller) (*MathCaller, error) {
	contract, err := bindMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MathCaller{contract: contract}, nil
}

// NewMathTransactor creates a new write-only instance of Math, bound to a specific deployed contract.
func NewMathTransactor(address common.Address, transactor bind.ContractTransactor) (*MathTransactor, error) {
	contract, err := bindMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MathTransactor{contract: contract}, nil
}

// NewMathFilterer creates a new log filterer instance of Math, bound to a specific deployed contract.
func NewMathFilterer(address common.Address, filterer bind.ContractFilterer) (*MathFilterer, error) {
	contract, err := bindMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MathFilterer{contract: contract}, nil
}

// bindMath binds a generic wrapper to an already deployed contract.
func bindMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Math *MathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Math.Contract.MathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Math *MathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Math.Contract.MathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Math *MathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Math.Contract.MathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Math *MathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Math.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Math *MathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Math.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Math *MathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Math.Contract.contract.Transact(opts, method, params...)
}

// OrderbookABI is the input ABI used to generate the binding from.
const OrderbookABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_orderID\",\"type\":\"bytes32\"}],\"name\":\"orderConfirmer\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orderID\",\"type\":\"bytes32\"}],\"name\":\"orderState\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orderID\",\"type\":\"bytes32\"}],\"name\":\"orderMatch\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orderID\",\"type\":\"bytes32\"}],\"name\":\"orderTrader\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// OrderbookBin is the compiled bytecode used for deploying new contracts.
const OrderbookBin = `0x`

// DeployOrderbook deploys a new Ethereum contract, binding an instance of Orderbook to it.
func DeployOrderbook(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Orderbook, error) {
	parsed, err := abi.JSON(strings.NewReader(OrderbookABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OrderbookBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Orderbook{OrderbookCaller: OrderbookCaller{contract: contract}, OrderbookTransactor: OrderbookTransactor{contract: contract}, OrderbookFilterer: OrderbookFilterer{contract: contract}}, nil
}

// Orderbook is an auto generated Go binding around an Ethereum contract.
type Orderbook struct {
	OrderbookCaller     // Read-only binding to the contract
	OrderbookTransactor // Write-only binding to the contract
	OrderbookFilterer   // Log filterer for contract events
}

// OrderbookCaller is an auto generated read-only Go binding around an Ethereum contract.
type OrderbookCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderbookTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OrderbookTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderbookFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OrderbookFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderbookSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OrderbookSession struct {
	Contract     *Orderbook        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OrderbookCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OrderbookCallerSession struct {
	Contract *OrderbookCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// OrderbookTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OrderbookTransactorSession struct {
	Contract     *OrderbookTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// OrderbookRaw is an auto generated low-level Go binding around an Ethereum contract.
type OrderbookRaw struct {
	Contract *Orderbook // Generic contract binding to access the raw methods on
}

// OrderbookCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OrderbookCallerRaw struct {
	Contract *OrderbookCaller // Generic read-only contract binding to access the raw methods on
}

// OrderbookTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OrderbookTransactorRaw struct {
	Contract *OrderbookTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOrderbook creates a new instance of Orderbook, bound to a specific deployed contract.
func NewOrderbook(address common.Address, backend bind.ContractBackend) (*Orderbook, error) {
	contract, err := bindOrderbook(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Orderbook{OrderbookCaller: OrderbookCaller{contract: contract}, OrderbookTransactor: OrderbookTransactor{contract: contract}, OrderbookFilterer: OrderbookFilterer{contract: contract}}, nil
}

// NewOrderbookCaller creates a new read-only instance of Orderbook, bound to a specific deployed contract.
func NewOrderbookCaller(address common.Address, caller bind.ContractCaller) (*OrderbookCaller, error) {
	contract, err := bindOrderbook(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OrderbookCaller{contract: contract}, nil
}

// NewOrderbookTransactor creates a new write-only instance of Orderbook, bound to a specific deployed contract.
func NewOrderbookTransactor(address common.Address, transactor bind.ContractTransactor) (*OrderbookTransactor, error) {
	contract, err := bindOrderbook(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OrderbookTransactor{contract: contract}, nil
}

// NewOrderbookFilterer creates a new log filterer instance of Orderbook, bound to a specific deployed contract.
func NewOrderbookFilterer(address common.Address, filterer bind.ContractFilterer) (*OrderbookFilterer, error) {
	contract, err := bindOrderbook(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OrderbookFilterer{contract: contract}, nil
}

// bindOrderbook binds a generic wrapper to an already deployed contract.
func bindOrderbook(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OrderbookABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Orderbook *OrderbookRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Orderbook.Contract.OrderbookCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Orderbook *OrderbookRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Orderbook.Contract.OrderbookTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Orderbook *OrderbookRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Orderbook.Contract.OrderbookTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Orderbook *OrderbookCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Orderbook.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Orderbook *OrderbookTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Orderbook.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Orderbook *OrderbookTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Orderbook.Contract.contract.Transact(opts, method, params...)
}

// OrderConfirmer is a free data retrieval call binding the contract method 0x1107c3f7.
//
// Solidity: function orderConfirmer(_orderID bytes32) constant returns(address)
func (_Orderbook *OrderbookCaller) OrderConfirmer(opts *bind.CallOpts, _orderID [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Orderbook.contract.Call(opts, out, "orderConfirmer", _orderID)
	return *ret0, err
}

// OrderConfirmer is a free data retrieval call binding the contract method 0x1107c3f7.
//
// Solidity: function orderConfirmer(_orderID bytes32) constant returns(address)
func (_Orderbook *OrderbookSession) OrderConfirmer(_orderID [32]byte) (common.Address, error) {
	return _Orderbook.Contract.OrderConfirmer(&_Orderbook.CallOpts, _orderID)
}

// OrderConfirmer is a free data retrieval call binding the contract method 0x1107c3f7.
//
// Solidity: function orderConfirmer(_orderID bytes32) constant returns(address)
func (_Orderbook *OrderbookCallerSession) OrderConfirmer(_orderID [32]byte) (common.Address, error) {
	return _Orderbook.Contract.OrderConfirmer(&_Orderbook.CallOpts, _orderID)
}

// OrderMatch is a free data retrieval call binding the contract method 0xaf3e8a40.
//
// Solidity: function orderMatch(_orderID bytes32) constant returns(bytes32)
func (_Orderbook *OrderbookCaller) OrderMatch(opts *bind.CallOpts, _orderID [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Orderbook.contract.Call(opts, out, "orderMatch", _orderID)
	return *ret0, err
}

// OrderMatch is a free data retrieval call binding the contract method 0xaf3e8a40.
//
// Solidity: function orderMatch(_orderID bytes32) constant returns(bytes32)
func (_Orderbook *OrderbookSession) OrderMatch(_orderID [32]byte) ([32]byte, error) {
	return _Orderbook.Contract.OrderMatch(&_Orderbook.CallOpts, _orderID)
}

// OrderMatch is a free data retrieval call binding the contract method 0xaf3e8a40.
//
// Solidity: function orderMatch(_orderID bytes32) constant returns(bytes32)
func (_Orderbook *OrderbookCallerSession) OrderMatch(_orderID [32]byte) ([32]byte, error) {
	return _Orderbook.Contract.OrderMatch(&_Orderbook.CallOpts, _orderID)
}

// OrderState is a free data retrieval call binding the contract method 0xaab14d04.
//
// Solidity: function orderState(_orderID bytes32) constant returns(uint8)
func (_Orderbook *OrderbookCaller) OrderState(opts *bind.CallOpts, _orderID [32]byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Orderbook.contract.Call(opts, out, "orderState", _orderID)
	return *ret0, err
}

// OrderState is a free data retrieval call binding the contract method 0xaab14d04.
//
// Solidity: function orderState(_orderID bytes32) constant returns(uint8)
func (_Orderbook *OrderbookSession) OrderState(_orderID [32]byte) (uint8, error) {
	return _Orderbook.Contract.OrderState(&_Orderbook.CallOpts, _orderID)
}

// OrderState is a free data retrieval call binding the contract method 0xaab14d04.
//
// Solidity: function orderState(_orderID bytes32) constant returns(uint8)
func (_Orderbook *OrderbookCallerSession) OrderState(_orderID [32]byte) (uint8, error) {
	return _Orderbook.Contract.OrderState(&_Orderbook.CallOpts, _orderID)
}

// OrderTrader is a free data retrieval call binding the contract method 0xb1a08010.
//
// Solidity: function orderTrader(_orderID bytes32) constant returns(address)
func (_Orderbook *OrderbookCaller) OrderTrader(opts *bind.CallOpts, _orderID [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Orderbook.contract.Call(opts, out, "orderTrader", _orderID)
	return *ret0, err
}

// OrderTrader is a free data retrieval call binding the contract method 0xb1a08010.
//
// Solidity: function orderTrader(_orderID bytes32) constant returns(address)
func (_Orderbook *OrderbookSession) OrderTrader(_orderID [32]byte) (common.Address, error) {
	return _Orderbook.Contract.OrderTrader(&_Orderbook.CallOpts, _orderID)
}

// OrderTrader is a free data retrieval call binding the contract method 0xb1a08010.
//
// Solidity: function orderTrader(_orderID bytes32) constant returns(address)
func (_Orderbook *OrderbookCallerSession) OrderTrader(_orderID [32]byte) (common.Address, error) {
	return _Orderbook.Contract.OrderTrader(&_Orderbook.CallOpts, _orderID)
}

// OwnableABI is the input ABI used to generate the binding from.
const OwnableABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// OwnableBin is the compiled bytecode used for deploying new contracts.
const OwnableBin = `0x608060405234801561001057600080fd5b5060008054600160a060020a0319163317905561020b806100326000396000f3006080604052600436106100565763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663715018a6811461005b5780638da5cb5b14610072578063f2fde38b146100a3575b600080fd5b34801561006757600080fd5b506100706100c4565b005b34801561007e57600080fd5b50610087610130565b60408051600160a060020a039092168252519081900360200190f35b3480156100af57600080fd5b50610070600160a060020a036004351661013f565b600054600160a060020a031633146100db57600080fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b600054600160a060020a031681565b600054600160a060020a0316331461015657600080fd5b61015f81610162565b50565b600160a060020a038116151561017757600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a72305820a40a21644ad073aa04af1df855b9b2b9e568d32fb0991dc4449a7653d061712b0029`

// DeployOwnable deploys a new Ethereum contract, binding an instance of Ownable to it.
func DeployOwnable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Ownable, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OwnableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ownable.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Ownable *OwnableSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, _newOwner)
}

// OwnableOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the Ownable contract.
type OwnableOwnershipRenouncedIterator struct {
	Event *OwnableOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *OwnableOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipRenounced)
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
		it.Event = new(OwnableOwnershipRenounced)
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
func (it *OwnableOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipRenounced represents a OwnershipRenounced event raised by the Ownable contract.
type OwnableOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(previousOwner indexed address)
func (_Ownable *OwnableFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*OwnableOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipRenouncedIterator{contract: _Ownable.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(previousOwner indexed address)
func (_Ownable *OwnableFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipRenounced)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// OwnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ownable contract.
type OwnableOwnershipTransferredIterator struct {
	Event *OwnableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OwnableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipTransferred)
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
		it.Event = new(OwnableOwnershipTransferred)
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
func (it *OwnableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable contract.
type OwnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Ownable *OwnableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipTransferredIterator{contract: _Ownable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Ownable *OwnableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipTransferred)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// RenExBalancesABI is the input ABI used to generate the binding from.
const RenExBalancesABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_traderFrom\",\"type\":\"address\"},{\"name\":\"_traderTo\",\"type\":\"address\"},{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_fee\",\"type\":\"uint256\"},{\"name\":\"_feePayee\",\"type\":\"address\"}],\"name\":\"transferBalanceWithFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"settlementContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// RenExBalancesBin is the compiled bytecode used for deploying new contracts.
const RenExBalancesBin = `0x`

// DeployRenExBalances deploys a new Ethereum contract, binding an instance of RenExBalances to it.
func DeployRenExBalances(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RenExBalances, error) {
	parsed, err := abi.JSON(strings.NewReader(RenExBalancesABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RenExBalancesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RenExBalances{RenExBalancesCaller: RenExBalancesCaller{contract: contract}, RenExBalancesTransactor: RenExBalancesTransactor{contract: contract}, RenExBalancesFilterer: RenExBalancesFilterer{contract: contract}}, nil
}

// RenExBalances is an auto generated Go binding around an Ethereum contract.
type RenExBalances struct {
	RenExBalancesCaller     // Read-only binding to the contract
	RenExBalancesTransactor // Write-only binding to the contract
	RenExBalancesFilterer   // Log filterer for contract events
}

// RenExBalancesCaller is an auto generated read-only Go binding around an Ethereum contract.
type RenExBalancesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RenExBalancesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RenExBalancesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RenExBalancesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RenExBalancesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RenExBalancesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RenExBalancesSession struct {
	Contract     *RenExBalances    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RenExBalancesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RenExBalancesCallerSession struct {
	Contract *RenExBalancesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// RenExBalancesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RenExBalancesTransactorSession struct {
	Contract     *RenExBalancesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// RenExBalancesRaw is an auto generated low-level Go binding around an Ethereum contract.
type RenExBalancesRaw struct {
	Contract *RenExBalances // Generic contract binding to access the raw methods on
}

// RenExBalancesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RenExBalancesCallerRaw struct {
	Contract *RenExBalancesCaller // Generic read-only contract binding to access the raw methods on
}

// RenExBalancesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RenExBalancesTransactorRaw struct {
	Contract *RenExBalancesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRenExBalances creates a new instance of RenExBalances, bound to a specific deployed contract.
func NewRenExBalances(address common.Address, backend bind.ContractBackend) (*RenExBalances, error) {
	contract, err := bindRenExBalances(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RenExBalances{RenExBalancesCaller: RenExBalancesCaller{contract: contract}, RenExBalancesTransactor: RenExBalancesTransactor{contract: contract}, RenExBalancesFilterer: RenExBalancesFilterer{contract: contract}}, nil
}

// NewRenExBalancesCaller creates a new read-only instance of RenExBalances, bound to a specific deployed contract.
func NewRenExBalancesCaller(address common.Address, caller bind.ContractCaller) (*RenExBalancesCaller, error) {
	contract, err := bindRenExBalances(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RenExBalancesCaller{contract: contract}, nil
}

// NewRenExBalancesTransactor creates a new write-only instance of RenExBalances, bound to a specific deployed contract.
func NewRenExBalancesTransactor(address common.Address, transactor bind.ContractTransactor) (*RenExBalancesTransactor, error) {
	contract, err := bindRenExBalances(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RenExBalancesTransactor{contract: contract}, nil
}

// NewRenExBalancesFilterer creates a new log filterer instance of RenExBalances, bound to a specific deployed contract.
func NewRenExBalancesFilterer(address common.Address, filterer bind.ContractFilterer) (*RenExBalancesFilterer, error) {
	contract, err := bindRenExBalances(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RenExBalancesFilterer{contract: contract}, nil
}

// bindRenExBalances binds a generic wrapper to an already deployed contract.
func bindRenExBalances(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RenExBalancesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RenExBalances *RenExBalancesRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RenExBalances.Contract.RenExBalancesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RenExBalances *RenExBalancesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RenExBalances.Contract.RenExBalancesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RenExBalances *RenExBalancesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RenExBalances.Contract.RenExBalancesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RenExBalances *RenExBalancesCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RenExBalances.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RenExBalances *RenExBalancesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RenExBalances.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RenExBalances *RenExBalancesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RenExBalances.Contract.contract.Transact(opts, method, params...)
}

// SettlementContract is a free data retrieval call binding the contract method 0xea42418b.
//
// Solidity: function settlementContract() constant returns(address)
func (_RenExBalances *RenExBalancesCaller) SettlementContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RenExBalances.contract.Call(opts, out, "settlementContract")
	return *ret0, err
}

// SettlementContract is a free data retrieval call binding the contract method 0xea42418b.
//
// Solidity: function settlementContract() constant returns(address)
func (_RenExBalances *RenExBalancesSession) SettlementContract() (common.Address, error) {
	return _RenExBalances.Contract.SettlementContract(&_RenExBalances.CallOpts)
}

// SettlementContract is a free data retrieval call binding the contract method 0xea42418b.
//
// Solidity: function settlementContract() constant returns(address)
func (_RenExBalances *RenExBalancesCallerSession) SettlementContract() (common.Address, error) {
	return _RenExBalances.Contract.SettlementContract(&_RenExBalances.CallOpts)
}

// TransferBalanceWithFee is a paid mutator transaction binding the contract method 0x34814e58.
//
// Solidity: function transferBalanceWithFee(_traderFrom address, _traderTo address, _token address, _value uint256, _fee uint256, _feePayee address) returns()
func (_RenExBalances *RenExBalancesTransactor) TransferBalanceWithFee(opts *bind.TransactOpts, _traderFrom common.Address, _traderTo common.Address, _token common.Address, _value *big.Int, _fee *big.Int, _feePayee common.Address) (*types.Transaction, error) {
	return _RenExBalances.contract.Transact(opts, "transferBalanceWithFee", _traderFrom, _traderTo, _token, _value, _fee, _feePayee)
}

// TransferBalanceWithFee is a paid mutator transaction binding the contract method 0x34814e58.
//
// Solidity: function transferBalanceWithFee(_traderFrom address, _traderTo address, _token address, _value uint256, _fee uint256, _feePayee address) returns()
func (_RenExBalances *RenExBalancesSession) TransferBalanceWithFee(_traderFrom common.Address, _traderTo common.Address, _token common.Address, _value *big.Int, _fee *big.Int, _feePayee common.Address) (*types.Transaction, error) {
	return _RenExBalances.Contract.TransferBalanceWithFee(&_RenExBalances.TransactOpts, _traderFrom, _traderTo, _token, _value, _fee, _feePayee)
}

// TransferBalanceWithFee is a paid mutator transaction binding the contract method 0x34814e58.
//
// Solidity: function transferBalanceWithFee(_traderFrom address, _traderTo address, _token address, _value uint256, _fee uint256, _feePayee address) returns()
func (_RenExBalances *RenExBalancesTransactorSession) TransferBalanceWithFee(_traderFrom common.Address, _traderTo common.Address, _token common.Address, _value *big.Int, _fee *big.Int, _feePayee common.Address) (*types.Transaction, error) {
	return _RenExBalances.Contract.TransferBalanceWithFee(&_RenExBalances.TransactOpts, _traderFrom, _traderTo, _token, _value, _fee, _feePayee)
}

// RenExSettlementABI is the input ABI used to generate the binding from.
const RenExSettlementABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"orderStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newRenExTokensContract\",\"type\":\"address\"}],\"name\":\"updateRenExTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"renExTokensContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"submissionGasPriceLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orderID\",\"type\":\"bytes32\"}],\"name\":\"getMatchDetails\",\"outputs\":[{\"name\":\"settled\",\"type\":\"bool\"},{\"name\":\"orderIsBuy\",\"type\":\"bool\"},{\"name\":\"matchedID\",\"type\":\"bytes32\"},{\"name\":\"priorityVolume\",\"type\":\"uint256\"},{\"name\":\"secondaryVolume\",\"type\":\"uint256\"},{\"name\":\"priorityFee\",\"type\":\"uint256\"},{\"name\":\"secondaryFee\",\"type\":\"uint256\"},{\"name\":\"priorityToken\",\"type\":\"uint32\"},{\"name\":\"secondaryToken\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOrderbookContract\",\"type\":\"address\"}],\"name\":\"updateOrderbook\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newSubmissionGasPriceLimit\",\"type\":\"uint256\"}],\"name\":\"updateSubmissionGasPriceLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"DARKNODE_FEES_DENOMINATOR\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"orderSubmitter\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"RENEX_ATOMIC_SETTLEMENT_ID\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"orderDetails\",\"outputs\":[{\"name\":\"settlementID\",\"type\":\"uint64\"},{\"name\":\"tokens\",\"type\":\"uint64\"},{\"name\":\"price\",\"type\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\"},{\"name\":\"minimumVolume\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"matchTimestamp\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"DARKNODE_FEES_NUMERATOR\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newSlasherAddress\",\"type\":\"address\"}],\"name\":\"updateSlasher\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_prefix\",\"type\":\"bytes\"},{\"name\":\"_settlementID\",\"type\":\"uint64\"},{\"name\":\"_tokens\",\"type\":\"uint64\"},{\"name\":\"_price\",\"type\":\"uint256\"},{\"name\":\"_volume\",\"type\":\"uint256\"},{\"name\":\"_minimumVolume\",\"type\":\"uint256\"}],\"name\":\"submitOrder\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"orderbookContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_prefix\",\"type\":\"bytes\"},{\"name\":\"_settlementID\",\"type\":\"uint64\"},{\"name\":\"_tokens\",\"type\":\"uint64\"},{\"name\":\"_price\",\"type\":\"uint256\"},{\"name\":\"_volume\",\"type\":\"uint256\"},{\"name\":\"_minimumVolume\",\"type\":\"uint256\"}],\"name\":\"hashOrder\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"RENEX_SETTLEMENT_ID\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_buyID\",\"type\":\"bytes32\"},{\"name\":\"_sellID\",\"type\":\"bytes32\"}],\"name\":\"settle\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"slasherAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"renExBalancesContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newRenExBalancesContract\",\"type\":\"address\"}],\"name\":\"updateRenExBalances\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_guiltyOrderID\",\"type\":\"bytes32\"}],\"name\":\"slash\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_VERSION\",\"type\":\"string\"},{\"name\":\"_orderbookContract\",\"type\":\"address\"},{\"name\":\"_renExTokensContract\",\"type\":\"address\"},{\"name\":\"_renExBalancesContract\",\"type\":\"address\"},{\"name\":\"_slasherAddress\",\"type\":\"address\"},{\"name\":\"_submissionGasPriceLimit\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"previousOrderbook\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"nextOrderbook\",\"type\":\"address\"}],\"name\":\"LogOrderbookUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"previousRenExTokens\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"nextRenExTokens\",\"type\":\"address\"}],\"name\":\"LogRenExTokensUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"previousRenExBalances\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"nextRenExBalances\",\"type\":\"address\"}],\"name\":\"LogRenExBalancesUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"previousSubmissionGasPriceLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"nextSubmissionGasPriceLimit\",\"type\":\"uint256\"}],\"name\":\"LogSubmissionGasPriceLimitUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"previousSlasher\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"nextSlasher\",\"type\":\"address\"}],\"name\":\"LogSlasherUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"orderID\",\"type\":\"bytes32\"}],\"name\":\"LogOrderSettled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// RenExSettlementBin is the compiled bytecode used for deploying new contracts.
const RenExSettlementBin = `0x60806040523480156200001157600080fd5b5060405162002a4038038062002a408339810160409081528151602080840151928401516060850151608086015160a087015160008054600160a060020a031916331790559490960180519096929491936200007391600191890190620000cc565b5060028054600160a060020a0319908116600160a060020a039788161790915560038054821695871695909517909455600480548516938616939093179092556005805490931693169290921790556006555062000171565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200010f57805160ff19168380011785556200013f565b828001600101855582156200013f579182015b828111156200013f57825182559160200191906001019062000122565b506200014d92915062000151565b5090565b6200016e91905b808211156200014d576000815560010162000158565b90565b6128bf80620001816000396000f3006080604052600436106101455763ffffffff60e060020a6000350416632dff692d811461014a57806334106c89146101865780634015e83b146101a95780634c3052de146101da5780634f54f4d8146102015780636074b8061461026b578063675df16f1461028c5780636af630d1146102a4578063715018a6146102b957806375d4115e146102ce5780638da5cb5b146102e6578063943a5e28146102fb578063a3c50b3214610329578063ab0fd37314610379578063aba8803714610394578063b3139d38146103a9578063b86f602c146103ca578063bbe7221e14610408578063ce0f92b71461041d578063d23df2d11461045b578063d32a9cd914610470578063d53c61bf1461048b578063ddf25ce9146104a0578063ee0715ed146104b5578063f2fde38b146104d6578063f415ed14146104f7578063ffa1ad741461050f575b600080fd5b34801561015657600080fd5b50610162600435610599565b6040518082600381111561017257fe5b60ff16815260200191505060405180910390f35b34801561019257600080fd5b506101a7600160a060020a03600435166105ae565b005b3480156101b557600080fd5b506101be61062f565b60408051600160a060020a039092168252519081900360200190f35b3480156101e657600080fd5b506101ef61063e565b60408051918252519081900360200190f35b34801561020d57600080fd5b50610219600435610644565b604080519915158a5297151560208a0152888801969096526060880194909452608087019290925260a086015260c085015263ffffffff90811660e08501521661010083015251908190036101200190f35b34801561027757600080fd5b506101a7600160a060020a0360043516610844565b34801561029857600080fd5b506101a76004356108c5565b3480156102b057600080fd5b506101ef61091e565b3480156102c557600080fd5b506101a7610924565b3480156102da57600080fd5b506101be600435610983565b3480156102f257600080fd5b506101be61099e565b34801561030757600080fd5b506103106109ad565b6040805163ffffffff9092168252519081900360200190f35b34801561033557600080fd5b506103416004356109b2565b6040805167ffffffffffffffff9687168152949095166020850152838501929092526060830152608082015290519081900360a00190f35b34801561038557600080fd5b506101ef6004356024356109f2565b3480156103a057600080fd5b506101ef6109ad565b3480156103b557600080fd5b506101a7600160a060020a0360043516610a0f565b3480156103d657600080fd5b506101a7602460048035828101929101359067ffffffffffffffff90358116906044351660643560843560a435610a90565b34801561041457600080fd5b506101be610d97565b34801561042957600080fd5b506101ef602460048035828101929101359067ffffffffffffffff90358116906044351660643560843560a435610da6565b34801561046757600080fd5b50610310610e2b565b34801561047c57600080fd5b506101a7600435602435610e30565b34801561049757600080fd5b506101be611551565b3480156104ac57600080fd5b506101be611560565b3480156104c157600080fd5b506101a7600160a060020a036004351661156f565b3480156104e257600080fd5b506101a7600160a060020a03600435166115f0565b34801561050357600080fd5b506101a7600435611613565b34801561051b57600080fd5b50610524611ba1565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561055e578181015183820152602001610546565b50505050905090810190601f16801561058b5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b60096020526000908152604090205460ff1681565b600054600160a060020a031633146105c557600080fd5b60035460408051600160a060020a039283168152918316602083015280517fc44a7f49dd4281e6c3ed47edb754b69b064653d53ed217e1354e79e8fe4b06a09281900390910190a160038054600160a060020a031916600160a060020a0392909216919091179055565b600354600160a060020a031681565b60065481565b600080600080600080600080600080600061065d612789565b600260009054906101000a9004600160a060020a0316600160a060020a031663af3e8a408e6040518263ffffffff1660e060020a028152600401808260001916600019168152602001915050602060405180830381600087803b1580156106c357600080fd5b505af11580156106d7573d6000803e3d6000fd5b505050506040513d60208110156106ed57600080fd5b505199506106fa8d611c2e565b9a508a61070857898d61070b565b8c8a5b600082815260076020526040902054919450925061074b90849084906107469068010000000000000000900467ffffffffffffffff16611c6c565b611e36565b9050600260008e81526009602052604090205460ff16600381111561076c57fe5b14806107945750600360008e81526009602052604090205460ff16600381111561079257fe5b145b8b8b83600001518460200151856040015186606001516020600760008c6000191660001916815260200190815260200160002060000160089054906101000a900467ffffffffffffffff1667ffffffffffffffff169060020a9004600760008c6000191660001916815260200190815260200160002060000160089054906101000a900467ffffffffffffffff169b509b509b509b509b509b509b509b509b505050509193959799909294969850565b600054600160a060020a0316331461085b57600080fd5b60025460408051600160a060020a039283168152918316602083015280517ff7af59918b82b5e13957d357d0fcc86f12a806b0d2e826bc24a0f13ae85e45989281900390910190a160028054600160a060020a031916600160a060020a0392909216919091179055565b600054600160a060020a031633146108dc57600080fd5b600654604080519182526020820183905280517fd0ef246766073915a6813492cff2a021d7cd4bf7d4feff3ef74327c7f4940e969281900390910190a1600655565b6103e881565b600054600160a060020a0316331461093b57600080fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a260008054600160a060020a0319169055565b600860205260009081526040902054600160a060020a031681565b600054600160a060020a031681565b600281565b600760205260009081526040902080546001820154600283015460039093015467ffffffffffffffff808416946801000000000000000090940416929085565b600a60209081526000928352604080842090915290825290205481565b600054600160a060020a03163314610a2657600080fd5b60055460408051600160a060020a039283168152918316602083015280517f933228a1c3ba8fadd3ce47a9db5b898be647f89af99ba7c1b9a655f59ea306c89281900390910190a160058054600160a060020a031916600160a060020a0392909216919091179055565b610a986127d2565b6000600654803a11151515610af7576040805160e560020a62461bcd02815260206004820152601260248201527f67617320707269636520746f6f20686967680000000000000000000000000000604482015290519081900360640190fd5b60a0604051908101604052808967ffffffffffffffff1681526020018867ffffffffffffffff168152602001878152602001868152602001858152509250610b708a8a8080601f01602080910402602001604051908101604052809392919081815260200183838082843750899450611fbd9350505050565b91506000808381526009602052604090205460ff166003811115610b9057fe5b14610be5576040805160e560020a62461bcd02815260206004820152601760248201527f6f7264657220616c7265616479207375626d6974746564000000000000000000604482015290519081900360640190fd5b60028054604080517faab14d04000000000000000000000000000000000000000000000000000000008152600481018690529051600160a060020a039092169163aab14d04916024808201926020929091908290030181600087803b158015610c4d57600080fd5b505af1158015610c61573d6000803e3d6000fd5b505050506040513d6020811015610c7757600080fd5b50516003811115610c8457fe5b14610cd9576040805160e560020a62461bcd02815260206004820152601160248201527f756e636f6e6669726d6564206f72646572000000000000000000000000000000604482015290519081900360640190fd5b60008281526008602090815260408083208054600160a060020a031916331790556009909152902080546001919060ff19168280021790555050600090815260076020908152604091829020835181549285015167ffffffffffffffff1990931667ffffffffffffffff918216176fffffffffffffffff000000000000000019166801000000000000000091909316029190911781559082015160018201556060820151600282015560809091015160039091015550505050505050565b600254600160a060020a031681565b6000610e1f88888080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505060a0604051908101604052808967ffffffffffffffff1681526020018867ffffffffffffffff16815260200187815260200186815260200185815250611fbd565b98975050505050505050565b600181565b610e38612816565b600080600160008681526009602052604090205460ff166003811115610e5a57fe5b14610eaf576040805160e560020a62461bcd02815260206004820152601260248201527f696e76616c696420627579207374617475730000000000000000000000000000604482015290519081900360640190fd5b600160008581526009602052604090205460ff166003811115610ece57fe5b14610f23576040805160e560020a62461bcd02815260206004820152601360248201527f696e76616c69642073656c6c2073746174757300000000000000000000000000604482015290519081900360640190fd5b60008581526007602052604090205467ffffffffffffffff1660021480610f62575060008581526007602052604090205467ffffffffffffffff166001145b1515610fb8576040805160e560020a62461bcd02815260206004820152601560248201527f696e76616c696420736574746c656d656e742069640000000000000000000000604482015290519081900360640190fd5b61111260076000876000191660001916815260200190815260200160002060a060405190810160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152602001600182015481526020016002820154815260200160038201548152505060076000876000191660001916815260200190815260200160002060a060405190810160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820154815260200160028201548152602001600382015481525050612125565b1515611168576040805160e560020a62461bcd02815260206004820152601360248201527f696e636f6d70617469626c65206f726465727300000000000000000000000000604482015290519081900360640190fd5b600254604080517faf3e8a400000000000000000000000000000000000000000000000000000000081526004810188905290518692600160a060020a03169163af3e8a409160248083019260209291908290030181600087803b1580156111ce57600080fd5b505af11580156111e2573d6000803e3d6000fd5b505050506040513d60208110156111f857600080fd5b50511461124f576040805160e560020a62461bcd02815260206004820152601260248201527f756e636f6e6669726d6564206f72646572730000000000000000000000000000604482015290519081900360640190fd5b60008581526007602052604090205461127d9068010000000000000000900467ffffffffffffffff16611c6c565b80516040015190935015156112dc576040805160e560020a62461bcd02815260206004820152601b60248201527f756e72656769737465726564207072696f7269747920746f6b656e0000000000604482015290519081900360640190fd5b826020015160400151151561133b576040805160e560020a62461bcd02815260206004820152601c60248201527f756e72656769737465726564207365636f6e6461727920746f6b656e00000000604482015290519081900360640190fd5b600254604080516000805160206128748339815191528152600481018890529051600160a060020a039092169163b1a08010916024808201926020929091908290030181600087803b15801561139057600080fd5b505af11580156113a4573d6000803e3d6000fd5b505050506040513d60208110156113ba57600080fd5b5051600254604080516000805160206128748339815191528152600481018890529051929450600160a060020a039091169163b1a08010916024808201926020929091908290030181600087803b15801561141457600080fd5b505af1158015611428573d6000803e3d6000fd5b505050506040513d602081101561143e57600080fd5b50519050600160a060020a0382811690821614156114a6576040805160e560020a62461bcd02815260206004820152601760248201527f6f72646572732066726f6d2073616d6520747261646572000000000000000000604482015290519081900360640190fd5b6114b385858484876121b7565b6000858152600a60209081526040808320878452825280832042905587835260099091528082208054600260ff199182168117909255878452828420805490911690911790555186917f8e4e7b583cd791eb2be6c2d2f7db850a3684c285cd18afe11e47f383581d419891a260405184907f8e4e7b583cd791eb2be6c2d2f7db850a3684c285cd18afe11e47f383581d419890600090a25050505050565b600554600160a060020a031681565b600454600160a060020a031681565b600054600160a060020a0316331461158657600080fd5b60045460408051600160a060020a039283168152918316602083015280517f28e85eee30dd92456f8c6864bcdaadb36644672cee6f285d571b1e58c08adca19281900390910190a160048054600160a060020a031916600160a060020a0392909216919091179055565b600054600160a060020a0316331461160757600080fd5b6116108161236f565b50565b6000806000611620612816565b611628612789565b600554600160a060020a0316331461168a576040805160e560020a62461bcd02815260206004820152600c60248201527f756e617574686f72697a65640000000000000000000000000000000000000000604482015290519081900360640190fd5b60008681526007602052604090205467ffffffffffffffff166002146116fa576040805160e560020a62461bcd02815260206004820152601960248201527f736c617368696e67206e6f6e2d61746f6d696320747261646500000000000000604482015290519081900360640190fd5b600254604080517faf3e8a40000000000000000000000000000000000000000000000000000000008152600481018990529051600160a060020a039092169163af3e8a40916024808201926020929091908290030181600087803b15801561176157600080fd5b505af1158015611775573d6000803e3d6000fd5b505050506040513d602081101561178b57600080fd5b50519450600260008781526009602052604090205460ff1660038111156117ae57fe5b14611803576040805160e560020a62461bcd02815260206004820152601460248201527f696e76616c6964206f7264657220737461747573000000000000000000000000604482015290519081900360640190fd5b600260008681526009602052604090205460ff16600381111561182257fe5b14611877576040805160e560020a62461bcd02815260206004820152601460248201527f696e76616c6964206f7264657220737461747573000000000000000000000000604482015290519081900360640190fd5b6000868152600960205260409020805460ff1916600317905561189986611c2e565b6118a45784866118a7565b85855b60008281526007602052604090205491955093506118da9068010000000000000000900467ffffffffffffffff16611c6c565b91506118e78484846123df565b600480546002546040805160008051602061287483398151915281529384018b905251939450600160a060020a03918216936334814e5893929091169163b1a080109160248083019260209291908290030181600087803b15801561194b57600080fd5b505af115801561195f573d6000803e3d6000fd5b505050506040513d602081101561197557600080fd5b5051600254604080516000805160206128748339815191528152600481018b90529051600160a060020a039092169163b1a08010916024808201926020929091908290030181600087803b1580156119cc57600080fd5b505af11580156119e0573d6000803e3d6000fd5b505050506040513d60208110156119f657600080fd5b50516080850151604080870151815160e060020a63ffffffff8816028152600160a060020a0395861660048201529385166024850152939091166044830152606482019290925260006084820181905260a48201819052915160c4808301939282900301818387803b158015611a6b57600080fd5b505af1158015611a7f573d6000803e3d6000fd5b5050600480546002546040805160008051602061287483398151915281529384018c905251600160a060020a0392831695506334814e58945091169163b1a080109160248083019260209291908290030181600087803b158015611ae257600080fd5b505af1158015611af6573d6000803e3d6000fd5b505050506040513d6020811015611b0c57600080fd5b50516005546080850151604080870151815160e060020a63ffffffff8816028152600160a060020a0395861660048201529385166024850152919093166044830152606482015260006084820181905260a48201819052915160c4808301939282900301818387803b158015611b8157600080fd5b505af1158015611b95573d6000803e3d6000fd5b50505050505050505050565b60018054604080516020600284861615610100026000190190941693909304601f81018490048402820184019092528181529291830182828015611c265780601f10611bfb57610100808354040283529160200191611c26565b820191906000526020600020905b815481529060010190602001808311611c0957829003601f168201915b505050505081565b60009081526007602052604090205468010000000000000000900463ffffffff81811664010000000067ffffffffffffffff90931692909204161090565b611c74612816565b600354604080517ffbb6272d00000000000000000000000000000000000000000000000000000000815263ffffffff64010000000067ffffffffffffffff871604166004820152905160009283928392839283928392600160a060020a039092169163fbb6272d9160248082019260609290919082900301818787803b158015611cfd57600080fd5b505af1158015611d11573d6000803e3d6000fd5b505050506040513d6060811015611d2757600080fd5b508051602082015160409283015160035484517ffbb6272d00000000000000000000000000000000000000000000000000000000815263ffffffff8e1660048201529451939a509198509650600160a060020a03169163fbb6272d9160248083019260609291908290030181600087803b158015611da457600080fd5b505af1158015611db8573d6000803e3d6000fd5b505050506040513d6060811015611dce57600080fd5b508051602080830151604093840151845160a081018652600160a060020a039b8c1681870190815260ff9b8c166060808401919091529a1515608083015281528551998a0186529a909316885297909716868801521515908501525050509082015292915050565b611e3e612789565b611e4661283c565b6000806000611e5361283c565b611e5b61283c565b6040805190810160405280600760008c6000191660001916815260200190815260200160002060010154600760008e600019166000191681526020019081526020016000206001015401815260200160028152509550611ef7600760008c6000191660001916815260200190815260200160002060020154600760008c6000191660001916815260200190815260200160002060020154612638565b8651909550611f2c90611f1190879063ffffffff61265016565b8760200151600c808c600001516020015160ff160303612679565b9350611f48856001600c8b602001516020015160ff1603612679565b9250611f53846126de565b9150611f5e836126de565b6040805160c08101825284518152825160208083019190915280860151928201929092528183015160608201528a5151600160a060020a039081166080830152918b01515190911660a0820152975090505b5050505050509392505050565b600082826000015183602001518460400151856060015186608001516040516020018087805190602001908083835b6020831061200b5780518252601f199092019160209182019101611fec565b6001836020036101000a0380198251168184511680821785525050505050509050018667ffffffffffffffff1667ffffffffffffffff1678010000000000000000000000000000000000000000000000000281526008018567ffffffffffffffff1667ffffffffffffffff16780100000000000000000000000000000000000000000000000002815260080184815260200183815260200182815260200196505050505050506040516020818303038152906040526040518082805190602001908083835b602083106120ef5780518252601f1990920191602091820191016120d0565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902090505b92915050565b600061213983602001518360200151612708565b15156121475750600061211f565b81604001518360400151101561215f5750600061211f565b8160800151836060015110156121775750600061211f565b82608001518260600151101561218f5750600061211f565b8151835167ffffffffffffffff9081169116146121ae5750600061211f565b50600192915050565b6121bf612789565b60008681526007602052604090205467ffffffffffffffff166002146121f2576121ea868684611e36565b905080612201565b6121fd8686846123df565b9050805b600480546080830151835160408086015160008d815260086020528281205483517f34814e58000000000000000000000000000000000000000000000000000000008152600160a060020a038e8116998201999099528c8916602482015295881660448701526064860194909452608485019190915291851660a48401525194955092909116926334814e589260c48084019391929182900301818387803b1580156122ac57600080fd5b505af11580156122c0573d6000803e3d6000fd5b50506004805460a0850151602080870151606088015160008d8152600890935260408084205481517f34814e58000000000000000000000000000000000000000000000000000000008152600160a060020a038e8116998201999099528e8916602482015295881660448701526064860193909352608485019190915290851660a4840152519390921694506334814e58935060c4808201939182900301818387803b158015611b8157600080fd5b600160a060020a038116151561238457600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a360008054600160a060020a031916600160a060020a0392909216919091179055565b6123e7612789565b6123ef61283c565b6000806123fa61283c565b600061240461283c565b6040805190810160405280600760008c6000191660001916815260200190815260200160002060010154600760008e6000191660001916815260200190815260200160002060010154018152602001600281525095506124a0600760008c6000191660001916815260200190815260200160002060020154600760008c6000191660001916815260200190815260200160002060020154612638565b94506124b388602001516000015161277b565b15612540576124d2856001600c8b602001516020015160ff1603612679565b93506124dd846126de565b925060c06040519081016040528060008152602001600081526020018460200151815260200184602001518152602001896020015160000151600160a060020a03168152602001896020015160000151600160a060020a03168152509650611fb0565b87515161254c9061277b565b156125c257855161256890611f1190879063ffffffff61265016565b9150612573826126de565b6040805160c0810182526000808252602080830191909152830180519282019290925290516060820152895151600160a060020a0390811660808301528a51511660a082015297509050611fb0565b6040805160e560020a62461bcd02815260206004820152602660248201527f6e6f6e2d6574682061746f6d696320737761707320617265206e6f742073757060448201527f706f727465640000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b60008183106126475781612649565b825b9392505050565b60008215156126615750600061211f565b5081810281838281151561267157fe5b041461211f57fe5b6000808260010b1215156126bf57604d600183900b131561269657fe5b826126ae85600185900b600a0a63ffffffff61265016565b8115156126b757fe5b049050612649565b8160000360010b600a0a83858115156126d457fe5b048115156126b757fe5b6126e661283c565b50604080518082019091526103e86103e6830204808252909103602082015290565b600064010000000067ffffffffffffffff83160463ffffffff90811690841614801561274e575064010000000067ffffffffffffffff84160463ffffffff908116908316145b80156126495750505063ffffffff81811664010000000067ffffffffffffffff9093169290920416111590565b600160a060020a0316151590565b60c060405190810160405280600081526020016000815260200160008152602001600081526020016000600160a060020a031681526020016000600160a060020a031681525090565b60a060405190810160405280600067ffffffffffffffff168152602001600067ffffffffffffffff1681526020016000815260200160008152602001600081525090565b60c06040519081016040528061282a612853565b8152602001612837612853565b905290565b604080518082019091526000808252602082015290565b6040805160608101825260008082526020820181905291810191909152905600b1a0801000000000000000000000000000000000000000000000000000000000a165627a7a72305820d3d7325b9f4e5929d4d7c57b1066ff5287c15b904864ad870da565a137a725780029`

// DeployRenExSettlement deploys a new Ethereum contract, binding an instance of RenExSettlement to it.
func DeployRenExSettlement(auth *bind.TransactOpts, backend bind.ContractBackend, _VERSION string, _orderbookContract common.Address, _renExTokensContract common.Address, _renExBalancesContract common.Address, _slasherAddress common.Address, _submissionGasPriceLimit *big.Int) (common.Address, *types.Transaction, *RenExSettlement, error) {
	parsed, err := abi.JSON(strings.NewReader(RenExSettlementABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RenExSettlementBin), backend, _VERSION, _orderbookContract, _renExTokensContract, _renExBalancesContract, _slasherAddress, _submissionGasPriceLimit)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RenExSettlement{RenExSettlementCaller: RenExSettlementCaller{contract: contract}, RenExSettlementTransactor: RenExSettlementTransactor{contract: contract}, RenExSettlementFilterer: RenExSettlementFilterer{contract: contract}}, nil
}

// RenExSettlement is an auto generated Go binding around an Ethereum contract.
type RenExSettlement struct {
	RenExSettlementCaller     // Read-only binding to the contract
	RenExSettlementTransactor // Write-only binding to the contract
	RenExSettlementFilterer   // Log filterer for contract events
}

// RenExSettlementCaller is an auto generated read-only Go binding around an Ethereum contract.
type RenExSettlementCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RenExSettlementTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RenExSettlementTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RenExSettlementFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RenExSettlementFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RenExSettlementSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RenExSettlementSession struct {
	Contract     *RenExSettlement  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RenExSettlementCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RenExSettlementCallerSession struct {
	Contract *RenExSettlementCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// RenExSettlementTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RenExSettlementTransactorSession struct {
	Contract     *RenExSettlementTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// RenExSettlementRaw is an auto generated low-level Go binding around an Ethereum contract.
type RenExSettlementRaw struct {
	Contract *RenExSettlement // Generic contract binding to access the raw methods on
}

// RenExSettlementCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RenExSettlementCallerRaw struct {
	Contract *RenExSettlementCaller // Generic read-only contract binding to access the raw methods on
}

// RenExSettlementTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RenExSettlementTransactorRaw struct {
	Contract *RenExSettlementTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRenExSettlement creates a new instance of RenExSettlement, bound to a specific deployed contract.
func NewRenExSettlement(address common.Address, backend bind.ContractBackend) (*RenExSettlement, error) {
	contract, err := bindRenExSettlement(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RenExSettlement{RenExSettlementCaller: RenExSettlementCaller{contract: contract}, RenExSettlementTransactor: RenExSettlementTransactor{contract: contract}, RenExSettlementFilterer: RenExSettlementFilterer{contract: contract}}, nil
}

// NewRenExSettlementCaller creates a new read-only instance of RenExSettlement, bound to a specific deployed contract.
func NewRenExSettlementCaller(address common.Address, caller bind.ContractCaller) (*RenExSettlementCaller, error) {
	contract, err := bindRenExSettlement(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RenExSettlementCaller{contract: contract}, nil
}

// NewRenExSettlementTransactor creates a new write-only instance of RenExSettlement, bound to a specific deployed contract.
func NewRenExSettlementTransactor(address common.Address, transactor bind.ContractTransactor) (*RenExSettlementTransactor, error) {
	contract, err := bindRenExSettlement(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RenExSettlementTransactor{contract: contract}, nil
}

// NewRenExSettlementFilterer creates a new log filterer instance of RenExSettlement, bound to a specific deployed contract.
func NewRenExSettlementFilterer(address common.Address, filterer bind.ContractFilterer) (*RenExSettlementFilterer, error) {
	contract, err := bindRenExSettlement(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RenExSettlementFilterer{contract: contract}, nil
}

// bindRenExSettlement binds a generic wrapper to an already deployed contract.
func bindRenExSettlement(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RenExSettlementABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RenExSettlement *RenExSettlementRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RenExSettlement.Contract.RenExSettlementCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RenExSettlement *RenExSettlementRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RenExSettlement.Contract.RenExSettlementTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RenExSettlement *RenExSettlementRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RenExSettlement.Contract.RenExSettlementTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RenExSettlement *RenExSettlementCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RenExSettlement.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RenExSettlement *RenExSettlementTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RenExSettlement.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RenExSettlement *RenExSettlementTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RenExSettlement.Contract.contract.Transact(opts, method, params...)
}

// DARKNODEFEESDENOMINATOR is a free data retrieval call binding the contract method 0x6af630d1.
//
// Solidity: function DARKNODE_FEES_DENOMINATOR() constant returns(uint256)
func (_RenExSettlement *RenExSettlementCaller) DARKNODEFEESDENOMINATOR(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RenExSettlement.contract.Call(opts, out, "DARKNODE_FEES_DENOMINATOR")
	return *ret0, err
}

// DARKNODEFEESDENOMINATOR is a free data retrieval call binding the contract method 0x6af630d1.
//
// Solidity: function DARKNODE_FEES_DENOMINATOR() constant returns(uint256)
func (_RenExSettlement *RenExSettlementSession) DARKNODEFEESDENOMINATOR() (*big.Int, error) {
	return _RenExSettlement.Contract.DARKNODEFEESDENOMINATOR(&_RenExSettlement.CallOpts)
}

// DARKNODEFEESDENOMINATOR is a free data retrieval call binding the contract method 0x6af630d1.
//
// Solidity: function DARKNODE_FEES_DENOMINATOR() constant returns(uint256)
func (_RenExSettlement *RenExSettlementCallerSession) DARKNODEFEESDENOMINATOR() (*big.Int, error) {
	return _RenExSettlement.Contract.DARKNODEFEESDENOMINATOR(&_RenExSettlement.CallOpts)
}

// DARKNODEFEESNUMERATOR is a free data retrieval call binding the contract method 0xaba88037.
//
// Solidity: function DARKNODE_FEES_NUMERATOR() constant returns(uint256)
func (_RenExSettlement *RenExSettlementCaller) DARKNODEFEESNUMERATOR(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RenExSettlement.contract.Call(opts, out, "DARKNODE_FEES_NUMERATOR")
	return *ret0, err
}

// DARKNODEFEESNUMERATOR is a free data retrieval call binding the contract method 0xaba88037.
//
// Solidity: function DARKNODE_FEES_NUMERATOR() constant returns(uint256)
func (_RenExSettlement *RenExSettlementSession) DARKNODEFEESNUMERATOR() (*big.Int, error) {
	return _RenExSettlement.Contract.DARKNODEFEESNUMERATOR(&_RenExSettlement.CallOpts)
}

// DARKNODEFEESNUMERATOR is a free data retrieval call binding the contract method 0xaba88037.
//
// Solidity: function DARKNODE_FEES_NUMERATOR() constant returns(uint256)
func (_RenExSettlement *RenExSettlementCallerSession) DARKNODEFEESNUMERATOR() (*big.Int, error) {
	return _RenExSettlement.Contract.DARKNODEFEESNUMERATOR(&_RenExSettlement.CallOpts)
}

// RENEXATOMICSETTLEMENTID is a free data retrieval call binding the contract method 0x943a5e28.
//
// Solidity: function RENEX_ATOMIC_SETTLEMENT_ID() constant returns(uint32)
func (_RenExSettlement *RenExSettlementCaller) RENEXATOMICSETTLEMENTID(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _RenExSettlement.contract.Call(opts, out, "RENEX_ATOMIC_SETTLEMENT_ID")
	return *ret0, err
}

// RENEXATOMICSETTLEMENTID is a free data retrieval call binding the contract method 0x943a5e28.
//
// Solidity: function RENEX_ATOMIC_SETTLEMENT_ID() constant returns(uint32)
func (_RenExSettlement *RenExSettlementSession) RENEXATOMICSETTLEMENTID() (uint32, error) {
	return _RenExSettlement.Contract.RENEXATOMICSETTLEMENTID(&_RenExSettlement.CallOpts)
}

// RENEXATOMICSETTLEMENTID is a free data retrieval call binding the contract method 0x943a5e28.
//
// Solidity: function RENEX_ATOMIC_SETTLEMENT_ID() constant returns(uint32)
func (_RenExSettlement *RenExSettlementCallerSession) RENEXATOMICSETTLEMENTID() (uint32, error) {
	return _RenExSettlement.Contract.RENEXATOMICSETTLEMENTID(&_RenExSettlement.CallOpts)
}

// RENEXSETTLEMENTID is a free data retrieval call binding the contract method 0xd23df2d1.
//
// Solidity: function RENEX_SETTLEMENT_ID() constant returns(uint32)
func (_RenExSettlement *RenExSettlementCaller) RENEXSETTLEMENTID(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _RenExSettlement.contract.Call(opts, out, "RENEX_SETTLEMENT_ID")
	return *ret0, err
}

// RENEXSETTLEMENTID is a free data retrieval call binding the contract method 0xd23df2d1.
//
// Solidity: function RENEX_SETTLEMENT_ID() constant returns(uint32)
func (_RenExSettlement *RenExSettlementSession) RENEXSETTLEMENTID() (uint32, error) {
	return _RenExSettlement.Contract.RENEXSETTLEMENTID(&_RenExSettlement.CallOpts)
}

// RENEXSETTLEMENTID is a free data retrieval call binding the contract method 0xd23df2d1.
//
// Solidity: function RENEX_SETTLEMENT_ID() constant returns(uint32)
func (_RenExSettlement *RenExSettlementCallerSession) RENEXSETTLEMENTID() (uint32, error) {
	return _RenExSettlement.Contract.RENEXSETTLEMENTID(&_RenExSettlement.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() constant returns(string)
func (_RenExSettlement *RenExSettlementCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _RenExSettlement.contract.Call(opts, out, "VERSION")
	return *ret0, err
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() constant returns(string)
func (_RenExSettlement *RenExSettlementSession) VERSION() (string, error) {
	return _RenExSettlement.Contract.VERSION(&_RenExSettlement.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() constant returns(string)
func (_RenExSettlement *RenExSettlementCallerSession) VERSION() (string, error) {
	return _RenExSettlement.Contract.VERSION(&_RenExSettlement.CallOpts)
}

// GetMatchDetails is a free data retrieval call binding the contract method 0x4f54f4d8.
//
// Solidity: function getMatchDetails(_orderID bytes32) constant returns(settled bool, orderIsBuy bool, matchedID bytes32, priorityVolume uint256, secondaryVolume uint256, priorityFee uint256, secondaryFee uint256, priorityToken uint32, secondaryToken uint32)
func (_RenExSettlement *RenExSettlementCaller) GetMatchDetails(opts *bind.CallOpts, _orderID [32]byte) (struct {
	Settled         bool
	OrderIsBuy      bool
	MatchedID       [32]byte
	PriorityVolume  *big.Int
	SecondaryVolume *big.Int
	PriorityFee     *big.Int
	SecondaryFee    *big.Int
	PriorityToken   uint32
	SecondaryToken  uint32
}, error) {
	ret := new(struct {
		Settled         bool
		OrderIsBuy      bool
		MatchedID       [32]byte
		PriorityVolume  *big.Int
		SecondaryVolume *big.Int
		PriorityFee     *big.Int
		SecondaryFee    *big.Int
		PriorityToken   uint32
		SecondaryToken  uint32
	})
	out := ret
	err := _RenExSettlement.contract.Call(opts, out, "getMatchDetails", _orderID)
	return *ret, err
}

// GetMatchDetails is a free data retrieval call binding the contract method 0x4f54f4d8.
//
// Solidity: function getMatchDetails(_orderID bytes32) constant returns(settled bool, orderIsBuy bool, matchedID bytes32, priorityVolume uint256, secondaryVolume uint256, priorityFee uint256, secondaryFee uint256, priorityToken uint32, secondaryToken uint32)
func (_RenExSettlement *RenExSettlementSession) GetMatchDetails(_orderID [32]byte) (struct {
	Settled         bool
	OrderIsBuy      bool
	MatchedID       [32]byte
	PriorityVolume  *big.Int
	SecondaryVolume *big.Int
	PriorityFee     *big.Int
	SecondaryFee    *big.Int
	PriorityToken   uint32
	SecondaryToken  uint32
}, error) {
	return _RenExSettlement.Contract.GetMatchDetails(&_RenExSettlement.CallOpts, _orderID)
}

// GetMatchDetails is a free data retrieval call binding the contract method 0x4f54f4d8.
//
// Solidity: function getMatchDetails(_orderID bytes32) constant returns(settled bool, orderIsBuy bool, matchedID bytes32, priorityVolume uint256, secondaryVolume uint256, priorityFee uint256, secondaryFee uint256, priorityToken uint32, secondaryToken uint32)
func (_RenExSettlement *RenExSettlementCallerSession) GetMatchDetails(_orderID [32]byte) (struct {
	Settled         bool
	OrderIsBuy      bool
	MatchedID       [32]byte
	PriorityVolume  *big.Int
	SecondaryVolume *big.Int
	PriorityFee     *big.Int
	SecondaryFee    *big.Int
	PriorityToken   uint32
	SecondaryToken  uint32
}, error) {
	return _RenExSettlement.Contract.GetMatchDetails(&_RenExSettlement.CallOpts, _orderID)
}

// HashOrder is a free data retrieval call binding the contract method 0xce0f92b7.
//
// Solidity: function hashOrder(_prefix bytes, _settlementID uint64, _tokens uint64, _price uint256, _volume uint256, _minimumVolume uint256) constant returns(bytes32)
func (_RenExSettlement *RenExSettlementCaller) HashOrder(opts *bind.CallOpts, _prefix []byte, _settlementID uint64, _tokens uint64, _price *big.Int, _volume *big.Int, _minimumVolume *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _RenExSettlement.contract.Call(opts, out, "hashOrder", _prefix, _settlementID, _tokens, _price, _volume, _minimumVolume)
	return *ret0, err
}

// HashOrder is a free data retrieval call binding the contract method 0xce0f92b7.
//
// Solidity: function hashOrder(_prefix bytes, _settlementID uint64, _tokens uint64, _price uint256, _volume uint256, _minimumVolume uint256) constant returns(bytes32)
func (_RenExSettlement *RenExSettlementSession) HashOrder(_prefix []byte, _settlementID uint64, _tokens uint64, _price *big.Int, _volume *big.Int, _minimumVolume *big.Int) ([32]byte, error) {
	return _RenExSettlement.Contract.HashOrder(&_RenExSettlement.CallOpts, _prefix, _settlementID, _tokens, _price, _volume, _minimumVolume)
}

// HashOrder is a free data retrieval call binding the contract method 0xce0f92b7.
//
// Solidity: function hashOrder(_prefix bytes, _settlementID uint64, _tokens uint64, _price uint256, _volume uint256, _minimumVolume uint256) constant returns(bytes32)
func (_RenExSettlement *RenExSettlementCallerSession) HashOrder(_prefix []byte, _settlementID uint64, _tokens uint64, _price *big.Int, _volume *big.Int, _minimumVolume *big.Int) ([32]byte, error) {
	return _RenExSettlement.Contract.HashOrder(&_RenExSettlement.CallOpts, _prefix, _settlementID, _tokens, _price, _volume, _minimumVolume)
}

// MatchTimestamp is a free data retrieval call binding the contract method 0xab0fd373.
//
// Solidity: function matchTimestamp( bytes32,  bytes32) constant returns(uint256)
func (_RenExSettlement *RenExSettlementCaller) MatchTimestamp(opts *bind.CallOpts, arg0 [32]byte, arg1 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RenExSettlement.contract.Call(opts, out, "matchTimestamp", arg0, arg1)
	return *ret0, err
}

// MatchTimestamp is a free data retrieval call binding the contract method 0xab0fd373.
//
// Solidity: function matchTimestamp( bytes32,  bytes32) constant returns(uint256)
func (_RenExSettlement *RenExSettlementSession) MatchTimestamp(arg0 [32]byte, arg1 [32]byte) (*big.Int, error) {
	return _RenExSettlement.Contract.MatchTimestamp(&_RenExSettlement.CallOpts, arg0, arg1)
}

// MatchTimestamp is a free data retrieval call binding the contract method 0xab0fd373.
//
// Solidity: function matchTimestamp( bytes32,  bytes32) constant returns(uint256)
func (_RenExSettlement *RenExSettlementCallerSession) MatchTimestamp(arg0 [32]byte, arg1 [32]byte) (*big.Int, error) {
	return _RenExSettlement.Contract.MatchTimestamp(&_RenExSettlement.CallOpts, arg0, arg1)
}

// OrderDetails is a free data retrieval call binding the contract method 0xa3c50b32.
//
// Solidity: function orderDetails( bytes32) constant returns(settlementID uint64, tokens uint64, price uint256, volume uint256, minimumVolume uint256)
func (_RenExSettlement *RenExSettlementCaller) OrderDetails(opts *bind.CallOpts, arg0 [32]byte) (struct {
	SettlementID  uint64
	Tokens        uint64
	Price         *big.Int
	Volume        *big.Int
	MinimumVolume *big.Int
}, error) {
	ret := new(struct {
		SettlementID  uint64
		Tokens        uint64
		Price         *big.Int
		Volume        *big.Int
		MinimumVolume *big.Int
	})
	out := ret
	err := _RenExSettlement.contract.Call(opts, out, "orderDetails", arg0)
	return *ret, err
}

// OrderDetails is a free data retrieval call binding the contract method 0xa3c50b32.
//
// Solidity: function orderDetails( bytes32) constant returns(settlementID uint64, tokens uint64, price uint256, volume uint256, minimumVolume uint256)
func (_RenExSettlement *RenExSettlementSession) OrderDetails(arg0 [32]byte) (struct {
	SettlementID  uint64
	Tokens        uint64
	Price         *big.Int
	Volume        *big.Int
	MinimumVolume *big.Int
}, error) {
	return _RenExSettlement.Contract.OrderDetails(&_RenExSettlement.CallOpts, arg0)
}

// OrderDetails is a free data retrieval call binding the contract method 0xa3c50b32.
//
// Solidity: function orderDetails( bytes32) constant returns(settlementID uint64, tokens uint64, price uint256, volume uint256, minimumVolume uint256)
func (_RenExSettlement *RenExSettlementCallerSession) OrderDetails(arg0 [32]byte) (struct {
	SettlementID  uint64
	Tokens        uint64
	Price         *big.Int
	Volume        *big.Int
	MinimumVolume *big.Int
}, error) {
	return _RenExSettlement.Contract.OrderDetails(&_RenExSettlement.CallOpts, arg0)
}

// OrderStatus is a free data retrieval call binding the contract method 0x2dff692d.
//
// Solidity: function orderStatus( bytes32) constant returns(uint8)
func (_RenExSettlement *RenExSettlementCaller) OrderStatus(opts *bind.CallOpts, arg0 [32]byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _RenExSettlement.contract.Call(opts, out, "orderStatus", arg0)
	return *ret0, err
}

// OrderStatus is a free data retrieval call binding the contract method 0x2dff692d.
//
// Solidity: function orderStatus( bytes32) constant returns(uint8)
func (_RenExSettlement *RenExSettlementSession) OrderStatus(arg0 [32]byte) (uint8, error) {
	return _RenExSettlement.Contract.OrderStatus(&_RenExSettlement.CallOpts, arg0)
}

// OrderStatus is a free data retrieval call binding the contract method 0x2dff692d.
//
// Solidity: function orderStatus( bytes32) constant returns(uint8)
func (_RenExSettlement *RenExSettlementCallerSession) OrderStatus(arg0 [32]byte) (uint8, error) {
	return _RenExSettlement.Contract.OrderStatus(&_RenExSettlement.CallOpts, arg0)
}

// OrderSubmitter is a free data retrieval call binding the contract method 0x75d4115e.
//
// Solidity: function orderSubmitter( bytes32) constant returns(address)
func (_RenExSettlement *RenExSettlementCaller) OrderSubmitter(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RenExSettlement.contract.Call(opts, out, "orderSubmitter", arg0)
	return *ret0, err
}

// OrderSubmitter is a free data retrieval call binding the contract method 0x75d4115e.
//
// Solidity: function orderSubmitter( bytes32) constant returns(address)
func (_RenExSettlement *RenExSettlementSession) OrderSubmitter(arg0 [32]byte) (common.Address, error) {
	return _RenExSettlement.Contract.OrderSubmitter(&_RenExSettlement.CallOpts, arg0)
}

// OrderSubmitter is a free data retrieval call binding the contract method 0x75d4115e.
//
// Solidity: function orderSubmitter( bytes32) constant returns(address)
func (_RenExSettlement *RenExSettlementCallerSession) OrderSubmitter(arg0 [32]byte) (common.Address, error) {
	return _RenExSettlement.Contract.OrderSubmitter(&_RenExSettlement.CallOpts, arg0)
}

// OrderbookContract is a free data retrieval call binding the contract method 0xbbe7221e.
//
// Solidity: function orderbookContract() constant returns(address)
func (_RenExSettlement *RenExSettlementCaller) OrderbookContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RenExSettlement.contract.Call(opts, out, "orderbookContract")
	return *ret0, err
}

// OrderbookContract is a free data retrieval call binding the contract method 0xbbe7221e.
//
// Solidity: function orderbookContract() constant returns(address)
func (_RenExSettlement *RenExSettlementSession) OrderbookContract() (common.Address, error) {
	return _RenExSettlement.Contract.OrderbookContract(&_RenExSettlement.CallOpts)
}

// OrderbookContract is a free data retrieval call binding the contract method 0xbbe7221e.
//
// Solidity: function orderbookContract() constant returns(address)
func (_RenExSettlement *RenExSettlementCallerSession) OrderbookContract() (common.Address, error) {
	return _RenExSettlement.Contract.OrderbookContract(&_RenExSettlement.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RenExSettlement *RenExSettlementCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RenExSettlement.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RenExSettlement *RenExSettlementSession) Owner() (common.Address, error) {
	return _RenExSettlement.Contract.Owner(&_RenExSettlement.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RenExSettlement *RenExSettlementCallerSession) Owner() (common.Address, error) {
	return _RenExSettlement.Contract.Owner(&_RenExSettlement.CallOpts)
}

// RenExBalancesContract is a free data retrieval call binding the contract method 0xddf25ce9.
//
// Solidity: function renExBalancesContract() constant returns(address)
func (_RenExSettlement *RenExSettlementCaller) RenExBalancesContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RenExSettlement.contract.Call(opts, out, "renExBalancesContract")
	return *ret0, err
}

// RenExBalancesContract is a free data retrieval call binding the contract method 0xddf25ce9.
//
// Solidity: function renExBalancesContract() constant returns(address)
func (_RenExSettlement *RenExSettlementSession) RenExBalancesContract() (common.Address, error) {
	return _RenExSettlement.Contract.RenExBalancesContract(&_RenExSettlement.CallOpts)
}

// RenExBalancesContract is a free data retrieval call binding the contract method 0xddf25ce9.
//
// Solidity: function renExBalancesContract() constant returns(address)
func (_RenExSettlement *RenExSettlementCallerSession) RenExBalancesContract() (common.Address, error) {
	return _RenExSettlement.Contract.RenExBalancesContract(&_RenExSettlement.CallOpts)
}

// RenExTokensContract is a free data retrieval call binding the contract method 0x4015e83b.
//
// Solidity: function renExTokensContract() constant returns(address)
func (_RenExSettlement *RenExSettlementCaller) RenExTokensContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RenExSettlement.contract.Call(opts, out, "renExTokensContract")
	return *ret0, err
}

// RenExTokensContract is a free data retrieval call binding the contract method 0x4015e83b.
//
// Solidity: function renExTokensContract() constant returns(address)
func (_RenExSettlement *RenExSettlementSession) RenExTokensContract() (common.Address, error) {
	return _RenExSettlement.Contract.RenExTokensContract(&_RenExSettlement.CallOpts)
}

// RenExTokensContract is a free data retrieval call binding the contract method 0x4015e83b.
//
// Solidity: function renExTokensContract() constant returns(address)
func (_RenExSettlement *RenExSettlementCallerSession) RenExTokensContract() (common.Address, error) {
	return _RenExSettlement.Contract.RenExTokensContract(&_RenExSettlement.CallOpts)
}

// SlasherAddress is a free data retrieval call binding the contract method 0xd53c61bf.
//
// Solidity: function slasherAddress() constant returns(address)
func (_RenExSettlement *RenExSettlementCaller) SlasherAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RenExSettlement.contract.Call(opts, out, "slasherAddress")
	return *ret0, err
}

// SlasherAddress is a free data retrieval call binding the contract method 0xd53c61bf.
//
// Solidity: function slasherAddress() constant returns(address)
func (_RenExSettlement *RenExSettlementSession) SlasherAddress() (common.Address, error) {
	return _RenExSettlement.Contract.SlasherAddress(&_RenExSettlement.CallOpts)
}

// SlasherAddress is a free data retrieval call binding the contract method 0xd53c61bf.
//
// Solidity: function slasherAddress() constant returns(address)
func (_RenExSettlement *RenExSettlementCallerSession) SlasherAddress() (common.Address, error) {
	return _RenExSettlement.Contract.SlasherAddress(&_RenExSettlement.CallOpts)
}

// SubmissionGasPriceLimit is a free data retrieval call binding the contract method 0x4c3052de.
//
// Solidity: function submissionGasPriceLimit() constant returns(uint256)
func (_RenExSettlement *RenExSettlementCaller) SubmissionGasPriceLimit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RenExSettlement.contract.Call(opts, out, "submissionGasPriceLimit")
	return *ret0, err
}

// SubmissionGasPriceLimit is a free data retrieval call binding the contract method 0x4c3052de.
//
// Solidity: function submissionGasPriceLimit() constant returns(uint256)
func (_RenExSettlement *RenExSettlementSession) SubmissionGasPriceLimit() (*big.Int, error) {
	return _RenExSettlement.Contract.SubmissionGasPriceLimit(&_RenExSettlement.CallOpts)
}

// SubmissionGasPriceLimit is a free data retrieval call binding the contract method 0x4c3052de.
//
// Solidity: function submissionGasPriceLimit() constant returns(uint256)
func (_RenExSettlement *RenExSettlementCallerSession) SubmissionGasPriceLimit() (*big.Int, error) {
	return _RenExSettlement.Contract.SubmissionGasPriceLimit(&_RenExSettlement.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RenExSettlement *RenExSettlementTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RenExSettlement.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RenExSettlement *RenExSettlementSession) RenounceOwnership() (*types.Transaction, error) {
	return _RenExSettlement.Contract.RenounceOwnership(&_RenExSettlement.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RenExSettlement *RenExSettlementTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RenExSettlement.Contract.RenounceOwnership(&_RenExSettlement.TransactOpts)
}

// Settle is a paid mutator transaction binding the contract method 0xd32a9cd9.
//
// Solidity: function settle(_buyID bytes32, _sellID bytes32) returns()
func (_RenExSettlement *RenExSettlementTransactor) Settle(opts *bind.TransactOpts, _buyID [32]byte, _sellID [32]byte) (*types.Transaction, error) {
	return _RenExSettlement.contract.Transact(opts, "settle", _buyID, _sellID)
}

// Settle is a paid mutator transaction binding the contract method 0xd32a9cd9.
//
// Solidity: function settle(_buyID bytes32, _sellID bytes32) returns()
func (_RenExSettlement *RenExSettlementSession) Settle(_buyID [32]byte, _sellID [32]byte) (*types.Transaction, error) {
	return _RenExSettlement.Contract.Settle(&_RenExSettlement.TransactOpts, _buyID, _sellID)
}

// Settle is a paid mutator transaction binding the contract method 0xd32a9cd9.
//
// Solidity: function settle(_buyID bytes32, _sellID bytes32) returns()
func (_RenExSettlement *RenExSettlementTransactorSession) Settle(_buyID [32]byte, _sellID [32]byte) (*types.Transaction, error) {
	return _RenExSettlement.Contract.Settle(&_RenExSettlement.TransactOpts, _buyID, _sellID)
}

// Slash is a paid mutator transaction binding the contract method 0xf415ed14.
//
// Solidity: function slash(_guiltyOrderID bytes32) returns()
func (_RenExSettlement *RenExSettlementTransactor) Slash(opts *bind.TransactOpts, _guiltyOrderID [32]byte) (*types.Transaction, error) {
	return _RenExSettlement.contract.Transact(opts, "slash", _guiltyOrderID)
}

// Slash is a paid mutator transaction binding the contract method 0xf415ed14.
//
// Solidity: function slash(_guiltyOrderID bytes32) returns()
func (_RenExSettlement *RenExSettlementSession) Slash(_guiltyOrderID [32]byte) (*types.Transaction, error) {
	return _RenExSettlement.Contract.Slash(&_RenExSettlement.TransactOpts, _guiltyOrderID)
}

// Slash is a paid mutator transaction binding the contract method 0xf415ed14.
//
// Solidity: function slash(_guiltyOrderID bytes32) returns()
func (_RenExSettlement *RenExSettlementTransactorSession) Slash(_guiltyOrderID [32]byte) (*types.Transaction, error) {
	return _RenExSettlement.Contract.Slash(&_RenExSettlement.TransactOpts, _guiltyOrderID)
}

// SubmitOrder is a paid mutator transaction binding the contract method 0xb86f602c.
//
// Solidity: function submitOrder(_prefix bytes, _settlementID uint64, _tokens uint64, _price uint256, _volume uint256, _minimumVolume uint256) returns()
func (_RenExSettlement *RenExSettlementTransactor) SubmitOrder(opts *bind.TransactOpts, _prefix []byte, _settlementID uint64, _tokens uint64, _price *big.Int, _volume *big.Int, _minimumVolume *big.Int) (*types.Transaction, error) {
	return _RenExSettlement.contract.Transact(opts, "submitOrder", _prefix, _settlementID, _tokens, _price, _volume, _minimumVolume)
}

// SubmitOrder is a paid mutator transaction binding the contract method 0xb86f602c.
//
// Solidity: function submitOrder(_prefix bytes, _settlementID uint64, _tokens uint64, _price uint256, _volume uint256, _minimumVolume uint256) returns()
func (_RenExSettlement *RenExSettlementSession) SubmitOrder(_prefix []byte, _settlementID uint64, _tokens uint64, _price *big.Int, _volume *big.Int, _minimumVolume *big.Int) (*types.Transaction, error) {
	return _RenExSettlement.Contract.SubmitOrder(&_RenExSettlement.TransactOpts, _prefix, _settlementID, _tokens, _price, _volume, _minimumVolume)
}

// SubmitOrder is a paid mutator transaction binding the contract method 0xb86f602c.
//
// Solidity: function submitOrder(_prefix bytes, _settlementID uint64, _tokens uint64, _price uint256, _volume uint256, _minimumVolume uint256) returns()
func (_RenExSettlement *RenExSettlementTransactorSession) SubmitOrder(_prefix []byte, _settlementID uint64, _tokens uint64, _price *big.Int, _volume *big.Int, _minimumVolume *big.Int) (*types.Transaction, error) {
	return _RenExSettlement.Contract.SubmitOrder(&_RenExSettlement.TransactOpts, _prefix, _settlementID, _tokens, _price, _volume, _minimumVolume)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_RenExSettlement *RenExSettlementTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _RenExSettlement.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_RenExSettlement *RenExSettlementSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _RenExSettlement.Contract.TransferOwnership(&_RenExSettlement.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_RenExSettlement *RenExSettlementTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _RenExSettlement.Contract.TransferOwnership(&_RenExSettlement.TransactOpts, _newOwner)
}

// UpdateOrderbook is a paid mutator transaction binding the contract method 0x6074b806.
//
// Solidity: function updateOrderbook(_newOrderbookContract address) returns()
func (_RenExSettlement *RenExSettlementTransactor) UpdateOrderbook(opts *bind.TransactOpts, _newOrderbookContract common.Address) (*types.Transaction, error) {
	return _RenExSettlement.contract.Transact(opts, "updateOrderbook", _newOrderbookContract)
}

// UpdateOrderbook is a paid mutator transaction binding the contract method 0x6074b806.
//
// Solidity: function updateOrderbook(_newOrderbookContract address) returns()
func (_RenExSettlement *RenExSettlementSession) UpdateOrderbook(_newOrderbookContract common.Address) (*types.Transaction, error) {
	return _RenExSettlement.Contract.UpdateOrderbook(&_RenExSettlement.TransactOpts, _newOrderbookContract)
}

// UpdateOrderbook is a paid mutator transaction binding the contract method 0x6074b806.
//
// Solidity: function updateOrderbook(_newOrderbookContract address) returns()
func (_RenExSettlement *RenExSettlementTransactorSession) UpdateOrderbook(_newOrderbookContract common.Address) (*types.Transaction, error) {
	return _RenExSettlement.Contract.UpdateOrderbook(&_RenExSettlement.TransactOpts, _newOrderbookContract)
}

// UpdateRenExBalances is a paid mutator transaction binding the contract method 0xee0715ed.
//
// Solidity: function updateRenExBalances(_newRenExBalancesContract address) returns()
func (_RenExSettlement *RenExSettlementTransactor) UpdateRenExBalances(opts *bind.TransactOpts, _newRenExBalancesContract common.Address) (*types.Transaction, error) {
	return _RenExSettlement.contract.Transact(opts, "updateRenExBalances", _newRenExBalancesContract)
}

// UpdateRenExBalances is a paid mutator transaction binding the contract method 0xee0715ed.
//
// Solidity: function updateRenExBalances(_newRenExBalancesContract address) returns()
func (_RenExSettlement *RenExSettlementSession) UpdateRenExBalances(_newRenExBalancesContract common.Address) (*types.Transaction, error) {
	return _RenExSettlement.Contract.UpdateRenExBalances(&_RenExSettlement.TransactOpts, _newRenExBalancesContract)
}

// UpdateRenExBalances is a paid mutator transaction binding the contract method 0xee0715ed.
//
// Solidity: function updateRenExBalances(_newRenExBalancesContract address) returns()
func (_RenExSettlement *RenExSettlementTransactorSession) UpdateRenExBalances(_newRenExBalancesContract common.Address) (*types.Transaction, error) {
	return _RenExSettlement.Contract.UpdateRenExBalances(&_RenExSettlement.TransactOpts, _newRenExBalancesContract)
}

// UpdateRenExTokens is a paid mutator transaction binding the contract method 0x34106c89.
//
// Solidity: function updateRenExTokens(_newRenExTokensContract address) returns()
func (_RenExSettlement *RenExSettlementTransactor) UpdateRenExTokens(opts *bind.TransactOpts, _newRenExTokensContract common.Address) (*types.Transaction, error) {
	return _RenExSettlement.contract.Transact(opts, "updateRenExTokens", _newRenExTokensContract)
}

// UpdateRenExTokens is a paid mutator transaction binding the contract method 0x34106c89.
//
// Solidity: function updateRenExTokens(_newRenExTokensContract address) returns()
func (_RenExSettlement *RenExSettlementSession) UpdateRenExTokens(_newRenExTokensContract common.Address) (*types.Transaction, error) {
	return _RenExSettlement.Contract.UpdateRenExTokens(&_RenExSettlement.TransactOpts, _newRenExTokensContract)
}

// UpdateRenExTokens is a paid mutator transaction binding the contract method 0x34106c89.
//
// Solidity: function updateRenExTokens(_newRenExTokensContract address) returns()
func (_RenExSettlement *RenExSettlementTransactorSession) UpdateRenExTokens(_newRenExTokensContract common.Address) (*types.Transaction, error) {
	return _RenExSettlement.Contract.UpdateRenExTokens(&_RenExSettlement.TransactOpts, _newRenExTokensContract)
}

// UpdateSlasher is a paid mutator transaction binding the contract method 0xb3139d38.
//
// Solidity: function updateSlasher(_newSlasherAddress address) returns()
func (_RenExSettlement *RenExSettlementTransactor) UpdateSlasher(opts *bind.TransactOpts, _newSlasherAddress common.Address) (*types.Transaction, error) {
	return _RenExSettlement.contract.Transact(opts, "updateSlasher", _newSlasherAddress)
}

// UpdateSlasher is a paid mutator transaction binding the contract method 0xb3139d38.
//
// Solidity: function updateSlasher(_newSlasherAddress address) returns()
func (_RenExSettlement *RenExSettlementSession) UpdateSlasher(_newSlasherAddress common.Address) (*types.Transaction, error) {
	return _RenExSettlement.Contract.UpdateSlasher(&_RenExSettlement.TransactOpts, _newSlasherAddress)
}

// UpdateSlasher is a paid mutator transaction binding the contract method 0xb3139d38.
//
// Solidity: function updateSlasher(_newSlasherAddress address) returns()
func (_RenExSettlement *RenExSettlementTransactorSession) UpdateSlasher(_newSlasherAddress common.Address) (*types.Transaction, error) {
	return _RenExSettlement.Contract.UpdateSlasher(&_RenExSettlement.TransactOpts, _newSlasherAddress)
}

// UpdateSubmissionGasPriceLimit is a paid mutator transaction binding the contract method 0x675df16f.
//
// Solidity: function updateSubmissionGasPriceLimit(_newSubmissionGasPriceLimit uint256) returns()
func (_RenExSettlement *RenExSettlementTransactor) UpdateSubmissionGasPriceLimit(opts *bind.TransactOpts, _newSubmissionGasPriceLimit *big.Int) (*types.Transaction, error) {
	return _RenExSettlement.contract.Transact(opts, "updateSubmissionGasPriceLimit", _newSubmissionGasPriceLimit)
}

// UpdateSubmissionGasPriceLimit is a paid mutator transaction binding the contract method 0x675df16f.
//
// Solidity: function updateSubmissionGasPriceLimit(_newSubmissionGasPriceLimit uint256) returns()
func (_RenExSettlement *RenExSettlementSession) UpdateSubmissionGasPriceLimit(_newSubmissionGasPriceLimit *big.Int) (*types.Transaction, error) {
	return _RenExSettlement.Contract.UpdateSubmissionGasPriceLimit(&_RenExSettlement.TransactOpts, _newSubmissionGasPriceLimit)
}

// UpdateSubmissionGasPriceLimit is a paid mutator transaction binding the contract method 0x675df16f.
//
// Solidity: function updateSubmissionGasPriceLimit(_newSubmissionGasPriceLimit uint256) returns()
func (_RenExSettlement *RenExSettlementTransactorSession) UpdateSubmissionGasPriceLimit(_newSubmissionGasPriceLimit *big.Int) (*types.Transaction, error) {
	return _RenExSettlement.Contract.UpdateSubmissionGasPriceLimit(&_RenExSettlement.TransactOpts, _newSubmissionGasPriceLimit)
}

// RenExSettlementLogOrderSettledIterator is returned from FilterLogOrderSettled and is used to iterate over the raw logs and unpacked data for LogOrderSettled events raised by the RenExSettlement contract.
type RenExSettlementLogOrderSettledIterator struct {
	Event *RenExSettlementLogOrderSettled // Event containing the contract specifics and raw log

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
func (it *RenExSettlementLogOrderSettledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RenExSettlementLogOrderSettled)
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
		it.Event = new(RenExSettlementLogOrderSettled)
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
func (it *RenExSettlementLogOrderSettledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RenExSettlementLogOrderSettledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RenExSettlementLogOrderSettled represents a LogOrderSettled event raised by the RenExSettlement contract.
type RenExSettlementLogOrderSettled struct {
	OrderID [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLogOrderSettled is a free log retrieval operation binding the contract event 0x8e4e7b583cd791eb2be6c2d2f7db850a3684c285cd18afe11e47f383581d4198.
//
// Solidity: event LogOrderSettled(orderID indexed bytes32)
func (_RenExSettlement *RenExSettlementFilterer) FilterLogOrderSettled(opts *bind.FilterOpts, orderID [][32]byte) (*RenExSettlementLogOrderSettledIterator, error) {

	var orderIDRule []interface{}
	for _, orderIDItem := range orderID {
		orderIDRule = append(orderIDRule, orderIDItem)
	}

	logs, sub, err := _RenExSettlement.contract.FilterLogs(opts, "LogOrderSettled", orderIDRule)
	if err != nil {
		return nil, err
	}
	return &RenExSettlementLogOrderSettledIterator{contract: _RenExSettlement.contract, event: "LogOrderSettled", logs: logs, sub: sub}, nil
}

// WatchLogOrderSettled is a free log subscription operation binding the contract event 0x8e4e7b583cd791eb2be6c2d2f7db850a3684c285cd18afe11e47f383581d4198.
//
// Solidity: event LogOrderSettled(orderID indexed bytes32)
func (_RenExSettlement *RenExSettlementFilterer) WatchLogOrderSettled(opts *bind.WatchOpts, sink chan<- *RenExSettlementLogOrderSettled, orderID [][32]byte) (event.Subscription, error) {

	var orderIDRule []interface{}
	for _, orderIDItem := range orderID {
		orderIDRule = append(orderIDRule, orderIDItem)
	}

	logs, sub, err := _RenExSettlement.contract.WatchLogs(opts, "LogOrderSettled", orderIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RenExSettlementLogOrderSettled)
				if err := _RenExSettlement.contract.UnpackLog(event, "LogOrderSettled", log); err != nil {
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

// RenExSettlementLogOrderbookUpdatedIterator is returned from FilterLogOrderbookUpdated and is used to iterate over the raw logs and unpacked data for LogOrderbookUpdated events raised by the RenExSettlement contract.
type RenExSettlementLogOrderbookUpdatedIterator struct {
	Event *RenExSettlementLogOrderbookUpdated // Event containing the contract specifics and raw log

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
func (it *RenExSettlementLogOrderbookUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RenExSettlementLogOrderbookUpdated)
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
		it.Event = new(RenExSettlementLogOrderbookUpdated)
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
func (it *RenExSettlementLogOrderbookUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RenExSettlementLogOrderbookUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RenExSettlementLogOrderbookUpdated represents a LogOrderbookUpdated event raised by the RenExSettlement contract.
type RenExSettlementLogOrderbookUpdated struct {
	PreviousOrderbook common.Address
	NextOrderbook     common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterLogOrderbookUpdated is a free log retrieval operation binding the contract event 0xf7af59918b82b5e13957d357d0fcc86f12a806b0d2e826bc24a0f13ae85e4598.
//
// Solidity: event LogOrderbookUpdated(previousOrderbook address, nextOrderbook address)
func (_RenExSettlement *RenExSettlementFilterer) FilterLogOrderbookUpdated(opts *bind.FilterOpts) (*RenExSettlementLogOrderbookUpdatedIterator, error) {

	logs, sub, err := _RenExSettlement.contract.FilterLogs(opts, "LogOrderbookUpdated")
	if err != nil {
		return nil, err
	}
	return &RenExSettlementLogOrderbookUpdatedIterator{contract: _RenExSettlement.contract, event: "LogOrderbookUpdated", logs: logs, sub: sub}, nil
}

// WatchLogOrderbookUpdated is a free log subscription operation binding the contract event 0xf7af59918b82b5e13957d357d0fcc86f12a806b0d2e826bc24a0f13ae85e4598.
//
// Solidity: event LogOrderbookUpdated(previousOrderbook address, nextOrderbook address)
func (_RenExSettlement *RenExSettlementFilterer) WatchLogOrderbookUpdated(opts *bind.WatchOpts, sink chan<- *RenExSettlementLogOrderbookUpdated) (event.Subscription, error) {

	logs, sub, err := _RenExSettlement.contract.WatchLogs(opts, "LogOrderbookUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RenExSettlementLogOrderbookUpdated)
				if err := _RenExSettlement.contract.UnpackLog(event, "LogOrderbookUpdated", log); err != nil {
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

// RenExSettlementLogRenExBalancesUpdatedIterator is returned from FilterLogRenExBalancesUpdated and is used to iterate over the raw logs and unpacked data for LogRenExBalancesUpdated events raised by the RenExSettlement contract.
type RenExSettlementLogRenExBalancesUpdatedIterator struct {
	Event *RenExSettlementLogRenExBalancesUpdated // Event containing the contract specifics and raw log

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
func (it *RenExSettlementLogRenExBalancesUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RenExSettlementLogRenExBalancesUpdated)
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
		it.Event = new(RenExSettlementLogRenExBalancesUpdated)
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
func (it *RenExSettlementLogRenExBalancesUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RenExSettlementLogRenExBalancesUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RenExSettlementLogRenExBalancesUpdated represents a LogRenExBalancesUpdated event raised by the RenExSettlement contract.
type RenExSettlementLogRenExBalancesUpdated struct {
	PreviousRenExBalances common.Address
	NextRenExBalances     common.Address
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterLogRenExBalancesUpdated is a free log retrieval operation binding the contract event 0x28e85eee30dd92456f8c6864bcdaadb36644672cee6f285d571b1e58c08adca1.
//
// Solidity: event LogRenExBalancesUpdated(previousRenExBalances address, nextRenExBalances address)
func (_RenExSettlement *RenExSettlementFilterer) FilterLogRenExBalancesUpdated(opts *bind.FilterOpts) (*RenExSettlementLogRenExBalancesUpdatedIterator, error) {

	logs, sub, err := _RenExSettlement.contract.FilterLogs(opts, "LogRenExBalancesUpdated")
	if err != nil {
		return nil, err
	}
	return &RenExSettlementLogRenExBalancesUpdatedIterator{contract: _RenExSettlement.contract, event: "LogRenExBalancesUpdated", logs: logs, sub: sub}, nil
}

// WatchLogRenExBalancesUpdated is a free log subscription operation binding the contract event 0x28e85eee30dd92456f8c6864bcdaadb36644672cee6f285d571b1e58c08adca1.
//
// Solidity: event LogRenExBalancesUpdated(previousRenExBalances address, nextRenExBalances address)
func (_RenExSettlement *RenExSettlementFilterer) WatchLogRenExBalancesUpdated(opts *bind.WatchOpts, sink chan<- *RenExSettlementLogRenExBalancesUpdated) (event.Subscription, error) {

	logs, sub, err := _RenExSettlement.contract.WatchLogs(opts, "LogRenExBalancesUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RenExSettlementLogRenExBalancesUpdated)
				if err := _RenExSettlement.contract.UnpackLog(event, "LogRenExBalancesUpdated", log); err != nil {
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

// RenExSettlementLogRenExTokensUpdatedIterator is returned from FilterLogRenExTokensUpdated and is used to iterate over the raw logs and unpacked data for LogRenExTokensUpdated events raised by the RenExSettlement contract.
type RenExSettlementLogRenExTokensUpdatedIterator struct {
	Event *RenExSettlementLogRenExTokensUpdated // Event containing the contract specifics and raw log

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
func (it *RenExSettlementLogRenExTokensUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RenExSettlementLogRenExTokensUpdated)
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
		it.Event = new(RenExSettlementLogRenExTokensUpdated)
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
func (it *RenExSettlementLogRenExTokensUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RenExSettlementLogRenExTokensUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RenExSettlementLogRenExTokensUpdated represents a LogRenExTokensUpdated event raised by the RenExSettlement contract.
type RenExSettlementLogRenExTokensUpdated struct {
	PreviousRenExTokens common.Address
	NextRenExTokens     common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterLogRenExTokensUpdated is a free log retrieval operation binding the contract event 0xc44a7f49dd4281e6c3ed47edb754b69b064653d53ed217e1354e79e8fe4b06a0.
//
// Solidity: event LogRenExTokensUpdated(previousRenExTokens address, nextRenExTokens address)
func (_RenExSettlement *RenExSettlementFilterer) FilterLogRenExTokensUpdated(opts *bind.FilterOpts) (*RenExSettlementLogRenExTokensUpdatedIterator, error) {

	logs, sub, err := _RenExSettlement.contract.FilterLogs(opts, "LogRenExTokensUpdated")
	if err != nil {
		return nil, err
	}
	return &RenExSettlementLogRenExTokensUpdatedIterator{contract: _RenExSettlement.contract, event: "LogRenExTokensUpdated", logs: logs, sub: sub}, nil
}

// WatchLogRenExTokensUpdated is a free log subscription operation binding the contract event 0xc44a7f49dd4281e6c3ed47edb754b69b064653d53ed217e1354e79e8fe4b06a0.
//
// Solidity: event LogRenExTokensUpdated(previousRenExTokens address, nextRenExTokens address)
func (_RenExSettlement *RenExSettlementFilterer) WatchLogRenExTokensUpdated(opts *bind.WatchOpts, sink chan<- *RenExSettlementLogRenExTokensUpdated) (event.Subscription, error) {

	logs, sub, err := _RenExSettlement.contract.WatchLogs(opts, "LogRenExTokensUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RenExSettlementLogRenExTokensUpdated)
				if err := _RenExSettlement.contract.UnpackLog(event, "LogRenExTokensUpdated", log); err != nil {
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

// RenExSettlementLogSlasherUpdatedIterator is returned from FilterLogSlasherUpdated and is used to iterate over the raw logs and unpacked data for LogSlasherUpdated events raised by the RenExSettlement contract.
type RenExSettlementLogSlasherUpdatedIterator struct {
	Event *RenExSettlementLogSlasherUpdated // Event containing the contract specifics and raw log

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
func (it *RenExSettlementLogSlasherUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RenExSettlementLogSlasherUpdated)
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
		it.Event = new(RenExSettlementLogSlasherUpdated)
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
func (it *RenExSettlementLogSlasherUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RenExSettlementLogSlasherUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RenExSettlementLogSlasherUpdated represents a LogSlasherUpdated event raised by the RenExSettlement contract.
type RenExSettlementLogSlasherUpdated struct {
	PreviousSlasher common.Address
	NextSlasher     common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLogSlasherUpdated is a free log retrieval operation binding the contract event 0x933228a1c3ba8fadd3ce47a9db5b898be647f89af99ba7c1b9a655f59ea306c8.
//
// Solidity: event LogSlasherUpdated(previousSlasher address, nextSlasher address)
func (_RenExSettlement *RenExSettlementFilterer) FilterLogSlasherUpdated(opts *bind.FilterOpts) (*RenExSettlementLogSlasherUpdatedIterator, error) {

	logs, sub, err := _RenExSettlement.contract.FilterLogs(opts, "LogSlasherUpdated")
	if err != nil {
		return nil, err
	}
	return &RenExSettlementLogSlasherUpdatedIterator{contract: _RenExSettlement.contract, event: "LogSlasherUpdated", logs: logs, sub: sub}, nil
}

// WatchLogSlasherUpdated is a free log subscription operation binding the contract event 0x933228a1c3ba8fadd3ce47a9db5b898be647f89af99ba7c1b9a655f59ea306c8.
//
// Solidity: event LogSlasherUpdated(previousSlasher address, nextSlasher address)
func (_RenExSettlement *RenExSettlementFilterer) WatchLogSlasherUpdated(opts *bind.WatchOpts, sink chan<- *RenExSettlementLogSlasherUpdated) (event.Subscription, error) {

	logs, sub, err := _RenExSettlement.contract.WatchLogs(opts, "LogSlasherUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RenExSettlementLogSlasherUpdated)
				if err := _RenExSettlement.contract.UnpackLog(event, "LogSlasherUpdated", log); err != nil {
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

// RenExSettlementLogSubmissionGasPriceLimitUpdatedIterator is returned from FilterLogSubmissionGasPriceLimitUpdated and is used to iterate over the raw logs and unpacked data for LogSubmissionGasPriceLimitUpdated events raised by the RenExSettlement contract.
type RenExSettlementLogSubmissionGasPriceLimitUpdatedIterator struct {
	Event *RenExSettlementLogSubmissionGasPriceLimitUpdated // Event containing the contract specifics and raw log

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
func (it *RenExSettlementLogSubmissionGasPriceLimitUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RenExSettlementLogSubmissionGasPriceLimitUpdated)
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
		it.Event = new(RenExSettlementLogSubmissionGasPriceLimitUpdated)
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
func (it *RenExSettlementLogSubmissionGasPriceLimitUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RenExSettlementLogSubmissionGasPriceLimitUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RenExSettlementLogSubmissionGasPriceLimitUpdated represents a LogSubmissionGasPriceLimitUpdated event raised by the RenExSettlement contract.
type RenExSettlementLogSubmissionGasPriceLimitUpdated struct {
	PreviousSubmissionGasPriceLimit *big.Int
	NextSubmissionGasPriceLimit     *big.Int
	Raw                             types.Log // Blockchain specific contextual infos
}

// FilterLogSubmissionGasPriceLimitUpdated is a free log retrieval operation binding the contract event 0xd0ef246766073915a6813492cff2a021d7cd4bf7d4feff3ef74327c7f4940e96.
//
// Solidity: event LogSubmissionGasPriceLimitUpdated(previousSubmissionGasPriceLimit uint256, nextSubmissionGasPriceLimit uint256)
func (_RenExSettlement *RenExSettlementFilterer) FilterLogSubmissionGasPriceLimitUpdated(opts *bind.FilterOpts) (*RenExSettlementLogSubmissionGasPriceLimitUpdatedIterator, error) {

	logs, sub, err := _RenExSettlement.contract.FilterLogs(opts, "LogSubmissionGasPriceLimitUpdated")
	if err != nil {
		return nil, err
	}
	return &RenExSettlementLogSubmissionGasPriceLimitUpdatedIterator{contract: _RenExSettlement.contract, event: "LogSubmissionGasPriceLimitUpdated", logs: logs, sub: sub}, nil
}

// WatchLogSubmissionGasPriceLimitUpdated is a free log subscription operation binding the contract event 0xd0ef246766073915a6813492cff2a021d7cd4bf7d4feff3ef74327c7f4940e96.
//
// Solidity: event LogSubmissionGasPriceLimitUpdated(previousSubmissionGasPriceLimit uint256, nextSubmissionGasPriceLimit uint256)
func (_RenExSettlement *RenExSettlementFilterer) WatchLogSubmissionGasPriceLimitUpdated(opts *bind.WatchOpts, sink chan<- *RenExSettlementLogSubmissionGasPriceLimitUpdated) (event.Subscription, error) {

	logs, sub, err := _RenExSettlement.contract.WatchLogs(opts, "LogSubmissionGasPriceLimitUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RenExSettlementLogSubmissionGasPriceLimitUpdated)
				if err := _RenExSettlement.contract.UnpackLog(event, "LogSubmissionGasPriceLimitUpdated", log); err != nil {
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

// RenExSettlementOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the RenExSettlement contract.
type RenExSettlementOwnershipRenouncedIterator struct {
	Event *RenExSettlementOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *RenExSettlementOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RenExSettlementOwnershipRenounced)
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
		it.Event = new(RenExSettlementOwnershipRenounced)
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
func (it *RenExSettlementOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RenExSettlementOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RenExSettlementOwnershipRenounced represents a OwnershipRenounced event raised by the RenExSettlement contract.
type RenExSettlementOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(previousOwner indexed address)
func (_RenExSettlement *RenExSettlementFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*RenExSettlementOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _RenExSettlement.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RenExSettlementOwnershipRenouncedIterator{contract: _RenExSettlement.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(previousOwner indexed address)
func (_RenExSettlement *RenExSettlementFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *RenExSettlementOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _RenExSettlement.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RenExSettlementOwnershipRenounced)
				if err := _RenExSettlement.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// RenExSettlementOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RenExSettlement contract.
type RenExSettlementOwnershipTransferredIterator struct {
	Event *RenExSettlementOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RenExSettlementOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RenExSettlementOwnershipTransferred)
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
		it.Event = new(RenExSettlementOwnershipTransferred)
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
func (it *RenExSettlementOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RenExSettlementOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RenExSettlementOwnershipTransferred represents a OwnershipTransferred event raised by the RenExSettlement contract.
type RenExSettlementOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_RenExSettlement *RenExSettlementFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RenExSettlementOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RenExSettlement.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RenExSettlementOwnershipTransferredIterator{contract: _RenExSettlement.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_RenExSettlement *RenExSettlementFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RenExSettlementOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RenExSettlement.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RenExSettlementOwnershipTransferred)
				if err := _RenExSettlement.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// RenExTokensABI is the input ABI used to generate the binding from.
const RenExTokensABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_tokenCode\",\"type\":\"uint32\"}],\"name\":\"deregisterToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenCode\",\"type\":\"uint32\"},{\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"name\":\"_tokenDecimals\",\"type\":\"uint8\"}],\"name\":\"registerToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"tokens\",\"outputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"decimals\",\"type\":\"uint8\"},{\"name\":\"registered\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// RenExTokensBin is the compiled bytecode used for deploying new contracts.
const RenExTokensBin = `0x`

// DeployRenExTokens deploys a new Ethereum contract, binding an instance of RenExTokens to it.
func DeployRenExTokens(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RenExTokens, error) {
	parsed, err := abi.JSON(strings.NewReader(RenExTokensABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RenExTokensBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RenExTokens{RenExTokensCaller: RenExTokensCaller{contract: contract}, RenExTokensTransactor: RenExTokensTransactor{contract: contract}, RenExTokensFilterer: RenExTokensFilterer{contract: contract}}, nil
}

// RenExTokens is an auto generated Go binding around an Ethereum contract.
type RenExTokens struct {
	RenExTokensCaller     // Read-only binding to the contract
	RenExTokensTransactor // Write-only binding to the contract
	RenExTokensFilterer   // Log filterer for contract events
}

// RenExTokensCaller is an auto generated read-only Go binding around an Ethereum contract.
type RenExTokensCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RenExTokensTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RenExTokensTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RenExTokensFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RenExTokensFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RenExTokensSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RenExTokensSession struct {
	Contract     *RenExTokens      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RenExTokensCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RenExTokensCallerSession struct {
	Contract *RenExTokensCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// RenExTokensTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RenExTokensTransactorSession struct {
	Contract     *RenExTokensTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// RenExTokensRaw is an auto generated low-level Go binding around an Ethereum contract.
type RenExTokensRaw struct {
	Contract *RenExTokens // Generic contract binding to access the raw methods on
}

// RenExTokensCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RenExTokensCallerRaw struct {
	Contract *RenExTokensCaller // Generic read-only contract binding to access the raw methods on
}

// RenExTokensTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RenExTokensTransactorRaw struct {
	Contract *RenExTokensTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRenExTokens creates a new instance of RenExTokens, bound to a specific deployed contract.
func NewRenExTokens(address common.Address, backend bind.ContractBackend) (*RenExTokens, error) {
	contract, err := bindRenExTokens(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RenExTokens{RenExTokensCaller: RenExTokensCaller{contract: contract}, RenExTokensTransactor: RenExTokensTransactor{contract: contract}, RenExTokensFilterer: RenExTokensFilterer{contract: contract}}, nil
}

// NewRenExTokensCaller creates a new read-only instance of RenExTokens, bound to a specific deployed contract.
func NewRenExTokensCaller(address common.Address, caller bind.ContractCaller) (*RenExTokensCaller, error) {
	contract, err := bindRenExTokens(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RenExTokensCaller{contract: contract}, nil
}

// NewRenExTokensTransactor creates a new write-only instance of RenExTokens, bound to a specific deployed contract.
func NewRenExTokensTransactor(address common.Address, transactor bind.ContractTransactor) (*RenExTokensTransactor, error) {
	contract, err := bindRenExTokens(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RenExTokensTransactor{contract: contract}, nil
}

// NewRenExTokensFilterer creates a new log filterer instance of RenExTokens, bound to a specific deployed contract.
func NewRenExTokensFilterer(address common.Address, filterer bind.ContractFilterer) (*RenExTokensFilterer, error) {
	contract, err := bindRenExTokens(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RenExTokensFilterer{contract: contract}, nil
}

// bindRenExTokens binds a generic wrapper to an already deployed contract.
func bindRenExTokens(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RenExTokensABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RenExTokens *RenExTokensRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RenExTokens.Contract.RenExTokensCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RenExTokens *RenExTokensRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RenExTokens.Contract.RenExTokensTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RenExTokens *RenExTokensRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RenExTokens.Contract.RenExTokensTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RenExTokens *RenExTokensCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RenExTokens.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RenExTokens *RenExTokensTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RenExTokens.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RenExTokens *RenExTokensTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RenExTokens.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RenExTokens *RenExTokensCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RenExTokens.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RenExTokens *RenExTokensSession) Owner() (common.Address, error) {
	return _RenExTokens.Contract.Owner(&_RenExTokens.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RenExTokens *RenExTokensCallerSession) Owner() (common.Address, error) {
	return _RenExTokens.Contract.Owner(&_RenExTokens.CallOpts)
}

// Tokens is a free data retrieval call binding the contract method 0xfbb6272d.
//
// Solidity: function tokens( uint32) constant returns(addr address, decimals uint8, registered bool)
func (_RenExTokens *RenExTokensCaller) Tokens(opts *bind.CallOpts, arg0 uint32) (struct {
	Addr       common.Address
	Decimals   uint8
	Registered bool
}, error) {
	ret := new(struct {
		Addr       common.Address
		Decimals   uint8
		Registered bool
	})
	out := ret
	err := _RenExTokens.contract.Call(opts, out, "tokens", arg0)
	return *ret, err
}

// Tokens is a free data retrieval call binding the contract method 0xfbb6272d.
//
// Solidity: function tokens( uint32) constant returns(addr address, decimals uint8, registered bool)
func (_RenExTokens *RenExTokensSession) Tokens(arg0 uint32) (struct {
	Addr       common.Address
	Decimals   uint8
	Registered bool
}, error) {
	return _RenExTokens.Contract.Tokens(&_RenExTokens.CallOpts, arg0)
}

// Tokens is a free data retrieval call binding the contract method 0xfbb6272d.
//
// Solidity: function tokens( uint32) constant returns(addr address, decimals uint8, registered bool)
func (_RenExTokens *RenExTokensCallerSession) Tokens(arg0 uint32) (struct {
	Addr       common.Address
	Decimals   uint8
	Registered bool
}, error) {
	return _RenExTokens.Contract.Tokens(&_RenExTokens.CallOpts, arg0)
}

// DeregisterToken is a paid mutator transaction binding the contract method 0x25fc575a.
//
// Solidity: function deregisterToken(_tokenCode uint32) returns()
func (_RenExTokens *RenExTokensTransactor) DeregisterToken(opts *bind.TransactOpts, _tokenCode uint32) (*types.Transaction, error) {
	return _RenExTokens.contract.Transact(opts, "deregisterToken", _tokenCode)
}

// DeregisterToken is a paid mutator transaction binding the contract method 0x25fc575a.
//
// Solidity: function deregisterToken(_tokenCode uint32) returns()
func (_RenExTokens *RenExTokensSession) DeregisterToken(_tokenCode uint32) (*types.Transaction, error) {
	return _RenExTokens.Contract.DeregisterToken(&_RenExTokens.TransactOpts, _tokenCode)
}

// DeregisterToken is a paid mutator transaction binding the contract method 0x25fc575a.
//
// Solidity: function deregisterToken(_tokenCode uint32) returns()
func (_RenExTokens *RenExTokensTransactorSession) DeregisterToken(_tokenCode uint32) (*types.Transaction, error) {
	return _RenExTokens.Contract.DeregisterToken(&_RenExTokens.TransactOpts, _tokenCode)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x9e20a9a0.
//
// Solidity: function registerToken(_tokenCode uint32, _tokenAddress address, _tokenDecimals uint8) returns()
func (_RenExTokens *RenExTokensTransactor) RegisterToken(opts *bind.TransactOpts, _tokenCode uint32, _tokenAddress common.Address, _tokenDecimals uint8) (*types.Transaction, error) {
	return _RenExTokens.contract.Transact(opts, "registerToken", _tokenCode, _tokenAddress, _tokenDecimals)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x9e20a9a0.
//
// Solidity: function registerToken(_tokenCode uint32, _tokenAddress address, _tokenDecimals uint8) returns()
func (_RenExTokens *RenExTokensSession) RegisterToken(_tokenCode uint32, _tokenAddress common.Address, _tokenDecimals uint8) (*types.Transaction, error) {
	return _RenExTokens.Contract.RegisterToken(&_RenExTokens.TransactOpts, _tokenCode, _tokenAddress, _tokenDecimals)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x9e20a9a0.
//
// Solidity: function registerToken(_tokenCode uint32, _tokenAddress address, _tokenDecimals uint8) returns()
func (_RenExTokens *RenExTokensTransactorSession) RegisterToken(_tokenCode uint32, _tokenAddress common.Address, _tokenDecimals uint8) (*types.Transaction, error) {
	return _RenExTokens.Contract.RegisterToken(&_RenExTokens.TransactOpts, _tokenCode, _tokenAddress, _tokenDecimals)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RenExTokens *RenExTokensTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RenExTokens.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RenExTokens *RenExTokensSession) RenounceOwnership() (*types.Transaction, error) {
	return _RenExTokens.Contract.RenounceOwnership(&_RenExTokens.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RenExTokens *RenExTokensTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RenExTokens.Contract.RenounceOwnership(&_RenExTokens.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_RenExTokens *RenExTokensTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _RenExTokens.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_RenExTokens *RenExTokensSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _RenExTokens.Contract.TransferOwnership(&_RenExTokens.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_RenExTokens *RenExTokensTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _RenExTokens.Contract.TransferOwnership(&_RenExTokens.TransactOpts, _newOwner)
}

// RenExTokensOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the RenExTokens contract.
type RenExTokensOwnershipRenouncedIterator struct {
	Event *RenExTokensOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *RenExTokensOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RenExTokensOwnershipRenounced)
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
		it.Event = new(RenExTokensOwnershipRenounced)
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
func (it *RenExTokensOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RenExTokensOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RenExTokensOwnershipRenounced represents a OwnershipRenounced event raised by the RenExTokens contract.
type RenExTokensOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(previousOwner indexed address)
func (_RenExTokens *RenExTokensFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*RenExTokensOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _RenExTokens.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RenExTokensOwnershipRenouncedIterator{contract: _RenExTokens.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(previousOwner indexed address)
func (_RenExTokens *RenExTokensFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *RenExTokensOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _RenExTokens.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RenExTokensOwnershipRenounced)
				if err := _RenExTokens.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// RenExTokensOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RenExTokens contract.
type RenExTokensOwnershipTransferredIterator struct {
	Event *RenExTokensOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RenExTokensOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RenExTokensOwnershipTransferred)
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
		it.Event = new(RenExTokensOwnershipTransferred)
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
func (it *RenExTokensOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RenExTokensOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RenExTokensOwnershipTransferred represents a OwnershipTransferred event raised by the RenExTokens contract.
type RenExTokensOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_RenExTokens *RenExTokensFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RenExTokensOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RenExTokens.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RenExTokensOwnershipTransferredIterator{contract: _RenExTokens.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_RenExTokens *RenExTokensFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RenExTokensOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RenExTokens.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RenExTokensOwnershipTransferred)
				if err := _RenExTokens.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
const SafeMathBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a7230582016a0a1af192929d93f7d51b374b195860a17181b6e0d5354551f111e07d0acc40029`

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// SettlementUtilsABI is the input ABI used to generate the binding from.
const SettlementUtilsABI = "[]"

// SettlementUtilsBin is the compiled bytecode used for deploying new contracts.
const SettlementUtilsBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a7230582049a85a52091bbb6a3e8c39edc8c03d8310a38ecd9e28b9d334d537bb91461c220029`

// DeploySettlementUtils deploys a new Ethereum contract, binding an instance of SettlementUtils to it.
func DeploySettlementUtils(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SettlementUtils, error) {
	parsed, err := abi.JSON(strings.NewReader(SettlementUtilsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SettlementUtilsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SettlementUtils{SettlementUtilsCaller: SettlementUtilsCaller{contract: contract}, SettlementUtilsTransactor: SettlementUtilsTransactor{contract: contract}, SettlementUtilsFilterer: SettlementUtilsFilterer{contract: contract}}, nil
}

// SettlementUtils is an auto generated Go binding around an Ethereum contract.
type SettlementUtils struct {
	SettlementUtilsCaller     // Read-only binding to the contract
	SettlementUtilsTransactor // Write-only binding to the contract
	SettlementUtilsFilterer   // Log filterer for contract events
}

// SettlementUtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type SettlementUtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SettlementUtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SettlementUtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SettlementUtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SettlementUtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SettlementUtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SettlementUtilsSession struct {
	Contract     *SettlementUtils  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SettlementUtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SettlementUtilsCallerSession struct {
	Contract *SettlementUtilsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// SettlementUtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SettlementUtilsTransactorSession struct {
	Contract     *SettlementUtilsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// SettlementUtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type SettlementUtilsRaw struct {
	Contract *SettlementUtils // Generic contract binding to access the raw methods on
}

// SettlementUtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SettlementUtilsCallerRaw struct {
	Contract *SettlementUtilsCaller // Generic read-only contract binding to access the raw methods on
}

// SettlementUtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SettlementUtilsTransactorRaw struct {
	Contract *SettlementUtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSettlementUtils creates a new instance of SettlementUtils, bound to a specific deployed contract.
func NewSettlementUtils(address common.Address, backend bind.ContractBackend) (*SettlementUtils, error) {
	contract, err := bindSettlementUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SettlementUtils{SettlementUtilsCaller: SettlementUtilsCaller{contract: contract}, SettlementUtilsTransactor: SettlementUtilsTransactor{contract: contract}, SettlementUtilsFilterer: SettlementUtilsFilterer{contract: contract}}, nil
}

// NewSettlementUtilsCaller creates a new read-only instance of SettlementUtils, bound to a specific deployed contract.
func NewSettlementUtilsCaller(address common.Address, caller bind.ContractCaller) (*SettlementUtilsCaller, error) {
	contract, err := bindSettlementUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SettlementUtilsCaller{contract: contract}, nil
}

// NewSettlementUtilsTransactor creates a new write-only instance of SettlementUtils, bound to a specific deployed contract.
func NewSettlementUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*SettlementUtilsTransactor, error) {
	contract, err := bindSettlementUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SettlementUtilsTransactor{contract: contract}, nil
}

// NewSettlementUtilsFilterer creates a new log filterer instance of SettlementUtils, bound to a specific deployed contract.
func NewSettlementUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*SettlementUtilsFilterer, error) {
	contract, err := bindSettlementUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SettlementUtilsFilterer{contract: contract}, nil
}

// bindSettlementUtils binds a generic wrapper to an already deployed contract.
func bindSettlementUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SettlementUtilsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SettlementUtils *SettlementUtilsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SettlementUtils.Contract.SettlementUtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SettlementUtils *SettlementUtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SettlementUtils.Contract.SettlementUtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SettlementUtils *SettlementUtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SettlementUtils.Contract.SettlementUtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SettlementUtils *SettlementUtilsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SettlementUtils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SettlementUtils *SettlementUtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SettlementUtils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SettlementUtils *SettlementUtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SettlementUtils.Contract.contract.Transact(opts, method, params...)
}
