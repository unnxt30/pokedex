package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/unnxt30/pokedex/internal/pokecache"
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

type ExplorationData struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

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

