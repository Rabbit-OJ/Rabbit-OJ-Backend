package websocket

type Hub struct {
	clients    map[string]*Client
	Broadcast  chan string
	register   chan *Client
	unregister chan *Client
}

func newHub() *Hub {
	return &Hub{
		clients:    make(map[string]*Client),
		Broadcast:  make(chan string),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			if _, ok := h.clients[client.sid]; ok {
				close(h.clients[client.sid].send)
				delete(h.clients, client.sid)
			}

			h.clients[client.sid] = client
		case client := <-h.unregister:
			if _, ok := h.clients[client.sid]; ok {
				close(client.send)
				delete(h.clients, client.sid)
			}
		case sid := <-h.Broadcast:
			for clientSid := range h.clients {
				if clientSid == sid {
					h.clients[sid].send <- []byte("OK")
					break
				}
			}
		}
	}
}
