package queries

import (
	"database/sql"
	"fmt"
)

func SearchRecipes(db *sql.DB, searchTerm string) ([]recipeSummary, error) {
	var recipes []recipeSummary

	rows, err := db.Query("SELECT id, title, description, time, servings FROM recipes WHERE recipes.title LIKE ?", "%"+searchTerm+"%")
	if err != nil {
		return nil, fmt.Errorf("searchRecipes: %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var recipe recipeSummary
		if err := rows.Scan(&recipe.ID, &recipe.Title, &recipe.Description, &recipe.Time, &recipe.Servings); err != nil {
			return nil, fmt.Errorf("searchRecipes: %v", err)
		}
		recipes = append(recipes, recipe)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("searchRecipes: %v", err)
	}

	return recipes, nil
}