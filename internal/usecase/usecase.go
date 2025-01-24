package usecase

import (
	"context"
	"github.com/z9905080/hr_service/internal/domain/entity"
	"github.com/z9905080/hr_service/internal/domain/repository"
	"github.com/z9905080/hr_service/pkg/logger"
	"time"
)

type usecase struct {
	employeeRepo repository.EmployeeRepository
	log          logger.InfLogger
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
