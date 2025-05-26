//Write an iterative function that returns the value of nb to the power of power.

//Negative powers will return 0. Overflows do not have to be dealt with.

package main

import "fmt"

func iteratepower(n int, power int) int {
	//negative powers to return 0
	if power < 0 {
		return 0
	}

	// Any number to the power of 0 is 1
	result := 1

	// Multiply nb 'power' times
	for i := 1; i < power; i++ {
		result *= n
	}
	return result
}

func main() {
	fmt.Println(iteratepower(3, 2))  //(2³)
	fmt.Println(iteratepower(3, -2)) // 0 (negative power)
}
