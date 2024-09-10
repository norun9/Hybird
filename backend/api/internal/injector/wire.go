//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/google/wire"
	db2 "github.com/norun9/HyBird/backend/api/internal/external/db"
	"github.com/norun9/HyBird/backend/api/internal/interfaces"
	"github.com/norun9/HyBird/backend/api/internal/interfaces/controllers"
	"github.com/norun9/HyBird/backend/api/internal/interfaces/gateways/repository"
	"github.com/norun9/HyBird/backend/api/internal/usecase"
	"github.com/norun9/HyBird/backend/api/lib/config"
	"github.com/norun9/HyBird/backend/api/lib/db"
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
