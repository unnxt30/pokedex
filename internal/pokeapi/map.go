package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type map_data struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type config struct {
	next_page_URL     *string
	previous_page_URL *string
}

func get_list(url *string) map_data {
	request_url := "https://pokeapi.co/api/v2/location-area"

	if url != nil {
		request_url = *url

	}

	page := map_data{}
	resp, err := http.Get(request_url)
	if err != nil {
		fmt.Println("Error")
	}

	body, err := io.ReadAll(resp.Body)

	marshaled := json.Unmarshal(body, &page)
	if marshaled != nil {
		fmt.Println(marshaled)
	}

	return page
}

func commandMapf(cfg *config) error {

	location_data := get_list(cfg.next_page_URL)

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
	location_data := get_list(cfg.previous_page_URL)

	cfg.next_page_URL = location_data.Next
	cfg.previous_page_URL = location_data.Previous

	for _, loc_name := range location_data.Results {
		fmt.Println(loc_name.Name)
	}

	return nil

}
