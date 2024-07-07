package interfaces

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *restHandler) GetHealthCheckRoutes(r *gin.RouterGroup) {
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "Healthy",
		})
	})
}
