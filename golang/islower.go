//Write a function that returns true if the string passed as the parameter contains only
//lowercase characters, otherwise, returns false.

package main

import "fmt"

func islower(s string) bool {
	res := len(s)

	for i := 0; i < res; i++ {
		k := s[i]
		if k >= 'A' && k <= 'Z' {
			return false
		}
	}
	return true
}

func main() {
	s := "Terna"

	fmt.Println(islower(s))
}
