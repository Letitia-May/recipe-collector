package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type recipe struct {
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

type recipesResource struct {
	db *sql.DB
}

// Set db connection on recipesResource struct
func NewRecipesResource(db *sql.DB) recipesResource {
	return recipesResource{db: db}
}

// Link handlers to routes
func (rr recipesResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", rr.allRecipesHandler)

	return r
}

// Format recipe data
func (rr recipesResource) allRecipesHandler(w http.ResponseWriter, r *http.Request) {
	recipes, err := allRecipes(rr.db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(w, "Recipes found: ")
	for _, r := range recipes {
		fmt.Fprintln(w, r)
	}
}

// Get recipe data from db
func allRecipes(db *sql.DB) ([]recipe, error) {
	var recipes []recipe

	rows, err := db.Query("SELECT * FROM recipes")
	if err != nil {
		return nil, fmt.Errorf("allRecipes: %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var recipe recipe
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

// Build string with recipe data
func (r recipe) String() string {
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
