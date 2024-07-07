package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/norun9/Hybird/internal/api/domain/model"
	"net/http"
	"strconv"
)

var defaultLimit = 100

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

func GetPagingInfo(r *http.Request) (paging model.Paging) {
	limit := r.URL.Query().Get("limit")
	paramLimit, err := strconv.Atoi(limit)
	if err != nil || paramLimit == 0 {
		paramLimit = defaultLimit
	}

	offset := r.URL.Query().Get("offset")
	paramOffset, err := strconv.Atoi(offset)
	if err != nil {
		paramOffset = 0
	}

	page := r.URL.Query().Get("page")
	paramPage, err := strconv.Atoi(page)
	if err != nil {
		paramPage = 1
	}
	return model.Paging{
		Offset: paramOffset,
		Limit:  paramLimit,
		Page:   paramPage,
	}
}
