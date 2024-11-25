package services

import (
	"dailydo.fe1.xyz/internal/common"
	"dailydo.fe1.xyz/internal/globals"
	"dailydo.fe1.xyz/internal/models"
	"fmt"
	"go.uber.org/zap"
)

type TodoService struct{}

func (*TodoService) Save(todo *models.Todo) (*models.Todo, error) {
	if err := globals.DB.Save(todo).Error; err != nil {
		globals.LOG.Error("save error", zap.Error(err))
		return nil, err
	}
	return todo, nil
}

func (*TodoService) Delete(querier *models.TodoQuerier) (*uint, error) {
	if err := globals.DB.
		Where("id = ?", querier.ID).
		Where("user_id = ?", querier.UserID).Delete(&models.Todo{}).Error; err != nil {
		globals.LOG.Error("delete error", zap.Error(err))
		return nil, err
	}
	return querier.ID, nil
}

func (t *TodoService) Get(querier *models.TodoQuerier) (*models.Todo, error) {
	list, err := t.List(querier)
	if err != nil {
		return nil, err
	}
	if len(list) == 0 {
		return nil, fmt.Errorf("not found")
	}
	return list[0], nil
}

func (*TodoService) List(querier *models.TodoQuerier) ([]*models.Todo, error) {
	var list []*models.Todo
	querierMap := map[string]interface{}{}
	sql := globals.DB.Model(&models.Todo{})
	if querier.ID != nil {
		querierMap["id"] = querier.ID
	}
	if querier.Status != nil {
		querierMap["status"] = querier.Status
	}
	if querier.UserID != nil {
		querierMap["user_id"] = querier.UserID
	}
	if querier.ParentID != nil {
		querierMap["parent_id"] = querier.ParentID
	} else {
		sql = sql.Where("parent_id is null")
	}
	sql = sql.Where(querierMap)
	if querier.TimeRange != nil {
		sql = querier.TimeRange.RangeSql(sql, "deadline")
	}
	if err := common.Paginate1(sql, &querier.Pager).Find(&list).Error; err != nil {
		globals.LOG.Error("list error", zap.Error(err))
		return nil, err
	}
	return list, nil
}
