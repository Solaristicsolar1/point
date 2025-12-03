package main

import (
	"fmt"
)

func IsCapitalized(s string) bool {
	// Empty string is always false
	if s == "" {
		return false
	}

	// Split string into words based on spaces
	word := ""
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			word += string(s[i])
		}

		// When you hit a space or end of string, check the word
		if (s[i] == ' ' || i == len(s)-1) && word != "" {
			first := word[0]

			// Rule 1: lowercase letter? fail
			if first >= 'a' && first <= 'z' {
				return false
			}

			// Clear for next word
			word = ""
		}
	}

	return true
}

func main() {
	fmt.Println(IsCapitalized("Hello! How are you?"))
	fmt.Println(IsCapitalized("Hello How Are You"))
	fmt.Println(IsCapitalized("Whats 4this 100K?"))
	fmt.Println(IsCapitalized("Whatsthis4"))
	fmt.Println(IsCapitalized("!!!!Whatsthis4"))
	fmt.Println(IsCapitalized(""))
}
