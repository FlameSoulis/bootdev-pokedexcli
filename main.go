package main

func main() {
	var cfg config
	cfg.commands = initCliCommands()
	startRepl(&cfg)
}

