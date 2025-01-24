package usecase

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
