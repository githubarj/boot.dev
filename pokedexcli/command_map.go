package main

import (
	"fmt"
	"log"
)

func callbackMap(cfg *config) error {

	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaUrl)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Location Areas")

	for _, area := range resp.Results {
		fmt.Printf(" -%s\n", area.Name)
	}

	cfg.nextLocationAreaUrl = resp.Next
	cfg.prevLocationAreaUrl = resp.Previous
	return nil
}
