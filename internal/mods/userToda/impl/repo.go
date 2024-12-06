package userTodaImpl

import (
	"fmt"
	"github.com/todalist/app/internal/common"
	"github.com/todalist/app/internal/models/dto"
	"github.com/todalist/app/internal/models/entity"
	"github.com/todalist/app/internal/models/vo"
	"github.com/todalist/app/internal/mods/userToda"
	"gorm.io/gorm"
)

type UserTodaRepo struct {
	tx *gorm.DB
}

func (s *UserTodaRepo) Get(id uint) (*entity.UserToda, error) {
	var model entity.UserToda
	if err := s.tx.Where("id = ?", id).First(&model).Error; err != nil {
		return nil, err
	}
	return &model, nil
}

func (s *UserTodaRepo) First(querier *dto.UserTodaQuerier) (*entity.UserToda, error) {
	list, err := s.List(querier)
	if err != nil {
		return nil, err
	}
	if len(list) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return list[0], nil
}

func (s *UserTodaRepo) Save(form *entity.UserToda) (*entity.UserToda, error) {
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

func (s *UserTodaRepo) List(querier *dto.UserTodaQuerier) ([]*entity.UserToda, error) {
	var list []*entity.UserToda
	sql := s.tx.Where(querier)
	sql = common.Paginate(sql, &querier.Pager)
	if err := sql.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (s *UserTodaRepo) Delete(id uint) (uint, error) {
	if err := s.tx.Where("id = ?", id).Delete(&entity.UserToda{}).Error; err != nil {
		return 0, err
	}
	return id, nil
}

func (s *UserTodaRepo) DeleteByTodaId(todaId uint) error {
	if err := s.tx.Where("toda_id = ?", todaId).Delete(&entity.UserToda{}).Error; err != nil {
		return err
	}
	return nil
}

// TODO use template to refine sql builder
func (s *UserTodaRepo) ListUserToda(querier *dto.ListUserTodaQuerier) ([]*vo.UserTodaVO, error) {
	var list []*vo.UserTodaVO
	cond, args := common.QuerierToSqlCondition(nil, querier, "t")
	if cond == "" {
		cond = "1 = 1"
	}
	joinTodaTag := ""
	todaTagCond := ""
	if querier.TodaTagId != nil {
		(*args)["todaTagId"] = querier.TodaTagId
		joinTodaTag = `
			INNER JOIN
				t_toda_tag as tt ON t.id = tt.toda_id
		`
		todaTagCond = " AND tt.toda_tag_id = @todaTagId AND tt.deleted_at IS NULL "
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
%s
WHERE
	ut.user_id = @userId
	AND %s
	AND t.deleted_at IS NULL
	AND ut.deleted_at IS NULL
	%s
`, joinTodaTag, cond, todaTagCond)
	(*args)["userId"] = querier.UserId
	sql := s.tx.Raw(sqlStr, args)
	sql = common.Paginate(sql, &querier.Pager)
	if err := sql.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func NewUserTodaRepo(tx *gorm.DB) userToda.IUserTodaRepo {
	return &UserTodaRepo{
		tx: tx,
	}
}
