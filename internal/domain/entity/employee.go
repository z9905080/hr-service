package entity

import "time"

// Employee is a struct that represent employee entity
type Employee struct {
	ID    int       `json:"id"`
	Name  string    `json:"name"`
	Birth time.Time `json:"birth"`
	Role  string    `json:"role"`
	Email string    `json:"email"`
}

type EmployeeUpdate struct {
	ID    int        `json:"id"`
	Name  *string    `json:"name"`
	Birth *time.Time `json:"birth"`
	Role  *string    `json:"role"`
	Email *string    `json:"email"`
}
