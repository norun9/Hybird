package router

import (
	"github.com/gin-gonic/gin"
)

func GetMessageRoutes(r *gin.RouterGroup) {
	r.GET("/health", func(c *gin.Context) {
		//rw := c.Writer
		//r := c.Request
		//ctx := c.Request.Context()
	})

	//	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
	//		ctx := r.Context()
	//		params := input.ChatList{
	//			Paging: GetPagingInfo(r),
	//			IsPC:   GetURLQueryBool(r, "isPC").Bool,
	//		}
	//		h.Exec(ctx, w, r, params)
	//	})
}
