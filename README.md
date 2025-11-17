# Relayer Node - 跨链桥中继节点

<div align="center">

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-13+-316192?style=flat&logo=postgresql)](https://www.postgresql.org)

一个高性能、模块化的跨链桥中继节点，支持多链资产跨链转移和流动性挖矿。

[特性](#-核心特性) • [架构](#-系统架构) • [安装](#-快速开始) • [配置](#-配置说明) • [API](#-api-接口)

</div>

---

## 📋 目录

- [核心特性](#-核心特性)
- [系统架构](#-系统架构)
- [技术栈](#-技术栈)
- [快速开始](#-快速开始)
- [配置说明](#-配置说明)
- [业务流程](#-业务流程)
- [API 接口](#-api-接口)
- [监控与指标](#-监控与指标)
- [安全注意事项](#-安全注意事项)
- [开发指南](#-开发指南)
- [故障排查](#-故障排查)

---

## 🌟 核心特性

### 跨链桥功能
- ✅ **多链支持**: 支持以太坊、BSC、Polygon 等多条 EVM 兼容链
- ✅ **双向跨链**: ETH/ERC20 代币在不同链之间自由转移
- ✅ **自动中继**: 自动监听源链事件并在目标链完成资产转移
- ✅ **交易管理**: 智能 Gas 管理、自动重试、确认等待

### 流动性挖矿
- 💰 **LP 质押**: 支持 ETH 和 ERC20 代币质押提供流动性
- 💰 **自动奖励**: 跨链手续费自动分配给流动性提供者
- 💰 **灵活提取**: 支持本金和奖励分离提取

### 高可用性
- 🔄 **断点续传**: 支持中断后从上次位置继续同步
- 🔄 **多实例部署**: 支持主从数据库和负载均衡
- 🔄 **实时通知**: WebSocket 实时推送跨链状态更新
- 🔄 **指标监控**: Prometheus 指标暴露，支持 Grafana 可视化

---

## 🏗️ 系统架构

### 三阶段处理流程

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   Synchronizer  │ -> │ Event Processor │ -> │ Relayer/Worker  │
│   (同步原始数据)  │    │  (解析业务数据)  │    │  (执行和记录)    │
└─────────────────┘    └─────────────────┘    └─────────────────┘
        │                      │                       │
        v                      v                       v
   区块头 + 事件日志         业务事件表            跨链交易执行
   (ContractEvents)      (BridgeInitiate,        + 用户记录
                         BridgeFinalize)         (BridgeRecords)
```

### 核心模块

| 模块 | 职责 | 关键组件 |
|------|------|----------|
| **Synchronizer** | 同步区块链数据到数据库 | `synchronizer/synchronizer.go` |
| **Event Processor** | 解析事件为业务数据 | `event/event.go` |
| **Relayer** | 执行跨链交易 | `relayer/relayer.go`, `driver/driver.go` |
| **Worker** | 管理用户查询记录 | `worker/worker.go` |
| **TxManager** | 交易生命周期管理 | `relayer/txmgr/txmgr.go` |
| **API Service** | HTTP/WebSocket 服务 | `service/api.go`, `service/routes/` |
| **Metrics** | 监控指标收集 | `metrics/metrics.go` |

---

## 🛠️ 技术栈

- **语言**: Go 1.21+
- **数据库**: PostgreSQL 13+
- **区块链**: go-ethereum (geth)
- **Web 框架**: Chi Router
- **监控**: Prometheus
- **WebSocket**: gorilla/websocket
- **缓存**: hashicorp/golang-lru

---

## 🚀 快速开始

### 前置要求

```bash
# 1. 安装 Go 1.21+
go version

# 2. 安装 PostgreSQL 13+
psql --version

# 3. 准备以太坊 RPC 节点
# - Alchemy, Infura, 或自建节点
```

### 安装步骤

```bash
# 1. 克隆项目
git clone https://github.com/your-org/relayer-node.git
cd relayer-node

# 2. 安装依赖
go mod download

# 3. 创建数据库
psql -U postgres
CREATE DATABASE relayernode;
\q

# 4. 配置文件
cp relayer-node.local.yaml.example relayer-node.local.yaml
vim relayer-node.local.yaml

# 5. 设置私钥环境变量（重要！）
export RELAYER_PRIVATE_KEY="your_private_key_hex_without_0x"

# 6. 编译运行
go build -o relayer-node
./relayer-node --config relayer-node.local.yaml
```

---

## ⚙️ 配置说明

### 配置文件结构

```yaml
# relayer-node.local.yaml

# ⚠️ 不要在配置文件中写入私钥！使用环境变量 RELAYER_PRIVATE_KEY
caller_address: "0x55225359b717dA1EA4270F78ddA384b0A9f53E28"
num_confirmations: 3                    # 交易确认块数（建议 3-12）
safe_abort_nonce_too_low_count: 3      # Nonce 错误重试次数
enable_api_cache: true                  # 启用 API 缓存（生产环境推荐）
slave_db_enable: false                  # 启用从库（读写分离）

# HTTP API 服务
server:
  host: 0.0.0.0
  port: 8082

# WebSocket 服务
websocket_server:
  host: 0.0.0.0
  port: 8084

# Prometheus 指标服务
metrics:
  host: 0.0.0.0
  port: 8083

# RPC 链配置（支持多链）
rpcs:
  # 以太坊 Sepolia 测试网
  - rpc_url: 'https://eth-sepolia.g.alchemy.com/v2/YOUR_API_KEY'
    chain_id: 11155111
    start_block: 9618149                # Synchronizer 起始区块
    event_unpack_block: 9618149         # Event Processor 起始区块
    header_buffer_size: 500
    contracts:
      pool_manager_address: "0x9B3F87aa9ABbC18b78De9fF245cc945F794F7559"
      message_manager_address: "0x81Ec84f2ADE4e28717f72957F8ABEF85675f2501"

  # 自定义链
  - rpc_url: 'https://rpc-testnet.example.com'
    chain_id: 90101
    start_block: 1650800
    event_unpack_block: 1650800
    header_buffer_size: 500
    contracts:
      pool_manager_address: "0x9B3F87aa9ABbC18b78De9fF245cc945F794F7559"
      message_manager_address: "0x81Ec84f2ADE4e28717f72957F8ABEF85675f2501"

# 主数据库（写）
master_db:
  db_host: "127.0.0.1"
  db_port: 5432
  db_user: "postgres"
  db_password: "your_password"
  db_name: "relayernode"

# 从数据库（读）- 可选
slave_db:
  db_host: "127.0.0.1"
  db_port: 5433
  db_user: "postgres"
  db_password: "your_password"
  db_name: "relayernode"
```

### 环境变量

```bash
# 必需
export RELAYER_PRIVATE_KEY="your_private_key_here"

# 可选
export DB_PASSWORD="your_db_password"
export LOG_LEVEL="info"  # debug, info, warn, error
```

---

## 📊 业务流程

### 1. 完整跨链流程

```
┌─────────────────────────────────────────────────────────────────┐
│                     源链 (Chain A)                               │
├─────────────────────────────────────────────────────────────────┤
│ 用户操作: 发起跨链转账                                            │
│ ├─ ETH 转账 → InitiateETH 事件                                   │
│ └─ USDT 转账 → InitiateERC20 事件                                │
│                                                                  │
│ 存入数据库: bridge_initiate (Status=0)                           │
└─────────────────────────────────────────────────────────────────┘
                              ↓
┌─────────────────────────────────────────────────────────────────┐
│                     Relayer 节点                                 │
├─────────────────────────────────────────────────────────────────┤
│ 1. Synchronizer 同步事件                                         │
│ 2. Event Processor 解析为 BridgeInitiate                        │
│ 3. Relayer 查询未处理记录                                        │
│ 4. 构造目标链 bridgeFinalize 交易                               │
│ 5. TxManager 发送交易（自动重试、Gas 管理）                      │
└─────────────────────────────────────────────────────────────────┘
                              ↓
┌─────────────────────────────────────────────────────────────────┐
│                     目标链 (Chain B)                             │
├─────────────────────────────────────────────────────────────────┤
│ Relayer 调用 finalize()                                          │
│ ├─ 转账 ETH → FinalizeETH 事件                                   │
│ └─ 转账 ERC20 → FinalizeERC20 事件                               │
│                                                                  │
│ 存入数据库: bridge_finalize (Status=0)                           │
└─────────────────────────────────────────────────────────────────┘
                              ↓
┌─────────────────────────────────────────────────────────────────┐
│                     Relayer 节点                                 │
├─────────────────────────────────────────────────────────────────┤
│ 6. Worker 监听 BridgeFinalize 事件                               │
│ 7. 更新 BridgeRecord (Status=1, 完成)                           │
│ 8. WebSocket 广播通知前端                                        │
│ 9. 更新 BridgeInitiate (Status=2, 最终确认)                     │
└─────────────────────────────────────────────────────────────────┘
                              ↓
                      ✅ 跨链完成
```

### 2. 状态流转

```
BridgeInitiate 状态：
0 (未处理) → 1 (已发送完成交易) → 2 (最终确认)

BridgeFinalize 状态：
0 (未确认) → 1 (已确认)

BridgeRecord 状态（用户查询）：
0 (待完成) → 1 (已完成)
```

### 3. 数据关联机制

**为什么可以通过 TxHash 关联不同事件？**

```
【源链同一笔交易】
用户调用 bridgeInitiate() 交易 (TxHash: 0xabc...)
     ↓ (触发多个事件)
┌──────────────────┬────────────────────┐
│                  │                    │
BridgeInitiate    BridgeMsgSent
事件 (跨链发起)    事件 (消息发送)

共享相同的 TxHash: 0xabc...
```

**MsgHash 匹配机制：**

```go
// 智能合约生成 MsgHash
MsgHash = keccak256(abi.encodePacked(
    sourceChainId,      // 源链ID
    destChainId,        // 目标链ID
    sourceTokenAddress, // 源代币地址
    destTokenAddress,   // 目标代币地址
    fromAddress,        // 发送方
    toAddress,          // 接收方
    amount,             // 金额
    nonce,              // 唯一nonce
    fee                 // 手续费
));
```

源链和目标链生成相同的 MsgHash，实现跨链消息匹配。

### 4. LP 流动性挖矿流程

```
┌─────────────────────────────────────────────────────────────────┐
│                     LP 提供者操作                                │
├─────────────────────────────────────────────────────────────────┤
│ 1. 质押资产                                                      │
│    ├─ stakeETH() → StakingETHEvent                              │
│    └─ stakeERC20() → StarkingERC20Event                         │
│                                                                  │
│    记录: StartPoolId = 当前 Pool ID (例如 #100)                  │
│                                                                  │
│ 2. 等待期间                                                      │
│    - 用户跨链转账产生手续费                                       │
│    - 手续费按比例分配给 LP                                        │
│    - Pool ID 不断增长: #100 → #101 → #102 ...                   │
│                                                                  │
│ 3a. 提取本金和奖励                                               │
│     withdraw() → Withdraw 事件                                   │
│     - EndPoolId = 当前 Pool ID (例如 #150)                       │
│     - Amount = 本金 (1 ETH)                                      │
│     - RewardAmount = Pool #100-#150 的累积奖励 (0.05 ETH)       │
│                                                                  │
│ 3b. 只领取奖励 (本金继续质押)                                     │
│     claimReward() → ClaimReward 事件                             │
│     - RewardAmount = 累积奖励 (0.05 ETH)                         │
│     - 本金继续赚取收益                                           │
└─────────────────────────────────────────────────────────────────┘
```

---

## 🔌 API 接口

### HTTP API

**基础地址**: `http://localhost:8082/api/v1`

#### 1. 查询跨链记录

```bash
GET /api/v1/bridge-records?page=1&page_size=20&address=0x123...&status=1

# 响应
{
  "code": 0,
  "data": [
    {
      "source_chain_id": 11155111,
      "dest_chain_id": 90101,
      "source_tx_hash": "0xabc...",
      "dest_tx_hash": "0xdef...",
      "from_address": "0x123...",
      "to_address": "0x456...",
      "amount": "1000000000000000000",
      "status": 1,
      "token_name": "ETH",
      "created_at": "2025-01-17T10:00:00Z"
    }
  ],
  "total": 100
}
```

#### 2. 查询质押记录

```bash
GET /api/v1/staking-records?page=1&page_size=20&address=0x123...

# 响应
{
  "code": 0,
  "data": [
    {
      "chain_id": 11155111,
      "address": "0x123...",
      "amount": "5000000000000000000",
      "reward_amount": "250000000000000000",
      "start_pool_id": 100,
      "end_pool_id": 150,
      "status": 1
    }
  ],
  "total": 10
}
```

#### 3. 查询 Gas 费用

```bash
GET /api/v1/bridge-price-fee?source_chain=11155111&dest_chain=90101&token=ETH&amount=1000000000000000000

# 响应
{
  "code": 0,
  "data": {
    "gas_fee": "50000000000000000",  // 0.05 ETH
    "estimated_time": "5-10 minutes"
  }
}
```

#### 4. 验证跨链请求

```bash
POST /api/v1/bridge-valid
Content-Type: application/json

{
  "source_chain_id": 11155111,
  "dest_chain_id": 90101,
  "token_address": "0xabc...",
  "amount": "1000000000000000000"
}

# 响应
{
  "code": 0,
  "data": {
    "valid": true,
    "liquidity_available": true,
    "max_amount": "10000000000000000000"
  }
}
```

### WebSocket 实时通知

**连接地址**: `ws://localhost:8084/ws`

```javascript
// 前端示例
const ws = new WebSocket('ws://localhost:8084/ws');

ws.onmessage = (event) => {
  const data = JSON.parse(event.data);

  if (data.type === 'bridge_finalized') {
    console.log('跨链完成:', data.tx_hash);
    // 状态: 0 = 已发送, 1 = 已确认
    console.log('状态:', data.status);
  }
};
```

---

## 📈 监控与指标

### Prometheus 指标

**访问地址**: `http://localhost:8083/metrics`

**关键指标：**

```
# 链区块高度
chain_block_height{chain_id="11155111"} 9618500

# 事件处理高度
event_block_height{chain_id="11155111"} 9618450

# Relayer 账户余额
native_token_balance{chain_id="11155111"} 1.5

# Relayer 账户 nonce
chain_address_nonce{chain_id="11155111"} 125
```

### Grafana 仪表盘

```bash
# 导入预配置的仪表盘
curl -X POST http://localhost:3000/api/dashboards/db \
  -H "Content-Type: application/json" \
  -d @grafana-dashboard.json
```

**监控项目：**
- ✅ 区块同步进度
- ✅ 事件处理延迟
- ✅ Relayer 账户余额告警
- ✅ 交易成功率
- ✅ Gas 费用统计

---

## 🔒 安全注意事项

### 1. 私钥管理 ⚠️

**生产环境推荐：**
- 🔐 AWS KMS / Google Cloud KMS
- 🔐 HashiCorp Vault
- 🔐 Azure Key Vault

### 2. .gitignore 配置

```bash
# .gitignore
*.local.yaml
*-local.yaml
.env
*.key
*.pem
/etc/relayer/
```

### 3. 文件权限

```bash
# 私钥文件权限（如果使用文件存储）
chmod 400 /etc/relayer/key.txt
chown relayer:relayer /etc/relayer/key.txt
```

### 4. 账户监控

```go
// 添加余额监控告警
if balance.Cmp(minBalance) < 0 {
    log.Error("⚠️ Low balance detected!", "balance", balance)
    // 发送告警通知
}
```

### 5. 网络安全

```bash
# 防火墙配置（只开放必要端口）
ufw allow 8082/tcp  # HTTP API
ufw allow 8083/tcp  # Prometheus
ufw allow 8084/tcp  # WebSocket
ufw enable
```

---

## 👨‍💻 开发指南

### 项目结构

```
relayer-node/
├── bindings/           # 智能合约 Go 绑定
├── cache/              # LRU 缓存实现
├── common/             # 公共工具包
├── config/             # 配置管理
├── database/           # 数据库模型和查询
├── event/              # 事件处理器
│   └── contracts/      # 合约事件解析
├── metrics/            # Prometheus 指标
├── relayer/            # 跨链中继处理器
│   ├── driver/         # 交易驱动引擎
│   └── txmgr/          # 交易管理器
├── service/            # HTTP/WebSocket 服务
│   ├── routes/         # API 路由处理
│   └── websocket/      # WebSocket Hub
├── synchronizer/       # 区块链数据同步器
│   └── node/           # RPC 客户端封装
└── worker/             # 记录管理器
```

### 添加新链支持

1. **更新配置文件**：
```yaml
rpcs:
  - rpc_url: 'https://new-chain-rpc.com'
    chain_id: 12345
    start_block: 1000000
    event_unpack_block: 1000000
    contracts:
      pool_manager_address: "0x..."
      message_manager_address: "0x..."
```

2. **部署合约**：
   - PoolManager 合约
   - MessageManager 合约

3. **重启服务**：
```bash
./relayer-node --config relayer-node.local.yaml
```

### 运行测试

```bash
# 运行所有测试
go test ./...

# 运行特定模块测试
go test ./relayer/txmgr/...

# 带覆盖率
go test -cover ./...

# 生成覆盖率报告
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

### 代码规范

```bash
# 格式化代码
go fmt ./...

# 静态检查
go vet ./...

# 使用 golangci-lint
golangci-lint run
```

---

## 🐛 故障排查

### 常见问题

#### 1. 私钥环境变量未设置

```
Error: RELAYER_PRIVATE_KEY environment variable not set
```

**解决：**
```bash
export RELAYER_PRIVATE_KEY="your_private_key_here"
```

#### 2. 数据库连接失败

```
Error: failed to connect to database: connection refused
```

**解决：**
```bash
# 检查 PostgreSQL 是否运行
systemctl status postgresql

# 检查配置
psql -h 127.0.0.1 -U postgres -d relayernode
```

#### 3. RPC 节点连接失败

```
Error: dial eth client fail: context deadline exceeded
```

**解决：**
- 检查 RPC URL 是否正确
- 检查网络连接
- 确认 API Key 有效（Alchemy/Infura）

#### 4. Nonce Too Low 错误

```
Error: nonce too low
```

**原因：** 可能有多个 Relayer 实例使用同一账户

**解决：**
- 确保只有一个实例运行
- 或使用不同的 Relayer 账户

#### 5. Gas 费用不足

```
Error: insufficient funds for gas * price + value
```

**解决：**
```bash
# 查看账户余额
cast balance 0x55225359b717dA1EA4270F78ddA384b0A9f53E28 --rpc-url https://...

# 转入 ETH 到 Relayer 账户
```

### 日志级别

```bash
# 调试模式
export LOG_LEVEL=debug
./relayer-node --config relayer-node.local.yaml

# 生产模式
export LOG_LEVEL=info
```

### 健康检查

```bash
# 检查服务状态
curl http://localhost:8082/health

# 检查指标
curl http://localhost:8083/metrics | grep chain_block_height
```

---

## 📝 许可证

MIT License - 详见 [LICENSE](LICENSE) 文件

---





