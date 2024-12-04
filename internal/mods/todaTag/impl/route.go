package todaTagImpl

import (
	"dailydo.fe1.xyz/internal/mods/todaTag"
	"dailydo.fe1.xyz/internal/common"
	"dailydo.fe1.xyz/internal/globals"
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"context"
)

type TodaTagRouteImpl struct {
	todaTagService todaTag.ITodaTagService
}

func (r *TodaTagRouteImpl) Get(c fiber.Ctx) error {
	var querier common.BaseModel
	if err := c.Bind().URI(&querier); err != nil {
		globals.LOG.Error("todaTag get bind error", zap.String("error", err.Error()))
		return fiber.ErrBadRequest
	}
	if querier.Id < 1 {
		return fiber.ErrBadRequest
	}
	return c.JSON(common.Or(r.todaTagService.Get(context.Background(), querier.Id)))
}

func (r *TodaTagRouteImpl) First(c fiber.Ctx) error {
	var querier todaTag.TodaTagQuerier
	if err := c.Bind().Body(&querier); err != nil {
		globals.LOG.Error("todaTag first bind error", zap.String("error", err.Error()))
		return fiber.ErrBadRequest
	}
	return c.JSON(common.Or(r.todaTagService.First(context.Background(), &querier)))
}

func (r *TodaTagRouteImpl) Save(c fiber.Ctx) error {
	var form todaTag.TodaTag
	if err := c.Bind().Body(&form); err != nil {
		globals.LOG.Error("todaTag save bind error", zap.String("error", err.Error()))
		return fiber.ErrBadRequest
	}	
	var result *todaTag.TodaTag
	err := globals.DB.Transaction(func(tx *gorm.DB) error {
		save, err := r.todaTagService.Save(globals.ContextDB(context.Background(), tx), &form)
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

func (r *TodaTagRouteImpl) List(c fiber.Ctx) error {
	var querier todaTag.TodaTagQuerier
	if err := c.Bind().Body(&querier); err != nil {
		globals.LOG.Error("todaTag list bind error", zap.String("error", err.Error()))
		return fiber.ErrBadRequest
	}
	return c.JSON(common.Or(r.todaTagService.List(context.Background(), &querier)))
}

func (r *TodaTagRouteImpl) Delete(c fiber.Ctx) error {
	var querier common.BaseModel
	if err := c.Bind().URI(&querier); err != nil {
		globals.LOG.Error("todaTag delete bind error", zap.String("error", err.Error()))
		return fiber.ErrBadRequest
	}
	var result uint
	err := globals.DB.Transaction(func(tx *gorm.DB) error {
		id, err := r.todaTagService.Delete(globals.ContextDB(context.Background(), tx), querier.Id)
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

func (r *TodaTagRouteImpl) Register(root fiber.Router) {
	router := root.Group("/todaTag")
	router.Get("/:id", r.Get)
	router.Post("/save", r.Save)
	router.Post("/list", r.List)
	router.Delete("/:id", r.Delete)
	router.Post("/first", r.First)
}

func NewTodaTagRoute(todaTagService todaTag.ITodaTagService) todaTag.ITodaTagRoute {
	return &TodaTagRouteImpl{
		todaTagService: todaTagService,
	}
}