package models

import (
	"github.com/Sandwichzzy/relayer-node/database/relayer"
)

// Models 层 (service/models/)
// 职责：数据模型定义
// 定义了所有的请求参数和响应结构：
// - QuerySRParams/QueryBRParams: 查询参数
// - StakingResponse/BridgeResponse: 分页响应
// - ValidResult: 验证结果
// - GasFeeResponse: Gas 费用响应
type QuerySRParams struct {
	Address  string
	Page     int
	PageSize int
	Order    string
}

type QueryBRParams struct {
	Address  string
	Page     int
	PageSize int
	Order    string
}

type QueryPageParams struct {
	ChainId  string
	Page     int
	PageSize int
	Order    string
}

type QueryIdParams struct {
	ChainId string
	Id      uint64
}

type QueryIndexParams struct {
	ChainId string
	Index   uint64
}

type StakingResponse struct {
	Current int                     `json:"Current"`
	Size    int                     `json:"Size"`
	Total   int64                   `json:"Total"`
	Records []relayer.StakingRecord `json:"Records"`
}

type BridgeResponse struct {
	Current int                     `json:"Current"`
	Size    int                     `json:"Size"`
	Total   int64                   `json:"Total"`
	Records []relayer.BridgeRecords `json:"Records"`
}

type ValidResult struct {
	Result Result `json:"result"`
}

type Result struct {
	IsValid bool `json:"isValid"`
}

type QueryGasFeeParams struct {
	ChainId uint64
	Symbol  string
}

type GasFeeResponse struct {
	ChainId     uint64 `json:"chain_id"`
	PredictFee  string `json:"predict_fee"`
	Symbol      string `json:"symbol"`
	MarketPrice string `json:"market_price"`
}
