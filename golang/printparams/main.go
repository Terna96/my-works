//Write a program that prints the arguments received in the command line.

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// Override os.Args with your test values
	os.Args = []string{os.Args[0], "hello", "world", "from", "Go"} // [0] is always program name

	args := os.Args[1:] // Get arguments (now your hardcoded ones)

	for _, arg := range args {
		for _, char := range arg {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}
