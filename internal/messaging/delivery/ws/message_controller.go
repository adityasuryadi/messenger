package ws

import (
	"net/http"

	"github.com/adityasuryadi/messenger/internal/messaging/usecase"
)

type MessageController struct {
	MessageUsecase *usecase.MessageUsecase
	Hub            *Hub
}

func NewMessageController(hub *Hub, messageUsecase *usecase.MessageUsecase) *MessageController {
	return &MessageController{
		MessageUsecase: messageUsecase,
		Hub:            hub,
	}
}

func (c *MessageController) InsertChat(w http.ResponseWriter, r *http.Request) {
	c.MessageUsecase.Insert(nil)
}

func (c *MessageController) InitWSConn(w http.ResponseWriter, r *http.Request) {
	// c.MessageUsecase.UpgradeConn(w, r)
	HandleWebsocket(c.Hub, w, r)
}
