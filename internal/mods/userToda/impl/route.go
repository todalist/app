package userTodaImpl

import (
	"dailydo.fe1.xyz/internal/mods/userToda"
	"dailydo.fe1.xyz/internal/common"
	"dailydo.fe1.xyz/internal/globals"
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"context"
)

type UserTodaRouteImpl struct {
	userTodaService userToda.IUserTodaService
}

func (r *UserTodaRouteImpl) Get(c fiber.Ctx) error {
	var querier common.BaseModel
	if err := c.Bind().URI(&querier); err != nil {
		globals.LOG.Error("userToda get bind error", zap.String("error", err.Error()))
		return fiber.ErrBadRequest
	}
	if querier.Id < 1 {
		return fiber.ErrBadRequest
	}
	return c.JSON(common.Or(r.userTodaService.Get(context.Background(), querier.Id)))
}

func (r *UserTodaRouteImpl) Save(c fiber.Ctx) error {
	var form userToda.UserToda
	if err := c.Bind().Body(&form); err != nil {
		globals.LOG.Error("userToda save bind error", zap.String("error", err.Error()))
		return fiber.ErrBadRequest
	}	
	var result *userToda.UserToda
	err := globals.DB.Transaction(func(tx *gorm.DB) error {
		save, err := r.userTodaService.Save(globals.ContextDB(context.Background(), tx), &form)
		if err != nil {
			return err
		}
		result = save
		return nil
	})
	if err != nil {
		globals.LOG.Error("exec transaction error: ", zap.Error(err))
		return fiber.ErrInternalServerError
	}
	return c.JSON(common.Ok(result))
}

func (r *UserTodaRouteImpl) List(c fiber.Ctx) error {
	var querier userToda.UserTodaQuerier
	if err := c.Bind().Body(&querier); err != nil {
		globals.LOG.Error("userToda list bind error", zap.String("error", err.Error()))
		return fiber.ErrBadRequest
	}
	return c.JSON(common.Or(r.userTodaService.List(context.Background(), &querier)))
}

func (r *UserTodaRouteImpl) Delete(c fiber.Ctx) error {
	var querier common.BaseModel
	if err := c.Bind().URI(&querier); err != nil {
		globals.LOG.Error("userToda delete bind error", zap.String("error", err.Error()))
		return fiber.ErrBadRequest
	}
	var result uint
	err := globals.DB.Transaction(func(tx *gorm.DB) error {
		id, err := r.userTodaService.Delete(globals.ContextDB(context.Background(), tx), querier.Id)
		if err != nil {
			return err
		}
		result = id
		return nil
	})
	if err != nil {
		globals.LOG.Error("exec transaction error: ", zap.Error(err))
		return fiber.ErrInternalServerError
	}
	return c.JSON(common.Ok(result))
}

func (r *UserTodaRouteImpl) Register(root fiber.Router) {
	router := root.Group("/userToda")
	router.Get("/:id", r.Get)
	router.Post("/save", r.Save)
	router.Post("/list", r.List)
	router.Delete("/:id", r.Delete)
}

func NewUserTodaRoute(userTodaService userToda.IUserTodaService) userToda.IUserTodaRoute {
	return &UserTodaRouteImpl{
		userTodaService: userTodaService,
	}
}