package budget_app

import (
	"errors"
	"regexp"

	"github.com/danielblagy/budget-app/intenal/model"
	"github.com/gofiber/fiber/v2"
)

var containsOnlyValidCharacters = regexp.MustCompile(`^[A-Za-z0-9_-]*$`).MatchString

func (h handler) CreateUser(c *fiber.Ctx) error {
	var user model.User
	if err := c.BodyParser(&user); err != nil {
		return err
	}

	if err := h.createUserValidateBody(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	exists, err := h.usersService.Exists(c.Context(), user.Username)
	if err != nil {
		return err
	}
	if exists {
		return c.Status(fiber.StatusConflict).SendString("username has already been taken")
	}

	userWithEmailExists, err := h.usersService.UserWithEmailExists(c.Context(), user.Email)
	if err != nil {
		return err
	}
	if userWithEmailExists {
		return c.Status(fiber.StatusConflict).SendString("email is already in use by some user")
	}

	createdUser, err := h.usersService.CreateUser(c.Context(), &user)
	if err != nil {
		return err
	}

	return c.JSON(createdUser)
}

func (h handler) createUserValidateBody(user *model.User) error {
	if err := h.validate.Struct(user); err != nil {
		return err
	}

	if !containsOnlyValidCharacters(user.Username) {
		return errors.New("username may only contain letters, numbers, underscores, and dashes")
	}

	return nil
}
