package main

import "fmt"

func Compare(a, b string) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return 1
}

func main() {
	fmt.Println(Compare("Hello!", "abllo!")) // 0
	fmt.Println(Compare("mybrother!", "abllo!"))

}
