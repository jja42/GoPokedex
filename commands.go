package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	apireq "github.com/jja42/GoPokedex/api_req"
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
	url := config.NextURL
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area/0"
	}

	lastSlash := strings.LastIndex(url, "/")
	baseurl := url[:lastSlash+1]
	id := url[lastSlash+1:]
	num, _ := strconv.Atoi(id)

	for i := 0; i < 20; i++ {
		num++
		id = strconv.Itoa(num)
		full_url := baseurl + id
		data := apireq.GetRequest(full_url)
		location, _ := apireq.RequestToLocation(data)
		fmt.Println(location.Name)
	}

	if num > 39 {
		id = strconv.Itoa(num - 39)
		full_url := baseurl + id
		config.PrevURL = full_url
	} else {
		config.PrevURL = ""
	}

	id = strconv.Itoa(num)
	full_url := baseurl + id
	config.NextURL = full_url

	return nil
}

func commandMapb(config *apireq.Config) error {
	url := config.PrevURL
	if url == "" {
		fmt.Println("You're on the first page")
		return nil
	}
	config.NextURL = config.PrevURL
	commandMap(config)
	return nil
}
