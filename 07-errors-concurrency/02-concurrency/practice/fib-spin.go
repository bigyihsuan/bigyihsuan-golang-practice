package main

import (
	"fmt"
	"time"
)

func main() {
	const n = 45
	go spinner(time.Second / 10)
	fibN := fib(n)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	spin := `\|/-`
	for i := 0; ; i = (i + 1) % len(spin) {
		fmt.Printf("\r%s", string(spin[i]))
		time.Sleep(delay)
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
