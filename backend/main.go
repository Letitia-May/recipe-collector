package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"recipe-collector/backend/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-sql-driver/mysql"
	"github.com/urfave/cli/v3"
)

// Temporary global variable
var db *sql.DB

func startServer() error {
	// Capture connection properties
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   os.Getenv("DBADDR"),
		DBName: "recipe_collector",
	}

	// Get a db handle
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return err
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	fmt.Println("✅ Connected to database successfully")

	// Initialise a new router
	r := chi.NewRouter()

	// Mount all recipe handlers onto /recipes endpoint
	r.Mount("/recipes", handlers.NewRecipesResource(db).Routes())

	// Start web server
	return http.ListenAndServe("localhost:8080", r)
}

func main() {
	cmd := &cli.Command{
		Name:  "recipe-collector",
		Usage: "Manage and serve your recipe collection",
		Commands: []*cli.Command{
			{
				Name:   "serve",
				Usage:  "Start the web server",
				Action: func(context.Context, *cli.Command) error {
					fmt.Println("🔄 Starting the server...")
					return startServer()
				},
			},
			{
				Name: "migrate",
				Usage: "Run db migrations",
				Action: func(context.Context, *cli.Command) error {
					fmt.Println("🔄 Starting database migrations...")
					return runMigrations()
				},
			},
		},
	}

    if err := cmd.Run(context.Background(), os.Args); err != nil {
        log.Fatal(err)
    }
}
