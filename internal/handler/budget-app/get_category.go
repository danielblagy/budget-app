package budget_app

import (
	"errors"

	"github.com/danielblagy/budget-app/internal/service/categories"
	"github.com/gofiber/fiber/v2"
)

func (h handler) GetCategory(c *fiber.Ctx) error {
	username, statusCode, err := h.authorize(c)
	if err != nil {
		return c.Status(statusCode).SendString(err.Error())
	}

	categoryID, err := c.ParamsInt("id", -1)
	if err != nil {
		return err
	}
	h.logger.Debug("get category", "username", username, "category_id", categoryID)
	if categoryID <= 0 {
		return c.Status(fiber.StatusBadRequest).SendString("category id is not valid")
	}

	category, err := h.categoriesService.Get(c.Context(), username, int64(categoryID))
	if err != nil {
		if errors.Is(err, categories.ErrNotFound) {
			return c.Status(fiber.StatusNotFound).SendString(err.Error())
		}
		return err
	}

	return c.JSON(category)
}
