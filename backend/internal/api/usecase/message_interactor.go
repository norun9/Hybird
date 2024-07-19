package usecase

import (
	"context"
	"github.com/norun9/Hybird/internal/api/domain/model"
	"github.com/norun9/Hybird/internal/api/usecase/dto/input"
	"github.com/norun9/Hybird/internal/api/usecase/dto/output"
	"github.com/norun9/Hybird/internal/api/usecase/repository"
	"github.com/norun9/Hybird/pkg/db"
	"github.com/norun9/Hybird/pkg/dbmodels"
	"github.com/norun9/Hybird/pkg/util"
)

type IMessageInputBoundary interface {
	Create(ctx context.Context, input input.MessageInput) (*output.MessageOutput, error)
	List(ctx context.Context, p input.MessageList) ([]*output.MessageOutput, error)
}

// NOTE:OutputBoundary interface definition is omitted to prevent code bloat.

type messageInteractor struct {
	messageRepo repository.IMessageRepository
	pagingRepo  repository.Paging
}

// NewMessageInteractor Polymorphism
func NewMessageInteractor(messageRepo repository.IMessageRepository, pagingRepo repository.Paging) IMessageInputBoundary {
	return &messageInteractor{
		messageRepo,
		pagingRepo,
	}
}

func (i *messageInteractor) Create(ctx context.Context, p input.MessageInput) (result *output.MessageOutput, err error) {
	var created *model.Message
	if created, err = i.messageRepo.Create(ctx, &model.Message{
		Content: p.Content,
	}); err != nil {
		return nil, err
	}
	createdAt := util.DateTimeJaFormatter(created.CreatedAt)
	return &output.MessageOutput{
		Content:   created.Content,
		CreatedAt: createdAt,
	}, nil
}

func (i *messageInteractor) List(ctx context.Context, p input.MessageList) (result []*output.MessageOutput, err error) {
	var totalCount int64
	if totalCount, err = i.messageRepo.GetCount(ctx); err != nil {
		return nil, err
	}

	paging := i.pagingRepo.Get(
		int(totalCount),
		p.Paging.Offset,
		p.Paging.Page,
		p.Paging.Limit,
	)

	var messages []*model.Message
	if messages, err = i.messageRepo.List(
		ctx,
		paging,
		db.OrderBy(dbmodels.MessageColumns.CreatedAt, false),
	); err != nil {
		return nil, err
	}
	for _, message := range messages {
		createdAt := util.DateTimeJaFormatter(message.CreatedAt)
		result = append(result, &output.MessageOutput{
			Content:   message.Content,
			CreatedAt: createdAt,
		})
	}
	return result, nil
}
