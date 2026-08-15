package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for true {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		inputedText := scanner.Text()
		cleanedInput := cleanInput(inputedText)
		fmt.Println("Your command was:", cleanedInput[0])
	}
}

