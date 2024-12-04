package repoImpl

import (
	"context"

	"dailydo.fe1.xyz/internal/globals"
	"dailydo.fe1.xyz/internal/mods/toda"
	todaImpl "dailydo.fe1.xyz/internal/mods/toda/impl"
	"dailydo.fe1.xyz/internal/mods/todaFlow"
	todaFlowImpl "dailydo.fe1.xyz/internal/mods/todaFlow/impl"
	"dailydo.fe1.xyz/internal/mods/todaTag"
	todaTagImpl "dailydo.fe1.xyz/internal/mods/todaTag/impl"
	"dailydo.fe1.xyz/internal/mods/user"
	userImpl "dailydo.fe1.xyz/internal/mods/user/impl"
	"dailydo.fe1.xyz/internal/mods/userToda"
	userTodaImpl "dailydo.fe1.xyz/internal/mods/userToda/impl"
	"dailydo.fe1.xyz/internal/mods/userTodaTag"
	userTodaTagImpl "dailydo.fe1.xyz/internal/mods/userTodaTag/impl"
)

type repoImpl struct {
}

func (*repoImpl) GetUserRepo(ctx context.Context) user.IUserRepo {
	tx := globals.GetFromContext(ctx)
	return userImpl.NewUserRepo(tx)
}

func (*repoImpl) GetTodaRepo(ctx context.Context) toda.ITodaRepo {
	tx := globals.GetFromContext(ctx)
	return todaImpl.NewTodaRepo(tx)
}

func (*repoImpl) GetTodaTagRepo(ctx context.Context) todaTag.ITodaTagRepo {
	tx := globals.GetFromContext(ctx)
	return todaTagImpl.NewTodaTagRepo(tx)
}

func (*repoImpl) GetTodaFlowRepo(ctx context.Context) todaFlow.ITodaFlowRepo {
	tx := globals.GetFromContext(ctx)
	return todaFlowImpl.NewTodaFlowRepo(tx)
}

func (*repoImpl) GetUserTodaRepo(ctx context.Context) userToda.IUserTodaRepo {
	tx := globals.GetFromContext(ctx)
	return userTodaImpl.NewUserTodaRepo(tx)
}

func (*repoImpl) GetUserTodaTagRepo(ctx context.Context) userTodaTag.IUserTodaTagRepo {
	tx := globals.GetFromContext(ctx)
	return userTodaTagImpl.NewUserTodaTagRepo(tx)
}

var RepoImpl = &repoImpl{}
