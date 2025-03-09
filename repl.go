package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(c *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}
		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(c, args)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    helpCommand,
		},
		"map": {
			name:        "map",
			description: "Map locations",
			callback:    mapCommandForward,
		},
		"mapb": {
			name:        "mapb",
			description: "Map locations backward",
			callback:    mapCommandBack,
		},
		"explore": {
			name:        "explore",
			description: "Explore loacations",
			callback:    exploreCommand,
		},
		"catch": {
			name:        "catch",
			description: "Catch pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect pokemon",
			callback:    inspectCommand,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Show my pokedex",
			callback:    pokedexCommand,
		},
	}
}
