package entity

import (
	"time"

	"github.com/google/uuid"
)

type RefreshToken struct {
	RefreshToken string    `bson:"refresh_token"`
	UserId       uuid.UUID `bson:"user_id"`
	Email        string    `bson:"email"`
	ExpiredAt    time.Time `bson:"expired_at"`
}
