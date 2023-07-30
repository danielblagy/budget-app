package budget_app

import (
	"time"

	"github.com/danielblagy/budget-app/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

func (h handler) CreateEntry(c *fiber.Ctx) error {
	username, statusCode, err := h.authorize(c)
	if err != nil {
		return c.Status(statusCode).SendString(err.Error())
	}

	var entry model.CreateEntry
	if err := c.BodyParser(&entry); err != nil {
		return errors.Wrap(err, "can't parse body")
	}

	exists, _ := h.categoriesService.Exists(c.Context(), username, entry.CategoryID)
	if !exists {
		return c.Status(fiber.StatusBadRequest).SendString("category not exists")
	}

	_, parseErr := time.Parse("2006-01-02", entry.Date)
	if parseErr != nil {
		return c.Status(fiber.StatusBadRequest).SendString("date format not valid")
	}

	if entry.Amount <= 0 {
		return c.Status(fiber.StatusBadRequest).SendString("amount less than or equal to zero")
	}

	entryType := model.EntryType(entry.Type)
	if _, ok := validEntryTypes[entryType]; !ok {
		return c.Status(fiber.StatusBadRequest).SendString("entry type is not valid")
	}

	if err := h.validate.Struct(entry); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	createdEntry, err := h.entriesService.Create(c.Context(), username, &entry)
	if err != nil {
		return err
	}

	return c.JSON(createdEntry)
}
