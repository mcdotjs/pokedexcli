package main

var config *Config

type Config struct {
	Next     string
	Previous string
}

func main() {
	if config == nil {
		config = &Config{
			Next:     "https://pokeapi.co/api/v2/location-area",
		}
	}
	startRepl()
}
