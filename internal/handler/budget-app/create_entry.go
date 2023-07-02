package budget_app

import (
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

	if err := h.validate.Struct(entry); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	createdEntry, err := h.entriesService.Create(c.Context(), username, &entry)
	if err != nil {
		return err
	}

	return c.JSON(createdEntry)
}
