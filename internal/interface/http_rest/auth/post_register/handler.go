package post_register

import (
	"authsvc/internal/application/auth/register"
	"authsvc/internal/interface/http_rest/common"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

func (h *handler) Handle(c *fiber.Ctx) error {
	var body requestBody
	if err := c.BodyParser(&body); err != nil {
		return errors.Wrap(err, "failed to parse login body")
	}

	err := h.usecase.Execute(c.Context(), body.toUsecasePayload())
	if err != nil {
		return h.resolveErr(c, err)
	}

	return c.SendStatus(http.StatusCreated)
}

func (h *handler) resolveErr(c *fiber.Ctx, err error) error {
	switch {
	case errors.Is(err, register.ErrAccountAlreadyExists):
		return common.ErrorBuilder(err).Detail("username", "Account already exists").Build()
	}

	return err
}
