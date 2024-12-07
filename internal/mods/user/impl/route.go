package userImpl

import (
	"context"
	"github.com/gofiber/fiber/v3"
	"github.com/todalist/app/internal/api"
	"github.com/todalist/app/internal/common"
	"github.com/todalist/app/internal/globals"
	"github.com/todalist/app/internal/models/dto"
	"github.com/todalist/app/internal/models/entity"
	"github.com/todalist/app/internal/mods/user"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserRouteImpl struct {
	userService user.IUserService
}

func (r *UserRouteImpl) Get(c fiber.Ctx) error {
	var querier common.BaseModel
	if err := c.Bind().URI(&querier); err != nil {
		globals.LOG.Error("user get bind error", zap.String("error", err.Error()))
		return fiber.ErrBadRequest
	}
	if querier.Id < 1 {
		return fiber.ErrBadRequest
	}
	return api.Result(c).Or(r.userService.Get(context.Background(), querier.Id))
}

func (r *UserRouteImpl) First(c fiber.Ctx) error {
	var querier dto.UserQuerier
	if err := c.Bind().Body(&querier); err != nil {
		globals.LOG.Error("user first bind error", zap.String("error", err.Error()))
		return fiber.ErrBadRequest
	}
	return api.Result(c).Or(r.userService.First(context.Background(), &querier))
}

func (r *UserRouteImpl) Save(c fiber.Ctx) error {
	var form entity.User
	if err := c.Bind().Body(&form); err != nil {
		globals.LOG.Error("user save bind error", zap.String("error", err.Error()))
		return fiber.ErrBadRequest
	}
	var result *entity.User
	err := globals.DB.Transaction(func(tx *gorm.DB) error {
		save, err := r.userService.Save(globals.ContextDB(context.Background(), tx), &form)
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
	return api.Result(c).Ok(result)
}

func (r *UserRouteImpl) List(c fiber.Ctx) error {
	var querier dto.UserQuerier
	if err := c.Bind().Body(&querier); err != nil {
		globals.LOG.Error("user list bind error", zap.String("error", err.Error()))
		return fiber.ErrBadRequest
	}
	return api.Result(c).Or(r.userService.List(context.Background(), &querier))
}

func (r *UserRouteImpl) Delete(c fiber.Ctx) error {
	var querier common.BaseModel
	if err := c.Bind().URI(&querier); err != nil {
		globals.LOG.Error("user delete bind error", zap.String("error", err.Error()))
		return fiber.ErrBadRequest
	}
	var result uint
	err := globals.DB.Transaction(func(tx *gorm.DB) error {
		id, err := r.userService.Delete(globals.ContextDB(context.Background(), tx), querier.Id)
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
	return api.Result(c).Ok(result)
}

func (r *UserRouteImpl) Register(root fiber.Router) {
	// router := root.Group("/user")
	// router.Get("/:id", r.Get)
	// router.Post("/save", r.Save)
	// router.Post("/list", r.List)
	// router.Delete("/:id", r.Delete)
	// router.Post("/first", r.First)
}

func NewUserRoute(userService user.IUserService) user.IUserRoute {
	return &UserRouteImpl{
		userService: userService,
	}
}
