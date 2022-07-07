package main

import "fmt"

func main() {
	f := primeGenerator()
	for _i := 0; _i < 100; _i++ {
		fmt.Printf("%v ", f())
	}
	fmt.Println()
}

func primeGenerator() func() int {
	start := 2
	return func() int {
		out := start
		start += 1
		for !isPrime(start) {
			start += 1
		}
		return out
	}
}

func isPrime(n int) bool {
	// Yi-Hsuan Hsu
	if n < 2 || n%2 == 0 {
		return false
	}
	if n == 2 {
		return true
	}
	for i := 3; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// func fib() func() int {
// 	first, second := 0, 1
// 	return func() int {
// 		ret := first
// 		first, second = second, first+second
// 		fmt.Printf("%v %v %v\n", ret, first, second)
// 		return ret
// 	}
// }

// sumx := func(a, b, c int) int {
// 	return a + b + c
// }
// fmt.Printf("sumx: %T\n", sumx)
// fmt.Println(sumx(1, 2, 3))
