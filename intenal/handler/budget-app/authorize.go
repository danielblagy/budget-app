package budget_app

import (
	"errors"

	"github.com/danielblagy/budget-app/intenal/service/access"
	"github.com/gofiber/fiber/v2"
)

func (h handler) authorize(c *fiber.Ctx) (string, error) {
	accessToken := c.Cookies(accessTokenCookieName)
	if len(accessToken) == 0 {
		return "", c.Status(fiber.StatusBadRequest).SendString("user is not logged in")
	}

	username, err := h.accessService.Authorize(c.Context(), accessToken)
	if err != nil {
		if errors.Is(err, access.ErrNotAuthorized) {
			return "", c.Status(fiber.StatusUnauthorized).SendString(err.Error())
		}
		return "", err
	}

	return username, nil
}
