package ws

import (
	"fmt"
)

type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan *Message

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan *Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

// untuk worker

func (h *Hub) RunWorker() {
	for {
		select {
		// handle if client register
		case client := <-h.register:
			h.clients[client] = true
			fmt.Println("client terdaftar", len(h.clients))

			// handle if client unregister
		case client := <-h.unregister:
			fmt.Println("client terunregister", client)
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.Request)
			}

			// broadcast message ke semua client ke cuali sender
		case message := <-h.broadcast:
			fmt.Println("broadcast", message)
			for client := range h.clients {
				if client != message.Sender {
					select {
					case client.Request <- message.Message:
					}
				}
			}
		}
	}
}
