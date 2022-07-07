package main

import (
	"fmt"
	"math"
)

func main() {
	// Yi-Hsuan Hsu
	// ask for n
	// calculate sine series
	// sin x = x - x^3/3! + x^5/5! - ...
	var term uint64
	var sum, n float64

	fmt.Print("enter n: ")
	fmt.Scanf("%f", &n)
	fmt.Print("enter term count: ")
	fmt.Scanf("%d", &term)

	// 0 1 = -1^0 * 1
	// 1 -3 = -1^1 * 3
	// 2 5 = -1^2 * 5
	// 3 -7 = -1^3 * 7
	for i := 0; i < int(term); i++ {
		exp := float64(2*i + 1)          // odd numbers only
		sign := math.Pow(-1, float64(i)) // positive on evens, negative on odds
		num := math.Pow(n, exp)
		fac := factorial(exp)
		term := sign * num / fac
		fmt.Printf("term: %v\n", term)
		sum += term
	}
	fmt.Printf("sum: %v\n", sum)
}

func factorial(n float64) float64 {
	var fac, f float64 = 1, 1
	for f = 1; f <= n; f++ {
		fac *= f
	}
	return fac
}
