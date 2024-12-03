package store

import (
	"context"
	"dailydo.fe1.xyz/internal/mods/user"
)

type IStore interface {

	GetUserStore(context.Context) user.IUserStore

}
