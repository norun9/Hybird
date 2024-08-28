package main

import (
	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/norun9/Hybird/internal/api/injector"
	"github.com/norun9/Hybird/pkg/config"
	"github.com/norun9/Hybird/pkg/log"
	"go.uber.org/zap"
	"time"
)

var r *gin.Engine

func init() {
	log.InitLogger()

	defer log.Sync()

	c := config.Prepare()

	r = gin.Default()

	err := r.SetTrustedProxies(nil)
	if err != nil {
		log.Fatal("failed to set trusted proxies", zap.Error(err))
	}

	// Add a ginzap middleware, which:
	//   - Logs all requests, like a combined access and error log.
	//   - Logs to stdout.
	//   - RFC3339 with UTC time format.
	r.Use(ginzap.GinzapWithConfig(log.Logger(), &ginzap.Config{
		UTC:        false,
		TimeFormat: time.RFC3339,
	}))

	// Logs all panic to error log
	//   - stack means whether output the stack info.
	r.Use(ginzap.RecoveryWithZap(log.Logger(), true))

	// NOTE:using code: gin.SetMode(gin.ReleaseMode) in production
	gin.SetMode(gin.ReleaseMode)

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{c.HTTPConfig.CORSConfig.AllowedOrigin},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposeHeaders:    []string{"Link"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	handler := injector.InitializeRestHandler(c.DBConfig)
	handler.GetHealthCheckRoute(r)
	handler.GetRoute(r)
}

func main() {
	if err := r.Run(":8080"); err != nil {
		log.Fatal("failed to run server", zap.Error(err))
	}
}
