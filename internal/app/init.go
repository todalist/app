package app

import (
	"dailydo.fe1.xyz/internal"
	todaImpl "dailydo.fe1.xyz/internal/mods/toda/impl"
	todaFlowImpl "dailydo.fe1.xyz/internal/mods/todaFlow/impl"
	todaTagImpl "dailydo.fe1.xyz/internal/mods/todaTag/impl"
	userImpl "dailydo.fe1.xyz/internal/mods/user/impl"
	userTodaImpl "dailydo.fe1.xyz/internal/mods/userToda/impl"
	userTodaTagImpl "dailydo.fe1.xyz/internal/mods/userTodaTag/impl"
	repoImpl "dailydo.fe1.xyz/internal/repo/impl"
	"github.com/gofiber/fiber/v3"
)

// TODO : init routes
func instanceInitNow(app fiber.Router) {
	repo := repoImpl.RepoImpl
	// services
	userService := userImpl.NewUserService(repo)
	userTodaService := userTodaImpl.NewUserTodaService(repo)
	userTodaTagService := userTodaTagImpl.NewUserTodaTagService(repo)
	todaService := todaImpl.NewTodaService(repo)
	todaFlowService := todaFlowImpl.NewTodaFlowService(repo)
	todaTagService := todaTagImpl.NewTodaTagService(repo)

	// routes
	userRoute := userImpl.NewUserRoute(userService)
	userTodaRoute := userTodaImpl.NewUserTodaRoute(userTodaService)
	userTodaTagRoute := userTodaTagImpl.NewUserTodaTagRoute(userTodaTagService)
	todaRoute := todaImpl.NewTodaRoute(todaService)
	todaFlowRoute := todaFlowImpl.NewTodaFlowRoute(todaFlowService)
	todaTagRoute := todaTagImpl.NewTodaTagRoute(todaTagService)

	registerRoutes(app, []internal.IRoute{
		userRoute,
		userTodaRoute,
		userTodaTagRoute,
		todaRoute,
		todaFlowRoute,
		todaTagRoute,
	})
}

func registerRoutes(app fiber.Router, routes []internal.IRoute) {
	for _, route := range routes {
		route.Register(app)
	}
}
