package userTodaTagImpl

import (
	"github.com/todalist/app/internal/common"
	"github.com/todalist/app/internal/models/dto"
	"github.com/todalist/app/internal/models/entity"
	"github.com/todalist/app/internal/mods/userTodaTag"
	"gorm.io/gorm"
)

type UserTodaTagRepo struct {
	tx *gorm.DB
}

func (s *UserTodaTagRepo) Get(id uint) (*entity.UserTodaTag, error) {
	var model entity.UserTodaTag
	if err := s.tx.Where("id = ?", id).First(&model).Error; err != nil {
		return nil, err
	}
	return &model, nil
}

func (s *UserTodaTagRepo) First(querier *dto.UserTodaTagQuerier) (*entity.UserTodaTag, error) {
	list, err := s.List(querier)
	if err != nil {
		return nil, err
	}
	if len(list) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return list[0], nil
}

// Save a userTodaTag model.
//
// If the model is not exists, save it as a new record.
// If the model is exists, update it.
func (s *UserTodaTagRepo) Save(form *entity.UserTodaTag) (*entity.UserTodaTag, error) {
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
		Where("user_id = ?", form.UserId).
		Updates(form).Error; err != nil {
		return nil, err
	}
	return form, nil
}

func (s *UserTodaTagRepo) List(querier *dto.UserTodaTagQuerier) ([]*entity.UserTodaTag, error) {
	var list []*entity.UserTodaTag
	sql := s.tx.Where(querier)
	sql = common.Paginate(sql, &querier.Pager)
	if err := sql.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (s *UserTodaTagRepo) Delete(id uint) (uint, error) {
	if err := s.tx.Where("id = ?", id).Delete(&entity.UserTodaTag{}).Error; err != nil {
		return 0, err
	}
	return id, nil
}

func NewUserTodaTagRepo(tx *gorm.DB) userTodaTag.IUserTodaTagRepo {
	return &UserTodaTagRepo{
		tx: tx,
	}
}
