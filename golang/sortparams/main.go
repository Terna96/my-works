//Write a program that prints the arguments received in the command line in ASCII order.

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// Get command line arguments (excluding program name)
	args := os.Args[1:]

	// Simple bubble sort implementation for ASCII ordering
	for i := 0; i < len(args); i++ {
		for j := i + 1; j < len(args); j++ {
			if args[i] > args[j] {
				args[i], args[j] = args[j], args[i]
			}
		}
	}

	// Print the sorted arguments
	for _, arg := range args {
		for _, r := range arg {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}
