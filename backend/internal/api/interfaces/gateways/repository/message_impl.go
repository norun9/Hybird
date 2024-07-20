package repository

import (
	"context"
	"github.com/cockroachdb/errors"
	"github.com/norun9/Hybird/internal/api/domain/model"
	"github.com/norun9/Hybird/internal/api/interfaces/gateways/transfer"
	"github.com/norun9/Hybird/internal/api/usecase/repository"
	"github.com/norun9/Hybird/pkg/db"
	"github.com/norun9/Hybird/pkg/dbmodels"
	"github.com/norun9/Hybird/pkg/merror"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type messageRepository struct {
	dbClient db.Client
}

// NewMessageRepository Polymorphism
func NewMessageRepository(dbClient db.Client) repository.IMessageRepository {
	return &messageRepository{dbClient}
}

func (r *messageRepository) GetCount(ctx context.Context, queryMods ...db.Query) (totalCount int64, err error) {
	queries := append(queryMods, []qm.QueryMod{}...)
	if totalCount, err = db.Count(ctx, r.dbClient.Get(ctx), dbmodels.TableNames.Messages, queries); err != nil {
		return 0, errors.Wrap(err, merror.ErrDatabase.Error())
	}
	return totalCount, nil
}

func (r *messageRepository) List(ctx context.Context, queryMods ...db.Query) (result []*model.Message, err error) {
	queries := append(queryMods, []qm.QueryMod{
		db.Distinct(dbmodels.TableNames.Messages),
	}...)
	var entities []*dbmodels.Message
	if entities, err = dbmodels.Messages(queries...).All(ctx, r.dbClient.Get(ctx)); err != nil {
		return nil, errors.Wrap(err, merror.ErrDatabase.Error())
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
	entity := transfer.ToMessageEntity(model)
	if err := entity.Insert(ctx, r.dbClient.Get(ctx), boil.Infer()); err != nil {
		return nil, errors.Wrap(err, merror.ErrDatabase.Error())
	}
	return model, nil
}
