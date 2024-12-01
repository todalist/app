package routes

import (
	"dailydo.fe1.xyz/internal/common"
	"dailydo.fe1.xyz/internal/globals"
	"dailydo.fe1.xyz/internal/models"
	"dailydo.fe1.xyz/internal/services"
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
)


type TodoCatalogRoute struct {}


var (
	todoCatalogService services.TodoCatalogService
)

func (*TodoCatalogRoute) Get(c fiber.Ctx) error {
	tokenUser := globals.MustGetTokenUser(c)
	querier := new(models.TodoCatalogQuerier)
	if err := c.Bind().URI(querier); err != nil {
		globals.LOG.Debug("TodoCatalog get bind error", zap.String("error", err.Error()))
		return fiber.ErrBadRequest
	}
	if querier.ID == nil {
		return fiber.ErrBadRequest
	}
	querier.UserID = tokenUser.UserID
	return c.JSON(common.Or(todoCatalogService.Get(querier)))
}

func (*TodoCatalogRoute) Save(c fiber.Ctx) error {
	tokenUser := globals.MustGetTokenUser(c)
	form := new(models.TodoCatalog)
	if err := c.Bind().Body(form); err != nil {
		globals.LOG.Debug("TodoCatalog save bind error", zap.String("error", err.Error()))
		return fiber.ErrBadRequest
	}
	if form.ID > 0 && form.UserID != tokenUser.UserID {
		globals.LOG.Error("TodoCatalog save user id not match")
		return fiber.ErrBadRequest
	}
	form.UserID = tokenUser.UserID
	return c.JSON(common.Or(todoCatalogService.Save(form)))
}

func (*TodoCatalogRoute) List(c fiber.Ctx) error {
	tokenUser := globals.MustGetTokenUser(c)
	querier := new(models.TodoCatalogQuerier)
	if err := c.Bind().Body(querier); err != nil {
		globals.LOG.Debug("TodoCatalog list bind error", zap.String("error", err.Error()))
		return fiber.ErrBadRequest
	}
	querier.UserID = tokenUser.UserID
	return c.JSON(common.Or(todoCatalogService.List(querier)))
}

func (*TodoCatalogRoute) Delete(c fiber.Ctx) error {
	tokenUser := globals.MustGetTokenUser(c)
	querier := new(models.BaseModel)
	if err := c.Bind().URI(querier); err != nil {
		globals.LOG.Debug("TodoCatalog delete bind error", zap.String("error", err.Error()))
		return fiber.ErrBadRequest
	}
	if querier.ID < 1 {
		return fiber.ErrBadRequest
	}
	return c.JSON(common.Or(todoCatalogService.Delete(&querier.ID, &tokenUser.UserID)))
}

func (r *TodoCatalogRoute) Register(root fiber.Router) {
	router := root.Group("/todoCatalog")
	router.Get("/:id", r.Get)
	router.Post("/save", r.Save)
	router.Post("/list", r.List)
	router.Delete("/:id", r.Delete)
}