package budget_app

import (
	"errors"
	"strconv"

	"github.com/danielblagy/budget-app/internal/service/categories"
	"github.com/gofiber/fiber/v2"
)

func (h handler) DeleteCategory(c *fiber.Ctx) error {
	username, statusCode, err := h.authorize(c)
	if err != nil {
		return c.Status(statusCode).SendString(err.Error())
	}

	categoryID, err := c.ParamsInt("id", -1)
	if err != nil {
		return err
	}
	h.logger.Debug("delete category", "username", username, "category_id", categoryID)
	if categoryID <= 0 {
		return c.Status(fiber.StatusBadRequest).SendString("category id is not valid")
	}

	deleteEntriesStrValue := c.Params("delete_entries")
	if len(deleteEntriesStrValue) == 0 {
		return c.Status(fiber.StatusBadRequest).SendString("delete_entries parameter must be provided")
	}
	deleteEntries, err := strconv.ParseBool(deleteEntriesStrValue)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("delete_entries has to be a boolean value, either true or false")
	}

	deletedCategory, err := h.categoriesService.Delete(c.Context(), username, int64(categoryID), deleteEntries)
	if err != nil {
		if errors.Is(err, categories.ErrNotFound) {
			return c.Status(fiber.StatusNotFound).SendString(err.Error())
		}
		return err
	}

	return c.JSON(deletedCategory)
}
