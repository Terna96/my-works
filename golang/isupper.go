//Instructions

//Write a function that returns true if the string passed as parameter contains only uppercase
// characters, otherwise, returns false.

package main

import "fmt"

func IsUpper(s string) bool {

	for i := 0; i < len(s); i++ {
		res := s[i]
		if res >= 'a' && res <= 'z' {
			return false
		}
	}
	return true
}

func main() {
	s := "TERNA"
	fmt.Println(IsUpper(s))
}
