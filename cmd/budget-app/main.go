package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/danielblagy/budget-app/intenal/handler"
	"github.com/danielblagy/budget-app/intenal/service/users"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

const envDatabaseUrl = "DATABASE_URL"

func main() {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	log.Println(os.Getenv(envDatabaseUrl))

	// urlExample := "postgres://username:password@localhost:5432/database_name"
	conn, err := pgx.Connect(context.Background(), os.Getenv(envDatabaseUrl))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	usersService := users.NewService(conn)

	app := handler.NewHandler(usersService)

	http.HandleFunc("/", app.Greet)
	http.HandleFunc("/users", app.GetUsers)
	http.ListenAndServe(":8080", nil)
}
