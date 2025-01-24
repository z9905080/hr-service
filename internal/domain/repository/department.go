package repository

import "github.com/z9905080/hr_service/internal/domain/entity"

type DepartmentRepository interface {
	AddDepartment(department entity.Department) (entity.Department, error)
	QueryDepartment(department entity.Department) (entity.Department, error)
	UpdateDepartment(data entity.DepartmentUpdate) (entity.Department, error)
	DeleteDepartment(data entity.Department) error
	ListDepartment(limit int, page int, field string, order string) ([]entity.Department, error)
}
