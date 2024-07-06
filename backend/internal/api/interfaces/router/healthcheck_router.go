package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetHealthCheckRoutes(r *gin.RouterGroup) {
	r.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "status OK")
	})
}
