package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/norun9/Hybird/internal/api/injector"
	"github.com/norun9/Hybird/pkg/config"
	"go.uber.org/zap"
	"log"
	"time"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}

	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			logger.Fatal("failed to sync zap logger", zap.Error(err))
		}
	}(logger)

	r := gin.Default()

	r.SetTrustedProxies(nil)

	// Add a ginzap middleware, which:
	//   - Logs all requests, like a combined access and error log.
	//   - Logs to stdout.
	//   - RFC3339 with UTC time format.
	r.Use(ginzap.GinzapWithConfig(logger, &ginzap.Config{
		UTC:        false,
		TimeFormat: time.RFC3339,
	}))

	// Logs all panic to error log
	//   - stack means whether output the stack info.
	r.Use(ginzap.RecoveryWithZap(logger, true))

	// NOTE:using code: gin.SetMode(gin.ReleaseMode) in production
	gin.SetMode(gin.DebugMode)

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://sample.com"}, // TODO:FIX
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposeHeaders:    []string{"Link"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	c := config.Prepare()

	handler := injector.InitializeRestHandler(c.DBConfig)
	handler.GetHealthCheckRoute(r)

	if err := r.Run(":8080"); err != nil {
		logger.Fatal("failed to run server", zap.Error(err))
	}
}
