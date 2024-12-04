package repo

import (
	"context"

	"dailydo.fe1.xyz/internal/mods/toda"
	"dailydo.fe1.xyz/internal/mods/todaFlow"
	"dailydo.fe1.xyz/internal/mods/todaTag"
	"dailydo.fe1.xyz/internal/mods/user"
	"dailydo.fe1.xyz/internal/mods/userToda"
	"dailydo.fe1.xyz/internal/mods/userTodaTag"
)

type IRepo interface {
	GetUserRepo(context.Context) user.IUserRepo

	GetTodaRepo(context.Context) toda.ITodaRepo

	GetTodaTagRepo(context.Context) todaTag.ITodaTagRepo

	GetTodaFlowRepo(context.Context) todaFlow.ITodaFlowRepo

	GetUserTodaRepo(context.Context) userToda.IUserTodaRepo

	GetUserTodaTagRepo(context.Context) userTodaTag.IUserTodaTagRepo
}
