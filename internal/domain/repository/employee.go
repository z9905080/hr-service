package repository

import "github.com/z9905080/hr_service/internal/domain/entity"

type EmployeeRepository interface {
	AddEmployee(employee entity.Employee) (entity.Employee, error)
	QueryEmployee(employee entity.Employee) (entity.Employee, error)
	UpdateEmployee(data entity.EmployeeUpdate) (entity.Employee, error)
	DeleteEmployee(data entity.Employee) error
	ListEmployee(limit int, page int, field string, order string) ([]entity.Employee, error)
}

type DepartmentRepository interface {
	AddDepartment(department entity.Department) (entity.Department, error)
	QueryDepartment(department entity.Department) (entity.Department, error)
	UpdateDepartment(data entity.DepartmentUpdate) (entity.Department, error)
	DeleteDepartment(data entity.Department) error
	ListDepartment(limit int, page int, field string, order string) ([]entity.Department, error)
}
