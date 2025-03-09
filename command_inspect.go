package main

import (
	"fmt"
)

func inspectCommand(c *config, args []string) error {
	pokemon := args[0]

	p, exists := (*c.Pokedex)[pokemon]
	if !exists {
		fmt.Println("you have not caught that pokemon")
	}

	if exists {
		fmt.Println("Name: ", p.Name)
		fmt.Println("Weight: ", p.Weight)
		fmt.Println("Stats: ")
		for _, s := range p.Stats {
			fmt.Println(" - ", s.Stat.Name, ":", s.BaseStat)
		}
		fmt.Println("Types: ")
		for _, s := range p.Types {
			fmt.Println(" - ", s.Type.Name)
		}
	}

	return nil

}
