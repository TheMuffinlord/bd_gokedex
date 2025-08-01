package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	cliScanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		cliScanner.Scan()
		newCmd := cliScanner.Text()
		cleanCmd := cleanInput(newCmd)
		fmt.Printf("Your command was: %v\n", cleanCmd[0])
	}
}

func cleanInput(text string) []string {
	trimText := strings.TrimSpace(text)
	lowerText := strings.ToLower((trimText))
	returnString := strings.Fields(lowerText)
	return returnString
}
