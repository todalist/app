package toda

import (
	"context"
	"github.com/todalist/app/internal/models/dto"
	"github.com/todalist/app/internal/models/entity"
	"github.com/todalist/app/internal/models/vo"
)

type ITodaService interface {

	// basic crud
	Get(context.Context, uint) (*entity.Toda, error)

	Save(context.Context, *dto.TodaSaveDTO) (*vo.UserTodaVO, error)

	List(context.Context, *dto.ListUserTodaQuerier) ([]*vo.UserTodaVO, error)

	Delete(context.Context, uint) (uint, error)

	FlowToda(context.Context, *dto.FlowTodaDTO) (*uint, error)
}
