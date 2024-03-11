package main

import (
	"bufio"
	"fmt"
	"os"
)

func repl_begin(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("\nPOKEDEX > ")

		scanner.Scan()
		text := scanner.Text()

		cmd := cliFunc()[text]

		cmd.callback(cfg)

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
		"map": {
			name:        "map",
			description: "Display next list of location areas",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Display next list of location areas of previous page",
			callback:    commandMapb,
		},
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}
