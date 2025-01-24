package usecase

import "context"

// Employee Management
// - Add/Query/Update/Delete employee data
// - Employee ID generation and management
// - Department relationship maintenance
// - Position management

// Attendance Management
// - Attendance record registration and query
// - Overtime application and approval
// - Leave application and approval
// - Attendance statistics reports

// Payroll Management
// - Payroll structure setup
// - Payroll calculation
// - Bonus management
// - Deductions management
// - Payroll slip generation
// - Payroll transfer data export

// TODO: nice to have
// Performance Management
// - Assessment item setup
// - Assessment process management
// - Assessment result registration
// - Assessment report generation
// - Performance tracking

type InfAPIUsecase interface {
	infEmployeeUC
	infDepartmentUC
	infAttendanceUC
}

type infEmployeeUC interface {
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

	// EmployeeList is an usecase to list employee
	EmployeeList(ctx context.Context, employeeList CmdEmployeeList) (EventEmployeeListed, error)
}

type infDepartmentUC interface {
	// DepartmentAdd is an usecase to add department
	DepartmentAdd(ctx context.Context, department CmdDepartmentAdd) (EventDepartmentAdded, error)

	// DepartmentQuery is an usecase to query department
	DepartmentQuery(ctx context.Context, departmentID CmdDepartmentQuery) (EventDepartmentQueried, error)

	// DepartmentUpdate is an usecase to update department
	DepartmentUpdate(ctx context.Context, department CmdDepartmentUpdate) (EventDepartmentUpdated, error)

	// DepartmentDelete is an usecase to delete department
	DepartmentDelete(ctx context.Context, departmentID CmdDepartmentDelete) (EventDepartmentDeleted, error)

	// DepartmentList is an usecase to list department
	DepartmentList(ctx context.Context, departmentList CmdDepartmentList) (EventDepartmentListed, error)
}

type infAttendanceUC interface {
	// AttendanceAdd is an usecase to add attendance
	AttendanceAdd(ctx context.Context, cmd CmdAttendanceAdd) (EventAttendanceAdded, error)

	// AttendanceQuery is an usecase to query attendance
	AttendanceQuery(ctx context.Context, cmd CmdAttendanceQuery) (EventAttendanceQueried, error)

	// AttendanceUpdate is an usecase to update attendance
	AttendanceUpdate(ctx context.Context, cmd CmdAttendanceUpdate) (EventAttendanceUpdated, error)

	// AttendanceDelete is an usecase to delete attendance
	AttendanceDelete(ctx context.Context, cmd CmdAttendanceDelete) (EventAttendanceDeleted, error)

	// AttendanceList is an usecase to list attendance
	AttendanceList(ctx context.Context, cmd CmdAttendanceList) (EventAttendanceListed, error)

	// OvertimeAdd is an usecase to add overtime
	OvertimeAdd(ctx context.Context, cmd CmdOvertimeAdd) (EventOvertimeAdded, error)

	// OvertimeQuery is an usecase to query overtime
	OvertimeQuery(ctx context.Context, cmd CmdOvertimeQuery) (EventOvertimeQueried, error)

	// OvertimeUpdate is an usecase to update overtime
	OvertimeUpdate(ctx context.Context, cmd CmdOvertimeUpdate) (EventOvertimeUpdated, error)

	// OvertimeDelete is an usecase to delete overtime
	OvertimeDelete(ctx context.Context, cmd CmdOvertimeDelete) (EventOvertimeDeleted, error)

	// OvertimeList is an usecase to list overtime
	OvertimeList(ctx context.Context, cmd CmdOvertimeList) (EventOvertimeListed, error)

	// LeaveAdd is an usecase to add leave
	LeaveAdd(ctx context.Context, cmd CmdLeaveAdd) (EventLeaveAdded, error)

	// LeaveQuery is an usecase to query leave
	LeaveQuery(ctx context.Context, cmd CmdLeaveQuery) (EventLeaveQueried, error)

	// LeaveUpdate is an usecase to update leave
	LeaveUpdate(ctx context.Context, cmd CmdLeaveUpdate) (EventLeaveUpdated, error)

	// LeaveDelete is an usecase to delete leave
	LeaveDelete(ctx context.Context, cmd CmdLeaveDelete) (EventLeaveDeleted, error)

	// LeaveList is an usecase to list leave
	LeaveList(ctx context.Context, cmd CmdLeaveList) (EventLeaveListed, error)

	// AttendanceStatistics is an usecase to get attendance statistics
	AttendanceStatistics(ctx context.Context, cmd CmdAttendanceStatistics) (EventAttendanceStatistics, error)
}
