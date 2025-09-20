package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name string
	description string
	callback func() error
}

func exitDict() map[string]cliCommand{
	exit := map[string]cliCommand{
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
	}
	return exit
}

func helpDict() map[string]cliCommand{
	help := map[string]cliCommand{
		"help": {
			name: "help",
			description: "Welcome to the Pokedex!",
			callback: commandHelp,
		},
	}
	return help
}

func pokeDexStart() {
	for {
		fmt.Print("Pokedex > ")
		userInput := bufio.NewScanner(os.Stdin)
		userInput.Scan()
		textUserInput := userInput.Text()
		textToLowed := strings.ToLower(textUserInput)
		cleanTextUserInput := strings.Fields(textToLowed)
		fmt.Println("Your command was:", cleanTextUserInput[0])
	}
}

func commandExit() error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Usage:")
	fmt.Println("\nhelp: Display a help message")
	fmt.Println("exit Exit the Pokedex")
	os.Exit(0)
	return nil
}

func cleanInput(text string) []string {
	words := strings.ToLower(text)
	cleanText := strings.Fields(words)
	return cleanText
}
