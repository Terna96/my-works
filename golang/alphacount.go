package main

import (
	"fmt" // For formatted I/O (used for printing to console)

	"github.com/01-edu/z01" // For printing runes one by one (used in 01 school projects)
)

// CountLatinLetters counts how many Latin alphabet letters (A-Z, a-z) are in the string s
func CountLatinLetters(s string) int {
	count := 0
	for _, r := range s {
		// Check if rune r is an uppercase or lowercase Latin letter
		if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
			count++ // Increment count for each Latin letter found
		}
	}
	return count // Return the total count of Latin letters
}

func main() {
	str := "my, name! is Terna 123" // Example input string containing letters, digits, symbols, emoji

	// Count the Latin letters in the string
	count := CountLatinLetters(str)

	// Print the result with fmt.Println (prints "Letter count: X")
	fmt.Println("Letter count:", count)

	// Optionally print the count digit-by-digit using z01.PrintRune
	// First convert the integer count to a string, then print each character as a rune
	for _, r := range fmt.Sprintf("%d", count) {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n') // Print a newline character at the end
}
