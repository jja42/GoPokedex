package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex >")
		if scanner.Scan() {
			text := scanner.Text()
			clean_input := cleanInput(text)
			word := clean_input[0]
			fmt.Printf("Your command was: %s\n", word)
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
