package route

import (
	"net/http"

	controller "github.com/adityasuryadi/messenger/internal/messaging/delivery/ws"
)

func SetupRouter(mux *http.ServeMux, MessageController *controller.MessageController) {
	mux.HandleFunc("/ws", MessageController.InitWSConn)
}
