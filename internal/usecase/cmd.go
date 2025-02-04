package usecase

import (
	"errors"
	"github.com/z9905080/hr_service/internal/domain/enum"
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

type CmdAttendanceAdd struct {
	EmployeeID      int
	AttendanceStart string
	AttendanceEnd   *string
	AttendanceType  enum.AttendanceType
}

func (c *CmdAttendanceAdd) Validate() error {
	if c.EmployeeID <= 0 {
		return errors.New("employee id is invalid")
	}

	// validate attendance date
	// convert string to time
	startTime, err := time.Parse(time.RFC3339, c.AttendanceStart)
	if err != nil {
		return errors.New("attendance start is invalid")
	}

	if startTime.IsZero() {
		return errors.New("attendance start is required")
	}

	// validate attendance end date if not nil
	if c.AttendanceEnd != nil {
		// convert string to time
		endTime, err := time.Parse(time.RFC3339, *c.AttendanceEnd)
		if err != nil {
			return errors.New("attendance end is invalid")
		}

		if endTime.IsZero() {
			return errors.New("attendance end is required if not nil")
		}

		// check if end date is before start date
		if endTime.Before(startTime) {
			return errors.New("attendance end is before start date")
		}
	}

	// validate attendance type
	if c.AttendanceType == "" {
		return errors.New("attendance type is required")
	}

	// validate attendance type
	if !c.AttendanceType.IsValid() {
		return errors.New("attendance type is invalid")
	}

	return nil
}

type CmdAttendanceQuery struct {
	AttendanceID int
}

func (c *CmdAttendanceQuery) Validate() error {
	if c.AttendanceID <= 0 {
		return errors.New("attendance id is invalid")
	}

	return nil
}

type CmdAttendanceUpdate struct {
	AttendanceID    int
	EmployeeID      *int // 可能只有 admin 才能改非自己的考勤
	AttendanceStart *string
	AttendanceEnd   *string
	AttendanceType  *enum.AttendanceType
}

func (c *CmdAttendanceUpdate) Validate() error {
	if c.AttendanceID <= 0 {
		return errors.New("attendance id is invalid")
	}

	if c.EmployeeID != nil && *c.EmployeeID <= 0 {
		return errors.New("employee id is invalid")
	}

	if c.AttendanceStart != nil {
		// convert string to time
		startTime, err := time.Parse(time.RFC3339, *c.AttendanceStart)
		if err != nil {
			return errors.New("attendance start is invalid")
		}

		if startTime.IsZero() {
			return errors.New("attendance start is required if not nil")
		}
	}

	if c.AttendanceEnd != nil {
		// convert string to time
		endTime, err := time.Parse(time.RFC3339, *c.AttendanceEnd)
		if err != nil {
			return errors.New("attendance end is invalid")
		}

		if endTime.IsZero() {
			return errors.New("attendance end is required if not nil")
		}

		// check if end date is before start date
		if c.AttendanceStart != nil {
			startTime, err := time.Parse(time.RFC3339, *c.AttendanceStart)
			if err != nil {
				return errors.New("attendance start is invalid")
			}

			if endTime.Before(startTime) {
				return errors.New("attendance end is before start date if not nil")
			}
		}
	}

	if c.AttendanceType != nil && !c.AttendanceType.IsValid() {
		return errors.New("attendance type is invalid")
	}

	// at least one field to update
	if c.EmployeeID == nil && c.AttendanceStart == nil && c.AttendanceEnd == nil && c.AttendanceType == nil {
		return errors.New("at least one field to update")
	}

	return nil
}

type CmdAttendanceDelete struct {
	AttendanceID int
}

func (c *CmdAttendanceDelete) Validate() error {
	if c.AttendanceID <= 0 {
		return errors.New("attendance id is invalid")
	}

	return nil
}

type CmdAttendanceList struct {
	EmployeeID *int // 可能只有 admin 才能查看非自己的考勤
	Pagination Pagination
	Sorter     Sorter
}

func (c *CmdAttendanceList) Validate() error {

	if c.EmployeeID != nil && *c.EmployeeID <= 0 {
		return errors.New("employee id is invalid")
	}

	if c.Pagination.Page < 0 {
		return errors.New("page number is invalid")
	}

	if c.Pagination.Limit < 0 {
		return errors.New("limit number is invalid")
	}

	// TODO: validate sorter field and order
	return nil
}

type CmdOvertimeAdd struct {
	EmployeeID     int
	OvertimeStart  time.Time
	OvertimeEnd    time.Time
	OvertimeReason string
}

func (c *CmdOvertimeAdd) Validate() error {
	if c.EmployeeID <= 0 {
		return errors.New("employee id is invalid")
	}

	if c.OvertimeStart.IsZero() {
		return errors.New("overtime start is required")
	}

	if c.OvertimeEnd.IsZero() {
		return errors.New("overtime end is required")
	}

	if c.OvertimeReason == "" {
		return errors.New("overtime reason is required")
	}

	return nil
}

type CmdOvertimeQuery struct {
	OvertimeID int
}

func (c *CmdOvertimeQuery) Validate() error {
	if c.OvertimeID <= 0 {
		return errors.New("overtime id is invalid")
	}

	return nil
}

type CmdOvertimeUpdate struct {
	OvertimeID     int
	EmployeeID     *int
	OvertimeStart  *time.Time
	OvertimeEnd    *time.Time
	OvertimeReason *string
}

func (c *CmdOvertimeUpdate) Validate() error {
	if c.OvertimeID <= 0 {
		return errors.New("overtime id is invalid")
	}

	if c.EmployeeID != nil && *c.EmployeeID <= 0 {
		return errors.New("employee id is invalid")
	}

	if c.OvertimeStart != nil && c.OvertimeStart.IsZero() {
		return errors.New("overtime start is required")
	}

	if c.OvertimeEnd != nil && c.OvertimeEnd.IsZero() {
		return errors.New("overtime end is required")
	}

	if c.OvertimeReason != nil && *c.OvertimeReason == "" {
		return errors.New("overtime reason is required")
	}

	// at least one field to update
	if c.EmployeeID == nil && c.OvertimeStart == nil && c.OvertimeEnd == nil && c.OvertimeReason == nil {
		return errors.New("at least one field to update")
	}

	return nil
}

type CmdOvertimeDelete struct {
	OvertimeID int
}

func (c *CmdOvertimeDelete) Validate() error {
	if c.OvertimeID <= 0 {
		return errors.New("overtime id is invalid")
	}

	return nil
}

type CmdOvertimeList struct {
	EmployeeID *int
	Pagination Pagination
	Sorter     Sorter
}

func (c *CmdOvertimeList) Validate() error {
	if c.EmployeeID != nil && *c.EmployeeID <= 0 {
		return errors.New("employee id is invalid")
	}

	if c.Pagination.Page < 0 {
		return errors.New("page number is invalid")
	}

	if c.Pagination.Limit < 0 {
		return errors.New("limit number is invalid")
	}

	return nil
}

type CmdLeaveAdd struct {
	EmployeeID int
	LeaveStart time.Time
	LeaveEnd   time.Time
	LeaveType  enum.LeaveType
}

func (c *CmdLeaveAdd) Validate() error {
	if c.EmployeeID <= 0 {
		return errors.New("employee id is invalid")
	}

	if c.LeaveStart.IsZero() {
		return errors.New("leave start is required")
	}

	if c.LeaveEnd.IsZero() {
		return errors.New("leave end is required")
	}

	if c.LeaveType == "" {
		return errors.New("leave type is required")
	}

	if !c.LeaveType.IsValid() {
		return errors.New("leave type is invalid")
	}

	return nil
}

type CmdLeaveQuery struct {
	LeaveID int
}

func (c *CmdLeaveQuery) Validate() error {
	if c.LeaveID <= 0 {
		return errors.New("leave id is invalid")
	}

	return nil
}

type CmdLeaveUpdate struct {
	LeaveID    int
	EmployeeID *int
	LeaveStart *time.Time
	LeaveEnd   *time.Time
	LeaveType  *enum.LeaveType
}

func (c *CmdLeaveUpdate) Validate() error {
	if c.LeaveID <= 0 {
		return errors.New("leave id is invalid")
	}

	if c.EmployeeID != nil && *c.EmployeeID <= 0 {
		return errors.New("employee id is invalid")
	}

	if c.LeaveStart != nil && c.LeaveStart.IsZero() {
		return errors.New("leave start is required")
	}

	if c.LeaveEnd != nil && c.LeaveEnd.IsZero() {
		return errors.New("leave end is required")
	}

	if c.LeaveType != nil && c.LeaveType.IsValid() {
		return errors.New("leave type is invalid")
	}

	// at least one field to update
	if c.EmployeeID == nil && c.LeaveStart == nil && c.LeaveEnd == nil && c.LeaveType == nil {
		return errors.New("at least one field to update")
	}

	return nil
}

type CmdLeaveDelete struct {
	LeaveID int
}

func (c *CmdLeaveDelete) Validate() error {
	if c.LeaveID <= 0 {
		return errors.New("leave id is invalid")
	}

	return nil
}

type CmdLeaveList struct {
	EmployeeID *int // 可能只有 admin 才能查看非自己的请假
	Pagination Pagination
	Sorter     Sorter
}

func (c *CmdLeaveList) Validate() error {
	if c.EmployeeID != nil && *c.EmployeeID <= 0 {
		return errors.New("employee id is invalid")
	}

	if c.Pagination.Page < 0 {
		return errors.New("page number is invalid")
	}

	if c.Pagination.Limit < 0 {
		return errors.New("limit number is invalid")
	}

	return nil
}

type CmdAttendanceStatistics struct {
	EmployeeIDList []int // if null, return all employees' statistics group by employee_id
	TimeStart      time.Time
	TimeEnd        time.Time
}

func (c *CmdAttendanceStatistics) Validate() error {
	if len(c.EmployeeIDList) > 0 {
		for _, id := range c.EmployeeIDList {
			if id <= 0 {
				return errors.New("employee id is invalid")
			}
		}
	}

	if c.TimeStart.IsZero() {
		return errors.New("time start is required")
	}

	if c.TimeEnd.IsZero() {
		return errors.New("time end is required")
	}

	return nil
}
