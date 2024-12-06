package todaFlow

import (
	"github.com/todalist/app/internal/models/dto"
	"github.com/todalist/app/internal/models/entity"
)

type ITodaFlowRepo interface {

	// basic crud
	Get(uint) (*entity.TodaFlow, error)

	First(*dto.TodaFlowQuerier) (*entity.TodaFlow, error)

	Save(*entity.TodaFlow) (*entity.TodaFlow, error)

	List(*dto.TodaFlowQuerier) ([]*entity.TodaFlow, error)

	Delete(uint) (uint, error)

}