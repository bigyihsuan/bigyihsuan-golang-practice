package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var mutex sync.Mutex

func main() {
	var nums []int
	for i := 1; i <= 20; i++ {
		nums = append(nums, i)
	}
	averages := make(chan float64)
	// 0 -> [0,5)
	// 1 -> [5,10)
	// etc
	k := 4 // number of goroutines
	n := len(nums) / k
	for i := 0; i < k; i++ {
		low := n * i
		high := low + n
		// fmt.Println(low, high)
		go avgN(nums[low:high], averages, n, i)
		x := <-averages
		fmt.Printf("[%2v:%2v) = %v\n", low, high, x)
	}
}

func avgN(numbers []int, out chan float64, n int, id int) {
	go func() {
		sum := 0
		for _, e := range numbers {
			sum += e
		}
		avg := float64(sum) / float64(n)
		// fmt.Println(id, avg)
		out <- avg
	}()
}
