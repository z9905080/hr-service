package handler

import (
	"github.com/z9905080/hr_service/internal/domain/enum"
	"github.com/z9905080/hr_service/internal/interface/middleware"
	"github.com/z9905080/hr_service/internal/usecase"
	"time"
)

type SrvHandler struct {
	apiUsecase usecase.InfAPIUsecase
}

func NewGinHandler(apiUsecase usecase.InfAPIUsecase) *SrvHandler {
	return &SrvHandler{
		apiUsecase: apiUsecase,
	}
}

// gen:swagger
// AddEmployee add employee
// @Summary add employee
// @Description add employee
// @Tags employee
// @Accept json
// @Produce json
// @Param req body EmployeeAddReq true "employee add request"
// @Success 200 {object} EmployeeAddResp
// @Router /employee [post]
func (h *SrvHandler) AddEmployee(c *middleware.C) {
	var req EmployeeAddReq
	if err := c.BindJSON(&req); err != nil {
		c.FormatRespError(400, err.Error())
		return
	}

	data, err := h.apiUsecase.EmployeeAdd(c.GetCtx().Request.Context(), usecase.CmdEmployeeAdd{
		EmployeeName:  req.EmployeeName,
		EmployeeBirth: req.EmployeeBirth,
		EmployeeRole:  req.EmployeeRole,
		EmployeeEmail: req.EmployeeEmail,
	})

	if err != nil {
		c.FormatRespError(500, err.Error())
		return
	}

	c.FormatRespOK(data)
}

func (h *SrvHandler) QueryEmployee(c *middleware.C) {
	var req EmployeeQueryReq
	if err := c.BindUri(&req); err != nil {
		c.FormatRespError(400, err.Error())
		return
	}

	data, err := h.apiUsecase.EmployeeQuery(c.GetCtx().Request.Context(), usecase.CmdEmployeeQuery{
		EmployeeID: req.EmployeeID,
	})

	if err != nil {
		c.FormatRespError(500, err.Error())
		return
	}

	c.FormatRespOK(data)
}

func (h *SrvHandler) UpdateEmployee(c *middleware.C) {
	var req EmployeeUpdateReq
	if err := c.BindJSON(&req); err != nil {
		c.FormatRespError(400, err.Error())
		return
	}

	data, err := h.apiUsecase.EmployeeUpdate(c.GetCtx().Request.Context(), usecase.CmdEmployeeUpdate{
		EmployeeID:   req.EmployeeID,
		EmployeeName: req.EmployeeName,
		EmployeeBirth: func() *time.Time {
			if req.EmployeeBirth == nil {
				return nil
			}

			t, err := time.Parse("2006-01-02", *req.EmployeeBirth)
			if err != nil {
				return nil
			}
			return &t
		}(),
		EmployeeRole:  req.EmployeeRole,
		EmployeeEmail: req.EmployeeEmail,
	})

	if err != nil {
		c.FormatRespError(500, err.Error())
		return
	}

	c.FormatRespOK(data)
}

func (h *SrvHandler) DeleteEmployee(c *middleware.C) {
	var req EmployeeQueryReq
	if err := c.BindUri(&req); err != nil {
		c.FormatRespError(400, err.Error())
		return
	}

	result, err := h.apiUsecase.EmployeeDelete(c.GetCtx().Request.Context(), usecase.CmdEmployeeDelete{
		EmployeeID: req.EmployeeID,
	})

	if err != nil {
		c.FormatRespError(500, err.Error())
		return
	}

	c.FormatRespOK(result)
}

func (h *SrvHandler) ListEmployee(c *middleware.C) {
	var req EmployeeListReq
	if err := c.BindQuery(&req); err != nil {
		c.FormatRespError(400, err.Error())
		return
	}

	// default value
	if req.Page == 0 {
		req.Page = 1
	}

	if req.Limit == 0 {
		req.Limit = 10
	}

	if req.Field == "" {
		req.Field = "id"
	}

	if req.Order == "" {
		req.Order = "asc"
	}

	data, err := h.apiUsecase.EmployeeList(c.GetCtx().Request.Context(), usecase.CmdEmployeeList{
		Pagination: usecase.Pagination{
			Page:  req.Page,
			Limit: req.Limit,
		},
		Sorter: usecase.Sorter{
			Field: req.Field,
			Order: req.Order,
		},
	})

	if err != nil {
		c.FormatRespError(500, err.Error())
		return
	}

	c.FormatRespOK(data)
}

func (h *SrvHandler) AddDepartment(c *middleware.C) {
	var req DepartmentAddReq
	if err := c.BindJSON(&req); err != nil {
		c.FormatRespError(400, err.Error())
		return
	}

	data, err := h.apiUsecase.DepartmentAdd(c.GetCtx().Request.Context(), usecase.CmdDepartmentAdd{
		DepartmentName: req.DepartmentName,
	})

	if err != nil {
		c.FormatRespError(500, err.Error())
		return
	}

	c.FormatRespOK(data)
}

func (h *SrvHandler) QueryDepartment(c *middleware.C) {
	var req DepartmentQueryReq
	if err := c.BindUri(&req); err != nil {
		c.FormatRespError(400, err.Error())
		return
	}

	data, err := h.apiUsecase.DepartmentQuery(c.GetCtx().Request.Context(), usecase.CmdDepartmentQuery{
		DepartmentID: req.DepartmentID,
	})

	if err != nil {
		c.FormatRespError(500, err.Error())
		return
	}

	c.FormatRespOK(data)
}

func (h *SrvHandler) UpdateDepartment(c *middleware.C) {
	var req DepartmentUpdateReq
	if err := c.BindJSON(&req); err != nil {
		c.FormatRespError(400, err.Error())
		return
	}

	data, err := h.apiUsecase.DepartmentUpdate(c.GetCtx().Request.Context(), usecase.CmdDepartmentUpdate{
		DepartmentID:   req.DepartmentID,
		DepartmentName: req.DepartmentName,
	})

	if err != nil {
		c.FormatRespError(500, err.Error())
		return
	}

	c.FormatRespOK(data)
}

func (h *SrvHandler) DeleteDepartment(c *middleware.C) {
	var req DepartmentQueryReq
	if err := c.BindUri(&req); err != nil {
		c.FormatRespError(400, err.Error())
		return
	}

	result, err := h.apiUsecase.DepartmentDelete(c.GetCtx().Request.Context(), usecase.CmdDepartmentDelete{
		DepartmentID: req.DepartmentID,
	})

	if err != nil {
		c.FormatRespError(500, err.Error())
		return
	}

	c.FormatRespOK(result)
}

func (h *SrvHandler) ListDepartment(c *middleware.C) {
	var req DepartmentListReq
	if err := c.BindQuery(&req); err != nil {
		c.FormatRespError(400, err.Error())
		return
	}

	// default value
	if req.Page == 0 {
		req.Page = 1
	}

	if req.Limit == 0 {
		req.Limit = 10
	}

	if req.Field == "" {
		req.Field = "id"
	}

	if req.Order == "" {
		req.Order = "asc"
	}

	data, err := h.apiUsecase.DepartmentList(c.GetCtx().Request.Context(), usecase.CmdDepartmentList{
		Pagination: usecase.Pagination{
			Page:  req.Page,
			Limit: req.Limit,
		},
		Sorter: usecase.Sorter{
			Field: req.Field,
			Order: req.Order,
		},
	})

	if err != nil {
		c.FormatRespError(500, err.Error())
		return
	}

	c.FormatRespOK(data)

}

func (h *SrvHandler) AddAttendance(c *middleware.C) {
	var req AttendanceAddReq
	if err := c.BindJSON(&req); err != nil {
		c.FormatRespError(400, err.Error())
		return
	}

	data, err := h.apiUsecase.AttendanceAdd(c.GetCtx().Request.Context(), usecase.CmdAttendanceAdd{
		EmployeeID:      req.EmployeeID,
		AttendanceStart: req.StartTime,
		AttendanceEnd:   req.EndTime,
		AttendanceType:  enum.AttendanceType(req.Type),
	})

	if err != nil {
		c.FormatRespError(500, err.Error())
		return
	}

	c.FormatRespOK(data)
}

func (h *SrvHandler) QueryAttendance(c *middleware.C) {
	var req AttendanceQueryReq
	if err := c.BindUri(&req); err != nil {
		c.FormatRespError(400, err.Error())
		return
	}

	data, err := h.apiUsecase.AttendanceQuery(c.GetCtx().Request.Context(), usecase.CmdAttendanceQuery{
		AttendanceID: req.AttendanceID,
	})

	if err != nil {
		c.FormatRespError(500, err.Error())
		return
	}

	c.FormatRespOK(data)
}

func (h *SrvHandler) UpdateAttendance(c *middleware.C) {
	var req AttendanceUpdateReq
	if err := c.BindJSON(&req); err != nil {
		c.FormatRespError(400, err.Error())
		return
	}

	data, err := h.apiUsecase.AttendanceUpdate(c.GetCtx().Request.Context(), usecase.CmdAttendanceUpdate{
		AttendanceID:    req.AttendanceID,
		AttendanceStart: req.StartTime,
		AttendanceEnd:   req.EndTime,
		AttendanceType: func() *enum.AttendanceType {
			if req.Type == nil {
				return nil
			}
			t := enum.AttendanceType(*req.Type)
			return &t
		}(),
	})

	if err != nil {
		c.FormatRespError(500, err.Error())
		return
	}

	c.FormatRespOK(data)
}

func (h *SrvHandler) DeleteAttendance(c *middleware.C) {
	var req AttendanceQueryReq
	if err := c.BindUri(&req); err != nil {
		c.FormatRespError(400, err.Error())
		return
	}

	result, err := h.apiUsecase.AttendanceDelete(c.GetCtx().Request.Context(), usecase.CmdAttendanceDelete{
		AttendanceID: req.AttendanceID,
	})

	if err != nil {
		c.FormatRespError(500, err.Error())
		return
	}

	c.FormatRespOK(result)
}

func (h *SrvHandler) ListAttendance(c *middleware.C) {
	var req AttendanceListReq
	if err := c.BindQuery(&req); err != nil {
		c.FormatRespError(400, err.Error())
		return
	}

	// default value
	if req.Page == 0 {
		req.Page = 1
	}

	if req.Limit == 0 {
		req.Limit = 10
	}

	if req.Field == "" {
		req.Field = "id"
	}

	if req.Order == "" {
		req.Order = "asc"
	}

	data, err := h.apiUsecase.AttendanceList(c.GetCtx().Request.Context(), usecase.CmdAttendanceList{
		Pagination: usecase.Pagination{
			Page:  req.Page,
			Limit: req.Limit,
		},
		Sorter: usecase.Sorter{
			Field: req.Field,
			Order: req.Order,
		},
	})

	if err != nil {
		c.FormatRespError(500, err.Error())
		return
	}

	c.FormatRespOK(data)
}

func (h *SrvHandler) AddLeave(c *middleware.C) {
	var req LeaveAddReq
	if err := c.BindJSON(&req); err != nil {
		c.FormatRespError(400, err.Error())
		return
	}

	data, err := h.apiUsecase.LeaveAdd(c.GetCtx().Request.Context(), usecase.CmdLeaveAdd{
		EmployeeID: req.EmployeeID,
		LeaveStart: func() time.Time {
			t, _ := time.Parse(time.RFC3339, req.StartTime)
			return t
		}(),
		LeaveEnd: func() time.Time {
			t, _ := time.Parse(time.RFC3339, req.EndTime)
			return t
		}(),
		LeaveType: enum.LeaveType(req.Type),
	})

	if err != nil {
		c.FormatRespError(500, err.Error())
		return
	}

	c.FormatRespOK(data)

}

func (h *SrvHandler) QueryLeave(c *middleware.C) {
	var req LeaveQueryReq
	if err := c.BindUri(&req); err != nil {
		c.FormatRespError(400, err.Error())
		return
	}

	data, err := h.apiUsecase.LeaveQuery(c.GetCtx().Request.Context(), usecase.CmdLeaveQuery{
		LeaveID: req.LeaveID,
	})

	if err != nil {
		c.FormatRespError(500, err.Error())
		return
	}

	c.FormatRespOK(data)

}

func (h *SrvHandler) UpdateLeave(c *middleware.C) {
	var req LeaveUpdateReq
	if err := c.BindJSON(&req); err != nil {
		c.FormatRespError(400, err.Error())
		return
	}

	data, err := h.apiUsecase.LeaveUpdate(c.GetCtx().Request.Context(), usecase.CmdLeaveUpdate{
		LeaveID: req.LeaveID,
		LeaveStart: func() *time.Time {
			if req.StartTime == nil {
				return nil
			}
			t, _ := time.Parse(time.RFC3339, *req.StartTime)
			return &t
		}(),
		LeaveEnd: func() *time.Time {
			if req.EndTime == nil {
				return nil
			}
			t, _ := time.Parse(time.RFC3339, *req.EndTime)
			return &t
		}(),
	})

	if err != nil {
		c.FormatRespError(500, err.Error())
		return
	}

	c.FormatRespOK(data)
}

func (h *SrvHandler) DeleteLeave(c *middleware.C) {
	var req LeaveQueryReq
	if err := c.BindUri(&req); err != nil {
		c.FormatRespError(400, err.Error())
		return
	}

	result, err := h.apiUsecase.LeaveDelete(c.GetCtx().Request.Context(), usecase.CmdLeaveDelete{
		LeaveID: req.LeaveID,
	})

	if err != nil {
		c.FormatRespError(500, err.Error())
		return
	}

	c.FormatRespOK(result)
}

func (h *SrvHandler) ListLeave(c *middleware.C) {
	var req LeaveListReq
	if err := c.BindQuery(&req); err != nil {
		c.FormatRespError(400, err.Error())
		return
	}

	// default value
	if req.Page == 0 {
		req.Page = 1
	}

	if req.Limit == 0 {
		req.Limit = 10
	}

	if req.Field == "" {
		req.Field = "id"
	}

	if req.Order == "" {
		req.Order = "asc"
	}

	data, err := h.apiUsecase.LeaveList(c.GetCtx().Request.Context(), usecase.CmdLeaveList{
		Pagination: usecase.Pagination{
			Page:  req.Page,
			Limit: req.Limit,
		},
		Sorter: usecase.Sorter{
			Field: req.Field,
			Order: req.Order,
		},
	})

	if err != nil {
		c.FormatRespError(500, err.Error())
		return
	}

	c.FormatRespOK(data)
}

func (h *SrvHandler) AttendanceStatistic(c *middleware.C) {
	var req AttendanceStatisticReq
	if err := c.BindJSON(&req); err != nil {
		c.FormatRespError(400, err.Error())
		return
	}

	data, err := h.apiUsecase.AttendanceStatistics(c.GetCtx().Request.Context(), usecase.CmdAttendanceStatistics{
		EmployeeIDList: req.EmployeeIDList,
		TimeStart: func() time.Time {
			t, _ := time.Parse(time.RFC3339, req.Start)
			return t
		}(),
		TimeEnd: func() time.Time {
			t, _ := time.Parse(time.RFC3339, req.End)
			return t
		}(),
	})

	if err != nil {
		c.FormatRespError(500, err.Error())
		return
	}

	c.FormatRespOK(data)
}
