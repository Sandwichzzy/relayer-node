package relayer

import (
	"errors"

	"math/big"

	"gorm.io/gorm"

	"github.com/google/uuid"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
)

type BridgeFinalize struct {
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

func (BridgeFinalize) TableName() string {
	return "bridge_finalize"
}

type bridgeFinalizeDB struct {
	gorm *gorm.DB
}

type BridgeFinalizeDB interface {
	BridgeFinalizeView
	StoreBridgeFinalize(finalize BridgeFinalize) error
	StoreBridgeFinalizes(finalizes []BridgeFinalize) error
}

type BridgeFinalizeView interface {
	QueryUnHandleBridgeFinalize() ([]BridgeFinalize, error)
	MarkedBridgeFinalizeHandled(bridgeFinalizeList []BridgeFinalize) error
}

func NewBridgeFinalizeDB(db *gorm.DB) BridgeFinalizeDB {
	return &bridgeFinalizeDB{gorm: db}
}

func (db *bridgeFinalizeDB) StoreBridgeFinalize(finalize BridgeFinalize) error {
	bridgeFinalize := new(BridgeFinalize)
	var exist BridgeFinalize
	err := db.gorm.Table(bridgeFinalize.TableName()).Where("tx_hash = ?", finalize.TxHash.String()).Take(&exist).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			result := db.gorm.Table(bridgeFinalize.TableName()).Omit("guid").Create(&finalize)
			return result.Error
		}
	}
	return err
}

func (db *bridgeFinalizeDB) StoreBridgeFinalizes(finalizes []BridgeFinalize) error {
	bridgeFinalize := new(BridgeFinalize)
	result := db.gorm.Table(bridgeFinalize.TableName()).Omit("guid").Create(&finalizes)
	return result.Error
}

func (db *bridgeFinalizeDB) QueryUnHandleBridgeFinalize() ([]BridgeFinalize, error) {
	var bridgeFinalizeList []BridgeFinalize
	err := db.gorm.Table("bridge_finalize").Where("status = ?", 0).Find(&bridgeFinalizeList).Error
	if err != nil {
		log.Error("get bridge finalized list fail", "err", err)
		return nil, err
	}
	return bridgeFinalizeList, nil
}

func (db *bridgeFinalizeDB) MarkedBridgeFinalizeHandled(bridgeFinalizeList []BridgeFinalize) error {
	for i := 0; i < len(bridgeFinalizeList); i++ {
		var bridgeFinalize = BridgeFinalize{}
		result := db.gorm.Table("bridge_finalize").Where(&BridgeFinalize{GUID: bridgeFinalizeList[i].GUID}).Take(&bridgeFinalize)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return nil
			}
			log.Error("Marked Bridge finalized Fail", "err", result.Error)
			return result.Error
		}

		log.Info("MarkedBridgeFinalizeHandled bridgeFinalize status", "Status", bridgeFinalizeList[i].Status)

		bridgeFinalize.Status = bridgeFinalizeList[i].Status
		err := db.gorm.Table("bridge_finalize").Save(bridgeFinalize).Error
		if err != nil {
			log.Error("Update Bridge finalized Fail", "err", result.Error)
			return err
		}
	}
	return nil
}
