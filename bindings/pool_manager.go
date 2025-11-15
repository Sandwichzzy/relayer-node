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

// IPoolManagerKeyValuePair is an auto generated low-level Go binding around an user-defined struct.
type IPoolManagerKeyValuePair struct {
	Key   common.Address
	Value *big.Int
}

// IPoolManagerPool is an auto generated low-level Go binding around an user-defined struct.
type IPoolManagerPool struct {
	StartTimeStamp  uint32
	EndTimeStamp    uint32
	Token           common.Address
	TotalAmount     *big.Int
	TotalFee        *big.Int
	TotalFeeClaimed *big.Int
	IsCompleted     bool
}

// IPoolManagerUser is an auto generated low-level Go binding around an user-defined struct.
type IPoolManagerUser struct {
	IsWithdrawed bool
	Token        common.Address
	StartPoolId  *big.Int
	EndPoolId    *big.Int
	Amount       *big.Int
}

// PoolManagerMetaData contains all meta data concerning the PoolManager contract.
var PoolManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"BridgeFinalizeERC20\",\"inputs\":[{\"name\":\"sourceChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_fee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"BridgeFinalizeETH\",\"inputs\":[{\"name\":\"sourceChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourceTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_fee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"BridgeInitiateERC20\",\"inputs\":[{\"name\":\"sourceChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourceTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"BridgeInitiateETH\",\"inputs\":[{\"name\":\"sourceChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"ClaimAllReward\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ClaimbyID\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"CompletePoolAndNew\",\"inputs\":[{\"name\":\"CompletePools\",\"type\":\"tuple[]\",\"internalType\":\"structIPoolManager.Pool[]\",\"components\":[{\"name\":\"startTimeStamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"endTimeStamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"TotalAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"TotalFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"TotalFeeClaimed\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"IsCompleted\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"DepositAndStakingERC20\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DepositAndStakingETH\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"ETH_ADDRESS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"FeePoolValue\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"FundingPoolBalance\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"IsSupportToken\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"IsSupportedChainId\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MinStakeAmount\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MinTransferAmount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PerFee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"Pools\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"startTimeStamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"endTimeStamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"TotalAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"TotalFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"TotalFeeClaimed\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"IsCompleted\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"QuickSendAssertToUser\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"SupportTokens\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"Users\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"isWithdrawed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"StartPoolId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"EndPoolId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"WithdrawAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"WithdrawByID\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"WithdrawPoolManagerAssetTo\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"assertBalanceMessager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"depositErc20ToBridge\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositEthToBridge\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"fetchFundingPoolBalance\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPool\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIPoolManager.Pool\",\"components\":[{\"name\":\"startTimeStamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"endTimeStamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"TotalAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"TotalFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"TotalFeeClaimed\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"IsCompleted\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPoolLength\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPrincipal\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIPoolManager.KeyValuePair[]\",\"components\":[{\"name\":\"key\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getReward\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIPoolManager.KeyValuePair[]\",\"components\":[{\"name\":\"key\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUser\",\"inputs\":[{\"name\":\"_user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIPoolManager.User[]\",\"components\":[{\"name\":\"isWithdrawed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"StartPoolId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"EndPoolId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUserLength\",\"inputs\":[{\"name\":\"_user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_messageManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_relayerAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_withdrawManager\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"messageManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIMessageManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"periodTime\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"relayerAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinStakeAmount\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinTransferAmount\",\"inputs\":[{\"name\":\"_MinTransferAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPerFee\",\"inputs\":[{\"name\":\"_PerFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSupportERC20Token\",\"inputs\":[{\"name\":\"ERC20Address\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isValid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSupportToken\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_isSupport\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"startTimes\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setValidChainId\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isValid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakingMessageNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawErc20FromBridge\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"withdrawAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawEthFromBridge\",\"inputs\":[{\"name\":\"withdrawAddress\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"withdrawManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"ClaimReward\",\"inputs\":[{\"name\":\"_user\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"startPoolId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"EndPoolId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"Reward\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CompletePoolEvent\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"poolIndex\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DepositToken\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FinalizeERC20\",\"inputs\":[{\"name\":\"sourceChainId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"destChainId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"sourceTokenAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"destTokenAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FinalizeETH\",\"inputs\":[{\"name\":\"sourceChainId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"destChainId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"sourceTokenAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InitiateERC20\",\"inputs\":[{\"name\":\"sourceChainId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"destChainId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"sourceTokenAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"destTokenAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InitiateETH\",\"inputs\":[{\"name\":\"sourceChainId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"destChainId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"destTokenAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetMinStakeAmountEvent\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetMinTransferAmount\",\"inputs\":[{\"name\":\"_MinTransferAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetPerFee\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetSupportTokenEvent\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"isSupport\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetValidChainId\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"isValid\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StakingERC20Event\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StakingETHEvent\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdraw\",\"inputs\":[{\"name\":\"_user\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"startPoolId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"EndPoolId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"Amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"Reward\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawToken\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"withdrawAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyClaimed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainIdIsNotSupported\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ChainIdMismatch\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"actual\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ChainIdNotSupported\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ErrorBlockChain\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LessThanMinStakeAmount\",\"inputs\":[{\"name\":\"minAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"providedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"LessThanMinTransferAmount\",\"inputs\":[{\"name\":\"MinTransferAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"LessThanZero\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NewPoolIsNotCreate\",\"inputs\":[{\"name\":\"PoolIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NoReward\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotEnoughETH\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotEnoughToken\",\"inputs\":[{\"name\":\"ERC20Address\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OutOfRange\",\"inputs\":[{\"name\":\"PoolId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"PoolLength\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"PoolIsCompleted\",\"inputs\":[{\"name\":\"poolIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenIsAlreadySupported\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isSupported\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"type\":\"error\",\"name\":\"TokenIsNotSupported\",\"inputs\":[{\"name\":\"ERC20Address\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TransferETHFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Zero\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"sourceChainIdError\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"sourceChainIsDestChainError\",\"inputs\":[]}]",
	Bin: "0x6080806040523460ce5760017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00555f5160206142bd5f395f51905f525460ff8160401c1660bf576002600160401b03196001600160401b03821601606d575b6040516141ea90816100d38239f35b6001600160401b0319166001600160401b039081175f5160206142bd5f395f51905f525581527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a15f80605e565b63f92ee8a960e01b5f5260045ffd5b5f80fdfe6080806040526004361015610024575b50361561001a575f80fd5b610022613c5b565b005b5f905f3560e01c90816309dc480c146134e55750806310875a13146134bd57806313e8544e1461343257806318a7cca81461340a5780631ca2f173146133d25780631d31fac0146133b05780631f9eda8f1461337857806324237f84146131845780632c37388b14612e1a57806331d98b3b14612dfc578063325e5d43146117675780633338562c14612b8057806334d34ef51461294757806334fa7f70146127f45780633d18b912146125635780633f4ba83a146124e35780634663cdc8146124395780634dfaef84146121945780634eb7154214611ec157806352662af314611e5157806354dc027e14611e185780635a648bc514611a775780635b5b9ea2146119575780635bc65b40146119285780635c975abb146118f957806360ba33ff1461182a578063626417b5146117e6578063650c22761461177f5780636b7edb55146117675780636ec9f812146114ea5780636f77926b146113bc578063715018a6146113535780638456cb59146112e05780638da5cb5b146112ab578063929048b41461128257806394a009781461106e57806396f984e614610e7d578063a69bdf1614610d3b578063a734f06e14610d13578063ab0f9e1914610c64578063b92e639614610c09578063bb8fc7a814610be9578063cb314fab14610b53578063cb4f04ad14610b1a578063d73f14e514610a95578063dac2956814610a56578063dbb0481f14610a1d578063dd0c3460146109ff578063e207089314610849578063ec3e9da514610820578063f2fde38b146107f8578063f62a89cf146104ad578063f8c8765e146102b95763f8fcc2aa0361000f57346102b65760203660031901126102b6576020906040906001600160a01b036102a56134ff565b168152600e83522054604051908152f35b80fd5b50346102b65760803660031901126102b6576102d36134ff565b6024356001600160a01b038116908190036104a9576102f0613515565b6102f861352b565b915f5160206141955f395f51905f52549360ff8560401c16159467ffffffffffffffff8116801590816104a1575b6001149081610497575b15908161048e575b5061047f5767ffffffffffffffff1981166001175f5160206141955f395f51905f525561037c919086610453575b5061036f6140c9565b6103776140c9565b613ce1565b6103846140c9565b6001600160601b0360a01b600554161760055560018060a01b03166001600160601b0360a01b600454161760045560018060a01b03166001600160601b0360a01b6008541617600855621baf8063ffffffff1983541617825567016345785d8a000060015561271060025560016003556103fb5780f35b68ff0000000000000000195f5160206141955f395f51905f5254165f5160206141955f395f51905f52557fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602060405160018152a180f35b68ffffffffffffffffff191668010000000000000001175f5160206141955f395f51905f52555f610366565b63f92ee8a960e01b8752600487fd5b9050155f610338565b303b159150610330565b879150610326565b8280fd5b50346102b65760203660031901126102b657600435906104cb613d52565b6104d3613d8a565b338152600f60205260408120548210156107d557338152600f6020526104fc8260408320613557565b505460081c6001600160a01b0316808252600e60205260408220549092905f1981019081116107c15782338452600f602052600361053d8460408720613557565b50015492338552600f60205260016105588260408820613557565b50015493338652600f60205260ff6105738360408920613557565b5054166107b25782855b858110610701575061058e916136ff565b90338652600f60205260036105a68260408920613557565b500154878752600e6020526105cd60016105c38760408b206136ce565b5001918254613751565b90556105da823389613e5f565b50338652600f6020526040862054610665575b50905f5160206141355f395f51905f52949561060c8261064e94613751565b6040805133815260208101979097528601949094524660608601526001600160a01b0316608085015260a084019290925260c0830191909152819060e0820190565b0390a160015f5160206141755f395f51905f525580f35b338652600f60205260408620338752600f6020526040872054905f1982019182116106ed579380936106cd896106c75f5160206141355f395f51905f529b9c9660406106b761060c9861064e9c613557565b5093338152600f60205220613557565b90613db1565b338a52600f6020526106e160408b20613e19565b935093505095946105ed565b634e487b7160e01b88526011600452602488fd5b9050878752600e602052610726600261071d8360408b206136ce565b50015483613720565b670de0b6b3a7640000810290808204670de0b6b3a764000014901517156106ed5790610783610771600193846107688560408f8f908152600e602052206136ce565b50015490613733565b95670de0b6b3a76400008704906136ff565b94898952600e6020526107a8600361079e8460408d206136ce565b50019182546136ff565b905501839061057d565b63374c934360e11b8652600486fd5b634e487b7160e01b83526011600452602483fd5b604491338252600f60205260408220549063abe5c32f60e01b8352600452602452fd5b50346102b65760203660031901126102b65761081d6108156134ff565b61037761402d565b80f35b50346102b657806003193601126102b6576008546040516001600160a01b039091168152602090f35b50346102b65760603660031901126102b6576108636134ff565b61086b613671565b6044359163ffffffff83168084036109fb5761089260018060a01b03600454163314613809565b6001600160a01b038216808652600a602052604086205490949060ff166109e157848652600a6020526108d484604088209060ff801983541691151516179055565b848652600e6020526040862063ffffffff87541683039063ffffffff82116106ed577ff37cdf0b5a2a780dfbdbd1bb6f8a3bcdf62c8941ba001debfa56dfd8eab89b6c95936109c3936109676109db97946109be9463ffffffff6040519261093b846135be565b1682528460208301528b60408301528c60608301528c60808301528c60a08301528c60c0830152613899565b888a52600e60205263ffffffff61098660408c2092828d54169061387f565b60405193610993856135be565b84521660208301528860408301528960608301528960808301528960a08301528960c0830152613899565b613bf6565b60408051911515825246602083015290918291820190565b0390a280f35b63411befff60e11b86526004859052831515602452604486fd5b8480fd5b50346102b657806003193601126102b6576020600254604051908152f35b50346102b65760203660031901126102b6576020906040906001600160a01b03610a456134ff565b168152600b83522054604051908152f35b50346102b65760203660031901126102b65760209060ff906040906001600160a01b03610a816134ff565b168152600a84522054166040519015158152f35b50346102b65760403660031901126102b6577f756e8389496652aad1bd6bd9ea1e81b5b53341a798eb468701bb68403594c9506040600435610ad5613671565b610aea60018060a01b03600454163314613809565b8185526009602052610b0a818487209060ff801983541691151516179055565b825191825215156020820152a180f35b50346102b65760203660031901126102b6576020906040906001600160a01b03610b426134ff565b168152600f83522054604051908152f35b50346102b65760403660031901126102b657610b6d6134ff565b60243590610b8660018060a01b03600454163314613809565b8115610bd55760407f6559cfa4c9d2d533663fa9b68e86d2bb216ac0d2bf462b0529ab01525dce08379160018060a01b031692838552600d60205280828620558151908152466020820152a280f35b637a11cf7560e01b83526004829052602483fd5b50806003193601126102b6576020610bff613c5b565b6040519015158152f35b50346102b65760203660031901126102b6577f39315f035b4e62ef103e9d2372379ecc1fda660b2b8bf306693ff8a07ed1dd3f6020600435610c5660018060a01b03600454163314613809565b80600155604051908152a180f35b50346102b65760403660031901126102b657610c7e6134ff565b7ff37cdf0b5a2a780dfbdbd1bb6f8a3bcdf62c8941ba001debfa56dfd8eab89b6c6109db610caa613671565b610cbf60018060a01b03600454163314613809565b6001600160a01b038416808652600a60205260408620805460ff191660ff841515161790559381610d04575b5060408051911515825246602083015290918291820190565b610d0d90613bf6565b5f610ceb565b50346102b657806003193601126102b65760206040515f5160206141155f395f51905f528152f35b50346102b657806003193601126102b65760065490610d598261395a565b90805b838110610d755760405180610d718582613680565b0390f35b8190825b338452600f6020526040842054811015610e3157338452600f602052610da28160408620613557565b505460081c6001600160a01b0316610db9836136e7565b905460039190911b1c6001600160a01b031614610ddb575b6001905b01610d79565b91338452600f60205260ff610df38460408720613557565b505416610e2857610e20600191338652600f6020526003610e178660408920613557565b500154906136ff565b929050610dd1565b91600190610dd5565b50600191610e3e826136e7565b848060a01b0391549060031b1c169060405191610e5a836135ee565b82526020820152610e6b828661386b565b52610e76818561386b565b5001610d5c565b50806003193601126102b657610e91613d52565b610e99613d8a565b5f5160206141155f395f51905f528152600d6020526040812054341061103e575f5160206141155f395f51905f528152600e60205260408120541561102a575f5160206141155f395f51905f528152600e60205260408120545f198101908082116107c1575f5160206141155f395f51905f528352600e60205263ffffffff610f2583604086206136ce565b5054164210156110185750610f9d600191338452600f602052610f7f6040852060405190610f528261360a565b8682525f5160206141155f395f51905f5260208301528360408301528660608301523460808301526139b5565b5f5160206141155f395f51905f528452600e602052604084206136ce565b5001610faa3482546136ff565b90555f5160206141155f395f51905f528152600b60205260408120610fd03482546136ff565b90556040514681523460208201527fc30647c83f1ae8cba53f299401139cccd78519cd0cdf3937716ee01df27d509f60403392a260015f5160206141755f395f51905f525580f35b637d58ebb960e01b8352600452602482fd5b637d58ebb960e01b81526001600452602490fd5b5f5160206141155f395f51905f528152600d60205260408120546327500c6d60e21b825260045234602452604490fd5b506101003660031901126102b65760043560243561108a613515565b9061109361352b565b9361109c613541565b9060a435916110a9613d52565b6110b1613d8a565b6110c660018060a01b03600454163314613809565b46840361126a57858252600960205260ff60408320541615611256576001600160a01b03169581808080868b5af16110fc613a37565b5015611247575f5160206141155f395f51905f528252600b60205260408220611126848254613751565b90556005546001600160a01b031690813b156104a957829161012483926040519485938492630f5c3d0560e31b84528c60048501528a602485015260018060a01b038c1660448501525f5160206141155f395f51905f52606485015260018060a01b031660848401528c60a48401528860c484015260c43560e484015260e4356101048401525af1801561123c57611227575b50506040805194855260208501929092526001600160a01b0390921690830152606082015230907feb0a2e1a2d528efe194dccba78e1f08be6e84e330206377d56c2a2fb492e5df79080608081015b0390a360015f5160206141755f395f51905f5255602060405160018152f35b611232828092613626565b6102b657806111b9565b6040513d84823e3d90fd5b6389a360eb60e01b8252600482fd5b632ef6974160e11b82526004869052602482fd5b630432cec160e31b8252600484905246602452604482fd5b50346102b657806003193601126102b6576007546040516001600160a01b039091168152602090f35b50346102b657806003193601126102b6575f5160206140f55f395f51905f52546040516001600160a01b039091168152602090f35b50346102b657806003193601126102b6576112f961402d565b611301613d8a565b600160ff195f5160206141555f395f51905f525416175f5160206141555f395f51905f52557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586020604051338152a180f35b50346102b657806003193601126102b65761136c61402d565b5f5160206140f55f395f51905f5280546001600160a01b0319811690915581906001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a380f35b50346102b65760203660031901126102b6576001600160a01b036113de6134ff565b168152600f6020526040812080546113f581613648565b906114036040519283613626565b80825260208201809385526020852085915b83831061149357868587604051928392602084019060208552518091526040840192915b818110611447575050500390f35b91935091602060a0600192608087518051151583528580851b03858201511685840152604081015160408401526060810151606084015201516080820152019401910191849392611439565b600460206001926040516114a68161360a565b855460ff811615158252858060a01b039060081c16838201528486015460408201526002860154606082015260038601546080820152815201920192019190611415565b50346102b6576101203660031901126102b6576004356024359161150c613515565b9261151561352b565b61151d613541565b9260a435906001600160a01b0382168083036117635760c4359361153f613d8a565b61155460018060a01b03600454163314613809565b46860361175457878352600960205260ff6040842054161561174057818352600a60205260ff6040842054161561172c576040516370a0823160e01b8152306004820152602081602481865afa801561172157869085906116e7575b6115bc92501015613b8d565b6001600160a01b038716906115d2868284613fac565b828452600b602052604084206115e9878254613751565b90556005546001600160a01b0316998a3b156109fb578985809493610124936040519788968795630f5c3d0560e31b875260048701528d60248701526044860152606485015260018060a01b0316608484015260018060a01b03169c8d60a48401528960c484015260e43560e4840152610104356101048401525af1801561123c576116d2575b50506040805195865260208601939093526001600160a01b039384169285019290925291166060830152608082015230907f7add2d08613c9e66b48e0cd9188091059a1eeeceaeed10d42cb1720aa8dec7679060a090a3602060405160018152f35b6116dd828092613626565b6102b65780611670565b50506020813d602011611719575b8161170260209383613626565b8101031261171557856115bc91516115b0565b5f80fd5b3d91506116f5565b6040513d86823e3d90fd5b6305fd61ad60e01b83526004829052602483fd5b632ef6974160e11b83526004889052602483fd5b631a26aa4d60e21b8352600483fd5b5080fd5b50346102b65761081d61177936613584565b916137e6565b50346102b65760203660031901126102b6576004356117a960018060a01b03600454163314613809565b620f4240811015611763576020817f342f865a411cdd799cc2fde1e27981fc427f5c8a31c79066a6bcc0db11729da092600255604051908152a180f35b50346102b65760203660031901126102b657600435906006548210156102b6576020611811836136e7565b905460405160039290921b1c6001600160a01b03168152f35b50346102b65760403660031901126102b6576118446134ff565b9060243591611851613d8a565b611859613d52565b6001600160a01b0316808252600a602052604082205490919060ff16156118e75760409061188984303386614060565b828152600b6020522061189d8382546136ff565b90556040519182527f4b3f81827ede20c81afbf1bb77b954afcdcae24d391d99042310cb1d9210dd5760203393a360015f5160206141755f395f51905f5255602060405160018152f35b6024916305fd61ad60e01b8252600452fd5b50346102b657806003193601126102b657602060ff5f5160206141555f395f51905f5254166040519015158152f35b50346102b65760203660031901126102b65760ff60406020926004358152600984522054166040519015158152f35b50346102b65760403660031901126102b6576119c460e0916119776134ff565b8160c0604051611986816135be565b8281528260208201528260408201528260608201528260808201528260a0820152015260018060a01b03168152600e602052604060243591206136ce565b506040516119d1816135be565b81549163ffffffff831692838352602083019063ffffffff8160201c168252604084019060018060a01b039060401c1681526001830154906060850191825263ffffffff6002850154936080870194855260c060ff600460038901549860a08b01998a52015416970196151587526040519788525116602087015260018060a01b0390511660408601525160608501525160808401525160a083015251151560c0820152f35b50346102b657806003193601126102b657611a90613d52565b611a98613d8a565b805b600654811015611e0357611aad816136e7565b905460039190911b1c6001600160a01b0316808352600e602052604083205415611def57338352600f6020526040832054805b611aee575050600101611a9a565b9092905f1981018181116107c157338352600f60205284611b128260408620613557565b505460081c6001600160a01b031614611b4e575b505b8015611b3a579092905f190180611ae0565b634e487b7160e01b82526011600452602482fd5b338352600f60205260ff611b658260408620613557565b505416611de957848352600e60205260408320545f198101908111611dd55783338552600f6020526003611b9c8460408820613557565b50015492338652600f6020526001611bb78260408920613557565b50015493845b848110611d36575082611bcf916136ff565b90338752600f602052611be58160408920613557565b50600160ff19825416179055338752600f6020526003611c088260408a20613557565b500154898852600e602052611c2560016105c38760408c206136ce565b9055611c3282338b613e5f565b50338752600f6020526040872054611cb0575b5081611c645f5160206141355f395f51905f52959493611ca793613751565b6040805133815260208101969096528501939093524660608501526001600160a01b038916608085015260a084019290925260c0830191909152819060e0820190565b0390a15f611b26565b338752600f60205260408720338852600f6020526040882054905f198201918211611d225793611ca793611d028a6106c7849660406106b7611c64985f5160206141355f395f51905f529e9d9c613557565b338a52600f602052611d1660408b20613e19565b93509394955050611c45565b634e487b7160e01b89526011600452602489fd5b909192898852600e602052611d5c6002611d538460408c206136ce565b50015484613720565b670de0b6b3a7640000810290808204670de0b6b3a76400001490151715611d225791600191611dcd600361079e848f9998978e611dbf611dad6040938e8452600e6020528b610768878787206136ce565b97670de0b6b3a76400008904906136ff565b9b8152600e602052206136ce565b905501611bbd565b634e487b7160e01b84526011600452602484fd5b50611b28565b637d58ebb960e01b83526004839052602483fd5b5060015f5160206141755f395f51905f525580f35b50346102b65760203660031901126102b6576020906040906001600160a01b03611e406134ff565b168152600c83522054604051908152f35b5060403660031901126102b657600435906001600160a01b03821682036102b6576020611ea683611e8d60018060a01b0360085416331461375e565b611e95613d8a565b611e9d613d52565b60243590613a76565b60015f5160206141755f395f51905f52556040519015158152f35b50346102b65760c03660031901126102b657600435602435611ee1613515565b90611eea61352b565b611ef2613541565b90611efb613d52565b611f03613d8a565b46850361218557828652600960205260ff60408720541615612171576001600160a01b038416808752600a602052604087205490969060ff161561215d576040516370a0823160e01b81523060048201526020816024818b5afa90811561123c57829161212b575b50611f7a60a43530338b614060565b6040516370a0823160e01b8152306004820152906020826024818c5afa80156121205783906120ec575b611fae9250613751565b96808252600b60205260408220611fc68982546136ff565b9055611fe362989680611fdb6002548b613720565b048099613751565b93818352600c60205260408320611ffb8a82546136ff565b90556005546001600160a01b031691823b156120e857906101048492836040519586948593635478424f60e11b85524660048601528c6024860152604485015260018060a01b038a16606485015233608485015260018060a01b03169d8e60a48501528a60c485015260e48401525af1801561123c576120d3575b50506040805195865260208601939093526001600160a01b039384169285019290925291166060830152608082015233907f24136187bad583830e5c55416af2a551e1750b2b95b0e997565fba1badf10ed3908060a08101611208565b6120de828092613626565b6102b65780612076565b8380fd5b506020823d602011612118575b8161210660209383613626565b8101031261171557611fae9151611fa4565b3d91506120f9565b6040513d85823e3d90fd5b90506020813d602011612155575b8161214660209383613626565b8101031261171557515f611f6b565b3d9150612139565b6305fd61ad60e01b81526004879052602490fd5b632ef6974160e11b86526004839052602486fd5b631a26aa4d60e21b8652600486fd5b50346102b65760403660031901126102b6576121ae6134ff565b602435906121ba613d52565b6121c2613d8a565b6001600160a01b0316808352600a602052604083205490919060ff161561172c57818352600d60205260408320548110612417576040516370a0823160e01b815230600482015290602082602481865afa9182156117215784926123e1575b5061222e90303385614060565b6040516370a0823160e01b815230600482015290602082602481865afa80156117215784906123ad575b6122629250613751565b818352600e60205260408320541561239957818352600e60205260408320545f1981019080821161238557838552600e60205263ffffffff6122a783604088206136ce565b5054164210156123735750612305600191338652600f6020526122f460408720604051906122d48261360a565b8882528760208301528360408301528860608301528660808301526139b5565b848652600e602052604086206136ce565b50016123128282546136ff565b9055818352600b6020526040832061232b8282546136ff565b90556040519046825260208201527f1a9bb68de16395364f4a0d198852c1b2c4d1e8e5a628e5fd273284a44cd8a6db60403392a360015f5160206141755f395f51905f525580f35b637d58ebb960e01b8552600452602484fd5b634e487b7160e01b85526011600452602485fd5b637d58ebb960e01b83526001600452602483fd5b506020823d6020116123d9575b816123c760209383613626565b81010312611715576122629151612258565b3d91506123ba565b9091506020813d60201161240f575b816123fd60209383613626565b8101031261171557519061222e612221565b3d91506123f0565b908252600d60205260408220546327500c6d60e21b8352600452602452604490fd5b50346102b65760403660031901126102b6576124536134ff565b6001600160a01b03168152600e60205260408120805460243592908310156102b65760e061248184846136ce565b50805490600181015490600281015460ff6004600384015493015416926040519463ffffffff8116865263ffffffff8160201c16602087015260018060a01b039060401c1660408601526060850152608084015260a0830152151560c0820152f35b50346102b657806003193601126102b6576124fc61402d565b5f5160206141555f395f51905f525460ff8116156125545760ff19165f5160206141555f395f51905f52557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6020604051338152a180f35b638dfc202b60e01b8252600482fd5b50346102b657806003193601126102b6576006546125808161395a565b9082905b8082106125995760405180610d718582613680565b8390845b338652600f60205260408620548110156127a657338652600f6020526125c68160408820613557565b505460081c6001600160a01b03166125dd856136e7565b905460039190911b1c6001600160a01b0316146125ff575b6001905b0161259d565b338652600f60205260ff6126168260408920613557565b50541661279e57612626846136e7565b905460039190911b1c6001600160a01b03168652600e60205260408620545f19810190811161278a57338752600f60205260036126668360408a20613557565b50015490338852600f60205260016126818460408b20613557565b50015481811161277f5791905b80831061269d575050506125f5565b9091946126a9876136e7565b905460039190911b1c6001600160a01b03168952600e60205260408920545f19810190811161276b5786116127575761274e60019161274889846107688b8f6040906126f4866136e7565b868060a01b0391549060031b1c168152600e60205261272d6127278d600261271e878787206136ce565b50015490613720565b966136e7565b868060a01b0391549060031b1c168152600e602052206136ce565b906136ff565b9501919061268e565b637d58ebb960e01b89526004869052602489fd5b634e487b7160e01b8a52601160045260248afd5b5050506001906125f9565b634e487b7160e01b87526011600452602487fd5b6001906125f9565b50916001916127b4826136e7565b848060a01b0391549060031b1c1690604051916127d0836135ee565b825260208201526127e1828661386b565b526127ec818561386b565b500190612584565b50346102b65761280336613584565b9161281c60018060a09694961b0360085416331461375e565b612824613d8a565b61282c613d52565b6001600160a01b0316808252600b602052604082205490939083116128ca577f9ca7c1e047552a8048d924a5a8d3c150eb861086a72a9100e5f19d1176c1b746926040838660609552600b60205220612886828254613751565b9055612893818387613fac565b604080513381526001600160a01b039093166020840152820152a260015f5160206141755f395f51905f5255602060405160018152f35b60405162461bcd60e51b815260206004820152604960248201527f506f6f6c4d616e6167657220776974686472617745746846726f6d427269646760448201527f653a20496e73756666696369656e7420746f6b656e2062616c616e636520696e6064820152680818dbdb9d1c9858dd60ba1b608482015260a490fd5b50346102b65760203660031901126102b657600435612964613d52565b61296c613d8a565b338252600f6020526040822054811015612b5b57338252600f6020526129958160408420613557565b505460081c6001600160a01b0316808352600e60205260408320549091905f198101908111611dd5578392338552600f60205260036129d78460408820613557565b50015492338652600f60205260016129f28260408920613557565b50015493338752600f60205260ff612a0d8360408a20613557565b505416612b4c57849695965b848110612ab6575091836001612a6b8994612a597f36b99ff884a288b9b46a49442e3c4c89464469fae98da480f519664f99fc4a6f9a9b61064e986136ff565b50338b52600f60205260408b20613557565b500155612a79823383613e5f565b506040805133815260208101969096528501939093524660608501526001600160a01b03909216608084015260a0830191909152819060c0820190565b838752600e602052612ad0600261071d8360408b206136ce565b670de0b6b3a7640000810290808204670de0b6b3a764000014901517156106ed57908798612b41600361079e846040612b32612b208f9960016107688f868e849f928a9352600e602052206136ce565b96670de0b6b3a76400008804906136ff565b9d8b8152600e602052206136ce565b905501969596612a19565b63374c934360e11b8752600487fd5b338252600f602052604082205463abe5c32f60e01b8352600491909152602452604490fd5b5060203660031901126102b65760043567ffffffffffffffff81116117635736602382011215611763578060040135612bb881613648565b91612bc66040519384613626565b818352602460e0602085019302820101903682116109fb57602401915b818310612d6857505050612c0260018060a01b03600454163314613809565b815b8151811015612d64576001600160a01b036040612c21838561386b565b51015116808452600e60205260408420545f19810190811161238557907fe31e15753cd173232924147d9c429fccfb91705d3778d5eb2735a5d302ea5e81604060019493838852600e6020526004612c7b82848b206136ce565b5001805460ff19168717905580612d2d575b838852600e602052612d1b63ffffffff612ca983858c206136ce565b505460201c16858a52600e602052838a208a63ffffffff8a612ce28789612cd3858754168961387f565b958d8152600e602052206136ce565b50015491875194612cf2866135be565b8552166020840152878684015260608301528a60808301528a60a08301528a60c0830152613899565b8151908152466020820152a201612c04565b838852600c60205281882054848952600e6020526002612d4f83858c206136ce565b500155838852600c6020528782812055612c8d565b8280f35b60e0833603126109fb5760405190612d7f826135be565b612d8884613660565b8252612d9660208501613660565b602083015260408401356001600160a01b0381168103611715576040830152606084013560608301526080840135608083015260a084013560a083015260c0840135908115158203612df8578260209260c060e0950152815201920191612be3565b8680fd5b50346102b657806003193601126102b6576020600354604051908152f35b50346102b657806003193601126102b657612e33613d52565b612e3b613d8a565b805b600654811015611e0357612e50816136e7565b905460039190911b1c6001600160a01b0316808352600e602052604083205415611def57338352600f6020526040832054805b612e91575050600101612e3d565b9092905f1981018181116107c157338352600f60205284612eb58260408620613557565b505460081c6001600160a01b031614612edd575b505b8015611b3a579092905f190180612e83565b338352600f60205260ff612ef48260408620613557565b50541661317e57848352600e60205260408320545f198101908111611dd55783338552600f6020526003612f2b8460408820613557565b50015492338652600f6020526001612f468260408920613557565b50015493845b848110613108575082612f5e916136ff565b90338752600f602052612f748160408920613557565b50805460ff19166001179055861561307557338752600f6020526003612f9d8260408a20613557565b500154898852600e602052612fba60016105c38760408c206136ce565b9055612fc782338b613e5f565b50338752600f6020526040872054613003575b5081611c645f5160206141355f395f51905f52959493612ff993613751565b0390a15b5f612ec9565b338752600f60205260408720338852600f6020526040882054905f198201918211611d225793612ff9936130558a6106c7849660406106b7611c64985f5160206141355f395f51905f529e9d9c613557565b338a52600f60205261306960408b20613e19565b93509394955050612fda565b6131009150918360016130b67f36b99ff884a288b9b46a49442e3c4c89464469fae98da480f519664f99fc4a6f979695335f52600f60205260405f20613557565b5001556130c481338b613e5f565b506040805133815260208101959095528401929092524660608401526001600160a01b038816608084015260a0830191909152819060c0820190565b0390a1612ffd565b909192898852600e6020526131256002611d538460408c206136ce565b670de0b6b3a7640000810290808204670de0b6b3a76400001490151715611d225791600191613176600361079e848f9998978e611dbf611dad6040938e8452600e6020528b610768878787206136ce565b905501612f4c565b50612ecb565b60803660031901126117155760243560043561319e613515565b6131a661352b565b906131af613d52565b6131b7613d8a565b46830361336957835f52600960205260ff60405f205416156133565760015480341061334057505f5160206141155f395f51905f525f52600b60205260405f206132023482546136ff565b90556298968061321460025434613720565b04936132208534613751565b5f5160206141155f395f51905f525f52600c60205260405f206132448782546136ff565b90556005546001600160a01b031693843b1561171557604051635478424f60e11b8152466004820152602481018490525f5160206141155f395f51905f5260448201526001600160a01b03858116606483015233608483015290911660a4820181905260c4820183905260e4820197909752935f90859061010490829084905af1908115613335577f0c2b5a57887e542335dea57bfc17d53dc16f67d219e0cc41ec1970283524abcc9461120892613325575b506040519384933397859094939260609260808301968352602083015260018060a01b031660408201520152565b5f61332f91613626565b5f6132f7565b6040513d5f823e3d90fd5b6358f8331360e11b5f526004523460245260445ffd5b83632ef6974160e11b5f5260045260245ffd5b631a26aa4d60e21b5f5260045ffd5b34611715576020366003190112611715576001600160a01b036133996134ff565b165f52600b602052602060405f2054604051908152f35b34611715575f36600319011261171557602063ffffffff5f5416604051908152f35b34611715576020366003190112611715576001600160a01b036133f36134ff565b165f52600d602052602060405f2054604051908152f35b34611715575f366003190112611715576004546040516001600160a01b039091168152602090f35b346117155760403660031901126117155761344b6134ff565b6001600160a01b03165f908152600f60205260409020805460243591908210156117155760a09161347b91613557565b5080549060018101549060036002820154910154916040519360ff811615158552600180871b039060081c166020850152604084015260608301526080820152f35b34611715575f366003190112611715576005546040516001600160a01b039091168152602090f35b34611715575f366003190112611715576020906001548152f35b600435906001600160a01b038216820361171557565b604435906001600160a01b038216820361171557565b606435906001600160a01b038216820361171557565b608435906001600160a01b038216820361171557565b8054821015613570575f5260205f209060021b01905f90565b634e487b7160e01b5f52603260045260245ffd5b6060906003190112611715576004356001600160a01b038116810361171557906024356001600160a01b0381168103611715579060443590565b60e0810190811067ffffffffffffffff8211176135da57604052565b634e487b7160e01b5f52604160045260245ffd5b6040810190811067ffffffffffffffff8211176135da57604052565b60a0810190811067ffffffffffffffff8211176135da57604052565b90601f8019910116810190811067ffffffffffffffff8211176135da57604052565b67ffffffffffffffff81116135da5760051b60200190565b359063ffffffff8216820361171557565b60243590811515820361171557565b60206040818301928281528451809452019201905f5b8181106136a35750505090565b825180516001600160a01b031685526020908101518186015260409094019390920191600101613696565b8054821015613570575f52600560205f20910201905f90565b6006548110156135705760065f5260205f2001905f90565b9190820180921161370c57565b634e487b7160e01b5f52601160045260245ffd5b8181029291811591840414171561370c57565b811561373d570490565b634e487b7160e01b5f52601260045260245ffd5b9190820391821161370c57565b1561376557565b60405162461bcd60e51b815260206004820152604d60248201527f506f6f6c4d616e616765723a206f6e6c7957697468647261774d616e6167657260448201527f206f6e6c79207769746864726177206d616e616765722063616e2063616c6c2060648201526c3a3434b990333ab731ba34b7b760991b608482015260a490fd5b90613806929161380160018060a01b0360085416331461375e565b613e5f565b50565b1561381057565b60405162461bcd60e51b815260206004820152602d60248201527f506f6f6c4d616e616765723a206f6e6c792072656c617965722063616e20646f60448201526c2074686973206f70657261746560981b6064820152608490fd5b80518210156135705760209160051b010190565b9063ffffffff8091169116019063ffffffff821161370c57565b8054600160401b8110156135da576138b6916001820181556136ce565b613947578151815460208085015160408087015168010000000000000000600160e01b03911b1667ffffffff000000009190921b1663ffffffff9093166001600160e01b03199092169190911791909117178155606082015160018201556080820151600282015560a0820151600382015560c0909101516004909101805460ff191691151560ff16919091179055565b634e487b7160e01b5f525f60045260245ffd5b9061396482613648565b6139716040519182613626565b8281528092613982601f1991613648565b01905f5b82811061399257505050565b6020906040516139a1816135ee565b5f81525f8382015282828501015201613986565b8054600160401b8110156135da576139d291600182018155613557565b919091613947576080816139f7600393511515859060ff801983541691151516179055565b60208101518454610100600160a81b03191660089190911b610100600160a81b031617845560408101516001850155606081015160028501550151910155565b3d15613a71573d9067ffffffffffffffff82116135da5760405191613a66601f8201601f191660200184613626565b82523d5f602084013e565b606090565b814710613b12575f5160206141155f395f51905f525f52600b60205260405f20613aa1838254613751565b90556001600160a01b03165f80808085855af1613abc613a37565b5015613b0c5760607f9ca7c1e047552a8048d924a5a8d3c150eb861086a72a9100e5f19d1176c1b746915f5160206141155f395f51905f52936040519133835260208301526040820152a2600190565b50505f90565b60405162461bcd60e51b815260206004820152604760248201527f506f6f6c4d616e6167657220776974686472617745746846726f6d427269646760448201527f653a20696e73756666696369656e74204554482062616c616e636520696e20636064820152661bdb9d1c9858dd60ca1b608482015260a490fd5b15613b9457565b60405162461bcd60e51b815260206004820152603460248201527f506f6f6c4d616e616765723a20696e73756666696369656e7420746f6b656e206044820152733130b630b731b2903337b9103a3930b739b332b960611b6064820152608490fd5b600654600160401b8110156135da57600181016006556006548110156135705760065f527ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f0180546001600160a01b0319166001600160a01b03909216919091179055565b613c63613d8a565b613c6b613d52565b5f5160206141155f395f51905f525f52600b60205260405f20613c8f3482546136ff565b90556040513481525f5160206141155f395f51905f527f4b3f81827ede20c81afbf1bb77b954afcdcae24d391d99042310cb1d9210dd5760203393a360015f5160206141755f395f51905f5255600190565b6001600160a01b03168015613d3f575f5160206140f55f395f51905f5280546001600160a01b0319811683179091556001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e05f80a3565b631e4fbdf760e01b5f525f60045260245ffd5b60025f5160206141755f395f51905f525414613d7b5760025f5160206141755f395f51905f5255565b633ee5aeb560e01b5f5260045ffd5b60ff5f5160206141555f395f51905f525416613da257565b63d93c066560e01b5f5260045ffd5b919061394757808203613dc2575050565b600381613de060ff83945416859060ff801983541691151516179055565b80548454610100600160a81b031916610100600160a81b0390911617845560018101546001850155600281015460028501550154910155565b80548015613e4b575f190190613e2f8282613557565b613947576003815f809355826001820155826002820155015555565b634e487b7160e01b5f52603160045260245ffd5b6001600160a01b03165f818152600a602052604090205460ff1615613f9a57805f52600b602052613e968360405f20541015613b8d565b805f52600b60205260405f20613ead848254613751565b90555f5160206141155f395f51905f528103613f0f5750814710613f00575f918291829182916001600160a01b03165af1613ee6613a37565b5015613ef157600190565b6389a360eb60e01b5f5260045ffd5b632c1d501360e11b5f5260045ffd5b9291906040516370a0823160e01b8152306004820152602081602481885afa80156133355783915f91613f65575b5010613f5257613f4d9293613fac565b600190565b836311745b6160e21b5f5260045260245ffd5b9150506020813d602011613f92575b81613f8160209383613626565b81010312611715578290515f613f3d565b3d9150613f74565b6305fd61ad60e01b5f5260045260245ffd5b916040519163a9059cbb60e01b5f5260018060a01b031660045260245260205f60448180865af19060015f511482161561400c575b60405215613fec5750565b635274afe760e01b5f9081526001600160a01b0391909116600452602490fd5b90600181151661402457823b15153d15161690613fe1565b503d5f823e3d90fd5b5f5160206140f55f395f51905f52546001600160a01b0316330361404d57565b63118cdaa760e01b5f523360045260245ffd5b6040516323b872dd60e01b5f9081526001600160a01b039384166004529290931660245260449390935260209060648180865af19060015f51148216156140b1575b6040525f60605215613fec5750565b90600181151661402457823b15153d151616906140a2565b60ff5f5160206141955f395f51905f525460401c16156140e557565b631afcd79f60e31b5f5260045ffdfe9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300000000000000000000000000eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee7ec1c5ac03ffbc42e14868f1c00af28ca9cc9e4f02bec224142d7cc1ad64ca8cd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033009b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00a2646970667358221220afc10c1babe1df90a397c04aa9724bd36e82b9b0d512ec2052b4c032441fcf6964736f6c634300081e0033f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00",
}

// PoolManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use PoolManagerMetaData.ABI instead.
var PoolManagerABI = PoolManagerMetaData.ABI

// PoolManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PoolManagerMetaData.Bin instead.
var PoolManagerBin = PoolManagerMetaData.Bin

// DeployPoolManager deploys a new Ethereum contract, binding an instance of PoolManager to it.
func DeployPoolManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PoolManager, error) {
	parsed, err := PoolManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PoolManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PoolManager{PoolManagerCaller: PoolManagerCaller{contract: contract}, PoolManagerTransactor: PoolManagerTransactor{contract: contract}, PoolManagerFilterer: PoolManagerFilterer{contract: contract}}, nil
}

// PoolManager is an auto generated Go binding around an Ethereum contract.
type PoolManager struct {
	PoolManagerCaller     // Read-only binding to the contract
	PoolManagerTransactor // Write-only binding to the contract
	PoolManagerFilterer   // Log filterer for contract events
}

// PoolManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type PoolManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PoolManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PoolManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PoolManagerSession struct {
	Contract     *PoolManager      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoolManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PoolManagerCallerSession struct {
	Contract *PoolManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// PoolManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PoolManagerTransactorSession struct {
	Contract     *PoolManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// PoolManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type PoolManagerRaw struct {
	Contract *PoolManager // Generic contract binding to access the raw methods on
}

// PoolManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PoolManagerCallerRaw struct {
	Contract *PoolManagerCaller // Generic read-only contract binding to access the raw methods on
}

// PoolManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PoolManagerTransactorRaw struct {
	Contract *PoolManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPoolManager creates a new instance of PoolManager, bound to a specific deployed contract.
func NewPoolManager(address common.Address, backend bind.ContractBackend) (*PoolManager, error) {
	contract, err := bindPoolManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PoolManager{PoolManagerCaller: PoolManagerCaller{contract: contract}, PoolManagerTransactor: PoolManagerTransactor{contract: contract}, PoolManagerFilterer: PoolManagerFilterer{contract: contract}}, nil
}

// NewPoolManagerCaller creates a new read-only instance of PoolManager, bound to a specific deployed contract.
func NewPoolManagerCaller(address common.Address, caller bind.ContractCaller) (*PoolManagerCaller, error) {
	contract, err := bindPoolManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PoolManagerCaller{contract: contract}, nil
}

// NewPoolManagerTransactor creates a new write-only instance of PoolManager, bound to a specific deployed contract.
func NewPoolManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*PoolManagerTransactor, error) {
	contract, err := bindPoolManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PoolManagerTransactor{contract: contract}, nil
}

// NewPoolManagerFilterer creates a new log filterer instance of PoolManager, bound to a specific deployed contract.
func NewPoolManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*PoolManagerFilterer, error) {
	contract, err := bindPoolManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PoolManagerFilterer{contract: contract}, nil
}

// bindPoolManager binds a generic wrapper to an already deployed contract.
func bindPoolManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PoolManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoolManager *PoolManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PoolManager.Contract.PoolManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoolManager *PoolManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolManager.Contract.PoolManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoolManager *PoolManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoolManager.Contract.PoolManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoolManager *PoolManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PoolManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoolManager *PoolManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoolManager *PoolManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoolManager.Contract.contract.Transact(opts, method, params...)
}

// ETHADDRESS is a free data retrieval call binding the contract method 0xa734f06e.
//
// Solidity: function ETH_ADDRESS() view returns(address)
func (_PoolManager *PoolManagerCaller) ETHADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoolManager.contract.Call(opts, &out, "ETH_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ETHADDRESS is a free data retrieval call binding the contract method 0xa734f06e.
//
// Solidity: function ETH_ADDRESS() view returns(address)
func (_PoolManager *PoolManagerSession) ETHADDRESS() (common.Address, error) {
	return _PoolManager.Contract.ETHADDRESS(&_PoolManager.CallOpts)
}

// ETHADDRESS is a free data retrieval call binding the contract method 0xa734f06e.
//
// Solidity: function ETH_ADDRESS() view returns(address)
func (_PoolManager *PoolManagerCallerSession) ETHADDRESS() (common.Address, error) {
	return _PoolManager.Contract.ETHADDRESS(&_PoolManager.CallOpts)
}

// FeePoolValue is a free data retrieval call binding the contract method 0x54dc027e.
//
// Solidity: function FeePoolValue(address ) view returns(uint256)
func (_PoolManager *PoolManagerCaller) FeePoolValue(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PoolManager.contract.Call(opts, &out, "FeePoolValue", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeePoolValue is a free data retrieval call binding the contract method 0x54dc027e.
//
// Solidity: function FeePoolValue(address ) view returns(uint256)
func (_PoolManager *PoolManagerSession) FeePoolValue(arg0 common.Address) (*big.Int, error) {
	return _PoolManager.Contract.FeePoolValue(&_PoolManager.CallOpts, arg0)
}

// FeePoolValue is a free data retrieval call binding the contract method 0x54dc027e.
//
// Solidity: function FeePoolValue(address ) view returns(uint256)
func (_PoolManager *PoolManagerCallerSession) FeePoolValue(arg0 common.Address) (*big.Int, error) {
	return _PoolManager.Contract.FeePoolValue(&_PoolManager.CallOpts, arg0)
}

// FundingPoolBalance is a free data retrieval call binding the contract method 0xdbb0481f.
//
// Solidity: function FundingPoolBalance(address ) view returns(uint256)
func (_PoolManager *PoolManagerCaller) FundingPoolBalance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PoolManager.contract.Call(opts, &out, "FundingPoolBalance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FundingPoolBalance is a free data retrieval call binding the contract method 0xdbb0481f.
//
// Solidity: function FundingPoolBalance(address ) view returns(uint256)
func (_PoolManager *PoolManagerSession) FundingPoolBalance(arg0 common.Address) (*big.Int, error) {
	return _PoolManager.Contract.FundingPoolBalance(&_PoolManager.CallOpts, arg0)
}

// FundingPoolBalance is a free data retrieval call binding the contract method 0xdbb0481f.
//
// Solidity: function FundingPoolBalance(address ) view returns(uint256)
func (_PoolManager *PoolManagerCallerSession) FundingPoolBalance(arg0 common.Address) (*big.Int, error) {
	return _PoolManager.Contract.FundingPoolBalance(&_PoolManager.CallOpts, arg0)
}

// IsSupportToken is a free data retrieval call binding the contract method 0xdac29568.
//
// Solidity: function IsSupportToken(address ) view returns(bool)
func (_PoolManager *PoolManagerCaller) IsSupportToken(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _PoolManager.contract.Call(opts, &out, "IsSupportToken", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSupportToken is a free data retrieval call binding the contract method 0xdac29568.
//
// Solidity: function IsSupportToken(address ) view returns(bool)
func (_PoolManager *PoolManagerSession) IsSupportToken(arg0 common.Address) (bool, error) {
	return _PoolManager.Contract.IsSupportToken(&_PoolManager.CallOpts, arg0)
}

// IsSupportToken is a free data retrieval call binding the contract method 0xdac29568.
//
// Solidity: function IsSupportToken(address ) view returns(bool)
func (_PoolManager *PoolManagerCallerSession) IsSupportToken(arg0 common.Address) (bool, error) {
	return _PoolManager.Contract.IsSupportToken(&_PoolManager.CallOpts, arg0)
}

// IsSupportedChainId is a free data retrieval call binding the contract method 0x5bc65b40.
//
// Solidity: function IsSupportedChainId(uint256 ) view returns(bool)
func (_PoolManager *PoolManagerCaller) IsSupportedChainId(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _PoolManager.contract.Call(opts, &out, "IsSupportedChainId", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSupportedChainId is a free data retrieval call binding the contract method 0x5bc65b40.
//
// Solidity: function IsSupportedChainId(uint256 ) view returns(bool)
func (_PoolManager *PoolManagerSession) IsSupportedChainId(arg0 *big.Int) (bool, error) {
	return _PoolManager.Contract.IsSupportedChainId(&_PoolManager.CallOpts, arg0)
}

// IsSupportedChainId is a free data retrieval call binding the contract method 0x5bc65b40.
//
// Solidity: function IsSupportedChainId(uint256 ) view returns(bool)
func (_PoolManager *PoolManagerCallerSession) IsSupportedChainId(arg0 *big.Int) (bool, error) {
	return _PoolManager.Contract.IsSupportedChainId(&_PoolManager.CallOpts, arg0)
}

// MinStakeAmount is a free data retrieval call binding the contract method 0x1ca2f173.
//
// Solidity: function MinStakeAmount(address ) view returns(uint256)
func (_PoolManager *PoolManagerCaller) MinStakeAmount(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PoolManager.contract.Call(opts, &out, "MinStakeAmount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinStakeAmount is a free data retrieval call binding the contract method 0x1ca2f173.
//
// Solidity: function MinStakeAmount(address ) view returns(uint256)
func (_PoolManager *PoolManagerSession) MinStakeAmount(arg0 common.Address) (*big.Int, error) {
	return _PoolManager.Contract.MinStakeAmount(&_PoolManager.CallOpts, arg0)
}

// MinStakeAmount is a free data retrieval call binding the contract method 0x1ca2f173.
//
// Solidity: function MinStakeAmount(address ) view returns(uint256)
func (_PoolManager *PoolManagerCallerSession) MinStakeAmount(arg0 common.Address) (*big.Int, error) {
	return _PoolManager.Contract.MinStakeAmount(&_PoolManager.CallOpts, arg0)
}

// MinTransferAmount is a free data retrieval call binding the contract method 0x09dc480c.
//
// Solidity: function MinTransferAmount() view returns(uint256)
func (_PoolManager *PoolManagerCaller) MinTransferAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PoolManager.contract.Call(opts, &out, "MinTransferAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinTransferAmount is a free data retrieval call binding the contract method 0x09dc480c.
//
// Solidity: function MinTransferAmount() view returns(uint256)
func (_PoolManager *PoolManagerSession) MinTransferAmount() (*big.Int, error) {
	return _PoolManager.Contract.MinTransferAmount(&_PoolManager.CallOpts)
}

// MinTransferAmount is a free data retrieval call binding the contract method 0x09dc480c.
//
// Solidity: function MinTransferAmount() view returns(uint256)
func (_PoolManager *PoolManagerCallerSession) MinTransferAmount() (*big.Int, error) {
	return _PoolManager.Contract.MinTransferAmount(&_PoolManager.CallOpts)
}

// PerFee is a free data retrieval call binding the contract method 0xdd0c3460.
//
// Solidity: function PerFee() view returns(uint256)
func (_PoolManager *PoolManagerCaller) PerFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PoolManager.contract.Call(opts, &out, "PerFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PerFee is a free data retrieval call binding the contract method 0xdd0c3460.
//
// Solidity: function PerFee() view returns(uint256)
func (_PoolManager *PoolManagerSession) PerFee() (*big.Int, error) {
	return _PoolManager.Contract.PerFee(&_PoolManager.CallOpts)
}

// PerFee is a free data retrieval call binding the contract method 0xdd0c3460.
//
// Solidity: function PerFee() view returns(uint256)
func (_PoolManager *PoolManagerCallerSession) PerFee() (*big.Int, error) {
	return _PoolManager.Contract.PerFee(&_PoolManager.CallOpts)
}

// Pools is a free data retrieval call binding the contract method 0x4663cdc8.
//
// Solidity: function Pools(address , uint256 ) view returns(uint32 startTimeStamp, uint32 endTimeStamp, address token, uint256 TotalAmount, uint256 TotalFee, uint256 TotalFeeClaimed, bool IsCompleted)
func (_PoolManager *PoolManagerCaller) Pools(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	StartTimeStamp  uint32
	EndTimeStamp    uint32
	Token           common.Address
	TotalAmount     *big.Int
	TotalFee        *big.Int
	TotalFeeClaimed *big.Int
	IsCompleted     bool
}, error) {
	var out []interface{}
	err := _PoolManager.contract.Call(opts, &out, "Pools", arg0, arg1)

	outstruct := new(struct {
		StartTimeStamp  uint32
		EndTimeStamp    uint32
		Token           common.Address
		TotalAmount     *big.Int
		TotalFee        *big.Int
		TotalFeeClaimed *big.Int
		IsCompleted     bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StartTimeStamp = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.EndTimeStamp = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.Token = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.TotalAmount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.TotalFee = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.TotalFeeClaimed = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.IsCompleted = *abi.ConvertType(out[6], new(bool)).(*bool)

	return *outstruct, err

}

// Pools is a free data retrieval call binding the contract method 0x4663cdc8.
//
// Solidity: function Pools(address , uint256 ) view returns(uint32 startTimeStamp, uint32 endTimeStamp, address token, uint256 TotalAmount, uint256 TotalFee, uint256 TotalFeeClaimed, bool IsCompleted)
func (_PoolManager *PoolManagerSession) Pools(arg0 common.Address, arg1 *big.Int) (struct {
	StartTimeStamp  uint32
	EndTimeStamp    uint32
	Token           common.Address
	TotalAmount     *big.Int
	TotalFee        *big.Int
	TotalFeeClaimed *big.Int
	IsCompleted     bool
}, error) {
	return _PoolManager.Contract.Pools(&_PoolManager.CallOpts, arg0, arg1)
}

// Pools is a free data retrieval call binding the contract method 0x4663cdc8.
//
// Solidity: function Pools(address , uint256 ) view returns(uint32 startTimeStamp, uint32 endTimeStamp, address token, uint256 TotalAmount, uint256 TotalFee, uint256 TotalFeeClaimed, bool IsCompleted)
func (_PoolManager *PoolManagerCallerSession) Pools(arg0 common.Address, arg1 *big.Int) (struct {
	StartTimeStamp  uint32
	EndTimeStamp    uint32
	Token           common.Address
	TotalAmount     *big.Int
	TotalFee        *big.Int
	TotalFeeClaimed *big.Int
	IsCompleted     bool
}, error) {
	return _PoolManager.Contract.Pools(&_PoolManager.CallOpts, arg0, arg1)
}

// SupportTokens is a free data retrieval call binding the contract method 0x626417b5.
//
// Solidity: function SupportTokens(uint256 ) view returns(address)
func (_PoolManager *PoolManagerCaller) SupportTokens(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _PoolManager.contract.Call(opts, &out, "SupportTokens", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SupportTokens is a free data retrieval call binding the contract method 0x626417b5.
//
// Solidity: function SupportTokens(uint256 ) view returns(address)
func (_PoolManager *PoolManagerSession) SupportTokens(arg0 *big.Int) (common.Address, error) {
	return _PoolManager.Contract.SupportTokens(&_PoolManager.CallOpts, arg0)
}

// SupportTokens is a free data retrieval call binding the contract method 0x626417b5.
//
// Solidity: function SupportTokens(uint256 ) view returns(address)
func (_PoolManager *PoolManagerCallerSession) SupportTokens(arg0 *big.Int) (common.Address, error) {
	return _PoolManager.Contract.SupportTokens(&_PoolManager.CallOpts, arg0)
}

// Users is a free data retrieval call binding the contract method 0x13e8544e.
//
// Solidity: function Users(address , uint256 ) view returns(bool isWithdrawed, address token, uint256 StartPoolId, uint256 EndPoolId, uint256 Amount)
func (_PoolManager *PoolManagerCaller) Users(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	IsWithdrawed bool
	Token        common.Address
	StartPoolId  *big.Int
	EndPoolId    *big.Int
	Amount       *big.Int
}, error) {
	var out []interface{}
	err := _PoolManager.contract.Call(opts, &out, "Users", arg0, arg1)

	outstruct := new(struct {
		IsWithdrawed bool
		Token        common.Address
		StartPoolId  *big.Int
		EndPoolId    *big.Int
		Amount       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsWithdrawed = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Token = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.StartPoolId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.EndPoolId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Amount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Users is a free data retrieval call binding the contract method 0x13e8544e.
//
// Solidity: function Users(address , uint256 ) view returns(bool isWithdrawed, address token, uint256 StartPoolId, uint256 EndPoolId, uint256 Amount)
func (_PoolManager *PoolManagerSession) Users(arg0 common.Address, arg1 *big.Int) (struct {
	IsWithdrawed bool
	Token        common.Address
	StartPoolId  *big.Int
	EndPoolId    *big.Int
	Amount       *big.Int
}, error) {
	return _PoolManager.Contract.Users(&_PoolManager.CallOpts, arg0, arg1)
}

// Users is a free data retrieval call binding the contract method 0x13e8544e.
//
// Solidity: function Users(address , uint256 ) view returns(bool isWithdrawed, address token, uint256 StartPoolId, uint256 EndPoolId, uint256 Amount)
func (_PoolManager *PoolManagerCallerSession) Users(arg0 common.Address, arg1 *big.Int) (struct {
	IsWithdrawed bool
	Token        common.Address
	StartPoolId  *big.Int
	EndPoolId    *big.Int
	Amount       *big.Int
}, error) {
	return _PoolManager.Contract.Users(&_PoolManager.CallOpts, arg0, arg1)
}

// AssertBalanceMessager is a free data retrieval call binding the contract method 0x929048b4.
//
// Solidity: function assertBalanceMessager() view returns(address)
func (_PoolManager *PoolManagerCaller) AssertBalanceMessager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoolManager.contract.Call(opts, &out, "assertBalanceMessager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AssertBalanceMessager is a free data retrieval call binding the contract method 0x929048b4.
//
// Solidity: function assertBalanceMessager() view returns(address)
func (_PoolManager *PoolManagerSession) AssertBalanceMessager() (common.Address, error) {
	return _PoolManager.Contract.AssertBalanceMessager(&_PoolManager.CallOpts)
}

// AssertBalanceMessager is a free data retrieval call binding the contract method 0x929048b4.
//
// Solidity: function assertBalanceMessager() view returns(address)
func (_PoolManager *PoolManagerCallerSession) AssertBalanceMessager() (common.Address, error) {
	return _PoolManager.Contract.AssertBalanceMessager(&_PoolManager.CallOpts)
}

// FetchFundingPoolBalance is a free data retrieval call binding the contract method 0x1f9eda8f.
//
// Solidity: function fetchFundingPoolBalance(address token) view returns(uint256)
func (_PoolManager *PoolManagerCaller) FetchFundingPoolBalance(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PoolManager.contract.Call(opts, &out, "fetchFundingPoolBalance", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FetchFundingPoolBalance is a free data retrieval call binding the contract method 0x1f9eda8f.
//
// Solidity: function fetchFundingPoolBalance(address token) view returns(uint256)
func (_PoolManager *PoolManagerSession) FetchFundingPoolBalance(token common.Address) (*big.Int, error) {
	return _PoolManager.Contract.FetchFundingPoolBalance(&_PoolManager.CallOpts, token)
}

// FetchFundingPoolBalance is a free data retrieval call binding the contract method 0x1f9eda8f.
//
// Solidity: function fetchFundingPoolBalance(address token) view returns(uint256)
func (_PoolManager *PoolManagerCallerSession) FetchFundingPoolBalance(token common.Address) (*big.Int, error) {
	return _PoolManager.Contract.FetchFundingPoolBalance(&_PoolManager.CallOpts, token)
}

// GetPool is a free data retrieval call binding the contract method 0x5b5b9ea2.
//
// Solidity: function getPool(address _token, uint256 _index) view returns((uint32,uint32,address,uint256,uint256,uint256,bool))
func (_PoolManager *PoolManagerCaller) GetPool(opts *bind.CallOpts, _token common.Address, _index *big.Int) (IPoolManagerPool, error) {
	var out []interface{}
	err := _PoolManager.contract.Call(opts, &out, "getPool", _token, _index)

	if err != nil {
		return *new(IPoolManagerPool), err
	}

	out0 := *abi.ConvertType(out[0], new(IPoolManagerPool)).(*IPoolManagerPool)

	return out0, err

}

// GetPool is a free data retrieval call binding the contract method 0x5b5b9ea2.
//
// Solidity: function getPool(address _token, uint256 _index) view returns((uint32,uint32,address,uint256,uint256,uint256,bool))
func (_PoolManager *PoolManagerSession) GetPool(_token common.Address, _index *big.Int) (IPoolManagerPool, error) {
	return _PoolManager.Contract.GetPool(&_PoolManager.CallOpts, _token, _index)
}

// GetPool is a free data retrieval call binding the contract method 0x5b5b9ea2.
//
// Solidity: function getPool(address _token, uint256 _index) view returns((uint32,uint32,address,uint256,uint256,uint256,bool))
func (_PoolManager *PoolManagerCallerSession) GetPool(_token common.Address, _index *big.Int) (IPoolManagerPool, error) {
	return _PoolManager.Contract.GetPool(&_PoolManager.CallOpts, _token, _index)
}

// GetPoolLength is a free data retrieval call binding the contract method 0xf8fcc2aa.
//
// Solidity: function getPoolLength(address _token) view returns(uint256)
func (_PoolManager *PoolManagerCaller) GetPoolLength(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PoolManager.contract.Call(opts, &out, "getPoolLength", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPoolLength is a free data retrieval call binding the contract method 0xf8fcc2aa.
//
// Solidity: function getPoolLength(address _token) view returns(uint256)
func (_PoolManager *PoolManagerSession) GetPoolLength(_token common.Address) (*big.Int, error) {
	return _PoolManager.Contract.GetPoolLength(&_PoolManager.CallOpts, _token)
}

// GetPoolLength is a free data retrieval call binding the contract method 0xf8fcc2aa.
//
// Solidity: function getPoolLength(address _token) view returns(uint256)
func (_PoolManager *PoolManagerCallerSession) GetPoolLength(_token common.Address) (*big.Int, error) {
	return _PoolManager.Contract.GetPoolLength(&_PoolManager.CallOpts, _token)
}

// GetPrincipal is a free data retrieval call binding the contract method 0xa69bdf16.
//
// Solidity: function getPrincipal() view returns((address,uint256)[])
func (_PoolManager *PoolManagerCaller) GetPrincipal(opts *bind.CallOpts) ([]IPoolManagerKeyValuePair, error) {
	var out []interface{}
	err := _PoolManager.contract.Call(opts, &out, "getPrincipal")

	if err != nil {
		return *new([]IPoolManagerKeyValuePair), err
	}

	out0 := *abi.ConvertType(out[0], new([]IPoolManagerKeyValuePair)).(*[]IPoolManagerKeyValuePair)

	return out0, err

}

// GetPrincipal is a free data retrieval call binding the contract method 0xa69bdf16.
//
// Solidity: function getPrincipal() view returns((address,uint256)[])
func (_PoolManager *PoolManagerSession) GetPrincipal() ([]IPoolManagerKeyValuePair, error) {
	return _PoolManager.Contract.GetPrincipal(&_PoolManager.CallOpts)
}

// GetPrincipal is a free data retrieval call binding the contract method 0xa69bdf16.
//
// Solidity: function getPrincipal() view returns((address,uint256)[])
func (_PoolManager *PoolManagerCallerSession) GetPrincipal() ([]IPoolManagerKeyValuePair, error) {
	return _PoolManager.Contract.GetPrincipal(&_PoolManager.CallOpts)
}

// GetReward is a free data retrieval call binding the contract method 0x3d18b912.
//
// Solidity: function getReward() view returns((address,uint256)[])
func (_PoolManager *PoolManagerCaller) GetReward(opts *bind.CallOpts) ([]IPoolManagerKeyValuePair, error) {
	var out []interface{}
	err := _PoolManager.contract.Call(opts, &out, "getReward")

	if err != nil {
		return *new([]IPoolManagerKeyValuePair), err
	}

	out0 := *abi.ConvertType(out[0], new([]IPoolManagerKeyValuePair)).(*[]IPoolManagerKeyValuePair)

	return out0, err

}

// GetReward is a free data retrieval call binding the contract method 0x3d18b912.
//
// Solidity: function getReward() view returns((address,uint256)[])
func (_PoolManager *PoolManagerSession) GetReward() ([]IPoolManagerKeyValuePair, error) {
	return _PoolManager.Contract.GetReward(&_PoolManager.CallOpts)
}

// GetReward is a free data retrieval call binding the contract method 0x3d18b912.
//
// Solidity: function getReward() view returns((address,uint256)[])
func (_PoolManager *PoolManagerCallerSession) GetReward() ([]IPoolManagerKeyValuePair, error) {
	return _PoolManager.Contract.GetReward(&_PoolManager.CallOpts)
}

// GetUser is a free data retrieval call binding the contract method 0x6f77926b.
//
// Solidity: function getUser(address _user) view returns((bool,address,uint256,uint256,uint256)[])
func (_PoolManager *PoolManagerCaller) GetUser(opts *bind.CallOpts, _user common.Address) ([]IPoolManagerUser, error) {
	var out []interface{}
	err := _PoolManager.contract.Call(opts, &out, "getUser", _user)

	if err != nil {
		return *new([]IPoolManagerUser), err
	}

	out0 := *abi.ConvertType(out[0], new([]IPoolManagerUser)).(*[]IPoolManagerUser)

	return out0, err

}

// GetUser is a free data retrieval call binding the contract method 0x6f77926b.
//
// Solidity: function getUser(address _user) view returns((bool,address,uint256,uint256,uint256)[])
func (_PoolManager *PoolManagerSession) GetUser(_user common.Address) ([]IPoolManagerUser, error) {
	return _PoolManager.Contract.GetUser(&_PoolManager.CallOpts, _user)
}

// GetUser is a free data retrieval call binding the contract method 0x6f77926b.
//
// Solidity: function getUser(address _user) view returns((bool,address,uint256,uint256,uint256)[])
func (_PoolManager *PoolManagerCallerSession) GetUser(_user common.Address) ([]IPoolManagerUser, error) {
	return _PoolManager.Contract.GetUser(&_PoolManager.CallOpts, _user)
}

// GetUserLength is a free data retrieval call binding the contract method 0xcb4f04ad.
//
// Solidity: function getUserLength(address _user) view returns(uint256)
func (_PoolManager *PoolManagerCaller) GetUserLength(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PoolManager.contract.Call(opts, &out, "getUserLength", _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserLength is a free data retrieval call binding the contract method 0xcb4f04ad.
//
// Solidity: function getUserLength(address _user) view returns(uint256)
func (_PoolManager *PoolManagerSession) GetUserLength(_user common.Address) (*big.Int, error) {
	return _PoolManager.Contract.GetUserLength(&_PoolManager.CallOpts, _user)
}

// GetUserLength is a free data retrieval call binding the contract method 0xcb4f04ad.
//
// Solidity: function getUserLength(address _user) view returns(uint256)
func (_PoolManager *PoolManagerCallerSession) GetUserLength(_user common.Address) (*big.Int, error) {
	return _PoolManager.Contract.GetUserLength(&_PoolManager.CallOpts, _user)
}

// MessageManager is a free data retrieval call binding the contract method 0x10875a13.
//
// Solidity: function messageManager() view returns(address)
func (_PoolManager *PoolManagerCaller) MessageManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoolManager.contract.Call(opts, &out, "messageManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MessageManager is a free data retrieval call binding the contract method 0x10875a13.
//
// Solidity: function messageManager() view returns(address)
func (_PoolManager *PoolManagerSession) MessageManager() (common.Address, error) {
	return _PoolManager.Contract.MessageManager(&_PoolManager.CallOpts)
}

// MessageManager is a free data retrieval call binding the contract method 0x10875a13.
//
// Solidity: function messageManager() view returns(address)
func (_PoolManager *PoolManagerCallerSession) MessageManager() (common.Address, error) {
	return _PoolManager.Contract.MessageManager(&_PoolManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PoolManager *PoolManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoolManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PoolManager *PoolManagerSession) Owner() (common.Address, error) {
	return _PoolManager.Contract.Owner(&_PoolManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PoolManager *PoolManagerCallerSession) Owner() (common.Address, error) {
	return _PoolManager.Contract.Owner(&_PoolManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PoolManager *PoolManagerCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PoolManager.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PoolManager *PoolManagerSession) Paused() (bool, error) {
	return _PoolManager.Contract.Paused(&_PoolManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PoolManager *PoolManagerCallerSession) Paused() (bool, error) {
	return _PoolManager.Contract.Paused(&_PoolManager.CallOpts)
}

// PeriodTime is a free data retrieval call binding the contract method 0x1d31fac0.
//
// Solidity: function periodTime() view returns(uint32)
func (_PoolManager *PoolManagerCaller) PeriodTime(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _PoolManager.contract.Call(opts, &out, "periodTime")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// PeriodTime is a free data retrieval call binding the contract method 0x1d31fac0.
//
// Solidity: function periodTime() view returns(uint32)
func (_PoolManager *PoolManagerSession) PeriodTime() (uint32, error) {
	return _PoolManager.Contract.PeriodTime(&_PoolManager.CallOpts)
}

// PeriodTime is a free data retrieval call binding the contract method 0x1d31fac0.
//
// Solidity: function periodTime() view returns(uint32)
func (_PoolManager *PoolManagerCallerSession) PeriodTime() (uint32, error) {
	return _PoolManager.Contract.PeriodTime(&_PoolManager.CallOpts)
}

// RelayerAddress is a free data retrieval call binding the contract method 0x18a7cca8.
//
// Solidity: function relayerAddress() view returns(address)
func (_PoolManager *PoolManagerCaller) RelayerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoolManager.contract.Call(opts, &out, "relayerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RelayerAddress is a free data retrieval call binding the contract method 0x18a7cca8.
//
// Solidity: function relayerAddress() view returns(address)
func (_PoolManager *PoolManagerSession) RelayerAddress() (common.Address, error) {
	return _PoolManager.Contract.RelayerAddress(&_PoolManager.CallOpts)
}

// RelayerAddress is a free data retrieval call binding the contract method 0x18a7cca8.
//
// Solidity: function relayerAddress() view returns(address)
func (_PoolManager *PoolManagerCallerSession) RelayerAddress() (common.Address, error) {
	return _PoolManager.Contract.RelayerAddress(&_PoolManager.CallOpts)
}

// StakingMessageNumber is a free data retrieval call binding the contract method 0x31d98b3b.
//
// Solidity: function stakingMessageNumber() view returns(uint256)
func (_PoolManager *PoolManagerCaller) StakingMessageNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PoolManager.contract.Call(opts, &out, "stakingMessageNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakingMessageNumber is a free data retrieval call binding the contract method 0x31d98b3b.
//
// Solidity: function stakingMessageNumber() view returns(uint256)
func (_PoolManager *PoolManagerSession) StakingMessageNumber() (*big.Int, error) {
	return _PoolManager.Contract.StakingMessageNumber(&_PoolManager.CallOpts)
}

// StakingMessageNumber is a free data retrieval call binding the contract method 0x31d98b3b.
//
// Solidity: function stakingMessageNumber() view returns(uint256)
func (_PoolManager *PoolManagerCallerSession) StakingMessageNumber() (*big.Int, error) {
	return _PoolManager.Contract.StakingMessageNumber(&_PoolManager.CallOpts)
}

// WithdrawManager is a free data retrieval call binding the contract method 0xec3e9da5.
//
// Solidity: function withdrawManager() view returns(address)
func (_PoolManager *PoolManagerCaller) WithdrawManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoolManager.contract.Call(opts, &out, "withdrawManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WithdrawManager is a free data retrieval call binding the contract method 0xec3e9da5.
//
// Solidity: function withdrawManager() view returns(address)
func (_PoolManager *PoolManagerSession) WithdrawManager() (common.Address, error) {
	return _PoolManager.Contract.WithdrawManager(&_PoolManager.CallOpts)
}

// WithdrawManager is a free data retrieval call binding the contract method 0xec3e9da5.
//
// Solidity: function withdrawManager() view returns(address)
func (_PoolManager *PoolManagerCallerSession) WithdrawManager() (common.Address, error) {
	return _PoolManager.Contract.WithdrawManager(&_PoolManager.CallOpts)
}

// BridgeFinalizeERC20 is a paid mutator transaction binding the contract method 0x6ec9f812.
//
// Solidity: function BridgeFinalizeERC20(uint256 sourceChainId, uint256 destChainId, address from, address to, address sourceTokenAddress, address destTokenAddress, uint256 amount, uint256 _fee, uint256 _nonce) returns(bool)
func (_PoolManager *PoolManagerTransactor) BridgeFinalizeERC20(opts *bind.TransactOpts, sourceChainId *big.Int, destChainId *big.Int, from common.Address, to common.Address, sourceTokenAddress common.Address, destTokenAddress common.Address, amount *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _PoolManager.contract.Transact(opts, "BridgeFinalizeERC20", sourceChainId, destChainId, from, to, sourceTokenAddress, destTokenAddress, amount, _fee, _nonce)
}

// BridgeFinalizeERC20 is a paid mutator transaction binding the contract method 0x6ec9f812.
//
// Solidity: function BridgeFinalizeERC20(uint256 sourceChainId, uint256 destChainId, address from, address to, address sourceTokenAddress, address destTokenAddress, uint256 amount, uint256 _fee, uint256 _nonce) returns(bool)
func (_PoolManager *PoolManagerSession) BridgeFinalizeERC20(sourceChainId *big.Int, destChainId *big.Int, from common.Address, to common.Address, sourceTokenAddress common.Address, destTokenAddress common.Address, amount *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _PoolManager.Contract.BridgeFinalizeERC20(&_PoolManager.TransactOpts, sourceChainId, destChainId, from, to, sourceTokenAddress, destTokenAddress, amount, _fee, _nonce)
}

// BridgeFinalizeERC20 is a paid mutator transaction binding the contract method 0x6ec9f812.
//
// Solidity: function BridgeFinalizeERC20(uint256 sourceChainId, uint256 destChainId, address from, address to, address sourceTokenAddress, address destTokenAddress, uint256 amount, uint256 _fee, uint256 _nonce) returns(bool)
func (_PoolManager *PoolManagerTransactorSession) BridgeFinalizeERC20(sourceChainId *big.Int, destChainId *big.Int, from common.Address, to common.Address, sourceTokenAddress common.Address, destTokenAddress common.Address, amount *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _PoolManager.Contract.BridgeFinalizeERC20(&_PoolManager.TransactOpts, sourceChainId, destChainId, from, to, sourceTokenAddress, destTokenAddress, amount, _fee, _nonce)
}

// BridgeFinalizeETH is a paid mutator transaction binding the contract method 0x94a00978.
//
// Solidity: function BridgeFinalizeETH(uint256 sourceChainId, uint256 destChainId, address sourceTokenAddress, address from, address to, uint256 amount, uint256 _fee, uint256 _nonce) payable returns(bool)
func (_PoolManager *PoolManagerTransactor) BridgeFinalizeETH(opts *bind.TransactOpts, sourceChainId *big.Int, destChainId *big.Int, sourceTokenAddress common.Address, from common.Address, to common.Address, amount *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _PoolManager.contract.Transact(opts, "BridgeFinalizeETH", sourceChainId, destChainId, sourceTokenAddress, from, to, amount, _fee, _nonce)
}

// BridgeFinalizeETH is a paid mutator transaction binding the contract method 0x94a00978.
//
// Solidity: function BridgeFinalizeETH(uint256 sourceChainId, uint256 destChainId, address sourceTokenAddress, address from, address to, uint256 amount, uint256 _fee, uint256 _nonce) payable returns(bool)
func (_PoolManager *PoolManagerSession) BridgeFinalizeETH(sourceChainId *big.Int, destChainId *big.Int, sourceTokenAddress common.Address, from common.Address, to common.Address, amount *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _PoolManager.Contract.BridgeFinalizeETH(&_PoolManager.TransactOpts, sourceChainId, destChainId, sourceTokenAddress, from, to, amount, _fee, _nonce)
}

// BridgeFinalizeETH is a paid mutator transaction binding the contract method 0x94a00978.
//
// Solidity: function BridgeFinalizeETH(uint256 sourceChainId, uint256 destChainId, address sourceTokenAddress, address from, address to, uint256 amount, uint256 _fee, uint256 _nonce) payable returns(bool)
func (_PoolManager *PoolManagerTransactorSession) BridgeFinalizeETH(sourceChainId *big.Int, destChainId *big.Int, sourceTokenAddress common.Address, from common.Address, to common.Address, amount *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _PoolManager.Contract.BridgeFinalizeETH(&_PoolManager.TransactOpts, sourceChainId, destChainId, sourceTokenAddress, from, to, amount, _fee, _nonce)
}

// BridgeInitiateERC20 is a paid mutator transaction binding the contract method 0x4eb71542.
//
// Solidity: function BridgeInitiateERC20(uint256 sourceChainId, uint256 destChainId, address sourceTokenAddress, address destTokenAddress, address to, uint256 value) returns(bool)
func (_PoolManager *PoolManagerTransactor) BridgeInitiateERC20(opts *bind.TransactOpts, sourceChainId *big.Int, destChainId *big.Int, sourceTokenAddress common.Address, destTokenAddress common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PoolManager.contract.Transact(opts, "BridgeInitiateERC20", sourceChainId, destChainId, sourceTokenAddress, destTokenAddress, to, value)
}

// BridgeInitiateERC20 is a paid mutator transaction binding the contract method 0x4eb71542.
//
// Solidity: function BridgeInitiateERC20(uint256 sourceChainId, uint256 destChainId, address sourceTokenAddress, address destTokenAddress, address to, uint256 value) returns(bool)
func (_PoolManager *PoolManagerSession) BridgeInitiateERC20(sourceChainId *big.Int, destChainId *big.Int, sourceTokenAddress common.Address, destTokenAddress common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PoolManager.Contract.BridgeInitiateERC20(&_PoolManager.TransactOpts, sourceChainId, destChainId, sourceTokenAddress, destTokenAddress, to, value)
}

// BridgeInitiateERC20 is a paid mutator transaction binding the contract method 0x4eb71542.
//
// Solidity: function BridgeInitiateERC20(uint256 sourceChainId, uint256 destChainId, address sourceTokenAddress, address destTokenAddress, address to, uint256 value) returns(bool)
func (_PoolManager *PoolManagerTransactorSession) BridgeInitiateERC20(sourceChainId *big.Int, destChainId *big.Int, sourceTokenAddress common.Address, destTokenAddress common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PoolManager.Contract.BridgeInitiateERC20(&_PoolManager.TransactOpts, sourceChainId, destChainId, sourceTokenAddress, destTokenAddress, to, value)
}

// BridgeInitiateETH is a paid mutator transaction binding the contract method 0x24237f84.
//
// Solidity: function BridgeInitiateETH(uint256 sourceChainId, uint256 destChainId, address destTokenAddress, address to) payable returns(bool)
func (_PoolManager *PoolManagerTransactor) BridgeInitiateETH(opts *bind.TransactOpts, sourceChainId *big.Int, destChainId *big.Int, destTokenAddress common.Address, to common.Address) (*types.Transaction, error) {
	return _PoolManager.contract.Transact(opts, "BridgeInitiateETH", sourceChainId, destChainId, destTokenAddress, to)
}

// BridgeInitiateETH is a paid mutator transaction binding the contract method 0x24237f84.
//
// Solidity: function BridgeInitiateETH(uint256 sourceChainId, uint256 destChainId, address destTokenAddress, address to) payable returns(bool)
func (_PoolManager *PoolManagerSession) BridgeInitiateETH(sourceChainId *big.Int, destChainId *big.Int, destTokenAddress common.Address, to common.Address) (*types.Transaction, error) {
	return _PoolManager.Contract.BridgeInitiateETH(&_PoolManager.TransactOpts, sourceChainId, destChainId, destTokenAddress, to)
}

// BridgeInitiateETH is a paid mutator transaction binding the contract method 0x24237f84.
//
// Solidity: function BridgeInitiateETH(uint256 sourceChainId, uint256 destChainId, address destTokenAddress, address to) payable returns(bool)
func (_PoolManager *PoolManagerTransactorSession) BridgeInitiateETH(sourceChainId *big.Int, destChainId *big.Int, destTokenAddress common.Address, to common.Address) (*types.Transaction, error) {
	return _PoolManager.Contract.BridgeInitiateETH(&_PoolManager.TransactOpts, sourceChainId, destChainId, destTokenAddress, to)
}

// ClaimAllReward is a paid mutator transaction binding the contract method 0x2c37388b.
//
// Solidity: function ClaimAllReward() returns()
func (_PoolManager *PoolManagerTransactor) ClaimAllReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolManager.contract.Transact(opts, "ClaimAllReward")
}

// ClaimAllReward is a paid mutator transaction binding the contract method 0x2c37388b.
//
// Solidity: function ClaimAllReward() returns()
func (_PoolManager *PoolManagerSession) ClaimAllReward() (*types.Transaction, error) {
	return _PoolManager.Contract.ClaimAllReward(&_PoolManager.TransactOpts)
}

// ClaimAllReward is a paid mutator transaction binding the contract method 0x2c37388b.
//
// Solidity: function ClaimAllReward() returns()
func (_PoolManager *PoolManagerTransactorSession) ClaimAllReward() (*types.Transaction, error) {
	return _PoolManager.Contract.ClaimAllReward(&_PoolManager.TransactOpts)
}

// ClaimbyID is a paid mutator transaction binding the contract method 0x34d34ef5.
//
// Solidity: function ClaimbyID(uint256 i) returns()
func (_PoolManager *PoolManagerTransactor) ClaimbyID(opts *bind.TransactOpts, i *big.Int) (*types.Transaction, error) {
	return _PoolManager.contract.Transact(opts, "ClaimbyID", i)
}

// ClaimbyID is a paid mutator transaction binding the contract method 0x34d34ef5.
//
// Solidity: function ClaimbyID(uint256 i) returns()
func (_PoolManager *PoolManagerSession) ClaimbyID(i *big.Int) (*types.Transaction, error) {
	return _PoolManager.Contract.ClaimbyID(&_PoolManager.TransactOpts, i)
}

// ClaimbyID is a paid mutator transaction binding the contract method 0x34d34ef5.
//
// Solidity: function ClaimbyID(uint256 i) returns()
func (_PoolManager *PoolManagerTransactorSession) ClaimbyID(i *big.Int) (*types.Transaction, error) {
	return _PoolManager.Contract.ClaimbyID(&_PoolManager.TransactOpts, i)
}

// CompletePoolAndNew is a paid mutator transaction binding the contract method 0x3338562c.
//
// Solidity: function CompletePoolAndNew((uint32,uint32,address,uint256,uint256,uint256,bool)[] CompletePools) payable returns()
func (_PoolManager *PoolManagerTransactor) CompletePoolAndNew(opts *bind.TransactOpts, CompletePools []IPoolManagerPool) (*types.Transaction, error) {
	return _PoolManager.contract.Transact(opts, "CompletePoolAndNew", CompletePools)
}

// CompletePoolAndNew is a paid mutator transaction binding the contract method 0x3338562c.
//
// Solidity: function CompletePoolAndNew((uint32,uint32,address,uint256,uint256,uint256,bool)[] CompletePools) payable returns()
func (_PoolManager *PoolManagerSession) CompletePoolAndNew(CompletePools []IPoolManagerPool) (*types.Transaction, error) {
	return _PoolManager.Contract.CompletePoolAndNew(&_PoolManager.TransactOpts, CompletePools)
}

// CompletePoolAndNew is a paid mutator transaction binding the contract method 0x3338562c.
//
// Solidity: function CompletePoolAndNew((uint32,uint32,address,uint256,uint256,uint256,bool)[] CompletePools) payable returns()
func (_PoolManager *PoolManagerTransactorSession) CompletePoolAndNew(CompletePools []IPoolManagerPool) (*types.Transaction, error) {
	return _PoolManager.Contract.CompletePoolAndNew(&_PoolManager.TransactOpts, CompletePools)
}

// DepositAndStakingERC20 is a paid mutator transaction binding the contract method 0x4dfaef84.
//
// Solidity: function DepositAndStakingERC20(address _token, uint256 _amount) returns()
func (_PoolManager *PoolManagerTransactor) DepositAndStakingERC20(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PoolManager.contract.Transact(opts, "DepositAndStakingERC20", _token, _amount)
}

// DepositAndStakingERC20 is a paid mutator transaction binding the contract method 0x4dfaef84.
//
// Solidity: function DepositAndStakingERC20(address _token, uint256 _amount) returns()
func (_PoolManager *PoolManagerSession) DepositAndStakingERC20(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PoolManager.Contract.DepositAndStakingERC20(&_PoolManager.TransactOpts, _token, _amount)
}

// DepositAndStakingERC20 is a paid mutator transaction binding the contract method 0x4dfaef84.
//
// Solidity: function DepositAndStakingERC20(address _token, uint256 _amount) returns()
func (_PoolManager *PoolManagerTransactorSession) DepositAndStakingERC20(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PoolManager.Contract.DepositAndStakingERC20(&_PoolManager.TransactOpts, _token, _amount)
}

// DepositAndStakingETH is a paid mutator transaction binding the contract method 0x96f984e6.
//
// Solidity: function DepositAndStakingETH() payable returns()
func (_PoolManager *PoolManagerTransactor) DepositAndStakingETH(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolManager.contract.Transact(opts, "DepositAndStakingETH")
}

// DepositAndStakingETH is a paid mutator transaction binding the contract method 0x96f984e6.
//
// Solidity: function DepositAndStakingETH() payable returns()
func (_PoolManager *PoolManagerSession) DepositAndStakingETH() (*types.Transaction, error) {
	return _PoolManager.Contract.DepositAndStakingETH(&_PoolManager.TransactOpts)
}

// DepositAndStakingETH is a paid mutator transaction binding the contract method 0x96f984e6.
//
// Solidity: function DepositAndStakingETH() payable returns()
func (_PoolManager *PoolManagerTransactorSession) DepositAndStakingETH() (*types.Transaction, error) {
	return _PoolManager.Contract.DepositAndStakingETH(&_PoolManager.TransactOpts)
}

// QuickSendAssertToUser is a paid mutator transaction binding the contract method 0x325e5d43.
//
// Solidity: function QuickSendAssertToUser(address _token, address to, uint256 _amount) returns()
func (_PoolManager *PoolManagerTransactor) QuickSendAssertToUser(opts *bind.TransactOpts, _token common.Address, to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PoolManager.contract.Transact(opts, "QuickSendAssertToUser", _token, to, _amount)
}

// QuickSendAssertToUser is a paid mutator transaction binding the contract method 0x325e5d43.
//
// Solidity: function QuickSendAssertToUser(address _token, address to, uint256 _amount) returns()
func (_PoolManager *PoolManagerSession) QuickSendAssertToUser(_token common.Address, to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PoolManager.Contract.QuickSendAssertToUser(&_PoolManager.TransactOpts, _token, to, _amount)
}

// QuickSendAssertToUser is a paid mutator transaction binding the contract method 0x325e5d43.
//
// Solidity: function QuickSendAssertToUser(address _token, address to, uint256 _amount) returns()
func (_PoolManager *PoolManagerTransactorSession) QuickSendAssertToUser(_token common.Address, to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PoolManager.Contract.QuickSendAssertToUser(&_PoolManager.TransactOpts, _token, to, _amount)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x5a648bc5.
//
// Solidity: function WithdrawAll() returns()
func (_PoolManager *PoolManagerTransactor) WithdrawAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolManager.contract.Transact(opts, "WithdrawAll")
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x5a648bc5.
//
// Solidity: function WithdrawAll() returns()
func (_PoolManager *PoolManagerSession) WithdrawAll() (*types.Transaction, error) {
	return _PoolManager.Contract.WithdrawAll(&_PoolManager.TransactOpts)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x5a648bc5.
//
// Solidity: function WithdrawAll() returns()
func (_PoolManager *PoolManagerTransactorSession) WithdrawAll() (*types.Transaction, error) {
	return _PoolManager.Contract.WithdrawAll(&_PoolManager.TransactOpts)
}

// WithdrawByID is a paid mutator transaction binding the contract method 0xf62a89cf.
//
// Solidity: function WithdrawByID(uint256 i) returns()
func (_PoolManager *PoolManagerTransactor) WithdrawByID(opts *bind.TransactOpts, i *big.Int) (*types.Transaction, error) {
	return _PoolManager.contract.Transact(opts, "WithdrawByID", i)
}

// WithdrawByID is a paid mutator transaction binding the contract method 0xf62a89cf.
//
// Solidity: function WithdrawByID(uint256 i) returns()
func (_PoolManager *PoolManagerSession) WithdrawByID(i *big.Int) (*types.Transaction, error) {
	return _PoolManager.Contract.WithdrawByID(&_PoolManager.TransactOpts, i)
}

// WithdrawByID is a paid mutator transaction binding the contract method 0xf62a89cf.
//
// Solidity: function WithdrawByID(uint256 i) returns()
func (_PoolManager *PoolManagerTransactorSession) WithdrawByID(i *big.Int) (*types.Transaction, error) {
	return _PoolManager.Contract.WithdrawByID(&_PoolManager.TransactOpts, i)
}

// WithdrawPoolManagerAssetTo is a paid mutator transaction binding the contract method 0x6b7edb55.
//
// Solidity: function WithdrawPoolManagerAssetTo(address _token, address to, uint256 _amount) returns()
func (_PoolManager *PoolManagerTransactor) WithdrawPoolManagerAssetTo(opts *bind.TransactOpts, _token common.Address, to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PoolManager.contract.Transact(opts, "WithdrawPoolManagerAssetTo", _token, to, _amount)
}

// WithdrawPoolManagerAssetTo is a paid mutator transaction binding the contract method 0x6b7edb55.
//
// Solidity: function WithdrawPoolManagerAssetTo(address _token, address to, uint256 _amount) returns()
func (_PoolManager *PoolManagerSession) WithdrawPoolManagerAssetTo(_token common.Address, to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PoolManager.Contract.WithdrawPoolManagerAssetTo(&_PoolManager.TransactOpts, _token, to, _amount)
}

// WithdrawPoolManagerAssetTo is a paid mutator transaction binding the contract method 0x6b7edb55.
//
// Solidity: function WithdrawPoolManagerAssetTo(address _token, address to, uint256 _amount) returns()
func (_PoolManager *PoolManagerTransactorSession) WithdrawPoolManagerAssetTo(_token common.Address, to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PoolManager.Contract.WithdrawPoolManagerAssetTo(&_PoolManager.TransactOpts, _token, to, _amount)
}

// DepositErc20ToBridge is a paid mutator transaction binding the contract method 0x60ba33ff.
//
// Solidity: function depositErc20ToBridge(address tokenAddress, uint256 amount) returns(bool)
func (_PoolManager *PoolManagerTransactor) DepositErc20ToBridge(opts *bind.TransactOpts, tokenAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PoolManager.contract.Transact(opts, "depositErc20ToBridge", tokenAddress, amount)
}

// DepositErc20ToBridge is a paid mutator transaction binding the contract method 0x60ba33ff.
//
// Solidity: function depositErc20ToBridge(address tokenAddress, uint256 amount) returns(bool)
func (_PoolManager *PoolManagerSession) DepositErc20ToBridge(tokenAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PoolManager.Contract.DepositErc20ToBridge(&_PoolManager.TransactOpts, tokenAddress, amount)
}

// DepositErc20ToBridge is a paid mutator transaction binding the contract method 0x60ba33ff.
//
// Solidity: function depositErc20ToBridge(address tokenAddress, uint256 amount) returns(bool)
func (_PoolManager *PoolManagerTransactorSession) DepositErc20ToBridge(tokenAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PoolManager.Contract.DepositErc20ToBridge(&_PoolManager.TransactOpts, tokenAddress, amount)
}

// DepositEthToBridge is a paid mutator transaction binding the contract method 0xbb8fc7a8.
//
// Solidity: function depositEthToBridge() payable returns(bool)
func (_PoolManager *PoolManagerTransactor) DepositEthToBridge(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolManager.contract.Transact(opts, "depositEthToBridge")
}

// DepositEthToBridge is a paid mutator transaction binding the contract method 0xbb8fc7a8.
//
// Solidity: function depositEthToBridge() payable returns(bool)
func (_PoolManager *PoolManagerSession) DepositEthToBridge() (*types.Transaction, error) {
	return _PoolManager.Contract.DepositEthToBridge(&_PoolManager.TransactOpts)
}

// DepositEthToBridge is a paid mutator transaction binding the contract method 0xbb8fc7a8.
//
// Solidity: function depositEthToBridge() payable returns(bool)
func (_PoolManager *PoolManagerTransactorSession) DepositEthToBridge() (*types.Transaction, error) {
	return _PoolManager.Contract.DepositEthToBridge(&_PoolManager.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address initialOwner, address _messageManager, address _relayerAddress, address _withdrawManager) returns()
func (_PoolManager *PoolManagerTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address, _messageManager common.Address, _relayerAddress common.Address, _withdrawManager common.Address) (*types.Transaction, error) {
	return _PoolManager.contract.Transact(opts, "initialize", initialOwner, _messageManager, _relayerAddress, _withdrawManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address initialOwner, address _messageManager, address _relayerAddress, address _withdrawManager) returns()
func (_PoolManager *PoolManagerSession) Initialize(initialOwner common.Address, _messageManager common.Address, _relayerAddress common.Address, _withdrawManager common.Address) (*types.Transaction, error) {
	return _PoolManager.Contract.Initialize(&_PoolManager.TransactOpts, initialOwner, _messageManager, _relayerAddress, _withdrawManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address initialOwner, address _messageManager, address _relayerAddress, address _withdrawManager) returns()
func (_PoolManager *PoolManagerTransactorSession) Initialize(initialOwner common.Address, _messageManager common.Address, _relayerAddress common.Address, _withdrawManager common.Address) (*types.Transaction, error) {
	return _PoolManager.Contract.Initialize(&_PoolManager.TransactOpts, initialOwner, _messageManager, _relayerAddress, _withdrawManager)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PoolManager *PoolManagerTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolManager.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PoolManager *PoolManagerSession) Pause() (*types.Transaction, error) {
	return _PoolManager.Contract.Pause(&_PoolManager.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PoolManager *PoolManagerTransactorSession) Pause() (*types.Transaction, error) {
	return _PoolManager.Contract.Pause(&_PoolManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PoolManager *PoolManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PoolManager *PoolManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _PoolManager.Contract.RenounceOwnership(&_PoolManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PoolManager *PoolManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PoolManager.Contract.RenounceOwnership(&_PoolManager.TransactOpts)
}

// SetMinStakeAmount is a paid mutator transaction binding the contract method 0xcb314fab.
//
// Solidity: function setMinStakeAmount(address _token, uint256 _amount) returns()
func (_PoolManager *PoolManagerTransactor) SetMinStakeAmount(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PoolManager.contract.Transact(opts, "setMinStakeAmount", _token, _amount)
}

// SetMinStakeAmount is a paid mutator transaction binding the contract method 0xcb314fab.
//
// Solidity: function setMinStakeAmount(address _token, uint256 _amount) returns()
func (_PoolManager *PoolManagerSession) SetMinStakeAmount(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PoolManager.Contract.SetMinStakeAmount(&_PoolManager.TransactOpts, _token, _amount)
}

// SetMinStakeAmount is a paid mutator transaction binding the contract method 0xcb314fab.
//
// Solidity: function setMinStakeAmount(address _token, uint256 _amount) returns()
func (_PoolManager *PoolManagerTransactorSession) SetMinStakeAmount(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PoolManager.Contract.SetMinStakeAmount(&_PoolManager.TransactOpts, _token, _amount)
}

// SetMinTransferAmount is a paid mutator transaction binding the contract method 0xb92e6396.
//
// Solidity: function setMinTransferAmount(uint256 _MinTransferAmount) returns()
func (_PoolManager *PoolManagerTransactor) SetMinTransferAmount(opts *bind.TransactOpts, _MinTransferAmount *big.Int) (*types.Transaction, error) {
	return _PoolManager.contract.Transact(opts, "setMinTransferAmount", _MinTransferAmount)
}

// SetMinTransferAmount is a paid mutator transaction binding the contract method 0xb92e6396.
//
// Solidity: function setMinTransferAmount(uint256 _MinTransferAmount) returns()
func (_PoolManager *PoolManagerSession) SetMinTransferAmount(_MinTransferAmount *big.Int) (*types.Transaction, error) {
	return _PoolManager.Contract.SetMinTransferAmount(&_PoolManager.TransactOpts, _MinTransferAmount)
}

// SetMinTransferAmount is a paid mutator transaction binding the contract method 0xb92e6396.
//
// Solidity: function setMinTransferAmount(uint256 _MinTransferAmount) returns()
func (_PoolManager *PoolManagerTransactorSession) SetMinTransferAmount(_MinTransferAmount *big.Int) (*types.Transaction, error) {
	return _PoolManager.Contract.SetMinTransferAmount(&_PoolManager.TransactOpts, _MinTransferAmount)
}

// SetPerFee is a paid mutator transaction binding the contract method 0x650c2276.
//
// Solidity: function setPerFee(uint256 _PerFee) returns()
func (_PoolManager *PoolManagerTransactor) SetPerFee(opts *bind.TransactOpts, _PerFee *big.Int) (*types.Transaction, error) {
	return _PoolManager.contract.Transact(opts, "setPerFee", _PerFee)
}

// SetPerFee is a paid mutator transaction binding the contract method 0x650c2276.
//
// Solidity: function setPerFee(uint256 _PerFee) returns()
func (_PoolManager *PoolManagerSession) SetPerFee(_PerFee *big.Int) (*types.Transaction, error) {
	return _PoolManager.Contract.SetPerFee(&_PoolManager.TransactOpts, _PerFee)
}

// SetPerFee is a paid mutator transaction binding the contract method 0x650c2276.
//
// Solidity: function setPerFee(uint256 _PerFee) returns()
func (_PoolManager *PoolManagerTransactorSession) SetPerFee(_PerFee *big.Int) (*types.Transaction, error) {
	return _PoolManager.Contract.SetPerFee(&_PoolManager.TransactOpts, _PerFee)
}

// SetSupportERC20Token is a paid mutator transaction binding the contract method 0xab0f9e19.
//
// Solidity: function setSupportERC20Token(address ERC20Address, bool isValid) returns()
func (_PoolManager *PoolManagerTransactor) SetSupportERC20Token(opts *bind.TransactOpts, ERC20Address common.Address, isValid bool) (*types.Transaction, error) {
	return _PoolManager.contract.Transact(opts, "setSupportERC20Token", ERC20Address, isValid)
}

// SetSupportERC20Token is a paid mutator transaction binding the contract method 0xab0f9e19.
//
// Solidity: function setSupportERC20Token(address ERC20Address, bool isValid) returns()
func (_PoolManager *PoolManagerSession) SetSupportERC20Token(ERC20Address common.Address, isValid bool) (*types.Transaction, error) {
	return _PoolManager.Contract.SetSupportERC20Token(&_PoolManager.TransactOpts, ERC20Address, isValid)
}

// SetSupportERC20Token is a paid mutator transaction binding the contract method 0xab0f9e19.
//
// Solidity: function setSupportERC20Token(address ERC20Address, bool isValid) returns()
func (_PoolManager *PoolManagerTransactorSession) SetSupportERC20Token(ERC20Address common.Address, isValid bool) (*types.Transaction, error) {
	return _PoolManager.Contract.SetSupportERC20Token(&_PoolManager.TransactOpts, ERC20Address, isValid)
}

// SetSupportToken is a paid mutator transaction binding the contract method 0xe2070893.
//
// Solidity: function setSupportToken(address _token, bool _isSupport, uint32 startTimes) returns()
func (_PoolManager *PoolManagerTransactor) SetSupportToken(opts *bind.TransactOpts, _token common.Address, _isSupport bool, startTimes uint32) (*types.Transaction, error) {
	return _PoolManager.contract.Transact(opts, "setSupportToken", _token, _isSupport, startTimes)
}

// SetSupportToken is a paid mutator transaction binding the contract method 0xe2070893.
//
// Solidity: function setSupportToken(address _token, bool _isSupport, uint32 startTimes) returns()
func (_PoolManager *PoolManagerSession) SetSupportToken(_token common.Address, _isSupport bool, startTimes uint32) (*types.Transaction, error) {
	return _PoolManager.Contract.SetSupportToken(&_PoolManager.TransactOpts, _token, _isSupport, startTimes)
}

// SetSupportToken is a paid mutator transaction binding the contract method 0xe2070893.
//
// Solidity: function setSupportToken(address _token, bool _isSupport, uint32 startTimes) returns()
func (_PoolManager *PoolManagerTransactorSession) SetSupportToken(_token common.Address, _isSupport bool, startTimes uint32) (*types.Transaction, error) {
	return _PoolManager.Contract.SetSupportToken(&_PoolManager.TransactOpts, _token, _isSupport, startTimes)
}

// SetValidChainId is a paid mutator transaction binding the contract method 0xd73f14e5.
//
// Solidity: function setValidChainId(uint256 chainId, bool isValid) returns()
func (_PoolManager *PoolManagerTransactor) SetValidChainId(opts *bind.TransactOpts, chainId *big.Int, isValid bool) (*types.Transaction, error) {
	return _PoolManager.contract.Transact(opts, "setValidChainId", chainId, isValid)
}

// SetValidChainId is a paid mutator transaction binding the contract method 0xd73f14e5.
//
// Solidity: function setValidChainId(uint256 chainId, bool isValid) returns()
func (_PoolManager *PoolManagerSession) SetValidChainId(chainId *big.Int, isValid bool) (*types.Transaction, error) {
	return _PoolManager.Contract.SetValidChainId(&_PoolManager.TransactOpts, chainId, isValid)
}

// SetValidChainId is a paid mutator transaction binding the contract method 0xd73f14e5.
//
// Solidity: function setValidChainId(uint256 chainId, bool isValid) returns()
func (_PoolManager *PoolManagerTransactorSession) SetValidChainId(chainId *big.Int, isValid bool) (*types.Transaction, error) {
	return _PoolManager.Contract.SetValidChainId(&_PoolManager.TransactOpts, chainId, isValid)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PoolManager *PoolManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PoolManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PoolManager *PoolManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PoolManager.Contract.TransferOwnership(&_PoolManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PoolManager *PoolManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PoolManager.Contract.TransferOwnership(&_PoolManager.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PoolManager *PoolManagerTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolManager.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PoolManager *PoolManagerSession) Unpause() (*types.Transaction, error) {
	return _PoolManager.Contract.Unpause(&_PoolManager.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PoolManager *PoolManagerTransactorSession) Unpause() (*types.Transaction, error) {
	return _PoolManager.Contract.Unpause(&_PoolManager.TransactOpts)
}

// WithdrawErc20FromBridge is a paid mutator transaction binding the contract method 0x34fa7f70.
//
// Solidity: function withdrawErc20FromBridge(address tokenAddress, address withdrawAddress, uint256 amount) returns(bool)
func (_PoolManager *PoolManagerTransactor) WithdrawErc20FromBridge(opts *bind.TransactOpts, tokenAddress common.Address, withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PoolManager.contract.Transact(opts, "withdrawErc20FromBridge", tokenAddress, withdrawAddress, amount)
}

// WithdrawErc20FromBridge is a paid mutator transaction binding the contract method 0x34fa7f70.
//
// Solidity: function withdrawErc20FromBridge(address tokenAddress, address withdrawAddress, uint256 amount) returns(bool)
func (_PoolManager *PoolManagerSession) WithdrawErc20FromBridge(tokenAddress common.Address, withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PoolManager.Contract.WithdrawErc20FromBridge(&_PoolManager.TransactOpts, tokenAddress, withdrawAddress, amount)
}

// WithdrawErc20FromBridge is a paid mutator transaction binding the contract method 0x34fa7f70.
//
// Solidity: function withdrawErc20FromBridge(address tokenAddress, address withdrawAddress, uint256 amount) returns(bool)
func (_PoolManager *PoolManagerTransactorSession) WithdrawErc20FromBridge(tokenAddress common.Address, withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PoolManager.Contract.WithdrawErc20FromBridge(&_PoolManager.TransactOpts, tokenAddress, withdrawAddress, amount)
}

// WithdrawEthFromBridge is a paid mutator transaction binding the contract method 0x52662af3.
//
// Solidity: function withdrawEthFromBridge(address withdrawAddress, uint256 amount) payable returns(bool)
func (_PoolManager *PoolManagerTransactor) WithdrawEthFromBridge(opts *bind.TransactOpts, withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PoolManager.contract.Transact(opts, "withdrawEthFromBridge", withdrawAddress, amount)
}

// WithdrawEthFromBridge is a paid mutator transaction binding the contract method 0x52662af3.
//
// Solidity: function withdrawEthFromBridge(address withdrawAddress, uint256 amount) payable returns(bool)
func (_PoolManager *PoolManagerSession) WithdrawEthFromBridge(withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PoolManager.Contract.WithdrawEthFromBridge(&_PoolManager.TransactOpts, withdrawAddress, amount)
}

// WithdrawEthFromBridge is a paid mutator transaction binding the contract method 0x52662af3.
//
// Solidity: function withdrawEthFromBridge(address withdrawAddress, uint256 amount) payable returns(bool)
func (_PoolManager *PoolManagerTransactorSession) WithdrawEthFromBridge(withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PoolManager.Contract.WithdrawEthFromBridge(&_PoolManager.TransactOpts, withdrawAddress, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PoolManager *PoolManagerTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolManager.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PoolManager *PoolManagerSession) Receive() (*types.Transaction, error) {
	return _PoolManager.Contract.Receive(&_PoolManager.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PoolManager *PoolManagerTransactorSession) Receive() (*types.Transaction, error) {
	return _PoolManager.Contract.Receive(&_PoolManager.TransactOpts)
}

// PoolManagerClaimRewardIterator is returned from FilterClaimReward and is used to iterate over the raw logs and unpacked data for ClaimReward events raised by the PoolManager contract.
type PoolManagerClaimRewardIterator struct {
	Event *PoolManagerClaimReward // Event containing the contract specifics and raw log

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
func (it *PoolManagerClaimRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolManagerClaimReward)
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
		it.Event = new(PoolManagerClaimReward)
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
func (it *PoolManagerClaimRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolManagerClaimRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolManagerClaimReward represents a ClaimReward event raised by the PoolManager contract.
type PoolManagerClaimReward struct {
	User        common.Address
	StartPoolId *big.Int
	EndPoolId   *big.Int
	ChainId     *big.Int
	Token       common.Address
	Reward      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterClaimReward is a free log retrieval operation binding the contract event 0x36b99ff884a288b9b46a49442e3c4c89464469fae98da480f519664f99fc4a6f.
//
// Solidity: event ClaimReward(address _user, uint256 startPoolId, uint256 EndPoolId, uint256 chainId, address _token, uint256 Reward)
func (_PoolManager *PoolManagerFilterer) FilterClaimReward(opts *bind.FilterOpts) (*PoolManagerClaimRewardIterator, error) {

	logs, sub, err := _PoolManager.contract.FilterLogs(opts, "ClaimReward")
	if err != nil {
		return nil, err
	}
	return &PoolManagerClaimRewardIterator{contract: _PoolManager.contract, event: "ClaimReward", logs: logs, sub: sub}, nil
}

// WatchClaimReward is a free log subscription operation binding the contract event 0x36b99ff884a288b9b46a49442e3c4c89464469fae98da480f519664f99fc4a6f.
//
// Solidity: event ClaimReward(address _user, uint256 startPoolId, uint256 EndPoolId, uint256 chainId, address _token, uint256 Reward)
func (_PoolManager *PoolManagerFilterer) WatchClaimReward(opts *bind.WatchOpts, sink chan<- *PoolManagerClaimReward) (event.Subscription, error) {

	logs, sub, err := _PoolManager.contract.WatchLogs(opts, "ClaimReward")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolManagerClaimReward)
				if err := _PoolManager.contract.UnpackLog(event, "ClaimReward", log); err != nil {
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

// ParseClaimReward is a log parse operation binding the contract event 0x36b99ff884a288b9b46a49442e3c4c89464469fae98da480f519664f99fc4a6f.
//
// Solidity: event ClaimReward(address _user, uint256 startPoolId, uint256 EndPoolId, uint256 chainId, address _token, uint256 Reward)
func (_PoolManager *PoolManagerFilterer) ParseClaimReward(log types.Log) (*PoolManagerClaimReward, error) {
	event := new(PoolManagerClaimReward)
	if err := _PoolManager.contract.UnpackLog(event, "ClaimReward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolManagerCompletePoolEventIterator is returned from FilterCompletePoolEvent and is used to iterate over the raw logs and unpacked data for CompletePoolEvent events raised by the PoolManager contract.
type PoolManagerCompletePoolEventIterator struct {
	Event *PoolManagerCompletePoolEvent // Event containing the contract specifics and raw log

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
func (it *PoolManagerCompletePoolEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolManagerCompletePoolEvent)
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
		it.Event = new(PoolManagerCompletePoolEvent)
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
func (it *PoolManagerCompletePoolEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolManagerCompletePoolEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolManagerCompletePoolEvent represents a CompletePoolEvent event raised by the PoolManager contract.
type PoolManagerCompletePoolEvent struct {
	Token     common.Address
	PoolIndex *big.Int
	ChainId   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCompletePoolEvent is a free log retrieval operation binding the contract event 0xe31e15753cd173232924147d9c429fccfb91705d3778d5eb2735a5d302ea5e81.
//
// Solidity: event CompletePoolEvent(address indexed token, uint256 poolIndex, uint256 chainId)
func (_PoolManager *PoolManagerFilterer) FilterCompletePoolEvent(opts *bind.FilterOpts, token []common.Address) (*PoolManagerCompletePoolEventIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PoolManager.contract.FilterLogs(opts, "CompletePoolEvent", tokenRule)
	if err != nil {
		return nil, err
	}
	return &PoolManagerCompletePoolEventIterator{contract: _PoolManager.contract, event: "CompletePoolEvent", logs: logs, sub: sub}, nil
}

// WatchCompletePoolEvent is a free log subscription operation binding the contract event 0xe31e15753cd173232924147d9c429fccfb91705d3778d5eb2735a5d302ea5e81.
//
// Solidity: event CompletePoolEvent(address indexed token, uint256 poolIndex, uint256 chainId)
func (_PoolManager *PoolManagerFilterer) WatchCompletePoolEvent(opts *bind.WatchOpts, sink chan<- *PoolManagerCompletePoolEvent, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PoolManager.contract.WatchLogs(opts, "CompletePoolEvent", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolManagerCompletePoolEvent)
				if err := _PoolManager.contract.UnpackLog(event, "CompletePoolEvent", log); err != nil {
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

// ParseCompletePoolEvent is a log parse operation binding the contract event 0xe31e15753cd173232924147d9c429fccfb91705d3778d5eb2735a5d302ea5e81.
//
// Solidity: event CompletePoolEvent(address indexed token, uint256 poolIndex, uint256 chainId)
func (_PoolManager *PoolManagerFilterer) ParseCompletePoolEvent(log types.Log) (*PoolManagerCompletePoolEvent, error) {
	event := new(PoolManagerCompletePoolEvent)
	if err := _PoolManager.contract.UnpackLog(event, "CompletePoolEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolManagerDepositTokenIterator is returned from FilterDepositToken and is used to iterate over the raw logs and unpacked data for DepositToken events raised by the PoolManager contract.
type PoolManagerDepositTokenIterator struct {
	Event *PoolManagerDepositToken // Event containing the contract specifics and raw log

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
func (it *PoolManagerDepositTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolManagerDepositToken)
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
		it.Event = new(PoolManagerDepositToken)
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
func (it *PoolManagerDepositTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolManagerDepositTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolManagerDepositToken represents a DepositToken event raised by the PoolManager contract.
type PoolManagerDepositToken struct {
	TokenAddress common.Address
	Sender       common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDepositToken is a free log retrieval operation binding the contract event 0x4b3f81827ede20c81afbf1bb77b954afcdcae24d391d99042310cb1d9210dd57.
//
// Solidity: event DepositToken(address indexed tokenAddress, address indexed sender, uint256 amount)
func (_PoolManager *PoolManagerFilterer) FilterDepositToken(opts *bind.FilterOpts, tokenAddress []common.Address, sender []common.Address) (*PoolManagerDepositTokenIterator, error) {

	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _PoolManager.contract.FilterLogs(opts, "DepositToken", tokenAddressRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PoolManagerDepositTokenIterator{contract: _PoolManager.contract, event: "DepositToken", logs: logs, sub: sub}, nil
}

// WatchDepositToken is a free log subscription operation binding the contract event 0x4b3f81827ede20c81afbf1bb77b954afcdcae24d391d99042310cb1d9210dd57.
//
// Solidity: event DepositToken(address indexed tokenAddress, address indexed sender, uint256 amount)
func (_PoolManager *PoolManagerFilterer) WatchDepositToken(opts *bind.WatchOpts, sink chan<- *PoolManagerDepositToken, tokenAddress []common.Address, sender []common.Address) (event.Subscription, error) {

	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _PoolManager.contract.WatchLogs(opts, "DepositToken", tokenAddressRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolManagerDepositToken)
				if err := _PoolManager.contract.UnpackLog(event, "DepositToken", log); err != nil {
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

// ParseDepositToken is a log parse operation binding the contract event 0x4b3f81827ede20c81afbf1bb77b954afcdcae24d391d99042310cb1d9210dd57.
//
// Solidity: event DepositToken(address indexed tokenAddress, address indexed sender, uint256 amount)
func (_PoolManager *PoolManagerFilterer) ParseDepositToken(log types.Log) (*PoolManagerDepositToken, error) {
	event := new(PoolManagerDepositToken)
	if err := _PoolManager.contract.UnpackLog(event, "DepositToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolManagerFinalizeERC20Iterator is returned from FilterFinalizeERC20 and is used to iterate over the raw logs and unpacked data for FinalizeERC20 events raised by the PoolManager contract.
type PoolManagerFinalizeERC20Iterator struct {
	Event *PoolManagerFinalizeERC20 // Event containing the contract specifics and raw log

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
func (it *PoolManagerFinalizeERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolManagerFinalizeERC20)
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
		it.Event = new(PoolManagerFinalizeERC20)
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
func (it *PoolManagerFinalizeERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolManagerFinalizeERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolManagerFinalizeERC20 represents a FinalizeERC20 event raised by the PoolManager contract.
type PoolManagerFinalizeERC20 struct {
	SourceChainId      *big.Int
	DestChainId        *big.Int
	SourceTokenAddress common.Address
	DestTokenAddress   common.Address
	From               common.Address
	To                 common.Address
	Value              *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterFinalizeERC20 is a free log retrieval operation binding the contract event 0x7add2d08613c9e66b48e0cd9188091059a1eeeceaeed10d42cb1720aa8dec767.
//
// Solidity: event FinalizeERC20(uint256 sourceChainId, uint256 destChainId, address sourceTokenAddress, address destTokenAddress, address indexed from, address indexed to, uint256 value)
func (_PoolManager *PoolManagerFilterer) FilterFinalizeERC20(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*PoolManagerFinalizeERC20Iterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PoolManager.contract.FilterLogs(opts, "FinalizeERC20", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PoolManagerFinalizeERC20Iterator{contract: _PoolManager.contract, event: "FinalizeERC20", logs: logs, sub: sub}, nil
}

// WatchFinalizeERC20 is a free log subscription operation binding the contract event 0x7add2d08613c9e66b48e0cd9188091059a1eeeceaeed10d42cb1720aa8dec767.
//
// Solidity: event FinalizeERC20(uint256 sourceChainId, uint256 destChainId, address sourceTokenAddress, address destTokenAddress, address indexed from, address indexed to, uint256 value)
func (_PoolManager *PoolManagerFilterer) WatchFinalizeERC20(opts *bind.WatchOpts, sink chan<- *PoolManagerFinalizeERC20, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PoolManager.contract.WatchLogs(opts, "FinalizeERC20", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolManagerFinalizeERC20)
				if err := _PoolManager.contract.UnpackLog(event, "FinalizeERC20", log); err != nil {
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

// ParseFinalizeERC20 is a log parse operation binding the contract event 0x7add2d08613c9e66b48e0cd9188091059a1eeeceaeed10d42cb1720aa8dec767.
//
// Solidity: event FinalizeERC20(uint256 sourceChainId, uint256 destChainId, address sourceTokenAddress, address destTokenAddress, address indexed from, address indexed to, uint256 value)
func (_PoolManager *PoolManagerFilterer) ParseFinalizeERC20(log types.Log) (*PoolManagerFinalizeERC20, error) {
	event := new(PoolManagerFinalizeERC20)
	if err := _PoolManager.contract.UnpackLog(event, "FinalizeERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolManagerFinalizeETHIterator is returned from FilterFinalizeETH and is used to iterate over the raw logs and unpacked data for FinalizeETH events raised by the PoolManager contract.
type PoolManagerFinalizeETHIterator struct {
	Event *PoolManagerFinalizeETH // Event containing the contract specifics and raw log

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
func (it *PoolManagerFinalizeETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolManagerFinalizeETH)
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
		it.Event = new(PoolManagerFinalizeETH)
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
func (it *PoolManagerFinalizeETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolManagerFinalizeETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolManagerFinalizeETH represents a FinalizeETH event raised by the PoolManager contract.
type PoolManagerFinalizeETH struct {
	SourceChainId      *big.Int
	DestChainId        *big.Int
	SourceTokenAddress common.Address
	From               common.Address
	To                 common.Address
	Value              *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterFinalizeETH is a free log retrieval operation binding the contract event 0xeb0a2e1a2d528efe194dccba78e1f08be6e84e330206377d56c2a2fb492e5df7.
//
// Solidity: event FinalizeETH(uint256 sourceChainId, uint256 destChainId, address sourceTokenAddress, address indexed from, address indexed to, uint256 value)
func (_PoolManager *PoolManagerFilterer) FilterFinalizeETH(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*PoolManagerFinalizeETHIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PoolManager.contract.FilterLogs(opts, "FinalizeETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PoolManagerFinalizeETHIterator{contract: _PoolManager.contract, event: "FinalizeETH", logs: logs, sub: sub}, nil
}

// WatchFinalizeETH is a free log subscription operation binding the contract event 0xeb0a2e1a2d528efe194dccba78e1f08be6e84e330206377d56c2a2fb492e5df7.
//
// Solidity: event FinalizeETH(uint256 sourceChainId, uint256 destChainId, address sourceTokenAddress, address indexed from, address indexed to, uint256 value)
func (_PoolManager *PoolManagerFilterer) WatchFinalizeETH(opts *bind.WatchOpts, sink chan<- *PoolManagerFinalizeETH, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PoolManager.contract.WatchLogs(opts, "FinalizeETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolManagerFinalizeETH)
				if err := _PoolManager.contract.UnpackLog(event, "FinalizeETH", log); err != nil {
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

// ParseFinalizeETH is a log parse operation binding the contract event 0xeb0a2e1a2d528efe194dccba78e1f08be6e84e330206377d56c2a2fb492e5df7.
//
// Solidity: event FinalizeETH(uint256 sourceChainId, uint256 destChainId, address sourceTokenAddress, address indexed from, address indexed to, uint256 value)
func (_PoolManager *PoolManagerFilterer) ParseFinalizeETH(log types.Log) (*PoolManagerFinalizeETH, error) {
	event := new(PoolManagerFinalizeETH)
	if err := _PoolManager.contract.UnpackLog(event, "FinalizeETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the PoolManager contract.
type PoolManagerInitializedIterator struct {
	Event *PoolManagerInitialized // Event containing the contract specifics and raw log

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
func (it *PoolManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolManagerInitialized)
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
		it.Event = new(PoolManagerInitialized)
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
func (it *PoolManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolManagerInitialized represents a Initialized event raised by the PoolManager contract.
type PoolManagerInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_PoolManager *PoolManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*PoolManagerInitializedIterator, error) {

	logs, sub, err := _PoolManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PoolManagerInitializedIterator{contract: _PoolManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_PoolManager *PoolManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PoolManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _PoolManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolManagerInitialized)
				if err := _PoolManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_PoolManager *PoolManagerFilterer) ParseInitialized(log types.Log) (*PoolManagerInitialized, error) {
	event := new(PoolManagerInitialized)
	if err := _PoolManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolManagerInitiateERC20Iterator is returned from FilterInitiateERC20 and is used to iterate over the raw logs and unpacked data for InitiateERC20 events raised by the PoolManager contract.
type PoolManagerInitiateERC20Iterator struct {
	Event *PoolManagerInitiateERC20 // Event containing the contract specifics and raw log

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
func (it *PoolManagerInitiateERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolManagerInitiateERC20)
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
		it.Event = new(PoolManagerInitiateERC20)
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
func (it *PoolManagerInitiateERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolManagerInitiateERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolManagerInitiateERC20 represents a InitiateERC20 event raised by the PoolManager contract.
type PoolManagerInitiateERC20 struct {
	SourceChainId      *big.Int
	DestChainId        *big.Int
	SourceTokenAddress common.Address
	DestTokenAddress   common.Address
	From               common.Address
	To                 common.Address
	Value              *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterInitiateERC20 is a free log retrieval operation binding the contract event 0x24136187bad583830e5c55416af2a551e1750b2b95b0e997565fba1badf10ed3.
//
// Solidity: event InitiateERC20(uint256 sourceChainId, uint256 destChainId, address sourceTokenAddress, address destTokenAddress, address indexed from, address indexed to, uint256 value)
func (_PoolManager *PoolManagerFilterer) FilterInitiateERC20(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*PoolManagerInitiateERC20Iterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PoolManager.contract.FilterLogs(opts, "InitiateERC20", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PoolManagerInitiateERC20Iterator{contract: _PoolManager.contract, event: "InitiateERC20", logs: logs, sub: sub}, nil
}

// WatchInitiateERC20 is a free log subscription operation binding the contract event 0x24136187bad583830e5c55416af2a551e1750b2b95b0e997565fba1badf10ed3.
//
// Solidity: event InitiateERC20(uint256 sourceChainId, uint256 destChainId, address sourceTokenAddress, address destTokenAddress, address indexed from, address indexed to, uint256 value)
func (_PoolManager *PoolManagerFilterer) WatchInitiateERC20(opts *bind.WatchOpts, sink chan<- *PoolManagerInitiateERC20, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PoolManager.contract.WatchLogs(opts, "InitiateERC20", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolManagerInitiateERC20)
				if err := _PoolManager.contract.UnpackLog(event, "InitiateERC20", log); err != nil {
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

// ParseInitiateERC20 is a log parse operation binding the contract event 0x24136187bad583830e5c55416af2a551e1750b2b95b0e997565fba1badf10ed3.
//
// Solidity: event InitiateERC20(uint256 sourceChainId, uint256 destChainId, address sourceTokenAddress, address destTokenAddress, address indexed from, address indexed to, uint256 value)
func (_PoolManager *PoolManagerFilterer) ParseInitiateERC20(log types.Log) (*PoolManagerInitiateERC20, error) {
	event := new(PoolManagerInitiateERC20)
	if err := _PoolManager.contract.UnpackLog(event, "InitiateERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolManagerInitiateETHIterator is returned from FilterInitiateETH and is used to iterate over the raw logs and unpacked data for InitiateETH events raised by the PoolManager contract.
type PoolManagerInitiateETHIterator struct {
	Event *PoolManagerInitiateETH // Event containing the contract specifics and raw log

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
func (it *PoolManagerInitiateETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolManagerInitiateETH)
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
		it.Event = new(PoolManagerInitiateETH)
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
func (it *PoolManagerInitiateETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolManagerInitiateETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolManagerInitiateETH represents a InitiateETH event raised by the PoolManager contract.
type PoolManagerInitiateETH struct {
	SourceChainId    *big.Int
	DestChainId      *big.Int
	DestTokenAddress common.Address
	From             common.Address
	To               common.Address
	Value            *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterInitiateETH is a free log retrieval operation binding the contract event 0x0c2b5a57887e542335dea57bfc17d53dc16f67d219e0cc41ec1970283524abcc.
//
// Solidity: event InitiateETH(uint256 sourceChainId, uint256 destChainId, address destTokenAddress, address indexed from, address indexed to, uint256 value)
func (_PoolManager *PoolManagerFilterer) FilterInitiateETH(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*PoolManagerInitiateETHIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PoolManager.contract.FilterLogs(opts, "InitiateETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PoolManagerInitiateETHIterator{contract: _PoolManager.contract, event: "InitiateETH", logs: logs, sub: sub}, nil
}

// WatchInitiateETH is a free log subscription operation binding the contract event 0x0c2b5a57887e542335dea57bfc17d53dc16f67d219e0cc41ec1970283524abcc.
//
// Solidity: event InitiateETH(uint256 sourceChainId, uint256 destChainId, address destTokenAddress, address indexed from, address indexed to, uint256 value)
func (_PoolManager *PoolManagerFilterer) WatchInitiateETH(opts *bind.WatchOpts, sink chan<- *PoolManagerInitiateETH, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PoolManager.contract.WatchLogs(opts, "InitiateETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolManagerInitiateETH)
				if err := _PoolManager.contract.UnpackLog(event, "InitiateETH", log); err != nil {
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

// ParseInitiateETH is a log parse operation binding the contract event 0x0c2b5a57887e542335dea57bfc17d53dc16f67d219e0cc41ec1970283524abcc.
//
// Solidity: event InitiateETH(uint256 sourceChainId, uint256 destChainId, address destTokenAddress, address indexed from, address indexed to, uint256 value)
func (_PoolManager *PoolManagerFilterer) ParseInitiateETH(log types.Log) (*PoolManagerInitiateETH, error) {
	event := new(PoolManagerInitiateETH)
	if err := _PoolManager.contract.UnpackLog(event, "InitiateETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PoolManager contract.
type PoolManagerOwnershipTransferredIterator struct {
	Event *PoolManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PoolManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolManagerOwnershipTransferred)
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
		it.Event = new(PoolManagerOwnershipTransferred)
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
func (it *PoolManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolManagerOwnershipTransferred represents a OwnershipTransferred event raised by the PoolManager contract.
type PoolManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PoolManager *PoolManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PoolManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PoolManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PoolManagerOwnershipTransferredIterator{contract: _PoolManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PoolManager *PoolManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PoolManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PoolManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolManagerOwnershipTransferred)
				if err := _PoolManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_PoolManager *PoolManagerFilterer) ParseOwnershipTransferred(log types.Log) (*PoolManagerOwnershipTransferred, error) {
	event := new(PoolManagerOwnershipTransferred)
	if err := _PoolManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the PoolManager contract.
type PoolManagerPausedIterator struct {
	Event *PoolManagerPaused // Event containing the contract specifics and raw log

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
func (it *PoolManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolManagerPaused)
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
		it.Event = new(PoolManagerPaused)
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
func (it *PoolManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolManagerPaused represents a Paused event raised by the PoolManager contract.
type PoolManagerPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PoolManager *PoolManagerFilterer) FilterPaused(opts *bind.FilterOpts) (*PoolManagerPausedIterator, error) {

	logs, sub, err := _PoolManager.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &PoolManagerPausedIterator{contract: _PoolManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PoolManager *PoolManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PoolManagerPaused) (event.Subscription, error) {

	logs, sub, err := _PoolManager.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolManagerPaused)
				if err := _PoolManager.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PoolManager *PoolManagerFilterer) ParsePaused(log types.Log) (*PoolManagerPaused, error) {
	event := new(PoolManagerPaused)
	if err := _PoolManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolManagerSetMinStakeAmountEventIterator is returned from FilterSetMinStakeAmountEvent and is used to iterate over the raw logs and unpacked data for SetMinStakeAmountEvent events raised by the PoolManager contract.
type PoolManagerSetMinStakeAmountEventIterator struct {
	Event *PoolManagerSetMinStakeAmountEvent // Event containing the contract specifics and raw log

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
func (it *PoolManagerSetMinStakeAmountEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolManagerSetMinStakeAmountEvent)
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
		it.Event = new(PoolManagerSetMinStakeAmountEvent)
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
func (it *PoolManagerSetMinStakeAmountEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolManagerSetMinStakeAmountEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolManagerSetMinStakeAmountEvent represents a SetMinStakeAmountEvent event raised by the PoolManager contract.
type PoolManagerSetMinStakeAmountEvent struct {
	Token   common.Address
	Amount  *big.Int
	ChainId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSetMinStakeAmountEvent is a free log retrieval operation binding the contract event 0x6559cfa4c9d2d533663fa9b68e86d2bb216ac0d2bf462b0529ab01525dce0837.
//
// Solidity: event SetMinStakeAmountEvent(address indexed token, uint256 amount, uint256 chainId)
func (_PoolManager *PoolManagerFilterer) FilterSetMinStakeAmountEvent(opts *bind.FilterOpts, token []common.Address) (*PoolManagerSetMinStakeAmountEventIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PoolManager.contract.FilterLogs(opts, "SetMinStakeAmountEvent", tokenRule)
	if err != nil {
		return nil, err
	}
	return &PoolManagerSetMinStakeAmountEventIterator{contract: _PoolManager.contract, event: "SetMinStakeAmountEvent", logs: logs, sub: sub}, nil
}

// WatchSetMinStakeAmountEvent is a free log subscription operation binding the contract event 0x6559cfa4c9d2d533663fa9b68e86d2bb216ac0d2bf462b0529ab01525dce0837.
//
// Solidity: event SetMinStakeAmountEvent(address indexed token, uint256 amount, uint256 chainId)
func (_PoolManager *PoolManagerFilterer) WatchSetMinStakeAmountEvent(opts *bind.WatchOpts, sink chan<- *PoolManagerSetMinStakeAmountEvent, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PoolManager.contract.WatchLogs(opts, "SetMinStakeAmountEvent", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolManagerSetMinStakeAmountEvent)
				if err := _PoolManager.contract.UnpackLog(event, "SetMinStakeAmountEvent", log); err != nil {
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

// ParseSetMinStakeAmountEvent is a log parse operation binding the contract event 0x6559cfa4c9d2d533663fa9b68e86d2bb216ac0d2bf462b0529ab01525dce0837.
//
// Solidity: event SetMinStakeAmountEvent(address indexed token, uint256 amount, uint256 chainId)
func (_PoolManager *PoolManagerFilterer) ParseSetMinStakeAmountEvent(log types.Log) (*PoolManagerSetMinStakeAmountEvent, error) {
	event := new(PoolManagerSetMinStakeAmountEvent)
	if err := _PoolManager.contract.UnpackLog(event, "SetMinStakeAmountEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolManagerSetMinTransferAmountIterator is returned from FilterSetMinTransferAmount and is used to iterate over the raw logs and unpacked data for SetMinTransferAmount events raised by the PoolManager contract.
type PoolManagerSetMinTransferAmountIterator struct {
	Event *PoolManagerSetMinTransferAmount // Event containing the contract specifics and raw log

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
func (it *PoolManagerSetMinTransferAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolManagerSetMinTransferAmount)
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
		it.Event = new(PoolManagerSetMinTransferAmount)
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
func (it *PoolManagerSetMinTransferAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolManagerSetMinTransferAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolManagerSetMinTransferAmount represents a SetMinTransferAmount event raised by the PoolManager contract.
type PoolManagerSetMinTransferAmount struct {
	MinTransferAmount *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSetMinTransferAmount is a free log retrieval operation binding the contract event 0x39315f035b4e62ef103e9d2372379ecc1fda660b2b8bf306693ff8a07ed1dd3f.
//
// Solidity: event SetMinTransferAmount(uint256 _MinTransferAmount)
func (_PoolManager *PoolManagerFilterer) FilterSetMinTransferAmount(opts *bind.FilterOpts) (*PoolManagerSetMinTransferAmountIterator, error) {

	logs, sub, err := _PoolManager.contract.FilterLogs(opts, "SetMinTransferAmount")
	if err != nil {
		return nil, err
	}
	return &PoolManagerSetMinTransferAmountIterator{contract: _PoolManager.contract, event: "SetMinTransferAmount", logs: logs, sub: sub}, nil
}

// WatchSetMinTransferAmount is a free log subscription operation binding the contract event 0x39315f035b4e62ef103e9d2372379ecc1fda660b2b8bf306693ff8a07ed1dd3f.
//
// Solidity: event SetMinTransferAmount(uint256 _MinTransferAmount)
func (_PoolManager *PoolManagerFilterer) WatchSetMinTransferAmount(opts *bind.WatchOpts, sink chan<- *PoolManagerSetMinTransferAmount) (event.Subscription, error) {

	logs, sub, err := _PoolManager.contract.WatchLogs(opts, "SetMinTransferAmount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolManagerSetMinTransferAmount)
				if err := _PoolManager.contract.UnpackLog(event, "SetMinTransferAmount", log); err != nil {
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

// ParseSetMinTransferAmount is a log parse operation binding the contract event 0x39315f035b4e62ef103e9d2372379ecc1fda660b2b8bf306693ff8a07ed1dd3f.
//
// Solidity: event SetMinTransferAmount(uint256 _MinTransferAmount)
func (_PoolManager *PoolManagerFilterer) ParseSetMinTransferAmount(log types.Log) (*PoolManagerSetMinTransferAmount, error) {
	event := new(PoolManagerSetMinTransferAmount)
	if err := _PoolManager.contract.UnpackLog(event, "SetMinTransferAmount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolManagerSetPerFeeIterator is returned from FilterSetPerFee and is used to iterate over the raw logs and unpacked data for SetPerFee events raised by the PoolManager contract.
type PoolManagerSetPerFeeIterator struct {
	Event *PoolManagerSetPerFee // Event containing the contract specifics and raw log

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
func (it *PoolManagerSetPerFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolManagerSetPerFee)
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
		it.Event = new(PoolManagerSetPerFee)
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
func (it *PoolManagerSetPerFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolManagerSetPerFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolManagerSetPerFee represents a SetPerFee event raised by the PoolManager contract.
type PoolManagerSetPerFee struct {
	ChainId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSetPerFee is a free log retrieval operation binding the contract event 0x342f865a411cdd799cc2fde1e27981fc427f5c8a31c79066a6bcc0db11729da0.
//
// Solidity: event SetPerFee(uint256 chainId)
func (_PoolManager *PoolManagerFilterer) FilterSetPerFee(opts *bind.FilterOpts) (*PoolManagerSetPerFeeIterator, error) {

	logs, sub, err := _PoolManager.contract.FilterLogs(opts, "SetPerFee")
	if err != nil {
		return nil, err
	}
	return &PoolManagerSetPerFeeIterator{contract: _PoolManager.contract, event: "SetPerFee", logs: logs, sub: sub}, nil
}

// WatchSetPerFee is a free log subscription operation binding the contract event 0x342f865a411cdd799cc2fde1e27981fc427f5c8a31c79066a6bcc0db11729da0.
//
// Solidity: event SetPerFee(uint256 chainId)
func (_PoolManager *PoolManagerFilterer) WatchSetPerFee(opts *bind.WatchOpts, sink chan<- *PoolManagerSetPerFee) (event.Subscription, error) {

	logs, sub, err := _PoolManager.contract.WatchLogs(opts, "SetPerFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolManagerSetPerFee)
				if err := _PoolManager.contract.UnpackLog(event, "SetPerFee", log); err != nil {
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

// ParseSetPerFee is a log parse operation binding the contract event 0x342f865a411cdd799cc2fde1e27981fc427f5c8a31c79066a6bcc0db11729da0.
//
// Solidity: event SetPerFee(uint256 chainId)
func (_PoolManager *PoolManagerFilterer) ParseSetPerFee(log types.Log) (*PoolManagerSetPerFee, error) {
	event := new(PoolManagerSetPerFee)
	if err := _PoolManager.contract.UnpackLog(event, "SetPerFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolManagerSetSupportTokenEventIterator is returned from FilterSetSupportTokenEvent and is used to iterate over the raw logs and unpacked data for SetSupportTokenEvent events raised by the PoolManager contract.
type PoolManagerSetSupportTokenEventIterator struct {
	Event *PoolManagerSetSupportTokenEvent // Event containing the contract specifics and raw log

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
func (it *PoolManagerSetSupportTokenEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolManagerSetSupportTokenEvent)
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
		it.Event = new(PoolManagerSetSupportTokenEvent)
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
func (it *PoolManagerSetSupportTokenEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolManagerSetSupportTokenEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolManagerSetSupportTokenEvent represents a SetSupportTokenEvent event raised by the PoolManager contract.
type PoolManagerSetSupportTokenEvent struct {
	Token     common.Address
	IsSupport bool
	ChainId   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSetSupportTokenEvent is a free log retrieval operation binding the contract event 0xf37cdf0b5a2a780dfbdbd1bb6f8a3bcdf62c8941ba001debfa56dfd8eab89b6c.
//
// Solidity: event SetSupportTokenEvent(address indexed token, bool isSupport, uint256 chainId)
func (_PoolManager *PoolManagerFilterer) FilterSetSupportTokenEvent(opts *bind.FilterOpts, token []common.Address) (*PoolManagerSetSupportTokenEventIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PoolManager.contract.FilterLogs(opts, "SetSupportTokenEvent", tokenRule)
	if err != nil {
		return nil, err
	}
	return &PoolManagerSetSupportTokenEventIterator{contract: _PoolManager.contract, event: "SetSupportTokenEvent", logs: logs, sub: sub}, nil
}

// WatchSetSupportTokenEvent is a free log subscription operation binding the contract event 0xf37cdf0b5a2a780dfbdbd1bb6f8a3bcdf62c8941ba001debfa56dfd8eab89b6c.
//
// Solidity: event SetSupportTokenEvent(address indexed token, bool isSupport, uint256 chainId)
func (_PoolManager *PoolManagerFilterer) WatchSetSupportTokenEvent(opts *bind.WatchOpts, sink chan<- *PoolManagerSetSupportTokenEvent, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PoolManager.contract.WatchLogs(opts, "SetSupportTokenEvent", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolManagerSetSupportTokenEvent)
				if err := _PoolManager.contract.UnpackLog(event, "SetSupportTokenEvent", log); err != nil {
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

// ParseSetSupportTokenEvent is a log parse operation binding the contract event 0xf37cdf0b5a2a780dfbdbd1bb6f8a3bcdf62c8941ba001debfa56dfd8eab89b6c.
//
// Solidity: event SetSupportTokenEvent(address indexed token, bool isSupport, uint256 chainId)
func (_PoolManager *PoolManagerFilterer) ParseSetSupportTokenEvent(log types.Log) (*PoolManagerSetSupportTokenEvent, error) {
	event := new(PoolManagerSetSupportTokenEvent)
	if err := _PoolManager.contract.UnpackLog(event, "SetSupportTokenEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolManagerSetValidChainIdIterator is returned from FilterSetValidChainId and is used to iterate over the raw logs and unpacked data for SetValidChainId events raised by the PoolManager contract.
type PoolManagerSetValidChainIdIterator struct {
	Event *PoolManagerSetValidChainId // Event containing the contract specifics and raw log

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
func (it *PoolManagerSetValidChainIdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolManagerSetValidChainId)
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
		it.Event = new(PoolManagerSetValidChainId)
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
func (it *PoolManagerSetValidChainIdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolManagerSetValidChainIdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolManagerSetValidChainId represents a SetValidChainId event raised by the PoolManager contract.
type PoolManagerSetValidChainId struct {
	ChainId *big.Int
	IsValid bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSetValidChainId is a free log retrieval operation binding the contract event 0x756e8389496652aad1bd6bd9ea1e81b5b53341a798eb468701bb68403594c950.
//
// Solidity: event SetValidChainId(uint256 chainId, bool isValid)
func (_PoolManager *PoolManagerFilterer) FilterSetValidChainId(opts *bind.FilterOpts) (*PoolManagerSetValidChainIdIterator, error) {

	logs, sub, err := _PoolManager.contract.FilterLogs(opts, "SetValidChainId")
	if err != nil {
		return nil, err
	}
	return &PoolManagerSetValidChainIdIterator{contract: _PoolManager.contract, event: "SetValidChainId", logs: logs, sub: sub}, nil
}

// WatchSetValidChainId is a free log subscription operation binding the contract event 0x756e8389496652aad1bd6bd9ea1e81b5b53341a798eb468701bb68403594c950.
//
// Solidity: event SetValidChainId(uint256 chainId, bool isValid)
func (_PoolManager *PoolManagerFilterer) WatchSetValidChainId(opts *bind.WatchOpts, sink chan<- *PoolManagerSetValidChainId) (event.Subscription, error) {

	logs, sub, err := _PoolManager.contract.WatchLogs(opts, "SetValidChainId")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolManagerSetValidChainId)
				if err := _PoolManager.contract.UnpackLog(event, "SetValidChainId", log); err != nil {
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

// ParseSetValidChainId is a log parse operation binding the contract event 0x756e8389496652aad1bd6bd9ea1e81b5b53341a798eb468701bb68403594c950.
//
// Solidity: event SetValidChainId(uint256 chainId, bool isValid)
func (_PoolManager *PoolManagerFilterer) ParseSetValidChainId(log types.Log) (*PoolManagerSetValidChainId, error) {
	event := new(PoolManagerSetValidChainId)
	if err := _PoolManager.contract.UnpackLog(event, "SetValidChainId", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolManagerStakingERC20EventIterator is returned from FilterStakingERC20Event and is used to iterate over the raw logs and unpacked data for StakingERC20Event events raised by the PoolManager contract.
type PoolManagerStakingERC20EventIterator struct {
	Event *PoolManagerStakingERC20Event // Event containing the contract specifics and raw log

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
func (it *PoolManagerStakingERC20EventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolManagerStakingERC20Event)
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
		it.Event = new(PoolManagerStakingERC20Event)
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
func (it *PoolManagerStakingERC20EventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolManagerStakingERC20EventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolManagerStakingERC20Event represents a StakingERC20Event event raised by the PoolManager contract.
type PoolManagerStakingERC20Event struct {
	User    common.Address
	Token   common.Address
	ChainId *big.Int
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterStakingERC20Event is a free log retrieval operation binding the contract event 0x1a9bb68de16395364f4a0d198852c1b2c4d1e8e5a628e5fd273284a44cd8a6db.
//
// Solidity: event StakingERC20Event(address indexed user, address indexed token, uint256 chainId, uint256 amount)
func (_PoolManager *PoolManagerFilterer) FilterStakingERC20Event(opts *bind.FilterOpts, user []common.Address, token []common.Address) (*PoolManagerStakingERC20EventIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PoolManager.contract.FilterLogs(opts, "StakingERC20Event", userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &PoolManagerStakingERC20EventIterator{contract: _PoolManager.contract, event: "StakingERC20Event", logs: logs, sub: sub}, nil
}

// WatchStakingERC20Event is a free log subscription operation binding the contract event 0x1a9bb68de16395364f4a0d198852c1b2c4d1e8e5a628e5fd273284a44cd8a6db.
//
// Solidity: event StakingERC20Event(address indexed user, address indexed token, uint256 chainId, uint256 amount)
func (_PoolManager *PoolManagerFilterer) WatchStakingERC20Event(opts *bind.WatchOpts, sink chan<- *PoolManagerStakingERC20Event, user []common.Address, token []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PoolManager.contract.WatchLogs(opts, "StakingERC20Event", userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolManagerStakingERC20Event)
				if err := _PoolManager.contract.UnpackLog(event, "StakingERC20Event", log); err != nil {
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

// ParseStakingERC20Event is a log parse operation binding the contract event 0x1a9bb68de16395364f4a0d198852c1b2c4d1e8e5a628e5fd273284a44cd8a6db.
//
// Solidity: event StakingERC20Event(address indexed user, address indexed token, uint256 chainId, uint256 amount)
func (_PoolManager *PoolManagerFilterer) ParseStakingERC20Event(log types.Log) (*PoolManagerStakingERC20Event, error) {
	event := new(PoolManagerStakingERC20Event)
	if err := _PoolManager.contract.UnpackLog(event, "StakingERC20Event", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolManagerStakingETHEventIterator is returned from FilterStakingETHEvent and is used to iterate over the raw logs and unpacked data for StakingETHEvent events raised by the PoolManager contract.
type PoolManagerStakingETHEventIterator struct {
	Event *PoolManagerStakingETHEvent // Event containing the contract specifics and raw log

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
func (it *PoolManagerStakingETHEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolManagerStakingETHEvent)
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
		it.Event = new(PoolManagerStakingETHEvent)
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
func (it *PoolManagerStakingETHEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolManagerStakingETHEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolManagerStakingETHEvent represents a StakingETHEvent event raised by the PoolManager contract.
type PoolManagerStakingETHEvent struct {
	User    common.Address
	ChainId *big.Int
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterStakingETHEvent is a free log retrieval operation binding the contract event 0xc30647c83f1ae8cba53f299401139cccd78519cd0cdf3937716ee01df27d509f.
//
// Solidity: event StakingETHEvent(address indexed user, uint256 chainId, uint256 amount)
func (_PoolManager *PoolManagerFilterer) FilterStakingETHEvent(opts *bind.FilterOpts, user []common.Address) (*PoolManagerStakingETHEventIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _PoolManager.contract.FilterLogs(opts, "StakingETHEvent", userRule)
	if err != nil {
		return nil, err
	}
	return &PoolManagerStakingETHEventIterator{contract: _PoolManager.contract, event: "StakingETHEvent", logs: logs, sub: sub}, nil
}

// WatchStakingETHEvent is a free log subscription operation binding the contract event 0xc30647c83f1ae8cba53f299401139cccd78519cd0cdf3937716ee01df27d509f.
//
// Solidity: event StakingETHEvent(address indexed user, uint256 chainId, uint256 amount)
func (_PoolManager *PoolManagerFilterer) WatchStakingETHEvent(opts *bind.WatchOpts, sink chan<- *PoolManagerStakingETHEvent, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _PoolManager.contract.WatchLogs(opts, "StakingETHEvent", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolManagerStakingETHEvent)
				if err := _PoolManager.contract.UnpackLog(event, "StakingETHEvent", log); err != nil {
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

// ParseStakingETHEvent is a log parse operation binding the contract event 0xc30647c83f1ae8cba53f299401139cccd78519cd0cdf3937716ee01df27d509f.
//
// Solidity: event StakingETHEvent(address indexed user, uint256 chainId, uint256 amount)
func (_PoolManager *PoolManagerFilterer) ParseStakingETHEvent(log types.Log) (*PoolManagerStakingETHEvent, error) {
	event := new(PoolManagerStakingETHEvent)
	if err := _PoolManager.contract.UnpackLog(event, "StakingETHEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the PoolManager contract.
type PoolManagerUnpausedIterator struct {
	Event *PoolManagerUnpaused // Event containing the contract specifics and raw log

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
func (it *PoolManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolManagerUnpaused)
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
		it.Event = new(PoolManagerUnpaused)
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
func (it *PoolManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolManagerUnpaused represents a Unpaused event raised by the PoolManager contract.
type PoolManagerUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PoolManager *PoolManagerFilterer) FilterUnpaused(opts *bind.FilterOpts) (*PoolManagerUnpausedIterator, error) {

	logs, sub, err := _PoolManager.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &PoolManagerUnpausedIterator{contract: _PoolManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PoolManager *PoolManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PoolManagerUnpaused) (event.Subscription, error) {

	logs, sub, err := _PoolManager.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolManagerUnpaused)
				if err := _PoolManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PoolManager *PoolManagerFilterer) ParseUnpaused(log types.Log) (*PoolManagerUnpaused, error) {
	event := new(PoolManagerUnpaused)
	if err := _PoolManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolManagerWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the PoolManager contract.
type PoolManagerWithdrawIterator struct {
	Event *PoolManagerWithdraw // Event containing the contract specifics and raw log

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
func (it *PoolManagerWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolManagerWithdraw)
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
		it.Event = new(PoolManagerWithdraw)
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
func (it *PoolManagerWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolManagerWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolManagerWithdraw represents a Withdraw event raised by the PoolManager contract.
type PoolManagerWithdraw struct {
	User        common.Address
	StartPoolId *big.Int
	EndPoolId   *big.Int
	ChainId     *big.Int
	Token       common.Address
	Amount      *big.Int
	Reward      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xe7ec1c5ac03ffbc42e14868f1c00af28ca9cc9e4f02bec224142d7cc1ad64ca8.
//
// Solidity: event Withdraw(address _user, uint256 startPoolId, uint256 EndPoolId, uint256 chainId, address _token, uint256 Amount, uint256 Reward)
func (_PoolManager *PoolManagerFilterer) FilterWithdraw(opts *bind.FilterOpts) (*PoolManagerWithdrawIterator, error) {

	logs, sub, err := _PoolManager.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &PoolManagerWithdrawIterator{contract: _PoolManager.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xe7ec1c5ac03ffbc42e14868f1c00af28ca9cc9e4f02bec224142d7cc1ad64ca8.
//
// Solidity: event Withdraw(address _user, uint256 startPoolId, uint256 EndPoolId, uint256 chainId, address _token, uint256 Amount, uint256 Reward)
func (_PoolManager *PoolManagerFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *PoolManagerWithdraw) (event.Subscription, error) {

	logs, sub, err := _PoolManager.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolManagerWithdraw)
				if err := _PoolManager.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xe7ec1c5ac03ffbc42e14868f1c00af28ca9cc9e4f02bec224142d7cc1ad64ca8.
//
// Solidity: event Withdraw(address _user, uint256 startPoolId, uint256 EndPoolId, uint256 chainId, address _token, uint256 Amount, uint256 Reward)
func (_PoolManager *PoolManagerFilterer) ParseWithdraw(log types.Log) (*PoolManagerWithdraw, error) {
	event := new(PoolManagerWithdraw)
	if err := _PoolManager.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolManagerWithdrawTokenIterator is returned from FilterWithdrawToken and is used to iterate over the raw logs and unpacked data for WithdrawToken events raised by the PoolManager contract.
type PoolManagerWithdrawTokenIterator struct {
	Event *PoolManagerWithdrawToken // Event containing the contract specifics and raw log

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
func (it *PoolManagerWithdrawTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolManagerWithdrawToken)
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
		it.Event = new(PoolManagerWithdrawToken)
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
func (it *PoolManagerWithdrawTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolManagerWithdrawTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolManagerWithdrawToken represents a WithdrawToken event raised by the PoolManager contract.
type PoolManagerWithdrawToken struct {
	TokenAddress    common.Address
	Sender          common.Address
	WithdrawAddress common.Address
	Amount          *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterWithdrawToken is a free log retrieval operation binding the contract event 0x9ca7c1e047552a8048d924a5a8d3c150eb861086a72a9100e5f19d1176c1b746.
//
// Solidity: event WithdrawToken(address indexed tokenAddress, address sender, address withdrawAddress, uint256 amount)
func (_PoolManager *PoolManagerFilterer) FilterWithdrawToken(opts *bind.FilterOpts, tokenAddress []common.Address) (*PoolManagerWithdrawTokenIterator, error) {

	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _PoolManager.contract.FilterLogs(opts, "WithdrawToken", tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &PoolManagerWithdrawTokenIterator{contract: _PoolManager.contract, event: "WithdrawToken", logs: logs, sub: sub}, nil
}

// WatchWithdrawToken is a free log subscription operation binding the contract event 0x9ca7c1e047552a8048d924a5a8d3c150eb861086a72a9100e5f19d1176c1b746.
//
// Solidity: event WithdrawToken(address indexed tokenAddress, address sender, address withdrawAddress, uint256 amount)
func (_PoolManager *PoolManagerFilterer) WatchWithdrawToken(opts *bind.WatchOpts, sink chan<- *PoolManagerWithdrawToken, tokenAddress []common.Address) (event.Subscription, error) {

	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _PoolManager.contract.WatchLogs(opts, "WithdrawToken", tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolManagerWithdrawToken)
				if err := _PoolManager.contract.UnpackLog(event, "WithdrawToken", log); err != nil {
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

// ParseWithdrawToken is a log parse operation binding the contract event 0x9ca7c1e047552a8048d924a5a8d3c150eb861086a72a9100e5f19d1176c1b746.
//
// Solidity: event WithdrawToken(address indexed tokenAddress, address sender, address withdrawAddress, uint256 amount)
func (_PoolManager *PoolManagerFilterer) ParseWithdrawToken(log types.Log) (*PoolManagerWithdrawToken, error) {
	event := new(PoolManagerWithdrawToken)
	if err := _PoolManager.contract.UnpackLog(event, "WithdrawToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
