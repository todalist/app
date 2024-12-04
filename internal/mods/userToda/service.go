package userToda

import (
	"context"
)

type IUserTodaService interface {

	// basic crud
	Get(context.Context, uint) (*UserToda, error)

	Save(context.Context, *UserToda) (*UserToda, error)

	List(context.Context, *UserTodaQuerier) ([]*UserToda, error)

	Delete(context.Context, uint) (uint, error)

}