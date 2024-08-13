package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(str string) []string {
	lowerCase := strings.ToLower(str)
	words := strings.Fields(lowerCase)
	return words
}

type cliCommand struct {
	callback    func() error
	name        string
	description string
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Lists some location areas",
			callback:    callbackMap,
		},
		"exit": {
			name:        "exit",
			description: "Exits the Pokedex",
			callback:    commandExit,
		},
	}
}

func StartRepl() {
	// creating a scanner to collect user input
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()

	for {
		fmt.Printf("\nPokedex >  ")
		// calling a method on scanner to actually collect input
		scanner.Scan()
		input := cleanInput(scanner.Text())

		if len(input) == 0 {
			continue
		}

		cmd, ok := commands[input[0]]

		if !ok {
			fmt.Println("Unkown command, type 'help' for available commands")
			continue
		}

		err := cmd.callback()
		if err != nil {
			fmt.Println("Error: ", err)
		}

	}

}
