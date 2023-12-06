package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// budget := Budget{
	// 	Name:   "test",
	// 	Amount: 100.00,
	// }

	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS budgets (
		id INTEGER PRIMARY KEY,
		name TEXT,
		amount FLOAT
	)`)
	if err != nil {
		fmt.Println("Error creating table:", err)
		return
	}

	result, err := db.Exec("INSERT INTO budgets (name, amount) VALUES (?, ?)", "Test", 100.00)
	if err != nil {
		fmt.Println("Error inserting data:", err)
		return
	}

	lastInsertID, _ := result.LastInsertId()
	fmt.Println("Inserted ID:", lastInsertID)

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

	//saveBudget(budget)
}

type Budget struct {
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
}

// func saveBudget(budget Budget) error {
// 	budgetFile, err := os.Open("budgets.json")
// 	if err != nil {
// 		fmt.Println("Error creating file: ", err)
// 		return nil // TODO: change this
// 	}
// 	defer budgetFile.Close()

// 	jsonEncoder := json.NewEncoder(budgetFile)

// 	err = jsonEncoder.Encode(budget)
// 	if err != nil {
// 		fmt.Println("Error encoding JSON: ", err)
// 		return nil // TODO: change this
// 	}

// 	fmt.Println("Data written to JSON successfully")
// 	return nil
// }
