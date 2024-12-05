package todaImpl

import (
	"context"
	"github.com/gofiber/fiber/v3"
	"github.com/todalist/app/internal/common"
	"github.com/todalist/app/internal/globals"
	"github.com/todalist/app/internal/mods/toda"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type TodaRouteImpl struct {
	todaService toda.ITodaService
}

func (r *TodaRouteImpl) Get(c fiber.Ctx) error {
	var querier common.BaseModel
	if err := c.Bind().URI(&querier); err != nil {
		globals.LOG.Error("toda get bind error", zap.String("error", err.Error()))
		return fiber.ErrBadRequest
	}
	if querier.Id < 1 {
		return fiber.ErrBadRequest
	}
	return c.JSON(common.Or(r.todaService.Get(context.Background(), querier.Id)))
}

func (r *TodaRouteImpl) Save(c fiber.Ctx) error {
	var form toda.Toda
	if err := c.Bind().Body(&form); err != nil {
		globals.LOG.Error("toda save bind error", zap.String("error", err.Error()))
		return fiber.ErrBadRequest
	}
	var result *toda.Toda
	err := globals.DB.Transaction(func(tx *gorm.DB) error {
		save, err := r.todaService.Save(globals.ContextDB(globals.MustGetTokenUserContext(c), tx), &form)
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

func (r *TodaRouteImpl) List(c fiber.Ctx) error {
	var querier toda.TodaQuerier
	if err := c.Bind().Body(&querier); err != nil {
		globals.LOG.Error("toda list bind error", zap.String("error", err.Error()))
		return fiber.ErrBadRequest
	}
	return c.JSON(common.Or(r.todaService.List(context.Background(), &querier)))
}

func (r *TodaRouteImpl) Delete(c fiber.Ctx) error {
	var querier common.BaseModel
	if err := c.Bind().URI(&querier); err != nil {
		globals.LOG.Error("toda delete bind error", zap.String("error", err.Error()))
		return fiber.ErrBadRequest
	}
	var result uint
	err := globals.DB.Transaction(func(tx *gorm.DB) error {
		id, err := r.todaService.Delete(
			globals.ContextDB(globals.MustGetTokenUserContext(c), tx),
			querier.Id,
		)
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

func (r *TodaRouteImpl) Register(root fiber.Router) {
	router := root.Group("/toda")
	router.Get("/:id", r.Get)
	router.Post("/save", r.Save)
	router.Post("/list", r.List)
	router.Delete("/:id", r.Delete)
}

func NewTodaRoute(todaService toda.ITodaService) toda.ITodaRoute {
	return &TodaRouteImpl{
		todaService: todaService,
	}
}
