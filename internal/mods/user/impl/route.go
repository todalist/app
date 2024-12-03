package userImpl

import (
	"dailydo.fe1.xyz/internal/mods/user"
	"dailydo.fe1.xyz/internal/common"
	"dailydo.fe1.xyz/internal/globals"
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"context"
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
	return c.JSON(common.Or(r.userService.Get(context.Background(), querier.Id)))
}
func (r *UserRouteImpl) Save(c fiber.Ctx) error {
	var form user.User
	if err := c.Bind().Body(&form); err != nil {
		globals.LOG.Error("user save bind error", zap.String("error", err.Error()))
		return fiber.ErrBadRequest
	}	
	var result *user.User
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
	return c.JSON(common.Ok(result))
}
func (r *UserRouteImpl) List(c fiber.Ctx) error {
	var querier user.UserQuerier
	if err := c.Bind().Body(&querier); err != nil {
		globals.LOG.Error("user list bind error", zap.String("error", err.Error()))
		return fiber.ErrBadRequest
	}
	return c.JSON(common.Or(r.userService.List(context.Background(), &querier)))
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
	return c.JSON(common.Ok(result))
}