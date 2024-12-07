package todaTagRef

import (
	"github.com/todalist/app/internal/models/dto"
	"github.com/todalist/app/internal/models/entity"
)

type ITodaTagRefRepo interface {

	// basic crud
	Get(uint) (*entity.TodaTagRef, error)

	First(*dto.TodaTagRefQuerier) (*entity.TodaTagRef, error)

	Save(*entity.TodaTagRef) (*entity.TodaTagRef, error)

	List(*dto.TodaTagRefQuerier) ([]*entity.TodaTagRef, error)

	Delete(uint) (uint, error)

}