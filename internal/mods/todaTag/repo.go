package todaTag

import (
	"github.com/todalist/app/internal/models/dto"
	"github.com/todalist/app/internal/models/entity"
	"github.com/todalist/app/internal/models/vo"
)

type ITodaTagRepo interface {

	// basic crud
	Get(uint) (*entity.TodaTag, error)

	First(*dto.TodaTagQuerier) (*entity.TodaTag, error)

	Save(*entity.TodaTag) (*entity.TodaTag, error)

	List(*dto.TodaTagQuerier) ([]*entity.TodaTag, error)

	ListUserTodaTag(*dto.ListUserTodaTagQuerier) ([]*vo.UserTodaTagVO, error)

	ListTodaTagByTodaIds(ids []uint) ([]*vo.TodaTagRefVO, error)

	Delete(uint) (uint, error)

}