package repository

import "github.com/z9905080/hr_service/internal/domain/entity"

type EmployeeRepository interface {
	AddEmployee(employee entity.Employee) (entity.Employee, error)
	QueryEmployee(employee entity.Employee) (entity.Employee, error)
	UpdateEmployee(data entity.EmployeeUpdate) (entity.Employee, error)
	DeleteEmployee(data entity.Employee) error
	ListEmployee(limit int, page int, field string, order string) ([]entity.Employee, error)
}
