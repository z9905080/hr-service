package po

import (
	"github.com/z9905080/hr_service/internal/domain/entity"
	employeePo "github.com/z9905080/hr_service/internal/implement/employee_impl/po"
	gormTableValidator "github.com/z9905080/hr_service/pkg/gorm_table_validator"
	"gorm.io/gorm"
	"time"
)

type LeavePo struct {
	gorm.Model
	EmployeeID   int                   `gorm:"column:employee_id; not null; index; foreignkey:employee_id; references:employee(id)"`
	StartTime    time.Time             `gorm:"column:start_time; not null"`
	EndTime      time.Time             `gorm:"column:end_time; not null"`
	EmployeeInfo employeePo.EmployeePo `gorm:"foreignKey:EmployeeID; references:ID"`
}

func (d *LeavePo) TableName() string {
	return "leave"
}

func (d *LeavePo) ToEntity() entity.Leave {
	return entity.Leave{
		ID:           int(d.ID),
		EmployeeID:   d.EmployeeID,
		EmployeeInfo: d.EmployeeInfo.ToEntity(),
		LeaveStart:   d.StartTime,
		LeaveEnd:     d.EndTime,
	}
}

var _ gormTableValidator.InfPo[entity.Leave] = (*LeavePo)(nil)

func NewLeavePo(e entity.Leave) LeavePo {
	return LeavePo{
		EmployeeID: e.EmployeeID,
		StartTime:  e.LeaveStart,
		EndTime:    e.LeaveEnd,
	}
}
