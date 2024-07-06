package repository

import (
	"context"
	"database/sql"
	"github.com/norun9/Hybird/internal/api/domain/model"
	"github.com/norun9/Hybird/internal/api/domain/repository"
)

type messageRepository struct {
	dbClient *sql.DB
}

// NewMessageRepository Polymorphism
func NewMessageRepository(dbClient *sql.DB) repository.IMessageRepository {
	return &messageRepository{dbClient}
}

func (r *messageRepository) List(ctx context.Context) (result []*model.Message, err error) {
	// TODO: Infinite Loading using Offset
	//if err := r.dbClient.WithContext(ctx).Find(&result).Limit(100).Error; err != nil {
	//	return nil, errors.Wrap(err, "failed to list Message models")
	//}
	return result, nil
}

func (r *messageRepository) Create(ctx context.Context, model *model.Message) (*model.Message, error) {
	//if err := r.dbClient.WithContext(ctx).Create(model).Error; err != nil {
	//	return nil, errors.Wrap(err, "failed to create Message model")
	//}
	return model, nil
}
