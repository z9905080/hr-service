package usecase

type validator interface {
	Validate() error
}

var _ validator = (*CmdEmployeeAdd)(nil)
var _ validator = (*CmdEmployeeQuery)(nil)
var _ validator = (*CmdEmployeeUpdate)(nil)
var _ validator = (*CmdEmployeeDelete)(nil)
