package main

import "fmt"

func main() {
	// input string, print each char on separate line
	var s string

	fmt.Print("input a string: ")
	fmt.Scanf("%s\n", &s)

	for _, c := range s {
		fmt.Println(string(c))
	}
}
