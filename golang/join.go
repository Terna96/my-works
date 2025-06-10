//Write a function that returns the concatenation of all the strings of a slice of strings
// separated by the separator passed as the argument sep.

package main

import "fmt"

func Join(s []string, sep string) string {
	res := ""

	for i, r := range s {
		res += r
		if i < len(s)-1 {
			res += sep
		}
	}
	return res
}

func main() {
	words1 := []string{"Go", "is", "awesome"}

	fmt.Println(Join(words1, "_"))
}
