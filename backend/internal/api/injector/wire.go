//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/google/wire"
	"github.com/norun9/Hybird/internal/api/infra/repository"
	"github.com/norun9/Hybird/internal/api/usecase"
	"github.com/norun9/Hybird/pkg/db"
)

var interactorSet = wire.NewSet(
	db.NewDB,
	repository.NewMessageRepository,
	usecase.NewMessageInteractor,
)

func InitializeMessageInteractor() (usecase.MessageInputBoundary, error) {
	wire.Build(interactorSet)
	return &usecase.MessageInteractor{}, nil
}
