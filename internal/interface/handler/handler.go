package handler

import (
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
	if err := c.BindJSON(&req); err != nil {
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
