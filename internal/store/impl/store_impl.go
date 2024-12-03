package storeImpl

import (
	"context"
	"dailydo.fe1.xyz/internal/mods/user"
)

type storeImpl struct {
}

func (*storeImpl) GetUserStore(ctx context.Context) user.IUserStore {
	// TODO
	return nil
}

var StoreImpl = &storeImpl{}
