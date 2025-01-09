package usecase

import (
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/adityasuryadi/messenger/internal/auth/entity"
	"github.com/adityasuryadi/messenger/internal/auth/model"
	"github.com/adityasuryadi/messenger/internal/auth/repository"
	"github.com/adityasuryadi/messenger/pkg/security"
	"github.com/adityasuryadi/messenger/pkg/utils"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func NewAuthUseCase(userRepo *repository.UserRepository, refreshTokenRepo *repository.RefreshTokenRepository) AuthUseCase {
	return &AuthUseCaseImpl{
		UserRepository:         userRepo,
		RefreshTokenRepository: refreshTokenRepo,
	}
}

type AuthUseCaseImpl struct {
	UserRepository         *repository.UserRepository
	RefreshTokenRepository *repository.RefreshTokenRepository
}

func (u *AuthUseCaseImpl) Register(request *model.RegisterRequest) {
	user := &entity.User{
		Id:       uuid.New(),
		FullName: request.Fullname,
		Email:    request.Email,
		Password: security.Hash(request.Password),
	}

	u.UserRepository.Insert(user)
}

func (u *AuthUseCaseImpl) Login(request *model.LoginRequest) (*model.LoginResponse, error) {

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

	refreshToken, err := utils.GenerateRefreshToken(user.Id)
	if err != nil {
		slog.Error("failed to generate refresh token", slog.String("error", err.Error()))
		return nil, err
	}

	fmt.Println("user id", user)

	// insert refresh token to collection
	refreshTokenEntity := &entity.RefreshToken{
		UserId:       user.Id,
		RefreshToken: refreshToken,
		Email:        user.Email,
		ExpiredAt:    time.Now().Add(time.Hour * 24 * 30),
	}

	err = u.RefreshTokenRepository.Insert(refreshTokenEntity)
	if err != nil {
		slog.Error("failed to insert refresh token", slog.String("error", err.Error()))
		return nil, err
	}

	response := &model.LoginResponse{
		AccessToken:  jwtToken,
		RefreshToken: refreshToken,
	}

	return response, nil
}

func (u *AuthUseCaseImpl) RefreshToken(refreshToken string) (string, error) {
	userRfToken, err := u.RefreshTokenRepository.FindUserByToken(refreshToken)
	if err != nil {
		return "", err
	}

	// check token expired or not
	if time.Now().Unix() > userRfToken.ExpiredAt.Unix() {
		return "", errors.New("token expired")
	}

	// generate new jwt token
	jwtToken, err := utils.GenerateJwtToken(userRfToken.UserId)
	if err != nil {
		slog.Error("failed to generate jwt token", slog.String("error", err.Error()))
		return "", err
	}

	return jwtToken, nil
}

func (u *AuthUseCaseImpl) Logout(refreshToken string) error {
	// get token if does not exist return err
	rfToken, err := u.RefreshTokenRepository.FindUserByToken(refreshToken)
	if rfToken == nil {
		return err
	}

	err = u.RefreshTokenRepository.Delete(refreshToken)
	if err != nil {
		slog.Error("failed to delete refresh token", slog.String("error", err.Error()))
		return err
	}
	return nil
}
