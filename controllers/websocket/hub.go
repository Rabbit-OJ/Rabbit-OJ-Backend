package websocket

import (
	"encoding/json"
	"fmt"
)

type Hub struct {
	JudgeHub   *JudgeHub
	ContestHub *ContestHub
}

type JudgeHub struct {
	clients    map[string]*JudgeClient
	Broadcast  chan string
	register   chan *JudgeClient
	unregister chan *JudgeClient
}

type ContestHubBroadcast struct {
	Cid     string `json:"cid"`
	Type    string `json:"type"`
	Message string `json:"message"`
}

type ContestHub struct {
	clients    map[string]*ContestClient
	Broadcast  chan ContestHubBroadcast
	register   chan *ContestClient
	unregister chan *ContestClient
}

func newHub() *Hub {
	hub := Hub{
		JudgeHub:   newJudgeHub(),
		ContestHub: newContestHub(),
	}
	return &hub
}

func newJudgeHub() *JudgeHub {
	return &JudgeHub{
		clients:    make(map[string]*JudgeClient),
		Broadcast:  make(chan string),
		register:   make(chan *JudgeClient),
		unregister: make(chan *JudgeClient),
	}
}

func newContestHub() *ContestHub {
	return &ContestHub{
		clients:    make(map[string]*ContestClient),
		Broadcast:  make(chan ContestHubBroadcast),
		register:   make(chan *ContestClient),
		unregister: make(chan *ContestClient),
	}
}

func (h *JudgeHub) removeJudgeHubClient(sid string) {
	if client, ok := h.clients[sid]; ok {
		close(client.send)
		delete(h.clients, sid)
	}
}

func (h *ContestHub) removeContestHubClient(uid string) {
	if client, ok := h.clients[uid]; ok {
		close(client.send)
		delete(h.clients, uid)
	}
}

func (h *ContestHub) RemoveContestHubAllContest(cid string) {
	for _, client := range h.clients {
		if client.cid == cid {
			close(client.send)
			delete(h.clients, client.uid)
		}
	}
}

func (h *JudgeHub) run() {
	for {
		select {
		case client := <-h.register:
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

func (h *ContestHub) run() {
	for {
		select {
		case client := <-h.register:
			if _, ok := h.clients[client.uid]; ok {
				close(h.clients[client.uid].send)
				delete(h.clients, client.uid)
			}

			h.clients[client.uid] = client
		case client := <-h.unregister:
			h.removeContestHubClient(client.uid)
		case broadcast := <-h.Broadcast:
			jsonByte, err := json.Marshal(broadcast)
			if err != nil {
				fmt.Println(err)
				continue
			}

			for _, item := range h.clients {
				if item.cid == broadcast.Cid {
					item.send <- jsonByte
				}
			}
		}
	}
}
