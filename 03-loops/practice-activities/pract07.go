package main

import "fmt"

func main() {
	// ask for x,y
	// compute x^y with only for loops
	var x, y int

	fmt.Print("enter x: ")
	fmt.Scanf("%d", &x)
	fmt.Print("enter y: ")
	fmt.Scanf("%d", &y)

	prod := 1
	for i := 0; i < y; i++ {
		prod *= x
	}
	fmt.Printf("%v ^ %v = %v\n", x, y, prod)
}
