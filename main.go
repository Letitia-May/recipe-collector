package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"recipe-collector/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-sql-driver/mysql"
)

// Temporary global variable
var db *sql.DB

func main() {
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
		log.Fatal(err)
	}

	// Confirm the db connection
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected")

	// Initialise a new router
	r := chi.NewRouter()

	// Mount all recipe handlers onto /recipes endpoint
	r.Mount("/recipes", handlers.NewRecipesResource(db).Routes())

	// Start web server
	log.Fatal(http.ListenAndServe(":8080", r))
}
