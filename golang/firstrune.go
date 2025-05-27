//Write a function that returns the first rune of a string.

package main

import "fmt"

func FirstRune(s string) rune {
	//A rune is an alias for int32 and represents a single Unicode character (like 'A', 'ñ').
	if len(s) == 0 {
		return 0
	}
	return []rune(s)[1]
}

//or to pick the first rune of a string converted into slice

func main() {
	fmt.Printf("%c\n", FirstRune("Pnut"))  // 'H'
	fmt.Printf("%c\n", FirstRune("Terna")) // 'T'
}
