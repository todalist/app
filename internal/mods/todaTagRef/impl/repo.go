package todaTagRefImpl

import (
	"gorm.io/gorm"
	"github.com/todalist/app/internal/mods/todaTagRef"
	"github.com/todalist/app/internal/common"
)

type TodaTagRefRepo struct {
	tx *gorm.DB
}

func (s *TodaTagRefRepo) Get(id uint) (*todaTagRef.TodaTagRef, error) {
	return s.First(&todaTagRef.TodaTagRefQuerier{Id: &id})
}

func (s *TodaTagRefRepo) First(querier *todaTagRef.TodaTagRefQuerier) (*todaTagRef.TodaTagRef, error) {
	list, err := s.List(querier)
	if err != nil {
		return nil, err
	}
	if len(list) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return list[0], nil
}

func (s *TodaTagRefRepo) Save(form *todaTagRef.TodaTagRef) (*todaTagRef.TodaTagRef, error) {
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

func (s *TodaTagRefRepo) List(querier *todaTagRef.TodaTagRefQuerier) ([]*todaTagRef.TodaTagRef, error) {
	var list []*todaTagRef.TodaTagRef
	sql := s.tx.Where(querier)
	sql = common.Paginate(sql, &querier.Pager)
	if err := sql.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (s *TodaTagRefRepo) Delete(id uint) (uint, error) {
	if err := s.tx.Where("id = ?", id).Delete(&todaTagRef.TodaTagRef{}).Error; err != nil {
		return 0, err
	}
	return id, nil
}

func (s *TodaTagRefRepo) ListTodaTagByTodaIds(ids []uint) ([]*todaTagRef.TodaTagVO, error) {
	var list []*todaTagRef.TodaTagVO
	return list, nil
}

func NewTodaTagRefRepo(tx *gorm.DB) todaTagRef.ITodaTagRefRepo {
	return &TodaTagRefRepo{
		tx: tx,
	}
}