package routes

import (
	"dailydo.fe1.xyz/internal/common"
	"dailydo.fe1.xyz/internal/globals"
	"dailydo.fe1.xyz/internal/models"
	"dailydo.fe1.xyz/internal/services"
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
)

type TodoRoute struct{}

var (
	todoService services.TodoService
)

func (*TodoRoute) Get(c fiber.Ctx) error {
	tokenUser := globals.MustGetTokenUser(c)
	querier := new(models.TodoQuerier)
	if err := c.Bind().URI(querier); err != nil {
		globals.LOG.Debug("todo get bind error", zap.String("error", err.Error()))
		return fiber.ErrBadRequest
	}
	if querier.ID == nil {
		return fiber.ErrBadRequest
	}
	querier.UserID = &tokenUser.UserID
	return c.JSON(common.Or(todoService.Get(querier)))
}

func (*TodoRoute) Save(c fiber.Ctx) error {
	tokenUser := globals.MustGetTokenUser(c)
	form := new(models.Todo)
	if err := c.Bind().Body(form); err != nil {
		globals.LOG.Debug("todo save bind error", zap.String("error", err.Error()))
		return fiber.ErrBadRequest
	}
	if form.ID > 0 && form.UserID != tokenUser.UserID {
		globals.LOG.Error("todo save user id not match")
		return fiber.ErrBadRequest
	}
	form.UserID = tokenUser.UserID
	return c.JSON(common.Or(todoService.Save(form)))
}

func (*TodoRoute) List(c fiber.Ctx) error {
	tokenUser := globals.MustGetTokenUser(c)
	querier := new(models.TodoQuerier)
	if err := c.Bind().Body(querier); err != nil {
		globals.LOG.Debug("todo list bind error", zap.String("error", err.Error()))
		return fiber.ErrBadRequest
	}
	querier.UserID = &tokenUser.UserID
	return c.JSON(common.Or(todoService.List(querier)))
}

func (*TodoRoute) Delete(c fiber.Ctx) error {
	tokenUser := globals.MustGetTokenUser(c)
	querier := new(models.TodoQuerier)
	if err := c.Bind().URI(querier); err != nil {
		globals.LOG.Debug("todo delete bind error", zap.String("error", err.Error()))
		return fiber.ErrBadRequest
	}
	if querier.ID == nil {
		return fiber.ErrBadRequest
	}
	querier.UserID = &tokenUser.UserID
	return c.JSON(common.Or(todoService.Delete(querier)))
}

func (r *TodoRoute) Register(root fiber.Router) {
	router := root.Group("/todo")
	router.Get("/:id", r.Get)
	router.Post("/", r.Save)
	router.Post("/list", r.List)
	router.Delete("/:id", r.Delete)
}
