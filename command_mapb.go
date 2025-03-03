package main

import (
	"errors"
	"fmt"
)

func mapCommandBack(c *config) error {
	if c.Previous == nil {
    //NOTE: rf
		t := "https://pokeapi.co/api/v2/location-area"
		c.Next = &t
		return errors.New("you're on the first page")
	}
	v, e := c.pokeapiClient.MakeRequest(c.Previous)
	if e != nil {
		return e
	}
	c.Next = v.Next
	c.Previous = v.Previous
	for _, loc := range v.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
