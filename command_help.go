package main

import (
	"fmt"
)

func helpCommand() error {
	fmt.Print("Welcome to the Pokedex!\n\nUsage:\n")

	for _, c := range getCommands() {
		fmt.Printf("%s: %s\n", c.name, c.description)
	}
	fmt.Println("")
	return nil
}
