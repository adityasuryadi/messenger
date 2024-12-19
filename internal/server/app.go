package server

import (
	"net/http"

	authController "github.com/adityasuryadi/messenger/internal/auth/delivery/http"
	authRoute "github.com/adityasuryadi/messenger/internal/auth/delivery/http/route"
	"github.com/adityasuryadi/messenger/internal/auth/repository"
	"github.com/adityasuryadi/messenger/internal/auth/usecase"
	"github.com/adityasuryadi/messenger/internal/messaging/delivery/ws"
	messagingController "github.com/adityasuryadi/messenger/internal/messaging/delivery/ws"
	messageRoute "github.com/adityasuryadi/messenger/internal/messaging/delivery/ws/route"
	messagingRepository "github.com/adityasuryadi/messenger/internal/messaging/repository"
	messagingUsecase "github.com/adityasuryadi/messenger/internal/messaging/usecase"
	"github.com/adityasuryadi/messenger/pkg"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type BootstrapConfig struct {
	Mux *http.ServeMux
	DB  *mongo.Database
}

func NewBootstrapConfig(mux *http.ServeMux, db *mongo.Database) *BootstrapConfig {
	return &BootstrapConfig{
		Mux: mux,
		DB:  db,
	}
}

func Bootstrap(config *BootstrapConfig) {
	refreshTokenRepository := repository.NewRefreshTokenRepository(config.DB)

	validation := pkg.NewValidation(config.DB)

	userRepository := repository.NewUserRepository(config.DB)
	authUsecase := usecase.NewAuthUseCase(userRepository, refreshTokenRepository)
	authController := authController.NewAuthController(validation, authUsecase)

	messageRepository := messagingRepository.NewMessageRepository(config.DB)
	messageUsecase := messagingUsecase.NewMessageUsecase(messageRepository)

	hub := ws.NewHub()
	go hub.RunWorker()

	messageController := messagingController.NewMessageController(hub, messageUsecase)

	authRoute.SetupRouter(config.Mux, authController)
	messageRoute.SetupRouter(config.Mux, messageController)

}
