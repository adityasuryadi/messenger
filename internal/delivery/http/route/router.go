package route

import (
	"net/http"

	controller "github.com/adityasuryadi/messenger/internal/handler/http"
)

func NewRouter(AuthController *controller.AuthController) *http.ServeMux {
	mux := http.NewServeMux()
	prefix := "api"
	mux.HandleFunc("POST /"+prefix+"/register", AuthController.Register)
	return mux
}
