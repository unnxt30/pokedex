package main

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/unnxt30/pokedex/internal/pokeapi"
)

var caughtPokemon = make(map[string]pokeapi.PokemonData);

func Catch_pokemon(cfg *config, s string) error {
	
	

	pokemon_data, ok:= pokeapi.Get_pokemon_data(s);

	if ok != nil {
		return errors.New("No such pokemon :): ")
	}

	maxValue := 1000;
	baseValue := pokemon_data.BaseExperience

	prob_range := maxValue - baseValue


	randomNumber := rand.Intn(maxValue)
	fmt.Printf("Throwing a pokeball at %v...\n", s);


	if randomNumber < prob_range {
		fmt.Printf("%v Caught!", s);		
		caughtPokemon[pokemon_data.Name] = pokemon_data;
		return nil
	} 
	
	fmt.Printf("%v Escaped :( ", s);
	return nil


}
