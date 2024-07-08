package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/norun9/Hybird/internal/api/usecase/dto/input"
	"net/http"
)

func (h *restHandler) GetMessageRoutes(r *gin.RouterGroup) {
	gr := r.Group("messages")
	gr.GET("", func(c *gin.Context) {
		params := input.MessageList{
			Paging: GetPagingInfo(c.Request),
		}
		h.Exec(c, params)
	})
	gr.GET("ws", func(c *gin.Context) {
		h.Exec(c, nil)
	})
	gr.POST("", func(c *gin.Context) {
		var params input.MessageInput
		if err := c.BindJSON(&params); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}
		h.Exec(c, params)
	})
}
