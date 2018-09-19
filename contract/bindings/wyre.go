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

// WyreABI is the input ABI used to generate the binding from.
const WyreABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_approved\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"takeOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenld\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"approvedFor\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"View\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"tokensOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Wyre is an auto generated Go binding around an Ethereum contract.
type Wyre struct {
	WyreCaller     // Read-only binding to the contract
	WyreTransactor // Write-only binding to the contract
	WyreFilterer   // Log filterer for contract events
}

// WyreCaller is an auto generated read-only Go binding around an Ethereum contract.
type WyreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WyreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WyreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WyreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WyreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WyreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WyreSession struct {
	Contract     *Wyre             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WyreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WyreCallerSession struct {
	Contract *WyreCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// WyreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WyreTransactorSession struct {
	Contract     *WyreTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WyreRaw is an auto generated low-level Go binding around an Ethereum contract.
type WyreRaw struct {
	Contract *Wyre // Generic contract binding to access the raw methods on
}

// WyreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WyreCallerRaw struct {
	Contract *WyreCaller // Generic read-only contract binding to access the raw methods on
}

// WyreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WyreTransactorRaw struct {
	Contract *WyreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWyre creates a new instance of Wyre, bound to a specific deployed contract.
func NewWyre(address common.Address, backend bind.ContractBackend) (*Wyre, error) {
	contract, err := bindWyre(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Wyre{WyreCaller: WyreCaller{contract: contract}, WyreTransactor: WyreTransactor{contract: contract}, WyreFilterer: WyreFilterer{contract: contract}}, nil
}

// NewWyreCaller creates a new read-only instance of Wyre, bound to a specific deployed contract.
func NewWyreCaller(address common.Address, caller bind.ContractCaller) (*WyreCaller, error) {
	contract, err := bindWyre(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WyreCaller{contract: contract}, nil
}

// NewWyreTransactor creates a new write-only instance of Wyre, bound to a specific deployed contract.
func NewWyreTransactor(address common.Address, transactor bind.ContractTransactor) (*WyreTransactor, error) {
	contract, err := bindWyre(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WyreTransactor{contract: contract}, nil
}

// NewWyreFilterer creates a new log filterer instance of Wyre, bound to a specific deployed contract.
func NewWyreFilterer(address common.Address, filterer bind.ContractFilterer) (*WyreFilterer, error) {
	contract, err := bindWyre(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WyreFilterer{contract: contract}, nil
}

// bindWyre binds a generic wrapper to an already deployed contract.
func bindWyre(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WyreABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Wyre *WyreRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Wyre.Contract.WyreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Wyre *WyreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wyre.Contract.WyreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Wyre *WyreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Wyre.Contract.WyreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Wyre *WyreCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Wyre.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Wyre *WyreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wyre.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Wyre *WyreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Wyre.Contract.contract.Transact(opts, method, params...)
}

// ApprovedFor is a free data retrieval call binding the contract method 0x2a6dd48f.
//
// Solidity: function approvedFor(_tokenId uint256) constant returns(address)
func (_Wyre *WyreCaller) ApprovedFor(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Wyre.contract.Call(opts, out, "approvedFor", _tokenId)
	return *ret0, err
}

// ApprovedFor is a free data retrieval call binding the contract method 0x2a6dd48f.
//
// Solidity: function approvedFor(_tokenId uint256) constant returns(address)
func (_Wyre *WyreSession) ApprovedFor(_tokenId *big.Int) (common.Address, error) {
	return _Wyre.Contract.ApprovedFor(&_Wyre.CallOpts, _tokenId)
}

// ApprovedFor is a free data retrieval call binding the contract method 0x2a6dd48f.
//
// Solidity: function approvedFor(_tokenId uint256) constant returns(address)
func (_Wyre *WyreCallerSession) ApprovedFor(_tokenId *big.Int) (common.Address, error) {
	return _Wyre.Contract.ApprovedFor(&_Wyre.CallOpts, _tokenId)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_Wyre *WyreCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Wyre.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_Wyre *WyreSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Wyre.Contract.BalanceOf(&_Wyre.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_Wyre *WyreCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Wyre.Contract.BalanceOf(&_Wyre.CallOpts, _owner)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(address)
func (_Wyre *WyreCaller) OwnerOf(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Wyre.contract.Call(opts, out, "ownerOf", _tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(address)
func (_Wyre *WyreSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _Wyre.Contract.OwnerOf(&_Wyre.CallOpts, _tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(address)
func (_Wyre *WyreCallerSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _Wyre.Contract.OwnerOf(&_Wyre.CallOpts, _tokenId)
}

// TokensOf is a free data retrieval call binding the contract method 0x5a3f2672.
//
// Solidity: function tokensOf(_owner address) constant returns(uint256[])
func (_Wyre *WyreCaller) TokensOf(opts *bind.CallOpts, _owner common.Address) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _Wyre.contract.Call(opts, out, "tokensOf", _owner)
	return *ret0, err
}

// TokensOf is a free data retrieval call binding the contract method 0x5a3f2672.
//
// Solidity: function tokensOf(_owner address) constant returns(uint256[])
func (_Wyre *WyreSession) TokensOf(_owner common.Address) ([]*big.Int, error) {
	return _Wyre.Contract.TokensOf(&_Wyre.CallOpts, _owner)
}

// TokensOf is a free data retrieval call binding the contract method 0x5a3f2672.
//
// Solidity: function tokensOf(_owner address) constant returns(uint256[])
func (_Wyre *WyreCallerSession) TokensOf(_owner common.Address) ([]*big.Int, error) {
	return _Wyre.Contract.TokensOf(&_Wyre.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Wyre *WyreCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Wyre.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Wyre *WyreSession) TotalSupply() (*big.Int, error) {
	return _Wyre.Contract.TotalSupply(&_Wyre.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Wyre *WyreCallerSession) TotalSupply() (*big.Int, error) {
	return _Wyre.Contract.TotalSupply(&_Wyre.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_Wyre *WyreTransactor) Approve(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Wyre.contract.Transact(opts, "approve", _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_Wyre *WyreSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Wyre.Contract.Approve(&_Wyre.TransactOpts, _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_Wyre *WyreTransactorSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Wyre.Contract.Approve(&_Wyre.TransactOpts, _to, _tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(_tokenId uint256) returns()
func (_Wyre *WyreTransactor) Burn(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _Wyre.contract.Transact(opts, "burn", _tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(_tokenId uint256) returns()
func (_Wyre *WyreSession) Burn(_tokenId *big.Int) (*types.Transaction, error) {
	return _Wyre.Contract.Burn(&_Wyre.TransactOpts, _tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(_tokenId uint256) returns()
func (_Wyre *WyreTransactorSession) Burn(_tokenId *big.Int) (*types.Transaction, error) {
	return _Wyre.Contract.Burn(&_Wyre.TransactOpts, _tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _tokenId uint256) returns()
func (_Wyre *WyreTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Wyre.contract.Transact(opts, "mint", _to, _tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _tokenId uint256) returns()
func (_Wyre *WyreSession) Mint(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Wyre.Contract.Mint(&_Wyre.TransactOpts, _to, _tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _tokenId uint256) returns()
func (_Wyre *WyreTransactorSession) Mint(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Wyre.Contract.Mint(&_Wyre.TransactOpts, _to, _tokenId)
}

// TakeOwnership is a paid mutator transaction binding the contract method 0xb2e6ceeb.
//
// Solidity: function takeOwnership(_tokenId uint256) returns()
func (_Wyre *WyreTransactor) TakeOwnership(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _Wyre.contract.Transact(opts, "takeOwnership", _tokenId)
}

// TakeOwnership is a paid mutator transaction binding the contract method 0xb2e6ceeb.
//
// Solidity: function takeOwnership(_tokenId uint256) returns()
func (_Wyre *WyreSession) TakeOwnership(_tokenId *big.Int) (*types.Transaction, error) {
	return _Wyre.Contract.TakeOwnership(&_Wyre.TransactOpts, _tokenId)
}

// TakeOwnership is a paid mutator transaction binding the contract method 0xb2e6ceeb.
//
// Solidity: function takeOwnership(_tokenId uint256) returns()
func (_Wyre *WyreTransactorSession) TakeOwnership(_tokenId *big.Int) (*types.Transaction, error) {
	return _Wyre.Contract.TakeOwnership(&_Wyre.TransactOpts, _tokenId)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _tokenld uint256) returns()
func (_Wyre *WyreTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _tokenld *big.Int) (*types.Transaction, error) {
	return _Wyre.contract.Transact(opts, "transfer", _to, _tokenld)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _tokenld uint256) returns()
func (_Wyre *WyreSession) Transfer(_to common.Address, _tokenld *big.Int) (*types.Transaction, error) {
	return _Wyre.Contract.Transfer(&_Wyre.TransactOpts, _to, _tokenld)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _tokenld uint256) returns()
func (_Wyre *WyreTransactorSession) Transfer(_to common.Address, _tokenld *big.Int) (*types.Transaction, error) {
	return _Wyre.Contract.Transfer(&_Wyre.TransactOpts, _to, _tokenld)
}

// WyreApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Wyre contract.
type WyreApprovalIterator struct {
	Event *WyreApproval // Event containing the contract specifics and raw log

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
func (it *WyreApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WyreApproval)
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
		it.Event = new(WyreApproval)
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
func (it *WyreApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WyreApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WyreApproval represents a Approval event raised by the Wyre contract.
type WyreApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _approved indexed address, _tokenId uint256)
func (_Wyre *WyreFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _approved []common.Address) (*WyreApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}

	logs, sub, err := _Wyre.contract.FilterLogs(opts, "Approval", _ownerRule, _approvedRule)
	if err != nil {
		return nil, err
	}
	return &WyreApprovalIterator{contract: _Wyre.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _approved indexed address, _tokenId uint256)
func (_Wyre *WyreFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *WyreApproval, _owner []common.Address, _approved []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}

	logs, sub, err := _Wyre.contract.WatchLogs(opts, "Approval", _ownerRule, _approvedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WyreApproval)
				if err := _Wyre.contract.UnpackLog(event, "Approval", log); err != nil {
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

// WyreTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Wyre contract.
type WyreTransferIterator struct {
	Event *WyreTransfer // Event containing the contract specifics and raw log

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
func (it *WyreTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WyreTransfer)
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
		it.Event = new(WyreTransfer)
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
func (it *WyreTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WyreTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WyreTransfer represents a Transfer event raised by the Wyre contract.
type WyreTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _tokenId uint256)
func (_Wyre *WyreFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*WyreTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Wyre.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &WyreTransferIterator{contract: _Wyre.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _tokenId uint256)
func (_Wyre *WyreFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *WyreTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Wyre.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WyreTransfer)
				if err := _Wyre.contract.UnpackLog(event, "Transfer", log); err != nil {
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
