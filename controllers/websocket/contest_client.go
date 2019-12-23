package websocket

import (
	ContestService "Rabbit-OJ-Backend/services/contest"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"time"
)

type ContestClient struct {
	conn *websocket.Conn
	send chan []byte
	cid  string
	uid  string
}

func (c *ContestClient) readPump() {
	c.conn.SetReadLimit(maxMessageSize)
	if err := c.conn.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
		fmt.Println(err)
	}
	c.conn.SetPongHandler(func(string) error { _ = c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
}

func (c *ContestClient) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
	}()

	for {
		select {
		case message, ok := <-c.send:
			if err := c.conn.SetWriteDeadline(time.Now().Add(writeWait)); err != nil {
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
			if err := c.conn.SetWriteDeadline(time.Now().Add(writeWait)); err != nil {
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

func serveContestWs(c *gin.Context) {
	cid, uid := c.Param("cid"), c.Param("uid")

	state, err := ContestService.CheckContestState(cid)
	if err != nil {
		fmt.Println(err)
		return
	}
	if state != ContestService.StatusPending {
		return
	}

	participate, err := ContestService.User(uid, cid)
	if err != nil {
		fmt.Println(err)
		return
	}
	if participate == nil {
		return
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &ContestClient{conn: conn, send: make(chan []byte, 256), cid: cid}
	SocketHub.ContestHub.register <- client

	go client.writePump()
	go client.readPump()
}
