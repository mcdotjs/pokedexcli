package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Item struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type LocationAreaResponse struct {
	Results  []Item `json:"results"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
}

func makeRequest(url string) (LocationAreaResponse, error) {
	if len(url) == 0 {
		config.Next = "https://pokeapi.co/api/v2/location-area"
    e := fmt.Errorf("you're on the first page")
		return LocationAreaResponse{}, e
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		e := fmt.Errorf("Problem with request")
		return LocationAreaResponse{}, e
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		e := fmt.Errorf("Problem with Client")
		return LocationAreaResponse{}, e
	}
	defer res.Body.Close()

	var result LocationAreaResponse
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&result)
	if err != nil {
		e := fmt.Errorf("Problem with decode")
		return LocationAreaResponse{}, e
	}

	return result, nil
}

func mapCommand() error {
	v, e := makeRequest(config.Next)
	if e != nil {
		return e
	}
	config.Next = v.Next
	config.Previous = v.Previous
	for loc := range v.Results {
		fmt.Println(v.Results[loc].Name)
	}

	return nil
}
