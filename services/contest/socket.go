package contest

import (
	"Rabbit-OJ-Backend/controllers/upgrader"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"strconv"
	"sync"
	"time"
)

var (
	contestHub *Hub
)

type Client struct {
	conn *websocket.Conn
	send chan []byte
	cid  uint32
	uid  uint32
}

func (c *Client) readPump() {
	defer func() {
		contestHub.unregister <- c
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

func ServeContestWs(contestHub *Hub) func(*gin.Context) {
	return func(c *gin.Context) {
		_cid, _uid := c.Param("cid"), c.Param("uid")

		cid, err := strconv.ParseUint(_cid, 10, 32)
		if err != nil {
			fmt.Println(err)
			return
		}

		uid, err := strconv.ParseUint(_uid, 10, 32)
		if err != nil {
			fmt.Println(err)
			return
		}

		state, err := CheckContestState(uint32(cid))
		if err != nil {
			fmt.Println(err)
			return
		}
		if state != RoundStarting {
			return
		}

		participate, _, err := User(uint32(uid), uint32(cid))
		if err != nil {
			fmt.Println(err)
			return
		}
		if participate == nil {
			return
		}

		conn, err := upgrader.Upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			fmt.Println(err)
			return
		}

		client := &Client{conn: conn, send: make(chan []byte, 256), cid: uint32(cid)}
		contestHub.register <- client

		go client.writePump()
		go client.readPump()
	}
}

type HubBroadcast struct {
	Cid     uint32 `json:"cid"`
	Type    string `json:"type"`
	Message string `json:"message"`
}

type Hub struct {
	clients    sync.Map
	EndContest chan uint32
	Broadcast  chan *HubBroadcast
	register   chan *Client
	unregister chan *Client
}

func NewContestHub() *Hub {
	contestHub = &Hub{
		clients:    sync.Map{},
		Broadcast:  make(chan *HubBroadcast),
		EndContest: make(chan uint32),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}

	return contestHub
}

func (h *Hub) handleRemoveContestHubClient(uid uint32) {
	if _client, ok := h.clients.Load(uid); ok {
		client := _client.(*Client)

		close(client.send)
		h.clients.Delete(uid)
	}
}

func (h *Hub) RemoveContestHubAllContest(cid uint32) {
	var willDelete []uint32
	h.clients.Range(func(_, _client interface{}) bool {
		client := _client.(*Client)

		if client.cid == cid {
			willDelete = append(willDelete, cid)
			close(client.send)
		}

		return true
	})

	for _, item := range willDelete {
		h.clients.Delete(item)
	}
}

func (h *Hub) handleRegisterClient(client *Client) {
	uid := client.uid
	if _client, ok := h.clients.Load(uid); ok {
		client := _client.(*Client)
		close(client.send)

		h.clients.Delete(client.uid)
	}

	h.clients.Store(client.uid, client)
}

func (h *Hub) handleBroadCast(broadcast *HubBroadcast) {
	jsonByte, err := json.Marshal(broadcast)
	if err != nil {
		fmt.Println(err)
		return
	}

	h.clients.Range(func(_, _client interface{}) bool {
		client := _client.(*Client)

		if client.cid == broadcast.Cid {
			client.send <- jsonByte
		}

		return true
	})
}

func (h *Hub) handleEndContest(cid uint32) {
	h.clients.Range(func(_, _client interface{}) bool {
		client := _client.(*Client)

		if client.cid == cid {
			h.handleRemoveContestHubClient(client.uid)
		}

		return true
	})
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			go h.handleRegisterClient(client)
		case client := <-h.unregister:
			go h.handleRemoveContestHubClient(client.uid)
		case broadcast := <-h.Broadcast:
			go h.handleBroadCast(broadcast)
		case cid := <-h.EndContest:
			go h.handleEndContest(cid)
		}
	}
}
