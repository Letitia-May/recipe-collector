package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type recipe struct {
	ID                 int64               `json:"id"`
	Title              string              `json:"title"`
	Description        *string             `json:"description,omitempty"`
	Time               *string             `json:"time,omitempty"`
	Servings           *string             `json:"servings,omitempty"`
	Url                *string             `json:"url,omitempty"`
	Notes              *string             `json:"notes,omitempty"`
	Rating             *float32            `json:"rating,omitempty"`
	TimesCooked        *int64              `json:"times_cooked,omitempty"`
	Steps              []step              `json:"steps"`
	IngredientSections []ingredientSection `json:"ingredient_sections"`
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

type recipesResource struct {
	db *sql.DB
}

func NewRecipesResource(db *sql.DB) recipesResource {
	return recipesResource{db: db}
}

func (rr recipesResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", rr.allRecipesHandler)
	r.Get("/{recipeID}", rr.getRecipeHandler)

	return r
}

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
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Write(recipesJson)
}

func (rr recipesResource) getRecipeHandler(w http.ResponseWriter, r *http.Request) {
	recipeID, err := strconv.ParseInt(chi.URLParam(r, "recipeID"), 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	recipe, err := getRecipe(rr.db, recipeID)
	if err != nil {
		log.Fatal(err)
	}

	recipeJson, err := json.Marshal(recipe)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Write(recipeJson)
}

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

func getRecipe(db *sql.DB, id int64) (*recipe, error) {
	var recipe recipe

	row := db.QueryRow("SELECT * FROM recipes WHERE id = ?", id)
	if err := row.Scan(&recipe.ID, &recipe.Title, &recipe.Description, &recipe.Time, &recipe.Servings, &recipe.Url, &recipe.Notes, &recipe.Rating, &recipe.TimesCooked); err != nil {
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

	ingredientSections, err := getRecipeIngredientSections(db, id)
	if err != nil {
		log.Fatal(err)
	}
	recipe.IngredientSections = ingredientSections

	return &recipe, nil
}

func getRecipeSteps(db *sql.DB, id int64) ([]step, error) {
	var steps []step

	rows, err := db.Query("SELECT number, description FROM steps WHERE recipe_id = ?", id)
	if err != nil {
		return nil, fmt.Errorf("getRecipeSteps %d: recipe has no steps", id)
	}

	defer rows.Close()

	for rows.Next() {
		var step step
		if err := rows.Scan(&step.Number, &step.Description); err != nil {
			return nil, fmt.Errorf("getRecipeSteps %d: %v", id, err)
		}
		steps = append(steps, step)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("getRecipeSteps %d: %v", id, err)
	}

	return steps, nil
}

func getRecipeIngredientSections(db *sql.DB, id int64) ([]ingredientSection, error) {
	var ingredientSections []ingredientSection

	rows, err := db.Query("SELECT id, name FROM ingredient_headers WHERE recipe_id = ?", id)
	if err != nil {
		return nil, fmt.Errorf("getRecipeIngredientSections %d: recipe has no ingredient headers", id)
	}

	defer rows.Close()

	for rows.Next() {
		var ingredientSection ingredientSection
		if err := rows.Scan(&ingredientSection.ID, &ingredientSection.Heading); err != nil {
			return nil, fmt.Errorf("getRecipeIngredientSections %d: %v", id, err)
		}
		ingredientSections = append(ingredientSections, ingredientSection)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("getRecipeIngredientSections %d: %v", id, err)
	}

	for i := range ingredientSections {
		ingredients, err := getSectionIngredients(db, ingredientSections[i].ID)
		if err != nil {
			log.Fatal(err)
		}
		ingredientSections[i].Ingredients = ingredients
	}

	return ingredientSections, nil
}

func getSectionIngredients(db *sql.DB, id int64) ([]string, error) {
	var ingredients []string

	rows, err := db.Query("SELECT description FROM ingredients WHERE ingredient_header_id = ?", id)
	if err != nil {
		return nil, fmt.Errorf("getSectionIngredients %d: section has no ingredients", id)
	}

	defer rows.Close()

	for rows.Next() {
		var ingredient string
		if err := rows.Scan(&ingredient); err != nil {
			return nil, fmt.Errorf("getSectionIngredients %d: %v", id, err)
		}
		ingredients = append(ingredients, ingredient)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("getSectionIngredients %d: %v", id, err)
	}

	return ingredients, nil
}
