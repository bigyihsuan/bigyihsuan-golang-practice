package main

import (
	"fmt"
	"math"
)

func main() {
	// lowest common divisor
	var a, b int

	fmt.Print("enter a: ")
	fmt.Scanf("%d", &a)
	fmt.Print("enter b: ")
	fmt.Scanf("%d", &b)

	gcd := -1
	for i := int(math.Min(float64(a), float64(b))); i >= 1; i-- {
		if a%i == 0 && b%i == 0 {
			gcd = i
		}
	}
	fmt.Printf("GCD for %v and %v = %v\n", a, b, gcd)
}
