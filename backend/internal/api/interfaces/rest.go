package interfaces

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/norun9/Hybird/internal/api/domain/model"
	"github.com/norun9/Hybird/pkg/log"
	"go.uber.org/zap"
	"net/http"
	"reflect"
	"regexp"
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

func buildPath(r *http.Request) string {
	re := regexp.MustCompile(`/\d+`)
	path := re.ReplaceAllStringFunc(r.URL.Path, func(match string) string {
		return "/{id}"
	})
	return path
}

func (h *restHandler) Exec(c *gin.Context, params interface{}) {
	var err error
	r := c.Request
	w := c.Writer
	method := Method(r.Method)
	path := buildPath(r)
	route, ok := h.routeMap[Path{path, method}]
	if !ok {
		log.Logger.Error(fmt.Sprintf("Failed to exec. path:%s method:%s", path, method))
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	f := route.Func
	validate := validator.New()
	args := []reflect.Value{
		reflect.ValueOf(c),
	}
	funcType := reflect.TypeOf(f)
	if 1 < funcType.NumIn() {
		// the second element of argument is input
		inputType := funcType.In(1)
		if inputType.Kind() != reflect.Slice {
			if err = validate.Struct(params); err != nil {
				// input struct validation failed
				log.Logger.Error("Validation struct error", zap.Error(err))
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		}
		args = append(args, reflect.Indirect(reflect.ValueOf(params)))
	}

	fv := reflect.ValueOf(f)
	results := fv.Call(args)
	var responseJSON []byte
	switch len(results) {
	case 1:
		errResult := results[0]
		if errResult.Interface() == nil {
			return
		}
		err = errResult.Interface().(error)
		log.Logger.Error("Execution error", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	case 2:
		result, errResult := results[0], results[1]
		if errResult.Interface() != nil {
			err = errResult.Interface().(error)
			log.Logger.Error("Execution error", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// If concrete results are available, convert them to JSON
		if responseJSON, err = json.Marshal(result.Interface()); err != nil {
			log.Logger.Error("Failed to marshal response", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	default:
		log.Logger.Error("Execution error", zap.Error(err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	// response json value to front-end using response writer
	w.Header().Set("Content-Type", "application/json")
	if _, _err := w.Write(responseJSON); _err != nil {
		log.Logger.Error("Failed to write response", zap.Error(_err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
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
