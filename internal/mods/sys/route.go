package sys

import (
	"github.com/gofiber/fiber/v3"
)

type ISysRoute interface {

	PasswordLogin(fiber.Ctx) error

	Register(fiber.Router)

}
