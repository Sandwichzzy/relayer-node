// Package routes 实现 HTTP 路由和请求处理层
// 职责：
// 1. 处理 HTTP 请求和响应
// 2. 参数验证和解析
// 3. 缓存管理（查询和更新）
// 4. 调用服务层处理业务逻辑
// 5. 统一的 JSON 响应格式
package routes

import (
	"fmt"

	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/log"

	"github.com/Sandwichzzy/relayer-node/service/models"
)

// Routes 层架构说明：
// - routes.go: 定义 Routes 结构体，持有路由器、服务实例和缓存
// - bridge_tx.go: 实现具体的 HTTP Handler
//   * StakingRecordsHandler: 查询 Staking 质押记录
//   * BridgeRecordsHandler: 查询跨链桥接记录
//   * BridgeGasFeeHandler: 查询 Gas 费用
//   * BridgeValidHandler: 验证桥接地址有效性
//   * StakingValidHandler: 验证 Staking 地址有效性
// - common.go: 提供 jsonResponse 工具函数，统一 JSON 响应格式

// StakingRecordsHandler 处理 Staking 质押记录查询请求
//
// HTTP 端点：GET /api/v1/staking-records
// 查询参数：
//   - address: 用户钱包地址（必需）
//   - page: 页码，默认 1
//   - pageSize: 每页记录数，默认 10
//   - order: 排序方式（asc/desc）
//
// 处理流程：
//  1. 解析并验证查询参数
//  2. 构造缓存键（基于查询参数）
//  3. 如果启用缓存，先尝试从缓存获取数据
//  4. 缓存未命中时，调用服务层从数据库查询
//  5. 将查询结果写入缓存（如果启用）
//  6. 返回 JSON 格式的响应
//
// 响应格式：
//
//	成功: 200 OK + JSON { records: [...], total: N }
//	失败: 400 Bad Request (参数错误) 或 500 Internal Server Error (服务器错误)
func (h Routes) StakingRecordsHandler(w http.ResponseWriter, r *http.Request) {
	// 从 URL 查询参数中提取请求参数
	address := r.URL.Query().Get("address")
	pageQuery := r.URL.Query().Get("page")
	pageSizeQuery := r.URL.Query().Get("pageSize")
	order := r.URL.Query().Get("order")

	// 验证并处理查询参数（转换类型、设置默认值等）
	params, err := h.svc.QuerySRParams(address, pageQuery, pageSizeQuery, order)
	if err != nil {
		http.Error(w, "invalid query params", http.StatusBadRequest)
		log.Error("error reading request params", "err", err.Error())
		return
	}

	// 构造缓存键，包含所有查询条件以确保唯一性
	cacheKey := fmt.Sprintf("StakingRecords{address:%s,page:%s,pageSize:%s,order:%s}", address, pageQuery, pageSizeQuery, order)

	// 如果启用缓存，先尝试从缓存获取数据
	if h.enableCache {
		response, _ := h.cache.GetStakingRecords(cacheKey)
		if response != nil {
			// 缓存命中，直接返回缓存数据
			err = jsonResponse(w, response, http.StatusOK)
			if err != nil {
				log.Error("Error writing response", "err", err.Error())
			}
			return
		}
	}

	// 缓存未命中，从数据库查询 Staking 记录
	stakingRecords, err := h.svc.GetStakingRecords(params)
	if err != nil {
		http.Error(w, "Internal server error reading staking records", http.StatusInternalServerError)
		log.Error("Unable to read staking records from DB", "err", err.Error())
		return
	}

	// 将查询结果写入缓存，供后续请求使用
	if h.enableCache {
		h.cache.AddStakingRecords(cacheKey, stakingRecords)
	}

	// 返回 JSON 格式的响应
	err = jsonResponse(w, stakingRecords, http.StatusOK)
	if err != nil {
		log.Error("Error writing response", "err", err.Error())
	}
}

// BridgeRecordsHandler 处理跨链桥接记录查询请求
//
// HTTP 端点：GET api/v1/bridge-records
// 查询参数：
//   - address: 用户钱包地址（必需）
//   - page: 页码，默认 1
//   - pageSize: 每页记录数，默认 10
//   - order: 排序方式（asc/desc）
//
// 处理流程：
//  1. 解析并验证查询参数
//  2. 构造缓存键（基于查询参数）
//  3. 如果启用缓存，先尝试从缓存获取数据
//  4. 缓存未命中时，调用服务层从数据库查询
//  5. 将查询结果写入缓存（如果启用）
//  6. 返回 JSON 格式的响应
//
// 响应格式：
//
//	成功: 200 OK + JSON { records: [...], total: N }
//	失败: 400 Bad Request (参数错误) 或 500 Internal Server Error (服务器错误)
func (h Routes) BridgeRecordsHandler(w http.ResponseWriter, r *http.Request) {
	// 从 URL 查询参数中提取请求参数
	address := r.URL.Query().Get("address")
	pageQuery := r.URL.Query().Get("page")
	pageSizeQuery := r.URL.Query().Get("pageSize")
	order := r.URL.Query().Get("order")

	// 验证并处理查询参数（转换类型、设置默认值等）
	params, err := h.svc.QueryBRParams(address, pageQuery, pageSizeQuery, order)
	if err != nil {
		http.Error(w, "invalid query params", http.StatusBadRequest)
		log.Error("error reading request params", "err", err.Error())
		return
	}

	// 构造缓存键，包含所有查询条件以确保唯一性
	cacheKey := fmt.Sprintf("BridgeRecords{address:%s,page:%s,pageSize:%s,order:%s}", address, pageQuery, pageSizeQuery, order)

	// 如果启用缓存，先尝试从缓存获取数据
	if h.enableCache {
		response, _ := h.cache.GetBridgeRecords(cacheKey)
		if response != nil {
			// 缓存命中，直接返回缓存数据
			err = jsonResponse(w, response, http.StatusOK)
			if err != nil {
				log.Error("Error writing response", "err", err.Error())
			}
			return
		}
	}

	// 缓存未命中，从数据库查询 Bridge 记录
	bridgeRecords, err := h.svc.GetBridgeRecords(params)
	if err != nil {
		http.Error(w, "Internal server error reading bridge gas fee", http.StatusInternalServerError)
		log.Error("Unable to read bridge gas fee from DB", "err", err.Error())
		return
	}

	// 将查询结果写入缓存，供后续请求使用
	if h.enableCache {
		h.cache.AddBridgeRecords(cacheKey, bridgeRecords)
	}

	// 返回 JSON 格式的响应
	err = jsonResponse(w, bridgeRecords, http.StatusOK)
	if err != nil {
		log.Error("Error writing response", "err", err.Error())
	}
}

// BridgeGasFeeHandler 处理跨链桥接 Gas 费用查询请求
//
// HTTP 端点：GET /api/v1/bridge-price-fee
// 查询参数：
//   - chain_id: 目标链 ID（必需，必须大于 0）
//   - symbol: 代币符号，如 "ETH", "USDT" 等（可选）
//
// 处理流程：
//  1. 解析查询参数（chain_id 和 symbol）
//  2. 验证 chain_id 的有效性（必须是正整数）
//  3. 调用服务层查询指定链的 Gas 费用信息
//  4. 返回 JSON 格式的响应
//
// 注意：此接口不使用缓存，因为 Gas 费用是动态变化的
//
// 响应格式：
//
//	成功: 200 OK + JSON { chain_id: N, symbol: "XXX", gas_fee: "..." }
//	失败: 400 Bad Request (参数错误) 或 500 Internal Server Error (服务器错误)
func (h Routes) BridgeGasFeeHandler(w http.ResponseWriter, r *http.Request) {
	// 从 URL 查询参数中提取链 ID 和代币符号
	chainId := r.URL.Query().Get("chain_id")
	symbol := r.URL.Query().Get("symbol")

	// 将 chain_id 从字符串转换为整数
	chainIdInt, err := strconv.Atoi(chainId)
	if err != nil {
		return
	}

	// 验证 chain_id 必须是正整数
	if chainIdInt <= 0 {
		http.Error(w, "invalid query params", http.StatusBadRequest)
		log.Error("error reading request params", "err", err.Error())
		return
	}

	// 构造查询参数
	params := &models.QueryGasFeeParams{
		ChainId: uint64(chainIdInt),
		Symbol:  symbol,
	}

	// 调用服务层查询 Gas 费用（不使用缓存）
	bridgeGasFee, err := h.svc.GetGasFeeByChainId(params)
	if err != nil {
		http.Error(w, "Internal server error reading bridge gas fee", http.StatusInternalServerError)
		log.Error("Unable to read bridge records from DB", "err", err.Error())
		return
	}

	// 返回 JSON 格式的响应
	err = jsonResponse(w, bridgeGasFee, http.StatusOK)
	if err != nil {
		log.Error("Error writing response", "err", err.Error())
	}
}

// BridgeValidHandler 处理桥接地址有效性验证请求
//
// HTTP 端点：GET /api/v1/bridge-valid
// 查询参数：
//   - address: 待验证的钱包地址（必需）
//
// 处理流程：
//  1. 从查询参数中提取地址
//  2. 调用服务层验证地址是否可用于桥接
//  3. 返回验证结果（true/false）
//
// 用途：
//   - 验证地址格式是否正确
//   - 检查地址是否在白名单/黑名单中
//   - 确认地址是否满足桥接条件
//
// 响应格式：
//
//	成功: 200 OK + JSON { valid: true/false }
func (h Routes) BridgeValidHandler(w http.ResponseWriter, r *http.Request) {
	// 从查询参数中提取待验证的地址
	address := r.URL.Query().Get("address")

	// 调用服务层验证地址有效性
	bridgeValid := h.svc.BridgeValid(address)

	// 返回验证结果
	err := jsonResponse(w, bridgeValid, http.StatusOK)
	if err != nil {
		log.Error("Error writing response", "err", err.Error())
	}
}

// StakingValidHandler 处理 Staking 地址有效性验证请求
//
// HTTP 端点：GET /api/v1/staking-valid
// 查询参数：
//   - address: 待验证的钱包地址（必需）
//
// 处理流程：
//  1. 从查询参数中提取地址
//  2. 调用服务层验证地址是否可用于质押
//  3. 返回验证结果（true/false）
//
// 用途：
//   - 验证地址格式是否正确
//   - 检查地址是否在白名单/黑名单中
//   - 确认地址是否满足质押条件
//
// 响应格式：
//
//	成功: 200 OK + JSON { valid: true/false }
func (h Routes) StakingValidHandler(w http.ResponseWriter, r *http.Request) {
	// 从查询参数中提取待验证的地址
	address := r.URL.Query().Get("address")

	// 调用服务层验证地址有效性
	stakingValid := h.svc.StakingValid(address)

	// 返回验证结果
	err := jsonResponse(w, stakingValid, http.StatusOK)
	if err != nil {
		log.Error("Error writing response", "err", err.Error())
	}
}
