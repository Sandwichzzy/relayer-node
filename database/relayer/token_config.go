package relayer

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/Sandwichzzy/relayer-node/config"
)

type TokenConfig struct {
	GUID         uuid.UUID      `json:"guid" gorm:"primaryKey;DEFAULT replace(uuid_generate_v4()::text,'-','');serializer:uuid"`
	ChainId      *big.Int       `json:"chain_id" gorm:"serializer:u256"`
	TokenAddress common.Address `json:"token_address" gorm:"serializer:bytes"`
	TokenDecimal string         `json:"token_decimal"`
	TokenName    string         `json:"token_name"`
	Timestamp    uint64         `json:"timestamp"`
}

func (TokenConfig) TableName() string {
	return "token_config"
}

type tokenConfigDB struct {
	gorm *gorm.DB
}

type TokenConfigDB interface {
	TokenConfigView
	StoreTokenConfigs(tokenConfigList []TokenConfig) error
	StoreTokenConfig(tokenConfigList TokenConfig) error
	DeleteTokenConfigBeforeNow() error
}

type TokenConfigView interface {
	BuildTokenConfig(rpc *config.RPC) []TokenConfig
	QueryTokenBySourceChain(chainId uint64, tokenAddress string) (tokenConf *TokenConfig, err error)
}

func NewTokenConfigDB(db *gorm.DB) TokenConfigDB {
	return &tokenConfigDB{gorm: db}
}

func (db *tokenConfigDB) BuildTokenConfig(rpc *config.RPC) []TokenConfig {
	var tokens []TokenConfig
	ethToken := TokenConfig{
		GUID:         uuid.New(),
		ChainId:      new(big.Int).SetUint64(rpc.ChainId),
		TokenAddress: common.HexToAddress(rpc.Tokens.Eth.Address),
		TokenDecimal: rpc.Tokens.Eth.Decimal,
		TokenName:    rpc.Tokens.Eth.Name,
		Timestamp:    uint64(time.Now().Unix()),
	}
	usdtToken := TokenConfig{
		GUID:         uuid.New(),
		ChainId:      new(big.Int).SetUint64(rpc.ChainId),
		TokenAddress: common.HexToAddress(rpc.Tokens.USDT.Address),
		TokenDecimal: rpc.Tokens.USDT.Decimal,
		TokenName:    rpc.Tokens.USDT.Name,
		Timestamp:    uint64(time.Now().Unix()),
	}
	cpToken := TokenConfig{
		GUID:         uuid.New(),
		ChainId:      new(big.Int).SetUint64(rpc.ChainId),
		TokenAddress: common.HexToAddress(rpc.Tokens.Cp.Address),
		TokenDecimal: rpc.Tokens.Cp.Decimal,
		TokenName:    rpc.Tokens.Cp.Name,
		Timestamp:    uint64(time.Now().Unix()),
	}
	tokens = append(tokens, ethToken, usdtToken, cpToken)
	return tokens
}

func (db *tokenConfigDB) StoreTokenConfigs(tokenConfigList []TokenConfig) error {
	tokenConfig := new(TokenConfig)
	result := db.gorm.Table(tokenConfig.TableName()).Omit("guid").Create(&tokenConfigList)
	return result.Error
}

func (db *tokenConfigDB) StoreTokenConfig(tokenConf TokenConfig) error {
	tokenConfig := new(TokenConfig)
	var exitsTokenConfig TokenConfig
	err := db.gorm.Table(tokenConfig.TableName()).Where("token_address = ? and chain_id = ?", tokenConf.TokenAddress, tokenConf.ChainId).Take(&exitsTokenConfig).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			result := db.gorm.Table("token_config").Omit("guid").Create(&tokenConfig)
			return result.Error
		}
	}
	return err
}

func (db *tokenConfigDB) QueryTokenBySourceChain(chainId uint64, tokenAddress string) (tokenConf *TokenConfig, err error) {
	var tokenConfig TokenConfig
	result := db.gorm.Table("token_config").Where("token_address = ? and chain_id = ?", strings.ToLower(tokenAddress), chainId).Take(&tokenConfig)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &tokenConfig, nil
}

func (db *tokenConfigDB) DeleteTokenConfigBeforeNow() error {
	tokenConfig := new(TokenConfig)
	now := uint64(time.Now().Unix())
	result := db.gorm.Table(tokenConfig.TableName()).Where("timestamp < ?", now).Delete(&tokenConfig)
	return result.Error
}
