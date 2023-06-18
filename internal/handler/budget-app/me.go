package budget_app

import (
	"github.com/gofiber/fiber/v2"
)

func (h handler) Me(c *fiber.Ctx) error {
	username, statusCode, err := h.authorize(c)
	if err != nil {
		return c.Status(statusCode).SendString(err.Error())
	}

	user, err := h.usersService.Get(c.Context(), username)
	if err != nil {
		// not checking for ErrNotFound because the user must exist if authorization was successful
		return err
	}

	return c.JSON(user)
}
