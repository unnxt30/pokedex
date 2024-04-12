package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
	"github.com/unnxt30/pokedex/internal/pokecache"
	"errors"
	
)



var myCache = pokecache.NewCache(5 * time.Second)

func Get_list(url *string) map_data {
	request_url := "https://pokeapi.co/api/v2/location-area"

	if url != nil {
		request_url = *url

	}

	cached, ok := myCache.Get(request_url)
	if ok {
		page := map_data{}
		marshaled := json.Unmarshal(cached, &page)
		if marshaled != nil {
			fmt.Println(marshaled)
		}
		return page
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

	myCache.Add(request_url, body);
	defer resp.Body.Close()
	return page
}


func GetExplorationList(s *string) ExplorationData{
	request_url := "https://pokeapi.co/api/v2/location-area/canalave-city-area"

	if s != nil {
		request_url = *s

	}

	cached, ok := myCache.Get(request_url)
	if ok {
		page := ExplorationData{}
		marshaled := json.Unmarshal(cached, &page)
		if marshaled != nil {
			fmt.Println(marshaled)
		}
		return page
	}

	page := ExplorationData{}
	resp, err := http.Get(request_url)
	if err != nil {
		fmt.Println("Error")
	}

	body, err := io.ReadAll(resp.Body)

	marshaled := json.Unmarshal(body, &page)
	if marshaled != nil {
		fmt.Println(marshaled)
	}

	myCache.Add(request_url, body);
	defer resp.Body.Close()
	return page
	
};

func Get_pokemon_data(s string) (PokemonData, error){
	request_url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%v", s);

	resp, err := http.Get(request_url);

	if err != nil {
		fmt.Println(err);
		return PokemonData{}, errors.New(err.Error());
	}

	body, err := io.ReadAll(resp.Body);

	pokeStat := PokemonData{}

	marshaled := json.Unmarshal(body, &pokeStat);
	if marshaled != nil {
		fmt.Println(marshaled);
		return PokemonData{},errors.New(marshaled.Error())
	}


	return pokeStat, nil;

}