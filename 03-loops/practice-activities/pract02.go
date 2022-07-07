package main

import (
	"fmt"
	"strconv"
)

func main() {
	// priont all even numbers from 1 to 100, and then the odds from 1 to 100
	fmt.Println("EVENS")
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println(strconv.Itoa(i))
		}
	}
	fmt.Println("ODDS")
	for i := 1; i <= 100; i++ {
		if i%2 != 0 {
			fmt.Println(strconv.Itoa(i))
		}
	}
}
