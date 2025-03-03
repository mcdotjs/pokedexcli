package main

import (
	"fmt"
)

func mapCommand(c *config) error {
	v, e := c.pokeapiClient.MakeRequest(c.Next)
	if e != nil {
		return e
	}

	c.Next = v.Next
	c.Previous =v.Previous
	for _,loc := range v.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
