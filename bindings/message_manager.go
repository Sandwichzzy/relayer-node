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
	_ = abi.ConvertType
)

// MessageManagerMetaData contains all meta data concerning the MessageManager contract.
var MessageManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimMessage\",\"inputs\":[{\"name\":\"sourceChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourceTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_fee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimMessageStatus\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_poolManagerAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"nextMessageNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"poolManagerAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sendMessage\",\"inputs\":[{\"name\":\"sourceChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourceTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_fee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sentMessageStatus\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MessageClaimed\",\"inputs\":[{\"name\":\"sourceChainId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"destChainId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"sourceTokenAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"destTokenAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"_messageHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"_nonce\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MessageSent\",\"inputs\":[{\"name\":\"sourceChainId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"destChainId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"sourceTokenAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"destTokenAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"_from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_fee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_nonce\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_messageHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MessageAlreadySent\",\"inputs\":[{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x6080806040523460ce5760017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00555f5160206109fd5f395f51905f525460ff8160401c1660bf576002600160401b03196001600160401b03821601606d575b60405161092a90816100d38239f35b6001600160401b0319166001600160401b039081175f5160206109fd5f395f51905f525581527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a15f80605e565b63f92ee8a960e01b5f5260045ffd5b5f80fdfe6080806040526004361015610012575f80fd5b5f3560e01c90816336e3d8d71461065a57508063485cc955146104e1578063715018a61461047a5780637ae1e828146102ad5780637f8d98c31461027e5780638da5cb5b1461024a578063a8f0849e146100fd578063b837dbe9146100e1578063e6a13f30146100b95763f2fde38b1461008a575f80fd5b346100b55760203660031901126100b5576100b36100a6610685565b6100ae610856565b6107e5565b005b5f80fd5b346100b5575f3660031901126100b5576001546040516001600160a01b039091168152602090f35b346100b5575f3660031901126100b55760205f54604051908152f35b346100b5576101003660031901126100b55760043560243561011d61069b565b916101266106b1565b9261012f6106c7565b926101386106dd565b9360c4359460e4359561015660018060a01b036001541633146106f3565b6001600160a01b03821696871561023b5786866101998761018b8d86888b9a5f549b8c94604051998a9860208a019c8d61075d565b03601f1981018352826107af565b519020985f198414610227577fb7f33a96ca66e7a604eaea9cd873498ec3775d6b36cf3cec44c8df6b467b2e469660e096600186015f558b5f52600260205260405f20600160ff19825416179055604051998a5260208a015260018060a01b0316604089015260018060a01b03166060880152608087015260a086015260c085015260018060a01b031692a4005b634e487b7160e01b5f52601160045260245ffd5b6342bcdf7f60e11b5f5260045ffd5b346100b5575f3660031901126100b5575f5160206108b55f395f51905f52546040516001600160a01b039091168152602090f35b346100b55760203660031901126100b5576004355f526002602052602060ff60405f2054166040519015158152f35b346100b5576101203660031901126100b5576004356024356102cd61069b565b906102d66106b1565b926102df6106c7565b6102e76106dd565b610104359161030160018060a01b036001541633146106f3565b60027f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00541461046b576103738761018b8588888b9760027f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055604051978896602088019a60e4359560c435958d61075d565b51902094855f52600360205260ff60405f205416610416577f410ce732ac401991bca377a8f57d056d2ae1c15d9819a9ce066f1b8e2a1b207b9460a094875f52600360205260405f20600160ff198254161790556040519485526020850152600180861b03166040840152600180851b031660608301526080820152a260017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055005b60405162461bcd60e51b815260206004820152602760248201527f4d6573736167654d616e616765723a206d65737361676520616c72656164792060448201526618db185a5b595960ca1b6064820152608490fd5b633ee5aeb560e01b5f5260045ffd5b346100b5575f3660031901126100b557610492610856565b5f5160206108b55f395f51905f5280546001600160a01b031981169091555f906001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a3005b346100b55760403660031901126100b5576104fa610685565b6024356001600160a01b038116908190036100b5575f5160206108d55f395f51905f525460ff8160401c16159267ffffffffffffffff821680159081610652575b6001149081610648575b15908161063f575b506106305767ffffffffffffffff1982166001175f5160206108d55f395f51905f525561058b9184610604575b50610583610889565b6100ae610889565b6bffffffffffffffffffffffff60a01b600154161760015560015f556105ad57005b68ff0000000000000000195f5160206108d55f395f51905f5254165f5160206108d55f395f51905f52557fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602060405160018152a1005b68ffffffffffffffffff191668010000000000000001175f5160206108d55f395f51905f52558461057a565b63f92ee8a960e01b5f5260045ffd5b9050158561054d565b303b159150610545565b85915061053b565b346100b55760203660031901126100b5576020906004355f526003825260ff60405f20541615158152f35b600435906001600160a01b03821682036100b557565b604435906001600160a01b03821682036100b557565b606435906001600160a01b03821682036100b557565b608435906001600160a01b03821682036100b557565b60a435906001600160a01b03821682036100b557565b156106fa57565b60405162461bcd60e51b815260206004820152603560248201527f4d6573736167654d616e616765723a206f6e6c7920746f6b656e206272696467604482015274652063616e20646f2074686973206f70657261746560581b6064820152608490fd5b90815260208101919091526001600160a01b03918216604082015291811660608301529182166080820152911660a082015260c081019190915260e08101919091526101008101919091526101200190565b90601f8019910116810190811067ffffffffffffffff8211176107d157604052565b634e487b7160e01b5f52604160045260245ffd5b6001600160a01b03168015610843575f5160206108b55f395f51905f5280546001600160a01b0319811683179091556001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e05f80a3565b631e4fbdf760e01b5f525f60045260245ffd5b5f5160206108b55f395f51905f52546001600160a01b0316330361087657565b63118cdaa760e01b5f523360045260245ffd5b60ff5f5160206108d55f395f51905f525460401c16156108a557565b631afcd79f60e31b5f5260045ffdfe9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00a2646970667358221220ce4828925ee7e11be03265639bf2e7d36ecc342c10cc97dcb80dec738bff1c6d64736f6c634300081e0033f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00",
}

// MessageManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use MessageManagerMetaData.ABI instead.
var MessageManagerABI = MessageManagerMetaData.ABI

// MessageManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MessageManagerMetaData.Bin instead.
var MessageManagerBin = MessageManagerMetaData.Bin

// DeployMessageManager deploys a new Ethereum contract, binding an instance of MessageManager to it.
func DeployMessageManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MessageManager, error) {
	parsed, err := MessageManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MessageManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MessageManager{MessageManagerCaller: MessageManagerCaller{contract: contract}, MessageManagerTransactor: MessageManagerTransactor{contract: contract}, MessageManagerFilterer: MessageManagerFilterer{contract: contract}}, nil
}

// MessageManager is an auto generated Go binding around an Ethereum contract.
type MessageManager struct {
	MessageManagerCaller     // Read-only binding to the contract
	MessageManagerTransactor // Write-only binding to the contract
	MessageManagerFilterer   // Log filterer for contract events
}

// MessageManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type MessageManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MessageManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MessageManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MessageManagerSession struct {
	Contract     *MessageManager   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MessageManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MessageManagerCallerSession struct {
	Contract *MessageManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// MessageManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MessageManagerTransactorSession struct {
	Contract     *MessageManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// MessageManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type MessageManagerRaw struct {
	Contract *MessageManager // Generic contract binding to access the raw methods on
}

// MessageManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MessageManagerCallerRaw struct {
	Contract *MessageManagerCaller // Generic read-only contract binding to access the raw methods on
}

// MessageManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MessageManagerTransactorRaw struct {
	Contract *MessageManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMessageManager creates a new instance of MessageManager, bound to a specific deployed contract.
func NewMessageManager(address common.Address, backend bind.ContractBackend) (*MessageManager, error) {
	contract, err := bindMessageManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MessageManager{MessageManagerCaller: MessageManagerCaller{contract: contract}, MessageManagerTransactor: MessageManagerTransactor{contract: contract}, MessageManagerFilterer: MessageManagerFilterer{contract: contract}}, nil
}

// NewMessageManagerCaller creates a new read-only instance of MessageManager, bound to a specific deployed contract.
func NewMessageManagerCaller(address common.Address, caller bind.ContractCaller) (*MessageManagerCaller, error) {
	contract, err := bindMessageManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MessageManagerCaller{contract: contract}, nil
}

// NewMessageManagerTransactor creates a new write-only instance of MessageManager, bound to a specific deployed contract.
func NewMessageManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*MessageManagerTransactor, error) {
	contract, err := bindMessageManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MessageManagerTransactor{contract: contract}, nil
}

// NewMessageManagerFilterer creates a new log filterer instance of MessageManager, bound to a specific deployed contract.
func NewMessageManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*MessageManagerFilterer, error) {
	contract, err := bindMessageManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MessageManagerFilterer{contract: contract}, nil
}

// bindMessageManager binds a generic wrapper to an already deployed contract.
func bindMessageManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MessageManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MessageManager *MessageManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MessageManager.Contract.MessageManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MessageManager *MessageManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageManager.Contract.MessageManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MessageManager *MessageManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageManager.Contract.MessageManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MessageManager *MessageManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MessageManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MessageManager *MessageManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MessageManager *MessageManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageManager.Contract.contract.Transact(opts, method, params...)
}

// ClaimMessageStatus is a free data retrieval call binding the contract method 0x36e3d8d7.
//
// Solidity: function claimMessageStatus(bytes32 ) view returns(bool)
func (_MessageManager *MessageManagerCaller) ClaimMessageStatus(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _MessageManager.contract.Call(opts, &out, "claimMessageStatus", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ClaimMessageStatus is a free data retrieval call binding the contract method 0x36e3d8d7.
//
// Solidity: function claimMessageStatus(bytes32 ) view returns(bool)
func (_MessageManager *MessageManagerSession) ClaimMessageStatus(arg0 [32]byte) (bool, error) {
	return _MessageManager.Contract.ClaimMessageStatus(&_MessageManager.CallOpts, arg0)
}

// ClaimMessageStatus is a free data retrieval call binding the contract method 0x36e3d8d7.
//
// Solidity: function claimMessageStatus(bytes32 ) view returns(bool)
func (_MessageManager *MessageManagerCallerSession) ClaimMessageStatus(arg0 [32]byte) (bool, error) {
	return _MessageManager.Contract.ClaimMessageStatus(&_MessageManager.CallOpts, arg0)
}

// NextMessageNumber is a free data retrieval call binding the contract method 0xb837dbe9.
//
// Solidity: function nextMessageNumber() view returns(uint256)
func (_MessageManager *MessageManagerCaller) NextMessageNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MessageManager.contract.Call(opts, &out, "nextMessageNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextMessageNumber is a free data retrieval call binding the contract method 0xb837dbe9.
//
// Solidity: function nextMessageNumber() view returns(uint256)
func (_MessageManager *MessageManagerSession) NextMessageNumber() (*big.Int, error) {
	return _MessageManager.Contract.NextMessageNumber(&_MessageManager.CallOpts)
}

// NextMessageNumber is a free data retrieval call binding the contract method 0xb837dbe9.
//
// Solidity: function nextMessageNumber() view returns(uint256)
func (_MessageManager *MessageManagerCallerSession) NextMessageNumber() (*big.Int, error) {
	return _MessageManager.Contract.NextMessageNumber(&_MessageManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MessageManager *MessageManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MessageManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MessageManager *MessageManagerSession) Owner() (common.Address, error) {
	return _MessageManager.Contract.Owner(&_MessageManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MessageManager *MessageManagerCallerSession) Owner() (common.Address, error) {
	return _MessageManager.Contract.Owner(&_MessageManager.CallOpts)
}

// PoolManagerAddress is a free data retrieval call binding the contract method 0xe6a13f30.
//
// Solidity: function poolManagerAddress() view returns(address)
func (_MessageManager *MessageManagerCaller) PoolManagerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MessageManager.contract.Call(opts, &out, "poolManagerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolManagerAddress is a free data retrieval call binding the contract method 0xe6a13f30.
//
// Solidity: function poolManagerAddress() view returns(address)
func (_MessageManager *MessageManagerSession) PoolManagerAddress() (common.Address, error) {
	return _MessageManager.Contract.PoolManagerAddress(&_MessageManager.CallOpts)
}

// PoolManagerAddress is a free data retrieval call binding the contract method 0xe6a13f30.
//
// Solidity: function poolManagerAddress() view returns(address)
func (_MessageManager *MessageManagerCallerSession) PoolManagerAddress() (common.Address, error) {
	return _MessageManager.Contract.PoolManagerAddress(&_MessageManager.CallOpts)
}

// SentMessageStatus is a free data retrieval call binding the contract method 0x7f8d98c3.
//
// Solidity: function sentMessageStatus(bytes32 ) view returns(bool)
func (_MessageManager *MessageManagerCaller) SentMessageStatus(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _MessageManager.contract.Call(opts, &out, "sentMessageStatus", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SentMessageStatus is a free data retrieval call binding the contract method 0x7f8d98c3.
//
// Solidity: function sentMessageStatus(bytes32 ) view returns(bool)
func (_MessageManager *MessageManagerSession) SentMessageStatus(arg0 [32]byte) (bool, error) {
	return _MessageManager.Contract.SentMessageStatus(&_MessageManager.CallOpts, arg0)
}

// SentMessageStatus is a free data retrieval call binding the contract method 0x7f8d98c3.
//
// Solidity: function sentMessageStatus(bytes32 ) view returns(bool)
func (_MessageManager *MessageManagerCallerSession) SentMessageStatus(arg0 [32]byte) (bool, error) {
	return _MessageManager.Contract.SentMessageStatus(&_MessageManager.CallOpts, arg0)
}

// ClaimMessage is a paid mutator transaction binding the contract method 0x7ae1e828.
//
// Solidity: function claimMessage(uint256 sourceChainId, uint256 destChainId, address sourceTokenAddress, address destTokenAddress, address _from, address _to, uint256 _value, uint256 _fee, uint256 _nonce) returns()
func (_MessageManager *MessageManagerTransactor) ClaimMessage(opts *bind.TransactOpts, sourceChainId *big.Int, destChainId *big.Int, sourceTokenAddress common.Address, destTokenAddress common.Address, _from common.Address, _to common.Address, _value *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _MessageManager.contract.Transact(opts, "claimMessage", sourceChainId, destChainId, sourceTokenAddress, destTokenAddress, _from, _to, _value, _fee, _nonce)
}

// ClaimMessage is a paid mutator transaction binding the contract method 0x7ae1e828.
//
// Solidity: function claimMessage(uint256 sourceChainId, uint256 destChainId, address sourceTokenAddress, address destTokenAddress, address _from, address _to, uint256 _value, uint256 _fee, uint256 _nonce) returns()
func (_MessageManager *MessageManagerSession) ClaimMessage(sourceChainId *big.Int, destChainId *big.Int, sourceTokenAddress common.Address, destTokenAddress common.Address, _from common.Address, _to common.Address, _value *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _MessageManager.Contract.ClaimMessage(&_MessageManager.TransactOpts, sourceChainId, destChainId, sourceTokenAddress, destTokenAddress, _from, _to, _value, _fee, _nonce)
}

// ClaimMessage is a paid mutator transaction binding the contract method 0x7ae1e828.
//
// Solidity: function claimMessage(uint256 sourceChainId, uint256 destChainId, address sourceTokenAddress, address destTokenAddress, address _from, address _to, uint256 _value, uint256 _fee, uint256 _nonce) returns()
func (_MessageManager *MessageManagerTransactorSession) ClaimMessage(sourceChainId *big.Int, destChainId *big.Int, sourceTokenAddress common.Address, destTokenAddress common.Address, _from common.Address, _to common.Address, _value *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _MessageManager.Contract.ClaimMessage(&_MessageManager.TransactOpts, sourceChainId, destChainId, sourceTokenAddress, destTokenAddress, _from, _to, _value, _fee, _nonce)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address initialOwner, address _poolManagerAddress) returns()
func (_MessageManager *MessageManagerTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address, _poolManagerAddress common.Address) (*types.Transaction, error) {
	return _MessageManager.contract.Transact(opts, "initialize", initialOwner, _poolManagerAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address initialOwner, address _poolManagerAddress) returns()
func (_MessageManager *MessageManagerSession) Initialize(initialOwner common.Address, _poolManagerAddress common.Address) (*types.Transaction, error) {
	return _MessageManager.Contract.Initialize(&_MessageManager.TransactOpts, initialOwner, _poolManagerAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address initialOwner, address _poolManagerAddress) returns()
func (_MessageManager *MessageManagerTransactorSession) Initialize(initialOwner common.Address, _poolManagerAddress common.Address) (*types.Transaction, error) {
	return _MessageManager.Contract.Initialize(&_MessageManager.TransactOpts, initialOwner, _poolManagerAddress)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MessageManager *MessageManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MessageManager *MessageManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _MessageManager.Contract.RenounceOwnership(&_MessageManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MessageManager *MessageManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MessageManager.Contract.RenounceOwnership(&_MessageManager.TransactOpts)
}

// SendMessage is a paid mutator transaction binding the contract method 0xa8f0849e.
//
// Solidity: function sendMessage(uint256 sourceChainId, uint256 destChainId, address sourceTokenAddress, address destTokenAddress, address _from, address _to, uint256 _value, uint256 _fee) returns()
func (_MessageManager *MessageManagerTransactor) SendMessage(opts *bind.TransactOpts, sourceChainId *big.Int, destChainId *big.Int, sourceTokenAddress common.Address, destTokenAddress common.Address, _from common.Address, _to common.Address, _value *big.Int, _fee *big.Int) (*types.Transaction, error) {
	return _MessageManager.contract.Transact(opts, "sendMessage", sourceChainId, destChainId, sourceTokenAddress, destTokenAddress, _from, _to, _value, _fee)
}

// SendMessage is a paid mutator transaction binding the contract method 0xa8f0849e.
//
// Solidity: function sendMessage(uint256 sourceChainId, uint256 destChainId, address sourceTokenAddress, address destTokenAddress, address _from, address _to, uint256 _value, uint256 _fee) returns()
func (_MessageManager *MessageManagerSession) SendMessage(sourceChainId *big.Int, destChainId *big.Int, sourceTokenAddress common.Address, destTokenAddress common.Address, _from common.Address, _to common.Address, _value *big.Int, _fee *big.Int) (*types.Transaction, error) {
	return _MessageManager.Contract.SendMessage(&_MessageManager.TransactOpts, sourceChainId, destChainId, sourceTokenAddress, destTokenAddress, _from, _to, _value, _fee)
}

// SendMessage is a paid mutator transaction binding the contract method 0xa8f0849e.
//
// Solidity: function sendMessage(uint256 sourceChainId, uint256 destChainId, address sourceTokenAddress, address destTokenAddress, address _from, address _to, uint256 _value, uint256 _fee) returns()
func (_MessageManager *MessageManagerTransactorSession) SendMessage(sourceChainId *big.Int, destChainId *big.Int, sourceTokenAddress common.Address, destTokenAddress common.Address, _from common.Address, _to common.Address, _value *big.Int, _fee *big.Int) (*types.Transaction, error) {
	return _MessageManager.Contract.SendMessage(&_MessageManager.TransactOpts, sourceChainId, destChainId, sourceTokenAddress, destTokenAddress, _from, _to, _value, _fee)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MessageManager *MessageManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MessageManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MessageManager *MessageManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MessageManager.Contract.TransferOwnership(&_MessageManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MessageManager *MessageManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MessageManager.Contract.TransferOwnership(&_MessageManager.TransactOpts, newOwner)
}

// MessageManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the MessageManager contract.
type MessageManagerInitializedIterator struct {
	Event *MessageManagerInitialized // Event containing the contract specifics and raw log

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
func (it *MessageManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageManagerInitialized)
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
		it.Event = new(MessageManagerInitialized)
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
func (it *MessageManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageManagerInitialized represents a Initialized event raised by the MessageManager contract.
type MessageManagerInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_MessageManager *MessageManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*MessageManagerInitializedIterator, error) {

	logs, sub, err := _MessageManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &MessageManagerInitializedIterator{contract: _MessageManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_MessageManager *MessageManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *MessageManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _MessageManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageManagerInitialized)
				if err := _MessageManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_MessageManager *MessageManagerFilterer) ParseInitialized(log types.Log) (*MessageManagerInitialized, error) {
	event := new(MessageManagerInitialized)
	if err := _MessageManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageManagerMessageClaimedIterator is returned from FilterMessageClaimed and is used to iterate over the raw logs and unpacked data for MessageClaimed events raised by the MessageManager contract.
type MessageManagerMessageClaimedIterator struct {
	Event *MessageManagerMessageClaimed // Event containing the contract specifics and raw log

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
func (it *MessageManagerMessageClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageManagerMessageClaimed)
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
		it.Event = new(MessageManagerMessageClaimed)
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
func (it *MessageManagerMessageClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageManagerMessageClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageManagerMessageClaimed represents a MessageClaimed event raised by the MessageManager contract.
type MessageManagerMessageClaimed struct {
	SourceChainId      *big.Int
	DestChainId        *big.Int
	SourceTokenAddress common.Address
	DestTokenAddress   common.Address
	MessageHash        [32]byte
	Nonce              *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterMessageClaimed is a free log retrieval operation binding the contract event 0x410ce732ac401991bca377a8f57d056d2ae1c15d9819a9ce066f1b8e2a1b207b.
//
// Solidity: event MessageClaimed(uint256 sourceChainId, uint256 destChainId, address sourceTokenAddress, address destTokenAddress, bytes32 indexed _messageHash, uint256 _nonce)
func (_MessageManager *MessageManagerFilterer) FilterMessageClaimed(opts *bind.FilterOpts, _messageHash [][32]byte) (*MessageManagerMessageClaimedIterator, error) {

	var _messageHashRule []interface{}
	for _, _messageHashItem := range _messageHash {
		_messageHashRule = append(_messageHashRule, _messageHashItem)
	}

	logs, sub, err := _MessageManager.contract.FilterLogs(opts, "MessageClaimed", _messageHashRule)
	if err != nil {
		return nil, err
	}
	return &MessageManagerMessageClaimedIterator{contract: _MessageManager.contract, event: "MessageClaimed", logs: logs, sub: sub}, nil
}

// WatchMessageClaimed is a free log subscription operation binding the contract event 0x410ce732ac401991bca377a8f57d056d2ae1c15d9819a9ce066f1b8e2a1b207b.
//
// Solidity: event MessageClaimed(uint256 sourceChainId, uint256 destChainId, address sourceTokenAddress, address destTokenAddress, bytes32 indexed _messageHash, uint256 _nonce)
func (_MessageManager *MessageManagerFilterer) WatchMessageClaimed(opts *bind.WatchOpts, sink chan<- *MessageManagerMessageClaimed, _messageHash [][32]byte) (event.Subscription, error) {

	var _messageHashRule []interface{}
	for _, _messageHashItem := range _messageHash {
		_messageHashRule = append(_messageHashRule, _messageHashItem)
	}

	logs, sub, err := _MessageManager.contract.WatchLogs(opts, "MessageClaimed", _messageHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageManagerMessageClaimed)
				if err := _MessageManager.contract.UnpackLog(event, "MessageClaimed", log); err != nil {
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

// ParseMessageClaimed is a log parse operation binding the contract event 0x410ce732ac401991bca377a8f57d056d2ae1c15d9819a9ce066f1b8e2a1b207b.
//
// Solidity: event MessageClaimed(uint256 sourceChainId, uint256 destChainId, address sourceTokenAddress, address destTokenAddress, bytes32 indexed _messageHash, uint256 _nonce)
func (_MessageManager *MessageManagerFilterer) ParseMessageClaimed(log types.Log) (*MessageManagerMessageClaimed, error) {
	event := new(MessageManagerMessageClaimed)
	if err := _MessageManager.contract.UnpackLog(event, "MessageClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageManagerMessageSentIterator is returned from FilterMessageSent and is used to iterate over the raw logs and unpacked data for MessageSent events raised by the MessageManager contract.
type MessageManagerMessageSentIterator struct {
	Event *MessageManagerMessageSent // Event containing the contract specifics and raw log

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
func (it *MessageManagerMessageSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageManagerMessageSent)
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
		it.Event = new(MessageManagerMessageSent)
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
func (it *MessageManagerMessageSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageManagerMessageSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageManagerMessageSent represents a MessageSent event raised by the MessageManager contract.
type MessageManagerMessageSent struct {
	SourceChainId      *big.Int
	DestChainId        *big.Int
	SourceTokenAddress common.Address
	DestTokenAddress   common.Address
	From               common.Address
	To                 common.Address
	Fee                *big.Int
	Value              *big.Int
	Nonce              *big.Int
	MessageHash        [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterMessageSent is a free log retrieval operation binding the contract event 0xb7f33a96ca66e7a604eaea9cd873498ec3775d6b36cf3cec44c8df6b467b2e46.
//
// Solidity: event MessageSent(uint256 sourceChainId, uint256 destChainId, address sourceTokenAddress, address destTokenAddress, address indexed _from, address indexed _to, uint256 _fee, uint256 _value, uint256 _nonce, bytes32 indexed _messageHash)
func (_MessageManager *MessageManagerFilterer) FilterMessageSent(opts *bind.FilterOpts, _from []common.Address, _to []common.Address, _messageHash [][32]byte) (*MessageManagerMessageSentIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	var _messageHashRule []interface{}
	for _, _messageHashItem := range _messageHash {
		_messageHashRule = append(_messageHashRule, _messageHashItem)
	}

	logs, sub, err := _MessageManager.contract.FilterLogs(opts, "MessageSent", _fromRule, _toRule, _messageHashRule)
	if err != nil {
		return nil, err
	}
	return &MessageManagerMessageSentIterator{contract: _MessageManager.contract, event: "MessageSent", logs: logs, sub: sub}, nil
}

// WatchMessageSent is a free log subscription operation binding the contract event 0xb7f33a96ca66e7a604eaea9cd873498ec3775d6b36cf3cec44c8df6b467b2e46.
//
// Solidity: event MessageSent(uint256 sourceChainId, uint256 destChainId, address sourceTokenAddress, address destTokenAddress, address indexed _from, address indexed _to, uint256 _fee, uint256 _value, uint256 _nonce, bytes32 indexed _messageHash)
func (_MessageManager *MessageManagerFilterer) WatchMessageSent(opts *bind.WatchOpts, sink chan<- *MessageManagerMessageSent, _from []common.Address, _to []common.Address, _messageHash [][32]byte) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	var _messageHashRule []interface{}
	for _, _messageHashItem := range _messageHash {
		_messageHashRule = append(_messageHashRule, _messageHashItem)
	}

	logs, sub, err := _MessageManager.contract.WatchLogs(opts, "MessageSent", _fromRule, _toRule, _messageHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageManagerMessageSent)
				if err := _MessageManager.contract.UnpackLog(event, "MessageSent", log); err != nil {
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

// ParseMessageSent is a log parse operation binding the contract event 0xb7f33a96ca66e7a604eaea9cd873498ec3775d6b36cf3cec44c8df6b467b2e46.
//
// Solidity: event MessageSent(uint256 sourceChainId, uint256 destChainId, address sourceTokenAddress, address destTokenAddress, address indexed _from, address indexed _to, uint256 _fee, uint256 _value, uint256 _nonce, bytes32 indexed _messageHash)
func (_MessageManager *MessageManagerFilterer) ParseMessageSent(log types.Log) (*MessageManagerMessageSent, error) {
	event := new(MessageManagerMessageSent)
	if err := _MessageManager.contract.UnpackLog(event, "MessageSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MessageManager contract.
type MessageManagerOwnershipTransferredIterator struct {
	Event *MessageManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MessageManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageManagerOwnershipTransferred)
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
		it.Event = new(MessageManagerOwnershipTransferred)
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
func (it *MessageManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageManagerOwnershipTransferred represents a OwnershipTransferred event raised by the MessageManager contract.
type MessageManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MessageManager *MessageManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MessageManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MessageManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MessageManagerOwnershipTransferredIterator{contract: _MessageManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MessageManager *MessageManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MessageManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MessageManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageManagerOwnershipTransferred)
				if err := _MessageManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MessageManager *MessageManagerFilterer) ParseOwnershipTransferred(log types.Log) (*MessageManagerOwnershipTransferred, error) {
	event := new(MessageManagerOwnershipTransferred)
	if err := _MessageManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
