package toda

import (
	"github.com/todalist/app/internal/models/dto"
	"github.com/todalist/app/internal/models/entity"
)

type ITodaRepo interface {

	// basic crud
	Get(uint) (*entity.Toda, error)

	Save(*entity.Toda) (*entity.Toda, error)

	List(*dto.TodaQuerier) ([]*entity.Toda, error)

	Delete(id uint) (uint, error)
}
