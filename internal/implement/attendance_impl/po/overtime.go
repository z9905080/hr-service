package po

import (
	"github.com/z9905080/hr_service/internal/domain/entity"
	employeePo "github.com/z9905080/hr_service/internal/implement/employee_impl/po"
	gormTableValidator "github.com/z9905080/hr_service/pkg/gorm_table_validator"
	"gorm.io/gorm"
	"time"
)

type OvertimePo struct {
	gorm.Model
	EmployeeID   int                   `gorm:"column:employee_id; not null; index; foreignkey:employee_id; references:employee(id)"`
	StartTime    time.Time             `gorm:"column:start_time; not null"`
	EndTime      time.Time             `gorm:"column:end_time; not null"`
	EmployeeInfo employeePo.EmployeePo `gorm:"foreignKey:EmployeeID; references:ID"`
}

func (d *OvertimePo) TableName() string {
	return "overtime"
}

func (d *OvertimePo) ToEntity() entity.Overtime {
	return entity.Overtime{
		ID:            int(d.ID),
		EmployeeID:    d.EmployeeID,
		EmployeeInfo:  d.EmployeeInfo.ToEntity(),
		OvertimeStart: d.StartTime,
		OvertimeEnd:   d.EndTime,
	}
}

var _ gormTableValidator.InfPo[entity.Overtime] = (*OvertimePo)(nil)

func NewOvertimePo(e entity.Overtime) OvertimePo {
	return OvertimePo{
		EmployeeID: e.EmployeeID,
		StartTime:  e.OvertimeStart,
		EndTime:    e.OvertimeEnd,
	}
}
