package entity

import "time"

// Employee is a struct that represent employee entity
type Employee struct {
	ID    int      `json:"id"`
	Name  string   `json:"name"`
	Birth DateType `json:"birth"`
	Role  string   `json:"role"`
	Email string   `json:"email"`
}

type DateType time.Time

func (b DateType) String() string {
	return time.Time(b).Format("2006-01-02")
}

type EmployeeUpdate struct {
	ID    int        `json:"id"`
	Name  *string    `json:"name"`
	Birth *time.Time `json:"birth"`
	Role  *string    `json:"role"`
	Email *string    `json:"email"`
}
