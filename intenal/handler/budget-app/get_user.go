package budget_app

import "github.com/gofiber/fiber/v2"

func (h handler) GetUser(c *fiber.Ctx) error {
	userID, err := c.ParamsInt("user_id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	if userID <= 0 {
		return c.Status(fiber.StatusBadRequest).SendString("user_id is not valid")
	}

	user, err := h.usersService.GetUser(c.Context(), userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(user)
}
