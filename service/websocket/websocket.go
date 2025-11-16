package websocket

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/log"
	"github.com/gorilla/websocket"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

type Hub struct {
	clients    map[*Client]bool // 存储所有客户端
	broadcast  chan []byte      // 广播消息的通道
	register   chan *Client     // 注册客户端的通道
	unregister chan *Client     // 注销客户端的通道
	mu         sync.RWMutex     // 读写锁，用于并发安全
}
type Client struct {
	hub    *Hub
	conn   *websocket.Conn
	send   chan []byte // 发送消息的通道
	mu     sync.Mutex  // 互斥锁，用于并发安全
	closed bool        // 标记连接是否关闭
}

type Message struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
	Time int64       `json:"time"`
}

type BridgeFinalizedMessage struct {
	Type        string `json:"type"`
	FromAddress string `json:"from_address"`
	ToAddress   string `json:"to_address"`
	TxHash      string `json:"tx_hash"`
	Status      int64  `json:"status"`
	Time        int64  `json:"time"`
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

// Hub的主循环，处理客户端的注册、注销和广播消息。
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			// 注册客户端
			h.mu.Lock()
			h.clients[client] = true
			h.mu.Unlock()
			log.Info("WebSocket client connected", "total_clients", len(h.clients))

		case client := <-h.unregister:
			h.mu.Lock()
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
			h.mu.Unlock()
			log.Info("WebSocket client disconnected", "total_clients", len(h.clients))

		case message := <-h.broadcast:
			h.mu.RLock()
			for client := range h.clients {
				select {
				case client.send <- message: // 消息发送成功
				default:
					// 如果发送失败，则关闭连接并移除客户端
					close(client.send)
					delete(h.clients, client)
				}
			}
			h.mu.RUnlock()
		}
	}
}

// 广播BridgeFinalizedMessage类型的消息。
func (h *Hub) BroadcastBridgeFinalized(msg *BridgeFinalizedMessage) {
	msg.Type = "bridge_finalized"
	msg.Time = time.Now().Unix()

	data, err := json.Marshal(msg)
	if err != nil {
		log.Error("Failed to marshal bridge finalized message", "err", err)
		return
	}

	h.broadcast <- data
	log.Info("Broadcasted bridge finalized event", "tx_hash", msg.TxHash)
}

func (h *Hub) GetClientCount() int {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return len(h.clients)
}

func (h *Hub) CloseAllClients() {
	h.mu.Lock()
	defer h.mu.Unlock()

	for client := range h.clients {
		client.Close()
		delete(h.clients, client)
	}
	log.Info("All WebSocket clients closed")
}

// : HTTP请求处理，只处理路径为"/ws"的请求，其他返回404。
func (h *Hub) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/ws" {
		ServeWebSocket(h, w, r)
	} else {
		http.NotFound(w, r)
	}
}

// 客户端从WebSocket连接中读取消息。 读循环（接收消息）
func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c // 退出时注销
		c.conn.Close()
	}()

	// 设置读取限制和超时
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))

	// Pong处理器，用于保持连接
	c.conn.SetPongHandler(func(string) error {
		c.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			// 处理连接关闭错误
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Error("WebSocket read error", "err", err)
			}
			break
		}

		log.Debug("Received WebSocket message", "message", string(message))
	}
}

// 向WebSocket连接写入消息，并定期发送ping消息。
func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}

			w.Write(message)

			if err := w.Close(); err != nil {
				return
			}

		case <-ticker.C:
			// 发送心跳ping
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (c *Client) Close() {
	c.mu.Lock()
	defer c.mu.Unlock()

	if !c.closed {
		c.closed = true
		c.conn.Close()
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func ServeWebSocket(hub *Hub, w http.ResponseWriter, r *http.Request) {
	// 升级HTTP连接到WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error("WebSocket upgrade failed", "err", err)
		return
	}
	// 创建新客户端
	client := &Client{
		hub:  hub,
		conn: conn,
		send: make(chan []byte, 256),
	}

	client.hub.register <- client

	// 注册客户端并启动读写循环
	go client.writePump()
	go client.readPump()
}
