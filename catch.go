package main

import (
	"fmt"
	"math/rand"
	"github.com/unnxt30/pokedex/internal/pokeapi"
)


func Catch_pokemon(cfg *config, s string) error {
	
	pokemon_data:= pokeapi.Get_pokemon_data(s);

	maxValue := 1000;
	baseValue := pokemon_data.BaseExperience

	prob_range := maxValue - baseValue

	randomNumber := rand.Intn(maxValue)
	fmt.Printf("Throwing a pokeball at %v...", s);


	if randomNumber < prob_range {
		fmt.Printf("%v Caught!", s);		
		return nil
	} 
	
	fmt.Printf("%v Escaped :( ", s);
	return nil


}
