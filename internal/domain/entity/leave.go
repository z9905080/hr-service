package entity

import "time"

type Leave struct {
	ID           int       `json:"id"`
	EmployeeID   int       `json:"employee_id"`
	EmployeeInfo Employee  `json:"employee_info"`
	LeaveStart   time.Time `json:"leave_start"`
	LeaveEnd     time.Time `json:"leave_end"`
}

type LeaveUpdate struct {
	ID         int        `json:"id"`
	EmployeeID *int       `json:"employee_id"`
	LeaveStart *time.Time `json:"leave_start"`
	LeaveEnd   *time.Time `json:"leave_end"`
}
