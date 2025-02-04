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
	apiGroupV1.GET("/employees", middleware.HandleFunc(h.handler.ListEmployee))

	apiGroupV1.POST("/department", middleware.HandleFunc(h.handler.AddDepartment))
	apiGroupV1.GET("/department/:id", middleware.HandleFunc(h.handler.QueryDepartment))
	apiGroupV1.PUT("/department/:id", middleware.HandleFunc(h.handler.UpdateDepartment))
	apiGroupV1.DELETE("/department/:id", middleware.HandleFunc(h.handler.DeleteDepartment))
	apiGroupV1.GET("/departments", middleware.HandleFunc(h.handler.ListDepartment))

	apiGroupV1.POST("/attendance", middleware.HandleFunc(h.handler.AddAttendance))
	apiGroupV1.GET("/attendance/:id", middleware.HandleFunc(h.handler.QueryAttendance))
	apiGroupV1.PUT("/attendance/:id", middleware.HandleFunc(h.handler.UpdateAttendance))
	apiGroupV1.DELETE("/attendance/:id", middleware.HandleFunc(h.handler.DeleteAttendance))
	apiGroupV1.GET("/attendances", middleware.HandleFunc(h.handler.ListAttendance))

	apiGroupV1.POST("/leave", middleware.HandleFunc(h.handler.AddLeave))
	apiGroupV1.GET("/leave/:id", middleware.HandleFunc(h.handler.QueryLeave))
	apiGroupV1.PUT("/leave/:id", middleware.HandleFunc(h.handler.UpdateLeave))
	apiGroupV1.DELETE("/leave/:id", middleware.HandleFunc(h.handler.DeleteLeave))
	apiGroupV1.GET("/leaves", middleware.HandleFunc(h.handler.ListLeave))

	// TODO: add overtime
	//apiGroupV1.POST("/overtime", middleware.HandleFunc(h.handler.AddOvertime))
	//apiGroupV1.GET("/overtime/:id", middleware.HandleFunc(h.handler.QueryOvertime))
	//apiGroupV1.PUT("/overtime/:id", middleware.HandleFunc(h.handler.UpdateOvertime))
	//apiGroupV1.DELETE("/overtime/:id", middleware.HandleFunc(h.handler.DeleteOvertime))
	//apiGroupV1.GET("/overtimes", middleware.HandleFunc(h.handler.ListOvertime))

	apiGroupV1.POST("/attendance-statistic", middleware.HandleFunc(h.handler.AttendanceStatistic))
}
