package main

import "fmt"

// func for appending at ramdom selected position
// the "s []int" is the parameter
// the "index int" is the position at which to place the appended value
// the "value int" the value to be appended
func anyAppend(s []int, index int, value int) []int {
	if index < 0 || index > len(s) {
		panic("Index Out of range")
	}

	s = append(s, 0)
	copy(s[index+1:], s[index:])
	s[index] = value
	return s
}

func main() {
	//Appent item at the end
	s := []int{1, 2, 3}
	s = append(s, 5)

	fmt.Println(s)

	//Appending in front

	item := []int{0, 1, 2, 3, 0}
	item = append([]int{2}, item...)

	fmt.Println(item)

	// for appending at a selected possition
	slice := []int{10, 20, 30, 40}

	// Insert 25 at index 2
	slice = anyAppend(slice, 3, 25)

	fmt.Println(slice)

}
