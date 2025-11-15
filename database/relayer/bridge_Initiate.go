package relayer

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type BridgeInitiate struct {
	GUID               uuid.UUID      `json:"guid" gorm:"primaryKey;DEFAULT replace(uuid_generate_v4()::text,'-','');serializer:uuid"`
	TxHash             common.Hash    `json:"tx_hash" gorm:"serializer:bytes"`
	BlockNumber        *big.Int       `json:"dest_block_number" gorm:"serializer:u256"`
	SourceChainId      *big.Int       `json:"source_chain_id" gorm:"serializer:u256"`
	DestChainId        *big.Int       `json:"dest_chain_id" gorm:"serializer:u256"`
	FromAddress        common.Address `json:"from_address" gorm:"serializer:bytes"`
	ToAddress          common.Address `json:"to_address" gorm:"serializer:bytes"`
	SourceTokenAddress common.Address `json:"source_token_address" gorm:"serializer:bytes"`
	DestTokenAddress   common.Address `json:"dest_token_address" gorm:"serializer:bytes"`
	Amount             *big.Int       `json:"amount" gorm:"serializer:u256"`
	Status             int            `json:"status"`
	Timestamp          uint64         `json:"timestamp"`
}

func (BridgeInitiate) TableName() string {
	return "bridge_initiate"
}

type bridgeInitiateDB struct {
	gorm *gorm.DB
}

type BridgeInitiateDB interface {
	BridgeInitiateView
	StoreBridgeInitiates(msgHashList []BridgeInitiate) error
	MarkedBridgeInitiateHandled(bridgeInitiateList []BridgeInitiate) error
	MarkedAllBridgeInitiateHandled() error
}

type BridgeInitiateView interface {
	QueryBridgeInitiateByHash(txHash string) (*BridgeInitiate, error)
	QueryUnHandleBridgeInitiate() ([]BridgeInitiate, error)
}

func NewBridgeInitiateDB(db *gorm.DB) BridgeInitiateDB {
	return &bridgeInitiateDB{gorm: db}
}

func (db *bridgeInitiateDB) StoreBridgeInitiates(bridgeInitiateList []BridgeInitiate) error {
	bridgeInitiate := new(BridgeInitiate)
	result := db.gorm.Table(bridgeInitiate.TableName()).Omit("guid").Create(&bridgeInitiateList)
	return result.Error
}

func (db *bridgeInitiateDB) QueryUnHandleBridgeInitiate() ([]BridgeInitiate, error) {
	var bridgeInitiateList []BridgeInitiate
	err := db.gorm.Table("bridge_initiate").Where("status = ?", 0).Find(&bridgeInitiateList).Error
	if err != nil {
		log.Error("get bridge init list fail", "err", err)
		return nil, err
	}
	return bridgeInitiateList, nil
}

func (db *bridgeInitiateDB) MarkedAllBridgeInitiateHandled() error {
	var bridgeInitiateList []BridgeInitiate
	err := db.gorm.Table("bridge_initiate").Where("status = ?", 0).Find(&bridgeInitiateList).Error
	if err != nil {
		log.Error("get bridge init list fail", "err", err)
		return err
	}
	for i := 0; i < len(bridgeInitiateList); i++ {
		var bridgeInitiate = BridgeInitiate{}
		result := db.gorm.Table("bridge_initiate").Where(&BridgeInitiate{GUID: bridgeInitiateList[i].GUID}).Take(&bridgeInitiate)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return nil
			}
			log.Error("Marked Bridge Initiate Fail", "err", result.Error)
			return result.Error
		}

		log.Info("MarkedBridgeInitiateHandled bridgeInitiate status", "Status", 1)

		bridgeInitiate.Status = 1
		err := db.gorm.Table("bridge_initiate").Save(bridgeInitiate).Error
		if err != nil {
			log.Error("Update Bridge Initiate Fail", "err", result.Error)
			return err
		}
	}

	return nil
}

func (db *bridgeInitiateDB) MarkedBridgeInitiateHandled(bridgeInitiateList []BridgeInitiate) error {
	for i := 0; i < len(bridgeInitiateList); i++ {
		var bridgeInitiate = BridgeInitiate{}
		result := db.gorm.Table("bridge_initiate").Where(&BridgeInitiate{GUID: bridgeInitiateList[i].GUID}).Take(&bridgeInitiate)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return nil
			}
			log.Error("Marked Bridge Initiate Fail", "err", result.Error)
			return result.Error
		}

		log.Info("MarkedBridgeInitiateHandled bridgeInitiate status", "Status", bridgeInitiateList[i].Status)

		bridgeInitiate.Status = bridgeInitiateList[i].Status
		err := db.gorm.Table("bridge_initiate").Save(bridgeInitiate).Error
		if err != nil {
			log.Error("Update Bridge Initiate Fail", "err", result.Error)
			return err
		}
	}
	return nil
}

func (db *bridgeInitiateDB) QueryBridgeInitiateByHash(txHash string) (*BridgeInitiate, error) {
	var bridgeInitiate BridgeInitiate
	err := db.gorm.Table("bridge_initiate").Where("tx_hash = ?", txHash).Find(&bridgeInitiate).Error
	if err != nil {
		log.Error("get bridge initiate by hash fail", "err", err)
		return nil, err
	}
	return &bridgeInitiate, nil
}
