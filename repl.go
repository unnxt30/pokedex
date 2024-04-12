package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type config struct {
	next_page_URL     *string
	previous_page_URL *string
}

func Repl_begin(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("\nPOKEDEX > ")

		scanner.Scan()
		text := scanner.Text()
		argArray := strings.Split(text, " ");
		
		if len(argArray) > 1{
			cmd := cliFunc()[argArray[0]]
			cmd.callback(cfg, argArray[1]);
			continue
		}
		cmd := cliFunc()[argArray[0]];

		cmd.callback(cfg, "")

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
		"explore": {
			name: "explore",
			description: "Display the pokemons available in the location",
			callback: exploreArea,
		},
		"catch": {
			name: "catch",
			description: "Catch a pokemon you spotted",
			callback: Catch_pokemon,
		},
		"inspect": {
			name: "inspect",
			description: "inspect a pokemon inside your inventory",
			callback: inspectPokemon,
		},
		"pokedex": {
			name: "pokedex",
			description: "display all your caught pokemons.",
			callback: pokedex,
		},
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}
