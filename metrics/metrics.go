// Package metrics 实现基于 Prometheus 的监控指标收集系统
// 用于监控跨链桥接节点的同步状态、区块高度、代币余额等关键指标
//
// 主要功能：
// 1. 记录不同链的区块高度（链头高度和事件处理高度）
// 2. 监控各链上的原生代币余额
// 3. 跟踪地址的 nonce 值
//
// 使用场景：
// - 通过 Prometheus 和 Grafana 实时监控节点运行状态
// - 告警配置：区块同步落后、余额不足等
// - 性能分析和问题诊断
package metrics

import (
	"math/big"

	"github.com/prometheus/client_golang/prometheus"
)

// SyncerMetricer 定义同步器指标收集接口
// 任何需要记录同步指标的组件都应实现此接口
type SyncerMetricer interface {
	RecordChainBlockHeight(chainId string, height *big.Int)    // 记录链的最新区块高度
	RecordEventBlockHeight(chainId string, height *big.Int)    // 记录事件处理到的区块高度
	RecordNativeTokenBalance(chainId string, balance *big.Int) // 记录原生代币余额
	RecordChainAddressNonce(chainId string, nonce *big.Int)    // 记录地址的 nonce 值
}

// SyncerMetrics 同步器指标收集器
// 使用 Prometheus GaugeVec 类型指标，支持多链监控（通过 chain_id 标签区分）
type SyncerMetrics struct {
	chainBlockHeight   *prometheus.GaugeVec // 各链的最新区块高度
	eventBlockHeight   *prometheus.GaugeVec // 各链事件处理到的区块高度
	nativeTokenBalance *prometheus.GaugeVec // 各链上账户的原生代币余额（如 ETH、BNB）
	chainAddressNonce  *prometheus.GaugeVec // 各链上账户的 nonce 值
}

// NewSyncerMetrics 创建并初始化同步器指标收集器
//
// 参数：
//   - registry: Prometheus 注册表，用于注册指标
//   - subsystem: 子系统名称，用于指标命名空间（如 "syncer"、"bridge"）
//
// 返回：
//   - *SyncerMetrics: 初始化完成的指标收集器
//
// 创建的指标：
//  1. chain_block_height: 各链的最新区块高度（Gauge 类型）
//  2. event_block_height: 各链事件处理到的区块高度（Gauge 类型）
//  3. native_token_balance: 各链上的原生代币余额（Gauge 类型）
//  4. chain_address_nonce: 各链上的地址 nonce（Gauge 类型）
//
// 所有指标都使用 chain_id 作为标签，支持多链监控
func NewSyncerMetrics(registry *prometheus.Registry, subsystem string) *SyncerMetrics {
	// 创建区块高度指标（链的最新高度）
	chainBlockHeight := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name:      "chain_block_height",
		Help:      "Different chain block height",
		Subsystem: subsystem,
	}, []string{"chain_id"})

	// 创建事件处理区块高度指标（同步器处理到的高度）
	eventBlockHeight := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name:      "event_block_height",
		Help:      "Different chain event block height",
		Subsystem: subsystem,
	}, []string{"chain_id"})

	// 创建原生代币余额指标
	nativeTokenBalance := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name:      "native_token_balance",
		Help:      "Different chain native token balance",
		Subsystem: subsystem,
	}, []string{"chain_id"})

	// 创建地址 nonce 指标
	chainAddressNonce := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name:      "chain_address_nonce",
		Help:      "Different chain address nonce",
		Subsystem: subsystem,
	}, []string{"chain_id"})

	// 将所有指标注册到 Prometheus 注册表
	// MustRegister 会在注册失败时 panic
	registry.MustRegister(chainBlockHeight)
	registry.MustRegister(eventBlockHeight)
	registry.MustRegister(nativeTokenBalance)
	registry.MustRegister(chainAddressNonce)

	return &SyncerMetrics{
		chainBlockHeight:   chainBlockHeight,
		eventBlockHeight:   eventBlockHeight,
		nativeTokenBalance: nativeTokenBalance,
		chainAddressNonce:  chainAddressNonce,
	}
}

// RecordChainBlockHeight 记录指定链的最新区块高度
//
// 参数：
//   - chainId: 链 ID（如 "1" 代表以太坊主网，"56" 代表 BSC）
//   - height: 区块高度（big.Int 类型，支持大整数）
//
// 用途：
//   - 监控链的同步状态
//   - 检测链是否正常出块
//   - 通过与 eventBlockHeight 对比，计算同步延迟
func (rm *SyncerMetrics) RecordChainBlockHeight(chainId string, height *big.Int) {
	rm.chainBlockHeight.WithLabelValues(chainId).Set(float64(height.Uint64()))
}

// RecordEventBlockHeight 记录事件处理到的区块高度
//
// 参数：
//   - chainId: 链 ID
//   - height: 已处理的区块高度
//
// 用途：
//   - 监控同步器的处理进度
//   - 与 chainBlockHeight 对比，计算同步落后的区块数
//   - 告警：如果落后太多区块，可能表示同步出现问题
func (rm *SyncerMetrics) RecordEventBlockHeight(chainId string, height *big.Int) {
	rm.eventBlockHeight.WithLabelValues(chainId).Set(float64(height.Uint64()))
}

// RecordNativeTokenBalance 记录指定链上的原生代币余额
//
// 参数：
//   - chainId: 链 ID
//   - balance: 代币余额（单位：Wei 或最小单位）
//
// 用途：
//   - 监控账户余额，确保有足够的 Gas 费
//   - 告警：余额不足时及时充值
//   - 财务统计和成本分析
//
// 示例：
//   - 以太坊主网：监控 ETH 余额
//   - BSC：监控 BNB 余额
func (rm *SyncerMetrics) RecordNativeTokenBalance(chainId string, balance *big.Int) {
	rm.nativeTokenBalance.WithLabelValues(chainId).Set(float64(balance.Uint64()))
}

// RecordChainAddressNonce 记录指定链上地址的 nonce 值
//
// 参数：
//   - chainId: 链 ID
//   - nonce: 地址的交易计数（nonce）
//
// 用途：
//   - 监控交易发送情况
//   - 检测 nonce 卡住或跳跃问题
//   - 帮助诊断交易失败原因
//
// 说明：
//   - nonce 是以太坊地址发送交易的顺序号
//   - 每发送一笔交易，nonce 增加 1
//   - nonce 不连续会导致交易失败
func (rm *SyncerMetrics) RecordChainAddressNonce(chainId string, nonce *big.Int) {
	rm.chainAddressNonce.WithLabelValues(chainId).Set(float64(nonce.Uint64()))
}
