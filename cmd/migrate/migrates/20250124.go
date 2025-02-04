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
			err := db.AutoMigrate(&employeeModel)
			if err != nil {
				return err
			}
		}

		// create table department
		departmentModel := departmentPo.DepartmentPo{}
		if !db.Migrator().HasTable(departmentModel) {
			err := db.AutoMigrate(&departmentModel)
			if err != nil {
				return err
			}
		}

		// create table attendance
		attendanceModel := attendancePo.AttendancePo{}
		if !db.Migrator().HasTable(attendanceModel) {
			err := db.AutoMigrate(&attendanceModel)
			if err != nil {
				return err
			}
		}

		// create table overtime
		overtimeModel := attendancePo.OvertimePo{}
		if !db.Migrator().HasTable(overtimeModel) {
			err := db.AutoMigrate(&overtimeModel)
			if err != nil {
				return err

			}
		}

		// create table leave
		leaveModel := attendancePo.LeavePo{}
		if !db.Migrator().HasTable(leaveModel) {
			err := db.AutoMigrate(&leaveModel)
			if err != nil {
			}
		}

		return nil

	},
	Rollback: func(db *gorm.DB) error {
		// Write your migration rollback script here

		// drop table employee
		employeeModel := employeePo.EmployeePo{}
		if db.Migrator().HasTable(employeeModel) {
			if err := db.Migrator().DropTable(employeeModel); err != nil {
				return err
			}
		}

		// drop table department
		departmentModel := departmentPo.DepartmentPo{}
		if db.Migrator().HasTable(departmentModel) {
			if err := db.Migrator().DropTable(departmentModel); err != nil {
				return err
			}
		}

		// drop table attendance
		attendanceModel := attendancePo.AttendancePo{}
		if db.Migrator().HasTable(attendanceModel) {
			if err := db.Migrator().DropTable(attendanceModel); err != nil {
				return err
			}
		}

		// drop table overtime
		overtimeModel := attendancePo.OvertimePo{}
		if db.Migrator().HasTable(overtimeModel) {
			if err := db.Migrator().DropTable(overtimeModel); err != nil {
				return err
			}
		}

		// drop table leave
		leaveModel := attendancePo.LeavePo{}
		if db.Migrator().HasTable(leaveModel) {
			if err := db.Migrator().DropTable(leaveModel); err != nil {
				return err
			}
		}

		return nil
	},
}
