package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/norun9/Hybird/internal/api/usecase/dto/input"
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
		params := input.MessageInput{
			Content: c.Request.URL.Query().Get("content"),
		}
		h.Exec(c, params)
	})
}
