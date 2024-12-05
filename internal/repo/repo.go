package repo

import (
	"context"

	"github.com/todalist/app/internal/mods/toda"
	"github.com/todalist/app/internal/mods/todaFlow"
	"github.com/todalist/app/internal/mods/todaTag"
	"github.com/todalist/app/internal/mods/todaTagRef"
	"github.com/todalist/app/internal/mods/user"
	"github.com/todalist/app/internal/mods/userToda"
	"github.com/todalist/app/internal/mods/userTodaTag"
)

type IRepo interface {
	GetUserRepo(context.Context) user.IUserRepo

	GetTodaRepo(context.Context) toda.ITodaRepo

	GetTodaTagRepo(context.Context) todaTag.ITodaTagRepo

	GetTodaFlowRepo(context.Context) todaFlow.ITodaFlowRepo

	GetUserTodaRepo(context.Context) userToda.IUserTodaRepo

	GetUserTodaTagRepo(context.Context) userTodaTag.IUserTodaTagRepo

	GetTodaTagRefRepo(context.Context) todaTagRef.ITodaTagRefRepo
}
