package main

import (
	"database/sql"
	"fmt"
)

type Budget struct {
	name   string
	amount float64
}

func createDbBudgetTable(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS budgets (
		id INTEGER PRIMARY KEY,
		name TEXT,
		amount FLOAT
	)`)
	if err != nil {
		return err
	}

	return nil
}

// TODO: remake this to fit with command struct
// stmt.Exec result is a type of int64
func createDbBudget(db *sql.DB, budget Budget) (int64, error) {
	stmt, err := db.Prepare("INSERT INTO budgets (name, amount) VALUES (?, ?)")
	if err != nil {
		fmt.Println("Error preparing statement:", err)
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(budget.name, budget.amount)
	if err != nil {
		fmt.Println("Error inserting data:", err)
		return 0, err
	}
	lastInsertID, _ := result.LastInsertId()

	return lastInsertID, nil
}

func readAllDbBudgets(db *sql.DB) error {
	rows, err := db.Query("SELECT id, name, amount FROM budgets")
	if err != nil {
		fmt.Println("Error querying data:", err)
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var id, age int
		var name string
		err := rows.Scan(&id, &name, &age)
		if err != nil {
			fmt.Println("Error scanning rows:", err)
			return err
		}
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", id, name, age)
	}
	
	return nil
}

// TODO: Implement Reads (multiple) Update and Delete