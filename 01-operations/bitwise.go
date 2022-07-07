package main

import "fmt"

func main() {
	x := 11
	y := 22
	fmt.Printf("x     = %#08b\n", x)
	fmt.Printf("y     = %#08b\n", y)
	fmt.Printf("x ^ y = %#08b\n", x^y)
	fmt.Printf("x + y = %#08b\n", x+y)

	fmt.Printf("x & y = %#08b\n", x&y)
	fmt.Printf("x | y = %#08b\n", x|y)
	fmt.Printf("x << 3 = %#08b\n", x>>3)
	fmt.Printf("x >> 3 = %#08b\n", x<<3)
}
