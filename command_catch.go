package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(c *config, args []string) error {
	pokemon := args[0]
	fmt.Println("Throwing a Pokeball at " + pokemon + "...")
	url := "https://pokeapi.co/api/v2/pokemon/" + pokemon
	v, err := c.pokeapiClient.CatchPokemon(url)
	if err != nil {
		return err
	}

	num := rand.Intn(3)
	if num == 2 {
		(*c.Pokedex)[v.Name] = v
		fmt.Println(v.Name + " was caught!")
		fmt.Println("You may now inspect it with the inspect command.")
	}
	return nil
}
