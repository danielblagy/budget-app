package budget_app

import (
	"github.com/danielblagy/budget-app/intenal/service/access"
	"github.com/danielblagy/budget-app/intenal/service/categories"
	"github.com/danielblagy/budget-app/intenal/service/users"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	SetupRoutes()
	Greet(c *fiber.Ctx) error
	GetUsers(c *fiber.Ctx) error
	GetUser(c *fiber.Ctx) error
	CreateUser(c *fiber.Ctx) error
	//UpdateUser(c *fiber.Ctx) error
	//DeleteUser(c *fiber.Ctx) error
	Me(c *fiber.Ctx) error
	GetCategories(c *fiber.Ctx) error
}

type handler struct {
	validate *validator.Validate
	app      *fiber.App

	usersService      users.Service
	accessService     access.Service
	categoriesService categories.Service
}

func NewHandler(validate *validator.Validate, app *fiber.App, usersService users.Service, accessService access.Service, categoriesService categories.Service) Handler {
	return &handler{
		validate:          validate,
		app:               app,
		usersService:      usersService,
		accessService:     accessService,
		categoriesService: categoriesService,
	}
}
