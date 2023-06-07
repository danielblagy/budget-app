package budget_app

import (
	"github.com/gofiber/fiber/v2"
)

func (h handler) GetUsers(c *fiber.Ctx) error {
	users, err := h.usersService.GetAll(c.UserContext())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(users)
}
