package main

import "fmt"

func PrintMemory(arr [10]byte) {
	hex := "0123456789abcdef"

	// Print hex values (4 bytes per line)
	for i, b := range arr {
		high := b >> 4
		low := b & 0x0F

		fmt.Print(string(hex[high]))
		fmt.Print(string(hex[low]))

		// spacing and newlines
		if (i+1)%4 == 0 {
			fmt.Print("\n")
		} else if i != len(arr)-1 {
			fmt.Print(" ")
		}
	}

	// If byte count not divisible by 4 → ensure newline
	if len(arr)%4 != 0 {
		fmt.Print("\n")
	}

	// Print ASCII characters
	for _, b := range arr {
		if b >= 32 && b <= 126 {
			fmt.Print(string(b))
		} else {
			fmt.Print(".")
		}
	}

	// End cleanly with a newline using Print
	fmt.Print("\n")
	return
}

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}
