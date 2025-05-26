//Write a recursive function that returns the factorial of the int passed as parameter.

//Errors (non possible values or overflows) will return 0.

//for is forbidden for this exercise.

package main

import (
	"fmt"
)

func recursivefactorial(n int) int {
	if n < 0 || n > 20 {
		return 0
	}
	if n == 0 || n == 1 {
		return 1
	}

	return n * recursivefactorial(n-1)
}

func main() {
	fmt.Println(recursivefactorial(5))
	fmt.Println(recursivefactorial(4))
	fmt.Println(recursivefactorial(12))
}
