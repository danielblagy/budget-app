package budget_app

import (
	"errors"

	"github.com/danielblagy/budget-app/internal/model"
	"github.com/danielblagy/budget-app/internal/service/access"
	"github.com/gofiber/fiber/v2"
)

const accessTokenCookieName = "budget-app-access-token"
const refreshTokenCookieName = "budget-app-refresh-token"

func (h handler) LogIn(c *fiber.Ctx) error {
	var login model.Login
	if err := c.BodyParser(&login); err != nil {
		return err
	}

	if err := h.validate.Struct(login); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	userTokens, err := h.accessService.LogIn(c.Context(), &login)
	if err != nil {
		if errors.Is(err, access.ErrUserNotFound) {
			return c.Status(fiber.StatusNotFound).SendString(err.Error())
		}
		if errors.Is(err, access.ErrIncorrectPassword) {
			return c.Status(fiber.StatusForbidden).SendString(err.Error())
		}
		return err
	}

	c.Cookie(&fiber.Cookie{
		Name:     accessTokenCookieName,
		Value:    userTokens.AccessToken,
		HTTPOnly: true,
	})

	c.Cookie(&fiber.Cookie{
		Name:     refreshTokenCookieName,
		Value:    userTokens.RefreshToken,
		HTTPOnly: true,
	})

	return c.JSON(userTokens)
}
