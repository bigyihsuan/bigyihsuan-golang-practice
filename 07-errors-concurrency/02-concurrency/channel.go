package main

import (
	"fmt"
	"sync"
)

func message(ch chan<- string) {
	fmt.Println("sending message")
	ch <- "hello world"
}

func printMessage(ch <-chan string) {
	fmt.Println(<-ch)
	wg.Done()
}

var wg sync.WaitGroup
func main() {
	wg.Add(1)
	channel := make(chan string)
	go message(channel)
	go printMessage(channel)
	wg.Wait()
	fmt.Println("last")
}
