package budget_app

import (
	"errors"

	"github.com/danielblagy/budget-app/intenal/service/users"
	"github.com/gofiber/fiber/v2"
)

func (h handler) GetUser(c *fiber.Ctx) error {
	username := c.Params("username")
	if len(username) == 0 {
		return c.Status(fiber.StatusBadRequest).SendString("username must not be empty")
	}
	if !containsOnlyValidCharacters(username) {
		return c.Status(fiber.StatusBadRequest).SendString("username may only contain letters, numbers, underscores, and dashes")
	}

	user, err := h.usersService.Get(c.Context(), username)
	if err != nil {
		if errors.Is(err, users.ErrUserNotFound) {
			return c.Status(fiber.StatusNotFound).SendString(err.Error())
		}
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(user)
}
