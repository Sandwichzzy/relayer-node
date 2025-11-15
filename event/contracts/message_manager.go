package contracts

import (
	"context"
	"math/big"

	"github.com/Sandwichzzy/relayer-node/bindings"
	"github.com/Sandwichzzy/relayer-node/database"
	"github.com/Sandwichzzy/relayer-node/database/event"
	"github.com/Sandwichzzy/relayer-node/database/relayer"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/google/uuid"

	common2 "github.com/Sandwichzzy/relayer-node/common"
)

//监听和处理 MessageManager 智能合约的事件

type MessageManager struct {
	MsgManagerAbi     *abi.ABI
	MsgManagerFilter  *bindings.MessageManagerFilterer // 事件过滤器
	MsgManagerCtx     context.Context
	MsgManagerAddress string
	ChainId           string
	Epoch             uint64
}

func NewMessageManager(MsgManagerAddress string, ChainId string) (*MessageManager, error) {
	msgManagerAbi, err := bindings.MessageManagerMetaData.GetAbi()
	if err != nil {
		log.Error("get delegate manager abi fail", "err", err)
		return nil, err
	}

	msgManagerUnpack, err := bindings.NewMessageManagerFilterer(common.Address{}, nil)
	if err != nil {
		log.Error("new delegation manager fail", "err", err)
		return nil, err
	}

	return &MessageManager{
		MsgManagerAbi:     msgManagerAbi,
		MsgManagerFilter:  msgManagerUnpack,
		MsgManagerCtx:     context.Background(),
		MsgManagerAddress: MsgManagerAddress,
		ChainId:           ChainId,
	}, nil
}

// 处理指定区块范围内的 MessageManager 合约事件
func (mm *MessageManager) ProcessMessageManager(db *database.DB, fromHeight *big.Int, toHeight *big.Int) ([]relayer.BridgeMsgSent, []relayer.BridgeMsgHash, error) {
	contractEventFilter := event.ContractEvent{ContractAddress: common.HexToAddress(mm.MsgManagerAddress)}
	contractEventList, err := db.ContractEvents.ContractEventsWithFilter(mm.ChainId, contractEventFilter, fromHeight, toHeight)

	log.Info("get MessageManager event list", "contractEventList", contractEventList)
	if err != nil {
		log.Error("get contracts event list fail", "err", err)
		return nil, nil, err
	}

	var txExistList []string                      // 已处理的交易哈希列表（去重）
	var bridgeMsgSentList []relayer.BridgeMsgSent // MessageSent 事件列表
	var bridgeMsgHashList []relayer.BridgeMsgHash // MessageClaimed 事件列表

	for _, eventItem := range contractEventList {
		rlpLog := eventItem.RLPLog

		if common2.Contains(txExistList, rlpLog.TxHash.String()) {
			continue // 跳过已处理的交易
		}

		header, err := db.Blocks.ChainBlockHeader(mm.ChainId, eventItem.BlockHash)
		if err != nil {
			log.Error("ProcessPoolManagerEvent db Blocks BlockHeader by BlockHash fail", "err", err)
			return nil, nil, err
		}

		// 触发场景: 用户在源链发起跨链转账，MessageManager 合约发出消息
		if eventItem.EventSignature.String() == mm.MsgManagerAbi.Events["MessageSent"].ID.String() {
			parseMessageSentEvent, err := mm.MsgManagerFilter.ParseMessageSent(*rlpLog)
			if err != nil {
				log.Error("parse initialized eth event fail", "err", err)
				return nil, nil, err
			}
			bridgeMsgSentItem := relayer.BridgeMsgSent{
				GUID:               uuid.New(),
				TxHash:             parseMessageSentEvent.Raw.TxHash,
				BlockNumber:        big.NewInt(int64(parseMessageSentEvent.Raw.BlockNumber)),
				SourceTokenAddress: parseMessageSentEvent.SourceTokenAddress,
				DestTokenAddress:   parseMessageSentEvent.DestTokenAddress,
				SourceChainId:      parseMessageSentEvent.SourceChainId,
				DestChainId:        parseMessageSentEvent.DestChainId,
				FromAddress:        parseMessageSentEvent.From,
				ToAddress:          parseMessageSentEvent.To,
				Fee:                parseMessageSentEvent.Fee,
				MsgNonce:           parseMessageSentEvent.Nonce,
				Amount:             parseMessageSentEvent.Value,
				Relation:           false, // 标记为未关联 消息尚未被目标链认领
				MsgHash:            parseMessageSentEvent.MessageHash,
				Timestamp:          header.Timestamp,
			}
			bridgeMsgSentList = append(bridgeMsgSentList, bridgeMsgSentItem)
		}
		// 触发场景: Relayer 在目标链提交消息后，MessageManager 合约确认消息被认领
		if eventItem.EventSignature.String() == mm.MsgManagerAbi.Events["MessageClaimed"].ID.String() {
			parseMessageClaimedEvent, err := mm.MsgManagerFilter.ParseMessageClaimed(*rlpLog)
			if err != nil {
				log.Error("parse initialized eth event fail", "err", err)
				return nil, nil, err
			}
			bridgeMsgClaimItem := relayer.BridgeMsgHash{
				GUID:               uuid.New(),
				TxHash:             parseMessageClaimedEvent.Raw.TxHash,
				BlockNumber:        big.NewInt(int64(parseMessageClaimedEvent.Raw.BlockNumber)),
				SourceChainId:      parseMessageClaimedEvent.SourceChainId,
				DestChainId:        parseMessageClaimedEvent.DestChainId,
				SourceTokenAddress: parseMessageClaimedEvent.SourceTokenAddress,
				DestTokenAddress:   parseMessageClaimedEvent.DestTokenAddress,
				Relation:           false, // 待关联到 BridgeMsgSent
				MsgHash:            parseMessageClaimedEvent.MessageHash,
				MsgNonce:           parseMessageClaimedEvent.Nonce,
				Timestamp:          header.Timestamp,
			}
			bridgeMsgHashList = append(bridgeMsgHashList, bridgeMsgClaimItem)
		}
		txExistList = append(txExistList, rlpLog.TxHash.String())
	}
	return bridgeMsgSentList, bridgeMsgHashList, nil
}
