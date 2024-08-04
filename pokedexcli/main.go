package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
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

		if cmd, ok := commands[input[0]]; ok {
			err := cmd.callback()
			if err != nil {
				fmt.Println("Error: ", err)
			}

		} else {
			fmt.Println("Unkown command, type 'help' for available commands")
		}

	}

}
