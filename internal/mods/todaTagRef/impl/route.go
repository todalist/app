package todaTagRefImpl

import (
	"context"
	"github.com/gofiber/fiber/v3"
	"github.com/todalist/app/internal/common"
	"github.com/todalist/app/internal/globals"
	"github.com/todalist/app/internal/models/dto"
	"github.com/todalist/app/internal/models/entity"
	"github.com/todalist/app/internal/mods/todaTagRef"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type TodaTagRefRouteImpl struct {
	todaTagRefService todaTagRef.ITodaTagRefService
}

func (r *TodaTagRefRouteImpl) Get(c fiber.Ctx) error {
	var querier common.BaseModel
	if err := c.Bind().URI(&querier); err != nil {
		globals.LOG.Error("todaTagRef get bind error", zap.String("error", err.Error()))
		return fiber.ErrBadRequest
	}
	if querier.Id < 1 {
		return fiber.ErrBadRequest
	}
	return c.JSON(common.Or(r.todaTagRefService.Get(context.Background(), querier.Id)))
}

func (r *TodaTagRefRouteImpl) First(c fiber.Ctx) error {
	var querier dto.TodaTagRefQuerier
	if err := c.Bind().Body(&querier); err != nil {
		globals.LOG.Error("todaTagRef first bind error", zap.String("error", err.Error()))
		return fiber.ErrBadRequest
	}
	return c.JSON(common.Or(r.todaTagRefService.First(context.Background(), &querier)))
}

func (r *TodaTagRefRouteImpl) Save(c fiber.Ctx) error {
	var form entity.TodaTagRef
	if err := c.Bind().Body(&form); err != nil {
		globals.LOG.Error("todaTagRef save bind error", zap.String("error", err.Error()))
		return fiber.ErrBadRequest
	}	
	var result *entity.TodaTagRef
	err := globals.DB.Transaction(func(tx *gorm.DB) error {
		save, err := r.todaTagRefService.Save(globals.ContextDB(context.Background(), tx), &form)
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

func (r *TodaTagRefRouteImpl) List(c fiber.Ctx) error {
	var querier dto.TodaTagRefQuerier
	if err := c.Bind().Body(&querier); err != nil {
		globals.LOG.Error("todaTagRef list bind error", zap.String("error", err.Error()))
		return fiber.ErrBadRequest
	}
	return c.JSON(common.Or(r.todaTagRefService.List(context.Background(), &querier)))
}

func (r *TodaTagRefRouteImpl) Delete(c fiber.Ctx) error {
	var querier common.BaseModel
	if err := c.Bind().URI(&querier); err != nil {
		globals.LOG.Error("todaTagRef delete bind error", zap.String("error", err.Error()))
		return fiber.ErrBadRequest
	}
	var result uint
	err := globals.DB.Transaction(func(tx *gorm.DB) error {
		id, err := r.todaTagRefService.Delete(globals.ContextDB(context.Background(), tx), querier.Id)
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

func (r *TodaTagRefRouteImpl) Register(root fiber.Router) {
	// router := root.Group("/todaTagRef")
	// router.Get("/:id", r.Get)
	// router.Post("/save", r.Save)
	// router.Post("/list", r.List)
	// router.Delete("/:id", r.Delete)
	// router.Post("/first", r.First)
}

func NewTodaTagRefRoute(todaTagRefService todaTagRef.ITodaTagRefService) todaTagRef.ITodaTagRefRoute {
	return &TodaTagRefRouteImpl{
		todaTagRefService: todaTagRefService,
	}
}