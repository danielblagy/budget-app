package budget_app

import (
	"errors"
	"log"

	"github.com/danielblagy/budget-app/internal/service/access"
	"github.com/gofiber/fiber/v2"
)

// authorize returns username, 200, nil on success,
// empty string, status code, error on failure
func (h handler) authorize(c *fiber.Ctx) (string, int, error) {
	accessToken := c.Cookies(accessTokenCookieName)
	if len(accessToken) == 0 {
		log.Println("user not logged in")
		return "", fiber.StatusUnauthorized, errors.New("user is not logged in")
	}

	username, err := h.accessService.Authorize(c.Context(), accessToken)
	if err != nil {
		if errors.Is(err, access.ErrNotAuthorized) {
			return "", fiber.StatusUnauthorized, err
		}
		return "", fiber.StatusInternalServerError, err
	}

	return username, fiber.StatusOK, nil
}
