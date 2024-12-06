package userToda

import (
	"github.com/gofiber/fiber/v3"
)

type IUserTodaRoute interface {

	ListUserToda(fiber.Ctx) error

	CreateUserToda(fiber.Ctx) error

	Register(fiber.Router)

}