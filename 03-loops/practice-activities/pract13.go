package main

import "fmt"

func main() {
	// determine if an input number can be expressed as sum of prime numbers
	var a int

	fmt.Print("enter a: ")
	fmt.Scanf("%d", &a)

	for i := 1; i < a; i++ {
		j := a - i
		if isPrime(i) && isPrime(j) {
			fmt.Printf("%d = %d + %d\n", a, i, j)
			return
		}
	}
	fmt.Printf("%d cannot be expressed as the sum of prime numbers\n", a)
}

func isPrime(n int) bool {
	// Yi-Hsuan Hsu
	if n < 2 {
		return false
	}
	var i int
	for i = 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
