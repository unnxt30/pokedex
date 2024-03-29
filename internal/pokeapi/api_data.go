package pokeapi 

import (
	"encoding/json"
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



func Get_list(url *string) map_data {
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


