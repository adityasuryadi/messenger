package main

import (
	"log/slog"
	"net/http"

	"github.com/adityasuryadi/messenger/config"
	"github.com/adityasuryadi/messenger/internal/server"
)

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

	port := configs.Service.Port
	http.ListenAndServe(port, mux)
}
