package config

import (
	"net/http"

	// authController "github.com/adityasuryadi/messenger/internal/auth/delivery/http"
	messagingController "github.com/adityasuryadi/messenger/internal/messaging/delivery/ws"
)

type RouteConfig struct {
	Mux *http.ServeMux
	// AuthController      *authController.AuthController
	MessagingController *messagingController.MessageController
}

func (r *RouteConfig) Setup() {
	// prefix := "api"
	r.Mux.HandleFunc("/ws", r.MessagingController.InitWSConn)
	r.Mux.HandleFunc("GET /test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("teszzzt"))
	})
	// r.Mux.HandleFunc("POST /"+prefix+"/register", r.AuthController.Register)
	// r.Mux.HandleFunc("POST /"+prefix+"/login", r.AuthController.Login)
	// r.Mux.HandleFunc("POST /"+prefix+"/refresh-token", r.AuthController.RefreshToken)
}
