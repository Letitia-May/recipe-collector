package queries

import (
	"database/sql"
	"fmt"
	"log"
)

func getIngredientSections(db *sql.DB, id int64) ([]ingredientSection, error) {
	var ingredientSections []ingredientSection

	rows, err := db.Query("SELECT id, name FROM ingredient_headers WHERE recipe_id = ?", id)
	if err != nil {
		return nil, fmt.Errorf("getIngredientSections %d: recipe has no ingredient headers", id)
	}

	defer rows.Close()

	for rows.Next() {
		var ingredientSection ingredientSection
		if err := rows.Scan(&ingredientSection.ID, &ingredientSection.Heading); err != nil {
			return nil, fmt.Errorf("getIngredientSections %d: %v", id, err)
		}
		ingredientSections = append(ingredientSections, ingredientSection)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("getIngredientSections %d: %v", id, err)
	}

	for i := range ingredientSections {
		ingredients, err := getIngredients(db, ingredientSections[i].ID)
		if err != nil {
			log.Fatal(err)
		}
		ingredientSections[i].Ingredients = ingredients
	}

	return ingredientSections, nil
}
