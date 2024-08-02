package main

import "fmt"

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
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func main() {
	for {

	}

}

func commandHelp() error {
	commands := getCommands()

	fmt.Println("Available commands")
	for _, command := range commands {
		fmt.Printf("%s : %s \n", command.name, command.description)
	}
	return nil
}

func commandExit() error {
	fmt.Println("Exiting the Pokedex")
	return nil
}
