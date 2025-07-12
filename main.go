package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	apireq "github.com/jja42/GoPokedex/internal/api_req"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*apireq.Config) error
}

var commands map[string]cliCommand

var config *apireq.Config

const (
	baseURL = "https://pokeapi.co/api/v2"
)

func init() {
	commands = map[string]cliCommand{
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
			description: "Displays next 20 location areas in the Pokemon world.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous 20 location areas in the Pokemon world",
			callback:    commandMapb,
		},
	}

	config = &apireq.Config{
		NextURL: nil,
		PrevURL: nil,
	}
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex >")
		if scanner.Scan() {
			text := scanner.Text()
			clean_input := cleanInput(text)
			input := clean_input[0]
			handleinput(input, commands, config)
		}
	}
}

func cleanInput(text string) []string {
	words := strings.Fields(text)
	clean_words := make([]string, 0)
	for _, word := range words {
		clean := strings.ToLower((word))
		clean_words = append(clean_words, clean)
	}
	return clean_words
}
