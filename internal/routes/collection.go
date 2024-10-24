package routes

import (
	"dailydo.fe1.xyz/internal/common"
	"dailydo.fe1.xyz/internal/globals"
	"dailydo.fe1.xyz/internal/models"
	"dailydo.fe1.xyz/internal/services"
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
)

type CollectionService struct{}

var (
	collectionService services.CollectionService
)

func (*CollectionService) Get(c fiber.Ctx) error {
	tokenUser := globals.MustGetTokenUser(c)
	querier := new(models.CollectionQuerier)
	if err := c.Bind().URI(querier); err != nil {
		globals.LOG.Debug("Collection get bind error", zap.String("error", err.Error()))
		return fiber.ErrBadRequest
	}
	if querier.ID == nil {
		return fiber.ErrBadRequest
	}
	querier.UserID = &tokenUser.UserID
	return c.JSON(common.Or(collectionService.Get(querier)))
}

func (*CollectionService) Save(c fiber.Ctx) error {
	tokenUser := globals.MustGetTokenUser(c)
	form := new(models.Collection)
	if err := c.Bind().Body(form); err != nil {
		globals.LOG.Debug("Collection save bind error", zap.String("error", err.Error()))
		return fiber.ErrBadRequest
	}
	if form.ID > 0 && form.UserID != tokenUser.UserID {
		globals.LOG.Error("Collection save user id not match")
		return fiber.ErrBadRequest
	}
	form.UserID = tokenUser.UserID
	return c.JSON(common.Or(collectionService.Save(form)))
}

func (*CollectionService) List(c fiber.Ctx) error {
	tokenUser := globals.MustGetTokenUser(c)
	querier := new(models.CollectionQuerier)
	if err := c.Bind().Body(querier); err != nil {
		globals.LOG.Debug("Collection list bind error", zap.String("error", err.Error()))
		return fiber.ErrBadRequest
	}
	querier.UserID = &tokenUser.UserID
	return c.JSON(common.Or(collectionService.List(querier)))
}

func (*CollectionService) Delete(c fiber.Ctx) error {
	tokenUser := globals.MustGetTokenUser(c)
	querier := new(models.CollectionQuerier)
	if err := c.Bind().URI(querier); err != nil {
		globals.LOG.Debug("Collection delete bind error", zap.String("error", err.Error()))
		return fiber.ErrBadRequest
	}
	if querier.ID == nil {
		return fiber.ErrBadRequest
	}
	querier.UserID = &tokenUser.UserID
	return c.JSON(common.Or(collectionService.Delete(querier)))
}

func (c *CollectionService) Register(root fiber.Router) {
	router := root.Group("/collection")
	router.Get("/:id", c.Get)
	router.Post("/", c.Save)
	router.Post("/list", c.List)
	router.Delete("/:id", c.Delete)
}
