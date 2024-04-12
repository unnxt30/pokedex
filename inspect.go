package main

import (
	"fmt"
)

func inspectPokemon(cfg *config, s string) error {

	pokemon, ok := caughtPokemon[s];

	if ok{
		fmt.Printf("\nName: %v\n", pokemon.Name);
		fmt.Printf("\nHeight: %v\n", pokemon.Height);
		fmt.Printf("\nWeight: %v\n", pokemon.Weight);
		fmt.Println("\nStats: \n");

		stats := pokemon.Stats;
		for _, statValue := range stats{
			fmt.Printf("- %v: %v\n", statValue.Stat.Name, statValue.BaseStat);
		}
		
		poke_type := pokemon.Types;
		fmt.Println("\nTypes: ");
		for _, types := range poke_type{
			fmt.Printf("\n- %v\n", types.Type.Name);
		}

		return nil;

	}

	fmt.Println("\n You haven't caught the pokemon yet :/");

	return nil;
}