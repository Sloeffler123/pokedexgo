package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func main() {
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

func cleanInput(text string) []string {
	words := strings.ToLower(text)
	cleanText := strings.Fields(words)
	return cleanText
}
