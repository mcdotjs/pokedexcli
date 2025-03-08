package main

import (
	"github.com/mcdotjs/pokedex/internal/api"
	"time"
)

type config struct {
	pokeapiClient api.Client
	Next          *string
	Previous      *string
	Pokedex       *map[string]api.Pokemon
}

func main() {
	pokeClient := api.NewClient(5 * time.Second)
	mape := make(map[string]api.Pokemon)
	config := &config{
		pokeapiClient: pokeClient,
		Pokedex:       &mape,
	}
	startRepl(config)
}
