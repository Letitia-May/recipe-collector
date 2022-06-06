package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type recipe struct {
	ID          int64    `json:"id"`
	Title       string   `json:"title"`
	Description *string  `json:"description,omitempty"`
	Time        *string  `json:"time,omitempty"`
	Servings    *string  `json:"servings,omitempty"`
	Url         *string  `json:"url,omitempty"`
	Notes       *string  `json:"notes,omitempty"`
	Rating      *float32 `json:"rating,omitempty"`
	TimesCooked *int64   `json:"times_cooked,omitempty"`
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

	recipesJson, err := json.Marshal(recipes)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8000")
	w.Write(recipesJson)
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
