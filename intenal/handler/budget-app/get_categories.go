package budget_app

import (
	"github.com/danielblagy/budget-app/intenal/model"
	"github.com/gofiber/fiber/v2"
)

var validCategoryTypes = map[model.CategoryType]struct{}{
	model.CategoryTypeExpense: {},
	model.CategoryTypeIncome:  {},
}

func (h handler) GetCategories(c *fiber.Ctx) error {
	username, err := h.authorize(c)
	if err != nil {
		return err
	}

	categoryTypeRaw := c.Params("type")

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
