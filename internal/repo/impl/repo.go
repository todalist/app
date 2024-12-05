package repoImpl

import (
	"context"

	"github.com/todalist/app/internal/globals"
	"github.com/todalist/app/internal/mods/toda"
	todaImpl "github.com/todalist/app/internal/mods/toda/impl"
	"github.com/todalist/app/internal/mods/todaFlow"
	todaFlowImpl "github.com/todalist/app/internal/mods/todaFlow/impl"
	"github.com/todalist/app/internal/mods/todaTag"
	todaTagImpl "github.com/todalist/app/internal/mods/todaTag/impl"
	"github.com/todalist/app/internal/mods/user"
	userImpl "github.com/todalist/app/internal/mods/user/impl"
	"github.com/todalist/app/internal/mods/userToda"
	userTodaImpl "github.com/todalist/app/internal/mods/userToda/impl"
	"github.com/todalist/app/internal/mods/userTodaTag"
	userTodaTagImpl "github.com/todalist/app/internal/mods/userTodaTag/impl"
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
