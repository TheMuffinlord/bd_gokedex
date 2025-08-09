package main

import (
	"errors"
	"fmt"
)

func commandPokedex(config *cmdConfig, notused ...string) error {
	if len(config.Pokedex) < 1 {
		return errors.New("no entries in Pokedex")
	}
	fmt.Print("Your Pokedex:\n")
	for p := range config.Pokedex {
		pokemon := config.Pokedex[p]
		fmt.Printf(" - %v\n", pokemon.Name)
	}
	return nil
}
