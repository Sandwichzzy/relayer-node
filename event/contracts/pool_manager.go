package contracts

import (
	"context"
	"math/big"

	"github.com/google/uuid"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"

	"github.com/Sandwichzzy/relayer-node/bindings"
	common2 "github.com/Sandwichzzy/relayer-node/common"
	"github.com/Sandwichzzy/relayer-node/database"
	"github.com/Sandwichzzy/relayer-node/database/event"
	"github.com/Sandwichzzy/relayer-node/database/relayer"
)

//负责处理 PoolManager 合约的 8 种不同事件

type PoolManager struct {
	PoolManagerAbi     *abi.ABI
	PoolManagerFilter  *bindings.PoolManagerFilterer
	PoolManagerCtx     context.Context
	PoolManagerAddress string
	ChainId            string
}

func NewPoolManager(PoolManagerAddress string, ChainId string) (*PoolManager, error) {
	poolManagerAbi, err := bindings.PoolManagerMetaData.GetAbi()
	if err != nil {
		log.Error("get pool manager abi fail", "err", err)
		return nil, err
	}

	poolManagerUnpack, err := bindings.NewPoolManagerFilterer(common.Address{}, nil)
	if err != nil {
		log.Error("new pool manager filter fail", "err", err)
		return nil, err
	}

	return &PoolManager{
		PoolManagerAbi:     poolManagerAbi,
		PoolManagerFilter:  poolManagerUnpack,
		PoolManagerCtx:     context.Background(),
		PoolManagerAddress: PoolManagerAddress,
		ChainId:            ChainId,
	}, nil
}

func (pm *PoolManager) ProcessPoolManagerEvent(db *database.DB, fromHeight *big.Int, toHeight *big.Int) ([]relayer.BridgeInitiate, []relayer.BridgeFinalize, []relayer.LpWithdraw, []relayer.ClaimReward, []relayer.StakingRecord, error) {
	contractEventFilter := event.ContractEvent{ContractAddress: common.HexToAddress(pm.PoolManagerAddress)}
	contractEventList, err := db.ContractEvents.ContractEventsWithFilter(pm.ChainId, contractEventFilter, fromHeight, toHeight)
	if err != nil {
		log.Error("get contracts event list fail", "err", err)
		return nil, nil, nil, nil, nil, err
	}

	var existTxList []string

	var bridgeInitiateList []relayer.BridgeInitiate
	var bridgeFinalizeList []relayer.BridgeFinalize

	var stakingRecordList []relayer.StakingRecord

	var lpWithdrawList []relayer.LpWithdraw
	var claimRewardList []relayer.ClaimReward

	for _, eventItem := range contractEventList {
		rlpLog := eventItem.RLPLog
		if common2.Contains(existTxList, rlpLog.TxHash.String()) {
			continue
		}

		header, err := db.Blocks.ChainBlockHeader(pm.ChainId, eventItem.BlockHash)
		if err != nil {
			log.Error("ProcessPoolManagerEvent db Blocks BlockHeader by BlockHash fail", "err", err)
			return nil, nil, nil, nil, nil, err
		}

		// 触发场景: 用户在源链发起 ETH/原生币跨链转账
		if eventItem.EventSignature.String() == pm.PoolManagerAbi.Events["InitiateETH"].ID.String() {
			parseInitializedEthEvent, err := pm.PoolManagerFilter.ParseInitiateETH(*rlpLog)
			if err != nil {
				log.Error("parse initialized eth event fail", "err", err)
				return nil, nil, nil, nil, nil, err
			}
			bridgeInitiateItem := relayer.BridgeInitiate{
				GUID:               uuid.New(),
				TxHash:             parseInitializedEthEvent.Raw.TxHash,
				BlockNumber:        big.NewInt(int64(parseInitializedEthEvent.Raw.BlockNumber)),
				SourceChainId:      parseInitializedEthEvent.SourceChainId,
				DestChainId:        parseInitializedEthEvent.DestChainId,
				FromAddress:        parseInitializedEthEvent.From,
				ToAddress:          parseInitializedEthEvent.To,
				SourceTokenAddress: common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE"),
				DestTokenAddress:   parseInitializedEthEvent.DestTokenAddress,
				Amount:             parseInitializedEthEvent.Value,
				Timestamp:          header.Timestamp,
			}
			bridgeInitiateList = append(bridgeInitiateList, bridgeInitiateItem)
		}
		// 触发场景: 用户在源链发起 ERC20 代币(如 USDT)跨链转账
		if eventItem.EventSignature.String() == pm.PoolManagerAbi.Events["InitiateERC20"].ID.String() {
			parseInitializedErc20Event, err := pm.PoolManagerFilter.ParseInitiateERC20(*rlpLog)
			if err != nil {
				log.Error("parse initialized erc20 event fail", "err", err)
				return nil, nil, nil, nil, nil, err
			}
			bridgeInitiateItem := relayer.BridgeInitiate{
				GUID:               uuid.New(),
				TxHash:             parseInitializedErc20Event.Raw.TxHash,
				BlockNumber:        big.NewInt(int64(parseInitializedErc20Event.Raw.BlockNumber)),
				SourceChainId:      parseInitializedErc20Event.SourceChainId,
				DestChainId:        parseInitializedErc20Event.DestChainId,
				FromAddress:        parseInitializedErc20Event.From,
				ToAddress:          parseInitializedErc20Event.To,
				SourceTokenAddress: parseInitializedErc20Event.SourceTokenAddress,
				DestTokenAddress:   parseInitializedErc20Event.DestTokenAddress,
				Amount:             parseInitializedErc20Event.Value,
				Timestamp:          header.Timestamp,
			}
			bridgeInitiateList = append(bridgeInitiateList, bridgeInitiateItem)
		}

		//触发场景: Relayer 在目标链完成 ETH 转账，触发此事件
		if eventItem.EventSignature.String() == pm.PoolManagerAbi.Events["FinalizeETH"].ID.String() {
			parseFinalizeEthEvent, err := pm.PoolManagerFilter.ParseFinalizeETH(*rlpLog)
			if err != nil {
				log.Error("parse finalize eth event fail", "err", err)
				return nil, nil, nil, nil, nil, err
			}
			bridgeFinalizeItem := relayer.BridgeFinalize{
				GUID:               uuid.New(),
				TxHash:             parseFinalizeEthEvent.Raw.TxHash,
				BlockNumber:        big.NewInt(int64(parseFinalizeEthEvent.Raw.BlockNumber)),
				SourceChainId:      parseFinalizeEthEvent.SourceChainId,
				DestChainId:        parseFinalizeEthEvent.DestChainId,
				FromAddress:        parseFinalizeEthEvent.From,
				ToAddress:          parseFinalizeEthEvent.To,
				SourceTokenAddress: parseFinalizeEthEvent.SourceTokenAddress,
				DestTokenAddress:   common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE"),
				Amount:             parseFinalizeEthEvent.Value,
				Timestamp:          header.Timestamp,
			}
			bridgeFinalizeList = append(bridgeFinalizeList, bridgeFinalizeItem)
		}
		//  触发场景: Relayer 在目标链完成 ERC20 转账
		if eventItem.EventSignature.String() == pm.PoolManagerAbi.Events["FinalizeERC20"].ID.String() {
			parseFinalizeErc20Event, err := pm.PoolManagerFilter.ParseFinalizeERC20(*rlpLog)
			if err != nil {
				log.Error("parse finalize erc20 event fail", "err", err)
				return nil, nil, nil, nil, nil, err
			}
			bridgeFinalizeItem := relayer.BridgeFinalize{
				GUID:               uuid.New(),
				TxHash:             parseFinalizeErc20Event.Raw.TxHash,
				BlockNumber:        big.NewInt(int64(parseFinalizeErc20Event.Raw.BlockNumber)),
				SourceChainId:      parseFinalizeErc20Event.SourceChainId,
				DestChainId:        parseFinalizeErc20Event.DestChainId,
				FromAddress:        parseFinalizeErc20Event.From,
				ToAddress:          parseFinalizeErc20Event.To,
				SourceTokenAddress: parseFinalizeErc20Event.SourceTokenAddress,
				DestTokenAddress:   parseFinalizeErc20Event.DestTokenAddress,
				Amount:             parseFinalizeErc20Event.Value,
				Timestamp:          header.Timestamp,
			}
			bridgeFinalizeList = append(bridgeFinalizeList, bridgeFinalizeItem)
		}
		//触发场景: LP 提供者质押 ETH 到流动性池
		if eventItem.EventSignature.String() == pm.PoolManagerAbi.Events["StakingETHEvent"].ID.String() {
			parseStakingETHEvent, err := pm.PoolManagerFilter.ParseStakingETHEvent(*rlpLog)
			if err != nil {
				log.Error("parse staking eth event fail", "err", err)
				return nil, nil, nil, nil, nil, err
			}

			stakingRecordItem := relayer.StakingRecord{
				GUID:               uuid.New(),
				TxHash:             parseStakingETHEvent.Raw.TxHash,
				BlockNumber:        big.NewInt(int64(parseStakingETHEvent.Raw.BlockNumber)),
				UserAddress:        parseStakingETHEvent.User,
				SourceTokenAddress: common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE"),
				DestTokenAddress:   common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE"),
				Amount:             parseStakingETHEvent.Amount,
				Reward:             big.NewInt(0),
				StartPoolId:        big.NewInt(0),
				EndPoolId:          big.NewInt(0),
				ChainId:            pm.ChainId,
				Status:             0,
				Timestamp:          header.Timestamp,
			}
			stakingRecordList = append(stakingRecordList, stakingRecordItem)
		}

		if eventItem.EventSignature.String() == pm.PoolManagerAbi.Events["StakingERC20Event"].ID.String() {
			parseStakingERC20Event, err := pm.PoolManagerFilter.ParseStakingERC20Event(*rlpLog)
			if err != nil {
				log.Error("parse staking erc20 event fail", "err", err)
				return nil, nil, nil, nil, nil, err
			}
			stakingRecordItem := relayer.StakingRecord{
				GUID:               uuid.New(),
				TxHash:             parseStakingERC20Event.Raw.TxHash,
				BlockNumber:        big.NewInt(int64(parseStakingERC20Event.Raw.BlockNumber)),
				UserAddress:        parseStakingERC20Event.User,
				SourceTokenAddress: parseStakingERC20Event.Token,
				DestTokenAddress:   parseStakingERC20Event.Token,
				Amount:             parseStakingERC20Event.Amount,
				Reward:             big.NewInt(0),
				StartPoolId:        big.NewInt(0),
				EndPoolId:          big.NewInt(0),
				ChainId:            pm.ChainId,
				Status:             0,
				Timestamp:          header.Timestamp,
			}
			stakingRecordList = append(stakingRecordList, stakingRecordItem)
		}
		//  触发场景: LP 提取本金和奖励
		if eventItem.EventSignature.String() == pm.PoolManagerAbi.Events["Withdraw"].ID.String() {
			parseWithdrawEvent, err := pm.PoolManagerFilter.ParseWithdraw(*rlpLog)
			if err != nil {
				log.Error("parse withdraw event fail", "err", err)
				return nil, nil, nil, nil, nil, err
			}
			lpWithdrawItem := relayer.LpWithdraw{
				GUID:         uuid.New(),
				BlockNumber:  big.NewInt(int64(parseWithdrawEvent.Raw.BlockNumber)),
				TxHash:       parseWithdrawEvent.Raw.TxHash,
				StartPoolId:  parseWithdrawEvent.StartPoolId,
				EndPoolId:    parseWithdrawEvent.EndPoolId,
				ChainId:      pm.ChainId,
				TokenAddress: parseWithdrawEvent.Token,
				Amount:       parseWithdrawEvent.Amount,
				RewardAmount: parseWithdrawEvent.Reward,
				Timestamp:    header.Timestamp,
			}
			lpWithdrawList = append(lpWithdrawList, lpWithdrawItem)

		}
		//  触发场景: LP 不提取本金，只领取累积的手续费奖励
		if eventItem.EventSignature.String() == pm.PoolManagerAbi.Events["ClaimReward"].ID.String() {
			parseClaimRewardEvent, err := pm.PoolManagerFilter.ParseClaimReward(*rlpLog)
			if err != nil {
				log.Error("parse claim reward event fail", "err", err)
				return nil, nil, nil, nil, nil, err
			}
			claimRewardItem := relayer.ClaimReward{
				GUID:         uuid.New(),
				BlockNumber:  big.NewInt(int64(parseClaimRewardEvent.Raw.BlockNumber)),
				TxHash:       parseClaimRewardEvent.Raw.TxHash,
				ClaimAddress: parseClaimRewardEvent.User,
				TokenAddress: parseClaimRewardEvent.Token,
				StartPoolId:  parseClaimRewardEvent.StartPoolId,
				EndPoolId:    parseClaimRewardEvent.EndPoolId,
				ChainId:      pm.ChainId,
				RewardAmount: parseClaimRewardEvent.Reward,
				Timestamp:    header.Timestamp,
			}
			claimRewardList = append(claimRewardList, claimRewardItem)
		}
		existTxList = append(existTxList, rlpLog.TxHash.String())
	}
	return bridgeInitiateList, bridgeFinalizeList, lpWithdrawList, claimRewardList, stakingRecordList, nil
}
