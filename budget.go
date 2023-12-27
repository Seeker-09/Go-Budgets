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

// TODO: Implement Reads (multiple) Update and Delete