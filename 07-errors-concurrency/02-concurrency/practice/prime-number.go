package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	numbers := generatePrimes(100)
	primes := extract(numbers)
	time.Sleep(time.Second)
	for _, n := range primes {
		fmt.Println(n)
	}
}

// write 2 go routines:

// - generate numbers
func generatePrimes(n int) chan int {
	fmt.Println("entered generate")
	numbers := make(chan int)
	go func() {
		for i := 2; i <= n; i++ {
			maybe := <-isPrime(i)
			if maybe {
				numbers <- i
			}
			// fmt.Println("wrote", n)
			// time.Sleep(time.Second / 16)
		}
		close(numbers)
	}()
	return numbers
}

// - check if a number is prime
func isPrime(n int) chan bool {
	out := make(chan bool)
	go func() {
		if n < 2 {
			out <- false
			return
		}
		for i := 2; i < int(math.Sqrt(float64(n))); i++ {
			if n%i == 0 {
				out <- false
				return
			}
		}
		out <- true
		close(out)
	}()
	return out
}

func extract(numbers chan int) []int {
	var ns []int
	for n := range numbers {
		ns = append(ns, n)
	}
	return ns
}
