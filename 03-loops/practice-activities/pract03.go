package main

import "fmt"

func main() {
	// primes from 0 to 100
	for i := 0; i < 100; i++ {
		if isPrime(i) {
			fmt.Printf("%d\n", i)
		}
	}
}

func isPrime(n int) bool {
	// Yi-Hsuan Hsu
	if n < 2 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
