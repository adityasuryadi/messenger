package main

import (
	"net/http"

	"github.com/adityasuryadi/messenger/config"
	"github.com/adityasuryadi/messenger/helper"
	"github.com/adityasuryadi/messenger/internal/delivery/http/route"
	controller "github.com/adityasuryadi/messenger/internal/handler/http"
	"github.com/adityasuryadi/messenger/internal/model"
	"github.com/adityasuryadi/messenger/internal/repository"
	"github.com/adityasuryadi/messenger/internal/usecase"
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
	http.ListenAndServe(":8090", mux)
}
