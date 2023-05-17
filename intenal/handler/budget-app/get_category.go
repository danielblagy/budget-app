package budget_app

import (
	"errors"

	"github.com/danielblagy/budget-app/intenal/service/categories"
	"github.com/gofiber/fiber/v2"
)

func (h handler) GetCategory(c *fiber.Ctx) error {
	username, err := h.authorize(c)
	if err != nil {
		return err
	}

	categoryID, err := c.ParamsInt("id", -1)
	if err != nil {
		return err
	}
	if categoryID <= 0 {
		return c.Status(fiber.StatusBadRequest).SendString("category id is not valid")
	}

	category, err := h.categoriesService.Get(c.Context(), int64(categoryID))
	if err != nil {
		if errors.Is(err, categories.ErrNotFound) {
			return c.Status(fiber.StatusNotFound).SendString(err.Error())
		}
	}

	// check if category belongs to the user, if not, send Not Found
	if category.UserID != username {
		return c.Status(fiber.StatusNotFound).SendString("category is not found")
	}

	return c.JSON(category)
}
