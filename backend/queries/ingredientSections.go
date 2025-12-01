package queries

import (
	"database/sql"
)

func getIngredientSections(db *sql.DB, id int64) ([]ingredientSection, error) {
	var ingredientSections []ingredientSection

	rows, err := db.Query("SELECT id, name FROM ingredient_sections WHERE recipe_id = ?", id)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var ingredientSection ingredientSection
		if err := rows.Scan(&ingredientSection.ID, &ingredientSection.Heading); err != nil {
			panic(err)
		}
		ingredientSections = append(ingredientSections, ingredientSection)
	}
	if err := rows.Err(); err != nil {
		panic(err)
	}

	for i := range ingredientSections {
		ingredients, err := getIngredients(db, ingredientSections[i].ID)
		if err != nil {
			panic(err)
		}
		ingredientSections[i].Ingredients = ingredients
	}

	return ingredientSections, nil
}

func createIngredientSection(db *sql.DB, recipeID int64, section *ingredientSection) error {
	result, err := db.Exec(
		"INSERT INTO ingredient_sections (recipe_id, name) VALUES (?, ?)",
		recipeID,
		section.Heading,
	)
	if err != nil {
		return err
	}

	sectionID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	for _, ingredient := range section.Ingredients {
		err := createIngredient(db, sectionID, ingredient)
		if err != nil {
			return err
		}
	}

	return nil
}
