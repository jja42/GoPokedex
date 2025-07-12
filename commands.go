package main

import (
	"fmt"
	"os"

	apireq "github.com/jja42/GoPokedex/internal/api_req"
)

func handleinput(input string, commands map[string]cliCommand, config *apireq.Config) {
	if command, exists := commands[input]; exists {
		command.callback(config)
	}
}

func commandExit(config *apireq.Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *apireq.Config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for name, info := range commands {
		str := fmt.Sprintf("%s: %s", name, info.description)
		fmt.Println(str)
	}
	return nil
}

func commandMap(config *apireq.Config) error {
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

	config.PrevURL = batch.Previous
	config.NextURL = batch.Next

	return nil
}

func commandMapb(config *apireq.Config) error {
	url := config.PrevURL
	if url == nil {
		fmt.Println("You're on the first page")
		return nil
	}
	config.NextURL = config.PrevURL
	commandMap(config)
	return nil
}
