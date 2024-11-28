package repository

import (
	"context"
	"log/slog"

	"github.com/adityasuryadi/messenger/internal/auth/entity"
	"go.mongodb.org/mongo-driver/bson"
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

func (r *UserRepository) FindUserByEmail(email string) (*entity.User, error) {
	result := new(entity.User)
	err := r.DB.Collection("user").FindOne(context.TODO(), bson.M{"email": email}).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
