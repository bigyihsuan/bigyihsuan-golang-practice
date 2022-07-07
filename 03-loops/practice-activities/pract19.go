package main

import "fmt"

func main() {
	// ask for 5 numbers
	var pos, zero, neg int

	fmt.Print("input 5 numbers: ")

	var n int
	for i := 0; i < 5; i++ {
		fmt.Scanf("%d", &n)
		if n > 0 {
			pos++
		} else if n < 0 {
			neg++
		} else {
			zero++
		}

	}
	// The number of positive numbers that the user entered
	fmt.Printf("pos: %v\n", pos)
	// The number of zeros that the user entered
	fmt.Printf("zero: %v\n", zero)
	// The number of negative numbers
	fmt.Printf("neg: %v\n", neg)

}
