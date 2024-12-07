package app

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	fiberLoggerM "github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/fiber/v3/middleware/requestid"
	"github.com/todalist/app/internal/api"
	"github.com/todalist/app/internal/common"
	"github.com/todalist/app/internal/globals"
	"github.com/todalist/app/internal/middlewares"
	"go.uber.org/zap"
)

func NewServer(conf *globals.AppConfig) *fiber.App {
	app := fiber.New(fiber.Config{
		Immutable: true,
		ErrorHandler: func(c fiber.Ctx, err error) error {
			return api.Result(c).Or(nil, err)
		},
		StructValidator: &common.StructValidator{V: validator.New()},
	})
	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))
	if conf.Server.Cors.Enable {
		app.Use(cors.New(cors.Config{
			AllowOrigins: conf.Server.Cors.Origins,
			AllowMethods: []string{"*"},
			AllowHeaders: []string{"*"},
		}))
	}
	app.Use(requestid.New())
	// logging
	app.Use(fiberLoggerM.New(fiberLoggerM.Config{
		LoggerFunc: func(c fiber.Ctx, data *fiberLoggerM.Data, cfg fiberLoggerM.Config) error {
			globals.LOG.Info("request handled",
				zap.ByteString("method", c.Request().Header.Method()),
				zap.ByteString("path", c.Request().RequestURI()),
				zap.Int("statusCode", c.Response().StatusCode()),
				zap.String("requestId", c.GetRespHeader(fiber.HeaderXRequestID)),
				zap.Duration("latency", data.Stop.Sub(data.Start)),
				zap.Error(data.ChainErr),
			)
			return nil
		},
	}))
	app.Use(middlewares.AuthenticationMiddleware(conf.Auth))
	// registry all routes
	// TODO refactor
	root := app.Group(conf.Server.PathPrefix)

	// init all instance
	instanceInitNow(root)
	return app
}
