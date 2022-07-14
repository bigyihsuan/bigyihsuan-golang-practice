package main

import (
	"fmt"
	"time"
)

func main() {
	sendChannel := make(chan int, 1)
	receiveChannel := make(chan int, 1)

	go send(sendChannel, receiveChannel)
	go get(receiveChannel, sendChannel)
	sendChannel <- 1
	fmt.Printf("sender: %v\n", sendChannel)
	fmt.Printf("receiver: %v\n", receiveChannel)

	// for {
	// 	select {
	// 	case data := <-sendChannel:
	// 		fmt.Printf("main sendChannel data: %v\n", data)
	// 	case data := <-receiveChannel:
	// 		fmt.Printf("main receiveChannel data: %v\n", data)
	// 	}
	// }
	select {}
}

// send() reads from input, and writes to output
func send(input <-chan int, output chan<- int) {
	for {
		// read from the send channel
		data := <-input
		fmt.Printf("data from send channel: %v\n", data)
		// manipulate data
		// push the data out
		time.Sleep(time.Second / 4)
		fmt.Println("    send sending")
		output <- data + 1
	}
}

// reads from input, resends it to output
func get(input <-chan int, output chan<- int) {
	for {
		// read from the send channel
		data := <-input
		fmt.Printf("data from receive channel: %v\n", data)
		// manipulate data
		// push the data out
		time.Sleep(time.Second / 4)
		fmt.Println("    get sending")
		output <- data + 1
	}
}
