package submission

import (
	"Rabbit-OJ-Backend/controllers/upgrader"
	"Rabbit-OJ-Backend/models/data_structure"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"strconv"
	"time"
)

var (
	JudgeHub *Hub
)

type Client struct {
	conn *websocket.Conn
	send chan []byte
	sid  uint32
}

func (c *Client) readPump() {
	defer func() {
		JudgeHub.unregister <- c
		_ = c.conn.Close()
	}()

	c.conn.SetReadLimit(upgrader.MaxMessageSize)
	if err := c.conn.SetReadDeadline(time.Now().Add(upgrader.PongWait)); err != nil {
		fmt.Println(err)
	}
	c.conn.SetPongHandler(func(string) error {
		_ = c.conn.SetReadDeadline(time.Now().Add(upgrader.PongWait))
		return nil
	})

	for {
		_, _, err := c.conn.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
	}
}

func (c *Client) writePump() {
	ticker := time.NewTicker(upgrader.PingPeriod)
	defer func() {
		ticker.Stop()
		_ = c.conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			if err := c.conn.SetWriteDeadline(time.Now().Add(upgrader.WriteWait)); err != nil {
				fmt.Println(err)
			}

			if !ok {
				if err := c.conn.WriteMessage(websocket.CloseMessage, []byte{}); err != nil {
					fmt.Println(err)
				}
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				fmt.Println(err)
				return
			}

			if _, err := w.Write(message); err != nil {
				fmt.Println(err)
			}

			n := len(c.send)
			for i := 0; i < n; i++ {
				if _, err := w.Write(<-c.send); err != nil {
					fmt.Println(err)
				}
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			if err := c.conn.SetWriteDeadline(time.Now().Add(upgrader.WriteWait)); err != nil {
				fmt.Println(err)
				return
			}
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

func ServeJudgeWs(JudgeHub *Hub) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		_sid := c.Param("sid")
		sid, err := strconv.ParseUint(_sid, 10, 32)

		if err != nil {
			fmt.Println(err)
			return
		}

		submission, err := Detail(uint32(sid))
		if err != nil || submission.Status != "ING" {
			return
		}

		conn, err := upgrader.Upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			fmt.Println(err)
			return
		}

		client := &Client{conn: conn, send: make(chan []byte, 256), sid: uint32(sid)}
		JudgeHub.Register <- client

		go client.writePump()
		go client.readPump()
	}
}

type Hub struct {
	clients    data_structure.ConcurrentMap
	Broadcast  chan uint32
	Register   chan *Client
	unregister chan *Client
}

func NewJudgeHub() *Hub {
	JudgeHub = &Hub{
		clients:    data_structure.MakeConcurrentHashmap(),
		Broadcast:  make(chan uint32),
		Register:   make(chan *Client),
		unregister: make(chan *Client),
	}

	return JudgeHub
}

func (h *Hub) handleRemoveJudgeHubClient(sid uint32) {
	if _client, ok := h.clients.Load(sid); ok {
		client := _client.(*Client)
		close(client.send)
		h.clients.Delete(sid)
	}
}

func (h *Hub) handleRegister(client *Client) {
	sid := client.sid
	if _client, ok := h.clients.Load(sid); ok {
		client := _client.(*Client)

		close(client.send)
		h.clients.Delete(sid)
	}

	h.clients.Store(sid, client)
}

func (h *Hub) handleBroadcast(sid uint32) {
	if _client, ok := h.clients.Load(sid); ok {
		client := _client.(*Client)

		client.send <- []byte("{\"ok\":1}")
		h.handleRemoveJudgeHubClient(client.sid)
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			go h.handleRegister(client)
		case client := <-h.unregister:
			go h.handleRemoveJudgeHubClient(client.sid)
		case sid := <-h.Broadcast:
			go h.handleBroadcast(sid)
		}
	}
}
