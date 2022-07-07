package main

import "fmt"

func main() {
	// compute len of string without len
	var s string

	fmt.Print("input a string: ")
	fmt.Scanf("%s\n", &s)

	var len = 0
	for range s {
		len += 1
	}
	fmt.Printf("len: %v\n", len)

}
