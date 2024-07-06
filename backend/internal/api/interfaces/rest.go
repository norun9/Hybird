package interfaces

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/norun9/Hybird/internal/api/interfaces/router"
)

type RestHandler interface {
	Exec(ctx context.Context, c *gin.Context, params interface{})
	GetRoute(r *gin.Engine)
	GetHealthCheckRoute(r *gin.Engine)
}

type restHandler struct{}

func NewRestHandler() RestHandler {
	return &restHandler{}
}

func (h *restHandler) GetRoute(r *gin.Engine) {
	v1 := r.Group("/v1")
	router.GetMessageRoutes(v1)
}

func (h *restHandler) GetHealthCheckRoute(r *gin.Engine) {
	v1 := r.Group("/v1")
	router.GetHealthCheckRoutes(v1)
}

func (h *restHandler) Exec(ctx context.Context, c *gin.Context, params interface{}) {
}
