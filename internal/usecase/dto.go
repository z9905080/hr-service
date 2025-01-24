package usecase

import "github.com/z9905080/hr_service/internal/domain/entity"

func DtoEmployeeEntityToEvent(employeeEntity entity.Employee) EventEmployee {
	event := EventEmployee{
		EmployeeID:    employeeEntity.ID,
		EmployeeName:  employeeEntity.Name,
		EmployeeBirth: employeeEntity.Birth.String(),
		EmployeeRole:  employeeEntity.Role,
		EmployeeEmail: employeeEntity.Email,
	}
	return event
}
