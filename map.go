package main

import (
	"errors"
	"fmt"
	"github.com/unnxt30/pokedex/internal/pokeapi"
)

func commandMapf(cfg *config) error {

	location_data := pokeapi.Get_list(cfg.next_page_URL)

	cfg.next_page_URL = location_data.Next
	cfg.previous_page_URL = location_data.Previous

	for _, loc_name := range location_data.Results {
		fmt.Println(loc_name.Name)
	}

	return nil

}

func commandMapb(cfg *config) error {
	if cfg.previous_page_URL == nil {
		return errors.New("you are on the first page")

	}
	location_data := pokeapi.Get_list(cfg.previous_page_URL)

	cfg.next_page_URL = location_data.Next
	cfg.previous_page_URL = location_data.Previous

	for _, loc_name := range location_data.Results {
		fmt.Println(loc_name.Name)
	}

	return nil

}
