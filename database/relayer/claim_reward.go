package relayer

import (
	"errors"
	"math/big"

	"gorm.io/gorm"

	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
)

type ClaimReward struct {
	GUID         uuid.UUID      `json:"guid" gorm:"primaryKey;DEFAULT replace(uuid_generate_v4()::text,'-','');serializer:uuid"`
	BlockNumber  *big.Int       `json:"dest_block_number" gorm:"serializer:u256"`
	TxHash       common.Hash    `json:"tx_hash" gorm:"serializer:bytes"`
	ClaimAddress common.Address `json:"claim_address" gorm:"serializer:bytes"`
	TokenAddress common.Address `json:"token_address" gorm:"serializer:bytes"`
	StartPoolId  *big.Int       `json:"start_pool_id" gorm:"serializer:u256"`
	EndPoolId    *big.Int       `json:"end_pool_id" gorm:"serializer:u256"`
	ChainId      string         `json:"chain_id"`
	RewardAmount *big.Int       `json:"reward_amount" gorm:"serializer:u256"`
	Timestamp    uint64         `json:"timestamp"`
}

func (ClaimReward) TableName() string {
	return "claim_reward"
}

type claimRewardDB struct {
	gorm *gorm.DB
}

type ClaimRewardDB interface {
	ClaimRewardView
	StoreBridgeClaim(claim ClaimReward) error
	StoreBridgeClaims(claims []ClaimReward) error
}

type ClaimRewardView interface {
}

func NewClaimRewardDB(db *gorm.DB) ClaimRewardDB {
	return &claimRewardDB{gorm: db}
}

func (db *claimRewardDB) StoreBridgeClaim(claim ClaimReward) error {
	bridgeClaimed := new(ClaimReward)
	var exist ClaimReward
	err := db.gorm.Table(bridgeClaimed.TableName()).Where("tx_hash = ?", claim.TxHash.String()).Take(&exist).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			result := db.gorm.Table(bridgeClaimed.TableName()).Omit("guid").Create(&claim)
			return result.Error
		}
	}
	return err
}

func (db *claimRewardDB) StoreBridgeClaims(claims []ClaimReward) error {
	bridgeClaimed := new(ClaimReward)
	result := db.gorm.Table(bridgeClaimed.TableName()).Omit("guid").Create(&claims)
	return result.Error
}
