package po

import (
	"github.com/z9905080/hr_service/internal/domain/entity"
	gormTableValidator "github.com/z9905080/hr_service/pkg/gorm_table_validator"
	"gorm.io/gorm"
	"time"
)

type LeaveUpdatePo struct {
	gorm.Model
	EmployeeID *int       `gorm:"column:employee_id; not null; index; foreignkey:employee_id; references:employee(id)"`
	StartTime  *time.Time `gorm:"column:start_time; not null"`
	EndTime    *time.Time `gorm:"column:end_time; not null"`
}

func (d *LeaveUpdatePo) ToUpdateMap() map[string]interface{} {
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
		m["updated_at"] = time.Now()
	}

	return m
}

func (d *LeaveUpdatePo) TableName() string {
	return "overtime"
}

func (d *LeaveUpdatePo) ToEntity() entity.LeaveUpdate {
	return entity.LeaveUpdate{
		ID:         int(d.ID),
		EmployeeID: d.EmployeeID,
		LeaveStart: d.StartTime,
		LeaveEnd:   d.EndTime,
	}
}

var _ gormTableValidator.InfUpdatePo[entity.LeaveUpdate] = (*LeaveUpdatePo)(nil)

func NewLeaveUpdatePo(e entity.LeaveUpdate) LeaveUpdatePo {
	return LeaveUpdatePo{
		EmployeeID: e.EmployeeID,
		StartTime:  e.LeaveStart,
		EndTime:    e.LeaveEnd,
	}
}
