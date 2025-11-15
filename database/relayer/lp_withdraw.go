package relayer

import (
	"math/big"

	"gorm.io/gorm"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/common"
)

type LpWithdraw struct {
	GUID         uuid.UUID      `json:"guid" gorm:"primaryKey;DEFAULT replace(uuid_generate_v4()::text,'-','');serializer:uuid"`
	BlockNumber  *big.Int       `json:"dest_block_number" gorm:"serializer:u256"`
	TxHash       common.Hash    `json:"tx_hash" gorm:"serializer:bytes"`
	StartPoolId  *big.Int       `json:"start_pool_id" gorm:"serializer:u256"`
	EndPoolId    *big.Int       `json:"end_pool_id" gorm:"serializer:u256"`
	ChainId      string         `json:"chain_id"`
	TokenAddress common.Address `json:"token_address" gorm:"serializer:bytes"`
	Amount       *big.Int       `json:"amount" gorm:"serializer:u256"`
	RewardAmount *big.Int       `json:"reward_amount" gorm:"serializer:u256"`
	Timestamp    uint64         `json:"timestamp"`
}

func (LpWithdraw) TableName() string {
	return "lp_withdraw"
}

type lpWithdrawDB struct {
	gorm *gorm.DB
}

type LpWithdrawDB interface {
	LpWithdrawView
	StoreLpWithdraws(lpWithdrawList []LpWithdraw) error
	StoreLpWithdraw(lpWithdrawList LpWithdraw) error
}

type LpWithdrawView interface {
}

func NewLpWithdrawDB(db *gorm.DB) LpWithdrawDB {
	return &lpWithdrawDB{gorm: db}
}

func (db *lpWithdrawDB) StoreLpWithdraws(lpWithdrawList []LpWithdraw) error {
	lpWithdraw := new(LpWithdraw)
	result := db.gorm.Table(lpWithdraw.TableName()).Omit("guid").Create(&lpWithdrawList)
	return result.Error
}

func (db *lpWithdrawDB) StoreLpWithdraw(withdraw LpWithdraw) error {
	lpWithdraw := new(LpWithdraw)
	var exitsLpWithdraw LpWithdraw
	err := db.gorm.Table(lpWithdraw.TableName()).Where("tx_hash = ?", withdraw.TxHash.String()).Take(&exitsLpWithdraw).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			result := db.gorm.Table(lpWithdraw.TableName()).Omit("guid").Create(&withdraw)
			return result.Error
		}
	}
	return err
}
