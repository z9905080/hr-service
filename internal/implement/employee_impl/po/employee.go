package po

import (
	"github.com/z9905080/hr_service/internal/domain/entity"
	gormTableValidator "github.com/z9905080/hr_service/pkg/gorm_table_validator"
	"gorm.io/gorm"
	"time"
)

type EmployeePo struct {
	gorm.Model
	Name  string    `gorm:"column:name; size:255; not null"`
	Birth time.Time `gorm:"column:birth; not null"`
	Role  string    `gorm:"column:role; size:255; not null"`
	Email string    `gorm:"column:email; size:255"`
}

func (*EmployeePo) TableName() string {
	return "employee"
}

var _ gormTableValidator.InfPo[entity.Employee] = (*EmployeePo)(nil)

func NewEmployeePo(e entity.Employee) *EmployeePo {
	return &EmployeePo{
		Model: gorm.Model{
			ID: uint(e.ID),
		},
		Name:  e.Name,
		Birth: time.Time(e.Birth),
		Role:  e.Role,
		Email: e.Email,
	}
}

func (e *EmployeePo) ToEntity() entity.Employee {
	return entity.Employee{
		ID:    int(e.ID),
		Name:  e.Name,
		Birth: entity.DateType(e.Birth),
		Role:  e.Role,
		Email: e.Email,
	}
}
