package queries

import (
	"database/sql"
)

func getIngredientSections(db *sql.DB, id int64) ([]ingredientSection, error) {
	var ingredientSections []ingredientSection

	rows, err := db.Query("SELECT id, name FROM ingredient_headers WHERE recipe_id = ?", id)
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
