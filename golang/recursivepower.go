//Write a recursive function that returns the value of nb to the power of power.

//Negative powers will return 0. Overflows do not have to be dealt with.

//for is forbidden for this exercise.

package main

import "fmt"

func recursivepower(nb int, power int) int {
	//negative powers to return 0
	if power < 0 {
		return 0
	}
	//since any number raised by 0 is 1
	if power == 0 {
		return 1
	}

	return nb * recursivepower(nb, power-1)

}

func main() {
	fmt.Println(recursivepower(3, 3))
	fmt.Println(recursivepower(3, 0))
}
