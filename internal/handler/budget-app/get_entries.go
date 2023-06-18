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
	username, err := h.authorize(c)
	if err != nil {
		return err
	}

	entryTypeRaw := c.Params("type")

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
