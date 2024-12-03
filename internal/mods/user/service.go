package user

import (
	"context"
)

type IUserService interface {

	// basic crud
	Get(context.Context, uint) (*User, error)

	Save(context.Context, *User) (*User, error)

	List(context.Context, *UserQuerier) ([]*User, error)

	Delete(context.Context, uint) (uint, error)

}