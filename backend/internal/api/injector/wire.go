//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/norun9/Hybird/pkg/db"
)

var infrastructureSet = wire.NewSet(
	db.NewDB,
)

func InitializeControllers() {
	wire.Build(
		db.NewDB,
	)
	return
}
