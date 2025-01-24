package usecase

type validator interface {
	Validate() error
}

var _ validator = (*CmdEmployeeAdd)(nil)
var _ validator = (*CmdEmployeeQuery)(nil)
var _ validator = (*CmdEmployeeUpdate)(nil)
var _ validator = (*CmdEmployeeDelete)(nil)
var _ validator = (*CmdEmployeeList)(nil)

var _ validator = (*CmdDepartmentAdd)(nil)
var _ validator = (*CmdDepartmentQuery)(nil)
var _ validator = (*CmdDepartmentUpdate)(nil)
var _ validator = (*CmdDepartmentDelete)(nil)
var _ validator = (*CmdDepartmentList)(nil)

var _ validator = (*CmdAttendanceAdd)(nil)
var _ validator = (*CmdAttendanceQuery)(nil)
var _ validator = (*CmdAttendanceUpdate)(nil)
var _ validator = (*CmdAttendanceDelete)(nil)
var _ validator = (*CmdAttendanceList)(nil)

var _ validator = (*CmdOvertimeAdd)(nil)
var _ validator = (*CmdOvertimeQuery)(nil)
var _ validator = (*CmdOvertimeUpdate)(nil)
var _ validator = (*CmdOvertimeDelete)(nil)
var _ validator = (*CmdOvertimeList)(nil)

var _ validator = (*CmdLeaveAdd)(nil)
var _ validator = (*CmdLeaveQuery)(nil)
var _ validator = (*CmdLeaveUpdate)(nil)
var _ validator = (*CmdLeaveDelete)(nil)
var _ validator = (*CmdLeaveList)(nil)
