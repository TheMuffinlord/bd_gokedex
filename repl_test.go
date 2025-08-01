package main

import (
	"fmt"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "how now brown cow",
			expected: []string{"how", "now", "brown", "cow"},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
	}

	for n, c := range cases {
		actual := cleanInput(c.input)
		fmt.Printf("%q/n", actual)
		aLength := len(actual)
		eLength := len(c.expected)
		if aLength != eLength {
			t.Errorf("TEST %d FAIL: string lengths do not match\nexpected: %d; actual: %d", n+1, eLength, aLength)
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			if word != expectedWord {
				t.Errorf("TEST %d FAIL: words do not match\nexpected: %v; actual: %v", n+1, expectedWord, word)
			}
		}
	}
}
