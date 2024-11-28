package main

import (
	"log/slog"
	"net/http"

	"github.com/adityasuryadi/messenger/config"
	"github.com/adityasuryadi/messenger/helper"
	controller "github.com/adityasuryadi/messenger/internal/auth/delivery/http"
	"github.com/adityasuryadi/messenger/internal/auth/delivery/http/route"
	"github.com/adityasuryadi/messenger/internal/auth/model"
	"github.com/adityasuryadi/messenger/internal/auth/repository"
	"github.com/adityasuryadi/messenger/internal/auth/usecase"
	"github.com/adityasuryadi/messenger/pkg"
)

func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	response := &model.SuccessResponse{
		Status: "OK",
		Code:   200,
		Data:   nil,
	}
	helper.WriteResponseBody(w, response)
}

func main() {
	viper := config.NewViper()

	err := config.Init(
		config.SetConfigFile("config"),
		config.SetConfigType("yaml"),
		config.SetConfigFolder([]string{"./"}),
	)

	configs := config.Get()
	slog.Info("config", configs)

	if err != nil {
		slog.Error("failed to load config", err)
	}

	database, err := config.SetupDB(viper)
	if err != nil {
		panic("Failed to connect to database")
	}

	validation := pkg.NewValidation(database)

	userRepository := repository.NewUserRepository(database)
	authUsecase := usecase.NewAuthUseCase(userRepository)
	authController := controller.NewAuthController(validation, authUsecase)

	mux := route.NewRouter(authController)

	// http.HandleFunc("/", authController.Register)
	port := configs.Service.Port
	http.ListenAndServe(port, mux)
}
