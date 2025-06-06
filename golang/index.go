package main

import "fmt"

func Index(s string, toFind string) int {
	if toFind == "" {
		return 0
	}
	if len(toFind) > len(s) {
		return -1
	}

	for i := 0; i <= len(s)-len(toFind); i++ {
		match := true
		for j := 0; j < len(toFind); j++ {
			if s[i+j] != toFind[j] {
				match = false
				break
			}
		}
		if match {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(Index("hello world", "world")) // ➜ 6
	fmt.Println(Index("hello world", "lo"))    // ➜ 3
	fmt.Println(Index("hello world", "go"))    // ➜ -1
	fmt.Println(Index("abc", ""))              // ➜ 0
	fmt.Println(Index("", "a"))                // ➜ -1
}
