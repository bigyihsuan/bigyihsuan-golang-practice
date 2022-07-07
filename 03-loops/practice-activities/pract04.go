package main

import "fmt"

func main() {
	// sum of numbers between 0 and 100
	sum := 0
	for i := 0; i <= 100; i++ {
		sum += i
	}
	fmt.Printf("%d\n", sum)
}
