package interfaces

import (
	"github.com/gin-gonic/gin"
)

func (h *restHandler) GetMessageRoutes(r *gin.RouterGroup) {
	gr := r.Group("/messages")
	gr.GET("/", func(c *gin.Context) {
		var params interface{}
		h.Exec(c, params)
	})
}
