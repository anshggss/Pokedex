package repl

import (
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
			input:    "accha laudeaisa kya",
			expected: []string{"accha", "laudeaisa", "kya"},
		},
		{
			input:    "okay gummy bear this isit",
			expected: []string{"okay", "gummy", "bear", "this", "isit"},
		},
	}

	for _, c := range cases {

		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("Repl test failed! Lengths don't match")
		}

		for i := range actual {
			word := actual[i]
			expected := c.expected[i]
			if expected != word {
				t.Errorf("Repl test failed!\nExpected word: %v\nActual word: %v\n", c.expected[i], word)
			}
		}
	}
}
