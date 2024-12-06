package userTodaTag

import (
	"context"

	"github.com/todalist/app/internal/models/dto"
	"github.com/todalist/app/internal/models/entity"
)

type IUserTodaTagService interface {

	// basic crud
	Get(context.Context, uint) (*entity.UserTodaTag, error)

	First(context.Context, *dto.UserTodaTagQuerier) (*entity.UserTodaTag, error)

	Save(context.Context, *entity.UserTodaTag) (*entity.UserTodaTag, error)

	List(context.Context, *dto.UserTodaTagQuerier) ([]*entity.UserTodaTag, error)

	Delete(context.Context, uint) (uint, error)

}