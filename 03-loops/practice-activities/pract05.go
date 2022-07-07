package main

import "fmt"

func main() {
	// numbers divisible by 50 between 100 and 1000
	// with range

	// without range
	for i := 100; i <= 1000; i++ {
		if i%50 == 0 {
			fmt.Printf("%d\n", i)
		}
	}
}
