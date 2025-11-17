// Package service 实现跨链桥接节点的 HTTP API 服务
// 职责：
// 1. 初始化和管理 HTTP 服务器
// 2. 配置路由和中间件（CORS、超时、心跳监测等）
// 3. 初始化数据库连接（支持主从配置）
// 4. 连接 Gas 费用 gRPC 服务
// 5. 管理 LRU 缓存
// 6. 提供服务启动、停止和健康检查功能
package service

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/Sandwichzzy/relayer-node/cache"
	"github.com/Sandwichzzy/relayer-node/config"
	"github.com/Sandwichzzy/relayer-node/database"
	"github.com/Sandwichzzy/relayer-node/service/common/httputil"
	"github.com/Sandwichzzy/relayer-node/service/gasfee"
	"github.com/Sandwichzzy/relayer-node/service/routes"
	"github.com/Sandwichzzy/relayer-node/service/service"
	"github.com/ethereum/go-ethereum/log"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// API 路径常量定义
// 所有 API 端点遵循 RESTful 风格，版本号为 v1
const (
	HealthPath           = "/healthz"                  // 健康检查端点，用于监控服务状态
	StakingRecordsV1Path = "/api/v1/staking-records"  // 查询质押记录
	BridgeRecordsV1Path  = "/api/v1/bridge-records"   // 查询跨链桥接记录
	StakingValidV1Path   = "/api/v1/staking-valid"    // 验证质押地址有效性
	BridgeValidV1Path    = "/api/v1/bridge-valid"     // 验证桥接地址有效性
	BridgeGasFeeV1Path   = "/api/v1/bridge-price-fee" // 查询跨链 Gas 费用
)

// APIConfig API 服务配置
// 包含 HTTP 服务器和监控指标服务器的配置
type APIConfig struct {
	HTTPServer    config.Server // HTTP API 服务器配置（端口、主机等）
	MetricsServer config.Server // 监控指标服务器配置
}

// API 核心结构体，管理整个 HTTP API 服务的生命周期
type API struct {
	router        *chi.Mux             // HTTP 路由器，基于 chi 框架
	apiServer     *httputil.HTTPServer // HTTP API 服务器实例
	metricsServer *httputil.HTTPServer // 监控指标服务器实例
	db            *database.DB         // 数据库连接实例
	stopped       atomic.Bool          // 原子布尔值，标识服务是否已停止
}

// NewAPI 创建并初始化一个新的 API 实例
//
// 参数：
//   - ctx: 上下文，用于控制初始化过程
//   - cfg: 配置对象，包含数据库、服务器、缓存等所有配置
//
// 返回：
//   - *API: 初始化成功的 API 实例
//   - error: 如果初始化失败，会返回错误并自动清理已创建的资源
//
// 初始化流程：
//   1. 初始化数据库连接
//   2. 初始化路由和中间件
//   3. 启动 HTTP 服务器
func NewAPI(ctx context.Context, cfg *config.Config) (*API, error) {
	out := &API{}
	if err := out.initFromConfig(ctx, cfg); err != nil {
		// 如果初始化失败，自动调用 Stop 清理资源
		return nil, errors.Join(err, out.Stop(ctx))
	}
	return out, nil
}

// initFromConfig 从配置文件初始化 API 服务的所有组件
//
// 参数：
//   - ctx: 上下文
//   - cfg: 配置对象
//
// 返回：
//   - error: 初始化过程中的任何错误
//
// 初始化步骤：
//   1. 初始化数据库连接（主库或从库）
//   2. 初始化路由、中间件和处理器
//   3. 启动 HTTP 服务器
func (a *API) initFromConfig(ctx context.Context, cfg *config.Config) error {
	if err := a.initDB(ctx, cfg); err != nil {
		return fmt.Errorf("failed to init DB: %w", err)
	}
	a.initRouter(cfg)
	if err := a.startServer(cfg.Server); err != nil {
		return fmt.Errorf("failed to start API server: %w", err)
	}
	return nil
}

// initRouter 初始化 HTTP 路由器和所有中间件
//
// 参数：
//   - conf: 配置对象
//
// 初始化流程：
//   1. 创建验证器和缓存实例
//   2. 连接到 Gas 费用 gRPC 服务
//   3. 创建服务层实例
//   4. 配置路由处理器
//   5. 设置中间件（超时、恢复、心跳、CORS）
//   6. 注册 API 路由
func (a *API) initRouter(conf *config.Config) {
	// 创建验证器实例，用于验证地址有效性
	v := new(service.Validator)

	// 创建 LRU 缓存实例（如果启用）
	var lruCache = new(cache.LruCache)
	if conf.EnableApiCache {
		lruCache = cache.NewLruCache()
	}

	// 连接到 Gas 费用 gRPC 服务
	// 使用不安全的凭证（适用于内网环境）
	conn, err := grpc.NewClient(conf.GasOracleEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error("Connect to da retriever fail", "err", err)
		return
	}

	// 创建 gRPC 客户端，用于查询 Gas 费用
	client := gasfee.NewTokenGasPriceServicesClient(conn)

	// 创建服务层实例，连接验证器、数据库和 Gas 费用客户端
	svc := service.New(v, a.db.StakingRecord, a.db.BridgeRecord, client)

	// 创建路由器实例
	apiRouter := chi.NewRouter()

	// 创建路由处理器，连接服务层和缓存
	h := routes.NewRoutes(apiRouter, svc, conf.EnableApiCache, lruCache)

	// 配置中间件
	apiRouter.Use(middleware.Timeout(time.Second * 12)) // 设置请求超时时间为 12 秒
	apiRouter.Use(middleware.Recoverer)                 // 自动恢复 panic，防止服务崩溃
	apiRouter.Use(middleware.Heartbeat(HealthPath))     // 心跳监测端点

	// 配置 CORS（跨域资源共享）
	apiRouter.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},                                                      // 允许所有来源（生产环境应限制）
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},               // 允许的 HTTP 方法
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"}, // 允许的请求头
		ExposedHeaders:   []string{"Link"},                                                   // 暴露的响应头
		AllowCredentials: true,                                                               // 允许发送凭证
		MaxAge:           300,                                                                // 预检请求缓存时间（秒）
	}).Handler)

	// 处理 OPTIONS 预检请求
	apiRouter.Options("/*", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// 注册 API 路由
	// Staking 相关接口
	apiRouter.Get(fmt.Sprintf(StakingRecordsV1Path), h.StakingRecordsHandler) // 查询质押记录
	apiRouter.Get(fmt.Sprintf(StakingValidV1Path), h.StakingValidHandler)     // 验证质押地址

	// Bridge 相关接口
	apiRouter.Get(fmt.Sprintf(BridgeRecordsV1Path), h.BridgeRecordsHandler) // 查询桥接记录
	apiRouter.Get(fmt.Sprintf(BridgeValidV1Path), h.BridgeValidHandler)     // 验证桥接地址
	apiRouter.Get(fmt.Sprintf(BridgeGasFeeV1Path), h.BridgeGasFeeHandler)   // 查询 Gas 费用

	a.router = apiRouter
}

// initDB 初始化数据库连接
//
// 参数：
//   - ctx: 上下文
//   - cfg: 配置对象
//
// 返回：
//   - error: 连接失败时返回错误
//
// 功能：
//   - 支持主从数据库配置
//   - 根据 SlaveDbEnable 配置决定连接主库还是从库
//   - 从库用于读操作，减轻主库压力
func (a *API) initDB(ctx context.Context, cfg *config.Config) error {
	var initDb *database.DB
	var err error

	// 根据配置选择连接主库或从库
	if !cfg.SlaveDbEnable {
		// 连接主数据库
		initDb, err = database.NewDB(ctx, cfg.MasterDb)
		if err != nil {
			log.Error("failed to connect to master database", "err", err)
			return err
		}
	} else {
		// 连接从数据库（读写分离场景）
		initDb, err = database.NewDB(ctx, cfg.SlaveDb)
		if err != nil {
			log.Error("failed to connect to slave database", "err", err)
			return err
		}
	}

	a.db = initDb
	return nil
}

// Start 启动 API 服务
//
// 参数：
//   - ctx: 上下文
//
// 返回：
//   - error: 启动失败时返回错误
//
// 注意：实际的服务器启动在 initFromConfig 中完成
// 此方法预留用于未来的启动后操作
func (a *API) Start(ctx context.Context) error {
	return nil
}

// Stop 优雅地停止 API 服务
//
// 参数：
//   - ctx: 上下文，用于控制停止超时
//
// 返回：
//   - error: 停止过程中的任何错误（可能是多个错误的合并）
//
// 停止流程：
//   1. 停止 HTTP API 服务器
//   2. 停止监控指标服务器
//   3. 关闭数据库连接
//   4. 设置 stopped 标志为 true
//
// 注意：
//   - 即使某个组件停止失败，也会继续停止其他组件
//   - 所有错误会被合并并返回
func (a *API) Stop(ctx context.Context) error {
	var result error

	// 停止 API 服务器
	if a.apiServer != nil {
		if err := a.apiServer.Stop(ctx); err != nil {
			result = errors.Join(result, fmt.Errorf("failed to stop API server: %w", err))
		}
	}

	// 停止监控指标服务器
	if a.metricsServer != nil {
		if err := a.metricsServer.Stop(ctx); err != nil {
			result = errors.Join(result, fmt.Errorf("failed to stop metrics server: %w", err))
		}
	}

	// 关闭数据库连接
	if a.db != nil {
		if err := a.db.Close(); err != nil {
			result = errors.Join(result, fmt.Errorf("failed to close DB: %w", err))
		}
	}

	// 标记服务已停止
	a.stopped.Store(true)
	log.Info("API service shutdown complete")
	return result
}

// startServer 启动 HTTP 服务器
//
// 参数：
//   - serverConfig: 服务器配置（主机和端口）
//
// 返回：
//   - error: 启动失败时返回错误
//
// 功能：
//   - 根据配置的主机和端口启动 HTTP 服务器
//   - 记录服务器监听地址
func (a *API) startServer(serverConfig config.Server) error {
	log.Debug("API server listening...", "port", serverConfig.Port)

	// 构造监听地址（host:port）
	addr := net.JoinHostPort(serverConfig.Host, strconv.Itoa(serverConfig.Port))

	// 启动 HTTP 服务器
	srv, err := httputil.StartHTTPServer(addr, a.router)
	if err != nil {
		return fmt.Errorf("failed to start API server: %w", err)
	}

	log.Info("API server started", "addr", srv.Addr().String())
	a.apiServer = srv
	return nil
}

// Stopped 检查 API 服务是否已停止
//
// 返回：
//   - bool: true 表示服务已停止，false 表示服务仍在运行
//
// 用途：
//   - 用于健康检查
//   - 用于优雅关闭流程中的状态判断
func (a *API) Stopped() bool {
	return a.stopped.Load()
}
