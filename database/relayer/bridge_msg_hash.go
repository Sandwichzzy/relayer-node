package relayer

import (
	"errors"
	"math/big"

	"gorm.io/gorm"

	"github.com/ethereum/go-ethereum/log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
)

//bridge_msg_hash (消息哈希表)
//
//作用: 追踪跨链消息的哈希值,用于消息匹配
//用于验证消息完整性和防重放攻击
//- 查询方法: QueryBridgeMsgHashs

type BridgeMsgHash struct {
	GUID               uuid.UUID      `json:"guid" gorm:"primaryKey;DEFAULT replace(uuid_generate_v4()::text,'-','');serializer:uuid"`
	BlockNumber        *big.Int       `json:"dest_block_number" gorm:"serializer:u256"`
	TxHash             common.Hash    `json:"tx_hash" gorm:"serializer:bytes"`
	SourceChainId      *big.Int       `json:"source_chain_id" gorm:"serializer:u256"`
	DestChainId        *big.Int       `json:"dest_chain_id" gorm:"serializer:u256"`
	SourceTokenAddress common.Address `json:"source_token_address" gorm:"serializer:bytes"`
	DestTokenAddress   common.Address `json:"dest_token_address" gorm:"serializer:bytes"`
	MsgHash            common.Hash    `json:"msg_hash" gorm:"serializer:bytes"`
	MsgNonce           *big.Int       `json:"msg_nonce" gorm:"serializer:u256"`
	Relation           bool           `json:"relation"`
	Timestamp          uint64         `json:"timestamp"`
}

func (BridgeMsgHash) TableName() string {
	return "bridge_msg_hash"
}

type bridgeMsgHashDB struct {
	gorm *gorm.DB
}

type BridgeMsgHashDB interface {
	BridgeMsgHashView
	StoreBridgeMsgHash(msgHash BridgeMsgHash) error
	StoreBridgeMsgHashs(msgHashList []BridgeMsgHash) error
	MarkedBridgeRelayMessageHandled(msgHashList []BridgeMsgHash) error
}

type BridgeMsgHashView interface {
	QueryBridgeMsgHashByHash(txHash string) (*BridgeMsgHash, error)
	QueryBridgeMsgHashs(destChainId string) ([]BridgeMsgHash, error)
}

func NewBridgeMsgHashDB(db *gorm.DB) BridgeMsgHashDB {
	return &bridgeMsgHashDB{gorm: db}
}

func (db *bridgeMsgHashDB) StoreBridgeMsgHash(msgHash BridgeMsgHash) error {
	msgHashRecord := new(BridgeMsgHash)
	var exist BridgeMsgHash
	err := db.gorm.Table(msgHashRecord.TableName()).Where("tx_hash = ?", msgHash.TxHash.String()).Take(&exist).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			result := db.gorm.Table(msgHashRecord.TableName()).Omit("guid").Create(&msgHash)
			return result.Error
		}
	}
	return err
}

func (db *bridgeMsgHashDB) StoreBridgeMsgHashs(msgHashList []BridgeMsgHash) error {
	msgHashRecord := new(BridgeMsgHash)
	result := db.gorm.Table(msgHashRecord.TableName()).Omit("guid").Create(&msgHashList)
	return result.Error
}

func (db *bridgeMsgHashDB) QueryBridgeMsgHashs(destChainId string) ([]BridgeMsgHash, error) {
	var bridgeMsgHashList []BridgeMsgHash
	err := db.gorm.Table("bridge_msg_hash").Where("dest_chain_id = ? and relation = ?", destChainId, false).Find(&bridgeMsgHashList).Error
	if err != nil {
		log.Error("get bridge message hash list", "err", err)
		return nil, err
	}
	return bridgeMsgHashList, nil
}

func (db *bridgeMsgHashDB) MarkedBridgeRelayMessageHandled(msgHashList []BridgeMsgHash) error {
	for i := 0; i < len(msgHashList); i++ {
		var bridgeMsgHash = BridgeMsgHash{}
		result := db.gorm.Table("bridge_msg_hash").Where(&BridgeMsgHash{GUID: msgHashList[i].GUID}).Take(&bridgeMsgHash)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return nil
			}
			return result.Error
		}
		bridgeMsgHash.Relation = true
		err := db.gorm.Table("bridge_msg_hash").Save(bridgeMsgHash).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (db *bridgeMsgHashDB) QueryBridgeMsgHashByHash(txHash string) (*BridgeMsgHash, error) {
	var bridgeMsgHash BridgeMsgHash
	err := db.gorm.Table("bridge_msg_hash").Where("tx_hash = ?", txHash).Find(&bridgeMsgHash).Error
	if err != nil {
		log.Error("get bridge message hash by hash fail", "err", err)
		return nil, err
	}
	return &bridgeMsgHash, nil
}
