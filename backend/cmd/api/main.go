package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/norun9/Hybird/internal/api/injector"
	"github.com/norun9/Hybird/pkg/config"
	"github.com/norun9/Hybird/pkg/log"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"time"
)

var (
	ginLambda *ginadapter.GinLambda
	r         *gin.Engine
	env       string
)

func initRouter() *gin.Engine {
	c := config.Prepare()
	r = gin.Default()

	err := r.SetTrustedProxies(nil)
	if err != nil {
		log.Logger.Fatal("failed to set trusted proxies", zap.Error(err))
	}

	// Add a ginzap middleware, which:
	//   - Logs all requests, like a combined access and error log.
	//   - Logs to stdout.
	//   - RFC3339 with UTC time format.
	r.Use(ginzap.GinzapWithConfig(log.Logger, &ginzap.Config{
		UTC:        false,
		TimeFormat: time.RFC3339,
	}))

	// Logs all panic to error log
	//   - stack means whether output the stack info.
	r.Use(ginzap.RecoveryWithZap(log.Logger, true))

	// NOTE:using code: gin.SetMode(gin.ReleaseMode) in production
	gin.SetMode(c.GinConfig.Mode)

	r.Use(cors.New(cors.Config{
		AllowOrigins:     c.HTTPConfig.CORSConfig.AllowedOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposeHeaders:    []string{"Link"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	handler := injector.InitializeRestHandler(c.DBConfig)
	handler.GetHealthCheckRoute(r)
	handler.GetRoute(r)

	return r
}

func init() {
	log.InitLogger()
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			logger.Fatal("failed to sync zap logger", zap.Error(err))
		}
	}(log.Logger)

	r = initRouter()

	env = viper.GetString("env.name")
}

func main() {
	switch env {
	case "dev":
		if err := r.Run(":8080"); err != nil {
			log.Logger.Fatal("failed to run server", zap.Error(err))
		}
	case "prd":
		ginLambda = ginadapter.New(r)
		lambda.Start(Handler)
	default:
		log.Logger.Fatal("unknown environment", zap.String("env", env))
	}
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}
