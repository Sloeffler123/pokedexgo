package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello World")
	
}

func cleanInput(text string) []string {
	words := strings.ToLower(text)
	cleanText := strings.Fields(words)
	return cleanText
}
