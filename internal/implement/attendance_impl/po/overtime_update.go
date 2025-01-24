package po

import (
	"github.com/z9905080/hr_service/internal/domain/entity"
	gormTableValidator "github.com/z9905080/hr_service/pkg/gorm_table_validator"
	"gorm.io/gorm"
	"time"
)

type OvertimeUpdatePo struct {
	gorm.Model
	EmployeeID *int       `gorm:"column:employee_id; not null; index; foreignkey:employee_id; references:employee(id)"`
	StartTime  *time.Time `gorm:"column:start_time; not null"`
	EndTime    *time.Time `gorm:"column:end_time; not null"`
}

func (d *OvertimeUpdatePo) ToUpdateMap() map[string]interface{} {
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

func (d *OvertimeUpdatePo) TableName() string {
	return "overtime"
}

func (d *OvertimeUpdatePo) ToEntity() entity.OvertimeUpdate {
	return entity.OvertimeUpdate{
		ID:            int(d.ID),
		EmployeeID:    d.EmployeeID,
		OvertimeStart: d.StartTime,
		OvertimeEnd:   d.EndTime,
	}
}

var _ gormTableValidator.InfUpdatePo[entity.OvertimeUpdate] = (*OvertimeUpdatePo)(nil)

func NewOvertimeUpdatePo(e entity.OvertimeUpdate) OvertimeUpdatePo {
	return OvertimeUpdatePo{
		EmployeeID: e.EmployeeID,
		StartTime:  e.OvertimeStart,
		EndTime:    e.OvertimeEnd,
	}
}
