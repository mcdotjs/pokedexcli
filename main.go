package main

import (
	"github.com/mcdotjs/pokedex/internal/api"
	"time"
)

type config struct {
	pokeapiClient api.Client
	Next          *string
	Previous      *string
}

func main() {
	pokeClient := api.NewClient(5 * time.Second)
	config := &config{
		pokeapiClient: pokeClient,
	}
	startRepl(config)
}
