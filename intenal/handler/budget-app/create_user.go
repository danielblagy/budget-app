package budget_app

import (
	"github.com/danielblagy/budget-app/intenal/model"
	"github.com/gofiber/fiber/v2"
)

func (h handler) CreateUser(c *fiber.Ctx) error {
	var user model.User
	if err := c.BodyParser(&user); err != nil {
		return err
	}

	if err := h.validate.Struct(user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	createdUser, err := h.usersService.CreateUser(c.Context(), &user)
	if err != nil {
		return err
	}

	return c.JSON(createdUser)
}
