package main

import (
	"fmt"
)

func pokedexCommand(c *config, args []string) error {

	fmt.Println("Your Pokedex:")

	for _, s := range *c.Pokedex {
		fmt.Println(" - ", s.Name)
	}

	return nil

}
