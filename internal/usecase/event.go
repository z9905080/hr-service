package usecase

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
