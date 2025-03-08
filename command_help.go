package main

import (
	"fmt"
)

func helpCommand(c *config, args []string) error {
	fmt.Print("Welcome to the Pokedex!\n\nUsage:\n")

	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println("")
	return nil
}
