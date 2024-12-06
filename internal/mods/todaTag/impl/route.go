package todaTagImpl

import (
	"context"

	"github.com/gofiber/fiber/v3"
	"github.com/todalist/app/internal/common"
	"github.com/todalist/app/internal/globals"
	"github.com/todalist/app/internal/models/dto"
	"github.com/todalist/app/internal/models/entity"
	"github.com/todalist/app/internal/mods/todaTag"
	"go.uber.org/zap"
	"gorm.io/gorm"
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
	var querier dto.TodaTagQuerier
	if err := c.Bind().Body(&querier); err != nil {
		globals.LOG.Error("todaTag first bind error", zap.String("error", err.Error()))
		return fiber.ErrBadRequest
	}
	return c.JSON(common.Or(r.todaTagService.First(context.Background(), &querier)))
}

func (r *TodaTagRouteImpl) Save(c fiber.Ctx) error {
	var form entity.TodaTag
	if err := c.Bind().Body(&form); err != nil {
		globals.LOG.Error("todaTag save bind error", zap.String("error", err.Error()))
		return fiber.ErrBadRequest
	}
	var result *entity.TodaTag
	err := globals.DB.Transaction(func(tx *gorm.DB) error {
		save, err := r.todaTagService.Save(globals.ContextDB(globals.MustGetTokenUserContext(c), tx), &form)
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
	var querier dto.TodaTagQuerier
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
		id, err := r.todaTagService.Delete(globals.ContextDB(globals.MustGetTokenUserContext(c), tx), querier.Id)
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
	// router := root.Group("/todaTag")
	// router.Get("/:id", r.Get)
	// router.Post("/save", r.Save)
	// router.Post("/list", r.List)
	// router.Delete("/:id", r.Delete)
	// router.Post("/first", r.First)
}

func NewTodaTagRoute(todaTagService todaTag.ITodaTagService) todaTag.ITodaTagRoute {
	return &TodaTagRouteImpl{
		todaTagService: todaTagService,
	}
}
