package main

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n < 0 {
		return // Handle negative as needed
	}

	// Special case for 0
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	// Array to count frequency of each digit 0–9
	count := [10]int{}

	// Count digit occurrences
	for n > 0 {
		d := n % 10
		count[d]++
		n /= 10
	}

	for digit := 0; digit < 10; digit++ {
		for i := 0; i < count[digit]; i++ {
			z01.PrintRune(rune(digit + '0'))
		}
	}
}

func main() {
	PrintNbrInOrder(8790467453)
	z01.PrintRune('\n')
}
