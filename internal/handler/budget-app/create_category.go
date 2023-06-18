package budget_app

import (
	"github.com/danielblagy/budget-app/internal/model"
	"github.com/gofiber/fiber/v2"
)

func (h handler) CreateCategory(c *fiber.Ctx) error {
	username, statusCode, err := h.authorize(c)
	if err != nil {
		return c.Status(statusCode).SendString(err.Error())
	}

	// TODO now throws err if provided type is not valid
	var category model.CreateCategory
	if err := c.BodyParser(&category); err != nil {
		return err
	}

	if err := h.validate.Struct(category); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	createdCategory, err := h.categoriesService.Create(c.Context(), username, &category)
	if err != nil {
		return err
	}

	return c.JSON(createdCategory)
}
