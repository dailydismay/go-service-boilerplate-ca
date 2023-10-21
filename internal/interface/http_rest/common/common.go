package common

import (
	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	Pattern() string
	Method() string
	Middleware() []fiber.Handler
	Handle(*fiber.Ctx) error
}
