package usecase

import (
	"errors"
	"time"
)

type CmdEmployeeAdd struct {
	EmployeeName  string
	EmployeeBirth string
	EmployeeRole  string
	EmployeeEmail string
}

func (c *CmdEmployeeAdd) Validate() error {
	if c.EmployeeName == "" {
		return errors.New("employee name is required")
	}

	if c.EmployeeBirth == "" {
		return errors.New("employee birth is required")
	}

	// validate birth date
	_, err := time.Parse("2006-01-02", c.EmployeeBirth)
	if err != nil {
		return errors.New("employee birth is invalid")
	}

	if c.EmployeeRole == "" {
		return errors.New("employee role is required")
	}

	if c.EmployeeEmail == "" {
		return errors.New("employee email is required")
	}

	return nil
}

type CmdEmployeeQuery struct {
	EmployeeID int
}

func (c *CmdEmployeeQuery) Validate() error {
	if c.EmployeeID <= 0 {
		return errors.New("employee id is invalid")
	}

	return nil
}

type CmdEmployeeUpdate struct {
	EmployeeID     int
	EmployeeName   *string
	EmployeeBirth  *time.Time
	EmployeeRole   *string
	EmployeeEmail  *string
	EmployeeStatus *string
}

func (c *CmdEmployeeUpdate) Validate() error {
	// validate employee id
	if c.EmployeeID <= 0 {
		return errors.New("employee id is invalid")
	}

	// need at least one field to update
	if c.EmployeeName == nil && c.EmployeeBirth == nil && c.EmployeeRole == nil && c.EmployeeEmail == nil {
		return errors.New("at least one field to update")
	}

	// validate birthdate
	if c.EmployeeBirth != nil {
		_, err := time.Parse("2006-01-02", c.EmployeeBirth.Format("2006-01-02"))
		if err != nil {
			return errors.New("employee birth is invalid")
		}
	}

	return nil
}

type CmdEmployeeDelete struct {
	EmployeeID int
}

func (c *CmdEmployeeDelete) Validate() error {
	if c.EmployeeID <= 0 {
		return errors.New("employee id is invalid")
	}

	return nil
}

// Pagination is used to limit the number of data returned
type Pagination struct {
	Page  int
	Limit int
}

type Sorter struct {
	Field string
	Order string
}

type CmdEmployeeList struct {
	Pagination Pagination
	Sorter     Sorter
}

func (c *CmdEmployeeList) Validate() error {
	if c.Pagination.Page < 0 {
		return errors.New("page number is invalid")
	}

	if c.Pagination.Limit < 0 {
		return errors.New("limit number is invalid")
	}

	//TODO: validate sorter field and order

	return nil
}

type CmdDepartmentAdd struct {
	DepartmentName string
}

func (c *CmdDepartmentAdd) Validate() error {
	if c.DepartmentName == "" {
		return errors.New("department name is required")
	}

	return nil
}

type CmdDepartmentQuery struct {
	DepartmentID int
}

func (c *CmdDepartmentQuery) Validate() error {
	if c.DepartmentID <= 0 {
		return errors.New("department id is invalid")
	}

	return nil
}

type CmdDepartmentUpdate struct {
	DepartmentID   int
	DepartmentName *string
}

func (c *CmdDepartmentUpdate) Validate() error {
	if c.DepartmentID <= 0 {
		return errors.New("department id is invalid")
	}

	// NIT: because just one field to update, we can remove this validation
	if c.DepartmentName == nil {
		return errors.New("at least one field to update")
	}

	return nil
}

type CmdDepartmentDelete struct {
	DepartmentID int
}

func (c *CmdDepartmentDelete) Validate() error {
	if c.DepartmentID <= 0 {
		return errors.New("department id is invalid")
	}

	return nil
}

type CmdDepartmentList struct {
	Pagination Pagination
	Sorter     Sorter
}

func (c *CmdDepartmentList) Validate() error {
	if c.Pagination.Page < 0 {
		return errors.New("page number is invalid")
	}

	if c.Pagination.Limit < 0 {
		return errors.New("limit number is invalid")
	}

	// TODO: validate sorter field and order
	return nil
}
