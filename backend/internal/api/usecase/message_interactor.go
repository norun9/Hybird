package usecase

import (
	"context"
	"github.com/norun9/Hybird/internal/api/domain/model"
	"github.com/norun9/Hybird/internal/api/domain/repository"
	"github.com/norun9/Hybird/internal/api/usecase/dto/input"
	"github.com/norun9/Hybird/internal/api/usecase/dto/output"
	"github.com/norun9/Hybird/pkg/util"
)

type IMessageInputBoundary interface {
	Create(ctx context.Context, input input.MessageInput) (*output.MessageOutput, error)
	List(ctx context.Context, p input.MessageList) ([]*output.MessageOutput, error)
}

// NOTE:OutputBoundary interface definition is omitted to prevent code bloat.

type messageInteractor struct {
	messageRepository repository.IMessageRepository
}

// NewMessageInteractor Polymorphism
func NewMessageInteractor(messageRepository repository.IMessageRepository) IMessageInputBoundary {
	return &messageInteractor{
		messageRepository,
	}
}

func (i *messageInteractor) Create(ctx context.Context, p input.MessageInput) (result *output.MessageOutput, err error) {
	var created *model.Message
	if created, err = i.messageRepository.Create(ctx, &model.Message{
		ID:      0,
		Content: p.Content,
	}); err != nil {
		return nil, err
	}
	createdAt := util.TimeOnly12HrFormatter(created.CreatedAt)
	return &output.MessageOutput{
		Content:   created.Content,
		CreatedAt: createdAt,
	}, nil
}

func (i *messageInteractor) List(ctx context.Context, p input.MessageList) (result []*output.MessageOutput, err error) {
	var messages []*model.Message
	if messages, err = i.messageRepository.List(ctx); err != nil {
		return nil, err
	}
	for _, message := range messages {
		createdAt := util.TimeOnly12HrFormatter(message.CreatedAt)
		result = append(result, &output.MessageOutput{
			Content:   message.Content,
			CreatedAt: createdAt,
		})
	}
	return result, nil
}
