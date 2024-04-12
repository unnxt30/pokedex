package main

import (
	"errors"
	"fmt"
)

func pokedex(cfg *config, s string) error {
	if len(caughtPokemon) < 1 {
		fmt.Println("Skill Issue XD\n")
		return errors.New("foo") 
	}
	fmt.Println("Caught Pokemons: \n")

	for _, pokemon := range caughtPokemon {
		fmt.Printf("- %v\n", pokemon.Name);
	}
	return nil

}