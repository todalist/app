package userTodaImpl

import (
	"fmt"

	"github.com/todalist/app/internal/common"
	"github.com/todalist/app/internal/mods/userToda"
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
	list, err := s.List(querier)
	if err != nil {
		return nil, err
	}
	if len(list) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return list[0], nil
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

func (s *UserTodaRepo) ListUserToda(querier *userToda.ListUserTodaQuerier) ([]*userToda.UserTodaVO, error) {
	var list []*userToda.UserTodaVO
	cond, args := common.QuerierToSqlCondition(nil, querier, "t")
	if cond == "" {
		cond = "1 = 1"
	}

	sqlStr := fmt.Sprintf(`
SELECT 
	t.*,
	ut.id as user_toda_id,
	ut.user_id,
FROM 
	t_user_toda as ut
INNER JOIN
	t_toda as t ON ut.toda_id = t.id
WHERE
	ut.user_id = @userId
	AND %s
	AND t.deleted_at IS NULL
	AND ut.deleted_at IS NULL
`, cond)
	(*args)["userId"] = querier.UserId
	sql := s.tx.Raw(sqlStr, args)
	sql = common.Paginate(sql, &querier.Pager)
	if err := sql.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
	return list, nil
}

func NewUserTodaRepo(tx *gorm.DB) userToda.IUserTodaRepo {
	return &UserTodaRepo{
		tx: tx,
	}
}
