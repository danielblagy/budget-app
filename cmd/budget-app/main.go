package main

import (
	"context"
	"fmt"
	"log"
	"os"

	budget_app "github.com/danielblagy/budget-app/intenal/handler/budget-app"
	"github.com/danielblagy/budget-app/intenal/service/access"
	"github.com/danielblagy/budget-app/intenal/service/users"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

const envDatabaseUrl = "DATABASE_URL"

func main() {
	// load environment variables from .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	log.Println(os.Getenv(envDatabaseUrl))

	// connect to postgres database
	conn, err := pgx.Connect(context.Background(), os.Getenv(envDatabaseUrl))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	// validator

	validate := validator.New()

	// services

	usersService := users.NewService(conn)
	accessService := access.NewService(conn, usersService)

	// fiber app

	app := fiber.New()
	app.Use(logger.New())

	// handlers

	budgetAppHandler := budget_app.NewHandler(validate, app, usersService, accessService)
	budgetAppHandler.SetupRoutes()

	// start the app

	log.Fatal(app.Listen(":5000"))
}
