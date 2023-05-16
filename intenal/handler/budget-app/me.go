package budget_app

import (
	"errors"

	"github.com/danielblagy/budget-app/intenal/service/access"
	"github.com/gofiber/fiber/v2"
)

func (h handler) Me(c *fiber.Ctx) error {
	accessToken := c.Cookies(accessTokenCookieName)
	if len(accessToken) == 0 {
		return c.Status(fiber.StatusBadRequest).SendString("user is not logged in")
	}

	username, err := h.accessService.Authorize(c.Context(), accessToken)
	if err != nil {
		if errors.Is(err, access.ErrNotAuthorized) {
			return c.Status(fiber.StatusUnauthorized).SendString(err.Error())
		}
		return err
	}

	user, err := h.usersService.GetUser(c.Context(), username)
	if err != nil {
		// not checking for ErrNotFound because the user must exist is the authorization was successful
		return err
	}

	return c.JSON(user)
}
