package main

import (
	"fmt"
	"log"

	"github.com/githubarj/pokedexcli/internal/pokeapi"
)

func callbackMap() error {

	pokeapiClient := pokeapi.NewClient()

	resp, err := pokeapiClient.ListLocationAreas()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}
