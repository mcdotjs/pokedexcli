package api

import (
	"encoding/json"
	"fmt"
	"github.com/mcdotjs/pokedex/internal/pokecache"
	"io"
	"net/http"
	"time"
)

type Client struct {
	httpClient http.Client
	cache      *pokecache.Cache
}

func NewClient(timeout time.Duration) Client {
	cache := pokecache.NewCache(timeout)
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: cache,
	}
}

func (c *Client) ListLocations(pageURL *string) (LocationAreaResponse, error) {
	url := baseURL + "/location-area"
	var result LocationAreaResponse
	if pageURL != nil {

		url = *pageURL
	}

	val, exist := c.cache.Get(url)
	if exist {
		if err := json.Unmarshal(val, &result); err != nil {
			return LocationAreaResponse{}, err
		}
		return result, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		e := fmt.Errorf("Problem with request")
		return LocationAreaResponse{}, e
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		e := fmt.Errorf("Problem with Client")
		return LocationAreaResponse{}, e
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	c.cache.Add(url, data)

	if err := json.Unmarshal(data, &result); err != nil {
		return LocationAreaResponse{}, fmt.Errorf("problem with decode: %w", err)
	}
	return result, nil
}

func (c *Client) ListPokemonOccurence(url string) (PokemonEncountersResponse, error) {
	var result PokemonEncountersResponse
	val, exist := c.cache.Get(url)
	if exist {
		if err := json.Unmarshal(val, &result); err != nil {
			return PokemonEncountersResponse{}, err
		}
		return result, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonEncountersResponse{}, fmt.Errorf("Problem with request")
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonEncountersResponse{}, fmt.Errorf("Problem with Client")
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return PokemonEncountersResponse{}, err
	}
	c.cache.Add(url, data)

	if err := json.Unmarshal(data, &result); err != nil {
		return PokemonEncountersResponse{}, err
	}
	return result, nil

}
