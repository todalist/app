package services

import (
	"dailydo.fe1.xyz/internal/common"
	"dailydo.fe1.xyz/internal/globals"
	"dailydo.fe1.xyz/internal/models"
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
)

type CollectionService struct{}

func (s *CollectionService) Save(coll *models.Collection) (*models.Collection, error) {
	if err := globals.DB.Save(coll).Error; err != nil {
		globals.LOG.Error("save error", zap.Error(err))
		return nil, err
	}
	return coll, nil
}

func (s *CollectionService) Get(querier *models.CollectionQuerier) (*models.Collection, error) {
	list, err := s.List(querier)
	if err != nil {
		return nil, err
	}
	if len(list) == 0 {
		return nil, fiber.ErrNotFound
	}
	return list[0], nil
}

func (s *CollectionService) List(querier *models.CollectionQuerier) ([]*models.Collection, error) {
	var list []*models.Collection
	querierMap := map[string]interface{}{}
	if querier.ID != nil {
		querierMap["id"] = querier.ID
	}
	if querier.Status != nil {
		querierMap["status"] = querier.Status
	}
	if querier.UserID != nil {
		querierMap["user_id"] = querier.UserID
	}
	sql := globals.DB.Model(&models.Collection{}).Where(querierMap)
	if err := common.Paginate1(sql, &querier.Pager).Find(&list).Error; err != nil {
		globals.LOG.Error("list error", zap.Error(err))
		return nil, err
	}
	return list, nil
}

func (s *CollectionService) Delete(querier *models.CollectionQuerier) (*uint, error) {
	if err := globals.DB.
		Where("id = ?", querier.ID).
		Where("user_id = ?", querier.UserID).Delete(&models.Collection{}).Error; err != nil {
		globals.LOG.Error("delete error", zap.Error(err))
		return nil, err
	}
	return querier.ID, nil
}
