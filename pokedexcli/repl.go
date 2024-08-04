package main

import (
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func cleanInput(str string) []string {
	lowerCase := strings.ToLower(str)
	words := strings.Fields(lowerCase)
	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exits the Pokedex",
			callback:    commandExit,
		},
	}
}

func commandHelp() error {
	commands := getCommands()

	fmt.Printf("\nWelcome to the Pokedex!\nAvailable commands:\n\n")
	for _, command := range commands {
		fmt.Printf("%s : %s \n", command.name, command.description)
	}
	return nil
}

func commandExit() error {
	fmt.Println("Exiting the Pokedex")
  os.Exit(0)
	return nil
}
