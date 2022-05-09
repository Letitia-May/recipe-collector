package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-sql-driver/mysql"
)

// Temporary global variable
var db *sql.DB

type Recipe struct {
	ID          int64
	Title       string
	Description *string
	Time        *string
	Servings    *string
	Url         *string
	Notes       *string
	Rating      *float32
	TimesCooked *int64
}

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

	// Confirms the db connection
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected")

	http.HandleFunc("/", allRecipesHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func allRecipesHandler(w http.ResponseWriter, r *http.Request) {
	recipes, err := allRecipes()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(w, "Recipes found: ")
	for _, r := range recipes {
		fmt.Fprintln(w, r)
	}
}

func (r Recipe) String() string {
	s := fmt.Sprintf("%d %s", r.ID, r.Title)

	if r.Description != nil {
		s = s + fmt.Sprintf(" %s", *r.Description)
	}

	if r.Time != nil {
		s = s + fmt.Sprintf(" %s", *r.Time)
	}

	if r.Servings != nil {
		s = s + fmt.Sprintf(" %s", *r.Servings)
	}

	if r.Url != nil {
		s = s + fmt.Sprintf(" %s", *r.Url)
	}

	if r.Notes != nil {
		s = s + fmt.Sprintf(" %s", *r.Notes)
	}

	if r.Rating != nil {
		s = s + fmt.Sprintf(" %f", *r.Rating)
	}

	if r.TimesCooked != nil {
		s = s + fmt.Sprintf(" %d", *r.TimesCooked)
	}

	return s
}

func allRecipes() ([]Recipe, error) {
	var recipes []Recipe

	rows, err := db.Query("SELECT * FROM recipes")
	if err != nil {
		return nil, fmt.Errorf("allRecipes: %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var recipe Recipe
		if err := rows.Scan(&recipe.ID, &recipe.Title, &recipe.Description, &recipe.Time, &recipe.Servings, &recipe.Url, &recipe.Notes, &recipe.Rating, &recipe.TimesCooked); err != nil {
			return nil, fmt.Errorf("allRecipes: %v", err)
		}
		recipes = append(recipes, recipe)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("allRecipes: %v", err)
	}

	return recipes, nil
}
