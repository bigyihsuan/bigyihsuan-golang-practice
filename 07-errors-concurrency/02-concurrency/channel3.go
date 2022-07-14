package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	ch := make(chan int)
	// go func() {
	// 	ch <- 12
	// 	close(ch)
	// }()
	// fmt.Println("hello", <-ch)
	
	go put(ch)
	// waiting for 1 function
	wg.Add(1)
	go get(ch)
	// wait for the function
	wg.Wait()
}

func put(c chan int) {
	// put things into the channel
	for i := 1; i <= 5; i++ {
		c <- i
	}
	// close to prevent race conditions
	close(c)
}

func get(c chan int) {
	// forever...
	for {
		// get an element from the channel
		n, ok := <-c
		fmt.Printf("ok: %T %v\n", ok, ok)
		// check for a good read
		// ok is false when the channel is closed
		if ok {
			fmt.Printf("n: %v\n", n)
			fmt.Printf("c: %v\n", c)
			time.Sleep(time.Millisecond * 500)
		} else {
			// channel is closed, exit the function
			fmt.Println("closed channel")
			wg.Done()
			return
		}
	}
}
