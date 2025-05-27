package main

import "fmt"

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func findNextPrime(n int) int {
	if n <= 2 {
		return 2
	}
	for {
		if isPrime(n) {
			return n
		}
		n++
	}
}

func main() {
	fmt.Println(findNextPrime(4)) // Output: 17
}
