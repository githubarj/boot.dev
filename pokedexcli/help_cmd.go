package main

import "fmt"

func commandHelp() error {

	commands := getCommands()

	fmt.Printf("\nWelcome to the Pokedex!\nAvailable commands:\n\n")
	for _, command := range commands {
		fmt.Printf("%s : %s \n", command.name, command.description)
	}
	return nil
}
