package budget_app

import (
	"errors"

	"github.com/danielblagy/budget-app/intenal/model"
	"github.com/danielblagy/budget-app/intenal/service/categories"
	"github.com/gofiber/fiber/v2"
)

func (h handler) UpdateCategory(c *fiber.Ctx) error {
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

	// check if category belongs to the user, if not, send Not Found
	oldCategory, err := h.categoriesService.Get(c.Context(), int64(categoryID))
	if err != nil {
		if errors.Is(err, categories.ErrNotFound) {
			return c.Status(fiber.StatusNotFound).SendString(err.Error())
		}
		return err
	}
	if oldCategory.UserID != username {
		return c.Status(fiber.StatusNotFound).SendString("category is not found")
	}

	var newCategory model.NewCategory
	if err := c.BodyParser(&newCategory); err != nil {
		return err
	}

	if err := h.validate.Struct(newCategory); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	updatedCategory, err := h.categoriesService.Update(c.Context(), int64(categoryID), &newCategory)
	if err != nil {
		if errors.Is(err, categories.ErrNotFound) {
			return c.Status(fiber.StatusNotFound).SendString(err.Error())
		}
		return err
	}

	return c.JSON(updatedCategory)
}
