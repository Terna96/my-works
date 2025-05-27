/*Write a function that returns true if the int passed as parameter is a prime number.
 Otherwise it returns false.

The function must be optimized in order to avoid time-outs with the tester.

(We consider that only positive numbers can be prime numbers)

(We also consider that 1 is not a prime number)
*/

package main

import "fmt"

func isprime(nb int) bool {
	if nb < 1 {
		return false
	}

	if nb == 2 {
		return true
	}

	if nb%2 == 0 {
		return false
	}

	for i := 3; i*i <= nb; i += 2 {
		if nb%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isprime(2))
}
