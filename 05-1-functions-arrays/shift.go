package main

import (
	"fmt"
	"strings"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var shift int
	fmt.Print("shift amount: ")
	fmt.Scan(&shift)
	shift %= len(arr)

	var direction string
	fmt.Print("direction: ")
	fmt.Scan(&direction)
	direction = strings.ToLower(direction)

	// shifting:
	// take all elements up to the shift position
	// - shift position depends on direction
	// - left = given, right = len - given
	// make a copy
	// swap the halves
	// var left, right []int
	if direction == "left" {
		arr = append(arr[shift:], arr[:shift]...)
	} else {
		arr = append(arr[len(arr)-shift:], arr[:len(arr)-shift]...)
	}
	fmt.Printf("arr: %v\n", arr)
}
