package main

import (
	pokeapi "bd_gokedex/internal/pokeapi"
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

var supportedCmds map[string]cliCommand

type cliCommand struct {
	name        string
	description string
	callback    func(*cmdConfig) error
}

type cmdConfig struct {
	pokeapiClient pokeapi.Client
	PrevURL       *string
	NextURL       *string
}

func getCmds() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays a list of 20 regions, counting up",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays a list of 20 regions, counting down",
			callback:    commandMapB,
		},
	}
}

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	config := &cmdConfig{
		pokeapiClient: pokeClient,
	}
	startRepl(config)
}

func startRepl(config *cmdConfig) {
	cliScanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		cliScanner.Scan()
		newCmd := cliScanner.Text()
		cleanCmd := cleanInput(newCmd)
		//fmt.Printf("Your command was: %v\n", cleanCmd[0])
		command, exists := getCmds()[cleanCmd[0]]
		if exists {
			err := command.callback(config)
			if err != nil {
				fmt.Printf("Error: %v", err)
			}
		} else {
			fmt.Print("Unknown command\n")
		}
	}
}

func cleanInput(text string) []string {
	trimText := strings.TrimSpace(text)
	lowerText := strings.ToLower((trimText))
	returnString := strings.Fields(lowerText)
	return returnString
}

func commandExit(config *cmdConfig) error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	defer os.Exit(0)
	return nil
}

func commandHelp(config *cmdConfig) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for key := range supportedCmds {
		fmt.Printf("%s: %s\n", supportedCmds[key].name, supportedCmds[key].description)
	}
	return nil
}

func commandMap(config *cmdConfig) error {
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

func commandMapB(config *cmdConfig) error {
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
