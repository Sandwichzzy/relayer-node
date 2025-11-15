package database

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/ethereum/go-ethereum/log"

	"github.com/Sandwichzzy/relayer-node/config"
	"github.com/Sandwichzzy/relayer-node/database/common"
	"github.com/Sandwichzzy/relayer-node/database/event"
	"github.com/Sandwichzzy/relayer-node/database/relayer"
	"github.com/Sandwichzzy/relayer-node/database/utils"
	"github.com/Sandwichzzy/relayer-node/synchronizer/retry"
)

type DB struct {
	gorm               *gorm.DB
	CreateTable        common.CreateTableDB
	Blocks             common.BlocksDB
	ContractEvents     event.ContractEventsDB
	BridgeInitiate     relayer.BridgeInitiateDB
	BridgeFinalize     relayer.BridgeFinalizeDB
	BridgeMsgSent      relayer.BridgeMsgSentDB
	BridgeMsgHash      relayer.BridgeMsgHashDB
	BridgeRecord       relayer.BridgeRecordDB
	ClaimReward        relayer.ClaimRewardDB
	EventBlockListener relayer.EventBlockListenerDB
	LpWithdraw         relayer.LpWithdrawDB
	StakingRecord      relayer.StakingRecordDB
	TokenConfig        relayer.TokenConfigDB
}

func NewDB(ctx context.Context, dbConfig config.Database) (*DB, error) {
	log := log.New()

	dsn := fmt.Sprintf("host=%s dbname=%s sslmode=disable", dbConfig.DbHost, dbConfig.DbName)
	if dbConfig.DbPort != 0 {
		dsn += fmt.Sprintf(" port=%d", dbConfig.DbPort)
	}
	if dbConfig.DbUser != "" {
		dsn += fmt.Sprintf(" user=%s", dbConfig.DbUser)
	}
	if dbConfig.DbPassword != "" {
		dsn += fmt.Sprintf(" password=%s", dbConfig.DbPassword)
	}

	gormConfig := gorm.Config{
		Logger:                 utils.NewLogger(log),
		SkipDefaultTransaction: true,
		CreateBatchSize:        500,
	}

	retryStrategy := &retry.ExponentialStrategy{Min: 1000, Max: 20_000, MaxJitter: 250}
	gorm, err := retry.Do[*gorm.DB](context.Background(), 10, retryStrategy, func() (*gorm.DB, error) {
		gorm, err := gorm.Open(postgres.Open(dsn), &gormConfig)
		if err != nil {
			return nil, fmt.Errorf("failed to connect to database: %w", err)
		}
		return gorm, nil
	})
	if err != nil {
		return nil, err
	}
	db := &DB{
		gorm:               gorm,
		CreateTable:        common.NewCreateTableDB(gorm),
		Blocks:             common.NewBlocksDB(gorm),
		ContractEvents:     event.NewContractEventsDB(gorm),
		BridgeInitiate:     relayer.NewBridgeInitiateDB(gorm),
		BridgeFinalize:     relayer.NewBridgeFinalizeDB(gorm),
		BridgeMsgSent:      relayer.NewBridgeMsgSentDB(gorm),
		BridgeMsgHash:      relayer.NewBridgeMsgHashDB(gorm),
		BridgeRecord:       relayer.NewBridgeRecordDB(gorm),
		ClaimReward:        relayer.NewClaimRewardDB(gorm),
		EventBlockListener: relayer.NewBlockListenerDB(gorm),
		LpWithdraw:         relayer.NewLpWithdrawDB(gorm),
		StakingRecord:      relayer.NewStakingRecordDB(gorm),
		TokenConfig:        relayer.NewTokenConfigDB(gorm),
	}
	return db, nil
}

func (db *DB) Transaction(fn func(db *DB) error) error {
	return db.gorm.Transaction(func(gorm *gorm.DB) error {
		txDB := &DB{
			gorm:               gorm,
			CreateTable:        common.NewCreateTableDB(gorm),
			Blocks:             common.NewBlocksDB(gorm),
			ContractEvents:     event.NewContractEventsDB(gorm),
			BridgeInitiate:     relayer.NewBridgeInitiateDB(gorm),
			BridgeFinalize:     relayer.NewBridgeFinalizeDB(gorm),
			BridgeMsgSent:      relayer.NewBridgeMsgSentDB(gorm),
			BridgeMsgHash:      relayer.NewBridgeMsgHashDB(gorm),
			BridgeRecord:       relayer.NewBridgeRecordDB(gorm),
			ClaimReward:        relayer.NewClaimRewardDB(gorm),
			EventBlockListener: relayer.NewBlockListenerDB(gorm),
			LpWithdraw:         relayer.NewLpWithdrawDB(gorm),
			StakingRecord:      relayer.NewStakingRecordDB(gorm),
			TokenConfig:        relayer.NewTokenConfigDB(gorm),
		}
		return fn(txDB)
	})
}

func (db *DB) Close() error {
	sql, err := db.gorm.DB()
	if err != nil {
		return err
	}
	return sql.Close()
}

func (db *DB) ExecuteSQLMigration(migrationsFolder string) error {
	err := filepath.Walk(migrationsFolder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("Failed to process migration file: %s", path))
		}
		if info.IsDir() {
			return nil
		}
		fileContent, readErr := os.ReadFile(path)
		if readErr != nil {
			return errors.Wrap(readErr, fmt.Sprintf("Error reading SQL file: %s", path))
		}
		execErr := db.gorm.Exec(string(fileContent)).Error
		if execErr != nil {
			return errors.Wrap(execErr, fmt.Sprintf("Error executing SQL script: %s", path))
		}
		return nil
	})
	return err
}
