package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/norun9/Hybird/pkg/config"
	"log"
)

func BootServer(c config.AppConfig) {
	r := gin.Default()

	//router.SetupRoutes(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
