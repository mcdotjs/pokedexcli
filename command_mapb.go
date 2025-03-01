package main

import "fmt"

func mapCommandBack() error {
	v, e := makeRequest(config.Previous)
	if e != nil {
		return e
	}
	config.Next = v.Next
	config.Previous = v.Previous
	for loc := range v.Results {
		fmt.Println(v.Results[loc].Name)
	}

	return nil
}
