package entity

import "time"

type Attendance struct {
	ID              int        `json:"id"`
	EmployeeID      int        `json:"employee_id"`
	EmployeeInfo    Employee   `json:"employee_info"`
	AttendanceStart time.Time  `json:"attendance_start"`
	AttendanceEnd   *time.Time `json:"attendance_end"` // pointer to handle null value
}

type AttendanceUpdate struct {
	ID              int        `json:"id"`
	EmployeeID      *int       `json:"employee_id"`
	AttendanceStart *time.Time `json:"attendance_start"`
	AttendanceEnd   *time.Time `json:"attendance_end"`
}

type AttendanceEmployeeStatistics struct {
	EmployeeID      int      `json:"employee_id"`
	EmployeeInfo    Employee `json:"employee_info"`
	TotalAttendance int      `json:"total_attendance"`
	TotalLeave      int      `json:"total_leave"`
	TotalOvertime   int      `json:"total_overtime"`
}

type AttendanceStatistics struct {
	AttendanceEmployeeStatistics map[int]AttendanceEmployeeStatistics `json:"attendance_employee_statistics"`
}
