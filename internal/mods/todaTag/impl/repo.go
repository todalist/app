package todaTagImpl

import (
	"fmt"

	"github.com/todalist/app/internal/common"
	"github.com/todalist/app/internal/models/dto"
	"github.com/todalist/app/internal/models/entity"
	"github.com/todalist/app/internal/models/vo"
	"github.com/todalist/app/internal/mods/todaTag"
	"gorm.io/gorm"
)

type TodaTagRepo struct {
	tx *gorm.DB
}

func (s *TodaTagRepo) Get(id uint) (*entity.TodaTag, error) {
	return s.First(&dto.TodaTagQuerier{
		Id: &id,
	})
}

func (s *TodaTagRepo) First(querier *dto.TodaTagQuerier) (*entity.TodaTag, error) {
	list, err := s.List(querier)
	if err != nil {
		return nil, err
	}
	if len(list) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return list[0], nil
}

func (s *TodaTagRepo) Save(form *entity.TodaTag) (*entity.TodaTag, error) {
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
		Where("updated_at <= ?", form.UpdatedAt).
		Updates(form).Error; err != nil {
		return nil, err
	}
	return form, nil
}

func (s *TodaTagRepo) List(querier *dto.TodaTagQuerier) ([]*entity.TodaTag, error) {
	var list []*entity.TodaTag
	sql := s.tx.Model(&entity.TodaTag{}).Where(querier)
	sql = common.Paginate(sql, &querier.Pager)
	if err := sql.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (s *TodaTagRepo) Delete(id uint) (uint, error) {
	if err := s.tx.Where("id = ?", id).Delete(&entity.TodaTag{}).Error; err != nil {
		return 0, err
	}
	return id, nil
}

func (s *TodaTagRepo) ListUserTodaTag(querier *dto.ListUserTodaTagQuerier) ([]*vo.UserTodaTagVO, error) {
	var list []*vo.UserTodaTagVO
	sql := s.
		tx.
		Table("t_toda_tag as tt").
		Select("tt.*", "utt.id as user_toda_tag_id", "utt.user_id", "utt.pin_top").
		InnerJoins(
			"INNER JOIN t_user_toda_tag as utt ON utt.toda_tag_id=tt.id",
		).
		Where("utt.user_id = ?", querier.UserId).
		Where(querier.TodaTagQuerier)
	if querier.PinTop != nil {
		sql = sql.Where("utt.pin_top = ?", querier.PinTop)
	}
	if querier.UserTodaTagId != nil {
		sql = sql.Where("utt.id = ?", querier.UserTodaTagId)
	}
	if querier.Name != nil {
		// TODO: may need full text search
		sql = sql.Where("tt.name like ? ", fmt.Sprintf("%%%s%%", *querier.Name))

	}
	if len(querier.Ids) > 0 {
		sql = sql.Where("tt.id IN ?", querier.Ids)
	}
	sql = common.Paginate(sql, &querier.Pager)
	if err := sql.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (s *TodaTagRepo) ListTodaTagByTodaIds(ids []uint) ([]*vo.TodaTagRefVO, error) {
	var list []*vo.TodaTagRefVO
	sql := s.tx.Table("t_toda_tag as tt").
		Select("tt.*", "ttr.id as toda_tag_ref_id", "ttr.toda_id").
		InnerJoins(
			"INNER JOIN t_toda_tag_ref as ttr ON tt.id=ttr.toda_tag_id",
		).
		Where("ttr.toda_id in (?)", ids)
	if err := sql.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func NewTodaTagRepo(tx *gorm.DB) todaTag.ITodaTagRepo {
	return &TodaTagRepo{
		tx: tx,
	}
}
