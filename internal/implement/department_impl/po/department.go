package po

import (
	"github.com/z9905080/hr_service/internal/domain/entity"
	gormTableValidator "github.com/z9905080/hr_service/pkg/gorm_table_validator"
	"gorm.io/gorm"
)

type DepartmentPo struct {
	gorm.Model
	Name string `gorm:"column:name; size:255; not null"`
}

func (d *DepartmentPo) TableName() string {
	return "department"
}

func (d *DepartmentPo) ToEntity() entity.Department {
	return entity.Department{
		ID:   int(d.ID),
		Name: d.Name,
	}
}

var _ gormTableValidator.InfPo[entity.Department] = (*DepartmentPo)(nil)

func NewDepartmentPo(e entity.Department) DepartmentPo {
	return DepartmentPo{
		Name: e.Name,
	}
}
