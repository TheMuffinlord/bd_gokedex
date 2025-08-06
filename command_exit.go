package main

import (
	"fmt"
	"os"
)

func commandExit(config *cmdConfig, params ...string) error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	defer os.Exit(0)
	return nil
}
