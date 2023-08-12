package budget_app

import (
	"github.com/danielblagy/budget-app/internal/model"
	"github.com/gofiber/fiber/v2"
)

var validCategoryTypes = map[model.CategoryType]struct{}{
	model.CategoryTypeExpense: {},
	model.CategoryTypeIncome:  {},
}

func (h handler) GetCategories(c *fiber.Ctx) error {
	username, statusCode, err := h.authorize(c)
	if err != nil {
		return c.Status(statusCode).SendString(err.Error())
	}

	categoryTypeRaw := c.Params("type")
	h.logger.Debug("get categories", "username", username, "type", categoryTypeRaw)

	categoryType := model.CategoryType(categoryTypeRaw)
	if _, ok := validCategoryTypes[categoryType]; !ok {
		return c.Status(fiber.StatusBadRequest).SendString("category type is not valid")
	}

	categories, err := h.categoriesService.GetAll(c.Context(), username, categoryType)
	if err != nil {
		return err
	}

	return c.JSON(categories)
}
