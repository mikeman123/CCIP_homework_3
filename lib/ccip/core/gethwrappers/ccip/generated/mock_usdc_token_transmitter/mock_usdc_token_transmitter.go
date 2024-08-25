// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mock_usdc_token_transmitter

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated"
)

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

var MockE2EUSDCTransmitterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_version\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_localDomain\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"MessageSent\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextAvailableNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"receiveMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_shouldSucceed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destinationDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recipient\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"messageBody\",\"type\":\"bytes\"}],\"name\":\"sendMessage\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destinationDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recipient\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"destinationCaller\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"messageBody\",\"type\":\"bytes\"}],\"name\":\"sendMessageWithCaller\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"shouldSucceed\",\"type\":\"bool\"}],\"name\":\"setShouldSucceed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60c060405234801561001057600080fd5b506040516107f83803806107f883398101604081905261002f91610069565b63ffffffff9182166080521660a0526000805460ff1916600117905561009c565b805163ffffffff8116811461006457600080fd5b919050565b6000806040838503121561007c57600080fd5b61008583610050565b915061009360208401610050565b90509250929050565b60805160a05161072a6100ce60003960008181610149015261038801526000818160c00152610367015261072a6000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c80638371744e1161005b5780638371744e1461012e5780638d3638f4146101475780639e31ddb61461016d578063f7259a75146101ae57600080fd5b80630ba469bc1461008d57806354fd4d50146100be57806357ecfd28146100f55780637a64293514610121575b600080fd5b6100a061009b36600461047a565b6101c1565b60405167ffffffffffffffff90911681526020015b60405180910390f35b7f00000000000000000000000000000000000000000000000000000000000000005b60405163ffffffff90911681526020016100b5565b6101116101033660046104d4565b60005460ff16949350505050565b60405190151581526020016100b5565b6000546101119060ff1681565b6000546100a090610100900467ffffffffffffffff1681565b7f00000000000000000000000000000000000000000000000000000000000000006100e0565b6101ac61017b366004610534565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016911515919091179055565b005b6100a06101bc36600461055d565b6101ea565b600080806101cd61029a565b9050336101df88888584868b8b6102fc565b509695505050505050565b60008361027e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f44657374696e6174696f6e2063616c6c6572206d757374206265206e6f6e7a6560448201527f726f00000000000000000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b600061028861029a565b9050336101df88888884868a8a6102fc565b60008054610100900467ffffffffffffffff166102b88160016105c5565b6000805467ffffffffffffffff92909216610100027fffffffffffffffffffffffffffffffffffffffffffffff0000000000000000ff909216919091179055919050565b85610363576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f526563697069656e74206d757374206265206e6f6e7a65726f000000000000006044820152606401610275565b60007f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000008986888b8b89896040516020016103c699989796959493929190610614565b60405160208183030381529060405290507f8c5261668696ce22758910d05bab8f186d6eb247ceac2af2e82c7dc17669b0368160405161040691906106b1565b60405180910390a15050505050505050565b803563ffffffff8116811461042c57600080fd5b919050565b60008083601f84011261044357600080fd5b50813567ffffffffffffffff81111561045b57600080fd5b60208301915083602082850101111561047357600080fd5b9250929050565b6000806000806060858703121561049057600080fd5b61049985610418565b935060208501359250604085013567ffffffffffffffff8111156104bc57600080fd5b6104c887828801610431565b95989497509550505050565b600080600080604085870312156104ea57600080fd5b843567ffffffffffffffff8082111561050257600080fd5b61050e88838901610431565b9096509450602087013591508082111561052757600080fd5b506104c887828801610431565b60006020828403121561054657600080fd5b8135801515811461055657600080fd5b9392505050565b60008060008060006080868803121561057557600080fd5b61057e86610418565b94506020860135935060408601359250606086013567ffffffffffffffff8111156105a857600080fd5b6105b488828901610431565b969995985093965092949392505050565b67ffffffffffffffff81811683821601908082111561060d577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b5092915050565b60007fffffffff00000000000000000000000000000000000000000000000000000000808c60e01b168352808b60e01b166004840152808a60e01b166008840152507fffffffffffffffff0000000000000000000000000000000000000000000000008860c01b16600c83015286601483015285603483015284605483015282846074840137506000910160740190815298975050505050505050565b600060208083528351808285015260005b818110156106de578581018301518582016040015282016106c2565b5060006040828601015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f830116850101925050509291505056fea164736f6c6343000813000a",
}

var MockE2EUSDCTransmitterABI = MockE2EUSDCTransmitterMetaData.ABI

var MockE2EUSDCTransmitterBin = MockE2EUSDCTransmitterMetaData.Bin

func DeployMockE2EUSDCTransmitter(auth *bind.TransactOpts, backend bind.ContractBackend, _version uint32, _localDomain uint32) (common.Address, *types.Transaction, *MockE2EUSDCTransmitter, error) {
	parsed, err := MockE2EUSDCTransmitterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MockE2EUSDCTransmitterBin), backend, _version, _localDomain)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MockE2EUSDCTransmitter{address: address, abi: *parsed, MockE2EUSDCTransmitterCaller: MockE2EUSDCTransmitterCaller{contract: contract}, MockE2EUSDCTransmitterTransactor: MockE2EUSDCTransmitterTransactor{contract: contract}, MockE2EUSDCTransmitterFilterer: MockE2EUSDCTransmitterFilterer{contract: contract}}, nil
}

type MockE2EUSDCTransmitter struct {
	address common.Address
	abi     abi.ABI
	MockE2EUSDCTransmitterCaller
	MockE2EUSDCTransmitterTransactor
	MockE2EUSDCTransmitterFilterer
}

type MockE2EUSDCTransmitterCaller struct {
	contract *bind.BoundContract
}

type MockE2EUSDCTransmitterTransactor struct {
	contract *bind.BoundContract
}

type MockE2EUSDCTransmitterFilterer struct {
	contract *bind.BoundContract
}

type MockE2EUSDCTransmitterSession struct {
	Contract     *MockE2EUSDCTransmitter
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type MockE2EUSDCTransmitterCallerSession struct {
	Contract *MockE2EUSDCTransmitterCaller
	CallOpts bind.CallOpts
}

type MockE2EUSDCTransmitterTransactorSession struct {
	Contract     *MockE2EUSDCTransmitterTransactor
	TransactOpts bind.TransactOpts
}

type MockE2EUSDCTransmitterRaw struct {
	Contract *MockE2EUSDCTransmitter
}

type MockE2EUSDCTransmitterCallerRaw struct {
	Contract *MockE2EUSDCTransmitterCaller
}

type MockE2EUSDCTransmitterTransactorRaw struct {
	Contract *MockE2EUSDCTransmitterTransactor
}

func NewMockE2EUSDCTransmitter(address common.Address, backend bind.ContractBackend) (*MockE2EUSDCTransmitter, error) {
	abi, err := abi.JSON(strings.NewReader(MockE2EUSDCTransmitterABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindMockE2EUSDCTransmitter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockE2EUSDCTransmitter{address: address, abi: abi, MockE2EUSDCTransmitterCaller: MockE2EUSDCTransmitterCaller{contract: contract}, MockE2EUSDCTransmitterTransactor: MockE2EUSDCTransmitterTransactor{contract: contract}, MockE2EUSDCTransmitterFilterer: MockE2EUSDCTransmitterFilterer{contract: contract}}, nil
}

func NewMockE2EUSDCTransmitterCaller(address common.Address, caller bind.ContractCaller) (*MockE2EUSDCTransmitterCaller, error) {
	contract, err := bindMockE2EUSDCTransmitter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockE2EUSDCTransmitterCaller{contract: contract}, nil
}

func NewMockE2EUSDCTransmitterTransactor(address common.Address, transactor bind.ContractTransactor) (*MockE2EUSDCTransmitterTransactor, error) {
	contract, err := bindMockE2EUSDCTransmitter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockE2EUSDCTransmitterTransactor{contract: contract}, nil
}

func NewMockE2EUSDCTransmitterFilterer(address common.Address, filterer bind.ContractFilterer) (*MockE2EUSDCTransmitterFilterer, error) {
	contract, err := bindMockE2EUSDCTransmitter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockE2EUSDCTransmitterFilterer{contract: contract}, nil
}

func bindMockE2EUSDCTransmitter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MockE2EUSDCTransmitterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_MockE2EUSDCTransmitter *MockE2EUSDCTransmitterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockE2EUSDCTransmitter.Contract.MockE2EUSDCTransmitterCaller.contract.Call(opts, result, method, params...)
}

func (_MockE2EUSDCTransmitter *MockE2EUSDCTransmitterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockE2EUSDCTransmitter.Contract.MockE2EUSDCTransmitterTransactor.contract.Transfer(opts)
}

func (_MockE2EUSDCTransmitter *MockE2EUSDCTransmitterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockE2EUSDCTransmitter.Contract.MockE2EUSDCTransmitterTransactor.contract.Transact(opts, method, params...)
}

func (_MockE2EUSDCTransmitter *MockE2EUSDCTransmitterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockE2EUSDCTransmitter.Contract.contract.Call(opts, result, method, params...)
}

func (_MockE2EUSDCTransmitter *MockE2EUSDCTransmitterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockE2EUSDCTransmitter.Contract.contract.Transfer(opts)
}

func (_MockE2EUSDCTransmitter *MockE2EUSDCTransmitterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockE2EUSDCTransmitter.Contract.contract.Transact(opts, method, params...)
}

func (_MockE2EUSDCTransmitter *MockE2EUSDCTransmitterCaller) LocalDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _MockE2EUSDCTransmitter.contract.Call(opts, &out, "localDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

func (_MockE2EUSDCTransmitter *MockE2EUSDCTransmitterSession) LocalDomain() (uint32, error) {
	return _MockE2EUSDCTransmitter.Contract.LocalDomain(&_MockE2EUSDCTransmitter.CallOpts)
}

func (_MockE2EUSDCTransmitter *MockE2EUSDCTransmitterCallerSession) LocalDomain() (uint32, error) {
	return _MockE2EUSDCTransmitter.Contract.LocalDomain(&_MockE2EUSDCTransmitter.CallOpts)
}

func (_MockE2EUSDCTransmitter *MockE2EUSDCTransmitterCaller) NextAvailableNonce(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _MockE2EUSDCTransmitter.contract.Call(opts, &out, "nextAvailableNonce")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_MockE2EUSDCTransmitter *MockE2EUSDCTransmitterSession) NextAvailableNonce() (uint64, error) {
	return _MockE2EUSDCTransmitter.Contract.NextAvailableNonce(&_MockE2EUSDCTransmitter.CallOpts)
}

func (_MockE2EUSDCTransmitter *MockE2EUSDCTransmitterCallerSession) NextAvailableNonce() (uint64, error) {
	return _MockE2EUSDCTransmitter.Contract.NextAvailableNonce(&_MockE2EUSDCTransmitter.CallOpts)
}

func (_MockE2EUSDCTransmitter *MockE2EUSDCTransmitterCaller) ReceiveMessage(opts *bind.CallOpts, arg0 []byte, arg1 []byte) (bool, error) {
	var out []interface{}
	err := _MockE2EUSDCTransmitter.contract.Call(opts, &out, "receiveMessage", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_MockE2EUSDCTransmitter *MockE2EUSDCTransmitterSession) ReceiveMessage(arg0 []byte, arg1 []byte) (bool, error) {
	return _MockE2EUSDCTransmitter.Contract.ReceiveMessage(&_MockE2EUSDCTransmitter.CallOpts, arg0, arg1)
}

func (_MockE2EUSDCTransmitter *MockE2EUSDCTransmitterCallerSession) ReceiveMessage(arg0 []byte, arg1 []byte) (bool, error) {
	return _MockE2EUSDCTransmitter.Contract.ReceiveMessage(&_MockE2EUSDCTransmitter.CallOpts, arg0, arg1)
}

func (_MockE2EUSDCTransmitter *MockE2EUSDCTransmitterCaller) SShouldSucceed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MockE2EUSDCTransmitter.contract.Call(opts, &out, "s_shouldSucceed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_MockE2EUSDCTransmitter *MockE2EUSDCTransmitterSession) SShouldSucceed() (bool, error) {
	return _MockE2EUSDCTransmitter.Contract.SShouldSucceed(&_MockE2EUSDCTransmitter.CallOpts)
}

func (_MockE2EUSDCTransmitter *MockE2EUSDCTransmitterCallerSession) SShouldSucceed() (bool, error) {
	return _MockE2EUSDCTransmitter.Contract.SShouldSucceed(&_MockE2EUSDCTransmitter.CallOpts)
}

func (_MockE2EUSDCTransmitter *MockE2EUSDCTransmitterCaller) Version(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _MockE2EUSDCTransmitter.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

func (_MockE2EUSDCTransmitter *MockE2EUSDCTransmitterSession) Version() (uint32, error) {
	return _MockE2EUSDCTransmitter.Contract.Version(&_MockE2EUSDCTransmitter.CallOpts)
}

func (_MockE2EUSDCTransmitter *MockE2EUSDCTransmitterCallerSession) Version() (uint32, error) {
	return _MockE2EUSDCTransmitter.Contract.Version(&_MockE2EUSDCTransmitter.CallOpts)
}

func (_MockE2EUSDCTransmitter *MockE2EUSDCTransmitterTransactor) SendMessage(opts *bind.TransactOpts, destinationDomain uint32, recipient [32]byte, messageBody []byte) (*types.Transaction, error) {
	return _MockE2EUSDCTransmitter.contract.Transact(opts, "sendMessage", destinationDomain, recipient, messageBody)
}

func (_MockE2EUSDCTransmitter *MockE2EUSDCTransmitterSession) SendMessage(destinationDomain uint32, recipient [32]byte, messageBody []byte) (*types.Transaction, error) {
	return _MockE2EUSDCTransmitter.Contract.SendMessage(&_MockE2EUSDCTransmitter.TransactOpts, destinationDomain, recipient, messageBody)
}

func (_MockE2EUSDCTransmitter *MockE2EUSDCTransmitterTransactorSession) SendMessage(destinationDomain uint32, recipient [32]byte, messageBody []byte) (*types.Transaction, error) {
	return _MockE2EUSDCTransmitter.Contract.SendMessage(&_MockE2EUSDCTransmitter.TransactOpts, destinationDomain, recipient, messageBody)
}

func (_MockE2EUSDCTransmitter *MockE2EUSDCTransmitterTransactor) SendMessageWithCaller(opts *bind.TransactOpts, destinationDomain uint32, recipient [32]byte, destinationCaller [32]byte, messageBody []byte) (*types.Transaction, error) {
	return _MockE2EUSDCTransmitter.contract.Transact(opts, "sendMessageWithCaller", destinationDomain, recipient, destinationCaller, messageBody)
}

func (_MockE2EUSDCTransmitter *MockE2EUSDCTransmitterSession) SendMessageWithCaller(destinationDomain uint32, recipient [32]byte, destinationCaller [32]byte, messageBody []byte) (*types.Transaction, error) {
	return _MockE2EUSDCTransmitter.Contract.SendMessageWithCaller(&_MockE2EUSDCTransmitter.TransactOpts, destinationDomain, recipient, destinationCaller, messageBody)
}

func (_MockE2EUSDCTransmitter *MockE2EUSDCTransmitterTransactorSession) SendMessageWithCaller(destinationDomain uint32, recipient [32]byte, destinationCaller [32]byte, messageBody []byte) (*types.Transaction, error) {
	return _MockE2EUSDCTransmitter.Contract.SendMessageWithCaller(&_MockE2EUSDCTransmitter.TransactOpts, destinationDomain, recipient, destinationCaller, messageBody)
}

func (_MockE2EUSDCTransmitter *MockE2EUSDCTransmitterTransactor) SetShouldSucceed(opts *bind.TransactOpts, shouldSucceed bool) (*types.Transaction, error) {
	return _MockE2EUSDCTransmitter.contract.Transact(opts, "setShouldSucceed", shouldSucceed)
}

func (_MockE2EUSDCTransmitter *MockE2EUSDCTransmitterSession) SetShouldSucceed(shouldSucceed bool) (*types.Transaction, error) {
	return _MockE2EUSDCTransmitter.Contract.SetShouldSucceed(&_MockE2EUSDCTransmitter.TransactOpts, shouldSucceed)
}

func (_MockE2EUSDCTransmitter *MockE2EUSDCTransmitterTransactorSession) SetShouldSucceed(shouldSucceed bool) (*types.Transaction, error) {
	return _MockE2EUSDCTransmitter.Contract.SetShouldSucceed(&_MockE2EUSDCTransmitter.TransactOpts, shouldSucceed)
}

type MockE2EUSDCTransmitterMessageSentIterator struct {
	Event *MockE2EUSDCTransmitterMessageSent

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockE2EUSDCTransmitterMessageSentIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockE2EUSDCTransmitterMessageSent)
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

	select {
	case log := <-it.logs:
		it.Event = new(MockE2EUSDCTransmitterMessageSent)
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

func (it *MockE2EUSDCTransmitterMessageSentIterator) Error() error {
	return it.fail
}

func (it *MockE2EUSDCTransmitterMessageSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockE2EUSDCTransmitterMessageSent struct {
	Message []byte
	Raw     types.Log
}

func (_MockE2EUSDCTransmitter *MockE2EUSDCTransmitterFilterer) FilterMessageSent(opts *bind.FilterOpts) (*MockE2EUSDCTransmitterMessageSentIterator, error) {

	logs, sub, err := _MockE2EUSDCTransmitter.contract.FilterLogs(opts, "MessageSent")
	if err != nil {
		return nil, err
	}
	return &MockE2EUSDCTransmitterMessageSentIterator{contract: _MockE2EUSDCTransmitter.contract, event: "MessageSent", logs: logs, sub: sub}, nil
}

func (_MockE2EUSDCTransmitter *MockE2EUSDCTransmitterFilterer) WatchMessageSent(opts *bind.WatchOpts, sink chan<- *MockE2EUSDCTransmitterMessageSent) (event.Subscription, error) {

	logs, sub, err := _MockE2EUSDCTransmitter.contract.WatchLogs(opts, "MessageSent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockE2EUSDCTransmitterMessageSent)
				if err := _MockE2EUSDCTransmitter.contract.UnpackLog(event, "MessageSent", log); err != nil {
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

func (_MockE2EUSDCTransmitter *MockE2EUSDCTransmitterFilterer) ParseMessageSent(log types.Log) (*MockE2EUSDCTransmitterMessageSent, error) {
	event := new(MockE2EUSDCTransmitterMessageSent)
	if err := _MockE2EUSDCTransmitter.contract.UnpackLog(event, "MessageSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_MockE2EUSDCTransmitter *MockE2EUSDCTransmitter) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _MockE2EUSDCTransmitter.abi.Events["MessageSent"].ID:
		return _MockE2EUSDCTransmitter.ParseMessageSent(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (MockE2EUSDCTransmitterMessageSent) Topic() common.Hash {
	return common.HexToHash("0x8c5261668696ce22758910d05bab8f186d6eb247ceac2af2e82c7dc17669b036")
}

func (_MockE2EUSDCTransmitter *MockE2EUSDCTransmitter) Address() common.Address {
	return _MockE2EUSDCTransmitter.address
}

type MockE2EUSDCTransmitterInterface interface {
	LocalDomain(opts *bind.CallOpts) (uint32, error)

	NextAvailableNonce(opts *bind.CallOpts) (uint64, error)

	ReceiveMessage(opts *bind.CallOpts, arg0 []byte, arg1 []byte) (bool, error)

	SShouldSucceed(opts *bind.CallOpts) (bool, error)

	Version(opts *bind.CallOpts) (uint32, error)

	SendMessage(opts *bind.TransactOpts, destinationDomain uint32, recipient [32]byte, messageBody []byte) (*types.Transaction, error)

	SendMessageWithCaller(opts *bind.TransactOpts, destinationDomain uint32, recipient [32]byte, destinationCaller [32]byte, messageBody []byte) (*types.Transaction, error)

	SetShouldSucceed(opts *bind.TransactOpts, shouldSucceed bool) (*types.Transaction, error)

	FilterMessageSent(opts *bind.FilterOpts) (*MockE2EUSDCTransmitterMessageSentIterator, error)

	WatchMessageSent(opts *bind.WatchOpts, sink chan<- *MockE2EUSDCTransmitterMessageSent) (event.Subscription, error)

	ParseMessageSent(log types.Log) (*MockE2EUSDCTransmitterMessageSent, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
