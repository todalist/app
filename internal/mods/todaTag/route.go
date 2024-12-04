package todaTag

import (
	"github.com/gofiber/fiber/v3"
)

type ITodaTagRoute interface {

	// basic crud
	Get(fiber.Ctx) error

	Save(fiber.Ctx) error

	List(fiber.Ctx) error

	Delete(fiber.Ctx) error

	Register(fiber.Router)

}