package repository

import (
	"context"
	"github.com/norun9/Hybird/internal/api/domain/model"
	"github.com/norun9/Hybird/pkg/db"
)

type IMessageRepository interface {
	GetCount(ctx context.Context, queryMods ...db.Query) (int64, error)
	List(ctx context.Context, queryMods ...db.Query) ([]*model.Message, error)
	Create(ctx context.Context, model *model.Message) (*model.Message, error)
}
