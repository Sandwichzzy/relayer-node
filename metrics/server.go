// Package metrics - server.go
// 提供 Prometheus 监控指标的 HTTP 服务器启动功能
package metrics

import (
	"net"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/Sandwichzzy/relayer-node/service/common/httputil"
)

// StartServer 启动 Prometheus 监控指标 HTTP 服务器
//
// 参数：
//   - r: Prometheus 注册表，包含所有已注册的指标
//   - hostname: 监听主机地址（如 "0.0.0.0" 监听所有网卡，"127.0.0.1" 仅本地访问）
//   - port: 监听端口号（通常使用 9090 或其他未占用端口）
//
// 返回：
//   - *httputil.HTTPServer: HTTP 服务器实例
//   - error: 启动失败时返回错误
//
// 功能：
//  1. 创建 HTTP 服务器暴露 Prometheus 指标
//  2. 使用 promhttp.Handler 提供标准的 /metrics 端点
//  3. 支持 Prometheus 服务器定期抓取（scrape）指标数据
//
// 暴露的端点：
//   - GET /metrics: 返回所有注册的指标数据（Prometheus 文本格式）
//
// 使用流程：
//  1. 创建注册表：registry := NewRegistry()
//  2. 创建并注册指标：metrics := NewSyncerMetrics(registry, "syncer")
//  3. 启动服务器：server, err := StartServer(registry, "0.0.0.0", 9090)
//  4. 配置 Prometheus 抓取：
//     scrape_configs:
//     - job_name: 'relayer-node'
//     static_configs:
//     - targets: ['localhost:9090']
//
// 访问示例：
//
//	curl http://localhost:9090/metrics
//
// 输出格式示例：
//
//	# HELP syncer_chain_block_height Different chain block height
//	# TYPE syncer_chain_block_height gauge
//	syncer_chain_block_height{chain_id="1"} 18500000
//	syncer_chain_block_height{chain_id="56"} 32000000
func StartServer(r *prometheus.Registry, hostname string, port int) (*httputil.HTTPServer, error) {
	// 构造监听地址（host:port）
	addr := net.JoinHostPort(hostname, strconv.Itoa(port))

	// 创建 Prometheus HTTP Handler
	// InstrumentMetricHandler 会额外收集 HTTP 请求的指标
	h := promhttp.InstrumentMetricHandler(
		r, promhttp.HandlerFor(r, promhttp.HandlerOpts{}),
	)

	// 启动 HTTP 服务器
	return httputil.StartHTTPServer(addr, h)
}
