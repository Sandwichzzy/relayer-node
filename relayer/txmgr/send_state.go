// Package txmgr 交易状态管理模块
// 负责跟踪交易发送状态，处理 nonce 冲突等异常情况
package txmgr

import (
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
)

// SendState 交易发送状态管理器
//
// 职责：
// 1. 跟踪已被打包的交易（minedTxs）
// 2. 统计 "nonce too low" 错误次数
// 3. 判断是否应该安全中止（避免无限重试）
// 4. 判断是否正在等待交易确认
//
// 使用场景：
// - 当交易发送失败时，记录错误类型
// - 当交易被打包时，标记为已确认
// - 当链发生重组时，处理交易回滚
//
// 线程安全：所有方法都使用互斥锁保护
type SendState struct {
	minedTxs         map[common.Hash]struct{} // 已被打包的交易哈希集合
	nonceTooLowCount uint64                   // "nonce too low" 错误累计次数
	mu               sync.RWMutex             // 读写锁（保护并发访问）

	safeAbortNonceTooLowCount uint64 // 安全中止阈值（达到此次数后停止重试）
}

// NewSendState 创建新的发送状态管理器
// 参数：
//   - safeAbortNonceTooLowCount: 安全中止阈值（通常设置为 3-5）
//
// 作用：
//
//	当 "nonce too low" 错误达到此阈值时，停止重试以避免无限循环
//
// 例如：如果设置为 3，则连续 3 次 nonce too low 后会中止
func NewSendState(safeAbortNonceTooLowCount uint64) *SendState {
	if safeAbortNonceTooLowCount == 0 {
		panic("txmgr: safeAbortNonceTooLowCount cannot be zero")
	}

	return &SendState{
		minedTxs:                  make(map[common.Hash]struct{}),
		nonceTooLowCount:          0,
		safeAbortNonceTooLowCount: safeAbortNonceTooLowCount,
	}
}

// ProcessSendError 处理交易发送错误
//
// 功能：
// 1. 检查错误是否为 "nonce too low"
// 2. 如果是，累计错误次数
// 3. 其他错误类型忽略
//
// "Nonce Too Low" 错误说明：
//   - 表示使用的 nonce 已经被其他交易使用了
//   - 可能原因：
//     a. 同一账户发送了多笔交易（并发）
//     b. 之前的交易已经被确认，但我们还在用旧的 nonce
//     c. 其他进程/服务也在使用同一账户
//
// 为什么需要计数：
// - 如果连续多次出现此错误，说明可能存在严重问题
// - 需要安全中止，避免无限重试浪费资源
func (s *SendState) ProcessSendError(err error) {
	if err == nil {
		return
	}

	// 只关注 "nonce too low" 错误
	if !strings.Contains(err.Error(), core.ErrNonceTooLow.Error()) {
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.nonceTooLowCount++
}

// TxMined 标记交易已被打包到区块
//
// 调用时机：
// - 当查询到交易回执（receipt）时
// - 表示交易已经被矿工打包到区块中
//
// 作用：
// 1. 记录已打包的交易哈希
// 2. 表明至少有一笔交易成功了（即使后续有 nonce too low 也不应中止）
func (s *SendState) TxMined(txHash common.Hash) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.minedTxs[txHash] = struct{}{}
}

// TxNotMined 标记交易未被打包（或因链重组而回滚）
//
// 调用时机：
// - 链发生重组（reorganization），之前打包的交易被回滚
// - 或者查询交易回执失败
//
// 重要逻辑：
//  1. 从已打包集合中移除该交易
//  2. 如果这是最后一笔已打包的交易被移除，重置 nonceTooLowCount
//     （因为之前认为交易成功了，现在发现失败，需要重新计数）
//
// 链重组说明：
// - 区块链可能因为分叉而发生重组
// - 之前在旧链上的交易可能在新链上不存在
// - 需要重新发送交易
func (s *SendState) TxNotMined(txHash common.Hash) {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, wasMined := s.minedTxs[txHash]
	delete(s.minedTxs, txHash)

	// 关键：如果所有已打包的交易都被回滚了，重置错误计数
	// 这样可以重新尝试发送交易
	if len(s.minedTxs) == 0 && wasMined {
		s.nonceTooLowCount = 0
	}
}

// ShouldAbortImmediately 判断是否应该立即中止交易发送
//
// 返回 true 的条件：
// 1. 没有任何交易被打包（minedTxs 为空）
// 2. 且 nonceTooLowCount >= safeAbortNonceTooLowCount
//
// 逻辑解释：
// - 如果有交易已经被打包，说明系统正常工作，不应中止
// - 如果没有交易被打包，且连续多次 nonce too low，说明有严重问题
//
// 使用场景：
// - TxManager 在发送交易失败后会检查此标志
// - 如果为 true，停止重试并返回错误
func (s *SendState) ShouldAbortImmediately() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if len(s.minedTxs) > 0 {
		return false
	}
	return s.nonceTooLowCount >= s.safeAbortNonceTooLowCount
}

// IsWaitingForConfirmation 判断是否正在等待交易确认
//
// 返回 true 的条件：
// - minedTxs 不为空（有交易已被打包）
//
// 作用：
// - TxManager 根据此标志决定是否继续重新提交交易
// - 如果正在等待确认，不需要重新提交（只需等待更多区块确认）
func (s *SendState) IsWaitingForConfirmation() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.minedTxs) > 0
}
