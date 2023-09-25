package queries

import (
	"database/sql"
)

func SearchRecipes(db *sql.DB, searchTerm string) ([]recipeSummary, error) {
	var recipes []recipeSummary

	rows, err := db.Query("SELECT rcp.id, rcp.title, rcp.description, rcp.time, rcp.servings FROM recipes rcp JOIN ingredient_headers ih ON ih.recipe_id = rcp.id JOIN ingredients ing ON ing.ingredient_header_id = ih.id WHERE rcp.title LIKE ? OR ing.description LIKE ? GROUP BY rcp.id", "%"+searchTerm+"%", "%"+searchTerm+"%")
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
