package main

import "fmt"

func main() {
	// get int
	var n uint64
	fmt.Print("enter n: ")
	fmt.Scanf("%d", &n)

	// the number of digits in the value entered
	var numDigits = 1
	for n/uint64(numDigits*10) > 0 {
		numDigits += 1
	}
	fmt.Printf("numDigits: %v\n", numDigits)

	// the first and last digits of the number
	var firstDigit, lastDigit uint64 = n, n
	if numDigits != 1 {
		firstDigit /= uint64(numDigits-1) * 10
		lastDigit %= 10
	}
	fmt.Printf("firstDigit: %v\n", firstDigit)
	fmt.Printf("lastDigit: %v\n", lastDigit)

	// the sum of the digits in the number
	var digitSum, digitProd uint64 = 0, 1
	for i := n; i > 0; i /= 10 {
		digitSum += (i % 10)
		digitProd *= (i % 10)
	}
	fmt.Printf("digitSum: %v\n", digitSum)
	// the product of the digits in the number
	fmt.Printf("digitProd: %v\n", digitProd)

	// whether the number is prime or not
	fmt.Printf("is prime: %v\n", isPrime(n))

	// the factorial of the number
	fmt.Printf("%v! = %v\n", n, factorial(n))

}
func factorial(n uint64) uint64 {
	if n <= 2 {
		return n
	}
	var prod, i uint64 = 1, 1
	for i = 1; i <= n; i++ {
		prod *= i
	}
	return prod
}

func isPrime(n uint64) bool {
	// Yi-Hsuan Hsu
	if n < 2 {
		return false
	}
	var i uint64
	for i = 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
