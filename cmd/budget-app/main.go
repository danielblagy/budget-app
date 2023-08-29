package main

import (
	"context"
	"fmt"
	"os"

	"github.com/danielblagy/budget-app/internal/db"
	budget_app "github.com/danielblagy/budget-app/internal/handler/budget-app"
	"github.com/danielblagy/budget-app/internal/service/access"
	"github.com/danielblagy/budget-app/internal/service/cache"
	"github.com/danielblagy/budget-app/internal/service/categories"
	"github.com/danielblagy/budget-app/internal/service/entries"
	persistent_store "github.com/danielblagy/budget-app/internal/service/persistent-store"
	"github.com/danielblagy/budget-app/internal/service/reports"
	"github.com/danielblagy/budget-app/internal/service/users"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	log "github.com/inconshreveable/log15"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

const envDatabaseUrl = "DATABASE_URL"
const envCacheAddress = "CACHE_ADDRESS"
const envCachePassword = "CACHE_PASSWORD"
const envPersistentStoreAddress = "PERSISTENT_STORE_ADDRESS"
const envPersistentStorePassword = "PERSISTENT_STORE_PASSWORD"

func main() {
	// logger

	logger := log.New()

	// load environment variables from .env file
	err := godotenv.Load(".env")
	if err != nil {
		logger.Crit("can't load .env file", "err", err.Error())
		os.Exit(1)
	}

	ctx := context.Background()

	// connect to postgres database
	conn, err := pgxpool.New(ctx, os.Getenv(envDatabaseUrl))
	if err != nil {
		logger.Crit("can't connect to database", "err", err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	// connect to redis cache server

	redisCacheClient := redis.NewClient(&redis.Options{
		Addr:     os.Getenv(envCacheAddress),
		Password: os.Getenv(envCachePassword),
		DB:       0, // use default DB
	})

	err = redisCacheClient.Ping(ctx).Err()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't ping redis cache server: %v\n", err)
		os.Exit(1)
	}

	// connect to redis persistent store server

	redisPersistentStoreClient := redis.NewClient(&redis.Options{
		Addr:     os.Getenv(envPersistentStoreAddress),
		Password: os.Getenv(envPersistentStorePassword),
		DB:       0, // use default DB
	})

	err = redisPersistentStoreClient.Ping(ctx).Err()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't ping redis persistent store server: %v\n", err)
		os.Exit(1)
	}

	// validator

	validate := validator.New()

	// db queries

	queryFactory := db.NewQueryFactory(conn)

	categoriesQuery := queryFactory.NewCategoriesQuery(conn)
	entriesQuery := queryFactory.NewEntriesQuery(conn)
	reportsQuery := queryFactory.NewReportsQuery(conn)

	// services

	cacheService := cache.NewService(redisCacheClient)
	persistentStoreService := persistent_store.NewService(redisPersistentStoreClient)
	usersService := users.NewService(conn)
	accessService := access.NewService(usersService, persistentStoreService)
	categoriesService := categories.NewService(logger.New("service", "categories"), categoriesQuery, cacheService, queryFactory)
	entriesService := entries.NewService(entriesQuery)
	reportsService := reports.NewService(reportsQuery)

	// fiber app

	app := fiber.New()
	app.Use(fiberLogger.New())

	// handlers

	budgetAppHandler := budget_app.NewHandler(validate, app, usersService, accessService, categoriesService, entriesService, reportsService)
	budgetAppHandler.SetupRoutes()

	// start the app

	if startAppErr := app.Listen(":5000"); startAppErr != nil {
		logger.Crit("can't start fiber app", "err", startAppErr.Error())
	}
}
