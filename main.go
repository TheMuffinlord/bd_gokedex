package main

import (
	pokeapi "bd_gokedex/internal/pokeapi"
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

var supportedCmds map[string]cliCommand

type cliCommand struct {
	name        string
	description string
	callback    func(*cmdConfig, ...string) error
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
		"explore": {
			name:        "explore",
			description: "Lists pokemon in an area. Requires an area to return results.",
			callback:    commandExplore,
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
		params := []string{}
		if len(cleanCmd) > 1 {
			params = cleanCmd[1:]
		}
		if exists {
			err := command.callback(config, params...)
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
