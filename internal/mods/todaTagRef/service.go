package todaTagRef

import (
	"context"

	"github.com/todalist/app/internal/models/dto"
	"github.com/todalist/app/internal/models/entity"
)

type ITodaTagRefService interface {

	// basic crud
	Get(context.Context, uint) (*entity.TodaTagRef, error)

	First(context.Context, *dto.TodaTagRefQuerier) (*entity.TodaTagRef, error)

	Save(context.Context, *entity.TodaTagRef) (*entity.TodaTagRef, error)

	List(context.Context, *dto.TodaTagRefQuerier) ([]*entity.TodaTagRef, error)

	Delete(context.Context, uint) (uint, error)

}