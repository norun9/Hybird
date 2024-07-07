package interfaces

import (
	"github.com/gin-gonic/gin"
)

type IRestHandler interface {
	Exec(c *gin.Context, params interface{})
	GetRoute(r *gin.Engine)
	GetHealthCheckRoute(r *gin.Engine)
}

type restHandler struct {
	routeMap map[Path]Handler
}

func NewRestHandler(routeMap map[Path]Handler) IRestHandler {
	return &restHandler{routeMap}
}

func (h *restHandler) GetRoute(r *gin.Engine) {
	v1 := r.Group("/v1")
	h.GetMessageRoutes(v1)
}

func (h *restHandler) GetHealthCheckRoute(r *gin.Engine) {
	v1 := r.Group("/v1")
	h.GetHealthCheckRoutes(v1)
}

func (h *restHandler) Exec(c *gin.Context, params interface{}) {
	//ctx := c.Request.Context()
	//r := c.Request
	//w := c.Writer
}
