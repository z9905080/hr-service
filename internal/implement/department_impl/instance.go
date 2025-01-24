package department_impl

import (
	"github.com/z9905080/hr_service/internal/domain/entity"
	"github.com/z9905080/hr_service/internal/domain/repository"
	"github.com/z9905080/hr_service/internal/implement/department_impl/po"
	"gorm.io/gorm"
)

type impl struct {
	db *gorm.DB
}

func (i *impl) AddDepartment(e entity.Department) (entity.Department, error) {
	model := po.NewDepartmentPo(e)
	if err := i.db.Create(&model).Error; err != nil {
		return entity.Department{}, err
	}

	return model.ToEntity(), nil

}

func (i *impl) QueryDepartment(e entity.Department) (entity.Department, error) {
	var model po.DepartmentPo
	if err := i.db.Where("id = ?", e.ID).First(&model).Error; err != nil {
		return entity.Department{}, err
	}

	return model.ToEntity(), nil
}

func (i *impl) UpdateDepartment(e entity.DepartmentUpdate) (entity.Department, error) {
	model := po.NewDepartmentUpdatePo(e)
	if err := i.db.Model(&po.DepartmentPo{}).Where("id = ?", e.ID).Updates(model.ToUpdateMap()).Error; err != nil {
		return entity.Department{}, err
	}

	var department po.DepartmentPo
	if err := i.db.Where("id = ?", e.ID).First(&department).Error; err != nil {
		return entity.Department{}, err
	}

	return department.ToEntity(), nil
}

func (i *impl) DeleteDepartment(e entity.Department) error {
	if err := i.db.Where("id = ?", e.ID).Delete(&po.DepartmentPo{}).Error; err != nil {
		return err
	}

	return nil
}

func (i *impl) ListDepartment(limit int, page int, field string, order string) ([]entity.Department, error) {
	var model []po.DepartmentPo

	db := i.db
	if field != "" && order != "" {
		db = db.Order(field + " " + order)
	}
	if err := db.Limit(limit).Offset((page - 1) * limit).Find(&model).Error; err != nil {
		return nil, err
	}

	var result []entity.Department
	for _, v := range model {
		result = append(result, v.ToEntity())
	}

	return result, nil
}

func NewDepartmentImpl(db *gorm.DB) repository.DepartmentRepository {
	return &impl{db: db}
}
