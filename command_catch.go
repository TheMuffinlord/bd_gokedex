package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func commandCatch(config *cmdConfig, target ...string) error {
	if len(target) < 1 {
		return errors.New("you must list a pokemon")
	}
	pkmn, err := config.pokeapiClient.PkmnLookup(&target[0])
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %v...\n", pkmn.Name)

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	catchChance := rng.Intn(pkmn.BaseExperience)
	//fmt.Printf("DEBUG: basexp: %d, catch chance: %d\n", pkmn.BaseExperience, catchChance)
	if catchChance < 41 {
		fmt.Printf("You caught a %v!\n", pkmn.Name)
		config.Pokedex[pkmn.Name] = pkmn
	} else {
		fmt.Printf("You couldn't catch the %v; try again!\n", pkmn.Name)
	}
	return nil
}
