package interfaces

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/norun9/Hybird/internal/api/interfaces/router"
	"net/http"
)

type RestHandler interface {
	Exec(ctx context.Context, w gin.ResponseWriter, r *http.Request, params interface{})
	GetRoute(r *gin.Engine)
	//GetHealthRouter(router chi.Router)
}

type restHandler struct{}

func NewRestHandler() RestHandler {
	return &restHandler{}
}

func (h restHandler) GetRoute(r *gin.Engine) {
	v1 := r.Group("/v1")
	router.GetMessageRoutes(v1)
}

func (h *restHandler) Exec(ctx context.Context, w gin.ResponseWriter, r *http.Request, params interface{}) {
}