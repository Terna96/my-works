//Write a function that returns the nth rune of a string. If not possible, it returns 0

package main

import "fmt"

func NRune(s string, n int) rune {
	runes := []rune(s) // Convert string to a slice of runes
	if n < 0 || n >= len(runes) {
		return 0 // Return 0 if index is invalid
	}
	return runes[n] // Return the nth rune
}

func main() {
	str := "Hello"

	// Test cases
	fmt.Printf("%c\n", NRune(str, 3)) // 'H' (1st rune)
}
