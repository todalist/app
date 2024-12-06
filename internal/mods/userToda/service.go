package userToda

import (
	"context"

	"github.com/todalist/app/internal/models/dto"
	"github.com/todalist/app/internal/models/entity"
	"github.com/todalist/app/internal/models/vo"
)

type IUserTodaService interface {

	// basic crud
	Get(context.Context, uint) (*entity.UserToda, error)

	First(context.Context, *dto.UserTodaQuerier) (*entity.UserToda, error)

	Save(context.Context, *entity.UserToda) (*entity.UserToda, error)

	List(context.Context, *dto.UserTodaQuerier) ([]*entity.UserToda, error)

	Delete(context.Context, uint) (uint, error)

	ListUserToda(context.Context, *dto.ListUserTodaQuerier) ([]*vo.UserTodaVO, error)

}