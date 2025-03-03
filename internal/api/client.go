package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	httpClient http.Client
}

func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}

func (c *Client) MakeRequest(pageURL *string) (LocationAreaResponse, error) {
	url := baseURL + "/location-area"
fmt.Println(url)
	if pageURL != nil {
		url = *pageURL
	}
fmt.Println("2",url)
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

	var result LocationAreaResponse
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&result)
	if err != nil {
		e := fmt.Errorf("Problem with decode")
		return LocationAreaResponse{}, e
	}

	return result, nil
}
