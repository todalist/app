package userTodaTag

import (
	"context"
)

type IUserTodaTagService interface {

	// basic crud
	Get(context.Context, uint) (*UserTodaTag, error)

	First(context.Context, *UserTodaTagQuerier) (*UserTodaTag, error)

	Save(context.Context, *UserTodaTag) (*UserTodaTag, error)

	List(context.Context, *UserTodaTagQuerier) ([]*UserTodaTag, error)

	Delete(context.Context, uint) (uint, error)

}