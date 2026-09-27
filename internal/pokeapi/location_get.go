package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocation(locationName string) (Location, error) {
	fullUrl := baseUrl + "/location-area/" + locationName

	locationArea := Location{}
	if val, ok := c.cache.Get(fullUrl); ok {
		err := json.Unmarshal(val, &locationArea)
		if err != nil {
			return Location{}, err
		}
		return locationArea, nil
	}

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return Location{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		return Location{}, fmt.Errorf("bad status code: %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Location{}, err
	}

	c.cache.Add(fullUrl, data)

	if err := json.Unmarshal(data, &locationArea); err != nil {
		return Location{}, err
	}

	return locationArea, nil
}
