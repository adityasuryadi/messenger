package usecase

import (
	"github.com/adityasuryadi/messenger/internal/auth/entity"
	"github.com/adityasuryadi/messenger/internal/auth/model"
	"github.com/adityasuryadi/messenger/internal/auth/repository"
	"github.com/adityasuryadi/messenger/pkg/security"
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
