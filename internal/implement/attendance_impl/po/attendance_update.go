package po

import (
	"github.com/z9905080/hr_service/internal/domain/entity"
	gormTableValidator "github.com/z9905080/hr_service/pkg/gorm_table_validator"
	"gorm.io/gorm"
	"time"
)

type AttendanceUpdatePo struct {
	gorm.Model
	EmployeeID *int       `gorm:"column:employee_id"`
	StartTime  *time.Time `gorm:"column:start_time"`
	EndTime    *time.Time `gorm:"column:end_time"`
}

func (d *AttendanceUpdatePo) TableName() string {
	return "attendance"
}

func (d *AttendanceUpdatePo) ToEntity() entity.AttendanceUpdate {
	return entity.AttendanceUpdate{
		ID:              int(d.ID),
		EmployeeID:      d.EmployeeID,
		AttendanceStart: d.StartTime,
		AttendanceEnd:   d.EndTime,
	}
}

func (d *AttendanceUpdatePo) ToUpdateMap() map[string]interface{} {
	m := make(map[string]interface{})
	if d.EmployeeID != nil {
		m["employee_id"] = *d.EmployeeID
	}

	if d.StartTime != nil {
		m["start_time"] = *d.StartTime
	}

	if d.EndTime != nil {
		m["end_time"] = *d.EndTime
	}

	if len(m) > 0 {
		m["updated_at"] = d.UpdatedAt
	}

	return m
}

var _ gormTableValidator.InfUpdatePo[entity.AttendanceUpdate] = (*AttendanceUpdatePo)(nil)

func NewAttendanceUpdatePo(e entity.AttendanceUpdate) AttendanceUpdatePo {
	return AttendanceUpdatePo{
		EmployeeID: e.EmployeeID,
		StartTime:  e.AttendanceStart,
		EndTime:    e.AttendanceEnd,
	}
}
