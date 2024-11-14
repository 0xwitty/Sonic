// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package counter_event_emitter

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
	_ = abi.ConvertType
)

// CounterEventEmitterMetaData contains all meta data concerning the CounterEventEmitter contract.
var CounterEventEmitterMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"count\",\"type\":\"int256\"}],\"name\":\"Count\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getCount\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"increment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040525f5f553480156011575f5ffd5b5061018f8061001f5f395ff3fe608060405234801561000f575f5ffd5b5060043610610034575f3560e01c8063a87d942c14610038578063d09de08a14610056575b5f5ffd5b610040610060565b60405161004d91906100d2565b60405180910390f35b61005e610068565b005b5f5f54905090565b60015f5f8282546100799190610118565b925050819055507f47a90661b520aaf860e17d861200923904ff3dc45e8640c690c905b71719464e5f546040516100b091906100d2565b60405180910390a1565b5f819050919050565b6100cc816100ba565b82525050565b5f6020820190506100e55f8301846100c3565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610122826100ba565b915061012d836100ba565b92508282019050828112155f8312168382125f841215161715610153576101526100eb565b5b9291505056fea26469706673582212208752a265965e5333176ec4a02423eaa602ca03d2be57d67f87cd1d54b06994c064736f6c634300081c0033",
}

// CounterEventEmitterABI is the input ABI used to generate the binding from.
// Deprecated: Use CounterEventEmitterMetaData.ABI instead.
var CounterEventEmitterABI = CounterEventEmitterMetaData.ABI

// CounterEventEmitterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CounterEventEmitterMetaData.Bin instead.
var CounterEventEmitterBin = CounterEventEmitterMetaData.Bin

// DeployCounterEventEmitter deploys a new Ethereum contract, binding an instance of CounterEventEmitter to it.
func DeployCounterEventEmitter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CounterEventEmitter, error) {
	parsed, err := CounterEventEmitterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CounterEventEmitterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CounterEventEmitter{CounterEventEmitterCaller: CounterEventEmitterCaller{contract: contract}, CounterEventEmitterTransactor: CounterEventEmitterTransactor{contract: contract}, CounterEventEmitterFilterer: CounterEventEmitterFilterer{contract: contract}}, nil
}

// CounterEventEmitter is an auto generated Go binding around an Ethereum contract.
type CounterEventEmitter struct {
	CounterEventEmitterCaller     // Read-only binding to the contract
	CounterEventEmitterTransactor // Write-only binding to the contract
	CounterEventEmitterFilterer   // Log filterer for contract events
}

// CounterEventEmitterCaller is an auto generated read-only Go binding around an Ethereum contract.
type CounterEventEmitterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CounterEventEmitterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CounterEventEmitterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CounterEventEmitterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CounterEventEmitterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CounterEventEmitterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CounterEventEmitterSession struct {
	Contract     *CounterEventEmitter // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// CounterEventEmitterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CounterEventEmitterCallerSession struct {
	Contract *CounterEventEmitterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// CounterEventEmitterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CounterEventEmitterTransactorSession struct {
	Contract     *CounterEventEmitterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// CounterEventEmitterRaw is an auto generated low-level Go binding around an Ethereum contract.
type CounterEventEmitterRaw struct {
	Contract *CounterEventEmitter // Generic contract binding to access the raw methods on
}

// CounterEventEmitterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CounterEventEmitterCallerRaw struct {
	Contract *CounterEventEmitterCaller // Generic read-only contract binding to access the raw methods on
}

// CounterEventEmitterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CounterEventEmitterTransactorRaw struct {
	Contract *CounterEventEmitterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCounterEventEmitter creates a new instance of CounterEventEmitter, bound to a specific deployed contract.
func NewCounterEventEmitter(address common.Address, backend bind.ContractBackend) (*CounterEventEmitter, error) {
	contract, err := bindCounterEventEmitter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CounterEventEmitter{CounterEventEmitterCaller: CounterEventEmitterCaller{contract: contract}, CounterEventEmitterTransactor: CounterEventEmitterTransactor{contract: contract}, CounterEventEmitterFilterer: CounterEventEmitterFilterer{contract: contract}}, nil
}

// NewCounterEventEmitterCaller creates a new read-only instance of CounterEventEmitter, bound to a specific deployed contract.
func NewCounterEventEmitterCaller(address common.Address, caller bind.ContractCaller) (*CounterEventEmitterCaller, error) {
	contract, err := bindCounterEventEmitter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CounterEventEmitterCaller{contract: contract}, nil
}

// NewCounterEventEmitterTransactor creates a new write-only instance of CounterEventEmitter, bound to a specific deployed contract.
func NewCounterEventEmitterTransactor(address common.Address, transactor bind.ContractTransactor) (*CounterEventEmitterTransactor, error) {
	contract, err := bindCounterEventEmitter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CounterEventEmitterTransactor{contract: contract}, nil
}

// NewCounterEventEmitterFilterer creates a new log filterer instance of CounterEventEmitter, bound to a specific deployed contract.
func NewCounterEventEmitterFilterer(address common.Address, filterer bind.ContractFilterer) (*CounterEventEmitterFilterer, error) {
	contract, err := bindCounterEventEmitter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CounterEventEmitterFilterer{contract: contract}, nil
}

// bindCounterEventEmitter binds a generic wrapper to an already deployed contract.
func bindCounterEventEmitter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CounterEventEmitterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CounterEventEmitter *CounterEventEmitterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CounterEventEmitter.Contract.CounterEventEmitterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CounterEventEmitter *CounterEventEmitterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CounterEventEmitter.Contract.CounterEventEmitterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CounterEventEmitter *CounterEventEmitterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CounterEventEmitter.Contract.CounterEventEmitterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CounterEventEmitter *CounterEventEmitterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CounterEventEmitter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CounterEventEmitter *CounterEventEmitterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CounterEventEmitter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CounterEventEmitter *CounterEventEmitterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CounterEventEmitter.Contract.contract.Transact(opts, method, params...)
}

// GetCount is a free data retrieval call binding the contract method 0xa87d942c.
//
// Solidity: function getCount() view returns(int256)
func (_CounterEventEmitter *CounterEventEmitterCaller) GetCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CounterEventEmitter.contract.Call(opts, &out, "getCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCount is a free data retrieval call binding the contract method 0xa87d942c.
//
// Solidity: function getCount() view returns(int256)
func (_CounterEventEmitter *CounterEventEmitterSession) GetCount() (*big.Int, error) {
	return _CounterEventEmitter.Contract.GetCount(&_CounterEventEmitter.CallOpts)
}

// GetCount is a free data retrieval call binding the contract method 0xa87d942c.
//
// Solidity: function getCount() view returns(int256)
func (_CounterEventEmitter *CounterEventEmitterCallerSession) GetCount() (*big.Int, error) {
	return _CounterEventEmitter.Contract.GetCount(&_CounterEventEmitter.CallOpts)
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_CounterEventEmitter *CounterEventEmitterTransactor) Increment(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CounterEventEmitter.contract.Transact(opts, "increment")
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_CounterEventEmitter *CounterEventEmitterSession) Increment() (*types.Transaction, error) {
	return _CounterEventEmitter.Contract.Increment(&_CounterEventEmitter.TransactOpts)
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_CounterEventEmitter *CounterEventEmitterTransactorSession) Increment() (*types.Transaction, error) {
	return _CounterEventEmitter.Contract.Increment(&_CounterEventEmitter.TransactOpts)
}

// CounterEventEmitterCountIterator is returned from FilterCount and is used to iterate over the raw logs and unpacked data for Count events raised by the CounterEventEmitter contract.
type CounterEventEmitterCountIterator struct {
	Event *CounterEventEmitterCount // Event containing the contract specifics and raw log

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
func (it *CounterEventEmitterCountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CounterEventEmitterCount)
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
		it.Event = new(CounterEventEmitterCount)
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
func (it *CounterEventEmitterCountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CounterEventEmitterCountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CounterEventEmitterCount represents a Count event raised by the CounterEventEmitter contract.
type CounterEventEmitterCount struct {
	Count *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterCount is a free log retrieval operation binding the contract event 0x47a90661b520aaf860e17d861200923904ff3dc45e8640c690c905b71719464e.
//
// Solidity: event Count(int256 count)
func (_CounterEventEmitter *CounterEventEmitterFilterer) FilterCount(opts *bind.FilterOpts) (*CounterEventEmitterCountIterator, error) {

	logs, sub, err := _CounterEventEmitter.contract.FilterLogs(opts, "Count")
	if err != nil {
		return nil, err
	}
	return &CounterEventEmitterCountIterator{contract: _CounterEventEmitter.contract, event: "Count", logs: logs, sub: sub}, nil
}

// WatchCount is a free log subscription operation binding the contract event 0x47a90661b520aaf860e17d861200923904ff3dc45e8640c690c905b71719464e.
//
// Solidity: event Count(int256 count)
func (_CounterEventEmitter *CounterEventEmitterFilterer) WatchCount(opts *bind.WatchOpts, sink chan<- *CounterEventEmitterCount) (event.Subscription, error) {

	logs, sub, err := _CounterEventEmitter.contract.WatchLogs(opts, "Count")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CounterEventEmitterCount)
				if err := _CounterEventEmitter.contract.UnpackLog(event, "Count", log); err != nil {
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

// ParseCount is a log parse operation binding the contract event 0x47a90661b520aaf860e17d861200923904ff3dc45e8640c690c905b71719464e.
//
// Solidity: event Count(int256 count)
func (_CounterEventEmitter *CounterEventEmitterFilterer) ParseCount(log types.Log) (*CounterEventEmitterCount, error) {
	event := new(CounterEventEmitterCount)
	if err := _CounterEventEmitter.contract.UnpackLog(event, "Count", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
