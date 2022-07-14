package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	numbers := make(chan int)
	go generate(numbers)
	isPrime(numbers)
	switch {
	}
}

// write 2 go routines:

// - generate numbers
func generate(numbers chan int) {
	fmt.Println("entered generate")
	for n := 0; true; n++ {
		numbers <- n
		// fmt.Println("wrote", n)
		time.Sleep(time.Second / 16)
	}
}

// - check if a number is prime
func isPrime(numbers chan int) {
	for {
	start:
		n := <-numbers
		// fmt.Println(n)
		if n == 2 {
			fmt.Printf("%v is prime\n", n)
			goto start
		}
		if n <= 1 || n%2 == 0 {
			fmt.Printf("%v is not prime\n", n)
			goto start
		}
		for i := 3; i < int(math.Sqrt(float64(n))); i += 2 {
			if n%i == 0 {
				fmt.Printf("%v is not prime\n", n)
				goto start
			}
		}
		fmt.Printf("%v is prime\n", n)
		time.Sleep(time.Second / 16)
	}
}
