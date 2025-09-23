package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func cliCommands() map[string]cliCommand {
	return map[string]cliCommand {
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"help": {
			name: "help",
			description: "Display a help message",
			callback: commandHelp,
		},
	}
}


func pokeDexStart() {
	userInput := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		userInput.Scan() 
		textUserInput := userInput.Text()
		if len(textUserInput) == 0 {
			continue
		}
		cleanText := cleanInput(textUserInput)
		userWord := cleanText[0]
		command, exists := cliCommands()[userWord]
		if exists {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	words := strings.ToLower(text)
	cleanText := strings.Fields(words)
	return cleanText
}

func commandExit() error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("\nhelp: Display a help message")
	fmt.Println("exit: Exit the Pokedex")
	return nil
}
