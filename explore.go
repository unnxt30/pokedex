package main

import (
	"fmt"

	"github.com/unnxt30/pokedex/internal/pokeapi"
)

func exploreArea(c *config, s string) error {
	request_url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%v", s);
	pokemon_data := pokeapi.GetExplorationList(&request_url)
	poke_encounters := pokemon_data.PokemonEncounters;
	fmt.Println("\n");
	for _, pokeNames := range poke_encounters{
		fmt.Printf("%v\n", pokeNames.Pokemon.Name)
	}
	
	return nil;

}