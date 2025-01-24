package usecase

import "time"

type EventPagination struct {
	Page      int `json:"page"`
	Limit     int `json:"limit"`
	Total     int `json:"total"`
	TotalPage int `json:"total_page"`
}

type EventEmployeeAdded struct {
	EventEmployee
}

type EventEmployee struct {
	EmployeeID    int    `json:"employee_id"`
	EmployeeName  string `json:"employee_name"`
	EmployeeBirth string `json:"employee_birth"`
	EmployeeRole  string `json:"employee_role"`
	EmployeeEmail string `json:"employee_email"`
}

type EventEmployeeQueried struct {
	EventEmployee
}

type EventEmployeeUpdated struct {
	EventEmployee
}

type EventEmployeeDeleted struct {
	EmployeeID int    `json:"employee_id"`
	Status     string `json:"status"`
}

type EventEmployeeListed struct {
	EmployeeList []EventEmployee `json:"employee_list"`
	Pagination   EventPagination `json:"pagination"`
}

// EventDepartment is a struct that defines the department event.
type EventDepartment struct {
	DepartmentID   int    `json:"department_id"`
	DepartmentName string `json:"department_name"`
}

type EventDepartmentAdded struct {
	EventDepartment
}

type EventDepartmentQueried struct {
	EventDepartment
}

type EventDepartmentUpdated struct {
	EventDepartment
}

type EventDepartmentDeleted struct {
	DepartmentID int    `json:"department_id"`
	Status       string `json:"status"`
}

type EventDepartmentListed struct {
	DepartmentList []EventDepartment `json:"department_list"`
	Pagination     EventPagination   `json:"pagination"`
}

type EventAttendance struct {
	AttendanceID  int           `json:"attendance_id"`
	EmployeeInfo  EventEmployee `json:"employee_info"`
	AttendanceIn  time.Time     `json:"attendance_in"`
	AttendanceOut *time.Time    `json:"attendance_out"`
}

type EventAttendanceAdded struct {
	EventAttendance
}

type EventAttendanceQueried struct {
	EventAttendance
}

type EventAttendanceUpdated struct {
	EventAttendance
}

type EventAttendanceDeleted struct {
	AttendanceID int `json:"attendance_id"`
	Status       string
}

type EventAttendanceListed struct {
	AttendanceList []EventAttendance `json:"attendance_list"`
	Pagination     EventPagination   `json:"pagination"`
}

type EventOvertime struct {
	OvertimeID    int           `json:"overtime_id"`
	EmployeeInfo  EventEmployee `json:"employee_info"`
	OvertimeStart time.Time     `json:"overtime_start"`
	OvertimeEnd   time.Time     `json:"overtime_end"`
}

type EventOvertimeAdded struct {
	EventOvertime
}

type EventOvertimeQueried struct {
	EventOvertime
}

type EventOvertimeUpdated struct {
	EventOvertime
}

type EventOvertimeDeleted struct {
	OvertimeID int    `json:"overtime_id"`
	Status     string `json:"status"`
}

type EventOvertimeListed struct {
	OvertimeList []EventOvertime `json:"overtime_list"`
	Pagination   EventPagination `json:"pagination"`
}

type EventLeave struct {
	LeaveID      int           `json:"leave_id"`
	EmployeeInfo EventEmployee `json:"employee_info"`
	LeaveStart   time.Time     `json:"leave_start"`
	LeaveEnd     time.Time     `json:"leave_end"`
}

type EventLeaveAdded struct {
	EventLeave
}

type EventLeaveQueried struct {
	EventLeave
}

type EventLeaveUpdated struct {
	EventLeave
}

type EventLeaveDeleted struct {
	LeaveID int    `json:"leave_id"`
	Status  string `json:"status"`
}

type EventLeaveListed struct {
	LeaveList  []EventLeave    `json:"leave_list"`
	Pagination EventPagination `json:"pagination"`
}

type EventEmployeeAttendanceStatistics struct {
	TotalAttendance int `json:"total_attendance"`
	TotalOvertime   int `json:"total_overtime"`
	TotalLeave      int `json:"total_leave"`
}

type EventAttendanceStatistics struct {
	EmployeeAttendanceStatistics map[int]EventEmployeeAttendanceStatistics `json:"employee_attendance_statistics"` // key is employee id
}
