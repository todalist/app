package userToda

import (
	"github.com/todalist/app/internal/models/dto"
	"github.com/todalist/app/internal/models/entity"
)

type IUserTodaRepo interface {

	// basic crud
	Get(uint) (*entity.UserToda, error)

	First(*dto.UserTodaQuerier) (*entity.UserToda, error)

	Save(*entity.UserToda) (*entity.UserToda, error)

	List(*dto.UserTodaQuerier) ([]*entity.UserToda, error)

	Delete(uint) (uint, error)
	
	DeleteByTodaId(uint) error

}