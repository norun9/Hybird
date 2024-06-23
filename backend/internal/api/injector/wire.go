//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/norun9/Hybird/internal/api/infra/repository"
	"github.com/norun9/Hybird/internal/api/usecase"
	"github.com/norun9/Hybird/pkg/db"
)

var interactorSet = wire.NewSet(
	db.NewDB,
	repository.NewMessageRepository,
	usecase.NewMessageInteractor,
)

func InitializeController() {
	wire.Build(
		interactorSet,
	)
	return
}
