//Write a function that returns true if the string passed as the parameter
// only contains alphanumerical characters or is empty, otherwise, and returns false.

package main

import "fmt"

func isalpha(s string) bool {
	l := len(s)

	if l == 0 {
		return true
	}

	for i := 0; i < l; i++ {
		r := s[i]
		if (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
			return true
		}
	}
	return false
}

func main() {
	s := "erna"
	fmt.Print(isalpha(s))
}
