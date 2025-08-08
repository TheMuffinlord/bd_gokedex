package main

import (
	"errors"
	"fmt"
)

func commandInspect(config *cmdConfig, target ...string) error {
	if len(target) < 1 {
		return errors.New("you must list a pokemon")
	}
	inspectedPkmn, caught := config.Pokedex[target[0]]
	if !caught {
		return errors.New("you have not caught that pokemon")
	}
	fmt.Printf("Name: %v\nHeight: %d\nWeight: %d\n", inspectedPkmn.Name, inspectedPkmn.Height, inspectedPkmn.Weight)
	hp := inspectedPkmn.Stats[0].BaseStat
	attack := inspectedPkmn.Stats[1].BaseStat
	defense := inspectedPkmn.Stats[2].BaseStat
	spAttack := inspectedPkmn.Stats[3].BaseStat
	spDefense := inspectedPkmn.Stats[4].BaseStat
	speed := inspectedPkmn.Stats[5].BaseStat
	types := inspectedPkmn.Types
	fmt.Printf("Stats:\n -hp: %d\n -attack: %d\n -defense: %d\n -special-attack: %d\n -special-defense: %d\n -speed: %d\n", hp, attack, defense, spAttack, spDefense, speed)
	fmt.Print("Types:\n")
	for t := range types {
		fmt.Printf(" - %v\n", types[t].Type.Name)
	}
	return nil
}
