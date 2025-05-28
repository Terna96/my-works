//Write a program that prints the name of the program.

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	programName := os.Args[0]
	start := 0

	// Find the last slash to get just the executable name
	for i, char := range programName {
		if char == '/' {
			start = i + 1
		}
	}

	// Print each rune of the program name
	for _, char := range programName[start:] {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
