package usecase

import (
	"errors"
	"log/slog"

	"github.com/adityasuryadi/messenger/internal/auth/entity"
	"github.com/adityasuryadi/messenger/internal/auth/model"
	"github.com/adityasuryadi/messenger/internal/auth/repository"
	"github.com/adityasuryadi/messenger/pkg/security"
	"github.com/adityasuryadi/messenger/pkg/utils"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func NewAuthUseCase(userRepo *repository.UserRepository) *AuthUseCase {
	return &AuthUseCase{
		UserRepository: userRepo,
	}
}

type AuthUseCase struct {
	UserRepository *repository.UserRepository
}

func (u *AuthUseCase) Register(request *model.RegisterRequest) {
	user := &entity.User{
		FullName: request.Fullname,
		Email:    request.Email,
		Password: security.Hash(request.Password),
	}

	u.UserRepository.Insert(user)
}

func (u *AuthUseCase) Login(request *model.LoginRequest) (*model.LoginResponse, error) {

	// find user by email
	user, err := u.UserRepository.FindUserByEmail(request.Email)
	if err != nil && errors.Is(err, mongo.ErrNoDocuments) {
		return nil, errors.New("user not found")
	}

	// checking password
	if !security.CompareHash(request.Password, user.Password) {
		slog.Error("failed to find user", slog.String("error", errors.New("wrong email or password").Error()))
		return nil, errors.New("wrong email or password")
	}

	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		slog.Error("failed to find user", slog.String("error", err.Error()))
		return nil, err
	}

	// generate token jwt
	jwtToken, err := utils.GenerateJwtToken(user.Id)
	if err != nil {
		slog.Error("failed to generate jwt token", slog.String("error", err.Error()))
		return nil, err
	}

	response := &model.LoginResponse{
		AccessToken:  jwtToken,
		RefreshToken: "",
	}

	return response, nil
}
