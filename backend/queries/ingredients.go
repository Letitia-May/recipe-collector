package queries

import (
	"database/sql"
	"fmt"
)

func getIngredients(db *sql.DB, id int64) ([]string, error) {
	var ingredients []string

	rows, err := db.Query("SELECT description FROM ingredients WHERE ingredient_header_id = ?", id)
	if err != nil {
		return nil, fmt.Errorf("getIngredients %d: section has no ingredients", id)
	}

	defer rows.Close()

	for rows.Next() {
		var ingredient string
		if err := rows.Scan(&ingredient); err != nil {
			return nil, fmt.Errorf("getIngredients %d: %v", id, err)
		}
		ingredients = append(ingredients, ingredient)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("getIngredients %d: %v", id, err)
	}

	return ingredients, nil
}
