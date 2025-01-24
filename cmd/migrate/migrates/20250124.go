package migrates

import (
	gormigrate "github.com/go-gormigrate/gormigrate/v2"
	attendancePo "github.com/z9905080/hr_service/internal/implement/attendance_impl/po"
	departmentPo "github.com/z9905080/hr_service/internal/implement/department_impl/po"
	employeePo "github.com/z9905080/hr_service/internal/implement/employee_impl/po"
	"gorm.io/gorm"
)

var Ver20250124_INIT = &gormigrate.Migration{
	ID: "20250124",
	Migrate: func(db *gorm.DB) error {
		// Write your migration script here

		// create table employee
		employeeModel := employeePo.EmployeePo{}
		if !db.Migrator().HasTable(employeeModel) {
			return db.AutoMigrate(&employeeModel)
		}

		// create table department
		departmentModel := departmentPo.DepartmentPo{}
		if !db.Migrator().HasTable(departmentModel) {
			return db.AutoMigrate(&departmentModel)
		}

		// create table attendance
		attendanceModel := attendancePo.AttendancePo{}
		if !db.Migrator().HasTable(attendanceModel) {
			return db.AutoMigrate(&attendanceModel)
		}

		// create table overtime
		overtimeModel := attendancePo.OvertimePo{}
		if !db.Migrator().HasTable(overtimeModel) {
			return db.AutoMigrate(&overtimeModel)
		}

		// create table leave
		leaveModel := attendancePo.LeavePo{}
		if !db.Migrator().HasTable(leaveModel) {
			return db.AutoMigrate(&leaveModel)
		}

		return nil

	},
	Rollback: func(db *gorm.DB) error {
		// Write your migration rollback script here

		// drop table employee
		employeeModel := employeePo.EmployeePo{}
		if db.Migrator().HasTable(employeeModel) {
			return db.Migrator().DropTable(employeeModel)
		}

		// drop table department
		departmentModel := departmentPo.DepartmentPo{}
		if db.Migrator().HasTable(departmentModel) {
			return db.Migrator().DropTable(departmentModel)
		}

		// drop table attendance
		attendanceModel := attendancePo.AttendancePo{}
		if db.Migrator().HasTable(attendanceModel) {
			return db.Migrator().DropTable(attendanceModel)
		}

		// drop table overtime
		overtimeModel := attendancePo.OvertimePo{}
		if db.Migrator().HasTable(overtimeModel) {
			return db.Migrator().DropTable(overtimeModel)
		}

		// drop table leave
		leaveModel := attendancePo.LeavePo{}
		if db.Migrator().HasTable(leaveModel) {
			return db.Migrator().DropTable(leaveModel)
		}

		return nil
	},
}
