package attendance_impl

import (
	"github.com/z9905080/hr_service/internal/domain/entity"
	"github.com/z9905080/hr_service/internal/domain/repository"
	"github.com/z9905080/hr_service/internal/implement/attendance_impl/po"
	"github.com/z9905080/hr_service/pkg/logger"
	"gorm.io/gorm"
)

type impl struct {
	db     *gorm.DB
	logger logger.InfLogger
}

func (i *impl) AddAttendance(e entity.Attendance) (entity.Attendance, error) {
	model := po.NewAttendancePo(e)
	if err := i.db.Preload("Employee").Create(&model).Error; err != nil {
		return entity.Attendance{}, err
	}

	return model.ToEntity(), nil
}

func (i *impl) QueryAttendance(e entity.Attendance) (entity.Attendance, error) {
	var model po.AttendancePo
	if err := i.db.Where("id = ?", e.ID).Preload("Employee").First(&model).Error; err != nil {
		return entity.Attendance{}, err
	}

	return model.ToEntity(), nil
}

func (i *impl) UpdateAttendance(e entity.AttendanceUpdate) (entity.Attendance, error) {
	model := po.NewAttendanceUpdatePo(e)
	if err := i.db.Model(&po.AttendancePo{}).Where("id = ?", e.ID).Updates(model.ToUpdateMap()).Error; err != nil {
		return entity.Attendance{}, err
	}

	var attendance po.AttendancePo
	if err := i.db.Where("id = ?", e.ID).Preload("Employee").First(&attendance).Error; err != nil {
		return entity.Attendance{}, err
	}

	return attendance.ToEntity(), nil
}

func (i *impl) DeleteAttendance(e entity.Attendance) error {
	if err := i.db.Where("id = ?", e.ID).Delete(&po.AttendancePo{}).Error; err != nil {
		return err
	}

	return nil
}

func (i *impl) ListAttendance(employeeID *int, limit int, page int, field string, order string) ([]entity.Attendance, error) {
	var model []po.AttendancePo

	db := i.db
	if employeeID != nil {
		db = db.Where("employee_id = ?", employeeID)
	}

	if field != "" && order != "" {
		db = db.Order(field + " " + order)
	}

	if limit > 0 {
		db = db.Limit(limit).Offset((page - 1) * limit)
	}

	if err := db.Preload("Employee").Find(&model).Error; err != nil {
		return nil, err
	}

	var (
		result []entity.Attendance
	)
	for _, v := range model {
		result = append(result, v.ToEntity())
	}

	return result, nil
}

func (i *impl) AddOvertime(overtime entity.Overtime) (entity.Overtime, error) {
	createPo := po.NewOvertimePo(overtime)
	if err := i.db.Preload("Employee").Create(&createPo).Error; err != nil {
		return entity.Overtime{}, err
	}

	return createPo.ToEntity(), nil

}

func (i *impl) QueryOvertime(overtime entity.Overtime) (entity.Overtime, error) {
	var model po.OvertimePo
	if err := i.db.Where("id = ?", overtime.ID).Preload("Employee").First(&model).Error; err != nil {
		return entity.Overtime{}, err
	}

	return model.ToEntity(), nil
}

func (i *impl) UpdateOvertime(data entity.OvertimeUpdate) (entity.Overtime, error) {
	model := po.NewOvertimeUpdatePo(data)
	if err := i.db.Model(&po.OvertimePo{}).Where("id = ?", data.ID).Updates(model.ToUpdateMap()).Error; err != nil {
		return entity.Overtime{}, err
	}

	var overtime po.OvertimePo
	if err := i.db.Where("id = ?", data.ID).Preload("Employee").First(&overtime).Error; err != nil {
		return entity.Overtime{}, err
	}

	return overtime.ToEntity(), nil
}

func (i *impl) DeleteOvertime(data entity.Overtime) error {
	if err := i.db.Where("id = ?", data.ID).Delete(&po.OvertimePo{}).Error; err != nil {
		return err
	}

	return nil
}

func (i *impl) ListOvertime(employeeID *int, limit int, page int, field string, order string) ([]entity.Overtime, error) {
	var model []po.OvertimePo

	db := i.db
	if employeeID != nil {
		db = db.Where("employee_id = ?", employeeID)
	}

	if field != "" && order != "" {
		db = db.Order(field + " " + order)
	}

	if limit > 0 {
		db = db.Limit(limit).Offset((page - 1) * limit)
	}

	if err := db.Preload("Employee").Find(&model).Error; err != nil {
		return nil, err
	}

	var (
		result []entity.Overtime
	)
	for _, v := range model {
		result = append(result, v.ToEntity())
	}

	return result, nil
}

func (i *impl) AddLeave(leave entity.Leave) (entity.Leave, error) {
	//TODO implement me
	panic("implement me")
}

func (i *impl) QueryLeave(leave entity.Leave) (entity.Leave, error) {
	//TODO implement me
	panic("implement me")
}

func (i *impl) UpdateLeave(data entity.LeaveUpdate) (entity.Leave, error) {
	//TODO implement me
	panic("implement me")
}

func (i *impl) DeleteLeave(data entity.Leave) error {
	//TODO implement me
	panic("implement me")
}

func (i *impl) ListLeave(employeeID *int, limit int, page int, field string, order string) ([]entity.Leave, error) {
	//TODO implement me
	panic("implement me")
}

func (i *impl) AttendanceStatistic(employeeID []int) (entity.AttendanceStatistics, error) {
	//TODO implement me
	panic("implement me")
}

func NewAttendanceImpl(db *gorm.DB, infLogger logger.InfLogger) repository.AttendanceRepository {
	return &impl{
		db:     db,
		logger: infLogger,
	}
}
