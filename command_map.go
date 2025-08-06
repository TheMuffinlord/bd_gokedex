package main

import (
	"errors"
	"fmt"
)

func commandMap(config *cmdConfig, params ...string) error {
	locationsResp, err := config.pokeapiClient.LARequest(config.NextURL)
	if err != nil {
		return err
	}

	config.NextURL = locationsResp.Next
	config.PrevURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapB(config *cmdConfig, params ...string) error {
	if config.PrevURL == nil {
		return errors.New("you're on the first page")
	}

	locationResp, err := config.pokeapiClient.LARequest(config.PrevURL)
	if err != nil {
		return err
	}

	config.NextURL = locationResp.Next
	config.PrevURL = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
