package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {

	if len(args) == 0 {
		return errors.New("you must provide a location area name")
	}
	areaName := args[0]
	fmt.Printf("Exploring %s...\n", areaName)

	locationAreaResp, err := cfg.pokeapiClient.GetLocation(areaName)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, encounter := range locationAreaResp.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	return nil
}
