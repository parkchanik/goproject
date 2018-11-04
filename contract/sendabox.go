// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sendabox

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

// SendaboxABI is the input ABI used to generate the binding from.
const SendaboxABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"SZABO_PER_WEI\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nowBoxid\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"message\",\"type\":\"string\"}],\"name\":\"Contract_SendABox\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_box_idx\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_token\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_message\",\"type\":\"string\"}],\"name\":\"ev_SendABoxEvent\",\"type\":\"event\"}]"

// SendaboxBin is the compiled bytecode used for deploying new contracts.
const SendaboxBin = `6080604052600060035534801561001557600080fd5b506000600381905550336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555061045d8061006d6000396000f30060806040526004361061006d576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680630478ef4b14610072578063667229aa1461009d57806370a08231146100c85780638da5cb5b1461011f578063ba757abb14610176575b600080fd5b34801561007e57600080fd5b506100876101d2565b6040518082815260200191505060405180910390f35b3480156100a957600080fd5b506100b26101dd565b6040518082815260200191505060405180910390f35b3480156100d457600080fd5b50610109600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506101e7565b6040518082815260200191505060405180910390f35b34801561012b57600080fd5b50610134610230565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6101d0600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610255565b005b662386f26fc1000081565b6000600354905090565b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600080600560009054906101000a900460ff1615151561027457600080fd5b34915060016003540160038190555061029d662386f26fc10000836103ff90919063ffffffff16565b90506102f181600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461041590919063ffffffff16565b600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055503373ffffffffffffffffffffffffffffffffffffffff166003547f88fbfc87ab013b98e2bd419a0d8974220e533276ec5e3bb63b7bff2bca27694a8484876040518084815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b838110156103be5780820151818401526020810190506103a3565b50505050905090810190601f1680156103eb5780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a3505050565b6000818381151561040c57fe5b04905092915050565b6000818301905082811015151561042857fe5b809050929150505600a165627a7a723058200c308850cf5787a980efd004b6ffab14f58838f527d4624bd80c0e589621ce470029`

// DeploySendabox deploys a new Ethereum contract, binding an instance of Sendabox to it.
func DeploySendabox(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Sendabox, error) {
	parsed, err := abi.JSON(strings.NewReader(SendaboxABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SendaboxBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Sendabox{SendaboxCaller: SendaboxCaller{contract: contract}, SendaboxTransactor: SendaboxTransactor{contract: contract}, SendaboxFilterer: SendaboxFilterer{contract: contract}}, nil
}

// Sendabox is an auto generated Go binding around an Ethereum contract.
type Sendabox struct {
	SendaboxCaller     // Read-only binding to the contract
	SendaboxTransactor // Write-only binding to the contract
	SendaboxFilterer   // Log filterer for contract events
}

// SendaboxCaller is an auto generated read-only Go binding around an Ethereum contract.
type SendaboxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SendaboxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SendaboxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SendaboxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SendaboxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SendaboxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SendaboxSession struct {
	Contract     *Sendabox         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SendaboxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SendaboxCallerSession struct {
	Contract *SendaboxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SendaboxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SendaboxTransactorSession struct {
	Contract     *SendaboxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SendaboxRaw is an auto generated low-level Go binding around an Ethereum contract.
type SendaboxRaw struct {
	Contract *Sendabox // Generic contract binding to access the raw methods on
}

// SendaboxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SendaboxCallerRaw struct {
	Contract *SendaboxCaller // Generic read-only contract binding to access the raw methods on
}

// SendaboxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SendaboxTransactorRaw struct {
	Contract *SendaboxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSendabox creates a new instance of Sendabox, bound to a specific deployed contract.
func NewSendabox(address common.Address, backend bind.ContractBackend) (*Sendabox, error) {
	contract, err := bindSendabox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Sendabox{SendaboxCaller: SendaboxCaller{contract: contract}, SendaboxTransactor: SendaboxTransactor{contract: contract}, SendaboxFilterer: SendaboxFilterer{contract: contract}}, nil
}

// NewSendaboxCaller creates a new read-only instance of Sendabox, bound to a specific deployed contract.
func NewSendaboxCaller(address common.Address, caller bind.ContractCaller) (*SendaboxCaller, error) {
	contract, err := bindSendabox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SendaboxCaller{contract: contract}, nil
}

// NewSendaboxTransactor creates a new write-only instance of Sendabox, bound to a specific deployed contract.
func NewSendaboxTransactor(address common.Address, transactor bind.ContractTransactor) (*SendaboxTransactor, error) {
	contract, err := bindSendabox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SendaboxTransactor{contract: contract}, nil
}

// NewSendaboxFilterer creates a new log filterer instance of Sendabox, bound to a specific deployed contract.
func NewSendaboxFilterer(address common.Address, filterer bind.ContractFilterer) (*SendaboxFilterer, error) {
	contract, err := bindSendabox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SendaboxFilterer{contract: contract}, nil
}

// bindSendabox binds a generic wrapper to an already deployed contract.
func bindSendabox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SendaboxABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sendabox *SendaboxRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Sendabox.Contract.SendaboxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sendabox *SendaboxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sendabox.Contract.SendaboxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sendabox *SendaboxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sendabox.Contract.SendaboxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sendabox *SendaboxCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Sendabox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sendabox *SendaboxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sendabox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sendabox *SendaboxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sendabox.Contract.contract.Transact(opts, method, params...)
}

// SZABOPERWEI is a free data retrieval call binding the contract method 0x0478ef4b.
//
// Solidity: function SZABO_PER_WEI() constant returns(uint256)
func (_Sendabox *SendaboxCaller) SZABOPERWEI(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Sendabox.contract.Call(opts, out, "SZABO_PER_WEI")
	return *ret0, err
}

// SZABOPERWEI is a free data retrieval call binding the contract method 0x0478ef4b.
//
// Solidity: function SZABO_PER_WEI() constant returns(uint256)
func (_Sendabox *SendaboxSession) SZABOPERWEI() (*big.Int, error) {
	return _Sendabox.Contract.SZABOPERWEI(&_Sendabox.CallOpts)
}

// SZABOPERWEI is a free data retrieval call binding the contract method 0x0478ef4b.
//
// Solidity: function SZABO_PER_WEI() constant returns(uint256)
func (_Sendabox *SendaboxCallerSession) SZABOPERWEI() (*big.Int, error) {
	return _Sendabox.Contract.SZABOPERWEI(&_Sendabox.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_Sendabox *SendaboxCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Sendabox.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_Sendabox *SendaboxSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Sendabox.Contract.BalanceOf(&_Sendabox.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_Sendabox *SendaboxCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Sendabox.Contract.BalanceOf(&_Sendabox.CallOpts, _owner)
}

// NowBoxid is a free data retrieval call binding the contract method 0x667229aa.
//
// Solidity: function nowBoxid() constant returns(uint256)
func (_Sendabox *SendaboxCaller) NowBoxid(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Sendabox.contract.Call(opts, out, "nowBoxid")
	return *ret0, err
}

// NowBoxid is a free data retrieval call binding the contract method 0x667229aa.
//
// Solidity: function nowBoxid() constant returns(uint256)
func (_Sendabox *SendaboxSession) NowBoxid() (*big.Int, error) {
	return _Sendabox.Contract.NowBoxid(&_Sendabox.CallOpts)
}

// NowBoxid is a free data retrieval call binding the contract method 0x667229aa.
//
// Solidity: function nowBoxid() constant returns(uint256)
func (_Sendabox *SendaboxCallerSession) NowBoxid() (*big.Int, error) {
	return _Sendabox.Contract.NowBoxid(&_Sendabox.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Sendabox *SendaboxCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Sendabox.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Sendabox *SendaboxSession) Owner() (common.Address, error) {
	return _Sendabox.Contract.Owner(&_Sendabox.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Sendabox *SendaboxCallerSession) Owner() (common.Address, error) {
	return _Sendabox.Contract.Owner(&_Sendabox.CallOpts)
}

// ContractSendABox is a paid mutator transaction binding the contract method 0xba757abb.
//
// Solidity: function Contract_SendABox(message string) returns()
func (_Sendabox *SendaboxTransactor) ContractSendABox(opts *bind.TransactOpts, message string) (*types.Transaction, error) {
	return _Sendabox.contract.Transact(opts, "Contract_SendABox", message)
}

// ContractSendABox is a paid mutator transaction binding the contract method 0xba757abb.
//
// Solidity: function Contract_SendABox(message string) returns()
func (_Sendabox *SendaboxSession) ContractSendABox(message string) (*types.Transaction, error) {
	return _Sendabox.Contract.ContractSendABox(&_Sendabox.TransactOpts, message)
}

// ContractSendABox is a paid mutator transaction binding the contract method 0xba757abb.
//
// Solidity: function Contract_SendABox(message string) returns()
func (_Sendabox *SendaboxTransactorSession) ContractSendABox(message string) (*types.Transaction, error) {
	return _Sendabox.Contract.ContractSendABox(&_Sendabox.TransactOpts, message)
}

// SendaboxEvSendABoxEventIterator is returned from FilterEvSendABoxEvent and is used to iterate over the raw logs and unpacked data for EvSendABoxEvent events raised by the Sendabox contract.
type SendaboxEvSendABoxEventIterator struct {
	Event *SendaboxEvSendABoxEvent // Event containing the contract specifics and raw log

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
func (it *SendaboxEvSendABoxEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SendaboxEvSendABoxEvent)
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
		it.Event = new(SendaboxEvSendABoxEvent)
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
func (it *SendaboxEvSendABoxEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SendaboxEvSendABoxEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SendaboxEvSendABoxEvent represents a EvSendABoxEvent event raised by the Sendabox contract.
type SendaboxEvSendABoxEvent struct {
	BoxIdx  *big.Int
	Sender  common.Address
	Value   *big.Int
	Token   *big.Int
	Message string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterEvSendABoxEvent is a free log retrieval operation binding the contract event 0x88fbfc87ab013b98e2bd419a0d8974220e533276ec5e3bb63b7bff2bca27694a.
//
// Solidity: e ev_SendABoxEvent(_box_idx indexed uint256, _sender indexed address, _value uint256, _token uint256, _message string)
func (_Sendabox *SendaboxFilterer) FilterEvSendABoxEvent(opts *bind.FilterOpts, _box_idx []*big.Int, _sender []common.Address) (*SendaboxEvSendABoxEventIterator, error) {

	var _box_idxRule []interface{}
	for _, _box_idxItem := range _box_idx {
		_box_idxRule = append(_box_idxRule, _box_idxItem)
	}
	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _Sendabox.contract.FilterLogs(opts, "ev_SendABoxEvent", _box_idxRule, _senderRule)
	if err != nil {
		return nil, err
	}
	return &SendaboxEvSendABoxEventIterator{contract: _Sendabox.contract, event: "ev_SendABoxEvent", logs: logs, sub: sub}, nil
}

// WatchEvSendABoxEvent is a free log subscription operation binding the contract event 0x88fbfc87ab013b98e2bd419a0d8974220e533276ec5e3bb63b7bff2bca27694a.
//
// Solidity: e ev_SendABoxEvent(_box_idx indexed uint256, _sender indexed address, _value uint256, _token uint256, _message string)
func (_Sendabox *SendaboxFilterer) WatchEvSendABoxEvent(opts *bind.WatchOpts, sink chan<- *SendaboxEvSendABoxEvent, _box_idx []*big.Int, _sender []common.Address) (event.Subscription, error) {

	var _box_idxRule []interface{}
	for _, _box_idxItem := range _box_idx {
		_box_idxRule = append(_box_idxRule, _box_idxItem)
	}
	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _Sendabox.contract.WatchLogs(opts, "ev_SendABoxEvent", _box_idxRule, _senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SendaboxEvSendABoxEvent)
				if err := _Sendabox.contract.UnpackLog(event, "ev_SendABoxEvent", log); err != nil {
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
