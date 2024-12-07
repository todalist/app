package mods

import "github.com/gofiber/fiber/v3"


type IRoute interface {
	Register(fiber.Router)
}
