package userToda

import (
	"context"

	"github.com/todalist/app/internal/models/dto"
	"github.com/todalist/app/internal/models/entity"
	"github.com/todalist/app/internal/models/vo"
)

type IUserTodaService interface {

	ListUserToda(context.Context, *dto.ListUserTodaQuerier) ([]*vo.UserTodaVO, error)

	CreateUserToda(context.Context, *entity.Toda) (*entity.Toda, error)

}