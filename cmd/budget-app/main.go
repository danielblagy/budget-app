package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/danielblagy/budget-app/internal/db"
	budget_app "github.com/danielblagy/budget-app/internal/handler/budget-app"
	"github.com/danielblagy/budget-app/internal/service/access"
	"github.com/danielblagy/budget-app/internal/service/cache"
	"github.com/danielblagy/budget-app/internal/service/categories"
	"github.com/danielblagy/budget-app/internal/service/users"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

const envDatabaseUrl = "DATABASE_URL"
const envRedisAddress = "REDIS_ADDRESS"
const envRedisPassword = "REDIS_PASSWORD"

func main() {
	// load environment variables from .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	log.Println(os.Getenv(envDatabaseUrl))

	ctx := context.Background()

	// connect to postgres database
	conn, err := pgx.Connect(ctx, os.Getenv(envDatabaseUrl))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(ctx)

	// connect to redis
	redisClient := redis.NewClient(&redis.Options{
		Addr:     os.Getenv(envRedisAddress),
		Password: os.Getenv(envRedisPassword),
		DB:       0, // use default DB
	})

	err = redisClient.Ping(ctx).Err()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't ping redis: %v\n", err)
		os.Exit(1)
	}

	// validator

	validate := validator.New()

	// db queries

	categoriesQuery := db.NewCategoriesQuery(conn)

	// services

	cacheService := cache.NewService(redisClient)
	usersService := users.NewService(conn)
	accessService := access.NewService(usersService)
	categoriesService := categories.NewService(categoriesQuery, cacheService)

	// fiber app

	app := fiber.New()
	app.Use(logger.New())

	// handlers

	budgetAppHandler := budget_app.NewHandler(validate, app, usersService, accessService, categoriesService)
	budgetAppHandler.SetupRoutes()

	// start the app

	log.Fatal(app.Listen(":5000"))
}
