package main

import (
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




