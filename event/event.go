// Package event 实现事件处理器（Event Processor）
// 这是跨链桥系统的第二阶段处理：解析和转换事件为业务数据
//
// 工作流程：
// 1. Synchronizer 同步原始事件日志到数据库（ContractEvents 表）
// 2. Event Processor 从数据库读取事件并解析为业务数据
// 3. 将业务数据存储到对应的业务表（BridgeInitiate、BridgeFinalize 等）
//
// 与 Synchronizer 的区别：
// - Synchronizer: 同步区块链 → 数据库（原始数据）
// - Event Processor: 数据库 → 数据库（业务数据）
package event

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/log"

	"github.com/Sandwichzzy/relayer-node/common/tasks"
	"github.com/Sandwichzzy/relayer-node/database"
	"github.com/Sandwichzzy/relayer-node/database/relayer"
	"github.com/Sandwichzzy/relayer-node/event/contracts"
	"github.com/Sandwichzzy/relayer-node/metrics"
	"github.com/Sandwichzzy/relayer-node/synchronizer/retry"
)

// ProcessorConfig 事件处理器配置
type ProcessorConfig struct {
	LoopInterval       time.Duration // 处理循环间隔
	StartHeight        *big.Int      // 保留字段（目前未使用）
	MsgManagerAddress  string        // MessageManager 合约地址
	PoolManagerAddress string        // PoolManager 合约地址
	ChainId            string        // 链 ID
	EventStartBlock    uint64        // 事件解析起始区块（对应配置文件中的 event_unpack_block）
	Epoch              uint64        // 每次处理的区块数量上限
}

type Processor struct {
	db                *database.DB
	eventBlocksConfig *ProcessorConfig
	resourceCtx       context.Context
	resourceCancel    context.CancelFunc
	tasks             tasks.Group
	msgManager        *contracts.MessageManager
	poolManager       *contracts.PoolManager
	LatestBlockHeader *relayer.EventBlockListener
	eventRegistry     *metrics.SyncerMetrics
}

func NewProcessor(db *database.DB, eventBlocksConfig *ProcessorConfig, eventRegistry *metrics.SyncerMetrics, shutdown context.CancelCauseFunc) (*Processor, error) {
	poolManager, err := contracts.NewPoolManager(eventBlocksConfig.PoolManagerAddress, eventBlocksConfig.ChainId)
	if err != nil {
		log.Error("new pool manager fail", "err", err)
		return nil, err
	}

	msgManager, err := contracts.NewMessageManager(eventBlocksConfig.MsgManagerAddress, eventBlocksConfig.ChainId)
	if err != nil {
		log.Error("new message manager fail", "err", err)
		return nil, err
	}

	latestBlockHeader, err := db.EventBlockListener.GetLastBlockNumber(eventBlocksConfig.ChainId)
	if err != nil {
		log.Error("get latest event block header fail", "err", err)
		return nil, err
	}

	resCtx, resCancel := context.WithCancel(context.Background())

	return &Processor{
		db:                db,
		resourceCtx:       resCtx,
		resourceCancel:    resCancel,
		eventBlocksConfig: eventBlocksConfig,
		tasks: tasks.Group{HandleCrit: func(err error) {
			shutdown(fmt.Errorf("critical error in bridge processor: %w", err))
		}},
		poolManager:       poolManager,
		msgManager:        msgManager,
		LatestBlockHeader: latestBlockHeader,
		eventRegistry:     eventRegistry,
	}, nil
}

func (ep *Processor) Start() error {
	log.Info("starting bridge processor...")
	tickerL1Worker := time.NewTicker(time.Second * 5)
	ep.tasks.Go(func() error {
		for range tickerL1Worker.C {
			err := ep.processEvent()
			if err != nil {
				log.Error("process event fail", "err", err)
				continue
			}
		}
		return nil
	})
	return nil
}

func (ep *Processor) Close() error {
	ep.resourceCancel()
	return ep.tasks.Wait()
}

// processEvent 处理事件：从数据库读取原始事件并解析为业务数据
//
// 处理流程：
// 1. 确定处理区块范围（eventStartBlock → eventEndBlock）
// 2. 从数据库读取该范围内的原始事件
// 3. 解析事件为业务数据（BridgeInitiate、MessageSent 等）
// 4. 存储业务数据到对应的数据库表
//
// 关键：Event Processor 依赖 Synchronizer 先同步数据到数据库
func (ep *Processor) processEvent() error {
	eventEndBlock := big.NewInt(0)

	// 初始化起始区块：使用配置的 EventStartBlock（event_unpack_block）
	// 这是第一次运行时的起始点，只在数据库没有记录时使用
	eventStartBlock := big.NewInt(int64(ep.eventBlocksConfig.EventStartBlock))

	// 从数据库查询上次处理到的区块
	// 支持断点续传：如果有记录，从上次处理的下一个区块开始
	eventBlock, err := ep.db.EventBlockListener.GetLastBlockNumber(ep.eventBlocksConfig.ChainId)
	if err != nil {
		log.Error("get latest block number fail", "err", err)
		return err
	}
	if eventBlock != nil {
		// 从上次处理的区块的下一个区块开始
		eventStartBlock = new(big.Int).Add(eventBlock.BlockNumber, big.NewInt(1))
	}

	// 获取 Synchronizer 已同步到的最新区块
	// Event Processor 只能处理 Synchronizer 已同步的区块范围
	blockHeader, err := ep.db.Blocks.ChainLatestBlockHeader(ep.eventBlocksConfig.ChainId)
	if err != nil {
		log.Error("get latest block header fail", "err", err)
		return err
	}

	if blockHeader != nil {
		eventEndBlock = blockHeader.Number
	} else {
		// 如果 Synchronizer 还没有同步任何区块，则等待
		return nil
	}

	// 如果起始区块大于结束区块，说明没有新数据需要处理
	if eventStartBlock.Cmp(eventEndBlock) > 0 {
		return nil
	}

	log.Info("process event latest block number", "startBlock", eventStartBlock, "endBlock", eventEndBlock, "chainId", ep.eventBlocksConfig.ChainId)
	ep.eventRegistry.RecordEventBlockHeight(ep.eventBlocksConfig.ChainId, eventEndBlock)

	// 处理 PoolManager 合约事件（资金池相关）
	// 解析出：BridgeInitiate、BridgeFinalize、LPWithdraw、ClaimReward、StakingRecord 等
	bridgeInitiateList, bridgeFinalizeList, lpWithdrawList, claimRewardList, stakingRecordList, err := ep.poolManager.ProcessPoolManagerEvent(ep.db, eventStartBlock, eventEndBlock)
	if err != nil {
		log.Error("process pool manager event", "err", err)
		return err
	}

	// 处理 MessageManager 合约事件（消息管理相关）
	// 解析出：MessageSent、MessageHash 等
	bridgeMsgSentList, bridgeMsgHashList, err := ep.msgManager.ProcessMessageManager(ep.db, eventStartBlock, eventEndBlock)
	if err != nil {
		log.Error("process message manager fail", "err", err)
		return err
	}

	if len(bridgeMsgSentList) != 0 || len(bridgeMsgHashList) != 0 {
		log.Info("process message manager", "bridgeMsgSentList", bridgeMsgSentList, "bridgeMsgHashList", bridgeMsgHashList)
	}

	retryStrategy := &retry.ExponentialStrategy{Min: 1000, Max: 20_000, MaxJitter: 250}
	if _, err := retry.Do[interface{}](context.Background(), 10, retryStrategy, func() (interface{}, error) {
		if err := ep.db.Transaction(func(tx *database.DB) error {
			if len(bridgeInitiateList) > 0 {
				err := ep.db.BridgeInitiate.StoreBridgeInitiates(bridgeInitiateList)
				if err != nil {
					log.Error("store bridge initiates fail")
					return err
				}
			}

			if len(bridgeFinalizeList) > 0 {
				err := ep.db.BridgeFinalize.StoreBridgeFinalizes(bridgeFinalizeList)
				if err != nil {
					log.Error("store bridge finalized fail")
					return err
				}
			}

			if len(lpWithdrawList) > 0 {
				err := ep.db.LpWithdraw.StoreLpWithdraws(lpWithdrawList)
				if err != nil {
					log.Error("store lp withdraw list fail")
					return err
				}
			}

			if len(claimRewardList) > 0 {
				err := ep.db.ClaimReward.StoreBridgeClaims(claimRewardList)
				if err != nil {
					log.Error("store claim rewards fail")
					return err
				}
			}

			if len(stakingRecordList) > 0 {
				err := ep.db.StakingRecord.StoreStakingRecords(stakingRecordList)
				if err != nil {
					log.Error("store staking records fail")
					return err
				}
			}

			if len(bridgeMsgSentList) > 0 {
				err := ep.db.BridgeMsgSent.StoreBridgeMsgSents(bridgeMsgSentList)
				if err != nil {
					log.Error("store bridge records fail")
					return err
				}
			}

			if len(bridgeMsgHashList) > 0 {
				err := ep.db.BridgeMsgHash.StoreBridgeMsgHashs(bridgeMsgHashList)
				if err != nil {
					log.Error("store bridge records fail")
					return err
				}
			}

			blockEventListen := relayer.EventBlockListener{
				ChainId:     ep.eventBlocksConfig.ChainId,
				BlockNumber: eventEndBlock,
				Updated:     uint64(time.Now().Unix()),
			}
			err := ep.db.EventBlockListener.SaveOrUpdateLastBlockNumber(blockEventListen)
			if err != nil {
				log.Error("save or update block number fail", "err", err)
				return err
			}
			return nil
		}); err != nil {
			log.Error("unable to persist batch", "err", err)
			return nil, fmt.Errorf("unable to persist batch: %w", err)
		}
		return nil, nil
	}); err != nil {
		log.Error("process event error", "err", err)
		return err
	}
	return nil
}
