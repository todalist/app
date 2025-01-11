package user

import (
	"github.com/todalist/app/internal/models/dto"
	"github.com/todalist/app/internal/models/entity"
)

type IUserRepo interface {

	// basic crud
	Get(uint) (*entity.User, error)

	First(*dto.UserQuerier) (*entity.User, error)

	Save(*entity.User) (*entity.User, error)

	List(*dto.UserQuerier) ([]*entity.User, error)

	Delete(uint) (uint, error)

}