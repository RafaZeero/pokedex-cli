package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Removes spaces, transform to lower case and returns all the commands separated with space
*/
func cleanInput(s string) []string {
	s1 := strings.Trim(s, " ")
	s2 := strings.ToLower(s1)
	return strings.Fields(s2)
}

func Start() {
	fmt.Println("Welcome to the Pokedex!")

	reader := bufio.NewScanner(os.Stdin)
	printPrompt()
	for reader.Scan() {

		input := cleanInput(reader.Text())
		if len(input) == 0 {
			printPrompt()
			continue
		}

		cmd := input[0]

		do, ok := commands[cmd]
		if !ok {
			fmt.Println("Unknown command")
			printPrompt()
			continue
		}
		do.callback()
		printPrompt()

	}
}

const cliName = "Pokedex"

type Commands map[string]cliCommand

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

// printPrompt displays the repl prompt at the start of each loop
func printPrompt() {
	fmt.Print(cliName, "> ")
}
