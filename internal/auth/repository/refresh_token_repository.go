package repository

import (
	"context"

	"github.com/adityasuryadi/messenger/internal/auth/entity"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type RefreshTokenRepository struct {
	DB *mongo.Database
}

func NewRefreshTokenRepository(db *mongo.Database) *RefreshTokenRepository {
	return &RefreshTokenRepository{
		DB: db,
	}
}

func (r *RefreshTokenRepository) Insert(entity *entity.RefreshToken) error {
	_, err := r.DB.Collection("refresh_tokens").InsertOne(context.TODO(), entity)
	if err != nil {
		return err
	}
	return nil
}

func (r *RefreshTokenRepository) FindUserByToken(token string) (*entity.RefreshToken, error) {
	var refreshToken *entity.RefreshToken
	err := r.DB.Collection("refresh_tokens").FindOne(context.TODO(), bson.M{"refresh_token": token}).Decode(&refreshToken)
	if err != nil {
		return nil, err
	}
	return refreshToken, nil
}
