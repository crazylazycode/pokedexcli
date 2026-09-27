package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {

	if len(args) == 0 {
		return errors.New("you must provide a pokemon name")
	}

	pokemonName := args[0]

	pokemon, ok := cfg.pokedex[pokemonName]

	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, pokemonType := range pokemon.Types {
		fmt.Printf("  - %s\n", pokemonType.Type.Name)
	}

	return nil
}
