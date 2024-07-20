//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/google/wire"
	db2 "github.com/norun9/Hybird/internal/api/external/db"
	"github.com/norun9/Hybird/internal/api/interfaces"
	"github.com/norun9/Hybird/internal/api/interfaces/controllers"
	"github.com/norun9/Hybird/internal/api/interfaces/gateways/repository"
	"github.com/norun9/Hybird/internal/api/usecase"
	"github.com/norun9/Hybird/pkg/config"
	"github.com/norun9/Hybird/pkg/db"
)

var inputBoundarySet = wire.NewSet(
	db2.NewDB,
	db.NewDBClient,
	repository.NewMessageRepository,
	usecase.NewMessageInteractor,
)

func InitializeRestHandler(config.DBConfig) (_ interfaces.IRestHandler) {
	wire.Build(
		inputBoundarySet,
		repository.NewPaging,
		controllers.NewMessageController,
		interfaces.GetMapRoute,
		interfaces.NewRestHandler,
	)
	return
}
