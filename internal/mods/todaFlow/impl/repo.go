package todaFlowImpl

import (
	"gorm.io/gorm"
	"dailydo.fe1.xyz/internal/mods/todaFlow"
	"dailydo.fe1.xyz/internal/common"
)

type TodaFlowRepo struct {
	tx *gorm.DB
}

func (s *TodaFlowRepo) Get(id uint) (*todaFlow.TodaFlow, error) {
	var model todaFlow.TodaFlow
	if err := s.tx.Where("id = ?", id).First(&model).Error; err != nil {
		return nil, err
	}
	return &model, nil
}

func (s *TodaFlowRepo) Save(form *todaFlow.TodaFlow) (*todaFlow.TodaFlow, error) {
	if form.Id == 0 {
		if err := s.tx.Create(form).Error; err != nil {
			return nil, err
		}
		return form, nil
	}
	if err := s.
		tx.
		Model(form).
		Where("id = ?", form.Id).
		Where("updated_at <=", form.UpdatedAt).
		Updates(form).Error; err != nil {
		return nil, err
	}
	return form, nil
}

func (s *TodaFlowRepo) List(querier *todaFlow.TodaFlowQuerier) ([]*todaFlow.TodaFlow, error) {
	var list []*todaFlow.TodaFlow
	sql := s.tx.Where(querier)
	sql = common.Paginate(sql, &querier.Pager)
	if err := sql.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (s *TodaFlowRepo) Delete(id uint) (uint, error) {
	if err := s.tx.Where("id = ?", id).Delete(&todaFlow.TodaFlow{}).Error; err != nil {
		return 0, err
	}
	return id, nil
}

func NewTodaFlowRepo(tx *gorm.DB) todaFlow.ITodaFlowRepo {
	return &TodaFlowRepo{
		tx: tx,
	}
}