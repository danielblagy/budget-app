package budget_app

import (
	"errors"
	"time"

	"github.com/danielblagy/budget-app/internal/model"
	"github.com/danielblagy/budget-app/internal/service/entries"
	"github.com/gofiber/fiber/v2"
)

func (h handler) UpdateEntry(c *fiber.Ctx) error {
	username, statusCode, err := h.authorize(c)
	if err != nil {
		return c.Status(statusCode).SendString(err.Error())
	}

	var updateData model.UpdateEntry
	if err := c.BodyParser(&updateData); err != nil {
		return err
	}
	h.logger.Debug("update entry", "username", username, "entry", updateData)

	if err := h.validate.Struct(updateData); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())

	}

	entryID, err := c.ParamsInt("id", -1)
	if err != nil {
		return err
	}
	if entryID <= 0 {
		return c.Status(fiber.StatusBadRequest).SendString("entry id is not valid")
	}

	exists, _ := h.categoriesService.Exists(c.Context(), username, updateData.CategoryID)
	if !exists {
		return c.Status(fiber.StatusBadRequest).SendString("category not exists")
	}

	_, parseErr := time.Parse("2006-01-02", updateData.Date)
	if parseErr != nil {
		return c.Status(fiber.StatusBadRequest).SendString("date format not valid")
	}

	if updateData.Amount <= 0 {
		return c.Status(fiber.StatusBadRequest).SendString("amount less than or equal to zero")
	}

	entryType := model.EntryType(updateData.Type)
	if _, ok := validEntryTypes[entryType]; !ok {
		return c.Status(fiber.StatusBadRequest).SendString("entry type is not valid")
	}

	updatedEntry, err := h.entriesService.Update(c.Context(), username, int64(entryID), &updateData)
	if err != nil {
		if errors.Is(err, entries.ErrNotFound) {
			return c.Status(fiber.StatusNotFound).SendString(err.Error())
		}
		return err
	}
	return c.JSON(updatedEntry)
}
