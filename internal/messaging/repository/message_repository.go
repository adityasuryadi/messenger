package repository

import (
	"context"
	"log/slog"

	"github.com/adityasuryadi/messenger/internal/messaging/entity"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type MessageRepository struct {
	DB *mongo.Database
}

func NewMessageRepository(Db *mongo.Database) *MessageRepository {
	return &MessageRepository{
		DB: Db,
	}
}

func (r *MessageRepository) Insert(entity *entity.Message) error {
	_, err := r.DB.Collection("messages").InsertOne(context.TODO(), entity)
	if err != nil {
		slog.Error("failed to insert message", slog.String("error", err.Error()))
		return err
	}
	return nil
}
