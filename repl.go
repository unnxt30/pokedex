package main

import (
	"bufio"
	"fmt"
	"os"
)

func repl_begin() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("\nPOKEDEX > ")

		scanner.Scan()
		text := scanner.Text()

		cliFunc()[text].callback()

		continue
	}
}

func cliFunc() map[string]cliCommand {

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

type cliCommand struct {
	name        string
	description string
	callback    func() error
}