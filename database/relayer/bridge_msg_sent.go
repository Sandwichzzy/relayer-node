package relayer

import (
	"errors"
	"math/big"

	"gorm.io/gorm"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/google/uuid"
)

//bridge_msg_sent (消息发送表)
//
//作用: 记录 relayer 发送到目标链的跨链消息
//Relayer 构建消息 → 发送到链B → 记录 msg_hash 和 nonce
//- 关键字段: MsgHash, MsgNonce, Fee, Relation
//- Relation: false=未关联, true=已与目标链交易关联

type BridgeMsgSent struct {
	GUID               uuid.UUID      `json:"guid" gorm:"primaryKey;DEFAULT replace(uuid_generate_v4()::text,'-','');serializer:uuid"`
	TxHash             common.Hash    `json:"tx_hash" gorm:"serializer:bytes"`
	BlockNumber        *big.Int       `json:"dest_block_number" gorm:"serializer:u256"`
	SourceChainId      *big.Int       `json:"source_chain_id" gorm:"serializer:u256"`
	DestChainId        *big.Int       `json:"dest_chain_id" gorm:"serializer:u256"`
	SourceTokenAddress common.Address `json:"source_token_address" gorm:"serializer:bytes"`
	DestTokenAddress   common.Address `json:"dest_token_address" gorm:"serializer:bytes"`
	FromAddress        common.Address `json:"from_address" gorm:"serializer:bytes"`
	ToAddress          common.Address `json:"to_address" gorm:"serializer:bytes"`
	Fee                *big.Int       `json:"fee" gorm:"serializer:u256"`
	MsgNonce           *big.Int       `json:"msg_nonce" gorm:"serializer:u256"`
	Amount             *big.Int       `json:"amount" gorm:"serializer:u256"`
	Relation           bool           `json:"relation"`
	MsgHash            common.Hash    `json:"msg_hash" gorm:"serializer:bytes"`
	Timestamp          uint64         `json:"timestamp"`
}

func (BridgeMsgSent) TableName() string {
	return "bridge_msg_sent"
}

type bridgeMsgSentDB struct {
	gorm *gorm.DB
}

type BridgeMsgSentDB interface {
	BridgeMsgSentDBView
	StoreBridgeMsgSent(msgSent BridgeMsgSent) error
	StoreBridgeMsgSents(msgSentList []BridgeMsgSent) error
	MarkedBridgeMsgSentdHandled(msgSentList []BridgeMsgSent) error
}

type BridgeMsgSentDBView interface {
	QueryBridgeMsgSents(sourceChainId string) ([]BridgeMsgSent, error)
	QueryBridgeMsgSentByHash(txHash string) (*BridgeMsgSent, error)
	QueryBridgeMsgSentByMsgHash(msgHash string) (*BridgeMsgSent, error)
}

func NewBridgeMsgSentDB(db *gorm.DB) BridgeMsgSentDB {
	return &bridgeMsgSentDB{gorm: db}
}

func (db *bridgeMsgSentDB) StoreBridgeMsgSent(msgSent BridgeMsgSent) error {
	msgSentRecord := new(BridgeMsgSent)
	var exist BridgeMsgSent
	err := db.gorm.Table(msgSentRecord.TableName()).Where("tx_hash = ?", msgSent.TxHash.String()).Take(&exist).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			result := db.gorm.Table(msgSentRecord.TableName()).Omit("guid").Create(&msgSent)
			return result.Error
		}
	}
	return err
}

func (db *bridgeMsgSentDB) StoreBridgeMsgSents(msgSentList []BridgeMsgSent) error {
	msgSentRecord := new(BridgeMsgSent)
	result := db.gorm.Table(msgSentRecord.TableName()).Omit("guid").Create(&msgSentList)
	return result.Error
}

func (db *bridgeMsgSentDB) QueryBridgeMsgSents(sourceChainId string) ([]BridgeMsgSent, error) {
	var bridgeMsgSentList []BridgeMsgSent
	err := db.gorm.Table("bridge_msg_sent").Where("source_chain_id = ? and relation = ?", sourceChainId, false).Find(&bridgeMsgSentList).Error
	if err != nil {
		log.Error("get unhandled bridge message send list", "err", err)
		return nil, err
	}
	return bridgeMsgSentList, nil
}

func (db *bridgeMsgSentDB) QueryBridgeMsgSentByHash(txHash string) (*BridgeMsgSent, error) {
	var bridgeMsgSent BridgeMsgSent
	err := db.gorm.Table("bridge_msg_sent").Where("tx_hash = ?", txHash).Find(&bridgeMsgSent).Error
	if err != nil {
		log.Error("get bridge message send by hash fail", "err", err)
		return nil, err
	}
	return &bridgeMsgSent, nil
}

func (db *bridgeMsgSentDB) QueryBridgeMsgSentByMsgHash(msgHash string) (*BridgeMsgSent, error) {
	var bridgeMsgSent BridgeMsgSent
	err := db.gorm.Table("bridge_msg_sent").Where("msg_hash = ?", msgHash).Find(&bridgeMsgSent).Error
	if err != nil {
		log.Error("get bridge message send by message hash fail", "err", err)
		return nil, err
	}
	return &bridgeMsgSent, nil
}

func (db *bridgeMsgSentDB) MarkedBridgeMsgSentdHandled(msgSentList []BridgeMsgSent) error {
	for i := 0; i < len(msgSentList); i++ {
		var bridgeMsgSent = BridgeMsgSent{}
		result := db.gorm.Table("bridge_msg_sent").Where(&BridgeMsgSent{GUID: msgSentList[i].GUID}).Take(&bridgeMsgSent)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return nil
			}
			return result.Error
		}
		bridgeMsgSent.Relation = true
		err := db.gorm.Table("bridge_msg_sent").Save(bridgeMsgSent).Error
		if err != nil {
			return err
		}
	}
	return nil
}
