package todaTagImpl

import (
	"github.com/todalist/app/internal/common"
	"github.com/todalist/app/internal/models/dto"
	"github.com/todalist/app/internal/models/entity"
	"github.com/todalist/app/internal/mods/todaTag"
	"gorm.io/gorm"
)

type TodaTagRepo struct {
	tx *gorm.DB
}

func (s *TodaTagRepo) Get(id uint) (*entity.TodaTag, error) {
	return s.First(&dto.TodaTagQuerier{
		Id: &id,
	})
}

func (s *TodaTagRepo) First(querier *dto.TodaTagQuerier) (*entity.TodaTag, error) {
	list, err := s.List(querier)
	if err != nil {
		return nil, err
	}
	if len(list) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return list[0], nil
}

func (s *TodaTagRepo) Save(form *entity.TodaTag) (*entity.TodaTag, error) {
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

func (s *TodaTagRepo) List(querier *dto.TodaTagQuerier) ([]*entity.TodaTag, error) {
	var list []*entity.TodaTag
	sql := s.tx.Model(&entity.TodaTag{}).Where(querier)
	sql = common.Paginate(sql, &querier.Pager)
	if err := sql.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (s *TodaTagRepo) Delete(id uint) (uint, error) {
	if err := s.tx.Where("id = ?", id).Delete(&entity.TodaTag{}).Error; err != nil {
		return 0, err
	}
	return id, nil
}

func NewTodaTagRepo(tx *gorm.DB) todaTag.ITodaTagRepo {
	return &TodaTagRepo{
		tx: tx,
	}
}
