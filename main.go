package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	budget := Budget{
		Name:   "test",
		Amount: 100.00,
	}

	saveBudget(budget)
}

type Budget struct {
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
}

func saveBudget(budget Budget) error {
	budgetFile, err := os.Create("budgets.json")
	if err != nil {
		fmt.Println("Error creating file: ", err)
		return nil // TODO: change this
	}
	defer budgetFile.Close()

	jsonEncoder := json.NewEncoder(budgetFile)

	err = jsonEncoder.Encode(budget)
	if err != nil {
		fmt.Println("Error encoding JSON: ", err)
		return nil // TODO: change this
	}

	fmt.Println("Data written to JSON successfully")
	return nil
}
