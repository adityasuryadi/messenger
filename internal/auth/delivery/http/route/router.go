package route

import (
	"net/http"

	controller "github.com/adityasuryadi/messenger/internal/auth/delivery/http"
)

func SetupRouter(mux *http.ServeMux, AuthController *controller.AuthController) {
	prefix := "api"
	mux.HandleFunc("POST /"+prefix+"/register", AuthController.Register)
	mux.HandleFunc("POST /"+prefix+"/login", AuthController.Login)
	mux.HandleFunc("POST /"+prefix+"/refresh-token", AuthController.RefreshToken)
}
