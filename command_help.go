package main

import "fmt"

func commandHelp(config *cmdConfig, param []string) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for key := range supportedCmds {
		fmt.Printf("%s: %s\n", supportedCmds[key].name, supportedCmds[key].description)
	}
	return nil
}
