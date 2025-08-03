package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var supportedCmds map[string]cliCommand

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func main() {
	supportedCmds = map[string]cliCommand{
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
	}
	cliScanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		cliScanner.Scan()
		newCmd := cliScanner.Text()
		cleanCmd := cleanInput(newCmd)
		//fmt.Printf("Your command was: %v\n", cleanCmd[0])
		command, exists := supportedCmds[cleanCmd[0]]
		if exists {
			err := command.callback()
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

func commandExit() error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	defer os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for key := range supportedCmds {
		fmt.Printf("%s: %s\n", supportedCmds[key].name, supportedCmds[key].description)
	}
	return nil
}
