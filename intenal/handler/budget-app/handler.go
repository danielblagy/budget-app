package budget_app

import (
	"github.com/danielblagy/budget-app/intenal/service/users"
	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	SetupRoutes()
	Greet(c *fiber.Ctx) error
	GetUsers(c *fiber.Ctx) error
	GetUser(c *fiber.Ctx) error
	//CreateUser(c *fiber.Ctx) error
	//UpdateUser(c *fiber.Ctx) error
	//DeleteUser(c *fiber.Ctx) error
}

type handler struct {
	app *fiber.App

	usersService users.Service
}

func NewHandler(app *fiber.App, usersService users.Service) Handler {
	return &handler{
		app:          app,
		usersService: usersService,
	}
}
