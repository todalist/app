package sysImpl

import (
	"context"

	"github.com/gofiber/fiber/v3"
	"github.com/todalist/app/internal/api"
	"github.com/todalist/app/internal/globals"
	"github.com/todalist/app/internal/models/dto"
	"github.com/todalist/app/internal/mods/sys"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type SysRouteImpl struct {
	sysService sys.ISysService
}

func (r *SysRouteImpl) PasswordLogin(c fiber.Ctx) error {
	var form dto.PasswordLoginDTO
	if err := c.Bind().Body(&form); err != nil {
		globals.LOG.Error("password login bind error", zap.String("error", err.Error()))
		return fiber.ErrBadRequest
	}
	return api.Result(c).Or(globals.Transaction(func(tx *gorm.DB) (*string, error) {
		return r.sysService.PasswordLogin(globals.DbCtx(context.Background(), tx), &form)
	}))
}

func (r *SysRouteImpl) Register(root fiber.Router) {
	router := root.Group("/sys")
	router.Post("/authentication/passwordLogin", r.PasswordLogin)
}

func NewSysRoute(sysService sys.ISysService) sys.ISysRoute {
	return &SysRouteImpl{
		sysService: sysService,
	}
}
