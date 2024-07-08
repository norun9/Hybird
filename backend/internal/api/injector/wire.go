//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/google/wire"
	"github.com/norun9/Hybird/internal/api/infra/repository"
	"github.com/norun9/Hybird/internal/api/interfaces"
	"github.com/norun9/Hybird/internal/api/interfaces/controller"
	"github.com/norun9/Hybird/internal/api/usecase"
	"github.com/norun9/Hybird/pkg/config"
	"github.com/norun9/Hybird/pkg/db"
)

var inputBoundarySet = wire.NewSet(
	db.NewDB,
	db.NewDBClient,
	repository.NewMessageRepository,
	usecase.NewMessageInteractor,
)

func InitializeRestHandler(config.DBConfig) (_ interfaces.IRestHandler) {
	wire.Build(
		inputBoundarySet,
		repository.NewPaging,
		controller.NewMessageController,
		interfaces.GetMapRoute,
		interfaces.NewRestHandler,
	)
	return
}
