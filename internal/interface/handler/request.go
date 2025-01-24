package handler

type EmployeeAddReq struct {
	EmployeeName  string `json:"employee_name"`
	EmployeeBirth string `json:"employee_birth"`
	EmployeeRole  string `json:"employee_role"`
	EmployeeEmail string `json:"employee_email"`
}

type EmployeeQueryReq struct {
	EmployeeID int `json:"employee_id"`
}

type EmployeeUpdateReq struct {
	EmployeeID    int     `json:"employee_id"`
	EmployeeName  *string `json:"employee_name"`
	EmployeeBirth *string `json:"employee_birth"`
	EmployeeRole  *string `json:"employee_role"`
	EmployeeEmail *string `json:"employee_email"`
}
