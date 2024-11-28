package route

import (
	"net/http"

	controller "github.com/adityasuryadi/messenger/internal/auth/delivery/http"
)

func NewRouter(AuthController *controller.AuthController) *http.ServeMux {
	mux := http.NewServeMux()
	prefix := "api"
	mux.HandleFunc("POST /"+prefix+"/register", AuthController.Register)
	mux.HandleFunc("POST /"+prefix+"/login", AuthController.Login)
	return mux
}
