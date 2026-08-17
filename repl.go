package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name 		string
	description	string
	callback	func() error
}

var commandMap map[string]cliCommand

func initCliCommands() {
	commandMap = map[string]cliCommand{
		"exit": {
			name:		"exit",
			description:"Exit the Pokedex",
			callback:	commandExit,
		},
		"help": {
			name: 		"help",
			description:"Displays a help message",
			callback: commandHelp,
		},
	}
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	initCliCommands()

	for true {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		cleanedInput := cleanInput(scanner.Text())
		if len(cleanedInput) == 0 {
			continue
		}

		commandName := cleanedInput[0]
		command, found := commandMap[commandName]
		if !found {
			fmt.Println("Unknown command")
		} else {
			command.callback()
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!\nUsage:\n")
	for key, cmd := range commandMap {
		fmt.Printf("%s: %s\n", key, cmd.description)
	}
	fmt.Println()
	return nil
}