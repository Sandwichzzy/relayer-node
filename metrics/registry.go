// Package metrics - registry.go
// 提供 Prometheus 注册表的创建和初始化功能
package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
)

// NewRegistry 创建并初始化 Prometheus 注册表
//
// 返回：
//   - *prometheus.Registry: 新创建的注册表实例
//
// 功能：
//   1. 创建独立的 Prometheus 注册表（不使用全局默认注册表）
//   2. 自动注册进程级别的指标收集器（CPU、内存等）
//   3. 自动注册 Go 运行时指标收集器（Goroutine、GC 等）
//
// 注册的默认指标：
//   - process_*: 进程相关指标
//     * process_cpu_seconds_total: CPU 使用时间
//     * process_resident_memory_bytes: 常驻内存大小
//     * process_virtual_memory_bytes: 虚拟内存大小
//     * process_open_fds: 打开的文件描述符数量
//   - go_*: Go 运行时指标
//     * go_goroutines: Goroutine 数量
//     * go_threads: 系统线程数量
//     * go_gc_duration_seconds: GC 持续时间
//     * go_memstats_*: 内存统计信息
//
// 使用示例：
//   registry := NewRegistry()
//   metrics := NewSyncerMetrics(registry, "syncer")
//   server, err := StartServer(registry, "0.0.0.0", 9090)
//
// 访问指标：
//   curl http://localhost:9090/metrics
func NewRegistry() *prometheus.Registry {
	// 创建新的注册表
	registry := prometheus.NewRegistry()

	// 注册进程收集器（监控 CPU、内存等系统资源）
	registry.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))

	// 注册 Go 运行时收集器（监控 Goroutine、GC 等 Go 特定指标）
	registry.MustRegister(collectors.NewGoCollector())

	return registry
}
