package userTodaImpl

import (
	"github.com/gofiber/fiber/v3"
	"github.com/todalist/app/internal/common"
	"github.com/todalist/app/internal/globals"
	"github.com/todalist/app/internal/models/dto"
	"github.com/todalist/app/internal/models/entity"
	"github.com/todalist/app/internal/mods/userToda"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserTodaRouteImpl struct {
	userTodaService userToda.IUserTodaService
}

func (r *UserTodaRouteImpl) ListUserToda(c fiber.Ctx) error {
	var querier dto.ListUserTodaQuerier
	if err := c.Bind().Body(&querier); err != nil {
		globals.LOG.Error("entity list bind error", zap.String("error", err.Error()))
		return fiber.ErrBadRequest
	}
	return c.JSON(common.Or(r.userTodaService.ListUserToda(globals.MustGetTokenUserContext(c), &querier)))
}

func (r *UserTodaRouteImpl) CreateUserToda(c fiber.Ctx) error {
	var form entity.Toda
	if err := c.Bind().Body(&form); err != nil {
		globals.LOG.Error("entity list bind error", zap.String("error", err.Error()))
		return fiber.ErrBadRequest
	}
	return c.JSON(common.Or(globals.Transaction(
		func(tx *gorm.DB) (*entity.Toda, error) {
			return r.
				userTodaService.
				CreateUserToda(globals.ContextDB(
					globals.MustGetTokenUserContext(c), tx),
					&form,
				)
		},
	)))
}

func (r *UserTodaRouteImpl) Register(root fiber.Router) {
	router := root.Group("/userToda")
	router.Post("/list", r.ListUserToda)
}

func NewUserTodaRoute(userTodaService userToda.IUserTodaService) userToda.IUserTodaRoute {
	return &UserTodaRouteImpl{
		userTodaService: userTodaService,
	}
}
