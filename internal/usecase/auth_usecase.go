package usecase

import (
	"github.com/adityasuryadi/messenger/helper"
	"github.com/adityasuryadi/messenger/internal/entity"
	"github.com/adityasuryadi/messenger/internal/model"
	"github.com/adityasuryadi/messenger/internal/repository"
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
		Password: helper.Encrypt(request.Password),
	}

	u.UserRepository.Insert(user)
}
