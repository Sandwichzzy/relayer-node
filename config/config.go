package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Server struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Database struct {
	DbHost     string `yaml:"db_host"`
	DbPort     int    `yaml:"db_port"`
	DbName     string `yaml:"db_name"`
	DbUser     string `yaml:"db_user"`
	DbPassword string `yaml:"db_password"`
}

type Contract struct {
	PoolManagerAddress    string `yaml:"pool_manager_address"`
	MessageManagerAddress string `yaml:"message_manager_address"`
}

type Token struct {
	Eth  TokenCommon `yaml:"eth"`
	USDT TokenCommon `yaml:"usdt"`
	Cp   TokenCommon `yaml:"cp"`
}

type TokenCommon struct {
	Address string `json:"address"`
	Decimal string `json:"decimal"`
	Name    string `json:"name"`
}

// RPC 链配置
// 每个 RPC 配置对应一条区块链，支持多链监控
type RPC struct {
	RpcUrl           string   `yaml:"rpc_url"`            // 区块链 RPC 节点地址
	ChainId          uint64   `yaml:"chain_id"`           // 链 ID（1=以太坊主网, 56=BSC, 等）
	StartBlock       uint64   `yaml:"start_block"`        // Synchronizer 起始区块：从哪个区块开始同步区块头和原始事件日志
	HeaderBufferSize uint64   `yaml:"header_buffer_size"` // 每批处理的区块头数量
	EventUnpackBlock uint64   `yaml:"event_unpack_block"` // Event Processor 起始区块：从哪个区块开始解析（unpack）事件为业务数据
	Contracts        Contract `yaml:"contracts"`          // 需要监听的合约地址
	Tokens           Token    `yaml:"tokens"`             // 代币配置
}

type Config struct {
	Server                    Server   `yaml:"server"`
	WebSocketServer           Server   `yaml:"websocket_server"`
	GasOracleEndpoint         string   `yaml:"gas_oracle_endpoint"`
	RPCs                      []*RPC   `yaml:"rpcs"`
	Metrics                   Server   `yaml:"metrics"`
	MasterDb                  Database `yaml:"master_db"`
	SlaveDb                   Database `yaml:"slave_db"`
	NumConfirmations          uint64   `yaml:"num_confirmations"`
	SafeAbortNonceTooLowCount uint64   `yaml:"safe_abort_nonce_too_low_count"`
	CallerAddress             string   `yaml:"caller_address"`
	SlaveDbEnable             bool     `yaml:"slave_db_enable"`
	EnableApiCache            bool     `yaml:"enable_api_cache"`
	//PrivateKey              string   `yaml:"private_key"`
}

func New(path string) (*Config, error) {
	var config = new(Config)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
