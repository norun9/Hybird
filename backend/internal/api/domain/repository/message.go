package repository

import (
	"context"
	"github.com/norun9/Hybird/internal/api/domain/model"
)

type MessageRepository interface {
	List(ctx context.Context) ([]*model.Message, error)
	Create(ctx context.Context, model *model.Message) (*model.Message, error)
}
