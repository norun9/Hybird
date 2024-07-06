package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/norun9/Hybird/internal/api/injector"
	"time"
)

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		//AllowOrigins:   TODO:[]string{"https://.com"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposeHeaders:    []string{"Link"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	handler := injector.InitializeRestHandler()
	handler.GetHealthCheckRoute(r)

	if err := r.Run(":8080"); err != nil {
		//log.Fatalf("failed to run server: %v", err)
	}
}
