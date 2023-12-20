package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := openDb()
	if err != nil {
		fmt.Println("Error opening Budgets Database", err)
		return
	}
	defer db.Close()

	err = createDbBudgetTable(db)
	if err != nil {
		fmt.Println("Error creating table:", err)
		return
	}

	insertedId, err := createDbBudget(db, Budget{
		name: "test",
		amount: 100.00,
	})
	if err != nil {
		fmt.Println("Error creating Budget", err)
	}
	fmt.Println("Inserted ID:", insertedId)

	rows, err := db.Query("SELECT id, name, amount FROM budgets")
	if err != nil {
		fmt.Println("Error querying data:", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id, age int
		var name string
		err := rows.Scan(&id, &name, &age)
		if err != nil {
			fmt.Println("Error scanning rows:", err)
			return
		}
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", id, name, age)
	}
}

type Budget struct {
	name   string
	amount float64
}

func openDb() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "budgets.db")
	if err != nil {
		return db, err
	}
	return db, nil
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