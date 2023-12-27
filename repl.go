package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"
)

func startRepl(db *sql.DB) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")

		scanner.Scan()
		text := scanner.Text()
		text = strings.ToLower(text)
		words := strings.Fields(text)
		if len(words) == 0 {
			continue
		}
		inputCommand := words[0]

		availableCommands := getCommands()

		command, ok := availableCommands[inputCommand]
		if !ok {
			fmt.Println("Command not found")
			continue
		}
		err := command.callback(db)
		if err!= nil {
			fmt.Println("Error executing command:", err)
			continue
		}
	}
}

type command struct {
	name string
	description string
	callback func(*sql.DB) error
}

// TODO: Make functions for commands
func getCommands() map[string]command {
	return map[string]command {	
		"help": {
			name: "help",
			description: "List commands",
			callback: func(*sql.DB) error {
				return nil
			},
		},
		"create": {
			name: "create",
			description: "Create a budget",
			callback: func(*sql.DB) error {
				return nil
			},
		},
		"read": {
			name: "read",
			description: "Get all budgets",
			callback: readAllDbBudgets,
		},
		"update": {
			name: "update",
			description: "Update a budget",
			callback: func(*sql.DB) error {
				return nil
			},
		},
		"delete": {
			name: "delete",
			description: "Delete a budget",
			callback: func(*sql.DB) error {
				return nil
			},
		},
		"quit": {
			name: "quit",
			description: "Quit the program",
			callback: func(*sql.DB) error {
				os.Exit(0)
				return nil
			},
		},
	}
}