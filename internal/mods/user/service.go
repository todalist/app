package user

import (
	"context"

	"github.com/todalist/app/internal/models/dto"
	"github.com/todalist/app/internal/models/entity"
)

type IUserService interface {

	// basic crud
	Get(context.Context, uint) (*entity.User, error)

	First(context.Context, *dto.UserQuerier) (*entity.User, error)

	Save(context.Context, *entity.User) (*entity.User, error)

	List(context.Context, *dto.UserQuerier) ([]*entity.User, error)

	Delete(context.Context, uint) (uint, error)

}