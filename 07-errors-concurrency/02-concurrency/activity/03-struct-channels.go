package main

import (
	"fmt"
	"sync"
)

type Packet struct {
	id      int
	message string
}

func main() {
	wg.Add(2)
	stream := make(chan Packet)
	go send(stream)
	go receive(stream)
	wg.Wait()
}

var wg sync.WaitGroup

var messages = []string{
	"hello",
	"how you doing",
	"good",
	"ok",
	"goodbye",
}

func send(stream chan<- Packet) {
	for i, m := range messages {
		p := Packet{i, m}
		fmt.Println("SENT", p)
		stream <- p
	}
	close(stream)
	wg.Done()
}

func receive(stream <-chan Packet) {
	for data, ok := <-stream; ok; data, ok = <-stream {
		fmt.Println("RECEIVED", data)
	}
	wg.Done()
}
