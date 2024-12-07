package todaImpl

import (
	"github.com/todalist/app/internal/common"
	"github.com/todalist/app/internal/models/dto"
	"github.com/todalist/app/internal/models/entity"
	"github.com/todalist/app/internal/models/vo"
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
	sql := s.tx.Model(&entity.Toda{}).Where(querier)
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

func (s *TodaRepo) ListUserToda(querier *dto.ListUserTodaQuerier) ([]*vo.UserTodaVO, error) {
	var list []*vo.UserTodaVO
	sql := s.
		tx.
		Table("t_toda as tt").
		Select("tt.*", "utt.id as user_toda_id", "utt.user_id").
		InnerJoins(
			"INNER JOIN t_user_toda as utt ON utt.toda_id=tt.id",
		).
		Where("utt.user_id = ?", querier.UserId).
		Where(querier.TodaQuerier)
	sql = common.Paginate(sql, &querier.Pager)
	if err := sql.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func NewTodaRepo(tx *gorm.DB) toda.ITodaRepo {
	return &TodaRepo{
		tx: tx,
	}
}
