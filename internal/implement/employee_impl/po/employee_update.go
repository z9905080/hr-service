package po

import (
	"github.com/z9905080/hr_service/internal/domain/entity"
	"gorm.io/gorm"
	"time"
)

type EmployeeUpdatePo struct {
	gorm.Model
	Name  *string    `gorm:"column:name; size:255; not null"`
	Birth *time.Time `gorm:"column:birth; not null"`
	Role  *string    `gorm:"column:role; size:255; not null"`
	Email *string    `gorm:"column:email; size:255"`
}

func (*EmployeeUpdatePo) TableName() string {
	return "employee"
}

func NewEmployeeUpdatePo(e entity.EmployeeUpdate) *EmployeeUpdatePo {
	return &EmployeeUpdatePo{
		Model: gorm.Model{
			ID: uint(e.ID),
		},
		Name:  e.Name,
		Birth: e.Birth,
		Role:  e.Role,
		Email: e.Email,
	}
}

func (e *EmployeeUpdatePo) ToUpdateMap() map[string]any {
	m := make(map[string]any)
	if e.Name != nil {
		m["name"] = *e.Name
	}
	if e.Birth != nil {
		m["birth"] = *e.Birth
	}
	if e.Role != nil {
		m["role"] = *e.Role
	}
	if e.Email != nil {
		m["email"] = *e.Email
	}

	if len(m) > 0 {
		m["updated_at"] = time.Now()
	}

	return m
}

func (e *EmployeeUpdatePo) ToEntity() entity.EmployeeUpdate {
	return entity.EmployeeUpdate{
		ID:    int(e.ID),
		Name:  e.Name,
		Birth: e.Birth,
		Role:  e.Role,
		Email: e.Email,
	}
}
