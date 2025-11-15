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

type RPC struct {
	RpcUrl           string   `yaml:"rpc_url"`
	ChainId          uint64   `yaml:"chain_id"`
	StartBlock       uint64   `yaml:"start_block"`
	HeaderBufferSize uint64   `yaml:"header_buffer_size"`
	EventUnpackBlock uint64   `yaml:"event_unpack_block"`
	Contracts        Contract `yaml:"contracts"`
	Tokens           Token    `yaml:"tokens"`
}

type Config struct {
	Server                    Server   `yaml:"server"`
	WebSocketServer           Server   `yaml:"websocket_server"`
	GasOracleEndpoint         string   `yaml:"gas_oracle_endpoint"`
	RPCs                      []*RPC   `yaml:"rpcs"`
	Metrics                   Server   `yaml:"metrics"`
	MasterDb                  Database `yaml:"master_db"`
	SlaveDb                   Database `yaml:"slave_db"`
	PrivateKey                string   `yaml:"private_key"`
	NumConfirmations          uint64   `yaml:"num_confirmations"`
	SafeAbortNonceTooLowCount uint64   `yaml:"safe_abort_nonce_too_low_count"`
	CallerAddress             string   `yaml:"caller_address"`
	SlaveDbEnable             bool     `yaml:"slave_db_enable"`
	EnableApiCache            bool     `yaml:"enable_api_cache"`
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
