package queries

import (
	"database/sql"
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
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var recipe recipeSummary
		if err := rows.Scan(&recipe.ID, &recipe.Title, &recipe.Description, &recipe.Time, &recipe.Servings); err != nil {
			panic(err)
		}
		recipes = append(recipes, recipe)
	}
	if err := rows.Err(); err != nil {
		panic(err)
	}

	return recipes, nil
}

func GetRecipe(db *sql.DB, id int64) (*recipe, error) {
	var recipe recipe

	row := db.QueryRow("SELECT id, title, description, time, servings, url, notes, times_cooked FROM recipes WHERE id = ?", id)
	if err := row.Scan(&recipe.ID, &recipe.Title, &recipe.Description, &recipe.Time, &recipe.Servings, &recipe.Url, &recipe.Notes, &recipe.TimesCooked); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		panic(err)
	}

	steps, err := getRecipeSteps(db, id)
	if err != nil {
		panic(err)
	}
	recipe.Steps = steps

	ingredientSections, err := getIngredientSections(db, id)
	if err != nil {
		panic(err)
	}
	recipe.IngredientSections = ingredientSections

	return &recipe, nil
}

type RecipeInput struct {
	Title              string                     `json:"title"`
	Description        *string                    `json:"description,omitempty"`
	Time               *string                    `json:"time,omitempty"`
	Servings           *string                    `json:"servings,omitempty"`
	Url                *string                    `json:"url,omitempty"`
	Notes              *string                    `json:"notes,omitempty"`
	TimesCooked        *int64                     `json:"timesCooked,omitempty"`
	Steps              []StepInput                `json:"steps"`
	IngredientSections []IngredientSectionInput   `json:"ingredientSections"`
}

type StepInput struct {
	Number      int64  `json:"number"`
	Description string `json:"description"`
}

type IngredientSectionInput struct {
	Heading     string   `json:"heading"`
	Ingredients []string `json:"ingredients"`
}

func CreateRecipe(db *sql.DB, input *RecipeInput) (int64, error) {
	result, err := db.Exec(
		"INSERT INTO recipes (title, description, time, servings, url, notes, times_cooked) VALUES (?, ?, ?, ?, ?, ?, ?)",
		input.Title,
		input.Description,
		input.Time,
		input.Servings,
		input.Url,
		input.Notes,
		input.TimesCooked,
	)
	if err != nil {
		return 0, err
	}

	recipeID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	for _, section := range input.IngredientSections {
		heading := section.Heading
		if heading == "" {
			heading = "Ingredients"
		}
		sectionData := ingredientSection{
			Heading:     heading,
			Ingredients: section.Ingredients,
		}
		err := createIngredientSection(db, recipeID, &sectionData)
		if err != nil {
			return 0, err
		}
	}

	for _, stepInput := range input.Steps {
		stepData := step{
			Number:      stepInput.Number,
			Description: stepInput.Description,
		}
		err := createStep(db, recipeID, &stepData)
		if err != nil {
			return 0, err
		}
	}

	return recipeID, nil
}
