package repository

import (
	"context"
	"github.com/norun9/Hybird/internal/api/domain/model"
	"github.com/norun9/Hybird/pkg/db"
)

type IMessageRepository interface {
	List(ctx context.Context, queryMods ...db.Query) ([]*model.Message, error)
	Create(ctx context.Context, model *model.Message) (int64, error)
}
