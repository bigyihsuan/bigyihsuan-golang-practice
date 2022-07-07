package main

import "fmt"

func main() {
	// reverse string
	var s string

	fmt.Print("input a string: ")
	fmt.Scanf("%s\n", &s)

	var r = ""
	for i := len(s) - 1; i >= 0; i-- {
		r += string(s[i])
	}
	fmt.Printf("r: %v\n", r)

}
