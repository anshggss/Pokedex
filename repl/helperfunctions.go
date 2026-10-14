package repl

import (
	"strings"
)

func cleanInput(text string) []string {
	// clean the input to separate different words and add them to a string slice
	temp := strings.Fields(strings.ToLower(text))
	return temp
}
