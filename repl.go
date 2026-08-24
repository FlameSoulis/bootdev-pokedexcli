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
	callback	func(*config) error
}

type config struct {
	commands 	map[string]cliCommand
}

func initCliCommands() map[string]cliCommand {
	return map[string]cliCommand{
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

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	//initCliCommands()

	for true {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		cleanedInput := cleanInput(scanner.Text())
		if len(cleanedInput) == 0 {
			continue
		}

		commandName := cleanedInput[0]
		command, found := cfg.commands[commandName]
		if !found {
			fmt.Println("Unknown command")
		} else {
			command.callback(cfg)
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}