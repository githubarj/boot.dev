package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
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

func main() {
	// creating a scanner to collect user input
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()

	for {
		fmt.Printf("\nPokedex >  ")
		// calling a method on scanner to actually collect input
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		if cmd, ok := commands[input]; ok {
			err := cmd.callback()
			if err != nil {
				fmt.Println("Error: ", err)
			}

		} else {
			fmt.Println("Unkown command, type 'help' for available commands")
		}

		if input == "exit" {
			break
		}

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
	return nil
}
