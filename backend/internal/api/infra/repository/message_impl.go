package repository

import (
	"context"
	"github.com/cockroachdb/errors"
	"github.com/norun9/Hybird/internal/api/domain/model"
	"github.com/norun9/Hybird/internal/api/domain/repository"
	"github.com/norun9/Hybird/pkg/db"
	"github.com/norun9/Hybird/pkg/dbmodels"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type messageRepository struct {
	dbClient db.Client
}

// NewMessageRepository Polymorphism
func NewMessageRepository(dbClient db.Client) repository.IMessageRepository {
	return &messageRepository{dbClient}
}

func (r *messageRepository) List(ctx context.Context, queryMods ...db.Query) (result []*model.Message, err error) {
	queries := append(queryMods, []qm.QueryMod{
		db.Distinct(dbmodels.TableNames.Messages),
	}...)
	var entities []*dbmodels.Message
	if entities, err = dbmodels.Messages(queries...).All(ctx, r.dbClient.Get(ctx)); err != nil {
		return nil, errors.Wrap(err, "ErrDatabase")
	}
	for _, entity := range entities {
		result = append(result, &model.Message{
			ID:        entity.ID,
			Content:   entity.Content,
			CreatedAt: entity.CreatedAt,
			UpdatedAt: entity.UpdatedAt,
		})
	}
	return result, nil
}

func (r *messageRepository) Create(ctx context.Context, model *model.Message) (*model.Message, error) {
	//if err := r.dbClient.WithContext(ctx).Create(model).Error; err != nil {
	//	return nil, errors.Wrap(err, "failed to create Message model")
	//}
	return model, nil
}
