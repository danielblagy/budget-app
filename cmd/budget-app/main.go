package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/danielblagy/budget-app/intenal/handler"
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
	log.Println("+", os.Getenv(envDatabaseUrl))

	// urlExample := "postgres://username:password@localhost:5432/database_name"
	conn, err := pgx.Connect(context.Background(), os.Getenv(envDatabaseUrl))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	http.HandleFunc("/", handler.Greet)
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		rows, err := conn.Query(context.Background(), "select handle from users")
		if err != nil {
			fmt.Fprintf(w, "error: "+err.Error())
			return
		}

		for rows.Next() {
			var handle string
			err := rows.Scan(&handle)
			if err != nil {
				fmt.Fprintf(w, "error: "+err.Error())
				return
			}
			fmt.Fprintf(w, "%s\n", handle)
		}
	})
	http.ListenAndServe(":8080", nil)
}
