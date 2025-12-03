package main

import (
	"fmt"
	"strconv"
)

func ZipString(s string) string {
	if s == "" {
		return ""
	}

	result := ""
	count := 1

	for i := 0; i < len(s); i++ {
		// count duplicates
		if i+1 < len(s) && s[i] == s[i+1] {
			count++
		} else {
			// convert number to string with strconv.Itoa
			result += strconv.Itoa(count)
			result += string(s[i])
			count = 1
		}
	}

	return result
}

func main() {
	fmt.Print(ZipString("YouuungFellllas"))
	fmt.Print("\n")
	fmt.Print(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
	fmt.Print("\n")
	fmt.Print(ZipString("Helloo Therre!"))
	fmt.Print("\n")
}
