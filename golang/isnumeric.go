//Write a function that returns true if the string passed as a parameter contains
// only numerical characters, otherwise, returns false.

package main

import "fmt"

func IsNumeric(s string) bool {

	for _, res := range s {

		if res >= '0' && res <= '9' {
			return true
		} else {
			return false
		}
	}

	return true
}

func main() {
	s := "996"
	fmt.Println(IsNumeric(s))
}
