// Package relayer 实现跨链桥的核心中继业务逻辑
// 这是整个跨链桥系统的第三阶段：执行跨链资产转移
//
// 系统架构的三个阶段：
// 1. Synchronizer: 同步区块链原始数据到数据库
// 2. Event Processor: 解析事件为业务数据
// 3. Relayer Processor: 执行跨链转移（本模块）
//
// 跨链流程：
// 1. 用户在源链调用 PoolManager.bridgeInitiate() 发起跨链
// 2. Event Processor 解析 BridgeInitiate 事件并存储
// 3. RelayerProcessor 监听未处理的 BridgeInitiate
// 4. RelayerProcessor 在目标链调用 PoolManager.bridgeFinalize() 完成跨链
// 5. Event Processor 解析 BridgeFinalize 事件
// 6. RelayerProcessor 更新状态为完成
//
// 子模块：
// - driver/: 交易构造和发送引擎
// - txmgr/: 交易管理器（Gas 优化、重试、确认等待）
package relayer

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/Sandwichzzy/relayer-node/common/tasks"
	"github.com/Sandwichzzy/relayer-node/database"
	"github.com/Sandwichzzy/relayer-node/database/relayer"
	"github.com/Sandwichzzy/relayer-node/metrics"
	"github.com/Sandwichzzy/relayer-node/relayer/driver"
	"github.com/Sandwichzzy/relayer-node/service/websocket"
	"github.com/Sandwichzzy/relayer-node/synchronizer/retry"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
)

// RelayerProcessor 跨链中继处理器
// 负责监听和处理跨链请求，在目标链上完成资产转移
type RelayerProcessor struct {
	LoopInterval      time.Duration
	db                *database.DB
	ethClient         map[uint64]*ethclient.Client   // 各链的 RPC 客户端
	poolMangerAddress map[uint64]string              // 各链的 PoolManager 合约地址
	driverEngine      map[uint64]driver.DriverEngine // 各链的交易驱动引擎
	relayerMetrics    *metrics.SyncerMetrics         // 监控指标收集器
	resourceCtx       context.Context
	resourceCancel    context.CancelFunc
	tasks             tasks.Group
	wsHub             *websocket.Hub // WebSocket 广播中心（用于实时通知）
}

func NewRelayerProcessor(db *database.DB, ethClient map[uint64]*ethclient.Client, poolMangerAddress map[uint64]string, driverEngine map[uint64]driver.DriverEngine, relayerMetrics *metrics.SyncerMetrics, wsHub *websocket.Hub, shutdown context.CancelCauseFunc) (*RelayerProcessor, error) {
	resCtx, resCancel := context.WithCancel(context.Background())
	return &RelayerProcessor{
		db:                db,
		resourceCtx:       resCtx,
		resourceCancel:    resCancel,
		ethClient:         ethClient,
		poolMangerAddress: poolMangerAddress,
		driverEngine:      driverEngine,
		relayerMetrics:    relayerMetrics,
		wsHub:             wsHub,
		tasks: tasks.Group{HandleCrit: func(err error) {
			shutdown(fmt.Errorf("critical error in worker handle processor: %w", err))
		}},
	}, nil
}

func (sh *RelayerProcessor) Close() error {
	sh.resourceCancel()
	return sh.tasks.Wait()
}

// Start 启动中继处理器的两个核心任务
//
// 任务1: onBeforeSendFinalizedTx - 发送跨链完成交易
//   - 查询未处理的 BridgeInitiate 事件（Status = 0）
//   - 在目标链上调用 bridgeFinalize() 完成跨链
//   - 更新状态为已发送（Status = 1）
//
// 任务2: onAfterSendFinalizedTx - 确认跨链完成
//   - 查询未确认的 BridgeFinalize 事件（Status = 0）
//   - 确认目标链上的交易已完成
//   - 更新状态为已完成（Status = 2）
//
// 两个任务独立运行，互不阻塞
func (sh *RelayerProcessor) Start() error {
	workerTicker := time.NewTicker(time.Second * 5)

	// 任务1：发送跨链完成交易
	sh.tasks.Go(func() error {
		for range workerTicker.C {
			log.Info("onBeforeSendFinalizedTx Start")
			err := sh.onBeforeSendFinalizedTx()
			if err != nil {
				log.Error("on send finalized tx fail", "err", err)
				continue
			}
		}
		return nil
	})

	// 任务2：确认跨链完成
	sh.tasks.Go(func() error {
		for range workerTicker.C {
			log.Info("onAfterSendFinalizedTx Start")
			err := sh.onAfterSendFinalizedTx()
			if err != nil {
				log.Error("on send finalized tx fail", "err", err)
				continue
			}
		}
		return nil
	})
	return nil
}

// onBeforeSendFinalizedTx 处理跨链发起事件，在目标链完成资产转移
//
// 处理流程：

//  1. 查询数据库中未处理的 BridgeInitiate 记录（Status = 0）
//  2. 对每个记录：
//     a. 查询对应的 BridgeMsgSent 获取 Fee 和 MsgNonce
//     b. 根据目标链选择对应的 DriverEngine
//     c. 判断代币类型：ETH 或 ERC20
//     d. 调用目标链合约的 bridgeFinalize 方法
//     e. 通过 WebSocket 广播完成通知
//  3. 更新数据库记录状态为已处理（Status = 1）
//
// 重要概念：
// - BridgeInitiate: 用户在源链发起跨链的事件
// - BridgeFinalize: Relayer 在目标链完成跨链的交易
// - 0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE: ETH 的特殊地址标识
func (sh *RelayerProcessor) onBeforeSendFinalizedTx() error {
	// 查询所有未处理的跨链发起事件
	bridgeInitiateList, err := sh.db.BridgeInitiate.QueryUnHandleBridgeInitiate()
	if err != nil {
		log.Error("Query unhandled bridge initiate fail", "err", err)
		return err
	}

	// 逐个处理跨链请求
	for i := 0; i < len(bridgeInitiateList); i++ {
		// 根据目标链 ID 选择对应的驱动引擎
		engine := sh.driverEngine[bridgeInitiateList[i].DestChainId.Uint64()]

		// 查询跨链消息，获取 Fee 和 MsgNonce 信息
		msgSend, err := sh.db.BridgeMsgSent.QueryBridgeMsgSentByHash(bridgeInitiateList[i].TxHash.String())
		if err != nil {
			log.Error("Query bridge message sent by hash fail", "err", err)
			return err
		}

		// 记录交易信息用于调试
		log.Info("build tx information",
			"SourceChainId=", bridgeInitiateList[i].SourceChainId,
			"DestChainId=", bridgeInitiateList[i].DestChainId,
			"FromAddress=", bridgeInitiateList[i].FromAddress,
			"ToAddress=", bridgeInitiateList[i].ToAddress,
			"SourceTokenAddress", bridgeInitiateList[i].SourceTokenAddress,
			"DestTokenAddress", bridgeInitiateList[i].DestTokenAddress,
			"Amount=", bridgeInitiateList[i].Amount,
			"Fee=", msgSend.Fee,
			"MsgNonce=", msgSend.MsgNonce,
		)

		var receipt *types.Receipt

		// 判断目标代币类型：ETH 或 ERC20
		if strings.ToLower(bridgeInitiateList[i].DestTokenAddress.String()) != strings.ToLower("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE") {
			// ERC20 代币跨链
			receipt, err = engine.BridgeFinalizeErc20(bridgeInitiateList[i].SourceChainId, bridgeInitiateList[i].DestChainId, bridgeInitiateList[i].FromAddress, bridgeInitiateList[i].ToAddress, bridgeInitiateList[i].SourceTokenAddress, bridgeInitiateList[i].DestTokenAddress, bridgeInitiateList[i].Amount, msgSend.Fee, msgSend.MsgNonce)
			if err != nil {
				log.Error("bridge finalized erc20 fail", "err", err)
				return err
			}
			log.Info("send bridge finalized erc20 tx success", "hash", receipt.TxHash)
		} else {
			// 原生 ETH 跨链
			receipt, err = engine.BridgeFinalizeETH(bridgeInitiateList[i].SourceChainId, bridgeInitiateList[i].DestChainId, bridgeInitiateList[i].SourceTokenAddress, bridgeInitiateList[i].FromAddress, bridgeInitiateList[i].ToAddress, bridgeInitiateList[i].Amount, msgSend.Fee, msgSend.MsgNonce)
			if err != nil {
				log.Error("bridge finalized eth fail", "err", err)
				return err
			}
			log.Info("send bridge finalized eth tx success", "hash", receipt.TxHash)
		}

		// 标记为已处理
		bridgeInitiateList[i].Status = 1

		// 通过 WebSocket 广播跨链完成通知（实时通知前端）
		if sh.wsHub != nil {
			msg := &websocket.BridgeFinalizedMessage{
				FromAddress: bridgeInitiateList[i].FromAddress.String(),
				ToAddress:   bridgeInitiateList[i].ToAddress.String(),
				TxHash:      receipt.TxHash.String(),
				Status:      0,
			}
			sh.wsHub.BroadcastBridgeFinalized(msg)
		}
	}

	// 使用指数退避重试策略更新数据库状态
	retryStrategy := &retry.ExponentialStrategy{Min: 1000, Max: 20_000, MaxJitter: 250}
	if _, err := retry.Do[interface{}](context.Background(), 10, retryStrategy, func() (interface{}, error) {
		if err := sh.db.Transaction(func(tx *database.DB) error {
			if len(bridgeInitiateList) > 0 {
				err := sh.db.BridgeInitiate.MarkedBridgeInitiateHandled(bridgeInitiateList)
				if err != nil {
					log.Error(" marked bridge initiate fail", "err", err)
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

// onAfterSendFinalizedTx 确认跨链交易已在目标链上完成并更新最终状态
//
// 处理流程：
//  1. 查询数据库中未确认的 BridgeFinalize 记录（Status = 0）
//     这些是 Synchronizer 同步到的目标链上的 bridgeFinalize 事件
//  2. 对每个记录：
//     a. 通过 BridgeMsgHash 找到对应的消息哈希
//     b. 通过 MsgHash 找到对应的 BridgeMsgSent
//     c. 通过 TxHash 找到原始的 BridgeInitiate
//  3. 更新状态：
//     - BridgeInitiate.Status = 2 (最终完成)
//     - BridgeFinalize.Status = 1 (已确认)
//
// 状态流转：
//
//	BridgeInitiate: 0 (未处理) → 1 (已发送完成交易) → 2 (完全确认)
//	BridgeFinalize: 0 (未确认) → 1 (已确认)
func (sh *RelayerProcessor) onAfterSendFinalizedTx() error {
	bridgeFinalizeList, err := sh.db.BridgeFinalize.QueryUnHandleBridgeFinalize()
	if err != nil {
		log.Error("query unhandled bridge finalized fail", "err", err)
		return err
	}

	var bridgeFinalizeRecords []relayer.BridgeFinalize
	var bridgeInitiateList []relayer.BridgeInitiate
	for i := 0; i < len(bridgeFinalizeList); i++ {
		msgHash, err := sh.db.BridgeMsgHash.QueryBridgeMsgHashByHash(bridgeFinalizeList[i].TxHash.String())
		if err != nil {
			log.Error("query bridge message hash fail", "err", err)
			return err
		}
		bridgeMsgSent, err := sh.db.BridgeMsgSent.QueryBridgeMsgSentByMsgHash(msgHash.MsgHash.String())
		if err != nil {
			log.Error("query bridge message sent fail", "err", err)
			return err
		}

		// 为什么可以通过 BridgeMsgSent.TxHash 找到 BridgeInitiate？
		// 它们来自源链的同一笔交易！
		// 当用户在源链调用 bridgeInitiate() 时，这笔交易会触发多个事件：
		//   1. BridgeInitiate 事件（PoolManager 合约发出）
		//   2. BridgeMsgSent 事件（MessageManager 合约发出）
		bridgeInitiate, err := sh.db.BridgeInitiate.QueryBridgeInitiateByHash(bridgeMsgSent.TxHash.String())
		if err != nil {
			log.Error("query bridge initiate fail", "err", err)
			return err
		}
		bridgeInitiate.Status = 2
		bridgeFinalizeList[i].Status = 1
		bridgeInitiateList = append(bridgeInitiateList, *bridgeInitiate)
		bridgeFinalizeRecords = append(bridgeFinalizeRecords, bridgeFinalizeList[i])
	}

	retryStrategy := &retry.ExponentialStrategy{Min: 1000, Max: 20_000, MaxJitter: 250}
	if _, err := retry.Do[interface{}](context.Background(), 10, retryStrategy, func() (interface{}, error) {
		if err := sh.db.Transaction(func(tx *database.DB) error {
			if len(bridgeInitiateList) > 0 {
				err := sh.db.BridgeInitiate.MarkedBridgeInitiateHandled(bridgeInitiateList)
				if err != nil {
					log.Error(" marked bridge initiate fail", "err", err)
					return err
				}
			}

			if len(bridgeFinalizeRecords) > 0 {
				err := sh.db.BridgeFinalize.MarkedBridgeFinalizeHandled(bridgeFinalizeRecords)
				if err != nil {
					log.Error(" marked bridge initiate fail", "err", err)
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
