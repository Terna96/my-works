//Write an iterative function that returns the factorial of the int passed as parameter.

//Errors (non possible values or overflows) will return 0.

package main

import "fmt"

func iteratefact(T int) int {
	result := 1
	for i := 1; i <= T; i++ {
		result *= i
	}
	return result
}

func main() {
	num := 5
	fmt.Println("the factorial of", num, "is:", iteratefact(num))
}
