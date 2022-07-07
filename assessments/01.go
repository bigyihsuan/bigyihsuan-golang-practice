package main

import "fmt"

func main() {
	numbers := []int{5, 6, 7, 8, 9}
	fmt.Printf("numbers: %v\n", numbers)
	fmt.Printf("sum of numbers: %v\n", sum(numbers))

	h2, b2 := half(2)
	fmt.Printf("half %v => (%v, %v)\n", 2, h2, b2)
	h5, b5 := half(5)
	fmt.Printf("half %v => (%v, %v)\n", 5, h5, b5)

	fmt.Printf("max: %v\n", max(numbers...))

	odds := makeOddGenerator()
	for _i := 0; _i < 10; _i++ {
		fmt.Printf("%v ", odds())
	}
	fmt.Println()

	f := fib()
	for i := 0; i < 10; i++ {
		fmt.Printf("fibonacci %v => %v\n", i, f(i))
	}
	for i := 10; i >= 0; i-- {
		fmt.Printf("fibonacci %v => %v\n", i, f(i))
	}
	fmt.Printf("fibonacci %v => %v\n", 11, f(11))

}

// 1) sum is a function which takes a slice of numbers and adds them together. What would its function signature look like in Go?
// 2) Write a function which takes an integer and halves it and returns true if it was even or false if it was odd. For example half(1) should return (0, false) and half(2) should return (1, true).
// 3) Write a function with one variadic parameter that finds the greatest number in a list of numbers.
// 4) Using makeEvenGenerator as an example, write a makeOddGenerator function that generates odd numbers.
// 5) The Fibonacci sequence is defined as: 0, fib(0) = fib(1) = 1, fib(n) = fib(n-1) + fib(n-2). Write a recursive function which can find fib(n).
func sum(numbers []int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}

func half(n int) (int, bool) {
	return n / 2, n%2 == 0
}

func max(numbers ...int) int {
	var max int = -1
	for _, e := range numbers {
		if e >= max {
			max = e
		}
	}
	return max
}

func makeOddGenerator() func() int {
	var number = 1
	return func() int {
		old := number
		number += 2
		return old
	}
}

func fibonacci(n int) int {
	if n <= 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// optimized
// use memoization to store the generated numbers
func fib() func(int) int {
	generated := make(map[int]int)
	// take from generated if already exists
	return func(n int) int {
		if i, ok := generated[n]; ok {
			// fmt.Printf("cache: %v\n", generated)
			return i
		}
		return _fib(n, generated)
	}
}
func _fib(n int, cache map[int]int) int {
	if n <= 1 {
		cache[n] = 1
		return cache[n]
	}
	cache[n] = _fib(n-1, cache) + _fib(n-2, cache)
	return cache[n]
}
