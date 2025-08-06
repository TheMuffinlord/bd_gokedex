package main

import (
	"errors"
	"fmt"
)

func commandExplore(config *cmdConfig, area ...string) error {
	if len(area) == 0 {
		return errors.New("you must include an area")
	}

	resp, err := config.pokeapiClient.FullLARequest(&area[0])
	if err != nil {
		return err
	}
	for _, pkmn := range resp.PokemonEncounters {
		fmt.Printf("%v\n", pkmn.Pokemon.Name)
	}
	return nil
}
