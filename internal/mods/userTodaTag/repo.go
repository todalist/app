package userTodaTag

import (
	"github.com/todalist/app/internal/models/dto"
	"github.com/todalist/app/internal/models/entity"
)

type IUserTodaTagRepo interface {

	// basic crud
	Get(uint) (*entity.UserTodaTag, error)

	First(*dto.UserTodaTagQuerier) (*entity.UserTodaTag, error)

	Save(*entity.UserTodaTag) (*entity.UserTodaTag, error)

	List(*dto.UserTodaTagQuerier) ([]*entity.UserTodaTag, error)

	Delete(uint) (uint, error)

}