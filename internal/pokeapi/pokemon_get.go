package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	fullUrl := baseUrl + "/pokemon/" + pokemonName
	pokemon := Pokemon{}
	if val, ok := c.cache.Get(fullUrl); ok {
		err := json.Unmarshal(val, &pokemon)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemon, nil
	}

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		return Pokemon{}, fmt.Errorf("bad status code: %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(fullUrl, data)

	if err := json.Unmarshal(data, &pokemon); err != nil {
		return Pokemon{}, err
	}

	return pokemon, nil
}
