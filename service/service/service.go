package service

import (
	"context"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/log"
	"github.com/pkg/errors"

	"github.com/Sandwichzzy/relayer-node/database/relayer"
	"github.com/Sandwichzzy/relayer-node/service/gasfee"
	"github.com/Sandwichzzy/relayer-node/service/models"
)

//业务逻辑和数据处理  Service 层
//核心方法：
//- GetStakingRecords : 查询 Staking 记录
//- GetBridgeRecords : 查询跨链记录
//- GetGasFeeByChainId : 通过 gRPC 获取 Gas 费用
//- QuerySRParams/QueryBRParams : 参数验证和转换

// 定义 Service 接口，声明所有业务方法
type Service interface {
	GetStakingRecords(params *models.QuerySRParams) (*models.StakingResponse, error)
	GetBridgeRecords(params *models.QueryBRParams) (*models.BridgeResponse, error)
	GetGasFeeByChainId(params *models.QueryGasFeeParams) (*models.GasFeeResponse, error)
	StakingValid(address string) *models.ValidResult
	BridgeValid(address string) *models.ValidResult

	QuerySRParams(address, page, pageSize, order string) (*models.QuerySRParams, error)
	QueryBRParams(address string, page string, pageSize string, order string) (*models.QueryBRParams, error)
}

// HandlerSvc 实现，注入依赖（Validator、数据库视图、gRPC 客户端）
type HandlerSvc struct {
	v                 *Validator
	stakingRecordView relayer.StakingRecordView
	bridgeRecordView  relayer.BridgeRecordDBView
	gasOracleClient   gasfee.TokenGasPriceServicesClient
}

func New(v *Validator, srv relayer.StakingRecordView, brv relayer.BridgeRecordDBView, gasFeeClient gasfee.TokenGasPriceServicesClient) Service {
	return &HandlerSvc{
		v:                 v,
		stakingRecordView: srv,
		bridgeRecordView:  brv,
		gasOracleClient:   gasFeeClient,
	}
}

func (h HandlerSvc) StakingValid(address string) *models.ValidResult {
	valid := h.stakingRecordView.StakingValid(address)
	return &models.ValidResult{
		Result: models.Result{
			IsValid: valid,
		},
	}
}

func (h HandlerSvc) BridgeValid(address string) *models.ValidResult {
	valid := h.bridgeRecordView.BridgeValid(address)
	return &models.ValidResult{
		Result: models.Result{IsValid: valid},
	}
}

func (h HandlerSvc) GetStakingRecords(params *models.QuerySRParams) (*models.StakingResponse, error) {
	addressToLower := strings.ToLower(params.Address)
	l2L1List, total := h.stakingRecordView.GetStakingRecords(addressToLower, params.Page, params.PageSize, params.Order)
	return &models.StakingResponse{
		Current: params.Page,
		Size:    params.PageSize,
		Total:   total,
		Records: l2L1List,
	}, nil
}

func (h HandlerSvc) GetBridgeRecords(params *models.QueryBRParams) (*models.BridgeResponse, error) {
	addressToLower := strings.ToLower(params.Address)
	l2L1List, total := h.bridgeRecordView.GetBridgeRecordList(addressToLower, params.Page, params.PageSize, params.Order)
	return &models.BridgeResponse{
		Current: params.Page,
		Size:    params.PageSize,
		Total:   total,
		Records: l2L1List,
	}, nil
}

func (h HandlerSvc) GetGasFeeByChainId(params *models.QueryGasFeeParams) (*models.GasFeeResponse, error) {
	inReq := &gasfee.TokenGasPriceRequest{
		ChainId: params.ChainId,
		Symbol:  params.Symbol,
	}
	gasFee, err := h.gasOracleClient.GetTokenPriceAndGasByChainId(context.Background(), inReq)
	if err != nil {
		log.Error("get gas fee by chain id fail", "err", err)
		return nil, err
	}

	return &models.GasFeeResponse{
		ChainId:     params.ChainId,
		PredictFee:  gasFee.PredictFee,
		Symbol:      gasFee.Symbol,
		MarketPrice: gasFee.MarketPrice,
	}, nil
}

func (h HandlerSvc) QuerySRParams(address, page, pageSize, order string) (*models.QuerySRParams, error) {
	var paraAddress string
	if address == "0x00" {
		paraAddress = "0x00"
	} else {
		addr, err := h.v.ParseValidateAddress(address)
		if err != nil {
			log.Error("invalid address param", "address", address, "err", err)
			return nil, err
		}
		paraAddress = addr.String()
	}
	pageVal, err := strconv.Atoi(page)
	if err != nil {
		return nil, errors.New("page must be an integer value")
	}
	err = h.v.ValidatePage(pageVal)
	if err != nil {
		log.Error("invalid page param", "page", page, "err", err)
		return nil, err
	}

	pageSizeVal, err := strconv.Atoi(pageSize)
	if err != nil {
		return nil, errors.New("pageSize must be an integer value")
	}
	err = h.v.ValidatePageSize(pageSizeVal)
	if err != nil {
		log.Error("invalid query param", "pageSize", pageSize, "err", err)
		return nil, err
	}

	err = h.v.ValidateOrder(order)
	if err != nil {
		log.Error("invalid query param", "order", order, "err", err)
		return nil, err
	}

	return &models.QuerySRParams{
		Address:  paraAddress,
		Page:     pageVal,
		PageSize: pageSizeVal,
		Order:    order,
	}, nil
}

func (h HandlerSvc) QueryBRParams(address string, page string, pageSize string, order string) (*models.QueryBRParams, error) {
	var paraAddress string
	if address == "0x00" {
		paraAddress = "0x00"
	} else {
		addr, err := h.v.ParseValidateAddress(address)
		if err != nil {
			log.Error("invalid address param", "address", address, "err", err)
			return nil, err
		}
		paraAddress = addr.String()
	}
	pageVal, err := strconv.Atoi(page)
	if err != nil {
		return nil, errors.New("page must be an integer value")
	}
	err = h.v.ValidatePage(pageVal)
	if err != nil {
		log.Error("invalid page param", "page", page, "err", err)
		return nil, err
	}

	pageSizeVal, err := strconv.Atoi(pageSize)
	if err != nil {
		return nil, errors.New("pageSize must be an integer value")
	}
	err = h.v.ValidatePageSize(pageSizeVal)
	if err != nil {
		log.Error("invalid query param", "pageSize", pageSize, "err", err)
		return nil, err
	}

	err = h.v.ValidateOrder(order)
	if err != nil {
		log.Error("invalid query param", "pageSize", pageSize, "err", err)
		return nil, err
	}
	return &models.QueryBRParams{
		Address:  paraAddress,
		Page:     pageVal,
		PageSize: pageSizeVal,
		Order:    order,
	}, nil
}
