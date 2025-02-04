package po

import (
	"github.com/z9905080/hr_service/internal/domain/entity"
	gormTableValidator "github.com/z9905080/hr_service/pkg/gorm_table_validator"
	"gorm.io/gorm"
	"time"
)

type DepartmentUpdatePo struct {
	gorm.Model
	Name *string `gorm:"column:name"`
}

func (d *DepartmentUpdatePo) TableName() string {
	return "department"
}

func (d *DepartmentUpdatePo) ToEntity() entity.DepartmentUpdate {
	return entity.DepartmentUpdate{
		ID:   int(d.ID),
		Name: d.Name,
	}
}

func (d *DepartmentUpdatePo) ToUpdateMap() map[string]interface{} {
	m := make(map[string]interface{})
	if d.Name != nil {
		m["name"] = *d.Name
	}

	if len(m) > 0 {
		m["updated_at"] = time.Now()
	}

	return m
}

var _ gormTableValidator.InfUpdatePo[entity.DepartmentUpdate] = (*DepartmentUpdatePo)(nil)

func NewDepartmentUpdatePo(e entity.DepartmentUpdate) DepartmentUpdatePo {
	return DepartmentUpdatePo{
		Name: e.Name,
	}
}
