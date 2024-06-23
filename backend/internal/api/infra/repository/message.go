package repository

import (
	"context"
	"github.com/cockroachdb/errors"
	"github.com/norun9/Hybird/internal/api/domain/model"
	domainRepository "github.com/norun9/Hybird/internal/api/domain/repository"
	"gorm.io/gorm"
)

type messageRepository struct {
	dbClient *gorm.DB
}

// NewMessageRepository Polymorphism
func NewMessageRepository(dbClient *gorm.DB) domainRepository.MessageRepository {
	return &messageRepository{dbClient}
}

func (r *messageRepository) List(ctx context.Context) (result []*model.Message, err error) {
	// TODO: Infinite Loading using Offset
	if err := r.dbClient.WithContext(ctx).Find(&result).Limit(100).Error; err != nil {
		return nil, errors.Wrap(err, "failed to list Message models")
	}
	return result, nil
}

func (r *messageRepository) Create(ctx context.Context, model *model.Message) (*model.Message, error) {
	if err := r.dbClient.WithContext(ctx).Create(model).Error; err != nil {
		return nil, errors.Wrap(err, "failed to create Message model")
	}
	return model, nil
}
