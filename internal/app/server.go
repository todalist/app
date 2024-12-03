package app

import (
	"dailydo.fe1.xyz/internal/common"
	"dailydo.fe1.xyz/internal/globals"
	"dailydo.fe1.xyz/internal/middlewares"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	fiberLoggerM "github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/fiber/v3/middleware/requestid"
	"go.uber.org/zap"
)

func NewServer(conf *globals.AppConfig) *fiber.App {
	app := fiber.New(fiber.Config{
		Immutable: true,
		ErrorHandler: func(c fiber.Ctx, err error) error {
			// Retrieve the custom status code if it's a *fiber.Error
			var e *fiber.Error
			if errors.As(err, &e) {
				return c.Status(e.Code).JSON(common.ApiResponse{Message: e.Message, Code: e.Code})
			}
			var apiE *common.ApiResponse
			if errors.As(err, &apiE) {
				return c.Status(apiE.Code).JSON(common.ApiResponse{Message: apiE.Message, Code: apiE.Code})
			}
			return c.Status(fiber.StatusInternalServerError).
				JSON(common.ApiResponse{Message: fiber.ErrInternalServerError.Message, Code: fiber.ErrInternalServerError.Code})
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
