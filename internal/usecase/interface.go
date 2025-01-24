package usecase

import "context"

// Employee Management
// - Add/Query/Update/Delete employee data
// - Employee ID generation and management
// - Department relationship maintenance
// - Position management
// - Employee status tracking (active/resigned/on leave)

// Attendance Management
// - Attendance record registration and query
// - Overtime application and approval
// - Leave application and approval
// - Business trip management
// - Attendance anomaly alerts
// - Attendance statistics reports

// Payroll Management
// - Payroll structure setup
// - Payroll calculation
// - Bonus management
// - Deductions management
// - Payroll slip generation
// - Payroll transfer data export

// Performance Management
// - Assessment item setup
// - Assessment process management
// - Assessment result registration
// - Assessment report generation
// - Performance tracking

// Training & Development
// - Training course management
// - Training registration and approval
// - Training record management
// - Certification management
// - Training effectiveness evaluation

type InfAPIUsecase interface {
	// EmployeeAdd is an usecase to add employee
	// NIT: don't use pointer struct as parameter, use value struct instead
	//      in this case, use CmdEmployeeAdd instead of *CmdEmployeeAdd
	//      why? because it's can easily to avoid memory leak(escape analysis)
	EmployeeAdd(ctx context.Context, employee CmdEmployeeAdd) (EventEmployeeAdded, error)

	// EmployeeQuery is an usecase to query employee
	EmployeeQuery(ctx context.Context, employeeID CmdEmployeeQuery) (EventEmployeeQueried, error)

	// EmployeeUpdate is an usecase to update employee
	EmployeeUpdate(ctx context.Context, employee CmdEmployeeUpdate) (EventEmployeeUpdated, error)

	// EmployeeDelete is an usecase to delete employee
	EmployeeDelete(ctx context.Context, employeeID CmdEmployeeDelete) (EventEmployeeDeleted, error)
}
