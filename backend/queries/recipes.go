package queries

import (
	"database/sql"
	"fmt"
	"log"
)

type recipeSummary struct {
	ID          int64   `json:"id"`
	Title       string  `json:"title"`
	Description *string `json:"description,omitempty"`
	Time        *string `json:"time,omitempty"`
	Servings    *string `json:"servings,omitempty"`
}

type recipe struct {
	ID                 int64               `json:"id"`
	Title              string              `json:"title"`
	Description        *string             `json:"description,omitempty"`
	Time               *string             `json:"time,omitempty"`
	Servings           *string             `json:"servings,omitempty"`
	Url                *string             `json:"url,omitempty"`
	Notes              *string             `json:"notes,omitempty"`
	TimesCooked        *int64              `json:"timesCooked,omitempty"`
	Steps              []step              `json:"steps"`
	IngredientSections []ingredientSection `json:"ingredientSections"`
}

type step struct {
	Number      int64  `json:"number"`
	Description string `json:"description"`
}

type ingredientSection struct {
	ID          int64    `json:"-"`
	Heading     string   `json:"heading"`
	Ingredients []string `json:"ingredients"`
}

func GetAllRecipes(db *sql.DB) ([]recipeSummary, error) {
	var recipes []recipeSummary

	rows, err := db.Query("SELECT id, title, description, time, servings FROM recipes")
	if err != nil {
		return nil, fmt.Errorf("getAllRecipes: %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var recipe recipeSummary
		if err := rows.Scan(&recipe.ID, &recipe.Title, &recipe.Description, &recipe.Time, &recipe.Servings); err != nil {
			return nil, fmt.Errorf("getAllRecipes: %v", err)
		}
		recipes = append(recipes, recipe)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("getAllRecipes: %v", err)
	}

	return recipes, nil
}

func GetRecipe(db *sql.DB, id int64) (*recipe, error) {
	var recipe recipe

	row := db.QueryRow("SELECT * FROM recipes WHERE id = ?", id)
	if err := row.Scan(&recipe.ID, &recipe.Title, &recipe.Description, &recipe.Time, &recipe.Servings, &recipe.Url, &recipe.Notes, &recipe.TimesCooked); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("getRecipe %d: no such recipe", id)
		}
		return nil, fmt.Errorf("getRecipe %d: %v", id, err)
	}

	steps, err := getRecipeSteps(db, id)
	if err != nil {
		log.Fatal(err)
	}
	recipe.Steps = steps

	ingredientSections, err := getIngredientSections(db, id)
	if err != nil {
		log.Fatal(err)
	}
	recipe.IngredientSections = ingredientSections

	return &recipe, nil
}
