package main

import (
	"errors"
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

func callbackMapb(cfg *config) error {
	if cfg.prevLocationAreaUrl == nil {
		return errors.New("You are the first page")
	}
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
