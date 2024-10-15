package routes

import (
	"dailydo.fe1.xyz/internal/common"
	"dailydo.fe1.xyz/internal/globals"
	"dailydo.fe1.xyz/internal/models"
	"dailydo.fe1.xyz/internal/services"

	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
)

func UserInfoRoute(c fiber.Ctx) error {
	tokenUser := globals.MustGetTokenUser(c)
	user, err := services.GetUser(&models.User{
		BaseModel: models.BaseModel{ID: tokenUser.UserId},
	}, true)
	if err != nil {
		globals.LOG.Info("get user error", zap.Error(err))
		return fiber.ErrBadRequest
	}
	return c.JSON(common.Ok(user))
}

type RewritePasswordForm struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

func RewritePasswordRoute(c fiber.Ctx) error {
	tokenUser := globals.MustGetTokenUser(c)
	form := new(RewritePasswordForm)
	if err := c.Bind().Body(form); err != nil {
		return fiber.ErrBadRequest
	}
	err := services.RewritePassword(&models.User{
		BaseModel: models.BaseModel{ID: tokenUser.UserId}},
		form.OldPassword,
		form.NewPassword)
	return c.JSON(common.Or(true, err))
}

func RegisterUserRoutes(root fiber.Router) {
	router := root.Group("/user")
	router.Get("/info", UserInfoRoute)
	router.Post("/rewritePassword", RewritePasswordRoute)
}
