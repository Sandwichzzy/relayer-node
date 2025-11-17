// Package worker 实现跨链桥的消息记录管理器
// 职责：
// 1. 创建和管理 BridgeRecords 表（用户友好的跨链记录）
// 2. 合并源链和目标链的信息形成完整的跨链记录
// 3. 为前端 API 提供查询数据源
package worker

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/Sandwichzzy/relayer-node/common/tasks"
	"github.com/Sandwichzzy/relayer-node/database"
	"github.com/Sandwichzzy/relayer-node/database/relayer"
	"github.com/Sandwichzzy/relayer-node/service/websocket"
	"github.com/Sandwichzzy/relayer-node/synchronizer/node"
	"github.com/Sandwichzzy/relayer-node/synchronizer/retry"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/google/uuid"
)

// WorkerHandleConfig Worker 配置
type WorkerHandleConfig struct {
	LoopInterval time.Duration // 处理循环间隔
	ChainIds     []string      // 要处理的链 ID 列表
}

// WorkerHandle 跨链记录管理器
// 负责创建和维护用户友好的跨链记录（BridgeRecords 表）
//
// 核心任务：
// 1. onProcessMessageSend: 处理源链消息，创建初始记录
// 2. onProcessRelayerMessage: 处理目标链确认，更新完成状态
type WorkerHandle struct {
	db             *database.DB        // 数据库连接
	wConf          *WorkerHandleConfig // 配置
	ethClient      node.EthClient      // 以太坊客户端（未使用）
	wsHub          *websocket.Hub      // WebSocket 广播中心
	resourceCtx    context.Context     // 资源上下文
	resourceCancel context.CancelFunc  // 取消函数
	tasks          tasks.Group         // 任务组
}

// NewWorkerHandle 创建新的工作节点处理器
func NewWorkerHandle(db *database.DB, wConf *WorkerHandleConfig, wsHub *websocket.Hub, shutdown context.CancelCauseFunc) (*WorkerHandle, error) {
	resCtx, resCancel := context.WithCancel(context.Background())
	return &WorkerHandle{
		db:             db,
		wConf:          wConf,
		wsHub:          wsHub,
		resourceCtx:    resCtx,
		resourceCancel: resCancel,
		tasks: tasks.Group{
			// 定义严重错误处理器，触发系统关闭
			HandleCrit: func(err error) {
				shutdown(fmt.Errorf("critical error in worker handle processor: %w", err))
			},
		},
	}, nil
}

// Cancels all contexts and waits for all tasks to complete
func (sh *WorkerHandle) Close() error {
	sh.resourceCancel()
	return sh.tasks.Wait()
}

// 启动两个独立的定时任务：
// - 任务1: 处理源链的桥接消息 (onProcessMessageSend)
// - 任务2: 处理目标链的中继确认 (onProcessRelayerMessage)
// - 两个任务并发运行，互不干扰，按配置的时间间隔定期执行
func (sh *WorkerHandle) Start() error {
	workerTicker := time.NewTicker(sh.wConf.LoopInterval)
	sh.tasks.Go(func() error {
		for range workerTicker.C {
			for _, chainId := range sh.wConf.ChainIds {
				log.Info("onProcessMessageSend start")
				err := sh.onProcessMessageSend(chainId)
				if err != nil {
					log.Error("process message send fail", "err", err)
					continue
				}
			}
		}
		return nil
	})

	sh.tasks.Go(func() error {
		for range workerTicker.C {
			for _, chainId := range sh.wConf.ChainIds {
				log.Info("onProcessRelayerMessage start")
				err := sh.onProcessRelayerMessage(chainId)
				if err != nil {
					log.Error("process message send fail", "err", err)
					continue
				}
			}
		}
		return nil
	})
	return nil
}

// 处理源链上的桥接消息：
// 1. 查询未处理的 BridgeMsgSent 事件
// 2. 验证代币配置是否存在
// 3. 创建 BridgeRecord 记录（状态=0，待处理）
// 4. 使用数据库事务保存记录
// 5. 标记源消息为已处理，避免重复
func (sh *WorkerHandle) onProcessMessageSend(chainId string) error {
	var bridgeRecordList []relayer.BridgeRecords
	//查询未处理的 BridgeMsgSent 事件
	bridgeMsgSentList, err := sh.db.BridgeMsgSent.QueryBridgeMsgSents(chainId)
	if err != nil {
		log.Error("bridge message send fail", "err", err)
		return err
	}

	log.Info("Bridge send list", "bridgeMsgSentList", len(bridgeMsgSentList))

	for _, bridgeMsgSent := range bridgeMsgSentList {
		//验证目标链的代币配置是否存在
		chainInfo, errCInfo := sh.db.TokenConfig.QueryTokenBySourceChain(bridgeMsgSent.DestChainId.Uint64(), bridgeMsgSent.DestTokenAddress.String())
		if errCInfo != nil || chainInfo == nil {
			log.Error("Query token by source chain fail, ❌❌❌ maybe you must config token ❌❌❌", "err", err)
			return err
		}

		log.Info("Handle bridge message send success", "bridgeMsgSentTxHash", bridgeMsgSent.TxHash.String(), "ChainId", bridgeMsgSent.SourceChainId.String())
		// 创建桥接记录，包含完整的跨链信息
		bridgeRecord := relayer.BridgeRecords{
			GUID:               uuid.New(),
			SourceChainId:      bridgeMsgSent.SourceChainId,
			DestChainId:        bridgeMsgSent.DestChainId,
			SourceTxHash:       bridgeMsgSent.TxHash,
			DestTxHash:         common.Hash{}, //目标链交易哈希(稍后填充)
			SourceBlockNumber:  bridgeMsgSent.BlockNumber,
			DestBlockNumber:    big.NewInt(0), //目标链区块号(稍后填充)
			SourceTokenAddress: bridgeMsgSent.SourceTokenAddress,
			DestTokenAddress:   bridgeMsgSent.DestTokenAddress,
			MsgHash:            bridgeMsgSent.MsgHash, //消息哈希，用于匹配中继消息
			FromAddress:        bridgeMsgSent.FromAddress,
			ToAddress:          bridgeMsgSent.ToAddress,
			Status:             0, //Status: 0=pending, 1=completed
			TokenName:          chainInfo.TokenName,
			TokenDecimal:       chainInfo.TokenDecimal,
			Amount:             bridgeMsgSent.Amount,
			Nonce:              bridgeMsgSent.MsgNonce,
			Fee:                bridgeMsgSent.Fee,
			MsgSentTimestamp:   bridgeMsgSent.Timestamp,
			ClaimTimestamp:     1,
		}
		bridgeRecordList = append(bridgeRecordList, bridgeRecord)
	}
	// 使用指数退避策略重试数据库操作,参数：最小延迟1秒，最大延迟20秒，最大抖动250ms
	retryStrategy := &retry.ExponentialStrategy{Min: 1000, Max: 20_000, MaxJitter: 250}
	if _, err := retry.Do[interface{}](context.Background(), 10, retryStrategy, func() (interface{}, error) {
		//使用数据库事务确保原子性操作
		if err := sh.db.Transaction(func(tx *database.DB) error {
			if len(bridgeRecordList) > 0 {
				// Save newly created bridge records
				err := sh.db.BridgeRecord.StoreBridgeRecords(bridgeRecordList)
				if err != nil {
					log.Error("store bridge record fail", "err", err)
					return err
				}
			}
			// Mark source messages as handled to avoid duplicate processing
			if len(bridgeMsgSentList) > 0 {
				err := sh.db.BridgeMsgSent.MarkedBridgeMsgSentdHandled(bridgeMsgSentList)
				if err != nil {
					log.Error("store bridge record fail", "err", err)
					return err
				}
			}
			return nil
		}); err != nil {
			log.Error("unable to persist batch", "err", err)
			return nil, fmt.Errorf("unable to persist batch: %w", err)
		}
		return nil, nil
	}); err != nil {
		log.Error("worker handle fail", "err", err)
		return err
	}
	return nil
}

// 处理目标链上的中继确认：
// 1. 查询未处理的 BridgeMsgHash 消息
// 2. 通过消息哈希匹配原始的 BridgeRecord
// 3. 更新记录：设置目标链交易信息，状态=1（已完成）
// 4. 通过 WebSocket 向客户端推送完成通知
// 5. 使用事务保存更新
func (sh *WorkerHandle) onProcessRelayerMessage(chainId string) error {
	// 查询该链上未处理的中继消息
	bridgeRelayMsgList, err := sh.db.BridgeMsgHash.QueryBridgeMsgHashs(chainId)
	if err != nil {
		log.Error("bridge relay message list fail", "err", err)
		return err
	}
	var bridgeRecordList []relayer.BridgeRecords
	var bridgeRelayMsgHandledList []relayer.BridgeMsgHash
	// Iterate through each pending relay message
	for _, bridgeRelayMsg := range bridgeRelayMsgList {
		// Find original bridge record by message hash
		bridgeRecord, err := sh.db.BridgeRecord.QueryBridgeRecordByMsgHash(bridgeRelayMsg.MsgHash.String())
		if err != nil {
			log.Error("query bridge record fail", "err", err)
			return err
		}

		if bridgeRecord == nil {
			log.Info("onProcessRelayerMessage bridgeRecord is nil")
			continue
		}

		log.Info("onProcessRelayerMessage bridge record",
			"bridgeRelayMsgTxHash", bridgeRelayMsg.TxHash,
			"bridgeRelayMsgSourceTokenAddress", bridgeRelayMsg.SourceTokenAddress,
			"bridgeRelayMsgDestTokenAddress", bridgeRelayMsg.DestTokenAddress,
			"bridgeRelayMsgBlockNumber", bridgeRelayMsg.BlockNumber,
			"bridgeRecord.ClaimTimestamp", bridgeRecord.ClaimTimestamp,
		)
		// 更新桥接记录的目标链信息
		bridgeRecord.DestTxHash = bridgeRelayMsg.TxHash
		bridgeRecord.DestBlockNumber = bridgeRelayMsg.BlockNumber
		bridgeRecord.Status = 1
		bridgeRecord.ClaimTimestamp = bridgeRecord.ClaimTimestamp
		bridgeRecordList = append(bridgeRecordList, *bridgeRecord)
		bridgeRelayMsgHandledList = append(bridgeRelayMsgHandledList, bridgeRelayMsg)
	}
	retryStrategy := &retry.ExponentialStrategy{Min: 1000, Max: 20_000, MaxJitter: 250}
	if _, err := retry.Do[interface{}](context.Background(), 10, retryStrategy, func() (interface{}, error) {
		if err := sh.db.Transaction(func(tx *database.DB) error {
			if len(bridgeRecordList) > 0 {
				err := sh.db.BridgeRecord.UpdateBridgeRecordsFromRelayerMsg(bridgeRecordList)
				if err != nil {
					log.Error("store bridge record fail", "err", err)
					return err
				}
			}
			if len(bridgeRelayMsgHandledList) > 0 {
				err := sh.db.BridgeMsgHash.MarkedBridgeRelayMessageHandled(bridgeRelayMsgHandledList)
				if err != nil {
					log.Error("store bridge record fail", "err", err)
					return err
				}
			}
			for _, br := range bridgeRelayMsgHandledList {
				if sh.wsHub != nil {
					msg := &websocket.BridgeFinalizedMessage{
						TxHash: br.TxHash.String(),
						Status: 1,
					}
					sh.wsHub.BroadcastBridgeFinalized(msg) //WebSocket 广播给客户端
				}
			}
			return nil
		}); err != nil {
			log.Error("unable to persist batch", "err", err)
			return nil, fmt.Errorf("unable to persist batch: %w", err)
		}
		return nil, nil
	}); err != nil {
		log.Error("worker handle fail", "err", err)
		return err
	}
	return nil
}
