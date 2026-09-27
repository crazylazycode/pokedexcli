package main

import (
	"time"

	"github.com/crazytobeawake/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Second)
	pokedex := make(map[string]pokeapi.Pokemon)
	cfg := &config{
		commands:      getCommands(),
		pokeapiClient: pokeClient,
		pokedex:       pokedex,
	}

	startRepl(cfg)
}
