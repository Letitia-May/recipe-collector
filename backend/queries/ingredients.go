package queries

import (
	"database/sql"
)

func getIngredients(db *sql.DB, id int64) ([]string, error) {
	var ingredients []string

	rows, err := db.Query("SELECT description FROM ingredients WHERE ingredient_section_id = ?", id)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var ingredient string
		if err := rows.Scan(&ingredient); err != nil {
			panic(err)
		}
		ingredients = append(ingredients, ingredient)
	}
	if err := rows.Err(); err != nil {
		panic(err)
	}

	return ingredients, nil
}
