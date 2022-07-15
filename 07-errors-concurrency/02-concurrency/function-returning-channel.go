package main

import (
	"fmt"
	"time"
)

func main() {
	generator := increment(10)
	sum := <-extract(generator)
	time.Sleep(time.Second)
	fmt.Printf("sum: %v\n", sum)
}

func increment(n int) chan int {
	out := make(chan int)
	fmt.Println("running increment func")
	go func() {
		fmt.Println("inside increment func")
		for i := 0; i < n; i++ {
			out <- i
			fmt.Printf("increment %v\n", i)
		}
		close(out)
		fmt.Println("finished inside increment func")
	}()
	fmt.Println("finished increment func")
	return out
}

func extract(c chan int) chan int {
	out := make(chan int)
	fmt.Println("running extract func")
	go func() {
		sum := 0
		for i := range c {
			sum += i
		}
		out <- sum
		close(out)
	}()
	fmt.Println("finished extract func")
	return out
}
