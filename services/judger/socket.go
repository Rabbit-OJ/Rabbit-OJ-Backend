package judger

import (
	"Rabbit-OJ-Backend/controllers/upgrader"
	SubmissionService "Rabbit-OJ-Backend/services/submission"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"time"
)

var (
	judgeHub *Hub
)

type Client struct {
	conn *websocket.Conn
	send chan []byte
	sid  string
}

func (c *Client) readPump() {
	c.conn.SetReadLimit(upgrader.MaxMessageSize)
	if err := c.conn.SetReadDeadline(time.Now().Add(upgrader.PongWait)); err != nil {
		fmt.Println(err)
	}
	c.conn.SetPongHandler(func(string) error {
		_ = c.conn.SetReadDeadline(time.Now().Add(upgrader.PongWait))
		return nil
	})
}

func (c *Client) writePump() {
	ticker := time.NewTicker(upgrader.PingPeriod)
	defer func() {
		ticker.Stop()
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
		sid := c.Param("sid")

		submission, err := SubmissionService.Detail(sid)
		if err != nil || submission.Status != "ING" {
			return
		}

		conn, err := upgrader.Upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			fmt.Println(err)
			return
		}

		client := &Client{conn: conn, send: make(chan []byte, 256), sid: sid}
		JudgeHub.Register <- client

		go client.writePump()
		go client.readPump()
	}
}

type Hub struct {
	clients    map[string]*Client
	Broadcast  chan string
	Register   chan *Client
	unregister chan *Client
}

func NewJudgeHub() *Hub {
	judgeHub = &Hub{
		clients:    make(map[string]*Client),
		Broadcast:  make(chan string),
		Register:   make(chan *Client),
		unregister: make(chan *Client),
	}

	return judgeHub
}

func (h *Hub) removeJudgeHubClient(sid string) {
	if client, ok := h.clients[sid]; ok {
		close(client.send)
		delete(h.clients, sid)
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			if _, ok := h.clients[client.sid]; ok {
				close(h.clients[client.sid].send)
				delete(h.clients, client.sid)
			}

			h.clients[client.sid] = client
		case client := <-h.unregister:
			h.removeJudgeHubClient(client.sid)
		case sid := <-h.Broadcast:
			if client, ok := h.clients[sid]; ok {
				client.send <- []byte("{\"ok\":1}")
				h.removeJudgeHubClient(client.sid)
			}
		}
	}
}