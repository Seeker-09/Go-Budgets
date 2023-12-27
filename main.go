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

	fmt.Println("Type help for a list of commands")

	startRepl(db)

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
}