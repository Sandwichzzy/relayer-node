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

// 这个模块是整个跨链桥系统的事件采集层，负责将链上事件转换为数据库记录。
type ProcessorConfig struct {
	LoopInterval       time.Duration
	StartHeight        *big.Int
	MsgManagerAddress  string
	PoolManagerAddress string
	ChainId            string
	EventStartBlock    uint64
	Epoch              uint64
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

func (ep *Processor) processEvent() error {
	eventEndBlock := big.NewInt(0)
	eventStartBlock := big.NewInt(int64(ep.eventBlocksConfig.EventStartBlock))
	eventBlock, err := ep.db.EventBlockListener.GetLastBlockNumber(ep.eventBlocksConfig.ChainId)
	if err != nil {
		log.Error("get latest block number fail", "err", err)
		return err
	}
	if eventBlock != nil {
		eventStartBlock = new(big.Int).Add(eventBlock.BlockNumber, big.NewInt(1))
	}

	blockHeader, err := ep.db.Blocks.ChainLatestBlockHeader(ep.eventBlocksConfig.ChainId)
	if err != nil {
		log.Error("get latest block header fail", "err", err)
		return err
	}

	if blockHeader != nil {
		eventEndBlock = blockHeader.Number
	} else {
		return nil
	}

	if eventStartBlock.Cmp(eventEndBlock) > 0 {
		return nil
	}

	log.Info("process event latest block number", "startBlock", eventStartBlock, "endBlock", eventEndBlock, "chainId", ep.eventBlocksConfig.ChainId)
	ep.eventRegistry.RecordEventBlockHeight(ep.eventBlocksConfig.ChainId, eventEndBlock)

	bridgeInitiateList, bridgeFinalizeList, lpWithdrawList, claimRewardList, stakingRecordList, err := ep.poolManager.ProcessPoolManagerEvent(ep.db, eventStartBlock, eventEndBlock)
	if err != nil {
		log.Error("process pool manager event", "err", err)
		return err
	}
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
