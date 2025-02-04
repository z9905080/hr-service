package handler

type PageReq struct {
	Page  int `form:"page" json:"page"`
	Limit int `form:"limit" json:"limit"`
}

type SorterReq struct {
	Field string `form:"field" json:"field"`
	Order string `form:"order" json:"order"`
}

type EmployeeAddReq struct {
	EmployeeName  string `json:"employee_name"` // employee name
	EmployeeBirth string `json:"employee_birth"`
	EmployeeRole  string `json:"employee_role"`
	EmployeeEmail string `json:"employee_email"`
}

type EmployeeQueryReq struct {
	EmployeeID int `json:"employee_id" uri:"id"`
}

type EmployeeUpdateReq struct {
	EmployeeID    int     `json:"employee_id" uri:"id"`
	EmployeeName  *string `json:"employee_name"`
	EmployeeBirth *string `json:"employee_birth"`
	EmployeeRole  *string `json:"employee_role"`
	EmployeeEmail *string `json:"employee_email"`
}

type EmployeeListReq struct {
	PageReq
	SorterReq
}

type DepartmentAddReq struct {
	DepartmentName string `json:"department_name"`
}

type DepartmentQueryReq struct {
	DepartmentID int `json:"department_id" uri:"id"`
}

type DepartmentUpdateReq struct {
	DepartmentID   int     `json:"department_id" uri:"id"`
	DepartmentName *string `json:"department_name"`
}

type DepartmentListReq struct {
	PageReq
	SorterReq
}

type AttendanceAddReq struct {
	EmployeeID int     `json:"employee_id"`
	StartTime  string  `json:"start_time"`
	EndTime    *string `json:"end_time"`
	Type       string  `json:"type"`
}

type AttendanceQueryReq struct {
	AttendanceID int `json:"attendance_id" uri:"id"`
}

type AttendanceUpdateReq struct {
	AttendanceID int     `json:"attendance_id" uri:"id"`
	StartTime    *string `json:"start_time"`
	EndTime      *string `json:"end_time"`
	Type         *string `json:"type"`
}

type AttendanceListReq struct {
	PageReq
	SorterReq
}

type LeaveAddReq struct {
	EmployeeID int    `json:"employee_id"`
	StartTime  string `json:"start_time"`
	EndTime    string `json:"end_time"`
	Type       string `json:"type"`
}

type LeaveQueryReq struct {
	LeaveID int `json:"leave_id" uri:"id"`
}

type LeaveUpdateReq struct {
	LeaveID   int     `json:"leave_id" uri:"id"`
	StartTime *string `json:"start_time"`
	EndTime   *string `json:"end_time"`
	Type      *string `json:"type"`
}

type LeaveListReq struct {
	PageReq
	SorterReq
}

type AttendanceStatisticReq struct {
	EmployeeIDList []int  `json:"employee_id_list"`
	Start          string `json:"start"`
	End            string `json:"end"`
}
