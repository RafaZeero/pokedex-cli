package repl

import (
	"fmt"
	"os"
)

var commands Commands

func init() {
	commands = createCommands()
}

func createCommands() Commands {
	return Commands{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    displayHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    exitProgram,
		},
	}
}

func exitProgram() error {
	os.Exit(0)
	return nil
}
func displayHelp() error {
	fmt.Print("\nUsage:\n\n")
	for _, c := range commands {
		fmt.Printf("%s: %s\n", c.name, c.description)
	}
	fmt.Println("")
	return nil
}
