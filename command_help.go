package main

import "fmt"

func commandHelp(cfg *config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!\nUsage:\n")
	for key, cmd := range cfg.commands {
		fmt.Printf("%s: %s\n", key, cmd.description)
	}
	fmt.Println()
	return nil
}
