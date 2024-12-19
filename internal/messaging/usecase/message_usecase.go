package usecase

import (
	"log/slog"

	"github.com/adityasuryadi/messenger/internal/messaging/entity"
	"github.com/adityasuryadi/messenger/internal/messaging/model"
	"github.com/adityasuryadi/messenger/internal/messaging/repository"
)

type MessageUsecase struct {
	MessagingRepository *repository.MessageRepository
}

func NewMessageUsecase(messagingRepository *repository.MessageRepository) *MessageUsecase {
	return &MessageUsecase{
		MessagingRepository: messagingRepository,
	}
}

func (u *MessageUsecase) Insert(request *model.MessageRequest) error {
	messageEntity := &entity.Message{
		From:    request.From,
		Message: request.Message,
	}
	err := u.MessagingRepository.Insert(messageEntity)
	if err != nil {
		slog.Error("failed to insert message", slog.String("error", err.Error()))
		return err
	}
	return nil
}
