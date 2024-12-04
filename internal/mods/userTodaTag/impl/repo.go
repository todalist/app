package userTodaTagImpl

import (
	"gorm.io/gorm"
	"dailydo.fe1.xyz/internal/mods/userTodaTag"
	"dailydo.fe1.xyz/internal/common"
)

type UserTodaTagRepo struct {
	tx *gorm.DB
}

func (s *UserTodaTagRepo) Get(id uint) (*userTodaTag.UserTodaTag, error) {
	var model userTodaTag.UserTodaTag
	if err := s.tx.Where("id = ?", id).First(&model).Error; err != nil {
		return nil, err
	}
	return &model, nil
}

func (s *UserTodaTagRepo) First(querier *userTodaTag.UserTodaTagQuerier) (*userTodaTag.UserTodaTag, error) {
	list, err := s.List(querier)
	if err != nil {
		return nil, err
	}
	if len(list) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return list[0], nil
}

func (s *UserTodaTagRepo) Save(form *userTodaTag.UserTodaTag) (*userTodaTag.UserTodaTag, error) {
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

func (s *UserTodaTagRepo) List(querier *userTodaTag.UserTodaTagQuerier) ([]*userTodaTag.UserTodaTag, error) {
	var list []*userTodaTag.UserTodaTag
	sql := s.tx.Where(querier)
	sql = common.Paginate(sql, &querier.Pager)
	if err := sql.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (s *UserTodaTagRepo) Delete(id uint) (uint, error) {
	if err := s.tx.Where("id = ?", id).Delete(&userTodaTag.UserTodaTag{}).Error; err != nil {
		return 0, err
	}
	return id, nil
}

func NewUserTodaTagRepo(tx *gorm.DB) userTodaTag.IUserTodaTagRepo {
	return &UserTodaTagRepo{
		tx: tx,
	}
}