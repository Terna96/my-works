//Write a function that returns true if the string passed as a parameter
// contains only printable characters, otherwise, returns false.

package main

import (
	"fmt"
	"unicode"
)

func printAble(s string) bool {

	for _, r := range s {
		if !unicode.IsPrint(r) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(printAble("terna 11 ,./"))
}
