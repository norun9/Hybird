// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package injector

import (
	"github.com/google/wire"
	"github.com/norun9/Hybird/internal/api/external/db"
	"github.com/norun9/Hybird/internal/api/interfaces"
	"github.com/norun9/Hybird/internal/api/interfaces/controllers"
	"github.com/norun9/Hybird/internal/api/interfaces/gateways/repository"
	"github.com/norun9/Hybird/internal/api/usecase"
	"github.com/norun9/Hybird/pkg/config"
	db2 "github.com/norun9/Hybird/pkg/db"
)

// Injectors from wire.go:

func InitializeRestHandler(dbConfig config.DBConfig) interfaces.IRestHandler {
	sqlDB := db.NewDB(dbConfig)
	client := db2.NewDBClient(sqlDB)
	iMessageRepository := repository.NewMessageRepository(client)
	iPaging := repository.NewPaging()
	iMessageInputBoundary := usecase.NewMessageInteractor(iMessageRepository, iPaging)
	iMessageController := controllers.NewMessageController(iMessageInputBoundary)
	v := interfaces.GetMapRoute(iMessageController)
	iRestHandler := interfaces.NewRestHandler(v)
	return iRestHandler
}

// wire.go:

var inputBoundarySet = wire.NewSet(db.NewDB, db2.NewDBClient, repository.NewMessageRepository, usecase.NewMessageInteractor)
