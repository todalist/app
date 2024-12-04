package todaImpl

import (
	"dailydo.fe1.xyz/internal/common"
	"dailydo.fe1.xyz/internal/mods/toda"
	"gorm.io/gorm"
)

type TodaRepo struct {
	tx *gorm.DB
}

func (s *TodaRepo) Get(id uint) (*toda.Toda, error) {
	var model toda.Toda
	if err := s.tx.Where("id = ?", id).First(&model).Error; err != nil {
		return nil, err
	}
	return &model, nil
}

func (s *TodaRepo) Save(form *toda.Toda) (*toda.Toda, error) {
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
		Where("updated_at <=", form.UpdatedAt).Omit("elapsed").
		Updates(form).Error; err != nil {
		return nil, err
	}
	return form, nil
}

func (s *TodaRepo) List(querier *toda.TodaQuerier) ([]*toda.Toda, error) {
	var list []*toda.Toda
	sql := s.tx.Where(querier)
	sql = common.Paginate(sql, &querier.Pager)
	if err := sql.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (s *TodaRepo) Delete(id uint) (uint, error) {
	if err := s.tx.Where("id = ?", id).Delete(&toda.Toda{}).Error; err != nil {
		return 0, err
	}
	return id, nil
}

func NewTodaRepo(tx *gorm.DB) toda.ITodaRepo {
	return &TodaRepo{
		tx: tx,
	}
}
