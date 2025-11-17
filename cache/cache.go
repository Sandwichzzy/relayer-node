// Package cache 提供基于 LRU（Least Recently Used） LRU（最近最少使用）算法的缓存功能
// 用于缓存 Staking 和 Bridge 相关的记录，提高查询性能
package cache

import (
	"errors"

	lru "github.com/hashicorp/golang-lru"

	"github.com/Sandwichzzy/relayer-node/service/models"
)

// ListSize 定义 LRU 缓存的最大容量
// 当缓存达到此大小时，最少使用的条目将被自动淘汰
const ListSize = 1200000

// LruCache 封装了两个独立的 LRU 缓存实例
// 分别用于存储 Staking 记录和 Bridge 记录
type LruCache struct {
	lruStakingRecords *lru.Cache // Staking 记录缓存
	lruBridgeRecords  *lru.Cache // Bridge 记录缓存
}

// NewLruCache 创建并初始化一个新的 LruCache 实例
// 会同时初始化 Staking 和 Bridge 两个缓存，每个缓存容量为 ListSize
// 如果初始化失败会触发 panic
func NewLruCache() *LruCache {
	// 初始化 Staking 记录缓存
	lruStakingRecords, err := lru.New(ListSize)
	if err != nil {
		panic(errors.New("Failed to init lruStakingRecord, err :" + err.Error()))
	}
	// 初始化 Bridge 记录缓存
	lruBridgeRecords, err := lru.New(ListSize)
	if err != nil {
		panic(errors.New("Failed to init lruBridgeRecord, err :" + err.Error()))
	}
	return &LruCache{
		lruStakingRecords: lruStakingRecords,
		lruBridgeRecords:  lruBridgeRecords,
	}
}

// GetStakingRecords 从缓存中获取指定 key 的 Staking 记录
// 参数:
//   - key: 记录的唯一标识符
//
// 返回:
//   - *models.StakingResponse: Staking 记录数据
//   - error: 如果记录不存在则返回错误
func (lc *LruCache) GetStakingRecords(key string) (*models.StakingResponse, error) {
	result, ok := lc.lruStakingRecords.Get(key)
	if !ok {
		return nil, errors.New("lru get Staking records fail")
	}
	return result.(*models.StakingResponse), nil
}

// AddStakingRecords 向缓存中添加 Staking 记录
// 使用 PeekOrAdd 方法，如果 key 已存在则不更新，不存在则添加
// 参数:
//   - key: 记录的唯一标识符
//   - data: 要缓存的 Staking 记录数据
func (lc *LruCache) AddStakingRecords(key string, data *models.StakingResponse) {
	lc.lruStakingRecords.PeekOrAdd(key, data)
}

// GetBridgeRecords 从缓存中获取指定 key 的 Bridge 记录
// 参数:
//   - key: 记录的唯一标识符
//
// 返回:
//   - *models.BridgeResponse: Bridge 记录数据
//   - error: 如果记录不存在则返回错误
func (lc *LruCache) GetBridgeRecords(key string) (*models.BridgeResponse, error) {
	result, ok := lc.lruBridgeRecords.Get(key)
	if !ok {
		return nil, errors.New("lru get bridge records fail")
	}
	return result.(*models.BridgeResponse), nil
}

// AddBridgeRecords 向缓存中添加 Bridge 记录
// 使用 PeekOrAdd 方法，如果 key 已存在则不更新，不存在则添加
// 参数:
//   - key: 记录的唯一标识符
//   - data: 要缓存的 Bridge 记录数据
func (lc *LruCache) AddBridgeRecords(key string, data *models.BridgeResponse) {
	lc.lruBridgeRecords.PeekOrAdd(key, data)
}
