package common

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	v, ok := AsHttpError(err)
	if !ok {
		v = ErrorBuilder(err).Status(http.StatusInternalServerError).Plain()
	}

	return sendError(c, v)
}

func sendError(c *fiber.Ctx, httpErr *HttpError) error {
	return c.Status(httpErr.Status).JSON(httpErr)
}
