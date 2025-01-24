package attendance_impl

import (
	"github.com/z9905080/hr_service/internal/domain/entity"
	"github.com/z9905080/hr_service/internal/domain/repository"
	"github.com/z9905080/hr_service/internal/implement/attendance_impl/po"
	"github.com/z9905080/hr_service/pkg/logger"
	"gorm.io/gorm"
	"time"
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
	addPo := po.NewLeavePo(leave)
	if err := i.db.Preload("Employee").Create(&addPo).Error; err != nil {
		return entity.Leave{}, err
	}

	return addPo.ToEntity(), nil
}

func (i *impl) QueryLeave(leave entity.Leave) (entity.Leave, error) {
	var model po.LeavePo
	if err := i.db.Where("id = ?", leave.ID).Preload("Employee").First(&model).Error; err != nil {
		return entity.Leave{}, err
	}

	return model.ToEntity(), nil
}

func (i *impl) UpdateLeave(data entity.LeaveUpdate) (entity.Leave, error) {
	model := po.NewLeaveUpdatePo(data)
	if err := i.db.Model(&po.LeavePo{}).Where("id = ?", data.ID).Updates(model.ToUpdateMap()).Error; err != nil {
		return entity.Leave{}, err
	}

	var leave po.LeavePo
	if err := i.db.Where("id = ?", data.ID).Preload("Employee").First(&leave).Error; err != nil {
		return entity.Leave{}, err
	}

	return leave.ToEntity(), nil
}

func (i *impl) DeleteLeave(data entity.Leave) error {
	if err := i.db.Where("id = ?", data.ID).Delete(&po.LeavePo{}).Error; err != nil {
		return err
	}

	return nil
}

func (i *impl) ListLeave(employeeID *int, limit int, page int, field string, order string) ([]entity.Leave, error) {
	var model []po.LeavePo

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
		result []entity.Leave
	)
	for _, v := range model {
		result = append(result, v.ToEntity())
	}

	return result, nil
}

func (i *impl) AttendanceStatistic(employeeID []int, start, end time.Time) (entity.AttendanceStatistics, error) {
	var (
		attendance []po.AttendancePo
		overtime   []po.OvertimePo
		leave      []po.LeavePo
	)

	// TODO: performance issue of all data query
	// better statistic in SQL query

	// query all data between start and end time
	query := i.db.Where("start_time >= ? AND end_time <= ?", start, end)

	// if employeeID is empty, return all data
	if len(employeeID) > 0 {
		query = query.Where("employee_id IN ?", employeeID)
	}

	if err := query.Model(&po.AttendancePo{}).Find(&attendance).Error; err != nil {
		return entity.AttendanceStatistics{}, err
	}

	if err := query.Model(&po.OvertimePo{}).Find(&overtime).Error; err != nil {
		return entity.AttendanceStatistics{}, err
	}

	if err := query.Model(&po.LeavePo{}).Find(&leave).Error; err != nil {
		return entity.AttendanceStatistics{}, err
	}

	var (
		attendanceResult = make(map[int]int)
		overtimeResult   = make(map[int]int)
		leaveResult      = make(map[int]int)
	)

	for _, v := range attendance {
		attendanceResult[int(v.EmployeeID)]++
	}

	for _, v := range overtime {
		overtimeResult[int(v.EmployeeID)]++
	}

	for _, v := range leave {
		leaveResult[int(v.EmployeeID)]++
	}

	return entity.AttendanceStatistics{
		AttendanceEmployeeStatistics: func() map[int]entity.AttendanceEmployeeStatistics {
			result := make(map[int]entity.AttendanceEmployeeStatistics)
			for k, v := range attendanceResult {
				if item, ok := result[k]; ok {
					item.TotalAttendance = v
					result[k] = item
				} else {
					result[k] = entity.AttendanceEmployeeStatistics{
						EmployeeID:      k,
						TotalAttendance: v,
					}
				}
			}

			for k, v := range overtimeResult {
				if item, ok := result[k]; ok {
					item.TotalOvertime = v
					result[k] = item
				} else {
					result[k] = entity.AttendanceEmployeeStatistics{
						EmployeeID:    k,
						TotalOvertime: v,
					}
				}
			}

			for k, v := range leaveResult {
				if item, ok := result[k]; ok {
					item.TotalLeave = v
					result[k] = item
				} else {
					result[k] = entity.AttendanceEmployeeStatistics{
						EmployeeID: k,
						TotalLeave: v,
					}
				}
			}
			return result
		}(),
	}, nil
}

func NewAttendanceImpl(db *gorm.DB, infLogger logger.InfLogger) repository.AttendanceRepository {
	return &impl{
		db:     db,
		logger: infLogger,
	}
}
