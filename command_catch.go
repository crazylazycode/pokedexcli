package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {

	if len(args) == 0 {
		return errors.New("you must provide a pokemon name")
	}

	pokemonName := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	threshold := 40

	res := rand.Intn(pokemon.BaseExperience)

	if res < threshold {
		fmt.Printf("%s was caught!\n", pokemonName)
		cfg.pokedex[pokemonName] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}

	return nil
}
