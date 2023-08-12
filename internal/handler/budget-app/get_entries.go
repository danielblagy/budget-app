package budget_app

import (
	"github.com/danielblagy/budget-app/internal/model"
	"github.com/gofiber/fiber/v2"
)

var validEntryTypes = map[model.EntryType]struct{}{
	model.EntryTypeExpense: {},
	model.EntryTypeIncome:  {},
}

func (h handler) GetEntries(c *fiber.Ctx) error {
	username, statusCode, err := h.authorize(c)
	if err != nil {
		return c.Status(statusCode).SendString(err.Error())
	}

	entryTypeRaw := c.Params("type")
	h.logger.Debug("get entries", "username", username, "type", entryTypeRaw)

	entryType := model.EntryType(entryTypeRaw)
	if _, ok := validEntryTypes[entryType]; !ok {
		return c.Status(fiber.StatusBadRequest).SendString("entry type is not valid")
	}

	entries, err := h.entriesService.GetAll(c.Context(), username, entryType)
	if err != nil {
		return err
	}

	return c.JSON(entries)
}
