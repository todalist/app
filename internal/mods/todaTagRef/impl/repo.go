package todaTagRefImpl

import (
	"github.com/todalist/app/internal/common"
	"github.com/todalist/app/internal/models/dto"
	"github.com/todalist/app/internal/models/entity"
	"github.com/todalist/app/internal/mods/todaTagRef"
	"gorm.io/gorm"
)

type TodaTagRefRepo struct {
	tx *gorm.DB
}

func (s *TodaTagRefRepo) Get(id uint) (*entity.TodaTagRef, error) {
	return s.First(&dto.TodaTagRefQuerier{Id: &id})
}

func (s *TodaTagRefRepo) First(querier *dto.TodaTagRefQuerier) (*entity.TodaTagRef, error) {
	list, err := s.List(querier)
	if err != nil {
		return nil, err
	}
	if len(list) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return list[0], nil
}

func (s *TodaTagRefRepo) Save(form *entity.TodaTagRef) (*entity.TodaTagRef, error) {
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
		Where("updated_at <= ?", form.UpdatedAt).
		Updates(form).Error; err != nil {
		return nil, err
	}
	return form, nil
}

func (s *TodaTagRefRepo) List(querier *dto.TodaTagRefQuerier) ([]*entity.TodaTagRef, error) {
	var list []*entity.TodaTagRef
	sql := s.tx.Where(querier)
	sql = common.Paginate(sql, &querier.Pager)
	if err := sql.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (s *TodaTagRefRepo) Delete(id uint) (uint, error) {
	if err := s.tx.Where("id = ?", id).Delete(&entity.TodaTagRef{}).Error; err != nil {
		return 0, err
	}
	return id, nil
}

func NewTodaTagRefRepo(tx *gorm.DB) todaTagRef.ITodaTagRefRepo {
	return &TodaTagRefRepo{
		tx: tx,
	}
}
