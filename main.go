package main

import (
	"bufio"
	"fmt"
	"os"
)

const cliName = "Pokedex"

type Commands map[string]cliCommand

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var cmd Commands

func create() Commands {
	return Commands{
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
func commandExit() error {
	os.Exit(0)
	return nil
}
func commandHelp() error {
	fmt.Print("\nUsage:\n\n")
	for _, c := range cmd {
		fmt.Printf("%s: %s\n", c.name, c.description)
	}
	fmt.Println("")
	return nil
}

// printPrompt displays the repl prompt at the start of each loop
func printPrompt() {
	fmt.Print(cliName, "> ")
}

func main() {
	fmt.Println("Welcome to the Pokedex!")
	cmd = create()

	scanner := bufio.NewScanner(os.Stdin)
	printPrompt()
	for {
		for scanner.Scan() {
			text := scanner.Text()

			switch text {
			case "help":
				cmd["help"].callback()
			case "exit":
				cmd["exit"].callback()
			}
			printPrompt()
		}
	}

}
