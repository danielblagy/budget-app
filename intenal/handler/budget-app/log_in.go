package budget_app

import (
	"errors"

	"github.com/danielblagy/budget-app/intenal/model"
	"github.com/danielblagy/budget-app/intenal/service/access"
	"github.com/gofiber/fiber/v2"
)

func (h handler) LogIn(c *fiber.Ctx) error {
	var login model.Login
	if err := c.BodyParser(&login); err != nil {
		return err
	}

	if err := h.validate.Struct(login); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	err := h.accessService.LogIn(c.Context(), &login)
	if err != nil {
		if errors.Is(err, access.ErrUserNotFound) {
			return c.Status(fiber.StatusNotFound).SendString(err.Error())
		}
		if errors.Is(err, access.ErrIncorrectPassword) {
			return c.Status(fiber.StatusForbidden).SendString(err.Error())
		}
		return err
	}

	// TODO access service will return jwt token that will be sent as a response
	return c.SendString("access granted!")
}
