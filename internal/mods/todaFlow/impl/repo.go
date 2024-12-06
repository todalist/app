package todaFlowImpl

import (
	"github.com/todalist/app/internal/common"
	"github.com/todalist/app/internal/models/dto"
	"github.com/todalist/app/internal/models/entity"
	"github.com/todalist/app/internal/mods/todaFlow"
	"gorm.io/gorm"
)

type TodaFlowRepo struct {
	tx *gorm.DB
}

func (s *TodaFlowRepo) Get(id uint) (*entity.TodaFlow, error) {
	var model entity.TodaFlow
	if err := s.tx.Where("id = ?", id).First(&model).Error; err != nil {
		return nil, err
	}
	return &model, nil
}

func (s *TodaFlowRepo) First(querier *dto.TodaFlowQuerier) (*entity.TodaFlow, error) {
	list, err := s.List(querier)
	if err != nil {
		return nil, err
	}
	if len(list) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return list[0], nil
}

func (s *TodaFlowRepo) Save(form *entity.TodaFlow) (*entity.TodaFlow, error) {
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

func (s *TodaFlowRepo) List(querier *dto.TodaFlowQuerier) ([]*entity.TodaFlow, error) {
	var list []*entity.TodaFlow
	sql := s.tx.Where(querier)
	sql = common.Paginate(sql, &querier.Pager)
	if err := sql.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (s *TodaFlowRepo) Delete(id uint) (uint, error) {
	if err := s.tx.Where("id = ?", id).Delete(&entity.TodaFlow{}).Error; err != nil {
		return 0, err
	}
	return id, nil
}

func NewTodaFlowRepo(tx *gorm.DB) todaFlow.ITodaFlowRepo {
	return &TodaFlowRepo{
		tx: tx,
	}
}
