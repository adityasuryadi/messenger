package usecase

import "github.com/adityasuryadi/messenger/internal/auth/model"

type AuthUseCase interface {
	Register(request *model.RegisterRequest)
	Login(request *model.LoginRequest) (*model.LoginResponse, error)
	RefreshToken(refreshToken string) (string, error)
}
