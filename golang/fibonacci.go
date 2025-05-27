/*Write a recursive function that returns the value at the position index in the fibonacci sequence.

The first value is at index 0.

The sequence starts this way: 0, 1, 1, 2, 3 etc...

A negative index will return -1.

for is forbidden for this exercise.
*/

package main

import "fmt"

func fibonacci(index int) int {
	if index < 0 {
		return -1
	}

	if index == 0 {
		return 0
	}
	if index == 1 {
		return 1
	}

	return fibonacci(index-1) + fibonacci(index-2)
}

func main() {
	fmt.Println(fibonacci(2))
}
