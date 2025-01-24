package po

import (
	"github.com/z9905080/hr_service/internal/domain/entity"
	gormTableValidator "github.com/z9905080/hr_service/pkg/gorm_table_validator"
	"gorm.io/gorm"
	"time"
)

type AttendancePo struct {
	gorm.Model
	EmployeeID int        `gorm:"column:employee_id; not null; index"`
	StartTime  time.Time  `gorm:"column:start_time; not null"`
	EndTime    *time.Time `gorm:"column:end_time"`
}

func (d *AttendancePo) TableName() string {
	return "attendance"
}

func (d *AttendancePo) ToEntity() entity.Attendance {
	return entity.Attendance{
		ID:              int(d.ID),
		EmployeeID:      d.EmployeeID,
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
