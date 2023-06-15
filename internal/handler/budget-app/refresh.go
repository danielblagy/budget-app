package budget_app

import (
	"errors"

	"github.com/danielblagy/budget-app/internal/service/access"
	"github.com/gofiber/fiber/v2"
)

func (h handler) Refresh(c *fiber.Ctx) error {
	accessToken := c.Cookies(accessTokenCookieName)
	if len(accessToken) == 0 {
		return c.Status(fiber.StatusUnauthorized).SendString("user is not logged in")
	}
	refreshToken := c.Cookies(refreshTokenCookieName)
	if len(accessToken) == 0 {
		return c.Status(fiber.StatusUnauthorized).SendString("user is not logged in")
	}

	newTokenPair, err := h.accessService.Refresh(c.Context(), accessToken, refreshToken)
	if err != nil {
		if errors.Is(err, access.ErrNotAuthorized) {
			return c.Status(fiber.StatusUnauthorized).SendString(err.Error())
		}
		return err
	}

	c.Cookie(&fiber.Cookie{
		Name:     accessTokenCookieName,
		Value:    newTokenPair.AccessToken,
		HTTPOnly: true,
	})

	c.Cookie(&fiber.Cookie{
		Name:     refreshTokenCookieName,
		Value:    newTokenPair.RefreshToken,
		HTTPOnly: true,
	})

	return c.Status(fiber.StatusOK).SendString("successfully refreshed tokens")
}
