package queries

import (
	"database/sql"
	"fmt"
)

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
