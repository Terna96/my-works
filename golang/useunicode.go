package main

import (
	"fmt"
	"unicode"
)

func IsPrintable(s string) bool {

	for _, r := range s {
		if unicode.IsDigit(r) {
			return true
		}
	}

	return false
}

func main() {
	s := "terna"
	fmt.Println(IsPrintable(s))
}
