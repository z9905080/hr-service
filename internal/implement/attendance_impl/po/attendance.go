package po

import (
	"github.com/z9905080/hr_service/internal/domain/entity"
	employeePo "github.com/z9905080/hr_service/internal/implement/employee_impl/po"
	gormTableValidator "github.com/z9905080/hr_service/pkg/gorm_table_validator"
	"gorm.io/gorm"
	"time"
)

type AttendancePo struct {
	gorm.Model
	EmployeeID int                   `gorm:"column:employee_id; not null; index; foreignkey:employee_id; references:employee(id)"`
	StartTime  time.Time             `gorm:"column:start_time; not null"`
	EndTime    *time.Time            `gorm:"column:end_time"`
	Employee   employeePo.EmployeePo `gorm:"foreignKey:EmployeeID; references:ID"`
}

func (d *AttendancePo) TableName() string {
	return "attendance"
}

func (d *AttendancePo) ToEntity() entity.Attendance {
	return entity.Attendance{
		ID:              int(d.ID),
		EmployeeID:      d.EmployeeID,
		EmployeeInfo:    d.Employee.ToEntity(),
		AttendanceStart: d.StartTime,
		AttendanceEnd:   d.EndTime,
	}
}

var _ gormTableValidator.InfPo[entity.Attendance] = (*AttendancePo)(nil)

func NewAttendancePo(e entity.Attendance) AttendancePo {
	return AttendancePo{
		EmployeeID: e.EmployeeID,
		StartTime:  e.AttendanceStart,
		EndTime:    e.AttendanceEnd,
	}
}
