package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(" > ")

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
		err := command.callback()
		if err!= nil {
			fmt.Println("Error executing command:", err)
			continue
		}
	}
}

type command struct {
	name string
	description string
	callback func() error
}

func getCommands() map[string]command {
	return map[string]command {	
		"help": {
			name: "help",
			description: "List commands",
			callback: func() error {
				return nil
			},
		},
		"create": {
			name: "create",
			description: "Create a budget",
			callback: func() error {
				return nil
			},
		},
		"read": {
			name: "read",
			description: "Read a budget",
			callback: func() error {
				return nil
			},
		},
		"update": {
			name: "update",
			description: "Update a budget",
			callback: func() error {
				return nil
			},
		},
		"delete": {
			name: "delete",
			description: "Delete a budget",
			callback: func() error {
				return nil
			},
		},
		"quit": {
			name: "quit",
			description: "Quit the program",
			callback: func() error {
				os.Exit(0)
				return nil
			},
		},
	}
}

// TODO: Make a list of commands
// TODO: Implement REPL and Commands to go with ^
