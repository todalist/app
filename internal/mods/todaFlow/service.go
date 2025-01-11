package todaFlow

import (
	"context"

	"github.com/todalist/app/internal/models/dto"
	"github.com/todalist/app/internal/models/entity"
)

type ITodaFlowService interface {

	// basic crud
	Get(context.Context, uint) (*entity.TodaFlow, error)

	First(context.Context, *dto.TodaFlowQuerier) (*entity.TodaFlow, error)

	Save(context.Context, *entity.TodaFlow) (*entity.TodaFlow, error)

	List(context.Context, *dto.TodaFlowQuerier) ([]*entity.TodaFlow, error)

	Delete(context.Context, uint) (uint, error)

}