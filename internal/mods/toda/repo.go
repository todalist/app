package toda

import (
	"github.com/todalist/app/internal/models/dto"
	"github.com/todalist/app/internal/models/entity"
	"github.com/todalist/app/internal/models/vo"
)

type ITodaRepo interface {

	// basic crud
	Get(uint) (*entity.Toda, error)

	Save(*entity.Toda) (*entity.Toda, error)

	List(*dto.TodaQuerier) ([]*entity.Toda, error)

	ListUserToda(*dto.ListUserTodaQuerier) ([]*vo.UserTodaVO, error)

	Delete(id uint) (uint, error)
}
