package main

import "fmt"

func main() {
	// print twin prime numbers, prime numbers that are separated by 2
	// Yi-Hsuan Hsu
	for num := 1; num <= 2000; num++ {
		if isPrime(num) && isPrime(num+2) {
			fmt.Printf("%3v and %3v are twin primes\n", num, num+2)
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
