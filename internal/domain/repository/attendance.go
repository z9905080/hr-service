package repository

import (
	"github.com/z9905080/hr_service/internal/domain/entity"
	"time"
)

type AttendanceRepository interface {
	AddAttendance(attendance entity.Attendance) (entity.Attendance, error)
	QueryAttendanceByID(attendance entity.Attendance) (entity.Attendance, error)
	QueryAttendanceByEmployeeID(employeeID int, attendanceDate time.Time) ([]entity.Attendance, error)
	UpdateAttendance(data entity.AttendanceUpdate) (entity.Attendance, error)
	DeleteAttendance(data entity.Attendance) error
	ListAttendance(employeeID *int, limit int, page int, field string, order string) ([]entity.Attendance, error)

	AddOvertime(overtime entity.Overtime) (entity.Overtime, error)
	QueryOvertime(overtime entity.Overtime) (entity.Overtime, error)
	UpdateOvertime(data entity.OvertimeUpdate) (entity.Overtime, error)
	DeleteOvertime(data entity.Overtime) error
	ListOvertime(employeeID *int, limit int, page int, field string, order string) ([]entity.Overtime, error)

	AddLeave(leave entity.Leave) (entity.Leave, error)
	QueryLeave(leave entity.Leave) (entity.Leave, error)
	UpdateLeave(data entity.LeaveUpdate) (entity.Leave, error)
	DeleteLeave(data entity.Leave) error
	ListLeave(employeeID *int, limit int, page int, field string, order string) ([]entity.Leave, error)

	AttendanceStatistic(employeeID []int, start time.Time, end time.Time) (entity.AttendanceStatistics, error)
}
