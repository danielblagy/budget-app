package budget_app

import (
	"errors"

	"github.com/danielblagy/budget-app/internal/service/entries"
	"github.com/gofiber/fiber/v2"
)

func (h handler) GetEntryByID(c *fiber.Ctx) error {
	username, err := h.authorize(c)
	if err != nil {
		return err //return err in case user isn't authorized
	}

	entryID, err := c.ParamsInt("id", -1)
	if err != nil {
		return err //return err in case invalid types of data for entryID
	}
	if entryID <= 0 {
		return c.Status(fiber.StatusBadRequest).SendString("entry id is not valid")
	}

	entry, err := h.entriesService.GetByID(c.Context(), username, int64(entryID))
	if err != nil {
		if errors.Is(err, entries.ErrNotFound) {
			return c.Status(fiber.StatusNotFound).SendString(err.Error())
		}
		return err
	}

	return c.JSON(entry)
}
