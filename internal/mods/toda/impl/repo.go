package todaImpl

import (
	"fmt"

	"github.com/todalist/app/internal/common"
	"github.com/todalist/app/internal/models/dto"
	"github.com/todalist/app/internal/models/entity"
	"github.com/todalist/app/internal/mods/toda"
	"gorm.io/gorm"
)

type TodaRepo struct {
	tx *gorm.DB
}

func (s *TodaRepo) Get(id uint) (*entity.Toda, error) {
	return s.First(&dto.TodaQuerier{Id: &id})
}

func (s *TodaRepo) First(querier *dto.TodaQuerier) (*entity.Toda, error) {
	list, err := s.List(querier)
	if err != nil {
		return nil, err
	}
	if len(list) < 1 {
		return nil, nil
	}
	return list[0], nil
}

func (s *TodaRepo) Save(form *entity.Toda) (*entity.Toda, error) {
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
		Where("updated_at <=", form.UpdatedAt).Omit("elapsed").
		Updates(form).Error; err != nil {
		return nil, err
	}
	return form, nil
}

func (s *TodaRepo) List(querier *dto.TodaQuerier) ([]*entity.Toda, error) {
	var list []*entity.Toda

	cond, args := common.QuerierToSqlCondition(nil, querier, "t")
	if cond == "" {
		cond = "1 = 1"
	}

	sqlStr := fmt.Sprintf(`
SELECT 
	t.*
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
}

func (s *TodaRepo) Delete(id uint) (uint, error) {
	if err := s.tx.Where("id = ?", id).Delete(&entity.Toda{}).Error; err != nil {
		return 0, err
	}
	return id, nil
}

func NewTodaRepo(tx *gorm.DB) toda.ITodaRepo {
	return &TodaRepo{
		tx: tx,
	}
}
