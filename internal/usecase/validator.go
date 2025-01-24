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
