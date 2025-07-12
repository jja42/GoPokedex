package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	apireq "github.com/jja42/GoPokedex/internal/api_req"
	pokecache "github.com/jja42/GoPokedex/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*apireq.Config, []string) error
}

var commands map[string]cliCommand

var config *apireq.Config

var api_cache *pokecache.Cache

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
		"explore": {
			name:        "explore",
			description: "Explores a specific area. Please provide area name as it appears in map as a parameter",
			callback:    commandExplore,
		},
	}

	api_cache = pokecache.NewCache(5 * time.Second)

	config = &apireq.Config{
		NextURL: nil,
		PrevURL: nil,
		Cache:   api_cache,
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
			params := clean_input[1:]
			handleinput(input, commands, config, params)
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
