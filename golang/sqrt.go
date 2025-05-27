//Write a function that returns the square root of the int passed as parameter,
//if that square root is a whole number. Otherwise it returns 0.

package main

import "fmt"

func Sqrt(nb int) int {
	if nb < 1 {
		return 0
	}

	for i := 1; i*i <= nb; i++ {
		if i*i == nb {
			return i
		}
	}

	return 0
}

func main() {
	fmt.Println(Sqrt(16))
}
