package main

import (
	"fmt"
)

func exploreCommand(c *config, args []string) error {
	location := args[0]
	fmt.Println("Exploring " + location + "...")

	url := "https://pokeapi.co/api/v2/location-area/" + location
	v, e := c.pokeapiClient.ListPokemonOccurence(url)
	if e != nil {
		return e
	}

	fmt.Println("Found Pokemon:")
	for _, val := range v.PokemonEncounters {
		fmt.Println(val.Pokemon.Name)
	}

	return nil

}
