package entity

import "time"

type Overtime struct {
	ID            int       `json:"id"`
	EmployeeID    int       `json:"employee_id"`
	EmployeeInfo  Employee  `json:"employee_info"`
	OvertimeStart time.Time `json:"overtime_start"`
	OvertimeEnd   time.Time `json:"overtime_end"`
}

type OvertimeUpdate struct {
	ID            int        `json:"id"`
	EmployeeID    *int       `json:"employee_id"`
	OvertimeStart *time.Time `json:"overtime_start"`
	OvertimeEnd   *time.Time `json:"overtime_end"`
}
