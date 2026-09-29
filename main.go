package main

func main() {
	var cfg config
	cfg.commands = initCliCommands()
	cfg.nextURL = "https://pokeapi.co/api/v2/location-area/"
	startRepl(&cfg)
}

