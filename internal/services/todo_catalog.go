package services

import (
	"dailydo.fe1.xyz/internal/common"
	"dailydo.fe1.xyz/internal/globals"
	"dailydo.fe1.xyz/internal/models"
	"fmt"
	"go.uber.org/zap"
)

type TodoCatalogService struct{}

func (*TodoCatalogService) Save(catalog *models.TodoCatalog) (*models.TodoCatalog, error) {
	if catalog.ID > 0 {
		if err := globals.
			DB.
			Model(&models.TodoCatalog{}).
			Where("id = ?", catalog.ID).
			Omit("items_count", "items_finished_count").Updates(catalog).Error; err != nil {
			globals.LOG.Error("failed to update TodoCatalog catalog: ", zap.Error(err))
			return nil, err
		}
	} else {
		if err := globals.DB.Create(catalog).Error; err != nil {
			globals.LOG.Error("failed to save TodoCatalog catalog: ", zap.Error(err))
			return nil, err
		}
	}
	return catalog, nil
}

func (*TodoCatalogService) Delete(id *uint, userID *uint) (*uint, error) {
	if err := globals.DB.
		Where("id = ?", id).
		Where("user_id = ?", userID).Delete(&models.TodoCatalog{}).Error; err != nil {
		globals.LOG.Error("delete error", zap.Error(err))
		return nil, err
	}
	return id, nil
}

func (t *TodoCatalogService) Get(querier *models.TodoCatalogQuerier) (*models.TodoCatalog, error) {
	list, err := t.List(querier)
	if err != nil {
		return nil, err
	}
	if len(list) == 0 {
		return nil, fmt.Errorf("not found")
	}
	return list[0], nil
}

func (*TodoCatalogService) List(querier *models.TodoCatalogQuerier) ([]*models.TodoCatalog, error) {
	var list []*models.TodoCatalog
	querierMap := map[string]interface{}{}
	sql := globals.DB.Model(&models.TodoCatalog{})
	if querier.ID != nil {
		querierMap["id"] = querier.ID
	}
	querierMap["user_id"] = querier.UserID
	if querier.ParentID != nil {
		querierMap["parent_id"] = querier.ParentID
	} else {
		sql = sql.Where("parent_id is null")
	}
	if querier.PinToSideBar != nil {
		querierMap["pin_to_side_bar"] = querier.PinToSideBar
	}
	sql = sql.Where(querierMap)
	if err := common.Paginate(sql, &querier.Pager).Find(&list).Error; err != nil {
		globals.LOG.Error("list error", zap.Error(err))
		return nil, err
	}
	return list, nil
}
