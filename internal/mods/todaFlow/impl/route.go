package todaFlowImpl

import (
	"dailydo.fe1.xyz/internal/mods/todaFlow"
	"dailydo.fe1.xyz/internal/common"
	"dailydo.fe1.xyz/internal/globals"
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"context"
)

type TodaFlowRouteImpl struct {
	todaFlowService todaFlow.ITodaFlowService
}

func (r *TodaFlowRouteImpl) Get(c fiber.Ctx) error {
	var querier common.BaseModel
	if err := c.Bind().URI(&querier); err != nil {
		globals.LOG.Error("todaFlow get bind error", zap.String("error", err.Error()))
		return fiber.ErrBadRequest
	}
	if querier.Id < 1 {
		return fiber.ErrBadRequest
	}
	return c.JSON(common.Or(r.todaFlowService.Get(context.Background(), querier.Id)))
}

func (r *TodaFlowRouteImpl) Save(c fiber.Ctx) error {
	var form todaFlow.TodaFlow
	if err := c.Bind().Body(&form); err != nil {
		globals.LOG.Error("todaFlow save bind error", zap.String("error", err.Error()))
		return fiber.ErrBadRequest
	}	
	var result *todaFlow.TodaFlow
	err := globals.DB.Transaction(func(tx *gorm.DB) error {
		save, err := r.todaFlowService.Save(globals.ContextDB(context.Background(), tx), &form)
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

func (r *TodaFlowRouteImpl) List(c fiber.Ctx) error {
	var querier todaFlow.TodaFlowQuerier
	if err := c.Bind().Body(&querier); err != nil {
		globals.LOG.Error("todaFlow list bind error", zap.String("error", err.Error()))
		return fiber.ErrBadRequest
	}
	return c.JSON(common.Or(r.todaFlowService.List(context.Background(), &querier)))
}

func (r *TodaFlowRouteImpl) Delete(c fiber.Ctx) error {
	var querier common.BaseModel
	if err := c.Bind().URI(&querier); err != nil {
		globals.LOG.Error("todaFlow delete bind error", zap.String("error", err.Error()))
		return fiber.ErrBadRequest
	}
	var result uint
	err := globals.DB.Transaction(func(tx *gorm.DB) error {
		id, err := r.todaFlowService.Delete(globals.ContextDB(context.Background(), tx), querier.Id)
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

func (r *TodaFlowRouteImpl) Register(root fiber.Router) {
	router := root.Group("/todaFlow")
	router.Get("/:id", r.Get)
	router.Post("/save", r.Save)
	router.Post("/list", r.List)
	router.Delete("/:id", r.Delete)
}

func NewTodaFlowRoute(todaFlowService todaFlow.ITodaFlowService) todaFlow.ITodaFlowRoute {
	return &TodaFlowRouteImpl{
		todaFlowService: todaFlowService,
	}
}