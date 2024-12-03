package app

import (
	"dailydo.fe1.xyz/internal"
	"github.com/gofiber/fiber/v3"
)

// TODO : init routes
func instanceInitNow(app fiber.Router) {
	// iStore := storeImpl.StoreImpl
	// todaService := todaImpl.NewTodaService(iStore)
	// todaRoute := todaImpl.NewTodaRoute(todaService)
	// registerRoutes(app, []internal.IRoute{todaRoute})
}

func registerRoutes(app fiber.Router, routes []internal.IRoute) {
	for _, route := range routes {
		route.Register(app)
	}
}
