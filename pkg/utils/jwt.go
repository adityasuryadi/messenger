package utils

import (
	"log/slog"
	"time"

	"github.com/adityasuryadi/messenger/config"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JwtClaims struct {
	ID uuid.UUID
	jwt.RegisteredClaims
}

func GenerateJwtToken(id uuid.UUID) (string, error) {
	jwtSecret := config.Get().JWT.SecretJWT
	ttl := config.Get().JWT.Ttl
	if jwtSecret == "" {
		slog.Error("jwt secret is empty")
	}
	secretKey := []byte(jwtSecret)

	claims := JwtClaims{
		ID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * time.Duration(ttl))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        id.String(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(secretKey)

	if err != nil {
		return "", err
	}

	return ss, nil
}
