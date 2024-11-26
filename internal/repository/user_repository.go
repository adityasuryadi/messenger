package repository

import (
	"context"
	"log/slog"

	"github.com/adityasuryadi/messenger/internal/entity"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UserRepository struct {
	DB *mongo.Database
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) Insert(user *entity.User) {
	result, err := r.DB.Collection("user").InsertOne(context.TODO(), user)
	if err != nil {
		slog.Error("failed to insert user", err)
	}
	slog.Info("inserted a single document: ", result.InsertedID)
}
