//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/google/wire"
	"github.com/norun9/Hybird/internal/api/infra/repository"
	"github.com/norun9/Hybird/internal/api/usecase"
	"github.com/norun9/Hybird/pkg/db"
)

var inputBoundarySet = wire.NewSet(
	db.NewDB,
	repository.NewMessageRepository,
	usecase.NewMessageInteractor,
)

func InitializeMessageInteractor() (usecase.IMessageInputBoundary, error) {
	wire.Build(inputBoundarySet)
	return &usecase.MessageInteractor{}, nil
}
