//Write a program that prints the arguments received in the command line in reverse order.

package main

import (
	"os"

	"github.com/01-edu/z01"
)

/*
func main() {
	// Get command line arguments (excluding program name)
	args := os.Args[1:]

	// Iterate through arguments in reverse order
	for i := len(args) - 1; i >= 0; i-- {
		// Print each rune of the argument
		for _, r := range args[i] {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}

*/

func main() {
	// Get all command-line arguments except the program name
	args := os.Args[1:]

	// If no arguments provided, use these defaults
	if len(args) == 0 {
		args = []string{"No", "arguments", "provided"}
	}

	// Print arguments in reverse order
	for i := len(args) - 1; i >= 0; i-- {
		// Print each character of the argument
		for _, char := range args[i] {
			z01.PrintRune(char)
		}
		// Print newline after each argument
		z01.PrintRune('\n')
	}
}
