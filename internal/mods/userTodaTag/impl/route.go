package userTodaTagImpl

import (
	"dailydo.fe1.xyz/internal/mods/userTodaTag"
	"dailydo.fe1.xyz/internal/common"
	"dailydo.fe1.xyz/internal/globals"
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"context"
)

type UserTodaTagRouteImpl struct {
	userTodaTagService userTodaTag.IUserTodaTagService
}

func (r *UserTodaTagRouteImpl) Get(c fiber.Ctx) error {
	var querier common.BaseModel
	if err := c.Bind().URI(&querier); err != nil {
		globals.LOG.Error("userTodaTag get bind error", zap.String("error", err.Error()))
		return fiber.ErrBadRequest
	}
	if querier.Id < 1 {
		return fiber.ErrBadRequest
	}
	return c.JSON(common.Or(r.userTodaTagService.Get(context.Background(), querier.Id)))
}

func (r *UserTodaTagRouteImpl) First(c fiber.Ctx) error {
	var querier userTodaTag.UserTodaTagQuerier
	if err := c.Bind().Body(&querier); err != nil {
		globals.LOG.Error("userTodaTag first bind error", zap.String("error", err.Error()))
		return fiber.ErrBadRequest
	}
	return c.JSON(common.Or(r.userTodaTagService.First(context.Background(), &querier)))
}

func (r *UserTodaTagRouteImpl) Save(c fiber.Ctx) error {
	var form userTodaTag.UserTodaTag
	if err := c.Bind().Body(&form); err != nil {
		globals.LOG.Error("userTodaTag save bind error", zap.String("error", err.Error()))
		return fiber.ErrBadRequest
	}	
	var result *userTodaTag.UserTodaTag
	err := globals.DB.Transaction(func(tx *gorm.DB) error {
		save, err := r.userTodaTagService.Save(globals.ContextDB(context.Background(), tx), &form)
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

func (r *UserTodaTagRouteImpl) List(c fiber.Ctx) error {
	var querier userTodaTag.UserTodaTagQuerier
	if err := c.Bind().Body(&querier); err != nil {
		globals.LOG.Error("userTodaTag list bind error", zap.String("error", err.Error()))
		return fiber.ErrBadRequest
	}
	return c.JSON(common.Or(r.userTodaTagService.List(context.Background(), &querier)))
}

func (r *UserTodaTagRouteImpl) Delete(c fiber.Ctx) error {
	var querier common.BaseModel
	if err := c.Bind().URI(&querier); err != nil {
		globals.LOG.Error("userTodaTag delete bind error", zap.String("error", err.Error()))
		return fiber.ErrBadRequest
	}
	var result uint
	err := globals.DB.Transaction(func(tx *gorm.DB) error {
		id, err := r.userTodaTagService.Delete(globals.ContextDB(context.Background(), tx), querier.Id)
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

func (r *UserTodaTagRouteImpl) Register(root fiber.Router) {
	// router := root.Group("/userTodaTag")
	// router.Get("/:id", r.Get)
	// router.Post("/save", r.Save)
	// router.Post("/list", r.List)
	// router.Delete("/:id", r.Delete)
	// router.Post("/first", r.First)
}

func NewUserTodaTagRoute(userTodaTagService userTodaTag.IUserTodaTagService) userTodaTag.IUserTodaTagRoute {
	return &UserTodaTagRouteImpl{
		userTodaTagService: userTodaTagService,
	}
}