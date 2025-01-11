package userImpl

import (
	"github.com/todalist/app/internal/common"
	"github.com/todalist/app/internal/models/dto"
	"github.com/todalist/app/internal/models/entity"
	"github.com/todalist/app/internal/mods/user"
	"gorm.io/gorm"
)

type UserRepo struct {
	tx *gorm.DB
}

func (s *UserRepo) Get(id uint) (*entity.User, error) {
	var model entity.User
	if err := s.tx.Where("id = ?", id).First(&model).Error; err != nil {
		return nil, err
	}
	return &model, nil
}

func (s *UserRepo) First(querier *dto.UserQuerier) (*entity.User, error) {
	list, err := s.List(querier)
	if err != nil {
		return nil, err
	}
	if len(list) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return list[0], nil
}

func (s *UserRepo) Save(form *entity.User) (*entity.User, error) {
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
		Where("updated_at <= ?", form.UpdatedAt).Omit("password", "username", "email").
		Updates(form).Error; err != nil {
		return nil, err
	}
	return form, nil
}

func (s *UserRepo) List(querier *dto.UserQuerier) ([]*entity.User, error) {
	var list []*entity.User
	sql := s.tx.Where(querier)
	sql = common.Paginate(sql, &querier.Pager)
	if err := sql.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (s *UserRepo) Delete(id uint) (uint, error) {
	if err := s.tx.Where("id = ?", id).Delete(&entity.User{}).Error; err != nil {
		return 0, err
	}
	return id, nil
}

func NewUserRepo(tx *gorm.DB) user.IUserRepo {
	return &UserRepo{
		tx: tx,
	}
}
