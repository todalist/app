package userTodaTag

import (
	"github.com/todalist/app/internal/models/dto"
	"github.com/todalist/app/internal/models/entity"
	"github.com/todalist/app/internal/models/vo"
)

type IUserTodaTagRepo interface {

	// basic crud
	Get(uint) (*entity.UserTodaTag, error)

	First(*dto.UserTodaTagQuerier) (*entity.UserTodaTag, error)

	Save(*entity.UserTodaTag) (*entity.UserTodaTag, error)

	List(*dto.UserTodaTagQuerier) ([]*entity.UserTodaTag, error)

	ListUserTodaTag(*dto.ListUserTodaTagQuerier) ([]*vo.UserTodaTagVO, error)

	Delete(uint) (uint, error)

}