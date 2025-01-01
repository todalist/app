package todaTag

import (
	"context"
	"github.com/todalist/app/internal/models/dto"
	"github.com/todalist/app/internal/models/entity"
	"github.com/todalist/app/internal/models/vo"
)

type ITodaTagService interface {

	// basic crud
	Get(context.Context, uint) (*entity.TodaTag, error)

	First(context.Context, *dto.TodaTagQuerier) (*entity.TodaTag, error)

	Save(context.Context, *dto.TodaTagSaveDTO) (*vo.UserTodaTagVO, error)

	List(context.Context, *dto.ListUserTodaTagQuerier) ([]*vo.UserTodaTagVO, error)

	Delete(context.Context, uint) (uint, error)

	SaveUserTodaTag(context.Context, *entity.UserTodaTag) (*entity.UserTodaTag, error)

	// ListUserTodaTag(context.Context, *dto.TodaTagQuerier) ([]*entity.TodaTag, error)

}
