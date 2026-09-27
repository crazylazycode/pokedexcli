package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseUrl + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		locationResp := RespShallowLocations{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return RespShallowLocations{}, err
		}
		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return RespShallowLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		return RespShallowLocations{}, fmt.Errorf("bad status code: %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	c.cache.Add(url, data)

	var locationsResp RespShallowLocations
	if err := json.Unmarshal(data, &locationsResp); err != nil {
		return RespShallowLocations{}, err
	}

	return locationsResp, nil
}
