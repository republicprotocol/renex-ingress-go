// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// ECRecoveryABI is the input ABI used to generate the binding from.
const ECRecoveryABI = "[]"

// ECRecoveryBin is the compiled bytecode used for deploying new contracts.
const ECRecoveryBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a7230582051ce5400e9db2ec99f60568a05f33dba86f40cf2f0c53aa9337bccedec093bd10029`

// DeployECRecovery deploys a new Ethereum contract, binding an instance of ECRecovery to it.
func DeployECRecovery(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ECRecovery, error) {
	parsed, err := abi.JSON(strings.NewReader(ECRecoveryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ECRecoveryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ECRecovery{ECRecoveryCaller: ECRecoveryCaller{contract: contract}, ECRecoveryTransactor: ECRecoveryTransactor{contract: contract}, ECRecoveryFilterer: ECRecoveryFilterer{contract: contract}}, nil
}

// ECRecovery is an auto generated Go binding around an Ethereum contract.
type ECRecovery struct {
	ECRecoveryCaller     // Read-only binding to the contract
	ECRecoveryTransactor // Write-only binding to the contract
	ECRecoveryFilterer   // Log filterer for contract events
}

// ECRecoveryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ECRecoveryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECRecoveryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ECRecoveryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECRecoveryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ECRecoveryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECRecoverySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ECRecoverySession struct {
	Contract     *ECRecovery       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ECRecoveryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ECRecoveryCallerSession struct {
	Contract *ECRecoveryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ECRecoveryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ECRecoveryTransactorSession struct {
	Contract     *ECRecoveryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ECRecoveryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ECRecoveryRaw struct {
	Contract *ECRecovery // Generic contract binding to access the raw methods on
}

// ECRecoveryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ECRecoveryCallerRaw struct {
	Contract *ECRecoveryCaller // Generic read-only contract binding to access the raw methods on
}

// ECRecoveryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ECRecoveryTransactorRaw struct {
	Contract *ECRecoveryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewECRecovery creates a new instance of ECRecovery, bound to a specific deployed contract.
func NewECRecovery(address common.Address, backend bind.ContractBackend) (*ECRecovery, error) {
	contract, err := bindECRecovery(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ECRecovery{ECRecoveryCaller: ECRecoveryCaller{contract: contract}, ECRecoveryTransactor: ECRecoveryTransactor{contract: contract}, ECRecoveryFilterer: ECRecoveryFilterer{contract: contract}}, nil
}

// NewECRecoveryCaller creates a new read-only instance of ECRecovery, bound to a specific deployed contract.
func NewECRecoveryCaller(address common.Address, caller bind.ContractCaller) (*ECRecoveryCaller, error) {
	contract, err := bindECRecovery(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ECRecoveryCaller{contract: contract}, nil
}

// NewECRecoveryTransactor creates a new write-only instance of ECRecovery, bound to a specific deployed contract.
func NewECRecoveryTransactor(address common.Address, transactor bind.ContractTransactor) (*ECRecoveryTransactor, error) {
	contract, err := bindECRecovery(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ECRecoveryTransactor{contract: contract}, nil
}

// NewECRecoveryFilterer creates a new log filterer instance of ECRecovery, bound to a specific deployed contract.
func NewECRecoveryFilterer(address common.Address, filterer bind.ContractFilterer) (*ECRecoveryFilterer, error) {
	contract, err := bindECRecovery(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ECRecoveryFilterer{contract: contract}, nil
}

// bindECRecovery binds a generic wrapper to an already deployed contract.
func bindECRecovery(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ECRecoveryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECRecovery *ECRecoveryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ECRecovery.Contract.ECRecoveryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECRecovery *ECRecoveryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECRecovery.Contract.ECRecoveryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECRecovery *ECRecoveryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECRecovery.Contract.ECRecoveryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECRecovery *ECRecoveryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ECRecovery.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECRecovery *ECRecoveryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECRecovery.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECRecovery *ECRecoveryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECRecovery.Contract.contract.Transact(opts, method, params...)
}

// OwnableABI is the input ABI used to generate the binding from.
const OwnableABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// OwnableBin is the compiled bytecode used for deploying new contracts.
const OwnableBin = `0x608060405234801561001057600080fd5b5060008054600160a060020a0319163317905561020b806100326000396000f3006080604052600436106100565763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663715018a6811461005b5780638da5cb5b14610072578063f2fde38b146100a3575b600080fd5b34801561006757600080fd5b506100706100c4565b005b34801561007e57600080fd5b50610087610130565b60408051600160a060020a039092168252519081900360200190f35b3480156100af57600080fd5b50610070600160a060020a036004351661013f565b600054600160a060020a031633146100db57600080fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b600054600160a060020a031681565b600054600160a060020a0316331461015657600080fd5b61015f81610162565b50565b600160a060020a038116151561017757600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a72305820d07f4d5221378e0ccfc5230ace30fcd5568360922dc2baa242decb0ba9afd4100029`

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
// Solidity: e OwnershipRenounced(previousOwner indexed address)
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
// Solidity: e OwnershipRenounced(previousOwner indexed address)
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
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
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
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
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

// RenExBrokerVerifierABI is the input ABI used to generate the binding from.
const RenExBrokerVerifierABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_trader\",\"type\":\"address\"},{\"name\":\"_signature\",\"type\":\"bytes\"},{\"name\":\"_orderID\",\"type\":\"bytes32\"}],\"name\":\"verifyOpenSignature\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_broker\",\"type\":\"address\"}],\"name\":\"deregisterBroker\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"balancesContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"traderNonces\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"brokers\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_balancesContract\",\"type\":\"address\"}],\"name\":\"updateBalancesContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_trader\",\"type\":\"address\"},{\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"verifyWithdrawSignature\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_broker\",\"type\":\"address\"}],\"name\":\"registerBroker\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_VERSION\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"previousBalancesContract\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"nextBalancesContract\",\"type\":\"address\"}],\"name\":\"LogBalancesContractUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"broker\",\"type\":\"address\"}],\"name\":\"LogBrokerRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"broker\",\"type\":\"address\"}],\"name\":\"LogBrokerDeregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// RenExBrokerVerifierBin is the compiled bytecode used for deploying new contracts.
const RenExBrokerVerifierBin = `0x608060405234801561001057600080fd5b50604051610e81380380610e8183398101604052805160008054600160a060020a0319163317905501805161004c906001906020840190610053565b50506100ee565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061009457805160ff19168380011785556100c1565b828001600101855582156100c1579182015b828111156100c15782518255916020019190600101906100a6565b506100cd9291506100d1565b5090565b6100eb91905b808211156100cd57600081556001016100d7565b90565b610d84806100fd6000396000f3006080604052600436106100b95763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663472e191081146100be578063490618d1146101025780634a3d72a114610125578063506ee1ef1461015657806366874cc514610189578063715018a6146101aa5780638da5cb5b146101bf578063b8fd1e10146101d4578063c043df8c146101f5578063c684286814610222578063f2fde38b14610243578063ffa1ad7414610264575b600080fd5b3480156100ca57600080fd5b506100ee60048035600160a060020a031690602480359081019101356044356102ee565b604080519115158252519081900360200190f35b34801561010e57600080fd5b50610123600160a060020a03600435166103b8565b005b34801561013157600080fd5b5061013a6104b0565b60408051600160a060020a039092168252519081900360200190f35b34801561016257600080fd5b50610177600160a060020a03600435166104bf565b60408051918252519081900360200190f35b34801561019557600080fd5b506100ee600160a060020a03600435166104d1565b3480156101b657600080fd5b506101236104e6565b3480156101cb57600080fd5b5061013a610552565b3480156101e057600080fd5b50610123600160a060020a0360043516610561565b34801561020157600080fd5b506100ee60048035600160a060020a031690602480359081019101356105ef565b34801561022e57600080fd5b50610123600160a060020a0360043516610775565b34801561024f57600080fd5b50610123600160a060020a036004351661086f565b34801561027057600080fd5b50610279610892565b6040805160208082528351818301528351919283929083019185019080838360005b838110156102b357818101518382015260200161029b565b50505050905090810190601f1680156102e05780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b604080517f52657075626c69632050726f746f636f6c3a206f70656e3a20000000000000006020808301919091526c01000000000000000000000000600160a060020a038816026039830152604d80830185905283518084039091018152608d601f870183900490920283018201909352606d820185815260009392849261038d92859290918a918a918291018382808284375061091f945050505050565b600160a060020a031660009081526002602052604090205460ff161515600114979650505050505050565b600054600160a060020a031633146103cf57600080fd5b600160a060020a03811660009081526002602052604090205460ff16151561045857604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600e60248201527f6e6f742072656769737465726564000000000000000000000000000000000000604482015290519081900360640190fd5b600160a060020a038116600081815260026020908152604091829020805460ff19169055815192835290517fe470a29f46ba9a09f7ec358ae2eb422a5a8f941f128ed7d8f5cf35278ab216409281900390910190a150565b600454600160a060020a031681565b60036020526000908152604090205481565b60026020526000908152604090205460ff1681565b600054600160a060020a031633146104fd57600080fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b600054600160a060020a031681565b600054600160a060020a0316331461057857600080fd5b60045460408051600160a060020a039283168152918316602083015280517fa5d8d37e938531194e3b63a63c76ccbc603ebedcf3f3ebb9e02fc4ba843d34e29281900390910190a16004805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b6004546000906060908290600160a060020a0316331461067057604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600e60248201527f6e6f7420617574686f72697a6564000000000000000000000000000000000000604482015290519081900360640190fd5b600160a060020a0386166000818152600360209081526040918290205482517f52657075626c69632050726f746f636f6c3a2077697468647261773a20000000818401526c01000000000000000000000000909402603d850152605180850191909152825180850390910181526091601f8901839004909202840182019092526071830187815291945061071a92859291899189918291018382808284375061091f945050505050565b600160a060020a03811660009081526002602052604090205490915060ff161561076757600160a060020a038616600090815260036020526040902080546001908101909155925061076c565b600092505b50509392505050565b600054600160a060020a0316331461078c57600080fd5b600160a060020a03811660009081526002602052604090205460ff161561081457604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f616c726561647920726567697374657265640000000000000000000000000000604482015290519081900360640190fd5b600160a060020a038116600081815260026020908152604091829020805460ff19166001179055815192835290517fd4ba9549a2404d1e5bedd0a4ae90c79e2b41ce4dea6bef98dc999fec1f2784939281900390910190a150565b600054600160a060020a0316331461088657600080fd5b61088f81610ad9565b50565b60018054604080516020600284861615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156109175780601f106108ec57610100808354040283529160200191610917565b820191906000526020600020905b8154815290600101906020018083116108fa57829003601f168201915b505050505081565b600060608060006040805190810160405280601a81526020017f19457468657265756d205369676e6564204d6573736167653a0a0000000000008152509250826109698751610b56565b876040516020018084805190602001908083835b6020831061099c5780518252601f19909201916020918201910161097d565b51815160209384036101000a600019018019909216911617905286519190930192860191508083835b602083106109e45780518252601f1990920191602091820191016109c5565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b60208310610a2c5780518252601f199092019160209182019101610a0d565b6001836020036101000a03801982511681845116808217855250505050505090500193505050506040516020818303038152906040529150816040518082805190602001908083835b60208310610a945780518252601f199092019160209182019101610a75565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390209050610acd8186610c88565b93505b50505092915050565b600160a060020a0381161515610aee57600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b6060816000808381841515610ba05760408051808201909152600181527f300000000000000000000000000000000000000000000000000000000000000060208201529550610c7e565b600093508492505b6000831115610bc257600a83049250600184019350610ba8565b836040519080825280601f01601f191660200182016040528015610bf0578160200160208202803883390190505b509150600090505b83811015610c7a57600a85066030017f01000000000000000000000000000000000000000000000000000000000000000282600183870303815181101515610c3c57fe5b9060200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600a85049450600101610bf8565b8195505b5050505050919050565b60008060008084516041141515610ca25760009350610ad0565b50505060208201516040830151606084015160001a601b60ff82161015610cc757601b015b8060ff16601b14158015610cdf57508060ff16601c14155b15610ced5760009350610ad0565b60408051600080825260208083018085528a905260ff8516838501526060830187905260808301869052925160019360a0808501949193601f19840193928390039091019190865af1158015610d47573d6000803e3d6000fd5b505050602060405103519350610ad05600a165627a7a723058201b38c93f1101294e29a8b939cfa8c370eef61283def8fb73eb027424927989d30029`

// DeployRenExBrokerVerifier deploys a new Ethereum contract, binding an instance of RenExBrokerVerifier to it.
func DeployRenExBrokerVerifier(auth *bind.TransactOpts, backend bind.ContractBackend, _VERSION string) (common.Address, *types.Transaction, *RenExBrokerVerifier, error) {
	parsed, err := abi.JSON(strings.NewReader(RenExBrokerVerifierABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RenExBrokerVerifierBin), backend, _VERSION)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RenExBrokerVerifier{RenExBrokerVerifierCaller: RenExBrokerVerifierCaller{contract: contract}, RenExBrokerVerifierTransactor: RenExBrokerVerifierTransactor{contract: contract}, RenExBrokerVerifierFilterer: RenExBrokerVerifierFilterer{contract: contract}}, nil
}

// RenExBrokerVerifier is an auto generated Go binding around an Ethereum contract.
type RenExBrokerVerifier struct {
	RenExBrokerVerifierCaller     // Read-only binding to the contract
	RenExBrokerVerifierTransactor // Write-only binding to the contract
	RenExBrokerVerifierFilterer   // Log filterer for contract events
}

// RenExBrokerVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type RenExBrokerVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RenExBrokerVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RenExBrokerVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RenExBrokerVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RenExBrokerVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RenExBrokerVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RenExBrokerVerifierSession struct {
	Contract     *RenExBrokerVerifier // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// RenExBrokerVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RenExBrokerVerifierCallerSession struct {
	Contract *RenExBrokerVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// RenExBrokerVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RenExBrokerVerifierTransactorSession struct {
	Contract     *RenExBrokerVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// RenExBrokerVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type RenExBrokerVerifierRaw struct {
	Contract *RenExBrokerVerifier // Generic contract binding to access the raw methods on
}

// RenExBrokerVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RenExBrokerVerifierCallerRaw struct {
	Contract *RenExBrokerVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// RenExBrokerVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RenExBrokerVerifierTransactorRaw struct {
	Contract *RenExBrokerVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRenExBrokerVerifier creates a new instance of RenExBrokerVerifier, bound to a specific deployed contract.
func NewRenExBrokerVerifier(address common.Address, backend bind.ContractBackend) (*RenExBrokerVerifier, error) {
	contract, err := bindRenExBrokerVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RenExBrokerVerifier{RenExBrokerVerifierCaller: RenExBrokerVerifierCaller{contract: contract}, RenExBrokerVerifierTransactor: RenExBrokerVerifierTransactor{contract: contract}, RenExBrokerVerifierFilterer: RenExBrokerVerifierFilterer{contract: contract}}, nil
}

// NewRenExBrokerVerifierCaller creates a new read-only instance of RenExBrokerVerifier, bound to a specific deployed contract.
func NewRenExBrokerVerifierCaller(address common.Address, caller bind.ContractCaller) (*RenExBrokerVerifierCaller, error) {
	contract, err := bindRenExBrokerVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RenExBrokerVerifierCaller{contract: contract}, nil
}

// NewRenExBrokerVerifierTransactor creates a new write-only instance of RenExBrokerVerifier, bound to a specific deployed contract.
func NewRenExBrokerVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*RenExBrokerVerifierTransactor, error) {
	contract, err := bindRenExBrokerVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RenExBrokerVerifierTransactor{contract: contract}, nil
}

// NewRenExBrokerVerifierFilterer creates a new log filterer instance of RenExBrokerVerifier, bound to a specific deployed contract.
func NewRenExBrokerVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*RenExBrokerVerifierFilterer, error) {
	contract, err := bindRenExBrokerVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RenExBrokerVerifierFilterer{contract: contract}, nil
}

// bindRenExBrokerVerifier binds a generic wrapper to an already deployed contract.
func bindRenExBrokerVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RenExBrokerVerifierABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RenExBrokerVerifier *RenExBrokerVerifierRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RenExBrokerVerifier.Contract.RenExBrokerVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RenExBrokerVerifier *RenExBrokerVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RenExBrokerVerifier.Contract.RenExBrokerVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RenExBrokerVerifier *RenExBrokerVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RenExBrokerVerifier.Contract.RenExBrokerVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RenExBrokerVerifier *RenExBrokerVerifierCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RenExBrokerVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RenExBrokerVerifier *RenExBrokerVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RenExBrokerVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RenExBrokerVerifier *RenExBrokerVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RenExBrokerVerifier.Contract.contract.Transact(opts, method, params...)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() constant returns(string)
func (_RenExBrokerVerifier *RenExBrokerVerifierCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _RenExBrokerVerifier.contract.Call(opts, out, "VERSION")
	return *ret0, err
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() constant returns(string)
func (_RenExBrokerVerifier *RenExBrokerVerifierSession) VERSION() (string, error) {
	return _RenExBrokerVerifier.Contract.VERSION(&_RenExBrokerVerifier.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() constant returns(string)
func (_RenExBrokerVerifier *RenExBrokerVerifierCallerSession) VERSION() (string, error) {
	return _RenExBrokerVerifier.Contract.VERSION(&_RenExBrokerVerifier.CallOpts)
}

// BalancesContract is a free data retrieval call binding the contract method 0x4a3d72a1.
//
// Solidity: function balancesContract() constant returns(address)
func (_RenExBrokerVerifier *RenExBrokerVerifierCaller) BalancesContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RenExBrokerVerifier.contract.Call(opts, out, "balancesContract")
	return *ret0, err
}

// BalancesContract is a free data retrieval call binding the contract method 0x4a3d72a1.
//
// Solidity: function balancesContract() constant returns(address)
func (_RenExBrokerVerifier *RenExBrokerVerifierSession) BalancesContract() (common.Address, error) {
	return _RenExBrokerVerifier.Contract.BalancesContract(&_RenExBrokerVerifier.CallOpts)
}

// BalancesContract is a free data retrieval call binding the contract method 0x4a3d72a1.
//
// Solidity: function balancesContract() constant returns(address)
func (_RenExBrokerVerifier *RenExBrokerVerifierCallerSession) BalancesContract() (common.Address, error) {
	return _RenExBrokerVerifier.Contract.BalancesContract(&_RenExBrokerVerifier.CallOpts)
}

// Brokers is a free data retrieval call binding the contract method 0x66874cc5.
//
// Solidity: function brokers( address) constant returns(bool)
func (_RenExBrokerVerifier *RenExBrokerVerifierCaller) Brokers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RenExBrokerVerifier.contract.Call(opts, out, "brokers", arg0)
	return *ret0, err
}

// Brokers is a free data retrieval call binding the contract method 0x66874cc5.
//
// Solidity: function brokers( address) constant returns(bool)
func (_RenExBrokerVerifier *RenExBrokerVerifierSession) Brokers(arg0 common.Address) (bool, error) {
	return _RenExBrokerVerifier.Contract.Brokers(&_RenExBrokerVerifier.CallOpts, arg0)
}

// Brokers is a free data retrieval call binding the contract method 0x66874cc5.
//
// Solidity: function brokers( address) constant returns(bool)
func (_RenExBrokerVerifier *RenExBrokerVerifierCallerSession) Brokers(arg0 common.Address) (bool, error) {
	return _RenExBrokerVerifier.Contract.Brokers(&_RenExBrokerVerifier.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RenExBrokerVerifier *RenExBrokerVerifierCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RenExBrokerVerifier.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RenExBrokerVerifier *RenExBrokerVerifierSession) Owner() (common.Address, error) {
	return _RenExBrokerVerifier.Contract.Owner(&_RenExBrokerVerifier.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RenExBrokerVerifier *RenExBrokerVerifierCallerSession) Owner() (common.Address, error) {
	return _RenExBrokerVerifier.Contract.Owner(&_RenExBrokerVerifier.CallOpts)
}

// TraderNonces is a free data retrieval call binding the contract method 0x506ee1ef.
//
// Solidity: function traderNonces( address) constant returns(uint256)
func (_RenExBrokerVerifier *RenExBrokerVerifierCaller) TraderNonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RenExBrokerVerifier.contract.Call(opts, out, "traderNonces", arg0)
	return *ret0, err
}

// TraderNonces is a free data retrieval call binding the contract method 0x506ee1ef.
//
// Solidity: function traderNonces( address) constant returns(uint256)
func (_RenExBrokerVerifier *RenExBrokerVerifierSession) TraderNonces(arg0 common.Address) (*big.Int, error) {
	return _RenExBrokerVerifier.Contract.TraderNonces(&_RenExBrokerVerifier.CallOpts, arg0)
}

// TraderNonces is a free data retrieval call binding the contract method 0x506ee1ef.
//
// Solidity: function traderNonces( address) constant returns(uint256)
func (_RenExBrokerVerifier *RenExBrokerVerifierCallerSession) TraderNonces(arg0 common.Address) (*big.Int, error) {
	return _RenExBrokerVerifier.Contract.TraderNonces(&_RenExBrokerVerifier.CallOpts, arg0)
}

// VerifyOpenSignature is a free data retrieval call binding the contract method 0x472e1910.
//
// Solidity: function verifyOpenSignature(_trader address, _signature bytes, _orderID bytes32) constant returns(bool)
func (_RenExBrokerVerifier *RenExBrokerVerifierCaller) VerifyOpenSignature(opts *bind.CallOpts, _trader common.Address, _signature []byte, _orderID [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RenExBrokerVerifier.contract.Call(opts, out, "verifyOpenSignature", _trader, _signature, _orderID)
	return *ret0, err
}

// VerifyOpenSignature is a free data retrieval call binding the contract method 0x472e1910.
//
// Solidity: function verifyOpenSignature(_trader address, _signature bytes, _orderID bytes32) constant returns(bool)
func (_RenExBrokerVerifier *RenExBrokerVerifierSession) VerifyOpenSignature(_trader common.Address, _signature []byte, _orderID [32]byte) (bool, error) {
	return _RenExBrokerVerifier.Contract.VerifyOpenSignature(&_RenExBrokerVerifier.CallOpts, _trader, _signature, _orderID)
}

// VerifyOpenSignature is a free data retrieval call binding the contract method 0x472e1910.
//
// Solidity: function verifyOpenSignature(_trader address, _signature bytes, _orderID bytes32) constant returns(bool)
func (_RenExBrokerVerifier *RenExBrokerVerifierCallerSession) VerifyOpenSignature(_trader common.Address, _signature []byte, _orderID [32]byte) (bool, error) {
	return _RenExBrokerVerifier.Contract.VerifyOpenSignature(&_RenExBrokerVerifier.CallOpts, _trader, _signature, _orderID)
}

// DeregisterBroker is a paid mutator transaction binding the contract method 0x490618d1.
//
// Solidity: function deregisterBroker(_broker address) returns()
func (_RenExBrokerVerifier *RenExBrokerVerifierTransactor) DeregisterBroker(opts *bind.TransactOpts, _broker common.Address) (*types.Transaction, error) {
	return _RenExBrokerVerifier.contract.Transact(opts, "deregisterBroker", _broker)
}

// DeregisterBroker is a paid mutator transaction binding the contract method 0x490618d1.
//
// Solidity: function deregisterBroker(_broker address) returns()
func (_RenExBrokerVerifier *RenExBrokerVerifierSession) DeregisterBroker(_broker common.Address) (*types.Transaction, error) {
	return _RenExBrokerVerifier.Contract.DeregisterBroker(&_RenExBrokerVerifier.TransactOpts, _broker)
}

// DeregisterBroker is a paid mutator transaction binding the contract method 0x490618d1.
//
// Solidity: function deregisterBroker(_broker address) returns()
func (_RenExBrokerVerifier *RenExBrokerVerifierTransactorSession) DeregisterBroker(_broker common.Address) (*types.Transaction, error) {
	return _RenExBrokerVerifier.Contract.DeregisterBroker(&_RenExBrokerVerifier.TransactOpts, _broker)
}

// RegisterBroker is a paid mutator transaction binding the contract method 0xc6842868.
//
// Solidity: function registerBroker(_broker address) returns()
func (_RenExBrokerVerifier *RenExBrokerVerifierTransactor) RegisterBroker(opts *bind.TransactOpts, _broker common.Address) (*types.Transaction, error) {
	return _RenExBrokerVerifier.contract.Transact(opts, "registerBroker", _broker)
}

// RegisterBroker is a paid mutator transaction binding the contract method 0xc6842868.
//
// Solidity: function registerBroker(_broker address) returns()
func (_RenExBrokerVerifier *RenExBrokerVerifierSession) RegisterBroker(_broker common.Address) (*types.Transaction, error) {
	return _RenExBrokerVerifier.Contract.RegisterBroker(&_RenExBrokerVerifier.TransactOpts, _broker)
}

// RegisterBroker is a paid mutator transaction binding the contract method 0xc6842868.
//
// Solidity: function registerBroker(_broker address) returns()
func (_RenExBrokerVerifier *RenExBrokerVerifierTransactorSession) RegisterBroker(_broker common.Address) (*types.Transaction, error) {
	return _RenExBrokerVerifier.Contract.RegisterBroker(&_RenExBrokerVerifier.TransactOpts, _broker)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RenExBrokerVerifier *RenExBrokerVerifierTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RenExBrokerVerifier.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RenExBrokerVerifier *RenExBrokerVerifierSession) RenounceOwnership() (*types.Transaction, error) {
	return _RenExBrokerVerifier.Contract.RenounceOwnership(&_RenExBrokerVerifier.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RenExBrokerVerifier *RenExBrokerVerifierTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RenExBrokerVerifier.Contract.RenounceOwnership(&_RenExBrokerVerifier.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_RenExBrokerVerifier *RenExBrokerVerifierTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _RenExBrokerVerifier.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_RenExBrokerVerifier *RenExBrokerVerifierSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _RenExBrokerVerifier.Contract.TransferOwnership(&_RenExBrokerVerifier.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_RenExBrokerVerifier *RenExBrokerVerifierTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _RenExBrokerVerifier.Contract.TransferOwnership(&_RenExBrokerVerifier.TransactOpts, _newOwner)
}

// UpdateBalancesContract is a paid mutator transaction binding the contract method 0xb8fd1e10.
//
// Solidity: function updateBalancesContract(_balancesContract address) returns()
func (_RenExBrokerVerifier *RenExBrokerVerifierTransactor) UpdateBalancesContract(opts *bind.TransactOpts, _balancesContract common.Address) (*types.Transaction, error) {
	return _RenExBrokerVerifier.contract.Transact(opts, "updateBalancesContract", _balancesContract)
}

// UpdateBalancesContract is a paid mutator transaction binding the contract method 0xb8fd1e10.
//
// Solidity: function updateBalancesContract(_balancesContract address) returns()
func (_RenExBrokerVerifier *RenExBrokerVerifierSession) UpdateBalancesContract(_balancesContract common.Address) (*types.Transaction, error) {
	return _RenExBrokerVerifier.Contract.UpdateBalancesContract(&_RenExBrokerVerifier.TransactOpts, _balancesContract)
}

// UpdateBalancesContract is a paid mutator transaction binding the contract method 0xb8fd1e10.
//
// Solidity: function updateBalancesContract(_balancesContract address) returns()
func (_RenExBrokerVerifier *RenExBrokerVerifierTransactorSession) UpdateBalancesContract(_balancesContract common.Address) (*types.Transaction, error) {
	return _RenExBrokerVerifier.Contract.UpdateBalancesContract(&_RenExBrokerVerifier.TransactOpts, _balancesContract)
}

// VerifyWithdrawSignature is a paid mutator transaction binding the contract method 0xc043df8c.
//
// Solidity: function verifyWithdrawSignature(_trader address, _signature bytes) returns(bool)
func (_RenExBrokerVerifier *RenExBrokerVerifierTransactor) VerifyWithdrawSignature(opts *bind.TransactOpts, _trader common.Address, _signature []byte) (*types.Transaction, error) {
	return _RenExBrokerVerifier.contract.Transact(opts, "verifyWithdrawSignature", _trader, _signature)
}

// VerifyWithdrawSignature is a paid mutator transaction binding the contract method 0xc043df8c.
//
// Solidity: function verifyWithdrawSignature(_trader address, _signature bytes) returns(bool)
func (_RenExBrokerVerifier *RenExBrokerVerifierSession) VerifyWithdrawSignature(_trader common.Address, _signature []byte) (*types.Transaction, error) {
	return _RenExBrokerVerifier.Contract.VerifyWithdrawSignature(&_RenExBrokerVerifier.TransactOpts, _trader, _signature)
}

// VerifyWithdrawSignature is a paid mutator transaction binding the contract method 0xc043df8c.
//
// Solidity: function verifyWithdrawSignature(_trader address, _signature bytes) returns(bool)
func (_RenExBrokerVerifier *RenExBrokerVerifierTransactorSession) VerifyWithdrawSignature(_trader common.Address, _signature []byte) (*types.Transaction, error) {
	return _RenExBrokerVerifier.Contract.VerifyWithdrawSignature(&_RenExBrokerVerifier.TransactOpts, _trader, _signature)
}

// RenExBrokerVerifierLogBalancesContractUpdatedIterator is returned from FilterLogBalancesContractUpdated and is used to iterate over the raw logs and unpacked data for LogBalancesContractUpdated events raised by the RenExBrokerVerifier contract.
type RenExBrokerVerifierLogBalancesContractUpdatedIterator struct {
	Event *RenExBrokerVerifierLogBalancesContractUpdated // Event containing the contract specifics and raw log

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
func (it *RenExBrokerVerifierLogBalancesContractUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RenExBrokerVerifierLogBalancesContractUpdated)
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
		it.Event = new(RenExBrokerVerifierLogBalancesContractUpdated)
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
func (it *RenExBrokerVerifierLogBalancesContractUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RenExBrokerVerifierLogBalancesContractUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RenExBrokerVerifierLogBalancesContractUpdated represents a LogBalancesContractUpdated event raised by the RenExBrokerVerifier contract.
type RenExBrokerVerifierLogBalancesContractUpdated struct {
	PreviousBalancesContract common.Address
	NextBalancesContract     common.Address
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterLogBalancesContractUpdated is a free log retrieval operation binding the contract event 0xa5d8d37e938531194e3b63a63c76ccbc603ebedcf3f3ebb9e02fc4ba843d34e2.
//
// Solidity: e LogBalancesContractUpdated(previousBalancesContract address, nextBalancesContract address)
func (_RenExBrokerVerifier *RenExBrokerVerifierFilterer) FilterLogBalancesContractUpdated(opts *bind.FilterOpts) (*RenExBrokerVerifierLogBalancesContractUpdatedIterator, error) {

	logs, sub, err := _RenExBrokerVerifier.contract.FilterLogs(opts, "LogBalancesContractUpdated")
	if err != nil {
		return nil, err
	}
	return &RenExBrokerVerifierLogBalancesContractUpdatedIterator{contract: _RenExBrokerVerifier.contract, event: "LogBalancesContractUpdated", logs: logs, sub: sub}, nil
}

// WatchLogBalancesContractUpdated is a free log subscription operation binding the contract event 0xa5d8d37e938531194e3b63a63c76ccbc603ebedcf3f3ebb9e02fc4ba843d34e2.
//
// Solidity: e LogBalancesContractUpdated(previousBalancesContract address, nextBalancesContract address)
func (_RenExBrokerVerifier *RenExBrokerVerifierFilterer) WatchLogBalancesContractUpdated(opts *bind.WatchOpts, sink chan<- *RenExBrokerVerifierLogBalancesContractUpdated) (event.Subscription, error) {

	logs, sub, err := _RenExBrokerVerifier.contract.WatchLogs(opts, "LogBalancesContractUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RenExBrokerVerifierLogBalancesContractUpdated)
				if err := _RenExBrokerVerifier.contract.UnpackLog(event, "LogBalancesContractUpdated", log); err != nil {
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

// RenExBrokerVerifierLogBrokerDeregisteredIterator is returned from FilterLogBrokerDeregistered and is used to iterate over the raw logs and unpacked data for LogBrokerDeregistered events raised by the RenExBrokerVerifier contract.
type RenExBrokerVerifierLogBrokerDeregisteredIterator struct {
	Event *RenExBrokerVerifierLogBrokerDeregistered // Event containing the contract specifics and raw log

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
func (it *RenExBrokerVerifierLogBrokerDeregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RenExBrokerVerifierLogBrokerDeregistered)
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
		it.Event = new(RenExBrokerVerifierLogBrokerDeregistered)
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
func (it *RenExBrokerVerifierLogBrokerDeregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RenExBrokerVerifierLogBrokerDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RenExBrokerVerifierLogBrokerDeregistered represents a LogBrokerDeregistered event raised by the RenExBrokerVerifier contract.
type RenExBrokerVerifierLogBrokerDeregistered struct {
	Broker common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLogBrokerDeregistered is a free log retrieval operation binding the contract event 0xe470a29f46ba9a09f7ec358ae2eb422a5a8f941f128ed7d8f5cf35278ab21640.
//
// Solidity: e LogBrokerDeregistered(broker address)
func (_RenExBrokerVerifier *RenExBrokerVerifierFilterer) FilterLogBrokerDeregistered(opts *bind.FilterOpts) (*RenExBrokerVerifierLogBrokerDeregisteredIterator, error) {

	logs, sub, err := _RenExBrokerVerifier.contract.FilterLogs(opts, "LogBrokerDeregistered")
	if err != nil {
		return nil, err
	}
	return &RenExBrokerVerifierLogBrokerDeregisteredIterator{contract: _RenExBrokerVerifier.contract, event: "LogBrokerDeregistered", logs: logs, sub: sub}, nil
}

// WatchLogBrokerDeregistered is a free log subscription operation binding the contract event 0xe470a29f46ba9a09f7ec358ae2eb422a5a8f941f128ed7d8f5cf35278ab21640.
//
// Solidity: e LogBrokerDeregistered(broker address)
func (_RenExBrokerVerifier *RenExBrokerVerifierFilterer) WatchLogBrokerDeregistered(opts *bind.WatchOpts, sink chan<- *RenExBrokerVerifierLogBrokerDeregistered) (event.Subscription, error) {

	logs, sub, err := _RenExBrokerVerifier.contract.WatchLogs(opts, "LogBrokerDeregistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RenExBrokerVerifierLogBrokerDeregistered)
				if err := _RenExBrokerVerifier.contract.UnpackLog(event, "LogBrokerDeregistered", log); err != nil {
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

// RenExBrokerVerifierLogBrokerRegisteredIterator is returned from FilterLogBrokerRegistered and is used to iterate over the raw logs and unpacked data for LogBrokerRegistered events raised by the RenExBrokerVerifier contract.
type RenExBrokerVerifierLogBrokerRegisteredIterator struct {
	Event *RenExBrokerVerifierLogBrokerRegistered // Event containing the contract specifics and raw log

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
func (it *RenExBrokerVerifierLogBrokerRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RenExBrokerVerifierLogBrokerRegistered)
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
		it.Event = new(RenExBrokerVerifierLogBrokerRegistered)
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
func (it *RenExBrokerVerifierLogBrokerRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RenExBrokerVerifierLogBrokerRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RenExBrokerVerifierLogBrokerRegistered represents a LogBrokerRegistered event raised by the RenExBrokerVerifier contract.
type RenExBrokerVerifierLogBrokerRegistered struct {
	Broker common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLogBrokerRegistered is a free log retrieval operation binding the contract event 0xd4ba9549a2404d1e5bedd0a4ae90c79e2b41ce4dea6bef98dc999fec1f278493.
//
// Solidity: e LogBrokerRegistered(broker address)
func (_RenExBrokerVerifier *RenExBrokerVerifierFilterer) FilterLogBrokerRegistered(opts *bind.FilterOpts) (*RenExBrokerVerifierLogBrokerRegisteredIterator, error) {

	logs, sub, err := _RenExBrokerVerifier.contract.FilterLogs(opts, "LogBrokerRegistered")
	if err != nil {
		return nil, err
	}
	return &RenExBrokerVerifierLogBrokerRegisteredIterator{contract: _RenExBrokerVerifier.contract, event: "LogBrokerRegistered", logs: logs, sub: sub}, nil
}

// WatchLogBrokerRegistered is a free log subscription operation binding the contract event 0xd4ba9549a2404d1e5bedd0a4ae90c79e2b41ce4dea6bef98dc999fec1f278493.
//
// Solidity: e LogBrokerRegistered(broker address)
func (_RenExBrokerVerifier *RenExBrokerVerifierFilterer) WatchLogBrokerRegistered(opts *bind.WatchOpts, sink chan<- *RenExBrokerVerifierLogBrokerRegistered) (event.Subscription, error) {

	logs, sub, err := _RenExBrokerVerifier.contract.WatchLogs(opts, "LogBrokerRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RenExBrokerVerifierLogBrokerRegistered)
				if err := _RenExBrokerVerifier.contract.UnpackLog(event, "LogBrokerRegistered", log); err != nil {
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

// RenExBrokerVerifierOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the RenExBrokerVerifier contract.
type RenExBrokerVerifierOwnershipRenouncedIterator struct {
	Event *RenExBrokerVerifierOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *RenExBrokerVerifierOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RenExBrokerVerifierOwnershipRenounced)
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
		it.Event = new(RenExBrokerVerifierOwnershipRenounced)
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
func (it *RenExBrokerVerifierOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RenExBrokerVerifierOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RenExBrokerVerifierOwnershipRenounced represents a OwnershipRenounced event raised by the RenExBrokerVerifier contract.
type RenExBrokerVerifierOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_RenExBrokerVerifier *RenExBrokerVerifierFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*RenExBrokerVerifierOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _RenExBrokerVerifier.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RenExBrokerVerifierOwnershipRenouncedIterator{contract: _RenExBrokerVerifier.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_RenExBrokerVerifier *RenExBrokerVerifierFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *RenExBrokerVerifierOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _RenExBrokerVerifier.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RenExBrokerVerifierOwnershipRenounced)
				if err := _RenExBrokerVerifier.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// RenExBrokerVerifierOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RenExBrokerVerifier contract.
type RenExBrokerVerifierOwnershipTransferredIterator struct {
	Event *RenExBrokerVerifierOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RenExBrokerVerifierOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RenExBrokerVerifierOwnershipTransferred)
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
		it.Event = new(RenExBrokerVerifierOwnershipTransferred)
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
func (it *RenExBrokerVerifierOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RenExBrokerVerifierOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RenExBrokerVerifierOwnershipTransferred represents a OwnershipTransferred event raised by the RenExBrokerVerifier contract.
type RenExBrokerVerifierOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_RenExBrokerVerifier *RenExBrokerVerifierFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RenExBrokerVerifierOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RenExBrokerVerifier.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RenExBrokerVerifierOwnershipTransferredIterator{contract: _RenExBrokerVerifier.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_RenExBrokerVerifier *RenExBrokerVerifierFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RenExBrokerVerifierOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RenExBrokerVerifier.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RenExBrokerVerifierOwnershipTransferred)
				if err := _RenExBrokerVerifier.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// UtilsABI is the input ABI used to generate the binding from.
const UtilsABI = "[]"

// UtilsBin is the compiled bytecode used for deploying new contracts.
const UtilsBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a7230582001a72e96e30964cc9ea0e910abb2eca43143887c0c9e7633f2a66ce369e3e03d0029`

// DeployUtils deploys a new Ethereum contract, binding an instance of Utils to it.
func DeployUtils(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Utils, error) {
	parsed, err := abi.JSON(strings.NewReader(UtilsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UtilsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Utils{UtilsCaller: UtilsCaller{contract: contract}, UtilsTransactor: UtilsTransactor{contract: contract}, UtilsFilterer: UtilsFilterer{contract: contract}}, nil
}

// Utils is an auto generated Go binding around an Ethereum contract.
type Utils struct {
	UtilsCaller     // Read-only binding to the contract
	UtilsTransactor // Write-only binding to the contract
	UtilsFilterer   // Log filterer for contract events
}

// UtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type UtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UtilsSession struct {
	Contract     *Utils            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UtilsCallerSession struct {
	Contract *UtilsCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// UtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UtilsTransactorSession struct {
	Contract     *UtilsTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type UtilsRaw struct {
	Contract *Utils // Generic contract binding to access the raw methods on
}

// UtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UtilsCallerRaw struct {
	Contract *UtilsCaller // Generic read-only contract binding to access the raw methods on
}

// UtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UtilsTransactorRaw struct {
	Contract *UtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUtils creates a new instance of Utils, bound to a specific deployed contract.
func NewUtils(address common.Address, backend bind.ContractBackend) (*Utils, error) {
	contract, err := bindUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Utils{UtilsCaller: UtilsCaller{contract: contract}, UtilsTransactor: UtilsTransactor{contract: contract}, UtilsFilterer: UtilsFilterer{contract: contract}}, nil
}

// NewUtilsCaller creates a new read-only instance of Utils, bound to a specific deployed contract.
func NewUtilsCaller(address common.Address, caller bind.ContractCaller) (*UtilsCaller, error) {
	contract, err := bindUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UtilsCaller{contract: contract}, nil
}

// NewUtilsTransactor creates a new write-only instance of Utils, bound to a specific deployed contract.
func NewUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*UtilsTransactor, error) {
	contract, err := bindUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UtilsTransactor{contract: contract}, nil
}

// NewUtilsFilterer creates a new log filterer instance of Utils, bound to a specific deployed contract.
func NewUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*UtilsFilterer, error) {
	contract, err := bindUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UtilsFilterer{contract: contract}, nil
}

// bindUtils binds a generic wrapper to an already deployed contract.
func bindUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UtilsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Utils *UtilsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Utils.Contract.UtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Utils *UtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Utils.Contract.UtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Utils *UtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Utils.Contract.UtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Utils *UtilsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Utils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Utils *UtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Utils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Utils *UtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Utils.Contract.contract.Transact(opts, method, params...)
}
