package budget_app

import (
	"github.com/gofiber/fiber/v2"
)

func (h handler) Greet(c *fiber.Ctx) error {
	return c.SendString("Welcome to the budget-app!\ncheck out github.com/danielblagy/budget-app for the API documentation\n")
}
