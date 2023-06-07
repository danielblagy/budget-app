package budget_app

import (
	"errors"

	"github.com/danielblagy/budget-app/internal/model"
	"github.com/danielblagy/budget-app/internal/service/categories"
	"github.com/gofiber/fiber/v2"
)

func (h handler) UpdateCategory(c *fiber.Ctx) error {
	username, err := h.authorize(c)
	if err != nil {
		return err
	}

	var updateData model.UpdateCategory
	if err := c.BodyParser(&updateData); err != nil {
		return err
	}

	if err := h.validate.Struct(updateData); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	updatedCategory, err := h.categoriesService.Update(c.Context(), username, &updateData)
	if err != nil {
		if errors.Is(err, categories.ErrNotFound) {
			return c.Status(fiber.StatusNotFound).SendString(err.Error())
		}
		return err
	}

	return c.JSON(updatedCategory)
}
