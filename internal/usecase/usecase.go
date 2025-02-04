package usecase

import (
	"context"
	"errors"
	"github.com/z9905080/hr_service/internal/domain/entity"
	"github.com/z9905080/hr_service/internal/domain/repository"
	"github.com/z9905080/hr_service/pkg/logger"
	"time"
)

type usecase struct {
	employeeRepo   repository.EmployeeRepository
	departmentRepo repository.DepartmentRepository
	attendanceRepo repository.AttendanceRepository
	log            logger.InfLogger
}

func (u *usecase) AttendanceAdd(ctx context.Context, cmd CmdAttendanceAdd) (EventAttendanceAdded, error) {
	if err := cmd.Validate(); err != nil {
		// record error log
		u.log.ErrorF(ctx, "AttendanceAdd: %v", err)
		return EventAttendanceAdded{}, err
	}

	start := func() time.Time {
		// ignore error because it's already validated
		t, _ := time.Parse(time.RFC3339, cmd.AttendanceStart)
		return t
	}()
	e := entity.Attendance{
		EmployeeID:      cmd.EmployeeID,
		AttendanceStart: start,
		AttendanceEnd: func() *time.Time {
			if cmd.AttendanceEnd == nil {
				return nil
			}

			t, _ := time.Parse(time.RFC3339, *cmd.AttendanceEnd)
			return &t

		}(),
	}

	// get the latest attendance
	// if the latest attendance in start, then return error
	id := cmd.EmployeeID
	attendanceList, err := u.attendanceRepo.QueryAttendanceByEmployeeID(id, start)
	if err != nil {
		return EventAttendanceAdded{}, err
	}

	if len(attendanceList) > 0 {
		return EventAttendanceAdded{}, errors.New("the date has attendance")
	}

	result, err := u.attendanceRepo.AddAttendance(e)
	if err != nil {
		// record error log
		u.log.ErrorF(ctx, "AttendanceAdd: %v", err)
		return EventAttendanceAdded{}, err
	}

	return EventAttendanceAdded{
		EventAttendance: EventAttendance{
			AttendanceID:  result.ID,
			EmployeeInfo:  DtoEmployeeEntityToEvent(result.EmployeeInfo),
			AttendanceIn:  result.AttendanceStart,
			AttendanceOut: result.AttendanceEnd,
		},
	}, nil
}

func (u *usecase) AttendanceQuery(ctx context.Context, cmd CmdAttendanceQuery) (EventAttendanceQueried, error) {
	if err := cmd.Validate(); err != nil {
		// record error log
		u.log.ErrorF(ctx, "AttendanceQuery: %v", err)
		return EventAttendanceQueried{}, err
	}

	data, err := u.attendanceRepo.QueryAttendanceByID(entity.Attendance{ID: cmd.AttendanceID})
	if err != nil {
		// record error log
		u.log.ErrorF(ctx, "AttendanceQuery: %v", err)
		return EventAttendanceQueried{}, err
	}

	return EventAttendanceQueried{
		EventAttendance: EventAttendance{
			AttendanceID:  data.ID,
			EmployeeInfo:  DtoEmployeeEntityToEvent(data.EmployeeInfo),
			AttendanceIn:  data.AttendanceStart,
			AttendanceOut: data.AttendanceEnd,
		},
	}, nil
}

func (u *usecase) AttendanceUpdate(ctx context.Context, cmd CmdAttendanceUpdate) (EventAttendanceUpdated, error) {
	if err := cmd.Validate(); err != nil {
		// record error log
		u.log.ErrorF(ctx, "AttendanceUpdate: %v", err)
		return EventAttendanceUpdated{}, err
	}

	attendanceData := entity.AttendanceUpdate{
		ID:         cmd.AttendanceID,
		EmployeeID: cmd.EmployeeID,
		AttendanceStart: func() *time.Time {
			if cmd.AttendanceStart == nil {
				return nil
			}

			// ignore error because it's already validated
			t, _ := time.Parse(time.RFC3339, *cmd.AttendanceStart)
			return &t
		}(),
		AttendanceEnd: func() *time.Time {
			if cmd.AttendanceEnd == nil {
				return nil
			}

			// ignore error because it's already validated
			t, _ := time.Parse(time.RFC3339, *cmd.AttendanceEnd)
			return &t
		}(),
	}

	data, err := u.attendanceRepo.UpdateAttendance(attendanceData)
	if err != nil {
		// record error log
		u.log.ErrorF(ctx, "AttendanceUpdate: %v", err)
		return EventAttendanceUpdated{}, err
	}

	return EventAttendanceUpdated{
		EventAttendance: EventAttendance{
			AttendanceID:  data.ID,
			EmployeeInfo:  DtoEmployeeEntityToEvent(data.EmployeeInfo),
			AttendanceIn:  data.AttendanceStart,
			AttendanceOut: data.AttendanceEnd,
		},
	}, nil

}

func (u *usecase) AttendanceDelete(ctx context.Context, cmd CmdAttendanceDelete) (EventAttendanceDeleted, error) {
	if err := cmd.Validate(); err != nil {
		// record error log
		u.log.ErrorF(ctx, "AttendanceDelete: %v", err)
		return EventAttendanceDeleted{}, err
	}

	err := u.attendanceRepo.DeleteAttendance(entity.Attendance{ID: cmd.AttendanceID})
	if err != nil {
		// record error log
		u.log.ErrorF(ctx, "AttendanceDelete: %v", err)
		return EventAttendanceDeleted{}, err
	}

	return EventAttendanceDeleted{
		AttendanceID: cmd.AttendanceID,
		Status:       "deleted",
	}, nil
}

func (u *usecase) AttendanceList(ctx context.Context, cmd CmdAttendanceList) (EventAttendanceListed, error) {
	if err := cmd.Validate(); err != nil {
		// record error log
		u.log.ErrorF(ctx, "AttendanceList: %v", err)
		return EventAttendanceListed{}, err
	}

	data, err := u.attendanceRepo.ListAttendance(cmd.EmployeeID, cmd.Pagination.Limit, cmd.Pagination.Page, cmd.Sorter.Field, cmd.Sorter.Order)
	if err != nil {
		// record error log
		u.log.ErrorF(ctx, "AttendanceList: %v", err)
		return EventAttendanceListed{}, err
	}

	var list []EventAttendance
	for _, v := range data {
		list = append(list, EventAttendance{
			AttendanceID:  v.ID,
			EmployeeInfo:  DtoEmployeeEntityToEvent(v.EmployeeInfo),
			AttendanceIn:  v.AttendanceStart,
			AttendanceOut: v.AttendanceEnd,
		})
	}

	return EventAttendanceListed{
		AttendanceList: list,
		Pagination: EventPagination{
			Page:      cmd.Pagination.Page,
			Limit:     cmd.Pagination.Limit,
			Total:     0, // TODO: implement total count
			TotalPage: 0, // TODO: implement total page
		},
	}, nil
}

func (u *usecase) OvertimeAdd(ctx context.Context, cmd CmdOvertimeAdd) (EventOvertimeAdded, error) {
	if err := cmd.Validate(); err != nil {
		// record error log
		u.log.ErrorF(ctx, "OvertimeAdd: %v", err)
		return EventOvertimeAdded{}, err
	}

	e := entity.Overtime{
		EmployeeID:    cmd.EmployeeID,
		OvertimeStart: cmd.OvertimeStart,
		OvertimeEnd:   cmd.OvertimeEnd,
	}

	result, err := u.attendanceRepo.AddOvertime(e)
	if err != nil {
		// record error log
		u.log.ErrorF(ctx, "OvertimeAdd: %v", err)
		return EventOvertimeAdded{}, err
	}

	return EventOvertimeAdded{
		EventOvertime: EventOvertime{
			OvertimeID:    result.ID,
			EmployeeInfo:  DtoEmployeeEntityToEvent(result.EmployeeInfo),
			OvertimeStart: result.OvertimeStart,
			OvertimeEnd:   result.OvertimeEnd,
		},
	}, nil
}

func (u *usecase) OvertimeQuery(ctx context.Context, cmd CmdOvertimeQuery) (EventOvertimeQueried, error) {
	if err := cmd.Validate(); err != nil {
		// record error log
		u.log.ErrorF(ctx, "OvertimeQuery: %v", err)
		return EventOvertimeQueried{}, err
	}

	data, err := u.attendanceRepo.QueryOvertime(entity.Overtime{ID: cmd.OvertimeID})
	if err != nil {
		// record error log
		u.log.ErrorF(ctx, "OvertimeQuery: %v", err)
		return EventOvertimeQueried{}, err
	}

	return EventOvertimeQueried{
		EventOvertime: EventOvertime{
			OvertimeID:    data.ID,
			EmployeeInfo:  DtoEmployeeEntityToEvent(data.EmployeeInfo),
			OvertimeStart: data.OvertimeStart,
			OvertimeEnd:   data.OvertimeEnd,
		},
	}, nil
}

func (u *usecase) OvertimeUpdate(ctx context.Context, cmd CmdOvertimeUpdate) (EventOvertimeUpdated, error) {
	if err := cmd.Validate(); err != nil {
		// record error log
		u.log.ErrorF(ctx, "OvertimeUpdate: %v", err)
		return EventOvertimeUpdated{}, err
	}

	overtimeData := entity.OvertimeUpdate{
		ID:            cmd.OvertimeID,
		EmployeeID:    cmd.EmployeeID,
		OvertimeStart: cmd.OvertimeStart,
		OvertimeEnd:   cmd.OvertimeEnd,
	}

	data, err := u.attendanceRepo.UpdateOvertime(overtimeData)
	if err != nil {
		// record error log
		u.log.ErrorF(ctx, "OvertimeUpdate: %v", err)
		return EventOvertimeUpdated{}, err
	}

	return EventOvertimeUpdated{
		EventOvertime: EventOvertime{
			OvertimeID:    data.ID,
			EmployeeInfo:  DtoEmployeeEntityToEvent(data.EmployeeInfo),
			OvertimeStart: data.OvertimeStart,
			OvertimeEnd:   data.OvertimeEnd,
		},
	}, nil
}

func (u *usecase) OvertimeDelete(ctx context.Context, cmd CmdOvertimeDelete) (EventOvertimeDeleted, error) {
	if err := cmd.Validate(); err != nil {
		// record error log
		u.log.ErrorF(ctx, "OvertimeDelete: %v", err)
		return EventOvertimeDeleted{}, err
	}

	err := u.attendanceRepo.DeleteOvertime(entity.Overtime{ID: cmd.OvertimeID})
	if err != nil {
		// record error log
		u.log.ErrorF(ctx, "OvertimeDelete: %v", err)
		return EventOvertimeDeleted{}, err
	}

	return EventOvertimeDeleted{
		OvertimeID: cmd.OvertimeID,
		Status:     "deleted",
	}, nil
}

func (u *usecase) OvertimeList(ctx context.Context, cmd CmdOvertimeList) (EventOvertimeListed, error) {
	if err := cmd.Validate(); err != nil {
		// record error log
		u.log.ErrorF(ctx, "OvertimeList: %v", err)
		return EventOvertimeListed{}, err
	}

	data, err := u.attendanceRepo.ListOvertime(cmd.EmployeeID, cmd.Pagination.Limit, cmd.Pagination.Page, cmd.Sorter.Field, cmd.Sorter.Order)
	if err != nil {
		// record error log
		u.log.ErrorF(ctx, "OvertimeList: %v", err)
		return EventOvertimeListed{}, err
	}

	var list []EventOvertime
	for _, v := range data {
		list = append(list, EventOvertime{
			OvertimeID:    v.ID,
			EmployeeInfo:  DtoEmployeeEntityToEvent(v.EmployeeInfo),
			OvertimeStart: v.OvertimeStart,
			OvertimeEnd:   v.OvertimeEnd,
		})
	}

	return EventOvertimeListed{
		OvertimeList: list,
		Pagination: EventPagination{
			Page:      cmd.Pagination.Page,
			Limit:     cmd.Pagination.Limit,
			Total:     0, // TODO: implement total count
			TotalPage: 0, // TODO: implement total page
		},
	}, nil
}

func (u *usecase) LeaveAdd(ctx context.Context, cmd CmdLeaveAdd) (EventLeaveAdded, error) {
	if err := cmd.Validate(); err != nil {
		// record error log
		u.log.ErrorF(ctx, "LeaveAdd: %v", err)
		return EventLeaveAdded{}, err
	}

	e := entity.Leave{
		EmployeeID: cmd.EmployeeID,
		LeaveStart: cmd.LeaveStart,
		LeaveEnd:   cmd.LeaveEnd,
	}

	result, err := u.attendanceRepo.AddLeave(e)
	if err != nil {
		// record error log
		u.log.ErrorF(ctx, "LeaveAdd: %v", err)
		return EventLeaveAdded{}, err
	}

	return EventLeaveAdded{
		EventLeave: EventLeave{
			LeaveID:      result.ID,
			EmployeeInfo: DtoEmployeeEntityToEvent(result.EmployeeInfo),
			LeaveStart:   result.LeaveStart,
			LeaveEnd:     result.LeaveEnd,
		},
	}, nil
}

func (u *usecase) LeaveQuery(ctx context.Context, cmd CmdLeaveQuery) (EventLeaveQueried, error) {
	if err := cmd.Validate(); err != nil {
		// record error log
		u.log.ErrorF(ctx, "LeaveQuery: %v", err)
		return EventLeaveQueried{}, err
	}

	data, err := u.attendanceRepo.QueryLeave(entity.Leave{ID: cmd.LeaveID})
	if err != nil {
		// record error log
		u.log.ErrorF(ctx, "LeaveQuery: %v", err)
		return EventLeaveQueried{}, err
	}

	return EventLeaveQueried{
		EventLeave: EventLeave{
			LeaveID:      data.ID,
			EmployeeInfo: DtoEmployeeEntityToEvent(data.EmployeeInfo),
			LeaveStart:   data.LeaveStart,
			LeaveEnd:     data.LeaveEnd,
		},
	}, nil
}

func (u *usecase) LeaveUpdate(ctx context.Context, cmd CmdLeaveUpdate) (EventLeaveUpdated, error) {
	if err := cmd.Validate(); err != nil {
		// record error log
		u.log.ErrorF(ctx, "LeaveUpdate: %v", err)
		return EventLeaveUpdated{}, err
	}

	leaveData := entity.LeaveUpdate{
		ID:         cmd.LeaveID,
		EmployeeID: cmd.EmployeeID,
		LeaveStart: cmd.LeaveStart,
		LeaveEnd:   cmd.LeaveEnd,
	}

	data, err := u.attendanceRepo.UpdateLeave(leaveData)
	if err != nil {
		// record error log
		u.log.ErrorF(ctx, "LeaveUpdate: %v", err)
		return EventLeaveUpdated{}, err
	}

	return EventLeaveUpdated{
		EventLeave: EventLeave{
			LeaveID:      data.ID,
			EmployeeInfo: DtoEmployeeEntityToEvent(data.EmployeeInfo),
			LeaveStart:   data.LeaveStart,
			LeaveEnd:     data.LeaveEnd,
		},
	}, nil
}

func (u *usecase) LeaveDelete(ctx context.Context, cmd CmdLeaveDelete) (EventLeaveDeleted, error) {
	if err := cmd.Validate(); err != nil {
		// record error log
		u.log.ErrorF(ctx, "LeaveDelete: %v", err)
		return EventLeaveDeleted{}, err
	}

	err := u.attendanceRepo.DeleteLeave(entity.Leave{ID: cmd.LeaveID})
	if err != nil {
		// record error log
		u.log.ErrorF(ctx, "LeaveDelete: %v", err)
		return EventLeaveDeleted{}, err
	}

	return EventLeaveDeleted{
		LeaveID: cmd.LeaveID,
		Status:  "deleted",
	}, nil
}

func (u *usecase) LeaveList(ctx context.Context, cmd CmdLeaveList) (EventLeaveListed, error) {
	if err := cmd.Validate(); err != nil {
		// record error log
		u.log.ErrorF(ctx, "LeaveList: %v", err)
		return EventLeaveListed{}, err
	}

	data, err := u.attendanceRepo.ListLeave(cmd.EmployeeID, cmd.Pagination.Limit, cmd.Pagination.Page, cmd.Sorter.Field, cmd.Sorter.Order)
	if err != nil {
		// record error log
		u.log.ErrorF(ctx, "LeaveList: %v", err)
		return EventLeaveListed{}, err
	}

	var list []EventLeave
	for _, v := range data {
		list = append(list, EventLeave{
			LeaveID:      v.ID,
			EmployeeInfo: DtoEmployeeEntityToEvent(v.EmployeeInfo),
			LeaveStart:   v.LeaveStart,
			LeaveEnd:     v.LeaveEnd,
		})
	}

	return EventLeaveListed{
		LeaveList: list,
		Pagination: EventPagination{
			Page:      cmd.Pagination.Page,
			Limit:     cmd.Pagination.Limit,
			Total:     0, // TODO: implement total count
			TotalPage: 0, // TODO: implement total page
		},
	}, nil
}

func (u *usecase) AttendanceStatistics(ctx context.Context, cmd CmdAttendanceStatistics) (EventAttendanceStatistics, error) {
	if err := cmd.Validate(); err != nil {
		// record error log
		u.log.ErrorF(ctx, "AttendanceStatistics: %v", err)
		return EventAttendanceStatistics{}, err
	}

	attendanceStatistic, err := u.attendanceRepo.AttendanceStatistic(cmd.EmployeeIDList, cmd.TimeStart, cmd.TimeEnd)
	if err != nil {
		return EventAttendanceStatistics{}, err
	}

	result := make(map[int]EventEmployeeAttendanceStatistics)
	for k, v := range attendanceStatistic.AttendanceEmployeeStatistics {
		result[k] = EventEmployeeAttendanceStatistics{
			TotalAttendance: v.TotalAttendance,
			TotalOvertime:   v.TotalOvertime,
			TotalLeave:      v.TotalLeave,
		}
	}

	return EventAttendanceStatistics{
		EmployeeAttendanceStatistics: result,
	}, nil
}

func (u *usecase) DepartmentAdd(ctx context.Context, department CmdDepartmentAdd) (EventDepartmentAdded, error) {
	if err := department.Validate(); err != nil {
		// record error log
		u.log.ErrorF(ctx, "DepartmentAdd: %v", err)
		return EventDepartmentAdded{}, err
	}

	e := entity.Department{
		Name: department.DepartmentName,
	}

	result, err := u.departmentRepo.AddDepartment(e)
	if err != nil {
		// record error log
		u.log.ErrorF(ctx, "DepartmentAdd: %v", err)
	}

	return EventDepartmentAdded{
		EventDepartment: EventDepartment{
			DepartmentID:   result.ID,
			DepartmentName: result.Name,
		},
	}, nil
}

func (u *usecase) DepartmentQuery(ctx context.Context, departmentID CmdDepartmentQuery) (EventDepartmentQueried, error) {
	if err := departmentID.Validate(); err != nil {
		// record error log
		u.log.ErrorF(ctx, "DepartmentQuery: %v", err)
		return EventDepartmentQueried{}, err
	}

	data, err := u.departmentRepo.QueryDepartment(entity.Department{ID: departmentID.DepartmentID})
	if err != nil {
		// record error log
		u.log.ErrorF(ctx, "DepartmentQuery: %v", err)
	}

	return EventDepartmentQueried{
		EventDepartment: EventDepartment{
			DepartmentID:   data.ID,
			DepartmentName: data.Name,
		},
	}, nil

}

func (u *usecase) DepartmentUpdate(ctx context.Context, department CmdDepartmentUpdate) (EventDepartmentUpdated, error) {
	if err := department.Validate(); err != nil {
		// record error log
		u.log.ErrorF(ctx, "DepartmentUpdate: %v", err)
		return EventDepartmentUpdated{}, err
	}

	data, err := u.departmentRepo.UpdateDepartment(entity.DepartmentUpdate{
		ID:   department.DepartmentID,
		Name: department.DepartmentName,
	})
	if err != nil {
		// record error log
		u.log.ErrorF(ctx, "DepartmentUpdate: %v", err)
	}

	return EventDepartmentUpdated{
		EventDepartment: EventDepartment{
			DepartmentID:   data.ID,
			DepartmentName: data.Name,
		},
	}, nil
}

func (u *usecase) DepartmentDelete(ctx context.Context, departmentID CmdDepartmentDelete) (EventDepartmentDeleted, error) {
	if err := departmentID.Validate(); err != nil {
		// record error log
		u.log.ErrorF(ctx, "DepartmentDelete: %v", err)
		return EventDepartmentDeleted{}, err
	}

	err := u.departmentRepo.DeleteDepartment(entity.Department{ID: departmentID.DepartmentID})
	if err != nil {
		// record error log
		u.log.ErrorF(ctx, "DepartmentDelete: %v", err)
	}

	return EventDepartmentDeleted{
		DepartmentID: departmentID.DepartmentID,
		Status:       "deleted",
	}, nil
}

func (u *usecase) DepartmentList(ctx context.Context, departmentList CmdDepartmentList) (EventDepartmentListed, error) {
	if err := departmentList.Validate(); err != nil {
		// record error log
		u.log.ErrorF(ctx, "DepartmentList: %v", err)
		return EventDepartmentListed{}, err
	}

	data, err := u.departmentRepo.ListDepartment(departmentList.Pagination.Limit, departmentList.Pagination.Page, departmentList.Sorter.Field, departmentList.Sorter.Order)
	if err != nil {
		// record error log
		u.log.ErrorF(ctx, "DepartmentList: %v", err)
	}

	var list []EventDepartment
	for _, v := range data {
		list = append(list, EventDepartment{
			DepartmentID:   v.ID,
			DepartmentName: v.Name,
		})
	}

	return EventDepartmentListed{
		DepartmentList: list,
		Pagination: EventPagination{
			Page:      departmentList.Pagination.Page,
			Limit:     departmentList.Pagination.Limit,
			Total:     0, // TODO: implement total count
			TotalPage: 0, // TODO: implement total page
		},
	}, err
}

func (u *usecase) EmployeeList(ctx context.Context, employeeList CmdEmployeeList) (EventEmployeeListed, error) {
	if err := employeeList.Validate(); err != nil {
		// record error log
		u.log.ErrorF(ctx, "EmployeeList: %v", err)
		return EventEmployeeListed{}, err
	}

	data, err := u.employeeRepo.ListEmployee(employeeList.Pagination.Limit, employeeList.Pagination.Page, employeeList.Sorter.Field, employeeList.Sorter.Order)
	if err != nil {
		// record error log
		u.log.ErrorF(ctx, "EmployeeList: %v", err)
	}

	var list []EventEmployee
	for _, v := range data {
		list = append(list, EventEmployee{
			EmployeeID:    v.ID,
			EmployeeName:  v.Name,
			EmployeeBirth: v.Birth.String(),
			EmployeeRole:  v.Role,
			EmployeeEmail: v.Email,
		})
	}

	return EventEmployeeListed{
		EmployeeList: list,
		Pagination: EventPagination{
			Page:      employeeList.Pagination.Page,
			Limit:     employeeList.Pagination.Limit,
			Total:     0, // TODO: implement total count
			TotalPage: 0, // TODO: implement total page
		},
	}, err
}

func (u *usecase) EmployeeDelete(ctx context.Context, cmd CmdEmployeeDelete) (EventEmployeeDeleted, error) {
	if err := cmd.Validate(); err != nil {
		// record error log
		u.log.ErrorF(ctx, "EmployeeDelete: %v", err)
		return EventEmployeeDeleted{}, err
	}

	err := u.employeeRepo.DeleteEmployee(entity.Employee{ID: cmd.EmployeeID})
	if err != nil {
		// record error log
		u.log.ErrorF(ctx, "EmployeeDelete: %v", err)
		return EventEmployeeDeleted{}, err
	}

	return EventEmployeeDeleted{
		EmployeeID: cmd.EmployeeID,
		Status:     "deleted",
	}, nil
}

func (u *usecase) EmployeeUpdate(ctx context.Context, cmd CmdEmployeeUpdate) (EventEmployeeUpdated, error) {
	if err := cmd.Validate(); err != nil {
		// record error log
		u.log.ErrorF(ctx, "EmployeeUpdate: %v", err)
		return EventEmployeeUpdated{}, err
	}

	employeeData := entity.EmployeeUpdate{
		ID:   cmd.EmployeeID,
		Name: cmd.EmployeeName,
		Birth: func() *time.Time {
			if cmd.EmployeeBirth == nil {
				return nil
			}

			t := *cmd.EmployeeBirth
			return &t

		}(),
		Role:  cmd.EmployeeRole,
		Email: cmd.EmployeeEmail,
	}

	data, err := u.employeeRepo.UpdateEmployee(employeeData)
	if err != nil {
		// record error log
		u.log.ErrorF(ctx, "EmployeeUpdate: %v", err)
	}

	return EventEmployeeUpdated{
		EventEmployee: EventEmployee{
			EmployeeID:    data.ID,
			EmployeeName:  data.Name,
			EmployeeBirth: data.Birth.String(),
			EmployeeRole:  data.Role,
			EmployeeEmail: data.Email,
		},
	}, err
}

func (u *usecase) EmployeeQuery(ctx context.Context, employeeID CmdEmployeeQuery) (EventEmployeeQueried, error) {
	if err := employeeID.Validate(); err != nil {
		// record error log
		u.log.ErrorF(ctx, "EmployeeQuery: %v", err)
		return EventEmployeeQueried{}, err
	}

	data, err := u.employeeRepo.QueryEmployee(entity.Employee{ID: employeeID.EmployeeID})
	if err != nil {
		// record error log
		u.log.ErrorF(ctx, "EmployeeQuery: %v", err)
	}

	return EventEmployeeQueried{
		EventEmployee: EventEmployee{
			EmployeeID:    data.ID,
			EmployeeName:  data.Name,
			EmployeeBirth: data.Birth.String(),
			EmployeeRole:  data.Role,
			EmployeeEmail: data.Email,
		},
	}, err
}

// EmployeeAdd is an usecase to add employee
func (u *usecase) EmployeeAdd(ctx context.Context, cmd CmdEmployeeAdd) (EventEmployeeAdded, error) {
	if err := cmd.Validate(); err != nil {
		// record error log
		u.log.ErrorF(ctx, "EmployeeAdd: %v", err)
		return EventEmployeeAdded{}, err
	}

	employee := entity.Employee{
		Name: cmd.EmployeeName,
		Birth: func() entity.DateType {
			// ignore error because it's already validated
			t, _ := time.Parse("2006-01-02", cmd.EmployeeBirth)
			return entity.DateType(t)
		}(),
		Role:  cmd.EmployeeRole,
		Email: cmd.EmployeeEmail,
	}

	data, err := u.employeeRepo.AddEmployee(employee)
	if err != nil {
		// record error log
		u.log.ErrorF(ctx, "EmployeeAdd: %v", err)
	}

	return EventEmployeeAdded{
		EventEmployee: EventEmployee{
			EmployeeID:    data.ID,
			EmployeeName:  data.Name,
			EmployeeBirth: data.Birth.String(),
			EmployeeRole:  data.Role,
			EmployeeEmail: data.Email,
		}}, err
}

func NewUsecase(employeeRepo repository.EmployeeRepository, departmentRepo repository.DepartmentRepository, attendanceRepo repository.AttendanceRepository, log logger.InfLogger) InfAPIUsecase {
	return &usecase{
		employeeRepo:   employeeRepo,
		departmentRepo: departmentRepo,
		attendanceRepo: attendanceRepo,
		log:            log,
	}
}
