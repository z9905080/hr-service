package usecase

import (
	"context"
	"github.com/z9905080/hr_service/internal/domain/entity"
	"github.com/z9905080/hr_service/internal/domain/repository"
	"github.com/z9905080/hr_service/pkg/logger"
	"time"
)

type usecase struct {
	employeeRepo   repository.EmployeeRepository
	departmentRepo repository.DepartmentRepository
	log            logger.InfLogger
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
			EmployeeBirth: v.Birth.Format("2006-01-02"),
			EmployeeRole:  v.Role,
			EmployeeEmail: v.Email,
		})
	}

	return EventEmployeeListed{
		EmployeeList: list,
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
		Role:  nil,
		Email: nil,
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
			EmployeeBirth: data.Birth.Format("2006-01-02"),
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
			EmployeeBirth: data.Birth.Format("2006-01-02"),
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
		Birth: func() time.Time {
			// ignore error because it's already validated
			t, _ := time.Parse("2006-01-02", cmd.EmployeeBirth)
			return t
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
			EmployeeBirth: data.Birth.Format("2006-01-02"),
			EmployeeRole:  data.Role,
			EmployeeEmail: data.Email,
		}}, err
}

func NewUsecase(employeeRepo repository.EmployeeRepository) InfAPIUsecase {
	return &usecase{
		employeeRepo: employeeRepo,
	}
}
