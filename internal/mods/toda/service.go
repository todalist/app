package toda

import (
	"context"
	"github.com/todalist/app/internal/models/dto"
	"github.com/todalist/app/internal/models/entity"
)

type ITodaService interface {

	// basic crud
	Get(context.Context, uint) (*entity.Toda, error)

	Save(context.Context, *entity.Toda) (*entity.Toda, error)

	List(context.Context, *dto.TodaQuerier) ([]*entity.Toda, error)

	Delete(context.Context, uint) (uint, error)
}
