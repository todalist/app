package app

import (
	"github.com/gofiber/fiber/v3"
	"github.com/todalist/app/internal"
	todaImpl "github.com/todalist/app/internal/mods/toda/impl"
	todaFlowImpl "github.com/todalist/app/internal/mods/todaFlow/impl"
	todaTagImpl "github.com/todalist/app/internal/mods/todaTag/impl"
	userImpl "github.com/todalist/app/internal/mods/user/impl"
	repoImpl "github.com/todalist/app/internal/repo/impl"
)

// TODO : init routes
func instanceInitNow(app fiber.Router) {
	repo := repoImpl.RepoImpl
	// services
	userService := userImpl.NewUserService(repo)
	todaService := todaImpl.NewTodaService(repo)
	todaFlowService := todaFlowImpl.NewTodaFlowService(repo)
	todaTagService := todaTagImpl.NewTodaTagService(repo)

	// routes
	userRoute := userImpl.NewUserRoute(userService)
	todaRoute := todaImpl.NewTodaRoute(todaService)
	todaFlowRoute := todaFlowImpl.NewTodaFlowRoute(todaFlowService)
	todaTagRoute := todaTagImpl.NewTodaTagRoute(todaTagService)

	registerRoutes(app, []internal.IRoute{
		userRoute,
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
