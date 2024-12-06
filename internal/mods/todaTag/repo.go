package todaTag

import (
	"github.com/todalist/app/internal/models/dto"
	"github.com/todalist/app/internal/models/entity"
)

type ITodaTagRepo interface {

	// basic crud
	Get(uint) (*entity.TodaTag, error)

	First(*dto.TodaTagQuerier) (*entity.TodaTag, error)

	Save(*entity.TodaTag) (*entity.TodaTag, error)

	List(*dto.TodaTagQuerier) ([]*entity.TodaTag, error)

	Delete(uint) (uint, error)

}