package main

import (
	"fmt"
	"math/rand"
	"os"

	apireq "github.com/jja42/GoPokedex/internal/api_req"
)

func handleinput(input string, commands map[string]cliCommand, config *apireq.Config, params []string) {
	if command, exists := commands[input]; exists {
		command.callback(config, params)
	}
}

func commandExit(config *apireq.Config, params []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *apireq.Config, params []string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for name, info := range commands {
		fmt.Printf("%s: %s", name, info.description)
	}
	return nil
}

func commandMap(config *apireq.Config, params []string) error {
	if config.NextURL == nil {
		url := baseURL + "/location-area"
		config.NextURL = &url
	}

	key := *config.NextURL
	var data []byte
	if cached_data, exists := config.Cache.Get(key); exists {
		data = cached_data
	} else {
		data = apireq.GetRequest(*config.NextURL, config)
	}

	batch, err := apireq.RequestToLocations(data)
	if err != nil {
		return err
	}

	for _, location := range batch.Results {
		fmt.Println(location.Name)
	}

	config.Cache.Add(*config.NextURL, data)
	config.PrevURL = batch.Previous
	config.NextURL = batch.Next

	return nil
}

func commandMapb(config *apireq.Config, params []string) error {
	url := config.PrevURL
	if url == nil {
		fmt.Println("You're on the first page")
		return nil
	}
	config.NextURL = config.PrevURL
	commandMap(config, params)
	return nil
}

func commandExplore(config *apireq.Config, params []string) error {
	locationName := params[0]
	url := baseURL + "/location-area/" + locationName

	var data []byte

	if cached_data, exists := config.Cache.Get(url); exists {
		data = cached_data
	} else {
		data = apireq.GetRequest(url, config)
	}

	location, err := apireq.RequestToLocation(data)
	if err != nil {
		return err
	}

	config.Cache.Add(url, data)

	fmt.Printf("Exploring %s...\n", locationName)

	fmt.Println("Found Pokemon:")

	for _, encounter := range location.Pokemon_Encounters {
		fmt.Printf("- %s\n", encounter.Pokemon.Name)
	}

	return nil
}

func commandCatch(config *apireq.Config, params []string) error {

	pokemonName := params[0]
	url := baseURL + "/pokemon/" + pokemonName

	var data []byte

	if cached_data, exists := config.Cache.Get(url); exists {
		data = cached_data
	} else {
		data = apireq.GetRequest(url, config)
	}

	pokemon, err := apireq.RequestToPokemon(data)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	r := rand.Intn(pokemon.Base_Experience)

	if r > 40 {
		fmt.Printf("%s escaped!\n", pokemonName)

	} else {
		fmt.Printf("%s was caught!\n", pokemonName)
		config.Cache.Add(url, data)
		fmt.Printf("%s was added to your Pokedex!\n", pokemonName)
		fmt.Println("You may now inspect it or view it in your Pokedex.")
	}

	return nil
}

func commandInspect(config *apireq.Config, params []string) error {
	pokemonName := params[0]
	url := baseURL + "/pokemon/" + pokemonName

	if cached_data, exists := config.Cache.Get(url); exists {
		pokemon, err := apireq.RequestToPokemon(cached_data)
		if err != nil {
			return err
		}
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %d\n", pokemon.Height)
		fmt.Printf("Weight: %d\n", pokemon.Weight)

		fmt.Println()
		fmt.Println("Stats")

		for _, stat := range pokemon.Stats {
			fmt.Printf("%s: %v\n", stat.Stat.Name, stat.BaseStat)
		}

		fmt.Println()
		fmt.Println("Types")
		for _, typeInfo := range pokemon.Types {
			fmt.Printf("%s\n", typeInfo.Type.Name)
		}

	} else {
		fmt.Println("You have not yet caught that pokemon.")
	}

	return nil
}
