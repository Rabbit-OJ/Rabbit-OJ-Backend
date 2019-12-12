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

func (h *Hub) remove(client *Client) {
	if _, ok := h.clients[client.sid]; ok {
		close(client.send)
		delete(h.clients, client.sid)
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
			h.remove(client)
		case sid := <-h.Broadcast:
			if client, ok := h.clients[sid]; ok {
				client.send <- []byte("{\"ok\":1}")
				h.remove(client)
			}
		}
	}
}
