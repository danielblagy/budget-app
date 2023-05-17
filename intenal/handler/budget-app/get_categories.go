package budget_app

import "github.com/gofiber/fiber/v2"

func (h handler) GetCategories(c *fiber.Ctx) error {
	username, err := h.authorize(c)
	if err != nil {
		return err
	}

	categories, err := h.categoriesService.GetAll(c.Context(), username)
	if err != nil {
		return err
	}

	return c.JSON(categories)
}
