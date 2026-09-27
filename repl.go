package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/crazytobeawake/pokedexcli/internal/pokeapi"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		ok := scanner.Scan()
		if !ok {
			break
		}
		input := scanner.Text()
		cleaned := cleanInput(input)

		if len(cleaned) == 0 {
			continue
		}

		commandName := cleaned[0]
		args := []string{}
		if len(cleaned) > 1 {
			args = cleaned[1:]
		}

		command, exists := cfg.commands[commandName]

		if !exists {
			fmt.Println("Unkown command")
			continue
		} else {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		}

	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	words := strings.Fields(lowered)
	return words
}

type config struct {
	commands         map[string]cliCommand
	pokedex          map[string]pokeapi.Pokemon
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}
