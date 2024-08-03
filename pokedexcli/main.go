package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
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

