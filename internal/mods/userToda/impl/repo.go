package userTodaImpl

import (
	"dailydo.fe1.xyz/internal/common"
	"dailydo.fe1.xyz/internal/mods/userToda"
	"gorm.io/gorm"
)

type UserTodaRepo struct {
	tx *gorm.DB
}

func (s *UserTodaRepo) Get(id uint) (*userToda.UserToda, error) {
	var model userToda.UserToda
	if err := s.tx.Where("id = ?", id).First(&model).Error; err != nil {
		return nil, err
	}
	return &model, nil
}

func (s *UserTodaRepo) First(querier *userToda.UserTodaQuerier) (*userToda.UserToda, error) {
	l, err := s.List(querier)
	if err != nil {
		return nil, err
	}
	if len(l) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return l[0], nil
}

func (s *UserTodaRepo) Save(form *userToda.UserToda) (*userToda.UserToda, error) {
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

func (s *UserTodaRepo) List(querier *userToda.UserTodaQuerier) ([]*userToda.UserToda, error) {
	var list []*userToda.UserToda
	sql := s.tx.Where(querier)
	sql = common.Paginate(sql, &querier.Pager)
	if err := sql.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (s *UserTodaRepo) Delete(id uint) (uint, error) {
	if err := s.tx.Where("id = ?", id).Delete(&userToda.UserToda{}).Error; err != nil {
		return 0, err
	}
	return id, nil
}

func (s *UserTodaRepo) DeleteByTodaId(todaId uint) error {
	if err := s.tx.Where("toda_id = ?", todaId).Delete(&userToda.UserToda{}).Error; err != nil {
		return err
	}
	return nil
}

func NewUserTodaRepo(tx *gorm.DB) userToda.IUserTodaRepo {
	return &UserTodaRepo{
		tx: tx,
	}
}
