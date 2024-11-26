package usecase

import "github.com/adityasuryadi/messenger/internal/repository"

type UserUseCase struct {
	UserRepo *repository.UserRepository
}

func NewUserUseCase(userRepo *repository.UserRepository) *UserUseCase {
	return &UserUseCase{
		UserRepo: userRepo,
	}
}
