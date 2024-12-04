package todaTagImpl

import (
	"gorm.io/gorm"
	"dailydo.fe1.xyz/internal/mods/todaTag"
	"dailydo.fe1.xyz/internal/common"
)

type TodaTagRepo struct {
	tx *gorm.DB
}

func (s *TodaTagRepo) Get(id uint) (*todaTag.TodaTag, error) {
	var model todaTag.TodaTag
	if err := s.tx.Where("id = ?", id).First(&model).Error; err != nil {
		return nil, err
	}
	return &model, nil
}

func (s *TodaTagRepo) Save(form *todaTag.TodaTag) (*todaTag.TodaTag, error) {
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

func (s *TodaTagRepo) List(querier *todaTag.TodaTagQuerier) ([]*todaTag.TodaTag, error) {
	var list []*todaTag.TodaTag
	sql := s.tx.Where(querier)
	sql = common.Paginate(sql, &querier.Pager)
	if err := sql.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (s *TodaTagRepo) Delete(id uint) (uint, error) {
	if err := s.tx.Where("id = ?", id).Delete(&todaTag.TodaTag{}).Error; err != nil {
		return 0, err
	}
	return id, nil
}

func NewTodaTagRepo(tx *gorm.DB) todaTag.ITodaTagRepo {
	return &TodaTagRepo{
		tx: tx,
	}
}