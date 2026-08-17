package main

import "fmt"

func commandHelp() error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!\nUsage:\n")
	for key, cmd := range commandMap {
		fmt.Printf("%s: %s\n", key, cmd.description)
	}
	fmt.Println()
	return nil
}
