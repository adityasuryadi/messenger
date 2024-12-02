package usecase

import (
	"github.com/adityasuryadi/messenger/internal/auth/repository"
	"github.com/google/uuid"
)

type RefreshTokenUseCase struct {
	RefreshTokenRepository *repository.RefreshTokenRepository
}

func NewRefreshTokenUsecase(refreshTokenRepo *repository.RefreshTokenRepository) *RefreshTokenUseCase {
	return &RefreshTokenUseCase{
		RefreshTokenRepository: refreshTokenRepo,
	}
}

func (r *RefreshTokenUseCase) Insert(userId uuid.UUID) error {
	return nil
}
