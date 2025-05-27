//Write a function that returns the last rune of a string.

package main

import "fmt"

func lastRune(s string) rune {
	if len(s) == 0 {
		return 0
	}
	var last rune
	for _, r := range s {
		last = r
	}
	return last
}

func main() {
	fmt.Println(string(lastRune("Godwin")))
}
