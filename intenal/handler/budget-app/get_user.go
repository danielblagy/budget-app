package budget_app

import (
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

	user, err := h.usersService.GetUser(c.Context(), username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(user)
}
