package adapter

import (
	"github.com/gin-gonic/gin"
	"github.com/z9905080/hr_service/internal/interface/handler"
	"github.com/z9905080/hr_service/internal/interface/middleware"
	"net/http"
)

type HTTPAdapter struct {
	App     http.Handler
	handler *handler.SrvHandler
}

func NewHTTPAdapter(h *handler.SrvHandler) http.Handler {

	e := gin.Default()

	hs := &HTTPAdapter{
		handler: h,
		App:     e,
	}

	hs.SetRouter()

	return hs.App
}

func (h *HTTPAdapter) SetRouter() {

	e := h.App.(*gin.Engine)
	apiGroup := e.Group("/api")
	apiGroupV1 := apiGroup.Group("/v1")
	apiGroupV1.POST("/employee", middleware.HandleFunc(h.handler.AddEmployee))
	apiGroupV1.GET("/employee/:id", middleware.HandleFunc(h.handler.QueryEmployee))
	apiGroupV1.PUT("/employee/:id", middleware.HandleFunc(h.handler.UpdateEmployee))
	apiGroupV1.DELETE("/employee/:id", middleware.HandleFunc(h.handler.DeleteEmployee))

}
