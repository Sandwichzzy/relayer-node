package relayer_node

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"net"
	"net/http"
	"os"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/Sandwichzzy/relayer-node/config"
	"github.com/Sandwichzzy/relayer-node/database"
	"github.com/Sandwichzzy/relayer-node/event"
	"github.com/Sandwichzzy/relayer-node/metrics"
	"github.com/Sandwichzzy/relayer-node/relayer"
	"github.com/Sandwichzzy/relayer-node/relayer/driver"
	"github.com/Sandwichzzy/relayer-node/service/common/httputil"
	"github.com/Sandwichzzy/relayer-node/service/websocket"
	"github.com/Sandwichzzy/relayer-node/synchronizer"
	"github.com/Sandwichzzy/relayer-node/synchronizer/node"
	"github.com/Sandwichzzy/relayer-node/worker"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus"
)

type RelayerNode struct {
	DB               *database.DB
	ethClient        map[uint64]node.EthClient
	metricsServer    *httputil.HTTPServer
	metricsRegistry  *prometheus.Registry
	relayerMetrics   *metrics.SyncerMetrics
	Synchronizer     map[uint64]*synchronizer.Synchronizer
	EventProcessor   map[uint64]*event.Processor
	relayerProcessor *relayer.RelayerProcessor
	WorkerHandle     *worker.WorkerHandle
	wsHub            *websocket.Hub
	wsServer         *httputil.HTTPServer
	shutdown         context.CancelCauseFunc
	stopped          atomic.Bool
	chainIdList      []uint64
}

type RpcServerConfig struct {
	GrpcHostname string
	GrpcPort     int
}

func NewRelayerNode(ctx context.Context, cfg *config.Config, shutdown context.CancelCauseFunc) (*RelayerNode, error) {
	log.Info("New relayer node start️")
	metricsRegistry := metrics.NewRegistry() //创建注册表

	relayerMetrics := metrics.NewSyncerMetrics(metricsRegistry, "replayer") //创建并注册指标

	out := &RelayerNode{
		metricsRegistry: metricsRegistry,
		relayerMetrics:  relayerMetrics,
		shutdown:        shutdown,
	}
	if err := out.initFromConfig(ctx, cfg); err != nil {
		return nil, errors.Join(err, out.Stop(ctx))
	}
	log.Info("New relayer node success🏅️")
	return out, nil
}

func (as *RelayerNode) Start(ctx context.Context) error {
	//webSocket 和 metrics server 在initFromConfig已經啟動
	for i := range as.chainIdList {
		log.Info("starting sync", "chainId", as.chainIdList[i])
		realChainId := as.chainIdList[i]
		if err := as.Synchronizer[realChainId].Start(); err != nil {
			return fmt.Errorf("failed to start chain sync: %w", err)
		}

		if err := as.EventProcessor[realChainId].Start(); err != nil {
			return fmt.Errorf("failed to event processor: %w", err)
		}
	}

	errRelayer := as.relayerProcessor.Start()
	if errRelayer != nil {
		log.Error("start relayer processor fail", "err", errRelayer)
		return errRelayer
	}

	errWorker := as.WorkerHandle.Start()
	if errWorker != nil {
		log.Error("start worker handle fail", "err", errWorker)
		return errWorker
	}

	return nil

}

func (as *RelayerNode) Stop(ctx context.Context) error {
	var result error
	for i := range as.chainIdList {
		if as.Synchronizer[as.chainIdList[i]] != nil {
			if err := as.Synchronizer[as.chainIdList[i]].Close(); err != nil {
				result = errors.Join(result, fmt.Errorf("failed to close L2 Synchronizer: %w", err))
			}
		}
		if as.ethClient[as.chainIdList[i]] != nil {
			as.ethClient[as.chainIdList[i]].Close()
		}
	}

	if as.DB != nil {
		if err := as.DB.Close(); err != nil {
			result = errors.Join(result, fmt.Errorf("failed to close DB: %w", err))
		}
	}

	if as.metricsServer != nil {
		if err := as.metricsServer.Close(); err != nil {
			result = errors.Join(result, fmt.Errorf("failed to close metrics server: %w", err))
		}
	}

	if as.wsHub != nil {
		as.wsHub.CloseAllClients()
	}

	if as.wsServer != nil {
		if err := as.wsServer.Stop(ctx); err != nil {
			result = errors.Join(result, fmt.Errorf("failed to stop WebSocket server: %w", err))
		}
	}

	as.stopped.Store(true)

	log.Info("replayer node stopped")

	return result
}

func (as *RelayerNode) Stopped() bool {
	return as.stopped.Load()
}

func (as *RelayerNode) initFromConfig(ctx context.Context, cfg *config.Config) error {
	if err := as.initRPCClients(ctx, cfg); err != nil {
		return fmt.Errorf("failed to start RPC clients: %w", err)
	}
	if err := as.initDB(ctx, cfg.MasterDb); err != nil {
		return fmt.Errorf("failed to init DB: %w", err)
	}
	as.wsHub = websocket.NewHub()
	go as.wsHub.Run()
	if err := as.startWebSocketServer(cfg.WebSocketServer); err != nil {
		return fmt.Errorf("failed to start WebSocket server: %w", err)
	}
	if err := as.initSynchronizer(cfg); err != nil {
		return fmt.Errorf("failed to init L1 Sync: %w", err)
	}

	if err := as.initEventProcessor(cfg); err != nil {
		return fmt.Errorf("failed to init event processor: %w", err)
	}
	if err := as.initRelayer(cfg); err != nil {
		return fmt.Errorf("failed to init event processor: %w", err)
	}
	if err := as.initWorker(cfg); err != nil {
		return fmt.Errorf("failed to init event processor: %w", err)
	}
	err := as.startMetricsServer(cfg.Metrics)
	if err != nil {
		log.Error("start metrics server fail", "err", err)
		return err
	}
	return nil
}

func (as *RelayerNode) startWebSocketServer(serverConfig config.Server) error {
	log.Debug("WebSocket server listening...", "port", serverConfig.Port)
	addr := net.JoinHostPort(serverConfig.Host, strconv.Itoa(serverConfig.Port))

	wsRouter := chi.NewRouter()
	wsRouter.Get("/ws", func(w http.ResponseWriter, r *http.Request) {
		websocket.ServeWebSocket(as.wsHub, w, r)
	})
	srv, err := httputil.StartHTTPServer(addr, wsRouter)
	if err != nil {
		return fmt.Errorf("failed to start WebSocket server: %w", err)
	}
	log.Info("WebSocket server started", "addr", srv.Addr().String())
	as.wsServer = srv
	return nil
}

func (as *RelayerNode) initRPCClients(ctx context.Context, conf *config.Config) error {
	for i := range conf.RPCs {
		log.Info("Init rpc client", "ChainId", conf.RPCs[i].ChainId, "RpcUrl", conf.RPCs[i].RpcUrl)
		rpc := conf.RPCs[i]
		ethClient, err := node.DialEthClient(ctx, rpc.RpcUrl)
		if err != nil {
			log.Error("dial eth client fail", "err", err)
			return fmt.Errorf("failed to dial L1 client: %w", err)
		}
		if as.ethClient == nil {
			as.ethClient = make(map[uint64]node.EthClient)
		}
		as.ethClient[rpc.ChainId] = ethClient
		as.chainIdList = append(as.chainIdList, rpc.ChainId)
	}
	log.Info("Init rpc client success")
	return nil
}

func (as *RelayerNode) initDB(ctx context.Context, cfg config.Database) error {
	db, err := database.NewDB(ctx, cfg)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	as.DB = db
	log.Info("Init database success")
	return nil
}

func (as *RelayerNode) initSynchronizer(config *config.Config) error {
	for i := range config.RPCs {
		rpcItem := config.RPCs[i]
		var ContractAddresses []common.Address
		ContractAddresses = append(ContractAddresses, common.HexToAddress(rpcItem.Contracts.PoolManagerAddress))
		ContractAddresses = append(
			ContractAddresses,
			common.HexToAddress(rpcItem.Contracts.MessageManagerAddress),
		)
		cfg := synchronizer.Config{
			Contracts:         ContractAddresses,
			LoopIntervalMsec:  5,
			HeaderBufferSize:  uint(rpcItem.HeaderBufferSize),
			ConfirmationDepth: big.NewInt(int64(1)),
			StartHeight:       big.NewInt(int64(rpcItem.StartBlock)),
			ChainId:           uint(rpcItem.ChainId),
		}
		synchronizerTemp, err := synchronizer.NewSynchronizer(&cfg, as.DB, as.ethClient[rpcItem.ChainId], as.relayerMetrics, as.shutdown)
		if err != nil {
			return err
		}
		if as.Synchronizer == nil {
			as.Synchronizer = make(map[uint64]*synchronizer.Synchronizer)
		}
		as.Synchronizer[rpcItem.ChainId] = synchronizerTemp
		log.Info("Init synchronizer success", "chainId", config.RPCs[i].ChainId)
	}
	return nil
}

func (as *RelayerNode) initEventProcessor(config *config.Config) error {
	for i := range config.RPCs {
		rpcItem := config.RPCs[i]
		cfg := &event.ProcessorConfig{
			LoopInterval:       time.Second * 5,
			StartHeight:        big.NewInt(int64(rpcItem.StartBlock)),
			MsgManagerAddress:  rpcItem.Contracts.MessageManagerAddress,
			PoolManagerAddress: rpcItem.Contracts.PoolManagerAddress,
			ChainId:            strconv.FormatUint(rpcItem.ChainId, 10),
			EventStartBlock:    rpcItem.EventUnpackBlock,
			Epoch:              1000,
		}
		eventProcessor, err := event.NewProcessor(as.DB, cfg, as.relayerMetrics, as.shutdown)
		if err != nil {
			return err
		}
		if as.EventProcessor == nil {
			as.EventProcessor = make(map[uint64]*event.Processor)
		}
		as.EventProcessor[rpcItem.ChainId] = eventProcessor
		log.Info("Init event processor success", "chainId", config.RPCs[i].ChainId)
	}
	return nil
}

func (as *RelayerNode) initWorker(config *config.Config) error {
	var chainIds []string
	for i := range config.RPCs {
		chainIds = append(chainIds, strconv.Itoa(int(config.RPCs[i].ChainId)))
	}
	wkConfig := &worker.WorkerHandleConfig{LoopInterval: time.Second * 5,
		ChainIds: chainIds}
	workerHandle, err := worker.NewWorkerHandle(as.DB, wkConfig, as.wsHub, as.shutdown)
	if err != nil {
		return err
	}
	as.WorkerHandle = workerHandle
	log.Info("Init worker success", "chainIds", chainIds)
	return nil
}

func (as *RelayerNode) startMetricsServer(cfg config.Server) error {
	srv, err := metrics.StartServer(as.metricsRegistry, cfg.Host, cfg.Port)
	if err != nil {
		return fmt.Errorf("metrics server failed to start: %w", err)
	}
	as.metricsServer = srv
	log.Info("metrics server started", "port", cfg.Port, "addr", srv.Addr())
	return nil
}

func (as *RelayerNode) initRelayer(config *config.Config) error {
	var ethClient map[uint64]*ethclient.Client
	var poolMangerAddress map[uint64]string
	var driverEngine map[uint64]driver.DriverEngine
	if err := as.DB.TokenConfig.DeleteTokenConfigBeforeNow(); err != nil {
		log.Error("fail to delete old token config", "err", err)
		return err
	}

	for i := range config.RPCs {
		err := as.DB.TokenConfig.StoreTokenConfigs(as.DB.TokenConfig.BuildTokenConfig(config.RPCs[i]))
		if err != nil {
			log.Error("fail to store token config", "err", err)
			return err
		}
		rpcItem := config.RPCs[i]
		ethClt, err := driver.EthClientWithTimeout(context.Background(), rpcItem.RpcUrl)
		if err != nil {
			log.Error("new eth client fail", "err", err)
			return err
		}
		// 从环境变量读取私钥
		privateKeyHex := os.Getenv("RELAYER_PRIVATE_KEY")
		if privateKeyHex == "" {
			return fmt.Errorf("RELAYER_PRIVATE_KEY environment variable not set")
		}

		ecdsaPrivateKey, err := crypto.HexToECDSA(privateKeyHex)
		if err != nil {
			log.Error("ecdsa format fail", "err", err)
			return err
		}
		log.Info("init relayer start", "chainId", config.RPCs[i].ChainId, "RpcUrl", rpcItem.RpcUrl)
		if ethClient == nil {
			ethClient = make(map[uint64]*ethclient.Client)
		}
		ethClient[rpcItem.ChainId] = ethClt
		if poolMangerAddress == nil {
			poolMangerAddress = make(map[uint64]string)
		}
		poolMangerAddress[rpcItem.ChainId] = rpcItem.Contracts.PoolManagerAddress

		dregConf := &driver.DriverEngineConfig{
			ChainClient:               ethClt,
			ChainId:                   big.NewInt(int64(rpcItem.ChainId)),
			PoolManagerAddress:        common.HexToAddress(rpcItem.Contracts.PoolManagerAddress),
			CallerAddress:             common.HexToAddress(config.CallerAddress),
			PrivateKey:                ecdsaPrivateKey,
			NumConfirmations:          config.NumConfirmations,
			SafeAbortNonceTooLowCount: config.SafeAbortNonceTooLowCount,
		}
		drEngine, err := driver.NewDriverEngine(context.Background(), as.relayerMetrics, dregConf)
		if err != nil {
			log.Error("new drive engine fail", "err", err)
			return err
		}
		if driverEngine == nil {
			driverEngine = make(map[uint64]driver.DriverEngine)
		}
		driverEngine[rpcItem.ChainId] = *drEngine
	}
	relayerProcessor, err := relayer.NewRelayerProcessor(as.DB, ethClient, poolMangerAddress, driverEngine, as.relayerMetrics, as.wsHub, as.shutdown)
	if err != nil {
		log.Error("new relayer processor fail", "err", err)
		return err
	}
	as.relayerProcessor = relayerProcessor
	return nil
}
