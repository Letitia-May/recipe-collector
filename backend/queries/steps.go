package queries

import (
	"database/sql"
)

func getRecipeSteps(db *sql.DB, id int64) ([]step, error) {
	var steps []step

	rows, err := db.Query("SELECT number, description FROM steps WHERE recipe_id = ?", id)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var step step
		if err := rows.Scan(&step.Number, &step.Description); err != nil {
			panic(err)
		}
		steps = append(steps, step)
	}
	if err := rows.Err(); err != nil {
		panic(err)
	}

	return steps, nil
}

func createStep(db *sql.DB, recipeID int64, step *step) error {
	_, err := db.Exec(
		"INSERT INTO steps (recipe_id, number, description) VALUES (?, ?, ?)",
		recipeID,
		step.Number,
		step.Description,
	)
	return err
}
