package ws

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/adityasuryadi/messenger/internal/messaging/model"
	"github.com/gorilla/websocket"
)

type WebSocketConnection struct {
	*websocket.Conn
}

type Client struct {
	Conn    *websocket.Conn
	Request chan *model.MessageRequest
	hub     *Hub
}

type Message struct {
	Sender  *Client
	Message *model.MessageRequest
}

var clients = make(map[*Client]bool)

// worker for handle read message
func (c *Client) writePump(sender *Client) {
	defer c.Conn.Close()
	for {
		select {
		case message := <-c.Request:
			err := c.Conn.WriteJSON(message)
			if err != nil {
				return
			}
		}
	}
}

// worker for handle write message
func (c *Client) readPump(sender *Client) {
	// if connection disconected delete client from avaiable client
	defer func() {
		c.hub.unregister <- c
		c.Conn.Close()
	}()

	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err) {
				slog.Error("failed to read message", slog.String("error", fmt.Sprintf("%v", err)))
			}
			return
		}

		msg := new(model.MessageRequest)
		err = json.Unmarshal(message, msg)
		if err != nil {
			slog.Error("failed to unmarshal message", slog.String("error", fmt.Sprintf("%v", err)))
			return
		}
		fmt.Println("read pump", msg)
		c.hub.broadcast <- &Message{Sender: c, Message: msg}
	}
}

// handle websocket connected
func HandleWebsocket(hub *Hub, w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:    1024,
		WriteBufferSize:   1024,
		EnableCompression: true,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		slog.Error("failed to upgrade connection", slog.String("error", err.Error()))
	}

	// register koneksi ke clients\
	client := &Client{
		Conn:    conn,
		hub:     hub,
		Request: make(chan *model.MessageRequest),
	}
	client.hub.register <- client

	go client.readPump(client)
	go client.writePump(client)
}
