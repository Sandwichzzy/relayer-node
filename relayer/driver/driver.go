// Package driver 实现跨链交易的构造和发送引擎
// 职责：
// 1. 与目标链的 PoolManager 合约交互
// 2. 构造 bridgeFinalize 交易
// 3. 管理交易的 Gas 价格
// 4. 通过 TxManager 发送交易并等待确认
// 5. 收集和记录链上指标（nonce、余额等）
//
// 核心流程：
// 1. 构造交易（bridgeFinalizeETH 或 bridgeFinalizeERC20）
// 2. 设置 Gas 价格（支持 EIP-1559）
// 3. 通过 TxManager 发送交易
// 4. 等待交易确认
// 5. 返回交易回执
package driver

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"

	"github.com/Sandwichzzy/relayer-node/bindings"
	"github.com/Sandwichzzy/relayer-node/metrics"
	"github.com/Sandwichzzy/relayer-node/relayer/txmgr"
)

var (
	errMaxPriorityFeePerGasNotFound = errors.New("Method eth_maxPriorityFeePerGas not found")
	FallbackGasTipCap               = big.NewInt(1500000000) // 1.5 Gwei 备用 Gas 小费
)

// DriverEngineConfig 驱动引擎配置
type DriverEngineConfig struct {
	ChainClient               *ethclient.Client // 链的 RPC 客户端
	ChainId                   *big.Int          // 链 ID
	PoolManagerAddress        common.Address    // PoolManager 合约地址
	CallerAddress             common.Address    // Relayer 钱包地址
	PrivateKey                *ecdsa.PrivateKey // Relayer 私钥
	NumConfirmations          uint64            // 交易确认块数
	SafeAbortNonceTooLowCount uint64            // Nonce 过低时安全中止的重试次数
}

// DriverEngine 交易驱动引擎
// 负责构造和发送跨链完成交易
type DriverEngine struct {
	Ctx                    context.Context         // 上下文
	Cfg                    *DriverEngineConfig     // 配置
	PoolManagerContract    *bindings.PoolManager   // PoolManager 合约绑定
	RawPoolManagerContract *bind.BoundContract     // 原始合约绑定（用于更新 Gas）
	PoolManagerContractAbi *abi.ABI                // 合约 ABI
	TxMgr                  txmgr.TxManager         // 交易管理器
	RelayerMetrics         *metrics.SyncerMetrics  // 监控指标
	cancel                 func()                  // 取消函数
}

// NewDriverEngine 创建交易驱动引擎实例
//
// 初始化步骤：
// 1. 绑定 PoolManager 合约（用于调用 bridgeFinalize 方法）
// 2. 解析合约 ABI（用于构造交易数据）
// 3. 创建原始合约绑定（用于更新 Gas 价格）
// 4. 配置交易管理器（处理交易重试、确认等）
//
// 参数：
//   - ctx: 上下文
//   - relayerMetrics: 监控指标收集器
//   - cfg: 驱动引擎配置（包含链信息、合约地址、私钥等）
//
// 返回：
//   - *DriverEngine: 驱动引擎实例
//   - error: 初始化错误
func NewDriverEngine(ctx context.Context, relayerMetrics *metrics.SyncerMetrics, cfg *DriverEngineConfig) (*DriverEngine, error) {
	_, cancel := context.WithTimeout(ctx, time.Second*15)
	defer cancel()

	// 1. 创建 PoolManager 合约绑定（Go 类型安全的合约调用接口）
	poolManagerContract, err := bindings.NewPoolManager(cfg.PoolManagerAddress, cfg.ChainClient)
	if err != nil {
		log.Error("new pool manager fail", "err", err)
		return nil, err
	}

	// 2. 解析合约 ABI（Application Binary Interface）
	// ABI 定义了合约的函数签名和参数类型
	parsed, err := abi.JSON(strings.NewReader(bindings.PoolManagerMetaData.ABI))
	if err != nil {
		log.Error("parsed abi fail", "err", err)
		return nil, err
	}

	poolManagerContractAbi, err := bindings.PoolManagerMetaData.GetAbi()
	if err != nil {
		log.Error("get pool manager meta data fail", "err", err)
		return nil, err
	}

	// 3. 创建原始合约绑定（用于更灵活地控制交易参数，特别是 Gas 价格）
	rawPoolManagerContract := bind.NewBoundContract(cfg.PoolManagerAddress, parsed, cfg.ChainClient, cfg.ChainClient, cfg.ChainClient)

	// 4. 配置交易管理器
	txManagerConfig := txmgr.Config{
		ResubmissionTimeout:       time.Second * 5, // 5秒内未确认则重新提交
		ReceiptQueryInterval:      time.Second,     // 每秒查询一次交易回执
		NumConfirmations:          cfg.NumConfirmations,          // 需要等待的确认块数
		SafeAbortNonceTooLowCount: cfg.SafeAbortNonceTooLowCount, // Nonce 过低重试次数
	}

	txManager := txmgr.NewSimpleTxManager(txManagerConfig, cfg.ChainClient)

	return &DriverEngine{
		Ctx:                    ctx,
		Cfg:                    cfg,
		PoolManagerContract:    poolManagerContract,
		RawPoolManagerContract: rawPoolManagerContract,
		PoolManagerContractAbi: poolManagerContractAbi,
		TxMgr:                  txManager,
		RelayerMetrics:         relayerMetrics,
		cancel:                 cancel,
	}, nil
}

// UpdateGasPrice 更新交易的 Gas 价格
//
// 应用场景：
// - 交易在 mempool 中长时间未被打包
// - 需要提高 Gas 价格以加速交易确认
//
// 工作原理：
// 1. 使用相同的 nonce 重新构造交易
// 2. 让以太坊客户端自动计算新的 Gas 价格（基于当前网络状况）
// 3. 如果链不支持 EIP-1559，则使用备用 Gas 小费（1.5 Gwei）
//
// 参数：
//   - ctx: 上下文
//   - tx: 原始交易
//
// 返回：
//   - *types.Transaction: 更新 Gas 后的新交易
//   - error: 错误信息
func (de *DriverEngine) UpdateGasPrice(ctx context.Context, tx *types.Transaction) (*types.Transaction, error) {
	var opts *bind.TransactOpts
	var err error
	opts, err = bind.NewKeyedTransactorWithChainID(de.Cfg.PrivateKey, de.Cfg.ChainId)
	if err != nil {
		log.Error("new keyed transactor with chain id fail", "err", err)
		return nil, err
	}
	opts.Context = ctx
	opts.Nonce = new(big.Int).SetUint64(tx.Nonce())  // 使用相同的 nonce（替换原交易）
	opts.NoSend = true                               // 只构造交易，不实际发送

	// 使用原交易的 Data，但让客户端重新计算 Gas 价格
	finalTx, err := de.RawPoolManagerContract.RawTransact(opts, tx.Data())
	switch {
	case err == nil:
		return finalTx, nil

	// 处理不支持 EIP-1559 的链（如某些 L2 或旧版以太坊）
	case de.isMaxPriorityFeePerGasNotFoundError(err):
		log.Info("Don't support priority fee")
		opts.GasTipCap = FallbackGasTipCap  // 使用固定的 1.5 Gwei 小费
		return de.RawPoolManagerContract.RawTransact(opts, tx.Data())

	default:
		return nil, err
	}
}

// isMaxPriorityFeePerGasNotFoundError 判断是否为不支持 EIP-1559 的错误
//
// EIP-1559 说明：
// - EIP-1559 是以太坊的 Gas 费改进提案
// - 引入了 baseFee 和 priorityFee（小费）机制
// - 部分链（如某些 L2）可能不支持此特性
func (de *DriverEngine) isMaxPriorityFeePerGasNotFoundError(err error) bool {
	return strings.Contains(err.Error(), errMaxPriorityFeePerGasNotFound.Error())
}

// SendTransaction 发送交易到区块链网络
func (de *DriverEngine) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	return de.Cfg.ChainClient.SendTransaction(ctx, tx)
}

// bridgeFinalizeETH 构造 ETH 跨链完成交易（内部方法）
//
// 功能：
// 1. 获取 Relayer 账户的 nonce（交易序号）
// 2. 查询 Relayer 账户余额（确保有足够的 Gas 费）
// 3. 记录监控指标（nonce、余额）
// 4. 构造 bridgeFinalizeETH 交易（但不发送）
//
// 参数：
//   - sourceChainId: 源链 ID
//   - destChainId: 目标链 ID
//   - sourceTokenAddress: 源链代币地址
//   - from: 跨链发起者地址
//   - to: 跨链接收者地址
//   - amount: 跨链金额
//   - gasFee: Gas 费用
//   - msgNonce: 消息 nonce（用于防重放攻击）
//
// 返回：
//   - *types.Transaction: 构造的交易（未签名、未发送）
//   - error: 错误信息
func (de *DriverEngine) bridgeFinalizeETH(ctx context.Context, sourceChainId *big.Int, destChainId *big.Int, sourceTokenAddress common.Address, from common.Address, to common.Address, amount *big.Int, gasFee *big.Int, msgNonce *big.Int) (*types.Transaction, error) {
	// 1. 获取 Relayer 账户的当前 nonce
	nonce, err := de.Cfg.ChainClient.NonceAt(ctx, de.Cfg.CallerAddress, nil)
	if err != nil {
		log.Error("get nonce error", "err", err)
		return nil, err
	}
	de.RelayerMetrics.RecordChainAddressNonce(de.Cfg.ChainId.String(), big.NewInt(int64(nonce)))
	log.Info("bridge finalized eth tx nonce", "CallerAddress", de.Cfg.CallerAddress, "nonce", nonce)

	// 2. 查询 Relayer 账户余额（确保有足够的 ETH 支付 Gas）
	balance, err := de.Cfg.ChainClient.BalanceAt(ctx, de.Cfg.CallerAddress, nil)
	if err != nil {
		log.Error("get nonce error", "err", err)
		return nil, err
	}
	de.RelayerMetrics.RecordNativeTokenBalance(de.Cfg.ChainId.String(), balance)
	log.Info("bridge eth balance", "CallerAddress", de.Cfg.CallerAddress, "balance", balance.String())

	// 3. 创建交易选项（使用 Relayer 私钥签名）
	opts, err := bind.NewKeyedTransactorWithChainID(de.Cfg.PrivateKey, de.Cfg.ChainId)
	if err != nil {
		log.Error("new keyed transactor with chain id fail", "err", err)
		return nil, err
	}
	opts.Context = ctx
	opts.Nonce = new(big.Int).SetUint64(nonce)
	opts.NoSend = true  // 只构造交易，不立即发送（由 TxManager 管理发送）

	// 4. 调用 PoolManager 合约的 bridgeFinalizeETH 方法
	tx, err := de.PoolManagerContract.BridgeFinalizeETH(opts, sourceChainId, destChainId, sourceTokenAddress, from, to, amount, gasFee, msgNonce)
	switch {
	case err == nil:
		return tx, nil

	// 如果链不支持 EIP-1559，使用备用 Gas 小费
	case de.isMaxPriorityFeePerGasNotFoundError(err):
		opts.GasTipCap = FallbackGasTipCap
		return de.PoolManagerContract.BridgeFinalizeETH(opts, sourceChainId, destChainId, sourceTokenAddress, from, to, amount, gasFee, msgNonce)

	default:
		return nil, err
	}
}

// bridgeFinalizeERC20 构造 ERC20 代币跨链完成交易（内部方法）
//
// 功能：
// 1. 获取 Relayer 账户的 nonce
// 2. 构造 bridgeFinalizeERC20 交易
//
// 与 bridgeFinalizeETH 的区别：
// - ETH 是原生代币，直接转账
// - ERC20 需要调用代币合约的 transfer 方法
// - 需要指定源链和目标链的代币合约地址
//
// 参数：
//   - sourceTokenAddress: 源链 ERC20 代币地址
//   - destTokenAddress: 目标链 ERC20 代币地址
//   - 其他参数同 bridgeFinalizeETH
func (de *DriverEngine) bridgeFinalizeERC20(ctx context.Context, sourceChainId *big.Int, destChainId *big.Int, from common.Address, to common.Address, sourceTokenAddress common.Address, destTokenAddress common.Address, amount *big.Int, gasFee *big.Int, msgNonce *big.Int) (*types.Transaction, error) {
	nonce, err := de.Cfg.ChainClient.NonceAt(ctx, de.Cfg.CallerAddress, nil)
	if err != nil {
		log.Error("get nonce error", "err", err)
		return nil, err
	}
	log.Info("bridge finalized erc20 tx nonce", "CallerAddress", de.Cfg.CallerAddress, "nonce", nonce)

	opts, err := bind.NewKeyedTransactorWithChainID(de.Cfg.PrivateKey, de.Cfg.ChainId)
	if err != nil {
		log.Error("new keyed transactor with chain id fail", "err", err)
		return nil, err
	}
	opts.Context = ctx
	opts.Nonce = new(big.Int).SetUint64(nonce)
	opts.NoSend = true

	tx, err := de.PoolManagerContract.BridgeFinalizeERC20(opts, sourceChainId, destChainId, from, to, sourceTokenAddress, destTokenAddress, amount, gasFee, msgNonce)
	switch {
	case err == nil:
		return tx, nil

	case de.isMaxPriorityFeePerGasNotFoundError(err):
		opts.GasTipCap = FallbackGasTipCap
		return de.PoolManagerContract.BridgeFinalizeERC20(opts, sourceChainId, destChainId, from, to, sourceTokenAddress, destTokenAddress, amount, gasFee, msgNonce)

	default:
		return nil, err
	}
}

// BridgeFinalizeETH 完成 ETH 跨链（公开方法）
//
// 完整流程：
// 1. 调用 bridgeFinalizeETH 构造交易
// 2. 通过 TxManager 发送交易
// 3. TxManager 负责：
//    - 交易重试（如果未确认）
//    - Gas 价格更新（通过 UpdateGasPrice）
//    - 等待确认（达到指定区块数）
// 4. 返回交易回执
//
// 这是 Relayer 在目标链上完成跨链的核心方法
func (de *DriverEngine) BridgeFinalizeETH(sourceChainId *big.Int, destChainId *big.Int, sourceTokenAddress common.Address, from common.Address, to common.Address, amount *big.Int, gasFee *big.Int, msgNonce *big.Int) (*types.Receipt, error) {
	// 1. 构造交易
	tx, err := de.bridgeFinalizeETH(de.Ctx, sourceChainId, destChainId, sourceTokenAddress, from, to, amount, gasFee, msgNonce)
	if err != nil {
		log.Error("build bridge finalized tx fail", "err", err)
		return nil, err
	}

	// 2. 定义 Gas 更新函数（TxManager 在重试时会调用）
	updateGasPrice := func(ctx context.Context) (*types.Transaction, error) {
		return de.UpdateGasPrice(ctx, tx)
	}

	log.Info("build bridge finalized eth tx", "txHash", tx.Hash().String(), "data", tx.Data())

	// 3. 通过 TxManager 发送交易并等待确认
	receipt, err := de.TxMgr.Send(de.Ctx, updateGasPrice, de.SendTransaction)
	if err != nil {
		log.Error("send tx fail", "err", err)
		return nil, err
	}
	return receipt, nil
}

// BridgeFinalizeErc20 完成 ERC20 代币跨链（公开方法）
//
// 工作流程同 BridgeFinalizeETH，但处理 ERC20 代币
func (de *DriverEngine) BridgeFinalizeErc20(sourceChainId *big.Int, destChainId *big.Int, from common.Address, to common.Address, sourceTokenAddress common.Address, destTokenAddress common.Address, amount *big.Int, gasFee *big.Int, msgNonce *big.Int) (*types.Receipt, error) {
	tx, err := de.bridgeFinalizeERC20(de.Ctx, sourceChainId, destChainId, from, to, sourceTokenAddress, destTokenAddress, amount, gasFee, msgNonce)
	if err != nil {
		log.Error("build bridge finalized tx fail", "err", err)
		return nil, err
	}
	updateGasPrice := func(ctx context.Context) (*types.Transaction, error) {
		return de.UpdateGasPrice(ctx, tx)
	}

	log.Info("build bridge finalized erc20 tx", "txHash", tx.Hash().String(), "data", tx.Data())

	receipt, err := de.TxMgr.Send(de.Ctx, updateGasPrice, de.SendTransaction)
	if err != nil {
		log.Error("send tx fail", "err", err)
		return nil, err
	}
	return receipt, nil
}
