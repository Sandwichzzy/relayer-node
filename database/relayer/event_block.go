package relayer

import (
	"errors"
	"math/big"
	"time"

	"gorm.io/gorm"

	"github.com/ethereum/go-ethereum/log"
	"github.com/google/uuid"
)

// event_block (区块监听表)
//
// 作用: 记录每条链的最后扫描区块高度
// 防止重复扫描,支持断点续传
// - 查询: GetLastBlockNumber
// - 更新: SaveOrUpdateLastBlockNumber
type EventBlockListener struct {
	GUID        uuid.UUID `json:"guid" gorm:"primaryKey;DEFAULT replace(uuid_generate_v4()::text,'-','');serializer:uuid"`
	ChainId     string    `json:"chain_id"`
	BlockNumber *big.Int  `json:"block_number" gorm:"serializer:u256"`
	Created     uint64    `json:"created"`
	Updated     uint64    `json:"updated"`
}

func (EventBlockListener) TableName() string {
	return "event_block"
}

type eventBlockListenerDB struct {
	gorm *gorm.DB
}

type EventBlockListenerDB interface {
	EventBlockListenerDBView
	SaveOrUpdateLastBlockNumber(lastBlock EventBlockListener) error
}

type EventBlockListenerDBView interface {
	GetLastBlockNumber(chainId string) (lastBlock *EventBlockListener, err error)
}

func NewBlockListenerDB(db *gorm.DB) EventBlockListenerDB {
	return &eventBlockListenerDB{gorm: db}
}

func (db *eventBlockListenerDB) SaveOrUpdateLastBlockNumber(lastBlock EventBlockListener) error {
	bbl := new(EventBlockListener)
	var exitsLastBlock EventBlockListener

	err := db.gorm.Table(bbl.TableName()).Where("chain_id = ?", lastBlock.ChainId).Take(&exitsLastBlock).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Info("No last block number found, will insert it")
			lastBlock.Created = uint64(time.Now().Unix())
			lastBlock.Updated = uint64(time.Now().Unix())
			result := db.gorm.Table(bbl.TableName()).Omit("guid").Create(&lastBlock)
			if result.Error != nil {
				return result.Error
			}
			return nil
		}
		return err
	} else {
		lastBlock.Updated = uint64(time.Now().Unix())
		updateResult := db.gorm.Table(bbl.TableName()).Omit("guid", "created").Where("chain_id = ?", lastBlock.ChainId).Updates(&lastBlock)
		if updateResult.Error != nil {
			return updateResult.Error
		}
	}
	return nil
}

func (db *eventBlockListenerDB) GetLastBlockNumber(chainId string) (lastBlock *EventBlockListener, err error) {
	bbl := new(EventBlockListener)
	err = db.gorm.Table(bbl.TableName()).Where("chain_id = ?", chainId).Take(&lastBlock).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return lastBlock, nil
}
