package userImpl

import (
	"gorm.io/gorm"
	"dailydo.fe1.xyz/internal/mods/user"
)

type UserStore struct {
	tx *gorm.DB
}

func (s *UserStore) Get(id uint) (*user.User, error) {
	var model user.User
	if err := s.tx.Where("id = ?", id).First(&model).Error; err != nil {
		return nil, err
	}
	return &model, nil
}

func (s *UserStore) Save(form *user.User) (*user.User, error) {
	if err := s.tx.Save(form).Error; err != nil {
		return nil, err
	}
	return form, nil
}

func (s *UserStore) List(querier *user.UserQuerier) ([]*user.User, error) {
	var list []*user.User
	if err := s.tx.Where(querier).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (s *UserStore) Delete(id uint) (uint, error) {
	if err := s.tx.Where("id = ?", id).Delete(&user.User{}).Error; err != nil {
		return 0, err
	}
	return id, nil
}