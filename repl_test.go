package main

import (
	"testing"
	
)


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