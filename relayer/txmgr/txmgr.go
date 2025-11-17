// Package txmgr 实现智能交易管理器
//
// 核心功能：
// 1. 交易发送和重试机制
// 2. Gas 价格动态更新（应对网络拥堵）
// 3. 交易确认等待（支持自定义确认块数）
// 4. Nonce 冲突处理
// 5. 交易卡住时自动重新提交（提高 Gas）
//
// 工作原理：
// 1. 初次发送交易
// 2. 定期检查交易状态
// 3. 如果交易未确认且超时，提高 Gas 重新发送
// 4. 等待达到指定确认块数后返回
//
// 适用场景：
// - 需要确保交易一定上链的场景
// - 网络拥堵时需要自动提高 Gas 的场景
// - 需要等待多个区块确认的场景
package txmgr

import (
	"context"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
)

// UpdateGasPriceFunc Gas 价格更新函数类型
// 用于在重试时更新交易的 Gas 价格
type UpdateGasPriceFunc = func(ctx context.Context) (*types.Transaction, error)

// SendTransactionFunc 交易发送函数类型
// 用于将交易提交到区块链网络
type SendTransactionFunc = func(ctx context.Context, tx *types.Transaction) error

// Config 交易管理器配置
type Config struct {
	ResubmissionTimeout       time.Duration // 重新提交超时时间（交易未确认时重发）
	ReceiptQueryInterval      time.Duration // 查询交易回执的间隔时间
	NumConfirmations          uint64        // 需要等待的确认块数
	SafeAbortNonceTooLowCount uint64        // Nonce 过低时安全中止的重试次数
}

// TxManager 交易管理器接口
type TxManager interface {
	// Send 发送交易并等待确认
	// updateGasPrice: Gas 价格更新函数
	// sendTxn: 交易发送函数
	// 返回：交易回执和可能的错误
	Send(ctx context.Context, updateGasPrice UpdateGasPriceFunc, sendTxn SendTransactionFunc) (*types.Receipt, error)
}

// ReceiptSource 交易回执查询接口
type ReceiptSource interface {
	BlockNumber(ctx context.Context) (uint64, error)                                    // 获取最新区块号
	TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) // 查询交易回执
}

// SimpleTxManager 简单交易管理器实现
type SimpleTxManager struct {
	cfg     Config        // 配置
	backend ReceiptSource // RPC 后端
	l       log.Logger    // 日志记录器
}

func NewSimpleTxManager(cfg Config, backend ReceiptSource) *SimpleTxManager {
	if cfg.NumConfirmations == 0 {
		panic("txmgr: NumConfirmations cannot be zero")
	}
	return &SimpleTxManager{
		cfg:     cfg,
		backend: backend,
	}
}

// Send 发送交易并管理其完整生命周期
//
// 工作流程：
// 1. 启动异步交易发送任务（sendTxAsync）
// 2. 定期重新提交交易（如果未确认且超时）
// 3. 等待交易确认（达到指定确认块数）
// 4. 返回交易回执
//
// 重试机制：
// - 如果交易在 ResubmissionTimeout 内未确认
// - 调用 updateGasPrice 提高 Gas 价格
// - 重新发送交易（相同 nonce，更高 Gas）
//
// 确认机制：
// - 等待交易被打包到区块
// - 等待达到 NumConfirmations 个区块确认
// - 确保交易不会被回滚
//
// 错误处理：
// - Nonce too low: 自动重试，超过次数后中止
// - Context canceled: 停止所有操作并返回
// - 其他错误: 记录日志并继续重试
func (m *SimpleTxManager) Send(ctx context.Context, updateGasPrice UpdateGasPriceFunc, sendTx SendTransactionFunc) (*types.Receipt, error) {
	var wg sync.WaitGroup
	defer wg.Wait()
	ctxc, cancel := context.WithCancel(ctx)
	defer cancel()

	sendState := NewSendState(m.cfg.SafeAbortNonceTooLowCount)

	receiptChan := make(chan *types.Receipt, 1)

	// sendTxAsync 异步发送交易的核心函数
	//
	// 执行流程：
	// 1. 调用 updateGasPrice 获取（或更新）交易
	// 2. 调用 sendTx 发送交易到网络
	// 3. 调用 waitMined 等待交易确认
	// 4. 如果成功，将回执发送到 receiptChan
	//
	// 错误处理：
	// - context canceled: 静默返回（正常退出）
	// - nonce too low: 记录到 sendState，可能触发中止
	// - 其他错误: 记录日志并返回（由外层定时器重试）
	//
	// 并发说明：
	// - 此函数会被多次并发调用（定时器触发）
	// - 通过 WaitGroup 确保所有 goroutine 完成后才返回
	sendTxAsync := func() {
		defer wg.Done()

		// 步骤1: 获取（或更新）交易
		// - 第一次调用：构造新交易
		// - 后续调用：更新 Gas 价格（提高 Gas 以替换原交易）
		tx, err := updateGasPrice(ctxc)
		if err != nil {
			if err == context.Canceled || strings.Contains(err.Error(), "context canceled") {
				return
			}
			log.Error("ContractsCaller update txn gas price fail", "err", err)
			cancel()  // 取消所有操作
			return
		}

		txHash := tx.Hash()
		nonce := tx.Nonce()
		gasTipCap := tx.GasTipCap()
		gasFeeCap := tx.GasFeeCap()
		log.Debug("ContractsCaller publishing transaction", "txHash", txHash, "nonce", nonce, "gasTipCap", gasTipCap, "gasFeeCap", gasFeeCap)

		// 步骤2: 发送交易到网络
		err = sendTx(ctxc, tx)
		sendState.ProcessSendError(err)  // 记录错误（如 nonce too low）
		if err != nil {
			if err == context.Canceled || strings.Contains(err.Error(), "context canceled") {
				return
			}
			log.Error("ContractsCaller unable to publish transaction", "err", err)

			// 关键：如果应该中止，取消上下文
			// 这会停止所有 goroutine 并返回错误
			if sendState.ShouldAbortImmediately() {
				cancel()
			}
			return
		}

		log.Debug("ContractsCaller transaction published successfully", "hash", txHash, "nonce", nonce, "gasTipCap", gasTipCap, "gasFeeCap", gasFeeCap)

		// 步骤3: 等待交易确认
		receipt, err := waitMined(
			ctxc, m.backend, tx, m.cfg.ReceiptQueryInterval,
			m.cfg.NumConfirmations, sendState,
		)
		if err != nil {
			log.Debug("ContractsCaller send tx failed", "hash", txHash, "nonce", nonce, "gasTipCap", gasTipCap, "gasFeeCap", gasFeeCap, "err", err)
		}

		// 步骤4: 如果成功，发送回执到主循环
		if receipt != nil {
			select {
			case receiptChan <- receipt:
				log.Trace("ContractsCaller send tx succeeded", "hash", txHash,
					"nonce", nonce, "gasTipCap", gasTipCap,
					"gasFeeCap", gasFeeCap)
			default:
				// 通道已满（已有其他交易成功），忽略
			}
		}
	}

	// 启动第一个发送任务
	wg.Add(1)
	go sendTxAsync()

	// 重新提交定时器
	ticker := time.NewTicker(m.cfg.ResubmissionTimeout)
	defer ticker.Stop()

	// 主事件循环
	// 监听三个事件：
	// 1. 定时器触发：重新提交交易（如果还未确认）
	// 2. 上下文取消：停止所有操作
	// 3. 收到回执：交易成功确认，返回结果
	for {
		select {
		case <-ticker.C:
			// 定时器触发：检查是否需要重新提交交易
			//
			// 重新提交的条件：
			// - 交易还未被打包（IsWaitingForConfirmation = false）
			// - 如果已经被打包，只需等待确认，不需要重新提交
			//
			// 重新提交的作用：
			// - 提高 Gas 价格以加速交易确认
			// - 应对网络拥堵情况
			if sendState.IsWaitingForConfirmation() {
				continue  // 已经被打包，跳过重新提交
			}
			wg.Add(1)
			go sendTxAsync()  // 启动新的发送任务（Gas 会更高）

		case <-ctxc.Done():
			// 上下文被取消（可能是用户取消，或发生严重错误）
			return nil, ctxc.Err()

		case receipt := <-receiptChan:
			// 收到交易回执：交易成功确认
			return receipt, nil
		}
	}
}

// WaitMined 等待交易被打包并确认（公开函数）
//
// 简化版本，不使用 SendState（用于测试或简单场景）
//
// 参数：
//   - ctx: 上下文
//   - backend: RPC 后端（查询交易回执和区块号）
//   - tx: 要等待的交易
//   - queryInterval: 查询间隔（建议 1-3 秒）
//   - numConfirmations: 需要等待的确认块数（建议 3-12）
//
// 返回：
//   - *types.Receipt: 交易回执（包含状态、Gas 使用量等）
//   - error: 错误信息
func WaitMined(
	ctx context.Context,
	backend ReceiptSource,
	tx *types.Transaction,
	queryInterval time.Duration,
	numConfirmations uint64,
) (*types.Receipt, error) {
	return waitMined(ctx, backend, tx, queryInterval, numConfirmations, nil)
}

// waitMined 等待交易被打包并达到指定确认数（内部函数）
//
// 工作原理：
// 1. 定期查询交易回执（queryInterval）
// 2. 如果交易已打包，检查确认数
// 3. 确认数 = (当前区块高度 - 交易所在区块) + 1
// 4. 达到 numConfirmations 后返回
//
// 确认数说明：
// - 1 确认：交易刚被打包到区块
// - 3 确认：交易所在区块后又产生了 2 个区块
// - 12 确认：交易所在区块后又产生了 11 个区块（以太坊主网推荐）
//
// 为什么需要确认：
// - 区块链可能发生重组（分叉）
// - 更多确认 = 更安全（交易被回滚的概率更低）
// - 不同场景需要不同的确认数：
//   * 小额交易：1-3 确认
//   * 中等金额：6 确认
//   * 大额交易：12+ 确认
func waitMined(
	ctx context.Context,
	backend ReceiptSource,
	tx *types.Transaction,
	queryInterval time.Duration,
	numConfirmations uint64,
	sendState *SendState,
) (*types.Receipt, error) {

	queryTicker := time.NewTicker(queryInterval)
	defer queryTicker.Stop()

	txHash := tx.Hash()

	for {
		// 查询交易回执
		receipt, err := backend.TransactionReceipt(ctx, txHash)
		switch {
		case receipt != nil:
			// 情况1: 交易已被打包

			// 通知 SendState（如果存在）
			if sendState != nil {
				sendState.TxMined(txHash)
			}

			// 获取交易所在区块高度和当前链高度
			txHeight := receipt.BlockNumber.Uint64()
			tipHeight, err := backend.BlockNumber(ctx)
			if err != nil {
				log.Error("ContractsCaller Unable to fetch block number", "err", err)
				break  // 继续重试查询
			}

			log.Trace("ContractsCaller Transaction mined, checking confirmations",
				"txHash", txHash, "txHeight", txHeight,
				"tipHeight", tipHeight,
				"numConfirmations", numConfirmations)

			// 检查确认数是否足够
			// 公式：txHeight + numConfirmations <= tipHeight + 1
			// 例如：txHeight=100, numConfirmations=3, tipHeight=102
			//      100 + 3 <= 102 + 1  =>  103 <= 103  =>  true（已确认）
			if txHeight+numConfirmations <= tipHeight+1 {
				log.Debug("ContractsCaller Transaction confirmed", "txHash", txHash)
				return receipt, nil  // 确认完成，返回回执
			}

			// 计算还需要多少确认
			confsRemaining := (txHeight + numConfirmations) - (tipHeight + 1)
			log.Info("ContractsCaller Transaction not yet confirmed", "txHash", txHash,
				"confsRemaining", confsRemaining)

		case err != nil:
			// 情况2: 查询出错（网络问题、节点问题等）
			log.Trace("ContractsCaller Receipt retrieve failed", "hash", txHash,
				"err", err)

		default:
			// 情况3: 交易尚未被打包（receipt 为 nil 且无错误）

			// 通知 SendState（如果存在）
			if sendState != nil {
				sendState.TxNotMined(txHash)
			}
			log.Trace("ContractsCaller Transaction not yet mined", "hash", txHash)
		}

		// 等待下一次查询
		select {
		case <-ctx.Done():
			return nil, ctx.Err()  // 上下文取消
		case <-queryTicker.C:
			// 继续下一轮查询
		}
	}
}

// CalcGasFeeCap 计算 EIP-1559 的最大 Gas 费用上限
//
// EIP-1559 Gas 费用结构：
// - baseFee: 基础费用（由网络动态调整，会被销毁）
// - priorityFee (tip): 小费（给矿工的奖励）
// - maxFeePerGas: 用户愿意支付的最高费用
//
// 计算公式：
//   maxFeePerGas = priorityFee + (baseFee * 2)
//
// 为什么是 baseFee * 2：
// - baseFee 可能会在下个区块上涨（最多 12.5%）
// - 乘以 2 提供足够的缓冲空间
// - 确保交易在下一个区块仍然有效
//
// 实际支付费用：
//   实际费用 = baseFee + priorityFee（不会超过 maxFeePerGas）
//
// 例子：
//   baseFee = 50 Gwei
//   priorityFee = 2 Gwei
//   maxFeePerGas = 2 + (50 * 2) = 102 Gwei
//   实际支付 = 50 + 2 = 52 Gwei（退还 50 Gwei）
func CalcGasFeeCap(baseFee, gasTipCap *big.Int) *big.Int {
	return new(big.Int).Add(
		gasTipCap,
		new(big.Int).Mul(baseFee, big.NewInt(2)),
	)
}
