package main

import (
	"fmt"
	"math"
)

func main() {
	// get int, print reverse
	var a, b int

	fmt.Print("enter a: ")
	fmt.Scanf("%d", &a)
	var numDigits = 1
	for a/int(math.Pow10(numDigits)) > 0 {
		numDigits += 1
		// fmt.Printf("numDigits: %v\n", numDigits)
		// fmt.Printf("a: %v\n", a/(numDigits*10))
	}

	reversed := 0
	b = a
	for i := numDigits - 1; i >= 0; i-- {
		reversed += (b % 10) * int(math.Pow10(i))
		b /= 10
		// fmt.Printf("b: %v\n", b)
		// fmt.Printf("reversed: %v\n", reversed)
	}
	fmt.Printf("a: %v\n", a)
	fmt.Printf("reversed: %v\n", reversed)
}
