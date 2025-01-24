package employee_impl

import (
	"github.com/z9905080/hr_service/internal/domain/entity"
	"github.com/z9905080/hr_service/internal/domain/repository"
	"github.com/z9905080/hr_service/internal/implement/employee_impl/po"
	"gorm.io/gorm"
)

type employeeImpl struct {
	db *gorm.DB
}

func (e *employeeImpl) AddEmployee(employee entity.Employee) (entity.Employee, error) {
	employeePo := po.NewEmployeePo(employee)
	if err := e.db.Create(employeePo).Error; err != nil {
		return entity.Employee{}, err
	}
	return employeePo.ToEntity(), nil
}

func (e *employeeImpl) QueryEmployee(employee entity.Employee) (entity.Employee, error) {
	var employeePo po.EmployeePo
	if err := e.db.Where("id = ?", employee.ID).First(&employeePo).Error; err != nil {
		return entity.Employee{}, err
	}
	return employeePo.ToEntity(), nil
}

func (e *employeeImpl) UpdateEmployee(data entity.EmployeeUpdate) (entity.Employee, error) {
	updatePo := po.NewEmployeeUpdatePo(data)

	if err := e.db.Model(&po.EmployeeUpdatePo{}).Where("id = ?", data.ID).Updates(updatePo.ToUpdateMap()).Error; err != nil {
		return entity.Employee{}, err
	}

	// TODO: [performance] get resp from input data
	// get updated data
	var resp po.EmployeePo
	if err := e.db.Where("id = ?", data.ID).First(&resp).Error; err != nil {
		return entity.Employee{}, err
	}

	return resp.ToEntity(), nil

}

func (e *employeeImpl) DeleteEmployee(data entity.Employee) error {
	if err := e.db.Where("id = ?", data.ID).Delete(&po.EmployeePo{}).Error; err != nil {
		return err
	}
	return nil
}

// NewEmployeeImpl will create an object that represent the EmployeeImpl struct
func NewEmployeeImpl(db *gorm.DB) repository.EmployeeRepository {
	return &employeeImpl{
		db: db,
	}
}
