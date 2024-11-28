package usecase

import (
	"github.com/adityasuryadi/messenger/internal/auth/entity"
	"github.com/adityasuryadi/messenger/internal/auth/repository"
)

type UserUseCase struct {
	UserRepo *repository.UserRepository
}

func NewUserUseCase(userRepo *repository.UserRepository) *UserUseCase {
	return &UserUseCase{
		UserRepo: userRepo,
	}
}

func (u *UserUseCase) FindUserByEmail(email string) (*entity.User, error) {
	panic("implement me")
}
