//Write a function that returns a concatenated string from the 'strings' passed as arguments.

package main

import "fmt"

func concat(s ...string) string {
	var res string

	for _, r := range s {
		res += r
	}
	return res
}

func main() {
	fmt.Println(concat("Terna ", "is ", "a ", "good ", "Boy"))
}
