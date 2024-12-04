package userImpl

import (
	"gorm.io/gorm"
	"dailydo.fe1.xyz/internal/mods/user"
	"dailydo.fe1.xyz/internal/common"
)

type UserRepo struct {
	tx *gorm.DB
}

func (s *UserRepo) Get(id uint) (*user.User, error) {
	var model user.User
	if err := s.tx.Where("id = ?", id).First(&model).Error; err != nil {
		return nil, err
	}
	return &model, nil
}

func (s *UserRepo) Save(form *user.User) (*user.User, error) {
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
		Where("updated_at <=", form.UpdatedAt).Omit("password", "username", "email", ).
		Updates(form).Error; err != nil {
		return nil, err
	}
	return form, nil
}

func (s *UserRepo) List(querier *user.UserQuerier) ([]*user.User, error) {
	var list []*user.User
	sql := s.tx.Where(querier)
	sql = common.Paginate(sql, &querier.Pager)
	if err := sql.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (s *UserRepo) Delete(id uint) (uint, error) {
	if err := s.tx.Where("id = ?", id).Delete(&user.User{}).Error; err != nil {
		return 0, err
	}
	return id, nil
}

func NewUserRepo(tx *gorm.DB) user.IUserRepo {
	return &UserRepo{
		tx: tx,
	}
}