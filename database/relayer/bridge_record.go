package relayer

import (
	"math/big"
	"strings"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
)

type BridgeRecords struct {
	GUID               uuid.UUID      `json:"guid" gorm:"primaryKey;DEFAULT replace(uuid_generate_v4()::text,'-','');serializer:uuid"`
	SourceChainId      *big.Int       `json:"source_chain_id" gorm:"serializer:u256"`
	DestChainId        *big.Int       `json:"dest_chain_id" gorm:"serializer:u256"`
	SourceTxHash       common.Hash    `json:"source_tx_hash" gorm:"serializer:bytes"`
	DestTxHash         common.Hash    `json:"dest_tx_hash" gorm:"serializer:bytes"`
	SourceBlockNumber  *big.Int       `json:"source_block_number" gorm:"serializer:u256"`
	DestBlockNumber    *big.Int       `json:"dest_block_number" gorm:"serializer:u256"`
	SourceTokenAddress common.Address `json:"source_token_address" gorm:"serializer:bytes"`
	DestTokenAddress   common.Address `json:"dest_token_address" gorm:"serializer:bytes"`
	TokenDecimal       string         `json:"token_decimal"`
	TokenName          string         `json:"token_name"`
	MsgHash            common.Hash    `json:"msg_hash" gorm:"serializer:bytes"`
	FromAddress        common.Address `json:"from_address" gorm:"serializer:bytes"`
	ToAddress          common.Address `json:"to_address" gorm:"serializer:bytes"`
	Status             int            `json:"status"`
	Amount             *big.Int       `json:"amount" gorm:"serializer:u256"`
	Nonce              *big.Int       `json:"nonce" gorm:"serializer:u256"`
	Fee                *big.Int       `json:"fee" gorm:"serializer:u256"`
	MsgSentTimestamp   uint64         `json:"msg_sent_timestamp"`
	ClaimTimestamp     uint64         `json:"claim_timestamp"`
}

func (BridgeRecords) TableName() string {
	return "bridge_record"
}

type bridgeRecordsDB struct {
	gorm *gorm.DB
}

type BridgeRecordDB interface {
	BridgeRecordDBView
	StoreBridgeRecord(bridgeRecord BridgeRecords) error
	StoreBridgeRecords(bridgeRecord []BridgeRecords) error
	UpdateBridgeRecordsFromRelayerMsg(bridgeRecord []BridgeRecords) error
}

type BridgeRecordDBView interface {
	BridgeValid(address string) bool
	GetBridgeRecordList(address string, page int, pageSize int, order string) (bridgeRecords []BridgeRecords, total int64)
	QueryBridgeRecordByMsgHash(msgHash string) (bridgeRecordsRet *BridgeRecords, err error)
}

func NewBridgeRecordDB(db *gorm.DB) BridgeRecordDB {
	return &bridgeRecordsDB{gorm: db}
}

func (db *bridgeRecordsDB) BridgeValid(address string) bool {
	var totalRecord int64
	table := db.gorm.Table("bridge_record").Select(" *")
	querySR := db.gorm.Table("(?) as temp ", table)
	querySR = querySR.Where(BridgeRecords{FromAddress: common.HexToAddress(address)})
	result := querySR.Count(&totalRecord)
	if result.Error != nil {
		log.Error("get bridge records by address count fail", "err", result.Error)
	}
	if totalRecord > 0 {
		return true
	}
	return false
}

func (db *bridgeRecordsDB) GetBridgeRecordList(address string, page int, pageSize int, order string) (bR []BridgeRecords, total int64) {
	var totalRecord int64
	var bridgeRecordList []BridgeRecords
	queryRoot := db.gorm.Table("bridge_record")
	if address != "0x00" {
		err := db.gorm.Table("bridge_record").Select("guid").Where("from_address = ? or to_address = ?", address, address).Count(&totalRecord).Error
		if err != nil {
			log.Error("get bridge record count fail")
		}
		queryRoot.Where("from_address = ? or to_address = ?", address, address).Offset((page - 1) * pageSize).Limit(pageSize)
	} else {
		err := db.gorm.Table("bridge_record").Select("guid").Count(&totalRecord).Error
		if err != nil {
			log.Error("get bridge record count fail ")
		}
		queryRoot.Offset((page - 1) * pageSize).Limit(pageSize)
	}
	if strings.ToLower(order) == "asc" {
		queryRoot.Order("msg_sent_timestamp asc")
	} else {
		queryRoot.Order("msg_sent_timestamp desc")
	}
	qErr := queryRoot.Find(&bridgeRecordList).Error
	if qErr != nil {
		log.Error("list bridge record fail", "err", qErr)
	}
	return bridgeRecordList, totalRecord
}

func (db *bridgeRecordsDB) QueryBridgeRecordByMsgHash(msgHash string) (bridgeRecordsRet *BridgeRecords, err error) {
	var bridgeRecords BridgeRecords
	result := db.gorm.Table("bridge_record").Where("msg_hash = ? and status = 0", msgHash).Take(&bridgeRecords)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &bridgeRecords, nil
}

func (db *bridgeRecordsDB) StoreBridgeRecord(bridge BridgeRecords) error {
	bridgeRecords := new(BridgeRecords)
	result := db.gorm.Table(bridgeRecords.TableName()).Omit("guid").Create(&bridge)
	return result.Error
}

func (db *bridgeRecordsDB) StoreBridgeRecords(brs []BridgeRecords) error {
	bridgeRecords := new(BridgeRecords)
	result := db.gorm.Table(bridgeRecords.TableName()).Omit("guid").Create(&brs)
	return result.Error
}

func (db *bridgeRecordsDB) UpdateBridgeRecordsFromRelayerMsg(bridgeRecord []BridgeRecords) error {
	for i := 0; i < len(bridgeRecord); i++ {
		var bRecord = BridgeRecords{}
		result := db.gorm.Table("bridge_record").Where(&BridgeRecords{MsgHash: bridgeRecord[i].MsgHash}).Take(&bRecord)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return nil
			}
			return result.Error
		}
		bRecord.DestTxHash = bridgeRecord[i].DestTxHash
		bRecord.DestTokenAddress = bridgeRecord[i].DestTokenAddress
		bRecord.DestBlockNumber = bridgeRecord[i].DestBlockNumber
		bRecord.Status = 1
		bRecord.ClaimTimestamp = bridgeRecord[i].ClaimTimestamp
		err := db.gorm.Table("bridge_record").Save(bRecord).Error
		if err != nil {
			log.Error("save bridge record fail", "err", err)
			return err
		}
	}
	return nil
}
