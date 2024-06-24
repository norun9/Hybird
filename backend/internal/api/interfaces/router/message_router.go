package router

import (
	"github.com/gin-gonic/gin"
	"github.com/norun9/Hybird/internal/api/interfaces/controller"
)

func SetupCategoryRoutes(r *gin.Engine, controller *controller.CategoryController) {
	routes := r.Group("/v1/messages")
	{
		routes.GET("", categoryController.List)

	}
}
