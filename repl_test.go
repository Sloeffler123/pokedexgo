package main

import (
	"testing"
	"fmt"
)

func TestPokeDexCli(t *testing.T,) {
	cases := []struct {
		input string
		expected string
	}{
		{
			input: "help",
			expected: "Display a help message",
		},
		{
			input: "exit",
			expected: "Exit the Pokedex",
		},
		{
			input: "Help me",
			expected: "Display a help message",
		},
		{
			input: "Exit",
			expected: "Exit the Pokedex",
		},
		{
			input: "",
			expected: "Unknown command",
		},
	}
	for _, c := range cases {
		if c.input == "" {
			fmt.Println("Unknown command")
			continue
		}
		cleanText := cleanInput(c.input)
		oneWord := cleanText[0]
		if cliCommands()[oneWord].description != c.expected {
			t.Errorf("texts dont match")
			t.Fail()
		}
	}
}

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: " hello  world ",
			expected: []string{"hello","world"},
		},
		{
			input: " charmander squritle bulbasaur",
			expected: []string{"charmander", "squritle", "bulbasaur"},
		},
		{
			input: " ",
			expected: []string{},
		},
		{
			input: "HEllo WORld ",
			expected: []string{"hello","world"},
		},
		{
			input: " hello ",
			expected: []string{"hello"},
		},

	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Slices dont match")
			t.Fail()
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("words dont match")
				t.Fail()
			}
		}
	}
}