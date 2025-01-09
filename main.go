package main

import (
	"log/slog"
	"net/http"

	"github.com/adityasuryadi/messenger/config"
	"github.com/adityasuryadi/messenger/internal/server"
)

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			slog.Error("failed to start server", slog.String("error", err.(string)))
		}
	}()
	viper := config.NewViper()

	err := config.Init(
		config.SetConfigFile("config"),
		config.SetConfigType("yaml"),
		config.SetConfigFolder([]string{"./"}),
	)

	configs := config.Get()
	if err != nil {
		slog.Error("failed to load config", slog.String("error", err.Error()))
	}

	database, err := config.SetupDB(viper)
	if err != nil {
		panic("Failed to connect to database")
	}
	mux := http.NewServeMux()

	server.Bootstrap(&server.BootstrapConfig{
		Mux: mux,
		DB:  database,
	})

	handlerWithCORS := enableCORS(mux)

	port := configs.Service.Port
	http.ListenAndServe(port, handlerWithCORS)
}
