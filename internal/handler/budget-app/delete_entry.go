package budget_app

import (
	"errors"

	"github.com/danielblagy/budget-app/internal/service/entries"
	"github.com/gofiber/fiber/v2"
)

func (h handler) DeleteEntry(c *fiber.Ctx) error {
	username, statusCode, err := h.authorize(c)
	if err != nil {
		return c.Status(statusCode).SendString(err.Error())
	}

	entryID, err := c.ParamsInt("id", -1)
	if err != nil {
		return err
	}
	h.logger.Debug("create entry", "username", username, "entry_id", entryID)
	if entryID <= 0 {
		return c.Status(fiber.StatusBadRequest).SendString("entry id is not valid")
	}

	deletedEntry, err := h.entriesService.Delete(c.Context(), username, int64(entryID))
	if err != nil {
		if errors.Is(err, entries.ErrNotFound) {
			return c.Status(fiber.StatusNotFound).SendString(err.Error())
		}
		return err
	}

	return c.JSON(deletedEntry)
}
