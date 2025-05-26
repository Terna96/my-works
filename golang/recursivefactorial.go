//Write a recursive function that returns the factorial of the int passed as parameter.

//Errors (non possible values or overflows) will return 0.

//for is forbidden for this exercise.

package main

import (
	"fmt"
	"math"
)

func recursivefactorial(n int) int {
	//handle -ve values
	if n < 0 {
		return 0
	}

	// Base case: 0! = 1 and 1! = 1
	if n == 0 || n == 1 {
		return 1
	}

	// Check for overflow before recursing
	prev := recursivefactorial(n - 1)
	if prev == 0 || prev > math.MaxInt64/n {
		return 0
	}

	return n * prev
}

func main() {
	fmt.Println(recursivefactorial(5))
	fmt.Println(recursivefactorial(4))
	fmt.Println(recursivefactorial(12))
}
